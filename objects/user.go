package objects

type User struct {
	ID       int64   `json:"id"`
	AuthId   int32   `json:"auth_id"`
	UID      int64   `json:"user_uid"`
	Name     string  `json:"name"`
	UserName *string `json:"username"`
	Photo    *string `json:"photo"`
	Private  bool    `json:"private"`
	Lang     string  `json:"lang"`
}
