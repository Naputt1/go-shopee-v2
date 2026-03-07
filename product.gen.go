package goshopee

type ProductService interface {
	// GetCategory Get category tree data. More detail please check https://open.shopee.com/developer-guide/209
	// https://open.shopee.com/documents/v2/v2.product.get_category?module=89&type=1
	GetCategory(sid uint64, opt GetCategoryRequest, tok string) (*GetCategoryResponse, error)
	// GetAttributeTree Get the attribute tree for categories
	// https://open.shopee.com/documents/v2/v2.product.get_attribute_tree?module=89&type=1
	GetAttributeTree(sid uint64, opt GetAttributeTreeRequest, tok string) (*GetAttributeTreeResponse, error)
	// GetBrandList Get the brand data of a leaf category. More detail please check: https://open.shopee.com/developer-guide/209
	// https://open.shopee.com/documents/v2/v2.product.get_brand_list?module=89&type=1
	GetBrandList(sid uint64, opt GetBrandListRequest, tok string) (*GetBrandListResponse, error)
	// GetItemLimit Get item upload control.
	// https://open.shopee.com/documents/v2/v2.product.get_item_limit?module=89&type=1
	GetItemLimit(sid uint64, opt GetItemLimitRequest, tok string) (*GetItemLimitResponse, error)
	// GetItemList Use this call to get a list of items.
	// https://open.shopee.com/documents/v2/v2.product.get_item_list?module=89&type=1
	GetItemList(sid uint64, opt GetItemListRequest, tok string) (*GetItemListResponse, error)
	// GetItemBaseInfo Use this api to get basic info of item by item_id list.
	// https://open.shopee.com/documents/v2/v2.product.get_item_base_info?module=89&type=1
	GetItemBaseInfo(sid uint64, opt GetItemBaseInfoRequest, tok string) (*GetItemBaseInfoResponse, error)
	// GetItemExtraInfo Use this api to get extra info of item by item_id list.
	// https://open.shopee.com/documents/v2/v2.product.get_item_extra_info?module=89&type=1
	GetItemExtraInfo(sid uint64, opt GetItemExtraInfoRequest, tok string) (*GetItemExtraInfoResponse, error)
	// AddItem Add a new item.
	// https://open.shopee.com/documents/v2/v2.product.add_item?module=89&type=1
	AddItem(sid uint64, req AddItemRequest, tok string) (*AddItemResponse, error)
	// UpdateItem Update item.
	// https://open.shopee.com/documents/v2/v2.product.update_item?module=89&type=1
	UpdateItem(sid uint64, req UpdateItemRequest, tok string) (*UpdateItemResponse, error)
	// DeleteItem Use this call to delete a product item.
	// https://open.shopee.com/documents/v2/v2.product.delete_item?module=89&type=1
	DeleteItem(sid uint64, req DeleteItemRequest, tok string) (*DeleteItemResponse, error)
	// InitTierVariation This API allows you to update the tier structure of a product. Defining only color creates one tier, while color + size creates two tiers (maximum supported). Supported changes include: no tier ↔ one/two tiers, one tier ↔ two/no tier, and two tiers ↔ one/no tier. For details, see Developer Guide.  Please wait at least 5 seconds after creating an item before creating variants, as processing may be delayed.
	// https://open.shopee.com/documents/v2/v2.product.init_tier_variation?module=89&type=1
	InitTierVariation(sid uint64, req InitTierVariationRequest, tok string) (*InitTierVariationResponse, error)
	// UpdateTierVariation This api can only be used without changing the tier structure, you can add options, delete options, and update the option image by this api. More detail please check: https://open.shopee.com/developer-guide/219
	// https://open.shopee.com/documents/v2/v2.product.update_tier_variation?module=89&type=1
	UpdateTierVariation(sid uint64, req UpdateTierVariationRequest, tok string) (*UpdateTierVariationResponse, error)
	// GetModelList Get model list of an item.
	// https://open.shopee.com/documents/v2/v2.product.get_model_list?module=89&type=1
	GetModelList(sid uint64, opt GetModelListRequest, tok string) (*GetModelListResponse, error)
	// AddModel Add model. More detail please check: https://open.shopee.com/developer-guide/219
	// https://open.shopee.com/documents/v2/v2.product.add_model?module=89&type=1
	AddModel(sid uint64, req AddModelRequest, tok string) (*AddModelResponse, error)
	// UpdateModel Update seller sku/ pre order/ model status for model.
	// https://open.shopee.com/documents/v2/v2.product.update_model?module=89&type=1
	UpdateModel(sid uint64, req UpdateModelRequest, tok string) (*UpdateModelResponse, error)
	// DeleteModel Delete item model.
	// https://open.shopee.com/documents/v2/v2.product.delete_model?module=89&type=1
	DeleteModel(sid uint64, req DeleteModelRequest, tok string) (*DeleteModelResponse, error)
	// UnlistItem Unlist item.
	// https://open.shopee.com/documents/v2/v2.product.unlist_item?module=89&type=1
	UnlistItem(sid uint64, req UnlistItemRequest, tok string) (*UnlistItemResponse, error)
	// UpdatePrice Update price.
	// https://open.shopee.com/documents/v2/v2.product.update_price?module=89&type=1
	UpdatePrice(sid uint64, req UpdatePriceRequest, tok string) (*UpdatePriceResponse, error)
	// UpdateStock Use this API to update one item_id for each call, but still can support updating multiple model_ids stock of the same item_id (If you need batch modification, please call multiple times)This API will update only "seller_stock".Whenever there is a promotion ongoing or upcoming, the total stock must be larger than or equal to real-time “reserved_stock” promotion stock (Please check v2.get_item_promotion API for more details). Items that are deleted will not be allowed to modify stock.
	// https://open.shopee.com/documents/v2/v2.product.update_stock?module=89&type=1
	UpdateStock(sid uint64, req UpdateStockRequest, tok string) (*UpdateStockResponse, error)
	// BoostItem Boost item.
	// https://open.shopee.com/documents/v2/v2.product.boost_item?module=89&type=1
	BoostItem(sid uint64, req BoostItemRequest, tok string) (*BoostItemResponse, error)
	// GetBoostedList Get boosted item list.
	// https://open.shopee.com/documents/v2/v2.product.get_boosted_list?module=89&type=1
	GetBoostedList(sid uint64, tok string) (*GetBoostedListResponse, error)
	// GetItemPromotion Get item promotion info.
	// https://open.shopee.com/documents/v2/v2.product.get_item_promotion?module=89&type=1
	GetItemPromotion(sid uint64, opt GetItemPromotionRequest, tok string) (*GetItemPromotionResponse, error)
	// UpdateSipItemPrice Update sip item price.
	// https://open.shopee.com/documents/v2/v2.product.update_sip_item_price?module=89&type=1
	UpdateSipItemPrice(sid uint64, req UpdateSipItemPriceRequest, tok string) (*UpdateSipItemPriceResponse, error)
	// SearchItem Use this call to search item.
	// https://open.shopee.com/documents/v2/v2.product.search_item?module=89&type=1
	SearchItem(sid uint64, opt SearchItemRequest, tok string) (*SearchItemResponse, error)
	// GetComment Use this api to get comment by shop_id, item_id, or comment_id, get up to 1000 comments.
	// https://open.shopee.com/documents/v2/v2.product.get_comment?module=89&type=1
	GetComment(sid uint64, opt GetCommentRequest, tok string) (*GetCommentResponse, error)
	// ReplyComment Use this api to reply comments from buyers in batch.
	// https://open.shopee.com/documents/v2/v2.product.reply_comment?module=89&type=1
	ReplyComment(sid uint64, req ReplyCommentRequest, tok string) (*ReplyCommentResponse, error)
	// CategoryRecommend Recommend category by item name.
	// https://open.shopee.com/documents/v2/v2.product.category_recommend?module=89&type=1
	CategoryRecommend(sid uint64, opt CategoryRecommendRequest, tok string) (*CategoryRecommendResponse, error)
	// RegisterBrand Use this call to register a brand.
	// https://open.shopee.com/documents/v2/v2.product.register_brand?module=89&type=1
	RegisterBrand(sid uint64, req RegisterBrandRequest, tok string) (*RegisterBrandResponse, error)
	// GetRecommendAttribute Get recommend attributes.
	// https://open.shopee.com/documents/v2/v2.product.get_recommend_attribute?module=89&type=1
	GetRecommendAttribute(sid uint64, opt GetRecommendAttributeRequest, tok string) (*GetRecommendAttributeResponse, error)
	// GetWeightRecommendation Get recommended weight. Now only BR shop support to use this api to get recommended weight.
	// https://open.shopee.com/documents/v2/v2.product.get_weight_recommendation?module=89&type=1
	GetWeightRecommendation(sid uint64, req GetWeightRecommendationRequest, tok string) (*GetWeightRecommendationResponse, error)
	// GetSizeChartList Get new size chat list. Now only support local shop to use new size chart.
	// https://open.shopee.com/documents/v2/v2.product.get_size_chart_list?module=89&type=1
	GetSizeChartList(sid uint64, opt GetSizeChartListRequest, tok string) (*GetSizeChartListResponse, error)
	// GetSizeChartDetail Get new size chart detail. Now only local shop support to use this api to get new size chart detail.
	// https://open.shopee.com/documents/v2/v2.product.get_size_chart_detail?module=89&type=1
	GetSizeChartDetail(sid uint64, opt GetSizeChartDetailRequest, tok string) (*GetSizeChartDetailResponse, error)
	// GetItemViolationInfo get item violation info
	// https://open.shopee.com/documents/v2/v2.product.get_item_violation_info?module=89&type=1
	GetItemViolationInfo(sid uint64, req GetItemViolationInfoRequest, tok string) (*GetItemViolationInfoResponse, error)
	// GetVariations Get the standardized tier variation defined by Shopee, which is currently a three-layer tree structure.
	// The top layer is variations, the second layer is groups, groups are used to divide options, and the third layer is options.
	// https://open.shopee.com/documents/v2/v2.product.get_variations?module=89&type=1
	GetVariations(sid uint64, opt GetVariationsRequest, tok string) (*GetVariationsResponse, error)
	// GetAllVehicleList Use this Open API to get all vehicle list.
	// https://open.shopee.com/documents/v2/v2.product.get_all_vehicle_list?module=89&type=1
	GetAllVehicleList(sid uint64, opt GetAllVehicleListRequest, tok string) (*GetAllVehicleListResponse, error)
	// GetVehicleListByCompatibilityDetail Use this Open API to get vehicle list by brand, model, year, and version.
	// https://open.shopee.com/documents/v2/v2.product.get_vehicle_list_by_compatibility_detail?module=89&type=1
	GetVehicleListByCompatibilityDetail(sid uint64, opt GetVehicleListByCompatibilityDetailRequest, tok string) (*GetVehicleListByCompatibilityDetailResponse, error)
	// GetItemContentDiagnosisResult Get the content quality details (including content quality level, content issues, and system suggestions) for specific product list.
	// https://open.shopee.com/documents/v2/v2.product.get_item_content_diagnosis_result?module=89&type=1
	GetItemContentDiagnosisResult(sid uint64, req GetItemContentDiagnosisResultRequest, tok string) (*GetItemContentDiagnosisResultResponse, error)
	// GetItemListByContentDiagnosis Query the list of products and their content quality details by content quality level or content issues.
	// https://open.shopee.com/documents/v2/v2.product.get_item_list_by_content_diagnosis?module=89&type=1
	GetItemListByContentDiagnosis(sid uint64, req GetItemListByContentDiagnosisRequest, tok string) (*GetItemListByContentDiagnosisResponse, error)
	// GetKitItemLimit Get the limit of Kit item.
	// https://open.shopee.com/documents/v2/v2.product.get_kit_item_limit?module=89&type=1
	GetKitItemLimit(sid uint64, opt GetKitItemLimitRequest, tok string) (*GetKitItemLimitResponse, error)
	// AddKitItem Create the kit item by selecting multiple items and setting main component and quantity per kit.
	// https://open.shopee.com/documents/v2/v2.product.add_kit_item?module=89&type=1
	AddKitItem(sid uint64, req AddKitItemRequest, tok string) (*AddKitItemResponse, error)
	// UpdateKitItem Update the kit basic information and kit components, only support adding kit variations and updating existing kit variation’s image, price, and model_sku, don’t support deleting existing kit variations and updating the items, main component and quantity per kit of existing kit variations.
	// https://open.shopee.com/documents/v2/v2.product.update_kit_item?module=89&type=1
	UpdateKitItem(sid uint64, req UpdateKitItemRequest, tok string) (*UpdateKitItemResponse, error)
	// GetKitItemInfo Get the kit basic information and kit components.
	// https://open.shopee.com/documents/v2/v2.product.get_kit_item_info?module=89&type=1
	GetKitItemInfo(sid uint64, opt GetKitItemInfoRequest, tok string) (*GetKitItemInfoResponse, error)
	// GetSspList Get the list of SSP with main info as ssp_id, product_name, gtin, oem codes.
	// https://open.shopee.com/documents/v2/v2.product.get_ssp_list?module=89&type=1
	GetSspList(sid uint64, opt GetSspListRequest, tok string) (*GetSspListResponse, error)
	// GetSspInfo Get the SSP details with all available info, searching through ssp_id, gtin or oem.
	// https://open.shopee.com/documents/v2/v2.product.get_ssp_info?module=89&type=1
	GetSspInfo(sid uint64, req GetSspInfoRequest, tok string) (*GetSspInfoResponse, error)
	// AddSspItem Create product based on SSP.
	// https://open.shopee.com/documents/v2/v2.product.add_ssp_item?module=89&type=1
	AddSspItem(sid uint64, req AddSspItemRequest, tok string) (*AddSspItemResponse, error)
	// LinkSsp Link the existing product to SSP, replacing its info for the ones from the SSP.
	// https://open.shopee.com/documents/v2/v2.product.link_ssp?module=89&type=1
	LinkSsp(sid uint64, req LinkSspRequest, tok string) (*LinkSspResponse, error)
	// UnlinkSsp Unlink the existing product from SSP, keeping the existing info.
	// https://open.shopee.com/documents/v2/v2.product.unlink_ssp?module=89&type=1
	UnlinkSsp(sid uint64, req UnlinkSspRequest, tok string) (*UnlinkSspResponse, error)
	// GetAitemByPitemId Get the list of A Items under SIP Affiliate Shop corresponding to P Items under SIP Primary Shop.
	// https://open.shopee.com/documents/v2/v2.product.get_aitem_by_pitem_id?module=89&type=1
	GetAitemByPitemId(sid uint64, opt GetAitemByPitemIdRequest, tok string) (*GetAitemByPitemIdResponse, error)
	// SearchAttributeValueList this api is for searching attribute value list for attribute with support_search_value flag
	// https://open.shopee.com/documents/v2/v2.product.search_attribute_value_list?module=89&type=1
	SearchAttributeValueList(sid uint64, req SearchAttributeValueListRequest, tok string) (*SearchAttributeValueListResponse, error)
	// GetMainItemList get main item by direct item.
	// https://open.shopee.com/documents/v2/v2.product.get_main_item_list?module=89&type=1
	GetMainItemList(sid uint64, opt GetMainItemListRequest, tok string) (*GetMainItemListResponse, error)
	// GetDirectItemList get direct item by main item.
	// https://open.shopee.com/documents/v2/v2.product.get_direct_item_list?module=89&type=1
	GetDirectItemList(sid uint64, opt GetDirectItemListRequest, tok string) (*GetDirectItemListResponse, error)
	// GetDirectShopRecommendedPrice get recommend price for direct shop.
	// https://open.shopee.com/documents/v2/v2.product.get_direct_shop_recommended_price?module=89&type=1
	GetDirectShopRecommendedPrice(sid uint64, opt GetDirectShopRecommendedPriceRequest, tok string) (*GetDirectShopRecommendedPriceResponse, error)
	// GetProductCertificationRule Get product certification rule
	// https://open.shopee.com/documents/v2/v2.product.get_product_certification_rule?module=89&type=1
	GetProductCertificationRule(sid uint64, req GetProductCertificationRuleRequest, tok string) (*GetProductCertificationRuleResponse, error)
	// PublishItemToOutletShop
	// https://open.shopee.com/documents/v2/v2.product.publish_item_to_outlet_shop?module=89&type=1
	PublishItemToOutletShop(sid uint64, tok string) (*PublishItemToOutletShopResponse, error)
	// GetMartItemMappingById Get the mapping information between a Mart item and its corresponding outlet item by item ID.
	// https://open.shopee.com/documents/v2/v2.product.get_mart_item_mapping_by_id?module=89&type=1
	GetMartItemMappingById(sid uint64, req GetMartItemMappingByIdRequest, tok string) (*GetMartItemMappingByIdResponse, error)
	// SearchUnpackagedModelList Use this API to retrieve Unpackaged SKU ID information for items that toggle on logistics channel 30029.
	// https://open.shopee.com/documents/v2/v2.product.search_unpackaged_model_list?module=89&type=1
	SearchUnpackagedModelList(sid uint64, req SearchUnpackagedModelListRequest, tok string) (*SearchUnpackagedModelListResponse, error)
	// GenerateKitImage This API generates a single consolidated image by combining the cover images of all selected items. It is typically used to create a unified product display image for kits or bundles.
	// https://open.shopee.com/documents/v2/v2.product.generate_kit_image?module=89&type=1
	GenerateKitImage(sid uint64, req GenerateKitImageRequest, tok string) (*GenerateKitImageResponse, error)
}

type ProductServiceOp[T any] struct {
	client *Client[T]
}

func (s *ProductServiceOp[T]) GetCategory(sid uint64, opt GetCategoryRequest, tok string) (*GetCategoryResponse, error) {
	path := "/product/get_category"
	resp := new(GetCategoryResponse)
	err := s.client.WithShop(sid, tok).Get(path, resp, opt)
	return resp, err
}

func (s *ProductServiceOp[T]) GetAttributeTree(sid uint64, opt GetAttributeTreeRequest, tok string) (*GetAttributeTreeResponse, error) {
	path := "/product/get_attribute_tree"
	resp := new(GetAttributeTreeResponse)
	err := s.client.WithShop(sid, tok).Get(path, resp, opt)
	return resp, err
}

func (s *ProductServiceOp[T]) GetBrandList(sid uint64, opt GetBrandListRequest, tok string) (*GetBrandListResponse, error) {
	path := "/product/get_brand_list"
	resp := new(GetBrandListResponse)
	err := s.client.WithShop(sid, tok).Get(path, resp, opt)
	return resp, err
}

func (s *ProductServiceOp[T]) GetItemLimit(sid uint64, opt GetItemLimitRequest, tok string) (*GetItemLimitResponse, error) {
	path := "/product/get_item_limit"
	resp := new(GetItemLimitResponse)
	err := s.client.WithShop(sid, tok).Get(path, resp, opt)
	return resp, err
}

func (s *ProductServiceOp[T]) GetItemList(sid uint64, opt GetItemListRequest, tok string) (*GetItemListResponse, error) {
	path := "/product/get_item_list"
	resp := new(GetItemListResponse)
	err := s.client.WithShop(sid, tok).Get(path, resp, opt)
	return resp, err
}

func (s *ProductServiceOp[T]) GetItemBaseInfo(sid uint64, opt GetItemBaseInfoRequest, tok string) (*GetItemBaseInfoResponse, error) {
	path := "/product/get_item_base_info"
	resp := new(GetItemBaseInfoResponse)
	err := s.client.WithShop(sid, tok).Get(path, resp, opt)
	return resp, err
}

func (s *ProductServiceOp[T]) GetItemExtraInfo(sid uint64, opt GetItemExtraInfoRequest, tok string) (*GetItemExtraInfoResponse, error) {
	path := "/product/get_item_extra_info"
	resp := new(GetItemExtraInfoResponse)
	err := s.client.WithShop(sid, tok).Get(path, resp, opt)
	return resp, err
}

func (s *ProductServiceOp[T]) AddItem(sid uint64, req AddItemRequest, tok string) (*AddItemResponse, error) {
	path := "/product/add_item"
	resp := new(AddItemResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

func (s *ProductServiceOp[T]) UpdateItem(sid uint64, req UpdateItemRequest, tok string) (*UpdateItemResponse, error) {
	path := "/product/update_item"
	resp := new(UpdateItemResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

func (s *ProductServiceOp[T]) DeleteItem(sid uint64, req DeleteItemRequest, tok string) (*DeleteItemResponse, error) {
	path := "/product/delete_item"
	resp := new(DeleteItemResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

func (s *ProductServiceOp[T]) InitTierVariation(sid uint64, req InitTierVariationRequest, tok string) (*InitTierVariationResponse, error) {
	path := "/product/init_tier_variation"
	resp := new(InitTierVariationResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

func (s *ProductServiceOp[T]) UpdateTierVariation(sid uint64, req UpdateTierVariationRequest, tok string) (*UpdateTierVariationResponse, error) {
	path := "/product/update_tier_variation"
	resp := new(UpdateTierVariationResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

func (s *ProductServiceOp[T]) GetModelList(sid uint64, opt GetModelListRequest, tok string) (*GetModelListResponse, error) {
	path := "/product/get_model_list"
	resp := new(GetModelListResponse)
	err := s.client.WithShop(sid, tok).Get(path, resp, opt)
	return resp, err
}

func (s *ProductServiceOp[T]) AddModel(sid uint64, req AddModelRequest, tok string) (*AddModelResponse, error) {
	path := "/product/add_model"
	resp := new(AddModelResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

func (s *ProductServiceOp[T]) UpdateModel(sid uint64, req UpdateModelRequest, tok string) (*UpdateModelResponse, error) {
	path := "/product/update_model"
	resp := new(UpdateModelResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

func (s *ProductServiceOp[T]) DeleteModel(sid uint64, req DeleteModelRequest, tok string) (*DeleteModelResponse, error) {
	path := "/product/delete_model"
	resp := new(DeleteModelResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

func (s *ProductServiceOp[T]) UnlistItem(sid uint64, req UnlistItemRequest, tok string) (*UnlistItemResponse, error) {
	path := "/product/unlist_item"
	resp := new(UnlistItemResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

func (s *ProductServiceOp[T]) UpdatePrice(sid uint64, req UpdatePriceRequest, tok string) (*UpdatePriceResponse, error) {
	path := "/product/update_price"
	resp := new(UpdatePriceResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

func (s *ProductServiceOp[T]) UpdateStock(sid uint64, req UpdateStockRequest, tok string) (*UpdateStockResponse, error) {
	path := "/product/update_stock"
	resp := new(UpdateStockResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

func (s *ProductServiceOp[T]) BoostItem(sid uint64, req BoostItemRequest, tok string) (*BoostItemResponse, error) {
	path := "/product/boost_item"
	resp := new(BoostItemResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

func (s *ProductServiceOp[T]) GetBoostedList(sid uint64, tok string) (*GetBoostedListResponse, error) {
	path := "/product/get_boosted_list"
	resp := new(GetBoostedListResponse)
	err := s.client.WithShop(sid, tok).Get(path, resp, nil)
	return resp, err
}

func (s *ProductServiceOp[T]) GetItemPromotion(sid uint64, opt GetItemPromotionRequest, tok string) (*GetItemPromotionResponse, error) {
	path := "/product/get_item_promotion"
	resp := new(GetItemPromotionResponse)
	err := s.client.WithShop(sid, tok).Get(path, resp, opt)
	return resp, err
}

func (s *ProductServiceOp[T]) UpdateSipItemPrice(sid uint64, req UpdateSipItemPriceRequest, tok string) (*UpdateSipItemPriceResponse, error) {
	path := "/product/update_sip_item_price"
	resp := new(UpdateSipItemPriceResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

func (s *ProductServiceOp[T]) SearchItem(sid uint64, opt SearchItemRequest, tok string) (*SearchItemResponse, error) {
	path := "/product/search_item"
	resp := new(SearchItemResponse)
	err := s.client.WithShop(sid, tok).Get(path, resp, opt)
	return resp, err
}

func (s *ProductServiceOp[T]) GetComment(sid uint64, opt GetCommentRequest, tok string) (*GetCommentResponse, error) {
	path := "/product/get_comment"
	resp := new(GetCommentResponse)
	err := s.client.WithShop(sid, tok).Get(path, resp, opt)
	return resp, err
}

func (s *ProductServiceOp[T]) ReplyComment(sid uint64, req ReplyCommentRequest, tok string) (*ReplyCommentResponse, error) {
	path := "/product/reply_comment"
	resp := new(ReplyCommentResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

func (s *ProductServiceOp[T]) CategoryRecommend(sid uint64, opt CategoryRecommendRequest, tok string) (*CategoryRecommendResponse, error) {
	path := "/product/category_recommend"
	resp := new(CategoryRecommendResponse)
	err := s.client.WithShop(sid, tok).Get(path, resp, opt)
	return resp, err
}

func (s *ProductServiceOp[T]) RegisterBrand(sid uint64, req RegisterBrandRequest, tok string) (*RegisterBrandResponse, error) {
	path := "/product/register_brand"
	resp := new(RegisterBrandResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

func (s *ProductServiceOp[T]) GetRecommendAttribute(sid uint64, opt GetRecommendAttributeRequest, tok string) (*GetRecommendAttributeResponse, error) {
	path := "/product/get_recommend_attribute"
	resp := new(GetRecommendAttributeResponse)
	err := s.client.WithShop(sid, tok).Get(path, resp, opt)
	return resp, err
}

func (s *ProductServiceOp[T]) GetWeightRecommendation(sid uint64, req GetWeightRecommendationRequest, tok string) (*GetWeightRecommendationResponse, error) {
	path := "/product/get_weight_recommendation"
	resp := new(GetWeightRecommendationResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

func (s *ProductServiceOp[T]) GetSizeChartList(sid uint64, opt GetSizeChartListRequest, tok string) (*GetSizeChartListResponse, error) {
	path := "/product/get_size_chart_list"
	resp := new(GetSizeChartListResponse)
	err := s.client.WithShop(sid, tok).Get(path, resp, opt)
	return resp, err
}

func (s *ProductServiceOp[T]) GetSizeChartDetail(sid uint64, opt GetSizeChartDetailRequest, tok string) (*GetSizeChartDetailResponse, error) {
	path := "/product/get_size_chart_detail"
	resp := new(GetSizeChartDetailResponse)
	err := s.client.WithShop(sid, tok).Get(path, resp, opt)
	return resp, err
}

func (s *ProductServiceOp[T]) GetItemViolationInfo(sid uint64, req GetItemViolationInfoRequest, tok string) (*GetItemViolationInfoResponse, error) {
	path := "/product/get_item_violation_info"
	resp := new(GetItemViolationInfoResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

func (s *ProductServiceOp[T]) GetVariations(sid uint64, opt GetVariationsRequest, tok string) (*GetVariationsResponse, error) {
	path := "/product/get_variation_tree"
	resp := new(GetVariationsResponse)
	err := s.client.WithShop(sid, tok).Get(path, resp, opt)
	return resp, err
}

func (s *ProductServiceOp[T]) GetAllVehicleList(sid uint64, opt GetAllVehicleListRequest, tok string) (*GetAllVehicleListResponse, error) {
	path := "/product/get_all_vehicle_list"
	resp := new(GetAllVehicleListResponse)
	err := s.client.WithShop(sid, tok).Get(path, resp, opt)
	return resp, err
}

func (s *ProductServiceOp[T]) GetVehicleListByCompatibilityDetail(sid uint64, opt GetVehicleListByCompatibilityDetailRequest, tok string) (*GetVehicleListByCompatibilityDetailResponse, error) {
	path := "/product/get_vehicle_list_by_compatibility_detail"
	resp := new(GetVehicleListByCompatibilityDetailResponse)
	err := s.client.WithShop(sid, tok).Get(path, resp, opt)
	return resp, err
}

func (s *ProductServiceOp[T]) GetItemContentDiagnosisResult(sid uint64, req GetItemContentDiagnosisResultRequest, tok string) (*GetItemContentDiagnosisResultResponse, error) {
	path := "/product/get_item_content_diagnosis_result"
	resp := new(GetItemContentDiagnosisResultResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

func (s *ProductServiceOp[T]) GetItemListByContentDiagnosis(sid uint64, req GetItemListByContentDiagnosisRequest, tok string) (*GetItemListByContentDiagnosisResponse, error) {
	path := "/product/get_item_list_by_content_diagnosis"
	resp := new(GetItemListByContentDiagnosisResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

func (s *ProductServiceOp[T]) GetKitItemLimit(sid uint64, opt GetKitItemLimitRequest, tok string) (*GetKitItemLimitResponse, error) {
	path := "/product/get_kit_item_limit"
	resp := new(GetKitItemLimitResponse)
	err := s.client.WithShop(sid, tok).Get(path, resp, opt)
	return resp, err
}

func (s *ProductServiceOp[T]) AddKitItem(sid uint64, req AddKitItemRequest, tok string) (*AddKitItemResponse, error) {
	path := "/product/add_kit_item"
	resp := new(AddKitItemResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

func (s *ProductServiceOp[T]) UpdateKitItem(sid uint64, req UpdateKitItemRequest, tok string) (*UpdateKitItemResponse, error) {
	path := "/product/update_kit_item"
	resp := new(UpdateKitItemResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

func (s *ProductServiceOp[T]) GetKitItemInfo(sid uint64, opt GetKitItemInfoRequest, tok string) (*GetKitItemInfoResponse, error) {
	path := "/product/get_kit_item_info"
	resp := new(GetKitItemInfoResponse)
	err := s.client.WithShop(sid, tok).Get(path, resp, opt)
	return resp, err
}

func (s *ProductServiceOp[T]) GetSspList(sid uint64, opt GetSspListRequest, tok string) (*GetSspListResponse, error) {
	path := "/product/get_ssp_list"
	resp := new(GetSspListResponse)
	err := s.client.WithShop(sid, tok).Get(path, resp, opt)
	return resp, err
}

func (s *ProductServiceOp[T]) GetSspInfo(sid uint64, req GetSspInfoRequest, tok string) (*GetSspInfoResponse, error) {
	path := "/product/get_ssp_info"
	resp := new(GetSspInfoResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

func (s *ProductServiceOp[T]) AddSspItem(sid uint64, req AddSspItemRequest, tok string) (*AddSspItemResponse, error) {
	path := "/product/add_ssp_item"
	resp := new(AddSspItemResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

func (s *ProductServiceOp[T]) LinkSsp(sid uint64, req LinkSspRequest, tok string) (*LinkSspResponse, error) {
	path := "/product/link_ssp"
	resp := new(LinkSspResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

func (s *ProductServiceOp[T]) UnlinkSsp(sid uint64, req UnlinkSspRequest, tok string) (*UnlinkSspResponse, error) {
	path := "/product/unlink_ssp"
	resp := new(UnlinkSspResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

func (s *ProductServiceOp[T]) GetAitemByPitemId(sid uint64, opt GetAitemByPitemIdRequest, tok string) (*GetAitemByPitemIdResponse, error) {
	path := "/product/get_aitem_by_pitem_id"
	resp := new(GetAitemByPitemIdResponse)
	err := s.client.WithShop(sid, tok).Get(path, resp, opt)
	return resp, err
}

func (s *ProductServiceOp[T]) SearchAttributeValueList(sid uint64, req SearchAttributeValueListRequest, tok string) (*SearchAttributeValueListResponse, error) {
	path := "/product/search_attribute_value_list"
	resp := new(SearchAttributeValueListResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

func (s *ProductServiceOp[T]) GetMainItemList(sid uint64, opt GetMainItemListRequest, tok string) (*GetMainItemListResponse, error) {
	path := "/product/get_main_item_list"
	resp := new(GetMainItemListResponse)
	err := s.client.WithShop(sid, tok).Get(path, resp, opt)
	return resp, err
}

func (s *ProductServiceOp[T]) GetDirectItemList(sid uint64, opt GetDirectItemListRequest, tok string) (*GetDirectItemListResponse, error) {
	path := "/product/get_direct_item_list"
	resp := new(GetDirectItemListResponse)
	err := s.client.WithShop(sid, tok).Get(path, resp, opt)
	return resp, err
}

func (s *ProductServiceOp[T]) GetDirectShopRecommendedPrice(sid uint64, opt GetDirectShopRecommendedPriceRequest, tok string) (*GetDirectShopRecommendedPriceResponse, error) {
	path := "/product/get_direct_shop_recommended_price"
	resp := new(GetDirectShopRecommendedPriceResponse)
	err := s.client.WithShop(sid, tok).Get(path, resp, opt)
	return resp, err
}

func (s *ProductServiceOp[T]) GetProductCertificationRule(sid uint64, req GetProductCertificationRuleRequest, tok string) (*GetProductCertificationRuleResponse, error) {
	path := "/product/get_product_certification_rule"
	resp := new(GetProductCertificationRuleResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

func (s *ProductServiceOp[T]) PublishItemToOutletShop(sid uint64, tok string) (*PublishItemToOutletShopResponse, error) {
	path := "/"
	resp := new(PublishItemToOutletShopResponse)
	err := s.client.WithMerchant(sid, tok).Post(path, nil, resp)
	return resp, err
}

func (s *ProductServiceOp[T]) GetMartItemMappingById(sid uint64, req GetMartItemMappingByIdRequest, tok string) (*GetMartItemMappingByIdResponse, error) {
	path := "/product/get_mart_item_mapping_by_id"
	resp := new(GetMartItemMappingByIdResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

func (s *ProductServiceOp[T]) SearchUnpackagedModelList(sid uint64, req SearchUnpackagedModelListRequest, tok string) (*SearchUnpackagedModelListResponse, error) {
	path := "/product/search_unpackaged_model_list"
	resp := new(SearchUnpackagedModelListResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

func (s *ProductServiceOp[T]) GenerateKitImage(sid uint64, req GenerateKitImageRequest, tok string) (*GenerateKitImageResponse, error) {
	path := "/product/generate_kit_image"
	resp := new(GenerateKitImageResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}
