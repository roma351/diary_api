package objects

type Account struct {
	ID            int64    `json:"id"`
	RoleId        int32    `json:"role_id"`
	SourceId      int32    `json:"source_id"`
	DataAccountId *int64   `json:"data_account_id"`
	Name          string   `json:"name"`
	NickName      *string  `json:"nickname"`
	Formatting    *string  `json:"formatting,omitempty"`
	Emoji         *string  `json:"emoji"`
	Sex           *bool    `json:"sex"`
	Private       bool     `json:"private"`
	Class         Class    `json:"class"`
	School        School   `json:"school"`
	Permissions   []string `json:"permissions"`
	RatingTitle   string   `json:"rating_title"`
}
