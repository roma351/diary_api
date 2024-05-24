package resources

import "github.com/roma351/diary_api/objects"

type Login struct {
	User     objects.UserContext     `json:"user"`
	Token    objects.Token           `json:"token"`
	Sign     *objects.Sign           `json:"sign"`
	Accounts objects.AccountsContext `json:"accounts"`
	Settings objects.SettingsContext `json:"settings"`
	New      bool                    `json:"new"`
	*ExtraResponse
}

type BySign Login
