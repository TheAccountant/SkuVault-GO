package sales

import "time"

// CreateHolds is a automatically generated struct from json provided by sku vault's api docs.
type CreateHolds struct {
	Holds []struct {
		Description       string    `json:"Description"`
		ExpirationDateUtc time.Time `json:"ExpirationDateUtc"`
		Quantity          int       `json:"Quantity"`
		Sku               string    `json:"Sku"`
	} `json:"Holds"`
}

// CreateHoldsResponse response from CreateHolds.
type CreateHoldsResponse []struct {
	Quantity int    `json:"Quantity"`
	Sku      string `json:"Sku"`
}
