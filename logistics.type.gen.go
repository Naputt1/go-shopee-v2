package goshopee

type AddressTimeSlot struct {
	Date         int64    `json:"date"`           // [Required] <p>The date of pickup time. In timestamp.<br /></p>
	TimeText     string   `json:"time_text"`      // [Required] <p>The text description of pickup time. Only applicable for certain channels.<br /></p>
	PickupTimeId string   `json:"pickup_time_id"` // [Required] <p>The identity of pickuptime.<br /></p>
	Flags        []string `json:"flags"`          // [Required] <p>This field will have the value&nbsp;<font color="#c24f4a"><b>“recommended”</b></font>&nbsp;for the time slot that Shopee suggests sellers choose.&nbsp;</p><p><br /></p><p>While it is advisable for sellers to choose the recommended time slot, they can also choose other time slots that do not have the recommended flag.</p>
	Error        string   `json:"error"`          // [Required] <p>return if error getting pickup time, otherwise omitted<br /></p>
	Msg          string   `json:"msg"`            // [Required] <p>return if error getting pickup time, otherwise omitted<br /></p>
}

type AddressTypeConfig struct {
	AddressId   *int64   `json:"address_id,omitempty"`   // [Optional] The identifier id of the address.
	AddressType []string `json:"address_type,omitempty"` // [Optional] The type of addres. Available values: DEFAULT_ADDRESS, PICKUP_ADDRESS, RETURN_ADDRESS
}

type BatchShipOrderRequest struct {
	OrderList     []SharedOrder   `json:"order_list"`               // [Required] The list of order.
	Pickup        *Pickup         `json:"pickup,omitempty"`         // [Optional] Required parameter ONLY if GetParameterForInit returns "pickup" or if GetLogisticsInfo returns "pickup" under "info_needed" for the same order. Developer should still include "pickup" field in the call even if "pickup" has empty value.
	Dropoff       *RequestDropoff `json:"dropoff,omitempty"`        // [Optional] Required parameter ONLY if GetParameterForInit returns "dropoff" or if GetLogisticsInfo returns "dropoff" under "info_needed" for the same order. Developer should still include "dropoff" field in the call even if "dropoff" has empty value. For logistic_id 80003 and 80004, both Regular and JOB shipping methods are supported. If you choose Regular shipping method, please use "tracking_no" to call Init API. If you choose JOB shipping method, please use "sender_real_name" to call Init API. Note that only one of "tracking_no" and "sender_real_name" can be selected.
	NonIntegrated *NonIntegrated  `json:"non_integrated,omitempty"` // [Optional] Optional parameter when GetParameterForInit returns "non-integrated" or GetLogisticsInfo returns "non-integrated" under "info_needed".
}

type BatchShipOrderResponse struct {
	BaseResponse `json:",inline"`           // Common response fields
	Response     BatchShipOrderResponseData `json:"response"` // Actual response data
}

type BatchShipOrderResponseData struct {
	ResultList []CreateShippingDocumentResponseDataResult `json:"result_list"` // [Required]
}

type BatchUpdateTpfWarehouseTrackingStatusRequest struct {
	TpfName           string                                                `json:"tpf_name"`            // [Required] <p>The name of 3PF Warehouse Vendor. Prohibit pure numbers and excessive abbreviations. Standardize naming for easy business recognition. Input priority: warehouse English name &gt; full pinyin of warehouse brand name &gt; warehouse Chinese name &gt; other officially recognized and prominent names.</p><br /><br />
	TpfTrackingStatus string                                                `json:"tpf_tracking_status"` // [Required] <p>The 3PF tracking status for the timestamp. All statuses are in lower case.&nbsp;</p><p>List of status:&nbsp;</p><p>- 3pf_warehouse_order_created</p><p>- 3pf_warehouse_outbound_done</p>
	PackageList       []BatchUpdateTpfWarehouseTrackingStatusRequestPackage `json:"package_list"`        // [Required]
}

type BatchUpdateTpfWarehouseTrackingStatusRequestPackage struct {
	OrderSn       string  `json:"order_sn"`                 // [Required] <p>Shopee's unique identifier for an order.</p>
	PackageNumber *string `json:"package_number,omitempty"` // [Optional] <p>Shopee's unique identifier for the package under an order.</p>
	UpdateTime    int64   `json:"update_time"`              // [Required] <p>This is to indicate timestamp of the 3PF tracking status.</p><p>Timestamp should be within order create time and order pick up by 3PL time.</p>
}

type BatchUpdateTpfWarehouseTrackingStatusResponse struct {
	BaseResponse `json:",inline"`                                  // Common response fields
	Response     BatchUpdateTpfWarehouseTrackingStatusResponseData `json:"response"` // Actual response data
}

type BatchUpdateTpfWarehouseTrackingStatusResponseData struct {
	SuccessList []SharedOrder                              `json:"success_list"` // [Required] <p>Update success order list.</p>
	FailList    []CreateShippingDocumentResponseDataResult `json:"fail_list"`    // [Required] <p>Update fail order list.</p>
}

type Booking struct {
	BookingSn string `json:"booking_sn"` // [Required] <p>Shopee's unique identifier for a booking.<br /></p>
}

type Branch struct {
	BranchId int64  `json:"branch_id"` // [Required] <p>The identity of logistics branch.<br /></p>
	Region   string `json:"region"`    // [Required] <p>The region of specify address.<br /></p>
	State    string `json:"state"`     // [Required] <p>The state of specify address.<br /></p>
	City     string `json:"city"`      // [Required] <p>The city of specify address.<br /></p>
	Address  string `json:"address"`   // [Required] <p>The address description of specify address.<br /></p>
	Zipcode  string `json:"zipcode"`   // [Required] <p>The zipcode of specify address.<br /></p>
	District string `json:"district"`  // [Required] <p>The district of specify address.<br /></p>
	Town     string `json:"town"`      // [Required] <p>The town of specify address.<br /></p>
}

type BuyerPreferDeliveryTime struct {
	SlotId      string `json:"slot_id"`     // [Required] <p>The slot which buyer choose<br /></p>
	StartTime   string `json:"start_time"`  // [Required] <p>The start time of a day buyer prefer to receive the packages<br /></p>
	EndTime     string `json:"end_time"`    // [Required] <p>The end time of a day buyer prefer to receive the packages.<br /></p>
	Description string `json:"description"` // [Required] <p>The detailed instructions of the package delivering.<br /></p>
}

type CheckPolygonUpdateStatusRequest struct {
	TaskId string `json:"task_id"` // [Required] <p>ID that needs to be checked. Please pass the task_id returned via the v2.logistics.upload_serviceable_polygon.</p>
}

type CheckPolygonUpdateStatusResponse struct {
	BaseResponse `json:",inline"`                     // Common response fields
	Response     CheckPolygonUpdateStatusResponseData `json:"response"` // Actual response data
}

type CheckPolygonUpdateStatusResponseData struct {
	Status  int64  `json:"status"`  // [Required] <p>Serviceable polygon file upload status. Applicable values:&nbsp;</p><p>0: Task completed<br />1: Task in progress<br />2: KML file related errors</p>
	Message string `json:"message"` // [Required] <p>Details of the upload status, e.g "task in progress".</p>
}

type CreateBookingShippingDocumentRequest struct {
	BookingList []RequestBooking `json:"booking_list"` // [Required] <p>The list of bookings you want to get. limit [1,50]<br /></p>
}

type CreateBookingShippingDocumentResponse struct {
	BaseResponse `json:",inline"`                          // Common response fields
	Response     CreateBookingShippingDocumentResponseData `json:"response"` // Actual response data
}

type CreateBookingShippingDocumentResponseData struct {
	ResultList []CreateBookingShippingDocumentResponseDataResult `json:"result_list"` // [Required] <p>The list of the result data.<br /></p>
}

type CreateBookingShippingDocumentResponseDataResult struct {
	BookingSn   string `json:"booking_sn"`   // [Required] <p>Shopee's unique identifier for a booking.<br /></p>
	FailError   string `json:"fail_error"`   // [Required] <p>Indicate error type if one element hit error.<br /></p>
	FailMessage string `json:"fail_message"` // [Required] <p>Indicate error details if one element hit error.<br /></p>
}

type CreateShippingDocumentJobRequest struct {
	ShippingDocumentType  string                  `json:"shipping_document_type"`            // [Required] <p>The type of shipping document. Available values: THERMAL_UNPACKAGED_LABEL</p>
	UnpackagedSkuRequests []UnpackagedSkuRequests `json:"unpackaged_sku_requests,omitempty"` // [Optional] <p>List of Unpackaged SKUs to generate labels for.</p><p><br /></p><p>Note: The unpackaged_sku_requests and package_list cannot be populated at the same time, please select one.</p>
	PackageList           []string                `json:"package_list,omitempty"`            // [Optional] <p>List of Package Numbers to generate labels for.&nbsp;(maximum 600 total)</p><p><br /></p><p>Note: The unpackaged_sku_requests and package_list cannot be populated at the same time, please select one.</p>
}

type CreateShippingDocumentJobResponse struct {
	BaseResponse `json:",inline"`                      // Common response fields
	Response     CreateShippingDocumentJobResponseData `json:"response"` // Actual response data
}

type CreateShippingDocumentJobResponseData struct {
	JobId         string             `json:"job_id"`          // [Required] <p>Generated Job ID which will be used for status tracking and download the Shipping Document</p>
	SuccessIdList []string           `json:"success_id_list"` // [Required] <p>List of Package Number or Unpackaged SKU ID that succeeds in generating Shipping Document</p>
	FailList      []ResponseDataFail `json:"fail_list"`       // [Required] <p>List of Package Numbers or Unpackaged SKUs that failed in generating Shipping Document</p>
}

type CreateShippingDocumentRequest struct {
	OrderList []RequestOrder `json:"order_list"` // [Required] The list of order you want to create shipping document. limit [1, 50]
}

type CreateShippingDocumentResponse struct {
	BaseResponse `json:",inline"`                   // Common response fields
	Response     CreateShippingDocumentResponseData `json:"response"` // Actual response data
}

type CreateShippingDocumentResponseData struct {
	ResultList []CreateShippingDocumentResponseDataResult `json:"result_list"` // [Required] The list of the result data.
}

type DeleteAddressRequest struct {
	AddressId int64 `json:"address_id"` // [Required] The identity of address you want to delete.
}

type DeleteAddressResponse struct {
	BaseResponse `json:",inline"` // Common response fields
}

type DeleteSpecialOperatingHourRequest struct {
	Name string `json:"name"` // [Required] <p>Name of the special operating hour which can be retrieved from v2.logistics.get_operating_hours</p>
}

type DeleteSpecialOperatingHourResponse struct {
	BaseResponse `json:",inline"` // Common response fields
}

type DownloadBookingShippingDocumentRequest struct {
	ShippingDocumentType *string   `json:"shipping_document_type,omitempty"` // [Optional] <p>The type of shipping document. Available values: NORMAL_AIR_WAYBILL,THERMAL_AIR_WAYBILL<br /></p>
	BookingList          []Booking `json:"booking_list"`                     // [Required] <p>The list of bookings you want to get. limit [1,50]<br /></p>
}

type DownloadBookingShippingDocumentResponse struct {
	BaseResponse `json:",inline"` // Common response fields
	Waybill      interface{}      `json:"waybill,omitempty"` // <p>The waybill file.<br /></p>
}

type DownloadShippingDocumentJobRequest struct {
	JobId string `json:"job_id"` // [Required] <p>Generated Job ID for status tracking and download the Shipping Document</p>
}

type DownloadShippingDocumentJobResponse struct {
	BaseResponse `json:",inline"` // Common response fields
	File         interface{}      `json:"file,omitempty"` //
}

type DownloadShippingDocumentRequest struct {
	ShippingDocumentType *string       `json:"shipping_document_type,omitempty"` // [Optional] <p>The type of shipping document. Available values: NORMAL_AIR_WAYBILL, THERMAL_AIR_WAYBILL, NORMAL_JOB_AIR_WAYBILL, THERMAL_JOB_AIR_WAYBILL, THERMAL_UNPACKAGED_LABEL</p>
	OrderList            []SharedOrder `json:"order_list"`                       // [Required] The list of orders you need to download it's shipping document.
}

type DownloadShippingDocumentResponse struct {
	BaseResponse `json:",inline"` // Common response fields
	Waybill      interface{}      `json:"waybill,omitempty"` // <p>The waybill file.</p>
}

type DownloadToLabelRequest struct {
	SortingGroup int64  `json:"sorting_group"`      // [Required] <p>Sorting Group of the TO. Available value:</p><p>1:North</p><p>2:South</p>
	Quantity     *int64 `json:"quantity,omitempty"` // [Optional] <p>Specifies the TO quantity, up to a maximum of 20 per request. If not specified, the default value is 1</p>
}

type DownloadToLabelResponse struct {
	BaseResponse `json:",inline"` // Common response fields
	Waybill      interface{}      `json:"waybill,omitempty"` // <p>The waybill file.</p>
}

type Dropoff struct {
	BranchId       *int64  `json:"branch_id,omitempty"`        // [Optional] The identity of branch.
	SenderRealName *string `json:"sender_real_name,omitempty"` // [Optional] The real name of sender.
	TrackingNumber *string `json:"tracking_number,omitempty"`  // [Optional] Need input this field when "tracking_number" is returned from "info_need". Please note that this tracking number is assigned by third-party shipping carrier for item shipment.
	Slug           *string `json:"slug,omitempty"`             // [Optional]  The selected 3PL partner to drop-off parcels with.
}

type Fail struct {
	PackageNumber string `json:"package_number"` // [Required] <p>Shopee's unique identifier for the package under an order.</p>
	FailReason    string `json:"fail_reason"`    // [Required] <p>Reason for failure.</p>
}

type GetAddressListResponse struct {
	BaseResponse `json:",inline"`           // Common response fields
	Response     GetAddressListResponseData `json:"response"` // Actual response data
}

type GetAddressListResponseData struct {
	ShowPickupAddress bool                  `json:"show_pickup_address"` // [Required] Show pickup address or not.
	AddressList       []ResponseDataAddress `json:"address_list"`        // [Required] The address list of you shop
}

type GetBookingShippingDocumentDataInfoRequest struct {
	BookingSn            string                 `json:"booking_sn"`                       // [Required] <p>Shopee's unique identifier for a booking.<br /></p>
	RecipientAddressInfo []RecipientAddressInfo `json:"recipient_address_info,omitempty"` // [Optional] <p>recipient address to query as image<br /></p>
}

type GetBookingShippingDocumentDataInfoResponse struct {
	BaseResponse `json:",inline"`                               // Common response fields
	Response     GetBookingShippingDocumentDataInfoResponseData `json:"response"` // Actual response data
}

type GetBookingShippingDocumentDataInfoResponseData struct {
	RecipientAddressInfo *ResponseDataRecipientAddressInfo `json:"recipient_address_info"` // [Required]
	ShippingDocumentInfo *ResponseDataShippingDocumentInfo `json:"shipping_document_info"` // [Required]
}

type GetBookingShippingDocumentParameterRequest struct {
	BookingList []Booking `json:"booking_list"` // [Required] <p>The list of bookings you want to get. limit [1,50]<br /></p>
}

type GetBookingShippingDocumentParameterResponse struct {
	BaseResponse `json:",inline"`                                // Common response fields
	Response     GetBookingShippingDocumentParameterResponseData `json:"response"` // Actual response data
}

type GetBookingShippingDocumentParameterResponseData struct {
	ResultList []GetBookingShippingDocumentParameterResponseDataResult `json:"result_list"` // [Required] <p>The list of the result data.<br /></p>
}

type GetBookingShippingDocumentParameterResponseDataResult struct {
	BookingSn                      string   `json:"booking_sn"`                        // [Required] <p>Shopee's unique identifier for a booking.<br /></p>
	SuggestShippingDocumentType    string   `json:"suggest_shipping_document_type"`    // [Required] <p>The shipping document type Shopee suggests. If you don't select any shipping document type, Shopee will use this as default shipping document type.<br /></p>
	SelectableShippingDocumentType []string `json:"selectable_shipping_document_type"` // [Required] <p>The shipping document type you can select of this booking.<br /></p>
	FailError                      string   `json:"fail_error"`                        // [Required] <p>Indicate error type if one element hit error.<br /></p>
	FailMessage                    string   `json:"fail_message"`                      // [Required] <p>Indicate error details if one element hit error.<br /></p>
}

type GetBookingShippingDocumentResultRequest struct {
	BookingList []GetBookingShippingDocumentResultRequestBooking `json:"booking_list"` // [Required] <p>The list of bookings you want to get. limit [1,50]<br /></p>
}

type GetBookingShippingDocumentResultRequestBooking struct {
	BookingSn            string  `json:"booking_sn"`                       // [Required] <p>Shopee's unique identifier for a booking.<br /></p>
	ShippingDocumentType *string `json:"shipping_document_type,omitempty"` // [Optional] <p>The type of shipping document. Available values: NORMAL_AIR_WAYBILL,THERMAL_AIR_WAYBILL<br /></p>
}

type GetBookingShippingDocumentResultResponse struct {
	BaseResponse `json:",inline"`                             // Common response fields
	Response     GetBookingShippingDocumentResultResponseData `json:"response"` // Actual response data
}

type GetBookingShippingDocumentResultResponseData struct {
	ResultList []GetBookingShippingDocumentResultResponseDataResult `json:"result_list"` // [Required] <p>The list of the result data.<br /></p>
}

type GetBookingShippingDocumentResultResponseDataResult struct {
	BookingSn   string `json:"booking_sn"`   // [Required] <p>Shopee's unique identifier for a booking.<br /></p>
	Status      string `json:"status"`       // [Required] <p>The status of the shipping document task you querying with booking_sn. Available values: READY/FAILED/PROCESSING<br /></p>
	FailError   string `json:"fail_error"`   // [Required] <p>Indicate error type if one element hit error.<br /></p>
	FailMessage string `json:"fail_message"` // [Required] <p>Indicate error details if one element hit error.<br /></p>
}

type GetBookingShippingParameterRequest struct {
	BookingSn string `json:"booking_sn" url:"booking_sn"` // [Required] <p>Shopee's unique identifier for a booking.<br /></p>
}

type GetBookingShippingParameterResponse struct {
	BaseResponse `json:",inline"`                        // Common response fields
	Response     GetBookingShippingParameterResponseData `json:"response"` // Actual response data
}

type GetBookingShippingParameterResponseData struct {
	InfoNeeded *ResponseDataInfoNeeded                        `json:"info_needed"` // [Required] <p>The parameters required based on each specific booking to Init. Must use the fields included under info_needed to call Init.<br /></p>
	Pickup     *GetBookingShippingParameterResponseDataPickup `json:"pickup"`      // [Required] <p>Logistics information for pickup mode booking.<br /></p>
}

type GetBookingShippingParameterResponseDataPickup struct {
	AddressList []ResponseDataPickupAddress `json:"address_list"` // [Required] <p>List of available pickup address info.<br /></p>
}

type GetBookingTrackingInfoRequest struct {
	BookingSn string `json:"booking_sn" url:"booking_sn"` // [Required] <p>Shopee's unique identifier for a booking.<br /></p>
}

type GetBookingTrackingInfoResponse struct {
	BaseResponse `json:",inline"`                   // Common response fields
	Response     GetBookingTrackingInfoResponseData `json:"response"` // Actual response data
}

type GetBookingTrackingInfoResponseData struct {
	BookingSn       string          `json:"booking_sn"`       // [Required] <p>Shopee's unique identifier for a booking.<br /></p>
	LogisticsStatus LogisticsStatus `json:"logistics_status"` // [Required] <p>The Shopee logistics status for the booking. Applicable values.<br /></p><p>LOGISTICS_REQUEST_CREATED:booking arranged shipment<br /></p><p>LOGISTICS_PICKUP_DONE:booking handed over to 3PL<br /></p><p>LOGISTICS_PICKUP_FAILED:booking&nbsp;cancelled by 3PL due to failed pickup or picked up but not able to proceed with delivery<br /></p><p>LOGISTICS_DELIVERY_DONE:successfully delivered<br /></p><p>LOGISTICS_REQUEST_CANCELED:cancelled when booking at LOGISTICS_REQUEST_CREATED<br /></p><p>LOGISTICS_READY:booking&nbsp;ready for fulfilment<br /></p><p>LOGISTICS_INVALID:cancelled when booking at LOGISTICS_READY<br /></p><p>LOGISTICS_LOST:booking cancelled due to 3PL lost the parcel<br /></p>
	TrackingInfo    []TrackingInfo  `json:"tracking_info"`    // [Required] <p>The tracking info of the booking.<br /></p>
}

type GetBookingTrackingNumberRequest struct {
	BookingSn string `json:"booking_sn" url:"booking_sn"` // [Required] <p>Shopee's unique identifier for a booking.<br /></p>
}

type GetBookingTrackingNumberResponse struct {
	BaseResponse   `json:",inline"` // Common response fields
	TrackingNumber string           `json:"tracking_number,omitempty"` // <p>The tracking number of this booking.<br /></p>
}

type GetChannelListResponse struct {
	BaseResponse `json:",inline"`           // Common response fields
	Response     GetChannelListResponseData `json:"response"` // Actual response data
}

type GetChannelListResponseData struct {
	LogisticsChannelList []LogisticsChannel `json:"logistics_channel_list"` // [Required] The list of logistics channel.
}

type GetMartPackagingInfoResponse struct {
	BaseResponse `json:",inline"`                 // Common response fields
	Response     GetMartPackagingInfoResponseData `json:"response"` // Actual response data
}

type GetMartPackagingInfoResponseData struct {
	Enable       bool              `json:"enable"`        // [Required] <p>Indicates whether the seller has enabled or disabled the packaging fee configuration.</p><p><b>True:</b>&nbsp;The seller charges a packaging fee.</p><p><b>False:</b>&nbsp;The seller does not charge a packaging fee.</p>
	Dimension    *RequestDimension `json:"dimension"`     // [Required] <p>Returned only if&nbsp;enabled&nbsp;is set to&nbsp;True.</p>
	PackagingFee *PackagingFee     `json:"packaging_fee"` // [Required] <p>Returned only if&nbsp;enabled&nbsp;is set to&nbsp;True.</p>
}

type GetMassShippingParameterRequest struct {
	LogisticsChannelId *int64           `json:"logistics_channel_id,omitempty"` // [Optional] <p><font color="#c24f4a">The API can only batch request the shipping parameter for multiple packages under the same product_location_id and same logistics_channel_id.&nbsp;</font></p><p><br /></p><p>Use this field to specify the logistics_channel_id for the request. <b>If not specified, will use the logistics_channel_id corresponds to the first package_number by default.</b></p>
	ProductLocationId  *string          `json:"product_location_id,omitempty"`  // [Optional] <p><font color="#c24f4a">The API can only batch request the shipping parameter for multiple packages under the same product_location_id and same logistics_channel_id.&nbsp;</font></p><p><br /></p><p>Use this field to specify the product_location_id for the request. <b>If not specified, will use the product_location_id corresponds to the first package_number by default.</b></p>
	PackageList        []RequestPackage `json:"package_list"`                   // [Required] <p>The list of packages you want to get shipping parameters. limit [1, 50].</p>
}

type GetMassShippingParameterResponse struct {
	BaseResponse `json:",inline"`                     // Common response fields
	Response     GetMassShippingParameterResponseData `json:"response"` // Actual response data
}

type GetMassShippingParameterResponseData struct {
	InfoNeeded  *InfoNeeded                                  `json:"info_needed"`  // [Required] <p>The parameters required based on each specific order to Init. Must use the fields included under info_needed to call Init. For logistic_id 80003 and 80004, both Regular and JOB shipping methods are supported. If you choose Regular shipping method, please use "tracking_no" to call Init API. If you choose JOB shipping method, please use "sender_real_name" to call Init API. Note that only one of "tracking_no" and "sender_real_name" can be selected.</p>
	Dropoff     *GetMassShippingParameterResponseDataDropoff `json:"dropoff"`      // [Required] <p>Logistics information for dropoff mode package.<br /></p>
	Pickup      *ResponseDataPickup                          `json:"pickup"`       // [Required] <p>Logistics information for pickup mode package.<br /></p>
	SuccessList []RequestPackage                             `json:"success_list"` // [Required] <p>Success package list.</p>
	FailList    []Fail                                       `json:"fail_list"`    // [Required] <p>Fail package list.</p>
}

type GetMassShippingParameterResponseDataDropoff struct {
	BranchList []Branch `json:"branch_list"` // [Required] <p>List of available dropoff branches info.<br /></p>
}

type GetMassTrackingNumberRequest struct {
	PackageList            []RequestPackage `json:"package_list"`                       // [Required] <p>The list of packages you want to get tracking number. limit [1, 50].</p>
	ResponseOptionalFields *string          `json:"response_optional_fields,omitempty"` // [Optional] <p>Indicate response fields you want to get. Please select from the below response parameters. If you input an object field, all the params under it will be included automatically in the response. If there are multiple response fields you want to get, you need to use English comma to connect them. Available values: plp_number, first_mile_tracking_number,last_mile_tracking_number.</p>
}

type GetMassTrackingNumberResponse struct {
	BaseResponse `json:",inline"`                  // Common response fields
	Response     GetMassTrackingNumberResponseData `json:"response"` // Actual response data
}

type GetMassTrackingNumberResponseData struct {
	SuccessList []GetMassTrackingNumberResponseDataSuccess `json:"success_list"` // [Required] <p>Success package list.</p>
	FailList    []Fail                                     `json:"fail_list"`    // [Required] <p>Fail package list.</p>
}

type GetMassTrackingNumberResponseDataSuccess struct {
	PackageNumber           string `json:"package_number"`             // [Required] <p>Shopee's unique identifier for the package under an order.</p>
	TrackingNumber          string `json:"tracking_number"`            // [Required] <p>The tracking number of this order.</p>
	PlpNumber               string `json:"plp_number"`                 // [Required] <p>The unique identifier for package of BR correios.</p>
	FirstMileTrackingNumber string `json:"first_mile_tracking_number"` // [Required] <p>The first mile tracking number of the order. Only for Cross Border Seller</p>
	LastMileTrackingNumber  string `json:"last_mile_tracking_number"`  // [Required] <p>The last mile tracking number of the order. Only for Cross Border BR seller.</p>
	Hint                    string `json:"hint"`                       // [Required] <p>Indicate hint information if cannot get some fields under special scenarios. For example, cannot get tracking_number when cvs store is closed.</p>
	PickupCode              string `json:"pickup_code"`                // [Required] <p>For drivers to quickly identify parcel to be picked up.&nbsp;Only returned for instant+sameday orders.</p>
}

type GetOperatingHourRestrictionsResponse struct {
	BaseResponse `json:",inline"`                         // Common response fields
	Response     GetOperatingHourRestrictionsResponseData `json:"response"` // Actual response data
}

type GetOperatingHourRestrictionsResponseData struct {
	RegularOperatingHourRestrictions        *RegularOperatingHourRestrictions        `json:"regular_operating_hour_restrictions"`         // [Required] <p>The restrictions for Pickup Operating Hours / Preferred Pickup Hours</p>
	InstantOperatingHourRestrictions        *RegularOperatingHourRestrictions        `json:"instant_operating_hour_restrictions"`         // [Required] <p>The restrictions for Instant Operating Hours</p>
	SpecialOperatingHourRestrictions        *SpecialOperatingHourRestrictions        `json:"special_operating_hour_restrictions"`         // [Required] <p>The restrictions for Special Operating Hours</p>
	ShopCollectionOperatingHourRestrictions *ShopCollectionOperatingHourRestrictions `json:"shop_collection_operating_hour_restrictions"` // [Required] <p>The restrictions for Shop Collection Operating Hours</p>
}

type GetOperatingHoursResponse struct {
	BaseResponse `json:",inline"` // Common response fields
	Repsonse     *Repsonse        `json:"repsonse,omitempty"` //
}

type GetShippingDocumentDataInfoRequest struct {
	OrderSn              string                 `json:"order_sn"`                         // [Required] <p>Shopee's unique identifier for an order.<br /></p>
	PackageNumber        *string                `json:"package_number,omitempty"`         // [Optional] <p>Shopee's unique identifier for the package under an order. You shouldn't fill the field with empty string when there isn't a package number.<br /></p>
	RecipientAddressInfo []RecipientAddressInfo `json:"recipient_address_info,omitempty"` // [Optional] <p>recipient address to query as image</p>
}

type GetShippingDocumentDataInfoResponse struct {
	BaseResponse `json:",inline"`                        // Common response fields
	Response     GetShippingDocumentDataInfoResponseData `json:"response"` // Actual response data
}

type GetShippingDocumentDataInfoResponseData struct {
	RecipientAddressInfo *ResponseDataRecipientAddressInfo `json:"recipient_address_info"` // [Required]
	ShippingDocumentInfo *ShippingDocumentInfo             `json:"shipping_document_info"` // [Required]
}

type GetShippingDocumentJobStatusRequest struct {
	JobId string `json:"job_id"` // [Required] <p>Generated Job ID for status tracking and download the Shipping Document</p>
}

type GetShippingDocumentJobStatusResponse struct {
	BaseResponse `json:",inline"`                         // Common response fields
	Response     GetShippingDocumentJobStatusResponseData `json:"response"` // Actual response data
}

type GetShippingDocumentJobStatusResponseData struct {
	JobId     string `json:"job_id"`     // [Required] <p>Generated Job ID for status tracking and download the Shipping Document</p>
	JobName   string `json:"job_name"`   // [Required] <p>Generated Shipping Document file name.</p>
	JobStatus string `json:"job_status"` // [Required] <p>Requested Shipping Document current status. Available values:&nbsp;<span style="font-size:14px;">PROCESSING,&nbsp;</span><span style="font-size:14px;">READY,&nbsp;</span><span style="font-size:14px;">EXPIRED,&nbsp;</span><span style="font-size:14px;">FAILED</span></p>
}

type GetShippingDocumentParameterRequest struct {
	OrderList []SharedOrder `json:"order_list"` // [Required] The list of orders you want to get. limit [1,50]
}

type GetShippingDocumentParameterResponse struct {
	BaseResponse `json:",inline"`                         // Common response fields
	Response     GetShippingDocumentParameterResponseData `json:"response"` // Actual response data
}

type GetShippingDocumentParameterResponseData struct {
	ResultList []ResponseDataResult `json:"result_list"` // [Required] The list of the result data.
}

type GetShippingDocumentResultRequest struct {
	OrderList []GetShippingDocumentResultRequestOrder `json:"order_list"` // [Required] The list of orders, limit [1,50]
}

type GetShippingDocumentResultRequestOrder struct {
	OrderSn              string  `json:"order_sn"`                         // [Required] Shopee's unique identifier for an order.
	PackageNumber        *string `json:"package_number,omitempty"`         // [Optional] Shopee's unique identifier for the package under an order. You should't fill the field with empty string when there is't a package number.
	ShippingDocumentType *string `json:"shipping_document_type,omitempty"` // [Optional] <p>The type of shipping document. Available values: NORMAL_AIR_WAYBILL, THERMAL_AIR_WAYBILL, NORMAL_JOB_AIR_WAYBILL, THERMAL_JOB_AIR_WAYBILL, THERMAL_UNPACKAGED_LABEL</p>
}

type GetShippingDocumentResultResponse struct {
	BaseResponse `json:",inline"`                      // Common response fields
	Response     GetShippingDocumentResultResponseData `json:"response"` // Actual response data
}

type GetShippingDocumentResultResponseData struct {
	ResultList []GetShippingDocumentResultResponseDataResult `json:"result_list"` // [Required] The result data list of the API response.
}

type GetShippingDocumentResultResponseDataResult struct {
	OrderSn       string `json:"order_sn"`       // [Required] Shopee's unique identifier for an order.
	PackageNumber string `json:"package_number"` // [Required] Shopee's unique identifier for the package under an order.
	Status        string `json:"status"`         // [Required] The status of the shipping document task you querying with order_sn. Available values: READY， FAILED， PROCESSING
	FailError     string `json:"fail_error"`     // [Required] Indicate error type if one element hit error.
	FailMessage   string `json:"fail_message"`   // [Required] Indicate error details if one element hit error.
}

type GetShippingParameterRequest struct {
	OrderSn       string  `json:"order_sn" url:"order_sn"`                                 // [Required] Shopee's unique identifier for an order.
	PackageNumber *string `json:"package_number,omitempty" url:"package_number,omitempty"` // [Optional] <p>Shopee's unique identifier for the package under an order. You should't fill the field with empty string when there is't a package number.<br /></p>
}

type GetShippingParameterResponse struct {
	BaseResponse `json:",inline"`                 // Common response fields
	Response     GetShippingParameterResponseData `json:"response"` // Actual response data
}

type GetShippingParameterResponseData struct {
	InfoNeeded *InfoNeeded          `json:"info_needed"` // [Required] The parameters required based on each specific order to Init. Must use the fields included under info_needed to call Init. For logistic_id 80003 and 80004, both Regular and JOB shipping methods are supported. If you choose Regular shipping method, please use "tracking_no" to call Init API. If you choose JOB shipping method, please use "sender_real_name" to call Init API. Note that only one of "tracking_no" and "sender_real_name" can be selected.
	Dropoff    *ResponseDataDropoff `json:"dropoff"`     // [Required] Logistics information for dropoff mode order.
	Pickup     *ResponseDataPickup  `json:"pickup"`      // [Required] Logistics information for pickup mode order.
}

type GetTrackingInfoRequest struct {
	OrderSn       string  `json:"order_sn" url:"order_sn"`                                 // [Required] <p>Shopee's unique identifier for an order.</p>
	PackageNumber *string `json:"package_number,omitempty" url:"package_number,omitempty"` // [Optional] <p>Shopee's unique identifier for the package under an order. You shouldn't fill the field with empty string when there is a package number.</p>
}

type GetTrackingInfoResponse struct {
	BaseResponse `json:",inline"`            // Common response fields
	Response     GetTrackingInfoResponseData `json:"response"` // Actual response data
}

type GetTrackingInfoResponseData struct {
	OrderSn                string                 `json:"order_sn"`                 // [Required] Shopee's unique identifier for an order.
	PackageNumber          string                 `json:"package_number"`           // [Required] Shopee's unique identifier for the package under an order.
	LogisticsStatus        LogisticsStatus        `json:"logistics_status"`         // [Required] <p>The logistics status for the order. Applicable values: See Data Definition- LogisticsStatus.</p>
	TrackingInfo           []TrackingInfo         `json:"tracking_info"`            // [Required] The tracking info of the order.
	ReversedTrackingNumber string                 `json:"reversed_tracking_number"` // [Required] <p>The tracking number of the reversed logistics.</p><p><br /></p><p>Note: Only apply to the cross-border segment of failed delivery parcels returned from the local return warehouse to the seller.</p>
	ReversedCourierName    string                 `json:"reversed_courier_name"`    // [Required] <p>The courier name of the reversed logistics.</p><p><br /></p><p>Note: Only apply to the cross-border segment of failed delivery parcels returned from the local return warehouse to the seller.</p>
	ReversedTrackingInfo   []ReversedTrackingInfo `json:"reversed_tracking_info"`   // [Required] <p>The tracking information of the reversed logistics.</p><p><br /></p><p>Note: Only apply to the cross-border segment of failed delivery parcels returned from the local return warehouse to the seller.</p>
}

type GetTrackingNumberRequest struct {
	OrderSn                string  `json:"order_sn" url:"order_sn"`                                                     // [Required] Shopee's unique identifier for an order.
	PackageNumber          *string `json:"package_number,omitempty" url:"package_number,omitempty"`                     // [Optional] Shopee's unique identifier for the package under an order. You should't fill the field with empty string when there isn't a package number.
	ResponseOptionalFields *string `json:"response_optional_fields,omitempty" url:"response_optional_fields,omitempty"` // [Optional] Indicate response fields you want to get. Please select from the below response parameters. If you input an object field, all the params under it will be included automatically in the response. If there are multiple response fields you want to get, you need to use English comma to connect them. Available values: plp_number, first_mile_tracking_number,last_mile_tracking_number
}

type GetTrackingNumberResponse struct {
	BaseResponse `json:",inline"`              // Common response fields
	Response     GetTrackingNumberResponseData `json:"response"` // Actual response data
}

type GetTrackingNumberResponseData struct {
	TrackingNumber          string `json:"tracking_number"`            // [Required] The tracking number of this order.
	PlpNumber               string `json:"plp_number"`                 // [Required] The unique identifier for package of BR correios.
	FirstMileTrackingNumber string `json:"first_mile_tracking_number"` // [Required] The first mile tracking number of the order. Only for Cross Border Seller
	LastMileTrackingNumber  string `json:"last_mile_tracking_number"`  // [Required] The last mile tracking number of the order. Only for Cross Border BR seller.
	Hint                    string `json:"hint"`                       // [Required] Indicate hint information if cannot get some fields under special scenarios. For example, cannot get tracking_number when cvs store is closed.
	PickupCode              string `json:"pickup_code"`                // [Required] <p>For drivers to quickly identify parcel to be picked up.&nbsp;Only returned for ID local orders who using instant+sameday for delivery.<br /></p>
}

type InfoNeeded struct {
	Dropoff       []string `json:"dropoff"`        // [Required] <p>Could contain 'branch_id', 'sender_real_name', or 'tracking_no'. If it contains 'branch_id', choose one to Init. If it contains 'sender_real_name' or 'tracking_no', should manually input these values in Init API. If it has empty value, developer should still include "dropoff" field in Init API.</p>
	Pickup        []string `json:"pickup"`         // [Required] <p>Could contain 'address_id' and 'pickup_time_id'. Choose one address_id and its corresponding pickup_time_id to Init. If it has empty value, developer should still include "pickup" field in Init API. It could contains "tracking_number" returned from "info_need"for some channels, please also add it when init.</p>
	NonIntegrated []string `json:"non_integrated"` // [Required] <p>Could contain 'tracking_no'. If it contains 'tracking_no', should manually input these values in Init API. If it has empty value, developer should still include "non-integrated" field in Init API.</p>
}

type InstantOperatingHour struct {
	Monday        *Monday `json:"monday,omitempty"`         // [Optional] <p>Operating hours for Monday: You can skip this information if you want to mark the day as closed.</p>
	Tuesday       *Monday `json:"tuesday,omitempty"`        // [Optional] <p>Operating hours for Tuesday: You can skip this information if you want to mark the day as closed.</p>
	Wednesday     *Monday `json:"wednesday,omitempty"`      // [Optional] <p>Operating hours for Wednesday: You can skip this information if you want to mark the day as closed.</p>
	Thrusday      *Monday `json:"thrusday,omitempty"`       // [Optional] <p>Operating hours for Thursday: You can skip this information if you want to mark the day as closed.</p>
	Friday        *Monday `json:"friday,omitempty"`         // [Optional] <p>Operating hours for Friday: You can skip this information if you want to mark the day as closed.</p>
	Saturday      *Monday `json:"saturday,omitempty"`       // [Optional] <p>Operating hours for Saturday: You can skip this information if you want to mark the day as closed.</p>
	Sunday        *Monday `json:"sunday,omitempty"`         // [Optional] <p>Operating hours for Sunday: You can skip this information if you want to mark the day as closed.</p>
	PublicHoliday *Monday `json:"public_holiday,omitempty"` // [Optional] <p>Operating hours for public holiday: You can skip this information if you want to mark the day as closed.</p>
}

type ItemMaxDimension struct {
	Height       float64 `json:"height"`        // [Required] The max height limit.
	Width        float64 `json:"width"`         // [Required] The max width limit.
	Length       float64 `json:"length"`        // [Required] The max length limit.
	Unit         string  `json:"unit"`          // [Required] The unit for the limit.
	DimensionSum float64 `json:"dimension_sum"` // [Required] The sum of the item's dimension
}

type LogisticsCapability struct {
	SellerLogistics bool `json:"seller_logistics"` // [Required] <p>Indicate If it's a Seller logistics channel, if it's a Seller logistics channel will return true, otherwise it will return false.<br /></p>
}

type LogisticsChannel struct {
	LogisticsChannelId             int64                        `json:"logistics_channel_id"`              // [Required] The identity of logistic channel.
	LogisticsChannelName           string                       `json:"logistics_channel_name"`            // [Required] The name of logistic channel.
	CodEnabled                     bool                         `json:"cod_enabled"`                       // [Required] This is to indicate whether this logistic channel supports COD
	Enabled                        bool                         `json:"enabled"`                           // [Required] Whether this logistic channel is enabled on shop level.
	FeeType                        string                       `json:"fee_type"`                          // [Required] <p>SIZE_SELECTION</p><p>SIZE_INPUT</p><p>FIXED_DEFAULT_PRICE</p><p>CUSTOM_PRICE<br /></p>
	SizeList                       []Size                       `json:"size_list"`                         // [Required] Only for fee_type is SIZE_SELECTION
	WeightLimit                    *LogisticsChannelWeightLimit `json:"weight_limit"`                      // [Required] The weight limit for this logistic channel.
	ItemMaxDimension               *ItemMaxDimension            `json:"item_max_dimension"`                // [Required] The dimension limit for this logistic channel.
	VolumeLimit                    *VolumeLimit                 `json:"volume_limit"`                      // [Required] The limit of item volume.
	LogisticsDescription           string                       `json:"logistics_description"`             // [Required] For checkout channels, this field indicates its corresponding fulfillment channels.
	ForceEnable                    bool                         `json:"force_enable"`                      // [Required] Indicates whether the logistic channel is force enabled on Shop Level. If true, sellers cannot close this channel.
	MaskChannelId                  int64                        `json:"mask_channel_id"`                   // [Required] Indicate the parent logistic channel ID. If it’s 0, it indicates the channel is a checkout(masked) channel; if it’s not 0, indicate the channel is a fulfillment channel and has a checkout channel(checkout channel’s channel_id equals this mask_channel_id) on top of it. Multiple channels may share the same mask_channel_id.
	BlockSellerCoverShippingFee    bool                         `json:"block_seller_cover_shipping_fee"`   // [Required] <p>Indicate whether the channel is blocked to use seller cover shipping fee function.<br /></p><p>if the channel does not allow sellers to cover shipping fee, then the block_seller_cover_shipping_fee field will return true, otherwise it will return false.<br /></p>
	SupportCrossBorder             bool                         `json:"support_cross_border"`              // [Required] <p>Indicate whether this channel support cross border shipping.<br /></p>
	SellerLogisticHasConfiguration bool                         `json:"seller_logistic_has_configuration"` // [Required] <p>Indicate If seller has set the Seller logistics configuration if set will return true, otherwise it will return false or null.<br /></p>
	LogisticsCapability            *LogisticsCapability         `json:"logistics_capability"`              // [Required] <p>The capability of one logistic channel.</p>
	Preprint                       bool                         `json:"preprint"`                          // [Required] <p>Indicate whether this channel support pre-print AWB</p>
	ServiceTypeIdentifier          string                       `json:"service_type_identifier"`           // [Required] <p>This parameter specifies the delivery service type of logistics channel. Applicable values:&nbsp;</p><p>- instant</p><p>- same_day</p><p>- null</p>
}

type LogisticsChannelWeightLimit struct {
	ItemMaxWeight float64 `json:"item_max_weight"` // [Required] The max weight for an item on this logistic channel.If the value is 0 or null, that means there is no limit.
	ItemMinWeight float64 `json:"item_min_weight"` // [Required] The min weight for an item on this logistic channel. If the value is 0 or null, that means there is no limit.
}

type MassShipOrderRequest struct {
	LogisticsChannelId *int64                `json:"logistics_channel_id,omitempty"` // [Optional] <p><font color="#c24f4a">The API can only batch arrange shipment for multiple packages under the same product_location_id and same logistics_channel_id.&nbsp;</font></p><p><br /></p><p>Use this field to specify the logistics_channel_id for the request.&nbsp;<b>If not specified, will use the logistics_channel_id corresponds to the first package_number by default.</b></p>
	ProductLocationId  *string               `json:"product_location_id,omitempty"`  // [Optional] <p><font color="#c24f4a">The API can only batch arrange shipment for multiple packages under the same product_location_id and same logistics_channel_id.&nbsp;</font></p><p><br /></p><p>Use this field to specify the product_location_id for the request.&nbsp;<b>If not specified, will use the product_location_id corresponds to the first package_number by default.</b></p>
	PackageList        []RequestPackage      `json:"package_list"`                   // [Required] <p>The list of packages you want to arrange shipment. limit [1, 50].</p>
	Pickup             *RequestPickup        `json:"pickup,omitempty"`               // [Optional] <p>Required parameter ONLY if GetParameterForInit returns "pickup" or if GetLogisticsInfo returns "pickup" under "info_needed" for the same order. Developer should still include "pickup" field in the call even if "pickup" has empty value.</p>
	Dropoff            *RequestDropoff       `json:"dropoff,omitempty"`              // [Optional] Required parameter ONLY if GetParameterForInit returns "dropoff" or if GetLogisticsInfo returns "dropoff" under "info_needed" for the same order. Developer should still include "dropoff" field in the call even if "dropoff" has empty value. For logistic_id 80003 and 80004, both Regular and JOB shipping methods are supported. If you choose Regular shipping method, please use "tracking_no" to call Init API. If you choose JOB shipping method, please use "sender_real_name" to call Init API. Note that only one of "tracking_no" and "sender_real_name" can be selected.
	NonIntegrated      *RequestNonIntegrated `json:"non_integrated,omitempty"`       // [Optional] <p>Optional parameter when get_mass_shipping_parameter returns "non-integrated" under "info_needed".</p>
}

type MassShipOrderResponse struct {
	BaseResponse `json:",inline"` // Common response fields
	SuccessList  []RequestPackage `json:"success_list,omitempty"` // <p>Success package list.</p>
	FailList     []Fail           `json:"fail_list,omitempty"`    // <p>Fail package list.</p>
}

type Monday struct {
	StartTime string `json:"start_time"` // [Required] <p>Start time for Public Holiday</p>
	EndTime   string `json:"end_time"`   // [Required] <p>End time for Public Holiday</p>
}

type NonIntegrated struct {
	TrackingNumber *string `json:"tracking_number,omitempty"` // [Optional] Optional parameter for non-integrated channel order. The tracking number assigned by the shipping carrier for item shipment.
}

type OperatingHours struct {
	Date      string `json:"date"`       // [Required] <p>Date: it should include all date from start_date until end_date</p>
	StartTime string `json:"start_time"` // [Required] <p>Start time for that date</p><p>&lt;path&gt;&lt;/path&gt;<br /></p>
	EndTime   string `json:"end_time"`   // [Required] <p>End time for that date</p>
	Enable    bool   `json:"enable"`     // [Required] <p>True: If it is open on that date.</p><p>False: If it is closed on that date.</p>
}

type PackagingFee struct {
	Value float64 `json:"value"` // [Required] <p>The packaging fee price in the seller's local currency.</p>
}

type Pickup struct {
	AddressId      *int64  `json:"address_id,omitempty"`      // [Optional] The identity of address. Retrieved from shopee.logistics.GetAddress.
	PickupTimeId   *string `json:"pickup_time_id,omitempty"`  // [Optional] The pickup time id. Retrieved from shopee.logistics.GetTimeSlot.
	TrackingNumber *string `json:"tracking_number,omitempty"` // [Optional] Need input this field when "tracking_number" is returned from "info_need". Please note that this tracking number is assigned by third-party shipping carrier for item shipment.
}

type PickupAddress struct {
	AddressId    int64      `json:"address_id"`     // [Required] <p>The identity of address.<br /></p>
	Region       string     `json:"region"`         // [Required] <p>The region of specify address.<br /></p>
	State        string     `json:"state"`          // [Required] <p>The state of specify address.<br /></p>
	City         string     `json:"city"`           // [Required] <p>The city of specify address.<br /></p>
	District     string     `json:"district"`       // [Required] <p>The district of specify address.</p>
	Town         string     `json:"town"`           // [Required] <p>The town of specify address.<br /></p>
	Address      string     `json:"address"`        // [Required] <p>The address description of specify address.<br /></p>
	Zipcode      string     `json:"zipcode"`        // [Required] <p>The zipcode of specify address.<br /></p>
	AddressFlag  []string   `json:"address_flag"`   // [Required] <p>The flag of shop address, applicable values: default_address, pickup_address, return_address, current_address (Multi-Warehouse sellers only)<br /></p>
	TimeSlotList []TimeSlot `json:"time_slot_list"` // [Required] <p>List of pickup_time information corresponding to the address_id.<br /></p><p><br /></p><p>Some logistics channels may not return any date or time for pickup time slots. In such cases, sellers can arrange shipment without selecting any time slot, and Shopee will arrange a suitable timing for these situations.</p>
}

type RecipientAddressInfo struct {
	Key   string `json:"key"`             // [Required] <p>fields to query in the recipient address, should be&nbsp;name,&nbsp;phone,&nbsp;full_address,&nbsp;town,&nbsp;district,&nbsp;city,&nbsp;state,&nbsp;region,&nbsp;zipcode.<br /></p>
	Style *Style `json:"style,omitempty"` // [Optional] <p>image style<br /></p>
}

type RecipientSortCode struct {
	FirstRecipientSortCode  string `json:"first_recipient_sort_code"`  // [Required] <p>The first-level sort_code of recipient.<br /></p>
	SecondRecipientSortCode string `json:"second_recipient_sort_code"` // [Required] <p>The second-level sort_code of recipient.<br /></p>
	ThirdRecipientSortCode  string `json:"third_recipient_sort_code"`  // [Required] <p>The third-level sort_code of recipient.<br /></p>
}

type RegularOperatingHour struct {
	Monday        *Monday `json:"monday"`         // [Required] <p>The Operating hours for Monday</p>
	Tuesday       *Monday `json:"tuesday"`        // [Required] <p>The Operating hours for Tuesday</p>
	Wednesday     *Monday `json:"wednesday"`      // [Required] <p>The Operating hours for Wednesday</p>
	Thursday      *Monday `json:"thursday"`       // [Required] <p>The Operating hours for Thursday</p>
	Friday        *Monday `json:"friday"`         // [Required] <p>The Operating hours for Friday</p>
	Saturday      *Monday `json:"saturday"`       // [Required] <p>The Operating hours for Saturday</p>
	Sunday        *Monday `json:"sunday"`         // [Required] <p>The Operating hours for Sunday</p>
	PublicHoliday *Monday `json:"public_holiday"` // [Required] <p>The Operating hours for Public Holiday</p>
}

type RegularOperatingHourRestrictions struct {
	MinimumWorkingDaysInWeek int64             `json:"minimum_working_days_in_week"` // [Required] <p>Minimum number of days the seller needs to designate as working days per week (including Monday to Sunday, but excluding public holidays from the count).</p>
	WorkingDayConfig         *WorkingDayConfig `json:"working_day_config"`           // [Required] <p>The restrictions specific to each day</p>
}

type Repsonse struct {
	RegularOperatingHour        *RegularOperatingHour `json:"regular_operating_hour"`         // [Required] <p>The details of Pickup Operating Hours/Preferred Pickup Hours</p>
	InstantOperatingHour        *RegularOperatingHour `json:"instant_operating_hour"`         // [Required] <p>The details of Instant Operating Hours</p>
	SpecialOperatingHour        *SpecialOperatingHour `json:"special_operating_hour"`         // [Required] <p>The details of Special Operating Hours</p><p>&lt;path&gt;&lt;/path&gt;<br /></p>
	ShopCollectionOperatingHour *RegularOperatingHour `json:"shop_collection_operating_hour"` // [Required] <p>The details of Shop Collection Operating Hours</p>
}

type RequestBooking struct {
	BookingSn            string  `json:"booking_sn"`                       // [Required] <p>Shopee's unique identifier for a booking.<br /></p>
	TrackingNumber       *string `json:"tracking_number,omitempty"`        // [Optional] <p>The tracking number of booking. Required except for the channel allow print before arrange shipment.<br /></p>
	ShippingDocumentType *string `json:"shipping_document_type,omitempty"` // [Optional] <p>The type of shipping document. Available values: NORMAL_AIR_WAYBILL,THERMAL_AIR_WAYBILL,NORMAL_JOB_AIR_WAYBILL,THERMAL_JOB_AIR_WAYBILL<br /></p>
}

type RequestDimension struct {
	Length int64 `json:"length"` // [Required] <p>The length of the packaging in centimetres (cm).</p>
	Width  int64 `json:"width"`  // [Required] <p>The width of the packaging in centimetres (cm).</p>
	Height int64 `json:"height"` // [Required] <p>The height of the packaging in centimetres (cm).</p>
}

type RequestDropoff struct {
	BranchId       *int64  `json:"branch_id,omitempty"`        // [Optional] The identity of branch. Retrieved from shopee.logistics.GetBranch branch.
	SenderRealName *string `json:"sender_real_name,omitempty"` // [Optional] The real name of sender.
	TrackingNumber *string `json:"tracking_number,omitempty"`  // [Optional] Need input this field when "tracking_number" is returned from "info_need". Please note that this tracking number is assigned by third-party shipping carrier for item shipment.
}

type RequestNonIntegrated struct {
	TrackingNumber []TrackingNumber `json:"tracking_number,omitempty"` // [Optional] <p>Optional parameter for non-integrated channel order. The tracking number assigned by the shipping carrier for item shipment.</p>
}

type RequestOrder struct {
	OrderSn              string  `json:"order_sn"`                         // [Required] Shopee's unique identifier for an order.
	PackageNumber        *string `json:"package_number,omitempty"`         // [Optional] Shopee's unique identifier for the package under an order. You should't fill the field with empty string when there is not a package number.
	TrackingNumber       *string `json:"tracking_number,omitempty"`        // [Optional] The tracking number of order. Required except for the channel allow print before arrange shipment.
	ShippingDocumentType *string `json:"shipping_document_type,omitempty"` // [Optional] <p>The type of shipping document. Available values: NORMAL_AIR_WAYBILL, THERMAL_AIR_WAYBILL, NORMAL_JOB_AIR_WAYBILL, THERMAL_JOB_AIR_WAYBILL,&nbsp;THERMAL_UNPACKAGED_LABEL</p>
}

type RequestPackage struct {
	PackageNumber string `json:"package_number"` // [Required] <p>Shopee's unique identifier for the package under an order.</p>
}

type RequestPickup struct {
	AddressId    int64   `json:"address_id"`               // [Required] <p>The identity of address. Retrieved from v2.logistics.get_booking_shipping_parameter.<br /></p>
	PickupTimeId *string `json:"pickup_time_id,omitempty"` // [Optional] <p>The pickup time id. Retrieved from v2.logistics.get_shipping_booking_parameter, you can only select one from the time_slot_list.<br /></p><p><br /></p><p>Some logistics channels may not return any date or time for pickup time slots. In such cases, sellers can arrange shipment without selecting any time slot, and Shopee will arrange a suitable timing for these situations.</p>
}

type ResponseDataAddress struct {
	AddressId   int64    `json:"address_id"`   // [Required] The identity of address.
	Region      string   `json:"region"`       // [Required] The region of specify address.
	State       string   `json:"state"`        // [Required] The state of specify address.
	City        string   `json:"city"`         // [Required] The city of specify address.
	Address     string   `json:"address"`      // [Required] The address description of specify address.
	Zipcode     string   `json:"zipcode"`      // [Required] The zipcode of specify address.
	District    string   `json:"district"`     // [Required] The district of specify address.
	Town        string   `json:"town"`         // [Required] The town of specify address.
	AddressType []string `json:"address_type"` // [Required] The flag of shop address.Available values: DEFAULT_ADDRESS, PICK_UP_ADDRESS, RETURN_ADDRESS.
}

type ResponseDataDropoff struct {
	BranchList []Branch `json:"branch_list"` // [Required] List of available dropoff branches info.
	SlugList   []Slug   `json:"slug_list"`   // [Required]  List of available TW 3PL drop-off partners.
}

type ResponseDataFail struct {
	Id          string `json:"id"`           // [Required] <p>Package Number or Unpackaged SKU ID that failed in generating Shipping Document</p>
	FailError   string `json:"fail_error"`   // [Required] <p>Indicate error type if one element hit error.</p>
	FailMessage string `json:"fail_message"` // [Required] <p>Indicate error details if one element hit error.</p>
}

type ResponseDataInfoNeeded struct {
	Dropoff []string `json:"dropoff"` // [Required] <p>Could contain 'branch_id', 'sender_real_name' or 'tracking_no'. If it contains 'branch_id', choose one to Init. If it contains 'sender_real_name' or 'tracking_no', should manually input these values in Init API. If it has empty value, developer should still include "dropoff" field in Init API.<br /></p>
	Pickup  []string `json:"pickup"`  // [Required] <p>Could contain 'address_id' and 'pickup_time_id'. Choose one address_id and its corresponding pickup_time_id to Init. If it has empty value, developer should still include "pickup" field in Init API.It could contains "tracking_number" returned from "info_need"for some channels, please also add it when init.<br /></p>
}

type ResponseDataPickup struct {
	AddressList []PickupAddress `json:"address_list"` // [Required] <p>List of available pickup address info. For Multi-Warehouse sellers, note that changing pickup address from Current may incur higher shipping fees.<br /></p>
}

type ResponseDataPickupAddress struct {
	AddressId    int64             `json:"address_id"`     // [Required] <p>The identity of address.<br /></p>
	Region       string            `json:"region"`         // [Required] <p>The region of specify address.<br /></p>
	State        string            `json:"state"`          // [Required] <p>The state of specify address.<br /></p>
	City         string            `json:"city"`           // [Required] <p>The city of specify address.<br /></p>
	District     string            `json:"district"`       // [Required] <p>The district of specify address.<br /></p>
	Town         string            `json:"town"`           // [Required] <p>The town of specify address.<br /></p>
	Address      string            `json:"address"`        // [Required] <p>The address description of specify address.<br /></p>
	Zipcode      string            `json:"zipcode"`        // [Required] <p>The zipcode of specify address.<br /></p>
	AddressFlag  []string          `json:"address_flag"`   // [Required] <p>The flag of shop address, applicable values: default_address, pickup_address, return_address, current_address(only for multi-warehouse sellers)<br /></p>
	TimeSlotList []AddressTimeSlot `json:"time_slot_list"` // [Required] <p>List of pickup_time information corresponding to the address_id.<br /></p><p><br /></p><p>Some logistics channels may not return any date or time for pickup time slots. In such cases, sellers can arrange shipment without selecting any time slot, and Shopee will arrange a suitable timing for these situations.</p>
}

type ResponseDataRecipientAddressInfo struct {
	Key   string `json:"key"`   // [Required] <p>queried field in recipient address<br /></p>
	Image string `json:"image"` // [Required] <p>base64 encoded png data string<br /></p>
}

type ResponseDataResult struct {
	OrderSn                        string   `json:"order_sn"`                          // [Required] Shopee's unique identifier for an order.
	PackageNumber                  string   `json:"package_number"`                    // [Required] Shopee's unique identifier for the package under an order.
	SuggestShippingDocumentType    string   `json:"suggest_shipping_document_type"`    // [Required] The shipping document type Shopee suggests. If you don't select any shipping document type, Shopee will use this as default shipping document type.
	SelectableShippingDocumentType []string `json:"selectable_shipping_document_type"` // [Required] The shipping document type you can select of this order.
	FailError                      string   `json:"fail_error"`                        // [Required] Indicate error type if one element hit error.
	FailMessage                    string   `json:"fail_message"`                      // [Required] Indicate error details if one element hit error.
}

type ResponseDataShippingDocumentInfo struct {
	BookingWeight        int64              `json:"booking_weight"`        // [Required] <p>Use this field to indicate booking weight when calculate the shipping fee. The unit of weigh is gram.<br /></p>
	LogisticsChannelId   int64              `json:"logistics_channel_id"`  // [Required] <p>The identity of logistic channel.<br /></p>
	ShippingCarrier      string             `json:"shipping_carrier"`      // [Required] <p>The logistics service provider for the booking.<br /></p>
	RecipientSortCode    *RecipientSortCode `json:"recipient_sort_code"`   // [Required] <p>The sort_code of recipient.<br /></p>
	SenderSortCode       *SenderSortCode    `json:"sender_sort_code"`      // [Required] <p>The sort_code of sender.<br /></p>
	ReturnSortCode       *ReturnSortCode    `json:"return_sort_code"`      // [Required] <p>The sort code for 3PL doing RTS.<br /></p>
	TrackingNumber       string             `json:"tracking_number"`       // [Required] <p>The tracking number assigned by the shipping carrier for item shipment.<br /></p>
	PickupHub            string             `json:"pickup_hub"`            // [Required] <p>The name of pickup hub.<br /></p>
	DeliveryHub          string             `json:"delivery_hub"`          // [Required] <p>The name of delivery hub.<br /></p>
	DeliverArea          string             `json:"deliver_area"`          // [Required] <p>Zone name.<br /></p>
	EcBookingNo          string             `json:"ec_booking_no"`         // [Required] <p>The name of ec booing.<br /></p>
	CreateDateYmdSl      string             `json:"create_date_ymd_sl"`    // [Required] <p>The date of create shipment booking.<br /></p>
	ManufacturersName    string             `json:"manufacturers_name"`    // [Required] <p>The name of manufacturer.<br /></p>
	ManufacturersWebsite string             `json:"manufacturers_website"` // [Required] <p>The website of manufacturer.<br /></p>
	IsLmDgBool           int64              `json:"is_lm_dg_bool"`         // [Required] <p>Use this field to indicate booking contains dangerous goods or not. dg:1 non-dg:0<br /></p>
	SpxSubDistrict       string             `json:"spx_sub_district"`      // [Required] <p>The sub-district of recipient's address.<br /></p>
	SpxReceiveStation    *SpxReceiveStation `json:"spx_receive_station"`   // [Required] <p>The spx receive station.<br /></p>
	Zone                 string             `json:"zone"`                  // [Required] <p>The zone of this booking.<br /></p>
	ZoneCode             string             `json:"zone_code"`             // [Required] <p>Delivery Sub Zone.<br /></p>
	DestinationBaseCode  string             `json:"destination_base_code"` // [Required] <p>Distribution Center Code.<br /></p>
}

type ResultList struct {
	RegularOperatingHour        *ResultListRegularOperatingHour `json:"regular_operating_hour"`         // [Required] <p>The result of create/update regular_operating_hour.&nbsp;</p>
	SpecialOperatingHour        *ResultListRegularOperatingHour `json:"special_operating_hour"`         // [Required] <p>The result of create/update speicial_operating_hour.</p>
	InstantOperatingHour        *ResultListRegularOperatingHour `json:"instant_operating_hour"`         // [Required] <p>The result of create/update instant_operating_hour.</p>
	ShopCollectionOperatingHour *ResultListRegularOperatingHour `json:"shop_collection_operating_hour"` // [Required] <p>The result of create/update shop_collection_operating_hour.</p>
}

type ResultListRegularOperatingHour struct {
	Status      string `json:"status"`       // [Required] <p>The system will return "Failed" if there are any validation errors. Otherwise, it will return a blank response.</p>
	FailMessage string `json:"fail_message"` // [Required] <p>Fail reason</p>
}

type ReturnSortCode struct {
	ReturnFirstSortCode string `json:"return_first_sort_code"` // [Required] <p>The first-level sort code for 3PL doing RTS.<br /></p>
}

type ReversedTrackingInfo struct {
	UpdateTime  int64  `json:"update_time"` // [Required] <p>The time when the reversed logistics&nbsp;tracking info is updated.</p>
	Description string `json:"description"` // [Required] <p>The description of the reversed logistics tracking info.</p>
}

type SenderSortCode struct {
	FirstSenderSortCode  string `json:"first_sender_sort_code"`  // [Required] <p>The first-level sort_code of sender.<br /></p>
	SecondSenderSortCode string `json:"second_sender_sort_code"` // [Required] <p>The second-level sort_code of sender.<br /></p>
	ThirdSenderSortCode  string `json:"third_sender_sort_code"`  // [Required] <p>The third-level sort_code of sender.<br /></p>
}

type SetAddressConfigRequest struct {
	ShowPickupAddress *bool              `json:"show_pickup_address,omitempty"` // [Optional] Definite show pickup address or not.
	AddressTypeConfig *AddressTypeConfig `json:"address_type_config,omitempty"` // [Optional] The config of your shop addres.
}

type SetAddressConfigResponse struct {
	BaseResponse `json:",inline"` // Common response fields
}

type SetMartPackagingInfoRequest struct {
	Enable       bool              `json:"enable"`                  // [Required] <p>Indicates whether the seller has enabled or disabled the packaging fee configuration.</p><p><b>True:</b>&nbsp;The seller charges a packaging fee.</p><p><b>False:</b>&nbsp;The seller does not charge a packaging fee.</p>
	Dimension    *RequestDimension `json:"dimension,omitempty"`     // [Optional] <p>Required if&nbsp;enabled&nbsp;is set to&nbsp;True.</p>
	PackagingFee *PackagingFee     `json:"packaging_fee,omitempty"` // [Optional] <p>Required if&nbsp;enabled&nbsp;is set to&nbsp;True.</p>
}

type SetMartPackagingInfoResponse struct {
	BaseResponse `json:",inline"`                 // Common response fields
	Response     SetMartPackagingInfoResponseData `json:"response"` // Actual response data
}

type SetMartPackagingInfoResponseData struct {
	Enable       bool              `json:"enable"`        // [Required] <p>Indicates whether the seller has enabled or disabled the packaging fee configuration.</p><p><b>True:</b>&nbsp;The seller charges a packaging fee.</p><p><b>False:</b>&nbsp;The seller does not charge a packaging fee.</p>
	Dimension    *RequestDimension `json:"dimension"`     // [Required] <p>Returned only if&nbsp;enabled&nbsp;is set to&nbsp;True.</p>
	PackagingFee *PackagingFee     `json:"packaging_fee"` // [Required] <p>Returned only if&nbsp;enabled&nbsp;is set to&nbsp;True.</p>
}

type ShipBookingRequest struct {
	BookingSn string         `json:"booking_sn"`        // [Required] <p>Shopee's unique identifier for a booking.<br /></p>
	Pickup    *RequestPickup `json:"pickup,omitempty"`  // [Optional] <p>Required parameter ONLY if get_shipping_parameter returns "pickup" under "info_needed". Developer should still include "pickup" field in the call even if "pickup" has empty value.<br /></p>
	Dropoff   *interface{}   `json:"dropoff,omitempty"` // [Optional] <p>Required parameter ONLY if get_shipping_parameter returns "dropoff" under "info_needed". Developer should still include "dropoff" field in the call even if "dropoff" has empty value. If you choose Regular shipping method, please use "tracking_no" to call Init API. If you choose JOB shipping method, please use "sender_real_name" to call Init API. Note that only one of "tracking_no" and "sender_real_name" can be selected.<br /></p>
}

type ShipBookingResponse struct {
	BaseResponse `json:",inline"` // Common response fields
}

type ShipOrderRequest struct {
	OrderSn       string         `json:"order_sn"`                 // [Required] Shopee's unique identifier for an order.
	PackageNumber *string        `json:"package_number,omitempty"` // [Optional] Shopee's unique identifier for the package under an order. You should't fill the field with empty string when there is't a package number.
	Pickup        *Pickup        `json:"pickup,omitempty"`         // [Optional] Required parameter ONLY if get_shipping_parameter returns "pickup" under "info_needed". Developer should still include "pickup" field in the call even if "pickup" has empty value.
	Dropoff       *Dropoff       `json:"dropoff,omitempty"`        // [Optional] Required parameter ONLY if get_shipping_parameter returns "dropoff" under "info_needed". Developer should still include "dropoff" field in the call even if "dropoff" has empty value. For logistic_id 80003 and 80004, both Regular and JOB shipping methods are supported. If you choose Regular shipping method, please use "tracking_no" to call Init API. If you choose JOB shipping method, please use "sender_real_name" to call Init API. Note that only one of "tracking_no" and "sender_real_name" can be selected.
	NonIntegrated *NonIntegrated `json:"non_integrated,omitempty"` // [Optional] Optional parameter when get_shipping_parameter returns "non-integrated" under "info_needed".
}

type ShipOrderResponse struct {
	BaseResponse `json:",inline"` // Common response fields
}

type ShippingDocumentInfo struct {
	Cod                       bool                    `json:"cod"`                           // [Required] <p>This value indicates whether the order was a COD (cash on delivery) order.<br /></p>
	CodAmount                 string                  `json:"cod_amount"`                    // [Required] <p>Use this field to indicate cod amount.<br /></p>
	OrderWeight               int64                   `json:"order_weight"`                  // [Required] <p>Use this field to indicate order weight when calculate the shipping fee. The unit of weigh is gram.<br /></p>
	LogisticsChannelId        int64                   `json:"logistics_channel_id"`          // [Required] <p>The identity of logistic channel.<br /></p>
	ShippingCarrier           string                  `json:"shipping_carrier"`              // [Required] <p>The logistics service provider that the buyer selected for the order to deliver items.<br /></p>
	ServiceCode               string                  `json:"service_code"`                  // [Required] <p>Only work for cross-border order. This code is required at some sorting hub. Please ensure the service_code is INCLUDED on your shipping label, otherwise the parcel cannot be processed by the warehouse. If you didn't retrieve service_code after you first called this API, please try few more times within 30 minutes.<br /></p>
	FirstMileName             string                  `json:"first_mile_name"`               // [Required] <p>Only work for cross-border order.The name of the carrier ships cross country or region.<br /></p>
	LastMileName              string                  `json:"last_mile_name"`                // [Required] <p>Only work for cross-border order.The name of the carrier delivers the parcels in local country or region.<br /></p>
	GoodsToDeclare            bool                    `json:"goods_to_declare"`              // [Required] <p>Only work for cross-border order.This value indicates whether the order contains goods that are required to declare at customs. "T" means true and it will mark as "T" on the shipping label; "F" means false and it will mark as "P" on the shipping label. This value is accurate ONLY AFTER the order trackingNo is generated, please capture this value AFTER your retrieve the trackingNo.<br /></p>
	LaneCode                  string                  `json:"lane_code"`                     // [Required] <p>Only work for cross-border order. The string use for waybill printing. The format is "S - region_code and lane_number". For example, S-TH01, S-TH02<br /></p>
	WarehouseAddress          string                  `json:"warehouse_address"`             // [Required] <p>Only work for cross-border order in some special shop. The address info of the warehouse.<br /></p>
	WarehouseId               string                  `json:"warehouse_id"`                  // [Required] <p>Only work for cross-border order in some special shop. The ID of the warehouse.<br /></p>
	RecipientSortCode         *RecipientSortCode      `json:"recipient_sort_code"`           // [Required] <p>The sort_code of recipient.<br /></p>
	SenderSortCode            *SenderSortCode         `json:"sender_sort_code"`              // [Required] <p>The sort_code of sender.<br /></p>
	ReturnSortCode            *ReturnSortCode         `json:"return_sort_code"`              // [Required] <p>The sort code for 3PL doing RTS.<br /></p>
	ThirdPartyLogisticInfo    *ThirdPartyLogisticInfo `json:"third_party_logistic_info"`     // [Required] <p>Only used for TW sellers.<br /></p>
	TrackingNumber            string                  `json:"tracking_number"`               // [Required] <p>The tracking number assigned by the shipping carrier for item shipment.<br /></p>
	ShopeeTrackingNumber      string                  `json:"shopee_tracking_number"`        // [Required] <p>First mile tracking NO. for CrossBoard BR seller can be used to self-design CB Brazil AWB.<br /></p>
	LastMileTrackingNumber    string                  `json:"last_mile_tracking_number"`     // [Required] <p>The last-mile tracking number. Only for Cross Board BR seller.<br /></p>
	PickupHub                 string                  `json:"pickup_hub"`                    // [Required] <p>The name of pickup hub.<br /></p>
	DeliveryHub               string                  `json:"delivery_hub"`                  // [Required] <p>The name of delivery hub.<br /></p>
	DeliverArea               string                  `json:"deliver_area"`                  // [Required] <p>Zone name.<br /></p>
	EcOrderNo                 string                  `json:"ec_order_no"`                   // [Required] <p>The name of ec order.<br /></p>
	CreateDateYmdSl           string                  `json:"create_date_ymd_sl"`            // [Required] <p>The date of create shipment order.<br /></p>
	ManufacturersName         string                  `json:"manufacturers_name"`            // [Required] <p>The name of manufacturer.<br /></p>
	ManufacturersWebsite      string                  `json:"manufacturers_website"`         // [Required] <p>The website of manufacturer.<br /></p>
	IsLmDgBool                int64                   `json:"is_lm_dg_bool"`                 // [Required] <p>Use this field to indicate order contains dangerous goods or not. dg:1 non-dg:0<br /></p>
	PreferredDeliveryOption   int64                   `json:"preferred_delivery_option"`     // [Required] <p>Use this field to indicate delivery address is residential or office address.<br /></p><p>0: not configured</p><p>1: office address</p><p>2:&nbsp;residential address</p>
	SpxSubDistrict            string                  `json:"spx_sub_district"`              // [Required] <p>The sub-district of recipient's address.<br /></p>
	SpxReceiveStation         *SpxReceiveStation      `json:"spx_receive_station"`           // [Required] <p>The spx receive station.<br /></p>
	Zone                      string                  `json:"zone"`                          // [Required] <p>The zone of this order.<br /></p>
	ZoneCode                  string                  `json:"zone_code"`                     // [Required] <p>Delivery Sub Zone.<br /></p>
	DestinationBaseCode       string                  `json:"destination_base_code"`         // [Required] <p>Distribution Center Code.<br /></p>
	LastThirdDigitsBuyerPhone string                  `json:"last_third_digits_buyer_phone"` // [Required] <p>Use this field indicates buyer phone number (last 3 digits). For non-TW local sellers<br /></p>
	ParcelSize                string                  `json:"parcel_size"`                   // [Required] <p>corresponding locker sizing for self-collection locker channels [only available for specific logistic channels: 148003 and 140006]<br /></p>
	Sod                       bool                    `json:"sod"`                           // [Required] <p>this value indicates whether the buyer select "scan on delivery" payment channel at checkout.<br /></p>
	BuyerCpfId                string                  `json:"buyer_cpf_id"`                  // [Required] <p>Buyer's CPF number for taxation and invoice purposes. Only for Brazil order.<br /></p>
	MutualCheck               int64                   `json:"mutual_check"`                  // [Required] <p>only apply for ID/VN shops.</p><p>mutual_check indicates whether the parcel is eligible for Return-on-the-Spot (RoS) co-check.&nbsp;</p><p><br /></p><p>If mutual_check=1, then the parcel is RoS eligible, where drivers and buyers can co-check the parcel. Buyer can then choose to accept or reject the parcel on the spot.</p><p><br /></p><p>If mutual_check=0, then the parcel is ineligible for RoS.<br /></p>
	DelyFriLabel              string                  `json:"dely_fri_label"`                // [Required] <p>Probability of Successful Friday Delivery.</p><p>The value of L(low), M(medium), H(high) represent the chances of successful delivery attempts on Friday.<br /></p>
	DelySatLabel              string                  `json:"dely_sat_label"`                // [Required] <p>Probability of Successful Saturday Delivery<br /></p><p>The value of L(low), M(medium), H(high) represent the chances of successful delivery attempts on Saturday.<br /></p>
	DelySunLabel              string                  `json:"dely_sun_label"`                // [Required] <p>Probability of Successful Sunday Delivery.<br /></p><p>The value of L(low), M(medium), H(high) represent the chances of successful delivery attempts on Sunday.<br /></p>
	PickupCode                string                  `json:"pickup_code"`                   // [Required] <p>For drivers to quickly identify parcel to be picked up.&nbsp;Only returned for ID and TH local orders which use instant+sameday for delivery.</p>
	SortingGroup              string                  `json:"sorting_group"`                 // [Required] <p>[Only for TW 30029 channel]&nbsp;This field indicate the sorting group value of the package. Available values:&nbsp;<br />- North<br />- South</p>
	UnpackagedSkuId           string                  `json:"unpackaged_sku_id"`             // [Required] <p>[Only for TW 30029 channel] Please refer to this number instead of tracking number for this this channel. This field will be empty for other channels.</p>
	UnpackagedSkuIdQrcode     string                  `json:"unpackaged_sku_id_qrcode"`      // [Required] <p>[Only for TW 30029 channel] Please refer to this field to generate the QR code for the shipping document for this channel.&nbsp;This field will be empty for other channels.</p>
}

type ShopCollectionOperatingHourRestrictions struct {
	MinimumWorkingDaysInWeek int64                                                    `json:"minimum_working_days_in_week"` // [Required] <p>Minimum number of days the seller needs to designate as working days per week (including Monday to Sunday, but excluding public holidays from the count).</p>
	WorkingDayConfig         *ShopCollectionOperatingHourRestrictionsWorkingDayConfig `json:"working_day_config"`           // [Required] <p>The restrictions specific to each day</p>
}

type ShopCollectionOperatingHourRestrictionsWorkingDayConfig struct {
	Monday        *WorkingDayConfigMonday `json:"monday"`         // [Required] <p>The restrictions specific for Monday</p>
	Tuesday       *Tuesday                `json:"tuesday"`        // [Required] <p>The restrictions specific for Tuesday</p>
	Wednesday     *WorkingDayConfigMonday `json:"wednesday"`      // [Required] <p>The restrictions specific for Wednesday</p>
	Thursday      *WorkingDayConfigMonday `json:"thursday"`       // [Required] <p>The restrictions specific for Thursday</p>
	Friday        *WorkingDayConfigMonday `json:"friday"`         // [Required] <p>The restrictions specific for Friday</p>
	Saturday      *WorkingDayConfigMonday `json:"saturday"`       // [Required] <p>The restrictions specific for Saturday</p>
	Sunday        *WorkingDayConfigMonday `json:"sunday"`         // [Required] <p>The restrictions specific for Sunday</p>
	PublicHoliday *WorkingDayConfigMonday `json:"public_holiday"` // [Required] <p>The restrictions specific for Public Holiday</p>
}

type Size struct {
	SizeId       string  `json:"size_id"`       // [Required] The identity of size.
	Name         string  `json:"name"`          // [Required] The name of size.
	DefaultPrice float64 `json:"default_price"` // [Required] The pre-defined shipping fee for the specific size.
}

type Slug struct {
	Slug     string `json:"slug"`      // [Required]  The identity of slug.
	SlugName string `json:"slug_name"` // [Required]  The name of slug.
}

type SpecialOperatingHour struct {
	Name           string           `json:"name"`            // [Required] <p>The name of Special Operating Hours</p>
	StartDate      string           `json:"start_date"`      // [Required] <p>The start date of special operating hours</p>
	EndDate        string           `json:"end_date"`        // [Required] <p>The end date of special operating hours</p>
	OperatingHours []OperatingHours `json:"operating_hours"` // [Required]
}

type SpecialOperatingHourRestrictions struct {
	SpecialDay *WorkingDayConfigMonday `json:"special_day"` // [Required]
}

type SpxReceiveStation struct {
	SpxFirstReceiveStation string `json:"spx_first_receive_station"` // [Required] <p>The first pickup station.<br /></p>
}

type Style struct {
	TextStyle  []string `json:"text_style,omitempty"`  // [Optional] <p>supports bold and italic<br /></p>
	FontSize   *int64   `json:"font_size,omitempty"`   // [Optional] <p>supports 1 - 108<br /></p>
	TextColor  *string  `json:"text_color,omitempty"`  // [Optional] <p>color string like "#AbCd12"<br /></p>
	ImageWidth *float64 `json:"image_width,omitempty"` // [Optional] <p>supports 0.1-30, in centimeters<br /></p>
	HAlign     *string  `json:"h_align,omitempty"`     // [Optional] <p>text horizontal align, supports left, center and right.<br /></p>
}

type ThirdPartyLogisticInfo struct {
	ServiceDescription            string                   `json:"service_description"`               // [Required] <p>Use this field to indicate the order category.<br /></p>
	Barcode                       string                   `json:"barcode"`                           // [Required] <p>The manufacturer barcode.<br /></p>
	PurchaseTime                  string                   `json:"purchase_time"`                     // [Required] <p>The purchase_time of the store.<br /></p>
	ReturnTime                    string                   `json:"return_time"`                       // [Required] <p>The return_time of the store.<br /></p>
	ManufacturersName             string                   `json:"manufacturers_name"`                // [Required] <p>The name of manufacturers.<br /></p>
	ManufacturersWebsite          string                   `json:"manufacturers_website"`             // [Required] <p>The website of manufacturers.<br /></p>
	RecipientArea                 string                   `json:"recipient_area"`                    // [Required] <p>The identification of recipient area.<br /></p>
	RouteStep                     string                   `json:"route_step"`                        // [Required] <p>The route code of the waybill.<br /></p>
	Suda5Code                     string                   `json:"suda5_code"`                        // [Required] <p>The tally code of the waybill.<br /></p>
	LargeLogisticsId              string                   `json:"large_logistics_id"`                // [Required] <p>The code of large logistics.<br /></p>
	ParentId                      string                   `json:"parent_id"`                         // [Required] <p>The parent code of the waybill.<br /></p>
	ReturnCycle                   string                   `json:"return_cycle"`                      // [Required] <p>Use this field to indicate the return cycle.<br /></p>
	ReturnMode                    string                   `json:"return_mode"`                       // [Required] <p>Use this field to indicate the return mode.<br /></p>
	Prompt                        string                   `json:"prompt"`                            // [Required] <p>The reminder of stork work.<br /></p>
	OrderSn                       string                   `json:"order_sn"`                          // [Required] <p>Shopee's unique identifier for an order.<br /></p>
	Qrcode                        string                   `json:"qrcode"`                            // [Required] <p>The QR code of the waybill.<br /></p>
	EcSupplierName                string                   `json:"ec_supplier_name"`                  // [Required] <p>The supplier name of channel.<br /></p>
	EcBarCode16                   string                   `json:"ec_bar_code16"`                     // [Required] <p>Use this field to indicate the first barcode.<br /></p>
	EquipmentId                   string                   `json:"equipment_id"`                      // [Required] <p>The device code.<br /></p>
	EshopId                       string                   `json:"eshop_id"`                          // [Required] <p>The child code for B2C Family-mart.<br /></p>
	EcBarCode9                    string                   `json:"ec_bar_code9"`                      // [Required] <p>Use this field to indicate the pick barcode.<br /></p>
	PelicanTrackingNo             string                   `json:"pelican_tracking_no"`               // [Required] <p>The tracking number of Shopee Delivery.<br /></p>
	PrintDate                     string                   `json:"print_date"`                        // [Required] <p>The date of printing the wayBill.<br /></p>
	Pzip                          string                   `json:"pzip"`                              // [Required] <p>The sort code of the order.<br /></p>
	PzipC                         string                   `json:"pzip_c"`                            // [Required] <p>The barcode of the sort code.<br /></p>
	DeliverAreaTxt                string                   `json:"deliver_area_txt"`                  // [Required] <p>The code of the delivery area.<br /></p>
	DeliverDateYmd                string                   `json:"deliver_date_ymd"`                  // [Required] <p>Expected delivery date of the order.<br /></p>
	SdDriverCode                  string                   `json:"sd_driver_code"`                    // [Required] <p>Lorry driver code of the order.<br /></p>
	MdDriverCode                  string                   `json:"md_driver_code"`                    // [Required] <p>Motorcycle driver code of the order.<br /></p>
	PutorderStackzoneCode         string                   `json:"putorder_stackzone_code"`           // [Required] <p>Stacking area of the order.<br /></p>
	CustomerCode                  string                   `json:"customer_code"`                     // [Required] <p>The customer code of Shopee.<br /></p>
	DeliverRouter                 string                   `json:"deliver_router"`                    // [Required] <p>Use this field to indicate the delivery router.<br /></p>
	StoreType                     string                   `json:"store_type"`                        // [Required] <p>Use this field to indicate the store type.<br /></p>
	PickRouter                    string                   `json:"pick_router"`                       // [Required] <p>Use this field to indicate the pick router.<br /></p>
	BarcodeDc                     string                   `json:"barcode_dc"`                        // [Required] <p>The main logistic barcode of the waybill.<br /></p>
	EcOrderNumber                 string                   `json:"ec_order_number"`                   // [Required] <p>Use this field to indicate the logistics order number.<br /></p>
	BarcodePr                     string                   `json:"barcode_pr"`                        // [Required] <p>The sorting barcode of the waybill.<br /></p>
	FirstPickBarcode              string                   `json:"first_pick_barcode"`                // [Required] <p>The first pick barcode of the waybill.<br /></p>
	SecondPickBarcode             string                   `json:"second_pick_barcode"`               // [Required] <p>The second pick barcode of the waybill.<br /></p>
	IsCodBool                     string                   `json:"is_cod_bool"`                       // [Required] <p>Use this field to indicate the service type.<br /></p>
	ReceiverName                  string                   `json:"receiver_name"`                     // [Required] <p>Use this field to indicate the receiver name.<br /></p>
	RcvStoreName                  string                   `json:"rcv_store_name"`                    // [Required] <p>Use this field to indicate the receiver store name.<br /></p>
	BranchCode                    string                   `json:"branch_code"`                       // [Required] <p>Use this field indicates destination service point code.<br /></p>
	BranchName                    string                   `json:"branch_name"`                       // [Required] <p>Use this field indicates destination service point name.<br /></p>
	LastThirdDigitsRecipientPhone string                   `json:"last_third_digits_recipient_phone"` // [Required] <p>Use this field indicates buyer phone number (last 3 digits).<br /></p>
	LastThirdDigitsSenderPhone    string                   `json:"last_third_digits_sender_phone"`    // [Required] <p>Use this field indicates seller phone number (last 3 digits).<br /></p>
	BarcodeNo1                    string                   `json:"barcode_no1"`                       // [Required] <p>First barcode no. sacnned when seller drop off<br /></p>
	BarcodeNo2                    string                   `json:"barcode_no2"`                       // [Required] <p>Second barcode no. sacnned when seller drop off<br /></p>
	PrintDatetime                 string                   `json:"print_datetime"`                    // [Required] <p>AWB Print date and time<br /></p>
	OkMidType                     string                   `json:"ok_mid_type"`                       // [Required] <p>Middle type used in OK Mart SOC<br /></p>
	OkAisleNo                     string                   `json:"ok_aisle_no"`                       // [Required] <p>Aisle no. used in OK Mart SOC<br /></p>
	OkGridNo                      string                   `json:"ok_grid_no"`                        // [Required] <p>Grid no used in OK Mart SOC<br /></p>
	OkTrackingNumber              string                   `json:"ok_tracking_number"`                // [Required] <p>The tracking number of OK Mart.</p>
	BarcodeNo3                    string                   `json:"barcode_no3"`                       // [Required] <p>OK SOC received no.<br /></p>
	ShipType                      string                   `json:"ship_type"`                         // [Required] <p>Ship type used by OK Mart<br /></p>
	Area                          string                   `json:"area"`                              // [Required] <p>The area of the collect OK branch used for OK sorting<br /></p>
	BarcodeNo4                    string                   `json:"barcode_no4"`                       // [Required] <p>First barcode no. sacnned when buyer collect<br /></p>
	BarcodeNo5                    string                   `json:"barcode_no5"`                       // [Required] <p>Second barcode no. sacnned when buyer collect<br /></p>
	TwLastThreeDigitsBuyerPhone   string                   `json:"tw_last_three_digits_buyer_phone"`  // [Required] <p>[Only for local TW orders]&nbsp;Last 3 digits of buyer's phone number, apply for channel_id: 30005, 30006, 30007,30014,30015<br /></p>
	TwStoreName                   string                   `json:"tw_store_name"`                     // [Required] <p>[Only for TW channel_id:30005 ]&nbsp;Store name for 7-ELEVEN orders.<br /></p>
	TwStoreNumber                 string                   `json:"tw_store_number"`                   // [Required] <p>[Only for TW channel_id:30005 ]Store number for 7-ELEVEN orders.<br /></p>
	BuyerPreferDeliveryTime       *BuyerPreferDeliveryTime `json:"buyer_prefer_delivery_time"`        // [Required] <p>[Only for TW channel:30017]&nbsp;The time buyer prefers to receive the packages.<br /></p>
}

type TimeSlot struct {
	Date         int64    `json:"date"`           // [Required] <p>The date of pickup time. In timestamp.<br /></p>
	TimeText     string   `json:"time_text"`      // [Required] <p>The text description of pickup time. Only applicable for certain channels.<br /></p>
	PickupTimeId string   `json:"pickup_time_id"` // [Required] <p>The identity of pickuptime.<br /></p>
	Flags        []string `json:"flags"`          // [Required] <p>This field will have the value <b><font color="#c24f4a">“recommended”</font></b> for the time slot that Shopee suggests sellers choose.&nbsp;</p><p><br /></p><p>While it is advisable for sellers to choose the recommended time slot, they can also choose other time slots that do not have the recommended flag.</p>
}

type TrackingInfo struct {
	UpdateTime      int64           `json:"update_time"`      // [Required] <p>The time when logistics info has been updated.<br /></p>
	Description     string          `json:"description"`      // [Required] <p>The description of booking logistics tracking info.logistics_status<br /></p>
	LogisticsStatus LogisticsStatus `json:"logistics_status"` // [Required] <p>The Shopee logistics status for the booking.&nbsp;</p><p>TrackingLogisticsStatus:<br /></p><p>INITIAL</p><p>ORDER_INIT</p><p>ORDER_SUBMITTED</p><p>ORDER_CREATED</p><p>PICKUP_REQUESTED</p><p>PICKUP_PENDING</p><p>PICKED_UP</p><p>DELIVERY_PENDING</p><p>DELIVERED</p><p>LOST</p><p>UPDATE</p><p>UPDATE_SUBMITTED</p><p>UPDATE_CREATED</p><p>RETURN_STARTED</p><p>RETURN_PENDING</p><p>CANCEL</p><p>CANCEL_CREATED</p><p>CANCELED</p><p>FAILED_ORDER_SUBMITTED</p><p>FAILED_ORDER_CREATED</p><p>FAILED_PICKUP_REQUESTED</p><p>FAILED_PICKED_UP</p><p>FAILED_DELIVERED</p><p>FAILED_UPDATE_SUBMITTED</p><p>FAILED_UPDATE_CREATED</p><p>FAILED_RETURNED</p><p>FAILED_CANCEL_CREATED</p><p>FAILED_CANCELED</p><p>RETURNED</p><p>RETURN_INTIATED</p>
}

type TrackingNumber struct {
	PackageNumber  string `json:"package_number"`  // [Required] <p>Shopee's unique identifier for the package under an order.</p>
	TrackingNumber string `json:"tracking_number"` // [Required] <p>Optional parameter for non-integrated channel order. The tracking number assigned by the shipping carrier for item shipment.</p>
}

type Tuesday struct {
	Mandatory              bool   `json:"mandatory"`                // [Required] <p>If the value is true, this day must be marked as open.</p>
	MinimumOperatingHour   int64  `json:"minimum_operating_hour"`   // [Required] <p>Minimum number of hours required for the seller to operate on that day.</p>
	MinimumStartTime       string `json:"minimum_start_time"`       // [Required] <p>The start hour for that day should not be set earlier than this time.</p>
	MaximumStartTime       string `json:"maximum_start_time"`       // [Required] <p>The start hour for that day should not be set later than this time.</p>
	MinimumEndTime         string `json:"minimum_end_time"`         // [Required] <p>The end hour for that day should not be set earlier than this time.</p>
	MaximumEndTime         string `json:"maximum_end_time"`         // [Required] <p>The end hour for that day should not be set later than this time.</p>
	Operating_24HourToggle string `json:"operating_24_hour_toggle"` // [Required] <p>If the toggle value is true, the user can set the&nbsp;start_time&nbsp;to&nbsp;00:00&nbsp;and the&nbsp;end_time&nbsp;to&nbsp;23:59&nbsp;to indicate that the shop is operating 24 hours a day.</p>
}

type UnpackagedSkuRequests struct {
	UnpackagedSkuId *string `json:"unpackaged_sku_id,omitempty"` // [Optional] <p>Unpackaged SKU ID of the model.</p>
	Quantity        *int64  `json:"quantity,omitempty"`          // [Optional] <p>Number of copies for the generated labels (maximum 600 total across all requested SKUs).</p>
}

type UnsupportWarehouse struct {
	WarehouseId   int64  `json:"warehouse_id"`   // [Required] <p>Unsupported warehouse ID<br /></p>
	WarehouseName string `json:"warehouse_name"` // [Required] <p>Unsupported warehouse name<br /></p>
}

type UpdateAddressRequest struct {
	AddressId int64   `json:"address_id"`         // [Required] <p>Unique identifier for the address. You can get the address_id via v2.logistics.get_address_list.</p>
	Region    *string `json:"region,omitempty"`   // [Optional] <p>The region of the address.</p><p><br /></p><p>Note: Do not allow to update the region of the address.</p>
	State     *string `json:"state,omitempty"`    // [Optional] <p>The state of the address.</p>
	City      *string `json:"city,omitempty"`     // [Optional] <p>The city of the address.</p>
	District  *string `json:"district,omitempty"` // [Optional] <p>The district of the address.</p>
	Town      *string `json:"town,omitempty"`     // [Optional] <p>The town of the address.</p>
	Address   *string `json:"address,omitempty"`  // [Optional] <p>The detailed address description of the address.</p>
	Zipcode   *string `json:"zipcode,omitempty"`  // [Optional] <p>The zipcode of the address.</p>
	Name      *string `json:"name,omitempty"`     // [Optional] <p>Recipient’s name at this address.</p>
	Phone     *string `json:"phone,omitempty"`    // [Optional] <p>Contact phone number for the recipient.</p>
	GeoInfo   *string `json:"geo_info,omitempty"` // [Optional] <p>Geolocation information for the address.&nbsp;</p><p><br /></p><p>Type: JSON string</p><p><br /></p><p>Note:&nbsp;</p><p>1) To clear existing geo info, <b>pass "" or {}</b>.</p><p>2) To keep existing geo info, <b>do not include this field</b>.</p><p>3) The JSON may include optional fields:</p><p>- <b>formattedAddress</b> (string):&nbsp;full formatted address.</p><p>- <b>region</b> (object) – contains <b>latitude</b> and <b>longitude</b> as floats.</p><p>- <b>user_verified</b> (boolean) – whether the geolocation is verified by the user.<br />- <b>user_adjusted</b> (boolean) – whether the geolocation was adjusted by the user.</p>
}

type UpdateAddressResponse struct {
	BaseResponse `json:",inline"` // Common response fields
}

type UpdateChannelRequest struct {
	LogisticsChannelId int64 `json:"logistics_channel_id"`  // [Required] The identity of logistic channel.
	Enabled            *bool `json:"enabled,omitempty"`     // [Optional] Whether to enable this logistic channel.
	CodEnabled         *bool `json:"cod_enabled,omitempty"` // [Optional] Whether to enable COD for this logistic channel. Only COD supported channels are applicable.
}

type UpdateChannelResponse struct {
	BaseResponse `json:",inline"`          // Common response fields
	Response     UpdateChannelResponseData `json:"response"` // Actual response data
}

type UpdateChannelResponseData struct {
	Enabled            bool              `json:"enabled"`              // [Required] Whether this logistic channel is enabled.
	CodEnabled         bool              `json:"cod_enabled"`          // [Required] Whether COD is enabled for this channel.
	LogisticsChannelId int64             `json:"logistics_channel_id"` // [Required] The identity of logistic channel.
	UpdatedChannels    []UpdatedChannels `json:"updated_channels"`     // [Required] <p>List of channels that are updated in the operation (inclusive of dependent logistics channels)<br /></p>
	IsMultiWarehouse   bool              `json:"is_multi_warehouse"`   // [Required]
}

type UpdateOperatingHoursRequest struct {
	RegularOperatingHour        *RegularOperatingHour `json:"regular_operating_hour,omitempty"`         // [Optional] <p>Details of Pickup Operating Hours / Preferred Pickup Hours: You can skip this parameter if you are not updating the Pickup Operating Hours / Preferred Pickup Hours<br /></p>
	SpecialOperatingHour        *SpecialOperatingHour `json:"special_operating_hour,omitempty"`         // [Optional] <p>Details of Special Operating Hours : You can skip this parameter if you are not creating Special Operating Hours or&nbsp;if you do not have access to create&nbsp;Special Operating Hours&nbsp;</p>
	InstantOperatingHour        *InstantOperatingHour `json:"instant_operating_hour,omitempty"`         // [Optional] <p>Details of Instant Operating Hours : You can skip this parameter if you are not creating/updating Instant Operating Hours or&nbsp;if you do not have access to create/update Instant Operating Hours</p>
	ShopCollectionOperatingHour *RegularOperatingHour `json:"shop_collection_operating_hour,omitempty"` // [Optional] <p>Details of Shop Collection Operating Hours : You can skip this parameter if you are not creating/updating Shop Collection Operating Hours or&nbsp;if you do not have access to create/update Shop Collection Operating Hours</p>
}

type UpdateOperatingHoursResponse struct {
	BaseResponse `json:",inline"`                 // Common response fields
	Response     UpdateOperatingHoursResponseData `json:"response"` // Actual response data
}

type UpdateOperatingHoursResponseData struct {
	ResultList *ResultList `json:"result_list"` // [Required]
}

type UpdateSelfCollectionOrderLogisticsRequest struct {
	PackageNumber                 string   `json:"package_number"`                   // [Required] <p>Shopee's unique identifier for the package under an order.</p>
	SelfCollectionLogisticsAction string   `json:"self_collection_logistics_action"` // [Required] <p>Order logistics action. available values:</p><p>- ready_for_collection</p><p>- order_collected</p>
	EpocImageList                 []string `json:"epoc_image_list,omitempty"`        // [Optional] <p>List of image_id for the proof that buyer already collected the order at the store.&nbsp;</p><p>Required when self_collection_logistics_action is order_collected. Max: 3.</p><p>You can call the v2.media.upload_image to upload image and get the image_id, for this scenario, please pass the business = 1 and scene = 1.</p>
}

type UpdateSelfCollectionOrderLogisticsResponse struct {
	BaseResponse `json:",inline"` // Common response fields
}

type UpdateShippingOrderRequest struct {
	OrderSn       string         `json:"order_sn"`                 // [Required] Shopee's unique identifier for an order.
	PackageNumber *string        `json:"package_number,omitempty"` // [Optional] Shopee's unique identifier for the package under an order. You should't fill the field with empty string when there is't a package number.
	Pickup        *RequestPickup `json:"pickup"`                   // [Required] Required parameter ONLY if GetParameterForInit returns "pickup" or if GetLogisticsInfo returns "pickup" under "info_needed" for the same order. Developer should still include "pickup" field in the call even if "pickup" has empty value.
}

type UpdateShippingOrderResponse struct {
	BaseResponse `json:",inline"` // Common response fields
}

type UpdateTrackingStatusRequest struct {
	OrderSn         string          `json:"order_sn"`                  // [Required] <p>Shopee's unique identifier for an order.<br /></p>
	TrackingNumber  *string         `json:"tracking_number,omitempty"` // [Optional] <p>Order tracking number, might help seller to identify his order on the tracking_URL.</p><p>Can only be sent when updating logistics_status to "logistic_pickup_done".<br /></p>
	TrackingUrl     *string         `json:"tracking_url,omitempty"`    // [Optional] <p>Website's URL for order tracking&nbsp;with maximum length of 2048 characters.</p><p>Can only be sent when updating logistics_status to "logistic_pickup_done".<br /></p>
	LogisticsStatus LogisticsStatus `json:"logistics_status"`          // [Required] <p>Order status update support:</p><p>- logistics_pickup_done</p><p>- logistics_delivery_done</p><p>- logistics_delivery_failed<br /></p>
	FailedReason    *string         `json:"failed_reason,omitempty"`   // [Optional] <p>Only required when updating logistics_status to "logistics_delivery_failed". Only required for BR Instant Delivery channel (logistics_channel_id: 90026). Only accept the following values.&nbsp;</p><p>-&nbsp;buyer_unreachable</p><p>-&nbsp;buyer_unresponsive</p><p>-&nbsp;no_delivery_location_consensus</p>
}

type UpdateTrackingStatusResponse struct {
	BaseResponse `json:",inline"`                 // Common response fields
	Response     UpdateTrackingStatusResponseData `json:"response"` // Actual response data
}

type UpdateTrackingStatusResponseData struct {
	UpdateResult string `json:"update_result"` // [Required] <p>Update results:</p><p>- succeed</p><p>- failed</p>
}

type UpdatedChannels struct {
	ChannelId          int64                `json:"channel_id"`           // [Required] <p>Logistics channel ID<br /></p>
	ChannelDisplayName string               `json:"channel_display_name"` // [Required] <p>Logistics channel name<br /></p>
	UnsupportWarehouse []UnsupportWarehouse `json:"unsupport_warehouse"`  // [Required] <p>List details of unsupported warehouses<br /></p>
}

type UploadServiceablePolygonRequest struct {
	File interface{} `json:"file"` // [Required] <p>The .kml file to be uploaded to denote the serviceability area of the shops.</p><p><br /></p><p>Note: Please refer to&nbsp;<a href="https://open.shopee.com/faq/715" target="_blank">“KML file format for v2.logistics.upload_serviceable_polygon”</a>&nbsp;to understand the structure specifications and upload requirements for KML files.</p>
}

type UploadServiceablePolygonResponse struct {
	BaseResponse `json:",inline"`                     // Common response fields
	Response     UploadServiceablePolygonResponseData `json:"response"` // Actual response data
}

type UploadServiceablePolygonResponseData struct {
	TaskId string `json:"task_id"` // [Required] <p>Use the task_id to call v2.logistics.check_polygon_update_status to check if the upload job has been completed.</p>
}

type VolumeLimit struct {
	ItemMaxVolume float64 `json:"item_max_volume"` // [Required] The max volume for an item on this logistic channel.If the value is 0 or null, that means there is no limit for the item weight.
	ItemMinVolume float64 `json:"item_min_volume"` // [Required] The min volume for an item on this logistic channel. If the value is 0 or null, that means there is no limit for the item weight.
}

type WorkingDayConfig struct {
	Monday        *WorkingDayConfigMonday `json:"monday"`         // [Required] <p>The restrictions specific for Monday</p>
	Tuesday       *WorkingDayConfigMonday `json:"tuesday"`        // [Required] <p>The restrictions specific for Tuesday</p>
	Wednesday     *WorkingDayConfigMonday `json:"wednesday"`      // [Required] <p>The restrictions specific for Wednesday<br /></p>
	Thursday      *WorkingDayConfigMonday `json:"thursday"`       // [Required] <p>The restrictions specific for Thursday</p>
	Friday        *WorkingDayConfigMonday `json:"friday"`         // [Required] <p>The restrictions specific for Friday</p>
	Saturday      *WorkingDayConfigMonday `json:"saturday"`       // [Required] <p>The restrictions specific for Saturday</p>
	Sunday        *WorkingDayConfigMonday `json:"sunday"`         // [Required] <p>The restrictions specific for Sunday</p>
	PublicHoliday *WorkingDayConfigMonday `json:"public_holiday"` // [Required] <p>The restrictions specific for public holiday</p>
}

type WorkingDayConfigMonday struct {
	Mandatory              bool   `json:"mandatory"`                // [Required] <p>If the value is true, this day must be marked as open.</p>
	MinimumOperatingHour   int64  `json:"minimum_operating_hour"`   // [Required] <p>Minimum number of hours required for the seller to operate on that day.</p>
	MinimumStartTime       string `json:"minimum_start_time"`       // [Required] <p>The start hour for that day should not be set earlier than this time.</p>
	MaximumStartTime       string `json:"maximum_start_time"`       // [Required] <p>The start hour for that day should not be set later than this time.</p>
	MinimumEndTime         string `json:"minimum_end_time"`         // [Required] <p>The end hour for that day should not be set earlier than this time.</p>
	MaximumEndTime         string `json:"maximum_end_time"`         // [Required] <p>The end hour for that day should not be set later than this time.</p>
	Operating_24HourToggle bool   `json:"operating_24_hour_toggle"` // [Required] <p>If the toggle value is true, the user can set the&nbsp;start_time&nbsp;to&nbsp;00:00&nbsp;and the&nbsp;end_time&nbsp;to&nbsp;23:59&nbsp;to indicate that the shop is operating 24 hours a day.</p>
}
