package resources

import "github.com/roma351/diary_api/objects"

type Last struct {
	Items []objects.LastMark `json:"items"`
	*ExtraResponse
}
