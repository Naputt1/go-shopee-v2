// Code generated. DO NOT EDIT.
package goshopee

type SBSService interface {
	// GetBoundWhsInfo get bound warehouse by shop id
	// https://open.shopee.com/documents/v2/v2.sbs.get_bound_whs_info?module=124&type=1
	GetBoundWhsInfo(sid uint64, tok string) (*GetBoundWhsInfoResponse, error)
	// GetCurrentInventory Get Seller Center Current Inventory Page Data
	// https://open.shopee.com/documents/v2/v2.sbs.get_current_inventory?module=124&type=1
	GetCurrentInventory(sid uint64, req GetCurrentInventoryRequest, tok string) (*GetCurrentInventoryResponse, error)
	// GetExpiryReport Seller Center Expiry Report page data
	// https://open.shopee.com/documents/v2/v2.sbs.get_expiry_report?module=124&type=1
	GetExpiryReport(sid uint64, req GetExpiryReportRequest, tok string) (*GetExpiryReportResponse, error)
	// GetStockAging Get Seller Center Stock Aging page data
	// https://open.shopee.com/documents/v2/v2.sbs.get_stock_aging?module=124&type=1
	GetStockAging(sid uint64, req GetStockAgingRequest, tok string) (*GetStockAgingResponse, error)
	// GetStockMovement Get Seller Center，Stock Movement page data
	// https://open.shopee.com/documents/v2/v2.sbs.get_stock_movement?module=124&type=1
	GetStockMovement(sid uint64, req GetStockMovementRequest, tok string) (*GetStockMovementResponse, error)
}

type SBSServiceOp struct {
	client *Client
}

func (s *SBSServiceOp) GetBoundWhsInfo(sid uint64, tok string) (*GetBoundWhsInfoResponse, error) {
	path := "/sbs/get_bound_whs_info"
	resp := new(GetBoundWhsInfoResponse)
	err := s.client.WithShop(sid, tok).Get(path, resp, nil)
	return resp, err
}

func (s *SBSServiceOp) GetCurrentInventory(sid uint64, req GetCurrentInventoryRequest, tok string) (*GetCurrentInventoryResponse, error) {
	path := "/sbs/get_current_inventory"
	resp := new(GetCurrentInventoryResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

func (s *SBSServiceOp) GetExpiryReport(sid uint64, req GetExpiryReportRequest, tok string) (*GetExpiryReportResponse, error) {
	path := "/sbs/get_expiry_report"
	resp := new(GetExpiryReportResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

func (s *SBSServiceOp) GetStockAging(sid uint64, req GetStockAgingRequest, tok string) (*GetStockAgingResponse, error) {
	path := "/sbs/get_stock_aging"
	resp := new(GetStockAgingResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

func (s *SBSServiceOp) GetStockMovement(sid uint64, req GetStockMovementRequest, tok string) (*GetStockMovementResponse, error) {
	path := "/sbs/get_stock_movement"
	resp := new(GetStockMovementResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}
