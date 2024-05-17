package resources

type ExtraResponse struct {
	Cache *bool `json:"cache,omitempty"`
	HasTG *bool `json:"has_tg"`
}
