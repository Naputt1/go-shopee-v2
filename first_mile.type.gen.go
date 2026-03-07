package goshopee

type BindCourierDeliveryFirstMileTrackingNumberRequest struct {
	ShipmentMethod string        `json:"shipment_method"` // [Required] <p>The shipment method for generate and bind orders. Available value:&nbsp;courier_delivery.<br /></p>
	BindingId      string        `json:"binding_id"`      // [Required] <p>If using courier_delivery as the shipment method, the "binding_id" field should pass the value generated from&nbsp;v2.first_mile.generate_and_bind_first_mile_tracking_number.<br /></p>
	OrderList      []SharedOrder `json:"order_list"`      // [Required] <p>The list of order_sn. You can specify up to 50 order_sns in this call.&nbsp;One fm_tn maximum number of total bind orders is 10000.<br /></p>
}

type BindCourierDeliveryFirstMileTrackingNumberResponse struct {
	BaseResponse `json:",inline"`                                       // Common response fields
	Response     BindCourierDeliveryFirstMileTrackingNumberResponseData `json:"response"` // Actual response data
}

type BindCourierDeliveryFirstMileTrackingNumberResponseData struct {
	BindingId   string                                     `json:"binding_id"`   // [Required] <p>Binding ID<br /></p>
	SuccessList []SharedOrder                              `json:"success_list"` // [Required]
	FailList    []CreateShippingDocumentResponseDataResult `json:"fail_list"`    // [Required]
}

type BindFirstMileTrackingNumberRequest struct {
	FirstMileTrackingNumber string        `json:"first_mile_tracking_number"` // [Required] <p>If using "pickup" or "self_deliver" as the shipment method the "first_mile_tracking_number" field should pass the value generated from v2.first_mile.generate_first_mile_tracking_number.</p><p><br /></p><p>If using "dropoff" as the shipment method the "first_mile_tracking_number" field should pass the tracking number provide by the supplier.</p>
	ShipmentMethod          string        `json:"shipment_method"`            // [Required] <p>The shipment method for bound orders, should be pickup, dropoff or self_deliver.</p>
	Region                  string        `json:"region"`                     // [Required] <p>Use this field to specify the region you want to ship parcel.</p><p>Available value: cn,kr.&nbsp;</p><p>Please fill in the field according to the region of the Merchant to which the shop belongs.</p><p></p>
	LogisticsChannelId      int64         `json:"logistics_channel_id"`       // [Required] <p>The identity of first-mile logistic channel.</p><p>If you using "dropoff" or "pickup" as shipment method, please call&nbsp;v2.first_mile.get_channel_list to get the logsitc_channel_id then fill it.</p><p><br /></p><p>If you using "self_deliver"as shipment method then the logistics_channel_id should be "null".<br /></p>
	Volume                  *float64      `json:"volume,omitempty"`           // [Optional] The volume of the parcel.
	Weight                  *float64      `json:"weight,omitempty"`           // [Optional] The weight of the parcel.
	Width                   *float64      `json:"width,omitempty"`            // [Optional] The width of the parcel.
	Length                  *float64      `json:"length,omitempty"`           // [Optional] The length of the parcel.
	Height                  *float64      `json:"height,omitempty"`           // [Optional] The height of the parcel.
	OrderList               []SharedOrder `json:"order_list"`                 // [Required] The set of ordersn. You can specify up to 50 ordersns in this call.one fm_tn maximum number of total bind orders is 10000.
}

type BindFirstMileTrackingNumberResponse struct {
	BaseResponse `json:",inline"`                        // Common response fields
	Response     BindFirstMileTrackingNumberResponseData `json:"response"` // Actual response data
}

type BindFirstMileTrackingNumberResponseData struct {
	FirstMileTrackingNumber string                                     `json:"first_mile_tracking_number"` // [Required] The first mile tracking number
	OrderList               []CreateShippingDocumentResponseDataResult `json:"order_list"`                 // [Required] The list of orders.
}

type Courier struct {
	CourierName        string `json:"courier_name"`         // [Required] <p>The name of the courier.<br /></p>
	CourierServiceId   string `json:"courier_service_id"`   // [Required] <p>The identity of the service provided by courier.<br /></p>
	CourierServiceName string `json:"courier_service_name"` // [Required] <p>The name of the service provided by courier.<br /></p>
}

type CourierDeliveryInfo struct {
	AddressId          int64  `json:"address_id"`                   // [Required] <p>The identity of address. Retrieved from v2.logistics.get_address_list, address_type need to be FIRST_MILE_PICKUP_ADDRESS.</p>
	WarehouseId        string `json:"warehouse_id"`                 // [Required] <p>The identity of transit warehouse address. Retrieved from&nbsp;v2.first_mile.get_transit_warehouse_list.<br /></p>
	LogisticsProductId int64  `json:"logistics_product_id"`         // [Required] <p>The definition of logistics product ID:&nbsp;</p><p>1010003 (kuaidi100 to C) - seller book courier pickup and pay offline</p><p>1010004 (kuaidi100 prepaid(MP)) -&nbsp;seller can use prepaid account to place courier order, so prepaid_account_id is required</p>
	PrepaidAccountId   *int64 `json:"prepaid_account_id,omitempty"` // [Optional] <p>ID of prepaid account. Retrieved from&nbsp;v2.merchant.get_merchant_prepaid_account_list.</p>
	CourierServiceId   string `json:"courier_service_id"`           // [Required] <p>The identity of courier service. Retrieved from v2.first_mile.get_courier_delivery_channel_list.</p>
}

type CreateShippingDocumentResponseDataResult struct {
	OrderSn       string `json:"order_sn"`       // [Required] <p>Shopee's unique identifier for an order.</p>
	PackageNumber string `json:"package_number"` // [Required] <p>Shopee's unique identifier for the package under an order.</p>
	FailError     string `json:"fail_error"`     // [Required] <p>Indicate error type if one element hit error.</p>
	FailMessage   string `json:"fail_message"`   // [Required] <p>Indicate error details if one element hit error.</p>
}

type FirstMileGetChannelListResponse struct {
	BaseResponse `json:",inline"`                    // Common response fields
	Response     FirstMileGetChannelListResponseData `json:"response"` // Actual response data
}

type FirstMileGetChannelListResponseData struct {
	LogisticsChannelList []ResponseDataLogisticsChannel `json:"logistics_channel_list"` // [Required]
}

type FirstMileTrackingNumber struct {
	FirstMileTrackingNumber string `json:"first_mile_tracking_number"` // [Required] <p>The first-mile tracking number.</p>
	Status                  string `json:"status"`                     // [Required] <p>The logistics status for bound orders.</p><p><br /></p><p>NOT_AVAILABLE status means that First Mile Tracking Number is not yet bound with any order.</p>
	DeclareDate             string `json:"declare_date"`               // [Required] <p>The specified delivery date.</p>
}

type GenerateAndBindFirstMileTrackingNumberRequest struct {
	ShipmentMethod      string               `json:"shipment_method"`       // [Required] <p>The shipment method for generate and bind orders. Available value:&nbsp;courier_delivery.</p>
	Region              *string              `json:"region,omitempty"`      // [Optional] <p>Use this field to specify the region you want to ship parcel. Available value:&nbsp;CN.</p>
	OrderList           []SharedOrder        `json:"order_list"`            // [Required] <p>The list of order_sn. You can specify up to 50 order_sns in this call. One fm_tn maximum number of total bind orders is 10000.</p>
	CourierDeliveryInfo *CourierDeliveryInfo `json:"courier_delivery_info"` // [Required] <p>The set of information you need to generate FM tracking number and bind orders.</p>
}

type GenerateAndBindFirstMileTrackingNumberResponse struct {
	BaseResponse `json:",inline"`                                   // Common response fields
	Response     GenerateAndBindFirstMileTrackingNumberResponseData `json:"response"` // Actual response data
}

type GenerateAndBindFirstMileTrackingNumberResponseData struct {
	BindingId   string                                     `json:"binding_id"`   // [Required] <p>Binding ID</p>
	SuccessList []SharedOrder                              `json:"success_list"` // [Required]
	FailList    []CreateShippingDocumentResponseDataResult `json:"fail_list"`    // [Required]
}

type GenerateFirstMileTrackingNumberRequest struct {
	DeclareDate string `json:"declare_date"`       // [Required] This field is used for seller to specify the declare time.
	Quantity    *int64 `json:"quantity,omitempty"` // [Optional] The number of first-mile tracking numbers generated. Up to 20 first-mile tracking numbers can be generated for one declaration day.
}

type GenerateFirstMileTrackingNumberResponse struct {
	BaseResponse `json:",inline"`                            // Common response fields
	Response     GenerateFirstMileTrackingNumberResponseData `json:"response"` // Actual response data
}

type GenerateFirstMileTrackingNumberResponseData struct {
	FirstMileTrackingNumberList []string `json:"first_mile_tracking_number_list"` // [Required] The list of first mile tracking number that you generate
}

type GetChannelListRequest struct {
	Region *string `json:"region,omitempty"` // [Optional] <p>Use this field to specify the region you want to ship parcel. Available value: CN, KR</p>
}

type GetCourierDeliveryChannelListRequest struct {
	Region *string `json:"region,omitempty" url:"region,omitempty"` // [Optional] <p>Use this field to specify the region you want to ship parcel. Available value: CN</p>
}

type GetCourierDeliveryChannelListResponse struct {
	BaseResponse `json:",inline"`                          // Common response fields
	Response     GetCourierDeliveryChannelListResponseData `json:"response"` // Actual response data
}

type GetCourierDeliveryChannelListResponseData struct {
	LogisticsChannelList []GetCourierDeliveryChannelListResponseDataLogisticsChannel `json:"logistics_channel_list"` // [Required]
}

type GetCourierDeliveryChannelListResponseDataLogisticsChannel struct {
	LogisticsProductId   int64     `json:"logistics_product_id"`   // [Required] <p>The identity of logistics product ID:&nbsp;</p><p>1010003: kuaidi100 to C</p><p>1010004:&nbsp;kuaidi100 prepaid(MP)</p>
	LogisticsProductName string    `json:"logistics_product_name"` // [Required] <p>The name of logistics product ID:&nbsp;</p><p>- kuaidi100 to C</p><p>- kuaidi100 prepaid(MP)</p>
	CourierList          []Courier `json:"courier_list"`           // [Required]
}

type GetCourierDeliveryDetailRequest struct {
	BindingId string  `json:"binding_id" url:"binding_id"`                   // [Required] <p>Binding ID</p>
	Cursor    *string `json:"cursor,omitempty" url:"cursor,omitempty"`       // [Optional] <p>Specifies the starting entry of data to return in the current call. Default is "". If data is more than one page, the offset can be some entry to start next call.</p>
	PageSize  *int64  `json:"page_size,omitempty" url:"page_size,omitempty"` // [Optional] <p>Each result set is returned as a page of entries. Use the "page_size" filters to control the maximum number of entries to retrieve per page (i.e., per call). This integer value is used to specify the maximum number of entries to return in a single "page" of data. limit [1, 50].<br /></p>
}

type GetCourierDeliveryDetailResponse struct {
	BaseResponse `json:",inline"`                     // Common response fields
	Response     GetCourierDeliveryDetailResponseData `json:"response"` // Actual response data
}

type GetCourierDeliveryDetailResponseData struct {
	BindingId               string                       `json:"binding_id"`                 // [Required] <p>Binding ID<br /></p>
	FirstMileTrackingNumber string                       `json:"first_mile_tracking_number"` // [Required] <p>The first mile tracking number.<br /></p>
	Status                  string                       `json:"status"`                     // [Required] <p>The logistics status for first-mile tracking number. Status could be:<br />CANCELED<br />CANCELING<br />DELIVERED<br />NOT_AVAILABLE<br />ORDER_CREATED<br />ORDER_RECEIVED<br />PICKED_UP</p>
	DeclareDate             string                       `json:"declare_date"`               // [Required] <p>The specified delivery date.</p>
	More                    bool                         `json:"more"`                       // [Required] <p>This is to indicate whether the item list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of items.<br /></p>
	NextCursor              string                       `json:"next_cursor"`                // [Required] <p>If more is true, you should pass the next_cursor in the next request as cursor. The value of next_cursor will be empty string when more is false.<br /></p>
	OrderList               []GetDetailResponseDataOrder `json:"order_list"`                 // [Required]
}

type GetCourierDeliveryTrackingNumberListRequest struct {
	FromDate string  `json:"from_date"`           // [Required] <p>The start time of creation time</p>
	ToDate   string  `json:"to_date"`             // [Required] <p>The end time of creation time</p>
	PageSize *int64  `json:"page_size,omitempty"` // [Optional] <p>Each result set is returned as a page of entries. Use the "page_size" filters to control the maximum number of entries to retrieve per page (i.e., per call). This integer value is used to specify the maximum number of entries to return in a single "page" of data. limit [1, 50]<br /></p>
	Cursor   *string `json:"cursor,omitempty"`    // [Optional] <p>Specifies the starting entry of data to return in the current call. Default is "". If data is more than one page, the offset can be some entry to start next call.<br /></p>
}

type GetCourierDeliveryTrackingNumberListResponse struct {
	BaseResponse `json:",inline"`                                 // Common response fields
	Response     GetCourierDeliveryTrackingNumberListResponseData `json:"response"` // Actual response data
}

type GetCourierDeliveryTrackingNumberListResponseData struct {
	TrackingNumberList []ResponseDataTrackingNumber `json:"tracking_number_list"` // [Required]
	More               bool                         `json:"more"`                 // [Required] <p>This is to indicate whether the tracking number list is more than one page. If this value is true, you may want to continue to check next page to retrieve tracking numbers.<br /></p>
	NextCursor         string                       `json:"next_cursor"`          // [Required] <p>If more is true, you should pass the next_cursor in the next request as cursor. The value of next_cursor will be empty string when more is false.<br /></p>
}

type GetCourierDeliveryWaybillRequest struct {
	BindingIdList []string `json:"binding_id_list"` // [Required] <p>Binding ID list of waybill. System limits maximum of Binding ID to 50.&nbsp;&nbsp;</p>
}

type GetCourierDeliveryWaybillResponse struct {
	BaseResponse `json:",inline"`                      // Common response fields
	Response     GetCourierDeliveryWaybillResponseData `json:"response"` // Actual response data
}

type GetCourierDeliveryWaybillResponseData struct {
	WaybillList []Waybill `json:"waybill_list"` // [Required]
}

type GetDetailRequest struct {
	FirstMileTrackingNumber string  `json:"first_mile_tracking_number" url:"first_mile_tracking_number"` // [Required] The first mile tracking number.
	Cursor                  *string `json:"cursor,omitempty" url:"cursor,omitempty"`                     // [Optional] Specifies the starting entry of data to return in the current call. Default is "". If data is more than one page, the offset can be some entry to start next call.
}

type GetDetailResponse struct {
	BaseResponse `json:",inline"`      // Common response fields
	Response     GetDetailResponseData `json:"response"` // Actual response data
}

type GetDetailResponseData struct {
	LogisticsChannelId      int64                        `json:"logistics_channel_id"`       // [Required] The identity of logistic channel.
	FirstMileTrackingNumber string                       `json:"first_mile_tracking_number"` // [Required] The first-mile tracking number.
	ShipmentMethod          string                       `json:"shipment_method"`            // [Required] The shipment method for bound orders, should be pickup or dropoff.
	Status                  string                       `json:"status"`                     // [Required] <p>The logistics status for first-mile tracking number. Status could be: NOT_AVAILABLE,ORDER_CREATED,PICKED_UP,DELIVERED,ORDER_RECEIVED,CANCELING,CANCELED.</p><p><br /></p><p>NOT_AVAILABLE status means that either of the two scenarios has happened:</p><p>1. First Mile Tracking Number in the request does not exist. (e.g., wrong format)</p><p>2. First Mile Tracking Number is not yet bound with any order.</p>
	DeclareDate             string                       `json:"declare_date"`               // [Required] The specified delivery date.
	More                    bool                         `json:"more"`                       // [Required] This is to indicate whether the item list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of items.
	OrderList               []GetDetailResponseDataOrder `json:"order_list"`                 // [Required] The list of order.
	NextCursor              string                       `json:"next_cursor"`                // [Required] If more is true, you should pass the next_cursor in the next request as cursor. The value of next_cursor will be empty string when more is false.
}

type GetDetailResponseDataOrder struct {
	OrderSn                 string `json:"order_sn"`                  // [Required] <p>Shopee's unique identifier for an order.<br /></p>
	PackageNumber           string `json:"package_number"`            // [Required] <p>Shopee's unique identifier for the package under an order.<br /></p>
	SlsTrackingNumber       string `json:"sls_tracking_number"`       // [Required] <p>The tracking number of SLS for orders/forders.<br /></p>
	PickUpDone              bool   `json:"pick_up_done"`              // [Required] <p>Use this filed to indicate whether the order has been picked up by carrier.<br /></p>
	ArrivedTransitWarehouse bool   `json:"arrived_transit_warehouse"` // [Required] <p>Use this filed to indicate whether the order has arrived at transit warehouse.<br /></p>
}

type GetTrackingNumberListRequest struct {
	FromDate string  `json:"from_date" url:"from_date"`                     // [Required] The start time of declare_date.
	ToDate   string  `json:"to_date" url:"to_date"`                         // [Required] The end time of declare_date.
	PageSize *int64  `json:"page_size,omitempty" url:"page_size,omitempty"` // [Optional] Each result set is returned as a page of entries. Use the "page_size" filters to control the maximum number of entries to retrieve per page (i.e., per call). This integer value is used to specify the maximum number of entries to return in a single "page" of data. limit [1, 50]
	Cursor   *string `json:"cursor,omitempty" url:"cursor,omitempty"`       // [Optional] Specifies the starting entry of data to return in the current call. Default is "". If data is more than one page, the offset can be some entry to start next call.
}

type GetTrackingNumberListResponse struct {
	BaseResponse `json:",inline"`                  // Common response fields
	Response     GetTrackingNumberListResponseData `json:"response"` // Actual response data
}

type GetTrackingNumberListResponseData struct {
	More                        bool                      `json:"more"`                            // [Required] This is to indicate whether the order list is more than one page. If this value is true, you may want to continue to check next page to retrieve orders.
	FirstMileTrackingNumberList []FirstMileTrackingNumber `json:"first_mile_tracking_number_list"` // [Required] The first-mile tracking number.
	NextCursor                  string                    `json:"next_cursor"`                     // [Required] If more is true, you should pass the next_cursor in the next request as cursor. The value of next_cursor will be empty string when more is false.
}

type GetTransitWarehouseListRequest struct {
	Region *string `json:"region,omitempty" url:"region,omitempty"` // [Optional] <p>Use this field to specify the region you want to ship parcel. Available value:&nbsp;CN.</p>
}

type GetTransitWarehouseListResponse struct {
	BaseResponse `json:",inline"`                    // Common response fields
	Response     GetTransitWarehouseListResponseData `json:"response"` // Actual response data
}

type GetTransitWarehouseListResponseData struct {
	TransitWarehouseList []TransitWarehouse `json:"transit_warehouse_list"` // [Required]
}

type GetUnbindOrderListRequest struct {
	Cursor                 *string `json:"cursor,omitempty" url:"cursor,omitempty"`                                     // [Optional] Specifies the starting entry of data to return in the current call. Default is "". If data is more than one page, the offset can be some entry to start next call.
	PageSize               *int64  `json:"page_size,omitempty" url:"page_size,omitempty"`                               // [Optional] Each result set is returned as a page of entries. Use the "page_size" filters to control the maximum number of entries to retrieve per page (i.e., per call). This integer value is used to specify the maximum number of entries to return in a single "page" of data. limit [1, 100]
	ResponseOptionalFields *string `json:"response_optional_fields,omitempty" url:"response_optional_fields,omitempty"` // [Optional] <p>Indicate response fields you want to get. Please select from the below response parameters. If you input an object field, all the params under it will be included automatically in the response. If there are multiple response fields you want to get, you need to use English comma to connect them.  Available values: logistics_status,package_number.</p>
}

type GetUnbindOrderListResponse struct {
	BaseResponse `json:",inline"`               // Common response fields
	Response     GetUnbindOrderListResponseData `json:"response"` // Actual response data
}

type GetUnbindOrderListResponseData struct {
	More       bool                                  `json:"more"`        // [Required] This is to indicate whether the item list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of items.
	OrderList  []GetUnbindOrderListResponseDataOrder `json:"order_list"`  // [Required] The result list of order you querying.
	NextCursor string                                `json:"next_cursor"` // [Required] If more is true, you should pass the next_cursor in the next request as cursor. The value of next_cursor will be empty string when more is false.
}

type GetUnbindOrderListResponseDataOrder struct {
	OrderSn         string          `json:"order_sn"`         // [Required] Shopee's unique identifier for an order.
	PackageNumber   string          `json:"package_number"`   // [Required] Shopee's unique identifier for the package under an order.
	LogisticsStatus LogisticsStatus `json:"logistics_status"` // [Required] The Shopee logistics status for the order. Applicable values: See Data Definition- LogisticsStatus.
}

type GetWaybillRequest struct {
	FirstMileTrackingNumberList []string `json:"first_mile_tracking_number_list"` // [Required] The first mile tracking number that you want to print waybill.limit [1, 50]
}

type GetWaybillResponse struct {
	BaseResponse `json:",inline"` // Common response fields
	Waybill      interface{}      `json:"waybill,omitempty"` // <p>The waybill file.</p>
}

type ResponseDataLogisticsChannel struct {
	LogisticsChannelId   int64  `json:"logistics_channel_id"`   // [Required] The identity of logistic channel.
	LogisticsChannelName string `json:"logistics_channel_name"` // [Required] The name of logistic channel.
	ShipmentMethod       string `json:"shipment_method"`        // [Required] <p>The shipment method for bound orders.Available values: pickup, dropoff, self_deliver.</p>
}

type ResponseDataTrackingNumber struct {
	BindingId               string `json:"binding_id"`                 // [Required] <p>The generated binding ID.</p>
	FirstMileTrackingNumber string `json:"first_mile_tracking_number"` // [Required] <p>The generate first-mile tracking number, value might be blank.</p>
	Status                  string `json:"status"`                     // [Required] <p>The logistics status for first-mile tracking number. Status could be:</p><p>CANCELED<br />CANCELING<br />DELIVERED<br />NOT_AVAILABLE<br />ORDER_CREATED<br />ORDER_RECEIVED<br />PICKED_UP</p><p><br /></p><p>Note:&nbsp;NOT_AVAILABLE status means that Binding ID / First Mile Tracking Number is not yet bound with any order.</p>
	Reason                  string `json:"reason"`                     // [Required] <p>Indicate the reason when Shopee failed to place courier order to 3PL (Kuaidi 100 supporting) or courier company cancelled the order.</p><p><br /></p><p>Note: Will be empty if status is not CANCELED.</p>
	DeclareDate             string `json:"declare_date"`               // [Required] <p>The declare date of binding ID/first-mile tracking number.</p>
}

type SharedOrder struct {
	OrderSn       string `json:"order_sn"`       // [Required] <p>Shopee's unique identifier for an order.<br /></p>
	PackageNumber string `json:"package_number"` // [Required] <p>Shopee's unique identifier for the package under an order.<br /></p>
}

type TransitWarehouse struct {
	WarehouseId     string `json:"warehouse_id"`      // [Required] <p>The identity of transit warehouse.<br /></p>
	WarehouseNameEn string `json:"warehouse_name_en"` // [Required] <p>The name of transit warehouse in English.<br /></p>
	WarehouseNameCn string `json:"warehouse_name_cn"` // [Required] <p>The name of transit warehouse in Chinese.<br /></p>
}

type UnbindFirstMileTrackingNumberAllRequest struct {
	OrderList []SharedOrder `json:"order_list"` // [Required] <p>The list of order info you want to unbind from the first mile tracking number or binding ID.&nbsp;You can specify up to 50 order_sns in this call.</p>
}

type UnbindFirstMileTrackingNumberAllResponse struct {
	BaseResponse `json:",inline"`                             // Common response fields
	Response     UnbindFirstMileTrackingNumberAllResponseData `json:"response"` // Actual response data
}

type UnbindFirstMileTrackingNumberAllResponseData struct {
	SuccessList []UnbindFirstMileTrackingNumberAllResponseDataSuccess `json:"success_list"` // [Required]
	FailList    []CreateShippingDocumentResponseDataResult            `json:"fail_list"`    // [Required]
}

type UnbindFirstMileTrackingNumberAllResponseDataSuccess struct {
	OrderSn                 string `json:"order_sn"`                   // [Required] <p>Shopee's unique identifier for an order.<br /></p>
	PackageNumber           string `json:"package_number"`             // [Required] <p>Shopee's unique identifier for the package under an order.<br /></p>
	BindingId               string `json:"binding_id"`                 // [Required] <p>Binding ID</p>
	FirstMileTrackingNumber string `json:"first_mile_tracking_number"` // [Required] <p>The generated first-mile tracking number, value might be blank.</p>
}

type UnbindFirstMileTrackingNumberRequest struct {
	FirstMileTrackingNumber string        `json:"first_mile_tracking_number"` // [Required] The identifier for an API request for error tracking.
	OrderList               []SharedOrder `json:"order_list"`                 // [Required] <p>The list of order info you want to unbind from the given first mile tracking number.</p><p>You can specify up to 50 orders in this call.<br /></p>
}

type UnbindFirstMileTrackingNumberResponse struct {
	BaseResponse `json:",inline"`                          // Common response fields
	Response     UnbindFirstMileTrackingNumberResponseData `json:"response"` // Actual response data
}

type UnbindFirstMileTrackingNumberResponseData struct {
	FirstMileTrackingNumber string                                     `json:"first_mile_tracking_number"` // [Required] The first mile tracking number.
	OrderList               []CreateShippingDocumentResponseDataResult `json:"order_list"`                 // [Required] The binding result list of each order.
}

type Waybill struct {
	BindingId        string `json:"binding_id"`         // [Required] <p>Binding ID</p>
	ShippingLabelUrl string `json:"shipping_label_url"` // [Required] <p>URL for downloading waybill.</p>
}
