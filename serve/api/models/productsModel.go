package models

// Products abc
type Products []struct {
	Name             string `json:"name"`
	Type             string `json:"type"`
	Description      string `json:"description"`
	ShortDescription string `json:"short_description"`
	Categories       []struct {
		ID int `json:"id"`
	} `json:"categories"`
	Images []struct {
		Src string `json:"src"`
	} `json:"images"`
	Attributes []struct {
		ID        int      `json:"id,omitempty"`
		Position  int      `json:"position"`
		Visible   bool     `json:"visible"`
		Variation bool     `json:"variation"`
		Options   []string `json:"options"`
		Name      string   `json:"name,omitempty"`
	} `json:"attributes"`
	DefaultAttributes []struct {
		ID     int    `json:"id,omitempty"`
		Option string `json:"option"`
		Name   string `json:"name,omitempty"`
	} `json:"default_attributes"`
}
