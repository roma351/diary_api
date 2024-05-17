package objects

type Class struct {
	ID       int64   `json:"id"`
	Title    string  `json:"title"`
	NickName *string `json:"nickname"`
	Parallel int32   `json:"parallel"`
	Emoji    *string `json:"emoji"`
}
