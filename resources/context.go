package resources

import "github.com/roma351/diary_api/objects"

type Context struct {
	User     objects.UserContext     `json:"user"`
	Accounts objects.AccountsContext `json:"accounts"`
	Settings objects.SettingsContext `json:"settings"`
	*ExtraResponse
}

type Context2 struct {
	Account objects.AccountContext `json:"account"`
}
