package goshopee

import (
	"fmt"
)

type ProductService interface {
	// https://open.shopee.com/documents/v2/v2.product.get_category?module=89&type=1
	GetCategory(shopID uint64, language string, accessToken string) (*GetCategoryResponse, error)

	// https://open.shopee.com/documents/v2/v2.product.get_category_tree?module=89&type=1
	GetCategoryTree(shopID uint64, language string, accessToken string) (*GetCategoryResponse, error)

	// https://open.shopee.com/documents/v2/v2.product.get_brand_list?module=89&type=1
	GetBrandList(shopID uint64, categoryID uint64, status int, offset int, pageSize int, accessToken string) (*GetBrandListResponse, error)

	// https://open.shopee.com/documents/v2/v2.product.get_dts_limit?module=89&type=1
	GetDTSLimit(shopID uint64, categoryID uint64, accessToken string) (*GetDTSLimitResponse, error)

	// https://open.shopee.com/documents/v2/v2.product.get_attributes?module=89&type=1
	GetAttributes(shopID uint64, categoryID uint64, language string, accessToken string) (*GetAttributesResponse, error)

	// https://open.shopee.com/documents/v2/v2.product.get_attribute_tree?module=89&type=1
	GetAttributeTree(shopID uint64, categoryIDList []uint64, language string, accessToken string) (*GetAttributeTreeResponse, error)

	// https://open.shopee.com/documents/v2/v2.product.get_size_chart_list?module=89&type=1
	GetSizeChartList(shopID uint64, categoryID uint64, accessToken string) (*GetSizeChartListResponse, error)

	// Deprecated: SupportSizeChart is deprecated.
	//
	// https://open.shopee.com/documents/v2/v2.product.support_size_chart?module=89&type=1
	SupportSizeChart(shopID uint64, categoryID uint64, accessToken string) (*SupportSizeChartResponse, error)

	// Deprecated: UpdateSizeChart is deprecated.
	//
	// https://open.shopee.com/documents/v2/v2.product.update_size_chart?module=89&type=1
	UpdateSizeChart(shopID uint64, itemID uint64, sizeChart string, accessToken string) (*UpdateSizeChartResponse, error)

	// https://open.shopee.com/documents/v2/v2.product.get_item_base_info?module=89&type=1
	GetItemBaseInfo(shopID uint64, itemIDList []uint64, accessToken string) (*GetItemBaseInfoResponse, error)

	// https://open.shopee.com/documents/v2/v2.product.add_item?module=89&type=1
	AddItem(shopID uint64, req AddItemRequest, accessToken string) (*AddItemResponse, error)

	// https://open.shopee.com/documents/v2/v2.product.delete_item?module=89&type=1
	DeleteItem(shopID uint64, itemID uint64, accessToken string) (*BaseResponse, error)

	// https://open.shopee.com/documents/v2/v2.product.update_item?module=89&type=1
	UpdateItem(shopID uint64, req UpdateItemRequest, accessToken string) (*UpdateItemResponse, error)

	// https://open.shopee.com/documents/v2/v2.product.unlist_item?module=89&type=1
	UnlistItem(shopID uint64, req UnlistItemRequest, accessToken string) (*UnlistItemResponse, error)

	// https://open.shopee.com/documents/v2/v2.product.search_item?module=89&type=1
	SearchItem(shopID uint64, opt SearchItemRequest, accessToken string) (*SearchItemResponse, error)

	// https://open.shopee.com/documents/v2/v2.product.init_tier_variation?module=89&type=1
	InitTierVariation(shopID uint64, req InitTierVariationRequest, accessToken string) (*InitTierVariationResponse, error)

	// https://open.shopee.com/documents/v2/v2.product.update_tier_variation?module=89&type=1
	UpdateTierVariation(shopID uint64, req UpdateTierVariationRequest, accessToken string) (*UpdateTierVariationResponse, error)

	// https://open.shopee.com/documents/v2/v2.product.get_model_list?module=89&type=1
	GetModelList(shopID uint64, itemID uint64, accessToken string) (*GetModelListResponse, error)

	// https://open.shopee.com/documents/v2/v2.product.add_model?module=89&type=1
	AddModel(shopID uint64, req AddModelRequest, accessToken string) (*AddModelResponse, error)

	// https://open.shopee.com/documents/v2/v2.product.delete_model?module=89&type=1
	DeleteModel(shopID uint64, itemID uint64, modelID uint64, accessToken string) (*BaseResponse, error)

	// https://open.shopee.com/documents/v2/v2.product.update_model?module=89&type=1
	UpdateModel(shopID uint64, req UpdateModelRequest, accessToken string) (*UpdateModelResponse, error)

	// https://open.shopee.com/documents/v2/v2.product.update_price?module=89&type=1
	UpdatePrice(shopID uint64, req UpdatePriceRequest, accessToken string) (*UpdatePriceResponse, error)

	// https://open.shopee.com/documents/v2/v2.product.update_stock?module=89&type=1
	UpdateStock(shopID uint64, req UpdateStockRequest, accessToken string) (*UpdateStockResponse, error)

	// https://open.shopee.com/documents/v2/v2.product.category_recommend?module=89&type=1
	CategoryRecommend(shopID uint64, itemName string, accessToken string) (*CategoryRecommendResponse, error)

	// https://open.shopee.com/documents/v2/v2.product.get_item_promotion?module=89&type=1
	GetItemPromotion(shopID uint64, itemIDList []uint64, accessToken string) (*GetItemPromotionResponse, error)

	// https://open.shopee.com/documents/v2/v2.product.get_item_list?module=89&type=1
	GetItemList(shopID uint64, opt GetItemListRequest, accessToken string) (*GetItemListResponse, error)

	// https://open.shopee.com/documents/v2/v2.product.boost_item?module=89&type=1
	BoostItem(shopID uint64, itemIDList []uint64, accessToken string) (*BaseResponse, error)

	// https://open.shopee.com/documents/v2/v2.product.get_boosted_list?module=89&type=1
	GetBoostedList(shopID uint64, accessToken string) (*GetBoostedListResponse, error)

	// https://open.shopee.com/documents/v2/v2.product.get_comment?module=89&type=1
	GetComment(shopID uint64, opt GetCommentRequest, accessToken string) (*GetCommentResponse, error)

	// https://open.shopee.com/documents/v2/v2.product.reply_comment?module=89&type=1
	ReplyComment(shopID uint64, req ReplyCommentRequest, accessToken string) (*BaseResponse, error)
}

type ProductServiceOp struct {
	client *Client
}

// Category related structs

type GetCategoryRequest struct {
	// Language is the language of the category names. If not provided, the default language of the shop will be used.
	Language string `url:"language"`
}

type GetCategoryResponse struct {
	BaseResponse
	Response GetCategoryResponseData `json:"response"`
}

type GetCategoryResponseData struct {
	// CategoryList is the list of categories.
	CategoryList []Category `json:"category_list"`
}

type Category struct {
	// CategoryID is the identifier of the category.
	CategoryID uint64 `json:"category_id"`
	// ParentCategoryID is the identifier of the parent category.
	ParentCategoryID uint64 `json:"parent_category_id"`
	// OriginalCategoryName is the original name of the category.
	OriginalCategoryName string `json:"original_category_name"`
	// DisplayCategoryName is the display name of the category.
	DisplayCategoryName string `json:"display_category_name"`
	// HasChildren indicates whether the category has sub-categories.
	HasChildren bool `json:"has_children"`
}

func (s *ProductServiceOp) GetCategory(sid uint64, lang, tok string) (*GetCategoryResponse, error) {
	path := "/product/get_category"
	opt := GetCategoryRequest{Language: lang}
	resp := new(GetCategoryResponse)
	err := s.client.WithShop(sid, tok).Get(path, resp, opt)
	return resp, err
}

func (s *ProductServiceOp) GetCategoryTree(sid uint64, lang, tok string) (*GetCategoryResponse, error) {
	path := "/product/get_category_tree"
	opt := GetCategoryRequest{Language: lang}
	resp := new(GetCategoryResponse)
	err := s.client.WithShop(sid, tok).Get(path, resp, opt)
	return resp, err
}

// Brand related structs

type GetBrandListRequest struct {
	// Offset is the offset of the brand list.
	Offset int `url:"offset"`
	// PageSize is the number of brands to retrieve per page. Max is 100.
	PageSize int `url:"page_size"`
	// CategoryID is the identifier of the category for which to retrieve brands.
	CategoryID uint64 `url:"category_id" mandatory:"true"`
	// Status is the status of the brands to retrieve. 1: normal, 2: pending.
	Status int `url:"status" mandatory:"true"`
}

type GetBrandListResponse struct {
	BaseResponse
	Response GetBrandListResponseData `json:"response"`
}

type GetBrandListResponseData struct {
	// BrandList is the list of brands.
	BrandList []Brand `json:"brand_list"`
	// HasNextPage indicates whether there are more brands to retrieve.
	HasNextPage bool `json:"has_next_page"`
	// NextOffset is the offset to be used for the next request to retrieve more brands.
	NextOffset int `json:"next_offset"`
	// IsMandatory indicates whether a brand is mandatory for the specified category.
	IsMandatory bool `json:"is_mandatory"`
	// InputType is the input type of the brand.
	InputType string `json:"input_type"`
}

type Brand struct {
	// BrandID is the identifier of the brand.
	BrandID uint64 `json:"brand_id"`
	// OriginalBrandName is the original name of the brand.
	OriginalBrandName string `json:"original_brand_name"`
	// DisplayBrandName is the display name of the brand.
	DisplayBrandName string `json:"display_brand_name"`
}

func (s *ProductServiceOp) GetBrandList(sid uint64, cid uint64, status int, offset int, pageSize int, tok string) (*GetBrandListResponse, error) {
	path := "/product/get_brand_list"
	opt := GetBrandListRequest{
		CategoryID: cid,
		Offset:     offset,
		PageSize:   pageSize,
		Status:     status,
	}
	resp := new(GetBrandListResponse)
	err := s.client.WithShop(sid, tok).Get(path, resp, opt)
	return resp, err
}

// DTS Limit related structs

type GetDTSLimitRequest struct {
	// CategoryID is the identifier of the category for which to retrieve the DTS limit.
	CategoryID uint64 `url:"category_id" mandatory:"true"`
}

type GetDTSLimitResponse struct {
	BaseResponse
	Response GetDTSLimitResponseData `json:"response"`
}

type GetDTSLimitResponseData struct {
	// DaysToShipLimit is the DTS limit for the specified category.
	DaysToShipLimit DaysToShipLimit `json:"days_to_ship_limit"`
	// NonPreOrderDaysToShip is the non-pre-order DTS limit for the specified category.
	NonPreOrderDaysToShip int `json:"non_pre_order_days_to_ship"`
}

type DaysToShipLimit struct {
	// MinLimit is the minimum DTS limit.
	MinLimit int `json:"min_limit"`
	// MaxLimit is the maximum DTS limit.
	MaxLimit int `json:"max_limit"`
}

func (s *ProductServiceOp) GetDTSLimit(sid uint64, cid uint64, tok string) (*GetDTSLimitResponse, error) {
	path := "/product/get_dts_limit"
	opt := GetDTSLimitRequest{CategoryID: cid}
	resp := new(GetDTSLimitResponse)
	err := s.client.WithShop(sid, tok).Get(path, resp, opt)
	return resp, err
}

// Attribute related structs

type GetAttributesRequest struct {
	// CategoryID is the identifier of the category for which to retrieve attributes.
	CategoryID uint64 `url:"category_id" mandatory:"true"`
	// Language is the language of the attribute names. If not provided, the default language of the shop will be used.
	Language string `url:"language"`
}

type GetAttributesResponse struct {
	BaseResponse
	Response GetAttributesResponseData `json:"response"`
}

type GetAttributesResponseData struct {
	// AttributeList is the list of attributes for the specified category.
	AttributeList []Attribute `json:"attribute_list"`
}

type Attribute struct {
	// AttributeID is the identifier of the attribute.
	AttributeID uint64 `json:"attribute_id"`
	// OriginalAttributeName is the original name of the attribute.
	OriginalAttributeName string `json:"original_attribute_name"`
	// DisplayAttributeName is the display name of the attribute.
	DisplayAttributeName string `json:"display_attribute_name"`
	// IsMandatory indicates whether the attribute is mandatory.
	IsMandatory bool `json:"is_mandatory"`
	// InputValidationType is the validation type for the input of the attribute.
	InputValidationType string `json:"input_validation_type"`
	// FormatType is the format type for the input of the attribute.
	FormatType string `json:"format_type"`
	// DateFormatType is the date format type for the input of the attribute.
	DateFormatType string `json:"date_format_type"`
	// InputType is the input type for the attribute.
	InputType string `json:"input_type"`
	// AttributeUnit is the unit of the attribute.
	AttributeUnit []string `json:"attribute_unit"`
	// AttributeValueList is the list of possible values for the attribute.
	AttributeValueList []AttributeValue `json:"attribute_value_list"`
}

type AttributeValue struct {
	// ValueID is the identifier of the attribute value.
	ValueID uint64 `json:"value_id"`
	// OriginalValueName is the original name of the attribute value.
	OriginalValueName string `json:"original_value_name"`
	// DisplayValueName is the display name of the attribute value.
	DisplayValueName string `json:"display_value_name"`
	// ValueUnit is the unit of the attribute value.
	ValueUnit string `json:"value_unit"`
	// ParentAttributeList is the list of parent attributes for the attribute value.
	ParentAttributeList []ParentAttribute `json:"parent_attribute_list"`
	// ParentBrandList is the list of parent brands for the attribute value.
	ParentBrandList []ParentBrand `json:"parent_brand_list"`
}

type ParentAttribute struct {
	// ParentAttributeID is the identifier of the parent attribute.
	ParentAttributeID uint64 `json:"parent_attribute_id"`
	// ParentValueID is the identifier of the parent value.
	ParentValueID uint64 `json:"parent_value_id"`
}

type ParentBrand struct {
	// ParentBrandID is the identifier of the parent brand.
	ParentBrandID uint64 `json:"parent_brand_id"`
}

func (s *ProductServiceOp) GetAttributes(sid uint64, cid uint64, lang, tok string) (*GetAttributesResponse, error) {
	path := "/product/get_attributes"
	opt := GetAttributesRequest{CategoryID: cid, Language: lang}
	resp := new(GetAttributesResponse)
	err := s.client.WithShop(sid, tok).Get(path, resp, opt)
	return resp, err
}

type GetAttributeTreeRequest struct {
	// CategoryIDList is the identifiers of the categories for which to retrieve the attribute tree.
	CategoryIDList []uint64 `url:"category_id_list" mandatory:"true"`
	// Language is the language of the attribute names. If not provided, the default language of the shop will be used.
	Language string `url:"language"`
}

type GetAttributeTreeResponse struct {
	BaseResponse
	Response GetAttributeTreeResponseData `json:"response"`
}

type GetAttributeTreeResponseData struct {
	// List is the attribute tree for the specified categories.
	List []AttributeTreeCategory `json:"list"`
}

type AttributeTreeCategory struct {
	// CategoryID is the identifier of the category.
	CategoryID uint64 `json:"category_id"`
	// AttributeTree is the attribute tree for the category.
	AttributeTree []AttributeTree `json:"attribute_tree"`
	// Warning is the warning message if any.
	Warning string `json:"warning"`
}

type AttributeTree struct {
	// AttributeID is the identifier of the attribute.
	AttributeID uint64 `json:"attribute_id"`
	// Mandatory indicates whether the attribute is mandatory.
	Mandatory bool `json:"mandatory"`
	// Name is the name of the attribute.
	Name string `json:"name"`
	// AttributeValueList is the list of attribute values.
	AttributeValueList []AttributeTreeAttributeValue `json:"attribute_value_list"`
	// AttributeInfo is the information about the attribute.
	AttributeInfo AttributeTreeAttributeInfo `json:"attribute_info"`
	// MultiLang is the multi-language information for the attribute.
	MultiLang []AttributeTreeMultiLang `json:"multi_lang"`
}

type AttributeTreeAttributeValue struct {
	// ValueID is the identifier of the attribute value.
	ValueID uint64 `json:"value_id"`
	// Name is the name of the attribute value.
	Name string `json:"name"`
	// ValueUnit is the unit of the attribute value.
	ValueUnit string `json:"value_unit"`
	// ChildAttributeList is the list of child attributes.
	ChildAttributeList []AttributeTree `json:"child_attribute_list"`
	// MultiLang is the multi-language information for the attribute value.
	MultiLang []AttributeTreeMultiLang `json:"multi_lang"`
}

type AttributeTreeInputType uint8

const (
	SINGLE_DROP_DOWN AttributeTreeInputType = iota + 1
	SINGLE_COMBO_BOX
	FREE_TEXT_FILED
	MULTI_DROP_DOWN
	MULTI_COMBO_BOX
)

type AttributeTreeInpuValidationType uint8

const (
	VALIDATOR_NO_VALIDATE_TYPE AttributeTreeInpuValidationType = iota
	VALIDATOR_INT_TYPE
	VALIDATOR_STRING_TYPE
	VALIDATOR_FLOAT_TYPE
	VALIDATOR_DATE_TYPE
)

type AttributeTreeFormatType uint8

const (
	FORMAT_NORMAL AttributeTreeFormatType = iota + 1
	FORMAT_QUANTITATIVE_WITH_UNIT
)

type AttributeTreeDateFormatType uint8

const (
	YEAR_MONTH_DATE AttributeTreeDateFormatType = iota // (DD/MM/YYYY)
	YEAR_MONTH_DAY                                     // (MM/YYYY)
)

type AttributeTreeAttributeInfo struct {
	// InputType is the input type of the attribute.
	InputType AttributeTreeInputType `json:"input_type"`
	// InputValidationType is the input validation type of the attribute.
	InputValidationType AttributeTreeInpuValidationType `json:"input_validation_type"`
	// FormatType is the format type of the attribute.
	FormatType AttributeTreeFormatType `json:"format_type"`
	// DateFormatType is the date format type of the attribute.
	DateFormatType AttributeTreeDateFormatType `json:"date_format_type"`
	// AttributeUnitList is the list of units for the attribute.
	AttributeUnitList []string `json:"attribute_unit_list"`
	// MaxValueCount is the maximum number of values that can be selected for the attribute.
	MaxValueCount int64 `json:"max_value_count"`
	// Introduction is the introduction for the attribute.
	Introduction string `json:"introduction"`
	// IsOem indicates whether the attribute is an OEM attribute.
	IsOem bool `json:"is_oem"`
	// SupportSearchValue indicates whether the attribute supports search value.
	SupportSearchValue bool `json:"support_search_value"`
}

type AttributeTreeMultiLang struct {
	// Language is the language of the name.
	Language string `json:"language"`
	// Value is the value of the name in the specified language.
	Value string `json:"value"`
}

func (s *ProductServiceOp) GetAttributeTree(sid uint64, cidList []uint64, lang, tok string) (*GetAttributeTreeResponse, error) {
	path := "/product/get_attribute_tree"
	opt := GetAttributeTreeRequest{CategoryIDList: cidList, Language: lang}
	resp := new(GetAttributeTreeResponse)
	err := s.client.WithShop(sid, tok).Get(path, resp, opt)
	return resp, err
}

// Size Chart related structs

type GetSizeChartListRequest struct {
	// CategoryID is the identifier of the category for which to retrieve size charts.
	CategoryID uint64 `url:"category_id" mandatory:"true"`
}

type GetSizeChartListResponse struct {
	BaseResponse
	Response GetSizeChartListResponseData `json:"response"`
}

type GetSizeChartListResponseData struct {
	// SizeChartList is the list of size charts.
	SizeChartList []SizeChart `json:"size_chart_list"`
}

type SizeChart struct {
	// SizeChartID is the identifier of the size chart.
	SizeChartID uint64 `json:"size_chart_id"`
	// SizeChartName is the name of the size chart.
	SizeChartName string `json:"size_chart_name"`
}

func (s *ProductServiceOp) GetSizeChartList(sid uint64, cid uint64, tok string) (*GetSizeChartListResponse, error) {
	path := "/product/get_size_chart_list"
	opt := GetSizeChartListRequest{CategoryID: cid}
	resp := new(GetSizeChartListResponse)
	err := s.client.WithShop(sid, tok).Get(path, resp, opt)
	return resp, err
}

type SupportSizeChartRequest struct {
	// CategoryID is the identifier of the category for which to check size chart support.
	CategoryID uint64 `url:"category_id" mandatory:"true"`
}

type SupportSizeChartResponse struct {
	BaseResponse
	Response SupportSizeChartResponseData `json:"response"`
}

type SupportSizeChartResponseData struct {
	// SupportSizeChart indicates whether the category supports size charts.
	SupportSizeChart bool `json:"support_size_chart"`
}

func (s *ProductServiceOp) SupportSizeChart(sid uint64, cid uint64, tok string) (*SupportSizeChartResponse, error) {
	return nil, fmt.Errorf("SupportSizeChart is deprecated. Use v2.global_product.get_global_item_limit instead.")
}

type UpdateSizeChartResponse struct {
	BaseResponse
}

func (s *ProductServiceOp) UpdateSizeChart(sid uint64, itemID uint64, sizeChart string, tok string) (*UpdateSizeChartResponse, error) {
	return nil, fmt.Errorf("UpdateSizeChart is deprecated. Use v2.global_product.add_global_item or v2.global_product.update_global_item instead.")
}

// Item related structs

type ItemStatus string

const (
	ItemStatusNormal       ItemStatus = "NORMAL"
	ItemStatusBanned       ItemStatus = "BANNED"
	ItemStatusUnlist       ItemStatus = "UNLIST"
	ItemStatusReviewing    ItemStatus = "REVIEWING"
	ItemStatusSellerDelete ItemStatus = "SELLER_DELETE"
	ItemStatusShopeeDelete ItemStatus = "SHOPEE_DELETE"
)

type AddItemRequest struct {
	// CategoryID is the identifier of the category for the item.
	CategoryID uint64 `json:"category_id" mandatory:"true"`
	// ItemName is the name of the item.
	ItemName string `json:"item_name" mandatory:"true"`
	// Description is the description of the item.
	Description string `json:"description" mandatory:"true"`
	// ItemSKU is the SKU for the item.
	ItemSKU string `json:"item_sku"`
	// AttributeList is the list of attributes for the item.
	AttributeList []ItemAttribute `json:"attribute_list"`
	// Image is the images for the item.
	Image ItemImage `json:"image" mandatory:"true"`
	// Weight is the weight of the item.
	Weight float64 `json:"weight" mandatory:"true"`
	// Dimension is the dimension of the item.
	Dimension *Dimension `json:"dimension"`
	// LogisticInfo is the logistic information for the item.
	LogisticInfo []LogisticInfo `json:"logistic_info" mandatory:"true"`
	// PreOrder is the pre-order information for the item.
	PreOrder ItemPreOrder `json:"pre_order"`
	// Wholesale is the wholesale information for the item.
	Wholesale []ItemWholesale `json:"wholesale"`
	// Condition is the condition of the item. NEW or USED.
	Condition string `json:"condition"`
	// SizeChart is the size chart identifier for the item.
	SizeChart string `json:"size_chart"`
	// ItemStatus is the status of the item. NORMAL or UNLIST.
	ItemStatus string `json:"item_status"`
	// HasModel indicates whether the item has models.
	HasModel bool `json:"has_model"`
	// Brand is the brand information for the item.
	Brand ItemBrand `json:"brand"`
	// ItemDangerous is the dangerous information for the item. 0: none, 1: dangerous.
	ItemDangerous int `json:"item_dangerous"`
	// SellerStock is the stock information for the item. Required if has_model is false.
	SellerStock []SellerStock `json:"seller_stock"`
	// OriginalPrice is the original price of the item. Required if has_model is false.
	OriginalPrice float64 `json:"original_price"`
	// VideoUploadID is the video identifiers for the item.
	VideoUploadID []string `json:"video_upload_id"`
	// DescriptionInfo is the extended description information for the item.
	DescriptionInfo *DescriptionInfo `json:"description_info"`
	// DescriptionType is the type of description. NORMAL or EXTENDED.
	DescriptionType string `json:"description_type"`
	// TaxInfo is the tax information for the item.
	TaxInfo *TaxInfo `json:"tax_info"`
	// ComplaintPolicy is the complaint policy for the item.
	ComplaintPolicy *ComplaintPolicy `json:"complaint_policy"`
}

type DescriptionInfo struct {
	// ExtendedDescription is the extended description for the item.
	ExtendedDescription []DescriptionElement `json:"extended_description"`
}

type DescriptionElement struct {
	// FieldType is the type of the field. text or image.
	FieldType string `json:"field_type"`
	// Text is the text content if field_type is text.
	Text string `json:"text"`
	// ImageInfo is the image information if field_type is image.
	ImageInfo *DescriptionImageInfo `json:"image_info"`
}

type DescriptionImageInfo struct {
	// ImageID is the identifier of the image.
	ImageID string `json:"image_id"`
}

type TaxInfo struct {
	// Ncm is the NCM code for the item. (Brazil only)
	Ncm string `json:"ncm"`
	// SameStateCfop is the CFOP code for the item in the same state. (Brazil only)
	SameStateCfop string `json:"same_state_cfop"`
	// DiffStateCfop is the CFOP code for the item in a different state. (Brazil only)
	DiffStateCfop string `json:"diff_state_cfop"`
	// Csosn is the CSOSN code for the item. (Brazil only)
	Csosn string `json:"csosn"`
	// Origin is the origin of the item. (Brazil only)
	Origin string `json:"origin"`
	// Cest is the CEST code for the item. (Brazil only)
	Cest string `json:"cest"`
	// MeasureUnit is the measure unit for the item. (Brazil only)
	MeasureUnit string `json:"measure_unit"`
	// TaxCode is the tax code for the item. (Brazil only)
	TaxCode string `json:"tax_code"`
}

type ComplaintPolicy struct {
	// WarrantyTime is the warranty time for the item.
	WarrantyTime string `json:"warranty_time"`
	// ExcludeLanguage is the language to exclude for the complaint policy.
	ExcludeLanguage []string `json:"exclude_language"`
}

type AddItemResponse struct {
	BaseResponse
	Response AddItemResponseData `json:"response"`
}

type AddItemResponseData struct {
	// ItemID is the identifier of the item created.
	ItemID uint64 `json:"item_id"`
	// ItemBase is the detailed information of the item.
	ItemBase ItemBase `json:"item_base"`
}

type ItemBase struct {
	// CategoryID is the identifier of the category.
	CategoryID uint64 `json:"category_id"`
	// ItemName is the name of the item.
	ItemName string `json:"item_name"`
	// Description is the description of the item.
	Description string `json:"description"`
	// ItemSKU is the SKU for the item.
	ItemSKU string `json:"item_sku"`
	// CreateTime is the create time of the item.
	CreateTime uint64 `json:"create_time"`
	// UpdateTime is the update time of the item.
	UpdateTime uint64 `json:"update_time"`
	// AttributeList is the list of attributes for the item.
	AttributeList []ItemAttribute `json:"attribute_list"`
	// Image is the images for the item.
	Image ItemImage `json:"image"`
	// Weight is the weight of the item.
	Weight float64 `json:"weight"`
	// Dimension is the dimension of the item.
	Dimension *Dimension `json:"dimension,omitempty"`
	// LogisticInfo is the logistic information for the item.
	LogisticInfo []LogisticInfo `json:"logistic_info"`
	// PreOrder is the pre-order information for the item.
	PreOrder ItemPreOrder `json:"pre_order"`
	// Wholesale is the wholesale information for the item.
	Wholesale []ItemWholesale `json:"wholesale"`
	// Condition is the condition of the item.
	Condition string `json:"condition"`
	// SizeChart is the size chart identifier for the item.
	SizeChart string `json:"size_chart"`
	// ItemStatus is the status of the item.
	ItemStatus string `json:"item_status"`
	// HasModel indicates whether the item has models.
	HasModel bool `json:"has_model"`
	// PromotionID is the identifier of the promotion.
	PromotionID uint64 `json:"promotion_id"`
	// VideoInfo is the video information for the item.
	VideoInfo []ItemVideo `json:"video_info"`
	// Brand is the brand information for the item.
	Brand ItemBrand `json:"brand"`
	// ItemDangerous is the dangerous information for the item.
	ItemDangerous int `json:"item_dangerous"`
}

func (s *ProductServiceOp) AddItem(sid uint64, req AddItemRequest, tok string) (*AddItemResponse, error) {
	path := "/product/add_item"
	resp := new(AddItemResponse)
	reqMap, err := StructToMap(req)
	if err != nil {
		return nil, err
	}
	err = s.client.WithShop(sid, tok).Post(path, reqMap, resp)
	return resp, err
}

func (s *ProductServiceOp) DeleteItem(sid uint64, itemID uint64, tok string) (*BaseResponse, error) {
	path := "/product/delete_item"
	req := map[string]interface{}{"item_id": itemID}
	resp := new(BaseResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

type UpdateItemRequest struct {
	// ItemID is the identifier of the item to update.
	ItemID uint64 `json:"item_id" mandatory:"true"`
	// CategoryID is the identifier of the category for the item.
	CategoryID uint64 `json:"category_id"`
	// ItemName is the name of the item.
	ItemName string `json:"item_name"`
	// Description is the description of the item.
	Description string `json:"description"`
	// ItemSKU is the SKU for the item.
	ItemSKU string `json:"item_sku"`
	// AttributeList is the list of attributes for the item.
	AttributeList []ItemAttribute `json:"attribute_list"`
	// Image is the images for the item.
	Image ItemImage `json:"image"`
	// Weight is the weight of the item.
	Weight float64 `json:"weight"`
	// Dimension is the dimension of the item.
	Dimension *Dimension `json:"dimension"`
	// LogisticInfo is the logistic information for the item.
	LogisticInfo []LogisticInfo `json:"logistic_info"`
	// PreOrder is the pre-order information for the item.
	PreOrder ItemPreOrder `json:"pre_order"`
	// Wholesale is the wholesale information for the item.
	Wholesale []ItemWholesale `json:"wholesale"`
	// Condition is the condition of the item. NEW or USED.
	Condition string `json:"condition"`
	// SizeChart is the size chart identifier for the item.
	SizeChart string `json:"size_chart"`
	// VideoUploadID is the video identifiers for the item.
	VideoUploadID []string `json:"video_upload_id"`
	// Brand is the brand information for the item.
	Brand ItemBrand `json:"brand"`
	// ItemDangerous is the dangerous information for the item. 0: none, 1: dangerous.
	ItemDangerous int `json:"item_dangerous"`
	// DescriptionInfo is the extended description information for the item.
	DescriptionInfo *DescriptionInfo `json:"description_info"`
	// DescriptionType is the type of description. NORMAL or EXTENDED.
	DescriptionType string `json:"description_type"`
	// TaxInfo is the tax information for the item.
	TaxInfo *TaxInfo `json:"tax_info"`
	// ComplaintPolicy is the complaint policy for the item.
	ComplaintPolicy *ComplaintPolicy `json:"complaint_policy"`
}

type UpdateItemResponse struct {
	BaseResponse
	Response UpdateItemResponseData `json:"response"`
}

type UpdateItemResponseData struct {
	// ItemID is the identifier of the item updated.
	ItemID uint64 `json:"item_id"`
	// ItemBase is the detailed information of the item.
	ItemBase ItemBase `json:"item_base"`
}

func (s *ProductServiceOp) UpdateItem(sid uint64, req UpdateItemRequest, tok string) (*UpdateItemResponse, error) {
	path := "/product/update_item"
	resp := new(UpdateItemResponse)
	reqMap, err := StructToMap(req)
	if err != nil {
		return nil, err
	}
	err = s.client.WithShop(sid, tok).Post(path, reqMap, resp)
	return resp, err
}

type UnlistItemRequest struct {
	// ItemIDList is the list of item identifiers to unlist or list.
	ItemIDList []uint64 `json:"item_id_list" mandatory:"true"`
	// Unlist indicates whether to unlist the items. true: unlist, false: list.
	Unlist bool `json:"unlist" mandatory:"true"`
}

type UnlistItemResponse struct {
	BaseResponse
	Response UnlistItemResponseData `json:"response"`
}

type UnlistItemResponseData struct {
	// FailureList is the list of items that failed to be unlisted or listed.
	FailureList []UnlistItemResponseDataFail `json:"failure_list"`
	// SuccessList is the list of items that were successfully unlisted or listed.
	SuccessList []UnlistItemResponseDataSuccess `json:"success_list"`
}

type UnlistItemResponseDataFail struct {
	// ItemID is the identifier of the item.
	ItemID uint64 `json:"item_id"`
	// FailedReason is the reason for the failure.
	FailedReason string `json:"failed_reason"`
}

type UnlistItemResponseDataSuccess struct {
	// ItemID is the identifier of the item.
	ItemID uint64 `json:"item_id"`
	// Unlist indicates whether the item is currently unlisted.
	Unlist bool `json:"unlist"`
}

func (s *ProductServiceOp) UnlistItem(sid uint64, req UnlistItemRequest, tok string) (*UnlistItemResponse, error) {
	path := "/product/unlist_item"
	resp := new(UnlistItemResponse)
	reqMap, err := StructToMap(req)
	if err != nil {
		return nil, err
	}
	err = s.client.WithShop(sid, tok).Post(path, reqMap, resp)
	return resp, err
}

type SearchItemRequest struct {
	// Offset is the offset of the item list.
	Offset int `url:"offset"`
	// PageSize is the number of items to retrieve per page. Max is 100.
	PageSize int `url:"page_size" mandatory:"true"`
	// ItemSKU is the SKU for searching items.
	ItemSKU string `url:"item_sku"`
	// ItemName is the name for searching items.
	ItemName string `url:"item_name"`
	// ItemIDList is the identifiers of the items for searching.
	ItemIDList []uint64 `url:"item_id_list"`
	// ItemStatus is the status of the items for searching.
	ItemStatus []ItemStatus `url:"item_status"`
	// AttributeID is the identifier of the attribute.
	AttributeID uint64 `url:"attribute_id"`
}

type SearchItemResponse struct {
	BaseResponse
	Response SearchItemResponseData `json:"response"`
}

type SearchItemResponseData struct {
	// ItemIDList is the list of items found.
	ItemIDList []uint64 `json:"item_id_list"`
	// TotalCount is the total count of all items.
	TotalCount int `json:"total_count"`
	// HasNextPage indicates whether the item list is more than one page.
	HasNextPage bool `json:"has_next_page"`
	// NextOffset is the offset to be used for the next request.
	NextOffset int `json:"next_offset"`
}

func (s *ProductServiceOp) SearchItem(sid uint64, opt SearchItemRequest, tok string) (*SearchItemResponse, error) {
	path := "/product/search_item"
	resp := new(SearchItemResponse)
	err := s.client.WithShop(sid, tok).Get(path, resp, opt)
	return resp, err
}

type GetItemBaseInfoRequest struct {
	// ItemIDList is the identifiers of the items for which to retrieve detailed information.
	ItemIDList []uint64 `url:"item_id_list" mandatory:"true"`
}

type GetItemBaseInfoResponse struct {
	BaseResponse
	Response GetItemBaseInfoResponseData `json:"response"`
}

type GetItemBaseInfoResponseData struct {
	// ItemList is the list of detailed information for the specified items.
	ItemList []ItemBaseInfoData `json:"item_list"`
}

type ItemBaseInfoData struct {
	// ItemID is the identifier of the item.
	ItemID uint64 `json:"item_id"`
	// ItemBaseInfo is the base information of the item.
	ItemBaseInfo ItemBaseInfo `json:"item_base_info"`
	// StockInfoV2 is the stock information for the item.
	StockInfoV2 StockInfoV2 `json:"stock_info_v2"`
	// PriceInfo is the price information for the item.
	PriceInfo PriceInfo `json:"price_info"`
}

type ItemBaseInfo struct {
	// CategoryID is the identifier of the category.
	CategoryID uint64 `json:"category_id"`
	// ItemName is the name of the item.
	ItemName string `json:"item_name"`
	// Description is the description of the item.
	Description string `json:"description"`
	// ItemSKU is the SKU for the item.
	ItemSKU string `json:"item_sku"`
	// CreateTime is the create time of the item.
	CreateTime uint64 `json:"create_time"`
	// UpdateTime is the update time of the item.
	UpdateTime uint64 `json:"update_time"`
	// AttributeList is the list of attributes for the item.
	AttributeList []ItemAttribute `json:"attribute_list"`
	// Image is the images for the item.
	Image ItemImage `json:"image"`
	// Weight is the weight of the item.
	Weight string `json:"weight"`
	// Dimension is the dimension of the item.
	Dimension *Dimension `json:"dimension,omitempty"`
	// LogisticInfo is the logistic information for the item.
	LogisticInfo []LogisticInfo `json:"logistic_info"`
	// PreOrder is the pre-order information for the item.
	PreOrder ItemPreOrder `json:"pre_order"`
	// Wholesale is the wholesale information for the item.
	Wholesale []ItemWholesale `json:"wholesale"`
	// Condition is the condition of the item.
	Condition string `json:"condition"`
	// SizeChart is the size chart identifier for the item.
	SizeChart string `json:"size_chart"`
	// ItemStatus is the status of the item.
	ItemStatus string `json:"item_status"`
	// HasModel indicates whether the item has models.
	HasModel bool `json:"has_model"`
	// PromotionID is the identifier of the promotion.
	PromotionID uint64 `json:"promotion_id"`
	// VideoInfo is the video information for the item.
	VideoInfo []ItemVideo `json:"video_info"`
	// Brand is the brand information for the item.
	Brand ItemBrand `json:"brand"`
	// ItemDangerous is the dangerous information for the item.
	ItemDangerous int `json:"item_dangerous"`
}

func (s *ProductServiceOp) GetItemBaseInfo(sid uint64, itemIDList []uint64, tok string) (*GetItemBaseInfoResponse, error) {
	path := "/product/get_item_base_info"
	opt := GetItemBaseInfoRequest{ItemIDList: itemIDList}
	resp := new(GetItemBaseInfoResponse)
	err := s.client.WithShop(sid, tok).Get(path, resp, opt)
	return resp, err
}

// Variation and Model related structs

type InitTierVariationRequest struct {
	// ItemID is the identifier of the item.
	ItemID uint64 `json:"item_id" mandatory:"true"`
	// TierVariation is the list of tier variations.
	TierVariation []TierVariation `json:"tier_variation" mandatory:"true"`
	// Model is the list of models.
	Model []InitTierVariationModel `json:"model" mandatory:"true"`
}

type InitTierVariationModel struct {
	// TierIndex is the list of indices of the tier variations.
	TierIndex []int `json:"tier_index"`
	// OriginalPrice is the original price of the model.
	OriginalPrice float64 `json:"original_price"`
	// ModelSKU is the SKU for the model.
	ModelSKU string `json:"model_sku"`
	// SellerStock is the stock information for the model.
	SellerStock []SellerStock `json:"seller_stock"`
}

type InitTierVariationResponse struct {
	BaseResponse
	Response InitTierVariationResponseData `json:"response"`
}

type InitTierVariationResponseData struct {
	// ItemID is the identifier of the item.
	ItemID uint64 `json:"item_id"`
	// TierVariation is the list of tier variations.
	TierVariation []InitTierVariationResponseDataTierVariation `json:"tier_variation"`
	// Model is the list of models.
	Model []Model `json:"model"`
}

type InitTierVariationResponseDataTierVariation struct {
	// Name is the name of the tier variation.
	Name string `json:"name"`
	// OptionList is the list of options for the tier variation.
	OptionList []InitTierVariationResponseDataTierVariationTierVariationOption `json:"option_list"`
}

type InitTierVariationResponseDataTierVariationTierVariationOption struct {
	// Option is the name of the option.
	Option string `json:"option"`
	// Image is the image information for the option.
	Image InitTierVariationResponseDataTierVariationTierVariationOptionImage `json:"image"`
}

type InitTierVariationResponseDataTierVariationTierVariationOptionImage struct {
	// ImageID is the identifier of the image.
	ImageID string `json:"image_id"`
	// ImageURL is the URL of the image.
	ImageURL string `json:"image_url"`
}

func (s *ProductServiceOp) InitTierVariation(sid uint64, req InitTierVariationRequest, tok string) (*InitTierVariationResponse, error) {
	path := "/product/init_tier_variation"
	resp := new(InitTierVariationResponse)
	reqMap, err := StructToMap(req)
	if err != nil {
		return nil, err
	}
	err = s.client.WithShop(sid, tok).Post(path, reqMap, resp)
	return resp, err
}

type UpdateTierVariationRequest struct {
	// ItemID is the identifier of the item.
	ItemID uint64 `json:"item_id" mandatory:"true"`
	// TierVariation is the list of tier variations.
	TierVariation []TierVariation `json:"tier_variation" mandatory:"true"`
}

type UpdateTierVariationResponse struct {
	BaseResponse
}

func (s *ProductServiceOp) UpdateTierVariation(sid uint64, req UpdateTierVariationRequest, tok string) (*UpdateTierVariationResponse, error) {
	path := "/product/update_tier_variation"
	resp := new(UpdateTierVariationResponse)
	reqMap, err := StructToMap(req)
	if err != nil {
		return nil, err
	}
	err = s.client.WithShop(sid, tok).Post(path, reqMap, resp)
	return resp, err
}

type GetModelListRequest struct {
	// ItemID is the identifier of the item.
	ItemID uint64 `url:"item_id" mandatory:"true"`
}

type GetModelListResponse struct {
	BaseResponse
	Response GetModelListResponseData `json:"response"`
}

type GetModelListResponseData struct {
	// TierVariation is the list of tier variations.
	TierVariation []TierVariation `json:"tier_variation"`
	// Model is the list of models.
	Model []Model `json:"model"`
}

func (s *ProductServiceOp) GetModelList(sid uint64, itemID uint64, tok string) (*GetModelListResponse, error) {
	path := "/product/get_model_list"
	opt := GetModelListRequest{ItemID: itemID}
	resp := new(GetModelListResponse)
	err := s.client.WithShop(sid, tok).Get(path, resp, opt)
	return resp, err
}

type AddModelRequest struct {
	// ItemID is the identifier of the item.
	ItemID uint64 `json:"item_id" mandatory:"true"`
	// ModelList is the list of models to add.
	ModelList []AddModelRequestModel `json:"model_list" mandatory:"true"`
}

type AddModelRequestModel struct {
	// TierIndex is the list of indices of the tier variations.
	TierIndex []int `json:"tier_index"`
	// OriginalPrice is the original price of the model.
	OriginalPrice float64 `json:"original_price"`
	// ModelSKU is the SKU for the model.
	ModelSKU string `json:"model_sku"`
	// SellerStock is the stock information for the model.
	SellerStock []SellerStock `json:"seller_stock"`
}

type AddModelResponse struct {
	BaseResponse
	Response AddModelResponseData `json:"response"`
}

type AddModelResponseData struct {
	// Model is the list of added models.
	Model []Model `json:"model"`
}

func (s *ProductServiceOp) AddModel(sid uint64, req AddModelRequest, tok string) (*AddModelResponse, error) {
	path := "/product/add_model"
	resp := new(AddModelResponse)
	reqMap, err := StructToMap(req)
	if err != nil {
		return nil, err
	}
	err = s.client.WithShop(sid, tok).Post(path, reqMap, resp)
	return resp, err
}

func (s *ProductServiceOp) DeleteModel(sid uint64, itemID uint64, modelID uint64, tok string) (*BaseResponse, error) {
	path := "/product/delete_model"
	req := map[string]interface{}{"item_id": itemID, "model_id": modelID}
	resp := new(BaseResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

type UpdateModelRequest struct {
	// ItemID is the identifier of the item.
	ItemID uint64 `json:"item_id" mandatory:"true"`
	// ModelList is the list of models to update.
	ModelList []UpdateModelRequestModel `json:"model_list" mandatory:"true"`
}

type UpdateModelRequestModel struct {
	// ModelID is the identifier of the model.
	ModelID uint64 `json:"model_id" mandatory:"true"`
	// ModelSKU is the SKU for the model.
	ModelSKU string `json:"model_sku"`
	// OriginalPrice is the original price of the model.
	OriginalPrice float64 `json:"original_price"`
	// SellerStock is the stock information for the model.
	SellerStock []SellerStock `json:"seller_stock"`
}

type UpdateModelResponse struct {
	BaseResponse
}

func (s *ProductServiceOp) UpdateModel(sid uint64, req UpdateModelRequest, tok string) (*UpdateModelResponse, error) {
	path := "/product/update_model"
	resp := new(UpdateModelResponse)
	reqMap, err := StructToMap(req)
	if err != nil {
		return nil, err
	}
	err = s.client.WithShop(sid, tok).Post(path, reqMap, resp)
	return resp, err
}

// Price and Stock related structs

type UpdatePriceRequest struct {
	// ItemID is the identifier of the item.
	ItemID uint64 `json:"item_id" mandatory:"true"`
	// PriceList is the list of model prices to update.
	PriceList []UpdatePriceRequestPrice `json:"price_list" mandatory:"true"`
}

type UpdatePriceRequestPrice struct {
	// ModelID is the identifier of the model.
	ModelID uint64 `json:"model_id"`
	// OriginalPrice is the new original price of the model.
	OriginalPrice float64 `json:"original_price"`
}

type UpdatePriceResponse struct {
	BaseResponse
	Response UpdatePriceResponseData `json:"response"`
}

type UpdatePriceResponseData struct {
	// FailureList is the list of models that failed to update price.
	FailureList []UpdatePriceResponseDataFail `json:"failure_list"`
	// SuccessList is the list of models that successfully updated price.
	SuccessList []UpdatePriceResponseDataSuccess `json:"success_list"`
}

type UpdatePriceResponseDataFail struct {
	// ModelID is the identifier of the model.
	ModelID uint64 `json:"model_id"`
	// FailedReason is the reason for the failure.
	FailedReason string `json:"failed_reason"`
}

type UpdatePriceResponseDataSuccess struct {
	// ModelID is the identifier of the model.
	ModelID uint64 `json:"model_id"`
	// OriginalPrice is the new original price of the model.
	OriginalPrice float64 `json:"original_price"`
}

func (s *ProductServiceOp) UpdatePrice(sid uint64, req UpdatePriceRequest, tok string) (*UpdatePriceResponse, error) {
	path := "/product/update_price"
	resp := new(UpdatePriceResponse)
	reqMap, err := StructToMap(req)
	if err != nil {
		return nil, err
	}
	err = s.client.WithShop(sid, tok).Post(path, reqMap, resp)
	return resp, err
}

type UpdateStockRequest struct {
	// ItemID is the identifier of the item.
	ItemID uint64 `json:"item_id" mandatory:"true"`
	// StockList is the list of model stocks to update.
	StockList []UpdateStockRequestStock `json:"stock_list" mandatory:"true"`
}

type UpdateStockRequestStock struct {
	// ModelID is the identifier of the model.
	ModelID uint64 `json:"model_id"`
	// SellerStock is the new stock information for the model.
	SellerStock []SellerStock `json:"seller_stock"`
}

type UpdateStockResponse struct {
	BaseResponse
	Response UpdateStockResponseData `json:"response"`
}

type UpdateStockResponseData struct {
	// FailureList is the list of models that failed to update stock.
	FailureList []UpdateStockResponseDataFail `json:"failure_list"`
	// SuccessList is the list of models that successfully updated stock.
	SuccessList []UpdateStockResponseDataSuccess `json:"success_list"`
}

type UpdateStockResponseDataFail struct {
	// ModelID is the identifier of the model.
	ModelID uint64 `json:"model_id"`
	// FailedReason is the reason for the failure.
	FailedReason string `json:"failed_reason"`
}

type UpdateStockResponseDataSuccess struct {
	// ModelID is the identifier of the model.
	ModelID uint64 `json:"model_id"`
	// NormalStock is the new normal stock level.
	NormalStock int `json:"normal_stock"`
}

func (s *ProductServiceOp) UpdateStock(sid uint64, req UpdateStockRequest, tok string) (*UpdateStockResponse, error) {
	path := "/product/update_stock"
	resp := new(UpdateStockResponse)
	reqMap, err := StructToMap(req)
	if err != nil {
		return nil, err
	}
	err = s.client.WithShop(sid, tok).Post(path, reqMap, resp)
	return resp, err
}

// Other structs

type CategoryRecommendRequest struct {
	// ItemName is the name of the item. Shopee's recommendation system will provide recommended category_id based on the item_name.
	ItemName string `url:"item_name" mandatory:"true"`
	// ProductCoverImage is the cover image of the product. Shopee's recommendation system will provide recommended category_id based on the product_cover_image.
	ProductCoverImage string `url:"product_cover_image"`
}

type CategoryRecommendResponse struct {
	BaseResponse
	Response CategoryRecommendResponseData `json:"response"`
}

type CategoryRecommendResponseData struct {
	// The list of recommended category IDs.
	CategoryID []uint64 `json:"category_id"`
}

func (s *ProductServiceOp) CategoryRecommend(sid uint64, itemName string, tok string) (*CategoryRecommendResponse, error) {
	path := "/product/category_recommend"
	opt := CategoryRecommendRequest{ItemName: itemName}
	resp := new(CategoryRecommendResponse)
	err := s.client.WithShop(sid, tok).Get(path, resp, opt)
	return resp, err
}

type GetItemPromotionRequest struct {
	// The list of item IDs.
	ItemIDList []uint64 `url:"item_id_list" mandatory:"true"`
}

type GetItemPromotionResponse struct {
	BaseResponse
	Response GetItemPromotionResponseData `json:"response"`
}

type GetItemPromotionResponseData struct {
	// The list of items and their promotions.
	SuccessList []ItemPromotion `json:"success_list"`
}

type ItemPromotion struct {
	// The identifier of the item.
	ItemID uint64 `json:"item_id"`
	// The list of promotions for the item.
	Promotion []Promotion `json:"promotion"`
}

type Promotion struct {
	// The type of the promotion.
	PromotionType string `json:"promotion_type"`
	// The identifier of the promotion.
	PromotionID uint64 `json:"promotion_id"`
	// The identifier of the model.
	ModelID uint64 `json:"model_id"`
	// The start time of the promotion.
	StartTime uint64 `json:"start_time"`
	// The end time of the promotion.
	EndTime uint64 `json:"end_time"`
	// The price information for the promotion.
	PromotionPriceInfo []PromotionPriceInfo `json:"promotion_price_info"`
	// The stock information for the promotion.
	PromotionStockInfoV2 []PromotionStockInfoV2 `json:"promotion_stock_info_v2"`
	// The staging status of the promotion.
	PromotionStaging string `json:"promotion_staging"`
}

type PromotionPriceInfo struct {
	// The promotional price of the item.
	PromotionPrice float64 `json:"promotion_price"`
}

type PromotionStockInfoV2 struct {
	// The type of stock.
	StockType int `json:"stock_type"`
	// The identifier of the stock location.
	StockLocationID string `json:"stock_location_id"`
	// The reserved stock level.
	ReservedStock int `json:"reserved_stock"`
}

func (s *ProductServiceOp) GetItemPromotion(sid uint64, itemIDList []uint64, tok string) (*GetItemPromotionResponse, error) {
	path := "/product/get_item_promotion"
	opt := GetItemPromotionRequest{ItemIDList: itemIDList}
	resp := new(GetItemPromotionResponse)
	err := s.client.WithShop(sid, tok).Get(path, resp, opt)
	return resp, err
}

type GetItemListRequest struct {
	// The offset of the item list.
	Offset int `url:"offset"`
	// The number of items to retrieve per page.
	PageSize int `url:"page_size" mandatory:"true"`
	// The status of the items to retrieve.
	ItemStatus []ItemStatus `url:"item_status"`
	// The start time of the update time range.
	UpdateTimeFrom int `url:"update_time_from"`
	// The end time of the update time range.
	UpdateTimeTo int `url:"update_time_to"`
}

type GetItemListResponse struct {
	BaseResponse
	Response GetItemListResponseData `json:"response"`
}

type GetItemListResponseDataItem struct {
	// The identifier of the item.
	ItemID uint64 `json:"item_id"`
	// The status of the item.
	ItemStatus ItemStatus `json:"item_status"`
	// The update time of the item.
	UpdatedTime uint64 `json:"updated_time"`
	// The tags of the item.
	Tag struct {
		// Indicates whether the item is a kit.
		Kit bool `json:"kit"`
	} `json:"tag"`
}

type GetItemListResponseData struct {
	// The list of items.
	Item []GetItemListResponseDataItem `json:"item"`
	// The total count of items.
	TotalCount int `json:"total_count"`
	// Indicates whether there are more items to retrieve.
	HasNextPage bool `json:"has_next_page"`
	// The offset to be used for the next request to retrieve more items.
	NextOffset int `json:"next_offset"`
}

func (s *ProductServiceOp) GetItemList(sid uint64, opt GetItemListRequest, tok string) (*GetItemListResponse, error) {
	path := "/product/get_item_list"
	resp := new(GetItemListResponse)
	err := s.client.WithShop(sid, tok).Get(path, resp, opt)
	return resp, err
}

func (s *ProductServiceOp) BoostItem(sid uint64, itemIDList []uint64, tok string) (*BaseResponse, error) {
	path := "/product/boost_item"
	req := map[string]interface{}{"item_id_list": itemIDList}
	resp := new(BaseResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

type GetBoostedListResponse struct {
	BaseResponse
	Response GetBoostedListResponseData `json:"response"`
}

type GetBoostedListResponseData struct {
	// The list of boosted item IDs.
	ItemIDList []uint64 `json:"item_id_list"`
}

func (s *ProductServiceOp) GetBoostedList(sid uint64, tok string) (*GetBoostedListResponse, error) {
	path := "/product/get_boosted_list"
	resp := new(GetBoostedListResponse)
	err := s.client.WithShop(sid, tok).Get(path, resp, nil)
	return resp, err
}

type GetCommentRequest struct {
	// The identifier of the item.
	ItemID uint64 `url:"item_id"`
	// The identifier of the comment.
	CommentID uint64 `url:"comment_id"`
	// The cursor for the next page.
	Cursor string `url:"cursor"`
	// The number of comments to retrieve per page.
	PageSize int `url:"page_size"`
	// The reply state of the comment. Y or N.
	ReplyState string `url:"reply_state"` // Y or N
}

type GetCommentResponse struct {
	BaseResponse
	Response GetCommentResponseData `json:"response"`
}

type GetCommentResponseData struct {
	// The list of comments.
	CommentList []Comment `json:"comment_list"`
	// Indicates whether there are more comments to retrieve.
	More bool `json:"more"`
	// The cursor for the next page.
	NextCursor string `json:"next_cursor"`
}

type Comment struct {
	// The identifier of the comment.
	CommentID uint64 `json:"comment_id"`
	// The identifier of the item.
	ItemID uint64 `json:"item_id"`
	// The identifier of the model.
	ModelID uint64 `json:"model_id"`
	// The identifier of the buyer.
	BuyerUserID uint64 `json:"buyer_user_id"`
	// The order serial number.
	OrderSN string `json:"order_sn"`
	// The rating star of the comment.
	RatingStar int `json:"rating_star"`
	// The content of the comment.
	CommentText string `json:"comment"`
	// The create time of the comment.
	CreateTime uint64 `json:"create_time"`
	// The update time of the comment.
	UpdateTime uint64 `json:"update_time"`
	// The reply to the comment.
	Reply *Reply `json:"comment_reply"`
	// The images of the comment.
	ImageURL []string `json:"image_url"`
	// Indicates whether the comment is hidden.
	IsHidden bool `json:"is_hidden"`
}

type Reply struct {
	// The content of the reply.
	ReplyText string `json:"reply"`
	// The create time of the reply.
	CreateTime uint64 `json:"create_time"`
	// The update time of the reply.
	UpdateTime uint64 `json:"update_time"`
}

func (s *ProductServiceOp) GetComment(sid uint64, opt GetCommentRequest, tok string) (*GetCommentResponse, error) {
	path := "/product/get_comment"
	resp := new(GetCommentResponse)
	err := s.client.WithShop(sid, tok).Get(path, resp, opt)
	return resp, err
}

type ReplyCommentRequest struct {
	// The identifier of the comment.
	CommentID uint64 `json:"comment_id" mandatory:"true"`
	// The content of the reply.
	ReplyText string `json:"comment_reply" mandatory:"true"`
}

func (s *ProductServiceOp) ReplyComment(sid uint64, req ReplyCommentRequest, tok string) (*BaseResponse, error) {
	path := "/product/reply_comment"
	resp := new(BaseResponse)
	reqMap, err := StructToMap(req)
	if err != nil {
		return nil, err
	}
	err = s.client.WithShop(sid, tok).Post(path, reqMap, resp)
	return resp, err
}

// Shared nested structs

type Model struct {
	// TierIndex is the list of indices of the tier variations.
	TierIndex []int `json:"tier_index"`
	// ModelID is the identifier of the model.
	ModelID uint64 `json:"model_id"`
	// ModelSKU is the SKU for the model.
	ModelSKU string `json:"model_sku"`
	// StockInfoV2 is the stock information for the model.
	StockInfoV2 []StockInfoV2 `json:"stock_info_v2"`
	// PriceInfo is the price information for the model.
	PriceInfo []PriceInfo `json:"price_info"`
	// PromotionID is the identifier of the promotion.
	PromotionID uint64 `json:"promotion_id"`
}

type StockInfoV2 struct {
	// StockType is the type of stock.
	StockType int `json:"stock_type"`
	// StockLocationID is the identifier of the stock location.
	StockLocationID string `json:"stock_location_id"`
	// NormalStock is the normal stock level.
	NormalStock int `json:"normal_stock"`
	// CurrentStock is the current stock level.
	CurrentStock int `json:"current_stock"`
	// ReservedStock is the reserved stock level.
	ReservedStock int `json:"reserved_stock"`
}

type PriceInfo struct {
	// Currency is the currency of the price.
	Currency string `json:"currency"`
	// OriginalPrice is the original price.
	OriginalPrice float64 `json:"original_price"`
	// CurrentPrice is the current price.
	CurrentPrice float64 `json:"current_price"`
	// InflatedPriceOfOriginalPrice is the inflated price of the original price.
	InflatedPriceOfOriginalPrice float64 `json:"inflated_price_of_original_price"`
	// InflatedPriceOfCurrentPrice is the inflated price of the current price.
	InflatedPriceOfCurrentPrice float64 `json:"inflated_price_of_current_price"`
	// SipItemPrice is the SIP item price.
	SipItemPrice float64 `json:"sip_item_price"`
	// SipItemPriceSource is the source of the SIP item price.
	SipItemPriceSource string `json:"sip_item_price_source"`
}

type ItemAttribute struct {
	// AttributeID is the identifier of the attribute.
	AttributeID uint64 `json:"attribute_id"`
	// AttributeValueList is the list of attribute values.
	AttributeValueList []ItemAttributeValue `json:"attribute_value_list"`
}

type ItemAttributeValue struct {
	// ValueId is the identifier of the attribute value.
	ValueId uint64 `json:"value_id"`
	// OriginalValueName is the original name of the attribute value.
	OriginalValueName string `json:"original_value_name"`
	// ValueUnit is the unit of the attribute value.
	ValueUnit string `json:"value_unit"`
}

type ItemImage struct {
	// ImageIDList is the list of image identifiers.
	ImageIDList []string `json:"image_id_list"`
	// ImageURLList is the list of image URLs.
	ImageURLList []string `json:"image_url_list"`
}

type Dimension struct {
	// PackageHeight is the height of the package.
	PackageHeight int `json:"package_height"`
	// PackageLength is the length of the package.
	PackageLength int `json:"package_length"`
	// PackageWidth is the width of the package.
	PackageWidth int `json:"package_width"`
}

type LogisticInfo struct {
	// LogisticID is the identifier of the logistic channel.
	LogisticID uint64 `json:"logistic_id"`
	// LogisticName is the name of the logistic channel.
	LogisticName string `json:"logistic_name"`
	// Enabled indicates whether the logistic channel is enabled.
	Enabled bool `json:"enabled"`
	// ShippingFee is the shipping fee for the logistic channel.
	ShippingFee float64 `json:"shipping_fee"`
	// SizeID is the identifier of the size.
	SizeID uint64 `json:"size_id"`
	// IsFree indicates whether the shipping is free.
	IsFree bool `json:"is_free"`
	// EstimatedShippingFee is the estimated shipping fee.
	EstimatedShippingFee float64 `json:"estimated_shipping_fee"`
}

type ItemPreOrder struct {
	// IsPreOrder indicates whether the item is a pre-order.
	IsPreOrder bool `json:"is_pre_order"`
	// DaysToShip is the number of days to ship the pre-order item.
	DaysToShip *int `json:"days_to_ship,omitempty"`
}

type ItemWholesale struct {
	// MinCount is the minimum count for wholesale.
	MinCount int `json:"min_count"`
	// MaxCount is the maximum count for wholesale.
	MaxCount int `json:"max_count"`
	// UnitPrice is the unit price for wholesale.
	UnitPrice float64 `json:"unit_price"`
	// InflatedPriceOfUnitPrice is the inflated unit price for wholesale.
	InflatedPriceOfUnitPrice float64 `json:"inflated_price_of_unit_price"`
}

type ItemBrand struct {
	// BrandID is the identifier of the brand.
	BrandID uint64 `json:"brand_id"`
	// OriginalBrandName is the original name of the brand.
	OriginalBrandName string `json:"original_brand_name"`
}

type SellerStock struct {
	// LocationID is the identifier of the stock location.
	LocationID *string `json:"location_id"`
	// Stock is the stock level.
	Stock int32 `json:"stock"`
}

type ItemVideo struct {
	// VideoURL is the URL of the video.
	VideoURL string `json:"video_url"`
	// ThumbnailURL is the URL of the thumbnail.
	ThumbnailURL string `json:"thumbnail_url"`
	// Duration is the duration of the video.
	Duration int `json:"duration"`
}

type TierVariation struct {
	// Name is the name of the tier variation.
	Name string `json:"name"`
	// OptionList is the list of options for the tier variation.
	OptionList []TierVariationOption `json:"option_list"`
}

type TierVariationOption struct {
	// Option is the name of the option.
	Option string `json:"option"`
	// Image is the image information for the option.
	Image *TierVariationOptionImage `json:"image"`
}

type TierVariationOptionImage struct {
	// ImageID is the identifier of the image.
	ImageID string `json:"image_id"`
}
