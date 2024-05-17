package objects

type LessonGeneric struct {
	ID       string       `json:"id"`
	SID      string       `json:"sid"`
	Date     DiaryDayInfo `json:"day_info"`
	Number   int32        `json:"number"`
	Subject  Subject      `json:"subject"`
	Homeworks []Homework     `json:"homeworks"`
	Teachers []Teacher    `json:"teachers"`
}

type DiaryLesson struct {
	ID          string          `json:"id"` // unix + | + id
	SID         string          `json:"sid"`
	Number      int32           `json:"number"`
	Place       string          `json:"place"`
	Type        string          `json:"type"` // самостоятельная работа, практика
	Theme       string          `json:"theme"`
	IsCanceled  bool            `json:"is_canceled"`
	Comment     string          `json:"comment"` // если пару отменили, причина
	Time        DiaryLessonTime `json:"time"`
	Subject     Subject         `json:"subject"`
	Homeworks    []Homework        `json:"homeworks"`
	Marks       []Mark          `json:"marks"`
	Attachments []Attachment    `json:"attachments"`
	Teachers    []Teacher       `json:"teachers"`
	Mood        *string         `json:"mood"`
}

type DiaryLessonTime struct {
	Start     string `json:"start"`
	StartUnix int64  `json:"start_unix"`
	End       string `json:"end"`
	EndUnix   int64  `json:"end_unix"`
	Duration  int64  `json:"duration"`
	Current   bool   `json:"current"`
}
