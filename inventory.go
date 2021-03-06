package skuvault

import "github.com/OuttaLineNomad/skuvault/inventory"

// postAddItem payload sent to Sku Vault.
type postAddItem struct {
	*inventory.AddItem
	*ILoginCredentials
}

// AddItem creates http request for this SKU vault endpoint.
// Moderate Throttle.
// Bulk version available.
func (lc *ILoginCredentials) AddItem(pld *inventory.AddItem) (*inventory.AddItemResponse, error) {
	credPld := &postAddItem{
		AddItem:           pld,
		ILoginCredentials: lc,
	}

	response := &inventory.AddItemResponse{}
	err := do(credPld, response, "inventory/addItem")
	if err != nil {
		return nil, &Error{"AddItem()", err}
	}

	return response, nil
}

// postAddItemBulk payload sent to Sku Vault.
type postAddItemBulk struct {
	*inventory.AddItemBulk
	*ILoginCredentials
}

// AddItemBulk creates http request for this SKU vault endpoint.
// Moderate Throttle.
// Add quantity to warehouse locations, 100 at a time..
func (lc *ILoginCredentials) AddItemBulk(pld *inventory.AddItemBulk) (*inventory.AddItemBulkResponse, error) {
	credPld := &postAddItemBulk{
		AddItemBulk:       pld,
		ILoginCredentials: lc,
	}

	response := &inventory.AddItemBulkResponse{}
	err := do(credPld, response, "inventory/addItemBulk")
	if err != nil {
		return nil, &Error{"AddItemBulk()", err}
	}

	return response, nil
}

// postGetAvailableQuantities payload sent to Sku Vault.
type postGetAvailableQuantities struct {
	*inventory.GetAvailableQuantities
	*ILoginCredentials
}

// GetAvailableQuantities creates http request for this SKU vault endpoint.
// Heavy Throttle.
// Retrieve a list of SKUs and their total available quantities across all warehouses. Available Quantity is the quantity that is actually available to sell across all your sales channels..
func (lc *ILoginCredentials) GetAvailableQuantities(pld *inventory.GetAvailableQuantities) (*inventory.GetAvailableQuantitiesResponse, error) {
	credPld := &postGetAvailableQuantities{
		GetAvailableQuantities: pld,
		ILoginCredentials:      lc,
	}

	response := &inventory.GetAvailableQuantitiesResponse{}
	err := do(credPld, response, "inventory/getAvailableQuantities")
	if err != nil {
		return nil, &Error{"GetAvailableQuantities()", err}
	}

	return response, nil
}

// postGetExternalWarehouseQuantities payload sent to Sku Vault.
type postGetExternalWarehouseQuantities struct {
	*inventory.GetExternalWarehouseQuantities
	*ILoginCredentials
}

// GetExternalWarehouseQuantities creates http request for this SKU vault endpoint.
// Heavy Throttle.
// This call gets the quantities in a designated External Warehouse. Please note these are different than ordinary warehouses..
func (lc *ILoginCredentials) GetExternalWarehouseQuantities(pld *inventory.GetExternalWarehouseQuantities) (*inventory.GetExternalWarehouseQuantitiesResponse, error) {
	credPld := &postGetExternalWarehouseQuantities{
		GetExternalWarehouseQuantities: pld,
		ILoginCredentials:              lc,
	}

	response := &inventory.GetExternalWarehouseQuantitiesResponse{}
	err := do(credPld, response, "inventory/getExternalWarehouseQuantities")
	if err != nil {
		return nil, &Error{"GetExternalWarehouseQuantities()", err}
	}

	return response, nil
}

// postGetExternalWarehouses payload sent to Sku Vault.
type postGetExternalWarehouses struct {
	*ILoginCredentials
}

// GetExternalWarehouses creates http request for this SKU vault endpoint.
// Moderate Throttle.
// Returns your external warehouses. No page parameters..
func (lc *ILoginCredentials) GetExternalWarehouses() (*inventory.GetExternalWarehousesResponse, error) {
	credPld := &postGetExternalWarehouses{
		ILoginCredentials: lc,
	}

	response := &inventory.GetExternalWarehousesResponse{}
	err := do(credPld, response, "inventory/getExternalWarehouses")
	if err != nil {
		return nil, &Error{"GetExternalWarehouses()", err}
	}

	return response, nil
}

// postGetInventoryByLocation payload sent to Sku Vault.
type postGetInventoryByLocation struct {
	*inventory.GetInventoryByLocation
	*ILoginCredentials
}

// GetInventoryByLocation creates http request for this SKU vault endpoint.
// Heavy Throttle.
// Returns location and warehouse per product..
func (lc *ILoginCredentials) GetInventoryByLocation(pld *inventory.GetInventoryByLocation) (*inventory.GetInventoryByLocationResponse, error) {
	credPld := &postGetInventoryByLocation{
		GetInventoryByLocation: pld,
		ILoginCredentials:      lc,
	}

	response := &inventory.GetInventoryByLocationResponse{}
	err := do(credPld, response, "inventory/getInventoryByLocation")
	if err != nil {
		return nil, &Error{"GetInventoryByLocation()", err}
	}

	return response, nil
}

// postGetItemQuantities payload sent to Sku Vault.
type postGetItemQuantities struct {
	*inventory.GetItemQuantities
	*ILoginCredentials
}

// GetItemQuantities creates http request for this SKU vault endpoint.
// Heavy Throttle.
// Returns product quantities..
func (lc *ILoginCredentials) GetItemQuantities(pld *inventory.GetItemQuantities) (*inventory.GetItemQuantitiesResponse, error) {
	credPld := &postGetItemQuantities{
		GetItemQuantities: pld,
		ILoginCredentials: lc,
	}

	response := &inventory.GetItemQuantitiesResponse{}
	err := do(credPld, response, "inventory/getItemQuantities")
	if err != nil {
		return nil, &Error{"GetItemQuantities()", err}
	}

	return response, nil
}

// postGetKitQuantities payload sent to Sku Vault.
type postGetKitQuantities struct {
	*inventory.GetKitQuantities
	*ILoginCredentials
}

// GetKitQuantities creates http request for this SKU vault endpoint.
// Heavy Throttle.
// Returns kit quantities..
func (lc *ILoginCredentials) GetKitQuantities(pld *inventory.GetKitQuantities) (*inventory.GetKitQuantitiesResponse, error) {
	credPld := &postGetKitQuantities{
		GetKitQuantities:  pld,
		ILoginCredentials: lc,
	}

	response := &inventory.GetKitQuantitiesResponse{}
	err := do(credPld, response, "inventory/getKitQuantities")
	if err != nil {
		return nil, &Error{"GetKitQuantities()", err}
	}

	return response, nil
}

// postGetLocations payload sent to Sku Vault.
type postGetLocations struct {
	*ILoginCredentials
}

// GetLocations creates http request for this SKU vault endpoint.
// Severe Throttle.
//  throttling. Returns all your locations in enabled warehouses..
func (lc *ILoginCredentials) GetLocations() (*inventory.GetLocationsResponse, error) {
	credPld := &postGetLocations{
		ILoginCredentials: lc,
	}

	response := &inventory.GetLocationsResponse{}
	err := do(credPld, response, "inventory/getLocations")
	if err != nil {
		return nil, &Error{"GetLocations()", err}
	}

	return response, nil
}

// postGetTransactions payload sent to Sku Vault.
type postGetTransactions struct {
	*inventory.GetTransactions
	*ILoginCredentials
}

// GetTransactions creates http request for this SKU vault endpoint.
// Heavy Throttle.
// Look at your transaction history..
func (lc *ILoginCredentials) GetTransactions(pld *inventory.GetTransactions) (*inventory.GetTransactionsResponse, error) {
	credPld := &postGetTransactions{
		GetTransactions:   pld,
		ILoginCredentials: lc,
	}

	response := &inventory.GetTransactionsResponse{}
	err := do(credPld, response, "inventory/getTransactions")
	if err != nil {
		return nil, &Error{"GetTransactions()", err}
	}

	return response, nil
}

// postGetWarehouseItemQuantities payload sent to Sku Vault.
type postGetWarehouseItemQuantities struct {
	*inventory.GetWarehouseItemQuantities
	*ILoginCredentials
}

// GetWarehouseItemQuantities creates http request for this SKU vault endpoint.
// Heavy Throttle.
// This call returns SKUs and quantities from a specified warehouse. 10,000 SKUs returned per page..
func (lc *ILoginCredentials) GetWarehouseItemQuantities(pld *inventory.GetWarehouseItemQuantities) (*inventory.GetWarehouseItemQuantitiesResponse, error) {
	credPld := &postGetWarehouseItemQuantities{
		GetWarehouseItemQuantities: pld,
		ILoginCredentials:          lc,
	}

	response := &inventory.GetWarehouseItemQuantitiesResponse{}
	err := do(credPld, response, "inventory/getWarehouseItemQuantities")
	if err != nil {
		return nil, &Error{"GetWarehouseItemQuantities()", err}
	}

	return response, nil
}

// postGetWarehouseItemQuantity payload sent to Sku Vault.
type postGetWarehouseItemQuantity struct {
	*inventory.GetWarehouseItemQuantity
	*ILoginCredentials
}

// GetWarehouseItemQuantity creates http request for this SKU vault endpoint.
// Heavy Throttle.
// Returns the quantity for a specified SKU..
func (lc *ILoginCredentials) GetWarehouseItemQuantity(pld *inventory.GetWarehouseItemQuantity) (*inventory.GetWarehouseItemQuantityResponse, error) {
	credPld := &postGetWarehouseItemQuantity{
		GetWarehouseItemQuantity: pld,
		ILoginCredentials:        lc,
	}

	response := &inventory.GetWarehouseItemQuantityResponse{}
	err := do(credPld, response, "inventory/getWarehouseItemQuantity")
	if err != nil {
		return nil, &Error{"GetWarehouseItemQuantity()", err}
	}

	return response, nil
}

// postGetWarehouses payload sent to Sku Vault.
type postGetWarehouses struct {
	*inventory.GetWarehouses
	*ILoginCredentials
}

// GetWarehouses creates http request for this SKU vault endpoint.
// Severe Throttle.
//  throttling. Returns all your regular warehouses..
func (lc *ILoginCredentials) GetWarehouses(pld *inventory.GetWarehouses) (*inventory.GetWarehousesResponse, error) {
	credPld := &postGetWarehouses{
		GetWarehouses:     pld,
		ILoginCredentials: lc,
	}

	response := &inventory.GetWarehousesResponse{}
	err := do(credPld, response, "inventory/getWarehouses")
	if err != nil {
		return nil, &Error{"GetWarehouses()", err}
	}

	return response, nil
}

// postPickItem payload sent to Sku Vault.
type postPickItem struct {
	*inventory.PickItem
	*ILoginCredentials
}

// PickItem creates http request for this SKU vault endpoint.
// Moderate Throttle.
// Perform a pick transaction through the API. [Bulk version available].
func (lc *ILoginCredentials) PickItem(pld *inventory.PickItem) (*inventory.PickItemResponse, error) {
	credPld := &postPickItem{
		PickItem:          pld,
		ILoginCredentials: lc,
	}

	response := &inventory.PickItemResponse{}
	err := do(credPld, response, "inventory/pickItem")
	if err != nil {
		return nil, &Error{"PickItem()", err}
	}

	return response, nil
}

// postPickItemBulk payload sent to Sku Vault.
type postPickItemBulk struct {
	*inventory.PickItemBulk
	*ILoginCredentials
}

// PickItemBulk creates http request for this SKU vault endpoint.
// Moderate Throttle.
// Perform a bulk pick transaction through the API..
func (lc *ILoginCredentials) PickItemBulk(pld *inventory.PickItemBulk) (*inventory.PickItemBulkResponse, error) {
	credPld := &postPickItemBulk{
		PickItemBulk:      pld,
		ILoginCredentials: lc,
	}

	response := &inventory.PickItemBulkResponse{}
	err := do(credPld, response, "inventory/pickItemBulk")
	if err != nil {
		return nil, &Error{"PickItemBulk()", err}
	}

	return response, nil
}

// postRemoveItem payload sent to Sku Vault.
type postRemoveItem struct {
	*inventory.RemoveItem
	*ILoginCredentials
}

// RemoveItem creates http request for this SKU vault endpoint.
// Moderate Throttle.
// Bulk version available.
func (lc *ILoginCredentials) RemoveItem(pld *inventory.RemoveItem) (*inventory.RemoveItemResponse, error) {
	credPld := &postRemoveItem{
		RemoveItem:        pld,
		ILoginCredentials: lc,
	}

	response := &inventory.RemoveItemResponse{}
	err := do(credPld, response, "inventory/removeItem")
	if err != nil {
		return nil, &Error{"RemoveItem()", err}
	}

	return response, nil
}

// postRemoveItemBulk payload sent to Sku Vault.
type postRemoveItemBulk struct {
	*inventory.RemoveItemBulk
	*ILoginCredentials
}

// RemoveItemBulk creates http request for this SKU vault endpoint.
// Moderate Throttle.
// Remove quantity from warehouse locations, 100 at a time..
func (lc *ILoginCredentials) RemoveItemBulk(pld *inventory.RemoveItemBulk) (*inventory.RemoveItemBulkResponse, error) {
	credPld := &postRemoveItemBulk{
		RemoveItemBulk:    pld,
		ILoginCredentials: lc,
	}

	response := &inventory.RemoveItemBulkResponse{}
	err := do(credPld, response, "inventory/removeItemBulk")
	if err != nil {
		return nil, &Error{"RemoveItemBulk()", err}
	}

	return response, nil
}

// postSetItemQuantities payload sent to Sku Vault.
type postSetItemQuantities struct {
	*inventory.SetItemQuantities
	*ILoginCredentials
}

// SetItemQuantities creates http request for this SKU vault endpoint.
// Moderate Throttle.
// Sets quantity for multiple products in one request..
func (lc *ILoginCredentials) SetItemQuantities(pld *inventory.SetItemQuantities) (*inventory.SetItemQuantitiesResponse, error) {
	credPld := &postSetItemQuantities{
		SetItemQuantities: pld,
		ILoginCredentials: lc,
	}

	response := &inventory.SetItemQuantitiesResponse{}
	err := do(credPld, response, "inventory/setItemQuantities")
	if err != nil {
		return nil, &Error{"SetItemQuantities()", err}
	}

	return response, nil
}

// postSetItemQuantity payload sent to Sku Vault.
type postSetItemQuantity struct {
	*inventory.SetItemQuantity
	*ILoginCredentials
}

// SetItemQuantity creates http request for this SKU vault endpoint.
// Moderate Throttle.
// This lets you explicitly set quantity for an item in a warehouse&#39;s location..
func (lc *ILoginCredentials) SetItemQuantity(pld *inventory.SetItemQuantity) (*inventory.SetItemQuantityResponse, error) {
	credPld := &postSetItemQuantity{
		SetItemQuantity:   pld,
		ILoginCredentials: lc,
	}

	response := &inventory.SetItemQuantityResponse{}
	err := do(credPld, response, "inventory/setItemQuantity")
	if err != nil {
		return nil, &Error{"SetItemQuantity()", err}
	}

	return response, nil
}

// postUpdateExternalWarehouseQuantities payload sent to Sku Vault.
type postUpdateExternalWarehouseQuantities struct {
	*inventory.UpdateExternalWarehouseQuantities
	*ILoginCredentials
}

// UpdateExternalWarehouseQuantities creates http request for this SKU vault endpoint.
// Severe Throttle.
// Set the quantity of SKUs in a specified external warehouse. Updating an external warehouse&#39;s quantities will overwrite any existing quantities. The limit is 200,000 SKUs per call..
func (lc *ILoginCredentials) UpdateExternalWarehouseQuantities(pld *inventory.UpdateExternalWarehouseQuantities) (*inventory.UpdateExternalWarehouseQuantitiesResponse, error) {
	credPld := &postUpdateExternalWarehouseQuantities{
		UpdateExternalWarehouseQuantities: pld,
		ILoginCredentials:                 lc,
	}

	response := &inventory.UpdateExternalWarehouseQuantitiesResponse{}
	err := do(credPld, response, "inventory/updateExternalWarehouseQuantities")
	if err != nil {
		return nil, &Error{"UpdateExternalWarehouseQuantities()", err}
	}

	return response, nil
}
