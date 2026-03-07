package goshopee

type AddGlobalItemRequest struct {
	CategoryId      int64            `json:"category_id"`                // [Required] Category id of global item.
	GlobalItemName  string           `json:"global_item_name"`           // [Required] Name of global item.
	Description     string           `json:"description"`                // [Required] Description of global item.
	GlobalItemSku   *string          `json:"global_item_sku,omitempty"`  // [Optional] Sku of global item.
	Image           *PromotionImages `json:"image,omitempty"`            // [Optional] Image information of global item.
	OriginalPrice   float64          `json:"original_price"`             // [Required] Original price of global item.
	NormalStock     *int64           `json:"normal_stock,omitempty"`     // [Optional] Normal stock of global item.
	Weight          float64          `json:"weight"`                     // [Required] Weight of global item.
	Dimension       *Dimension       `json:"dimension,omitempty"`        // [Optional] Dimension information of global item.
	PreOrder        *RequestPreOrder `json:"pre_order"`                  // [Required] Preorder information of global item.
	Condition       *string          `json:"condition,omitempty"`        // [Optional] Condition of global item, "NEW" or "USED" is available.
	VideoUploadId   []string         `json:"video_upload_id,omitempty"`  // [Optional] Video upload id of global item. Only accept one video_upload_id at most.
	Brand           *Brand           `json:"brand,omitempty"`            // [Optional]
	AttributeList   []Attribute      `json:"attribute_list,omitempty"`   // [Optional] Item attributes.
	DescriptionInfo *DescriptionInfo `json:"description_info,omitempty"` // [Optional] New description field. New description field. Only whitelist sellers can use it. If you use the field, please upload the description_type=extended otherwise api will return error. If you don't use this field, you don't need to upload the description_type or upload description_type=normal
	DescriptionType *DescriptionType `json:"description_type,omitempty"` // [Optional] Values: See Data Definition- description_type (normal , extended). If you want to use extended_description, this field must be inputed
	SellerStock     []SellerStock    `json:"seller_stock,omitempty"`     // [Optional] <p>seller_stock&nbsp;of global item.<br /></p>
	DsCatRcmdId     *string          `json:"ds_cat_rcmd_id,omitempty"`   // [Optional] <p>category recommendation service id<br /></p>
	SizeChartInfo   *SizeChartInfo   `json:"size_chart_info,omitempty"`  // [Optional]
}

type AddGlobalItemResponse struct {
	BaseResponse `json:",inline"`          // Common response fields
	Response     AddGlobalItemResponseData `json:"response"` // Actual response data
}

type AddGlobalItemResponseData struct {
	GlobalItemId int64 `json:"global_item_id"` // [Required] Id of added global item.
}

type AddGlobalModelRequest struct {
	GlobalItemId int64         `json:"global_item_id"` // [Required] ID of global item.
	GlobalModel  []GlobalModel `json:"global_model"`   // [Required] Global model setting list. Limit is  [1,50].
}

type AddGlobalModelResponse struct {
	BaseResponse `json:",inline"` // Common response fields
}

type AttributeAttributeValue struct {
	ValueId int64 `json:"value_id"` // [Required] ID of attribute value.
}

type AttributeTreeAttributeInfo struct {
	InputType           int64    `json:"input_type"`            // [Required] <p>SINGLE_DROP_DOWN = 1</p><p>SINGLE_COMBO_BOX = 2</p><p>FREE_TEXT_FILED&nbsp; &nbsp; &nbsp; &nbsp; = 3</p><p>MULTI_DROP_DOWN&nbsp; &nbsp;= 4</p><p>MULTI_COMBO_BOX&nbsp; &nbsp;= 5<br /></p>
	InputValidationType int64    `json:"input_validation_type"` // [Required] <p>VALIDATOR_NO_VALIDATE_TYPE =&nbsp; 0</p><p>VALIDATOR_INT_TYPE = 1&nbsp;</p><p>VALIDATOR_STRING_TYPE = 2</p><p>VALIDATOR_FLOAT_TYPE = 3&nbsp;</p><p>VALIDATOR_DATE_TYPE = 4</p>
	FormatType          int64    `json:"format_type"`           // [Required] <p>FORMAT_NORMAL = 1</p><p>FORMAT_QUANTITATIVE_WITH_UNIT = 2<br /></p>
	DateFormatType      int64    `json:"date_format_type"`      // [Required] <p>YEAR_MONTH_DATE = 0 (DD/MM/YYYY)</p><p>YEAR_MONTH = 1 (MM/YYYY)<br /></p>
	AttributeUnitList   []string `json:"attribute_unit_list"`   // [Required] <p>Attribute's available units list<br /></p>
	MandatoryRegion     []string `json:"mandatory_region"`      // [Required] <p>Attribute is mandatory for these regions</p>
	MaxValueCount       int64    `json:"max_value_count"`       // [Required] <p>Max selected value count<br /></p>
	Introduction        string   `json:"introduction"`          // [Required] <p>introduction of special attribute<br /></p>
	IsOem               bool     `json:"is_oem"`                // [Required]
	SupportSearchValue  bool     `json:"support_search_value"`  // [Required] <p>Indicates whether this attribute has searchable values.</p><p>If yes, please call v2.global_product.search_global_attribute_value_list&nbsp;to get the default values</p>
}

type AttributeValue struct {
	ValueId           int64  `json:"value_id"`            // [Required] Unique identifier for value of this item attribute.
	OriginalValueName string `json:"original_value_name"` // [Required] Value name of this item attribute.
	ValueUnit         string `json:"value_unit"`          // [Required] Value unit of this item attribute.
}

type Attributes struct {
	AttributeId           int64            `json:"attribute_id"`            // [Required] The Identify of each category.
	OriginalAttributeName string           `json:"original_attribute_name"` // [Required] The name of each attribute.
	AttributeValueList    []AttributeValue `json:"attribute_value_list"`    // [Required]
}

type Brand struct {
	BrandId           int64  `json:"brand_id"`            // [Required] Id of brand.
	OriginalBrandName string `json:"original_brand_name"` // [Required] Original name of brand.
}

type CancelOrderRequestItem struct {
	ItemId  int64 `json:"item_id"`  // [Required] Failed item id.
	ModelId int64 `json:"model_id"` // [Required] Failed model id.
}

type Category struct {
	CategoryId           int64  `json:"category_id"`            // [Required] <p>ID for category.<br /></p>
	ParentCategoryId     int64  `json:"parent_category_id"`     // [Required] <p>ID for parent category.<br /></p>
	OriginalCategoryName string `json:"original_category_name"` // [Required] <p>English category name.<br /></p>
	DisplayCategoryName  string `json:"display_category_name"`  // [Required] <p>Display category name, it depends on what language you have uploaded<br /></p>
	HasChildren          bool   `json:"has_children"`           // [Required] <p>Whether this category has active children category.<br /></p>
}

type Column struct {
	Measurement          *Measurement       `json:"measurement"`            // [Required] <p>this is the column header which means a kind of measurement<br /></p>
	MeasurementValueList []MeasurementValue `json:"measurement_value_list"` // [Required] <p>the list of measurement value<br /></p>
}

type CreatePublishTaskRequest struct {
	GlobalItemId int64        `json:"global_item_id"` // [Required] Id of global item.
	ShopRegion   string       `json:"shop_region"`    // [Required] Region of shop.
	Item         *RequestItem `json:"item,omitempty"` // [Optional] Item information.
}

type CreatePublishTaskResponse struct {
	BaseResponse `json:",inline"`              // Common response fields
	Response     CreatePublishTaskResponseData `json:"response"` // Actual response data
}

type CreatePublishTaskResponseData struct {
	PublishTaskId int64 `json:"publish_task_id"` // [Required] The id of publish task.
}

type DeleteBundleDealItemRequestItem struct {
	ItemId int64 `json:"item_id"` // [Required] The id of related item failed to delete.
}

type DeleteGlobalItemRequest struct {
	GlobalItemId int64 `json:"global_item_id"` // [Required] The id of global item to delete.
}

type DeleteGlobalItemResponse struct {
	BaseResponse `json:",inline"`             // Common response fields
	Response     DeleteGlobalItemResponseData `json:"response"` // Actual response data
}

type DeleteGlobalItemResponseData struct {
	FailureDeleteItem []DeleteBundleDealItemRequestItem `json:"failure_delete_item"` // [Required] If delete failed, this field shows the details.
}

type DeleteGlobalModelRequest struct {
	GlobalItemId  int64 `json:"global_item_id"`  // [Required] Shopee's unique identifier for an global item.
	GlobalModelId int64 `json:"global_model_id"` // [Required] Shopee's unique identifier for an global model.
}

type DeleteGlobalModelResponse struct {
	BaseResponse `json:",inline"`              // Common response fields
	Response     DeleteGlobalModelResponseData `json:"response"` // Actual response data
}

type DeleteGlobalModelResponseData struct {
	GlobalModelId int64                    `json:"global_model_id"` // [Required] Global model id.
	Failures      []CancelOrderRequestItem `json:"failures"`        // [Required]
}

type DescriptionInfoExtendedDescription struct {
	FieldList []ExtendedDescriptionField `json:"field_list"` // [Required] Field of extended description
}

type Dimension struct {
	PackageHeight int64 `json:"package_height"` // [Required] <p>The height of package for this global model, the unit is CM.<br /></p>
	PackageLength int64 `json:"package_length"` // [Required] <p>The length of package for this global model, the unit is CM.<br /></p>
	PackageWidth  int64 `json:"package_width"`  // [Required] <p>The width of package for this global model, the unit is CM.<br /></p>
}

type DimensionLimit struct {
	DimensionMandatory bool `json:"dimension_mandatory"` // [Required] <p>dimension is mandatory or not for the category<br /></p>
}

type ExtendedDescriptionField struct {
	FieldType DescriptionElementFieldType `json:"field_type"` // [Required] Type of extended description field: values: See Data Definition- description_field_type (text , image).
	Text      string                      `json:"text"`       // [Required] If field_type is text, text information will be returned through this field.
	ImageInfo *FieldImageInfo             `json:"image_info"` // [Required] If field_type is image, image url will be returned through this field.
}

type ExtendedDescriptionLimit struct {
	DescriptionTextLengthMin       int64   `json:"description_text_length_min"`        // [Required] length min limit for item extended description text part
	DescriptionTextLengthMax       int64   `json:"description_text_length_max"`        // [Required] length max limit for item extended description text part
	DescriptionImageNumMin         int64   `json:"description_image_num_min"`          // [Required] length min limit for item extended description image num
	DescriptionImageNumMax         int64   `json:"description_image_num_max"`          // [Required] length max limit for item extended description image num
	DescriptionImageWidthMin       int64   `json:"description_image_width_min"`        // [Required] length min limit for item extended description image width
	DescriptionImageHeightMin      int64   `json:"description_image_height_min"`       // [Required] length min limit for item extended description image hight
	DescriptionImageAspectRatioMin float64 `json:"description_image_aspect_ratio_min"` // [Required] length min limit for item extended description image aspect (image width / image hight )
	DescriptionImageAspectRatioMax float64 `json:"description_image_aspect_ratio_max"` // [Required] length max limit for item extended description image aspect (image width / image hight )
}

type Failed struct {
	FailedReason string `json:"failed_reason"` // [Required] Failed reason.
}

type GetAttributeTreeResponseDataList struct {
	AttributeTree []ListAttributeTree `json:"attribute_tree"` // [Required] <p>One category's attribute trees<br /></p>
	CategoryId    int64               `json:"category_id"`    // [Required] <p>Category ID<br /></p>
	Warning       string              `json:"warning"`        // [Required] <p>Warning msg</p>
}

type GetGlobalItemIdRequest struct {
	ItemIdList []int64 `json:"item_id_list" url:"item_id_list"` // [Required] Item id list. Length limit is [1,20].
}

type GetGlobalItemIdResponse struct {
	BaseResponse `json:",inline"`            // Common response fields
	Response     GetGlobalItemIdResponseData `json:"response"` // Actual response data
}

type GetGlobalItemIdResponseData struct {
	ItemIdMap []ItemIdMap `json:"item_id_map"` // [Required]
}

type GetGlobalItemInfoRequest struct {
	GlobalItemIdList int64 `json:"global_item_id_list" url:"global_item_id_list"` // [Required] Global item id list. Length limit is [1,20].
}

type GetGlobalItemInfoResponse struct {
	BaseResponse `json:",inline"`              // Common response fields
	Response     GetGlobalItemInfoResponseData `json:"response"` // Actual response data
}

type GetGlobalItemInfoResponseData struct {
	GlobalItemList []ResponseDataGlobalItem `json:"global_item_list"` // [Required]
}

type GetGlobalItemLimitRequest struct {
	CategoryId *int64 `json:"category_id,omitempty" url:"category_id,omitempty"` // [Optional]
}

type GetGlobalItemLimitResponse struct {
	BaseResponse   `json:",inline"`               // Common response fields
	Response       GetGlobalItemLimitResponseData `json:"response"`                   // Actual response data
	SizeChartLimit *SizeChartLimit                `json:"size_chart_limit,omitempty"` //
}

type GetGlobalItemLimitResponseData struct {
	PriceLimit                       *PriceLimit                             `json:"price_limit"`                          // [Required]
	StockLimit                       *WholesalePriceThresholdPercentage      `json:"stock_limit"`                          // [Required]
	GlobalItemNameLengthLimit        *WholesalePriceThresholdPercentage      `json:"global_item_name_length_limit"`        // [Required]
	GlobalItemImageCountLimit        *WholesalePriceThresholdPercentage      `json:"global_item_image_count_limit"`        // [Required]
	GlobalItemDescriptionLengthLimit *WholesalePriceThresholdPercentage      `json:"global_item_description_length_limit"` // [Required]
	TierVariationNameLengthLimit     *WholesalePriceThresholdPercentage      `json:"tier_variation_name_length_limit"`     // [Required]
	TierVariationOptionLengthLimit   *WholesalePriceThresholdPercentage      `json:"tier_variation_option_length_limit"`   // [Required]
	TextLengthMultiplier             float64                                 `json:"text_length_multiplier"`               // [Required] Length ratio of Chinese characters to English characters in parameter verification. len(text)=len(Chinese characters)*text_length_multiplier+len(English characters )
	ExtendedDescriptionLimit         *ExtendedDescriptionLimit               `json:"extended_description_limit"`           // [Required]
	DtsLimit                         *GetGlobalItemLimitResponseDataDtsLimit `json:"dts_limit"`                            // [Required]
	WeightLimit                      *WeightLimit                            `json:"weight_limit"`                         // [Required]
	DimensionLimit                   *DimensionLimit                         `json:"dimension_limit"`                      // [Required]
}

type GetGlobalItemLimitResponseDataDtsLimit struct {
	DaysToShipRangeList []WholesalePriceThresholdPercentage `json:"days_to_ship_range_list"` // [Required] <p>Allowed limit scope for Pre order</p>
}

type GetGlobalItemListRequest struct {
	Offset         *string `json:"offset,omitempty" url:"offset,omitempty"`                     // [Optional] Specifies the starting entry of data to return in the current call. Default is null. if data is more than one page, the offset can be some entry to start next call.
	PageSize       int64   `json:"page_size" url:"page_size"`                                   // [Required] The size of one page. Limit is [1,50].
	UpdateTimeFrom *int64  `json:"update_time_from,omitempty" url:"update_time_from,omitempty"` // [Optional] The update_time_from and update_time_to fields specify a date range for retrieving orders (based on the item update time). The update_time_from field is the starting date range.
	UpdateTimeTo   *int64  `json:"update_time_to,omitempty" url:"update_time_to,omitempty"`     // [Optional] The update_time_from and update_time_to fields specify a date range for retrieving orders (based on the item update time). The update_time_to field is the ending date range
}

type GetGlobalItemListResponse struct {
	BaseResponse `json:",inline"`              // Common response fields
	Response     GetGlobalItemListResponseData `json:"response"` // Actual response data
}

type GetGlobalItemListResponseData struct {
	GlobalItemList []GlobalItem `json:"global_item_list"` // [Required]
	TotalCount     int64        `json:"total_count"`      // [Required] Total global item count.
	HasNextPage    bool         `json:"has_next_page"`    // [Required] This is to indicate whether the item list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of items.
	Offset         string       `json:"offset"`           // [Required] If has_next_page is true, this value need set to next request.offset.
}

type GetGlobalModelListRequest struct {
	GlobalItemId int64 `json:"global_item_id" url:"global_item_id"` // [Required] The id of global item.
}

type GetGlobalModelListResponse struct {
	BaseResponse `json:",inline"`               // Common response fields
	Response     GetGlobalModelListResponseData `json:"response"` // Actual response data
}

type GetGlobalModelListResponseData struct {
	TierVariation            []GetModelListResponseDataTierVariation `json:"tier_variation"`             // [Required] Tier variation information of global item.
	GlobalModel              []ResponseDataGlobalModel               `json:"global_model"`               // [Required] Global models.
	StandardiseTierVariation []ResponseDataStandardiseTierVariation  `json:"standardise_tier_variation"` // [Required] <p>Standardise Tier variation information of global item.<br /></p>
}

type GetLocalAdjustmentRateResponse struct {
	BaseResponse `json:",inline"`                   // Common response fields
	Response     GetLocalAdjustmentRateResponseData `json:"response"` // Actual response data
}

type GetLocalAdjustmentRateResponseData struct {
	LocalAdjustmentRate float64 `json:"local_adjustment_rate"` // [Required] <p>The multiplier used to adjust the cross-border original price to local price</p>
}

type GetModelListResponseDataTierVariation struct {
	Name       string                            `json:"name"`        // [Required] Tier name.
	OptionList []ResponseDataTierVariationOption `json:"option_list"` // [Required] Tier option list for corresponding tier name.
}

type GetPublishTaskResultRequest struct {
	PublishTaskId int64 `json:"publish_task_id" url:"publish_task_id"` // [Required] Id of publish task.
}

type GetPublishTaskResultResponse struct {
	BaseResponse `json:",inline"`                 // Common response fields
	Response     GetPublishTaskResultResponseData `json:"response"` // Actual response data
}

type GetPublishTaskResultResponseData struct {
	PublishStatus string                                   `json:"publish_status"` // [Required] Status of publish task.
	Success       *GetPublishTaskResultResponseDataSuccess `json:"success"`        // [Required] If publish task is successful, this field shows the published results.
	Failed        *Failed                                  `json:"failed"`         // [Required] If publish task is failed, this field shows the failed reason.
}

type GetPublishTaskResultResponseDataSuccess struct {
	Region string `json:"region"`  // [Required] The region of published item.
	ItemId string `json:"item_id"` // [Required] The id of published item.
}

type GetPublishableShopRequest struct {
	GlobalItemId int64   `json:"global_item_id" url:"global_item_id"`                 // [Required] Id of global item.
	ShopIdList   []int64 `json:"shop_id_list,omitempty" url:"shop_id_list,omitempty"` // [Optional] <p>Shop id list for checking if the shop is publishable.If not input the list, will return the first 300 publishable shop list in response<br /></p>
}

type GetPublishableShopResponse struct {
	BaseResponse `json:",inline"`               // Common response fields
	Response     GetPublishableShopResponseData `json:"response"` // Actual response data
}

type GetPublishableShopResponseData struct {
	PublishableShop []PublishableShop `json:"publishable_shop"` // [Required] Detail of publishable shops.
}

type GetPublishedListRequest struct {
	GlobalItemId int64   `json:"global_item_id" url:"global_item_id"`                 // [Required] Id of global item.
	ShopIdList   []int64 `json:"shop_id_list,omitempty" url:"shop_id_list,omitempty"` // [Optional] <p>Shop id list for checking if the shop is publishable.If not input the list, will return the first 300 publishable shop list in response after the&nbsp;migration period.<br /></p>
}

type GetPublishedListResponse struct {
	BaseResponse `json:",inline"`             // Common response fields
	Response     GetPublishedListResponseData `json:"response"` // Actual response data
}

type GetPublishedListResponseData struct {
	PublishedItem []PublishedItem `json:"published_item"` // [Required] Detail of published items.
}

type GetShopPublishableStatusRequest struct {
	GlobalItemId int64 `json:"global_item_id" url:"global_item_id"` // [Required] <p>Id of global item.<br /></p>
	Offset       int64 `json:"offset" url:"offset"`                 // [Required] <p>Specifies the starting entry of data to return in the current call. Default is 0. if data is more than one page, the offset can be some entry to start next call.<br /></p>
	PageSize     int64 `json:"page_size" url:"page_size"`           // [Required] <p>the size of one page.Max=100<br /></p>
}

type GetShopPublishableStatusResponse struct {
	BaseResponse `json:",inline"`                     // Common response fields
	Response     GetShopPublishableStatusResponseData `json:"response"` // Actual response data
}

type GetShopPublishableStatusResponseData struct {
	ShopPublishableStatusList []ShopPublishableStatus `json:"shop_publishable_status_list"` // [Required] <p>Detail of publishable shops.<br /></p>
	HasNextPage               bool                    `json:"has_next_page"`                // [Required] <p>This is to indicate whether the item list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of items.<br /></p>
	NextOffset                int64                   `json:"next_offset"`                  // [Required] <p>if has_next_page is true, this value need set to next request.offset<br /></p>
}

type GlobalItem struct {
	GlobalItemId int64 `json:"global_item_id"` // [Required] Shopee's unique identifier for an global item.
	UpdateTime   int64 `json:"update_time"`    // [Required] Timestamp that indicates the last time that there was a change in value of the item, such as price/stock change.
}

type GlobalItemPriceInfo struct {
	Currency           string  `json:"currency"`              // [Required] The three-digit code representing the currency unit used for the item in Shopee Listings.
	OriginalPrice      float64 `json:"original_price"`        // [Required] The original price of the item in the listing currency.
	SipItemPrice       float64 `json:"sip_item_price"`        // [Required] SIP item price.
	SipItemPriceSource string  `json:"sip_item_price_source"` // [Required] source of sip' price. ( auto or manual).
}

type GlobalModel struct {
	GlobalModelSku *string          `json:"global_model_sku,omitempty"` // [Optional] Sku of global model. model_sku length information needs to be no more than 100 characters.
	TierIndex      interface{}      `json:"tier_index"`                 // [Required] Tier index of global model.
	SellerStock    []SellerStock    `json:"seller_stock,omitempty"`     // [Optional] <p>seller_stock of global item<br /></p>
	OriginalPrice  float64          `json:"original_price"`             // [Required] Original price of global item.
	Weight         *float64         `json:"weight,omitempty"`           // [Optional] <p>The weight of this global model, the unit is KG.</p><p>If don't set the weight of this global model, will use the weight of global item by default.</p><p>If set the dimension of this global model, them must set the weight of this global model.</p>
	Dimension      *Dimension       `json:"dimension,omitempty"`        // [Optional] <p>The dimension of this global model.</p><p>If don't set the dimension of this global model, will use the dimension of global item by default.</p>
	PreOrder       *RequestPreOrder `json:"pre_order,omitempty"`        // [Optional] <p>Pre-order information of this global model.</p><p><br /></p><p>Notes:&nbsp;</p><p>If don't set the DTS of this global model, will use the DTS of the global item by default.</p>
}

type GlobalModelStockInfo struct {
	StockType       int64  `json:"stock_type"`        // [Required] Stock type. "1" means wms on hand, "2" means seller on hand.
	StockLocationId string `json:"stock_location_id"` // [Required] Stock location id.
	CurrentStock    int64  `json:"current_stock"`     // [Required] Current stock.
	NormalStock     int64  `json:"normal_stock"`      // [Required] Normal stock.
	ReservedStock   int64  `json:"reserved_stock"`    // [Required] Reserved stock.
}

type GlobalProductCategoryRecommendRequest struct {
	GlobalItemName          string  `json:"global_item_name" url:"global_item_name"`                                         // [Required] name of item
	GlobalProductCoverImage *string `json:"global_product_cover_image,omitempty" url:"global_product_cover_image,omitempty"` // [Optional] <p>Please use the image id returned by v2.media_space.upload_image api, we will ignore if this field is empty string<br /></p>
}

type GlobalProductCategoryRecommendResponse struct {
	BaseResponse `json:",inline"`                           // Common response fields
	Response     GlobalProductCategoryRecommendResponseData `json:"response"` // Actual response data
}

type GlobalProductCategoryRecommendResponseData struct {
	CategoryId []int64 `json:"category_id"` // [Required] Shopee's unique identifier for a category.
}

type GlobalProductGetAttributeTreeRequest struct {
	CategoryIdList []int64 `json:"category_id_list" url:"category_id_list"`     // [Required] <p>Max count is 20</p>
	Language       *string `json:"language,omitempty" url:"language,omitempty"` // [Optional] <p>Language</p><p>Support Lanuage:</p><p>"SG": [ "en", "zh-Hans", "ms" ],&nbsp;</p><p>"MY": [ "en", "zh-Hans", "ms" ],</p><p>"PH": [ "en", "zh-Hans" ],</p><p>"VN": [ "vn", "en" ],</p><p>"ID": [ "id", "en" ],</p><p>"TH": [ "th", "en" ],</p><p>"BR": [ "pt-BR", "en" ],</p><p>"MX": [ "es-MX", "en" ],</p><p>"CO": [ "es-CO", "en" ],</p><p>"CL": [ "es-CL", "en" ],</p><p>"TW": [ "zh-Hant", "zh-Hans", "en" ],</p><p>"IN": [ "en", "hi" ]</p>
}

type GlobalProductGetAttributeTreeResponse struct {
	BaseResponse `json:",inline"`                          // Common response fields
	Response     GlobalProductGetAttributeTreeResponseData `json:"response"` // Actual response data
}

type GlobalProductGetAttributeTreeResponseData struct {
	List []GetAttributeTreeResponseDataList `json:"list"` // [Required] <p>Each result corresponds to one category in category_ids<br /></p>
}

type GlobalProductGetBrandListRequest struct {
	Offset     int64 `json:"offset" url:"offset"`           // [Required] Specifies the starting entry of data to return in the current call. Default is 0. if data is more than one page, the offset can be some entry to start next call.
	PageSize   int64 `json:"page_size" url:"page_size"`     // [Required] the size of one page.
	CategoryId int64 `json:"category_id" url:"category_id"` // [Required] ID of category.
	Status     int64 `json:"status" url:"status"`           // [Required] Brand status , 1: normal brand, 2: pending brand.
}

type GlobalProductGetBrandListResponse struct {
	BaseResponse `json:",inline"`                      // Common response fields
	Response     GlobalProductGetBrandListResponseData `json:"response"` // Actual response data
}

type GlobalProductGetBrandListResponseData struct {
	BrandList   []ResponseDataBrand `json:"brand_list"`    // [Required]
	HasNextPage bool                `json:"has_next_page"` // [Required]  This is to indicate whether the item list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of items.
	NextOffset  int64               `json:"next_offset"`   // [Required] If has_next_page is true, this value need set to next request.offset
	IsMandatory bool                `json:"is_mandatory"`  // [Required] Whether is mandatory.
	InputType   string              `json:"input_type"`    // [Required] <p>Input type: DROP_DOWN</p>
}

type GlobalProductGetCategoryRequest struct {
	Language *string `json:"language,omitempty" url:"language,omitempty"` // [Optional] <p>Display language. Language should be one of "zh-hans", "en"</p>
}

type GlobalProductGetCategoryResponse struct {
	BaseResponse `json:",inline"`                     // Common response fields
	Response     GlobalProductGetCategoryResponseData `json:"response"` // Actual response data
}

type GlobalProductGetCategoryResponseData struct {
	CategoryList []Category `json:"category_list"` // [Required]
}

type GlobalProductGetRecommendAttributeRequest struct {
	GlobalItemName string  `json:"global_item_name" url:"global_item_name"`                 // [Required] Name of item.
	CategoryId     int64   `json:"category_id" url:"category_id"`                           // [Required] ID of category.
	CoverImageId   *string `json:"cover_image_id,omitempty" url:"cover_image_id,omitempty"` // [Optional] ID of image.
}

type GlobalProductGetRecommendAttributeResponse struct {
	BaseResponse `json:",inline"`                               // Common response fields
	Response     GlobalProductGetRecommendAttributeResponseData `json:"response"` // Actual response data
}

type GlobalProductGetRecommendAttributeResponseData struct {
	AttributeList []ResponseDataAttribute `json:"attribute_list"` // [Required] Attribute info list.
}

type GlobalProductGetSizeChartDetailRequest struct {
	SizeChartId int64   `json:"size_chart_id"`      // [Required]
	Language    *string `json:"language,omitempty"` // [Optional] <p>language should be in the list: ["en", "zh-Hans"]<br /></p>
}

type GlobalProductGetSizeChartDetailResponse struct {
	BaseResponse `json:",inline"`                            // Common response fields
	Response     GlobalProductGetSizeChartDetailResponseData `json:"response"` // Actual response data
}

type GlobalProductGetSizeChartDetailResponseData struct {
	SizeChartId    int64           `json:"size_chart_id"`    // [Required] <p>ID of new size chart<br /></p>
	SizeChartName  string          `json:"size_chart_name"`  // [Required] <p>name of new size chart<br /></p>
	SizeChartTable *SizeChartTable `json:"size_chart_table"` // [Required] <p>new size chart is a table format which include multiple columns. each column has column header (measurement) and multiple values (measurement value) of this column.<br /></p>
}

type GlobalProductGetSizeChartListRequest struct {
	CategoryId int64  `json:"category_id"` // [Required]
	PageSize   int64  `json:"page_size"`   // [Required]
	Cursor     string `json:"cursor"`      // [Required]
}

type GlobalProductGetSizeChartListResponse struct {
	BaseResponse `json:",inline"`                          // Common response fields
	Response     GlobalProductGetSizeChartListResponseData `json:"response"` // Actual response data
}

type GlobalProductGetSizeChartListResponseData struct {
	SizeChartList []SizeChart `json:"size_chart_list"` // [Required]
	TotalCount    int64       `json:"total_count"`     // [Required]
	NextCursor    string      `json:"next_cursor"`     // [Required]
}

type GlobalProductGetVariationsRequest struct {
	CategoryId int64 `json:"category_id" url:"category_id"` // [Required] <p>Leaf category id<br /></p>
}

type GlobalProductGetVariationsResponse struct {
	BaseResponse             `json:",inline"`       // Common response fields
	Data                     interface{}            `json:"data,omitempty"`                       //
	StandardiseVariationList []StandardiseVariation `json:"standardise_variation_list,omitempty"` //
}

type GlobalProductInitTierVariationRequest struct {
	GlobalModel              []GlobalModel              `json:"global_model"`                         // [Required] Model info list, model number at most 50
	GlobalItemId             int64                      `json:"global_item_id"`                       // [Required] ID of global item.
	StandardiseTierVariation []StandardiseTierVariation `json:"standardise_tier_variation,omitempty"` // [Optional] <p>There is at least one standardise_tier_variation and&nbsp;tier_variation.<br /></p><p><br /></p><p>If you want to update one tier/two tier to no tier, can just pass the tier_variation and standardise_tier_variation as [], and pass the global_model &gt;&gt; tier_index as [], meanwhile pass the original_price, seller_stock, etc., to set the price and stock for the modified product with no tier structure.<br /></p>
}

type GlobalProductInitTierVariationResponse struct {
	BaseResponse `json:",inline"` // Common response fields
}

type GlobalProductUpdatePriceRequest struct {
	GlobalItemId int64          `json:"global_item_id"` // [Required] ID of global item.
	PriceList    []RequestPrice `json:"price_list"`     // [Required] Price setting for global model. Limit is [1,50].
}

type GlobalProductUpdatePriceResponse struct {
	BaseResponse `json:",inline"` // Common response fields
}

type GlobalProductUpdateStockRequest struct {
	GlobalItemId int64          `json:"global_item_id"` // [Required] ID of global item.
	StockList    []RequestStock `json:"stock_list"`     // [Required] Stock setting for global model. Limit is [1,50].
}

type GlobalProductUpdateStockResponse struct {
	BaseResponse `json:",inline"` // Common response fields
}

type GlobalProductUpdateTierVariationRequest struct {
	GlobalItemId             int64                             `json:"global_item_id"`                       // [Required] ID of global item.
	ModelList                []UpdateTierVariationRequestModel `json:"model_list,omitempty"`                 // [Optional]
	StandardiseTierVariation []StandardiseTierVariation        `json:"standardise_tier_variation,omitempty"` // [Optional] <p>item standardise tier variation&nbsp;</p><p>There is at least one standardise_tier_variation and&nbsp;tier_variation<br /></p>
}

type GlobalProductUpdateTierVariationResponse struct {
	BaseResponse `json:",inline"` // Common response fields
}

type Images struct {
	ImageIdList  []string `json:"image_id_list"`  // [Required] List of image url.
	ImageUrlList []string `json:"image_url_list"` // [Required] List of image id.
}

type ItemIdMap struct {
	ItemId       int64 `json:"item_id"`        // [Required] Id of item.
	GlobalItemId int64 `json:"global_item_id"` // [Required] Id of global item.
}

type ItemModel struct {
	TierIndex     interface{} `json:"tier_index"`             // [Required] Tier index of model.
	OriginalPrice float64     `json:"original_price"`         // [Required] <p>Original price of model.&nbsp;If you upload this field, we will take your value, so you should pass the value in local currency, if you don't upload this field, Shopee will automatically calculate the price.</p>
	ModelStatus   *string     `json:"model_status,omitempty"` // [Optional] <p>can be&nbsp;"NORMAL" or "UNAVAILABLE". Normal models can be sold on the buyer's side, and UNAVAILABLE models cannot be sold on the buyer's side. If you do not upload this field, the model status will be considered as "NORMAL".</p>
}

type ListAttributeTree struct {
	AttributeId        int64                             `json:"attribute_id"`         // [Required] <p>Attribute ID<br /></p>
	Mandatory          bool                              `json:"mandatory"`            // [Required] <p>Is mandatory or not<br /></p>
	Name               string                            `json:"name"`                 // [Required] <p>Attribute Name<br /></p>
	AttributeValueList []ListAttributeTreeAttributeValue `json:"attribute_value_list"` // [Required] <p>All available values for this attribute<br /></p>
	AttributeInfo      *AttributeTreeAttributeInfo       `json:"attribute_info"`       // [Required] <p>Attribute extra info<br /></p>
	MultiLang          []MultiLang                       `json:"multi_lang"`           // [Required] <p>Translate result for attribute name display<br /></p>
}

type ListAttributeTreeAttributeValue struct {
	ValueId            int64         `json:"value_id"`             // [Required] <p>Value ID<br /></p>
	Name               string        `json:"name"`                 // [Required] <p>Value name<br /></p>
	ValueUnit          string        `json:"value_unit"`           // [Required] <p>Value unit<br /></p>
	ChildAttributeList []interface{} `json:"child_attribute_list"` // [Required] <p>Child attributes for the value of parent attribute<br />The structure content is the same as attribute_tree<br /></p>
	MultiLang          *MultiLang    `json:"multi_lang"`           // [Required] <p>Translate results for value name display<br /></p>
}

type Measurement struct {
	InputType   string `json:"input_type"`   // [Required] <p>there are 3 kinds of measurement type: Single Dropdown, Input Single Number, Input Range Number.<br /></p>
	Unit        string `json:"unit"`         // [Required] <p>the unit of this size measurement.<br /></p>
	DisplayName string `json:"display_name"` // [Required] <p>name of column header (measurement)<br /></p>
}

type MeasurementValue struct {
	Value    float64 `json:"value"`     // [Required] <p>if the input_type of measurement is single input number, measurement will have one value which is returned by this field.<br /></p>
	MinValue float64 `json:"min_value"` // [Required] <p>if the input_type of measurement is input range number, measurement will be a range which is returned by 2 fields: min_value and max_value.<br /></p>
	MaxValue float64 `json:"max_value"` // [Required] <p>if the input_type of measurement is input range number, measurement will be a range which is returned by 2 fields: min_value and max_value.<br /></p>
	Option   string  `json:"option"`    // [Required] <p>if the input_type of measurement is single dropdown, measurement will have one value which is returned by this field.<br /></p>
}

type ModelPriceInfo struct {
	OriginalPrice float64 `json:"original_price"` // [Required] Original price of global model.
}

type MultiLang struct {
	Language string `json:"language"` // [Required] <p>Language<br /></p>
	Value    string `json:"value"`    // [Required] <p>Translate result<br /></p>
}

type PriceLimit struct {
	MinLimit float64 `json:"min_limit"` // [Required] Global item price min limit.
	MaxLimit float64 `json:"max_limit"` // [Required] Global item price max limit.
}

type PromotionImages struct {
	ImageIdList []string `json:"image_id_list"` // [Required] Image id list of item.
}

type PublishableShop struct {
	ShopRegion string `json:"shop_region"` // [Required] Region of published shop.
}

type PublishedItem struct {
	ShopRegion string     `json:"shop_region"` // [Required] Region of shop.
	ItemId     int64      `json:"item_id"`     // [Required] Id of published item.
	ItemStatus ItemStatus `json:"item_status"` // [Required] <p>Status of published item.Applicable values: 0.DELETED(Item is deleted by seller himself),1.NORMAL, 2.BANNED,3.REVIEWING,4.INVALID(Shopee Admin deleted),5.INVALID_HIDE(Shopee Admin delete confirmed),6.BLACKLISTED(Offensive_hide),8.NORMAL_UNLIST</p>
}

type RequestBrand struct {
	BrandId *int64 `json:"brand_id,omitempty"` // [Optional] Id of brand.
}

type RequestGlobalModel struct {
	GlobalModelSku string           `json:"global_model_sku"`    // [Required] Sku of global model.
	GlobalModelId  int64            `json:"global_model_id"`     // [Required] ID of global model.
	Weight         *float64         `json:"weight,omitempty"`    // [Optional] <p>The weight of this global model, the unit is KG.</p><p>If don't set the weight of this global model, will use the weight of global item by default.</p><p>If set the dimension of this global model, them must set the weight of this global model.</p>
	Dimension      *Dimension       `json:"dimension,omitempty"` // [Optional] <p>The dimension of this global model.</p><p>If don't set the dimension of this global model, will use the dimension of global item by default.</p>
	PreOrder       *RequestPreOrder `json:"pre_order,omitempty"` // [Optional] <p>Pre-order information of this global model.</p><p><br /></p><p>Notes:&nbsp;</p><p>If don't set the DTS of this global model, will use the DTS of the global item by default.</p>
}

type RequestItem struct {
	ItemName                 *string                    `json:"item_name,omitempty"`                  // [Optional] <p>Name of item.&nbsp;If you upload this field, we will take your value, so you should pass the value in the local language, if you don't upload this field, Shopee will automatically translate your global product name into the local language.</p>
	Description              *string                    `json:"description,omitempty"`                // [Optional] <p>Description of item.&nbsp;If you upload this field, we will take your value, so you should pass the value in the local language, if you don't upload this field, Shopee will automatically translate your global product description into the local language.</p>
	ItemStatus               *ItemStatus                `json:"item_status,omitempty"`                // [Optional] Status of item.
	OriginalPrice            *float64                   `json:"original_price,omitempty"`             // [Optional] <p>Original price of item.</p><p><b><font color="#c24f4a">For&nbsp;SG/MY/BR/MX/PL/ES/AR seller:</font></b>&nbsp;Sellers can set the price with two decimal place,&nbsp;other regions can only set the price as an integer.&nbsp;If you upload this field, we will take your value, so you should pass the value in local currency, if you don't upload this field, Shopee will automatically calculate the price.<br /></p>
	Image                    *PromotionImages           `json:"image,omitempty"`                      // [Optional] Image information of item.
	Model                    []ItemModel                `json:"model,omitempty"`                      // [Optional] Model information of item.
	SizeChart                *string                    `json:"size_chart,omitempty"`                 // [Optional] <p>Size chart of item. Only support image_id for now</p>
	Logistic                 []LogisticInfo             `json:"logistic,omitempty"`                   // [Optional] Logistic information of item.
	PreOrder                 *PreOrder                  `json:"pre_order,omitempty"`                  // [Optional] Preorder information of item.
	DescriptionInfo          *DescriptionInfo           `json:"description_info,omitempty"`           // [Optional] <p>New description field. Only whitelist sellers can use it. If you use the field, please upload the description_type=extended otherwise api will return error. If you don't use this field, you don't need to upload the description_type or upload description_type=normal.&nbsp;If you upload this field, we will take your value, so you should pass the value in the local language, if you don't upload this field, Shopee will automatically translate your global product description into the local language.</p>
	StandardiseTierVariation []StandardiseTierVariation `json:"standardise_tier_variation,omitempty"` // [Optional]
}

type RequestPreOrder struct {
	DaysToShip int64 `json:"days_to_ship"` // [Required] <p>Days to ship.</p>
}

type RequestPrice struct {
	GlobalModelId *int64  `json:"global_model_id,omitempty"` // [Optional] ID of global model.
	OriginalPrice float64 `json:"original_price"`            // [Required] Original price of global item.
}

type RequestStock struct {
	GlobalModelId *int64        `json:"global_model_id,omitempty"` // [Optional] ID of global model.
	SellerStock   []SellerStock `json:"seller_stock,omitempty"`    // [Optional]
}

type ResponseDataAttribute struct {
	AttributeId        int64                     `json:"attribute_id"`         // [Required] ID of attribute.
	AttributeValueList []AttributeAttributeValue `json:"attribute_value_list"` // [Required] Value list of this attribute.
}

type ResponseDataBrand struct {
	BrandId           int64  `json:"brand_id"`            // [Required]  Id of brand.
	OriginalBrandName string `json:"original_brand_name"` // [Required] Original name of brand
	DisplayBrandName  string `json:"display_brand_name"`  // [Required]  Display name of brand
}

type ResponseDataDescriptionInfo struct {
	ExtendedDescription *DescriptionInfoExtendedDescription `json:"extended_description"` // [Required] If description_type is extended , Description information will be returned through this field.
}

type ResponseDataGlobalItem struct {
	GlobalItemId          int64                        `json:"global_item_id"`           // [Required] Shopee's unique identifier for an global item.
	GlobalItemName        string                       `json:"global_item_name"`         // [Required] Name of the global item.
	Description           string                       `json:"description"`              // [Required] Description of the global item.
	GlobalItemSku         string                       `json:"global_item_sku"`          // [Required] An global item SKU (stock keeping unit) is an identifier defined by a seller, sometimes called parent SKU. Item SKU can be assigned to an item in Shopee Listings.
	GlobalItemStatus      string                       `json:"global_item_status"`       // [Required] The current status of the item. You can only query global product with normal status, otherwise api will return error.
	CreateTime            int64                        `json:"create_time"`              // [Required] Timestamp that indicates the date and time that the global item was created.
	UpdateTime            int64                        `json:"update_time"`              // [Required] Timestamp that indicates the last time that there was a change in value of the global item.
	StockInfo             []StockInfo                  `json:"stock_info"`               // [Required] If the item has models, this field will not be returned, please get it through get_model_list api.
	PriceInfo             []GlobalItemPriceInfo        `json:"price_info"`               // [Required] If the item has models, price_info will not be returned. Please get the price of each model through the get_global_model_list api.
	Image                 *Images                      `json:"image"`                    // [Required]
	Weight                string                       `json:"weight"`                   // [Required] <p>The weight of this global item, the unit is KG.</p><p>If set the weight of global models under this item, will return the max weight of all global models during the switching period to ensure system compatibility, please switch to call v2.global_product.get_global_model_list to get the weight of models.</p>
	Dimension             *Dimension                   `json:"dimension"`                // [Required] <p>The dimension of this global item.</p><p>If set the dimension of global models under this global item, will return the dimension with largest volume calculated by height*length*width during the switching period to ensure system compatibility, please switch to call v2.global_product.get_global_model_list to get the dimension of models.</p>
	PreOrder              *RequestPreOrder             `json:"pre_order"`                // [Required] <p>If set the DTS of global models under this item, will return the max DTS of all global models during the switching period to ensure system compatibility, please switch to call v2.global_product.get_global_model_list to get the DTS of models.<br /></p>
	SizeChart             string                       `json:"size_chart"`               // [Required] Url of size chart image.
	Condition             string                       `json:"condition"`                // [Required] Is it second-hand.
	HasModel              bool                         `json:"has_model"`                // [Required] Does it contain model.
	Video                 *VideoInfo                   `json:"video"`                    // [Required]
	CategoryId            int64                        `json:"category_id"`              // [Required] Shopee's unique identifier for a category.
	Brand                 *Brand                       `json:"brand"`                    // [Required]
	AttributeList         []Attributes                 `json:"attribute_list"`           // [Required]
	DescriptionInfo       *ResponseDataDescriptionInfo `json:"description_info"`         // [Required] New description field.New description field. Only whitelist sellers can use it. If you use the field, please upload the description_type=extended otherwise api will return error. If you don't use this field, you don't need to upload the description_type or upload description_type=normal
	DescriptionType       DescriptionType              `json:"description_type"`         // [Required] Type of description : values: See Data Definition- description_type (normal , extended).
	IsFulfillmentByShopee bool                         `json:"is_fulfillment_by_shopee"` // [Required] <p>whether item is fulfillment by shopee</p>
	SizeChartId           int64                        `json:"size_chart_id"`            // [Required] <p>size_chart 模板ID</p>
}

type ResponseDataGlobalModel struct {
	GlobalModelId         int64                  `json:"global_model_id"`          // [Required] Id of global model.
	GlobalModelSku        string                 `json:"global_model_sku"`         // [Required] Sku of global model.
	PriceInfo             *ModelPriceInfo        `json:"price_info"`               // [Required] Price info of global model.
	StockInfo             []GlobalModelStockInfo `json:"stock_info"`               // [Required] Stock info of global model.
	TierIndex             interface{}            `json:"tier_index"`               // [Required] Tier index of global model.
	Weight                string                 `json:"weight"`                   // [Required] <p>The weight of this global model, the unit is KG.</p><p>If don't set the weight of this global model, will use the weight of global item by default.</p>
	Dimension             *Dimension             `json:"dimension"`                // [Required] <p>The dimension of this global model.</p><p>If don't set the dimension of this global model, will use the dimension of global item by default.</p>
	PreOrder              *RequestPreOrder       `json:"pre_order"`                // [Required] <p>Pre-order information of this global model.</p><p><br /></p><p>Notes:&nbsp;</p><p>If don't set the DTS of this global model, will use the DTS of the global item by default.</p>
	IsFulfillmentByShopee bool                   `json:"is_fulfillment_by_shopee"` // [Required] <p>If it it a FBS model</p>
}

type ResponseDataPageInfo struct {
	Cursor  int64 `json:"cursor"`   // [Required]
	HasNext bool  `json:"has_next"` // [Required]
}

type ResponseDataStandardiseTierVariation struct {
	VariationId         int64                                     `json:"variation_id"`          // [Required] <p>Standardise Tier variation ID.<br /></p>
	VariationName       string                                    `json:"variation_name"`        // [Required] <p>Standardise Tier variation Name.<br /></p>
	VariationGroupId    int64                                     `json:"variation_group_id"`    // [Required] <p>Standardise Tier variation Group ID.<br /></p>
	VariationOptionList []StandardiseTierVariationVariationOption `json:"variation_option_list"` // [Required] <p>Standardise Tier variation Options List.<br /></p>
}

type ResponseDataTierVariationOption struct {
	Option string          `json:"option"` // [Required] Tier option.
	Image  *FieldImageInfo `json:"image"`  // [Required] Image information of tier.
}

type SearchGlobalAttributeValueListRequest struct {
	AttributeId int64   `json:"attribute_id"`         // [Required]
	ValueName   *string `json:"value_name,omitempty"` // [Optional]
	Cursor      int64   `json:"cursor"`               // [Required]
	Limit       int64   `json:"limit"`                // [Required] <p>The range is 1 to 100</p>
}

type SearchGlobalAttributeValueListResponse struct {
	BaseResponse `json:",inline"`                           // Common response fields
	Msg          string                                     `json:"msg,omitempty"`           //
	DebugMessage string                                     `json:"debug_message,omitempty"` //
	Response     SearchGlobalAttributeValueListResponseData `json:"response"`                // Actual response data
}

type SearchGlobalAttributeValueListResponseData struct {
	ValueList []Value               `json:"value_list"` // [Required]
	PageInfo  *ResponseDataPageInfo `json:"page_info"`  // [Required]
}

type SetSyncFieldRequest struct {
	ShopSyncList []ShopSync `json:"shop_sync_list"` // [Required] Length limit is [1,50].
}

type SetSyncFieldResponse struct {
	BaseResponse `json:",inline"` // Common response fields
}

type ShopPublishableStatus struct {
	Region                string `json:"region"`                  // [Required] <p>Region of published shop.<br /></p>
	ShopPublishableStatus bool   `json:"shop_publishable_status"` // [Required] <p>If the shop is publishable, ture means shop is publishable, fals means shop is unpublishable<br /></p>
	UnpublishableReason   string `json:"unpublishable_reason"`    // [Required] <p>Return the unpublishable reason. If the shop is publishable, will return empty for this field.<br /></p>
}

type ShopSync struct {
	ShopRegion                 string `json:"shop_region"`                    // [Required] TW TH MY BR IN SG VN
	NameAndDescription         bool   `json:"name_and_description"`           // [Required] sync name and description
	MediaInformation           bool   `json:"media_information"`              // [Required] sync media information
	TierVariationNameAndOption bool   `json:"tier_variation_name_and_option"` // [Required] sync tier variation
	Price                      bool   `json:"price"`                          // [Required] sync price
	DaysToShip                 bool   `json:"days_to_ship"`                   // [Required] sync days to ship info
}

type SizeChart struct {
	SizeChartId int64 `json:"size_chart_id"` // [Required]
}

type SizeChartInfo struct {
	SizeChart   *string `json:"size_chart,omitempty"`    // [Optional] <p>ID of size chart image. If you want to remove the image size chart of the item, please pass the "size_chart" empty.&nbsp;</p><p>You only need to fill out either the image or template. If both are filled, only the template will be kept.</p><p>Notes: Both CB shops and local shops are supported to set "size_chart".</p>
	SizeChartId *int64  `json:"size_chart_id,omitempty"` // [Optional] <p>ID of template size chart. If you want to remove the template size chart of the item, please pass the "size_chart_id" as 0.&nbsp;</p><p>You only need to fill out either the image or template. If both are filled, only the template will be kept.</p><p>Notes: Both local shops and CB shops are supported to set "size_chart_id" now and seller need set the size_chart template in CBSC in advance</p>
}

type SizeChartLimit struct {
	SizeChartMandatory       bool `json:"size_chart_mandatory"`        // [Required]
	SupportImageSizeChart    bool `json:"support_image_size_chart"`    // [Required]
	SupportTemplateSizeChart bool `json:"support_template_size_chart"` // [Required]
}

type SizeChartTable struct {
	ColumnList []Column `json:"column_list"` // [Required] <p>column list of new size chart table. it include one column (measurement) and multiple values (measurement value)<br /></p>
}

type StandardiseTierVariation struct {
	VariationId         *int64            `json:"variation_id,omitempty"`          // [Optional]
	VariationName       *string           `json:"variation_name,omitempty"`        // [Optional]
	VariationGroupId    *int64            `json:"variation_group_id,omitempty"`    // [Optional]
	VariationOptionList []VariationOption `json:"variation_option_list,omitempty"` // [Optional]
}

type StandardiseTierVariationVariationOption struct {
	VariationOptionId   int64  `json:"variation_option_id"`   // [Required] <p>Standardise Tier variation Option ID.<br /></p>
	VariationOptionName string `json:"variation_option_name"` // [Required] <p>Standardise Tier variation Option Name.<br /></p>
	ImageId             string `json:"image_id"`              // [Required] <p>ID of image<br /></p>
	ImageUrl            string `json:"image_url"`             // [Required] <p>URL of image<br /></p>
}

type StandardiseVariation struct {
	VariationId        int64            `json:"variation_id"`         // [Required]
	VariationName      string           `json:"variation_name"`       // [Required]
	VariationGroupList []VariationGroup `json:"variation_group_list"` // [Required]
}

type StockInfo struct {
	StockType       int64  `json:"stock_type"`        // [Required] The stock type.
	StockLocationId string `json:"stock_location_id"` // [Required] location_id of the stock.
	NormalStock     int64  `json:"normal_stock"`      // [Required] The normal stock quantity of the variation in the listing currency.
	ReservedStock   int64  `json:"reserved_stock"`    // [Required] The reserved stock quantity of the variation in the listing currency.
}

type SupportSizeChartRequest struct {
	CategoryId int64 `json:"category_id" url:"category_id"` // [Required] Id of category.
}

type SupportSizeChartResponse struct {
	BaseResponse `json:",inline"`             // Common response fields
	Response     SupportSizeChartResponseData `json:"response"` // Actual response data
}

type SupportSizeChartResponseData struct {
	SupportSizeChart bool `json:"support_size_chart"` // [Required] If category support size chart.
}

type UpdateGlobalItemRequest struct {
	GlobalItemId    int64            `json:"global_item_id"`             // [Required] Id of global item.
	CategoryId      *int64           `json:"category_id,omitempty"`      // [Optional] Category id of global item.
	GlobalItemName  *string          `json:"global_item_name,omitempty"` // [Optional] Name of global item.
	Description     *string          `json:"description,omitempty"`      // [Optional] Description of global item.
	GlobalItemSku   *string          `json:"global_item_sku,omitempty"`  // [Optional] Sku of global item.
	Weight          *float64         `json:"weight,omitempty"`           // [Optional] <p>The weight of this global item, the unit is KG.</p><p>Updating the weight of this&nbsp;global item will overwrite the weight of all global models under this global item.</p>
	Dimension       *Dimension       `json:"dimension,omitempty"`        // [Optional] <p>The dimension of this global item.</p><p>Updating the dimension of this global item will overwrite the dimension of all global models under this global item.</p>
	PreOrder        *RequestPreOrder `json:"pre_order,omitempty"`        // [Optional] <p>Preorder information of global item.</p><p><br /></p><p>Updating the DTS of global item will overwrite the DTS of all global models under the global item</p>
	Condition       *string          `json:"condition,omitempty"`        // [Optional] Condition of global item, "NEW" or "USED" is available.
	Image           *PromotionImages `json:"image,omitempty"`            // [Optional] Image information of global item.
	VideoUploadId   []string         `json:"video_upload_id,omitempty"`  // [Optional] Video upload id of global item.
	Brand           *RequestBrand    `json:"brand,omitempty"`            // [Optional]
	AttributeList   []Attribute      `json:"attribute_list,omitempty"`   // [Optional] Item attributes.
	DescriptionInfo *DescriptionInfo `json:"description_info,omitempty"` // [Optional] New description field. New description field. Only whitelist sellers can use it. If you use the field, please upload the description_type=extended otherwise api will return error. If you don't use this field, you don't need to upload the description_type or upload description_type=normal
	DescriptionType *DescriptionType `json:"description_type,omitempty"` // [Optional] Values: See Data Definition- description_type (normal , extended). If you want to use extended_description or change description type ,this field must be inputed
	SizeChartInfo   *SizeChartInfo   `json:"size_chart_info,omitempty"`  // [Optional]
}

type UpdateGlobalItemResponse struct {
	BaseResponse `json:",inline"`             // Common response fields
	Response     UpdateGlobalItemResponseData `json:"response"` // Actual response data
}

type UpdateGlobalItemResponseData struct {
	GlobalItemId int64 `json:"global_item_id"` // [Required] Id of updated global item.
}

type UpdateGlobalModelRequest struct {
	GlobalItemId int64                `json:"global_item_id"` // [Required] ID of global item.
	GlobalModel  []RequestGlobalModel `json:"global_model"`   // [Required] Sku setting for global model. Limit is [1,50].
}

type UpdateGlobalModelResponse struct {
	BaseResponse `json:",inline"` // Common response fields
}

type UpdateLocalAdjustmentRateRequest struct {
	AdjustmentRate float64 `json:"adjustment_rate"` // [Required] <p>The multiplier used to adjust the cross-border original price to local price</p>
}

type UpdateLocalAdjustmentRateResponse struct {
	BaseResponse `json:",inline"` // Common response fields
}

type UpdateSizeChartRequest struct {
	GlobalItemId int64  `json:"global_item_id"` // [Required] Id of global item.
	SizeChart    string `json:"size_chart"`     // [Required] Image id of size chart.
}

type UpdateSizeChartResponse struct {
	BaseResponse `json:",inline"` // Common response fields
}

type UpdateTierVariationRequestModel struct {
	ModelId   int64 `json:"model_id"`   // [Required] <p>ID of model<br /></p>
	TierIndex int64 `json:"tier_index"` // [Required] <p>Model's tier_variation<br /></p>
}

type Value struct {
	ValueId   int64  `json:"value_id"`   // [Required]
	ValueName string `json:"value_name"` // [Required]
}

type VariationGroup struct {
	VariationGroupId    int64                           `json:"variation_group_id"`    // [Required]
	VariationGroupName  string                          `json:"variation_group_name"`  // [Required]
	VariationOptionList []VariationGroupVariationOption `json:"variation_option_list"` // [Required]
}

type VariationGroupVariationOption struct {
	VariationOptionId   int64  `json:"variation_option_id"`   // [Required]
	VariationOptionName string `json:"variation_option_name"` // [Required]
}

type VariationOption struct {
	VariationOptionId   *int64  `json:"variation_option_id,omitempty"`   // [Optional]
	VariationOptionName *string `json:"variation_option_name,omitempty"` // [Optional]
	ImageId             *string `json:"image_id,omitempty"`              // [Optional]
}

type VideoInfo struct {
	VideoUrl     string `json:"video_url"`     // [Required] Url of video.
	ThumbnailUrl string `json:"thumbnail_url"` // [Required] Thumbnail of video.
	Duration     int64  `json:"duration"`      // [Required] Duration of video.
}

type WeightLimit struct {
	WeightMandatory bool `json:"weight_mandatory"` // [Required] <p>weight is mandatory or not<br /></p>
}

type WholesalePriceThresholdPercentage struct {
	MinLimit int64 `json:"min_limit"` // [Required]
	MaxLimit int64 `json:"max_limit"` // [Required]
}
