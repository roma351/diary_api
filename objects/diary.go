package objects

type DiaryDay struct {
	Day     DiaryDayInfo  `json:"day_info"`
	Date    string        `json:"date"`
	Lessons []DiaryLesson `json:"lessons"`
}
