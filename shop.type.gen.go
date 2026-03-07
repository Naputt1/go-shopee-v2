package goshopee

type AuthorisedBrand struct {
	BrandId   int64  `json:"brand_id"`   // [Required] <p>ID of the authorised brand, it may be the same in different regions.</p>
	BrandName string `json:"brand_name"` // [Required] <p>Name of the authorised brand.</p>
}

type BillingAddress struct {
	State        string `json:"state"`        // [Required] <p>State of the billing address.</p>
	City         string `json:"city"`         // [Required] <p>City of&nbsp;the billing address.</p>
	Address      string `json:"address"`      // [Required] <p>Specific detail of&nbsp;the billing address.</p>
	Zipcode      string `json:"zipcode"`      // [Required] <p>ZIP code&nbsp;of&nbsp;the billing address.</p>
	Neighborhood string `json:"neighborhood"` // [Required] <p>Neighborhood&nbsp;of&nbsp;the billing address.</p>
}

type GetAuthorisedResellerBrandRequest struct {
	PageNo   int64 `json:"page_no" url:"page_no"`     // [Required] <p>Specifies the page number of data to return in the current call. Starting from 1. if data is more than one page, the page_no can be some entry to start next call.<br /></p>
	PageSize int64 `json:"page_size" url:"page_size"` // [Required] <p>Each result set is returned as a page of entries. Use the "page_size" filters to control the maximum number of entries to retrieve per page (i.e., per call), and the "page_no" to start next call. This integer value is used to specify the maximum number of entries to return in a single "page" of data. The limit of page_size if between 1 and 30.<br /></p>
}

type GetAuthorisedResellerBrandResponse struct {
	BaseResponse `json:",inline"`                       // Common response fields
	Response     GetAuthorisedResellerBrandResponseData `json:"response"` // Actual response data
}

type GetAuthorisedResellerBrandResponseData struct {
	IsAuthorisedReseller bool              `json:"is_authorised_reseller"` // [Required] <p>This is to indicate whether the shop is authorised reseller.</p>
	TotalCount           int64             `json:"total_count"`            // [Required] <p>The number of authorised brand linked with the shop.</p>
	More                 bool              `json:"more"`                   // [Required] <p>This is to indicate whether the authorised brand list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of authorised brand.<br /></p>
	AuthorisedBrandList  []AuthorisedBrand `json:"authorised_brand_list"`  // [Required]
}

type GetBrShopOnboardingInfoResponse struct {
	BaseResponse `json:",inline"`                    // Common response fields
	Response     GetBrShopOnboardingInfoResponseData `json:"response"` // Actual response data
}

type GetBrShopOnboardingInfoResponseData struct {
	TaxIdType         int64           `json:"tax_id_type"`        // [Required] <p>Type of the shop’s tax ID. Applicable values:&nbsp;</p><p> </p><p>1: Personal seller (CPF)<br /> 2: Company seller (CNPJ)</p>
	TaxId             string          `json:"tax_id"`             // [Required] <p>The shop’s tax ID.<br />- When tax_id_type = 1 (Personal seller), it is CPF.</p><p>- When tax_id_type = 2 (Company seller), it is CNPJ.</p>
	CpfId             string          `json:"cpf_id"`             // [Required] <p>CPF number of the individual seller. Valid only when tax_id_type = 1.</p>
	CnpjId            string          `json:"cnpj_id"`            // [Required] <p>CNPJ number of the company seller. Valid only when tax_id_type = 2.</p>
	Name              string          `json:"name"`               // [Required] <p>Full name of the individual seller. Valid&nbsp;only when&nbsp;tax_id_type = 1.</p>
	LegalEntityName   string          `json:"legal_entity_name"`  // [Required] <p>Legal name of the company seller. Valid&nbsp;only when&nbsp;tax_id_type = 2.</p>
	Birthday          int64           `json:"birthday"`           // [Required] <p>Birthday of the individual seller (stored as Unix timestamp).&nbsp;Valid only when tax_id_type = 1.</p>
	BirthdayStr       string          `json:"birthday_str"`       // [Required] <p>Birthday of the individual seller (formatted as&nbsp;YYYY-MM-DD). Valid only when tax_id_type = 1.</p>
	StateRegistration string          `json:"state_registration"` // [Required] <p>State registration number of the shop.</p>
	BillingAddress    *BillingAddress `json:"billing_address"`    // [Required] <p>Shop’s billing address details.</p>
	OnboardingStatus  int64           `json:"onboarding_status"`  // [Required] <p>Status of the shop’s current KYC onboarding process. Applicable values:</p><p>0: None<br />1: Regis Processing</p><p>2: Regis Validated</p><p>3: Regis Rejected</p><p>4: KYC Pending<br />5: KYC Processing<br />6: KYC Processing Manually<br />7: KYC Validated<br />8: KYC Rejected</p>
	SubmissionTime    int64           `json:"submission_time"`    // [Required] <p>Timestamp when the onboarding information was submitted.</p>
	Nationality       string          `json:"nationality"`        // [Required] <p>Nationality of the individual seller. Valid only when tax_id_type=1.</p>
	CnaeMain          string          `json:"cnae_main"`          // [Required] <p>Main CNAE code.</p>
	CnaeSecondary     string          `json:"cnae_secondary"`     // [Required] <p>Secondary CNAE code.</p>
	MeiCheck          string          `json:"mei_check"`          // [Required] <p>MEI verification result. Applicable values:</p><p>0: No<br />1: Yes</p>
	OnboardingPassed  bool            `json:"onboarding_passed"`  // [Required] <p>Indicate if the shop has passed KYC verification.</p>
}

type GetProfileResponse struct {
	BaseResponse `json:",inline"`       // Common response fields
	Response     GetProfileResponseData `json:"response"` // Actual response data
}

type GetProfileResponseData struct {
	ShopLogo      string `json:"shop_logo"`      // [Required] The Image URL of the shop logo.
	Description   string `json:"description"`    // [Required] The content of the shop description.
	ShopName      string `json:"shop_name"`      // [Required] The content of the shop name.
	InvoiceIssuer string `json:"invoice_issuer"` // [Required] <p>The invoice issuer information for the shop. It could be "Shopee" or "Other" as the invoice issuer. This is for BR CNPJ seller only.<br /></p>
}

type GetShopHolidayModeResponse struct {
	BaseResponse `json:",inline"`               // Common response fields
	Response     GetShopHolidayModeResponseData `json:"response"` // Actual response data
}

type GetShopHolidayModeResponseData struct {
	HolidayModeOn    bool   `json:"holiday_mode_on"`    // [Required] <p>Indicate whether the shop has enabled holiday mode. true means ON, false means OFF.</p>
	HolidayModeMtime int64  `json:"holiday_mode_mtime"` // [Required] <p>The last time the holiday mode was modifies.</p>
	DebugMsg         string `json:"debug_msg"`          // [Required] <p>Debug message.</p>
}

type GetShopInfoResponse struct {
	BaseResponse         `json:",inline"`   // Common response fields
	ShopName             string             `json:"shop_name,omitempty"`               // Name of the shop.
	Region               string             `json:"region,omitempty"`                  // Shop's area.
	Status               string             `json:"status,omitempty"`                  // <p>Applicable status: <b>BANNED</b>, <b>FROZEN</b>, <b>NORMAL</b>.</p>
	SipAffiShops         []SipAffiShops     `json:"sip_affi_shops,omitempty"`          // SIP affiliate shops info list.If you request for SIP primary shop,this field will be returned, if you request for SIP affiliate shop,this field won't be returned
	IsCb                 bool               `json:"is_cb,omitempty"`                   // Use this filed to indicate whether the shop is a cross-border shop.
	AuthTime             int64              `json:"auth_time,omitempty"`               // The timestamp when the shop was authorized to the partner.
	ExpireTime           int64              `json:"expire_time,omitempty"`             // Use this field to indicate the expiration date for shop authorization.
	IsSip                bool               `json:"is_sip,omitempty"`                  // This filed will return "true" when SIP primary shop or affiliate shop calls
	IsUpgradedCbsc       bool               `json:"is_upgraded_cbsc,omitempty"`        // <p>Use this filed to indicate whether this merchant is upgraded to CBSC, including CNSC and KRSC.<br /></p>
	MerchantId           int64              `json:"merchant_id,omitempty"`             // <p>Shopee’s unique identifier for a merchant. If the shop won't under any merchant, then the value will be null.<br /></p>
	ShopFulfillmentFlag  string             `json:"shop_fulfillment_flag,omitempty"`   // <p>Use this field to indicate the fulfillment type of current shop, the applicable values:&nbsp;</p><p><b><br /></b></p><p><b>- Pure - FBS Shop:&nbsp;</b>Single mode, refer to Local/CB shops which only have Shopee official warehouse stock, orders are fulfilled by Shopee from Shopee official warehouse;&nbsp;</p><p><b><br /></b></p><p><b>- Pure - 3PF Shop:&nbsp;</b>Single mode, refer to CB shops which only have local seller warehouse stock, orders are fulfilled by seller from local seller warehouse via local logistics channels;&nbsp;</p><p><b><br /></b></p><p><b>- PFF - FBS Shop:&nbsp;</b></p><p>1) Hybird mode, refer to Local shops which have both Shopee official warehouse stock and local seller warehouse stock, orders can be fulfilled by Shopee from Shopee official warehouse and can also fulfilled by seller from local seller warehouse via local logistics channels;&nbsp;</p><p>2) Hybrid mode, refer to CB shops which have both Shopee official warehouse stock and CB seller warehouse stock, orders can be fulfilled by Shopee from Shopee official warehouse and can also fulfilled by seller from CB seller warehouse via CB logistics channels;&nbsp;</p><p><b><br /></b></p><p><b>-&nbsp;PFF - 3PF Shop:&nbsp;</b>Hybrid mode, refer to CB shops which have both local seller warehouse stock and CB seller warehouse stock, orders can be fulfilled by seller from local seller warehouse via local logistics channels and can also fulfilled by seller from CB seller warehouse via CB logistics channels;&nbsp;</p><p><br /></p><p><b>- LFF Hybrid Shop:</b> Hybrid mode, refer to CB shops which have 3 types of stock: FBS stock (Shopee official warehouse stock), 3PF stock (CB seller own stock in the local market) and CB SLS stock (CB seller own stock in CN/HK/KR);</p><p><b><br /></b></p><p><b>- Others</b></p><p><b><br /></b></p><p><b>- Unknown: </b>Returned when obtaining shop_fulfillment_flag information fails</p>
	IsMainShop           bool               `json:"is_main_shop,omitempty"`            // <p>Identifies if the current shop is a Local Shop linked to Cross Border Direct Shop.</p>
	IsDirectShop         bool               `json:"is_direct_shop,omitempty"`          // <p>Identifies if the current shop is a Cross Border Direct Shop.</p>
	LinkedMainShopId     int64              `json:"linked_main_shop_id,omitempty"`     // <p>Returns the Shop ID of the Local Shop linked to the Cross Border Direct Shop.</p>
	LinkedDirectShopList []LinkedDirectShop `json:"linked_direct_shop_list,omitempty"` // <p>Returns the list of Cross Border Direct Shops linked to the Local Shop.</p>
	IsOneAwb             bool               `json:"is_one_awb,omitempty"`              // <p>Use this filed to indicate if the shop is in 1-AWB whitelist.&nbsp;</p><p><br /></p><p>If is_one_awb return true, please use&nbsp;new AWB size (10cm x 15cm thermal paper) to print AWB. For more details, please refer to:&nbsp;<a href="https://open.shopee.com/announcements/1138?category=3&amp;is_top=false" target="_blank" style="font-size:14px;">https://open.shopee.com/announcements/1138?category=3&amp;is_top=false</a></p>
	IsMartShop           bool               `json:"is_mart_shop,omitempty"`            // <p>Indicates whether the current shop is a Mart Shop.<br /></p>
	IsOutletShop         bool               `json:"is_outlet_shop,omitempty"`          // <p>Indicates whether the current shop is an Outlet Shop.<br /></p>
	MartShopId           int64              `json:"mart_shop_id,omitempty"`            // <p>(Only returned when requesting an Outlet Shop) Refers to the Mart Shop ID this Outlet belongs to.<br /></p>
	OutletShopInfoList   []OutletShopInfo   `json:"outlet_shop_info_list,omitempty"`   // <p>(Only returned when requesting a Mart Shop) List of Outlet Shop IDs under this Mart Shop.<br /></p>
}

type GetShopNotificationRequest struct {
	Cursor   *int64 `json:"cursor,omitempty" url:"cursor,omitempty"`       // [Optional] <p>The last notification_id returned on the page. When using the cursor, notifications will start with the one following this cursor notification. If no cursor is provided, the latest message from the shop will be returned.<br /></p>
	PageSize *int64 `json:"page_size,omitempty" url:"page_size,omitempty"` // [Optional] <p>Default 10; maximum 50</p>
}

type GetShopNotificationResponse struct {
	BaseResponse `json:",inline"` // Common response fields
	Cursor       int64            `json:"cursor,omitempty"` // <p>Last notification_id returned in the page</p>
	Data         *ResponseData    `json:"data,omitempty"`   //
}

type GetWarehouseDetailRequest struct {
	WarehouseType *int64 `json:"warehouse_type,omitempty" url:"warehouse_type,omitempty"` // [Optional] <p>Type of warehouse. Applicable values:&nbsp;</p><p>- 1: Pickup Warehouse</p><p>- 2: Return Warehouse</p><p><br /></p><p>Default value is 1 (Pickup Warehouse)</p>
}

type GetWarehouseDetailResponse struct {
	BaseResponse `json:",inline"`               // Common response fields
	Response     GetWarehouseDetailResponseData `json:"response"` // Actual response data
}

type GetWarehouseDetailResponseData struct {
	WarehouseId      int64  `json:"warehouse_id"`       // [Required] <p>Warehouse address identifier. It should be unique for every warehouse address<br /></p>
	WarehouseName    string `json:"warehouse_name"`     // [Required] <p>The warehouse name filled in when creating the warehouse address</p>
	WarehouseType    int64  `json:"warehouse_type"`     // [Required] <p>Type of warehouse. Applicable values:&nbsp;</p><p>- 1: Pickup Warehouse</p><p>- 2: Return Warehouse</p>
	LocationId       string `json:"location_id"`        // [Required] <p>Location identifier for stocks. Different location_ids represent that your addresses are in different item stocks</p>
	AddressId        int64  `json:"address_id"`         // [Required] <p>Identity of address<br /></p>
	Region           string `json:"region"`             // [Required] <p>Region of your warehouse address</p>
	State            string `json:"state"`              // [Required] <p>State of your warehouse address<br /></p>
	City             string `json:"city"`               // [Required] <p>City of your warehouse address<br /></p>
	Address          string `json:"address"`            // [Required] <p>Detail address of your warehouse address</p>
	Zipcode          string `json:"zipcode"`            // [Required] <p>Zipcode of your warehouse address<br /></p>
	District         string `json:"district"`           // [Required] <p>Distinct of your warehouse address<br /></p>
	Town             string `json:"town"`               // [Required] <p>Town of your warehouse address<br /></p>
	StateCode        string `json:"state_code"`         // [Required] <p>State code of your warehouse address<br /></p>
	HolidayModeState int64  `json:"holiday_mode_state"` // [Required] <p>The holiday mode state of your address.<br />0: not in holiday mode</p><p>1: holiday mode active</p><p>2: holiday mode is turning of</p><p>3: holiday mode is turning on</p>
}

type LinkedDirectShop struct {
	DirectShopId     int64  `json:"direct_shop_id"`     // [Required] <p>Shop ID of the Cross Border Direct Shop.</p>
	DirectShopRegion string `json:"direct_shop_region"` // [Required] <p>Shop Region of the Cross Border Direct Shop.</p>
}

type OutletShopInfo struct {
	OutletShopId int64 `json:"outlet_shop_id"` // [Required] <p>Shop ID of the Outlet Shop.</p>
}

type ResponseData struct {
	CreateTime int64  `json:"create_time"` // [Required] <p>the notification create time</p>
	Content    string `json:"content"`     // [Required] <p>The content of the notification</p>
	Title      string `json:"title"`       // [Required] <p>The content of the notification<br /></p>
	Url        string `json:"url"`         // [Required] <p>Some notification may be attached with URL, it will redirect to seller center</p>
}

type SetShopHolidayModeRequest struct {
	HolidayModeOn bool `json:"holiday_mode_on"` // [Required] <p>Indicate whether to enable holiday mode for the shop. true means turn ON, false means turn OFF.</p>
}

type SetShopHolidayModeResponse struct {
	BaseResponse `json:",inline"`               // Common response fields
	Response     SetShopHolidayModeResponseData `json:"response"` // Actual response data
}

type SetShopHolidayModeResponseData struct {
	DebugMsg string `json:"debug_msg"` // [Required] <p>Debug message.</p>
}

type UpdateProfileRequest struct {
	ShopName    *string `json:"shop_name,omitempty"`   // [Optional] The new shop name
	ShopLogo    *string `json:"shop_logo,omitempty"`   // [Optional] The new shop logo url. Recommend to use images
	Description *string `json:"description,omitempty"` // [Optional] The new shop description.
}

type UpdateProfileResponse struct {
	BaseResponse `json:",inline"`          // Common response fields
	Response     UpdateProfileResponseData `json:"response"` // Actual response data
}

type UpdateProfileResponseData struct {
	ShopLogo    string `json:"shop_logo"`   // [Required] The Image URL of the shop logo after updated.
	Description string `json:"description"` // [Required] The content of the shop description after updated.
	ShopName    string `json:"shop_name"`   // [Required] The content of the shop name after updated.
}
