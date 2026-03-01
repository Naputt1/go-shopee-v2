// Code generated. DO NOT EDIT.
package goshopee

type VoucherService interface {
	// AddVoucher Add voucher
	// https://open.shopee.com/documents/v2/v2.voucher.add_voucher?module=112&type=1
	AddVoucher(sid uint64, req AddVoucherRequest, tok string) (*AddVoucherResponse, error)
	// DeleteVoucher Delete voucher
	// https://open.shopee.com/documents/v2/v2.voucher.delete_voucher?module=112&type=1
	DeleteVoucher(sid uint64, req DeleteVoucherRequest, tok string) (*DeleteVoucherResponse, error)
	// EndVoucher End Voucher
	// https://open.shopee.com/documents/v2/v2.voucher.end_voucher?module=112&type=1
	EndVoucher(sid uint64, req EndVoucherRequest, tok string) (*EndVoucherResponse, error)
	// UpdateVoucher Update voucher
	// https://open.shopee.com/documents/v2/v2.voucher.update_voucher?module=112&type=1
	UpdateVoucher(sid uint64, req UpdateVoucherRequest, tok string) (*UpdateVoucherResponse, error)
	// GetVoucher Get Voucher Detail
	// https://open.shopee.com/documents/v2/v2.voucher.get_voucher?module=112&type=1
	GetVoucher(sid uint64, req GetVoucherRequest, tok string) (*GetVoucherResponse, error)
	// GetVoucherList Get Voucher List
	// https://open.shopee.com/documents/v2/v2.voucher.get_voucher_list?module=112&type=1
	GetVoucherList(sid uint64, req GetVoucherListRequest, tok string) (*GetVoucherListResponse, error)
}

type VoucherServiceOp struct {
	client *Client
}

func (s *VoucherServiceOp) AddVoucher(sid uint64, req AddVoucherRequest, tok string) (*AddVoucherResponse, error) {
	path := "/voucher/add_voucher"
	resp := new(AddVoucherResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

func (s *VoucherServiceOp) DeleteVoucher(sid uint64, req DeleteVoucherRequest, tok string) (*DeleteVoucherResponse, error) {
	path := "/voucher/delete_voucher"
	resp := new(DeleteVoucherResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

func (s *VoucherServiceOp) EndVoucher(sid uint64, req EndVoucherRequest, tok string) (*EndVoucherResponse, error) {
	path := "/voucher/end_voucher"
	resp := new(EndVoucherResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

func (s *VoucherServiceOp) UpdateVoucher(sid uint64, req UpdateVoucherRequest, tok string) (*UpdateVoucherResponse, error) {
	path := "/voucher/update_voucher"
	resp := new(UpdateVoucherResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

func (s *VoucherServiceOp) GetVoucher(sid uint64, req GetVoucherRequest, tok string) (*GetVoucherResponse, error) {
	path := "/voucher/get_voucher"
	resp := new(GetVoucherResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

func (s *VoucherServiceOp) GetVoucherList(sid uint64, req GetVoucherListRequest, tok string) (*GetVoucherListResponse, error) {
	path := "/voucher/get_voucher_list"
	resp := new(GetVoucherListResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}
