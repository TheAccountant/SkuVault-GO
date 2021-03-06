package inventory

// RemoveItem is a automatically generated struct from json provided by sku vault's api docs.
type RemoveItem struct {
	Code         string `json:"Code"`
	LocationCode string `json:"LocationCode"`
	Quantity     int  `json:"Quantity"`
	Reason       string `json:"Reason"`
	Sku          string `json:"Sku"`
	WarehouseID  int  `json:"WarehouseId"`
}

// RemoveItemResponse is a automatically generated struct from json provided by sku vault's api docs.
type RemoveItemResponse struct {
	RemoveItemStatus string `json:"RemoveItemStatus"`
}
