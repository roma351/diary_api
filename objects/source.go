package objects

type Source struct {
	ID        int32   `json:"id"`
	Code      string  `json:"code"`
	Name      string  `json:"name"`
	ShortName string  `json:"short_name"`
	Image     string  `json:"image"`
	Types     []int32 `json:"types"`

	NeedVendor bool `json:"need_vendor"`

	Verified *bool `json:"verified"`
}
