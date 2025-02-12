package objects

type Estimate struct {
	Subject Subject          `json:"subject"`
	Rating  Rating           `json:"rating"`
	Groups  []EstimatesGroup `json:"groups"`
}

type EstimatesGroup struct {
	Title  string `json:"title"`
	Code   string `json:"code"`
	Rating Rating `json:"rating"`
	Marks  []Mark `json:"marks"`
}

type EstimatePeriod struct {
	Class   Class      `json:"class"`
	Period  Period     `json:"period"`
	Current bool       `json:"current"`
	Items   []Estimate `json:"items"`
}
