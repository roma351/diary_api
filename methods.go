package diary_api

import (
	"github.com/roma351/diary_api/resources"
)

func (api *DiaryAPI) Last(user *User, params map[string]interface{}) (res resources.Last, err error) {
	err = api.UserRequest("/school/last", user, params, &res)
	return
}

func (api *DiaryAPI) Info() (res struct {
	ID    int32   `json:"id"`
	Title string  `json:"title"`
	Image *string `json:"image"`
}, err error) {
	err = api.UserRequest("/service", nil, nil, &res)
	return
}

func (api *DiaryAPI) InfoSecure(user *User) (res struct {
	UserId int64 `json:"user_id"`
}, err error) {
	err = api.UserRequest("/service/secure", user, nil, &res)
	return
}
