package poll

type IPollRepository interface {
	GetPolls() ([]Poll, error)
	GetPollById(id string) (Poll, error)
}
