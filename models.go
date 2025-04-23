package diary_api

import (
	"net/http"
	"time"
)

type DiaryAPI struct {
	client *http.Client

	AppId       string
	AppSecret   string
	AppPrivate  string
	AppPrivate2 string

	Log bool

	CallUser bool

	Timeout time.Duration
}

type User struct {
	AuthId  int32
	UserUid int64
}
