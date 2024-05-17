package diary_api

import (
	"github.com/roma351/diary_api/resources"
)

func (api *DiaryAPI) Last(user *User, params map[string]interface{}) (res resources.Last, err error) {
	err = api.UserRequest("/school/last", user, params, &res)
	return
}
