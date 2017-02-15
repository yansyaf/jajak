package poll

type Poll struct {
	Title   string   `db:"title" json:"title"`
	Creator string   `db:"creator" json:"creator"`
	Items   []string `db:"items" json:"items"`
}
