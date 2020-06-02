package models

// Product abc
type Product struct {
	Title       string `json:"title"`
	BodyHTML    string `json:"body_html"`
	Vendor      string `json:"vendor"`
	ProductType string `json:"product_type"`
	Variants    []struct {
		Option1 string `json:"option1"`
		Option2 string `json:"option2"`
	} `json:"variants"`
	Options []struct {
		Name   string   `json:"name"`
		Values []string `json:"values"`
	} `json:"options"`
}
