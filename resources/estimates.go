package resources

import "github.com/roma351/diary_api/objects"

type Estimates struct {
	Period objects.Period     `json:"period"`
	Items  []objects.Estimate `json:"items"`
	*ExtraResponse
}
