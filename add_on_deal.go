// Code generated. DO NOT EDIT.
package goshopee

type AddOnDealService interface {
	// AddAddOnDeal Add Add-on Deal
	// https://open.shopee.com/documents/v2/v2.add_on_deal.add_add_on_deal?module=111&type=1
	AddAddOnDeal(sid uint64, req AddAddOnDealRequest, tok string) (*AddAddOnDealResponse, error)
	// AddAddOnDealMainItem Add Add-on Deal Main Item
	// https://open.shopee.com/documents/v2/v2.add_on_deal.add_add_on_deal_main_item?module=111&type=1
	AddAddOnDealMainItem(sid uint64, req AddAddOnDealMainItemRequest, tok string) (*AddAddOnDealMainItemResponse, error)
	// AddAddOnDealSubItem Add Add-on Deal Sub Item
	// https://open.shopee.com/documents/v2/v2.add_on_deal.add_add_on_deal_sub_item?module=111&type=1
	AddAddOnDealSubItem(sid uint64, req AddAddOnDealSubItemRequest, tok string) (*AddAddOnDealSubItemResponse, error)
	// DeleteAddOnDeal Delete Add-on Deal
	// https://open.shopee.com/documents/v2/v2.add_on_deal.delete_add_on_deal?module=111&type=1
	DeleteAddOnDeal(sid uint64, req DeleteAddOnDealRequest, tok string) (*DeleteAddOnDealResponse, error)
	// DeleteAddOnDealMainItem Delete Add-on Deal Main Item
	// https://open.shopee.com/documents/v2/v2.add_on_deal.delete_add_on_deal_main_item?module=111&type=1
	DeleteAddOnDealMainItem(sid uint64, req DeleteAddOnDealMainItemRequest, tok string) (*DeleteAddOnDealMainItemResponse, error)
	// DeleteAddOnDealSubItem Delete Add-on Deal Sub Item
	// https://open.shopee.com/documents/v2/v2.add_on_deal.delete_add_on_deal_sub_item?module=111&type=1
	DeleteAddOnDealSubItem(sid uint64, req DeleteAddOnDealSubItemRequest, tok string) (*DeleteAddOnDealSubItemResponse, error)
	// GetAddOnDealList Get Add-on Deal List
	// https://open.shopee.com/documents/v2/v2.add_on_deal.get_add_on_deal_list?module=111&type=1
	GetAddOnDealList(sid uint64, req GetAddOnDealListRequest, tok string) (*GetAddOnDealListResponse, error)
	// GetAddOnDeal Get Add-on Deal
	// https://open.shopee.com/documents/v2/v2.add_on_deal.get_add_on_deal?module=111&type=1
	GetAddOnDeal(sid uint64, req GetAddOnDealRequest, tok string) (*GetAddOnDealResponse, error)
	// GetAddOnDealMainItem Get Add-on Deal Main Item
	// https://open.shopee.com/documents/v2/v2.add_on_deal.get_add_on_deal_main_item?module=111&type=1
	GetAddOnDealMainItem(sid uint64, req GetAddOnDealMainItemRequest, tok string) (*GetAddOnDealMainItemResponse, error)
	// GetAddOnDealSubItem Get Add-on Deal Sub Item
	// https://open.shopee.com/documents/v2/v2.add_on_deal.get_add_on_deal_sub_item?module=111&type=1
	GetAddOnDealSubItem(sid uint64, req GetAddOnDealSubItemRequest, tok string) (*GetAddOnDealSubItemResponse, error)
	// UpdateAddOnDeal Update Add-on Deal
	// https://open.shopee.com/documents/v2/v2.add_on_deal.update_add_on_deal?module=111&type=1
	UpdateAddOnDeal(sid uint64, req UpdateAddOnDealRequest, tok string) (*UpdateAddOnDealResponse, error)
	// UpdateAddOnDealMainItem Update Add-on Deal Main Item
	// https://open.shopee.com/documents/v2/v2.add_on_deal.update_add_on_deal_main_item?module=111&type=1
	UpdateAddOnDealMainItem(sid uint64, req UpdateAddOnDealMainItemRequest, tok string) (*UpdateAddOnDealMainItemResponse, error)
	// UpdateAddOnDealSubItem Update Add-on Deal Sub Item
	// https://open.shopee.com/documents/v2/v2.add_on_deal.update_add_on_deal_sub_item?module=111&type=1
	UpdateAddOnDealSubItem(sid uint64, req UpdateAddOnDealSubItemRequest, tok string) (*UpdateAddOnDealSubItemResponse, error)
	// EndAddOnDeal End Add-on Deal
	// https://open.shopee.com/documents/v2/v2.add_on_deal.end_add_on_deal?module=111&type=1
	EndAddOnDeal(sid uint64, req EndAddOnDealRequest, tok string) (*EndAddOnDealResponse, error)
}

type AddOnDealServiceOp struct {
	client *Client
}

func (s *AddOnDealServiceOp) AddAddOnDeal(sid uint64, req AddAddOnDealRequest, tok string) (*AddAddOnDealResponse, error) {
	path := "/add_on_deal/add_add_on_deal"
	resp := new(AddAddOnDealResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

func (s *AddOnDealServiceOp) AddAddOnDealMainItem(sid uint64, req AddAddOnDealMainItemRequest, tok string) (*AddAddOnDealMainItemResponse, error) {
	path := "/add_on_deal/add_add_on_deal_main_item"
	resp := new(AddAddOnDealMainItemResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

func (s *AddOnDealServiceOp) AddAddOnDealSubItem(sid uint64, req AddAddOnDealSubItemRequest, tok string) (*AddAddOnDealSubItemResponse, error) {
	path := "/add_on_deal/add_add_on_deal_sub_item"
	resp := new(AddAddOnDealSubItemResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

func (s *AddOnDealServiceOp) DeleteAddOnDeal(sid uint64, req DeleteAddOnDealRequest, tok string) (*DeleteAddOnDealResponse, error) {
	path := "/add_on_deal/delete_add_on_deal"
	resp := new(DeleteAddOnDealResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

func (s *AddOnDealServiceOp) DeleteAddOnDealMainItem(sid uint64, req DeleteAddOnDealMainItemRequest, tok string) (*DeleteAddOnDealMainItemResponse, error) {
	path := "/add_on_deal/delete_add_on_deal_main_item"
	resp := new(DeleteAddOnDealMainItemResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

func (s *AddOnDealServiceOp) DeleteAddOnDealSubItem(sid uint64, req DeleteAddOnDealSubItemRequest, tok string) (*DeleteAddOnDealSubItemResponse, error) {
	path := "/add_on_deal/delete_add_on_deal_sub_item"
	resp := new(DeleteAddOnDealSubItemResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

func (s *AddOnDealServiceOp) GetAddOnDealList(sid uint64, req GetAddOnDealListRequest, tok string) (*GetAddOnDealListResponse, error) {
	path := "/add_on_deal/get_add_on_deal_list"
	resp := new(GetAddOnDealListResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

func (s *AddOnDealServiceOp) GetAddOnDeal(sid uint64, req GetAddOnDealRequest, tok string) (*GetAddOnDealResponse, error) {
	path := "/add_on_deal/get_add_on_deal"
	resp := new(GetAddOnDealResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

func (s *AddOnDealServiceOp) GetAddOnDealMainItem(sid uint64, req GetAddOnDealMainItemRequest, tok string) (*GetAddOnDealMainItemResponse, error) {
	path := "/add_on_deal/get_add_on_deal_main_item"
	resp := new(GetAddOnDealMainItemResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

func (s *AddOnDealServiceOp) GetAddOnDealSubItem(sid uint64, req GetAddOnDealSubItemRequest, tok string) (*GetAddOnDealSubItemResponse, error) {
	path := "/add_on_deal/get_add_on_deal_sub_item"
	resp := new(GetAddOnDealSubItemResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

func (s *AddOnDealServiceOp) UpdateAddOnDeal(sid uint64, req UpdateAddOnDealRequest, tok string) (*UpdateAddOnDealResponse, error) {
	path := "/add_on_deal/update_add_on_deal"
	resp := new(UpdateAddOnDealResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

func (s *AddOnDealServiceOp) UpdateAddOnDealMainItem(sid uint64, req UpdateAddOnDealMainItemRequest, tok string) (*UpdateAddOnDealMainItemResponse, error) {
	path := "/add_on_deal/update_add_on_deal_main_item"
	resp := new(UpdateAddOnDealMainItemResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

func (s *AddOnDealServiceOp) UpdateAddOnDealSubItem(sid uint64, req UpdateAddOnDealSubItemRequest, tok string) (*UpdateAddOnDealSubItemResponse, error) {
	path := "/add_on_deal/update_add_on_deal_sub_item"
	resp := new(UpdateAddOnDealSubItemResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

func (s *AddOnDealServiceOp) EndAddOnDeal(sid uint64, req EndAddOnDealRequest, tok string) (*EndAddOnDealResponse, error) {
	path := "/add_on_deal/end_add_on_deal"
	resp := new(EndAddOnDealResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}
