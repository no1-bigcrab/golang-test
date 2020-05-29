package models

// Product abc
type Product struct {
	Title             string      `json:"title"`
	BodyHTML          string      `json:"body_html"`
	Vendor            string      `json:"vendor"`
	ProductType       string      `json:"product_type"`
	Handle            string      `json:"handle"`
	TemplateSuffix    interface{} `json:"template_suffix"`
	PublishedScope    string      `json:"published_scope"`
	Tags              string      `json:"tags"`
	AdminGraphqlAPIID string      `json:"admin_graphql_api_id"`
	Variants          []struct {
		ID                   int         `json:"id"`
		ProductID            int         `json:"product_id"`
		Title                string      `json:"title"`
		Price                string      `json:"price"`
		Sku                  string      `json:"sku"`
		Position             int         `json:"position"`
		InventoryPolicy      string      `json:"inventory_policy"`
		CompareAtPrice       interface{} `json:"compare_at_price"`
		FulfillmentService   string      `json:"fulfillment_service"`
		InventoryManagement  interface{} `json:"inventory_management"`
		Option1              string      `json:"option1"`
		Option2              interface{} `json:"option2"`
		Option3              interface{} `json:"option3"`
		Taxable              bool        `json:"taxable"`
		Barcode              interface{} `json:"barcode"`
		Grams                int         `json:"grams"`
		ImageID              interface{} `json:"image_id"`
		Weight               float64     `json:"weight"`
		WeightUnit           string      `json:"weight_unit"`
		InventoryItemID      int         `json:"inventory_item_id"`
		InventoryQuantity    int         `json:"inventory_quantity"`
		OldInventoryQuantity int         `json:"old_inventory_quantity"`
		RequiresShipping     bool        `json:"requires_shipping"`
		AdminGraphqlAPIID    string      `json:"admin_graphql_api_id"`
		PresentmentPrices    []struct {
			Price struct {
				CurrencyCode string `json:"currency_code"`
				Amount       string `json:"amount"`
			} `json:"price"`
			CompareAtPrice interface{} `json:"compare_at_price"`
		} `json:"presentment_prices"`
	} `json:"variants"`
	Options []struct {
		ID        int      `json:"id"`
		ProductID int      `json:"product_id"`
		Name      string   `json:"name"`
		Position  int      `json:"position"`
		Values    []string `json:"values"`
	} `json:"options"`
	Images []struct {
		ID                int           `json:"id"`
		ProductID         int           `json:"product_id"`
		Position          int           `json:"position"`
		Alt               interface{}   `json:"alt"`
		Width             int           `json:"width"`
		Height            int           `json:"height"`
		Src               string        `json:"src"`
		VariantIds        []interface{} `json:"variant_ids"`
		AdminGraphqlAPIID string        `json:"admin_graphql_api_id"`
	} `json:"images"`
	Image struct {
		ID                int           `json:"id"`
		ProductID         int           `json:"product_id"`
		Position          int           `json:"position"`
		Alt               interface{}   `json:"alt"`
		Width             int           `json:"width"`
		Height            int           `json:"height"`
		Src               string        `json:"src"`
		VariantIds        []interface{} `json:"variant_ids"`
		AdminGraphqlAPIID string        `json:"admin_graphql_api_id"`
	} `json:"image"`
}
