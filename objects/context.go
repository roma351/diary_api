package objects

type UserContext struct {
	ID       int64   `json:"id"`
	AuthId   int32   `json:"auth_id"`
	UID      int64   `json:"user_uid"`
	Name     string  `json:"name"`
	UserName *string `json:"username"`
	Photo    *string `json:"photo"`
	Private  bool    `json:"private"`
	Lang     string  `json:"lang"`
}

type SettingsContext struct {
	Ad          bool `json:"ad"`
	CanLoginTwo bool `json:"can_login_two"`
}

type AccountsContext struct {
	Default  int64            `json:"default"`
	Accounts []AccountContext `json:"items"`
}

type AccountContext struct {
	ID          int64   `json:"id"`
	RoleId      int32   `json:"role_id"`
	SourceId    int32   `json:"source_id"`
	Name        string  `json:"name"`
	NickName    *string `json:"nickname"`
	Emoji       *string `json:"emoji"`
	Sex         *bool   `json:"sex"`
	Private     bool    `json:"private"`
	Class       Class   `json:"class"`
	School      School  `json:"school"`
	RatingTitle string  `json:"rating_title"`
}

type Token struct {
	Token     string `json:"token"`
	ExpiresAt int64  `json:"expires_at"`
}

type Sign struct {
	Sign string `json:"sign"`
}
