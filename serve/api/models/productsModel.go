package models

// Products abc
type Products struct {
	Links struct {
		Collection []struct {
			Href string `json:"href"`
		} `json:"collection"`
		Self []struct {
			Href string `json:"href"`
		} `json:"self"`
		Up []struct {
			Href string `json:"href"`
		} `json:"up"`
	} `json:"_links"`
	Attributes []struct {
		ID     int    `json:"id"`
		Name   string `json:"name"`
		Option string `json:"option"`
	} `json:"attributes"`
	Backordered       bool        `json:"backordered"`
	Backorders        string      `json:"backorders"`
	BackordersAllowed bool        `json:"backorders_allowed"`
	DateCreated       string      `json:"date_created"`
	DateCreatedGmt    string      `json:"date_created_gmt"`
	DateModified      string      `json:"date_modified"`
	DateModifiedGmt   string      `json:"date_modified_gmt"`
	DateOnSaleFrom    interface{} `json:"date_on_sale_from"`
	DateOnSaleFromGmt interface{} `json:"date_on_sale_from_gmt"`
	DateOnSaleTo      interface{} `json:"date_on_sale_to"`
	DateOnSaleToGmt   interface{} `json:"date_on_sale_to_gmt"`
	Description       string      `json:"description"`
	Dimensions        struct {
		Height string `json:"height"`
		Length string `json:"length"`
		Width  string `json:"width"`
	} `json:"dimensions"`
	DownloadExpiry int           `json:"download_expiry"`
	DownloadLimit  int           `json:"download_limit"`
	Downloadable   bool          `json:"downloadable"`
	Downloads      []interface{} `json:"downloads"`
	ID             int           `json:"id"`
	Image          struct {
		Alt             string `json:"alt"`
		DateCreated     string `json:"date_created"`
		DateCreatedGmt  string `json:"date_created_gmt"`
		DateModified    string `json:"date_modified"`
		DateModifiedGmt string `json:"date_modified_gmt"`
		ID              int    `json:"id"`
		Name            string `json:"name"`
		Src             string `json:"src"`
	} `json:"image"`
	ManageStock bool `json:"manage_stock"`
	MenuOrder   int  `json:"menu_order"`
	MetaData    []struct {
		ID    int    `json:"id"`
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"meta_data"`
	OnSale          bool        `json:"on_sale"`
	Permalink       string      `json:"permalink"`
	Price           string      `json:"price"`
	Purchasable     bool        `json:"purchasable"`
	RegularPrice    string      `json:"regular_price"`
	SalePrice       string      `json:"sale_price"`
	ShippingClass   string      `json:"shipping_class"`
	ShippingClassID int         `json:"shipping_class_id"`
	Sku             string      `json:"sku"`
	Status          string      `json:"status"`
	StockQuantity   interface{} `json:"stock_quantity"`
	StockStatus     string      `json:"stock_status"`
	TaxClass        string      `json:"tax_class"`
	TaxStatus       string      `json:"tax_status"`
	Virtual         bool        `json:"virtual"`
	Weight          string      `json:"weight"`
}
