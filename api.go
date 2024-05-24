package diary_api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/roma351/diary_api/resources"
	"github.com/roma351/utils"
	"log"
	"net/http"
	"strconv"
	"time"
)

var ApiUrl = "https://api.school-diary.ru/method"

var UserAgent = "DiaryAPI/1.0"

func (api *DiaryAPI) Request(method string, params map[string]interface{}, response interface{}) error {
	headers := http.Header{
		"X-Token-App": {fmt.Sprintf("%s.%s", api.AppId, api.AppPrivate)},
	}
	return api.request(ApiUrl+method, params, headers, response)
}

func (api *DiaryAPI) UserRequest(method string, user *User, params map[string]interface{}, response interface{}) error {
	t := time.Now().Unix()
	mySecureRaw := fmt.Sprintf("secure-%s-%d-user-uid-%d-%d", api.AppId, t, user.AuthId, user.UserUid)

	mySecure := utils.SHA1(fmt.Sprintf("%s-%s", mySecureRaw, api.AppPrivate2))

	headers := http.Header{
		"X-Token-App": {fmt.Sprintf("%s.%s", api.AppId, api.AppPrivate)},
		"X-Secure":    {fmt.Sprintf("%d.%d.%d.%s", user.AuthId, user.UserUid, t, mySecure)},
	}

	if params != nil {
		if p, ok := params["account_id"]; ok {
			switch v := p.(type) {
			case string:
				headers["X-Account-Id"] = []string{v}
			case int:
				headers["X-Account-Id"] = []string{strconv.Itoa(v)}
			case int64:
				headers["X-Account-Id"] = []string{strconv.FormatInt(v, 10)}
			}
			delete(params, "account_id")
		}
	}

	return api.request(ApiUrl+method, params, headers, response)
}

func (api *DiaryAPI) request(url string, payload map[string]interface{}, headers http.Header, response interface{}) error {
	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		return err
	}

	req.Header = headers
	//req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Connection", "Keep-Alive")
	req.Header.Set("User-Agent", UserAgent)
	if !api.CallUser {
		req.Header.Set("X-Call-User", "0")
	}

	if api.Log {
		log.Printf("DIARY API: URL: %s, body: %s\n", url, body)
	}

	if api.client == nil {
		api.client = &http.Client{}
	}

	res, err := api.client.Do(req)
	if err != nil {
		return newError(994, err, "")
	}
	defer res.Body.Close()

	requestId := res.Header.Get("X-Request-Id")

	if api.Log {
		log.Printf("DIARY API: [%d] %s\n", res.StatusCode, requestId)
	}

	if res.StatusCode != http.StatusOK {
		return newError(996, nil, requestId)
	}

	var coreResponse resources.Core
	if err := json.NewDecoder(res.Body).Decode(&coreResponse); err != nil {
		return newError(997, err, requestId)
	}

	if coreResponse.Success {
		if coreResponse.Data == nil {
			return newError(998, nil, requestId)
		}
		if err := json.Unmarshal(coreResponse.Data, &response); err != nil {
			return newError(999, err, requestId)
		}
		return nil
	}

	if coreResponse.Error == nil {
		return newError(995, nil, requestId)
	}
	return newErrorApi(coreResponse.Error, requestId)
}
