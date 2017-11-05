package survey

import (
	"database/sql"
	"log"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/satori/go.uuid"
)

const (
	query_select_survey = "SELECT id, title, creator FROM survey"
	query_select_option = "SELECT `value` FROM `option`"
	query_select_poll   = "SELECT `key`, `value` FROM poll"
	query_count_poll    = "SELECT count(*) FROM poll"
	insert_survey       = "INSERT INTO survey (id, title, creator, options) VALUES (?,?,?,?)"
	insert_option       = "INSERT INTO `option` (survey_id, `value`) VALUES (?,?)"
	insert_poll         = "INSERT INTO poll (survey_id, `key`, `value`) VALUES (?,?,?)"
	update_poll         = "UPDATE poll SET `value` = ?"
)

type MySQLRepository struct {
	SurveyRepository
	db *sqlx.DB
}

func NewMySQLRepository(database *sqlx.DB) *MySQLRepository {
	return &MySQLRepository{db: database}
}

func (r *MySQLRepository) txRollback(tx *sqlx.Tx, in Survey) {
	if err := tx.Rollback(); err != nil {
		log.Panicf("%v : %#v", err, in)
	}
}

func (r *MySQLRepository) GetSurveys() (models []Survey, err error) {
	if err = r.db.Select(&models, query_select_survey); err != nil {
		if err != sql.ErrNoRows {
			return
		}
	}

	//TODO: BUG, this models will only live in this function
	for _, model := range models {
		if err = r.db.Select(&model.Options, query_select_option+" WHERE survey_id = ?", model.ID.String()); err != nil {
			if err != sql.ErrNoRows {
				return
			}
		}

		polls := []Poll{}
		if err = r.db.Select(&polls, query_select_poll+" WHERE survey_id = ?", model.ID.String()); err != nil {
			if err != sql.ErrNoRows {
				return
			}
		}

		for _, poll := range polls {
			model.Polls[poll.Key] = poll.Value
		}
	}

	return
}

func (r *MySQLRepository) GetSurveyById(id uuid.UUID) (model Survey, err error) {
	if err = r.db.Get(&model, query_select_survey+" WHERE id = ?", id.String()); err != nil {
		if err != sql.ErrNoRows {
			return
		}
	}

	if err = r.db.Select(&model.Options, query_select_option+" WHERE survey_id = ?", model.ID.String()); err != nil {
		if err != sql.ErrNoRows {
			return
		}
	}

	polls := []Poll{}
	if err = r.db.Select(&polls, query_select_poll+" WHERE survey_id = ?", model.ID.String()); err != nil {
		if err != sql.ErrNoRows {
			return
		}
	}

	if len(polls) == 0 {
		return
	}

	model.Polls = make(map[string]string)
	for _, poll := range polls {
		model.Polls[poll.Key] = poll.Value
	}

	return
}

func (r *MySQLRepository) StoreSurvey(in Survey) (err error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return
	}
	if _, err = tx.Exec(insert_survey, in.ID, in.Title, in.Creator, strings.Join(in.Options, ":")); err != nil {
		r.txRollback(tx, in)
		return
	}

	for _, value := range in.Options {
		if _, err = tx.Exec(insert_option, in.ID, value); err != nil {
			r.txRollback(tx, in)
			return
		}
	}

	tx.Commit()
	return
}

func (r *MySQLRepository) StorePoll(id uuid.UUID, in map[string]string) (err error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return
	}
	for k, v := range in {
		var poll Poll

		if err = r.db.Get(&poll, query_select_poll+" WHERE survey_id = ? AND `key` = ?", id.String(), k); err != nil {
			if err != sql.ErrNoRows {
				return
			}
		}

		if poll.Key == "" {
			if _, err = tx.Exec(insert_poll, id.String(), k, v); err != nil {
				r.txRollback(tx, Survey{Polls: in})
				return
			}
			continue
		}

		if _, err = tx.Exec(update_poll+" WHERE survey_id = ? AND `key` = ?", v, id.String(), k); err != nil {
			r.txRollback(tx, Survey{Polls: in})
			return
		}
	}

	tx.Commit()
	return
}

/**
CREATE TABLE `survey` (
  `id` char(36) NOT NULL DEFAULT '',
  `title` varchar(128) DEFAULT NULL,
  `creator` varchar(128) DEFAULT NULL,
  `options` varchar(256) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `poll` (
  `key` varchar(256) NOT NULL DEFAULT '',
  `value` varchar(256) NOT NULL DEFAULT '',
  `survey_id` char(36) NOT NULL DEFAULT '',
  PRIMARY KEY (`key`,`survey_id`),
  KEY `fk_survey_poll` (`survey_id`),
  CONSTRAINT `fk_survey_poll` FOREIGN KEY (`survey_id`) REFERENCES `survey` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `option` (
  `survey_id` char(36) NOT NULL DEFAULT '',
  `value` varchar(64) NOT NULL DEFAULT '',
  PRIMARY KEY (`survey_id`,`value`),
  CONSTRAINT `fk_option_survey` FOREIGN KEY (`survey_id`) REFERENCES `survey` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
**/
