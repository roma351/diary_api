package objects

type School struct {
	ID         int64    `json:"id"`
	TypeId     int32    `json:"type_id"`
	SourceId   int32    `json:"source_id"`
	SourcePath string   `json:"source_code"`
	VendorId   *int32   `json:"vendor_id"`
	Title      string   `json:"title"`
	NickName   *string  `json:"nickname"`
	Emoji      *string  `json:"emoji"`
	Country    *Country `json:"country"`
	Region     *Region  `json:"region"`
	City       *City    `json:"city"`
}
