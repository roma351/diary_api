package objects

type Account struct {
	Id    int64   `json:"id"`
	Image *string `json:"image"`

	Average string  `json:"average"`
	Emoji   *string `json:"emoji"`
	Current bool    `json:"current,omitempty"`
	Private bool    `json:"private"`
	Mood    *string `json:"mood"`
}
