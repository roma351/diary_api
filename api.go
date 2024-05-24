package diary_api

import (
	"encoding/json"
	"fmt"
	"github.com/roma351/diary_api/resources"
	"github.com/roma351/utils"
	"log"
	"net/http"
	"strings"
	"time"
)

var ApiUrl = "https://api.school-diary.ru/method"

var UserAgent = "DiaryAPI/1.0"

func (api *DiaryAPI) Request(method string, params map[string]interface{}, response interface{}) error {
	headers := map[string][]string{
		"X-Token-App": {fmt.Sprintf("%s.%s", api.AppId, api.AppPrivate)},
	}
	return api.request(ApiUrl+method, params, headers, response)
}

func (api *DiaryAPI) UserRequest(method string, user *User, params map[string]interface{}, response interface{}) error {
	dd := time.Now().Unix()
	mySecureRaw := fmt.Sprintf("secure-%s-%d-user-uid-%d-%d", api.AppId, dd, user.AuthId, user.UserUid)

	mySecure := utils.SHA1(fmt.Sprintf("%s-%s", mySecureRaw, api.AppPrivate2))

	headers := map[string][]string{
		"X-Token-App": {fmt.Sprintf("%s.%s", api.AppId, api.AppPrivate)},
		"X-Secure":    {fmt.Sprintf("%d.%d.%d.%s", user.AuthId, user.UserUid, dd, mySecure)},
	}

	if params != nil {
		if p, ok := params["account_id"]; ok {
			if z, ok := p.(string); ok {
				headers["X-Account-Id"] = []string{z}
			} else if z, ok := p.(int); ok {
				headers["X-Account-Id"] = []string{utils.IntToString(z)}
			} else if z, ok := p.(int64); ok {
				headers["X-Account-Id"] = []string{utils.Int64ToString(z)}
			}
			delete(params, "account_id")
		}
	}

	return api.request(ApiUrl+method, params, headers, response)
}

func (api *DiaryAPI) request(url string, json2 map[string]interface{}, headers http.Header, response interface{}) error {
	body, err := json.Marshal(json2)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", url, strings.NewReader(string(body)))
	if err != nil {
		return err
	}
	req.Header = headers
	//req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Connection", "Keep-Alive")
	req.Header.Set("User-Agent", UserAgent)
	if api.CallUser == false {
		req.Header.Set("X-Call-User", "0")
	}
	if api.Log {
		log.Println("DIARY API: request", url, string(body))
	}
	if api.client == nil {
		api.client = &http.Client{}
	}
	res, err := api.client.Do(req)

	if err != nil {
		return newError(994, err.Error(), "", "")
	}
	defer res.Body.Close()

	requestId := res.Header.Get("X-Request-Id")

	bodyString, err := utils.IOToString(res.Body)
	if err != nil {
		return newError(995, err.Error(), "", requestId)
	}
	if api.Log {
		log.Println("DIARY API: response", bodyString)
	}
	if res.StatusCode == 200 {
		var response2 resources.APIResponse
		err := json.Unmarshal([]byte(bodyString), &response2)
		if err != nil {
			return newError(997, bodyString+" | "+err.Error(), "", requestId)
		}
		if response2.Success {
			if response2.Data == nil {
				return newError(998, bodyString, "", requestId)
			}
			if err := json.Unmarshal(response2.Data, &response); err != nil {
				return newError(999, bodyString+" | "+err.Error(), "", requestId)
			}
			return nil
		} else {
			if response2.Error2 != nil {
				response2.Error2.RequestId = requestId
			}
			return response2.Error2
		}
	}
	return newError(996, fmt.Sprintf("%d, %s", res.StatusCode, bodyString), fmt.Sprintf("server error %d", res.StatusCode), requestId)
}

func newError(code int, devMessage string, message string, requestId string) *resources.APIResponseError {
	if message == "" {
		message = "Школьный Дневник не отвечает. Повторите позже"
	}
	var err2 resources.APIResponseError
	err2.Code = code
	err2.Message = message
	err2.DevMessage = devMessage
	err2.RequestId = requestId
	return &err2
}
