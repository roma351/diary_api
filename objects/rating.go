package objects

type Rating struct {
	Average string  `json:"average"`
	Place   int     `json:"place"`
	Trend   *bool   `json:"trend"`
	Mood    *string `json:"mood"`
}

type RatingCard struct {
	ItemId  int64   `json:"item_id"`
	Place   int     `json:"place"`
	Image   *string `json:"image"`
	Title   string  `json:"title"`
	Average string  `json:"average"`
	Emoji   *string `json:"emoji"`
	Current bool    `json:"current"`
	Private bool    `json:"private"`
	Mood    *string `json:"mood"`

	Formatting *string `json:"formatting,omitempty"`

	Class    *Class `json:"class,omitempty"`
	Parallel *int32 `json:"parallel,omitempty"`
}

type RatingSubject struct {
	Subject Subject      `json:"subject"`
	Items   []RatingCard `json:"items"`
}
