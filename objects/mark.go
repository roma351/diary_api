package objects

type Mark struct {
	SID    string  `json:"sid"`
	Type   int32   `json:"type"`
	Date   *string `json:"date"`
	Value  string  `json:"value"`
	Value2 int32   `json:"value2"`
}

type LastMark struct {
	ID        string       `json:"id"`
	SID       string       `json:"sid"`
	Date      DiaryDayInfo `json:"date"`
	Subject   Subject      `json:"subject"`
	Type      string       `json:"type"`
	TypeShort string       `json:"type_short"`
	Marks     []Mark       `json:"marks"`
	Mood      *string      `json:"mood"`
}
