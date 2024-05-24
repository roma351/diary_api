package resources

import "github.com/roma351/diary_api/objects"

type Sources struct {
	Items           []*objects.Source `json:"items"`
	DefaultSourceId int32             `json:"default_source_id"`
	*ExtraResponse
}

type Regions struct {
	Items []*objects.Region `json:"items"`
	*ExtraResponse
}
