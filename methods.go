package diary_api

import (
	"github.com/roma351/diary_api/resources"
)

func (api *DiaryAPI) Context(user *User, params map[string]interface{}) (res resources.Context, err error) {
	err = api.UserRequest("/user/context", user, params, &res)
	return
}

func (api *DiaryAPI) Account(user *User, params map[string]interface{}) (res interface{}, err error) {
	err = api.UserRequest("/account/context", user, params, &res)
	return
}

func (api *DiaryAPI) Last(user *User, params map[string]interface{}) (res resources.Last, err error) {
	err = api.UserRequest("/school/last", user, params, &res)
	return
}

func (api *DiaryAPI) Logout(user *User, params map[string]interface{}) (res resources.Last, err error) {
	err = api.UserRequest("/account/logout", user, params, &res)
	return
}

func (api *DiaryAPI) App() (res struct {
	ID    int32   `json:"id"`
	Title string  `json:"title"`
	Image *string `json:"image"`
}, err error) {
	err = api.UserRequest("/service", nil, nil, &res)
	return
}

func (api *DiaryAPI) AppSecure(user *User) (res struct {
	UserId int64 `json:"user_id"`
}, err error) {
	err = api.UserRequest("/service/secure", user, nil, &res)
	return
}
