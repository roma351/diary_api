package objects

type Period struct {
	Type     int32  `json:"type"`
	TypeText string `json:"type_text"`
	Number   int32  `json:"number"`
	Start    string `json:"start"`
	End      string `json:"end"`
}
