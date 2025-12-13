package objects

type Friend struct {
	UserId   int64  `json:"user_id"`
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Image    string `json:"image"`
}
