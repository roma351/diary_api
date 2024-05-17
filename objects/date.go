package objects

type DiaryDayInfo struct {
	Day   string `json:"day"`
	Date  string `json:"date"`
	Month string `json:"month"`
	Year  string `json:"year"`
	Text  string `json:"text"`
	Week  struct {
		ID    int    `json:"id"`
		Title string `json:"title"`
		Text  string `json:"text"`
	} `json:"week"`
}
