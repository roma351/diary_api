package resources

import "github.com/roma351/diary_api/objects"

type Diary struct {
	Days    []objects.DiaryDay    `json:"days"`
	LastDay *objects.DiaryDayInfo `json:"last_day"`
	NextDay *objects.DiaryDayInfo `json:"next_day"`
	*ExtraResponse
}
