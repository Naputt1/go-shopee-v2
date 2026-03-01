// Code generated. DO NOT EDIT.
package goshopee

type PaymentService interface {
	// GetEscrowDetail Use this API to fetch the accounting detail of order.
	// https://open.shopee.com/documents/v2/v2.payment.get_escrow_detail?module=97&type=1
	GetEscrowDetail(sid uint64, req GetEscrowDetailRequest, tok string) (*GetEscrowDetailResponse, error)
	// SetShopInstallmentStatus Sets the staging capability of shop level.
	// https://open.shopee.com/documents/v2/v2.payment.set_shop_installment_status?module=97&type=1
	SetShopInstallmentStatus(sid uint64, req SetShopInstallmentStatusRequest, tok string) (*SetShopInstallmentStatusResponse, error)
	// GetShopInstallmentStatus Get the installment state of shop.
	// https://open.shopee.com/documents/v2/v2.payment.get_shop_installment_status?module=97&type=1
	GetShopInstallmentStatus(sid uint64, tok string) (*GetShopInstallmentStatusResponse, error)
	// GetPayoutDetail This API is applicable for Cross Border (CB) sellers only to get the shop's payout data, such as the payout amount, currency, FX rate, the payout's associated order income and adjustment records etc.
	// https://open.shopee.com/documents/v2/v2.payment.get_payout_detail?module=97&type=1
	GetPayoutDetail(sid uint64, req GetPayoutDetailRequest, tok string) (*GetPayoutDetailResponse, error)
	// SetItemInstallmentStatus Set item installment.Only for TH、TW.
	// https://open.shopee.com/documents/v2/v2.payment.set_item_installment_status?module=97&type=1
	SetItemInstallmentStatus(sid uint64, req SetItemInstallmentStatusRequest, tok string) (*SetItemInstallmentStatusResponse, error)
	// GetItemInstallmentStatus Get item installment tenures.Only for TH、TW.
	// https://open.shopee.com/documents/v2/v2.payment.get_item_installment_status?module=97&type=1
	GetItemInstallmentStatus(sid uint64, req GetItemInstallmentStatusRequest, tok string) (*GetItemInstallmentStatusResponse, error)
	// GetPaymentMethodList Obtain payment method (no authentication required)
	// https://open.shopee.com/documents/v2/v2.payment.get_payment_method_list?module=97&type=1
	GetPaymentMethodList(sid uint64, tok string) (*GetPaymentMethodListResponse, error)
	// GetWalletTransactionList Use this API to get the transaction records of wallet. Only applicable for local shops
	// https://open.shopee.com/documents/v2/v2.payment.get_wallet_transaction_list?module=97&type=1
	GetWalletTransactionList(sid uint64, req GetWalletTransactionListRequest, tok string) (*GetWalletTransactionListResponse, error)
	// GetEscrowList Use this API to fetch the accounting list of order.
	// https://open.shopee.com/documents/v2/v2.payment.get_escrow_list?module=97&type=1
	GetEscrowList(sid uint64, req GetEscrowListRequest, tok string) (*GetEscrowListResponse, error)
	// GetPayoutInfo This is a new API which applicable for Cross Border (CB) sellers only to get the shop's payout data, will be used for the original API v2.get_payout_details replacement, we provide data such as the payout amount, currency, FX rate, the payout's associated order income and adjustment records etc.
	// https://open.shopee.com/documents/v2/v2.payment.get_payout_info?module=97&type=1
	GetPayoutInfo(sid uint64, req GetPayoutInfoRequest, tok string) (*GetPayoutInfoResponse, error)
	// GetBillingTransactionInfo This API is applicable for Cross Border (CB) sellers only to get the detailed payout transaction data, both released and to-be released transaction can be found in here
	// https://open.shopee.com/documents/v2/v2.payment.get_billing_transaction_info?module=97&type=1
	GetBillingTransactionInfo(sid uint64, req GetBillingTransactionInfoRequest, tok string) (*GetBillingTransactionInfoResponse, error)
	// GetEscrowDetailBatch Use this API to fetch the details of order income by batch.
	// https://open.shopee.com/documents/v2/v2.payment.get_escrow_detail_batch?module=97&type=1
	GetEscrowDetailBatch(sid uint64, req GetEscrowDetailBatchRequest, tok string) (*GetEscrowDetailBatchResponse, error)
	// GenerateIncomeStatement Trigger income statement generation.
	// https://open.shopee.com/documents/v2/v2.payment.generate_income_statement?module=97&type=1
	GenerateIncomeStatement(sid uint64, opt GenerateIncomeStatementRequest, tok string) (*GenerateIncomeStatementResponse, error)
	// GetIncomeStatement To query income statement status and provide file link if the income statement is ready to be downloaded.
	// https://open.shopee.com/documents/v2/v2.payment.get_income_statement?module=97&type=1
	GetIncomeStatement(sid uint64, opt GetIncomeStatementRequest, tok string) (*GetIncomeStatementResponse, error)
	// GenerateIncomeReport Trigger income report generation.
	// https://open.shopee.com/documents/v2/v2.payment.generate_income_report?module=97&type=1
	GenerateIncomeReport(sid uint64, req GenerateIncomeReportRequest, tok string) (*GenerateIncomeReportResponse, error)
	// GetIncomeReport To query income report status and provide file link if the income report is ready to be downloaded.
	// https://open.shopee.com/documents/v2/v2.payment.get_income_report?module=97&type=1
	GetIncomeReport(sid uint64, req GetIncomeReportRequest, tok string) (*GetIncomeReportResponse, error)
	// GetIncomeOverview Retrieves a consolidated snapshot of the seller’s income amounts categorized by income status for a specified shop. This API provides a holistic overview similar to Seller Center’s “Income Overview” section, allowing external systems to reflect the same current payout view.
	//
	// Data is dynamically determined based on the shop type (Local or Cross Border) and the income status requested. Historical income results are not retrievable, providing consistent information as Seller Centre.
	// https://open.shopee.com/documents/v2/v2.payment.get_income_overview?module=97&type=1
	GetIncomeOverview(sid uint64, opt GetIncomeOverviewRequest, tok string) (*GetIncomeOverviewResponse, error)
	// GetIncomeDetail Retrieves detailed order-level income information across various income statuses for a specified time period. This API enables partners to display granular transaction-level income data consistent with Seller Center’s “Income Details” view, segmented by income status and payout stage.
	//
	// The API dynamically adapts data fields based on the seller’s shop type (Local or Cross Border) and the selected income status (e.g., Pending, To Release, Released).
	// https://open.shopee.com/documents/v2/v2.payment.get_income_detail?module=97&type=1
	GetIncomeDetail(sid uint64, req GetIncomeDetailRequest, tok string) (*GetIncomeDetailResponse, error)
}

type PaymentServiceOp struct {
	client *Client
}

func (s *PaymentServiceOp) GetEscrowDetail(sid uint64, req GetEscrowDetailRequest, tok string) (*GetEscrowDetailResponse, error) {
	path := "/payment/get_escrow_detail"
	resp := new(GetEscrowDetailResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

func (s *PaymentServiceOp) SetShopInstallmentStatus(sid uint64, req SetShopInstallmentStatusRequest, tok string) (*SetShopInstallmentStatusResponse, error) {
	path := "/payment/set_shop_installment_status"
	resp := new(SetShopInstallmentStatusResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

func (s *PaymentServiceOp) GetShopInstallmentStatus(sid uint64, tok string) (*GetShopInstallmentStatusResponse, error) {
	path := "/payment/get_shop_installment_status"
	resp := new(GetShopInstallmentStatusResponse)
	err := s.client.WithShop(sid, tok).Post(path, nil, resp)
	return resp, err
}

func (s *PaymentServiceOp) GetPayoutDetail(sid uint64, req GetPayoutDetailRequest, tok string) (*GetPayoutDetailResponse, error) {
	path := "/payment/get_payout_detail"
	resp := new(GetPayoutDetailResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

func (s *PaymentServiceOp) SetItemInstallmentStatus(sid uint64, req SetItemInstallmentStatusRequest, tok string) (*SetItemInstallmentStatusResponse, error) {
	path := "/payment/set_item_installment_status"
	resp := new(SetItemInstallmentStatusResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

func (s *PaymentServiceOp) GetItemInstallmentStatus(sid uint64, req GetItemInstallmentStatusRequest, tok string) (*GetItemInstallmentStatusResponse, error) {
	path := "/payment/get_item_installment_status"
	resp := new(GetItemInstallmentStatusResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

func (s *PaymentServiceOp) GetPaymentMethodList(sid uint64, tok string) (*GetPaymentMethodListResponse, error) {
	path := "/payment/get_payment_method_list"
	resp := new(GetPaymentMethodListResponse)
	err := s.client.WithMerchant(sid, tok).Post(path, nil, resp)
	return resp, err
}

func (s *PaymentServiceOp) GetWalletTransactionList(sid uint64, req GetWalletTransactionListRequest, tok string) (*GetWalletTransactionListResponse, error) {
	path := "/payment/get_wallet_transaction_list"
	resp := new(GetWalletTransactionListResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

func (s *PaymentServiceOp) GetEscrowList(sid uint64, req GetEscrowListRequest, tok string) (*GetEscrowListResponse, error) {
	path := "/payment/get_escrow_list"
	resp := new(GetEscrowListResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

func (s *PaymentServiceOp) GetPayoutInfo(sid uint64, req GetPayoutInfoRequest, tok string) (*GetPayoutInfoResponse, error) {
	path := "/payment/get_payout_info"
	resp := new(GetPayoutInfoResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

func (s *PaymentServiceOp) GetBillingTransactionInfo(sid uint64, req GetBillingTransactionInfoRequest, tok string) (*GetBillingTransactionInfoResponse, error) {
	path := "/payment/get_billing_transaction_info"
	resp := new(GetBillingTransactionInfoResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

func (s *PaymentServiceOp) GetEscrowDetailBatch(sid uint64, req GetEscrowDetailBatchRequest, tok string) (*GetEscrowDetailBatchResponse, error) {
	path := "/payment/get_escrow_detail_batch"
	resp := new(GetEscrowDetailBatchResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

func (s *PaymentServiceOp) GenerateIncomeStatement(sid uint64, opt GenerateIncomeStatementRequest, tok string) (*GenerateIncomeStatementResponse, error) {
	path := "/payment/generate_income_statement"
	resp := new(GenerateIncomeStatementResponse)
	err := s.client.WithShop(sid, tok).Get(path, resp, opt)
	return resp, err
}

func (s *PaymentServiceOp) GetIncomeStatement(sid uint64, opt GetIncomeStatementRequest, tok string) (*GetIncomeStatementResponse, error) {
	path := "/payment/get_income_statement"
	resp := new(GetIncomeStatementResponse)
	err := s.client.WithShop(sid, tok).Get(path, resp, opt)
	return resp, err
}

func (s *PaymentServiceOp) GenerateIncomeReport(sid uint64, req GenerateIncomeReportRequest, tok string) (*GenerateIncomeReportResponse, error) {
	path := "/payment/generate_income_report"
	resp := new(GenerateIncomeReportResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

func (s *PaymentServiceOp) GetIncomeReport(sid uint64, req GetIncomeReportRequest, tok string) (*GetIncomeReportResponse, error) {
	path := "/payment/get_income_report"
	resp := new(GetIncomeReportResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

func (s *PaymentServiceOp) GetIncomeOverview(sid uint64, opt GetIncomeOverviewRequest, tok string) (*GetIncomeOverviewResponse, error) {
	path := "/payment/get_income_overview"
	resp := new(GetIncomeOverviewResponse)
	err := s.client.WithShop(sid, tok).Get(path, resp, opt)
	return resp, err
}

func (s *PaymentServiceOp) GetIncomeDetail(sid uint64, req GetIncomeDetailRequest, tok string) (*GetIncomeDetailResponse, error) {
	path := "/payment/get_income_detail"
	resp := new(GetIncomeDetailResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}
