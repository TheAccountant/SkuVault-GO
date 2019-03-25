package sales

// CreateHolds is a automatically generated struct from json provided by sku vault's api docs.
type CreateHolds struct {
	Holds []struct {
		Description       string `json:"Description"`
		ExpirationDateUtc string `json:"ExpirationDateUtc"`
		Quantity          int64  `json:"Quantity"`
		Sku               string `json:"Sku"`
	} `json:"Holds"`
}

type CreateHoldsResponse []struct {
	Quantity int64  `json:"Quantity"`
	Sku      string `json:"Sku"`
}