package objects

type Estimate struct {
	Subject Subject          `json:"subject"`
	Groups  []EstimatesGroup `json:"groups"`
}

type EstimatesGroup struct {
	Title  string `json:"title"`
	Code   string `json:"code"`
	Rating Rating `json:"rating"`
	Marks  []Mark `json:"marks"`
}
