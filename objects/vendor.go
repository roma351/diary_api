package objects

type Vendor struct {
	ID       int32   `json:"id"`
	Title    *string `json:"school"`
	Url      *string `json:"url"`
	Country  *string `json:"country"`
	Region   *string `json:"region"`
	City     *string `json:"city"`
	Image    *string `json:"image"`
	Verified *bool   `json:"verified"`
}
