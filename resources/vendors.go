package resources

import "github.com/roma351/diary_api/objects"

type Vendors struct {
	Items []*objects.Vendor `json:"items"`
	*ExtraResponse
}
