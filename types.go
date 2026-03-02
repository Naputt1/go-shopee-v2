// Code generated. DO NOT EDIT.
package goshopee

type AcceptOfferRequest struct {
	ReturnSn string `json:"return_sn"` // [Required] The serial number of return.
}

type AcceptOfferResponse struct {
	BaseResponse `json:",inline"`        // Common response fields
	Response     AcceptOfferResponseData `json:"response"` // Actual response data
}

type AcceptOfferResponseData struct {
	ReturnSn string `json:"return_sn"` // [Required] <p>The serial number of return.</p>
}

type Activity struct {
	ActivityId      int64           `json:"activity_id"`      // [Required] The id of activity.
	ActivityType    string          `json:"activity_type"`    // [Required] The type of activity.
	OriginalPrice   string          `json:"original_price"`   // [Required] activity's origin price
	DiscountedPrice string          `json:"discounted_price"` // [Required] activity's discount price
	Items           []ActivityItems `json:"items"`            // [Required]
	RefundAmount    float64         `json:"refund_amount"`    // [Required] <p>item's refund amount for bundle deal cases, only for shops whitelisted for Partial Qty RR.</p>
}

type ActivityItems struct {
	ItemId            int64  `json:"item_id"`            // [Required] The id of item.
	VariationId       int64  `json:"variation_id"`       // [Required] Shopee's unique identifier for a variation of an item.
	QuantityPurchased int64  `json:"quantity_purchased"` // [Required] item's quantity purchase
	OriginalPrice     string `json:"original_price"`     // [Required] item's origin price
}

type AddAddOnDealMainItemRequest struct {
	AddOnDealId  int64                          `json:"add_on_deal_id"` // [Required] Shopee's unique identifier for add on deal activity.
	MainItemList []AddBundleDealItemRequestItem `json:"main_item_list"` // [Required] The main items added in this add on deal promotion.
}

type AddAddOnDealMainItemResponse struct {
	BaseResponse `json:",inline"`                 // Common response fields
	Response     AddAddOnDealMainItemResponseData `json:"response"` // Actual response data
}

type AddAddOnDealMainItemResponseData struct {
	MainItemList []AddBundleDealItemRequestItem `json:"main_item_list"` // [Required] The main items added in this add on deal promotion.
	AddOnDealId  int64                          `json:"add_on_deal_id"` // [Required] Shopee's unique identifier for add on deal activity.
}

type AddAddOnDealRequest struct {
	AddOnDealName          string   `json:"add_on_deal_name"`                   // [Required] Title of the add on deal
	StartTime              int64    `json:"start_time"`                         // [Required] The time when add on deal activity start.
	EndTime                int64    `json:"end_time"`                           // [Required] The time when add on deal activity end
	PromotionType          int64    `json:"promotion_type"`                     // [Required] The type of add on deal：add on discount =0；gift with mini spend=1
	PurchaseMinSpend       *float64 `json:"purchase_min_spend,omitempty"`       // [Optional] The minimum purchase amount that needs to be met to buy the gift with min.Spend
	PerGiftNum             *int64   `json:"per_gift_num,omitempty"`             // [Optional] Number of gifts that buyers can get
	PromotionPurchaseLimit *int64   `json:"promotion_purchase_limit,omitempty"` // [Optional] promotion_purchase_limit
}

type AddAddOnDealResponse struct {
	BaseResponse `json:",inline"`         // Common response fields
	Response     AddAddOnDealResponseData `json:"response"` // Actual response data
}

type AddAddOnDealResponseData struct {
	AddOnDealId int64 `json:"add_on_deal_id"` // [Required] Shopee's unique identifier for an add on deal activity.
}

type AddAddOnDealSubItemRequest struct {
	AddOnDealId int64     `json:"add_on_deal_id"` // [Required] Shopee's unique identifier for add on deal activity.
	SubItemList []SubItem `json:"sub_item_list"`  // [Required] The sub items added in this add on deal promotion.
}

type AddAddOnDealSubItemResponse struct {
	BaseResponse `json:",inline"`                // Common response fields
	Response     AddAddOnDealSubItemResponseData `json:"response"` // Actual response data
}

type AddAddOnDealSubItemResponseData struct {
	SubItemList []ResponseDataSubItem `json:"sub_item_list"`  // [Required] The sub items added in this add on deal promotion.
	AddOnDealId int64                 `json:"add_on_deal_id"` // [Required] Shopee's unique identifier for add on deal activity.
}

type AddBundleDealItemRequest struct {
	BundleDealId int64                          `json:"bundle_deal_id"` // [Required] Shopee's unique identifier for a bundle deal activity.
	ItemList     []AddBundleDealItemRequestItem `json:"item_list"`      // [Required] The items added in this bundle deal promotion.
}

type AddBundleDealItemRequestItem struct {
	ItemId int64 `json:"item_id"` // [Required] Shopee's unique identifier for an item.
	Status int64 `json:"status"`  // [Required] The status of add on deal item：enable = 1；disable =2
}

type AddBundleDealItemResponse struct {
	BaseResponse `json:",inline"`              // Common response fields
	Response     AddBundleDealItemResponseData `json:"response"` // Actual response data
}

type AddBundleDealItemResponseData struct {
	FailedList  []ResponseDataFailed `json:"failed_list"`  // [Required] Indicate error details.
	SuccessList []interface{}        `json:"success_list"` // [Required] The list of succeed added items
}

type AddBundleDealRequest struct {
	RuleType           int64             `json:"rule_type"`                  // [Required] The bundle deal rule type：FIX_PRICE = 1 ；DISCOUNT_PERCENTAGE = 2； DISCOUNT_VALUE = 3
	DiscountValue      float64           `json:"discount_value"`             // [Required] The deducted price when when buying a bundle deal. Need to input it when the bundle deal rule type is 3
	FixPrice           float64           `json:"fix_price"`                  // [Required] The amount of the buyer needs to spend to purchase a bundle deal. Need to input it when the bundle deal rule type is 1
	DiscountPercentage int64             `json:"discount_percentage"`        // [Required] The discount that the buyer can get when buying a bundle deal. Need to input it when the bundle deal rule type is 2
	MinAmount          int64             `json:"min_amount"`                 // [Required] The quantity of items that need buyer to combine purchased
	StartTime          int64             `json:"start_time"`                 // [Required] The time when bundle deal activity start.The start time must be later than current time.
	EndTime            int64             `json:"end_time"`                   // [Required] The time when bundle deal activity end. The end time must be 1 hour later than start time.
	Name               string            `json:"name"`                       // [Required] Title of the bundle deal
	PurchaseLimit      int64             `json:"purchase_limit"`             // [Required] Maximum number of bundle deals that can be bought by a buyer.
	AdditionalTiers    []AdditionalTiers `json:"additional_tiers,omitempty"` // [Optional] <p>Use to create tiered discount for bundle deals, a max of 2 additional tiers are allowed to create.</p><p>the rule of multiple tiers needs to follow this faq&nbsp;<a href="https://open.shopee.com/faq/53" target="_blank" style="font-size:14px;">https://open.shopee.com/faq/53</a></p><p>For additional tiers, the fix price, discount_percentage, discount_value should be consistent with tier 1</p>
}

type AddBundleDealResponse struct {
	BaseResponse `json:",inline"`          // Common response fields
	Response     AddBundleDealResponseData `json:"response"` // Actual response data
}

type AddBundleDealResponseData struct {
	BundleDealId int64 `json:"bundle_deal_id"` // [Required] Shopee's unique identifier for a bundle deal activity.
}

type AddDiscountItemRequest struct {
	DiscountId int64                        `json:"discount_id"` // [Required] Shopee's unique identifier for a discount activity.
	ItemList   []AddDiscountItemRequestItem `json:"item_list"`   // [Required] The items added in this discount promotion.
}

type AddDiscountItemRequestItem struct {
	ItemId             int64              `json:"item_id"`                        // [Required] Shopee's unique identifier for an item.
	ItemPromotionPrice *float64           `json:"item_promotion_price,omitempty"` // [Optional] The discount price of the item. If the item has no variation, this param is necessary.
	ItemPromotionStock *int64             `json:"item_promotion_stock,omitempty"` // [Optional] The reserved stock of the item.
	ModelList          []RequestItemModel `json:"model_list,omitempty"`           // [Optional] The models which belongs to this item.
	PurchaseLimit      int64              `json:"purchase_limit"`                 // [Required] The max number of this product in the promotion price. If it's No Limit, please input the 0 for this request data.
}

type AddDiscountItemResponse struct {
	BaseResponse `json:",inline"`            // Common response fields
	Response     AddDiscountItemResponseData `json:"response"` // Actual response data
}

type AddDiscountItemResponseData struct {
	DiscountId int64   `json:"discount_id"` // [Required] Shopee's unique identifier for a discount activity.
	Count      int64   `json:"count"`       // [Required] The number of items that add successfully.
	ErrorList  []Error `json:"error_list"`  // [Required] Indicate error details.
}

type AddDiscountRequest struct {
	DiscountName string `json:"discount_name"` // [Required] Title of the discount.
	StartTime    int64  `json:"start_time"`    // [Required] The time when discount activity start.The start time must be 1 hour later than current time.
	EndTime      int64  `json:"end_time"`      // [Required] <p>The time when discount activity end.The end time must be 1 hour later than start time,and the discount period must be less than 180 days.</p>
}

type AddDiscountResponse struct {
	BaseResponse `json:",inline"`        // Common response fields
	Response     AddDiscountResponseData `json:"response"` // Actual response data
}

type AddDiscountResponseData struct {
	DiscountId int64 `json:"discount_id"` // [Required] Shopee's unique identifier for a discount activity.
}

type AddFollowPrizeRequest struct {
	FollowPrizeName string   `json:"follow_prize_name"`         // [Required] <p>The name of the follow prize,The follow prize name length max limit is 20.<br /></p>
	StartTime       int64    `json:"start_time"`                // [Required] <p>The timing from when the follow prize is valid,the start time later than the current time.If the start time and end time passed in by the seller overlap with other upcoming/ongoing activities, it will prompt "Another Follow Prize voucher already exists during this time period, please set another period."<br /></p>
	EndTime         int64    `json:"end_time"`                  // [Required] <p>The timing until when the follow prize is still valid,the end time must be greater than the start time by at least 1 day and end time cannot exceed 3 months after start time.If the start time and end time passed in by the seller overlap with other upcoming/ongoing activities, it will prompt "Another Follow Prize voucher already exists during this time period, please set another period."<br /></p>
	UsageQuantity   int64    `json:"usage_quantity"`            // [Required] <p>Please enter a value between 1 and 200000.</p>
	MinSpend        float64  `json:"min_spend"`                 // [Required] <p>The minimum spend required for using this follow prize.<br /></p>
	RewardType      int64    `json:"reward_type"`               // [Required] <p>The reward type of the follow prize.The available values are:1:discount---fix amount,2:discount---by percentage,3:coin cash back.<br /></p>
	DiscountAmount  *float64 `json:"discount_amount,omitempty"` // [Optional] <p>The discount amount set for this particular follow prize.Only fill in when you are creating a fix amount follow prize.<br /></p>
	Percentage      *int64   `json:"percentage,omitempty"`      // [Optional] <p>The discount percentage set for this particular follow prize. Only fill in when you are creating a discount percentage follow prize or coins cashback follow prize.Discount percentage (reward_type ==2) or Percentage of coins cash back (reward_type==3).<br /></p>
	MaxPrice        *float64 `json:"max_price,omitempty"`       // [Optional] <p>The max amount of discount/value a user can enjoy by using this particular follow prize. It is required to fill in when you are creating a discount percentage follow prize or coins cashback follow prize. max_price &gt;=1<br /></p>
}

type AddFollowPrizeResponse struct {
	BaseResponse `json:",inline"`           // Common response fields
	Response     AddFollowPrizeResponseData `json:"response"` // Actual response data
}

type AddFollowPrizeResponseData struct {
	CampaginId int64 `json:"campagin_id"` // [Required] <p>The unique identifier for the created follow prize.<br /></p>
}

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

type AddItemListRequest struct {
	ShopCategoryId int64   `json:"shop_category_id"` // [Required] ShopCategory's unique identifier.
	ItemList       []int64 `json:"item_list"`        // [Required] Shopee's unique identifiers list for an item. Max. 100 items to be deleted per request.
}

type AddItemListResponse struct {
	BaseResponse `json:",inline"`        // Common response fields
	Response     AddItemListResponseData `json:"response"` // Actual response data
}

type AddItemListResponseData struct {
	InvalidItemIdList []ResponseDataFailed `json:"invalid_item_id_list"` // [Required] List of invalid item ids.
	ShopCategoryId    int64                `json:"shop_category_id"`     // [Required] ShopCategory's unique identifier.
	CurrentCount      int64                `json:"current_count"`        // [Required] Count of items under this shop category after deletion.
}

type AddItemRequest struct {
	OriginalPrice        float64            `json:"original_price"`                   // [Required] Item price
	Description          string             `json:"description"`                      // [Required] if description_type is normal , Description information should be set by this field.
	Weight               float64            `json:"weight"`                           // [Required] <p>The weight of this item, the unit is KG.</p>
	ItemName             string             `json:"item_name"`                        // [Required] Item name
	ItemStatus           *ItemStatus        `json:"item_status,omitempty"`            // [Optional] Item status, could be UNLIST or NORMAL
	Dimension            *Dimension         `json:"dimension,omitempty"`              // [Optional] <p>The dimension of this item.</p>
	LogisticInfo         []LogisticInfo     `json:"logistic_info"`                    // [Required] Logistic channel setting
	AttributeList        []Attribute        `json:"attribute_list,omitempty"`         // [Optional] This field is optional(expect Indonesia) depending on the specific attribute under different categories. Should call shopee.item.GetAttributes to get attribute first. Must contain all all mandatory attribute.
	CategoryId           int64              `json:"category_id"`                      // [Required] ID of category
	Image                *Image             `json:"image"`                            // [Required] Item images
	PreOrder             *PreOrder          `json:"pre_order,omitempty"`              // [Optional] Pre order setting
	ItemSku              *string            `json:"item_sku,omitempty"`               // [Optional] SKU tag of item
	Condition            *string            `json:"condition,omitempty"`              // [Optional] Condition of item, could be USED or NEW
	Wholesale            []Wholesale        `json:"wholesale,omitempty"`              // [Optional] Wholesale setting
	VideoUploadId        []string           `json:"video_upload_id,omitempty"`        // [Optional] Video upload ID returned from video uploading API. Only accept one video_upload_id.
	Brand                *Brand             `json:"brand,omitempty"`                  // [Optional]
	ItemDangerous        *int64             `json:"item_dangerous,omitempty"`         // [Optional] This field is only applicable for local sellers in Indonesia and Malaysia. Use this field to identify whether a product is a dangerous product. 0 for non-dangerous product and 1 for dangerous product. For more information, please visit the market's respective Seller Education Hub.
	TaxInfo              *TaxInfo           `json:"tax_info,omitempty"`               // [Optional] Tax information
	ComplaintPolicy      *ComplaintPolicy   `json:"complaint_policy,omitempty"`       // [Optional] Complaint Policy for item. Only required for local PL sellers, ignored otherwise.
	DescriptionInfo      *DescriptionInfo   `json:"description_info,omitempty"`       // [Optional] New description field. Only whitelist sellers can use it. If you use the field, please upload the description_type=extended otherwise api will return error. If you don't use this field, you don't need to upload the description_type or upload description_type=normal
	DescriptionType      *DescriptionType   `json:"description_type,omitempty"`       // [Optional] Values: See Data Definition- description_type (normal , extended). If you want to use extended_description, this field must be inputed
	SellerStock          []SellerStock      `json:"seller_stock,omitempty"`           // [Optional] <p>seller stock（Please notice that stock(including Seller Stock and Shopee Stock) should be larger than or equal to real-time reserved stock）</p>
	GtinCode             *string            `json:"gtin_code,omitempty"`              // [Optional] <p>- GTIN is an identifier for trade items, developed by the international organization GS1.<br />- They have 8 to 14 digits. The most common are UPC, EAN, JAN and ISBN.<br />- GTIN will help boost positioning in online marketing channels like Google and Facebook.<br />- That incorporation with GTIN will also aid in Search and Recommendation in Shopee itself allowing buyers to have higher likelihood of finding one's listing.<br /></p><p><br /></p><p>Note: If you want to set “Item without GTIN”, please pass the gtin_code as "00".<br /><br />The validation rule is based on the value return in gtin_validation_rule" field in v2.product.get_item_limit API</p><p><b>- Mandatory</b>:&nbsp;This field is required and must contain a correctly formatted GTiN number.</p><p><b>- Flexible</b>: This field is required and must contain either a correctly formatted GTlN number or "00" to declare that the item/model has no valid GTlN.<br /><b>- Optional:</b> This field is optional and can contain a correctly formatted GTiN number, "00" or be omitted entirely.</p>
	DsCatRcmdId          *string            `json:"ds_cat_rcmd_id,omitempty"`         // [Optional] <p>category recommendation service id</p>
	PromotionImages      *PromotionImages   `json:"promotion_images,omitempty"`       // [Optional] <p>Promotion Image<br />Currently only allow one promoton image<br />You could set promotion image only if the product images' ratio is 3:4<br /></p>
	CompatibilityInfo    *CompatibilityInfo `json:"compatibility_info,omitempty"`     // [Optional]
	ScheduledPublishTime *int64             `json:"scheduled_publish_time,omitempty"` // [Optional] <p>Scheduled publish time of this item:&nbsp;</p><p>1) Can only set scheduled_publish_time for item with UNLIST status</p><p>2) Can only set the time from current time +1hour to current time +90days, and the time is only allowed to be accurate to the minute</p>
	AuthorisedBrandId    *int64             `json:"authorised_brand_id,omitempty"`    // [Optional] <p>ID of authorised reseller brand.</p>
	SizeChartInfo        *SizeChartInfo     `json:"size_chart_info,omitempty"`        // [Optional]
	CertificationInfo    *CertificationInfo `json:"certification_info,omitempty"`     // [Optional] <p>For PH product certification input<br />Required for some category and attribute option</p>
	PurchaseLimitInfo    *PurchaseLimitInfo `json:"purchase_limit_info,omitempty"`    // [Optional] <p>purchase limit info</p>
}

type AddItemResponse struct {
	BaseResponse `json:",inline"`    // Common response fields
	Response     AddItemResponseData `json:"response"` // Actual response data
}

type AddItemResponseData struct {
	Description     string                 `json:"description"`      // [Required] Description of item
	Weight          float64                `json:"weight"`           // [Required] <p>The weight of this item, the unit is KG.<br /></p>
	PreOrder        *PreOrder              `json:"pre_order"`        // [Required] Pre order setting
	ItemName        string                 `json:"item_name"`        // [Required] Item name
	Images          *Images                `json:"images"`           // [Required] Item images
	ItemStatus      ItemStatus             `json:"item_status"`      // [Required] Item status
	PriceInfo       *ResponseDataPriceInfo `json:"price_info"`       // [Required] Item price info
	LogisticInfo    []LogisticInfo         `json:"logistic_info"`    // [Required] Logistic setting
	ItemId          int64                  `json:"item_id"`          // [Required] Item ID
	Attribute       []Attribute            `json:"attribute"`        // [Required] Item attributes
	CategoryId      int64                  `json:"category_id"`      // [Required] Category ID
	Dimension       *Dimension             `json:"dimension"`        // [Required] <p>The dimension of this item.<br /></p>
	Condition       string                 `json:"condition"`        // [Required] Item condition, could be NEW or USED
	VideoInfo       []VideoInfo            `json:"video_info"`       // [Required] Item video
	Wholesale       []Wholesale            `json:"wholesale"`        // [Required] Wholesale setting
	Brand           *Brand                 `json:"brand"`            // [Required]
	ItemDangerous   int64                  `json:"item_dangerous"`   // [Required] This field is only applicable for local sellers in Indonesia and Malaysia. Use this field to identify whether a product is a dangerous product. 0 for non-dangerous product and 1 for dangerous product. For more information, please visit the market's respective Seller Education Hub.
	DescriptionInfo *DescriptionInfo       `json:"description_info"` // [Required] New description field. Only whitelist sellers can use it. If item with extended_description this field will return, otherwise do not return.
	DescriptionType DescriptionType        `json:"description_type"` // [Required] Values: See Data Definition- description_type (normal , extended).
	ComplaintPolicy *ComplaintPolicy       `json:"complaint_policy"` // [Required] Complaint Policy for item. Only returned for local PL sellers.
	SellerStock     []SellerStock          `json:"seller_stock"`     // [Required] seller stock
}

type AddKitItemRequest struct {
	ItemSetting *ItemSetting `json:"item_setting"`           // [Required]
	SyncSetting *SyncSetting `json:"sync_setting,omitempty"` // [Optional]
}

type AddKitItemResponse struct {
	BaseResponse `json:",inline"`       // Common response fields
	Response     AddKitItemResponseData `json:"response"` // Actual response data
}

type AddKitItemResponseData struct {
	ItemId int64 `json:"item_id"` // [Required]
}

type AddModelRequest struct {
	ItemId    int64   `json:"item_id"`    // [Required] ID of item
	ModelList []Model `json:"model_list"` // [Required] Model list
}

type AddModelResponse struct {
	BaseResponse `json:",inline"`     // Common response fields
	Response     AddModelResponseData `json:"response"` // Actual response data
}

type AddModelResponseData struct {
	Model []AddModelResponseDataModel `json:"model"` // [Required]
}

type AddModelResponseDataModel struct {
	TierIndex   interface{}      `json:"tier_index"`   // [Required] model tier index
	ModelId     int64            `json:"model_id"`     // [Required] ID of model
	ModelSku    string           `json:"model_sku"`    // [Required] Seller SKU of this model, model_sku length information needs to be no more than 100 characters.
	PriceInfo   []ModelPriceInfo `json:"price_info"`   // [Required]
	SellerStock []SellerStock    `json:"seller_stock"` // [Required] <p>new stock info</p>
	Weight      float64          `json:"weight"`       // [Required] <p>The weight of this model, the unit is KG.</p><p>If don't set the weight of this model, will use the weight of item by default.</p><p>If set the dimension of this model, them must set the weight of this model.</p>
	Dimension   *Dimension       `json:"dimension"`    // [Required] <p>The dimension of this model.</p><p>If don't set the dimension of this model, will use the dimension of item by default.</p>
}

type AddOnDeal struct {
	StartTime              int64   `json:"start_time"`               // [Required] The time when add on deal activity start.
	EndTime                int64   `json:"end_time"`                 // [Required] The time when add on deal activity end
	PromotionType          int64   `json:"promotion_type"`           // [Required] The type of add on deal：add on discount =0；gift with mini spend=1
	PurchaseMinSpend       float64 `json:"purchase_min_spend"`       // [Required] The minimum purchase amount that needs to be met to buy the gift with min.Spend
	AddOnDealId            int64   `json:"add_on_deal_id"`           // [Required] Shopee's unique identifier for an add on deal activity.
	PerGiftNum             int64   `json:"per_gift_num"`             // [Required] Number of gifts that buyers can get
	PromotionPurchaseLimit int64   `json:"promotion_purchase_limit"` // [Required] Max. number of add-on products that a customer can purchase per order.
	AddOnDealName          string  `json:"add_on_deal_name"`         // [Required] Title of the add on deal
	Source                 int64   `json:"source"`                   // [Required] The create source of bundle deal：Seller=1，shopee admin=0
	SubItemPrioriry        []int64 `json:"sub_item_prioriry"`        // [Required] The display sequence of sub item in buyer side
}

type AddShopCategoryRequest struct {
	Name       string `json:"name"`                  // [Required] ShopCategory's name.
	SortWeight *int64 `json:"sort_weight,omitempty"` // [Optional] ShopCategory's sort weight. The maximum number should be 2147483546.
}

type AddShopCategoryResponse struct {
	BaseResponse `json:",inline"`            // Common response fields
	Response     AddShopCategoryResponseData `json:"response"` // Actual response data
}

type AddShopCategoryResponseData struct {
	ShopCategoryId int64 `json:"shop_category_id"` // [Required] ShopCategory's unique identifier.
}

type AddShopFlashSaleItemsRequest struct {
	FlashSaleId int64          `json:"flash_sale_id"` // [Required]
	Items       []RequestItems `json:"items"`         // [Required]
}

type AddShopFlashSaleItemsResponse struct {
	BaseResponse `json:",inline"`                  // Common response fields
	Response     AddShopFlashSaleItemsResponseData `json:"response"` // Actual response data
}

type AddShopFlashSaleItemsResponseData struct {
	FailedItems []FailedItems `json:"failed_items"` // [Required]
}

type AddSspItemRequest struct {
	SspId             int64                     `json:"ssp_id"`                        // [Required] <p>Shopee's unique identifier for Shopee&nbsp;Standard Product.<br /></p>
	OriginalPrice     *float64                  `json:"original_price,omitempty"`      // [Optional] <p>The price of this item.<br /></p>
	ItemStatus        *ItemStatus               `json:"item_status,omitempty"`         // [Optional] <p>Item status, could be UNLIST or NORMAL.<br /></p>
	Dimension         *Dimension                `json:"dimension,omitempty"`           // [Optional] <p>The dimension of this item.<br /></p>
	LogisticInfo      []LogisticInfo            `json:"logistic_info"`                 // [Required] <p>Logistic channel setting of this item.<br /></p>
	AttributeList     []Attribute               `json:"attribute_list,omitempty"`      // [Optional]
	PreOrder          *PreOrder                 `json:"pre_order,omitempty"`           // [Optional] <p>Pre order setting of this item.</p>
	ItemSku           *string                   `json:"item_sku,omitempty"`            // [Optional] <p>SKU tag of this item.<br /></p>
	Condition         *string                   `json:"condition,omitempty"`           // [Optional] <p>Condition of item, could be USED or NEW.<br /></p>
	Wholesale         []Wholesale               `json:"wholesale,omitempty"`           // [Optional] <p>Wholesale setting of this item.<br /></p>
	VideoUploadId     []string                  `json:"video_upload_id,omitempty"`     // [Optional] <p>Video upload ID returned from video uploading API. Only accept one video_upload_id.<br /></p>
	ItemDangerous     *int64                    `json:"item_dangerous,omitempty"`      // [Optional] <p>This field is only applicable for&nbsp;local sellers&nbsp;in Indonesia and Malaysia. Use this field to identify whether a product is a dangerous product. 0 for non-dangerous product and 1 for dangerous product. For more information, please visit the market's respective Seller Education Hub.<br /></p>
	TaxInfo           *AddSspItemRequestTaxInfo `json:"tax_info,omitempty"`            // [Optional] <p>Tax information for this item.<br /></p>
	SellerStock       []SellerStock             `json:"seller_stock,omitempty"`        // [Optional] <p>seller stock（Please notice that stock(including Seller Stock and Shopee Stock) should be larger than or equal to real-time reserved stock）<br /></p>
	SizeChartInfo     *SizeChartInfo            `json:"size_chart_info,omitempty"`     // [Optional]
	AuthorisedBrandId *int64                    `json:"authorised_brand_id,omitempty"` // [Optional] <p>ID of authorised reseller brand.<br /></p>
	ModelList         []AddSspItemRequestModel  `json:"model_list,omitempty"`          // [Optional] <p>Model info list.<br /></p>
}

type AddSspItemRequestModel struct {
	TierIndex     interface{}   `json:"tier_index"`             // [Required] <p>Tier index of this model.</p>
	ModelSku      *string       `json:"model_sku,omitempty"`    // [Optional] <p>Seller SKU of this model, model_sku length information needs to be no more than 100 characters.<br /></p>
	OriginalPrice float64       `json:"original_price"`         // [Required] <p>Original price of this model.</p>
	SellerStock   []SellerStock `json:"seller_stock,omitempty"` // [Optional] <p>seller stock（Please notice that stock(including Seller Stock and Shopee Stock) should be larger than or equal to real-time reserved stock）<br /></p>
	Dimension     *Dimension    `json:"dimension,omitempty"`    // [Optional] <p>The dimension of this model.</p><p>If don't set the dimension of this model, will use the dimension of item by default.</p>
	PreOrder      *PreOrder     `json:"pre_order,omitempty"`    // [Optional] <p>Pre-order information of this model.</p><p><br /></p><p>Notes:&nbsp;</p><p>If don't set the DTS of this model, will use the DTS of the item by default.</p>
}

type AddSspItemRequestTaxInfo struct {
	Ncm               *string        `json:"ncm,omitempty"`                 // [Optional] <p>Mercosur Common Nomenclature, it is a convention between Mercosur member countries to easily recognize goods, services and productive factors negotiated among themselves.&nbsp;(BR region)<br /></p><p><br />NCM must have 8 digits, OR, if your item doesn't have a NCM enter the value "00"</p>
	SameStateCfop     *string        `json:"same_state_cfop,omitempty"`     // [Optional] <p>Tax Code of Operations and Installments for orders that seller and buyer are in the same state. It identifies a specific operation by category at the time of issuing the invoice.(BR region)<br /></p>
	DiffStateCfop     *string        `json:"diff_state_cfop,omitempty"`     // [Optional] <p>Tax Code of Operations and Installments for orders that seller and buyer are in different states. It identifies a specific operation by category at the time of issuing the invoice.(BR region)<br /></p>
	Csosn             *string        `json:"csosn,omitempty"`               // [Optional] <p>Code of Operation Status – Simples Nacional, code for company operations to identify the origin of the goods and the taxation regime of the operations.(BR region)<br /></p>
	Origin            *string        `json:"origin,omitempty"`              // [Optional] <p>Product source, domestic or foreig (BR region).<br /></p>
	Cest              *string        `json:"cest,omitempty"`                // [Optional] <p>Tax Replacement Specifying Code (CEST), to separate within the same NCM products that do or do not have ICMS tax substitution. (BR region)<br /><br />CEST must have 7 digits, OR, if your item doesn't have a CEST enter the value "00".<br /></p>
	MeasureUnit       *string        `json:"measure_unit,omitempty"`        // [Optional] <p>(BR region)<br /></p>
	InvoiceOption     *InvoiceOption `json:"invoice_option,omitempty"`      // [Optional] <p>Value shuold be one of NO_INVOICES VAT_MARGIN_SCHEME_INVOICES VAT_INVOICES NON_VAT_INVOICES and if value is NON_VAT_INVOICE vat_rate should be null (PL region)<br /></p>
	VatRate           *string        `json:"vat_rate,omitempty"`            // [Optional] <p>Value should be one of 0% 5% 8% 23% NO_VAT_RATE (PL region)<br /></p>
	HsCode            string         `json:"hs_code"`                       // [Required] <p>HS Code (Only for IN region)<br /></p>
	TaxCode           string         `json:"tax_code"`                      // [Required] <p>Tax Code (Only for IN region)<br /></p>
	TaxType           *TaxType       `json:"tax_type,omitempty"`            // [Optional] <p>tax_type only for TW whitelist shop. Shopee will referred Tax type when substitute sellers for issuing e-receipts to buyers. All variations share the same tax type. The meaning of value:&nbsp;</p><p>0: no tax type</p><p>1: tax-able</p><p>2: tax-free</p>
	Pis               *string        `json:"pis,omitempty"`                 // [Optional] <p>Only for BR shop.</p><p>PIS - Programa de Integração Social (Social Integration Program). It is a government tax to collect resources for the payment of unemployment insurance and other employee related rights.</p><p>PIS % - the tax applied to this product</p>
	Cofins            *string        `json:"cofins,omitempty"`              // [Optional] <p>Only for BR shop.<br /></p><p>COFINS – Contribuição para Financiamento da Seguridade Social (Contribution for Social Security Funding). It&nbsp;is a government tax to collect resources for public health system and social security.</p><p>COFINS&nbsp;% - the tax applied to this product</p>
	IcmsCst           *string        `json:"icms_cst,omitempty"`            // [Optional] <p>Only for BR shop.<br /></p><p>ICMS - Imposto sobre Circulação de Mercadorias e Serviços (Circulation of Goods and Services Tax).&nbsp;</p><p>CST - Código da Situação Tributária (Tax Situation Code) is represented by a combination of 3 numbers with the purpose of demonstrating the origin of a product and determining the form of taxation that will apply to it. Therefore, each digit in the CST Table has a specific meaning: the first digit indicates the origin of the operation, the second digit represents the ICMS taxation on the operation and the third digit provides additional information about the form of taxation.</p>
	PisCofinsCst      *string        `json:"pis_cofins_cst,omitempty"`      // [Optional] <p>Only for BR shop.</p><p>The CST PIS/Cofins is a code on the Electronic Invoice (NF-e) that identifies the tax situation of PIS (Programa de Integração Social) and Cofins (Contribuição para o Financiamento da Seguridade Social) in sales of goods.</p>
	FederalStateTaxes *string        `json:"federal_state_taxes,omitempty"` // [Optional] <p>Only for BR shop.</p><p>Enter the total percentage of the combination of federal, state, and municipal taxes, using up to two decimals.</p>
	OperationType     *OperationType `json:"operation_type,omitempty"`      // [Optional] <p>Only for BR shop.</p><p>1: Retailer</p><p>2: Manufacturer</p>
	ExTipi            *string        `json:"ex_tipi,omitempty"`             // [Optional] <p>Only for BR shop.<br /></p><p>The EXTIPI field in the NF-e (Nota Fiscal Eletrônica) is used to indicate if there's an exception to the IPI (Imposto sobre Produtos Industrializados) tax rate for a specific product.</p>
	FciNum            *string        `json:"fci_num,omitempty"`             // [Optional] <p>Only for BR shop.<br /></p><p>The FCI Control Number is a unique identifier assigned to each import FCI (Import Content Form). It's mandatory on the corresponding NF-e (electronic invoice) to ensure compliance with Brazilian import tax regulations.</p>
	RecopiNum         *string        `json:"recopi_num,omitempty"`          // [Optional] <p>Only for BR shop.</p><p>RECOPI NACIONAL is a Brazilian government system that facilitates the registration and management of tax-exempt operations involving paper destined for printing books, newspapers, and periodicals (known as "papel imune" in Portuguese).</p>
	AdditionalInfo    *string        `json:"additional_info,omitempty"`     // [Optional] <p>Only for BR shop.</p><p>Include relevant information to display on Invoice.</p>
	GroupItemInfo     *GroupItemInfo `json:"group_item_info,omitempty"`     // [Optional] <p>Only for BR shop.</p><p>Required if the item is a group item.</p>
}

type AddSspItemResponse struct {
	BaseResponse `json:",inline"`       // Common response fields
	Response     AddSspItemResponseData `json:"response"` // Actual response data
}

type AddSspItemResponseData struct {
	ItemId int64 `json:"item_id"` // [Required] <p>Shopee's unique identifier for an item.<br /></p>
}

type AddTopPicksRequest struct {
	Name        string  `json:"name"`         // [Required]
	ItemIdList  []int64 `json:"item_id_list"` // [Required]
	IsActivated bool    `json:"is_activated"` // [Required]
}

type AddTopPicksResponse struct {
	BaseResponse `json:",inline"`        // Common response fields
	Response     AddTopPicksResponseData `json:"response"` // Actual response data
}

type AddTopPicksResponseData struct {
	CollectionList []Collection `json:"collection_list"` // [Required] The top picks list in this shop.
}

type AddVoucherRequest struct {
	VoucherName        string       `json:"voucher_name"`                   // [Required] The name of the voucher.
	VoucherCode        string       `json:"voucher_code"`                   // [Required] The code of the voucher.
	StartTime          int64        `json:"start_time"`                     // [Required] <p>The timing from when the voucher is valid; so buyer is allowed to claim and to use. Field can only be updated if voucher has not started.</p>
	EndTime            int64        `json:"end_time"`                       // [Required] The timing until when the voucher is still valid. Any time after this end_time, buyer is not allowed to claim or to use.
	VoucherType        int64        `json:"voucher_type"`                   // [Required] The type of the voucher. The available values are: 1: shop voucher, 2: product voucher.
	RewardType         int64        `json:"reward_type"`                    // [Required] The reward type of the voucher. The available values are: 1: fix_amount voucher, 2: discount_percentage voucher, 3: coin_cashback voucher.
	UsageQuantity      int64        `json:"usage_quantity"`                 // [Required] The number of times for this particular voucher could be used.
	MinBasketPrice     float64      `json:"min_basket_price"`               // [Required] The minimum spend required for using this voucher.
	DiscountAmount     *float64     `json:"discount_amount,omitempty"`      // [Optional] The discount amount set for this particular voucher. Only fill in when you are creating a fix amount voucher.
	Percentage         *int64       `json:"percentage,omitempty"`           // [Optional] The discount percentage set for this particular voucher. Only fill in when you are creating a discount percentage voucher or coins cashback voucher.
	MaxPrice           *float64     `json:"max_price,omitempty"`            // [Optional] <p>The max amount of discount/value a user can enjoy by using this particular voucher. Only fill in when you are creating a discount percentage voucher or coins cashback voucher.If no cap limit, can set to be 0.</p>
	DisplayChannelList *interface{} `json:"display_channel_list,omitempty"` // [Optional] <p>The FE channel where the voucher will be displayed. The available values are: 1: display_all&nbsp; 3: feed, 4: live streaming,   [] (empty - which is hidden).</p>
	ItemIdList         []int64      `json:"item_id_list,omitempty"`         // [Optional] The list of items which is applicable for the voucher. Only fill in when you are creating a product type of voucher.
	DisplayStartTime   *int64       `json:"display_start_time,omitempty"`   // [Optional] <p>The timing of when voucher is displayed on shop pages; so buyer is allowed to claim.display_start_time must be left empty if the&nbsp;display_channel_list is empty (when voucher is hidden), otherwise it will show error.</p>
}

type AddVoucherResponse struct {
	BaseResponse `json:",inline"`       // Common response fields
	Response     AddVoucherResponseData `json:"response"` // Actual response data
}

type AddVoucherResponseData struct {
	VoucherId int64 `json:"voucher_id"` // [Required] The unique identifier for the created voucher.
}

type AdditionalTiers struct {
	MinAmount          int64   `json:"min_amount"`          // [Required] <p>The quantity of items that the buyers need to purchase for additional tier<br /></p>
	FixPrice           float64 `json:"fix_price"`           // [Required] <p>The bundle price when the buyers purchase a bundle deal for additional tiers. Need to input it when the bundle deal rule type is 1.<br /></p>
	DiscountValue      float64 `json:"discount_value"`      // [Required] <p>The bundle deal discount amount the buyer can save when purchasing a bundle deal. Need to input it when the bundle deal rule type is 3<br /></p>
	DiscountPercentage int64   `json:"discount_percentage"` // [Required] <p>The bundle deal discount% that the buyer can get when buying a bundle deal for additional tiers. Need to input it when the bundle deal rule type is 2<br /></p>
}

type Address struct {
	AddressName string `json:"address_name"` // [Required] <p>The address name filled in when creating the warehouse.<br /></p>
	Region      string `json:"region"`       // [Required] <p>Region of your warehouse address.<br /></p>
	Address     string `json:"address"`      // [Required] <p>Detail address of your warehouse address.<br /></p>
	City        string `json:"city"`         // [Required] <p>City of your warehouse address.<br /></p>
	District    string `json:"district"`     // [Required] <p>Distinct of your warehouse address.<br /></p>
	State       string `json:"state"`        // [Required] <p>State of your warehouse address.<br /></p>
	Town        string `json:"town"`         // [Required] <p>Town of your warehouse address.<br /></p>
	ZipCode     string `json:"zip_code"`     // [Required] <p>Zipcode of your warehouse address.<br /></p>
}

type AddressBreakdown struct {
	Region          string `json:"region"`           // [Required] <p>Return region value</p><p>- PH, TH only<br /></p>
	State           string `json:"state"`            // [Required] <p>Return value</p><p>- TH: Province<br /></p>
	City            string `json:"city"`             // [Required] <p>Return value</p><p>- TH: District<br /></p>
	Town            string `json:"town"`             // [Required] <p>Return value</p><p>- TH: Sub district<br /></p>
	Postcode        string `json:"postcode"`         // [Required] <p>Return value</p><p>- TH: Postal code</p><p>- PH: Postal code<br /></p>
	DetailedAddress string `json:"detailed_address"` // [Required] <p>Return value</p><p>- PH: Additional details, i.e. street name, building</p><p>- TH: Additional details, i.e. house number<br /></p>
	AdditionalInfo  string `json:"additional_info"`  // [Required] <p>Return value:</p><p>- Empty for PH, TH<br /></p>
	FullAddress     string `json:"full_address"`     // [Required] <p>- only has value when invoice_type is personal</p><p>- Buyer address in format "detailed_address, town, district, state, postcode, additional_info" for all regions</p><p>--- for TH: leave the 'additional_info' as empty</p>
}

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

type AdjustQty struct {
	AdjustTotal     int64 `json:"adjust_total"`      // [Required] <p>"Total number of SKU changes during the selected time period."</p>
	AdjustLostFound int64 `json:"adjust_lost_found"` // [Required] <p>"Total quantity of lost or recovered items during the selected time period."</p>
	AdjustTransWhs  int64 `json:"adjust_trans_whs"`  // [Required] <p>Total quantity of transfer orders created by the warehouse during the selected time period<br /></p>
}

type AdvanceStock struct {
	SellableAdvanceStock  int64 `json:"sellable_advance_stock"`   // [Required] <p>Refers to Advance Fulfillment&nbsp;stock that Seller has shipped out and is available&nbsp;to be used to fulfill an order.<br /></p>
	InTransitAdvanceStock int64 `json:"in_transit_advance_stock"` // [Required] <p>Refers to Advance Fulfillment stock that seller has shipped out and is still in transit and unavailable to be used to fulfill an order.<br /></p>
}

type AffiliateInfo struct {
	CommissionRate    float64 `json:"commission_rate"`     // [Required] <p>The commission rate that the streamer can get, for example, 0.1 means 10%.</p>
	IsCampaign        bool    `json:"is_campaign"`         // [Required] <p>Whether participate in a campaign project (generally, the commission will be higher)</p>
	CampaignMcnName   string  `json:"campaign_mcn_name"`   // [Required] <p>MCN agency that initiated this campaign</p>
	CampaignStartTime int64   `json:"campaign_start_time"` // [Required] <p>Campaign start time, it's unix timestamp in seconds.</p>
	CampaignEndTime   int64   `json:"campaign_end_time"`   // [Required] <p>Campaign end time, it's unix timestamp in seconds.</p>
}

type Aitem struct {
	AshopId          int64          `json:"ashop_id"`           // [Required] <p>ID of SIP Affiliate Shop.</p>
	AshopRegion      string         `json:"ashop_region"`       // [Required] <p>Region of SIP Affiliate Shop.</p>
	AitemId          int64          `json:"aitem_id"`           // [Required] <p>ID of item under SIP Affiliate Shop corresponding to the P Item.</p>
	ModelMappingList []ModelMapping `json:"model_mapping_list"` // [Required] <p>If the P Item does not have model, then the model_mapping_list will not be returned.</p>
}

type ApplyItemSetRequest struct {
	SessionId  int64   `json:"session_id"`   // [Required] <p>The identifier of livestream session.</p>
	ItemSetIds []int64 `json:"item_set_ids"` // [Required] <p>List of item set id to apply.</p>
}

type ApplyItemSetResponse struct {
	BaseResponse `json:",inline"` // Common response fields
	Response     interface{}      `json:"response"` // Actual response data
}

type AptOrder struct {
	OrderSn             string  `json:"order_sn"`               // [Required] <p>Order SN.</p>
	OrderCreateTime     int64   `json:"order_create_time"`      // [Required] <p>Order Paid Time.</p>
	ArrangePickUpTime   int64   `json:"arrange_pick_up_time"`   // [Required] <p>Seller arrange pick up time.</p>
	ActualPickUpTime    int64   `json:"actual_pick_up_time"`    // [Required] <p>Courier actual pick up time.</p>
	PreparationDays     float64 `json:"preparation_days"`       // [Required] <p>Preparation Days.</p>
	ShippingChannel     string  `json:"shipping_channel"`       // [Required] <p>Logistics Company.</p>
	FirstMileType       string  `json:"first_mile_type"`        // [Required] <p>First mile shipping type. Applicable values:</p><p>Pickup</p><p>Drop off</p>
	FirstMileTrackingNo string  `json:"first_mile_tracking_no"` // [Required] <p>First Mile Tracking No.</p>
}

type Attribute struct {
	AttributeId        int64            `json:"attribute_id"`         // [Required] Attribute ID
	AttributeValueList []AttributeValue `json:"attribute_value_list"` // [Required]
}

type AttributeAttributeValue struct {
	ValueId int64 `json:"value_id"` // [Required] ID of attribute value.
}

type AttributeInfo struct {
	InputType           int64    `json:"input_type"`            // [Required] <p>SINGLE_DROP_DOWN = 1 <br />SINGLE_COMBO_BOX = 2 <br />FREE_TEXT_FILED&nbsp; &nbsp; &nbsp; &nbsp; = 3 <br />MULTI_DROP_DOWN&nbsp; &nbsp;= 4 <br />MULTI_COMBO_BOX&nbsp; &nbsp;= 5&nbsp;<br /></p>
	InputValidationType int64    `json:"input_validation_type"` // [Required] <p>VALIDATOR_NO_VALIDATE_TYPE =&nbsp; 0 <br />VALIDATOR_INT_TYPE = 1&nbsp;</p><p>VALIDATOR_STRING_TYPE = 2</p><p>VALIDATOR_FLOAT_TYPE = 3&nbsp;</p><p>VALIDATOR_DATE_TYPE = 4<br /></p>
	FormatType          int64    `json:"format_type"`           // [Required] <p>FORMAT_NORMAL = 1</p><p>FORMAT_QUANTITATIVE_WITH_UNIT = 2<br /></p>
	DateFormatType      int64    `json:"date_format_type"`      // [Required] <p>YEAR_MONTH_DATE = 0 (DD/MM/YYYY)</p><p>YEAR_MONTH = 1 (MM/YYYY)<br /></p>
	AttributeUnitList   []string `json:"attribute_unit_list"`   // [Required] <p>Attribute's available units list</p>
	MaxValueCount       int64    `json:"max_value_count"`       // [Required] <p>Max selected value count&nbsp;</p>
	Introduction        string   `json:"introduction"`          // [Required] <p>Introduction for special Attribute</p>
	IsOem               bool     `json:"is_oem"`                // [Required]
	SupportSearchValue  bool     `json:"support_search_value"`  // [Required] <p>Indicates whether this attribute has searchable values.</p><p>If yes, please call <b>v2.product.search_attribute_value_list</b> to get the default values</p>
}

type AttributeTree struct {
	AttributeId        int64                         `json:"attribute_id"`         // [Required] <p>Attribute ID</p>
	Mandatory          bool                          `json:"mandatory"`            // [Required] <p>Is mandatory or not</p>
	Name               string                        `json:"name"`                 // [Required] <p>Attribute Name</p>
	AttributeValueList []AttributeTreeAttributeValue `json:"attribute_value_list"` // [Required] <p>All available values for this attribute<br /></p>
	AttributeInfo      *AttributeInfo                `json:"attribute_info"`       // [Required] <p>Attribute extra info</p>
	MultiLang          []MultiLang                   `json:"multi_lang"`           // [Required] <p>Attribute translate info</p>
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

type AttributeTreeAttributeValue struct {
	ValueId            int64         `json:"value_id"`             // [Required] <p>Value ID</p>
	Name               string        `json:"name"`                 // [Required] <p>Value name</p>
	ValueUnit          string        `json:"value_unit"`           // [Required] <p>Value unit</p>
	ChildAttributeList []interface{} `json:"child_attribute_list"` // [Required] <p>Child attributes for the value of parent attribute<br />The structure content is the same as attribute_tree<br /></p>
	MultiLang          []MultiLang   `json:"multi_lang"`           // [Required] <p>Translate results for display</p>
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

type AuthedMerchant struct {
	Region     string `json:"region"`      // [Required] Merchant's area
	AuthTime   int64  `json:"auth_time"`   // [Required] The timestamp when the merchant was authorized to the partner.
	ExpireTime int64  `json:"expire_time"` // [Required] Use this field to indicate the expiration date for merchant authorization.
}

type AuthedShop struct {
	Region          string         `json:"region"`             // [Required] Shop's area
	AuthTime        int64          `json:"auth_time"`          // [Required] The timestamp when the shop was authorized to the partner.
	ExpireTime      int64          `json:"expire_time"`        // [Required] Use this field to indicate the expiration date for shop authorization.
	SipAffiShopList []SipAffiShops `json:"sip_affi_shop_list"` // [Required] SIP affiliate shops info list
}

type AuthorisedBrand struct {
	BrandId   int64  `json:"brand_id"`   // [Required] <p>ID of the authorised brand, it may be the same in different regions.</p>
	BrandName string `json:"brand_name"` // [Required] <p>Name of the authorised brand.</p>
}

type AutoBiddingInfo struct {
	RoasTarget float64 `json:"roas_target"` // [Required] <p>the ROAS target for each campaign with auto bidding</p>
}

type AutoProductAdsInfo struct {
	ProductName string `json:"product_name"` // [Required] <p>the name of product</p>
	Status      string `json:"status"`       // [Required] <p>learning, ongoing, paused, ended, unavailable<br /></p>
	ItemId      int64  `json:"item_id"`      // [Required] <p>Unique identifier for the product.</p>
}

type BanUserCommentRequest struct {
	SessionId int64 `json:"session_id"`  // [Required] <p>The identifier of livestream session.</p>
	BanUserId int64 `json:"ban_user_id"` // [Required] <p>The user id that will be banned from posting comments.</p>
}

type BanUserCommentResponse struct {
	BaseResponse `json:",inline"` // Common response fields
	Response     interface{}      `json:"response"` // Actual response data
}

type BatchDownload struct {
	Start          int64  `json:"start"`                     // [Required] <p>Format YYYYMMDD</p><p>e.g. 20240101</p>
	End            int64  `json:"end"`                       // [Required] <p>Format YYYYMMDD</p><p>e.g. 20240101</p>
	DocumentType   int64  `json:"document_type"`             // [Required] <p>1 = Remessa</p><p>2 = Return</p><p>3 = Symbolic Return</p><p>4 = Sale</p><p>5 = Entrada</p><p>6 = Symbolic Remessa</p><p>7 = all</p>
	FileType       int64  `json:"file_type"`                 // [Required] <p>1 = xml only</p><p>2 = pdf only</p><p>3 = both</p>
	DocumentStatus *int64 `json:"document_status,omitempty"` // [Optional] <p>1= authorized only</p><p>2= cancelled</p><p><br /></p><p>Default:&nbsp;If document_status not passed or passed empty, means documents under ALL status (both authorized and cancelled) must be included</p>
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

type BillingAddress struct {
	State        string `json:"state"`        // [Required] <p>State of the billing address.</p>
	City         string `json:"city"`         // [Required] <p>City of&nbsp;the billing address.</p>
	Address      string `json:"address"`      // [Required] <p>Specific detail of&nbsp;the billing address.</p>
	Zipcode      string `json:"zipcode"`      // [Required] <p>ZIP code&nbsp;of&nbsp;the billing address.</p>
	Neighborhood string `json:"neighborhood"` // [Required] <p>Neighborhood&nbsp;of&nbsp;the billing address.</p>
}

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

type Booking struct {
	BookingSn string `json:"booking_sn"` // [Required] <p>Shopee's unique identifier for a booking.<br /></p>
}

type BookingItem struct {
	ItemId                 int64        `json:"item_id"`                  // [Required] <p>Shopee's unique identifier for an item.<br /></p>
	ItemName               string       `json:"item_name"`                // [Required] <p>The name of the item.<br /></p>
	ItemSku                string       `json:"item_sku"`                 // [Required] <p>A item SKU (stock keeping unit) is an identifier defined by a seller, sometimes called parent SKU. Item SKU can be assigned to an item in Shopee Listings.<br /></p>
	ModelId                int64        `json:"model_id"`                 // [Required] <p>ID of the model that belongs to the same item.<br /></p>
	ModelName              string       `json:"model_name"`               // [Required] <p>Name of the model that belongs to the same item. A seller can offer models of the same item. For example, the seller could create a fixed-priced listing for a t-shirt design and offer the shirt in different colors and sizes. In this case, each color and size combination is a separate model. Each model can have a different quantity and price.<br /></p>
	ModelSku               string       `json:"model_sku"`                // [Required] <p>A model SKU (stock keeping unit) is an identifier defined by a seller. It is only intended for the seller's use. Many sellers assign a SKU to an item of a specific type, size, and color, which are models of one item in Shopee Listings.<br /></p>
	ModelQuantityPurchased int64        `json:"model_quantity_purchased"` // [Required] <p>The number of identical items from one listing/item in the same booking.<br /></p>
	Weight                 float64      `json:"weight"`                   // [Required] <p>The weight of the item<br /></p>
	ProductLocationId      string       `json:"product_location_id"`      // [Required] <p>The fulfilment warehouse ID(s) of the items in the booking. (Multi-Warehouse sellers only)</p>
	ImageInfo              *OptionImage `json:"image_info"`               // [Required] <p>Image info of the product.<br /></p>
}

type BookingRecipientAddress struct {
	Name        string `json:"name"`         // [Required] <p>Recipient's name for the address.<br /></p>
	Phone       string `json:"phone"`        // [Required] <p>Recipient's phone number input when booking was placed.<br /></p>
	Town        string `json:"town"`         // [Required] <p>The town of the recipient's address. Whether there is a town will depend on the region and/or country.<br /></p>
	District    string `json:"district"`     // [Required] <p>The district of the recipient's address. Whether there is a district will depend on the region and/or country.<br /></p>
	City        string `json:"city"`         // [Required] <p>The city of the recipient's address. Whether there is a city will depend on the region and/or country.<br /></p>
	State       string `json:"state"`        // [Required] <p>The state/province of the recipient's address. Whether there is a state/province will depend on the region and/or country.<br /></p>
	Region      string `json:"region"`       // [Required] <p>The two-digit code representing the region of the Recipient.<br /></p>
	Zipcode     string `json:"zipcode"`      // [Required] <p>Recipient's postal code.<br /></p>
	FullAddress string `json:"full_address"` // [Required] <p>The full address of the recipient, including country, state, even street, and etc.<br /></p>
}

type BoostItemRequest struct {
	ItemIdList []int64 `json:"item_id_list"` // [Required] Shopee's unique identifier for an item, limit:[1,5]
}

type BoostItemResponse struct {
	BaseResponse `json:",inline"`      // Common response fields
	Response     BoostItemResponseData `json:"response"` // Actual response data
}

type BoostItemResponseData struct {
	FailureList []Failure    `json:"failure_list"` // [Required]
	SuccessList *SuccessList `json:"success_list"` // [Required]
}

type BoundWhs struct {
	WhsRegion string `json:"whs_region"` // [Required] <p>the warehouse region bound with the shop</p>
	WhsIds    string `json:"whs_ids"`    // [Required] <p>the warehouse id bound with the shop</p>
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

type Brand struct {
	BrandId           int64  `json:"brand_id"`            // [Required] Id of brand.
	OriginalBrandName string `json:"original_brand_name"` // [Required] Original name of brand.
}

type Budget struct {
	RecommendedBudget float64 `json:"recommended_budget"` // [Required] <p>Recommended Suggested Budget</p>
	MinBudget         float64 `json:"min_budget"`         // [Required] <p>Minimun&nbsp;Suggested Budget</p>
	MaxBudget         float64 `json:"max_budget"`         // [Required] <p>Maximum&nbsp;Suggested Budget</p>
}

type BundleDeal struct {
	BundleDealId   int64           `json:"bundle_deal_id"`   // [Required] Shopee's unique identifier for a bundle deal activity.
	Name           string          `json:"name"`             // [Required] Title of the bundle deal
	StartTime      int64           `json:"start_time"`       // [Required] The time when bundle deal activity start.
	EndTime        int64           `json:"end_time"`         // [Required] The time when bundle deal activity end.
	BundleDealRule *BundleDealRule `json:"bundle_deal_rule"` // [Required]
	PurchaseLimit  int64           `json:"purchase_limit"`   // [Required] Maximum number of bundle deals that can be bought by a buyer.
}

type BundleDealRule struct {
	RuleType           int64             `json:"rule_type"`           // [Required] The bundle deal rule type：FIX_PRICE = 1 ；DISCOUNT_PERCENTAGE = 2； DISCOUNT_VALUE = 3
	DiscountValue      float64           `json:"discount_value"`      // [Required] The deducted price when when buying a bundle deal. Need to input it when the bundle deal rule type is 3
	FixPrice           float64           `json:"fix_price"`           // [Required] The amount of the buyer needs to spend to purchase a bundle deal. Need to input it when the bundle deal rule type is 1
	DiscountPercentage int64             `json:"discount_percentage"` // [Required]  The discount that the buyer can get when buying a bundle deal. Need to input it when the bundle deal rule type is 2
	MinAmount          int64             `json:"min_amount"`          // [Required]  The quantity of items that need buyer to combine purchased
	AdditionalTiers    []AdditionalTiers `json:"additional_tiers"`    // [Required] <p>Use to create tiered discount for bundle deals, a max of 2 additional tiers are allowed to create bundle deals like buy 2 get 10% off, buy 3 for 15% off, buy 4 for 20% off;&nbsp;For each tier, we will need to set the following 4 values based on bundle deal type&nbsp;+&nbsp; &nbsp; min_amount = IntAttribute()&nbsp;+&nbsp;&nbsp;&nbsp; fix_price = IntAttribute()&nbsp;+&nbsp;&nbsp;&nbsp; discount_percentage = IntAttribute()&nbsp;+&nbsp;&nbsp;&nbsp; discount_value = IntAttribute()Note: for additional tiers, the fix price, discount_percentage, discount_value should be consistent with tier 1<br /></p>
}

type BuyerPaymentInfo struct {
	BuyerPaymentMethod    string  `json:"buyer_payment_method"`     // [Required] <p>The payment method used by buyer.</p>
	BuyerServiceFee       float64 `json:"buyer_service_fee"`        // [Required] <p>An additional service fee charged by shopee to Buyer at checkout. currently only applicable to ID region.<br />it is an initial value (will not be updated after RR/cancellation)<br /></p>
	BuyerTaxAmount        float64 `json:"buyer_tax_amount"`         // [Required] <p>The tax amount paid by buyer.</p><p>currently this is a custom tax charged to CB orders in TW,CL,MX</p>
	BuyerTotalAmount      float64 `json:"buyer_total_amount"`       // [Required] <p>The total amount paid by buyer at checkout moment.<br /></p>
	CreditCardPromotion   float64 `json:"credit_card_promotion"`    // [Required] <p>The promotion provided by credit card.<br />it is an initial value (will not be updated after RR/cancellation)<br /></p>
	IcmsTaxAmount         float64 `json:"icms_tax_amount"`          // [Required] <p>The icms tax paid by buyer. this is only applicable to BR region<br />it is an initial value (will not be updated after RR/cancellation)<br /></p>
	ImportTaxAmount       float64 `json:"import_tax_amount"`        // [Required] <p>The import tax paid by buyer.&nbsp;<br />it is an initial value (will not be updated after RR/cancellation)<br /></p>
	InitialBuyerTxnFee    float64 `json:"initial_buyer_txn_fee"`    // [Required] <p>Transaction fee paid by buyer for the order. (Only display for non cb sip affiliate shop. ).&nbsp;Most regions would have this fee charged to buyer at checkout depending on the fee rules applied in each region.<br />it is an initial value (will not be updated after RR/cancellation)<br /></p>
	InsurancePremium      float64 `json:"insurance_premium"`        // [Required] <p>The insurance premium paid by buyer. Currently only applicable to some regions like PH, MY, ID, VN, SG and TH it is an initial value (will not be updated after RR/cancellation)</p>
	IofTaxAmount          float64 `json:"iof_tax_amount"`           // [Required] <p>The iof tax paid by buyer.&nbsp;</p><p>it is an initial value (will not be updated after RR/cancellation)<br /></p>
	IsPaidByCreditCard    bool    `json:"is_paid_by_credit_card"`   // [Required] <p>Whether this order is paid by credit card. it's related to payment channel used at checkout.</p><p>Value: false,true</p>
	MerchantSubtotal      float64 `json:"merchant_subtotal"`        // [Required] <p>The total item price paid by buyer at checkout.</p><p>it is an initial value and will not be updated after RR/cancellation<br /></p>
	SellerVoucher         float64 `json:"seller_voucher"`           // [Required] <p>The voucher provided by seller to offset some value needs to be paid by buyer.</p><p>it is an initial value (will not be updated after RR/cancellation)<br /></p>
	ShippingFee           float64 `json:"shipping_fee"`             // [Required] <p>The shipping fee paid by buyer. (Only display for non cb sip affiliate shop.&nbsp;<br /></p><p>it is an initial value (will not be updated after RR/cancellation)<br /></p>
	ShippingFeeSstAmount  float64 `json:"shipping_fee_sst_amount"`  // [Required] <p>The shipping fee sst paid by buyer. Currently apply to MY region only&nbsp;</p><p>it is an initial value (will not be updated after RR/cancellation)<br /></p>
	ShopeeVoucher         float64 `json:"shopee_voucher"`           // [Required] <p>The voucher provided by Shopee to offset the amount need to be paid by buyer.</p><p>it is an initial value (will not be updated after RR/cancellation)<br /></p>
	ShopeeCoinsRedeemed   float64 `json:"shopee_coins_redeemed"`    // [Required] <p>This is an amount of coin used by buyer at checkout to offset some amount to be paid.&nbsp;</p><p>it is an initial value (will not be updated after RR/cancellation)<br /></p>
	BuyerPaidPackagingFee float64 `json:"buyer_paid_packaging_fee"` // [Required] <p>The fee charged to the buyer for packaging materials</p>
	TradeInBonus          float64 `json:"trade_in_bonus"`           // [Required] <p>The total amount of all trade-in bonuses applied to a transaction. This value is the summation of the bonuses contributed by all four parties: Shopee, Seller, Bank and Trade-in Partner.</p>
	BulkyHandlingFee      float64 `json:"bulky_handling_fee"`       // [Required] <p>A fee charged to the buyer for the special handling and transportation required for items that exceed a specified weight or dimension threshold. Only for ID local seller</p>
	DiscountPix           float64 `json:"discount_pix"`             // [Required] <p>The discount provided by PIX&nbsp;(Only applicable for BR region)</p>
}

type BuyerPreferDeliveryTime struct {
	SlotId      string `json:"slot_id"`     // [Required] <p>The slot which buyer choose<br /></p>
	StartTime   string `json:"start_time"`  // [Required] <p>The start time of a day buyer prefer to receive the packages<br /></p>
	EndTime     string `json:"end_time"`    // [Required] <p>The end time of a day buyer prefer to receive the packages.<br /></p>
	Description string `json:"description"` // [Required] <p>The detailed instructions of the package delivering.<br /></p>
}

type BuyerVideos struct {
	ThumbnailUrl string `json:"thumbnail_url"` // [Required] <p>The thumbnail url of video</p>
	VideoUrl     string `json:"video_url"`     // [Required] <p>The url of video</p>
}

type Campaign struct {
	CampaignId        int64     `json:"campaign_id"`        // [Required] <p>the unique id per campaign</p>
	AdType            string    `json:"ad_type"`            // [Required] <p>auto, manual<br /></p>
	CampaignPlacement string    `json:"campaign_placement"` // [Required] <p>search, discovery, all<br /></p>
	AdName            string    `json:"ad_name"`            // [Required] <p>the name of each ad</p>
	MetricsList       []Metrics `json:"metrics_list"`       // [Required] <p>the performance metric list</p>
}

type CampaignDuration struct {
	StartTime int64 `json:"start_time"` // [Required] <p>The start date for each campaign. please kindly note that if this campaign is no end date, please pass today's date as the start date&nbsp;</p>
	EndTime   int64 `json:"end_time"`   // [Required] <p>the end date per campaign. please kindly note that if it's no limit, so you don't need pass anything and if it's unlimited, the end time would return 0</p>
}

type CampaignMetrics struct {
	Hour              int64   `json:"hour"`                // [Required] <p>This is the parameter to indicate each hour the performance record belongs to.<br /></p>
	Date              string  `json:"date"`                // [Required] <p>This is the parameter of the single date on which requestor wants to check the hourly performance<br /></p>
	Impression        int64   `json:"impression"`          // [Required] <p>The number of times shoppers see your ad.</p>
	Clicks            int64   `json:"clicks"`              // [Required] <p>The number of times shoppers click on your ad. (Note: Shopee filters out repeated clicks from the same shopper that occur within a short time frame.)<br /></p>
	Ctr               float64 `json:"ctr"`                 // [Required] <p>The click-through rate (CTR) measures how often shoppers end up clicking on your ad after seeing it. It is the number of clicks on your ad divided by the number of times your ad is seen. CTR = clicks ÷ impressions × 100%.</p>
	Expense           float64 `json:"expense"`             // [Required] <p>The amount spent on your ad.<br /></p>
	BroadGmv          float64 `json:"broad_gmv"`           // [Required] <p>The amount of sales revenue generated from shoppers purchasing products within 7 days of them clicking on your ad.<br /></p>
	BroadOrder        int64   `json:"broad_order"`         // [Required] <p>The number of times shoppers purchased any product from your shop within 7 days of them clicking on your ad.<br /></p>
	BroadOrderAmount  int64   `json:"broad_order_amount"`  // [Required] <p>The total quantity of products purchased by shoppers within 7 days of them clicking on your ad.<br /></p>
	BroadRoi          float64 `json:"broad_roi"`           // [Required] <p>Return on ad spend (ROAS) measures how much revenue is generated by your ad relative to the cost of the ad. It is the amount of sales revenue attributed to your ad divided by the amount spent on the ad. ROAS = GMV ÷ expense. (Note: We recommend monitoring ROAS trends on a weekly basis.)<br /></p>
	BroadCir          float64 `json:"broad_cir"`           // [Required] <p>The advertising cost of sales (ACOS) measures how much your ad costs relative to the revenue the ad generates. It is the amount spent on your ad divided by the amount of sales revenue attributed to the ad. ACOS = expense ÷ GMV × 100%.<br /></p>
	Cr                float64 `json:"cr"`                  // [Required] <p>The conversion rate measures how often shoppers end up purchasing something from your shop after clicking on your ad. It is the number of conversions attributed to your ad divided by the number of clicks on the ad. Conversion rate = conversions ÷ clicks × 100%.<br /></p>
	Cpc               float64 `json:"cpc"`                 // [Required] <p>The cost per conversion is how much each conversion costs, on average. It is the amount spent on your ad divided by the number of conversions attributed to the ad. Cost per conversion = expense ÷ conversions.<br /></p>
	DirectOrder       int64   `json:"direct_order"`        // [Required] <p>The number of times shoppers purchased the advertised product within 7 days of them clicking on the ad.<br /></p>
	DirectOrderAmount int64   `json:"direct_order_amount"` // [Required] <p>The total quantity of the advertised product purchased by shoppers within 7 days of them clicking on the ad.<br /></p>
	DirectGmv         float64 `json:"direct_gmv"`          // [Required] <p>The amount of sales revenue generated from shoppers purchasing the advertised product within 7 days of them clicking on the ad.<br /></p>
	DirectRoi         float64 `json:"direct_roi"`          // [Required] <p>The direct return on ad spend, or direct ROAS, measures how much revenue is generated from sales of the advertised product, relative to the cost of the ad. It is the amount of sales revenue for the advertised product attributed to the ad, divided by the amount spent on the ad. Direct ROAS = direct GMV ÷ expense.<br /></p>
	DirectCir         float64 `json:"direct_cir"`          // [Required] <p>The direct advertising cost of sales, or direct ACOS, measures how much your ad costs relative to the revenue generated from sales of the advertised product. It is the amount spent on the ad divided by the amount of sales revenue for the advertised product that is attributed to the ad. Direct ACOS = expense ÷ direct GMV × 100%.<br /></p>
	DirectCr          float64 `json:"direct_cr"`           // [Required] <p>The direct conversion rate measures how often shoppers end up purchasing the advertised product after clicking on the ad. Direct conversion rate is the number of direct conversions divided by the number of clicks. Direct conversion rate = direct conversions ÷ clicks × 100%.<br /></p>
	Cpdc              float64 `json:"cpdc"`                // [Required] <p>The cost per direct conversion is how much each direct conversion costs, on average. It is the amount spent on your ad divided by the number of direct conversions attributed to the ad. Cost per direct conversion = expense ÷ direct conversions.<br /></p>
}

type CancelDisputeRequest struct {
	ReturnSn string `json:"return_sn"` // [Required] <p>Shopee's unique serial number identifier for a Return Refund request.</p><p><br /></p><p>Note:&nbsp;Sellers can only cancel compensation disputes, not normal disputes. This means that sellers can only cancel disputes <b>when the return_status is ACCEPTED and the compensation_status is COMPENSATION_REQUESTED</b>.</p>
	Email    string `json:"email"`     // [Required] <p>The operator's email address. For operation record keeping purposes (same as v2.returns.dispute API).</p>
}

type CancelDisputeResponse struct {
	BaseResponse `json:",inline"`          // Common response fields
	Response     CancelDisputeResponseData `json:"response"` // Actual response data
}

type CancelDisputeResponseData struct {
	ReturnSn string `json:"return_sn"` // [Required] <p>Shopee's unique serial number identifier for a Return Refund request.</p>
	Message  string `json:"message"`   // [Required] <p>To indicate whether the cancel dispute operation is successful or failed.</p>
}

type CancelOrderRequest struct {
	OrderSn      string                   `json:"order_sn"`            // [Required] Shopee's unique identifier for an order.
	CancelReason string                   `json:"cancel_reason"`       // [Required] <p>The reason seller want to cancel this order. Applicable values:&nbsp;</p><p>OUT_OF_STOCK</p><p>CUSTOMER_REQUEST</p><p>UNDELIVERABLE_AREA (Note: Only apply for TW and MY)</p><p>COD_NOT_SUPPORTED</p>
	ItemList     []CancelOrderRequestItem `json:"item_list,omitempty"` // [Optional] Required when cancel_reason is OUT_OF_STOCK.
}

type CancelOrderRequestItem struct {
	ItemId  int64 `json:"item_id"`  // [Required] Failed item id.
	ModelId int64 `json:"model_id"` // [Required] Failed model id.
}

type CancelOrderResponse struct {
	BaseResponse `json:",inline"`        // Common response fields
	Response     CancelOrderResponseData `json:"response"` // Actual response data
}

type CancelOrderResponseData struct {
	UpdateTime int64 `json:"update_time"` // [Required] The time when the order is updated.
}

type CancelVideoUploadRequest struct {
	VideoUploadId string `json:"video_upload_id"` // [Required] The ID of this upload session, returned in init_video_upload.
}

type CancelVideoUploadResponse struct {
	BaseResponse `json:",inline"` // Common response fields
}

type CancellationOrder struct {
	OrderSn          string `json:"order_sn"`          // [Required] <p>Order SN.</p>
	CancellationType int64  `json:"cancellation_type"` // [Required] <p>Cancellation Type.&nbsp;Applicable values:&nbsp;</p><p>1: System Cancellation</p><p>2: Seller Cancellation<br /></p>
	DetailedReason   int64  `json:"detailed_reason"`   // [Required] <p>Reason. Applicable values:&nbsp;</p><p>1001: Return Refund</p><p>1002: Parcel Split Cancellation</p><p>1003: First Mile Pick up fail</p><p>1004: Order inclusion<br /></p><p>10005: Out of Stock<br />10006: Undeliverable area<br />10007: Cannot support COD<br />10008: Logistics request cancelled<br />10009: Logistics pickup failed<br />10010: Logistics not ready<br />10011: Inactive seller<br />10012: Seller did not ship order<br />10013: Order did not reach warehouse<br /></p><p>10014: Seller asked to cancel<br /></p><p>10015: Non-receipt<br />10016: Wrong item<br />10017: Damaged item<br />10018: Incomplete product<br />10019: Fake item<br />10020: Functional Damage<br />10021: Return Refund</p>
}

type Category struct {
	CategoryId           int64  `json:"category_id"`            // [Required] <p>ID for category.<br /></p>
	ParentCategoryId     int64  `json:"parent_category_id"`     // [Required] <p>ID for parent category.<br /></p>
	OriginalCategoryName string `json:"original_category_name"` // [Required] <p>English category name.<br /></p>
	DisplayCategoryName  string `json:"display_category_name"`  // [Required] <p>Display category name, it depends on what language you have uploaded<br /></p>
	HasChildren          bool   `json:"has_children"`           // [Required] <p>Whether this category has active children category.<br /></p>
}

type CategoryRecommendRequest struct {
	ItemName          string  `json:"item_name" url:"item_name"`                                         // [Required] name of item
	ProductCoverImage *string `json:"product_cover_image,omitempty" url:"product_cover_image,omitempty"` // [Optional] <p>Please use the image id returned by v2.media_space.upload_image api, we will ignore if this field is empty string<br /></p>
}

type CategoryRecommendResponse struct {
	BaseResponse `json:",inline"`              // Common response fields
	Response     CategoryRecommendResponseData `json:"response"` // Actual response data
}

type CategoryRecommendResponseData struct {
	CategoryId []int64 `json:"category_id"` // [Required] Shopee's unique identifier for a category.
}

type Certification struct {
	CertificationNo     string                `json:"certification_no"`      // [Required] <p>Certification No.</p>
	PermitId            int64                 `json:"permit_id"`             // [Required] <p>Permit ID, get from&nbsp;<span style="font-size:14px;">v2.product.get_product_certification_rule</span></p><br />
	ExpiryDate          *int64                `json:"expiry_date,omitempty"` // [Optional] <p>Expiry timestamp. Required for PH, but not needed for TW.</p>
	CertificationProofs []CertificationProofs `json:"certification_proofs"`  // [Required] <p>An array of proof documents for the certification; each element represents one proof file.&lt;path&gt;&lt;/path&gt;</p>
}

type CertificationCertificationProofs struct {
	ImageId  string  `json:"image_id"`  // [Required] <p>The unique image ID of the certification proof, returned by the image upload API.</p>
	FileName string  `json:"file_name"` // [Required] <p>The name of the uploaded certification proof file.</p>
	Ratio    float64 `json:"ratio"`     // [Required] <p>image weight/ image height<br /><br />Will be optional in the future; can input 0.75 by default</p>
}

type CertificationInfo struct {
	CertificationList []Certification `json:"certification_list,omitempty"` // [Optional] <p>Array of certification records for the product, each containing type, certificate number, permit ID, and proof documents.</p><p>  </p><p><br /></p>
}

type CertificationInfoCertification struct {
	CertificationNo     string                            `json:"certification_no"`               // [Required] <p>Certification number issued by the regulatory or certifying authority; uniquely identifies the certification.</p><p>refer to<br /><a href="https://seller.shopee.ph/edu/article/24236" target="_blank">https://seller.shopee.ph/edu/article/24236</a></p>
	PermitId            int64                             `json:"permit_id"`                      // [Required] <p>Platform-defined permit ID used to link to a specific certification template or rule.</p><p>get from&nbsp;v2.product.get_product_certification_rule</p>
	ExpiryDate          *int64                            `json:"expiry_date,omitempty"`          // [Optional] <p>Expiry timestamp. Required for PH, but not needed for TW.</p>
	CertificationProofs *CertificationCertificationProofs `json:"certification_proofs,omitempty"` // [Optional] <p>An array of proof documents for the certification; each element represents one proof file.&lt;path&gt;&lt;/path&gt;</p>
}

type CertificationInfoCertificationCertificationProofs struct {
	ImageId  string  `json:"image_id"`  // [Required] <p>The unique image ID of the certification proof, returned by the image upload API.</p>
	Ratio    float64 `json:"ratio"`     // [Required] <p>image weight/ image height.<br /></p>
	FileName string  `json:"file_name"` // [Required] <p>The name of the uploaded certification proof file.</p>
	ImageUrl string  `json:"image_url"` // [Required] <p>The image url of the proof</p>
}

type CertificationProofs struct {
	FileName string  `json:"file_name"` // [Required] <p>The name of the uploaded certification proof file.</p>
	ImageId  int64   `json:"image_id"`  // [Required] <p>The unique image ID of the certification proof, returned by the image upload API.</p>
	Ratio    float64 `json:"ratio"`     // [Required] <p>image weight/ image height<br /><br />Will be optional in the future; can input 0.75 by default</p>
}

type CertificationRule struct {
	CertificationType int64  `json:"certification_type"` // [Required] <p>type of certification; always=1</p>
	IsMandatory       bool   `json:"is_mandatory"`       // [Required] <p>if this type of certification is mandatory for product</p>
	PermitId          int64  `json:"permit_id"`          // [Required]
	Name              string `json:"name"`               // [Required] <p>Permit Type Name</p>
}

type CheckCreateGmsProductCampaignEligibilityResponse struct {
	BaseResponse `json:",inline"`                                     // Common response fields
	Response     CheckCreateGmsProductCampaignEligibilityResponseData `json:"response"` // Actual response data
}

type CheckCreateGmsProductCampaignEligibilityResponseData struct {
	IsEligible bool   `json:"is_eligible"` // [Required] <p>Indicates if the seller is eligible to create a GMS Campaign</p>
	Reason     string `json:"reason"`      // [Required] <p>The following are the list of reasons for not being able to create a GMS Campaign:&nbsp;</p><p>active_campaign</p><p>&nbsp;&nbsp;&nbsp;&nbsp;There is already an existing GMS Campaign that is active</p><p>not_whitelisted</p><p>&nbsp;&nbsp;&nbsp;&nbsp;The seller is not whitelisted to sz_shop_gmv_max_feature</p><p>not_have_enough_sku</p><p>&nbsp;&nbsp;&nbsp;&nbsp;The seller does not have enough valid items in the shop</p><p>exclusive_with_other_campaign</p><p>&nbsp;&nbsp;&nbsp;&nbsp;Seller is whitelisted to&nbsp;sz_ads_auto_boost</p>
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

type Collection struct {
	IsActivated bool             `json:"is_activated"` // [Required] whether is activated
	ItemList    []CollectionItem `json:"item_list"`    // [Required] a list of item
	TopPicksId  int64            `json:"top_picks_id"` // [Required] collection id
	Name        string           `json:"name"`         // [Required] collection name
}

type CollectionItem struct {
	ItemName                    string  `json:"item_name"`                       // [Required] The name of item.
	ItemId                      int64   `json:"item_id"`                         // [Required] The id of item.
	CurrentPrice                float64 `json:"current_price"`                   // [Required] The price before tax of item.
	InflatedPriceOfCurrentPrice float64 `json:"inflated_price_of_current_price"` // [Required] The price after tax of item.
	Sales                       int64   `json:"sales"`                           // [Required] The sales of item.
}

type Column struct {
	Measurement          *Measurement       `json:"measurement"`            // [Required] <p>this is the column header which means a kind of measurement<br /></p>
	MeasurementValueList []MeasurementValue `json:"measurement_value_list"` // [Required] <p>the list of measurement value<br /></p>
}

type Comment struct {
	CommentId int64  `json:"comment_id"` // [Required] The identity of comment.
	Comment   string `json:"comment"`    // [Required] The content of the comment.
}

type CommentReply struct {
	Reply      string `json:"reply"`       // [Required] The content of reply.
	Hidden     bool   `json:"hidden"`      // [Required] The comment reply is hidden or not.
	CreateTime int64  `json:"create_time"` // [Required] <p>The time the seller replied to the comment.</p>
}

type CommonInfo struct {
	AdType            string            `json:"ad_type"`            // [Required] <p>auto, manual<br /></p>
	AdName            string            `json:"ad_name"`            // [Required] <p>the name of each ad</p>
	CampaignStatus    CampaignStatus    `json:"campaign_status"`    // [Required] <p>ongoing, scheduled, ended, paused, deleted, closed<br /></p>
	BiddingMethod     string            `json:"bidding_method"`     // [Required] <p>auto, manual<br /></p>
	CampaignPlacement string            `json:"campaign_placement"` // [Required] <p>search, discovery, all<br /></p>
	CampaignBudget    float64           `json:"campaign_budget"`    // [Required] <p>The budget per campaign. Please kindly note that if the campaign budget = 0, it means the budget set for this campaign is unlimited</p>
	CampaignDuration  *CampaignDuration `json:"campaign_duration"`  // [Required] <p>the duration per campaign</p>
	ItemIdList        []int64           `json:"item_id_list"`       // [Required] <p>List of unique identifiers for all products under this campaign. If the campaign is using auto product selection it can have between zero and many products. If the campaign is using manual product selection, it has exactly one.</p>
}

type CompanyAddressBreakdown struct {
	CompanyRegion          string `json:"company_region"`           // [Required] <p>Return region value</p><p>- PH, TH only<br /></p>
	CompanyState           string `json:"company_state"`            // [Required] <p>Return value</p><p>- PH: Province</p><p>- TH: Province<br /></p>
	CompanyCity            string `json:"company_city"`             // [Required] <p>Return value</p><p>- PH: City<br /></p>
	CompanyDistrict        string `json:"company_district"`         // [Required] <p>Return value</p><p>- PH: Barangay</p><p>- TH: District<br /></p>
	CompanyTown            string `json:"company_town"`             // [Required] <p>Return value</p><p>- TH: Sub district<br /></p>
	CompanyPostcode        string `json:"company_postcode"`         // [Required] <p>Return postal code</p><p>- TH, PH only<br /></p>
	CompanyDetailedAddress string `json:"company_detailed_address"` // [Required] <p>Return value</p><p>- PH: Detailed address</p><p>- TH: Detailed address<br /></p>
	CompanyAdditionalInfo  string `json:"company_additional_info"`  // [Required] <p>Return value:</p><p>- Empty for PH, TH<br /></p>
	CompanyFullAddress     string `json:"company_full_address"`     // [Required] <p>Concatenation of company address breakdown<br /></p><p>- only has value when invoice_type is company</p><p><br /></p><p><br /></p>
}

type CompatibilityInfo struct {
	VehicleInfoList []VehicleInfo `json:"vehicle_info_list"` // [Required]
}

type ComplaintPolicy struct {
	WarrantyTime                WarrantyTime `json:"warranty_time"`                 // [Required] Time for a warranty claim. Could be ONE_YEAR, TWO_YEARS, OVER_TWO_YEARS.
	ExcludeEntrepreneurWarranty bool         `json:"exclude_entrepreneur_warranty"` // [Required] If True means "I exclude warranty complaints for entrepreneur"
	ComplaintAddressId          int64        `json:"complaint_address_id"`          // [Required] The identity of complaint address.
	AdditionalInformation       string       `json:"additional_information"`        // [Required]  Additional information for complaint policy.
}

type CompleteVideoUploadRequest struct {
	VideoUploadId string      `json:"video_upload_id"` // [Required] The ID of this upload session, returned in init_video_upload.
	PartSeqList   []int64     `json:"part_seq_list"`   // [Required] All uploaded sequence number.
	ReportData    *ReportData `json:"report_data"`     // [Required]
}

type CompleteVideoUploadResponse struct {
	BaseResponse `json:",inline"` // Common response fields
}

type Component struct {
	ComponentItemId  int64  `json:"component_item_id"`            // [Required] <p>ID of the item that composes this kit model.<br /></p>
	ComponentModelId *int64 `json:"component_model_id,omitempty"` // [Optional] <p>ID of the model that composes this kit model.<br /></p>
	Quantity         int64  `json:"quantity"`                     // [Required] <p>The amount of the item/model that composes this kit model.<br /></p>
	MainComponent    *bool  `json:"main_component,omitempty"`     // [Optional] <p>Whether this item/model is the main component for this kit.</p><p>One kit item can only have one item/model as main component.</p>
}

type ConfirmConsumedLostPushMessageRequest struct {
	LastMessageId int64 `json:"last_message_id"` // [Required] <p>The last_message_id returned by v2.push.get_lost_push_message.<br /></p>
}

type ConfirmConsumedLostPushMessageResponse struct {
	BaseResponse `json:",inline"` // Common response fields
}

type ConfirmRequest struct {
	ReturnSn string `json:"return_sn"` // [Required] The serial number of return.
}

type ConfirmResponse struct {
	BaseResponse `json:",inline"`    // Common response fields
	Response     ConfirmResponseData `json:"response"` // Actual response data
}

type ConfirmResponseData struct {
	ReturnSn string `json:"return_sn"` // [Required] The identifier for an API request for error tracking
}

type ConvertImageRequest struct {
	ReturnSn    string      `json:"return_sn"`    // [Required] <p>The serial number of return.<br /></p>
	UploadImage interface{} `json:"upload_image"` // [Required] <p>The proof picture to be uploaded must be within 10MB in size, and the format only supports .jpg, .jpeg, and .png. Only one picture is allowed to be uploaded per request. If multiple pictures are uploaded, only the first picture will be processed.<br /></p>
}

type ConvertImageResponse struct {
	BaseResponse `json:",inline"`         // Common response fields
	Response     ConvertImageResponseData `json:"response"` // Actual response data
}

type ConvertImageResponseData struct {
	Url       string `json:"url"`       // [Required] <p>The link uploaded to the image server can be used with the upload_proof interface.</p>
	Thumbnail string `json:"thumbnail"` // [Required] <p>The image thumbnail.</p>
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

type CreateGmsProductCampaignRequest struct {
	StartDate   string   `json:"start_date"`             // [Required] <p>Start date of Campaign e.g. "30-11-2025". Cannot be earlier than today.<br /></p>
	EndDate     *string  `json:"end_date,omitempty"`     // [Optional] <p>End date of Campaign e.g. "30-11-2025". Do not fill if no end date.&nbsp;<br /></p>
	DailyBudget float64  `json:"daily_budget"`           // [Required] <p>Daily budget for Campaign.<br /></p>
	ReferenceId *string  `json:"reference_id,omitempty"` // [Optional] <p>Input a string</p>
	RoasTarget  *float64 `json:"roas_target,omitempty"`  // [Optional] <p>No input will be GMV Max Auto Bidding (Shop).</p><p>Input 0 for GMV Max Auto Bidding (Shop).</p><p>Input greater than 0 for GMV Max Custom ROAS (Shop).</p><p>If value = 10.123456, it will be taken as 10.1</p><p>If value = 10.199999, it will be taken as 10.1</p>
}

type CreateGmsProductCampaignResponse struct {
	BaseResponse `json:",inline"`                     // Common response fields
	Response     CreateGmsProductCampaignResponseData `json:"response"` // Actual response data
}

type CreateGmsProductCampaignResponseData struct {
	CampaignId int64 `json:"campaign_id"` // [Required] <p>GMS Campaign ID.<br /></p>
}

type CreateManualProductAdsRequest struct {
	ReferenceId           string                  `json:"reference_id"`                      // [Required] <p>A random string used to prevent duplicate ads. If an ads is created successfully, subsequent request using the same reference id will fail<br /></p>
	Budget                float64                 `json:"budget"`                            // [Required] <p>The budget set for the Auto Product Ads<br /></p>
	StartDate             string                  `json:"start_date"`                        // [Required] <p>the start date per campaign. please kindly note that if you want to set unlimited date, you can just pass today's date as the start date</p>
	EndDate               *string                 `json:"end_date,omitempty"`                // [Optional] <p>the end date of each campaign. please kindly note that if you want to set an unlimited campaign, you can keep empty for the end date field</p>
	BiddingMethod         string                  `json:"bidding_method"`                    // [Required] <p>auto, manual<br /></p>
	ItemId                int64                   `json:"item_id"`                           // [Required] <p>Product ID</p>
	RoasTarget            *float64                `json:"roas_target,omitempty"`             // [Optional] <p>the ROAS target for each campaign with auto bidding. If 0, GMV Max / ROI feature is not enabled</p>
	SelectedKeywords      []SelectedKeywords      `json:"selected_keywords,omitempty"`       // [Optional] <p>selected keywords, required for manual bidding mode</p>
	DiscoveryAdsLocations []DiscoveryAdsLocations `json:"discovery_ads_locations,omitempty"` // [Optional] <p>the location settings for manual bidding method</p>
	EnhancedCpc           *bool                   `json:"enhanced_cpc,omitempty"`            // [Optional] <p>Enhanced CPC functionality toggle</p>
	SmartCreativeSetting  *string                 `json:"smart_creative_setting,omitempty"`  // [Optional] <p>Whether to use default or set on/off Smart Creative for this ad. Supported Values: "", "default", "on", "off". Empty string treated as default.</p>
}

type CreateManualProductAdsResponse struct {
	BaseResponse `json:",inline"`                   // Common response fields
	Response     CreateManualProductAdsResponseData `json:"response"` // Actual response data
}

type CreateManualProductAdsResponseData struct {
	CampaignId int64 `json:"campaign_id"` // [Required] <p>The unique identifier for a campaign<br /></p>
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

type CreateSessionRequest struct {
	Title         string  `json:"title"`                 // [Required] <p>The title of livestream session, cannot exceed 200 characters.</p>
	Description   *string `json:"description,omitempty"` // [Optional] <p>The description of livestream session, cannot exceed 200 characters.</p>
	CoverImageUrl string  `json:"cover_image_url"`       // [Required] <p>The cover image URL of livestream session.</p><p>Please call the v2.livestream.upload_image to upload the cover image file and get the cover_image_url.</p>
	IsTest        *bool   `json:"is_test,omitempty"`     // [Optional] <p>Indicate whether the livestream session is for testing purpose only.</p>
}

type CreateSessionResponse struct {
	BaseResponse `json:",inline"`          // Common response fields
	Response     CreateSessionResponseData `json:"response"` // Actual response data
}

type CreateSessionResponseData struct {
	SessionId int64 `json:"session_id"` // [Required] <p>The identifier of livestream session.</p>
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

type CreateShippingDocumentResponseDataResult struct {
	OrderSn       string `json:"order_sn"`       // [Required] <p>Shopee's unique identifier for an order.</p>
	PackageNumber string `json:"package_number"` // [Required] <p>Shopee's unique identifier for the package under an order.</p>
	FailError     string `json:"fail_error"`     // [Required] <p>Indicate error type if one element hit error.</p>
	FailMessage   string `json:"fail_message"`   // [Required] <p>Indicate error details if one element hit error.</p>
}

type CreateShopFlashSaleRequest struct {
	TimeslotId int64 `json:"timeslot_id"` // [Required] <p>can get it from v2.shop_flash_sale.get_time_slot_id API, and you can only use the timeslot which start_time &gt; now</p>
}

type CreateShopFlashSaleResponse struct {
	BaseResponse `json:",inline"`                // Common response fields
	Response     CreateShopFlashSaleResponseData `json:"response"` // Actual response data
}

type CreateShopFlashSaleResponseData struct {
	TimeslotId  int64 `json:"timeslot_id"`   // [Required]
	FlashSaleId int64 `json:"flash_sale_id"` // [Required]
	Status      int64 `json:"status"`        // [Required] <p>the status of shop flash sale</p><p>0: deleted</p><p>1: enabled</p><p>2: disabled</p><p>3: system_rejected</p>
}

type Criteria struct {
	CriteriaId       int64   `json:"criteria_id"`        // [Required]
	MinProductRating float64 `json:"min_product_rating"` // [Required] <p>Product Rating(0.0-5.0),&nbsp;-1 means no limit</p>
	MinLikes         int64   `json:"min_likes"`          // [Required] <p>Likes(s), -1 means no limit<br /></p>
	MustNotPreOrder  bool    `json:"must_not_pre_order"` // [Required] <p>Pre-Order(s)<br /></p>
	MinOrderTotal    int64   `json:"min_order_total"`    // [Required] <p>Orders in the last 30 day(s), -1 means no limit<br /></p>
	MaxDaysToShip    int64   `json:"max_days_to_ship"`   // [Required] Days to Ship,&nbsp;<span style="font-size:16px;">-1 means no limit</span>
	MinRepetitionDay int64   `json:"min_repetition_day"` // [Required] Repetition Control (Same Product cannot Join ISFS within N Days)<br /><p>, -1 means no limit<br /></p>
	MinPromoStock    int64   `json:"min_promo_stock"`    // [Required] <p>Promo Stock,&nbsp;-1 means no limit</p>
	MaxPromoStock    int64   `json:"max_promo_stock"`    // [Required] <p>Promo Stock,&nbsp;-1 means no limit<br /></p>
	MinDiscount      int64   `json:"min_discount"`       // [Required] <p>Discount Limit, 10 means 10%, -1 means no limit</p><p><br /></p>
	MaxDiscount      int64   `json:"max_discount"`       // [Required] <p>Discount Limit, 100 means 100%,&nbsp;-1 means no limit</p>
	MinDiscountPrice int64   `json:"min_discount_price"` // [Required] <p>Discount Limit, -1 means no limit, real min discount price = min_discount_price / 100000<br /></p>
	MaxDiscountPrice int64   `json:"max_discount_price"` // [Required] <p>Discount Limit, -1 means no limit, real max discount price = max_discount_price / 100000</p>
	NeedLowestPrice  bool    `json:"need_lowest_price"`  // [Required] <p>lower than lowest price in last 7 days (exclude Shopee Flash Deals)<br /></p>
}

type Cursor struct {
	NextId   int64 `json:"next_id"`   // [Required]
	PrevId   int64 `json:"prev_id"`   // [Required]
	PageSize int64 `json:"page_size"` // [Required]
}

type Data struct {
	StandardiseVariationList []StandardiseVariation `json:"standardise_variation_list"` // [Required] <p>standardized tier variation tree<br /></p>
}

type DeboostDetails struct {
	ViolationType     string              `json:"violation_type"`     // [Required] <p>Violation types defined by Shopee.&nbsp;Applicable values:&nbsp;</p><p>Prohibited Listing</p><p>Counterfeit and IP Infringement</p><p>Spam</p><p>Inappropriate Image</p><p>Insufficient Information</p><p>Mall Listing Improvement</p><p>Other Listing Improvement<br /></p>
	ViolationReason   string              `json:"violation_reason"`   // [Required] <p>The reason for violation.<br /></p>
	Suggestion        string              `json:"suggestion"`         // [Required] <p>Shopee provides you with suggestions for modifying items.<br /></p>
	SuggestedCategory []SuggestedCategory `json:"suggested_category"` // [Required]
	FixDeadlineTime   int64               `json:"fix_deadline_time"`  // [Required] <p>Action required deadline. Empty if no deadline.</p>
	UpdateTime        int64               `json:"update_time"`        // [Required] <p>Latest update time.<br /></p>
}

type DeleteAddOnDealMainItemRequest struct {
	AddOnDealId  int64   `json:"add_on_deal_id"` // [Required] Shopee's unique identifier for add on deal activity.
	MainItemList []int64 `json:"main_item_list"` // [Required]  The main items added in this add on deal promotion.
}

type DeleteAddOnDealMainItemResponse struct {
	BaseResponse `json:",inline"`                    // Common response fields
	Response     DeleteAddOnDealMainItemResponseData `json:"response"` // Actual response data
}

type DeleteAddOnDealMainItemResponseData struct {
	MainItemList []int64 `json:"main_item_list"` // [Required] The main items added in this add on deal promotion.
	AddOnDealId  int64   `json:"add_on_deal_id"` // [Required] Shopee's unique identifier for add on deal activity.
}

type DeleteAddOnDealRequest struct {
	AddOnDealId int64 `json:"add_on_deal_id"` // [Required] Shopee's unique identifier for an add on deal activity.
}

type DeleteAddOnDealResponse struct {
	BaseResponse `json:",inline"`            // Common response fields
	Response     DeleteAddOnDealResponseData `json:"response"` // Actual response data
}

type DeleteAddOnDealResponseData struct {
	AddOnDealId int64 `json:"add_on_deal_id"` // [Required] Shopee's unique identifier for an add on deal activity.
}

type DeleteAddOnDealSubItemRequest struct {
	AddOnDealId int64                    `json:"add_on_deal_id"` // [Required] Shopee's unique identifier for add on deal activity.
	SubItemList []CancelOrderRequestItem `json:"sub_item_list"`  // [Required] The sub items added in this add on deal promotion.
}

type DeleteAddOnDealSubItemResponse struct {
	BaseResponse `json:",inline"`                   // Common response fields
	Response     DeleteAddOnDealSubItemResponseData `json:"response"` // Actual response data
}

type DeleteAddOnDealSubItemResponseData struct {
	SubItemList []Error `json:"sub_item_list"`  // [Required] The sub items added in this add on deal promotion.
	AddOnDealId int64   `json:"add_on_deal_id"` // [Required] Shopee's unique identifier for add on deal activity.
}

type DeleteAddressRequest struct {
	AddressId int64 `json:"address_id"` // [Required] The identity of address you want to delete.
}

type DeleteAddressResponse struct {
	BaseResponse `json:",inline"` // Common response fields
}

type DeleteBundleDealItemRequest struct {
	BundleDealId int64                             `json:"bundle_deal_id"` // [Required] Shopee's unique identifier for a bundle deal activity.
	ItemList     []DeleteBundleDealItemRequestItem `json:"item_list"`      // [Required] The items deleted in this bundle deal promotion.
}

type DeleteBundleDealItemRequestItem struct {
	ItemId int64 `json:"item_id"` // [Required] The id of related item failed to delete.
}

type DeleteBundleDealItemResponse struct {
	BaseResponse `json:",inline"`                 // Common response fields
	Response     DeleteBundleDealItemResponseData `json:"response"` // Actual response data
}

type DeleteBundleDealItemResponseData struct {
	FailedList  []ResponseDataFailed `json:"failed_list"`  // [Required] Indicate error details.
	SuccessList []interface{}        `json:"success_list"` // [Required] The list of succeed added items
}

type DeleteBundleDealRequest struct {
	BundleDealId int64 `json:"bundle_deal_id"` // [Required] Shopee's unique identifier for a bundle deal activity.
}

type DeleteBundleDealResponse struct {
	BaseResponse `json:",inline"`             // Common response fields
	Response     DeleteBundleDealResponseData `json:"response"` // Actual response data
}

type DeleteBundleDealResponseData struct {
	BundleDealId int64 `json:"bundle_deal_id"` // [Required] Shopee's unique identifier for a bundle deal activity.
}

type DeleteDiscountItemRequest struct {
	DiscountId int64  `json:"discount_id"`        // [Required] Shopee's unique identifier for a discount activity.
	ItemId     int64  `json:"item_id"`            // [Required] Shopee's unique identifier for an item.
	ModelId    *int64 `json:"model_id,omitempty"` // [Optional] Shopee's unique identifier for a variation of an item. If there is no variation of this item, you don't need to input this param. Dafault is 0.
}

type DeleteDiscountItemResponse struct {
	BaseResponse `json:",inline"`               // Common response fields
	Response     DeleteDiscountItemResponseData `json:"response"` // Actual response data
}

type DeleteDiscountItemResponseData struct {
	DiscountId int64   `json:"discount_id"` // [Required] Shopee's unique identifier for a discount activity.
	ErrorList  []Error `json:"error_list"`  // [Required] Detail informations about error.
}

type DeleteDiscountRequest struct {
	DiscountId int64 `json:"discount_id"` // [Required] Shopee's unique identifier for a discount activity.
}

type DeleteDiscountResponse struct {
	BaseResponse `json:",inline"`           // Common response fields
	Response     DeleteDiscountResponseData `json:"response"` // Actual response data
}

type DeleteDiscountResponseData struct {
	DiscountId int64 `json:"discount_id"` // [Required] Shopee's unique identifier for a discount activity.
	ModifyTime int64 `json:"modify_time"` // [Required] The time when discount has been deleted.
}

type DeleteFollowPrizeRequest struct {
	CampaignId int64 `json:"campaign_id"` // [Required] <p>The unique identifier for the created follow prize.<br /></p>
}

type DeleteFollowPrizeResponse struct {
	BaseResponse `json:",inline"`              // Common response fields
	Response     DeleteFollowPrizeResponseData `json:"response"` // Actual response data
}

type DeleteFollowPrizeResponseData struct {
	CampaginId int64 `json:"campagin_id"` // [Required] <p>The unique identifier for the created follow prize.<br /></p>
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

type DeleteItemListRequest struct {
	ShopCategoryId int64   `json:"shop_category_id"` // [Required] The list of items need to be deleted. To note that the items which can be deleted successfully should be under this category.
	ItemList       []int64 `json:"item_list"`        // [Required] ShopCategory's unique identifier.
}

type DeleteItemListResponse struct {
	BaseResponse `json:",inline"`           // Common response fields
	Response     DeleteItemListResponseData `json:"response"` // Actual response data
}

type DeleteItemListResponseData struct {
	ShopCategoryId int64   `json:"shop_category_id"` // [Required] ShopCategory's unique identifier.
	InvalidItemId  []int64 `json:"invalid_item_id"`  // [Required] The list of item ids which are invalid; In other words, the item ids not being under the category.
	CurrentCount   int64   `json:"current_count"`    // [Required] count of items under this shop category after deleting
}

type DeleteItemRequest struct {
	ItemId int64 `json:"item_id"` // [Required] The identity of product item.
}

type DeleteItemResponse struct {
	BaseResponse `json:",inline"` // Common response fields
}

type DeleteModelRequest struct {
	ItemId  int64 `json:"item_id"`  // [Required] ID of item.
	ModelId int64 `json:"model_id"` // [Required] ID of model.
}

type DeleteModelResponse struct {
	BaseResponse `json:",inline"` // Common response fields
}

type DeleteShopCategoryRequest struct {
	ShopCategoryId int64 `json:"shop_category_id"` // [Required] ShopCategory's unique identifier.
}

type DeleteShopCategoryResponse struct {
	BaseResponse `json:",inline"`               // Common response fields
	Response     DeleteShopCategoryResponseData `json:"response"` // Actual response data
}

type DeleteShopCategoryResponseData struct {
	ShopCategoryId int64  `json:"shop_category_id"` // [Required] ShopCategory's unique identifier.
	Msg            string `json:"msg"`              // [Required] The return message of the operation result
}

type DeleteShopFlashSaleItemsRequest struct {
	FlashSaleId int64   `json:"flash_sale_id"` // [Required]
	ItemIds     []int64 `json:"item_ids"`      // [Required] <p>if you delete a item, will delete all models of the item</p>
}

type DeleteShopFlashSaleItemsResponse struct {
	BaseResponse `json:",inline"`                     // Common response fields
	Response     DeleteShopFlashSaleItemsResponseData `json:"response"` // Actual response data
}

type DeleteShopFlashSaleItemsResponseData struct {
	FailedItems []FailedItems `json:"failed_items"` // [Required]
}

type DeleteShopFlashSaleRequest struct {
	FlashSaleId int64 `json:"flash_sale_id"` // [Required] <p>cannot delete ongoing and expired shop flash sale</p>
}

type DeleteShopFlashSaleResponse struct {
	BaseResponse `json:",inline"`                // Common response fields
	Response     DeleteShopFlashSaleResponseData `json:"response"` // Actual response data
}

type DeleteShopFlashSaleResponseData struct {
	TimeslotId  int64 `json:"timeslot_id"`   // [Required]
	FlashSaleId int64 `json:"flash_sale_id"` // [Required]
	Status      int64 `json:"status"`        // [Required] <p>the status of shop flash sale</p><p>0: deleted</p><p>1: enabled</p><p>2: disabled</p><p>3: system_rejected</p>
}

type DeleteShowItemRequest struct {
	SessionId int64 `json:"session_id"` // [Required] <p>The identifier of livestream session.</p>
}

type DeleteShowItemResponse struct {
	BaseResponse `json:",inline"` // Common response fields
	Response     interface{}      `json:"response"` // Actual response data
}

type DeleteSipDiscountRequest struct {
	Region string `json:"region"` // [Required] <p>The region of SIP affiliate shop that needs to delete discount.</p>
}

type DeleteSipDiscountResponse struct {
	BaseResponse `json:",inline"`              // Common response fields
	Response     DeleteSipDiscountResponseData `json:"response"` // Actual response data
}

type DeleteSipDiscountResponseData struct {
	Region string `json:"region"` // [Required] <p>The region of SIP affiliate shop that needs to delete discount.<br /></p>
}

type DeleteSpecialOperatingHourRequest struct {
	Name string `json:"name"` // [Required] <p>Name of the special operating hour which can be retrieved from v2.logistics.get_operating_hours</p>
}

type DeleteSpecialOperatingHourResponse struct {
	BaseResponse `json:",inline"` // Common response fields
}

type DeleteTopPicksRequest struct {
	TopPicksId int64 `json:"top_picks_id"` // [Required] collection id
}

type DeleteTopPicksResponse struct {
	BaseResponse `json:",inline"`           // Common response fields
	Response     DeleteTopPicksResponseData `json:"response"` // Actual response data
}

type DeleteTopPicksResponseData struct {
	TopPicksId int64 `json:"top_picks_id"` // [Required] collection id
}

type DeleteVoucherRequest struct {
	VoucherId int64 `json:"voucher_id"` // [Required] The unique identifier for the voucher you want to delete.
}

type DeleteVoucherResponse struct {
	BaseResponse `json:",inline"`          // Common response fields
	Response     DeleteVoucherResponseData `json:"response"` // Actual response data
}

type DeleteVoucherResponseData struct {
	VoucherId int64 `json:"voucher_id"` // [Required] The unique identifier for the voucher it is being deleted.
}

type Description struct {
	DescriptionType     DescriptionType                 `json:"description_type"`     // [Required] <p>Type of description : values: See Data Definition- description_type (normal , extended).<br /></p>
	Description         string                          `json:"description"`          // [Required] <p>If description_type is normal , Description information will be returned through this field，else description will be empty.<br /></p>
	ExtendedDescription *DescriptionExtendedDescription `json:"extended_description"` // [Required] <p>If description_type is extended , Description information will be returned through this field.<br /></p>
}

type DescriptionExtendedDescription struct {
	FieldList []DescriptionExtendedDescriptionField `json:"field_list"` // [Required] <p>Field of extended description.<br /></p>
}

type DescriptionExtendedDescriptionField struct {
	Type      string          `json:"type"`       // [Required] <p>Type of extended description field ：values: See Data Definition- description_field_type (text , image).<br /></p>
	Text      string          `json:"text"`       // [Required] <p>If field_type is text, text information will be returned through this field.<br /></p>
	ImageInfo *FieldImageInfo `json:"image_info"` // [Required] <p>If field_type is image, image url will be returned through this field.<br /></p>
}

type DescriptionInfo struct {
	ExtendedDescription *ExtendedDescription `json:"extended_description"` // [Required] If description_type is extended , description information should be set by this field.
}

type DescriptionInfoExtendedDescription struct {
	FieldList []ExtendedDescriptionField `json:"field_list"` // [Required] Field of extended description
}

type DescriptionLimit struct {
	DescriptionLengthMin           int64   `json:"description_length_min"`             // [Required] <p>Item description length min limit.<br /></p>
	DescriptionLengthMax           int64   `json:"description_length_max"`             // [Required] <p>length max limit for item extended description text part.<br /></p>
	DescriptionTextLengthMin       int64   `json:"description_text_length_min"`        // [Required] <p>length min limit for item extended description text part, when one of the minimum limits for image and text is reached, the item can be added or updated successfully.<br /></p>
	DescriptionTextLengthMax       int64   `json:"description_text_length_max"`        // [Required] length max limit for item extended description text part
	DescriptionImageNumMin         int64   `json:"description_image_num_min"`          // [Required] <p>length min limit for item extended description image num, when one of the minimum limits for image and text is reached, the item can be added or updated successfully.</p>
	DescriptionImageNumMax         int64   `json:"description_image_num_max"`          // [Required] <p>length max limit for item extended description image num.</p>
	DescriptionImageWidthMin       int64   `json:"description_image_width_min"`        // [Required] <p>length min limit for item extended description image width.</p>
	DescriptionImageHeightMin      int64   `json:"description_image_height_min"`       // [Required] <p>length min limit for item extended description image hight.</p>
	DescriptionImageAspectRatioMin float64 `json:"description_image_aspect_ratio_min"` // [Required] <p>length min limit for item extended description image aspect ( aspect_ratio= image width / image hight ).</p>
	DescriptionImageAspectRatioMax float64 `json:"description_image_aspect_ratio_max"` // [Required] <p>length max limit for item extended description image aspect ( aspect_ratio= image width / image hight ).</p>
}

type Dimension struct {
	PackageHeight int64 `json:"package_height"` // [Required] <p>The height of package for this global model, the unit is CM.<br /></p>
	PackageLength int64 `json:"package_length"` // [Required] <p>The length of package for this global model, the unit is CM.<br /></p>
	PackageWidth  int64 `json:"package_width"`  // [Required] <p>The width of package for this global model, the unit is CM.<br /></p>
}

type DimensionLimit struct {
	DimensionMandatory bool `json:"dimension_mandatory"` // [Required] <p>dimension is mandatory or not for the category<br /></p>
}

type DirectItem struct {
	DirectShopId int64 `json:"direct_shop_id"` // [Required] <p>Id of direct shop.</p>
	DirectItemId int64 `json:"direct_item_id"` // [Required] <p>Item id of direct shop.</p>
}

type DirectItemPrice struct {
	Region             string           `json:"region"`                // [Required] <p>Region of direct shop.</p>
	HiddenPrice        float64          `json:"hidden_price"`          // [Required]
	ItemModelPriceList []ItemModelPrice `json:"item_model_price_list"` // [Required]
}

type Discount struct {
	Status       string `json:"status"`        // [Required] The status of discount.
	DiscountName string `json:"discount_name"` // [Required] Title of the discount.
	StartTime    int64  `json:"start_time"`    // [Required] The time when discount activity start.
	EndTime      int64  `json:"end_time"`      // [Required] The time when discount activity end.
	DiscountId   int64  `json:"discount_id"`   // [Required] Shopee's unique identifier for a discount activity.
	Source       int64  `json:"source"`        // [Required] Source of the discount. 7: live stream, 1: admin, 0: others
}

type DiscoveryAdsLocations struct {
	Location string  `json:"location"`  // [Required] <p>daily_discover, you_may_also_like<br /></p>
	BidPrice float64 `json:"bid_price"` // [Required] <p>bid price of the location</p>
}

type DisputeReason struct {
	DisputeReason      int64            `json:"dispute_reason"`       // [Required] <p>The dispute_reason for one specific case. See Data Definition - DisputeReason.<br /></p>
	DisputeRequirement string           `json:"dispute_requirement"`  // [Required] <p>Indicate the importance of uploading required proof.<br /></p>
	SampleEvidence     []SampleEvidence `json:"sample_evidence"`      // [Required] <p>The URL of sample evidence to upload.<br /></p>
	EvidenceModuleList []EvidenceModule `json:"evidence_module_list"` // [Required] <p>The associated evidence module list for current dispute reason.</p>
}

type DisputeRequest struct {
	ReturnSn          string                `json:"return_sn"`                     // [Required] The serial number of return.
	Email             string                `json:"email"`                         // [Required] <p>The email address.</p>
	DisputeReasonId   int64                 `json:"dispute_reason_id"`             // [Required] <p>The dispute reason id.<br /></p><p>Please call&nbsp;v2.returns.get_return_dispute_reason to get it.</p>
	ImageList         []DisputeRequestImage `json:"image_list,omitempty"`          // [Optional] <p>Determines whether image submission is mandatory for the dispute request - mandatory input field for all dispute reasons except "Did not receive the return product".</p>
	DisputeTextReason *string               `json:"dispute_text_reason,omitempty"` // [Optional] <p>The content of dispute reason.</p>
}

type DisputeRequestImage struct {
	ModuleIndex int64    `json:"module_index"` // [Required] <p>The module_index of current evidence module returned by get_return_dispute_reason API.</p>
	Requirement string   `json:"requirement"`  // [Required] <p>The requirement content of current evidence module returned by get_return_dispute_reason API.</p>
	ImageUrl    []string `json:"image_url"`    // [Required] <p>The image URLs of current evidence module. It is recommended to pass in the URL returned by convert_image API.<br /></p>
}

type DisputeResponse struct {
	BaseResponse `json:",inline"`    // Common response fields
	Response     DisputeResponseData `json:"response"` // Actual response data
}

type DisputeResponseData struct {
	ReturnSn string `json:"return_sn"` // [Required] <p>The serial number of return.</p>
	Msg      string `json:"msg"`       // [Required]
}

type DownloadBookingShippingDocumentRequest struct {
	ShippingDocumentType *string   `json:"shipping_document_type,omitempty"` // [Optional] <p>The type of shipping document. Available values: NORMAL_AIR_WAYBILL,THERMAL_AIR_WAYBILL<br /></p>
	BookingList          []Booking `json:"booking_list"`                     // [Required] <p>The list of bookings you want to get. limit [1,50]<br /></p>
}

type DownloadBookingShippingDocumentResponse struct {
	BaseResponse `json:",inline"` // Common response fields
	Waybill      interface{}      `json:"waybill,omitempty"` // <p>The waybill file.<br /></p>
}

type DownloadFbsInvoicesRequest struct {
	RequestIdList *RequestIdList `json:"request_id_list,omitempty"` // [Optional] <p>list of request id (task identifiers)</p>
}

type DownloadFbsInvoicesResponse struct {
	BaseResponse `json:",inline"`                // Common response fields
	Response     DownloadFbsInvoicesResponseData `json:"response"`            // Actual response data
	ErrorMsg     string                          `json:"error_msg,omitempty"` //
	Timestamp    int64                           `json:"timestamp,omitempty"` //
}

type DownloadFbsInvoicesResponseData struct {
	RequestId int64  `json:"request_id"` // [Required]
	FileLink  string `json:"file_link"`  // [Required]
}

type DownloadInvoiceDocRequest struct {
	OrderSn string `json:"order_sn" url:"order_sn"` // [Required] Shopee's unique identifier for an order.
}

type DownloadInvoiceDocResponse struct {
	BaseResponse `json:",inline"` // Common response fields
	InvoiceDoc   interface{}      `json:"invoice_doc,omitempty"` //
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

type DriverInfo struct {
	DriverName   string `json:"driver_name"`    // [Required] <p>Driver Name</p>
	DriverPhone  string `json:"driver_phone"`   // [Required] <p>Driver phone number</p>
	VehicleType  string `json:"vehicle_type"`   // [Required] <p>Delivery vehicle type</p>
	LicensePlate string `json:"license_plate"`  // [Required] <p>License plate number</p>
	CourierPhoto string `json:"courier_photo"`  // [Required] <p>URL of the driver's photo</p>
	EtaStartTime int64  `json:"eta_start_time"` // [Required] <p>Earliest estimated arrival time at pickup address</p>
	EtaEndTime   int64  `json:"eta_end_time"`   // [Required] <p>Latest estimated arrival time at pickup address</p>
	DriverStatus string `json:"driver_status"`  // [Required] <p>Driver status. Applicable values:</p><p>- Allocating Driver<br />- Driver assigned<br />- Driver is on the way<br />- Driver is arrived</p>
}

type Dropoff struct {
	BranchId       *int64  `json:"branch_id,omitempty"`        // [Optional] The identity of branch.
	SenderRealName *string `json:"sender_real_name,omitempty"` // [Optional] The real name of sender.
	TrackingNumber *string `json:"tracking_number,omitempty"`  // [Optional] Need input this field when "tracking_number" is returned from "info_need". Please note that this tracking number is assigned by third-party shipping carrier for item shipment.
	Slug           *string `json:"slug,omitempty"`             // [Optional]  The selected 3PL partner to drop-off parcels with.
}

type DtsLimit struct {
	DaysToShipLimit       *WholesalePriceThresholdPercentage `json:"days_to_ship_limit"`         // [Required] <p>Pre order limits for the category</p>
	NonPreOrderDaysToShip int64                              `json:"non_pre_order_days_to_ship"` // [Required]
}

type EditGmsItemProductCampaignRequest struct {
	CampaignId *int64  `json:"campaign_id,omitempty"` // [Optional] <p>The GMS Campaign ID. Provide if available.<br /></p>
	EditAction string  `json:"edit_action"`           // [Required] <p>The following is the list of possible actions:</p><p>&nbsp;&nbsp;&nbsp;&nbsp;add</p><p>&nbsp;&nbsp;&nbsp;&nbsp;remove</p>
	ItemIdList []int64 `json:"item_id_list"`          // [Required] <p>Item IDs to add / remove. Minimum 1 Item ID. Maximum 30 Item IDs.<br /></p>
}

type EditGmsItemProductCampaignResponse struct {
	BaseResponse `json:",inline"`                       // Common response fields
	Response     EditGmsItemProductCampaignResponseData `json:"response"` // Actual response data
}

type EditGmsItemProductCampaignResponseData struct {
	CampaignId int64 `json:"campaign_id"` // [Required] <p>GMS Campaign ID<br /></p>
}

type EditGmsProductCampaignRequest struct {
	CampaignId  *int64   `json:"campaign_id,omitempty"`  // [Optional] <p>The GMS Campaign ID. Provide if available.<br /></p>
	EditAction  string   `json:"edit_action"`            // [Required] <p>The following is the list of possible actions and their required fields:</p><p>1.change_budgetdaily_budget</p><p>2.change_durationstart_date</p><p>3.end_date</p><p>4.pause</p><p>5.resume</p><p>6.start</p><p>7.change_roas_targetroas_target: when edit_action = change_roas_target, you must provide:</p><p>roas_target (float) Value rules follow the same logic as in the create endpoint</p>
	DailyBudget *float64 `json:"daily_budget,omitempty"` // [Optional] <p>Daily budget for Campaign.<br /></p>
	StartDate   *string  `json:"start_date,omitempty"`   // [Optional] <p>Start date of Campaign e.g. "11-11-2025". Cannot be earlier than today.<br /></p>
	EndDate     *string  `json:"end_date,omitempty"`     // [Optional] <p>End date of Campaign e.g. "11-11-2025". Do not fill if no end date.&nbsp;<br /></p>
	RoasTarget  *float64 `json:"roas_target,omitempty"`  // [Optional] <p>No input will be GMV Max Auto Bidding (Shop).</p><p>Input 0 for GMV Max Auto Bidding (Shop).</p><p>Input greater than 0 for GMV Max Custom ROAS (Shop).</p><p>If value = 10.123456, it will be taken as 10.1</p><p>If value = 10.199999, it will be taken as 10.1</p>
	ReferenceId *string  `json:"reference_id,omitempty"` // [Optional] <p>Generated by developers, used to prevent duplicate requests</p><p>Submitting the same reference_id more than once will fail; a new reference_id must be generated to retry.</p><p></p><p>Example: 086a16bf-49e9-4103-b7fe-c0125beb9278</p>
}

type EditGmsProductCampaignResponse struct {
	BaseResponse `json:",inline"`                   // Common response fields
	Response     EditGmsProductCampaignResponseData `json:"response"` // Actual response data
}

type EditGmsProductCampaignResponseData struct {
	CampaignId int64 `json:"campaign_id"` // [Required] <p>GMS Campaign ID<br /></p>
}

type EditManualProductAdKeywordsRequest struct {
	ReferenceId      string                    `json:"reference_id"`      // [Required] <p>A random string used to prevent duplicate ads. If an ads is created successfully, subsequent request using the same reference id will fail<br /></p>
	CampaignId       int64                     `json:"campaign_id"`       // [Required] <p>The unique identifier for a campaign<br /></p>
	SelectedKeywords []RequestSelectedKeywords `json:"selected_keywords"` // [Required] <p>selected keywords, required for manual bidding mode.</p>
}

type EditManualProductAdKeywordsResponse struct {
	BaseResponse `json:",inline"`                        // Common response fields
	Response     EditManualProductAdKeywordsResponseData `json:"response"` // Actual response data
}

type EditManualProductAdKeywordsResponseData struct {
	CampaignId  int64         `json:"campaign_id"`  // [Required] <p>The unique identifier for a campaign<br /></p>
	FailedEdits []FailedEdits `json:"failed_edits"` // [Required] <p>failed edits are mentioned here</p>
}

type EditManualProductAdsRequest struct {
	ReferenceId           string                         `json:"reference_id"`                      // [Required] <p>A random string used to prevent duplicate ads. If an ads is created successfully, subsequent request using the same reference id will fail<br /></p>
	CampaignId            int64                          `json:"campaign_id"`                       // [Required] <p>The unique identifier for a campaign<br /></p>
	EditAction            string                         `json:"edit_action"`                       // [Required] <p>Actions supported:&nbsp;<br />"start", "pause", "resume", "stop", "delete", "change_budget", "change_duration", "change_smart_creative", "change_location", "change_enhanced_cpc", "change_roas_target"</p>
	Budget                *float64                       `json:"budget,omitempty"`                  // [Optional] <p>The budget set for the Auto Product Ads<br /></p>
	StartDate             *string                        `json:"start_date,omitempty"`              // [Optional] <p>the start date per campaign. please kindly note that if you want to set unlimited date, you can just pass today's date as the start date</p>
	EndDate               *string                        `json:"end_date,omitempty"`                // [Optional] <p>the end date of each campaign. please kindly note that if you want to set an unlimited campaign, you can keep empty for the end date field</p>
	RoasTarget            *float64                       `json:"roas_target,omitempty"`             // [Optional] <p>the ROAS target for each campaign with auto bidding</p>
	DiscoveryAdsLocations []RequestDiscoveryAdsLocations `json:"discovery_ads_locations,omitempty"` // [Optional] <p>the location settings for manual bidding method</p>
	EnhancedCpc           *bool                          `json:"enhanced_cpc,omitempty"`            // [Optional] <p>Enhanced CPC functionality toggle</p>
	SmartCreativeSetting  *string                        `json:"smart_creative_setting,omitempty"`  // [Optional] <p>Whether to use default or set on/off Smart Creative for this ad. Supported Values: "", "default", "on", "off". Empty string treated as default.<br /></p>
}

type EditManualProductAdsResponse struct {
	BaseResponse `json:",inline"`                 // Common response fields
	Response     EditManualProductAdsResponseData `json:"response"` // Actual response data
}

type EditManualProductAdsResponseData struct {
	CampaignId int64 `json:"campaign_id"` // [Required] <p>The unique identifier for a campaign<br /></p>
}

type EndAddOnDealRequest struct {
	AddOnDealId int64 `json:"add_on_deal_id"` // [Required] The identifier of the API request for error tracking
}

type EndAddOnDealResponse struct {
	BaseResponse `json:",inline"`         // Common response fields
	Response     EndAddOnDealResponseData `json:"response"` // Actual response data
}

type EndAddOnDealResponseData struct {
	AddOnDealId int64 `json:"add_on_deal_id"` // [Required] The identifier of the API request for error tracking
}

type EndBundleDealRequest struct {
	BundleDealId int64 `json:"bundle_deal_id"` // [Required] Shopee's unique identifier for a bundle deal activity.
}

type EndBundleDealResponse struct {
	BaseResponse `json:",inline"`          // Common response fields
	Response     EndBundleDealResponseData `json:"response"` // Actual response data
}

type EndBundleDealResponseData struct {
	BundleDealId int64 `json:"bundle_deal_id"` // [Required] Shopee's unique identifier for a bundle deal activity.
}

type EndDiscountRequest struct {
	DiscountId int64 `json:"discount_id"` // [Required] Shopee's unique identifier for a discount activity.
}

type EndDiscountResponse struct {
	BaseResponse `json:",inline"`        // Common response fields
	Response     EndDiscountResponseData `json:"response"` // Actual response data
}

type EndDiscountResponseData struct {
	DiscountId int64 `json:"discount_id"` // [Required] Shopee's unique identifier for a discount activity.
	ModifyTime int64 `json:"modify_time"` // [Required] The time to track the modified time.
}

type EndFollowPrizeRequest struct {
	CampaignId int64 `json:"campaign_id"` // [Required] <p>The unique identifier for the created follow prize.<br /></p>
}

type EndFollowPrizeResponse struct {
	BaseResponse `json:",inline"`           // Common response fields
	Response     EndFollowPrizeResponseData `json:"response"` // Actual response data
}

type EndFollowPrizeResponseData struct {
	CampaignId int64 `json:"campaign_id"` // [Required] <p>The unique identifier for the created follow prize.<br /></p>
}

type EndQty struct {
	EndOnHandTotal int64 `json:"end_on_hand_total"` // [Required] <p>Total inventory at the end time.<br /></p>
	EndSellable    int64 `json:"end_sellable"`      // [Required]
	EndReserved    int64 `json:"end_reserved"`      // [Required]
	EndUnsellable  int64 `json:"end_unsellable"`    // [Required]
}

type EndSessionRequest struct {
	SessionId int64 `json:"session_id"` // [Required] <p>The identifier of livestream session.</p>
}

type EndSessionResponse struct {
	BaseResponse `json:",inline"` // Common response fields
	Response     interface{}      `json:"response"` // Actual response data
}

type EndVoucherRequest struct {
	VoucherId int64 `json:"voucher_id"` // [Required] The unique identifier for the voucher you want to end now.
}

type EndVoucherResponse struct {
	BaseResponse `json:",inline"`       // Common response fields
	Response     EndVoucherResponseData `json:"response"` // Actual response data
}

type EndVoucherResponseData struct {
	VoucherId int64 `json:"voucher_id"` // [Required] The unique identifier for the voucher it is being ended.
}

type EnterpriseInfo struct {
	CompanyName             string `json:"company_name"`              // [Required]
	Cnpj                    string `json:"cnpj"`                      // [Required]
	StateRegistrationNumber string `json:"state_registration_number"` // [Required]
	IsFreightPayer          bool   `json:"is_freight_payer"`          // [Required]
}

type Error struct {
	ItemId      int64  `json:"item_id"`      // [Required] Shopee's unique identifier for an item.
	ModelId     int64  `json:"model_id"`     // [Required] Shopee's unique identifier for a model.
	FailError   string `json:"fail_error"`   // [Required]
	FailMessage string `json:"fail_message"` // [Required]
}

type Escrow struct {
	EscrowAmount float64 `json:"escrow_amount"` // [Required] The total amount that the seller is expected to receive for the order.
	Currency     string  `json:"currency"`      // [Required] The currency used for calculating escrow amount.
	OrderSn      string  `json:"order_sn"`      // [Required] Shopee's unique identifier for an order.
}

type EscrowDetail struct {
	OrderSn           string                        `json:"order_sn"`             // [Required] <p>Shopee's unique identifier for an order.</p><p>&lt;path&gt;&lt;/path&gt;<br /></p>
	BuyerUserName     string                        `json:"buyer_user_name"`      // [Required] <p>The username of buyer.</p><p>&lt;path&gt;&lt;/path&gt;<br /></p>
	ReturnOrderSnList []string                      `json:"return_order_sn_list"` // [Required] <p>The list of the serial number of return.</p><p>&lt;path&gt;&lt;/path&gt;<br /></p>
	OrderIncome       *EscrowDetailOrderIncome      `json:"order_income"`         // [Required]
	BuyerPaymentInfo  *EscrowDetailBuyerPaymentInfo `json:"buyer_payment_info"`   // [Required] <p>The buyer payment info at order checkout moment. (snapshot value)</p>
}

type EscrowDetailBuyerPaymentInfo struct {
	BuyerPaymentMethod    string  `json:"buyer_payment_method"`     // [Required] <p>The payment method used by buyer.</p>
	BuyerServiceFee       string  `json:"buyer_service_fee"`        // [Required] <p>An additional service fee charged by shopee to Buyer at checkout. currently only applicable to ID region.<br />it is an initial value (will not be updated after RR/cancellation)<br /></p>
	BuyerTaxAmount        float64 `json:"buyer_tax_amount"`         // [Required] <p>The tax amount paid by buyer.</p><p>currently this is a custom tax charged to CB orders in TW,CL,MX</p>
	BuyerTotalAmount      float64 `json:"buyer_total_amount"`       // [Required] <p>The total amount paid by buyer at checkout moment.<br /></p>
	CreditCardPromotion   float64 `json:"credit_card_promotion"`    // [Required] <p>The promotion provided by credit card.<br />it is an initial value (will not be updated after RR/cancellation)<br /></p>
	IcmsTaxAmount         float64 `json:"icms_tax_amount"`          // [Required] <p>The icms tax paid by buyer. this is only applicable to BR region.<br />it is an initial value (will not be updated after RR/cancellation)<br /></p>
	ImportTaxAmount       float64 `json:"import_tax_amount"`        // [Required] <p>The import tax paid by buyer.&nbsp;<br />it is an initial value (will not be updated after RR/cancellation)<br /></p>
	InitialBuyerTxnFee    float64 `json:"initial_buyer_txn_fee"`    // [Required] <p>Transaction fee paid by buyer for the order. (Only display for non cb sip affiliate shop. ).&nbsp;Most regions would have this fee charged to buyer at checkout depending on the fee rules applied in each region.<br />it is an initial value (will not be updated after RR/cancellation)<br /></p>
	InsurancePremium      float64 `json:"insurance_premium"`        // [Required] <p>The insurance premium paid by buyer. Currently only applicable to some regions like ID &amp; BR</p><p>it is an initial value (will not be updated after RR/cancellation)</p>
	IofTaxAmount          float64 `json:"iof_tax_amount"`           // [Required] <p>The iof tax paid by buyer. this is only applicable for BR region.</p><p>it is an initial value (will not be updated after RR/cancellation)</p>
	IsPaidByCreditCard    bool    `json:"is_paid_by_credit_card"`   // [Required] <p>Whether this order is paid by credit card. it's related to payment channel used at checkout<br /></p>
	MerchantSubtotal      float64 `json:"merchant_subtotal"`        // [Required] <p>The total item price paid by buyer at checkout.</p><p>it is an initial value and will not be updated after RR/cancellation</p>
	SellerVoucher         float64 `json:"seller_voucher"`           // [Required] <p>The voucher provided by seller to offset some value needs to be paid by buyer.</p><p>it is an initial value (will not be updated after RR/cancellation)</p>
	ShippingFee           float64 `json:"shipping_fee"`             // [Required] <p>The shipping fee paid by buyer. (Only display for non cb sip affiliate shop.&nbsp;<br /></p><p>it is an initial value (will not be updated after RR/cancellation)</p>
	ShippingFeeSstAmount  float64 `json:"shipping_fee_sst_amount"`  // [Required] <p>The shipping fee sst paid by buyer. Currently apply to MY region only&nbsp;</p><p>it is an initial value (will not be updated after RR/cancellation)</p>
	ShopeeVoucher         float64 `json:"shopee_voucher"`           // [Required] <p>The voucher provided by Shopee to offset the amount need to be paid by buyer.</p><p>it is an initial value (will not be updated after RR/cancellation)</p>
	ShopeeCoinsRedeemed   float64 `json:"shopee_coins_redeemed"`    // [Required] <p>This is an amount of coin used by buyer at checkout to offset some amount to be paid.&nbsp;</p><p>it is an initial value (will not be updated after RR/cancellation)</p>
	BuyerPaidPackagingFee float64 `json:"buyer_paid_packaging_fee"` // [Required] <p>The fee charged to the buyer for packaging materials.</p>
	TradeInBonus          float64 `json:"trade_in_bonus"`           // [Required] <p>The total amount of all trade-in bonuses applied to a transaction. This value is the summation of the bonuses contributed by all four parties: Shopee, Seller, Bank and Trade-in Partner.</p>
	BulkyHandlingFee      float64 `json:"bulky_handling_fee"`       // [Required] <p>A fee charged to the buyer for the special handling and transportation required for items that exceed a specified weight or dimension threshold.Only for local ID sellers.</p>
	DiscountPix           float64 `json:"discount_pix"`             // [Required] <p>The discount provided by PIX&nbsp;(Only applicable for BR region)</p>
}

type EscrowDetailOrderIncome struct {
	EscrowAmount                                       float64                   `json:"escrow_amount"`                                             // [Required] <p>The total amount that the seller is expected to receive for the order and will change before order is completed.&nbsp;</p><p>For non cb sip affiliate shop (new formula):&nbsp;</p><p>escrow_amount=&nbsp;original_cost_of_goods_sold-original_shopee_discount+seller_return_refund+ shopee_discounts- voucher_from_seller- seller_coin_cash_back+ buyer_paid_shipping_fee- actual_shipping_fee+ shopee_shipping_rebate+ shipping_fee_discount_from_3pl- reverse_shipping_fee+ rsf_seller_protection_fee_claim_amount+ fsf_seller_protection_fee_claim_amount- final_return_to_seller_shipping_fee- seller_transaction_fee- service_fee- commission_fee- campaign_fee- shipping_seller_protection_fee_premium_amount- delivery_seller_protection_fee_premium_amount-final_escrow_product_gst- order_ams_commission fee- escrow_tax-sales_tax_on_lvg-reverse_shipping_fee_sst-shipping_fee_sst-withholding_tax-overseas_return_service_fee-vat_on_imported_goods - withholding_vat_tax - withholding_pit_tax - seller_order_processing_fee + buyer_paid_packaging_fee - trade_in_bonus_seller - fbs_fee -&nbsp;ads_escrow_top_up_fee_or_technical_support_fee</p><br /><p><br /></p><p>For cb sip affiliate shop:&nbsp;escrow_amount=escrow_amount_pri * exchange_rate</p><p><br /></p><p>note: Return refund amount = if adjustable RR, will equal to drc_adjustable_refund</p>
	OrderOriginalPrice                                 float64                   `json:"order_original_price"`                                      // [Required] <p>The original price of the item before ANY promotion/discount in the listing currency. It returns the subtotal of that specific item if quantity exceeds 1.<br /></p>
	OriginalPrice                                      float64                   `json:"original_price"`                                            // [Required] <p>The original price of the item before ANY promotion/discount in the listing currency. It returns the subtotal of that specific item if quantity exceeds 1.<br /></p>
	OrderSellingPrice                                  float64                   `json:"order_selling_price"`                                       // [Required] <p>This field value = sum of item unit price.selling price comes from the sum up of each item's unit price. For non-bundle deal case, this value will be same like order_original_price; For bundle deal case, this value will be price of sum of item price before bundle deal promo but after item promo.It returns the subtotal of that specific item if quantity exceeds 1. (Only display for non cb sip affiliate shop. )<br /></p>
	OrderSellerDiscount                                float64                   `json:"order_seller_discount"`                                     // [Required] <p>The total discount seller provided for this order. It returns the subtotal of that specific item if quantity exceeds 1. (Only display for non cb sip affiliate shop. )<br /></p>
	SellerDiscount                                     float64                   `json:"seller_discount"`                                           // [Required] <p>Final sum of each item seller discount of a specific order. (Only display for non cb sip affiliate shop. )<br /></p>
	OrderDiscountedPrice                               float64                   `json:"order_discounted_price"`                                    // [Required] <p>The total discounted price for this order. It returns the subtotal of that specific item if quantity exceeds 1. (Only display for non cb sip affiliate shop. )<br /></p>
	ShopeeDiscount                                     float64                   `json:"shopee_discount"`                                           // [Required] <p>Final sum of each item Shopee discount of a specific order. This amount will return remaining rebate value (i.e. remaining Shopee Item Rebate +&nbsp;remaining Shopee PIX Rebate) to seller. (Only display for non cb sip affiliate order. )<br /></p>
	VoucherFromSeller                                  float64                   `json:"voucher_from_seller"`                                       // [Required] <p>Final value of voucher provided by Seller for the order. (Only display for non cb sip affiliate shop. )<br /></p>
	VoucherFromShopee                                  float64                   `json:"voucher_from_shopee"`                                       // [Required] <p>Final value of voucher provided by Shopee for the order. (Only display for non cb sip affiliate shop. )<br /></p>
	Coins                                              float64                   `json:"coins"`                                                     // [Required] <p>This value indicates the total amount offset when the buyer consumed Shopee Coins upon checkout. (Only display for non cb sip affiliate shop. )<br /></p>
	CrossBorderTax                                     float64                   `json:"cross_border_tax"`                                          // [Required] <p>Amount incurred by Buyer for purchasing items outside of home country. Amount may change after Return Refund. (Only display for non cb sip affiliate shop. )<br /></p>
	PaymentPromotion                                   float64                   `json:"payment_promotion"`                                         // [Required] <p>The amount offset via payment promotion. (Only display for non cb sip affiliate shop. )<br /></p>
	CommissionFee                                      float64                   `json:"commission_fee"`                                            // [Required] <p>The commission fee charged by Shopee platform if applicable.</p><p>For CB SIP affiliate shop:&nbsp;commission_fee=commission_fee_pri * exchange_rate</p>
	ServiceFee                                         float64                   `json:"service_fee"`                                               // [Required] <p>Amount charged by Shopee to seller for additional services.</p><p>For CB SIP affiliate shop:&nbsp;service_fee=service_fee_pri * exchange_rate</p>
	SellerTransactionFee                               float64                   `json:"seller_transaction_fee"`                                    // [Required] <p>Tansaction fee paid by seller for the order. (Only display for non cb sip affiliate shop. )<br /></p>
	SellerLostCompensation                             float64                   `json:"seller_lost_compensation"`                                  // [Required] <p>Compensation to seller in case of lost parcel. (Only display for non cb sip affiliate shop. )<br /></p>
	SellerCoinCashBack                                 float64                   `json:"seller_coin_cash_back"`                                     // [Required] <p>Value of coins provided by Seller for purchasing with his or her store for the order. (Only display for non cb sip affiliate shop. )<br /></p>
	EscrowTax                                          float64                   `json:"escrow_tax"`                                                // [Required] <p>Cross-border tax imposed by the Indonesian government on sellers. (Only display for non cb sip affiliate shop. )<br /></p>
	FinalShippingFee                                   float64                   `json:"final_shipping_fee"`                                        // [Required] <p>Final adjusted amount that seller has to bear as part of escrow. This amount could be negative or positive. (Only display for non cb sip affiliate shop. )<br /></p>
	ActualShippingFee                                  float64                   `json:"actual_shipping_fee"`                                       // [Required] <p>The final shipping cost of order and it is positive. For Non-integrated logistics channel is 0. (Only display for non cb sip affiliate shop. )<br /></p>
	ShopeeShippingRebate                               float64                   `json:"shopee_shipping_rebate"`                                    // [Required] <p>The platform shipping subsidy to the seller. (Only display for non cb sip affiliate shop. )<br /></p>
	ShippingFeeSst                                     float64                   `json:"shipping_fee_sst"`                                          // [Required] <p>The Service Tax charged on Seller Paid Shipping Fee for forward shipping, based on Malaysia SST regulations&nbsp;for shipping fees on all orders. Definition of Seller Paid Shipping Fee is Actual Shipping Fee subtracted by sum of Shipping Fee Paid by Buyer and Shipping Rebate From Shopee. (Only applicable for non cb sip affiliate shop)</p>
	ShippingFeeDiscountFromPl                          float64                   `json:"shipping_fee_discount_from_pl"`                             // [Required] <p>The discount of shipping fee from 3PL. Currently only applicable to ID. (Only display for non cb sip affiliate shop. )<br /></p>
	SellerShippingDiscount                             float64                   `json:"seller_shipping_discount"`                                  // [Required] <p>The shipping discount defined by seller. (Only display for non cb sip affiliate shop. )<br /></p>
	EstimatedShippingFee                               float64                   `json:"estimated_shipping_fee"`                                    // [Required] <p>The estimated shipping fee is an estimation calculated by Shopee based on specific logistics courier's standard. (Only display for non cb sip affiliate shop. )<br /></p>
	SellerVoucherCode                                  float64                   `json:"seller_voucher_code"`                                       // [Required] <p>The list of voucher code provided by seller. (Only display for non cb sip affiliate shop. )<br /></p>
	DrcAdjustableRefund                                float64                   `json:"drc_adjustable_refund"`                                     // [Required] <p>The adjustable refund amount from Shopee Dispute Resolution Center.<br /></p>
	RefundAmountToBuyer                                float64                   `json:"refund_amount_to_buyer"`                                    // [Required] <p>Amount returned to Seller in the event of Partial Return.<br /></p>
	CostOfGoodsSold                                    float64                   `json:"cost_of_goods_sold"`                                        // [Required] <p>Final amount paid by the buyer for the items in a specific order. (Only display for non cb sip affiliate shop. )<br /></p>
	OriginalCostOfGoodsSold                            float64                   `json:"original_cost_of_goods_sold"`                               // [Required] <p>Amount paid by the buyer for the items in a specific order. (Only display for non cb sip affiliate shop. )<br /></p>
	OriginalShopeeDiscount                             float64                   `json:"original_shopee_discount"`                                  // [Required] <p>Sum of each item Shopee discount of a specific order. This amount will return initial rebate value (i.e. remaining Shopee Item Rebate +&nbsp;remaining Shopee PIX Rebate) to seller. (Only display for non cb sip affiliate order. )<br /></p>
	Items                                              []OrderIncomeItems        `json:"items"`                                                     // [Required] <p>This object contains the detailed breakdown for all the items in this order, including regular items(non-activity) and activity items.<br /></p>
	EscrowAmountPri                                    float64                   `json:"escrow_amount_pri"`                                         // [Required] <p>The total amount in the prim currency that the seller is expected to receive for the order and will change before order completed . escrow_amount_pri=original_price_pri-seller_return_refund_pri-commission_fee_pri-service_fee_pri-drc_adjustable_refund_pri.(Only display for cb sip affiliate order. )<br /></p>
	BuyerTotalAmountPri                                float64                   `json:"buyer_total_amount_pri"`                                    // [Required] <p>The total amount that paid by buyer in the primary currency. (Only display for cb sip affiliate order. )</p>
	OriginalPricePri                                   float64                   `json:"original_price_pri"`                                        // [Required] <p>The original price of the item before ANY promotion/discount in the primary currency. It returns the subtotal of that specific item if quantity exceeds 1.(Only display for cb sip affiliate order. )</p><p>&lt;path&gt;&lt;/path&gt;<br /></p>
	SellerReturnRefundPri                              float64                   `json:"seller_return_refund_pri"`                                  // [Required] <p>Amount returned to Seller in the event of Partial Return in the primary currency. (Only display for cb sip affiliate shop. )</p>
	CommissionFeePri                                   float64                   `json:"commission_fee_pri"`                                        // [Required] <p>The commission fee charged by Shopee platform if applicable in the primary currency. (Only display for cb sip affiliate shop. )</p>
	ServiceFeePri                                      float64                   `json:"service_fee_pri"`                                           // [Required] <p>Amount charged by Shopee to seller for additional services in the primary currency. (Only display for cb sip affiliate shop. )</p>
	DrcAdjustableRefundPri                             float64                   `json:"drc_adjustable_refund_pri"`                                 // [Required] <p>The adjustable refund amount from Shopee Dispute Resolution Center in the primary currency&nbsp;after proration. (Only applicable for cb sip affiliate shop.)</p>
	PriCurrency                                        string                    `json:"pri_currency"`                                              // [Required] <p>The currency of the country or region where the shop that real seller operates. (Only display for cb sip affiliate shop. )</p>
	AffCurrency                                        string                    `json:"aff_currency"`                                              // [Required] <p>The currency of the country or region where shop opened in. (Only display for cb sip affiliate shop. )</p>
	ExchangeRate                                       float64                   `json:"exchange_rate"`                                             // [Required] <p>Exchange rate from primary shop currency to affiliate shop currency.<br /></p>
	ReverseShippingFee                                 float64                   `json:"reverse_shipping_fee"`                                      // [Required] <p>Shopee charges the reverse shipping fee for the returned order.The value of this field will be non-negative.<br /></p>
	ReverseShippingFeeSst                              float64                   `json:"reverse_shipping_fee_sst"`                                  // [Required] <p>The Service Tax charged on Reverse Shipping Fee for reverse shipping, based on Malaysia SST regulations&nbsp;for shipping fees on all orders. (Only applicable for non cb sip affiliate shop)<br /></p>
	FinalProductProtection                             float64                   `json:"final_product_protection"`                                  // [Required] <p>The total amount of product protection purchased during placing an order.&nbsp;</p>
	CreditCardPromotion                                float64                   `json:"credit_card_promotion"`                                     // [Required] <p>This value indicate the offset via credit card promotion.<br /></p>
	CreditCardTransactionFee                           float64                   `json:"credit_card_transaction_fee"`                               // [Required] <p>This value indicate the total transaction fee.</p><p>credit_card_transaction_fee=buyer_transaction_fee+seller_transaction_fee</p>
	FinalProductVatTax                                 float64                   `json:"final_product_vat_tax"`                                     // [Required] <p>Value-added Tax is required for online purchases based on EU Value-added Tax regulations . (Only display for non cb sip affiliate shop. )<br /></p>
	FinalShippingVatTax                                float64                   `json:"final_shipping_vat_tax"`                                    // [Required] <p>Value-added Tax for product price is required for online purchases based on EU Value-added Tax regulations. (Only applicable for non cb sip affiliate shop. )<br /></p>
	CampaignFee                                        float64                   `json:"campaign_fee"`                                              // [Required] <p>The campaign fee charged by Shopee platform. Only available for some local Indonesian shops.<br /></p>
	SipSubsidy                                         float64                   `json:"sip_subsidy"`                                               // [Required] <p>The SIP subsidy amount is the difference between the item settlement price set by seller and item price actually paid by buyer. (Only available for CB SIP A Shops)<br /></p>
	SipSubsidyPri                                      float64                   `json:"sip_subsidy_pri"`                                           // [Required] <p>The SIP subsidy amount is the difference between the item settlement price set by seller and item price actually paid by buyer. This value is in the primary currency. (Only available for CB SIP A Shops)<br /></p>
	RsfSellerProtectionFeeClaimAmount                  float64                   `json:"rsf_seller_protection_fee_claim_amount"`                    // [Required] <p>The insurance claim amount if seller opt in to insurance program. this covers Reverse shipping Fee in the case of RR. As per Jun 2024:<br />- For ID &amp; MY Local: After Extension on coverage to FSF due to RR. this claim amount will consist of FSF + RSF claim.<br />- For PH local: This will only cover RSF claim<br /><br />will be updated if there's any RR/cancellation<br /></p>
	RsfSellerProtectionFeePremiumAmount                float64                   `json:"rsf_seller_protection_fee_premium_amount"`                  // [Required] <p>Amount charged by Shopee to seller for additional services. Only apply for PH local orders.<br /></p>
	FinalEscrowProductGst                              float64                   `json:"final_escrow_product_gst"`                                  // [Required] <p>Goods and Service Tax for product price is required for imported goods (overseas orders) based on Singapore GST regulations. (Only applicable for non cb sip affiliate shop selling in Singapore)<br /></p>
	FinalEscrowShippingGst                             float64                   `json:"final_escrow_shipping_gst"`                                 // [Required] <p>Goods and Service Tax for shipping fee is required for imported goods (overseas orders) based on Singapore GST regulations. (Only applicable for non cb sip affiliate shop selling in Singapore.)<br /></p>
	DeliverySellerProtectionFeePremiumAmount           float64                   `json:"delivery_seller_protection_fee_premium_amount"`             // [Required] <p>[Currently apply to ID &amp; local orders only] An insurance premium charged to seller at the time parcel is picked up by 3PL for insurance in case of parcel lost/damaged by 3PL.<br /></p>
	OrderAmsCommissionFee                              float64                   `json:"order_ams_commission_fee"`                                  // [Required] <p>The amount of affiliate commission fee for this order. Applicable for orders sold via the Affiliate Program.<br /></p>
	BuyerPaymentMethod                                 float64                   `json:"buyer_payment_method"`                                      // [Required] <p>The payment method buyer used when do the order checkout.<br /></p>
	InstalmentPlan                                     float64                   `json:"instalment_plan"`                                           // [Required] <p>The instalment plan buyer chosen when do the order checkout. Only applicable when payment method support instalment.<br /></p>
	SalesTaxOnLvg                                      float64                   `json:"sales_tax_on_lvg"`                                          // [Required] <p>Sales Tax on Low Value Goods (LVG) is required for imported goods (overseas orders) based on Malaysia SST regulations&nbsp;for selective orders (e.g. CB LVG orders in MY)<br /></p>
	WithholdingTax                                     float64                   `json:"withholding_tax"`                                           // [Required] <p>Only for PH and ID local shops.<br /><br /></p><p><b>PH:&nbsp;</b>According to regulations issued by Bureau of Internal Revenue in PH, the Withholding Tax is applied to the gross remittances sent by Shopee to online suppliers of goods and services.<br /></p><p><br /></p><p><b>ID:&nbsp;</b>According to regulations issued by Directorate General of Taxation in ID, the Withholding Tax is applied to the income stated in the invoice generated via Shopee related to Seller and/or Merchants' sales in Shopee's platform.</p>
	OverseasReturnServiceFee                           float64                   `json:"overseas_return_service_fee"`                               // [Required] <p>This is overseas return service fee charged to sellers who register ORS program.(Only applicable for non cb sip affiliate shop)<br /></p>
	ProratedCoinsValueOffsetReturnItems                float64                   `json:"prorated_coins_value_offset_return_items"`                  // [Required] <p>This is the prorated value from cash equivalent of coin offset due to adjustable RR.This field will only be updated when there is an adjustable RR. If it's a full RR or normal order will response 0.<br /></p>
	ProratedShopeeVoucherOffsetReturnItems             float64                   `json:"prorated_shopee_voucher_offset_return_items"`               // [Required] <p>This is the prorated refund value from shopee voucher discount due to adjustable RR.This field will only be updated when there is an adjustable RR.&nbsp;If it's a full RR or normal order will response 0.<br /></p>
	ProratedSellerVoucherOffsetReturnItems             float64                   `json:"prorated_seller_voucher_offset_return_items"`               // [Required] <p>This is the prorated refund value from seller voucher discount due to adjustable RR.This field will only be updated when there is an adjustable RR.&nbsp;If it's a full RR or normal order will response 0.<br /></p>
	ProratedPaymentChannelPromoBankOffsetReturnItems   float64                   `json:"prorated_payment_channel_promo_bank_offset_return_items"`   // [Required] <p>This is the prorated value from bank payment channel promo due to adjustable RR.This field will only be updated when there is an adjustable RR.If it's a full RR or normal order will response 0.<br /></p>
	ProratedPaymentChannelPromoShopeeOffsetReturnItems float64                   `json:"prorated_payment_channel_promo_shopee_offset_return_items"` // [Required] <p>This is the prorated value from shopee payment channel promo due to adjustable RR.This field will only be updated when there is an adjustable RR.If it's a full RR or normal order will response 0.<br /></p>
	FsfSellerProtectionFeeClaimAmount                  float64                   `json:"fsf_seller_protection_fee_claim_amount"`                    // [Required] <p>The claim amount given to seller if seller opt in to shipping fee service program. this covers Forward Shipping Fee cost in the case of cancelled due to Failed delivery. only apply to PH Local as per Jun 2024.<br /></p>
	ShippingSellerProtectionFeeAmount                  float64                   `json:"shipping_seller_protection_fee_amount"`                     // [Required] <p>Service fee charged to seller in MY,ID,PH Local (as per Jun 2024) due to additional program purchased<br /></p>
	FinalReturnToSellerShippingFee                     float64                   `json:"final_return_to_seller_shipping_fee"`                       // [Required] <p>The amount of fee charged to seller (if any) for the failed delivery order&nbsp;</p>
	VatOnImportedGoods                                 float64                   `json:"vat_on_imported_goods"`                                     // [Required] <p>7% VAT is charged for imported goods entering Thailand<br /></p>
	WithholdingVatTax                                  float64                   `json:"withholding_vat_tax"`                                       // [Required] <p>By VN law, E-commerce platforms need to Withhold VAT tax applicable to all VN business household and VN individual sellers</p>
	WithholdingPitTax                                  float64                   `json:"withholding_pit_tax"`                                       // [Required] <p>By VN law, E-commerce platforms need to Withhold Personal Income Tax applicable to all VN business household and VN individual sellers</p>
	TaxRegistrationCode                                string                    `json:"tax_registration_code"`                                     // [Required] <p>For VN Withholding Tax. This is the Tax Registration Number (TRN) from Seller Info (Business information) of the seller at the point of order creation</p>
	SellerOrderProcessingFee                           float64                   `json:"seller_order_processing_fee"`                               // [Required] <p>Order Processing Fee is the amount charged to sellers for every order created.</p>
	BuyerPaidPackagingFee                              float64                   `json:"buyer_paid_packaging_fee"`                                  // [Required] <p>The fee charged to the buyer for packaging materials.</p>
	TradeInBonusSeller                                 float64                   `json:"trade_in_bonus_seller"`                                     // [Required] <p>The discount provided by Seller/Retailers for buyers who opt for trade-in.</p>
	FbsFee                                             float64                   `json:"fbs_fee"`                                                   // [Required] <p>Fulfilled by Shopee (FBS) Fee applied to this order, covering costs such as handling and storage and packaging. Only for PH Local Orders.</p>
	NetCommissionFee                                   float64                   `json:"net_commission_fee"`                                        // [Required] <p>The respective fee amounts after the proportional rebate deduction.The total net commission fee applied to the order, calculated as the sum of all net commission fee items.</p><p>-only for BR local sellers.</p>
	NetServiceFee                                      float64                   `json:"net_service_fee"`                                           // [Required] <p>The respective fee amounts after the proportional rebate deduction.The total net service fee applied to the order, calculated as the sum of all net service fee items.</p><p>-only for BR local sellers.</p>
	NetCommissionFeeInfoList                           *NetCommissionFeeInfoList `json:"net_commission_fee_info_list"`                              // [Required] <p>Returns a breakdown of the net commission fees.</p><p>-only for BR local sellers.</p>
	NetServiceFeeInfoList                              *NetServiceFeeInfoList    `json:"net_service_fee_info_list"`                                 // [Required] <p>Returns a breakdown of the net service fees.</p><p>-only for BR local sellers.</p>
	SellerProductRebate                                *SellerProductRebate      `json:"seller_product_rebate"`                                     // [Required] <p>The shopee rebate borne by seller.</p><p>-only for BR local sellers.</p>
	PixDiscount                                        float64                   `json:"pix_discount"`                                              // [Required] <p>[BR only]Final sum of pix discount of a specific order. (Only display for non cb sip affiliate shop.)</p>
	ProratedPixDiscountOffsetReturnItems               float64                   `json:"prorated_pix_discount_offset_return_items"`                 // [Required] <p>[BR only]This is the prorated value from pix discount due to adjustable RR. This field will only be updated when there is an adjustable RR. If it's a full RR or normal order, will response 0.</p>
	AdsEscrowTopUpFeeOrTechnicalSupportFee             float64                   `json:"ads_escrow_top_up_fee_or_technical_support_fee"`            // [Required] <p>Includes both ads escrow top up fee (auto escrow top up to your ads balance) and technical support fee (charged by Shopee)</p><p>The actual fee type included in this field varies depending on the seller type and selling region, and may represent one of the following in Shopee Seller Center:</p><p>Ads Escrow Top-Up Fee</p><p>For local MY TH SG VN PH ID sellers and CNCB / JPCB / KRCB sellers selling in PH and ID</p><p>For JPCB sellers selling in SG, MY, TH, and VN</p><p>Technical Support Fee</p><p>For CNCB sellers selling in SG, MY, TH, and VN</p><p>Traffic Growth Fee</p><p>For KRCB sellers selling in SG, MY, TH, and VN</p>
}

type EvidenceModule struct {
	ModuleIndex int64  `json:"module_index"` // [Required] <p>The index of current evidence module.</p>
	Requirement string `json:"requirement"`  // [Required] <p>The specific requirement of current evidence module.</p>
	IsRequired  bool   `json:"is_required"`  // [Required] <p>Indicate if current evidence module is mandatory or not.<br /></p>
}

type ExtendedDescription struct {
	FieldList []Field `json:"field_list"` // [Required]  Field of extended description.
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

type Fail struct {
	PackageNumber string `json:"package_number"` // [Required] <p>Shopee's unique identifier for the package under an order.</p>
	FailReason    string `json:"fail_reason"`    // [Required] <p>Reason for failure.</p>
}

type Failed struct {
	FailedReason string `json:"failed_reason"` // [Required] Failed reason.
}

type FailedEdits struct {
	Keyword string `json:"keyword"` // [Required] <p>keyword that failed to update</p>
	Error   string `json:"error"`   // [Required] <p>Error code</p>
	Message string `json:"message"` // [Required] <p>error description</p>
}

type FailedItems struct {
	ItemId                int64                   `json:"item_id"`                // [Required]
	ModelId               int64                   `json:"model_id"`               // [Required] <p>If the item has no variation, this field will be empty</p>
	ErrCode               int64                   `json:"err_code"`               // [Required]
	ErrMsg                string                  `json:"err_msg"`                // [Required] <p>the reason why the model cannot be added</p>
	UnqualifiedConditions []UnqualifiedConditions `json:"unqualified_conditions"` // [Required] <p>if the model doesn't meet a criteria, will show the detail in this field<br /></p>
}

type Failure struct {
	ItemId       int64  `json:"item_id"`       // [Required] <p>Shopee's unique identifier for an item.<br /></p>
	FailedReason string `json:"failed_reason"` // [Required] <p>Item's failure reason.</p>
}

type FhrOrder struct {
	OrderSn             string   `json:"order_sn"`               // [Required] <p>Order SN.</p>
	ParcelId            int64    `json:"parcel_id"`              // [Required] <p>Parcel ID.</p>
	ParcelDisplayId     string   `json:"parcel_display_id"`      // [Required] <p>Display Parcel ID.</p>
	ConfirmTime         int64    `json:"confirm_time"`           // [Required] <p>Confirmed Date.</p>
	HandoverDeadline    int64    `json:"handover_deadline"`      // [Required] <p>Handover Deadline.<br /></p>
	FastHandoverDueDate int64    `json:"fast_handover_due_date"` // [Required] <p>Fast Handover Due Date.</p>
	ArrangePickUpTime   int64    `json:"arrange_pick_up_time"`   // [Required] <p>Seller arrange pick up time.</p>
	HandoverTime        int64    `json:"handover_time"`          // [Required] <p>Parcel drop off / pickup time.<br /></p>
	ShippingChannel     string   `json:"shipping_channel"`       // [Required] <p>Logistics Company.</p>
	FirstMileType       string   `json:"first_mile_type"`        // [Required] <p>First mile shipping type. Applicable values:</p><p>Pickup</p><p>Drop off</p>
	FirstMileTrackingNo string   `json:"first_mile_tracking_no"` // [Required] <p>First Mile Tracking No.</p>
	DiagnosisScenario   []string `json:"diagnosis_scenario"`     // [Required] <p>Diagnosis of the issue.</p>
}

type Field struct {
	FieldType DescriptionElementFieldType `json:"field_type"` // [Required] Type of extended description field ：values: See Data Definition- description_field_type (text , image).
	Text      string                      `json:"text"`       // [Required] If field_type is text， text information will be set by this field.
	ImageInfo *SharedImageInfo            `json:"image_info"` // [Required] If field_type is image，image url will be set by this field.
}

type FieldImageInfo struct {
	ImageId  string `json:"image_id"`  // [Required] <p>Unique ID of the uploaded image.</p>
	ImageUrl string `json:"image_url"` // [Required] <p>URL of the uploaded image.</p>
}

type Filter struct {
	PackageStatus       *int64       `json:"package_status,omitempty"`        // [Optional] <p>Use this field to filter the packages of specific status. Applicable values:</p><p>0: All</p><p>1: Pending</p><p>2: ToProcess</p><p>3: Processed</p><p><br /></p><p><font color="#c24f4a">Default value = 2 (ToProcess)</font></p>
	ProductLocationIds  []string     `json:"product_location_ids,omitempty"`  // [Optional] <p>List of product_location_id. Use this field to filter the packages under specific warehouses.</p>
	LogisticsChannelIds *interface{} `json:"logistics_channel_ids,omitempty"` // [Optional] <p>List of logistics_channel_id. Use this field to filter the packages under specific logistics channels.</p>
	FulfillmentType     *int64       `json:"fulfillment_type,omitempty"`      // [Optional] <p>Use this field to filter the packages fulfilled by shopee or seller. Applicable values:&nbsp;</p><p>0: None (not apply filter)</p><p>1: Shopee</p><p>2: Seller</p><p><br /></p><p><font color="#c24f4a">Default value = 2 (Seller)</font></p>
	InvoicePending      *bool        `json:"invoice_pending,omitempty"`       // [Optional] <p>Use this field to filter the packages under invoice_pending.</p><p><font color="#c24f4a"><br /></font></p><p><font color="#c24f4a">Default value = false</font></p>
	SortingGroup        *int64       `json:"sorting_group,omitempty"`         // [Optional] <p>[Only for TW 30029 channel] Use this field to filter the sorting group of parcel. This field is only available for package with package_status = 3 and logistics_channel_id = 30029. Applicable values:&nbsp;<br />0: All</p><p>1: North<br />2: South</p>
	OrderType           *int64       `json:"order_type,omitempty"`            // [Optional] <p>Use this field to filter packages by order type. Applicable values:</p><p>0: All</p><p>1: Regular Order</p><p>2: Instant Order</p><p><br /></p><p><font color="#c24f4a">Default value = 0 (All)</font></p><p><br /></p><p>Note:&nbsp;For VN shops, using 2: Instant Order will return both Instant Delivery and Same-day Delivery packages.</p>
	IsPreOrder          *int64       `json:"is_pre_order,omitempty"`          // [Optional] <p>Use this field to filter packages by pre-order status. Applicable values:</p><p>0: All</p><p>1:&nbsp;Pre-Order</p><p>2:&nbsp;Non Pre-Order</p><p><br /></p><p><font color="#c24f4a">Default value = 0 (All)</font></p>
	ShippingPriority    *int64       `json:"shipping_priority,omitempty"`     // [Optional] <p>Use this field to filter packages by shipping priority. Applicable values:&nbsp;</p><p><br /></p><p>- For MY/PH/TW/TH shops, and VN Preferred/Preferred Plus/Shopee Mall shops:</p><p>0: All</p><p>1: Overdue</p><p>2: Ship by Today</p><p>3: Ship by Tomorrow</p><p><br /></p><p>- For other shops：<br />0: All<br />1: Overdue<br />2: Within 24h<br />3: Beyond 24h<br /></p><p><br /></p><p><font color="#c24f4a">Default value = 0 (All)</font></p>
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

type FlashSale struct {
	TimeslotId       int64 `json:"timeslot_id"`        // [Required]
	FlashSaleId      int64 `json:"flash_sale_id"`      // [Required]
	Status           int64 `json:"status"`             // [Required] <p>the status of shop flash sale</p><p>0: deleted</p><p>1: enabled</p><p>2: disabled</p><p>3: system_rejected, you cannot edit the shop flash sale in 'system_rejected' status</p>
	StartTime        int64 `json:"start_time"`         // [Required] <p>the start time of shop flash sale<br /></p>
	EndTime          int64 `json:"end_time"`           // [Required] <p>the end time of shop flash sale<br /></p>
	EnabledItemCount int64 `json:"enabled_item_count"` // [Required] <p>the number of enabled items in shop flash sale</p>
	ItemCount        int64 `json:"item_count"`         // [Required] <p>the number of items in shop flash sale<br /></p>
	Type             int64 `json:"type"`               // [Required] <p>the state of shop flash sale</p><p>1: upcoming</p><p>2: ongoing</p><p>3: expired</p>
	RemindmeCount    int64 `json:"remindme_count"`     // [Required] <p>No. of Reminders Set<br /></p>
	ClickCount       int64 `json:"click_count"`        // [Required] <p>No. of Product Clicks<br /></p>
}

type FollowPrizeList struct {
	CampaignId      int64          `json:"campaign_id"`       // [Required] <p>The unique identifier for the created follow prize.<br /></p>
	CampaignStatus  CampaignStatus `json:"campaign_status"`   // [Required] <p>The status of follow prize,the campagin status have upcoming/ongoing/expired.<br /></p>
	FollowPrizeName string         `json:"follow_prize_name"` // [Required] <p>The name of the follow prize,The follow prize name length max limit is 20.<br /></p>
	StartTime       int64          `json:"start_time"`        // [Required] <p>The timing from when the follow prize is valid,the start time later than the current time.If the start time and end time passed in by the seller overlap with other upcoming/ongoing activities, it will prompt "Another Follow Prize voucher already exists during this time period, please set another period."<br /></p>
	EndTime         int64          `json:"end_time"`          // [Required] <p>The timing until when the follow prize is still valid,the end time must be greater than the start time by at least 1 day and end time cannot exceed 3 months after start time.If the start time and end time passed in by the seller overlap with other upcoming/ongoing activities, it will prompt "Another Follow Prize voucher already exists during this time period, please set another period."<br /></p>
	UsageQuantity   int64          `json:"usage_quantity"`    // [Required] <p>Please enter a value between 1 and 200000.<br /></p>
	Claimed         int64          `json:"claimed"`           // [Required] <p>This is to indicate the quantity of voucher claimed.<br /></p>
}

type FollowUpAction struct {
	ItemId               int64    `json:"item_id"`                 // [Required] <p>Unique identifier of the item.</p>
	ModelId              int64    `json:"model_id"`                // [Required] <p>Unique identifier of the model under the item.</p>
	Qty                  int64    `json:"qty"`                     // [Required] <p>Quantity of items or models under the same current status.</p>
	CurrentStatus        int64    `json:"current_status"`          // [Required] <p>Current status for the item/model within the warehouse.</p><p><br /></p><p>Applicable values:</p><p>1：Dispose</p><p>2：Return to Seller</p><p>7：Received and Putaway</p><p>8：Return to Buyer</p><p>9：Shortage</p><p><br /></p><p>Note: Since Resell is currently applicable only to Failed Delivery parcels, the following values will not be returned for now, and will be returned once Resell becomes applicable to Return Refund parcels in the future:</p><p>3：Putaway for Resell</p><p>4：Resell Outbound</p><p>5：Resell Failed</p><p>6：Resell Exit</p>
	RelatedOrderSnList   []string `json:"related_order_sn_list"`   // [Required] <p>List of order_sn generated from the Resell process. Returned only when current_status = 4 (Resell Outbound).</p><p><br /></p><p>Note: Since Resell is currently applicable only to Failed Delivery parcels, this field will remain empty for now, and valid values will be returned once Resell becomes applicable to Return Refund parcels in the future.</p>
	ResellFailedNextStep string   `json:"resell_failed_next_step"` // [Required] <p>Next step after a Resell failure. Returned only when current_status = 5 (Resell Failed).</p><p><br /></p><p>Note: Since Resell is currently applicable only to Failed Delivery parcels, this field will remain empty for now, and valid values will be returned once Resell becomes applicable to Return Refund parcels in the future.</p>
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

type GenerateFbsInvoicesRequest struct {
	BatchDownload *BatchDownload `json:"batch_download,omitempty"` // [Optional]
}

type GenerateFbsInvoicesResponse struct {
	BaseResponse `json:",inline"` // Common response fields
	ErrorMsg     string           `json:"error_msg,omitempty"`   // <p>Error messages</p>
	ResultList   []ResponseResult `json:"result_list,omitempty"` //
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

type GenerateIncomeReportRequest struct {
	ReleaseTimeFrom int64 `json:"release_time_from"` // [Required] <p>Start time in epoch<br /></p>
	ReleaseTimeTo   int64 `json:"release_time_to"`   // [Required] <p>End time in epoch</p>
}

type GenerateIncomeReportResponse struct {
	BaseResponse `json:",inline"`                 // Common response fields
	Response     GenerateIncomeReportResponseData `json:"response"`      // Actual response data
	Msg          string                           `json:"msg,omitempty"` // <p>error message</p>
}

type GenerateIncomeReportResponseData struct {
	Id int64 `json:"id"` // [Required] <p>Identifier of income report file.</p>
}

type GenerateIncomeStatementRequest struct {
	ReleaseTimeFrom int64 `json:"release_time_from" url:"release_time_from"` // [Required] <p>The release_time_from must be</p><p><b>- Monday</b>&nbsp;(local time) for a weekly report</p><p>- <b>The 1st day</b> (local time) of a Month for a monthly report</p>
	ReleaseTimeTo   int64 `json:"release_time_to" url:"release_time_to"`     // [Required] <p>The release_time_to must be</p><p>- <b>Sunday</b> (local time) for a weekly report</p><p>-&nbsp;<b>The last day</b>&nbsp;(local time) of a Month for a monthly report</p>
	StatementType   int64 `json:"statement_type" url:"statement_type"`       // [Required] <p>STATEMENT_TYPE_WEEKLY = 1;</p><p>STATEMENT_TYPE_MONTHLY = 2;</p><p><br /></p><p>Local seller Income statement requires this value to be set.</p><p>CB seller income statement does not require this.</p>
}

type GenerateIncomeStatementResponse struct {
	BaseResponse `json:",inline"`                    // Common response fields
	Response     GenerateIncomeStatementResponseData `json:"response"` // Actual response data
}

type GenerateIncomeStatementResponseData struct {
	Id int64 `json:"id"` // [Required] <p>Identifier of income statement file.</p>
}

type GenerateKitImageRequest struct {
	ComponentList []RequestComponent `json:"component_list"` // [Required] <p>Please send up until 9 components.</p>
}

type GenerateKitImageResponse struct {
	BaseResponse `json:",inline"`             // Common response fields
	Response     GenerateKitImageResponseData `json:"response"` // Actual response data
}

type GenerateKitImageResponseData struct {
	KitImage string `json:"kit_image"` // [Required] <p>generated kit image</p>
}

type Geolocation struct {
	Latitude  float64 `json:"latitude"`  // [Required] <p>Latitude.</p>
	Longitude float64 `json:"longitude"` // [Required] <p>Longitude.</p>
}

type GetAddOnDealListRequest struct {
	PromotionStatus PromotionStatus `json:"promotion_status"`    // [Required] The Status of add on deal，default status is all
	PageNo          *int64          `json:"page_no,omitempty"`   // [Optional] The default page number is 1
	PageSize        *int64          `json:"page_size,omitempty"` // [Optional] The default page size is 100
}

type GetAddOnDealListResponse struct {
	BaseResponse `json:",inline"`             // Common response fields
	Response     GetAddOnDealListResponseData `json:"response"` // Actual response data
}

type GetAddOnDealListResponseData struct {
	AddOnDealList []AddOnDeal `json:"add_on_deal_list"` // [Required] The list of add on deal id
	More          bool        `json:"more"`             // [Required] This is to indicate whether the promotion list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of promotions.
}

type GetAddOnDealMainItemRequest struct {
	AddOnDealId int64 `json:"add_on_deal_id"` // [Required] Shopee's unique identifier for add on deal activity.
}

type GetAddOnDealMainItemResponse struct {
	BaseResponse `json:",inline"`                 // Common response fields
	Response     GetAddOnDealMainItemResponseData `json:"response"` // Actual response data
}

type GetAddOnDealMainItemResponseData struct {
	MainItemList []AddBundleDealItemRequestItem `json:"main_item_list"` // [Required] The main items added in this add on deal promotion.
	AddOnDealId  int64                          `json:"add_on_deal_id"` // [Required] Shopee's unique identifier for add on deal activity.
}

type GetAddOnDealRequest struct {
	AddOnDealId int64 `json:"add_on_deal_id"` // [Required] Shopee's unique identifier for an add on deal activity.
}

type GetAddOnDealResponse struct {
	BaseResponse `json:",inline"`         // Common response fields
	Response     GetAddOnDealResponseData `json:"response"` // Actual response data
}

type GetAddOnDealResponseData struct {
	StartTime              int64   `json:"start_time"`               // [Required] The time when add on deal activity start.
	EndTime                int64   `json:"end_time"`                 // [Required] The time when add on deal activity end
	PromotionType          int64   `json:"promotion_type"`           // [Required] The type of add on deal：add on discount =0；gift with mini spend=1
	PurchaseMinSpend       float64 `json:"purchase_min_spend"`       // [Required] The minimum purchase amount that needs to be met to buy the gift with min.Spend
	AddOnDealId            int64   `json:"add_on_deal_id"`           // [Required] Shopee's unique identifier for an add on deal activity.
	PerGiftNum             int64   `json:"per_gift_num"`             // [Required] Number of gifts that buyers can get
	SubItemPriority        []int64 `json:"sub_item_priority"`        // [Required] The order of the sub item
	PromotionPurchaseLimit int64   `json:"promotion_purchase_limit"` // [Required] Max. number of add-on products that a customer can purchase per order.
	AddOnDealName          string  `json:"add_on_deal_name"`         // [Required] Title of the add on deal
	Source                 int64   `json:"source"`                   // [Required]
}

type GetAddOnDealSubItemRequest struct {
	AddOnDealId int64 `json:"add_on_deal_id"` // [Required] Shopee's unique identifier for add on deal activity.
}

type GetAddOnDealSubItemResponse struct {
	BaseResponse `json:",inline"`                // Common response fields
	Response     GetAddOnDealSubItemResponseData `json:"response"` // Actual response data
}

type GetAddOnDealSubItemResponseData struct {
	SubItemList []GetAddOnDealSubItemResponseDataSubItem `json:"sub_item_list"`  // [Required] The sub items added in this add on deal promotion.
	AddOnDealId int64                                    `json:"add_on_deal_id"` // [Required] Shopee's unique identifier for add on deal activity.
}

type GetAddOnDealSubItemResponseDataSubItem struct {
	ItemId       int64         `json:"item_id"`        // [Required] Shopee's unique identifier for an item.
	Status       int64         `json:"status"`         // [Required] The status of add on deal item：enable = 1；disable =2
	SubItemLimit int64         `json:"sub_item_limit"` // [Required] The purchase limit of each sub item. Only the add on discount can be set and the default limit of gift with mini.spend is 1
	ModelId      int64         `json:"model_id"`       // [Required] Shopee's unique identifier for a model.
	Price        *SubItemPrice `json:"price"`          // [Required]
}

type GetAddressListResponse struct {
	BaseResponse `json:",inline"`           // Common response fields
	Response     GetAddressListResponseData `json:"response"` // Actual response data
}

type GetAddressListResponseData struct {
	ShowPickupAddress bool                  `json:"show_pickup_address"` // [Required] Show pickup address or not.
	AddressList       []ResponseDataAddress `json:"address_list"`        // [Required] The address list of you shop
}

type GetAdsFcilShopRateResponse struct {
	BaseResponse `json:",inline"` // Common response fields
	Rate         float64          `json:"rate,omitempty"`      // <p>The rate of the shop who choose to participate in this program</p>
	UpdateAt     int64            `json:"update_at,omitempty"` // <p>The update time in timestamp format</p>
}

type GetAitemByPitemIdRequest struct {
	PitemId int64 `json:"pitem_id" url:"pitem_id"` // [Required] <p>ID of item under SIP Primary Shop.</p>
}

type GetAitemByPitemIdResponse struct {
	BaseResponse `json:",inline"`              // Common response fields
	Response     GetAitemByPitemIdResponseData `json:"response"` // Actual response data
}

type GetAitemByPitemIdResponseData struct {
	AitemList []Aitem `json:"aitem_list"` // [Required]
}

type GetAllCpcAdsDailyPerformanceRequest struct {
	StartDate string `json:"start_date" url:"start_date"` // [Required] <p>This is the parameter to indicate the start date of the time length of performance.<br /></p>
	EndDate   string `json:"end_date" url:"end_date"`     // [Required] <p>This is the parameter to indicate the end date of the time length of performance<br /></p>
}

type GetAllCpcAdsDailyPerformanceResponse struct {
	BaseResponse `json:",inline"`                         // Common response fields
	Response     GetAllCpcAdsDailyPerformanceResponseData `json:"response"` // Actual response data
}

type GetAllCpcAdsDailyPerformanceResponseData struct {
	Date              string  `json:"date"`                // [Required] <p>This is the parameter to indicate which date the performance record belongs to.<br /></p>
	Impression        int64   `json:"impression"`          // [Required] <p>Number of times buyers see ads<br /></p>
	Clicks            int64   `json:"clicks"`              // [Required] <p>Total number of clicks on the Ad<br /></p>
	Ctr               float64 `json:"ctr"`                 // [Required] <p>Ctr, click-through rate measures how often shoppers who see an ad end up clicking it. CTR = Clicks / Impressions<br /></p>
	DirectOrder       int64   `json:"direct_order"`        // [Required] <p>Buyer place an order within 7 days after clicking on the ads (item gets purchased from the clicked ads)<br /></p><p><u>Please kindly note that the&nbsp;<b>direct_order</b>&nbsp;in the API reflected to Seller Center - Shopee Ads Module FE is&nbsp;<b>Direct Conversions.</b></u><br /></p>
	BroadOrder        int64   `json:"broad_order"`         // [Required] <p>Buyer place an order within 7 days after clicking on the ads; (the item gets purchased as long as there are other items from the same shops got click.)<br /></p><p><u>Please kindly note that the <b>broad_order</b>&nbsp;in the API reflected to Seller Center - Shopee Ads Module FE</u><u>&nbsp;is <b>Conversions</b>.</u></p>
	DirectConversions float64 `json:"direct_conversions"`  // [Required] <p>Ad orders / total number of clicks on the Ad. (item gets purchased from the clicked ads.)<br /></p><p><u>Please kindly note that the&nbsp;<b>direct_conversions</b> in the API reflected to Seller Center - Shopee Ads Module FE is&nbsp;<b>Direct Conversions Rate</b>.</u><br /></p>
	BroadConversions  float64 `json:"broad_conversions"`   // [Required] <p>Ad orders / total number of clicks on the Ad. (the item gets purchased as long as there are other items from the same shops got click.)<br /></p><p><u>Please kindly note that the <b>broad_conversions</b>&nbsp;in the API reflected to Seller Center - Shopee Ads Module FE is <b>Conversions Rate</b>.</u><br /></p>
	DirectItemSold    int64   `json:"direct_item_sold"`    // [Required] <p>item sold within 7 days after clicking on the ads. (item gets purchased from the clicked ads.)<br /></p>
	BroadItemSold     int64   `json:"broad_item_sold"`     // [Required] <p>item sold within 7 days after clicking on the ads.(the item gets purchased as long as there are other items from the same shops got click.)<br /></p><p><u>Please kindly note that the <b>broad_conversions</b> in the API reflected to Advertiser Platform is <b>Conversion Rate</b>.</u><br /></p>
	DirectGmv         float64 `json:"direct_gmv"`          // [Required] <p>Total sales generated from Ad over a certain time frame Typically 7 days. (item gets purchased from the clicked ads.)<br /></p>
	BroadGmv          float64 `json:"broad_gmv"`           // [Required] <p>Total sales generated from Ad over a certain time frame (the item gets purchased as long as there are other items from the same shops got click.)<br /></p>
	Expense           float64 `json:"expense"`             // [Required] <p>Ad Expenditure<br /></p>
	CostPerConversion float64 `json:"cost_per_conversion"` // [Required] <p>(Cost Per Conversion) Ad's average cost per sales conversion<br /></p>
	DirectRoas        float64 `json:"direct_roas"`         // [Required] <p>Ad GMV/Ad Expenditure. (item gets purchased from the clicked ads.)<br /></p>
	BroadRoas         float64 `json:"broad_roas"`          // [Required] <p>Ad GMV/Ad Expenditure. (the item gets purchased as long as there are other items from the same shops got click.)<br /></p>
}

type GetAllCpcAdsHourlyPerformanceRequest struct {
	PerformanceDate string `json:"performance_date" url:"performance_date"` // [Required] <p>This is the parameter of the single date on which requester wants to check the hourly performance. Date in DD-MM-YYYY format.<br /></p>
}

type GetAllCpcAdsHourlyPerformanceResponse struct {
	BaseResponse `json:",inline"`                          // Common response fields
	Response     GetAllCpcAdsHourlyPerformanceResponseData `json:"response"` // Actual response data
}

type GetAllCpcAdsHourlyPerformanceResponseData struct {
	Hour              int64   `json:"hour"`                // [Required] <p>This is the parameter to indicate each hour the performance record belongs to.<br /></p>
	Date              string  `json:"date"`                // [Required] <p>This is the parameter to indicate which date the performance record belongs to.<br /></p>
	Impression        int64   `json:"impression"`          // [Required] <p>Number of times buyers see ads<br /></p>
	Clicks            int64   `json:"clicks"`              // [Required] <p>Total number of clicks on the Ad<br /></p>
	Ctr               float64 `json:"ctr"`                 // [Required] <p>Ctr, click-through rate measures how often shoppers who see an ad end up clicking it. CTR = Clicks / Impressions<br /></p>
	DirectOrder       int64   `json:"direct_order"`        // [Required] <p>Buyer place an order within 7 days after clicking on the ads (item gets purchased from the clicked ads).<br /></p><p><u>Please kindly note that the <b>direct_order</b>&nbsp;in the API reflected to Seller Center Shopee Ads Module FE is <b>Direct Conversions.</b></u></p>
	BroadOrder        int64   `json:"broad_order"`         // [Required] <p>Buyer place an order within 7 days after clicking on the ads; (the item gets purchased as long as there are other items from the same shops got click.)<br /></p><p><u>Please kindly note that the<b> broad_order </b>in the API<b>&nbsp;</b>reflected to Seller Center Shopee Ads Module FE is <b>Conversions.</b></u></p>
	DirectConversions float64 `json:"direct_conversions"`  // [Required] <p>Ad orders / total number of clicks on the Ad. (item gets purchased from the clicked ads.)<br /></p><p><u>Please kindly note that the <b>direct_conversions</b> in the API reflected to Seller Center Shopee Ads Module FE is the&nbsp;<b>Direct Conversion Rate</b></u></p>
	BroadConversions  float64 `json:"broad_conversions"`   // [Required] <p>Ad orders / total number of clicks on the Ad. (the item gets purchased as long as there are other items from the same shops got click.)<br /></p><p><u>Please kindly note that the <b>broad conversions</b> in the API reflected to Seller Center Shopee Ads Module FE is <b>Conversion Rate</b></u></p>
	DirectItemSold    int64   `json:"direct_item_sold"`    // [Required] <p>item sold within 7 days after clicking on the ads. (item gets purchased from the clicked ads.)<br /></p>
	BroadItemSold     int64   `json:"broad_item_sold"`     // [Required] <p>item sold within 7 days after clicking on the ads.(the item gets purchased as long as there are other items from the same shops got click.)<br /></p>
	DirectGmv         float64 `json:"direct_gmv"`          // [Required] <p>Total sales generated from Ad over a certain time frame Typically 7 days. (item gets purchased from the clicked ads.)<br /></p>
	BroadGmv          float64 `json:"broad_gmv"`           // [Required] <p>Total sales generated from Ad over a certain time frame (the item gets purchased as long as there are other items from the same shops got click.)<br /></p>
	Expense           float64 `json:"expense"`             // [Required] <p>Ad Expenditure<br /></p>
	CostPerConversion float64 `json:"cost_per_conversion"` // [Required] <p>(Cost Per Conversion) Ad's average cost per sales conversion<br /></p>
	DirectRoas        float64 `json:"direct_roas"`         // [Required] <p>Ad GMV/Ad Expenditure. (item gets purchased from the clicked ads.)<br /></p>
	BroadRoas         float64 `json:"broad_roas"`          // [Required] <p>Ad GMV/Ad Expenditure. (the item gets purchased as long as there are other items from the same shops got click.)<br /></p>
}

type GetAllVehicleListRequest struct {
	PageSize int64   `json:"page_size" url:"page_size"`                   // [Required] <p>The size of one page. Max=100</p>
	Offset   *int64  `json:"offset,omitempty" url:"offset,omitempty"`     // [Optional] <p>Specifies the starting entry of data to return in the current call. Default is 0, if data is more than one page, the offset can be some entry to start next call.<br /></p>
	Language *string `json:"language,omitempty" url:"language,omitempty"` // [Optional] <p>If language is not uploaded, the default language=en, the following are the languages supported by different markets SG: en ; MY: en / ms-my / zh-hans ; TH: en / th ; VN: en / vi ; PH: en ; TW: en / zh-hant ; ID: en / id ; BR: en / pt-br ; MX: en / es-mx ; CO: en/es-CO ; CL: en/es-CL. Note: For markets that have already launched global tree, Crossboard shop only support returning en and zh-hans language data<br /></p>
}

type GetAllVehicleListResponse struct {
	BaseResponse `json:",inline"`              // Common response fields
	Response     GetAllVehicleListResponseData `json:"response"` // Actual response data
}

type GetAllVehicleListResponseData struct {
	VehicleList []Vehicle `json:"vehicle_list"`  // [Required]
	HasNextPage bool      `json:"has_next_page"` // [Required] <p>This is to indicate whether the item list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of items.<br /></p>
	NextOffset  int64     `json:"next_offset"`   // [Required] <p>If has_next_page is true, this value need set to next request offset<br /></p>
}

type GetAppPushConfigResponse struct {
	BaseResponse `json:",inline"`             // Common response fields
	Response     GetAppPushConfigResponseData `json:"response"` // Actual response data
}

type GetAppPushConfigResponseData struct {
	CallbackUrl       string  `json:"callback_url"`         // [Required] <p>The callback url of push mechanism. It is the address where the Shopee will send the push message to. If you don't set any callback_url before, this parameters is required.<br /></p>
	LivePushStatus    string  `json:"live_push_status"`     // [Required] <p>live push status:Normal,Warning,Suspended<br /></p>
	SuspendedTime     int64   `json:"suspended_time"`       // [Required] <p>The live push suspended time caused by low successful rate of push mechanism.Only when live push status is suspended, this parameters will response.<br /></p>
	BlockedShopId     []int64 `json:"blocked_shop_id"`      // [Required] <p>Use this filed to indicate that Shopee won't send any push message created by this shop.<br /></p>
	PushConfigOnList  []int64 `json:"push_config_on_list"`  // [Required] <p>Use this field to indicate which push config turn on, and you can receive the push message.</p><p>1=Shop authorization for partners</p><p>2=Shop deauthorization for partners</p><p>3=Order status update push</p><p>4=TrackingNo push</p><p>5=Shopee Updates</p><p>6=Banned item push</p><p>7=item promotion push</p><p>8=reserved stock change push</p><p>9=promotionn update push</p><p>10=webchat push</p><p>11=video upload push</p><p>12=openapi authorization expiry push</p><p>13=brand register result<br /></p>
	PushConfigOffList []int64 `json:"push_config_off_list"` // [Required] <p>Use this field to indicate which push config turn on, and you can receive the push message.</p><p>1=Shop authorization for partners</p><p>2=Shop deauthorization for partners</p><p>3=Order status update push</p><p>4=TrackingNo push</p><p>5=Shopee Updates</p><p>6=Banned item push</p><p>7=item promotion push</p><p>8=reserved stock change push</p><p>9=promotionn update push</p><p>10=webchat push</p><p>11=video upload push</p><p>12=openapi authorization expiry push</p><p>13=brand register result<br /></p>
}

type GetAttributeTreeRequest struct {
	CategoryIdList []int64 `json:"category_id_list" url:"category_id_list"`     // [Required] <p>max count is 20</p>
	Language       *string `json:"language,omitempty" url:"language,omitempty"` // [Optional] <p>Language<br />Support Lanuage:<br />"SG": [        "en",        "zh-Hans",        "ms"      ],&nbsp;</p><p>"MY": [        "en",        "zh-Hans",        "ms"      ], <br />"PH": [        "en",        "zh-Hans"      ], <br />"VN": [        "vn",        "en"      ], <br />"ID": [        "id",        "en"      ], <br />"TH": [        "th",        "en"      ], <br />"BR": [        "pt-BR",        "en"      ], <br />"MX": [        "es-MX",        "en"      ], <br />"CO": [        "es-CO",        "en"      ], <br />"CL": [        "es-CL",        "en"      ], <br />"TW": [        "zh-Hant",        "zh-Hans",        "en"      ],<br />"IN": [        "en",        "hi"      ]<br /></p>
}

type GetAttributeTreeResponse struct {
	BaseResponse `json:",inline"`             // Common response fields
	Response     GetAttributeTreeResponseData `json:"response"` // Actual response data
}

type GetAttributeTreeResponseData struct {
	List []List `json:"list"` // [Required] <p>Each result corresponds to one category in category_ids<br /></p>
}

type GetAttributeTreeResponseDataList struct {
	AttributeTree []ListAttributeTree `json:"attribute_tree"` // [Required] <p>One category's attribute trees<br /></p>
	CategoryId    int64               `json:"category_id"`    // [Required] <p>Category ID<br /></p>
	Warning       string              `json:"warning"`        // [Required] <p>Warning msg</p>
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

type GetAvailableSolutionsRequest struct {
	ReturnSn string `json:"return_sn" url:"return_sn"` // [Required] The serial number of return.
}

type GetAvailableSolutionsResponse struct {
	BaseResponse `json:",inline"`                  // Common response fields
	Response     GetAvailableSolutionsResponseData `json:"response"` // Actual response data
}

type GetAvailableSolutionsResponseData struct {
	ReturnSn          string             `json:"return_sn"`           // [Required] <p>The serial number of return.<br /></p>
	OfferReturnRefund *OfferReturnRefund `json:"offer_return_refund"` // [Required]
	OfferRefund       *OfferReturnRefund `json:"offer_refund"`        // [Required]
}

type GetBillingTransactionInfoRequest struct {
	BillingTransactionInfoType int64    `json:"billing_transaction_info_type"`  // [Required] <p>Billing transaction types: 1: TO_RELEASE, 2:&nbsp;RELEASED</p>
	EncryptedPayoutIds         []string `json:"encrypted_payout_ids,omitempty"` // [Optional] <p>encrypted_payout_id get from API: v2.get_payout_info</p><p><br /></p><p><br /></p><p>when&nbsp;encrypted_payout_id provided and billing_transaction_info_type=2, we will return the "released" billing items under this payout.</p><p>when&nbsp;encrypted_payout_id not provided, we will return the "to release" billing items under hasn't form payout yet<br /></p><p><br /></p><p>Max length: 100</p>
	Cursor                     string   `json:"cursor"`                         // [Required] <p>Specifies the starting entry of data to return in the current call. Default is "". If data is more than one page, the offset can be some entry to start next call.<br /></p>
	PageSize                   int64    `json:"page_size"`                      // [Required] <p>Number of pages returned max:100<br /></p>
}

type GetBillingTransactionInfoResponse struct {
	BaseResponse `json:",inline"`                      // Common response fields
	Response     GetBillingTransactionInfoResponseData `json:"response"` // Actual response data
}

type GetBillingTransactionInfoResponseData struct {
	Transactions *Transactions `json:"transactions"` // [Required]
	More         bool          `json:"more"`         // [Required]
	NextCursor   string        `json:"next_cursor"`  // [Required]
}

type GetBookingDetailRequest struct {
	BookingSnList          string  `json:"booking_sn_list" url:"booking_sn_list"`                                       // [Required] <p>The set of booking_sn. If there are multiple booking_sn, you need to use English comma to connect them. limit [1,50]<br /></p>
	ResponseOptionalFields *string `json:"response_optional_fields,omitempty" url:"response_optional_fields,omitempty"` // [Optional] <p>The response fields you want to get. Please select from the below response parameters. If you input an object field, all the params under it will be included automatically in the response. If there are multiple response fields you want to get, you need to use English comma to connect them. Available values:&nbsp;item_list,cancel_by,cancel_reason,fulfillment_flag,pickup_done_time,shipping_carrier, recipient_address, dropshipper, dropshipper_phone<br /></p>
}

type GetBookingDetailResponse struct {
	BaseResponse `json:",inline"`             // Common response fields
	Response     GetBookingDetailResponseData `json:"response"` // Actual response data
}

type GetBookingDetailResponseData struct {
	BookingList []GetBookingDetailResponseDataBooking `json:"booking_list"` // [Required] <p>The list of bookings.<br /></p>
}

type GetBookingDetailResponseDataBooking struct {
	BookingSn        string                   `json:"booking_sn"`        // [Required] <p>Return by default. Shopee's unique identifier for a booking.<br /></p>
	OrderSn          string                   `json:"order_sn"`          // [Required] <p>Shopee's unique identifier for an order. Only return if booking_status is MATCHED.<br /></p>
	Region           string                   `json:"region"`            // [Required] <p>Return by default. The two-digit code representing the region where the booking was made.<br /></p>
	BookingStatus    BookingStatus            `json:"booking_status"`    // [Required] <p>Return by default.&nbsp;Enumerated type that defines the current status of the booking. Available value:&nbsp;READY_TO_SHIP/PROCESSED/SHIPPED/CANCELLED/MATCHED<br /></p>
	MatchStatus      string                   `json:"match_status"`      // [Required] <p>MATCH_PENDING/MATCH_SUCCESSFUL/MATCH_FAILED<br /></p>
	ShippingCarrier  string                   `json:"shipping_carrier"`  // [Required] <p>The logistics service provider&nbsp;that will deliver the booking.<br /></p>
	CreateTime       int64                    `json:"create_time"`       // [Required] <p>Return by default. Timestamp that indicates the date and time that the booking was created.<br /></p>
	UpdateTime       int64                    `json:"update_time"`       // [Required] <p>Return by default. Timestamp that indicates the last time that there was a change in value of booking, such as booking status changed from 'Processed' to 'Shipped'.<br /></p>
	ShipByDate       int64                    `json:"ship_by_date"`      // [Required] <p>Return by default. The deadline to ship out the parcel.<br /></p>
	RecipientAddress *BookingRecipientAddress `json:"recipient_address"` // [Required] <p>This object contains detailed breakdown for the recipient address.<br /></p>
	ItemList         []BookingItem            `json:"item_list"`         // [Required] <p>This object contains the detailed breakdown for the result of this API call.<br /></p>
	Dropshipper      string                   `json:"dropshipper"`       // [Required] <p>For Indonesia bookings only. The name of the dropshipper.<br /></p>
	DropshipperPhone string                   `json:"dropshipper_phone"` // [Required] <p>The phone number of dropshipper, could be empty.<br /></p>
	CancelBy         string                   `json:"cancel_by"`         // [Required] <p>Could be one of buyer, seller, system or Ops.<br /></p>
	CancelReason     string                   `json:"cancel_reason"`     // [Required] <p>Use this field to get reason for buyer, seller, and system cancellation.<br /></p>
	FulfillmentFlag  string                   `json:"fulfillment_flag"`  // [Required] <p>Use this field to indicate the booking is fulfilled by shopee or seller. Applicable values:&nbsp;fulfilled_by_shopee, fulfilled_by_cb_seller, fulfilled_by_local_seller.<br /></p>
	PickupDoneTime   int64                    `json:"pickup_done_time"`  // [Required] <p>The timestamp when pickup is done.<br /></p>
}

type GetBookingListRequest struct {
	TimeRangeField string         `json:"time_range_field" url:"time_range_field"`                 // [Required] <p>The kind of time_from and time_to.&nbsp;Available value: create_time, update_time.<br /></p>
	TimeFrom       int64          `json:"time_from" url:"time_from"`                               // [Required] <p>The time_from and time_to fields specify a date range for retrieving bookings (based on the time_range_field). The time_from field is the starting date range. The maximum date range that may be specified with the time_from and time_to fields is 15 days.<br /></p>
	TimeTo         int64          `json:"time_to" url:"time_to"`                                   // [Required] <p>The time_from and time_to fields specify a date range for retrieving bookings (based on the time_range_field). The time_from field is the starting date range. The maximum date range that may be specified with the time_from and time_to fields is 15 days.</p>
	PageSize       int64          `json:"page_size" url:"page_size"`                               // [Required] <p>Each result set is returned as a page of entries. Use the "page_size" filters to control the maximum number of entries to retrieve per page (i.e., per call). This integer value is used to specify the maximum number of entries to return in a single "page" of data.The limit of page_size if between 1 and 100.<br /></p>
	Cursor         *string        `json:"cursor,omitempty" url:"cursor,omitempty"`                 // [Optional] <p>Specifies the starting entry of data to return in the current call. Default is "". If data is more than one page, the offset can be some entry to start next call.<br /></p>
	BookingStatus  *BookingStatus `json:"booking_status,omitempty" url:"booking_status,omitempty"` // [Optional] <p>The booking_status filter for retrieving bookings and each one only every request. Available value: READY_TO_SHIP/PROCESSED/SHIPPED/CANCELLED/MATCHED<br /></p>
}

type GetBookingListResponse struct {
	BaseResponse `json:",inline"`           // Common response fields
	Response     GetBookingListResponseData `json:"response"` // Actual response data
}

type GetBookingListResponseData struct {
	More        bool                  `json:"more"`         // [Required] <p>This is to indicate whether the booking list is more than one page. If this value is true, you may want to continue to check next page to retrieve bookings.<br /></p>
	BookingList []ResponseDataBooking `json:"booking_list"` // [Required]
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

type GetBoostedListResponse struct {
	BaseResponse `json:",inline"`           // Common response fields
	Response     GetBoostedListResponseData `json:"response"` // Actual response data
}

type GetBoostedListResponseData struct {
	ItemList []GetBoostedListResponseDataItem `json:"item_list"` // [Required]
}

type GetBoostedListResponseDataItem struct {
	ItemId         int64 `json:"item_id"`          // [Required] Shopee's unique identifier for an item
	CoolDownSecond int64 `json:"cool_down_second"` // [Required] Remain cool down time
}

type GetBoundWhsInfoResponse struct {
	BaseResponse `json:",inline"`            // Common response fields
	Response     GetBoundWhsInfoResponseData `json:"response"` // Actual response data
}

type GetBoundWhsInfoResponseData struct {
	List []GetBoundWhsInfoResponseDataList `json:"list"` // [Required]
}

type GetBoundWhsInfoResponseDataList struct {
	BoundWhs []BoundWhs `json:"bound_whs"` // [Required]
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

type GetBrandListRequest struct {
	Offset     int64   `json:"offset" url:"offset"`                         // [Required] Specifies the starting entry of data to return in the current call. Default is 0. If data is more than one page,this field needs to be replaced with "next_offset" to request,and the offset can be some entry to start next call.
	PageSize   int64   `json:"page_size" url:"page_size"`                   // [Required] the size of one page.Max=100
	CategoryId int64   `json:"category_id" url:"category_id"`               // [Required] ID of category.
	Status     int64   `json:"status" url:"status"`                         // [Required] Brand status , 1: normal brand, 2: pending brand
	Language   *string `json:"language,omitempty" url:"language,omitempty"` // [Optional] <p>If language is not uploaded, the default language=en, the following are the languages supported by different markets SG: en ; MY: en / ms-my / zh-hans ; TH: en / th ; VN: en / vi ; PH: en ; TW: en / zh-hant ; ID: en / id ;  BR: en / pt-br ;  MX: en / es-mx ; CO: en/es-CO ; CL: en/es-CL. Note: For markets that have already launched global tree, Crossboard shop only support returning en and zh-hans language data</p>
}

type GetBrandListResponse struct {
	BaseResponse `json:",inline"`         // Common response fields
	Response     GetBrandListResponseData `json:"response"` // Actual response data
}

type GetBrandListResponseData struct {
	BrandList   []ResponseDataBrand `json:"brand_list"`    // [Required]
	HasNextPage bool                `json:"has_next_page"` // [Required]  This is to indicate whether the item list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of items.
	NextOffset  int64               `json:"next_offset"`   // [Required] If has_next_page is true, this value need set to next request.offset
	IsMandatory bool                `json:"is_mandatory"`  // [Required] Whether is mandatory.
	InputType   string              `json:"input_type"`    // [Required] <p>Input type: DROP_DOWN</p>
}

type GetBundleDealItemRequest struct {
	BundleDealId int64 `json:"bundle_deal_id"` // [Required] Shopee's unique identifier for a bundle deal activity.
}

type GetBundleDealItemResponse struct {
	BaseResponse `json:",inline"`              // Common response fields
	Response     GetBundleDealItemResponseData `json:"response"` // Actual response data
}

type GetBundleDealItemResponseData struct {
	ItemList   []AddBundleDealItemRequestItem `json:"item_list"`   // [Required] The list of bundle deal item
	TotalCount int64                          `json:"total_count"` // [Required] The number of  items in this bundle deal
}

type GetBundleDealListRequest struct {
	PageSize   *int64 `json:"page_size,omitempty" url:"page_size,omitempty"`     // [Optional] Data paging, representing the data size of each page, the maximum is 1000, the default is 20
	TimeStatus *int64 `json:"time_status,omitempty" url:"time_status,omitempty"` // [Optional] The Status of bundle deal，all=1；upcoming=2；ongoing=3，expired=4 , the default is 1
	PageNo     *int64 `json:"page_no,omitempty" url:"page_no,omitempty"`         // [Optional] Data paging, represents the page number, starting from 1, the default is 1
}

type GetBundleDealListResponse struct {
	BaseResponse `json:",inline"`              // Common response fields
	Response     GetBundleDealListResponseData `json:"response"` // Actual response data
}

type GetBundleDealListResponseData struct {
	BundleDealList []BundleDeal `json:"bundle_deal_list"` // [Required] The list of bundle deal id
	More           bool         `json:"more"`             // [Required] this field shows whether there are more bundle deals in next page or not
}

type GetBundleDealRequest struct {
	BundleDealId int64 `json:"bundle_deal_id" url:"bundle_deal_id"` // [Required] Shopee's unique identifier for a bundle deal activity.
}

type GetBundleDealResponse struct {
	BaseResponse `json:",inline"`          // Common response fields
	Response     GetBundleDealResponseData `json:"response"` // Actual response data
}

type GetBundleDealResponseData struct {
	BundleDealId   int64                       `json:"bundle_deal_id"`   // [Required] Shopee's unique identifier for a bundle deal activity.
	Name           string                      `json:"name"`             // [Required] Title of the bundle deal
	StartTime      int64                       `json:"start_time"`       // [Required] The time when bundle deal activity start.
	EndTime        int64                       `json:"end_time"`         // [Required] The time when bundle deal activity end.
	BundleDealRule *ResponseDataBundleDealRule `json:"bundle_deal_rule"` // [Required]
	PurchaseLimit  int64                       `json:"purchase_limit"`   // [Required] Maximum number of bundle deals that can be bought by a buyer.
}

type GetBuyerInvoiceInfoRequest struct {
	Queries []Queries `json:"queries"` // [Required]
}

type GetBuyerInvoiceInfoResponse struct {
	BaseResponse    `json:",inline"` // Common response fields
	InvoiceInfoList []InvoiceInfo    `json:"invoice_info_list,omitempty"` //
}

type GetCategoryRequest struct {
	Language *string `json:"language,omitempty" url:"language,omitempty"` // [Optional] <p>If language is not uploaded, the default language=en, the following are the languages supported by different markets SG: en ; MY: en / ms-my / zh-hans ; TH: en / th ; VN: en / vi ; PH: en ; TW: en / zh-hant ; ID: en / id ;  BR: en / pt-br ;  MX: en / es-mx ; CO: en/es-CO ; CL: en/es-CL .Note: For markets that have already launched global tree, Crossboard shop only support returning en and zh-hans language data</p>
}

type GetCategoryResponse struct {
	BaseResponse `json:",inline"`        // Common response fields
	Response     GetCategoryResponseData `json:"response"` // Actual response data
}

type GetCategoryResponseData struct {
	CategoryList []Category `json:"category_list"` // [Required]
}

type GetChannelListRequest struct {
	Region *string `json:"region,omitempty"` // [Optional] <p>Use this field to specify the region you want to ship parcel. Available value: CN, KR</p>
}

type GetChannelListResponse struct {
	BaseResponse `json:",inline"`           // Common response fields
	Response     GetChannelListResponseData `json:"response"` // Actual response data
}

type GetChannelListResponseData struct {
	LogisticsChannelList []LogisticsChannel `json:"logistics_channel_list"` // [Required] The list of logistics channel.
}

type GetCommentRequest struct {
	ItemId    *int64 `json:"item_id,omitempty" url:"item_id,omitempty"`       // [Optional] The identity of product item.
	CommentId *int64 `json:"comment_id,omitempty" url:"comment_id,omitempty"` // [Optional] The identity of comment.
	Cursor    string `json:"cursor" url:"cursor"`                             // [Required] Specifies the starting entry of data to return in the current call. Default is "". If data is more than one page, the offset can be some entry to start next call.
	PageSize  int64  `json:"page_size" url:"page_size"`                       // [Required] Each result set is returned as a page of entries. Use the "page_size" filters to control the maximum number of entries to retrieve per page (i.e., per call). This integer value is used to specify the maximum number of entries to return in a single "page" of data. The limit of page_size if between 1 and 100.
}

type GetCommentResponse struct {
	BaseResponse `json:",inline"`       // Common response fields
	Response     GetCommentResponseData `json:"response"` // Actual response data
}

type GetCommentResponseData struct {
	More            bool          `json:"more"`              // [Required] <p>This is to indicate whether the comment list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of comments. But only respond 500 comments at most through OpenAPI, if there are more than 500, this field "more" also respond "true".</p>
	ItemCommentList []ItemComment `json:"item_comment_list"` // [Required] The comment data list of the items.
	NextCursor      string        `json:"next_cursor"`       // [Required] If more is true, you should pass the next_cursor in the next request as cursor. The value of next_cursor will be empty string when more is false.
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

type GetCreateProductAdBudgetSuggestionRequest struct {
	ReferenceId               string   `json:"reference_id" url:"reference_id"`                                                     // [Required] <p>A random string used to prevent duplicate ads. If an ads is created successfully, subsequent request using the same reference id will fail<br /></p>
	ProductSelection          string   `json:"product_selection" url:"product_selection"`                                           // [Required] <p>auto,manual - for Auto product ads or Manual Product Ads</p>
	CampaignPlacement         string   `json:"campaign_placement" url:"campaign_placement"`                                         // [Required] <p>search, discovery, all<br /></p>
	BiddingMethod             string   `json:"bidding_method" url:"bidding_method"`                                                 // [Required] <p>Bidding Method of product ad: auto, manual</p>
	EnhancedCpc               *string  `json:"enhanced_cpc,omitempty" url:"enhanced_cpc,omitempty"`                                 // [Optional] <p>Enhanced CPC functionality toggle. Values supported "true"/"false". Mandatory for product_selection=manual, bidding_method=manual</p>
	DiscoveryAdsLocationNames *string  `json:"discovery_ads_location_names,omitempty" url:"discovery_ads_location_names,omitempty"` // [Optional] <p>List of comma separated location values from: daily_discover, you_may_also_like.<br /></p><p>Mandatory for product_selection=manual, product_placement={all|discovery}, bidding_method=manual</p>
	RoasTarget                *float64 `json:"roas_target,omitempty" url:"roas_target,omitempty"`                                   // [Optional] <p>the ROAS target for each campaign with auto bidding. If 0, GMV Max / ROI feature is not enabled</p>
	ItemId                    *int64   `json:"item_id,omitempty" url:"item_id,omitempty"`                                           // [Optional] <p>Product ID. Mandatory for product_selection=manual</p>
}

type GetCreateProductAdBudgetSuggestionResponse struct {
	BaseResponse `json:",inline"`                               // Common response fields
	Response     GetCreateProductAdBudgetSuggestionResponseData `json:"response"` // Actual response data
}

type GetCreateProductAdBudgetSuggestionResponseData struct {
	Budget *Budget `json:"budget"` // [Required] <p>Budget data</p>
}

type GetCurrentInventoryRequest struct {
	PageNo                 *int64  `json:"page_no,omitempty"`                  // [Optional] <p>Specifies the page number of data to return in the current call. Starting from 1. if data is more than one page, the page_no can be some entry to start next call. If empty, the default value is 1.</p>
	PageSize               *int64  `json:"page_size,omitempty"`                // [Optional] <p>Each result set is returned as a page of entries. Use the "page_size" filters to control the maximum number of entries to retrieve per page (i.e., per call), and the "page_no" to start next call. This integer value is used to specify the maximum number of entries to return in a single "page" of data.</p><p>If empty, the default value is 10. The value should be between 1 and 100.</p>
	SearchType             *int64  `json:"search_type,omitempty"`              // [Optional] <p>0-All data；1-Product Name；2-SKU ID；3-Variations；4-Item ID</p>
	Keyword                *string `json:"keyword,omitempty"`                  // [Optional] <p>Bind Value and Search_type</p>
	WhsIds                 *string `json:"whs_ids,omitempty"`                  // [Optional] <p>Whs ID list, comma-separated</p>
	NotMovingTag           *int64  `json:"not_moving_tag,omitempty"`           // [Optional] <p>Blank-All；0-No；1-Yes</p>
	InboundPendingApproval *int64  `json:"inbound_pending_approval,omitempty"` // [Optional] <p>Blank-All；0-No；1-Yes</p>
	ProductsWithInventory  *int64  `json:"products_with_inventory,omitempty"`  // [Optional] <p>Blank-All；0-No；1-Yes</p>
	CategoryId             *int64  `json:"category_id,omitempty"`              // [Optional] <p>Category id. Here you need to call the&nbsp;get_category&nbsp;API to retrieve the first-tier category_id.<br /></p>
	StockLevels            *string `json:"stock_levels,omitempty"`             // [Optional] <p>1-Low Stock &amp; No Sellable stock; 2-Low Stock &amp; To replenish; 3-Low Stock &amp; Replenished; 4-Excess<br /></p>
	WhsRegion              string  `json:"whs_region"`                         // [Required] <p>The warehouse region you want to query, can only query one region in a request<br />Optional value: BR、CN、ID、MY、MX、TH、TW、PH、VN、SG<br /><br /><b>If do not pass, will get error "block by gateway due to invalid cid"</b></p>
}

type GetCurrentInventoryResponse struct {
	BaseResponse `json:",inline"`                // Common response fields
	Response     GetCurrentInventoryResponseData `json:"response"` // Actual response data
}

type GetCurrentInventoryResponseData struct {
	ItemList []GetCurrentInventoryResponseDataItem `json:"item_list"` // [Required] <p>Data list of item sku</p>
}

type GetCurrentInventoryResponseDataItem struct {
	WarehouseItemId string `json:"warehouse_item_id"` // [Required] <p>Warehouse item id; To indicate an unique item in a warehouse<br /><br />one warehouse item id can match with multiple shop_item_id<br /><br />For Global Item,&nbsp;warehouse_item_id=Global Item id<br />For Local Item, shop_item_id=item_id</p>
	ItemName        string `json:"item_name"`         // [Required] <p>item name</p>
	ItemImage       string `json:"item_image"`        // [Required] <p>item image</p>
	SkuList         []Sku  `json:"sku_list"`          // [Required] <p>Data list of mtsku</p>
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

type GetDirectItemListRequest struct {
	MainItemId []int64 `json:"main_item_id" url:"main_item_id"` // [Required] <p>Item id of main shop.</p>
}

type GetDirectItemListResponse struct {
	BaseResponse `json:",inline"`              // Common response fields
	Response     GetDirectItemListResponseData `json:"response"` // Actual response data
}

type GetDirectItemListResponseData struct {
	List []GetDirectItemListResponseDataList `json:"list"` // [Required]
}

type GetDirectItemListResponseDataList struct {
	MainItemId     int64        `json:"main_item_id"`     // [Required] <p>Item id of main shop.</p>
	DirectItemList []DirectItem `json:"direct_item_list"` // [Required]
}

type GetDirectShopRecommendedPriceRequest struct {
	MainItemId           int64                                       `json:"main_item_id" url:"main_item_id"`                                           // [Required]
	DirectShopRegions    []string                                    `json:"direct_shop_regions" url:"direct_shop_regions"`                             // [Required] <p>Direct shop regions.</p>
	CategoryId           *int64                                      `json:"category_id,omitempty" url:"category_id,omitempty"`                         // [Optional] <pre>Main_item's category.</pre>
	ModelList            []GetDirectShopRecommendedPriceRequestModel `json:"model_list,omitempty" url:"model_list,omitempty"`                           // [Optional] <p>Main model model info.</p>
	EnabledChannelIdList []int64                                     `json:"enabled_channel_id_list,omitempty" url:"enabled_channel_id_list,omitempty"` // [Optional] <p>direct shop enabled channel</p>
}

type GetDirectShopRecommendedPriceRequestModel struct {
	ModelId    *int64   `json:"model_id,omitempty"`    // [Optional] <p>Id of main model.</p>
	TierIndex  []int64  `json:"tier_index,omitempty"`  // [Optional] <p>Tier index of main model. Index starts from 0.</p><p><br /></p>
	InputPrice *int64   `json:"input_price,omitempty"` // [Optional]
	Weight     *float64 `json:"weight,omitempty"`      // [Optional]
}

type GetDirectShopRecommendedPriceResponse struct {
	BaseResponse `json:",inline"`                          // Common response fields
	Response     GetDirectShopRecommendedPriceResponseData `json:"response"` // Actual response data
}

type GetDirectShopRecommendedPriceResponseData struct {
	DirectItemPrice []DirectItemPrice `json:"direct_item_price"` // [Required]
}

type GetDiscountListRequest struct {
	DiscountStatus string `json:"discount_status"`            // [Required] The status filter for retriveing discount list. Available value: upcoming/ongoing/expired/all.
	PageNo         int64  `json:"page_no"`                    // [Required] <p>Specifies the page number of data to return in the current call. Starting from 1. if data is more than one page, the page_no can be some entry to start next call.</p>
	PageSize       int64  `json:"page_size"`                  // [Required] If many items are available to retrieve, you may need to call GetDiscountsList multiple times to retrieve all the data. Each result set is returned as a page of entries. Use the Pagination filters to control the maximum number of entries (<= 100) to retrieve per page (i.e., per call), the offset number to start next call. This integer value is used to specify the maximum number of entries to return in a single "page" of data.
	UpdateTimeFrom *int64 `json:"update_time_from,omitempty"` // [Optional] The update_time_from and update_time_to fields specify a date range for retrieving orders (based on the discount update time). The maximum date range that may be specified with the update_time_from and update_time_to fields is 30 days.
	UpdateTimeTo   *int64 `json:"update_time_to,omitempty"`   // [Optional] The update_time_from and update_time_to fields specify a date range for retrieving orders (based on the discount update time). The maximum date range that may be specified with the update_time_from and update_time_to fields is 30 days.
}

type GetDiscountListResponse struct {
	BaseResponse `json:",inline"`            // Common response fields
	Response     GetDiscountListResponseData `json:"response"` // Actual response data
}

type GetDiscountListResponseData struct {
	DiscountList []Discount `json:"discount_list"` // [Required] The discounts created in this shop.
	More         bool       `json:"more"`          // [Required] <p>This is to indicate whether the item list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of items.<br /></p>
}

type GetDiscountRequest struct {
	DiscountId int64 `json:"discount_id" url:"discount_id"` // [Required] Shopee's unique identifier for a discount activity.
	PageNo     int64 `json:"page_no" url:"page_no"`         // [Required] Specifies the page number of data to return in the current call. Starting from 1. if data is more than one page, the page_no can be some entry to start next call.
	PageSize   int64 `json:"page_size" url:"page_size"`     // [Required] Each result set is returned as a page of entries. Use the "page_size" filters to control the maximum number of entries to retrieve per page (i.e., per call), and the "page_no" to start next call. This integer value is used to specify the maximum number of entries to return in a single "page" of data.
}

type GetDiscountResponse struct {
	BaseResponse `json:",inline"`        // Common response fields
	Response     GetDiscountResponseData `json:"response"` // Actual response data
}

type GetDiscountResponseData struct {
	Status       string                        `json:"status"`        // [Required] The status of discount promotion
	DiscountName string                        `json:"discount_name"` // [Required] Title of the discount.
	ItemList     []GetDiscountResponseDataItem `json:"item_list"`     // [Required] The items selected in this discount.
	StartTime    int64                         `json:"start_time"`    // [Required] The time when discount activity start.
	DiscountId   int64                         `json:"discount_id"`   // [Required] Shopee's unique identifier for a discount activity.
	EndTime      int64                         `json:"end_time"`      // [Required] The time when discount activity end.
	More         bool                          `json:"more"`          // [Required] This is to indicate whether the item list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of items.
}

type GetDiscountResponseDataItem struct {
	ItemId                            int64                   `json:"item_id"`                                // [Required] Shopee's unique identifier for an item.
	ItemName                          string                  `json:"item_name"`                              // [Required] Name of the item in local language.
	NormalStock                       int64                   `json:"normal_stock"`                           // [Required] The current stock quantity of the item.
	ItemPromotionStock                int64                   `json:"item_promotion_stock"`                   // [Required] <p>The reserved stock of the item. If the item has no variation, this param is necessary.</p>
	ItemOriginalPrice                 float64                 `json:"item_original_price"`                    // [Required] The original price before discount of the item. If there is variation, this value is 0.
	ItemPromotionPrice                float64                 `json:"item_promotion_price"`                   // [Required] The discount price of the item. If there is variation, this value is 0.
	ItemInflatedPriceOfOriginalPrice  float64                 `json:"item_inflated_price_of_original_price"`  // [Required] <p>The original price after tax of item (Only for taxable Shop).</p>
	ItemInflatedPriceOfPromotionPrice float64                 `json:"item_inflated_price_of_promotion_price"` // [Required] <p>The discount price after tax of item (Only for taxable Shop).</p>
	ItemLocalPrice                    float64                 `json:"item_local_price"`                       // [Required] <p>The local price of item calculated as: Local Price = CB Original Price × Local Adjustment Rate.<br />Reflects the final local price derived from shop-level adjustment rules and is denominated in local currency.</p>
	ItemLocalPromotionPrice           float64                 `json:"item_local_promotion_price"`             // [Required] <p>The local discount price of item calculated as: Local Discount Price = Local Price × Discount Rate.<br />Reflects the final local seller discount price derived from setting a seller discount&nbsp;and is denominated in local currency.</p>
	ItemLocalPriceInflated            float64                 `json:"item_local_price_inflated"`              // [Required] <p>The local price after tax of item (Only for taxable Shop).</p>
	ItemLocalPromotionPriceInflated   float64                 `json:"item_local_promotion_price_inflated"`    // [Required] <p>The local discount price after tax of item (Only for taxable Shop).</p>
	ModelList                         []ResponseDataItemModel `json:"model_list"`                             // [Required] The models belong to this item.
	PurchaseLimit                     int64                   `json:"purchase_limit"`                         // [Required] The max number of this product in the promotion price.
}

type GetEscrowDetailBatchRequest struct {
	OrderSnList []string `json:"order_sn_list"` // [Required] <p>Shopee's unique identifier for an order.</p><p>limit [1,50]&nbsp;<br /></p><p>The number of recommended requests ranges from 1 to 20 orders.</p>
}

type GetEscrowDetailBatchResponse struct {
	BaseResponse `json:",inline"`                 // Common response fields
	Response     GetEscrowDetailBatchResponseData `json:"response"` // Actual response data
}

type GetEscrowDetailBatchResponseData struct {
	EscrowDetail *EscrowDetail `json:"escrow_detail"` // [Required] <p>The escrow detail for one order</p>
}

type GetEscrowDetailRequest struct {
	OrderSn string `json:"order_sn"` // [Required] Shopee's unique identifier for an order.
}

type GetEscrowDetailResponse struct {
	BaseResponse `json:",inline"`            // Common response fields
	Response     GetEscrowDetailResponseData `json:"response"` // Actual response data
}

type GetEscrowDetailResponseData struct {
	OrderSn           string            `json:"order_sn"`             // [Required]  Shopee's unique identifier for an order.
	BuyerUserName     string            `json:"buyer_user_name"`      // [Required] The username of buyer.
	ReturnOrderSnList []string          `json:"return_order_sn_list"` // [Required] The list of the serial number of return.
	OrderIncome       *OrderIncome      `json:"order_income"`         // [Required]
	BuyerPaymentInfo  *BuyerPaymentInfo `json:"buyer_payment_info"`   // [Required] <p>The buyer payment info at order checkout moment. (snapshot value)</p>
}

type GetEscrowListRequest struct {
	ReleaseTimeFrom int64  `json:"release_time_from"`   // [Required] Query start time
	ReleaseTimeTo   int64  `json:"release_time_to"`     // [Required] Query end time
	PageSize        *int64 `json:"page_size,omitempty"` // [Optional] Number of pages returned  max:100  default:40
	PageNo          *int64 `json:"page_no,omitempty"`   // [Optional] The page number  min:1  default:1
}

type GetEscrowListResponse struct {
	BaseResponse `json:",inline"`          // Common response fields
	Response     GetEscrowListResponseData `json:"response"` // Actual response data
}

type GetEscrowListResponseData struct {
	EscrowList []ResponseDataEscrow `json:"escrow_list"` // [Required] <p>The list of escrow order sn.</p>
	More       bool                 `json:"more"`        // [Required] <p>This is to indicate whether the escrow list is more than one page. If this value is true, you may want to continue to check next page to retrieve escrow orders.<br /></p>
}

type GetExpiryReportRequest struct {
	PageNo       *int64  `json:"page_no,omitempty"`        // [Optional] <p>Specifies the page number of data to return in the current call. Starting from 1. if data is more than one page, the page_no can be some entry to start next call. If empty, the default value is 1.<br /></p>
	PageSize     *int64  `json:"page_size,omitempty"`      // [Optional] <p>Each result set is returned as a page of entries. Use the "page_size" filters to control the maximum number of entries to retrieve per page (i.e., per call), and the "page_no" to start next call. This integer value is used to specify the maximum number of entries to return in a single "page" of data.</p><p>If empty, the default value is 10. The value should be between 1 and 40.</p>
	WhsIds       *string `json:"whs_ids,omitempty"`        // [Optional]
	ExpiryStatus *string `json:"expiry_status,omitempty"`  // [Optional] <p>0-Expired，2-Expiring，4-expiry_blocked，5-damaged，6-normal。Multiple selections allowed, separated by commas.<br /></p>
	CategoryIdL1 *int64  `json:"category_id_l1,omitempty"` // [Optional] <p>Only <b>Level 1 Category</b> can be filtered</p>
	SkuId        *string `json:"sku_id,omitempty"`         // [Optional]
	ItemId       *string `json:"item_id,omitempty"`        // [Optional]
	Variation    *string `json:"variation,omitempty"`      // [Optional]
	ItemName     *string `json:"item_name,omitempty"`      // [Optional]
	WhsRegion    string  `json:"whs_region"`               // [Required] <p>Num value: BR、CN、ID、MY、MX、TH、TW、PH、VN、SG<br /><br /></p><p><b>If do not pass, will get error "block by gateway due to invalid cid"</b></p>
}

type GetExpiryReportResponse struct {
	BaseResponse `json:",inline"`            // Common response fields
	Response     GetExpiryReportResponseData `json:"response"` // Actual response data
}

type GetExpiryReportResponseData struct {
	ItemList []GetExpiryReportResponseDataItem `json:"item_list"` // [Required]
}

type GetExpiryReportResponseDataItem struct {
	WarehouseItemId string    `json:"warehouse_item_id"` // [Required] <p>Warehouse item id; To indicate an unique item in a warehouse<br />one warehouse item id can match with multiple shop_item_id<br /><br />For Global Item,&nbsp;warehouse_item_id=Global Item id<br />For Local Item, shop_item_id=item_id</p>
	ItemName        string    `json:"item_name"`         // [Required]
	ItemImage       string    `json:"item_image"`        // [Required]
	SkuList         []ItemSku `json:"sku_list"`          // [Required]
}

type GetFbsInvoicesResultRequest struct {
	RequestIdList *RequestIdList `json:"request_id_list"` // [Required] <p>-</p>
}

type GetFbsInvoicesResultResponse struct {
	BaseResponse `json:",inline"`                     // Common response fields
	ErrorMsg     string                               `json:"error_msg,omitempty"`   // <p>Indicate error details if hit error. Empty if no error happened.</p>
	ResultList   []GetFbsInvoicesResultResponseResult `json:"result_list,omitempty"` //
}

type GetFbsInvoicesResultResponseResult struct {
	RequestId int64  `json:"request_id"` // [Required] <p>Represents the current status of the request</p>
	FileName  string `json:"file_name"`  // [Required] <p>Name of the file&nbsp;to be downloaded</p>
	Status    string `json:"status"`     // [Required] <p>Represents the current status of the request</p>
}

type GetFollowPrizeDetailRequest struct {
	CampaignId *int64 `json:"campaign_id,omitempty"` // [Optional] <p>The unique identifier for the created follow prize.<br /></p>
}

type GetFollowPrizeDetailResponse struct {
	BaseResponse `json:",inline"`                 // Common response fields
	Response     GetFollowPrizeDetailResponseData `json:"response"` // Actual response data
}

type GetFollowPrizeDetailResponseData struct {
	CampaignStatus  CampaignStatus `json:"campaign_status"`   // [Required] <p>The status of follow prize,the campagin status have upcoming/ongoing/expired.<br /></p>
	CampaignId      int64          `json:"campaign_id"`       // [Required] <p>The unique identifier for the created follow prize.<br /></p>
	UsageQuantity   int64          `json:"usage_quantity"`    // [Required] <p>Please enter a value between 1 and 200000.<br /></p>
	StartTime       int64          `json:"start_time"`        // [Required] <p>The timing from when the follow prize is valid,the start time later than the current time.If the start time and end time passed in by the seller overlap with other upcoming/ongoing activities, it will prompt "Another Follow Prize voucher already exists during this time period, please set another period."<br /></p>
	EndTime         int64          `json:"end_time"`          // [Required] <p>The timing until when the follow prize is still valid,the end time must be greater than the start time by at least 1 day and end time cannot exceed 3 months after start time.If the start time and end time passed in by the seller overlap with other upcoming/ongoing activities, it will prompt "Another Follow Prize voucher already exists during this time period, please set another period."<br /></p>
	MinSpend        int64          `json:"min_spend"`         // [Required] <p>The minimum spend required for using this follow prize.<br /></p>
	RewardType      int64          `json:"reward_type"`       // [Required] <p>The reward type of the follow prize.The available values are:1:discount---fix amount,2:discount---by percentage,3:coin cash back.<br /></p>
	FollowPrizeName string         `json:"follow_prize_name"` // [Required] <p>The name of the follow prize,The follow prize name length max limit is 20.<br /></p>
	DiscountAmount  int64          `json:"discount_amount"`   // [Required] <p>The discount amount set for this particular follow prize.Only fill in when you are creating a fix amount follow prize.<br /></p>
	Percentage      int64          `json:"percentage"`        // [Required] <p>The discount percentage set for this particular follow prize. Only fill in when you are creating a discount percentage follow prize or coins cashback follow prize.Discount percentage (reward_type ==2) or Percentage of coins cash back (reward_type==3).<br /></p>
	MaxPrice        int64          `json:"max_price"`         // [Required] <p>The max amount of discount/value a user can enjoy by using this particular follow prize. Only fill in when you are creating a discount percentage follow prize or coins cashback follow prize.<br /></p>
}

type GetFollowPrizeListRequest struct {
	PageNo   *int64 `json:"page_no,omitempty"`   // [Optional] <p>Specifies the page number of data to return in the current call. Default to be 1.<br /></p>
	PageSize *int64 `json:"page_size,omitempty"` // [Optional] <p>Use the 'page_size' filters to control the maximum number of entries to retrieve per page (i.e., per call). Default to be 20 and allowed input is from 1- 100.<br /></p>
	Status   string `json:"status"`              // [Required] <p>The status filter for retrieving follow prize list. Available value: upcoming/ongoing/expired/all.<br /></p>
}

type GetFollowPrizeListResponse struct {
	BaseResponse `json:",inline"`               // Common response fields
	Response     GetFollowPrizeListResponseData `json:"response"` // Actual response data
}

type GetFollowPrizeListResponseData struct {
	More            bool             `json:"more"`              // [Required] <p>This is to indicate whether the comment list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of comments.<br /></p>
	FollowPrizeList *FollowPrizeList `json:"follow_prize_list"` // [Required] <p>The list of follow prize.<br /></p>
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

type GetGmsCampaignPerformanceRequest struct {
	CampaignId *int64 `json:"campaign_id,omitempty"` // [Optional] <p>The GMS Campaign ID. Provide if available.<br /></p>
	StartDate  string `json:"start_date"`            // [Required] <p>Start date of Campaign e.g. "11-11-2025". Maximum duration of 3 months between start_date &amp; end_date. Earliest start_date is 6 months before today.<br /></p>
	EndDate    string `json:"end_date"`              // [Required] <p>End date of Campaign e.g. "11-11-2025". Maximum duration of 3 months between start_date &amp; end_date.<br /></p>
}

type GetGmsCampaignPerformanceResponse struct {
	BaseResponse `json:",inline"`                      // Common response fields
	Response     GetGmsCampaignPerformanceResponseData `json:"response"` // Actual response data
}

type GetGmsCampaignPerformanceResponseData struct {
	CampaignId int64   `json:"campaign_id"` // [Required] <p>GMS Campaign ID<br /></p>
	Report     *Report `json:"report"`      // [Required]
}

type GetGmsItemPerformanceRequest struct {
	CampaignId *int64 `json:"campaign_id,omitempty"` // [Optional] <p>The GMS Campaign ID. Provide if available.<br /></p>
	StartDate  string `json:"start_date"`            // [Required] <p>Start date of Campaign e.g. "11-11-2025". Maximum duration of 3 months between start_date &amp; end_date. Earliest start_date is 6 months before today.<br /></p>
	EndDate    string `json:"end_date"`              // [Required] <p>End date of Campaign e.g. "11-11-2025". Maximum duration of 3 months between start_date &amp; end_date.<br /></p>
	Offset     *int64 `json:"offset,omitempty"`      // [Optional] <p>Specifies the starting point, or the number of records to skip. Default is 0.</p>
	Limit      *int64 `json:"limit,omitempty"`       // [Optional] <p>Specifies the maximum number of records to show. Default is 50. Maximum is 100.<br /></p>
}

type GetGmsItemPerformanceResponse struct {
	BaseResponse `json:",inline"`                  // Common response fields
	Response     GetGmsItemPerformanceResponseData `json:"response"` // Actual response data
}

type GetGmsItemPerformanceResponseData struct {
	CampaignId  int64                                     `json:"campaign_id"`   // [Required] <p>GMS Campaign ID<br /></p>
	ResultList  []GetGmsItemPerformanceResponseDataResult `json:"result_list"`   // [Required]
	Total       int64                                     `json:"total"`         // [Required] <p>Total number of Item ID reports.<br /></p>
	HasNextPage bool                                      `json:"has_next_page"` // [Required] <p>Indicate that there are more item ID reports.<br /></p>
}

type GetGmsItemPerformanceResponseDataResult struct {
	ItemId int64         `json:"item_id"` // [Required] <p>Item ID. Results are sorted by this.<br /></p>
	Report *ResultReport `json:"report"`  // [Required]
}

type GetIncomeDetailRequest struct {
	DateFrom     string  `json:"date_from"`        // [Required] <p>Start date (YYYY-MM-DD) of the income reference period. This field is only used for Income Status = Released, the other statuses will display all records currently in that status.<br /></p><p><br /></p><p>For income Status = Released,&nbsp;For Released → Payout released date:</p><p>1. date_to must be later than date_from</p><p>2. date range cannot exceed 14 days</p><p>3. Input must follow valid date format.</p>
	DateTo       string  `json:"date_to"`          // [Required] <p>End date (YYYY-MM-DD) of the income reference period. Must be later than date_from. This field is only used for Income Status = Released, the other statuses will display all records currently in that status.<br /></p><p><br /></p><p>For income Status = Released,&nbsp;For Released → Payout released date:</p><p>1. date_to must be later than date_from</p><p>2. date range cannot exceed 14 days</p><p>3. Input must follow valid date format.</p>
	IncomeStatus int64   `json:"income_status"`    // [Required] <p>Status of Seller Income payout (Enum - Desc)</p><p><br /></p><p>Local</p><p>1 -Released</p><p>2 - Pending</p><p>CB</p><p>0 - To Release</p><p>1 -&nbsp;Released</p><p>2 - Pending</p>
	Cursor       *string `json:"cursor,omitempty"` // [Optional] <p>Pagination token for the next set of results. Use an empty string "" for the first request.</p>
	PageSize     int64   `json:"page_size"`        // [Required] <p>Number of income detail records to retrieve per page</p>
}

type GetIncomeDetailResponse struct {
	BaseResponse     `json:",inline"`  // Common response fields
	IncomeDetailList *IncomeDetailList `json:"income_detail_list,omitempty"` // <p>List of income detail records returned for the specified time range and status.</p>
}

type GetIncomeOverviewRequest struct {
	IncomeStatus *int64 `json:"income_status,omitempty" url:"income_status,omitempty"` // [Optional] <p><b><u>Status of Seller Income payout (Enum - Desc)</u></b></p><p><br /></p><p><b><u>Local Shop</u></b></p><p>1 -Released</p><p>2 - Pending</p><p><b><u><br /></u></b></p><p><b><u>CB Shop</u></b></p><p>0 - To Release</p><p>1 -&nbsp;Released</p><p>2 - Pending</p><p><br /></p><p>Note: By default, if Income Status was not provided in the request params (non mandatory), API response will return all values for all Income status based on either Local/CB</p>
}

type GetIncomeOverviewResponse struct {
	BaseResponse `json:",inline"`              // Common response fields
	Response     GetIncomeOverviewResponseData `json:"response"` // Actual response data
}

type GetIncomeOverviewResponseData struct {
	LatestPayoutDate string       `json:"latest_payout_date"` // [Required] <p>The latest payout date for the released income. Format: YYYY-MM-DD. Only for CN shops.</p>
	TotalIncome      *TotalIncome `json:"total_income"`       // [Required] <p>Object containing total income components.</p>
}

type GetIncomeReportRequest struct {
	IncomeReportId int64 `json:"income_report_id"` // [Required] <p>The identifier for income report file request.<br /></p>
}

type GetIncomeReportResponse struct {
	BaseResponse `json:",inline"`            // Common response fields
	Response     GetIncomeReportResponseData `json:"response"`      // Actual response data
	Msg          string                      `json:"msg,omitempty"` // <p>Error Message</p>
}

type GetIncomeReportResponseData struct {
	Id            int64  `json:"id"`             // [Required] <p>The identifier for income statement file request.<br /></p>
	FileName      string `json:"file_name"`      // [Required] <p>Income report file name.<br /></p>
	Status        int64  `json:"status"`         // [Required] <p>STATUS_INVALID = 0;</p><p>STATUS_PROCESSING = 1;</p><p>STATUS_DOWNLOADABLE = 2;</p><p>STATUS_DOWNLOADED = 3;</p><p>STATUS_FAILED = 4;</p>
	GeneratedTime int64  `json:"generated_time"` // [Required] <p>File generation time.<br /></p>
	FileLink      string `json:"file_link"`      // [Required] <p>Link to download income report file.<br /></p>
}

type GetIncomeStatementRequest struct {
	IncomeStatementId int64 `json:"income_statement_id" url:"income_statement_id"` // [Required] <p>The identifier for income statement file request.<br />return from the API&nbsp;<span style="font-size:14px;">v2.payment.generate_income_statement</span></p><br />
}

type GetIncomeStatementResponse struct {
	BaseResponse `json:",inline"`               // Common response fields
	Response     GetIncomeStatementResponseData `json:"response"` // Actual response data
}

type GetIncomeStatementResponseData struct {
	Id            int64  `json:"id"`             // [Required] <p>The identifier for income statement file request.<br /></p>
	FileName      string `json:"file_name"`      // [Required] <p>Income statement file name.</p>
	Status        int64  `json:"status"`         // [Required] <p>STATUS_INVALID = 0;</p><p>STATUS_PROCESSING = 1;</p><p>STATUS_DOWNLOADABLE = 2;</p><p>STATUS_DOWNLOADED = 3;</p><p>STATUS_FAILED = 4;</p>
	GeneratedTime int64  `json:"generated_time"` // [Required] <p>File generation time.</p>
	FileLink      string `json:"file_link"`      // [Required] <p>Link to download income statement file.</p>
}

type GetItemBaseInfoRequest struct {
	ItemIdList          []int64 `json:"item_id_list" url:"item_id_list"`                                       // [Required] item_id list; limit [0,50]
	NeedTaxInfo         *bool   `json:"need_tax_info,omitempty" url:"need_tax_info,omitempty"`                 // [Optional] if true will response tax_info
	NeedComplaintPolicy *bool   `json:"need_complaint_policy,omitempty" url:"need_complaint_policy,omitempty"` // [Optional] if true will response complaint_policy
}

type GetItemBaseInfoResponse struct {
	BaseResponse      `json:",inline"`            // Common response fields
	Response          GetItemBaseInfoResponseData `json:"response"`                     // Actual response data
	CertificationInfo *ResponseCertificationInfo  `json:"certification_info,omitempty"` // <p>For PH product certification input<br />Required for some category and attribute option</p>
}

type GetItemBaseInfoResponseData struct {
	ItemList        []GetItemBaseInfoResponseDataItem `json:"item_list"`        // [Required]
	ComplaintPolicy *ComplaintPolicy                  `json:"complaint_policy"` // [Required]  Complaint policy.Only returned for local PL sellers, and need_complaint_policy in request is true.
	TaxInfo         *ResponseDataTaxInfo              `json:"tax_info"`         // [Required] Tax information
	DescriptionInfo *ResponseDataDescriptionInfo      `json:"description_info"` // [Required] New description  field. Only whitelist sellers can use it.
	DescriptionType DescriptionType                   `json:"description_type"` // [Required] Type of description : values: See Data Definition- description_type (normal , extended).
	StockInfoV2     *StockInfoV2                      `json:"stock_info_v2"`    // [Required] <p>new stock object<br /></p>
}

type GetItemBaseInfoResponseDataItem struct {
	ItemId                int64              `json:"item_id"`                  // [Required] Shopee's unique identifier for an item.
	CategoryId            int64              `json:"category_id"`              // [Required] Shopee's unique identifier for a category.
	ItemName              string             `json:"item_name"`                // [Required] Name of the item in local language.
	Description           string             `json:"description"`              // [Required] if description_type is normal , Description information will be returned through this field，else description will be empty
	ItemSku               string             `json:"item_sku"`                 // [Required] An item SKU (stock keeping unit) is an identifier defined by a seller, sometimes called parent SKU. Item SKU can be assigned to an item in Shopee Listings.
	CreateTime            int64              `json:"create_time"`              // [Required] Timestamp that indicates the date and time that the item was created.
	UpdateTime            int64              `json:"update_time"`              // [Required] Timestamp that indicates the last time that there was a change in value of the item, such as price/stock change.
	AttributeList         []ItemAttribute    `json:"attribute_list"`           // [Required]
	PriceInfo             []PriceInfo        `json:"price_info"`               // [Required] If the item has models, price_info will not be returned. Please get the price of each model through the get_model_list api
	Image                 *ItemImage         `json:"image"`                    // [Required]
	Weight                string             `json:"weight"`                   // [Required] <p>The weight of this item, the unit is KG.</p><p>If set the weight of models under this item, will return the max weight of all models during the switching period to ensure system compatibility, please switch to call v2.product.get_model_list to get the weight of models.</p>
	Dimension             *Dimension         `json:"dimension"`                // [Required] <p>The dimension of this item.</p><p>If set the dimension of models under this item, will return the dimension with largest volume calculated by height*length*width during the switching period to ensure system compatibility, please switch to call v2.product.get_model_list to get the dimension of models.</p>
	LogisticInfo          []ItemLogisticInfo `json:"logistic_info"`            // [Required] The logistics list.
	PreOrder              *PreOrder          `json:"pre_order"`                // [Required]
	Wholesales            []Wholesales       `json:"wholesales"`               // [Required] The wholesales tier list.
	Condition             string             `json:"condition"`                // [Required] Is it second-hand.
	SizeChart             string             `json:"size_chart"`               // [Required] Url of size chart image.
	ItemStatus            ItemStatus         `json:"item_status"`              // [Required] <p>Enumerated type that defines the current status of the item. Applicable values: NORMAL, BANNED, UNLIST,&nbsp;<b><font color="#c24f4a" style>SELLER_DELETE, SHOPEE_DELETE, REVIEWING</font></b>.<br /></p>
	Deboost               BoolString         `json:"deboost"`                  // [Required] <p>If deboost is true, means that the item's search ranking is lowered.<br /></p>
	HasModel              bool               `json:"has_model"`                // [Required] Does it contain model.
	PromotionId           int64              `json:"promotion_id"`             // [Required] <p>The unique identifier of the promotion applied to the item.</p>
	HasPromotion          bool               `json:"has_promotion"`            // [Required] <p>Indicates whether the item is currently under any ongoing promotion.</p>
	VideoInfo             []VideoInfo        `json:"video_info"`               // [Required] Info of video list.
	Brand                 *Brand             `json:"brand"`                    // [Required]
	ItemDangerous         int64              `json:"item_dangerous"`           // [Required] This field is only applicable for local sellers in Indonesia and Malaysia. Use this field to identify whether a product is a dangerous product. 0 for non-dangerous product and 1 for dangerous product. For more information, please visit the market's respective Seller Education Hub.
	GtinCode              string             `json:"gtin_code"`                // [Required] <p>gtin code for br region, will return this code only item has default model</p><p><br /></p><p>Note: gtin_code = "00" means that this item is&nbsp;“Item without GTIN”<br /></p>
	SizeChartId           int64              `json:"size_chart_id"`            // [Required] <p>id of new size chart.<br /></p>
	PromotionImage        *ItemImage         `json:"promotion_image"`          // [Required]
	CompatibilityInfo     *CompatibilityInfo `json:"compatibility_info"`       // [Required]
	ScheduledPublishTime  int64              `json:"scheduled_publish_time"`   // [Required] <p>Scheduled publish time of this item.</p>
	AuthorisedBrandId     int64              `json:"authorised_brand_id"`      // [Required] <p>ID of authorised reseller brand.<br /></p>
	SspId                 int64              `json:"ssp_id"`                   // [Required] <p>Shopee's unique identifier for Shopee&nbsp;Standard Product.<br /></p>
	IsFulfillmentByShopee bool               `json:"is_fulfillment_by_shopee"` // [Required] <p>return true if the item only has a default model and it is FBS model</p>
	Tag                   *Tag               `json:"tag"`                      // [Required]
	PurchaseLimitInfo     *PurchaseLimitInfo `json:"purchase_limit_info"`      // [Required] <p>purchase limit info</p>
}

type GetItemContentDiagnosisResultRequest struct {
	ItemIdList []int64 `json:"item_id_list"` // [Required] <p>item_id list; limit [1,48]</p>
}

type GetItemContentDiagnosisResultResponse struct {
	BaseResponse `json:",inline"`                          // Common response fields
	Response     GetItemContentDiagnosisResultResponseData `json:"response"` // Actual response data
}

type GetItemContentDiagnosisResultResponseData struct {
	SuccessItemList []SuccessItem `json:"success_item_list"` // [Required]
	FailureItemList []Failure     `json:"failure_item_list"` // [Required]
}

type GetItemCountRequest struct {
	SessionId int64 `json:"session_id" url:"session_id"` // [Required] <p>The identifier of livestream session.</p>
}

type GetItemCountResponse struct {
	BaseResponse `json:",inline"`         // Common response fields
	Response     GetItemCountResponseData `json:"response"` // Actual response data
}

type GetItemCountResponseData struct {
	ItemCount    int64 `json:"item_count"`     // [Required] <p>The number of items in the shopping bag of this session.</p>
	MaxItemCount int64 `json:"max_item_count"` // [Required] <p>The maximum number of items allowed in the shopping bag of this session.</p>
}

type GetItemCriteriaResponse struct {
	BaseResponse `json:",inline"`            // Common response fields
	Response     GetItemCriteriaResponseData `json:"response"` // Actual response data
}

type GetItemCriteriaResponseData struct {
	Criteria                []Criteria `json:"criteria"`                   // [Required] <p>criteria detail</p>
	PairIds                 []PairIds  `json:"pair_ids"`                   // [Required] <p>the mapping relationship between criteria and category</p>
	OverlapBlockCategoryIds []int64    `json:"overlap_block_category_ids"` // [Required] <p>Due to regulations, the promotion of some products in these categories are prohibited in this region<br /></p>
}

type GetItemExtraInfoRequest struct {
	ItemIdList []int64 `json:"item_id_list" url:"item_id_list"` // [Required]  item_id list, limit [0,50]
}

type GetItemExtraInfoResponse struct {
	BaseResponse `json:",inline"`             // Common response fields
	Response     GetItemExtraInfoResponseData `json:"response"` // Actual response data
}

type GetItemExtraInfoResponseData struct {
	ItemList []GetItemExtraInfoResponseDataItem `json:"item_list"` // [Required] extra info of item list.
}

type GetItemExtraInfoResponseDataItem struct {
	ItemId       int64   `json:"item_id"`       // [Required] Shopee's unique identifier for an item.
	Sale         int64   `json:"sale"`          // [Required] The sales volume of item.
	Views        int64   `json:"views"`         // [Required] The page view of item.
	Likes        int64   `json:"likes"`         // [Required] The collection number of item.
	RatingStar   float64 `json:"rating_star"`   // [Required] The rating star scores of this item.
	CommentCount int64   `json:"comment_count"` // [Required] Count of comments for the item.
}

type GetItemInstallmentStatusRequest struct {
	ItemIdList []int64 `json:"item_id_list"` // [Required] Item id array, Max :100
}

type GetItemInstallmentStatusResponse struct {
	BaseResponse `json:",inline"`                     // Common response fields
	Response     GetItemInstallmentStatusResponseData `json:"response"` // Actual response data
}

type GetItemInstallmentStatusResponseData struct {
	ItemInstallmentList []ItemInstallment           `json:"item_installment_list"` // [Required]
	ItemPlanAhoraList   []ResponseDataItemPlanAhora `json:"item_plan_ahora_list"`  // [Required] Only applicable for local AR sellers.
}

type GetItemLimitRequest struct {
	CategoryId *int64 `json:"category_id,omitempty" url:"category_id,omitempty"` // [Optional] <p>Shopee's unique identifier for a category.<br /></p>
}

type GetItemLimitResponse struct {
	BaseResponse `json:",inline"`         // Common response fields
	Response     GetItemLimitResponseData `json:"response"`             // Actual response data
	GtinLimit    *GtinLimit               `json:"gtin_limit,omitempty"` //
}

type GetItemLimitResponseData struct {
	PriceLimit                        *PriceLimit                        `json:"price_limit"`                          // [Required]
	WholesalePriceThresholdPercentage *WholesalePriceThresholdPercentage `json:"wholesale_price_threshold_percentage"` // [Required]
	StockLimit                        *WholesalePriceThresholdPercentage `json:"stock_limit"`                          // [Required]
	ItemNameLengthLimit               *WholesalePriceThresholdPercentage `json:"item_name_length_limit"`               // [Required]
	ItemImageCountLimit               *WholesalePriceThresholdPercentage `json:"item_image_count_limit"`               // [Required]
	ItemDescriptionLengthLimit        *WholesalePriceThresholdPercentage `json:"item_description_length_limit"`        // [Required]
	TierVariationNameLengthLimit      *WholesalePriceThresholdPercentage `json:"tier_variation_name_length_limit"`     // [Required]
	TierVariationOptionLengthLimit    *WholesalePriceThresholdPercentage `json:"tier_variation_option_length_limit"`   // [Required]
	ItemCountLimit                    *ItemCountLimit                    `json:"item_count_limit"`                     // [Required]
	ExtendedDescriptionLimit          *ExtendedDescriptionLimit          `json:"extended_description_limit"`           // [Required]
	DtsLimit                          *DtsLimit                          `json:"dts_limit"`                            // [Required]
	WeightLimit                       *WeightLimit                       `json:"weight_limit"`                         // [Required]
	DimensionLimit                    *DimensionLimit                    `json:"dimension_limit"`                      // [Required]
	SizeChartLimit                    *SizeChartLimit                    `json:"size_chart_limit"`                     // [Required]
}

type GetItemListByContentDiagnosisRequest struct {
	PageSize     int64        `json:"page_size"`               // [Required] <p>the size of one page.&nbsp;Max=48<br /></p>
	Offset       *string      `json:"offset,omitempty"`        // [Optional] <p>Specifies the starting entry of data to return in the current call. Default is empty. if data is more than one page, the offset can be some entry to start next call.</p>
	QualityLevel *interface{} `json:"quality_level,omitempty"` // [Optional] <p>Item's latest content quality level. Applicable values:</p><p>1: TO_BE_IMPROVED<br />2: QUALIFIED<br />3: EXCELLENT</p>
	IssueType    *interface{} `json:"issue_type,omitempty"`    // [Optional] <p>Item's content issue. Applicable values:&nbsp;</p><p>1: TOO_FEW_IMAGES<br />2: WRONG_CATEGORY<br />3: TOO_FEW_ATTRIBUTES_FOR_QUALIFIED<br />4: LACK_OF_SIZE_CHART<br />5: LACK_OF_STANDARD_VARIATION<br />6: LACK_BRAND<br />7: TOO_SHORT_DESCRIPTION<br />8: TOO_SHORT_OR_TOO_LONG_NAME<br />9: WRONG_WEIGHT<br />10: LACK_OF_VIDEO<br />11: TOO_FEW_ATTRIBUTES_FOR_EXCELLENT</p><p><br /></p><p>If you need to pass both quality_level and issue_type, the logic are as follows:<br />- When quality_level is 1, issue_type can only be 1, 2, 3, 4, 5<br />- When quality_level is 2, issue_type can only be 6, 7, 8, 9, 10, 11<br />- When quality_level is 3, issue_type can only be empty</p>
}

type GetItemListByContentDiagnosisResponse struct {
	BaseResponse `json:",inline"`                          // Common response fields
	Response     GetItemListByContentDiagnosisResponseData `json:"response"` // Actual response data
}

type GetItemListByContentDiagnosisResponseData struct {
	ItemList    []SuccessItem `json:"item_list"`     // [Required]
	TotalCount  int64         `json:"total_count"`   // [Required] <p>Total num of items match condition.<br /></p>
	HasNextPage bool          `json:"has_next_page"` // [Required] <p>This is to indicate whether the item list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of items.<br /></p>
	NextOffset  string        `json:"next_offset"`   // [Required] <p>If has_next_page is true, this value need set to next request.offset<br /></p>
}

type GetItemListRequest struct {
	Offset         int64        `json:"offset" url:"offset"`                                         // [Required] Specifies the starting entry of data to return in the current call. Default is 0. if data is more than one page, the offset can be some entry to start next call.
	PageSize       int64        `json:"page_size" url:"page_size"`                                   // [Required] the size of one page.Max=100
	UpdateTimeFrom *int64       `json:"update_time_from,omitempty" url:"update_time_from,omitempty"` // [Optional]  The update_time_from and update_time_to fields specify a date range for retrieving orders (based on the item update time). The update_time_from field is the starting date range.
	UpdateTimeTo   *int64       `json:"update_time_to,omitempty" url:"update_time_to,omitempty"`     // [Optional] The update_time_from and update_time_to fields specify a date range for retrieving orders (based on the item update time). The update_time_to field is the ending date range
	ItemStatus     []ItemStatus `json:"item_status" url:"item_status"`                               // [Required] <p>NORMAL/BANNED/UNLIST/<b><font color="#c24f4a">REVIEWING/SELLER_DELETE/SHOPEE_DELETE</font></b></p><p>If you want to search multiple status, please upload the url like this: item_status=NORMAL&amp;item_status=BANNED</p>
}

type GetItemListResponse struct {
	BaseResponse `json:",inline"`        // Common response fields
	Response     GetItemListResponseData `json:"response"` // Actual response data
}

type GetItemListResponseData struct {
	Item        []ResponseDataItem `json:"item"`          // [Required] list of item info with item_id/ item_status/ update_time
	TotalCount  int64              `json:"total_count"`   // [Required] total count of all items
	HasNextPage bool               `json:"has_next_page"` // [Required] This is to indicate whether the item list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of items.
	NextOffset  int64              `json:"next_offset"`   // [Required] if has_next_page is true, this value need set to next request.offset
}

type GetItemListResponseDataList struct {
	ItemNo        int64          `json:"item_no"`        // [Required] <p>The order of this item in the shopping bag of current session, start from 1.</p>
	ItemId        int64          `json:"item_id"`        // [Required] <p>Shopee's unique identifier for an item.</p>
	Name          string         `json:"name"`           // [Required] <p>Name of the item in local language.<br /></p>
	ImageUrl      string         `json:"image_url"`      // [Required] <p>The image url of this item.</p>
	PriceInfo     *ListPriceInfo `json:"price_info"`     // [Required]
	AffiliateInfo *AffiliateInfo `json:"affiliate_info"` // [Required]
}

type GetItemPromotionRequest struct {
	ItemIdList []int64 `json:"item_id_list" url:"item_id_list"` // [Required] Item ID list, can send 1 to 50 items.
}

type GetItemPromotionResponse struct {
	BaseResponse `json:",inline"`             // Common response fields
	Response     GetItemPromotionResponseData `json:"response"` // Actual response data
}

type GetItemPromotionResponseData struct {
	SuccessList []ResponseDataSuccess `json:"success_list"` // [Required] Success item promotion info.
	FailureList []Failure             `json:"failure_list"` // [Required] Fail item promotion info.
}

type GetItemSetItemListRequest struct {
	ItemSetId int64 `json:"item_set_id" url:"item_set_id"` // [Required] <p>The identifier of the item set.</p>
	Offset    int64 `json:"offset" url:"offset"`           // [Required] <p>Specifies the starting entry of data to return in the current call. Default is 0, if data is more than one page, the offset can be some entry to start next call.</p>
	PageSize  int64 `json:"page_size" url:"page_size"`     // [Required] <p>Each result set is returned as a page of entries. Use the "page_size" filters to control the maximum number of entries to retrieve per page (i.e., per call). This integer value is used to specify the maximum number of entries to return in a single "page" of data. The limit of page_size if between 1 and 100.</p>
}

type GetItemSetItemListResponse struct {
	BaseResponse `json:",inline"`               // Common response fields
	Response     GetItemSetItemListResponseData `json:"response"` // Actual response data
}

type GetItemSetItemListResponseData struct {
	More       bool                              `json:"more"`        // [Required] <p>This is to indicate whether the list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of data.</p>
	NextOffset int64                             `json:"next_offset"` // [Required] <p>If more is true, this value need set to next request offset.</p>
	List       []GetLikeItemListResponseDataList `json:"list"`        // [Required]
}

type GetItemSetListRequest struct {
	Offset   int64   `json:"offset" url:"offset"`                       // [Required] <p>Specifies the starting entry of data to return in the current call. Default is 0, if data is more than one page, the offset can be some entry to start next call.</p>
	PageSize int64   `json:"page_size" url:"page_size"`                 // [Required] <p>Each result set is returned as a page of entries. Use the "page_size" filters to control the maximum number of entries to retrieve per page (i.e., per call). This integer value is used to specify the maximum number of entries to return in a single "page" of data. The limit of page_size if between 1 and 100.</p>
	Keyword  *string `json:"keyword,omitempty" url:"keyword,omitempty"` // [Optional] <p>Search the item set with it's name matching the keyword.</p>
}

type GetItemSetListResponse struct {
	BaseResponse `json:",inline"`           // Common response fields
	Response     GetItemSetListResponseData `json:"response"` // Actual response data
}

type GetItemSetListResponseData struct {
	More       bool                             `json:"more"`        // [Required] <p>This is to indicate whether the list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of data.</p>
	NextOffset int64                            `json:"next_offset"` // [Required] <p>If more is true, this value need set to next request offset.</p>
	List       []GetItemSetListResponseDataList `json:"list"`        // [Required]
}

type GetItemSetListResponseDataList struct {
	ItemSetId   int64  `json:"item_set_id"`   // [Required] <p>The identifier of the item set.</p>
	ItemSetName string `json:"item_set_name"` // [Required] <p>The name of the item set.</p>
	ItemCount   int64  `json:"item_count"`    // [Required] <p>The number of items in this item set.</p>
}

type GetItemViolationInfoRequest struct {
	ItemIdList []int64 `json:"item_id_list"` // [Required] <p>item_id list; limit [0,50]</p>
}

type GetItemViolationInfoResponse struct {
	BaseResponse `json:",inline"`                 // Common response fields
	Response     GetItemViolationInfoResponseData `json:"response"` // Actual response data
}

type GetItemViolationInfoResponseData struct {
	ItemList []GetItemViolationInfoResponseDataItem `json:"item_list"` // [Required]
}

type GetItemViolationInfoResponseDataItem struct {
	ItemId            int64               `json:"item_id"`             // [Required] <p>Shopee's unique identifier for an item.<br /></p>
	ItemName          string              `json:"item_name"`           // [Required] <p>Name of the item.<br /></p>
	ItemStatus        ItemStatus          `json:"item_status"`         // [Required] <p>Enumerated type that defines the current status of the item. Applicable values: NORMAL, BANNED, UNLIST, SELLER_DELETE, SHOPEE_DELETE, REVIEWING.<br /></p>
	Deboost           bool                `json:"deboost"`             // [Required] <p>If deboost is true, means that the item's search ranking is lowered.<br /></p>
	ItemStatusDetails []ItemStatusDetails `json:"item_status_details"` // [Required]
	DeboostDetails    []DeboostDetails    `json:"deboost_details"`     // [Required]
	FailError         string              `json:"fail_error"`          // [Required] <p>Indicate error type if one element hit error.<br /></p>
	FailMessage       string              `json:"fail_message"`        // [Required] <p>Indicate error details if one element hit error.<br /></p>
}

type GetKitItemInfoRequest struct {
	ItemId int64 `json:"item_id" url:"item_id"` // [Required] <p>ID of kit item.</p>
}

type GetKitItemInfoResponse struct {
	BaseResponse `json:",inline"`           // Common response fields
	Response     GetKitItemInfoResponseData `json:"response"` // Actual response data
}

type GetKitItemInfoResponseData struct {
	ProductInfo *ProductInfo `json:"product_info"` // [Required]
}

type GetKitItemLimitRequest struct {
	CategoryId *int64 `json:"category_id,omitempty" url:"category_id,omitempty"` // [Optional] <p>Shopee's unique identifier for a category.<br /></p>
}

type GetKitItemLimitResponse struct {
	BaseResponse `json:",inline"`            // Common response fields
	Response     GetKitItemLimitResponseData `json:"response"` // Actual response data
}

type GetKitItemLimitResponseData struct {
	PriceLimit                       *PriceLimit                        `json:"price_limit"`                           // [Required]
	ItemNameLengthLimit              *WholesalePriceThresholdPercentage `json:"item_name_length_limit"`                // [Required]
	ItemImageCountLimit              *WholesalePriceThresholdPercentage `json:"item_image_count_limit"`                // [Required]
	DescriptionLimit                 *DescriptionLimit                  `json:"description_limit"`                     // [Required]
	TierVariationNameLengthLimit     *WholesalePriceThresholdPercentage `json:"tier_variation_name_length_limit"`      // [Required]
	TierVariationOptionLengthLimit   *WholesalePriceThresholdPercentage `json:"tier_variation_option_length_limit"`    // [Required]
	WeightLimit                      *WeightLimit                       `json:"weight_limit"`                          // [Required]
	DimensionLimit                   *DimensionLimit                    `json:"dimension_limit"`                       // [Required]
	DtsLimit                         *ResponseDataDtsLimit              `json:"dts_limit"`                             // [Required]
	ComponentCountLimitOfSingleModel *WholesalePriceThresholdPercentage `json:"component_count_limit_of_single_model"` // [Required]
}

type GetLateOrdersRequest struct {
	PageNo   *int64 `json:"page_no,omitempty" url:"page_no,omitempty"`     // [Optional] <p>Specifies the page number of data to return in the current call. Starting from 1. if data is more than one page, the page_no can be some entry to start next call.&nbsp;</p><p><br /></p><p>Default is 1.</p>
	PageSize *int64 `json:"page_size,omitempty" url:"page_size,omitempty"` // [Optional] <p>Each result set is returned as a page of entries. Use the "page_size" filters to control the maximum number of entries to retrieve per page (i.e., per call), and the "page_no" to start next call. This integer value is used to specify the maximum number of entries to return in a single "page" of data.&nbsp;</p><p><br /></p><p>The limit of page_size if between 1 and 100. Default is 10.</p>
}

type GetLateOrdersResponse struct {
	BaseResponse `json:",inline"`          // Common response fields
	Response     GetLateOrdersResponseData `json:"response"` // Actual response data
}

type GetLateOrdersResponseData struct {
	LateOrderList []LateOrder `json:"late_order_list"` // [Required] <p>Late Orders.</p>
	TotalCount    int64       `json:"total_count"`     // [Required] <p>Total number of late orders.<br /></p>
}

type GetLatestCommentListRequest struct {
	SessionId int64  `json:"session_id" url:"session_id"`             // [Required] <p>The identifier of livestream session.</p>
	Offset    *int64 `json:"offset,omitempty" url:"offset,omitempty"` // [Optional] <p>Specifies the starting entry of data to return in the current call. Default is 0, if data is more than one page, the offset can be some entry to start next call.</p>
}

type GetLatestCommentListResponse struct {
	BaseResponse `json:",inline"`                 // Common response fields
	Response     GetLatestCommentListResponseData `json:"response"` // Actual response data
}

type GetLatestCommentListResponseData struct {
	NextOffset int64                                  `json:"next_offset"` // [Required] <p>The offset for next page request.</p>
	List       []GetLatestCommentListResponseDataList `json:"list"`        // [Required]
}

type GetLatestCommentListResponseDataList struct {
	CommentId int64  `json:"comment_id"` // [Required] <p>The identifier of comment.</p>
	Content   string `json:"content"`    // [Required] <p>The content of comment.</p>
	UserId    int64  `json:"user_id"`    // [Required] <p>The user id of the one who posted the comment.</p>
	Username  string `json:"username"`   // [Required] <p>The username of the one who posted comment.</p>
}

type GetLikeItemListRequest struct {
	Offset   int64   `json:"offset" url:"offset"`                       // [Required] <p>Specifies the starting entry of data to return in the current call. Default is 0, if data is more than one page, the offset can be some entry to start next call.</p>
	PageSize int64   `json:"page_size" url:"page_size"`                 // [Required] <p>Each result set is returned as a page of entries. Use the "page_size" filters to control the maximum number of entries to retrieve per page (i.e., per call). This integer value is used to specify the maximum number of entries to return in a single "page" of data. The limit of page_size if between 1 and 100.</p>
	Keyword  *string `json:"keyword,omitempty" url:"keyword,omitempty"` // [Optional] <p>Search items with name matching this keyword.</p>
}

type GetLikeItemListResponse struct {
	BaseResponse `json:",inline"`            // Common response fields
	Response     GetLikeItemListResponseData `json:"response"` // Actual response data
}

type GetLikeItemListResponseData struct {
	More       bool                              `json:"more"`        // [Required] <p>This is to indicate whether the list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of data.</p>
	NextOffset int64                             `json:"next_offset"` // [Required] <p>If more is true, this value need set to next request offset.</p>
	List       []GetLikeItemListResponseDataList `json:"list"`        // [Required]
}

type GetLikeItemListResponseDataList struct {
	ItemId        int64          `json:"item_id"`        // [Required] <p>Shopee's unique identifier for an item.</p>
	Name          string         `json:"name"`           // [Required] <p>Name of the item in local language.<br /></p>
	ImageUrl      string         `json:"image_url"`      // [Required] <p>The image url of this item.</p>
	PriceInfo     *ListPriceInfo `json:"price_info"`     // [Required]
	AffiliateInfo *AffiliateInfo `json:"affiliate_info"` // [Required]
}

type GetListingsWithIssuesRequest struct {
	PageNo   *int64 `json:"page_no,omitempty" url:"page_no,omitempty"`     // [Optional] <p>Specifies the page number of data to return in the current call. Starting from 1. if data is more than one page, the page_no can be some entry to start next call.&nbsp;</p><p><br /></p><p>Default is 1.</p>
	PageSize *int64 `json:"page_size,omitempty" url:"page_size,omitempty"` // [Optional] <p>Each result set is returned as a page of entries. Use the "page_size" filters to control the maximum number of entries to retrieve per page (i.e., per call), and the "page_no" to start next call. This integer value is used to specify the maximum number of entries to return in a single "page" of data.&nbsp;</p><p><br /></p><p>The limit of page_size if between 1 and 100. Default is 10.</p>
}

type GetListingsWithIssuesResponse struct {
	BaseResponse `json:",inline"`                  // Common response fields
	Response     GetListingsWithIssuesResponseData `json:"response"` // Actual response data
}

type GetListingsWithIssuesResponseData struct {
	ListingList []Listing `json:"listing_list"` // [Required] <p>Listing with issues.</p>
	TotalCount  int64     `json:"total_count"`  // [Required] <p>Total number of listing with issues.</p>
}

type GetLocalAdjustmentRateResponse struct {
	BaseResponse `json:",inline"`                   // Common response fields
	Response     GetLocalAdjustmentRateResponseData `json:"response"` // Actual response data
}

type GetLocalAdjustmentRateResponseData struct {
	LocalAdjustmentRate float64 `json:"local_adjustment_rate"` // [Required] <p>The multiplier used to adjust the cross-border original price to local price</p>
}

type GetLostPushMessageResponse struct {
	BaseResponse `json:",inline"`               // Common response fields
	Response     GetLostPushMessageResponseData `json:"response"` // Actual response data
}

type GetLostPushMessageResponseData struct {
	PushMessageList []PushMessage `json:"push_message_list"` // [Required] <p>Returns the earliest 100 lost push messages that were lost within 3 days of the current time and not confirmed to have been consumed.</p>
	HasNextPage     bool          `json:"has_next_page"`     // [Required] <p>This is to indicate whether the lost push message to be consumed is more than 100. If this value is true, you may need to continue calling to get the remaining lost push messages to be consumed.<br /></p>
	LastMessageId   int64         `json:"last_message_id"`   // [Required] <p>Specifies the end entry of data returned in the current call.<br /></p>
}

type GetMainItemListRequest struct {
	DirectItemId []int64 `json:"direct_item_id" url:"direct_item_id"` // [Required] <p>Item id of direct shop.</p>
}

type GetMainItemListResponse struct {
	BaseResponse `json:",inline"`            // Common response fields
	Response     GetMainItemListResponseData `json:"response"` // Actual response data
}

type GetMainItemListResponseData struct {
	List []ResponseDataList `json:"list"` // [Required]
}

type GetMartItemMappingByIdRequest struct {
	MartItemId       int64   `json:"mart_item_id"`        // [Required] <p>The item ID of the item in the Mart shop.</p>
	OutletShopIdList []int64 `json:"outlet_shop_id_list"` // [Required] <p>A list of outlet shop IDs used to filter the mapping results.</p>
}

type GetMartItemMappingByIdResponse struct {
	BaseResponse `json:",inline"`                   // Common response fields
	Response     GetMartItemMappingByIdResponseData `json:"response"` // Actual response data
}

type GetMartItemMappingByIdResponseData struct {
	ItemMappingList []ItemMapping `json:"item_mapping_list"` // [Required] <p>A list of item mapping records between the Mart item and its corresponding outlet items.</p>
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

type GetMerchantInfoResponse struct {
	BaseResponse     `json:",inline"` // Common response fields
	MerchantName     string           `json:"merchant_name,omitempty"`     // Name of the merchant.
	AuthTime         int64            `json:"auth_time,omitempty"`         // The timestamp when the merchant was authorized to the partner.
	ExpireTime       int64            `json:"expire_time,omitempty"`       // Use this field to indicate the expiration date for merchant authorization.
	MerchantCurrency string           `json:"merchant_currency,omitempty"` // <p>The three-digit code representing the currency unit used for the item in this merchant. If currency haven't been setting in CNSC/KRSC, it will be empty.</p><p>China merchant support CNY and USD currently.</p><p>Korea merchant support KRW and USD currently.&nbsp;</p><p>Hong kong merchant support USD currently, will support HKD later.</p>
	MerchantRegion   string           `json:"merchant_region,omitempty"`   // <p>Region of the merchant. Including KR, HK and CN.<br /></p>
	IsUpgradedCbsc   bool             `json:"is_upgraded_cbsc,omitempty"`  // <p>Use this filed to indicate whether this merchant is upgraded to cbsc.<br /></p>
}

type GetMerchantPrepaidAccountListRequest struct {
	PageNo   int64 `json:"page_no" url:"page_no"`     // [Required] <p>Specifies the page number of data to return in the current call. Starting from 1. if data is more than one page, the page_no can be some entry to start next call.</p>
	PageSize int64 `json:"page_size" url:"page_size"` // [Required] <p>Each result set is returned as a page of entries. Use the "page_size" filters to control the maximum number of entries to retrieve per page (i.e., per call), and the "page_no" to start next call. This integer value is used to specify the maximum number of entries to return in a single "page" of data.<br /></p><p>Min: 1 Max:10<br /></p>
}

type GetMerchantPrepaidAccountListResponse struct {
	BaseResponse `json:",inline"`                          // Common response fields
	Response     GetMerchantPrepaidAccountListResponseData `json:"response"` // Actual response data
}

type GetMerchantPrepaidAccountListResponseData struct {
	Total int64                                           `json:"total"` // [Required]
	List  []GetMerchantPrepaidAccountListResponseDataList `json:"list"`  // [Required]
	More  bool                                            `json:"more"`  // [Required] <p>This is to indicate whether the list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of datas.<br /></p>
}

type GetMerchantPrepaidAccountListResponseDataList struct {
	PrepaidAccountId            int64  `json:"prepaid_account_id"`             // [Required] <p>Record ID</p>
	PrepaidAccountCourierKey    string `json:"prepaid_account_courier_key"`    // [Required] <p>Courier Company Key (快递公司编码).</p>
	PrepaidAccountCourierName   string `json:"prepaid_account_courier_name"`   // [Required] <p>Courier Company Name (快递公司名称).</p>
	PrepaidAccountPartnerId     string `json:"prepaid_account_partner_id"`     // [Required] <p>Prepaid Account Number (电子面单账户号码).</p>
	PrepaidAccountPartnerKey    string `json:"prepaid_account_partner_key"`    // [Required] <p>Prepaid Account Password (电子面单账户密码).</p>
	PrepaidAccountPartnerSecret string `json:"prepaid_account_partner_secret"` // [Required] <p>Partner Secret (电子面单密钥).</p>
	PrepaidAccountPartnerName   string `json:"prepaid_account_partner_name"`   // [Required] <p>Partner Name (电子面单客户账户名称).</p>
	PrepaidAccountPartnerNet    string `json:"prepaid_account_partner_net"`    // [Required] <p>Branch Name (网点名称).</p>
	PrepaidAccountPartnerCode   string `json:"prepaid_account_partner_code"`   // [Required] <p>Partner Code (电子面单承载编号).</p>
	PrepaidAccountCheckMan      string `json:"prepaid_account_check_man"`      // [Required] <p>Delivery Agent Name (电子面单承载快递员名).</p>
	PrepaidAccountIsDefault     bool   `json:"prepaid_account_is_default"`     // [Required] <p>This is to indicate whether the prepaid account is Default Prepaid Account.</p>
}

type GetMerchantWarehouseListRequest struct {
	Cursor        *Cursor `json:"cursor"`         // [Required] <p>// how to use DoubleSidedCursor</p><p>// Get data for the first page: Please pass next_id = 0 or nil, page_size = {your page size}.</p><p>// Get data for the next page: Please pass the Cursor from the previous response, and set prev_id=nil;</p><p>// Get data for the prev page: Please pass the Cursor from the previous response, and set next_id=nil;</p><p>// Stop fetching next data: The Cursor.next_id in the previous response is nil.</p><p>// Stop fetching prev data: The Cursor.prev_id in the previous response is nil.<br /></p>
	WarehouseType int64   `json:"warehouse_type"` // [Required] <p>1 means pickup warehouse</p><p>2 means return warehouse<br /></p>
}

type GetMerchantWarehouseListResponse struct {
	BaseResponse `json:",inline"`                     // Common response fields
	Response     GetMerchantWarehouseListResponseData `json:"response"` // Actual response data
}

type GetMerchantWarehouseListResponseData struct {
	TotalCount    int64       `json:"total_count"`    // [Required] <p>Total count of all warehouses.<br /></p>
	WarehouseList []Warehouse `json:"warehouse_list"` // [Required]
	Cursor        *Cursor     `json:"cursor"`         // [Required]
}

type GetMerchantWarehouseLocationListResponse struct {
	BaseResponse `json:",inline"`                             // Common response fields
	Response     GetMerchantWarehouseLocationListResponseData `json:"response"` // Actual response data
}

type GetMerchantWarehouseLocationListResponseData struct {
	LocationId    string `json:"location_id"`    // [Required] <p>Location identifier for stocks. Different location_ids represent that your addresses are in different item stocks<br /></p>
	WarehouseName string `json:"warehouse_name"` // [Required] <p>The warehouse name filled in when creating the warehouse address<br /></p>
}

type GetMerchantsByPartnerRequest struct {
	PageSize *int64 `json:"page_size,omitempty" url:"page_size,omitempty"` // [Optional] Each result set is returned as a page of entries. Use the "page_size" filters to control the maximum number of entries to retrieve per page (i.e., per call), and the "page_no" to start next call. This integer value is used to specify the maximum number of entries to return in a single "page" of data.
	PageNo   *int64 `json:"page_no,omitempty" url:"page_no,omitempty"`     // [Optional] Specifies the page number of data to return in the current call. Starting from 1. if data is more than one page, the page_no can be some entry to start next call.
}

type GetMerchantsByPartnerResponse struct {
	BaseResponse       `json:",inline"` // Common response fields
	AuthedMerchantList []AuthedMerchant `json:"authed_merchant_list,omitempty"` // A list of merchants authorized to the partner.
	More               bool             `json:"more,omitempty"`                 // This is to indicate whether the list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of datas.
}

type GetMetricSourceDetailRequest struct {
	MetricId int64  `json:"metric_id" url:"metric_id"`                     // [Required] <p>ID of metric. Supported values:&nbsp;</p><p>1: Late Shipment Rate (All Channels)</p><p>3:&nbsp;Non-Fulfilment Rate (All Channels)</p><p>4:&nbsp;Preparation Time</p><p>12:&nbsp;Pre-order Listing %</p><p>15:&nbsp;Days of Pre-order Listing Violation</p><p>25:&nbsp;Fast Handover Rate</p><p>28:&nbsp;On-time Pickup Failure Rate Violation Value<br /></p><p>42:&nbsp;Cancellation Rate (All Channels)</p><p>43:&nbsp;Return-refund Rate (All Channels)</p><p>52:&nbsp;Severe Listing Violations</p><p>53:&nbsp;Other Listing Violations</p><p>85: Late Shipment Rate (NDD)<br /></p><p>88:&nbsp;Non-fulfilment Rate (NDD<br /></p><p>91:&nbsp;Cancellation Rate (NDD)<br /></p><p>92:&nbsp;Return-refund Rate (NDD)</p><p>96: % SDD Listings</p><p>97: % NDD Listings</p><p>2001: Fast Handover Rate - SLS</p><p>2002: Fast Handover Rate - FBS<br /></p><p>2003: Fast Handover Rate - 3PF</p><p>2030: % HD Listings</p><p>2031: % HD Free Shipping Enabled</p><p>2032:&nbsp;Saturday Shipment</p><p>2033:&nbsp;Preparation Time PS</p>
	PageNo   *int64 `json:"page_no,omitempty" url:"page_no,omitempty"`     // [Optional] <p>Specifies the page number of data to return in the current call. Starting from 1. if data is more than one page, the page_no can be some entry to start next call.&nbsp;</p><p><br /></p><p>Default is 1.<br /></p>
	PageSize *int64 `json:"page_size,omitempty" url:"page_size,omitempty"` // [Optional] <p>Each result set is returned as a page of entries. Use the "page_size" filters to control the maximum number of entries to retrieve per page (i.e., per call), and the "page_no" to start next call. This integer value is used to specify the maximum number of entries to return in a single "page" of data.&nbsp;</p><p><br /></p><p>The limit of page_size if between 1 and 100. Default is 10.</p>
}

type GetMetricSourceDetailResponse struct {
	BaseResponse `json:",inline"`                  // Common response fields
	Response     GetMetricSourceDetailResponseData `json:"response"` // Actual response data
}

type GetMetricSourceDetailResponseData struct {
	MetricId                         int64                          `json:"metric_id"`                             // [Required] <p>ID of metric.</p>
	NfrOrderList                     []NfrOrder                     `json:"nfr_order_list"`                        // [Required] <p>Affected Orders for Non-fulfilment Rate.</p><p><br /></p><p>Supported metric_id:&nbsp;</p><p>3:&nbsp;Non-Fulfilment Rate (All Channels)</p><p>88:&nbsp;Non-fulfilment Rate (NDD)</p>
	CancellationOrderList            []CancellationOrder            `json:"cancellation_order_list"`               // [Required] <p>Affected Orders for Cancellation Rate.&nbsp;</p><p><br /></p><p>Supported metric_id:&nbsp;</p><p>42:&nbsp;Cancellation Rate (All Channels)</p><p>91:&nbsp;Cancellation Rate (NDD)</p>
	ReturnRefundOrderList            []ReturnRefundOrder            `json:"return_refund_order_list"`              // [Required] <p>Affected Orders for Return-refund Rate.</p><p><br /></p><p>Supported metric_id:&nbsp;</p><p>43:&nbsp;Return-refund Rate (All Channels)</p><p>92:&nbsp;Return-refund Rate (NDD)</p>
	LsrOrderList                     []LsrOrder                     `json:"lsr_order_list"`                        // [Required] <p>Affected Orders for Late Shipment Rate.</p><p><br /></p><p>Supported metric_id:&nbsp;</p><p>1:&nbsp;Late Shipment Rate (All Channels)</p><p>85:&nbsp;Late Shipment Rate (NDD)</p>
	FhrOrderList                     []FhrOrder                     `json:"fhr_order_list"`                        // [Required] <p>Affected Orders for Fast Handover Rate.</p><p><br /></p><p>Supported metric_id:&nbsp;</p><p>25:&nbsp;Fast Handover Rate</p><p>2001: Fast Handover Rate - SLS</p><p>2002: Fast Handover Rate - FBS<br /></p><p>2003: Fast Handover Rate - 3PF</p>
	OpfrDayDetailDataList            []OpfrDayDetailData            `json:"opfr_day_detail_data_list"`             // [Required] <p>Relevant Violations for OPFR Violation Value.</p><p><br /></p><p>Supported metric_id:&nbsp;</p><p>28:&nbsp;On-time Pickup Failure Rate Violation Value</p>
	ViolationListingList             []ViolationListing             `json:"violation_listing_list"`                // [Required] <p>Relevant Listings for Severe Listing Violations and Other Listing Violations.</p><p><br /></p><p>Supported metric_id:&nbsp;</p><p>52:&nbsp;Severe Listing Violations</p><p>53:&nbsp;Other Listing Violations</p>
	PreOrderListingViolationDataList []PreOrderListingViolationData `json:"pre_order_listing_violation_data_list"` // [Required] <p>Relevant Listings for Days of Pre-order Listing Violation.</p><p><br /></p><p>Supported metric_id:&nbsp;</p><p>15:&nbsp;Days of Pre-order Listing Violation</p>
	PreOrderListingList              []PreOrderListing              `json:"pre_order_listing_list"`                // [Required] <p>Relevant Listings for Pre-order Listing.</p><p><br /></p><p>Supported metric_id:&nbsp;</p><p>12:&nbsp;Pre-order Listing %</p>
	SddListingList                   []SddListing                   `json:"sdd_listing_list"`                      // [Required] <p>Relevant Listings for % SDD Listings.</p><p><br /></p><p>Supported metric_id:&nbsp;</p><p>96: % SDD Listings.</p>
	NddListingList                   []NddListing                   `json:"ndd_listing_list"`                      // [Required] <p>Relevant Listings for % NDD Listings.</p><p><br /></p><p>Supported metric_id:&nbsp;</p><p>97: % NDD Listings.</p>
	AptOrderList                     []AptOrder                     `json:"apt_order_list"`                        // [Required] <p>Affected Parcels for Preparation Time.</p><p><br /></p><p>Supported metric_id:&nbsp;</p><p>4:&nbsp;Preparation Time</p>
	HdListingList                    *HdListingList                 `json:"hd_listing_list"`                       // [Required] <p>Relevant Listings for % HD Listings and % HD Free Shipping Enabled.</p><p><br /></p><p>Supported metric_id:&nbsp;</p><p>2030: % HD Listings</p><p>2031: % HD Free Shipping Enabled</p>
	SaturdayShipmentList             []AptOrder                     `json:"saturday_shipment_list"`                // [Required] <p>Affected Parcels for Saturday Shipment</p><p><br />Supported metric_id:<br />2032: Saturday Shipment</p>
	TotalCount                       int64                          `json:"total_count"`                           // [Required] <p>Total number of Affected Orders or Relevant Listings.</p>
}

type GetModelListRequest struct {
	ItemId int64 `json:"item_id" url:"item_id"` // [Required] The ID of the item
}

type GetModelListResponse struct {
	BaseResponse `json:",inline"`         // Common response fields
	Response     GetModelListResponseData `json:"response"` // Actual response data
}

type GetModelListResponseData struct {
	TierVariation            []GetModelListResponseDataTierVariation `json:"tier_variation"`             // [Required] Variation config of item.
	Model                    []GetModelListResponseDataModel         `json:"model"`                      // [Required] Model list.
	StandardiseTierVariation []ResponseDataStandardiseTierVariation  `json:"standardise_tier_variation"` // [Required] <p>Standardise Variation config of item.<br /></p>
}

type GetModelListResponseDataModel struct {
	PriceInfo             []ResponseDataModelPriceInfo `json:"price_info"`               // [Required] <p>Price info.</p><p>For&nbsp;<b><font color="#c24f4a">SG/MY/BR/MX/PL/ES/AR</font></b> seller:&nbsp;Sellers can set the price with two decimal place,&nbsp;other regions can only set the price as an integer.<br /></p>
	ModelId               int64                        `json:"model_id"`                 // [Required] Model ID.
	TierIndex             interface{}                  `json:"tier_index"`               // [Required] Tier index of this model.
	PromotionId           int64                        `json:"promotion_id"`             // [Required] Current promotion ID of this model.
	HasPromotion          bool                         `json:"has_promotion"`            // [Required] <p>Indicates whether the model is currently under any ongoing promotion.</p><p>  </p><p><br /></p>
	ModelSku              string                       `json:"model_sku"`                // [Required] SKU of this model. the length should be under 100.
	ModelStatus           string                       `json:"model_status"`             // [Required] <p>The model status. Should be&nbsp;MODEL_NORMAL or&nbsp;MODEL_UNAVAILABLE.&nbsp;MODEL_NORMAL models can be sold on the buyer's side, and MODEL_UNAVAILABLE models cannot be sold on the buyer's side.</p>
	PreOrder              *PreOrder                    `json:"pre_order"`                // [Required] (Only whitelisted users can use)
	StockInfoV2           *ModelStockInfoV2            `json:"stock_info_v2"`            // [Required] <p>new stock info.<br /></p><p>Please check this FAQ for more detail:&nbsp;<a href="https://open.shopee.com/faq?top=162&amp;sub=166&amp;page=1&amp;faq=230" target="_blank" style="font-size:14px;">https://open.shopee.com/faq?top=162&amp;sub=166&amp;page=1&amp;faq=230</a></p>
	GtinCode              string                       `json:"gtin_code"`                // [Required] <p><b><font color="#c24f4a">(Only TW seller and BR local seller available)</font></b> gtin code.</p>
	Weight                string                       `json:"weight"`                   // [Required] <p>The weight of this model, the unit is KG.</p><p>If don't set the weight of this model, will use the weight of item by default.</p>
	Dimension             *Dimension                   `json:"dimension"`                // [Required] <p>The dimension of this model.</p><p>If don't set the dimension of this model, will use the dimension of item by default.</p>
	IsFulfillmentByShopee bool                         `json:"is_fulfillment_by_shopee"` // [Required] <p>whether model is fulfillment by shopee</p>
}

type GetModelListResponseDataTierVariation struct {
	Name       string                            `json:"name"`        // [Required] Tier name.
	OptionList []ResponseDataTierVariationOption `json:"option_list"` // [Required] Tier option list for corresponding tier name.
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

type GetOrderDetailRequest struct {
	OrderSnList               string  `json:"order_sn_list" url:"order_sn_list"`                                                   // [Required] The set of order_sn. If there are multiple order_sn, you need to use English comma to connect them. limit [1,50]
	RequestOrderStatusPending *bool   `json:"request_order_status_pending,omitempty" url:"request_order_status_pending,omitempty"` // [Optional] <p>Compatible parameter during migration period, send True will let API support PENDING status and return&nbsp; pending_terms, send False or don’t send will fallback to old logic</p>
	ResponseOptionalFields    *string `json:"response_optional_fields,omitempty" url:"response_optional_fields,omitempty"`         // [Optional] <p>a response fields you want to get. Please select from the below response parameters. If you input an object field, all the params under it will be included automatically in the response. If there are multiple response fields you want to get, you need to use English comma to connect them.  Available values: buyer_user_id,buyer_username,estimated_shipping_fee,recipient_address,actual_shipping_fee ,goods_to_declare,note,note_update_time,item_list,pay_time,dropshipper, dropshipper_phone,split_up,buyer_cancel_reason,cancel_by,cancel_reason,actual_shipping_fee_confirmed,buyer_cpf_id,fulfillment_flag,pickup_done_time,package_list,shipping_carrier,payment_method,total_amount,buyer_username,invoice_data,order_chargeable_weight_gram,return_request_due_date,edt,payment_info,international_label</p>
}

type GetOrderDetailResponse struct {
	BaseResponse `json:",inline"`           // Common response fields
	Response     GetOrderDetailResponseData `json:"response"` // Actual response data
}

type GetOrderDetailResponseData struct {
	OrderList []GetOrderDetailResponseDataOrder `json:"order_list"` // [Required] The list of orders.
}

type GetOrderDetailResponseDataOrder struct {
	OrderSn                    string            `json:"order_sn"`                      // [Required] Return by default. Shopee's unique identifier for an order.
	Region                     string            `json:"region"`                        // [Required] Return by default. The two-digit code representing the region where the order was made.
	Currency                   string            `json:"currency"`                      // [Required] Return by default. The three-digit code representing the currency unit for which the order was paid.
	Cod                        bool              `json:"cod"`                           // [Required] Return by default. This value indicates whether the order was a COD (cash on delivery) order.
	TotalAmount                float64           `json:"total_amount"`                  // [Required] The total amount paid by the buyer for the order. This amount includes the total sale price of items, shipping cost beared by buyer; and offset by Shopee promotions if applicable. This value will only return after the buyer has completed payment for the order.
	PendingTerms               []string          `json:"pending_terms"`                 // [Required] <p>The list of pending terms. Applicable values:</p><p>- SYSTEM_PENDING: Under Shopee internal processing.</p><p>- KYC_PENDING: Under KYC checking (TW CB order only).</p><p>- ARRANGE_SHIPMENT_PENDING: Temporarily held due to 3PL capacity constraints.</p>
	PendingDescription         []string          `json:"pending_description"`           // [Required] <p>The value of this field is the description of pending reason corresponding with pending terms. Applicable values:&nbsp;</p><p>- For SYSTEM_PENDING: Order is being processed by Shopee.</p><p>- For KYC_PENDING: Order is pending buyer TW KYC pre-authorization.</p><p>- For ARRANGE_SHIPMENT_PENDING: Allocating delivery resources due to high order volume. Label print will be available within 4 days after buyer paid.</p>
	OrderStatus                OrderStatus       `json:"order_status"`                  // [Required] Return by default. Enumerated type that defines the current status of the order.
	ShippingCarrier            string            `json:"shipping_carrier"`              // [Required] <p>The logistics service provider that the buyer selected for the order to deliver items.&nbsp;</p><p><br /></p><p>Note: If logistics_channel_id is 90021, 90025 or 90026, service_code will be appended,&nbsp;e.g., Entrega Turbo - M1020.</p>
	PaymentMethod              string            `json:"payment_method"`                // [Required] The payment method that the buyer selected to pay for the order. Applicable values: See Data Definition- Payment Methods.
	EstimatedShippingFee       float64           `json:"estimated_shipping_fee"`        // [Required] The estimated shipping fee is an estimation calculated by Shopee based on specific logistics courier's standard.
	MessageToSeller            string            `json:"message_to_seller"`             // [Required] Return by default. Message to seller.
	CreateTime                 int64             `json:"create_time"`                   // [Required] Return by default. Timestamp that indicates the date and time that the order was created.
	UpdateTime                 int64             `json:"update_time"`                   // [Required] Return by default. Timestamp that indicates the last time that there was a change in value of order, such as order status changed from 'Paid' to 'Completed'.
	DaysToShip                 int64             `json:"days_to_ship"`                  // [Required] Return by default. Shipping preparation time set by the seller when listing item on Shopee.
	ShipByDate                 int64             `json:"ship_by_date"`                  // [Required] Return by default. The deadline to ship out the parcel.
	BuyerUserId                int64             `json:"buyer_user_id"`                 // [Required] <p>The user id of buyer of this order, will be empty if it is a non-integrated order in TW region.</p>
	BuyerUsername              string            `json:"buyer_username"`                // [Required] <p>The name of buyer, will be masked as "****" if it is a non-integrated order in TW region.</p>
	RecipientAddress           *RecipientAddress `json:"recipient_address"`             // [Required] <p>This object contains detailed breakdown for the recipient address.<br />Different parameters might be masked according to each market and kind of seller.<br /><br />For TW region integrated channel orders will be all masked as "****". More details may refer the announcement.<br /></p>
	ActualShippingFee          float64           `json:"actual_shipping_fee"`           // [Required] The actual shipping fee of the order if available from external logistics partners.
	GoodsToDeclare             bool              `json:"goods_to_declare"`              // [Required] Only work for cross-border order.This value indicates whether the order contains goods that are required to declare at customs. "T" means true and it will mark as "T" on the shipping label; "F" means false and it will mark as "P" on the shipping label. This value is accurate ONLY AFTER the order trackingNo is generated, please capture this value AFTER your retrieve the trackingNo.
	Note                       string            `json:"note"`                          // [Required] The note seller made for own reference.
	NoteUpdateTime             int64             `json:"note_update_time"`              // [Required] Update time for the note.
	ItemList                   []OrderItem       `json:"item_list"`                     // [Required] This object contains the detailed breakdown for the result of this API call.
	PayTime                    int64             `json:"pay_time"`                      // [Required] The time when the order status is updated from UNPAID to PAID. This value is NULL when order is not paid yet.
	Dropshipper                string            `json:"dropshipper"`                   // [Required] For Indonesia orders only. The name of the dropshipper.
	DropshipperPhone           string            `json:"dropshipper_phone"`             // [Required] The phone number of dropshipper, could be empty.
	SplitUp                    bool              `json:"split_up"`                      // [Required] To indicate whether this order is split to fullfil order(forder) level. Call GetForderInfo if it's "true".
	BuyerCancelReason          string            `json:"buyer_cancel_reason"`           // [Required] Cancel reason from buyer, could be empty.
	CancelBy                   string            `json:"cancel_by"`                     // [Required] Could be one of buyer, seller, system or Ops.
	CancelReason               string            `json:"cancel_reason"`                 // [Required] Use this field to get reason for buyer, seller, and system cancellation.
	ActualShippingFeeConfirmed bool              `json:"actual_shipping_fee_confirmed"` // [Required] Use this filed to judge whether the actual_shipping_fee is confirmed.
	BuyerCpfId                 string            `json:"buyer_cpf_id"`                  // [Required] Buyer's CPF number for taxation and invoice purposes. Only for Brazil order.
	FulfillmentFlag            string            `json:"fulfillment_flag"`              // [Required] Use this field to indicate the order is fulfilled by shopee or seller. Applicable values: fulfilled_by_shopee, fulfilled_by_cb_seller, fulfilled_by_local_seller.
	PickupDoneTime             int64             `json:"pickup_done_time"`              // [Required] The timestamp when pickup is done.
	PackageList                []OrderPackage    `json:"package_list"`                  // [Required] The list of package under an order
	InvoiceData                *InvoiceData      `json:"invoice_data"`                  // [Required] The invoice data of the order.
	CheckoutShippingCarrier    string            `json:"checkout_shipping_carrier"`     // [Required] For non masking order, the logistics service provider that the buyer selected for the order to deliver items.  For masking order, the logistics service type that the buyer selected for the order to deliver items.
	ReverseShippingFee         float64           `json:"reverse_shipping_fee"`          // [Required] Shopee charges the reverse shipping fee for the returned order.The value of this field will be non-negative.
	OrderChargeableWeightGram  int64             `json:"order_chargeable_weight_gram"`  // [Required] display weight used to calculate ASF for this order
	PrescriptionImages         []string          `json:"prescription_images"`           // [Required] <p>Return prescription images of this order, only for ID and PH whitelist sellers.</p><p><br /></p><p>Please add the prefix to review:</p><p>for ID:&nbsp;<a href="https://cf.shopee.co.id/file/+prescription_image" target="_blank" style="font-size:14px;">https://cf.shopee.co.id/file/+prescription_image</a></p><p>for PH:<a href="https://cf.shopee.co.id/file/+prescription_image" target="_blank" style="font-size:14px;">https://cf.shopee.ph/file/+prescription_image</a></p>
	PrescriptionCheckStatus    int64             `json:"prescription_check_status"`     // [Required] enum OrderPrescriptionCheckStatus: NONE = 0; PASSED = 1; FAILED = 2; Only for ID and PH whitelist sellers.
	PharmacistName             string            `json:"pharmacist_name"`               // [Required] <p>Name of the Pharmacist for Prescription Order.</p>
	PrescriptionApprovalTime   int64             `json:"prescription_approval_time"`    // [Required] <p>Time of when the prescription is approved.</p>
	PrescriptionRejectionTime  int64             `json:"prescription_rejection_time"`   // [Required] <p>Time of when the prescription is rejected.</p>
	EdtFrom                    int64             `json:"edt_from"`                      // [Required] <p>Earliest estimated delivery date of orders (only available for BR region)<br /></p>
	EdtTo                      int64             `json:"edt_to"`                        // [Required] <p>Latest estimated delivery time of orders (only available for BR region)<br /></p>
	BookingSn                  string            `json:"booking_sn"`                    // [Required] <p>Return by default. Shopee's unique identifier for a booking.</p><p>Only returned for advance fulfilment matched order only.</p>
	AdvancePackage             bool              `json:"advance_package"`               // [Required] <p>Indicate whether order will be fulfilled using advance fulfilment stock or not. If value is true, order will be matched with a booking and seller should not arrange shipment.</p>
	ReturnRequestDueDate       int64             `json:"return_request_due_date"`       // [Required] <p>This field represents the deadline for buyers to initiate returns and refunds after order is completed.</p><p><br /></p><p>The “return_request_due_date” response parameter will be returned if the requested order meets&nbsp;<b>ALL&nbsp;the conditions</b>&nbsp;below:</p><p>- The status of the order is COMPLETED</p><p>- The return refund eligibility of the order is true</p><p><br /></p><p>If you have any questions related to the function of "returns and refunds after order is completed," please refer to the following link:&nbsp;https://seller.shopee.tw/edu/article/18474</p>
	PaymentInfo                []PaymentInfo     `json:"payment_info"`                  // [Required] <p>[Only for BR] List of payment information, to follow&nbsp;<a href="https://drive.google.com/file/d/1VfqlbmXr3XR6BkpKOPUbLCgjiqvsBbLd/view?usp=sharing" target="_blank">NT 2025.001</a>&nbsp;(BR government invoice rules).</p>
	IsBuyerShopCollection      bool              `json:"is_buyer_shop_collection"`      // [Required] <p>To indicate if this order is buyer self collection at store order</p>
	BuyerProofOfCollection     []string          `json:"buyer_proof_of_collection"`     // [Required] <p>The image url of the buyer self collection at the store.</p>
	HotListingOrder            bool              `json:"hot_listing_order"`             // [Required] <p>[Only for PH,TH,VN,MY,BR,TW] True if the order includes hot listing item.</p>
	IsInternational            bool              `json:"is_international"`              // [Required] <p>[Only for BR] Indicate if the order is SIP order. This field will only be returned if international_label is included in response_optional_field in the request.</p>
}

type GetOrderListRequest struct {
	TimeRangeField            string       `json:"time_range_field" url:"time_range_field"`                                             // [Required] The kind of time_from and time_to. Available value: create_time, update_time.
	TimeFrom                  int64        `json:"time_from" url:"time_from"`                                                           // [Required] The time_from and time_to fields specify a date range for retrieving orders (based on the time_range_field). The time_from field is the starting date range. The maximum date range that may be specified with the time_from and time_to fields is 15 days.
	TimeTo                    int64        `json:"time_to" url:"time_to"`                                                               // [Required] The time_from and time_to fields specify a date range for retrieving orders (based on the time_range_field). The time_from field is the starting date range. The maximum date range that may be specified with the time_from and time_to fields is 15 days.
	PageSize                  int64        `json:"page_size" url:"page_size"`                                                           // [Required] Each result set is returned as a page of entries. Use the "page_size" filters to control the maximum number of entries to retrieve per page (i.e., per call). This integer value is used to specify the maximum number of entries to return in a single "page" of data.The limit of page_size if between 1 and 100.
	Cursor                    *string      `json:"cursor,omitempty" url:"cursor,omitempty"`                                             // [Optional] <p>Specifies the starting entry of data to return in the current call. The default is empty. If the data is more than one page, the offset can be some entry to start the next call.</p>
	OrderStatus               *OrderStatus `json:"order_status,omitempty" url:"order_status,omitempty"`                                 // [Optional] The order_status filter for retriveing orders and each one only every request. Available value: UNPAID/READY_TO_SHIP/PROCESSED/SHIPPED/COMPLETED/IN_CANCEL/CANCELLED/INVOICE_PENDING
	ResponseOptionalFields    *string      `json:"response_optional_fields,omitempty" url:"response_optional_fields,omitempty"`         // [Optional] Optional fields in response. Available value: order_status.
	RequestOrderStatusPending *bool        `json:"request_order_status_pending,omitempty" url:"request_order_status_pending,omitempty"` // [Optional] <p>Compatible parameter during migration period, send True will let API support PENDING status, send False or don’t send will fallback to old logic.<br /></p>
	LogisticsChannelId        *int64       `json:"logistics_channel_id,omitempty" url:"logistics_channel_id,omitempty"`                 // [Optional] <p>The identity of logistic channel. Valid only for BR.<br /></p>
}

type GetOrderListResponse struct {
	BaseResponse `json:",inline"`         // Common response fields
	Response     GetOrderListResponseData `json:"response"` // Actual response data
}

type GetOrderListResponseData struct {
	More       bool                `json:"more"`        // [Required] This is to indicate whether the order list is more than one page. If this value is true, you may want to continue to check next page to retrieve orders.
	OrderList  []ResponseDataOrder `json:"order_list"`  // [Required]
	NextCursor string              `json:"next_cursor"` // [Required] If  more is true, you should pass the next_cursor in the next request as cursor. The value of next_cursor will be empty string when more is false.
}

type GetPackageDetailRequest struct {
	PackageNumberList string `json:"package_number_list" url:"package_number_list"` // [Required] <p>The set of package_number. If there are multiple package_number, you need to use English comma to connect them. limit [1,50]</p>
}

type GetPackageDetailResponse struct {
	BaseResponse `json:",inline"`             // Common response fields
	Response     GetPackageDetailResponseData `json:"response"` // Actual response data
}

type GetPackageDetailResponseData struct {
	PackageList []ResponseDataPackage `json:"package_list"` // [Required] <p>The list of packages.</p>
}

type GetPaymentMethodListResponse struct {
	BaseResponse `json:",inline"`                 // Common response fields
	Response     GetPaymentMethodListResponseData `json:"response"` // Actual response data
}

type GetPaymentMethodListResponseData struct {
	PaymentMethod []string `json:"payment_method"` // [Required]
	Region        string   `json:"region"`         // [Required]
}

type GetPayoutDetailRequest struct {
	PageSize       int64 `json:"page_size"`        // [Required] Number of pages returned  max:100
	PageNo         int64 `json:"page_no"`          // [Required] The page number  min:1  default:1
	PayoutTimeFrom int64 `json:"payout_time_from"` // [Required] <p>Strat time. Maximum time range is 15 days</p>
	PayoutTimeTo   int64 `json:"payout_time_to"`   // [Required] End time
}

type GetPayoutDetailResponse struct {
	BaseResponse `json:",inline"`            // Common response fields
	Response     GetPayoutDetailResponseData `json:"response"` // Actual response data
}

type GetPayoutDetailResponseData struct {
	More       bool     `json:"more"`        // [Required]
	PayoutList []Payout `json:"payout_list"` // [Required]
}

type GetPayoutInfoRequest struct {
	PayoutTimeFrom int64  `json:"payout_time_from"` // [Required] <p>Start time. Maximum time range is 15 days<br /></p>
	PayoutTimeTo   int64  `json:"payout_time_to"`   // [Required] <p>Payout End time<br /></p>
	PageSize       int64  `json:"page_size"`        // [Required] <p>Number of pages returned max:100<br /></p>
	Cursor         string `json:"cursor"`           // [Required] <p>Specifies the starting entry of data to return in the current call. Default is "". If data is more than one page, the offset can be some entry to start next call.<br /></p>
}

type GetPayoutInfoResponse struct {
	BaseResponse `json:",inline"`          // Common response fields
	Response     GetPayoutInfoResponseData `json:"response"` // Actual response data
}

type GetPayoutInfoResponseData struct {
	PayoutList *PayoutList `json:"payout_list"` // [Required]
	More       bool        `json:"more"`        // [Required] <p>True or False<br /></p>
	NextCursor string      `json:"next_cursor"` // [Required] <p>used for next batch data query. will return empty when all data been returned<br /></p>
}

type GetPenaltyPointHistoryRequest struct {
	PageNo        *int64 `json:"page_no,omitempty" url:"page_no,omitempty"`               // [Optional] <p>Specifies the page number of data to return in the current call. Starting from 1. if data is more than one page, the page_no can be some entry to start next call.&nbsp;</p><p><br /></p><p>Default is 1.</p>
	PageSize      *int64 `json:"page_size,omitempty" url:"page_size,omitempty"`           // [Optional] <p>Each result set is returned as a page of entries. Use the "page_size" filters to control the maximum number of entries to retrieve per page (i.e., per call), and the "page_no" to start next call. This integer value is used to specify the maximum number of entries to return in a single "page" of data.&nbsp;</p><p><br /></p><p>The limit of page_size if between 1 and 100. Default is 10.</p>
	ViolationType *int64 `json:"violation_type,omitempty" url:"violation_type,omitempty"` // [Optional] <p>Applicable values:&nbsp;</p><p>5: High Late Shipment Rate</p><p>6: High Non-fulfilment Rate</p><p>7: High number of non-fulfilled orders</p><p>8: High number of late shipped orders</p><p>9: Prohibited Listings</p><p>10: Counterfeit / IP infringement</p><p>11: Spam</p><p>12: Copy/Steal images</p><p>13: Re-uploading deleted listings with no change</p><p>14: Bought counterfeit from mall</p><p>15: Counterfeit caught by Shopee</p><p>16: High percentage of pre-order listings</p><p>17: Confirmed Fraud attempts (total)</p><p>18: Confirmed Fraud attempts per week (All with vouchers only)</p><p>19: Fake return address</p><p>20: Shipping fraud/abuse</p><p>21: High No. of Non-responded Chat</p><p>22: Rude chat replies</p><p>23: Request buyer to cancel order</p><p>24: Rude reply to buyer's review</p><p>25: Violate Return/Refund policy</p><p>101: Tier Reason</p><p>3026: Misuse of Shopee’s IP</p><p>3028: Violate Shop Name Regulations</p><p>3030: Direct transactions outside of the Shopee platform</p><p>3032: Shipping empty / incomplete parcels</p><p>3034: Severe Violations on Shopee Feed</p><p>3036: Severe Violations on Shopee LIVE</p><p>3038: Misuse of Local Vendor Tag</p><p>3040: Use of misleading shop tag in listing image</p><p>3042: Counterfeit / IP Infringement test</p><p>3044: Repeat Offender - IP infringement and Counterfeit listings</p><p>3046: Violation of Live Animals Selling Policy</p><p>3048: Chat Spam</p><p>3050: High Overseas Return Refunds Rate</p><p>3052: Privacy breach in buyer's review reply</p><p>3054: Order Brushing</p><p>3056: porn image</p><p>3058: Incorrect Product Categories</p><p>3060: Extremely High Non-Fulfilment Rate</p><p>3062: Penalty of Affiliate Marketing Solution (AMS) Overdue Invoice Payment</p><p>3064: Government-related listing</p><p>3066: Listing invalid gifted items</p><p>3068: High non-fulfilment rate (Next Day Delivery Orders)</p><p>3070: High Late Shipment Rate (Next Day Delivery Orders)</p><p>3072: OPFR Violation Value</p><p>3074: Direct transactions outside Shopee platform via chat</p><p>3090: Prohibited Listings-Extreme Violations</p><p>3091: Prohibited Listings-High Violations</p><p>3092: Prohibited Listings-Mid Violations</p><p>3093: Prohibited Listings-Low Violations</p><p>3094: Counterfeit Listings-Extreme Violations</p><p>3095: Counterfeit Listings-High Violations</p><p>3096: Counterfeit Listings-Mid Violations</p><p>3097: Counterfeit Listings-Low Violations</p><p>3098: Spam Listings-Extreme Violations</p><p>3099: Spam Listings-High Violations</p><p>3100: Spam Listings-Mid Violations</p><p>3101: Spam Listings-Low Violations</p><p>3145: Return/Refund Rate (Non-integrated Channel)</p><p>4130: Poor Product Quality</p>
}

type GetPenaltyPointHistoryResponse struct {
	BaseResponse `json:",inline"`                   // Common response fields
	Response     GetPenaltyPointHistoryResponseData `json:"response"` // Actual response data
}

type GetPenaltyPointHistoryResponseData struct {
	PenaltyPointList []PenaltyPoint `json:"penalty_point_list"` // [Required] <p>The penalty point records generated in the current quarter.<br /></p>
	TotalCount       int64          `json:"total_count"`        // [Required] <p>Total number of penalty point records.</p>
}

type GetPendingBuyerInvoiceOrderListRequest struct {
	Cursor   *string `json:"cursor,omitempty" url:"cursor,omitempty"` // [Optional] Specifies the starting entry of data to return in the current call. Default is "". If data is more than one page, the offset can be some entry to start next call.
	PageSize int64   `json:"page_size" url:"page_size"`               // [Required] Each result set is returned as a page of entries. Use the "page_size" filters to control the maximum number of entries to retrieve per page (i.e., per call). This integer value is used to specify the maximum number of entries to return in a single "page" of data.The limit of page_size if between 1 and 100.
}

type GetPendingBuyerInvoiceOrderListResponse struct {
	BaseResponse `json:",inline"`                            // Common response fields
	Response     GetPendingBuyerInvoiceOrderListResponseData `json:"response"` // Actual response data
}

type GetPendingBuyerInvoiceOrderListResponseData struct {
	More       bool      `json:"more"`        // [Required] This is to indicate whether the order list is more than one page. If this value is true, you may want to continue to check next page to retrieve orders.
	NextCursor string    `json:"next_cursor"` // [Required] If more is true, you should pass the next_cursor in the next request as cursor. The value of next_cursor will be empty string when more is false.
	OrderList  []Queries `json:"order_list"`  // [Required]
}

type GetProductCampaignDailyPerformanceRequest struct {
	StartDate      string `json:"start_date" url:"start_date"`             // [Required] <p>This is the parameter to indicate the start date of the time length of performance.<br /></p>
	EndDate        string `json:"end_date" url:"end_date"`                 // [Required] <p>This is the parameter to indicate the end date of the time length of performance<br /></p>
	CampaignIdList string `json:"campaign_id_list" url:"campaign_id_list"` // [Required] <p>The campaign ids (comma separated) you want to fetch the performance. (max 100)<br /></p>
}

type GetProductCampaignDailyPerformanceResponse struct {
	BaseResponse `json:",inline"`                               // Common response fields
	Response     GetProductCampaignDailyPerformanceResponseData `json:"response"` // Actual response data
}

type GetProductCampaignDailyPerformanceResponseData struct {
	Region       string     `json:"region"`        // [Required] <p>the region where each shop is under</p>
	CampaignList []Campaign `json:"campaign_list"` // [Required] <p>the list of campaign</p>
}

type GetProductCampaignHourlyPerformanceRequest struct {
	PerformanceDate string `json:"performance_date" url:"performance_date"` // [Required] <p>This is the parameter to indicate the start date of the time length of performance.<br /></p>
	CampaignIdList  string `json:"campaign_id_list" url:"campaign_id_list"` // [Required] <p>The campaign ids (comma separated) you want to fetch the performance. (max 100)<br /></p>
}

type GetProductCampaignHourlyPerformanceResponse struct {
	BaseResponse `json:",inline"`                                // Common response fields
	Response     GetProductCampaignHourlyPerformanceResponseData `json:"response"` // Actual response data
}

type GetProductCampaignHourlyPerformanceResponseData struct {
	Region       string                 `json:"region"`        // [Required] <p>The region where this Shop is under<br /></p>
	CampaignList []ResponseDataCampaign `json:"campaign_list"` // [Required] <p>the list of campaign</p>
}

type GetProductCertificationRuleRequest struct {
	AttributeList []Attribute `json:"attribute_list,omitempty"` // [Optional] Item attributes.
	CategoryId    *int64      `json:"category_id,omitempty"`    // [Optional] ID of category.
}

type GetProductCertificationRuleResponse struct {
	BaseResponse `json:",inline"`                        // Common response fields
	Response     GetProductCertificationRuleResponseData `json:"response"` // Actual response data
}

type GetProductCertificationRuleResponseData struct {
	CertificationRuleList []CertificationRule `json:"certification_rule_list"` // [Required] New description field. Only whitelist sellers can use it. If you use the field, please upload the description_type=extended otherwise api will return error. If you don't use this field, you don't need to upload the description_type or upload description_type=normal
}

type GetProductLevelCampaignIdListRequest struct {
	AdType *string `json:"ad_type,omitempty" url:"ad_type,omitempty"` // [Optional] <p>Any of ["","all","auto","manual"]</p>
	Offset *int64  `json:"offset,omitempty" url:"offset,omitempty"`   // [Optional] <p>offset</p>
	Limit  *int64  `json:"limit,omitempty" url:"limit,omitempty"`     // [Optional] <p>limit</p>
}

type GetProductLevelCampaignIdListResponse struct {
	BaseResponse `json:",inline"`                          // Common response fields
	Response     GetProductLevelCampaignIdListResponseData `json:"response"` // Actual response data
}

type GetProductLevelCampaignIdListResponseData struct {
	Region       string                                              `json:"region"`        // [Required] <p>Region the shop belongs to</p>
	HasNextPage  bool                                                `json:"has_next_page"` // [Required] <p>there are more campaigns on next page</p>
	CampaignList []GetProductLevelCampaignIdListResponseDataCampaign `json:"campaign_list"` // [Required] <p>the list of campaigns</p>
}

type GetProductLevelCampaignIdListResponseDataCampaign struct {
	AdType     string `json:"ad_type"`     // [Required] <p>auto/manual</p>
	CampaignId int64  `json:"campaign_id"` // [Required] <p>the unique id per campaign</p>
}

type GetProductLevelCampaignSettingInfoRequest struct {
	InfoTypeList   string `json:"info_type_list" url:"info_type_list"`     // [Required] <p>Info type values: 1.Common Info 2.Manual Bidding Info 3.Auto Bidding Info 4.Auto Product Ads Info</p>
	CampaignIdList string `json:"campaign_id_list" url:"campaign_id_list"` // [Required] <p>list of campaign ids comma separated (max 100 campaign ids)</p>
}

type GetProductLevelCampaignSettingInfoResponse struct {
	BaseResponse `json:",inline"`                               // Common response fields
	Response     GetProductLevelCampaignSettingInfoResponseData `json:"response"` // Actual response data
}

type GetProductLevelCampaignSettingInfoResponseData struct {
	Region       string                                                   `json:"region"`        // [Required] <p>Region the shop belongs to</p>
	CampaignList []GetProductLevelCampaignSettingInfoResponseDataCampaign `json:"campaign_list"` // [Required] <p>-</p>
}

type GetProductLevelCampaignSettingInfoResponseDataCampaign struct {
	CampaignId         int64                `json:"campaign_id"`           // [Required] <p>The unique ID per campaign</p>
	CommonInfo         *CommonInfo          `json:"common_info"`           // [Required] <p>common_info body</p>
	ManualBiddingInfo  *ManualBiddingInfo   `json:"manual_bidding_info"`   // [Required] <p>manual bidding info</p>
	AutoBiddingInfo    *AutoBiddingInfo     `json:"auto_bidding_info"`     // [Required] <p>bidding info</p>
	AutoProductAdsInfo []AutoProductAdsInfo `json:"auto_product_ads_info"` // [Required] <p>selected products info</p>
}

type GetProductRecommendedRoiTargetRequest struct {
	ReferenceId string `json:"reference_id" url:"reference_id"` // [Required] <p>A random string used to prevent duplicate ads. If an ads is created successfully, subsequent requests using the same reference id will fail - in this case, a new one must be generated.<br /><br />Use the same string for calling suggestion/recommendation API before the actual request to create an ads.<br /><br /><br /></p>
	ItemId      int64  `json:"item_id" url:"item_id"`           // [Required] <p>Unique identifier for a product.</p>
}

type GetProductRecommendedRoiTargetResponse struct {
	BaseResponse `json:",inline"`                           // Common response fields
	Response     GetProductRecommendedRoiTargetResponseData `json:"response"` // Actual response data
}

type GetProductRecommendedRoiTargetResponseData struct {
	LowerBound *LowerBound `json:"lower_bound"` // [Required] <pre><span style="font-family:Roboto, &quot;Helvetica Neue&quot;, Helvetica, &quot;Droid Sans&quot;, Arial, sans-serif;">Lower bound recommendation.  e.g., value=3.5 and percentile=80 mean that setting an ROI target of 3.5 </span></pre><pre><span style="font-family:Roboto, &quot;Helvetica Neue&quot;, Helvetica, &quot;Droid Sans&quot;, Arial, sans-serif;">makes the ads more competitive than 80% of similar ads.</span></pre>
	Exact      *LowerBound `json:"exact"`       // [Required] <p><span style="font-size:14px;">Mid-level recommendation&nbsp;</span></p><p><span style="font-size:14px;">e.g., value=5.9 and percentile=50 mean that setting an ROI target of 5.9 makes the ads more competitive than 50% of similar ads.</span></p>
	UpperBound *LowerBound `json:"upper_bound"` // [Required] <p>Higher bound recommendation.</p><p>e.g., value=10.8 and percentile=20 mean that setting an ROI target of 10.8 makes the ads more competitive than 20% of similar ads.</p>
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

type GetPunishmentHistoryRequest struct {
	PageNo           *int64 `json:"page_no,omitempty" url:"page_no,omitempty"`     // [Optional] <p>Specifies the page number of data to return in the current call. Starting from 1. if data is more than one page, the page_no can be some entry to start next call.&nbsp;</p><p><br /></p><p>Default is 1.</p>
	PageSize         *int64 `json:"page_size,omitempty" url:"page_size,omitempty"` // [Optional] <p>Each result set is returned as a page of entries. Use the "page_size" filters to control the maximum number of entries to retrieve per page (i.e., per call), and the "page_no" to start next call. This integer value is used to specify the maximum number of entries to return in a single "page" of data.&nbsp;</p><p><br /></p><p>The limit of page_size if between 1 and 100. Default is 10.</p>
	PunishmentStatus int64  `json:"punishment_status" url:"punishment_status"`     // [Required] <p>The status of punishment. Applicable values:&nbsp;</p><p>1: Ongoing</p><p>2: Ended<br /></p>
}

type GetPunishmentHistoryResponse struct {
	BaseResponse `json:",inline"`                 // Common response fields
	Response     GetPunishmentHistoryResponseData `json:"response"` // Actual response data
}

type GetPunishmentHistoryResponseData struct {
	PunishmentList []Punishment `json:"punishment_list"` // [Required] <p>The punishment records generated in the current quarter.<br /></p>
	TotalCount     int64        `json:"total_count"`     // [Required] <p>Total number of punishment records.<br /></p>
}

type GetRecentItemListRequest struct {
	Offset   int64 `json:"offset" url:"offset"`       // [Required] <p>Specifies the starting entry of data to return in the current call. Default is 0, if data is more than one page, the offset can be some entry to start next call.</p>
	PageSize int64 `json:"page_size" url:"page_size"` // [Required] <p>Each result set is returned as a page of entries. Use the "page_size" filters to control the maximum number of entries to retrieve per page (i.e., per call). This integer value is used to specify the maximum number of entries to return in a single "page" of data. The limit of page_size if between 1 and 100.</p>
}

type GetRecentItemListResponse struct {
	BaseResponse `json:",inline"`              // Common response fields
	Response     GetRecentItemListResponseData `json:"response"` // Actual response data
}

type GetRecentItemListResponseData struct {
	More       bool                              `json:"more"`        // [Required] <p>This is to indicate whether the list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of data.</p>
	NextOffset int64                             `json:"next_offset"` // [Required] <p>If more is true, this value need set to next request offset.</p>
	List       []GetLikeItemListResponseDataList `json:"list"`        // [Required]
}

type GetRecommendAttributeRequest struct {
	ItemName     string `json:"item_name" url:"item_name"`                               // [Required] name of item
	CoverImageId *int64 `json:"cover_image_id,omitempty" url:"cover_image_id,omitempty"` // [Optional] Cover image id of item
	CategoryId   int64  `json:"category_id" url:"category_id"`                           // [Required] ID of category
}

type GetRecommendAttributeResponse struct {
	BaseResponse `json:",inline"`                  // Common response fields
	Response     GetRecommendAttributeResponseData `json:"response"` // Actual response data
}

type GetRecommendAttributeResponseData struct {
	AttributeList []ResponseDataAttribute `json:"attribute_list"` // [Required] Attribute info list.
}

type GetRecommendedItemListResponse struct {
	BaseResponse `json:",inline"`                   // Common response fields
	Response     GetRecommendedItemListResponseData `json:"response"` // Actual response data
}

type GetRecommendedItemListResponseData struct {
	ItemId            int64    `json:"item_id"`              // [Required] <p>Recommended SKU's item id<br /></p>
	ItemStatusList    []string `json:"item_status_list"`     // [Required] <p>This is param to indicate the status of items, so sellers can know whether an item is eligible for ads or not.<br /></p>
	SkuTagList        []string `json:"sku_tag_list"`         // [Required] <p>The&nbsp;corresponding tag (or tags) that belong to item_id, sequences follow as best selling&gt;best ROI&gt;top search<br /></p>
	OngoingAdTypeList []string `json:"ongoing_ad_type_list"` // [Required] <p>Current status of the ad on this item. For example- no ongoing promotion, search ads, discovery ads, boost ads<br /></p>
}

type GetRecommendedKeywordListRequest struct {
	ItemId       int64   `json:"item_id" url:"item_id"`                                 // [Required] <p>Shopee's unique identifier for an item.<br /></p>
	InputKeyword *string `json:"input_keyword,omitempty" url:"input_keyword,omitempty"` // [Optional] <p>The keyword seller typed in the manually add keyword window.<br /></p>
}

type GetRecommendedKeywordListResponse struct {
	BaseResponse `json:",inline"`                      // Common response fields
	Response     GetRecommendedKeywordListResponseData `json:"response"` // Actual response data
}

type GetRecommendedKeywordListResponseData struct {
	ItemId            int64               `json:"item_id"`            // [Required] <p>Shopee's unique identifier for an item.<br /></p>
	InputKeyword      string              `json:"input_keyword"`      // [Required] <p>The keyword seller typed in the manually add keyword window.<br /></p>
	SuggestedKeywords []SuggestedKeywords `json:"suggested_keywords"` // [Required] <p>Suggested keywords recommended from product.<br /></p>
}

type GetReturnDetailRequest struct {
	ReturnSn string `json:"return_sn" url:"return_sn"` // [Required] The serial number of return.
}

type GetReturnDetailResponse struct {
	BaseResponse `json:",inline"`            // Common response fields
	Response     GetReturnDetailResponseData `json:"response"` // Actual response data
}

type GetReturnDetailResponseData struct {
	Image                               []string                          `json:"image"`                                   // [Required] Image URLs of return.
	BuyerVideos                         []BuyerVideos                     `json:"buyer_videos"`                            // [Required]
	Reason                              string                            `json:"reason"`                                  // [Required] <p>Indicates the original return reason submitted by the buyer when initiating the return request.</p><p><br /></p><p>Applicable values: See Data Definition- ReturnReason and Reassessed Request Reason.</p><p><br /></p><p> </p><p>Note: There may be cases where Shopee Agent updates the return request with a "Reassessed Return Reason" after reviewing more details about the buyer's return request and potentially after requesting evidence from the seller.&nbsp;If the platform updates the return reason during this process, the reassessed outcome will be provided separately in the <b>reassessed_request_reason</b> field.</p>
	TextReason                          string                            `json:"text_reason"`                             // [Required] Reason that buyer provide.
	ReassessedRequestReason             string                            `json:"reassessed_request_reason"`               // [Required] <p>Indicates the return reason reassessed by the platform as more suitable.</p><p><br /></p><p>There may be cases where Shopee Agent updates the return request with a "Reassessed Return Reason" after reviewing more details about the buyer's return request and potentially after requesting evidence from the seller.</p><p><br /></p><p>Applicable values: See Data Definition- ReturnReason and Reassessed Request Reason. If no reassessment has been made, the value will be NONE.</p>
	ReturnSn                            string                            `json:"return_sn"`                               // [Required] The serial number of return.
	RefundAmount                        float64                           `json:"refund_amount"`                           // [Required] Amount of the refund.
	Currency                            string                            `json:"currency"`                                // [Required] Currency of the return.
	CreateTime                          int64                             `json:"create_time"`                             // [Required] The time of return create.
	UpdateTime                          int64                             `json:"update_time"`                             // [Required] The time of modify return.
	Status                              string                            `json:"status"`                                  // [Required] Enumerated type that defines the current status of the return. Applicable values: See Data Definition- ReturnStatus.
	DueDate                             int64                             `json:"due_date"`                                // [Required] The last time seller deal with this return.
	TrackingNumber                      string                            `json:"tracking_number"`                         // [Required] The tracking number assigned by the shipping carrier for item shipment.
	DisputeReason                       int64                             `json:"dispute_reason"`                          // [Required] The reason of seller dispute return. While the return has been disputed, this field is useful. Applicable values: See Data Definition- ReturnDisputeReason.
	DisputeTextReason                   string                            `json:"dispute_text_reason"`                     // [Required] The reason that seller provide. While the return has been disputed, this field is useful.
	NeedsLogistics                      bool                              `json:"needs_logistics"`                         // [Required] Items to be sent back to seller. Can be either integrated/non-integrated.
	AmountBeforeDiscount                float64                           `json:"amount_before_discount"`                  // [Required] Order price before discount.
	User                                *User                             `json:"user"`                                    // [Required]
	Item                                []GetReturnDetailResponseDataItem `json:"item"`                                    // [Required]
	OrderSn                             string                            `json:"order_sn"`                                // [Required] Shopee's unique identifier for an order.
	ReturnShipDueDate                   int64                             `json:"return_ship_due_date"`                    // [Required] The due date for buyer to ship order.
	ReturnSellerDueDate                 int64                             `json:"return_seller_due_date"`                  // [Required] The due date for seller to deal with this return when buyer have shipped order.
	Activity                            []Activity                        `json:"activity"`                                // [Required]
	SellerProof                         *SellerProof                      `json:"seller_proof"`                            // [Required]
	SellerCompensation                  *SellerCompensation               `json:"seller_compensation"`                     // [Required]
	Negotiation                         *Negotiation                      `json:"negotiation"`                             // [Required]
	LogisticsStatus                     LogisticsStatus                   `json:"logistics_status"`                        // [Required] <p>To indicate the reverse logistics status. See "Data Definition - LogisticsStatus".</p><p><br /></p><p>Note:&nbsp;</p><p>- This is a legacy field that only reflects the reverse logistics status of Normal RR. To determine whether the RR is a Normal RR, check if return_refund_request_type = 0.</p><p>- If you need the reverse logistics status for Normal RR, In-transit RR, or Return-on-the-Spot, please use the newly released field reverse_logistic_status instead.</p>
	ReverseLogisticStatus               string                            `json:"reverse_logistic_status"`                 // [Required] <p>To indicate the latest reverse logistic status of a return, referring to the current status of the buyer shipping the return parcel back to the validation point (seller or warehouse),&nbsp;including Normal RR, In-transit RR, and Return-on-the-Spot.</p><p><br /></p><p>See "Data Definition - ReverseLogisticsStatus" as status displayed for Normal RR and In-transit RR or Return-on-the-Spot are different.</p>
	ReturnPickupAddress                 *ReturnPickupAddress              `json:"return_pickup_address"`                   // [Required] To indicate the buyer's pickup address
	VirtualContactNumber                string                            `json:"virtual_contact_number"`                  // [Required] <p>[Only for TW non-integrated channel] The virtual phone number to contact the recipient.</p>
	PackageQueryNumber                  string                            `json:"package_query_number"`                    // [Required] <p>[Only for TW non-integrated channel] The query number used in virtual phone number calls to contact the recipient of this return.</p>
	ReturnAddress                       *ReturnAddress                    `json:"return_address"`                          // [Required]
	ReturnRefundType                    string                            `json:"return_refund_type"`                      // [Required] <p>To indicate whether the return is RRBOC (Return/Refund request raised before Order Complete) or RRAOC&nbsp;(Return/Refund request raised after Order Complete).</p>
	ReturnSolution                      int64                             `json:"return_solution"`                         // [Required] <p>To indicate the most updated solution of the Return/Refund request (NOTE: this is not the solution during negotiation). Applicable value:&nbsp;</p><p>- 0: Return and Refund</p><p>- 1: Refund Only<br /></p>
	IsSellerArrange                     bool                              `json:"is_seller_arrange"`                       // [Required] <p>To indicate whether the return_sn is using the “Seller Arrange” return method. This would only be True for TW and BR.</p>
	IsShippingProofMandatory            bool                              `json:"is_shipping_proof_mandatory"`             // [Required] <p>To indicate whether uploading shipping proof is mandatory for seller to confirm "Arrange Pickup" when is_seller_arrange = true.</p>
	HasUploadedShippingProof            bool                              `json:"has_uploaded_shipping_proof"`             // [Required] <p>To indicate whether seller has already uploaded shipping proof for this return.</p>
	IsReverseLogisticsChannelIntegrated bool                              `json:"is_reverse_logistics_channel_integrated"` // [Required] <p>To indicate whether the reverse logistic channel type selected is integrated or non-integrated.</p>
	ReverseLogisticChannelName          string                            `json:"reverse_logistic_channel_name"`           // [Required] <p>To indicate reverse logistic carrier name.</p>
	ReturnRefundRequestType             int64                             `json:"return_refund_request_type"`              // [Required] <p>To indicate the type of return refund request, whether it is a Normal RR request, an In-transit RR request, and a Return on the Spot:&nbsp;</p><p>0:&nbsp;Normal RR (RR is raised by the buyer after delivery done / estimated delivery date)</p><p>1: In-transit RR (RR is raised by the buyer while item is still in-transit to buyer)</p><p>2: Return-on-the-Spot (RR is raised by the driver after buyer rejected parcel at delivery)</p><p><br /></p><p>For more details, see Data Definition- Return Refund Request Type.</p>
	ValidationType                      string                            `json:"validation_type"`                         // [Required] <p>To indicate whether seller or warehouse will expect to receive the return parcel from buyer and validate the condition of the parcel:&nbsp;</p><p>- seller_validation&nbsp;</p><p>- warehouse_validation</p><p><br /></p><p>For more details, see Data Definition- ValidationType.</p>
	IsArrivedAtWarehouse                int64                             `json:"is_arrived_at_warehouse"`                 // [Required] <p><b>[Only for validation_type =&nbsp;warehouse_validation]</b> Indicates the parcel’s check-in status at the warehouse. This field helps sellers quickly determine whether the parcel has arrived at the warehouse or has been rejected.&nbsp;</p><p><br /></p><p>Applicable values:</p><p>1: Pending Inbound</p><p>2: Rejected</p><p>3: Inbound</p><p>4: Cancelled</p>
	FollowUpActionList                  []FollowUpAction                  `json:"follow_up_action_list"`                   // [Required] <p><b>[Only for validation_type =&nbsp;warehouse_validation]</b> Warehouse handling actions for each item in the parcel.</p>
}

type GetReturnDetailResponseDataItem struct {
	ModelId      int64    `json:"model_id"`       // [Required] Shopee's unique identifier for a variation of an item.
	Name         string   `json:"name"`           // [Required] Name of item in local language.
	Images       []string `json:"images"`         // [Required] Image URLs of item.
	Amount       int64    `json:"amount"`         // [Required] Amount of this item.
	ItemPrice    float64  `json:"item_price"`     // [Required] The price of item.
	IsAddOnDeal  bool     `json:"is_add_on_deal"` // [Required] To indicate if this item belongs to an addon deal.
	IsMainItem   bool     `json:"is_main_item"`   // [Required] To indicate if this item is main item or sub item. True means main item, false means sub item.
	AddOnDealId  int64    `json:"add_on_deal_id"` // [Required] The unique identity of an addon deal.
	ItemId       int64    `json:"item_id"`        // [Required] The id of item.
	ItemSku      string   `json:"item_sku"`       // [Required] The sku of item.
	VariationSku string   `json:"variation_sku"`  // [Required] the variation sku of item
	RefundAmount float64  `json:"refund_amount"`  // [Required] <p>item's refund amount.&nbsp;only for shops whitelisted for Partial Qty RR.</p><p>If not available, refer to item_price</p>
}

type GetReturnDisputeReasonRequest struct {
	ReturnSn string `json:"return_sn" url:"return_sn"` // [Required] <p>The serial number of return.</p>
}

type GetReturnDisputeReasonResponse struct {
	BaseResponse `json:",inline"`                   // Common response fields
	Response     GetReturnDisputeReasonResponseData `json:"response"` // Actual response data
}

type GetReturnDisputeReasonResponseData struct {
	DisputeReasonList []DisputeReason `json:"dispute_reason_list"` // [Required] <p>The dispute_reason and associated evidence list.<br /></p>
}

type GetReturnListRequest struct {
	PageNo                   int64   `json:"page_no" url:"page_no"`                                                           // [Required] Specifies the starting entry of data to return in the current call. Default is 0. if data is more than one page, the offset can be some entry to start next call.
	PageSize                 int64   `json:"page_size" url:"page_size"`                                                       // [Required] if many items are available to retrieve, you may need to call GetReturnList multiple times to retrieve all the data. Each result set is returned as a page of entries. Default is 40. Use the Pagination filters to control the maximum number of entries (<= 100) to retrieve per page (i.e., per call), the offset number to start next call. This integer value is usUed to specify the maximum number of entries to return in a single ""page"" of data.
	CreateTimeFrom           *int64  `json:"create_time_from,omitempty" url:"create_time_from,omitempty"`                     // [Optional] The create_time_from and create_time_to fields specify a date range for retrieving orders (based on the order create time). The create_time_from field is the starting date range. The maximum date range that may be specified with the create_time_from and create_time_to fields is 15 days.
	CreateTimeTo             *int64  `json:"create_time_to,omitempty" url:"create_time_to,omitempty"`                         // [Optional] The create_time_from and create_time_to fields specify a date range for retrieving orders (based on the order create time). The create_time_from field is the starting date range. The maximum date range that may be specified with the create_time_from and create_time_to fields is 15 days.
	UpdateTimeFrom           *int64  `json:"update_time_from,omitempty" url:"update_time_from,omitempty"`                     // [Optional] <p>The update_time_from and update_time_to fields specify a date range for retrieving orders (based on the last return updated time). The update_time_from field is the starting date range. The maximum date range that may be specified with the update_time_from and update_time_to fields is 15 days. update_time_from should be &gt;= create_time_from<br /></p>
	UpdateTimeTo             *int64  `json:"update_time_to,omitempty" url:"update_time_to,omitempty"`                         // [Optional] <p>The update_time_from and update_time_to fields specify a date range for retrieving orders (based on the last return updated time). The update_time_from field is the starting date range. The maximum date range that may be specified with the update_time_from and update_time_to fields is 15 days. update_time_from should be &gt;= create_time_from<br /></p>
	Status                   *string `json:"status,omitempty" url:"status,omitempty"`                                         // [Optional] This is for filtering return request by return status. See "Data Definition - ReturnStatus"
	NegotiationStatus        *string `json:"negotiation_status,omitempty" url:"negotiation_status,omitempty"`                 // [Optional] This is for filtering return request by counter status. See "Data Definition - NegotiationStatus"
	SellerProofStatus        *string `json:"seller_proof_status,omitempty" url:"seller_proof_status,omitempty"`               // [Optional] <p>This is for filtering return request by proof status. See "Data Definition - SellerProofStatus"</p>
	SellerCompensationStatus *string `json:"seller_compensation_status,omitempty" url:"seller_compensation_status,omitempty"` // [Optional] This is for filtering return request by compensation status. See "Data Definition - SellerCompensationStatus"
}

type GetReturnListResponse struct {
	BaseResponse `json:",inline"`          // Common response fields
	Response     GetReturnListResponseData `json:"response"` // Actual response data
}

type GetReturnListResponseData struct {
	More   bool     `json:"more"`   // [Required] Whether has next page
	Return []Return `json:"return"` // [Required]
}

type GetReverseTrackingInfoRequest struct {
	ReturnSn string `json:"return_sn" url:"return_sn"` // [Required] <p>Shopee's unique identifier for a return/refund request (serial number of return).</p>
}

type GetReverseTrackingInfoResponse struct {
	BaseResponse `json:",inline"`                   // Common response fields
	Response     GetReverseTrackingInfoResponseData `json:"response"` // Actual response data
}

type GetReverseTrackingInfoResponseData struct {
	ReturnSn                        string                     `json:"return_sn"`                           // [Required] <p>Shopee's unique identifier for a return/refund request (serial number of return).<br /></p>
	ReturnRefundRequestType         int64                      `json:"return_refund_request_type"`          // [Required] <p>To indicate the type of return refund request, whether it is a Normal RR request, an In-transit RR request, and a Return on the Spot:&nbsp;</p><p>0:&nbsp;Normal RR (RR is raised by the buyer after delivery done / estimated delivery date)</p><p>1: In-transit RR (RR is raised by the buyer while item is still in-transit to buyer)</p><p>2: Return-on-the-Spot (RR is raised by the driver after buyer rejected parcel at delivery)</p><p><br /></p><p>For more details, see Data Definition- Return Refund Request Type.</p>
	ValidationType                  string                     `json:"validation_type"`                     // [Required] <p>To indicate whether seller or warehouse will expect to receive the return parcel from buyer and validate the condition of the parcel:&nbsp;</p><p>- seller_validation&nbsp;</p><p>- warehouse_validation</p><p><br /></p><p>For more details, see Data Definition- ValidationType.</p>
	ReverseLogisticsStatus          string                     `json:"reverse_logistics_status"`            // [Required] <p>To indicate the latest reverse logistic status of a return, referring to the current status of the buyer shipping the return parcel back to the validation point (seller or warehouse),&nbsp;including Normal RR, In-transit RR, and Return-on-the-Spot.</p><p><br /></p><p>See "Data Definition - ReverseLogisticsStatus" as status displayed for Normal RR and In-transit RR or Return-on-the-Spot are different.</p><p><br /></p><p>Note: If <b>validation_type = seller_validation</b>, there is <b>only one segment of reverse logistics (The buyer ships the return parcel directly back to the seller).</b> Please use the fields reverse_logistics_status, reverse_logistics_update_time, tracking_number, and tracking_info to obtain the reverse logistics tracking information.</p>
	ReverseLogisticsUpdateTime      int64                      `json:"reverse_logistics_update_time"`       // [Required] <p>The last update time of the reverse logistics status including Normal RR, In-transit RR, and Return-on-the-Spot.</p><p><br /></p><p>Note: If <b>validation_type = seller_validation</b>, there is <b>only one segment of reverse logistics (The buyer ships the return parcel directly back to the seller).</b> Please use the fields reverse_logistics_status, reverse_logistics_update_time, tracking_number, and tracking_info to obtain the reverse logistics tracking information.</p>
	EstimatedDeliveryDateMax        int64                      `json:"estimated_delivery_date_max"`         // [Required] <p>The maximum estimated delivery date for the reverse logistics. This is calculated by Shopee Logistics Services once buyer ships out if there is historical tracking data available from third party logistics provider.&nbsp;</p><p><br /></p><p>Note: Only available for <b>Normal RR with integrated reverse logistics.</b></p>
	EstimatedDeliveryDateMin        int64                      `json:"estimated_delivery_date_min"`         // [Required] <p>The minimum estimated delivery date for the reverse logistics.&nbsp;This is calculated by Shopee Logistics Services once buyer ships out if there is historical tracking data available from third party logistics provider.</p><p><br /></p><p>Note: Only available for <b>Normal RR with integrated reverse logistics.</b></p>
	TrackingNumber                  string                     `json:"tracking_number"`                     // [Required] <p>The tracking number for the reverse logistics&nbsp;(the logistics tracking number provided when the buyer ships the item back).</p><p><br /></p><p>Note:&nbsp;</p><p>- Only available for <b>Normal RR with integrated reverse logistics.</b></p><p>- If <b>validation_type = seller_validation</b>, there is <b>only one segment of reverse logistics (The buyer ships the return parcel directly back to the seller)</b>. Please use the fields reverse_logistics_status, reverse_logistics_update_time, tracking_number, and tracking_info to obtain the reverse logistics tracking information.</p>
	TrackingInfo                    []ResponseDataTrackingInfo `json:"tracking_info"`                       // [Required] <p>The detailed tracking information list for the reverse logistics.<b></b></p><p><b><br /></b></p><p>Note:&nbsp;</p><p>- Only available for <b>Normal RR with integrated reverse logistics</b>, with the tracking information pushed by third party logistics provider to Shopee.</p><p>- If <b>validation_type = seller_validation</b>, there is <b>only one segment of reverse logistics (The buyer ships the return parcel directly back to the seller)</b>. Please use the fields reverse_logistics_status, reverse_logistics_update_time, tracking_number, and tracking_info to obtain the reverse logistics tracking information.</p>
	PostReturnLogisticsStatus       string                     `json:"post_return_logistics_status"`        // [Required] <p>Post-return logistics status, referring to the current status of the warehouse shipping the return parcel back to the seller in warehouse validation mode. See&nbsp;"Data Definition - Post Return Logistics Status".</p><p><br /></p><p>Note:&nbsp;</p><p>- Only&nbsp;available for <b>Normal RR with return_solution = 0 (Return and Refund)</b> and<b> validation_type = warehouse_validation</b>, and the warehouse ships the return parcel back to seller <b>using integrated reverse logistics.</b></p><p>- If <b>validation_type = warehouse_validation</b> AND the warehouse <b>uses an integrated logistics channel</b> to ship the return parcel back to the seller, there are <b>two segments of reverse logistics</b>:&nbsp;</p><p style="padding-left:2em;">- <b>The buyer first ships the return parcel back to the warehouse.</b> Use the fields reverse_logistics_status, reverse_logistics_update_time, tracking_number, and tracking_info to obtain tracking information for this first segment.</p><p style="padding-left:2em;">- <b>The warehouse then ships the return parcel back to the seller.</b> Use the fields post_return_logistics_status, post_return_logistics_update_time, rts_tracking_number, and post_return_logistics_tracking_info to obtain tracking information for this second segment (post-return logistics).</p><p style="padding-left:2em;">- For Cross-Border Returns, if the second segment exists, the API <b>returns information for both the first and second segments.</b> For Local Returns, if the second segment exists, the API <b>prioritizes and returns only the second segment information.</b></p>
	PostReturnLogisticsUpdateTime   int64                      `json:"post_return_logistics_update_time"`   // [Required] <p>The last update time of the post-return logistics status where warehouse sends return parcel from warehouse to seller.</p><p><br /></p><p>Note:&nbsp;</p><p>-&nbsp;Only&nbsp;available for <b>Normal RR with return_solution = 0 (Return and Refund)</b> and <b>validation_type = warehouse_validation</b>, and the warehouse ships the return parcel back to seller <b>using integrated reverse logistics.</b></p><p>-&nbsp;If <b>validation_type = warehouse_validation</b> AND the warehouse <b>uses an integrated logistics channel</b> to ship the return parcel back to the seller, there are <b>two segments of reverse logistics:&nbsp;</b></p><p style="padding-left:2em;">-&nbsp;<b>The buyer first ships the return parcel back to the warehouse.</b>&nbsp;Use the fields reverse_logistics_status, reverse_logistics_update_time, tracking_number, and tracking_info to obtain tracking information for this first segment.</p><p style="padding-left:2em;">-&nbsp;<b>The warehouse then ships the return parcel back to the seller.</b>&nbsp;Use the fields post_return_logistics_status, post_return_logistics_update_time, rts_tracking_number, and post_return_logistics_tracking_info to obtain tracking information for this second segment (post-return logistics).</p><p style="padding-left:2em;">- For Cross-Border Returns, if the second segment exists, the API&nbsp;<b>returns information for both the first and second segments</b>.&nbsp;For Local Returns, if the second segment exists, the API&nbsp;<b>prioritizes and returns only the second segment information.</b></p>
	RtsTrackingNumber               string                     `json:"rts_tracking_number"`                 // [Required] <p>The tracking number for the post-return logistics (the logistics tracking number used when the warehouse ships the parcel back to the seller). RTS stands for "Return to Seller".</p><p><br /></p><p>Note:&nbsp;</p><p>-&nbsp;Only&nbsp;available for <b>Normal RR with return_solution = 0 (Return and Refund)</b> and <b>validation_type = warehouse_validation</b>, and the warehouse ships the return parcel back to seller <b>using integrated reverse logistics.</b></p><p>-<b>&nbsp;If validation_type = warehouse_validation</b> AND the warehouse <b>uses an integrated logistics channel</b> to ship the return parcel back to the seller, there are <b>two segments of reverse logistics:&nbsp;</b></p><p style="padding-left:2em;">-&nbsp;<b>The buyer first ships the return parcel back to the warehouse.</b>&nbsp;Use the fields reverse_logistics_status, reverse_logistics_update_time, tracking_number, and tracking_info to obtain tracking information for this first segment.</p><p style="padding-left:2em;">-&nbsp;<b>The warehouse then ships the return parcel back to the seller.</b>&nbsp;Use the fields post_return_logistics_status, post_return_logistics_update_time, rts_tracking_number, and post_return_logistics_tracking_info to obtain tracking information for this second segment (post-return logistics).</p><p style="padding-left:2em;">- For Cross-Border Returns, if the second segment exists, the API&nbsp;<b>returns information for both the first and second segments.</b>&nbsp;For Local Returns, if the second segment exists, the API&nbsp;<b>prioritizes and returns only the second segment information.</b></p>
	PostReturnLogisticsTrackingInfo []ResponseDataTrackingInfo `json:"post_return_logistics_tracking_info"` // [Required] <p>Only&nbsp;available for <b>Normal RR with return_solution = 0 (Return and Refund)</b> and <b>validation_type = warehouse_validation</b>, and the warehouse ships the return parcel back to seller <b>using integrated reverse logistics.</b></p><p><b><br /></b></p><p>In this scenario, the tracking logistics from warehouse to seller is called "post-return logistics", with the tracking information pushed by third party logistics provider to Shopee.</p><p><br /></p><p>Note:&nbsp;</p><p>-&nbsp;If <b>validation_type = warehouse_validation</b>&nbsp;AND the warehouse&nbsp;<b>uses an integrated logistics channel</b>&nbsp;to ship the return parcel back to the seller, there are&nbsp;<b>two segments of reverse logistics:&nbsp;</b></p><p style="padding-left:2em;">-&nbsp;<b>The buyer first ships the return parcel back to the warehouse.</b>&nbsp;Use the fields reverse_logistics_status, reverse_logistics_update_time, tracking_number, and tracking_info to obtain tracking information for this first segment.</p><p style="padding-left:2em;">-&nbsp;<b>The warehouse then ships the return parcel back to the seller.</b>&nbsp;Use the fields post_return_logistics_status, post_return_logistics_update_time, rts_tracking_number, and post_return_logistics_tracking_info to obtain tracking information for this second segment (post-return logistics).</p><p style="padding-left:2em;">- For Cross-Border Returns, if the second segment exists, the API&nbsp;<b>returns information for both the first and second segments.</b>&nbsp;For Local Returns, if the second segment exists, the API&nbsp;<b>prioritizes and returns only the second segment information.</b></p>
}

type GetSessionDetailRequest struct {
	SessionId int64 `json:"session_id" url:"session_id"` // [Required] <p>The identifier of livestream session.</p>
}

type GetSessionDetailResponse struct {
	BaseResponse `json:",inline"`             // Common response fields
	Response     GetSessionDetailResponseData `json:"response"` // Actual response data
}

type GetSessionDetailResponseData struct {
	SessionId     int64          `json:"session_id"`      // [Required] <p>The identifier of livestream session.</p>
	Title         string         `json:"title"`           // [Required] <p>The title of the livestream session.</p>
	Description   string         `json:"description"`     // [Required] <p>The description of the livestream session.</p>
	CoverImageUrl string         `json:"cover_image_url"` // [Required] <p>The cover image URL of the livestream session.</p>
	Status        int64          `json:"status"`          // [Required] <p>The status of the livestream session, the enumeration values are as follows:</p><p>0 - Initial</p><p>1 - Ongoing</p><p>2 - Ended</p>
	ShareUrl      string         `json:"share_url"`       // [Required] <p>The share link of the livestream session.</p>
	IsTest        bool           `json:"is_test"`         // [Required] <p>Indicate whether this livestream session if for testing purpose only.</p>
	CreateTime    int64          `json:"create_time"`     // [Required] <p>The creation time of the livestream session. It's unix timestamp in seconds.</p>
	UpdateTime    int64          `json:"update_time"`     // [Required] <p>The update time of the livestream session. It's unix timestamp in seconds.</p>
	StartTime     int64          `json:"start_time"`      // [Required] <p>The start time of the livestream session, 0 if session is not started yet.&nbsp;It's unix timestamp in seconds.</p>
	EndTime       int64          `json:"end_time"`        // [Required] <p>The end time of livestream session, 0 if session is not ended yet.&nbsp;It's unix timestamp in seconds.</p>
	StreamUrlList *StreamUrlList `json:"stream_url_list"` // [Required]
}

type GetSessionItemMetricRequest struct {
	SessionId int64 `json:"session_id" url:"session_id"` // [Required] <p>The identifier of livestream session.</p>
	Offset    int64 `json:"offset" url:"offset"`         // [Required] <p>Specifies the starting entry of data to return in the current call. Default is 0, if data is more than one page, the offset can be some entry to start next call.</p>
	PageSize  int64 `json:"page_size" url:"page_size"`   // [Required] <p>Each result set is returned as a page of entries. Use the "page_size" filters to control the maximum number of entries to retrieve per page (i.e., per call). This integer value is used to specify the maximum number of entries to return in a single "page" of data. The limit of page_size if between 1 and 100.</p>
}

type GetSessionItemMetricResponse struct {
	BaseResponse `json:",inline"`                 // Common response fields
	Response     GetSessionItemMetricResponseData `json:"response"` // Actual response data
}

type GetSessionItemMetricResponseData struct {
	More       bool                                   `json:"more"`        // [Required] <p>This is to indicate whether the list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of data.</p>
	NextOffset int64                                  `json:"next_offset"` // [Required] <p>If more is true, this value need set to next request offset.</p>
	List       []GetSessionItemMetricResponseDataList `json:"list"`        // [Required]
}

type GetSessionItemMetricResponseDataList struct {
	Item   *ListItem   `json:"item"`   // [Required]
	Metric *ListMetric `json:"metric"` // [Required]
}

type GetSessionMetricRequest struct {
	SessionId int64 `json:"session_id" url:"session_id"` // [Required] <p>The identifier of livestream session.</p>
}

type GetSessionMetricResponse struct {
	BaseResponse `json:",inline"`             // Common response fields
	Response     GetSessionMetricResponseData `json:"response"` // Actual response data
}

type GetSessionMetricResponseData struct {
	Gmv                float64 `json:"gmv"`                  // [Required] <p>Value of placed orders (paid and unpaid) during Livestream, including sales from cancelled orders.<br /></p>
	Atc                int64   `json:"atc"`                  // [Required] <p>Number of "Add To Cart" button clicked for all products in the orange bag during livestream.<br /></p>
	Ctr                float64 `json:"ctr"`                  // [Required] <p>Number of products clicks divided by Number of Livestream views.<br /></p>
	Co                 float64 `json:"co"`                   // [Required] <p>Amount of product orders from the stream divided by Amount of product clicks from the stream.<br /></p>
	Orders             int64   `json:"orders"`               // [Required] <p>Number of placed orders (paid and unpaid) during Livestream, including cancelled orders.<br /></p>
	Ccu                int64   `json:"ccu"`                  // [Required] <p>Number of viewers during stream.</p>
	EngageCcu_1m       int64   `json:"engage_ccu_1m"`        // [Required] <p>Number of Concurrent viewers in the stream that have watched for more than 1 minute.<br /></p>
	PeakCcu            int64   `json:"peak_ccu"`             // [Required] <p>Highest number of viewers during stream.<br /></p>
	Likes              int64   `json:"likes"`                // [Required] <p>Number of "Like" clicked during livestream.<br /></p>
	Comments           int64   `json:"comments"`             // [Required] <p>Number of comments acquired during the stream.<br /></p>
	Shares             int64   `json:"shares"`               // [Required] <p>Number of shares created during the stream.<br /></p>
	Views              int64   `json:"views"`                // [Required] <p>Number of views from the stream.<br /></p>
	AvgViewingDuration int64   `json:"avg_viewing_duration"` // [Required] <p>Average of Viewer duration watching in the stream.<br /></p>
}

type GetShipmentListRequest struct {
	Cursor   *string `json:"cursor,omitempty" url:"cursor,omitempty"` // [Optional] Specifies the starting entry of data to return in the current call. Default is "". If data is more than one page, the offset can be some entry to start next call.
	PageSize int64   `json:"page_size" url:"page_size"`               // [Required] Each result set is returned as a page of entries. Use the "page_size" filters to control the maximum number of entries to retrieve per page (i.e., per call). This integer value is used to specify the maximum number of entries to return in a single "page" of data.The limit of page_size if between 1 and 100.
}

type GetShipmentListResponse struct {
	BaseResponse `json:",inline"`            // Common response fields
	Response     GetShipmentListResponseData `json:"response"` // Actual response data
}

type GetShipmentListResponseData struct {
	OrderList  []SharedOrder `json:"order_list"`  // [Required] The list of  shipment orders
	More       bool          `json:"more"`        // [Required] This is to indicate whether the order list is more than one page. If this value is true, you may want to continue to check next page to retrieve orders.
	NextCursor string        `json:"next_cursor"` // [Required] If more is true, you should pass the next_cursor in the next request as cursor. The value of next_cursor will be empty string when more is false.
}

type GetShippingCarrierRequest struct {
	ReturnSn string `json:"return_sn" url:"return_sn"` // [Required] <p>The serial number of return.<br /></p>
}

type GetShippingCarrierResponse struct {
	BaseResponse `json:",inline"`               // Common response fields
	Response     GetShippingCarrierResponseData `json:"response"` // Actual response data
}

type GetShippingCarrierResponseData struct {
	IsShippingProofMandatory      bool                      `json:"is_shipping_proof_mandatory"`       // [Required] <p>To indicate whether uploading shipping proof is mandatory for seller to confirm "Arrange Pickup" when is_seller_arrange = true.</p>
	HasUploadedSellerArrangeProof bool                      `json:"has_uploaded_seller_arrange_proof"` // [Required] <p>To indicate whether seller has already uploaded shipping proof for this return.</p>
	ShippingProofTemplate         []ShippingProofTemplate   `json:"shipping_proof_template"`           // [Required] <p>To display list of request parameters needed to upload shipping proof.</p>
	ReverseLogisticsCarrierList   []ReverseLogisticsCarrier `json:"reverse_logistics_carrier_list"`    // [Required] <p>The list of logistics carriers available for sellers to choose.</p>
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

type GetShopCategoryListRequest struct {
	PageSize int64 `json:"page_size"` // [Required] Specifies the starting entry of data to return in the current call. The parameter range of page_size should be [1, 2147483647]
	PageNo   int64 `json:"page_no"`   // [Required] Specifies the total returned data per entry. The parameter range of page_no should be [1, 100]
}

type GetShopCategoryListResponse struct {
	BaseResponse `json:",inline"`                // Common response fields
	Response     GetShopCategoryListResponseData `json:"response"` // Actual response data
}

type GetShopCategoryListResponseData struct {
	ShopCategorys []ShopCategorys `json:"shop_categorys"` // [Required] ShopCategory's unique identifier.
	TotalCount    int64           `json:"total_count"`    // [Required] This is to indicate the whole number of  in-shop categories under the shop.
	More          bool            `json:"more"`           // [Required] This is to indicate whether the list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest.
}

type GetShopFlashSaleItemsRequest struct {
	FlashSaleId int64 `json:"flash_sale_id"` // [Required]
	Offset      int64 `json:"offset"`        // [Required] <pre>min=0,max=1000</pre>
	Limit       int64 `json:"limit"`         // [Required] <pre>min=1,max=100</pre>
}

type GetShopFlashSaleItemsResponse struct {
	BaseResponse `json:",inline"`                  // Common response fields
	Response     GetShopFlashSaleItemsResponseData `json:"response"` // Actual response data
}

type GetShopFlashSaleItemsResponseData struct {
	TotalCount int64                `json:"total_count"` // [Required]
	Models     []ResponseDataModels `json:"models"`      // [Required] <p>If the item has variation, the infomation of model will be in this field</p>
	ItemInfo   []ItemInfo           `json:"item_info"`   // [Required]
}

type GetShopFlashSaleListRequest struct {
	Type      int64  `json:"type"`                 // [Required] <p>you can use this filed to search different state of shop flash sale</p><p><br /></p><p>0: all state</p><p>1: upcoming state</p><p>2: ongoing state</p><p>3: expired state</p>
	StartTime *int64 `json:"start_time,omitempty"` // [Optional] <p>you should use start_time and end_time together, and start_time shoule be &lt;&nbsp;end_time</p>
	EndTime   *int64 `json:"end_time,omitempty"`   // [Optional] <p>you should use start_time and end_time together, and start_time shoule be &lt;&nbsp;end_time<br /></p>
	Offset    int64  `json:"offset"`               // [Required] <pre>min=0,max=1000</pre>
	Limit     int64  `json:"limit"`                // [Required] <pre>min=1,max=100</pre>
}

type GetShopFlashSaleListResponse struct {
	BaseResponse  `json:",inline"`                 // Common response fields
	Response      GetShopFlashSaleListResponseData `json:"response"`                  // Actual response data
	FlashSaleList []FlashSale                      `json:"flash_sale_list,omitempty"` //
}

type GetShopFlashSaleListResponseData struct {
	TotalCount int64 `json:"total_count"` // [Required] <p>the number of shop flash sale that the shop has</p>
}

type GetShopFlashSaleRequest struct {
	FlashSaleId int64 `json:"flash_sale_id"` // [Required]
}

type GetShopFlashSaleResponse struct {
	BaseResponse `json:",inline"`             // Common response fields
	Response     GetShopFlashSaleResponseData `json:"response"` // Actual response data
}

type GetShopFlashSaleResponseData struct {
	TimeslotId       int64 `json:"timeslot_id"`        // [Required]
	FlashSaleId      int64 `json:"flash_sale_id"`      // [Required]
	Status           int64 `json:"status"`             // [Required] <p>the status of shop flash sale</p><p>0: deleted</p><p>1: enabled</p><p>2: disabled</p>
	StartTime        int64 `json:"start_time"`         // [Required] <p>the start time of shop flash sale</p>
	EndTime          int64 `json:"end_time"`           // [Required] <p>the end time of shop flash sale<br /></p>
	EnabledItemCount int64 `json:"enabled_item_count"` // [Required] <p>the number of enabled items in shop flash sale</p>
	ItemCount        int64 `json:"item_count"`         // [Required] <p>the number of items in shop flash sale<br /></p>
	Type             int64 `json:"type"`               // [Required] <p>the state of shop flash sale</p><p>1: upcoming</p><p>2: ongoing</p><p>3: expired</p>
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

type GetShopInstallmentStatusResponse struct {
	BaseResponse `json:",inline"`                     // Common response fields
	Response     GetShopInstallmentStatusResponseData `json:"response"` // Actual response data
}

type GetShopInstallmentStatusResponseData struct {
	InstallmentStatus int64 `json:"installment_status"` // [Required] <p>The installment status for the shop</p>
}

type GetShopListByMerchantRequest struct {
	PageNo   int64 `json:"page_no" url:"page_no"`     // [Required] Specifies the page number of data to return in the current call. Starting from 1. if data is more than one page, the page_no can be some entry to start next call.
	PageSize int64 `json:"page_size" url:"page_size"` // [Required] Each result set is returned as a page of entries. Use the "page_size" filters to control the maximum number of entries to retrieve per page (i.e., per call), and the "page_no" to start next call. This integer value is used to specify the maximum number of entries to return in a single "page" of data.No more than 500.
}

type GetShopListByMerchantResponse struct {
	BaseResponse `json:",inline"` // Common response fields
	ShopList     []Shop           `json:"shop_list,omitempty"` // list of shop authorized to the partner and bound to the merchant.
	More         bool             `json:"more,omitempty"`      // This is to indicate whether the list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of datas.
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

type GetShopPerformanceResponse struct {
	BaseResponse `json:",inline"`               // Common response fields
	Response     GetShopPerformanceResponseData `json:"response"` // Actual response data
}

type GetShopPerformanceResponseData struct {
	OverallPerformance *OverallPerformance `json:"overall_performance"` // [Required]
	MetricList         []Metric            `json:"metric_list"`         // [Required]
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

type GetShopToggleInfoResponse struct {
	BaseResponse `json:",inline"`              // Common response fields
	Response     GetShopToggleInfoResponseData `json:"response"` // Actual response data
}

type GetShopToggleInfoResponseData struct {
	DataTimestamp int64 `json:"data_timestamp"` // [Required] <p>Timestamp of data in response<br /></p>
	AutoTopUp     bool  `json:"auto_top_up"`    // [Required] <p>auto_top_up toggle on/off<br /></p>
	CampaignSurge bool  `json:"campaign_surge"` // [Required] <p>campaign_surge toggle on/off<br /></p>
}

type GetShopeeIpRangesResponse struct {
	BaseResponse `json:",inline"` // Common response fields
	IpList       []string         `json:"ip_list,omitempty"` // <p>IP address ranges of Shopee</p>
}

type GetShopsByPartnerRequest struct {
	PageSize *int64 `json:"page_size,omitempty" url:"page_size,omitempty"` // [Optional] Each result set is returned as a page of entries. Use the "page_size" filters to control the maximum number of entries to retrieve per page (i.e., per call), and the "page_no" to start next call. This integer value is used to specify the maximum number of entries to return in a single "page" of data.
	PageNo   *int64 `json:"page_no,omitempty" url:"page_no,omitempty"`     // [Optional] Specifies the page number of data to return in the current call. Starting from 1. if data is more than one page, the page_no can be some entry to start next call.
}

type GetShopsByPartnerResponse struct {
	BaseResponse   `json:",inline"` // Common response fields
	AuthedShopList []AuthedShop     `json:"authed_shop_list,omitempty"` // A list of shops authorized to the partner.
	More           bool             `json:"more,omitempty"`             // This is to indicate whether the list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of datas.
}

type GetShowItemRequest struct {
	SessionId int64 `json:"session_id" url:"session_id"` // [Required] <p>The identifier of livestream session.</p>
}

type GetShowItemResponse struct {
	BaseResponse `json:",inline"`        // Common response fields
	Response     GetShowItemResponseData `json:"response"` // Actual response data
}

type GetShowItemResponseData struct {
	HasShowItem bool                         `json:"has_show_item"` // [Required] <p>Whether has the showing item.</p>
	Item        *GetShowItemResponseDataItem `json:"item"`          // [Required]
}

type GetShowItemResponseDataItem struct {
	ItemNo    int64          `json:"item_no"`    // [Required] <p>The order of this item in the shopping bag of current session, start from 1. Only return item_no when showing item is in the shopping bag of current session.</p>
	ItemId    int64          `json:"item_id"`    // [Required] <p>Shopee's unique identifier for an item.</p>
	Name      string         `json:"name"`       // [Required] <p>Name of the item in local language.<br /></p>
	ImageUrl  string         `json:"image_url"`  // [Required] <p>The image url of this item.</p>
	PriceInfo *ListPriceInfo `json:"price_info"` // [Required]
}

type GetSipDiscountsRequest struct {
	Region *string `json:"region,omitempty" url:"region,omitempty"` // [Optional] <p>The region of SIP affiliate shop that needs to get discount information.</p><p>If do not pass, will return the discount information set for all SIP affiliate shops.</p>
}

type GetSipDiscountsResponse struct {
	BaseResponse `json:",inline"`            // Common response fields
	Response     GetSipDiscountsResponseData `json:"response"` // Actual response data
}

type GetSipDiscountsResponseData struct {
	DiscountList []ResponseDataDiscount `json:"discount_list"` // [Required] <p>List of discounts in each region. Will be filtered based on the "region" request parameter.<br /></p>
}

type GetSizeChartDetailRequest struct {
	SizeChartId int64 `json:"size_chart_id" url:"size_chart_id"` // [Required] <p>ID of new size chart<br /></p>
}

type GetSizeChartDetailResponse struct {
	BaseResponse `json:",inline"`               // Common response fields
	Response     GetSizeChartDetailResponseData `json:"response"` // Actual response data
}

type GetSizeChartDetailResponseData struct {
	SizeChartId    int64           `json:"size_chart_id"`    // [Required] <p>ID of new size chart<br /></p>
	SizeChartName  string          `json:"size_chart_name"`  // [Required] <p>name of new size chart<br /></p>
	SizeChartTable *SizeChartTable `json:"size_chart_table"` // [Required] <p>new size chart is a table format which include multiple columns. each column has column header (measurement) and multiple values (measurement value) of this column.<br /></p>
}

type GetSizeChartListRequest struct {
	CategoryId int64   `json:"category_id" url:"category_id"`           // [Required] <p>category id under this shop<br /></p>
	PageSize   int64   `json:"page_size" url:"page_size"`               // [Required] <p>the size of one page.&nbsp;Max=50.<br /></p>
	Cursor     *string `json:"cursor,omitempty" url:"cursor,omitempty"` // [Optional] <p>Specifies the starting entry of data to return in the current call. Default is "". If data is more than one page, the cursor can be some entry to start next call.<br /></p>
}

type GetSizeChartListResponse struct {
	BaseResponse `json:",inline"`             // Common response fields
	Response     GetSizeChartListResponseData `json:"response"` // Actual response data
}

type GetSizeChartListResponseData struct {
	SizeChartList []SizeChart `json:"size_chart_list"` // [Required]
	TotalCount    int64       `json:"total_count"`     // [Required] <p>total number of new size chart under requested category_id<br /></p>
	NextCursor    string      `json:"next_cursor"`     // [Required] <p>if next_cursor has value, this value need set to next request.cursor<br /></p>
}

type GetSspInfoRequest struct {
	SspIdList []int64  `json:"ssp_id_list,omitempty"` // [Optional] <p>The ssp_id list; limit [0,50]</p>
	GtinList  []string `json:"gtin_list,omitempty"`   // [Optional] <p>The gtin_list; limit [0,50]</p>
	OemList   []string `json:"oem_list,omitempty"`    // [Optional] <p>The oem_list; limit [0,50]</p>
}

type GetSspInfoResponse struct {
	BaseResponse `json:",inline"`       // Common response fields
	Response     GetSspInfoResponseData `json:"response"` // Actual response data
}

type GetSspInfoResponseData struct {
	SspList []ResponseDataSsp `json:"ssp_list"` // [Required]
}

type GetSspListRequest struct {
	PageSize *int64  `json:"page_size,omitempty" url:"page_size,omitempty"` // [Optional] <p>The size of one page.</p><p>The limit of page_size is [1,20], and default page_size is 10.<br /></p>
	Offset   *string `json:"offset,omitempty" url:"offset,omitempty"`       // [Optional] <p>Specifies the starting entry of data to return in the current call. Default is "", if data is more than one page, the offset can be some entry to start next call.<br /></p>
}

type GetSspListResponse struct {
	BaseResponse `json:",inline"`       // Common response fields
	Response     GetSspListResponseData `json:"response"` // Actual response data
}

type GetSspListResponseData struct {
	PageInfo *PageInfo `json:"page_info"` // [Required]
	SspList  []Ssp     `json:"ssp_list"`  // [Required]
}

type GetStockAgingRequest struct {
	PageNo           *int64  `json:"page_no,omitempty"`            // [Optional] <p>Specifies the page number of data to return in the current call. Starting from 1. if data is more than one page, the page_no can be some entry to start next call. If empty, the default value is 1.<br /></p>
	PageSize         *int64  `json:"page_size,omitempty"`          // [Optional] <p>Each result set is returned as a page of entries. Use the "page_size" filters to control the maximum number of entries to retrieve per page (i.e., per call), and the "page_no" to start next call. This integer value is used to specify the maximum number of entries to return in a single "page" of data.</p><p>If empty, the default value is 10. The value should be between 1 and 100.</p>
	SearchType       *int64  `json:"search_type,omitempty"`        // [Optional] <p>1-Product Name；2-SKU ID；3-Variations；4-Item ID<br /></p>
	Keyword          *string `json:"keyword,omitempty"`            // [Optional] <p>bound with search_type<br /></p>
	WhsIds           *string `json:"whs_ids,omitempty"`            // [Optional] <p>split by comma<br /></p>
	AgingStorageTag  *int64  `json:"aging_storage_tag,omitempty"`  // [Optional] <p>0-false；1-true<br /></p>
	ExcessStorageTag *int64  `json:"excess_storage_tag,omitempty"` // [Optional] <p>0-false；1-true<br /></p>
	CategoryId       *int64  `json:"category_id,omitempty"`        // [Optional] <p>L1-level product category ID. You need to call the get_category API to obtain the first-level category_id<br /></p>
	WhsRegion        string  `json:"whs_region"`                   // [Required] <p>BR、CN、ID、MY、MX、TH、TW、PH、VN、SG<br /></p><p><b>If do not pass, will get error "block by gateway due to invalid cid"</b></p>
}

type GetStockAgingResponse struct {
	BaseResponse `json:",inline"`          // Common response fields
	Response     GetStockAgingResponseData `json:"response"` // Actual response data
}

type GetStockAgingResponseData struct {
	ItemList []GetStockAgingResponseDataItem `json:"item_list"` // [Required] <p>Data list of item sku<br /></p>
}

type GetStockAgingResponseDataItem struct {
	WarehouseItemId string                `json:"warehouse_item_id"` // [Required] <p>Warehouse item id; To indicate an unique item in a warehouse<br />one warehouse item id can match with multiple shop_item_id<br /><br /><b>For Global Item,&nbsp;warehouse_item_id=Global Item id<br />For Local Item, shop_item_id=item_id</b></p>
	ItemName        string                `json:"item_name"`         // [Required] <p>item name<br /></p>
	ItemImage       string                `json:"item_image"`        // [Required] <p>item image<br /></p>
	SkuList         []ResponseDataItemSku `json:"sku_list"`          // [Required] <p>Data list of mtsku<br /></p>
}

type GetStockMovementRequest struct {
	PageNo       *int64  `json:"page_no,omitempty"`        // [Optional] <p>Specifies the page number of data to return in the current call. Starting from 1. if data is more than one page, the page_no can be some entry to start next call. If empty, the default value is 1.<br /></p>
	PageSize     *int64  `json:"page_size,omitempty"`      // [Optional] <p>Each result set is returned as a page of entries. Use the "page_size" filters to control the maximum number of entries to retrieve per page (i.e., per call), and the "page_no" to start next call. This integer value is used to specify the maximum number of entries to return in a single "page" of data.</p><p>If empty, the default value is 10. The value should be between 1 and 20.<br /></p>
	StartTime    string  `json:"start_time"`               // [Required] <p>Start date in YYYY-MM-DD format. Only data within the past 1 year can be queried, and the time range must not exceed 90 days.<br /></p>
	EndTime      string  `json:"end_time"`                 // [Required] <p>End date in YYYY-MM-DD format. Only data within the past 1 year can be queried, and the time range must not exceed 90 days.<br /></p>
	WhsIds       *string `json:"whs_ids,omitempty"`        // [Optional] <p>Multiple warehouse_id values should be separated by commas.<br /></p>
	CategoryIdL1 *int64  `json:"category_id_l1,omitempty"` // [Optional] <p>L1-level category_id. You need to call the get_category API to retrieve the first-level category_id.<br /></p>
	SkuId        *string `json:"sku_id,omitempty"`         // [Optional]
	ItemId       *string `json:"item_id,omitempty"`        // [Optional]
	ItemName     *string `json:"item_name,omitempty"`      // [Optional] <p>Product Name Filter<br /></p>
	Variation    *string `json:"variation,omitempty"`      // [Optional]
	WhsRegion    string  `json:"whs_region"`               // [Required] <p>Warehouse Region. Enum values: BR, CN, ID, MY, MX, TH, TW, PH, VN, SG<br /></p><p><b>If do not pass, will get error "block by gateway due to invalid cid"</b></p>
}

type GetStockMovementResponse struct {
	BaseResponse `json:",inline"`             // Common response fields
	Response     GetStockMovementResponseData `json:"response"` // Actual response data
}

type GetStockMovementResponseData struct {
	Total        int64                              `json:"total"`          // [Required]
	StartTime    string                             `json:"start_time"`     // [Required]
	EndTime      string                             `json:"end_time"`       // [Required]
	QueryEndTime string                             `json:"query_end_time"` // [Required]
	ItemList     []GetStockMovementResponseDataItem `json:"item_list"`      // [Required] <p>Data list of item sku<br /></p>
}

type GetStockMovementResponseDataItem struct {
	WarehouseItemId string                                `json:"warehouse_item_id"` // [Required] <p>Warehouse item id; To indicate an unique item in a warehouse<br /><br />one warehouse item id can match with multiple shop_item_id<br /></p>
	ItemName        string                                `json:"item_name"`         // [Required] <p>item name<br /></p>
	ItemImage       string                                `json:"item_image"`        // [Required] <p>item image<br /></p>
	SkuList         []GetStockMovementResponseDataItemSku `json:"sku_list"`          // [Required] <p>Data list of mtsku<br /></p>
}

type GetStockMovementResponseDataItemSku struct {
	MtskuId            string                   `json:"mtsku_id"`             // [Required] <p>mtsku id<br /></p>
	ModelId            string                   `json:"model_id"`             // [Required] <p>model sku id<br /></p>
	Variation          string                   `json:"variation"`            // [Required]
	FulfillMappingMode int64                    `json:"fulfill_mapping_mode"` // [Required] <p>0-Null；1-Bundle SKU；2-Parent SKU<br /></p>
	Barcode            string                   `json:"barcode"`              // [Required]
	WhsList            []ResponseDataItemSkuWhs `json:"whs_list"`             // [Required] <p>Info of whs<br /></p>
	StartQty           *StartQty                `json:"start_qty"`            // [Required] <p>Inventory information at the start time.<br /></p>
	EndQty             *EndQty                  `json:"end_qty"`              // [Required]
	InboundQty         *InboundQty              `json:"inbound_qty"`          // [Required] <p>Inbound information during the selected time period<br /></p>
	OutboundQty        *OutboundQty             `json:"outbound_qty"`         // [Required]
	AdjustQty          *AdjustQty               `json:"adjust_qty"`           // [Required] <p>"SKU change information during the selected time period."</p><p> </p><p><br /></p>
	ShopSkuList        []ShopSku                `json:"shop_sku_list"`        // [Required]
}

type GetTimeSlotIdRequest struct {
	StartTime int64 `json:"start_time"` // [Required] <pre><span style="font-family:Roboto, &quot;Helvetica Neue&quot;, Helvetica, &quot;Droid Sans&quot;, Arial, sans-serif;">min = now, max=2145887999, should be &lt; end_time</span><br /></pre>
	EndTime   int64 `json:"end_time"`   // [Required] <pre><span style="font-family:Roboto, &quot;Helvetica Neue&quot;, Helvetica, &quot;Droid Sans&quot;, Arial, sans-serif;">should be &gt; start_time, max=2145887999</span><br /></pre>
}

type GetTimeSlotIdResponse struct {
	BaseResponse `json:",inline"`          // Common response fields
	Response     GetTimeSlotIdResponseData `json:"response"` // Actual response data
}

type GetTimeSlotIdResponseData struct {
	TimeslotId int64 `json:"timeslot_id"` // [Required]
	StartTime  int64 `json:"start_time"`  // [Required] <p>the start time of time slot</p>
	EndTime    int64 `json:"end_time"`    // [Required] <p>the end time of time slot<br /></p>
}

type GetTokenByResendCodeRequest struct {
	ResendCode string `json:"resend_code"` // [Required] the code in redirect url after you resend code in shop authorization management page. valid for one-time use, expires in 10minutes.
}

type GetTokenByResendCodeResponse struct {
	BaseResponse   `json:",inline"` // Common response fields
	ShopIdList     []int64          `json:"shop_id_list,omitempty"`     // Return when resend code in shop module
	MerchantIdList []int64          `json:"merchant_id_list,omitempty"` // Return when resend code in merchant module
	RefreshToken   string           `json:"refresh_token,omitempty"`    // Use refresh_token to obtain new access_token. Valid for each shop_id and merchant_id respectively one-time use, expires in 30 days.
	AccessToken    string           `json:"access_token,omitempty"`     // The token for API access, using to identify your permission to the api. Valid for multiple use and expires in 4 hours.
	ExpireIn       int64            `json:"expire_in,omitempty"`        // Access_token expiration time, unit is second.
}

type GetTopPicksListResponse struct {
	BaseResponse `json:",inline"`            // Common response fields
	Response     GetTopPicksListResponseData `json:"response"` // Actual response data
}

type GetTopPicksListResponseData struct {
	CollectionList []Collection `json:"collection_list"` // [Required] The top picks list in this shop.
}

type GetTotalBalanceResponse struct {
	BaseResponse `json:",inline"`            // Common response fields
	Response     GetTotalBalanceResponseData `json:"response"` // Actual response data
}

type GetTotalBalanceResponseData struct {
	DataTimestamp int64   `json:"data_timestamp"` // [Required] <p>This is param to indicate the time of the snapshot of total balance<br /></p>
	TotalBalance  float64 `json:"total_balance"`  // [Required] <p>This is seller's ads credit balance, including paid credits and free credits.<br /></p>
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

type GetVariationsRequest struct {
	CategoryId int64 `json:"category_id" url:"category_id"` // [Required] <p>Leaf category id</p>
}

type GetVariationsResponse struct {
	BaseResponse `json:",inline"` // Common response fields
	Data         *Data            `json:"data,omitempty"` // <p>standardized tier variation data<br /></p>
}

type GetVehicleListByCompatibilityDetailRequest struct {
	CompatibilityDetails string  `json:"compatibility_details" url:"compatibility_details"` // [Required] <p>To inform compatibility list, can be equal to Brand, Model, Year, or Version.<br /></p><p><br /></p><p>Pass the&nbsp;<b>compatibility_details="Brand"</b> to get all brand list;</p><p>Pass the <b>compatibility_details="Model" and brand_id=1234</b> to get all model list under brand_id=1234;<br /></p><p>Pass the&nbsp;<b>compatibility_details="Year" and brand_id=1234 and model_id=2345</b> to get all year list under brand_id=1234 and model_id=2345;</p><p>Pass the&nbsp;<b>compatibility_details="Version" and brand_id=1234 and model_id=2345 and year_id=3456</b> to get all version list under brand_id=1234 and model_id=2345 and year_id=3456.<br /></p>
	BrandId              *int64  `json:"brand_id,omitempty" url:"brand_id,omitempty"`       // [Optional] <p>ID of the brand.<br /></p>
	ModelId              *int64  `json:"model_id,omitempty" url:"model_id,omitempty"`       // [Optional] <p>ID of the model.<br /></p>
	YearId               *int64  `json:"year_id,omitempty" url:"year_id,omitempty"`         // [Optional] <p>ID of the year.<br /></p>
	Language             *string `json:"language,omitempty" url:"language,omitempty"`       // [Optional] <p>If language is not uploaded, the default language=en, the following are the languages supported by different markets SG: en ; MY: en / ms-my / zh-hans ; TH: en / th ; VN: en / vi ; PH: en ; TW: en / zh-hant ; ID: en / id ; BR: en / pt-br ; MX: en / es-mx ; CO: en/es-CO ; CL: en/es-CL. Note: For markets that have already launched global tree, Crossboard shop only support returning en and zh-hans language data.<br /></p>
}

type GetVehicleListByCompatibilityDetailResponse struct {
	BaseResponse `json:",inline"`                                // Common response fields
	Response     GetVehicleListByCompatibilityDetailResponseData `json:"response"` // Actual response data
}

type GetVehicleListByCompatibilityDetailResponseData struct {
	VehicleList []Vehicle `json:"vehicle_list"` // [Required]
}

type GetVideoUploadResultRequest struct {
	VideoUploadId string `json:"video_upload_id" url:"video_upload_id"` // [Required]
}

type GetVideoUploadResultResponse struct {
	BaseResponse `json:",inline"`                 // Common response fields
	Response     GetVideoUploadResultResponseData `json:"response"` // Actual response data
}

type GetVideoUploadResultResponseData struct {
	Status    string                 `json:"status"`     // [Required] Current status of this video upload session. could be: INITIATED(waiting for part uploading and/or the complete_video_upload API call), TRANSCODING(has received all video parts, and is transcoding the video file), SUCCEEDED(transcoding completed, and this upload_id can now be used for item adding/updating), FAILED(this upload failed, see the message filed for some info), CANCELLED(this upload is cancelled)
	VideoInfo *ResponseDataVideoInfo `json:"video_info"` // [Required] Transcoded video info, will be present if status is SUCCEEDED.
	Message   string                 `json:"message"`    // [Required] Detail error message if video uploading/transcoding failed.
}

type GetVideoUploadResultResponseDataVideoInfo struct {
	VideoUrl          string `json:"video_url"`           // [Required] <p>Video playback url.</p>
	VideoThumbnailUrl string `json:"video_thumbnail_url"` // [Required] <p>Video thumbnail image url.</p>
	ThumbnailWidth    int64  `json:"thumbnail_width"`     // [Required] <p>Video thumbnail image width.</p>
	ThumbnailHeight   int64  `json:"thumbnail_height"`    // [Required] <p>Video thumbnail image width.</p>
	Duration          int64  `json:"duration"`            // [Required] <p>Video duration in seconds.</p>
	Resolution        string `json:"resolution"`          // [Required] <p>Video resolution, e.g., "1280x1280".</p>
}

type GetVoucherListRequest struct {
	PageNo   *int64 `json:"page_no,omitempty"`   // [Optional] Specifies the page number of data to return in the current call. Default to be 1 and allowed input is from 1 - 5000.
	PageSize *int64 `json:"page_size,omitempty"` // [Optional] Use the 'page_size' filters to control the maximum number of entries to retrieve per page (i.e., per call). Default to be 20 and allowed input is from 1- 100.
	Status   string `json:"status"`              // [Required] The status filter for retrieving voucher list. Available value: upcoming/ongoing/expired/all.
}

type GetVoucherListResponse struct {
	BaseResponse `json:",inline"`           // Common response fields
	Response     GetVoucherListResponseData `json:"response"` // Actual response data
}

type GetVoucherListResponseData struct {
	More        bool      `json:"more"`         // [Required] This is to indicate whether the comment list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of comments.
	VoucherList []Voucher `json:"voucher_list"` // [Required] The list of voucher.
}

type GetVoucherRequest struct {
	VoucherId int64 `json:"voucher_id"` // [Required] The unique identifier of a voucher used to query the voucher details.
}

type GetVoucherResponse struct {
	BaseResponse `json:",inline"`       // Common response fields
	Response     GetVoucherResponseData `json:"response"` // Actual response data
}

type GetVoucherResponseData struct {
	VoucherId          int64       `json:"voucher_id"`           // [Required] The unique identifier of the voucher whose details are returned.
	VoucherCode        string      `json:"voucher_code"`         // [Required] Voucher Code
	VoucherName        string      `json:"voucher_name"`         // [Required] Voucher Name
	VoucherType        int64       `json:"voucher_type"`         // [Required] The type of the voucher. The available values are: 1: shop voucher, 2: product voucher.
	RewardType         int64       `json:"reward_type"`          // [Required] The reward type of the voucher. The available values are: 1: fix_amount voucher, 2: discount_percentage voucher, 3: coin_cashback voucher.
	UsageQuantity      int64       `json:"usage_quantity"`       // [Required] The number of times for this particular voucher could be used.
	CurrentUsage       int64       `json:"current_usage"`        // [Required] Up till now, how many times has this particular voucher already been used.
	StartTime          int64       `json:"start_time"`           // [Required] The timing from when the voucher is valid; so buyer is allowed to claim and to use.
	EndTime            int64       `json:"end_time"`             // [Required] The timing until when the voucher is still valid. Any time after this end_time, buyer is not allowed to claim or to use.
	IsAdmin            bool        `json:"is_admin"`             // [Required] If the voucher is created by Shopee or not.
	VoucherPurpose     int64       `json:"voucher_purpose"`      // [Required] <p>The use case for the voucher. The available values are: 0: normal; 1: welcome, 2: referral; 3: shop_follow; 4:shop_game, 5: free_gift, 6: membership，7: Ads</p>
	DisplayChannelList interface{} `json:"display_channel_list"` // [Required] The FE channel where the voucher will be displayed. The available values are: 1: display_all, 2: order page, 3: feed, 4: live streaming,   [] (empty - which is hidden).
	MinBasketPrice     float64     `json:"min_basket_price"`     // [Required] The minimum spend required for using this voucher.
	Percentage         int64       `json:"percentage"`           // [Required] <p>The discount percentage is set for this particular voucher. Only when it is a discount percentage voucher or coins cashback voucher, api will return a value.</p>
	MaxPrice           float64     `json:"max_price"`            // [Required] The max amount of discount/value a user can enjoy by using this particular voucher. Only when it is a discount percentage voucher or coins cashback voucher, api will return a value.
	DiscountAmount     float64     `json:"discount_amount"`      // [Required] The discount amount set for this particular voucher. Only when it is a fix amount voucher, api will return a value.
	CmtVoucherStatus   int64       `json:"cmt_voucher_status"`   // [Required] The voucher status in CMT. The available values are: 1:review, 2: approved, 3:reject. Only when this voucher is attending CMT campaign and not being rejected, api will return a value.
	ItemIdList         []int64     `json:"item_id_list"`         // [Required] The list of items which is applicable for the voucher. Only return a value when it is a product type of voucher.
	DisplayStartTime   int64       `json:"display_start_time"`   // [Required] The timing of when voucher is displayed on shop pages; so buyer is allowed to claim.
	TargetVoucher      int64       `json:"target_voucher"`       // [Required] <p>The field is used to mark new user voucher/ repeat buyer voucher&nbsp;</p><p>1: new user voucher&nbsp;</p><p>2: repeat buyer voucher: with 1 orders</p><p>3. repeat buyer voucher: with 2 orders<br /></p>
	Usecase            int64       `json:"usecase"`              // [Required] <p>usecase indicates a specific business scenario that the voucher is created and used for, i.e., new buyer voucher, live voucher, follow shop voucher, etc.</p><p>shop voucher:1</p><p>product voucher:2</p><p>new buyer voucher:3</p><p>repeat buyer voucher:4</p><p>private voucher:5</p><p>live voucher:6</p><p>video voucher:7</p><p>campaign voucher:8</p><p>follow prize voucher:9</p><p>membership voucher:10</p><p>game prize voucher:11</p><p>sample voucher:12<br /></p>
}

type GetWalletTransactionListRequest struct {
	PageNo             int64   `json:"page_no"`                        // [Required] Specifies the starting entry of data to return in the current call. Default is 0. if data is more than one page, the offset can be some entry to start next call.
	PageSize           int64   `json:"page_size"`                      // [Required] If many transactions are available to retrieve, you may need to call GetTransactionList multiple times to retrieve all the data. Each result set is returned as a page of entries. Default is 40. Use the Pagination filters to control the maximum number of entries (<= 100) to retrieve per page (i.e., per call), the offset number to start next call. This integer value is usUed to specify the maximum number of entries to return in a single ""page"" of data.
	CreateTimeFrom     *int64  `json:"create_time_from,omitempty"`     // [Optional] The create_time_from field is the starting date range. The maximum date range that may be specified with the create_time_from and create_time_to fields is 15 days.
	CreateTimeTo       *int64  `json:"create_time_to,omitempty"`       // [Optional] The create_time_to field is the ending date range. The maximum date range that may be specified with the create_time_from and create_time_to fields is 15 days.
	WalletType         *string `json:"wallet_type,omitempty"`          // [Optional] This field indicates the wallet type.
	TransactionType    *string `json:"transaction_type,omitempty"`     // [Optional] <p>Transaction type APIs:&nbsp;<br /></p><p>ESCROW_VERIFIED_ADD = 101;&nbsp; // Escrow has been verified and paid to seller&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</p><p>ESCROW_VERIFIED_MINUS = 102; // Escrow has been verified and charged from seller as escrow amount is negative&nbsp;&nbsp;&nbsp;&nbsp;</p><p>WITHDRAWAL_CREATED = 201; // The seller has created a withdrawal, so it’s deducted from balance &nbsp; &nbsp; &nbsp;<br /></p><p>WITHDRAWAL_COMPLETED = 202; // The withdrawal has been completed, so the ongoing amount decreases.&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</p><p>WITHDRAWAL_CANCELLED = 203; // The withdrawal has been canceled, so the amount is added back to the seller balance. Ongoing amount decreases as well.&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</p><p>ADJUSTMENT_ADD = 401; // One adjustment item has been paid to seller&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;ADJUSTMENT_MINUS = 402; // One adjustment item has been charged from seller&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;FBS_ADJUSTMENT_ADD = 404; //One adjustment item related to Shopee fulfillment order is added to seller&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</p><p>FBS_ADJUSTMENT_MINUS = 405; // One adjustment item related to Shopee fulfillment order is deducted from seller&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</p><p>ADJUSTMENT_CENTER_ADD = 406; // One adjustment item has been added to seller wallet&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;ADJUSTMENT_CENTER_DEDUCT = 407; // One adjustment item has been deducted from seller wallet&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</p><p>FSF_COST_PASSING_DEDUCT = 408; FSF cost passing for canceled/invalid orders&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;PERCEPTION_VAT_TAX_DEDUCT = 409; Extra charge for perception regime VAT tax (Argentina)&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;PERCEPTION_TURNOVER_TAX_DEDUCT = 410; Extra charge for perception regime turnover tax (Argentina)&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</p><p>PAID_ADS_CHARGE = 450; // Paid ads are charged from seller&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</p><p>PAID_ADS_REFUND = 451; // Paid ads are refunded to seller&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</p><p>FAST_ESCROW_DISBURSE = 452; // ADD. // The first disbursement of fast escrow has been paid to seller&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</p><p>AFFILIATE_ADS_SELLER_FEE = 455; // DEDUCT // Affiliate ads seller fee is charged from seller&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;AFFILIATE_ADS_SELLER_FEE_REFUND = 456; // ADD // Affiliate ads seller fee is refunded to seller</p><p>FAST_ESCROW_DEDUCT = 458; // Fast escrow is deducted from seller balance in the event of return and refund&nbsp;</p><p>FAST_ESCROW_DISBURSE_REMAIN = 459; // The second disbursement of fast escrow has been paid to seller&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</p><p>AFFILIATE_FEE_DEDUCT = 460; // Affiliate MKT fee is charged from seller for using affiliate MKT services<br /></p>
	MoneyFlow          *string `json:"money_flow,omitempty"`           // [Optional] <p>It's to indicate whether user wants to only return :&nbsp;</p><p>MONEY_IN = addition&nbsp;<br />MONEY_OUT = Deduction</p><p><br /></p><p>if not specified, we will return all</p><p>Note special case for TW JKO Pay, we will ignore Money_flow&nbsp;</p>
	TransactionTabType *string `json:"transaction_tab_type,omitempty"` // [Optional] <p><b>NOTE: Only 1 'transaction tab type' value should be passed in.</b></p><p>Passing in more than 1 value (eg: comma separated values) will return default response. This is because the request param treats the value passed in as a single string.</p><p><br /></p><p>This to indicates the updated filtering type that client can use to specify which transaction type we want to return. it will have :&nbsp;<br /><br />Default<br />wallet_order_income<br />wallet_adjustment_filter</p><p>wallet_wallet_payment</p><p>wallet_refund_from_order</p><p>wallet_withdrawals</p><p>fast_escrow_repayment</p><p>fast_pay</p><p>seller_loan<br />corporate_loan<br />pix_transactions_filter</p><p>open_finance_transactions_filter&nbsp;<br /><br />Note for BR, wallet txn type that linked to&nbsp;pix_transactions_filter&nbsp; and&nbsp;open_finance_transactions_filter are classified as&nbsp;default&nbsp;&nbsp;type tab instead. therefore for Open API client who wants to query these 2 trx can put default as the filter in this type<br /></p>
}

type GetWalletTransactionListResponse struct {
	BaseResponse `json:",inline"`                     // Common response fields
	Response     GetWalletTransactionListResponseData `json:"response"` // Actual response data
}

type GetWalletTransactionListResponseData struct {
	TransactionList []Transaction `json:"transaction_list"` // [Required]
	More            bool          `json:"more"`             // [Required]
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

type GetWarehouseEligibleShopListRequest struct {
	WarehouseId   int64   `json:"warehouse_id"`   // [Required] <p>Warehouse address identifier.<br /></p>
	WarehouseType int64   `json:"warehouse_type"` // [Required] <p>1 means pickup warehouse</p><p>2 means return warehouse<br /></p>
	Cursor        *Cursor `json:"cursor"`         // [Required] <p>// how to use DoubleSidedCursor</p><p>// Get data for the first page: Please pass next_id = 0 or nil, page_size = {your page size}.</p><p>// Get data for the next page: Please pass the Cursor from the previous response, and set prev_id=nil;</p><p>// Get data for the prev page: Please pass the Cursor from the previous response, and set next_id=nil;</p><p>// Stop fetching next data: The Cursor.next_id in the previous response is nil.</p><p>// Stop fetching prev data: The Cursor.prev_id in the previous response is nil.<br /></p>
}

type GetWarehouseEligibleShopListResponse struct {
	BaseResponse `json:",inline"`                         // Common response fields
	Response     GetWarehouseEligibleShopListResponseData `json:"response"` // Actual response data
}

type GetWarehouseEligibleShopListResponseData struct {
	ShopList []ResponseDataShop `json:"shop_list"` // [Required] <p>Eligible shop list of the warehouse</p>
	Cursor   *Cursor            `json:"cursor"`    // [Required]
}

type GetWarehouseFilterConfigResponse struct {
	BaseResponse `json:",inline"`                     // Common response fields
	Response     GetWarehouseFilterConfigResponseData `json:"response"` // Actual response data
}

type GetWarehouseFilterConfigResponseData struct {
	WarehouseFilters []WarehouseFilters `json:"warehouse_filters"` // [Required]
}

type GetWaybillRequest struct {
	FirstMileTrackingNumberList []string `json:"first_mile_tracking_number_list"` // [Required] The first mile tracking number that you want to print waybill.limit [1, 50]
}

type GetWaybillResponse struct {
	BaseResponse `json:",inline"` // Common response fields
	Waybill      interface{}      `json:"waybill,omitempty"` // <p>The waybill file.</p>
}

type GetWeightRecommendationRequest struct {
	ItemName        string           `json:"item_name"`                  // [Required] <p>Name of the item in local language.<br /></p>
	CoverImageId    string           `json:"cover_image_id"`             // [Required] <p>Image id of first product image.&nbsp;<br /></p>
	CategoryId      int64            `json:"category_id"`                // [Required] <p>Shopee's unique identifier for a category.<br /></p>
	AttributeList   []Attribute      `json:"attribute_list"`             // [Required]
	BrandId         int64            `json:"brand_id"`                   // [Required] <p>Id of brand.<br /></p>
	DescriptionType DescriptionType  `json:"description_type"`           // [Required] <p>Type of description, values: See Data Definition- description_type (normal , extended).<br /></p>
	Description     *string          `json:"description,omitempty"`      // [Optional] <p>If description_type is normal , Description information should be set by this field.<br /></p>
	DescriptionInfo *DescriptionInfo `json:"description_info,omitempty"` // [Optional] <p>New description field. Only whitelist sellers can use it. If you use the field, please upload the description_type=extended.<br /></p>
}

type GetWeightRecommendationResponse struct {
	BaseResponse `json:",inline"`                    // Common response fields
	Response     GetWeightRecommendationResponseData `json:"response"` // Actual response data
}

type GetWeightRecommendationResponseData struct {
	NormalWeightRange interface{} `json:"normal_weight_range"` // [Required] <p>Recommended weight range, in kg. If there are no recommended results, return empty.<br /></p>
}

type GlobalCode struct {
	Gtin string `json:"gtin"` // [Required] <p>GTIN of Shopee&nbsp;Standard Product.</p>
	Oem  string `json:"oem"`  // [Required] <p>OEM of Shopee&nbsp;Standard Product.<br /></p>
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

type GroupItemInfo struct {
	GroupQtd           string `json:"group_qtd"`            // [Required] <p>Example: The package contains 6 soda cans. Whether you are selling a pack of 6 cans (fardo) or a single can (unit), enter 6.<br /></p>
	GroupUnit          string `json:"group_unit"`           // [Required] <p>Example: The package contains 6 soda cans. Whether you are selling a pack of 6 cans (fardo) or a single can (unit), enter UNI for the individual can.<br /></p>
	GroupUnitValue     string `json:"group_unit_value"`     // [Required] <p>Example: The package contains 6 soda cans. Whether you are selling a pack of 6 cans (fardo) or a single can (unity), enter the value of the individual can.<br /></p>
	OriginalGroupPrice string `json:"original_group_price"` // [Required] <p>Example: The item is a package that contains 6 soda cans. Enter the price of the whole package.<br /></p>
	GroupGtinSscc      string `json:"group_gtin_sscc"`      // [Required] <p>Example: The item is a package that contains 6 soda cans. Please inform the GTIN SSCC code for the package.<br /></p>
	GroupGraiGtinSscc  string `json:"group_grai_gtin_sscc"` // [Required] <p>Example: The item is box, that contain 6 packages. Each package contains 6 soda cans. Please inform the GRAI GTIN SSCC code for the Box.<br /></p>
}

type GtinLimit struct {
	GtinValidationRule string `json:"gtin_validation_rule"` // [Required] <p>Indicate gtin_code validation logic in&nbsp;</p><p>v2.product.add_item</p><p>v2.product.update_item<br />v2.product.init_tier_variation<br />v2.product.add_model<br />v2.product.update_model</p><p>- <b>Mandatory</b>: This field is required and must contain a correctly formatted GTiN number.</p><p>- <b>Flexible</b>: This field is required and must contain either a correctly formatted GTlN number or "00" todeclare that the item/model has no valid GTlN.<br />- <b>Optional</b>: This field is optional and can contain a correctly formatted GTiN number, "00" or be omittedentirely.&nbsp;&nbsp;</p>
}

type HandleBuyerCancellationRequest struct {
	OrderSn   string `json:"order_sn"`  // [Required] Shopee's unique identifier for an order.
	Operation string `json:"operation"` // [Required] The operation you want to handle.Avaiable value: ACCEPT, REJECT
}

type HandleBuyerCancellationResponse struct {
	BaseResponse `json:",inline"`                    // Common response fields
	Response     HandleBuyerCancellationResponseData `json:"response"` // Actual response data
}

type HandleBuyerCancellationResponseData struct {
	UpdateTime int64 `json:"update_time"` // [Required] The time when the order is updated.
}

type HandlePrescriptionCheckRequest struct {
	OrderSn          string  `json:"order_sn"`                     // [Required] <p>Shopee's unique identifier for an order.<br /></p>
	IsApproved       bool    `json:"is_approved"`                  // [Required] <p>Approve or reject the prescription. Available values: TRUE, FALSE.<br /></p>
	RejectReasonCode *int64  `json:"reject_reason_code,omitempty"` // [Optional] <p>Reject reason code. Available values:&nbsp;<br /></p><p>1 = Invalid Prescription (counterfeit/incorrect format)</p><p>2 = Incorrect Dosage</p><p>3 = No Prescription</p><p>4 = Unclear Image</p>
	Items            []Items `json:"items,omitempty"`              // [Optional] <p>The list of invalid items that make the prescription get rejected<br /></p>
	PharmacistName   *string `json:"pharmacist_name,omitempty"`    // [Optional] <p>Full name of the pharmacist.&nbsp;Required for PH and ID Prescription Orders.</p>
}

type HandlePrescriptionCheckResponse struct {
	BaseResponse `json:",inline"`                    // Common response fields
	Response     HandlePrescriptionCheckResponseData `json:"response"` // Actual response data
}

type HandlePrescriptionCheckResponseData struct {
	IsSuccess bool `json:"is_success"` // [Required] <p>This is to indicate whether the request has been executed successfully.<br /></p>
}

type HdListingList struct {
	ItemId        int64 `json:"item_id"`        // [Required] <p>Item ID.</p>
	CurrentStatus int64 `json:"current_status"` // [Required] <p>For 2030: % HD Listings, it refer to Current HD Status.</p><p>For 2031: % HD Free Shipping Enabled, it refer to Free Shipping Enabled Status.</p><p><br /></p><p>Applicable values:&nbsp;</p><p>1: Yes</p><p>2: No</p>
}

type Image struct {
	ImageIdList []string `json:"image_id_list"`         // [Required] ID of image
	ImageRatio  *string  `json:"image_ratio,omitempty"` // [Optional] <p>Ratio of image, <br />OptionalAllowed ratios :<br />"1:1" (default)&nbsp;<br />"3:4"<br /></p><p>only applicable to whitelisted seller.<br /></p>
}

type Images struct {
	ImageIdList  []string `json:"image_id_list"`  // [Required] List of image url.
	ImageUrlList []string `json:"image_url_list"` // [Required] List of image id.
}

type InboundQty struct {
	InboundTotal    int64 `json:"inbound_total"`    // [Required] <p>Total inbound quantity during the selected time period<br /></p>
	InboundMy       int64 `json:"inbound_my"`       // [Required] <p>Total merchant procurement quantity during the selected time period.<br /></p>
	InboundReturned int64 `json:"inbound_returned"` // [Required] <p>Total number of SKUs returned by buyers and received into the warehouse during the selected time period.<br /></p>
}

type IncomeDetailList struct {
	NextPage             *Pagination           `json:"next_page"`               // [Required] <p>Contains pagination metadata for fetching the next page.</p>
	IncomeDetailListItem *IncomeDetailListItem `json:"income_detail_list_item"` // [Required] <p>List of income detail objects</p>
}

type IncomeDetailListItem struct {
	PaymentMethod         string  `json:"payment_method"`          // [Required] <p>Payment channel or method used for the order</p>
	OrderSn               string  `json:"order_sn"`                // [Required] <p>Unique order serial number associated with the income record.</p>
	Description           string  `json:"description"`             // [Required] <p>Type of income or billing item — e.g., Order Income, Adjustment etc</p>
	Status                string  `json:"status"`                  // [Required] <p>Status description of the order income or payout.</p>
	Currency              string  `json:"currency"`                // [Required] <p>Currency in which the income was transacted.</p>
	EstimatedEscrowAmount float64 `json:"estimated_escrow_amount"` // [Required] <p>Estimated escrow amount pending release for the order.</p>
	EstimatedPayoutTime   int64   `json:"estimated_payout_time"`   // [Required] <p>Estimated payout time (Unix timestamp). Applicable for Pending/To Release status.</p>
	ToReleaseAmount       float64 `json:"to_release_amount"`       // [Required] <p>Amount that is queued for release to seller (Cross Border only).</p>
	CreationDate          int64   `json:"creation_date"`           // [Required] <p>Order creation timestamp (Unix format).</p>
	ReleasedAmount        float64 `json:"released_amount"`         // [Required] <p>Amount successfully released to the seller.</p>
	ActualPayoutTime      int64   `json:"actual_payout_time"`      // [Required] <p>Actual payout time (Unix timestamp) when funds were transferred.</p>
}

type InfoNeeded struct {
	Dropoff       []string `json:"dropoff"`        // [Required] <p>Could contain 'branch_id', 'sender_real_name', or 'tracking_no'. If it contains 'branch_id', choose one to Init. If it contains 'sender_real_name' or 'tracking_no', should manually input these values in Init API. If it has empty value, developer should still include "dropoff" field in Init API.</p>
	Pickup        []string `json:"pickup"`         // [Required] <p>Could contain 'address_id' and 'pickup_time_id'. Choose one address_id and its corresponding pickup_time_id to Init. If it has empty value, developer should still include "pickup" field in Init API. It could contains "tracking_number" returned from "info_need"for some channels, please also add it when init.</p>
	NonIntegrated []string `json:"non_integrated"` // [Required] <p>Could contain 'tracking_no'. If it contains 'tracking_no', should manually input these values in Init API. If it has empty value, developer should still include "non-integrated" field in Init API.</p>
}

type InitTierVariationRequest struct {
	ItemId                   int64                      `json:"item_id"`                              // [Required] ID of item
	Model                    []Model                    `json:"model"`                                // [Required] Model info list, model number at most 50
	StandardiseTierVariation []StandardiseTierVariation `json:"standardise_tier_variation,omitempty"` // [Optional] <p>There is at least one standardise_tier_variation and&nbsp;tier_variation.</p><p><br /></p><p>If you want to update one tier/two tier to no tier, can just pass the tier_variation and standardise_tier_variation as [], and pass the model &gt;&gt; tier_index as [], meanwhile pass the original_price, seller_stock, etc., to set the price and stock for the modified product with no tier structure.<br /></p>
}

type InitTierVariationResponse struct {
	BaseResponse `json:",inline"`              // Common response fields
	Response     InitTierVariationResponseData `json:"response"` // Actual response data
}

type InitTierVariationResponseData struct {
	ItemId        int64                       `json:"item_id"`        // [Required] ID of item
	TierVariation []ResponseDataTierVariation `json:"tier_variation"` // [Required] Variations of item
	Model         []ResponseDataModel         `json:"model"`          // [Required]
}

type InitVideoUploadRequest struct {
	FileMd5  string `json:"file_md5"`  // [Required] md5 of video file
	FileSize int64  `json:"file_size"` // [Required] size of video file, in bytes, maximum is 30MB
}

type InitVideoUploadResponse struct {
	BaseResponse `json:",inline"`            // Common response fields
	Response     InitVideoUploadResponseData `json:"response"` // Actual response data
}

type InitVideoUploadResponseData struct {
	VideoUploadId string `json:"video_upload_id"` // [Required] The identifier of this upload session, used in following video upload request and item creating and/or updating
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

type InvoiceData struct {
	Number             string  `json:"number"`               // [Required] The number of the invoice.
	SeriesNumber       string  `json:"series_number"`        // [Required] The series number of the invoice.
	AccessKey          string  `json:"access_key"`           // [Required] The access key of the invoice.
	IssueDate          int64   `json:"issue_date"`           // [Required] The issue date of the invoice.
	TotalValue         float64 `json:"total_value"`          // [Required] The total value of the invoice.
	ProductsTotalValue float64 `json:"products_total_value"` // [Required] The products total value of the invoice.
	TaxCode            string  `json:"tax_code"`             // [Required] The tax code for the invoice.
}

type InvoiceDetail struct {
	Name                    string                   `json:"name"`                      // [Required] <p>Buyer name (only has value when invoice_type is personal)<br /></p><p>- VN, TH, PH only<br /></p>
	Email                   string                   `json:"email"`                     // [Required] <p>Buyer email address&nbsp;(only has value when invoice_type is personal)<br /></p><p>- VN, TH, PH only<br /></p>
	PhoneNumber             string                   `json:"phone_number"`              // [Required] <p>Buyer phone number<br /></p><p>- TH only<br /></p>
	TaxId                   string                   `json:"tax_id"`                    // [Required] <p>only has value when invoice_type is personal.&nbsp;<br /></p><p>- VN, TH, PH only<br /></p>
	Address                 string                   `json:"address"`                   // [Required] <p>Buyer address in format "Street &amp; number, city, zipcode, any additional info provided by buyer"&nbsp;(only has value when invoice_type is personal)<br /></p><p>- PH, VN only<br /></p>
	IdCardAddress           string                   `json:"id_card_address"`           // [Required] <p>Same function as the address, only having a different field name for TH.Buyer address in format "Street &amp; number, city, zipcode, any additional info provided by buyer"&nbsp;(only has value when invoice_type is personal).<br /></p>
	AddressBreakdown        *AddressBreakdown        `json:"address_breakdown"`         // [Required] <p>Buyer address breakdown.<br /></p><p>- TH, PH only<br /></p>
	CompanyHeadOffice       string                   `json:"company_head_office"`       // [Required] <p>- return value for TH only&nbsp;(only has value when invoice_type is company)<br /></p>
	CompanyName             string                   `json:"company_name"`              // [Required] <p>- Only return value when invoice type is company</p><p>- VN, TH, PH only<br /></p>
	CompanyBranchName       string                   `json:"company_branch_name"`       // [Required] <p>- Only return value when invoice type is company</p><p>- TH only<br /></p>
	CompanyBranchId         string                   `json:"company_branch_id"`         // [Required] <p>- Only return value when invoice type is company</p><p>- TH only<br /></p>
	CompanyType             string                   `json:"company_type"`              // [Required] <p>- Only return value when invoice type is company</p><p>- TH only<br /></p>
	CompanyEmail            string                   `json:"company_email"`             // [Required] <p>- Only return value when invoice type is company</p><p>- VN, TH, PH only<br /></p>
	CompanyTaxId            string                   `json:"company_tax_id"`            // [Required] <p>- Only return value when invoice type is company</p><p>- VN, TH, PH only<br /></p>
	CompanyAddress          string                   `json:"company_address"`           // [Required] <p>Buyer address in format "Street &amp; number,city, zipcode, any additional info provided by buyer"&nbsp;(only has value when invoice_type is company)<br /></p><p>- VN, TH only<br /></p>
	CompanyAddressBreakdown *CompanyAddressBreakdown `json:"company_address_breakdown"` // [Required] <p>Company address breakdown</p><p>- PH, TH only<br /></p>
}

type InvoiceInfo struct {
	OrderSn       string         `json:"order_sn"`       // [Required] <p>Shopee's unique identifier for an order.<br /></p>
	InvoiceType   string         `json:"invoice_type"`   // [Required] <p>Type of invoice requested: {1: personal, 2: company}.<br /></p>
	InvoiceDetail *InvoiceDetail `json:"invoice_detail"` // [Required] <p>Invoice info submitted by buyer. Might be masked, e.g. A****b, depending on order status.<br /></p>
	Error         string         `json:"error"`          // [Required] <p>Error in retrieving the receipt setting of a particular order.<br /></p>
	IsRequested   bool           `json:"is_requested"`   // [Required] <p>To identify order with and without buyer request, applicable to PL.<br /></p>
}

type ItemAttribute struct {
	AttributeId           int64            `json:"attribute_id"`            // [Required] The Identify of each category.
	OriginalAttributeName string           `json:"original_attribute_name"` // [Required] The name of each attribute.
	IsMandatory           bool             `json:"is_mandatory"`            // [Required] This is to indicate whether this attribute is mandantory.
	AttributeValueList    []AttributeValue `json:"attribute_value_list"`    // [Required]
}

type ItemComment struct {
	OrderSn       string        `json:"order_sn"`       // [Required] Shopee's unique identifier for an order.
	CommentId     string        `json:"comment_id"`     // [Required] The identity of comment.
	Comment       string        `json:"comment"`        // [Required] The content of the comment.
	BuyerUsername string        `json:"buyer_username"` // [Required] The username of the buyer who posted the comment.
	ItemId        int64         `json:"item_id"`        // [Required] The commented item's id
	ModelId       int64         `json:"model_id"`       // [Required] <p>Shopee's unique identifier for a model of an item. It will only return 0 now.&nbsp;</p><p><b>Will be offline on <font color="#c24f4a">2024-12-27</font>, please switch to use model_id_list.</b></p>
	RatingStar    int64         `json:"rating_star"`    // [Required] Buyer's rating for the item.
	Editable      string        `json:"editable"`       // [Required] The editable status of the comment. The value may be one of  EXPIRED/EDITABLE/HAVE_EDIT_ONCE.
	Hidden        bool          `json:"hidden"`         // [Required] The comment is hidden or not.
	CreateTime    int64         `json:"create_time"`    // [Required] The create time of the comment.
	CommentReply  *CommentReply `json:"comment_reply"`  // [Required] The reply of the comment.
	ModelIdList   []int64       `json:"model_id_list"`  // [Required] <p>List of model id of the buyer's purchase corresponding to the comment.&nbsp;</p>
	Media         *Media        `json:"media"`          // [Required]
}

type ItemCountLimit struct {
	MaxLimit int64 `json:"max_limit"` // [Required] Item count max limit.
}

type ItemIdMap struct {
	ItemId       int64 `json:"item_id"`        // [Required] Id of item.
	GlobalItemId int64 `json:"global_item_id"` // [Required] Id of global item.
}

type ItemImage struct {
	ImageIdList  []string `json:"image_id_list"`  // [Required] <p>List of image id.<br /></p>
	ImageUrlList []string `json:"image_url_list"` // [Required] <p>List of image url.<br /></p>
	ImageRatio   string   `json:"image_ratio"`    // [Required] <p>Image ratio.<br /></p>
}

type ItemInfo struct {
	ItemId                int64                  `json:"item_id"`                  // [Required]
	ItemName              string                 `json:"item_name"`                // [Required]
	Status                int64                  `json:"status"`                   // [Required] <p>item status</p><p><br /></p><p>0: Deleted<br />1: Normal<br /></p><p>2: reviewing</p><p>3: banned</p><p>4: invalid</p><p>5: invalid hide</p><p>6: offensive hide</p><p>7: auditing</p><p>8: normal unlist</p>
	Image                 string                 `json:"image"`                    // [Required] <p>item image</p>
	ItemStatus            ItemStatus             `json:"item_status"`              // [Required] <p>the status of item in shop flash sale.&nbsp;If the item has variation, this field will be empty</p><p>0: disable<br /></p><p>1: enable</p><p>2: delete</p><p>4: system_rejected,&nbsp;the item is rejected by system</p><p>5: manual_rejected, the item is rejected manually</p>
	OriginalPrice         float64                `json:"original_price"`           // [Required] <p>original price of item, if the item has variation, this field will be empty</p>
	InputPromotionPrice   float64                `json:"input_promotion_price"`    // [Required] <p>promotion price without tax of item, if the item has variation, this field will be empty</p>
	PromotionPriceWithTax float64                `json:"promotion_price_with_tax"` // [Required] <p>promotion price with tax of item,&nbsp;if the item has no variation, this field will has value</p>
	PurchaseLimit         int64                  `json:"purchase_limit"`           // [Required] <p>0 means NO LIMIT</p><p><br /></p><p>purchase limit of item,&nbsp;if the item has variation, this field will be empty</p>
	CampaignStock         int64                  `json:"campaign_stock"`           // [Required] <p>campaign stock of item,&nbsp;if the item has no variation, this field will has value</p>
	Stock                 int64                  `json:"stock"`                    // [Required] <p>Active inventory&nbsp;of item,&nbsp;if the item has no variation, this field will has value</p>
	RejectReason          string                 `json:"reject_reason"`            // [Required] <p>if the item_status is 4 or 5, this field will show the reason why this item was rejected</p><p><br /></p><p>if the item has variation, this field will be empty</p>
	UnqualifiedConditions *UnqualifiedConditions `json:"unqualified_conditions"`   // [Required] <p>if the item doesn't meet a criteria, will show the detail in this field</p><p><br /></p><p>if the item has variation, this field will be empty</p>
}

type ItemInstallment struct {
	ItemId     int64   `json:"item_id"`     // [Required] Item unique id
	TenureList []int64 `json:"tenure_list"` // [Required] The tenures of item support installment. [] represents with no installment
}

type ItemLogisticInfo struct {
	LogisticId           int64   `json:"logistic_id"`            // [Required] The identity of logistic channel.
	LogisticName         string  `json:"logistic_name"`          // [Required] The name of logistic.
	Enabled              bool    `json:"enabled"`                // [Required] Related to shopee.logistics.GetLogistics result.logistics.enabled only affect current item.
	ShippingFee          float64 `json:"shipping_fee"`           // [Required] Only needed when logistics fee_type = CUSTOM_PRICE.
	SizeId               int64   `json:"size_id"`                // [Required] If specify logistic fee_type is SIZE_SELECTION size_id is required.
	IsFree               bool    `json:"is_free"`                // [Required] when seller chooses this option, the shipping fee of this channel on item will be set to 0. Default value is False.
	EstimatedShippingFee float64 `json:"estimated_shipping_fee"` // [Required] Estimated shipping fee calculated by weight. Don't exist if channel is no-integrated.
}

type ItemMapping struct {
	MartItemId   int64                     `json:"mart_item_id"`   // [Required] <p>The item ID of the item in the Mart shop.</p>
	OutletItemId int64                     `json:"outlet_item_id"` // [Required] <p>The item ID of the corresponding item in the outlet shop.</p>
	ModelMapping []ItemMappingModelMapping `json:"model_mapping"`  // [Required] <p>The mapping relationship between Mart models and outlet models under the mapped items.</p>
}

type ItemMappingModelMapping struct {
	MartModelId   int64 `json:"mart_model_id"`   // [Required] <p>The model ID of the product in the Mart shop.</p>
	OutletModelId int64 `json:"outlet_model_id"` // [Required] <p>The model ID of the corresponding product in the outlet shop.</p>
}

type ItemMaxDimension struct {
	Height       float64 `json:"height"`        // [Required] The max height limit.
	Width        float64 `json:"width"`         // [Required] The max width limit.
	Length       float64 `json:"length"`        // [Required] The max length limit.
	Unit         string  `json:"unit"`          // [Required] The unit for the limit.
	DimensionSum float64 `json:"dimension_sum"` // [Required] The sum of the item's dimension
}

type ItemModel struct {
	TierIndex     interface{} `json:"tier_index"`             // [Required] Tier index of model.
	OriginalPrice float64     `json:"original_price"`         // [Required] <p>Original price of model.&nbsp;If you upload this field, we will take your value, so you should pass the value in local currency, if you don't upload this field, Shopee will automatically calculate the price.</p>
	ModelStatus   *string     `json:"model_status,omitempty"` // [Optional] <p>can be&nbsp;"NORMAL" or "UNAVAILABLE". Normal models can be sold on the buyer's side, and UNAVAILABLE models cannot be sold on the buyer's side. If you do not upload this field, the model status will be considered as "NORMAL".</p>
}

type ItemModelPrice struct {
	ModelId   int64   `json:"model_id"`   // [Required] <p>Id of main model.</p>
	TierIndex []int64 `json:"tier_index"` // [Required] <p>Tier index of main model. Index starts from 0.</p>
	Price     float64 `json:"price"`      // [Required]
}

type ItemPlanAhora struct {
	ItemId              int64 `json:"item_id"`               // [Required] Only applicable for local AR sellers.
	ParticipatePlanAhor bool  `json:"participate_plan_ahor"` // [Required] Only applicable for local AR sellers.
}

type ItemPromotion struct {
	PromotionType string `json:"promotion_type"` // [Required] <p>Indicates the type of item- or package-level promotion applied to a product. Each item can be associated with at most one item promotion and one package promotion at a time.<br /><br />Item Promotions:</p><p>low_price_promotion</p><p>deep_discount</p><p>platform_sale</p><p>seller_discount</p><p>flash_sale</p><p>wholesale</p><p>welcome_package_free_gift</p><p>brand_flash_sale</p><p>in_shop_flash_sale</p><p>synced_promo</p><p>platform_streaming_price</p><p>seller_streaming_price</p><p>exclusive_streamer_price</p><p>price_bidding_with_rebate</p><p>price_bidding_without_rebate</p><p>seller_advisor_price</p><p>selling_price</p><p>settlement_price</p><p>campaign_settlement_price</p><p>local_sip_settlement_price</p><p>platform_exclusive_price</p><p>seller_exclusive_price</p><p>seller_member_exclusive_sku</p><p>item_price</p><p>order_sync_price</p><p><br />Package Promotions:<br />bundle_deal<br />add_on_deal_main<br />add_on_deal_sub<br /><br /><br /><br /><br /></p>
	PromotionId   int64  `json:"promotion_id"`   // [Required] <p>Represents the unique identifier of a specific promotion applied to an item. Each promotion_id corresponds to a distinct promotion rule or campaign, defined under a particular promotion_type. The value is expressed in a numeric string format.</p>
}

type ItemSetting struct {
	ItemName          string             `json:"item_name"`                  // [Required] <p>The name of this kit item.<br /></p>
	Images            *PromotionImages   `json:"images"`                     // [Required] <p>Item images with 1:1 ratio.<br /></p>
	LongImages        *PromotionImages   `json:"long_images,omitempty"`      // [Optional] <p>Item images with 3:4 ratio.<br /></p>
	VideoUploadId     []string           `json:"video_upload_id,omitempty"`  // [Optional] <p>Video upload ID returned from video uploading API. Only accept one video_upload_id.<br /></p>
	Description       *string            `json:"description,omitempty"`      // [Optional] <p>If description_type is normal, description information should be set by this field.<br /></p>
	DescriptionInfo   *DescriptionInfo   `json:"description_info,omitempty"` // [Optional] <p>Rich text description field. Only whitelist sellers can use it. If you use the field, please upload the description_type=extended otherwise api will return error. If you don't use this field, you don't need to upload the description_type or upload description_type=normal<br /></p>
	DescriptionType   DescriptionType    `json:"description_type"`           // [Required] <p>See Data Definition- description_type (normal , extended). If you want to use extended_description, this field must be inputed.<br /></p>
	LogisticInfo      []LogisticInfo     `json:"logistic_info"`              // [Required] <p>Logistic channel setting.<br /></p>
	Unlisted          *bool              `json:"unlisted,omitempty"`         // [Optional] <p>Unlist or not.<br /></p>
	ItemSku           *string            `json:"item_sku,omitempty"`         // [Optional] <p>SKU tag of item<br /></p>
	Weight            float64            `json:"weight"`                     // [Required] <p>The weight of this kit item, the unit is KG.<br /></p>
	Dimension         *Dimension         `json:"dimension,omitempty"`        // [Optional] <p>The dimension of this kit item.<br /></p>
	PreOrder          *PreOrder          `json:"pre_order,omitempty"`        // [Optional] <p>Pre order setting.<br /></p>
	ModelList         []ItemSettingModel `json:"model_list"`                 // [Required] <p>Model info list, model number at most 9.<br /></p>
	TierVariationList []TierVariation    `json:"tier_variation_list"`        // [Required] <p>Tier variation info list.&nbsp;</p><p>Only support one tier variation, and each kit item can have from 1 to 9 kit variations.</p>
}

type ItemSettingModel struct {
	TierIndex     interface{} `json:"tier_index"`          // [Required] <p>Tier index of this kit model.</p>
	ModelSku      *string     `json:"model_sku,omitempty"` // [Optional] <p>Seller SKU of this kit model, model_sku length information needs to be no more than 100 characters.<br /></p>
	OriginalPrice float64     `json:"original_price"`      // [Required] <p>Original price of this kit model.</p>
	ComponentList []Component `json:"component_list"`      // [Required] <p>Please get the amount of item/model that one kit model support from get_kit_item_limit.</p>
}

type ItemSku struct {
	MtskuId            string    `json:"mtsku_id"`             // [Required] <p>Unique ID for a warehouse SKU<br />"warehouse_item_id"_"warehouse_model_id"</p>
	ModelId            string    `json:"model_id"`             // [Required] <p>Warehouse model SKU ID</p><p>For CB global items, this is equal to the global model_id.</p><p><br /></p><p>For local items, it differs from model_id; use shop_model_id to match the model_id</p>
	FulfillMappingMode int64     `json:"fulfill_mapping_mode"` // [Required] <p>0-Null；1-Bundle SKU；2-Parent SKU<br /></p>
	Variation          string    `json:"variation"`            // [Required]
	ShopSkuList        []ShopSku `json:"shop_sku_list"`        // [Required]
	WhsList            []SkuWhs  `json:"whs_list"`             // [Required]
}

type ItemSkuWhs struct {
	WhsId              string `json:"whs_id"`                 // [Required] <p>Whs id<br /></p>
	QtyOfStockAgeOne   int64  `json:"qty_of_stock_age_one"`   // [Required] <p>0-30Days<br /></p>
	QtyOfStockAgeTwo   int64  `json:"qty_of_stock_age_two"`   // [Required] <p>31-60Days<br /></p>
	QtyOfStockAgeThree int64  `json:"qty_of_stock_age_three"` // [Required] <p>61-90Days<br /></p>
	QtyOfStockAgeFour  int64  `json:"qty_of_stock_age_four"`  // [Required] <p>91-120Days<br /></p>
	QtyOfStockAgeFive  int64  `json:"qty_of_stock_age_five"`  // [Required] <p>121-180Days<br /></p>
	QtyOfStockAgeSix   int64  `json:"qty_of_stock_age_six"`   // [Required] <p>&gt;180Days<br /></p>
	ExcessStock        int64  `json:"excess_stock"`           // [Required] <p>expired stock<br /></p>
	AgingStorageTag    int64  `json:"aging_storage_tag"`      // [Required]
}

type ItemStatusDetails struct {
	ViolationType   string `json:"violation_type"`    // [Required] <p>Violation types defined by Shopee.&nbsp;Applicable values:&nbsp;</p><p>Prohibited Listing</p><p>Counterfeit and IP Infringement</p><p>Spam</p><p>Inappropriate Image</p><p>Insufficient Information</p><p>Mall Listing Improvement</p><p>Other Listing Improvement<br /></p>
	ViolationReason string `json:"violation_reason"`  // [Required] <p>The reason for violation.<br /></p>
	Suggestion      string `json:"suggestion"`        // [Required] <p>Shopee provides you with suggestions for modifying items.<br /></p>
	FixDeadlineTime int64  `json:"fix_deadline_time"` // [Required] <p>Action required deadline. Empty if no deadline.<br /></p>
	UpdateTime      int64  `json:"update_time"`       // [Required] <p>Latest update time.<br /></p>
}

type Items struct {
	ItemId  int64 `json:"item_id"`  // [Required] <p>Shopee's unique identifier for an item.<br /></p>
	ModelId int64 `json:"model_id"` // [Required] <p>Shopee's unique identifier for a model of an item.<br /></p>
	GroupId int64 `json:"group_id"` // [Required] <p>The identify of product promotion. For items in one same add on deal promotion, the group_id should share the same id. For items not in add on deal promotion, the group_id should be 0. And the data is from group_id of shopee.orders.GetOrderDetails.<br /></p>
}

type ItemsModels struct {
	ModelId         int64    `json:"model_id"`                    // [Required] <p>If the item has variation, this param is necessary.</p>
	Status          int64    `json:"status"`                      // [Required] <p>you can use this field to set the status of model</p><p><br /></p><p>0: disable<br /></p><p>1: enable</p>
	InputPromoPrice *float64 `json:"input_promo_price,omitempty"` // [Optional] <p>promotion price without tax</p><p><br /></p><p>if the model is enabled(status&nbsp; = 1) now, you can't set this field, you can only disable the model</p><p>if the model is disabled(status&nbsp; = 0) now and you want to set this field, you should also set status to 1</p>
	Stock           *int64   `json:"stock,omitempty"`             // [Optional] <p>min=1, Campaign Stock, Campaign stock can only be reserved from either Shopee stock or Seller stock<br /></p><p><br /></p><p>if the model is enabled(status&nbsp; = 1) now, you can't set this field,&nbsp;you can only disable the model</p><p>if the model is disabled(status&nbsp; = 0) now and you want to set this field, you should also set status to 1</p>
}

type KitItems struct {
	MtItemId          float64 `json:"mt_item_id"`          // [Required] <p>The merchant item identifier of the product within the kit&nbsp;(Only for BR Local Sellers)</p>
	MtModelId         float64 `json:"mt_model_id"`         // [Required] <p>The merchant product model of the item within the kit&nbsp;(Only for BR Local Sellers)</p>
	TotalQty          float64 `json:"total_qty"`           // [Required] <p>The quantity of this specific component within the kit.&nbsp;(Only for BR Local Sellers)</p>
	ParentItemPrice   float64 `json:"parent_item_price"`   // [Required] <p>The price of the item when it is listed as a standalone item.</p>
	ItemPriceProrated float64 `json:"item_price_prorated"` // [Required] <p>The price of the item when it is listed within the kit (i.e. proportionally distributed)&nbsp;(Only for BR Local Sellers)</p>
}

type LateOrder struct {
	OrderSn          string `json:"order_sn"`          // [Required] <p>Order SN.</p>
	ShippingDeadline int64  `json:"shipping_deadline"` // [Required] <p>Shipping Deadline of this order.</p>
	LateByDays       int64  `json:"late_by_days"`      // [Required] <p>Late-by Days of this order.</p>
}

type Licenses struct {
	FileName *string `json:"file_name,omitempty"` // [Optional] <p>Brand registration certificate image name, len &lt;&nbsp;254<br /></p>
	FileHash *string `json:"file_hash,omitempty"` // [Optional] <p>Image id of brand registration certificate image , max input num of file = 1 , each file's length&lt;=498<br /></p>
}

type LinkSspRequest struct {
	ItemId            int64                     `json:"item_id"`                       // [Required] <p>ID of this item.<br /></p>
	SspId             int64                     `json:"ssp_id"`                        // [Required] <p>Shopee's unique identifier for Shopee&nbsp;Standard Product.<br /></p>
	OriginalPrice     *float64                  `json:"original_price,omitempty"`      // [Optional] <p>The price of this item.<br /></p>
	ItemStatus        *ItemStatus               `json:"item_status,omitempty"`         // [Optional] <p>Item status, could be UNLIST or NORMAL.<br /></p>
	Dimension         *Dimension                `json:"dimension,omitempty"`           // [Optional] <p>The dimension of this item.<br /></p>
	LogisticInfo      []LogisticInfo            `json:"logistic_info,omitempty"`       // [Optional] <p>Logistic channel setting of this item.<br /></p>
	AttributeList     []Attribute               `json:"attribute_list,omitempty"`      // [Optional]
	PreOrder          *PreOrder                 `json:"pre_order,omitempty"`           // [Optional] <p>Pre order setting of this item.</p>
	ItemSku           *string                   `json:"item_sku,omitempty"`            // [Optional] <p>SKU tag of this item.<br /></p>
	Condition         *string                   `json:"condition,omitempty"`           // [Optional] <p>Condition of item, could be USED or NEW.<br /></p>
	Wholesale         []Wholesale               `json:"wholesale,omitempty"`           // [Optional] <p>Wholesale setting of this item.<br /></p>
	VideoUploadId     []string                  `json:"video_upload_id,omitempty"`     // [Optional] <p>Video upload ID returned from video uploading API. Only accept one video_upload_id.<br /></p>
	ItemDangerous     *int64                    `json:"item_dangerous,omitempty"`      // [Optional] <p>This field is only applicable for&nbsp;local sellers&nbsp;in Indonesia and Malaysia. Use this field to identify whether a product is a dangerous product. 0 for non-dangerous product and 1 for dangerous product. For more information, please visit the market's respective Seller Education Hub.<br /></p>
	TaxInfo           *AddSspItemRequestTaxInfo `json:"tax_info,omitempty"`            // [Optional] <p>Tax information for this item.<br /></p>
	SellerStock       []SellerStock             `json:"seller_stock,omitempty"`        // [Optional] <p>seller stock（Please notice that stock(including Seller Stock and Shopee Stock) should be larger than or equal to real-time reserved stock）<br /></p>
	SizeChartInfo     *SizeChartInfo            `json:"size_chart_info,omitempty"`     // [Optional]
	AuthorisedBrandId *int64                    `json:"authorised_brand_id,omitempty"` // [Optional] <p>ID of authorised reseller brand.<br /></p>
	ModelList         []AddSspItemRequestModel  `json:"model_list,omitempty"`          // [Optional] <p>Model info list.<br /></p>
}

type LinkSspResponse struct {
	BaseResponse `json:",inline"`    // Common response fields
	Response     LinkSspResponseData `json:"response"` // Actual response data
}

type LinkSspResponseData struct {
	ItemId int64 `json:"item_id"` // [Required] <p>Shopee's unique identifier for an item.<br /></p>
}

type LinkedDirectShop struct {
	DirectShopId     int64  `json:"direct_shop_id"`     // [Required] <p>Shop ID of the Cross Border Direct Shop.</p>
	DirectShopRegion string `json:"direct_shop_region"` // [Required] <p>Shop Region of the Cross Border Direct Shop.</p>
}

type List struct {
	AttributeTree []AttributeTree `json:"attribute_tree"` // [Required] <p>One category's attribute trees</p>
	CategoryId    int64           `json:"category_id"`    // [Required] <p>Category ID</p>
	Warning       string          `json:"warning"`        // [Required] <p>Warning msg</p>
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

type ListGmsUserDeletedItemRequest struct {
	Offset *int64 `json:"offset,omitempty"` // [Optional] <p>Specifies the starting point, or the number of records to skip. Default is 0.</p>
	Limit  *int64 `json:"limit,omitempty"`  // [Optional] <p>Specifies the maximum number of records to show. Default is 50. Maximum is 100.<br /></p>
}

type ListGmsUserDeletedItemResponse struct {
	BaseResponse `json:",inline"`                   // Common response fields
	Response     ListGmsUserDeletedItemResponseData `json:"response"` // Actual response data
}

type ListGmsUserDeletedItemResponseData struct {
	CampaignId  int64   `json:"campaign_id"`   // [Required] <p>GMS Campaign ID</p>
	ItemIdList  []int64 `json:"item_id_list"`  // [Required] <p>List of Item IDs</p>
	Total       int64   `json:"total"`         // [Required] <p>Total number of Item IDs</p>
	HasNextPage bool    `json:"has_next_page"` // [Required] <p>Indicate that there are more item IDs.</p>
}

type ListItem struct {
	ItemId    int64          `json:"item_id"`    // [Required] <p>Shopee's unique identifier for an item.<br /></p>
	Name      string         `json:"name"`       // [Required] <p>Name of the item in local language.<br /></p>
	ImageUrl  string         `json:"image_url"`  // [Required] <p>The image url of the item.</p>
	PriceInfo *ListPriceInfo `json:"price_info"` // [Required]
}

type ListMetric struct {
	ItemClicks int64 `json:"item_clicks"` // [Required] <p>Number of product clicks.</p>
	Atc        int64 `json:"atc"`         // [Required] <p>Number of "Add To Cart" button clicked for all products in the orange bag during livestream.<br /></p>
	SoldItems  int64 `json:"sold_items"`  // [Required] <p>Number of product sold.</p>
}

type ListPriceInfo struct {
	Currency      string  `json:"currency"`       // [Required] <p>The three-digit code representing the currency unit used for the item.<br /></p>
	CurrentPrice  float64 `json:"current_price"`  // [Required] <p>The current price of the item in the listing currency. If product under an ongoing promotion, current_price will be the promotion price.<br /></p>
	OriginalPrice float64 `json:"original_price"` // [Required] <p>The original price of the item in the listing currency.<br /></p>
}

type ListShopSku struct {
	ShopItemId    int64  `json:"shop_item_id"`    // [Required] <p>ID of item</p>
	ShopModelId   int64  `json:"shop_model_id"`   // [Required] <p>ID of model</p>
	ShopItemName  string `json:"shop_item_name"`  // [Required] <p>Name of item</p>
	ShopModelName string `json:"shop_model_name"` // [Required] <p>Name of model</p>
	FailReason    string `json:"fail_reason"`     // [Required] <p>Invoice issuance failed reason.</p>
}

type Listing struct {
	ItemId int64 `json:"item_id"` // [Required] <p>Item ID.</p>
	Reason int64 `json:"reason"`  // [Required] <p>Reason of this item. Applicable values:&nbsp;</p><p>1: Prohibited</p><p>2: Counterfeit</p><p>3: Spam</p><p>4: Inappropriate Image</p><p>5: Insufficient Info</p><p>6: Mall Listing Improvement</p><p>7: Other Listing Improvement<br /></p>
}

type LivestreamAddItemListRequest struct {
	SessionId int64                             `json:"session_id"` // [Required] <p>The identifier of livestream session.</p>
	ItemList  []DeleteBundleDealItemRequestItem `json:"item_list"`  // [Required] <p>The list of item to add.</p>
}

type LivestreamAddItemListResponse struct {
	BaseResponse `json:",inline"` // Common response fields
	Response     interface{}      `json:"response"` // Actual response data
}

type LivestreamDeleteItemListRequest struct {
	SessionId int64                             `json:"session_id"` // [Required] <p>The identifier of livestream session.</p>
	ItemList  []DeleteBundleDealItemRequestItem `json:"item_list"`  // [Required] <p>The list of item to delete.</p>
}

type LivestreamDeleteItemListResponse struct {
	BaseResponse `json:",inline"` // Common response fields
	Response     interface{}      `json:"response"` // Actual response data
}

type LivestreamGetItemListRequest struct {
	SessionId int64 `json:"session_id" url:"session_id"` // [Required] <p>The identifier of livestream session.</p>
	Offset    int64 `json:"offset" url:"offset"`         // [Required] <p>Specifies the starting entry of data to return in the current call. Default is 0, if data is more than one page, the offset can be some entry to start next call.</p>
	PageSize  int64 `json:"page_size" url:"page_size"`   // [Required] <p>Each result set is returned as a page of entries. Use the "page_size" filters to control the maximum number of entries to retrieve per page (i.e., per call). This integer value is used to specify the maximum number of entries to return in a single "page" of data. The limit of page_size if between 1 and 100.</p>
}

type LivestreamGetItemListResponse struct {
	BaseResponse `json:",inline"`                  // Common response fields
	Response     LivestreamGetItemListResponseData `json:"response"` // Actual response data
}

type LivestreamGetItemListResponseData struct {
	More       bool                          `json:"more"`        // [Required] <p>This is to indicate whether the list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of data.</p>
	NextOffset int64                         `json:"next_offset"` // [Required] <p>If more is true, this value need set to next request offset.</p>
	List       []GetItemListResponseDataList `json:"list"`        // [Required]
}

type LivestreamUploadImageRequest struct {
	Image interface{} `json:"image"` // [Required] <p>The image file to upload.</p>
}

type LivestreamUploadImageResponse struct {
	BaseResponse `json:",inline"`                  // Common response fields
	Response     LivestreamUploadImageResponseData `json:"response"` // Actual response data
}

type LivestreamUploadImageResponseData struct {
	ImageUrl string `json:"image_url"` // [Required] <p>The image URL</p>
}

type LogisticInfo struct {
	SizeId      int64   `json:"size_id"`      // [Required] Size ID
	ShippingFee float64 `json:"shipping_fee"` // [Required] Shipping fee
	Enabled     bool    `json:"enabled"`      // [Required] Whether this channel is enabled for this item
	LogisticId  int64   `json:"logistic_id"`  // [Required] Logistic channel ID
	IsFree      bool    `json:"is_free"`      // [Required] Whether cover shipping fee for buyer
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

type LowerBound struct {
	Value      float64 `json:"value"`      // [Required] <p>The ROI target value.</p>
	Percentile int64   `json:"percentile"` // [Required] <p>Competitiveness over similar ads.</p>
}

type LsrOrder struct {
	OrderSn            string   `json:"order_sn"`             // [Required] <p>Order SN.</p>
	ShippingDeadline   int64    `json:"shipping_deadline"`    // [Required] <p>Ship by date.</p>
	ActualShippingTime int64    `json:"actual_shipping_time"` // [Required] <p>Seller arrange shipment time.</p>
	LateByDays         int64    `json:"late_by_days"`         // [Required] <p>Late-by Days.</p>
	ActualPickUpTime   int64    `json:"actual_pick_up_time"`  // [Required] <p>Courier actual pick up time.</p>
	ShippingChannel    string   `json:"shipping_channel"`     // [Required] <p>Logistics Company.</p>
	FirstMileType      string   `json:"first_mile_type"`      // [Required] <p>First mile shipping type. Applicable values:</p><p>Pickup</p><p>Drop off</p>
	DiagnosisScenario  []string `json:"diagnosis_scenario"`   // [Required] <p>Diagnosis of the issue.</p>
}

type ManualBiddingInfo struct {
	EnhancedCpc           bool                                `json:"enhanced_cpc"`            // [Required] <p>Enhanced CPC functionality&nbsp;</p>
	SelectedKeywords      []ManualBiddingInfoSelectedKeywords `json:"selected_keywords"`       // [Required] <p>selected keywords</p>
	DiscoveryAdsLocations []RequestDiscoveryAdsLocations      `json:"discovery_ads_locations"` // [Required] <p>the location settings&nbsp;</p>
}

type ManualBiddingInfoSelectedKeywords struct {
	Keyword          string  `json:"keyword"`             // [Required] <p>bid keywords for each campaign with search placement</p>
	Status           string  `json:"status"`              // [Required] <p>deleted, normal, reserved, blacklist&nbsp;<br /></p>
	MatchType        string  `json:"match_type"`          // [Required] <p>exact, broad<br /></p>
	BidPricePerClick float64 `json:"bid_price_per_click"` // [Required] <p>the bid price&nbsp;</p>
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

type MaxPurchaseLimit struct {
	PurchaseLimit int64 `json:"purchase_limit"` // [Required] <p>maximum purchase limit for each order</p>
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

type Media struct {
	ImageUrlList []string `json:"image_url_list"` // [Required] <p>List of image url uploaded by the buyer in the comment.</p>
	VideoUrlList []string `json:"video_url_list"` // [Required] <p>List of video url uploaded by the buyer in the comment.</p>
}

type MediaCancelVideoUploadRequest struct {
	VideoUploadId string `json:"video_upload_id"` // [Required] <p>The unique ID of the upload task, returned by v2.media.init_video_upload.</p>
}

type MediaCancelVideoUploadResponse struct {
	BaseResponse `json:",inline"` // Common response fields
}

type MediaCompleteVideoUploadRequest struct {
	VideoUploadId string `json:"video_upload_id"` // [Required] <p>The unique ID of the upload task, returned by v2.media.init_video_upload.</p>
}

type MediaCompleteVideoUploadResponse struct {
	BaseResponse `json:",inline"` // Common response fields
}

type MediaGetVideoUploadResultRequest struct {
	VideoUploadId string `json:"video_upload_id" url:"video_upload_id"` // [Required] <p>The unique ID of the upload task, returned by v2.media.init_video_upload.</p>
}

type MediaGetVideoUploadResultResponse struct {
	BaseResponse `json:",inline"`                      // Common response fields
	Response     MediaGetVideoUploadResultResponseData `json:"response"` // Actual response data
}

type MediaGetVideoUploadResultResponseData struct {
	Status     string                                     `json:"status"`      // [Required] <p>Current status of the upload task. Possible values:<br /></p><p>- INITIATED: Upload task has been created (via init_video_upload) but no parts have been uploaded yet.</p><p>- UPLOADING: Video file parts are being uploaded. The upload has started but is not yet completed.</p><p>- UPLOADED: All video parts have been uploaded successfully, waiting for complete_video_upload to trigger processing.</p><p>- PROCESSING: Video is being transcoded/validated by the system (duration, format, resolution checks).</p><p>- SUCCEEDED: Video upload and transcoding completed successfully. Video URL and cover URL are available for use.</p><p>- FAILED: Upload or processing failed (e.g., invalid format, duration not within allowed range, transcoding error).</p><p>- CANCELLED: Upload task was explicitly canceled by the client (cancel_video_upload), and the video is discarded.</p>
	Reason     string                                     `json:"reason"`      // [Required] <p>Detailed fail or cancel reason, will be returned if status is FAILED or CANCELLED.</p>
	UpdateTime int64                                      `json:"update_time"` // [Required] <p>The time of video status updates.</p>
	VideoInfo  *GetVideoUploadResultResponseDataVideoInfo `json:"video_info"`  // [Required] <p>Transcoded video info, will be returned if status is SUCCEEDED.</p>
}

type MediaInitVideoUploadRequest struct {
	Business int64  `json:"business"`  // [Required] <p>Defines the business type of the uploaded image. Supported values:&nbsp;</p><p>3 = Video</p>
	Scene    int64  `json:"scene"`     // [Required] <p>Defines the purpose of the uploaded image under the specified business type.&nbsp;Supported values:&nbsp;</p><p>- If business = 3 (Video): 1 = Shopee Video</p>
	FileName string `json:"file_name"` // [Required] <p>Original video file name.</p>
	FileSize int64  `json:"file_size"` // [Required] <p>Total video file size in bytes.&nbsp;Rules and restrictions by business and scene:&nbsp;</p><p>- If business = 3 (Video) and scene = 1 (Shopee Video):&nbsp;Maximum <b>1GB</b>.</p>
	Duration int64  `json:"duration"`  // [Required] <p>Video duration in seconds.&nbsp;Rules and restrictions by business and scene:</p><p>- If business = 3 (Video) and scene = 1 (Shopee Video): <b>1s~180s</b>.</p>
}

type MediaInitVideoUploadResponse struct {
	BaseResponse `json:",inline"`                 // Common response fields
	Response     MediaInitVideoUploadResponseData `json:"response"` // Actual response data
}

type MediaInitVideoUploadResponseData struct {
	VideoUploadId string `json:"video_upload_id"` // [Required] <p>Unique upload session ID.</p>
	PartSize      int64  `json:"part_size"`       // [Required] <p>The size of each part. When uploading video chunks, the video must be split according to this part size for each upload request.</p>
}

type MediaUploadImageRequest struct {
	Business int64       `json:"business"` // [Required] <p>Defines the business type of the uploaded image. Supported values:&nbsp;</p><p>2 = Returns</p>
	Scene    int64       `json:"scene"`    // [Required] <p>Defines the purpose of the uploaded image under the specified business type.&nbsp;Supported values:</p><p>- If&nbsp;business = 2 (Returns): 1 = Return Seller Self Arrange Pickup Proof Image</p>
	Images   interface{} `json:"images"`   // [Required] <p>The image files to be uploaded. Rules and restrictions by business and scene:</p><p>- If business = 2 (Returns) and scene = 1 (Return Seller Self Arrange Pickup Proof Image): Up to&nbsp;<b>3</b>&nbsp;images can be uploaded. Each image must not exceed&nbsp;<b>10MB</b>. Supported formats:&nbsp;<b>JPG, JPEG, PNG</b>.</p>
}

type MediaUploadImageResponse struct {
	BaseResponse `json:",inline"`             // Common response fields
	Response     MediaUploadImageResponseData `json:"response"` // Actual response data
}

type MediaUploadImageResponseData struct {
	ImageList []FieldImageInfo `json:"image_list"` // [Required] <p>List of uploaded images.</p>
}

type MediaUploadVideoPartRequest struct {
	VideoUploadId string      `json:"video_upload_id"` // [Required] <p>The unique ID of the upload task, returned by v2.media.init_video_upload.</p>
	PartSeq       int64       `json:"part_seq"`        // [Required] <p>Sequence number of this part, starting from 0.</p>
	PartContent   interface{} `json:"part_content"`    // [Required] <p>The content of this part of file. Part size should be exactly equal to part_size returned in v2.media.init_video_upload, except last part of file.</p>
	PartMd5       string      `json:"part_md5"`        // [Required] <p>MD5 checksum of this part for data integrity validation.</p>
}

type MediaUploadVideoPartResponse struct {
	BaseResponse `json:",inline"` // Common response fields
}

type Metric struct {
	MetricType       int64   `json:"metric_type"`        // [Required] <p>Type of metric:&nbsp;</p><p>Fulfillment Performance = 1</p><p>Listing Performance = 2</p><p>Customer Service Performance = 3</p>
	MetricId         int64   `json:"metric_id"`          // [Required] <p>ID of metric.</p><p>If metric_id &lt; 0, it means that this is not a real metric, but a group of metrics.</p><p><br /></p><p>Non-Responded Chats = -1</p><p>Late Shipment Rate (All Channels) = 1</p><p>Non-Fulfilment Rate (All Channels) = 3</p><p>Preparation Time = 4</p><p>Chat Response Rate = 11</p><p>Pre-order Listing % = 12<br /></p><p>Days of Pre-order Listing Violation = 15</p><p>Response Time = 21</p><p>Shop Rating = 22</p><p>No. of Non-Responded Chats = 23</p><p>Fast Handover Rate = 25</p><p>On-time Pickup Failure Rate = 27</p><p>On-time Pickup Failure Rate Violation Value = 28</p><p>Average Response Time = 29</p><p>Cancellation Rate (All Channels) = 42</p><p>Return-refund Rate (All Channels) = 43</p><p>Severe Listing Violations = 52</p><p>Other Listing Violations = 53</p><p>Prohibited Listings = 54</p><p>Counterfeit/IP infringement = 55</p><p>Spam Listings = 56</p><p>Late Shipment Rate (NDD) = 85</p><p>Non-fulfilment Rate (NDD) = 88</p><p>Cancellation Rate (NDD) = 91</p><p>Return-refund Rate (NDD) = 92<br /></p><p>Customer Satisfaction = 95</p><p>% SDD Listings = 96</p><p>% NDD Listings = 97</p><p>Fast Handover Rate - SLS = 2001</p><p>Fast Handover Rate - FBS = 2002<br /></p><p>Fast Handover Rate - 3PF = 2003</p><p>Poor Quality Products = 2011</p><p>% HD Listings = 2030</p><p>% HD Free Shipping Enabled = 2031</p><p>Saturday Shipment = 2032<br />Preparation Time PS = 2033</p>
	ParentMetricId   int64   `json:"parent_metric_id"`   // [Required] <p>ID of parent metric.<br /></p>
	MetricName       string  `json:"metric_name"`        // [Required] <p>Default name of metric.</p>
	CurrentPeriod    float64 `json:"current_period"`     // [Required] <p>The performance of the metric at current period.</p>
	LastPeriod       float64 `json:"last_period"`        // [Required] <p>The performance of the metric at last period.<br /></p>
	Unit             int64   `json:"unit"`               // [Required] <p>Unit of metric:&nbsp;</p><p>Number = 1</p><p>Percentage = 2</p><p>Second = 3</p><p>Day = 4</p><p>Hour = 5</p>
	Target           *Target `json:"target"`             // [Required]
	ExemptionEndDate string  `json:"exemption_end_date"` // [Required] <p>(Only for whitelist TW sellers) The exemption_end_date value will not be empty if ALL conditions are met:&nbsp;</p><p>- The shop is in the "POL Shop Whitelist"</p><p>- Within the "Exemption Period"</p><p>- The metric_id is 12 (Pre-order Listing %) or 15 (Days of Pre-order Listing Violation)</p>
}

type Metrics struct {
	Date              string  `json:"date"`                // [Required] <p>the given date for the performance&nbsp;</p>
	Impression        int64   `json:"impression"`          // [Required] <p>The number of times shoppers see your ad.</p>
	Clicks            int64   `json:"clicks"`              // [Required] <p>The number of times shoppers click on your ad. (Note: Shopee filters out repeated clicks from the same shopper that occur within a short time frame.)<br /></p>
	Ctr               float64 `json:"ctr"`                 // [Required] <p>The click-through rate (CTR) measures how often shoppers end up clicking on your ad after seeing it. It is the number of clicks on your ad divided by the number of times your ad is seen. CTR = clicks ÷ impressions × 100%.</p>
	Expense           float64 `json:"expense"`             // [Required] <p>The amount spent on your ad.<br /></p>
	BroadGmv          float64 `json:"broad_gmv"`           // [Required] <p>The amount of sales revenue generated from shoppers purchasing products within 7 days of them clicking on your ad.<br /></p>
	BroadOrder        int64   `json:"broad_order"`         // [Required] <p>The number of times shoppers purchased any product from your shop within 7 days of them clicking on your ad.<br /></p>
	BroadOrderAmount  int64   `json:"broad_order_amount"`  // [Required] <p>The total quantity of products purchased by shoppers within 7 days of them clicking on your ad.<br /></p>
	BroadRoi          float64 `json:"broad_roi"`           // [Required] <p>Return on ad spend (ROAS) measures how much revenue is generated by your ad relative to the cost of the ad. It is the amount of sales revenue attributed to your ad divided by the amount spent on the ad. ROAS = GMV ÷ expense. (Note: We recommend monitoring ROAS trends on a weekly basis.)<br /></p>
	BroadCir          float64 `json:"broad_cir"`           // [Required] <p>The advertising cost of sales (ACOS) measures how much your ad costs relative to the revenue the ad generates. It is the amount spent on your ad divided by the amount of sales revenue attributed to the ad. ACOS = expense ÷ GMV × 100%.<br /></p>
	Cr                float64 `json:"cr"`                  // [Required] <p>The conversion rate measures how often shoppers end up purchasing something from your shop after clicking on your ad. It is the number of conversions attributed to your ad divided by the number of clicks on the ad. Conversion rate = conversions ÷ clicks × 100%.<br /></p>
	Cpc               float64 `json:"cpc"`                 // [Required] <p>The cost per conversion is how much each conversion costs, on average. It is the amount spent on your ad divided by the number of conversions attributed to the ad. Cost per conversion = expense ÷ conversions.<br /></p>
	DirectOrder       int64   `json:"direct_order"`        // [Required] <p>The number of times shoppers purchased the advertised product within 7 days of them clicking on the ad.<br /></p>
	DirectOrderAmount int64   `json:"direct_order_amount"` // [Required] <p>The total quantity of the advertised product purchased by shoppers within 7 days of them clicking on the ad.<br /></p>
	DirectGmv         float64 `json:"direct_gmv"`          // [Required] <p>The amount of sales revenue generated from shoppers purchasing the advertised product within 7 days of them clicking on the ad.<br /></p>
	DirectRoi         float64 `json:"direct_roi"`          // [Required] <p>The direct return on ad spend, or direct ROAS, measures how much revenue is generated from sales of the advertised product, relative to the cost of the ad. It is the amount of sales revenue for the advertised product attributed to the ad, divided by the amount spent on the ad. Direct ROAS = direct GMV ÷ expense.<br /></p>
	DirectCir         float64 `json:"direct_cir"`          // [Required] <p>The direct advertising cost of sales, or direct ACOS, measures how much your ad costs relative to the revenue generated from sales of the advertised product. It is the amount spent on the ad divided by the amount of sales revenue for the advertised product that is attributed to the ad. Direct ACOS = expense ÷ direct GMV × 100%.<br /></p>
	DirectCr          float64 `json:"direct_cr"`           // [Required] <p>The direct conversion rate measures how often shoppers end up purchasing the advertised product after clicking on the ad. Direct conversion rate is the number of direct conversions divided by the number of clicks. Direct conversion rate = direct conversions ÷ clicks × 100%.<br /></p>
	Cpdc              float64 `json:"cpdc"`                // [Required] <p>The cost per direct conversion is how much each direct conversion costs, on average. It is the amount spent on your ad divided by the number of direct conversions attributed to the ad. Cost per direct conversion = expense ÷ direct conversions.<br /></p>
}

type Model struct {
	TierIndex     interface{}   `json:"tier_index"`          // [Required] Tier index of model
	OriginalPrice float64       `json:"original_price"`      // [Required] Normal stock for price
	ModelSku      *string       `json:"model_sku,omitempty"` // [Optional] Seller sku, model_sku length information needs to be no more than 100 characters.
	SellerStock   []SellerStock `json:"seller_stock"`        // [Required] <p>new stock info for model（Please notice that stock(including Seller Stock and Shopee Stock) should be larger than or equal to real-time reserved stock）<br /></p>
	GtinCode      *string       `json:"gtin_code,omitempty"` // [Optional] <p>- GTIN is an identifier for trade items, developed by the international organization GS1.<br />- They have 8 to 14 digits. The most common are UPC, EAN, JAN and ISBN.<br />- GTIN will help boost positioning in online marketing channels like Google and Facebook.<br />- That incorporation with GTIN will also aid in Search and Recommendation in Shopee itself allowing buyers to have higher likelihood of finding one's listing.<br /></p><p><br /></p><p>Note: If you want to set “Item without GTIN”, please pass the gtin_code as "00".<br /><br />The validation rule is based on the value return in gtin_validation_rule" field in v2.product.get_item_limit API</p><p>- <b>Mandatory</b>:&nbsp;This field is required and must contain a correctly formatted GTiN number.</p><p>- <b>Flexible</b>: This field is required and must contain either a correctly formatted GTlN number or "00" to declare that the item/model has no valid GTlN.<br />- <b>Optional</b>: This field is optional and can contain a correctly formatted GTiN number, "00" or be omitted entirely.</p>
	Weight        *float64      `json:"weight,omitempty"`    // [Optional] <p>The weight of this model, the unit is KG.</p><p>If don't set the weight of this model, will use the weight of item by default.</p><p>If set the dimension of this model, them must set the weight of this model.</p>
	Dimension     *Dimension    `json:"dimension,omitempty"` // [Optional] <p>The dimension of this model.</p><p>If don't set the dimension of this model, will use the dimension of item by default.</p>
	PreOrder      *PreOrder     `json:"pre_order,omitempty"` // [Optional] <p>Pre-order information of this model.</p><p><br /></p><p>Notes:&nbsp;</p><p>If don't set the DTS of this model, will use the DTS of the item by default.</p>
}

type ModelComponent struct {
	ComponentItemId           int64  `json:"component_item_id"`             // [Required] <p>ID of the item that composes this kit model.<br /></p>
	ComponentItemName         string `json:"component_item_name"`           // [Required] <p>Name of the item that composes this kit model.<br /></p>
	ComponentModelId          int64  `json:"component_model_id"`            // [Required] <p>ID of the model that composes this kit model.<br /></p>
	ComponentModelName        string `json:"component_model_name"`          // [Required] <p>Name of the model that composes this kit model.<br /></p>
	Quantity                  int64  `json:"quantity"`                      // [Required] <p>The amount of the item/model that composes this kit model.<br /></p>
	MainComponent             bool   `json:"main_component"`                // [Required] <p>Whether this item/model is the main component for this kit.</p>
	ComponentItemOrModelImage string `json:"component_item_or_model_image"` // [Required]
	ComponentItemOrModelSku   string `json:"component_item_or_model_sku"`   // [Required]
}

type ModelMapping struct {
	PmodelId int64 `json:"pmodel_id"` // [Required] <p>ID of model for the P Item.</p>
	AmodelId int64 `json:"amodel_id"` // [Required] <p>ID of model for the A Item.<br /></p>
}

type ModelPriceInfo struct {
	OriginalPrice float64 `json:"original_price"` // [Required] Original price of global model.
}

type ModelSetting struct {
	TierIndex interface{} `json:"tier_index"` // [Required] <p>Tier index of this model.<br /></p>
	Gtin      string      `json:"gtin"`       // [Required] <p>GTIN of this model.</p>
}

type ModelStockInfoV2 struct {
	SummaryInfo  *SummaryInfo             `json:"summary_info"`  // [Required] <p>stock summary Info<br /></p>
	SellerStock  []StockInfoV2SellerStock `json:"seller_stock"`  // [Required] <p>Seller-managed stock<br /></p>
	ShopeeStock  []ShopeeStock            `json:"shopee_stock"`  // [Required] <p>Shopee warehouse stock<br /></p>
	AdvanceStock *AdvanceStock            `json:"advance_stock"` // [Required] <p>Only for PH/VN/ID/MY local selected shops.</p>
}

type Models struct {
	ModelId         int64   `json:"model_id"`          // [Required] <p>If the item has variation, this param is necessary.</p>
	InputPromoPrice float64 `json:"input_promo_price"` // [Required] <p>promotion price without tax<br /></p>
	Stock           int64   `json:"stock"`             // [Required] <p>min=1, Campaign Stock, Campaign stock can only be reserved from either Shopee stock or Seller stock<br /></p>
}

type Monday struct {
	StartTime string `json:"start_time"` // [Required] <p>Start time for Public Holiday</p>
	EndTime   string `json:"end_time"`   // [Required] <p>End time for Public Holiday</p>
}

type MultiLang struct {
	Language string `json:"language"` // [Required] <p>Language<br /></p>
	Value    string `json:"value"`    // [Required] <p>Translate result<br /></p>
}

type NddListing struct {
	ItemId           int64 `json:"item_id"`            // [Required] <p>Item ID.</p>
	CurrentNddStatus int64 `json:"current_ndd_status"` // [Required] <p>Current NDD Status. Applicable values:&nbsp;</p><p>1: Yes</p><p>0: No</p>
}

type Negotiation struct {
	NegotiationStatus  string  `json:"negotiation_status"`   // [Required] To indicate whether the seller can negotiate with the buyer. See "Data Definition - NegotiationStatus"
	LatestSolution     string  `json:"latest_solution"`      // [Required] To indicate what is the offer solution. See "Data Definition - ReturnSolution"
	LatestOfferAmount  float64 `json:"latest_offer_amount"`  // [Required] To indicate the refund amount in the latest offer solution
	LatestOfferCreator string  `json:"latest_offer_creator"` // [Required] To indicate which party made the latest offer
	CounterLimit       int64   `json:"counter_limit"`        // [Required] To indicate the remaining counter limit
	OfferDueDate       int64   `json:"offer_due_date"`       // [Required] To indicate offer_due_date
}

type NetCommissionFeeInfoList struct {
	RuleId          float64 `json:"rule_id"`           // [Required] <p>The unique identifier of the commission rule applied to calculate the net commission fee.</p>
	FeeAmount       float64 `json:"fee_amount"`        // [Required] <p>The net commission fee amount calculated based on the corresponding commission rule.</p>
	RuleDisplayName string  `json:"rule_display_name"` // [Required] <p>The display name of the commission rule for reference and readability.</p>
}

type NetServiceFeeInfoList struct {
	RuleId          float64 `json:"rule_id"`           // [Required] <p>The unique identifier of the service fee rule applied to calculate the net service fee.</p>
	FeeAmount       float64 `json:"fee_amount"`        // [Required] <p>The net service fee amount calculated based on the corresponding service fee rule.</p>
	RuleDisplayName string  `json:"rule_display_name"` // [Required] <p>The display name of the service fee rule for reference and readability.</p>
	Category        string  `json:"category"`          // [Required] <p>The category of the service fee, indicating the type of service the fee is charged for.</p>
}

type NfrOrder struct {
	OrderSn            string `json:"order_sn"`             // [Required] <p>Order SN.</p>
	NonFulfillmentType int64  `json:"non_fulfillment_type"` // [Required] <p>Non-fulfilment type. Applicable values:&nbsp;<br /></p><p>1: System Cancellation</p><p>2: Seller Cancellation</p><p>3: Return Refunds<br /></p>
	DetailedReason     int64  `json:"detailed_reason"`      // [Required] <p>Reason. Applicable values:&nbsp;</p><p>1001: Return Refund</p><p>1002: Parcel Split Cancellation</p><p>1003: First Mile Pick up fail</p><p>1004: Order inclusion<br /></p><p>10005: Out of Stock<br />10006: Undeliverable area<br />10007: Cannot support COD<br />10008: Logistics request cancelled<br />10009: Logistics pickup failed<br />10010: Logistics not ready<br />10011: Inactive seller<br />10012: Seller did not ship order<br />10013: Order did not reach warehouse<br /></p><p>10014: Seller asked to cancel<br /></p><p>10015: Non-receipt<br />10016: Wrong item<br />10017: Damaged item<br />10018: Incomplete product<br />10019: Fake item<br />10020: Functional Damage<br />10021: Return Refund<br /></p>
}

type NonIntegrated struct {
	TrackingNumber *string `json:"tracking_number,omitempty"` // [Optional] Optional parameter for non-integrated channel order. The tracking number assigned by the shipping carrier for item shipment.
}

type OfferRequest struct {
	ReturnSn                     string   `json:"return_sn"`                                 // [Required] The serial number of return.
	ProposedSolution             string   `json:"proposed_solution"`                         // [Required] The new solution to be offered. See "Data Definition - ReturnSolution"
	ProposedAdjustedRefundAmount *float64 `json:"proposed_adjusted_refund_amount,omitempty"` // [Optional] The new refund amount to be offered
}

type OfferResponse struct {
	BaseResponse `json:",inline"`  // Common response fields
	Response     OfferResponseData `json:"response"` // Actual response data
}

type OfferResponseData struct {
	ReturnSn string `json:"return_sn"` // [Required] <p>The serial number of return.</p>
}

type OfferReturnRefund struct {
	Eligibility            bool    `json:"eligibility"`              // [Required] <p>To indicate whether Refund solution is available for sellers to select.</p>
	RefundAmountAdjustable bool    `json:"refund_amount_adjustable"` // [Required] <p>To indicate whether refund is adjustable for Refund solution.</p>
	MaxRefundAmount        float64 `json:"max_refund_amount"`        // [Required] <p>The max refund amount for Refund solution.&nbsp;Returned when refund_amount_adjustable is true.</p>
	MinRefundAmount        float64 `json:"min_refund_amount"`        // [Required] <p>The min refund amount for Refund solution.&nbsp;Returned when refund_amount_adjustable is true.</p>
}

type OfflineAdjustment struct {
	AdjustmentAmount float64 `json:"adjustment_amount"` // [Required] The amount of offline adjustments.
	Module           string  `json:"module"`            // [Required] The reason for offline adjustment.
	Remark           string  `json:"remark"`            // [Required] The remark for the reason.
	Scenario         string  `json:"scenario"`          // [Required] The scenario of adjustment.
	AdjustmentLevel  string  `json:"adjustment_level"`  // [Required] Dimension of offline adjustment. Available value: shop, order.
	OrderSn          string  `json:"order_sn"`          // [Required] Shopee's unique identifier for an order.
}

type OperatingHours struct {
	Date      string `json:"date"`       // [Required] <p>Date: it should include all date from start_date until end_date</p>
	StartTime string `json:"start_time"` // [Required] <p>Start time for that date</p><p>&lt;path&gt;&lt;/path&gt;<br /></p>
	EndTime   string `json:"end_time"`   // [Required] <p>End time for that date</p>
	Enable    bool   `json:"enable"`     // [Required] <p>True: If it is open on that date.</p><p>False: If it is closed on that date.</p>
}

type OpfrDayDetailData struct {
	Date               string `json:"date"`                 // [Required] <p>Date.</p>
	ScheduledPickupNum int64  `json:"scheduled_pickup_num"` // [Required] <p>Number of scheduled pickups.<br /></p>
	FailedPickupNum    int64  `json:"failed_pickup_num"`    // [Required] <p>Number of failed pickups.<br /></p>
	Opfr               int64  `json:"opfr"`                 // [Required] <p>OPFR.<br /></p>
	Target             string `json:"target"`               // [Required] <p>Target.</p>
}

type OptionImage struct {
	ImageUrl string `json:"image_url"` // [Required] <p>The image url of the product. Default to be variation image, if the model does not have a variation image, will use an item main image instead.<br /></p>
}

type OrderAdjustment struct {
	Amount           float64 `json:"amount"`            // [Required] <p>adjustment transaction amount.<br /></p>
	Date             int64   `json:"date"`              // [Required] <p>adjustment transaction complete date.<br /></p>
	Currency         string  `json:"currency"`          // [Required] <p>order level adjustment transaction's currency type.<br /></p>
	AdjustmentReason string  `json:"adjustment_reason"` // [Required] <p>Reason for adjustment.&nbsp;</p><p>Possible cases could be:</p><p>- Return Refund deduction or compensation</p><p>- logistic issue deduction or compensation</p><p>- marketing fee deduction</p><p>- payment related fee<br /></p>
}

type OrderIncome struct {
	EscrowAmount                                       float64                   `json:"escrow_amount"`                                             // [Required] <p>The total amount that the seller is expected to receive for the order and will change before order is completed.&nbsp;</p><p>For non cb sip affiliate shop (new formula):&nbsp;</p><p>escrow_amount=&nbsp;original_cost_of_goods_sold-original_shopee_discount+seller_return_refund+ shopee_discount- voucher_from_seller- seller_coin_cash_back+ buyer_paid_shipping_fee- actual_shipping_fee+ shopee_shipping_rebate+ shipping_fee_discount_from_3pl- reverse_shipping_fee+ rsf_seller_protection_fee_claim_amount- final_return_to_seller_shipping_fee- seller_transaction_fee- service_fee- commission_fee- campaign_fee- shipping_seller_protection_fee_amount- delivery_seller_protection_fee_premium_amount-final_escrow_product_gst- order_ams_commission_fee- escrow_tax-sales_tax_on_lvg-reverse_shipping_fee_sst-shipping_fee_sst-withholding_tax-overseas_return_service_fee-vat_on_imported_goods - withholding_vat_tax - withholding_pit_tax - seller_order_processing_fee&nbsp;+ buyer_paid_packaging_fee - trade_in_bonus_by_seller - fbs_fee -&nbsp;ads_escrow_top_up_fee_or_technical_support_fee<br /></p><p><br /></p><p>For cb sip affiliate shop:&nbsp;escrow_amount=escrow_amount_pri * exchange_rate</p><p><br /></p><p>note: Return refund amount = if adjustable RR, will equal to drc_adjustable_refund<br /></p>
	BuyerTotalAmount                                   float64                   `json:"buyer_total_amount"`                                        // [Required] <p>The snapshot value of total amount that paid by buyer at checkout moment.<br /></p>
	OrderOriginalPrice                                 float64                   `json:"order_original_price"`                                      // [Required] <p>The original price of the item before ANY promotion/discount in the listing currency. It returns the subtotal of that specific item if quantity exceeds 1.<br /></p>
	OriginalPrice                                      float64                   `json:"original_price"`                                            // [Required] The original price of the item before ANY promotion/discount in the listing currency. It returns the subtotal of that specific item if quantity exceeds 1.
	OrderDiscountedPrice                               float64                   `json:"order_discounted_price"`                                    // [Required] <p>The total discounted price for this order. It returns the subtotal of that specific item if quantity exceeds 1. (Only display for non cb sip affiliate shop. )<br /></p>
	OrderSellingPrice                                  float64                   `json:"order_selling_price"`                                       // [Required] <p>This field value = sum of item unit price.selling price comes from the sum up of each item's unit price. For non-bundle deal case, this value will be same like order_original_price; For bundle deal case, this value will be price of sum of item price before bundle deal promo but after item promo.It returns the subtotal of that specific item if quantity exceeds 1. (Only display for non cb sip affiliate shop. )<br /></p>
	OrderSellerDiscount                                float64                   `json:"order_seller_discount"`                                     // [Required] <p>The total discount seller provided for this order. It returns the subtotal of that specific item if quantity exceeds 1. (Only display for non cb sip affiliate shop. )<br /></p>
	SellerDiscount                                     float64                   `json:"seller_discount"`                                           // [Required] Final sum of each item seller discount of a specific order. (Only display for non cb sip affiliate shop. )
	ShopeeDiscount                                     float64                   `json:"shopee_discount"`                                           // [Required] <p>Final sum of each item Shopee discount of a specific order. This amount will Return remaining rebate value to seller. (Only display for non cb sip affiliate order. )</p>
	VoucherFromSeller                                  float64                   `json:"voucher_from_seller"`                                       // [Required] Final value of voucher provided by Seller for the order. (Only display for non cb sip affiliate shop. )
	VoucherFromShopee                                  float64                   `json:"voucher_from_shopee"`                                       // [Required] Final value of voucher provided by Shopee for the order. (Only display for non cb sip affiliate shop. )
	Coins                                              float64                   `json:"coins"`                                                     // [Required] This value indicates the total amount offset when the buyer consumed Shopee Coins upon checkout. (Only display for non cb sip affiliate shop. )
	BuyerPaidShippingFee                               float64                   `json:"buyer_paid_shipping_fee"`                                   // [Required] The shipping fee paid by buyer. (Only display for non cb sip affiliate shop. )
	BuyerTransactionFee                                float64                   `json:"buyer_transaction_fee"`                                     // [Required] Tansaction fee paid by buyer for the order. (Only display for non cb sip affiliate shop. )
	CrossBorderTax                                     float64                   `json:"cross_border_tax"`                                          // [Required] <p>Amount incurred by Buyer for purchasing items outside of home country or region. Amount may change after Return Refund. (Only display for non cb sip affiliate shop. )</p>
	PaymentPromotion                                   float64                   `json:"payment_promotion"`                                         // [Required] The amount offset via payment promotion. (Only display for non cb sip affiliate shop. )
	CommissionFee                                      float64                   `json:"commission_fee"`                                            // [Required] <p>The commission fee charged by Shopee platform if applicable.&nbsp;</p><p><br /></p><p>For cb sip affiliate shop:&nbsp;commission_fee=commission_fee_pri * exchange_rate</p>
	ServiceFee                                         float64                   `json:"service_fee"`                                               // [Required] <p>Amount charged by Shopee to seller for additional services.</p><p><br /></p><p>For cb sip affiliate shop:&nbsp;service_fee=service_fee_pri * exchange_rate</p><p><br /></p><p>For tw shop, there will be pre-order service fee included</p>
	SellerTransactionFee                               float64                   `json:"seller_transaction_fee"`                                    // [Required] Tansaction fee paid by seller for the order. (Only display for non cb sip affiliate shop. )
	SellerLostCompensation                             float64                   `json:"seller_lost_compensation"`                                  // [Required] Compensation to seller in case of lost parcel. (Only display for non cb sip affiliate shop. )
	SellerCoinCashBack                                 float64                   `json:"seller_coin_cash_back"`                                     // [Required] Value of coins provided by Seller for purchasing with his or her store for the order. (Only display for non cb sip affiliate shop. )
	EscrowTax                                          float64                   `json:"escrow_tax"`                                                // [Required] Cross-border tax imposed by the Indonesian government on sellers. (Only display for non cb sip affiliate shop. )
	EstimatedShippingFee                               float64                   `json:"estimated_shipping_fee"`                                    // [Required] The estimated shipping fee is an estimation calculated by Shopee based on specific logistics courier's standard. (Only display for non cb sip affiliate shop. )
	FinalShippingFee                                   float64                   `json:"final_shipping_fee"`                                        // [Required] Final adjusted amount that seller has to bear as part of escrow. This amount could be negative or positive. (Only display for non cb sip affiliate shop. )
	ActualShippingFee                                  float64                   `json:"actual_shipping_fee"`                                       // [Required] The final shipping cost of order and it is positive. For Non-integrated logistics channel is 0. (Only display for non cb sip affiliate shop. )
	ShippingFeeSst                                     float64                   `json:"shipping_fee_sst"`                                          // [Required] <p>The Service Tax charged on Seller Paid Shipping Fee for forward shipping, based on Malaysia SST regulations&nbsp;for shipping fees on all orders. Definition of Seller Paid Shipping Fee is Actual Shipping Fee subtracted by sum of Shipping Fee Paid by Buyer and Shipping Rebate From Shopee. (Only applicable for non cb sip affiliate shop)<br /></p>
	OrderChargeableWeight                              int64                     `json:"order_chargeable_weight"`                                   // [Required] <p>For CB shop, display weight used to calculate actual_shipping_fee for this order.<br /></p>
	ShopeeShippingRebate                               float64                   `json:"shopee_shipping_rebate"`                                    // [Required] The platform shipping subsidy to the seller. (Only display for non cb sip affiliate shop. )
	ShippingFeeDiscountFrom_3pl                        float64                   `json:"shipping_fee_discount_from_3pl"`                            // [Required] The discount of shipping fee from 3PL. Currently only applicable to ID. (Only display for non cb sip affiliate shop. )
	SellerShippingDiscount                             float64                   `json:"seller_shipping_discount"`                                  // [Required] The shipping discount defined by seller. (Only display for non cb sip affiliate shop. )
	SellerVoucherCode                                  []string                  `json:"seller_voucher_code"`                                       // [Required] The list of voucher code provided by seller. (Only display for non cb sip affiliate shop. )
	DrcAdjustableRefund                                float64                   `json:"drc_adjustable_refund"`                                     // [Required] The adjustable refund amount from Shopee Dispute Resolution Center.
	CostOfGoodsSold                                    float64                   `json:"cost_of_goods_sold"`                                        // [Required] Final amount paid by the buyer for the items in a specific order. (Only display for non cb sip affiliate shop. )
	OriginalCostOfGoodsSold                            float64                   `json:"original_cost_of_goods_sold"`                               // [Required] Amount paid by the buyer for the items in a specific order. (Only display for non cb sip affiliate shop. )
	OriginalShopeeDiscount                             float64                   `json:"original_shopee_discount"`                                  // [Required] <p>Sum of each item Shopee discount of a specific order. This amount will return initial rebate value (i.e. remaining Shopee Item Rebate +&nbsp;remaining Shopee PIX Rebate) to seller. (Only display for non cb sip affiliate order. )</p>
	SellerReturnRefund                                 float64                   `json:"seller_return_refund"`                                      // [Required] Amount returned to Seller in the event of Partial Return.
	Items                                              []OrderIncomeItems        `json:"items"`                                                     // [Required] This object contains the detailed breakdown for all the items in this order, including regular items(non-activity) and activity items.
	EscrowAmountPri                                    float64                   `json:"escrow_amount_pri"`                                         // [Required] <p>The total amount in the prim currency that the seller is expected to receive for the order and will change before order completed . escrow_amount_pri=original_price_pri-seller_return_refund_pri-commission_fee_pri-service_fee_pri-drc_adjustable_refund_pri.(Only display for cb sip affiliate order. )</p>
	BuyerTotalAmountPri                                float64                   `json:"buyer_total_amount_pri"`                                    // [Required] <p>The total amount that paid by buyer in the primary currency. (Only display for cb sip affiliate order. )</p>
	OriginalPricePri                                   float64                   `json:"original_price_pri"`                                        // [Required] <p>The original price of the item before ANY promotion/discount in the primary currency. It returns the subtotal of that specific item if quantity exceeds 1.(Only display for cb sip affiliate order. )</p>
	SellerReturnRefundPri                              float64                   `json:"seller_return_refund_pri"`                                  // [Required] <p>Amount returned to Seller in the event of Partial Return in the primary currency. (Only display for cb sip affiliate shop. )</p>
	CommissionFeePri                                   float64                   `json:"commission_fee_pri"`                                        // [Required] <p>The commission fee charged by Shopee platform if applicable in the primary currency. (Only display for cb sip affiliate shop. )</p>
	ServiceFeePri                                      float64                   `json:"service_fee_pri"`                                           // [Required] <p>Amount charged by Shopee to seller for additional services in the primary currency. (Only display for cb sip affiliate shop. )</p>
	DrcAdjustableRefundPri                             float64                   `json:"drc_adjustable_refund_pri"`                                 // [Required] <p>The adjustable refund amount from Shopee Dispute Resolution Center in the primary currency&nbsp;after proration. (Only applicable for cb sip affiliate shop.)</p>
	PriCurrency                                        string                    `json:"pri_currency"`                                              // [Required] <p>The currency of the country or region where the shop that real seller operates. (Only display for cb sip affiliate shop. )</p>
	AffCurrency                                        string                    `json:"aff_currency"`                                              // [Required] <p>The currency of the country or region where shop opened in. (Only display for cb sip affiliate shop. )</p>
	ExchangeRate                                       float64                   `json:"exchange_rate"`                                             // [Required] Exchange rate from primary shop currency to affiliate shop currency.
	ReverseShippingFee                                 float64                   `json:"reverse_shipping_fee"`                                      // [Required] Shopee charges the reverse shipping fee for the returned order.The value of this field will be non-negative.
	ReverseShippingFeeSst                              float64                   `json:"reverse_shipping_fee_sst"`                                  // [Required] <p>The Service Tax charged on Reverse Shipping Fee for reverse shipping, based on Malaysia SST regulations&nbsp;for shipping fees on all orders. (Only applicable for non cb sip affiliate shop)<br /></p>
	FinalProductProtection                             float64                   `json:"final_product_protection"`                                  // [Required] <p>The total amount of product protection purchased during placing an order.&nbsp;</p>
	CreditCardPromotion                                float64                   `json:"credit_card_promotion"`                                     // [Required] This value indicate the offset via credit card promotion.
	CreditCardTransactionFee                           float64                   `json:"credit_card_transaction_fee"`                               // [Required] <p>This value indicate the total transaction fee.</p><p>credit_card_transaction_fee=buyer_transaction_fee+seller_transaction_fee<br /></p>
	FinalProductVatTax                                 float64                   `json:"final_product_vat_tax"`                                     // [Required] Value-added Tax is required for online purchases based on EU Value-added Tax regulations . (Only display for non cb sip affiliate shop. )
	FinalShippingVatTax                                float64                   `json:"final_shipping_vat_tax"`                                    // [Required] <p>Value-added Tax for product price is required for online purchases based on EU Value-added Tax regulations. (Only applicable for non cb sip affiliate shop. )<br /></p>
	CampaignFee                                        float64                   `json:"campaign_fee"`                                              // [Required] <p>The campaign fee charged by Shopee platform. Only available for some local Indonesian shops.<br /></p>
	SipSubsidy                                         float64                   `json:"sip_subsidy"`                                               // [Required] <p>The SIP subsidy amount is the difference between the item settlement price set by seller and item price actually paid by buyer. (Only available for CB SIP A Shops)<br /></p>
	SipSubsidyPri                                      float64                   `json:"sip_subsidy_pri"`                                           // [Required] <p>The SIP subsidy amount is the difference between the item settlement price set by seller and item price actually paid by buyer. This value is in the primary currency. (Only available for CB SIP A Shops)<br /></p>
	RsfSellerProtectionFeeClaimAmount                  float64                   `json:"rsf_seller_protection_fee_claim_amount"`                    // [Required] <p>The insurance claim amount if seller opt in to insurance program. this covers Reverse shipping Fee in the case of RR. As per Jun 2024:<br />- For ID &amp; MY Local: After Extension on coverage to FSF due to RR. this claim amount will consist of FSF + RSF claim.<br />- For PH local: This will only cover RSF claim<br /><br />will be updated if there's any RR/cancellation<br /></p>
	ShippingSellerProtectionFeeAmount                  float64                   `json:"shipping_seller_protection_fee_amount"`                     // [Required] <p>Service fee charged to seller in MY,ID,PH Local (as per Jun 2024) due to additional program purchased.<br /></p>
	FinalEscrowProductGst                              float64                   `json:"final_escrow_product_gst"`                                  // [Required] <p>Goods and Service Tax for product price is required for imported goods (overseas orders) based on Singapore GST regulations. (Only applicable for non cb sip affiliate shop selling in Singapore)<br /></p>
	FinalEscrowShippingGst                             float64                   `json:"final_escrow_shipping_gst"`                                 // [Required] <p>Goods and Service Tax for shipping fee is required for imported goods (overseas orders) based on Singapore GST regulations. (Only applicable for non cb sip affiliate shop selling in Singapore.)<br /></p>
	DeliverySellerProtectionFeePremiumAmount           float64                   `json:"delivery_seller_protection_fee_premium_amount"`             // [Required] <p>[Currently apply to ID &amp; local orders only] An insurance premium charged to seller at the time parcel is picked up by 3PL for insurance in case of parcel lost/damaged by 3PL.<br /></p>
	OrderAdjustment                                    []OrderAdjustment         `json:"order_adjustment"`                                          // [Required] <p>Order level adjustment transaction information.<br /></p><p>If the&nbsp;order without adjustment, no returned of the field.</p>
	TotalAdjustmentAmount                              float64                   `json:"total_adjustment_amount"`                                   // [Required] <p>Total adjustment made to the order.<br /></p>
	EscrowAmountAfterAdjustment                        float64                   `json:"escrow_amount_after_adjustment"`                            // [Required] <p>Final income seller can get from this order after deduct the order-level adjustment.<br /></p>
	OrderAmsCommissionFee                              float64                   `json:"order_ams_commission_fee"`                                  // [Required] <p>The amount of affiliate commission fee for this order. Applicable for orders sold via the Affiliate Program.<br /></p>
	BuyerPaymentMethod                                 string                    `json:"buyer_payment_method"`                                      // [Required] <p>The payment method buyer used when do the order checkout.<br /></p>
	InstalmentPlan                                     string                    `json:"instalment_plan"`                                           // [Required] <p>The instalment plan buyer chosen when do the order checkout. Only applicable when payment method support instalment.<br /></p>
	SalesTaxOnLvg                                      float64                   `json:"sales_tax_on_lvg"`                                          // [Required] <p>Sales Tax on Low Value Goods (LVG) is required for imported goods (overseas orders) based on Malaysia SST regulations&nbsp;for selective orders (e.g. CB LVG orders in MY)<br /></p>
	FinalReturnToSellerShippingFee                     float64                   `json:"final_return_to_seller_shipping_fee"`                       // [Required] <p>The amount of fee charged to seller (if any) for the failed delivery order.</p>
	WithholdingTax                                     float64                   `json:"withholding_tax"`                                           // [Required] <p>Only for PH and ID local shops.</p><p><br /></p><p><b>PH</b>: According to regulations issued by Bureau of Internal Revenue in PH, the Withholding Tax is applied to the gross remittances sent by Shopee to online suppliers of goods and services.<br /><br /><b>ID</b>: According to regulations issued by Directorate General of Taxation in ID, the Withholding Tax is applied to the income stated in the invoice generated via Shopee related to Seller and/or Merchants' sales in Shopee's platform.</p>
	OverseasReturnServiceFee                           float64                   `json:"overseas_return_service_fee"`                               // [Required] <p>This is overseas return service fee charged to sellers who register ORS program.(Only applicable for non cb sip affiliate shop)<br /></p>
	ProratedCoinsValueOffsetReturnItems                float64                   `json:"prorated_coins_value_offset_return_items"`                  // [Required] <p>This is the prorated value from cash equivalent of coin offset due to adjustable RR.This field will only be updated when there is an adjustable RR. If it's a full RR or normal order will response 0.<br /></p>
	ProratedShopeeVoucherOffsetReturnItems             float64                   `json:"prorated_shopee_voucher_offset_return_items"`               // [Required] <p>This is the prorated refund value from shopee voucher discount due to adjustable RR.This field will only be updated when there is an adjustable RR.&nbsp;If it's a full RR or normal order will response 0.<br /></p>
	ProratedSellerVoucherOffsetReturnItems             float64                   `json:"prorated_seller_voucher_offset_return_items"`               // [Required] <p>This is the prorated refund value from seller voucher discount due to adjustable RR.This field will only be updated when there is an adjustable RR.&nbsp;If it's a full RR or normal order will response 0.<br /></p>
	ProratedPaymentChannelPromoBankOffsetReturnItems   float64                   `json:"prorated_payment_channel_promo_bank_offset_return_items"`   // [Required] <p>This is the prorated value from bank payment channel promo due to adjustable RR.This field will only be updated when there is an adjustable RR.If it's a full RR or normal order will response 0.<br /></p>
	ProratedPaymentChannelPromoShopeeOffsetReturnItems float64                   `json:"prorated_payment_channel_promo_shopee_offset_return_items"` // [Required] <p>This is the prorated value from shopee payment channel promo due to adjustable RR.This field will only be updated when there is an adjustable RR.If it's a full RR or normal order will response 0.<br /></p>
	FsfSellerProtectionFeeClaimAmount                  float64                   `json:"fsf_seller_protection_fee_claim_amount"`                    // [Required] <p>The claim amount given to seller if seller opt in to shipping fee service program. this covers Forward Shipping Fee cost in the case of cancelled due to Failed delivery.&nbsp;<br /></p>
	VatOnImportedGoods                                 float64                   `json:"vat_on_imported_goods"`                                     // [Required] <p>7% VAT is charged for imported goods entering Thailand.<br /></p><p>8% VAT is charged for imported goods entering Vietnam</p>
	TenureInfoList                                     *TenureInfoList           `json:"tenure_info_list"`                                          // [Required]
	WithholdingVatTax                                  float64                   `json:"withholding_vat_tax"`                                       // [Required] <p>By VN law, E-commerce platforms need to Withhold VAT tax applicable to all VN business household and VN individual sellers</p>
	WithholdingPitTax                                  float64                   `json:"withholding_pit_tax"`                                       // [Required] <p>By VN law, E-commerce platforms need to Withhold Personal Income Tax applicable to all VN business household and VN individual sellers</p>
	TaxRegistrationCode                                string                    `json:"tax_registration_code"`                                     // [Required] <p>For VN Withholding Tax. This is the Tax Registration Number (TRN) from Seller Info (Business information) of the seller at the point of order creation</p>
	SellerOrderProcessingFee                           float64                   `json:"seller_order_processing_fee"`                               // [Required] <p>Order Processing Fee is the amount charged to sellers for every order created.</p>
	BuyerPaidPackagingFee                              float64                   `json:"buyer_paid_packaging_fee"`                                  // [Required] <p>The fee charged to the buyer for packaging materials</p>
	TradeInBonusBySeller                               float64                   `json:"trade_in_bonus_by_seller"`                                  // [Required] <p>The discount provided by Seller/ Retailers for buyers who opt for trade-in.</p>
	FbsFee                                             float64                   `json:"fbs_fee"`                                                   // [Required] <p>Fulfilled by Shopee (FBS) Fee applied to this order, covering costs such as handling and storage and packaging. Only for PH Local Orders.</p>
	NetCommissionFee                                   float64                   `json:"net_commission_fee"`                                        // [Required] <p>The respective fee amounts after the proportional rebate deduction.The total net commission fee applied to the order, calculated as the sum of all net commission fee items.</p><p>-only for BR local sellers.</p><br />
	NetServiceFee                                      float64                   `json:"net_service_fee"`                                           // [Required] <p>The respective fee amounts after the proportional rebate deduction.The total net service fee applied to the order, calculated as the sum of all net service fee items.</p><p>-only for BR local sellers.</p>
	NetCommissionFeeInfoList                           *NetCommissionFeeInfoList `json:"net_commission_fee_info_list"`                              // [Required] <p>Returns a breakdown of the net commission fees.</p><p>-only for BR local sellers.</p>
	NetServiceFeeInfoList                              *NetServiceFeeInfoList    `json:"net_service_fee_info_list"`                                 // [Required] <p>Returns a breakdown of the net service fees.<br /></p><p>-only for BR local sellers.</p>
	SellerProductRebate                                *SellerProductRebate      `json:"seller_product_rebate"`                                     // [Required] <p>The shopee rebate borne by seller.</p><p>-only for BR local sellers.</p>
	PixDiscount                                        float64                   `json:"pix_discount"`                                              // [Required] <p>[BR only]Final sum of pix discount of a specific order. (Only display for non cb sip affiliate shop.)</p>
	ProratedPixDiscountOffsetReturnItems               float64                   `json:"prorated_pix_discount_offset_return_items"`                 // [Required] <p>[BR only]This is the prorated value from pix discount due to adjustable RR. This field will only be updated when there is an adjustable RR. If it's a full RR or normal order, will response 0.</p>
	AdsEscrowTopUpFeeOrTechnicalSupportFee             float64                   `json:"ads_escrow_top_up_fee_or_technical_support_fee"`            // [Required] <p>Includes both ads escrow top up fee (auto escrow top up to your ads balance) and technical support fee (charged by Shopee)</p><p>The actual fee type included in this field varies depending on the seller type and selling region, and may represent one of the following in Shopee Seller Center:</p>   <p>Ads Escrow Top-Up Fee</p>   <p>For local MY TH SG VN PH ID sellers and CNCB / JPCB / KRCB sellers selling in PH and ID</p>   <p>For JPCB sellers selling in SG, MY, TH, and VN</p>     <p>Technical Support Fee</p>   <p>For CNCB sellers selling in SG, MY, TH, and VN</p>     <p>Traffic Growth Fee</p>   <p>For KRCB sellers selling in SG, MY, TH, and VN</p>
}

type OrderIncomeItems struct {
	ItemId                    int64           `json:"item_id"`                      // [Required] <p>ID of item<br /></p>
	ItemName                  string          `json:"item_name"`                    // [Required] <p>Name of item<br /></p>
	ItemSku                   string          `json:"item_sku"`                     // [Required] <p>A item SKU (stock keeping unit) is an identifier defined by a seller, sometimes called parent SKU. Item SKU can be assigned to an item in Shopee Listings.<br /></p>
	ModelId                   int64           `json:"model_id"`                     // [Required] <p>ID of the model that belongs to the same item.<br /></p>
	ModelName                 string          `json:"model_name"`                   // [Required] <p>Name of the model that belongs to the same item. A seller can offer variations of the same item. For example, the seller could create a fixed-priced listing for a t-shirt design and offer the shirt in different colors and sizes. In this case, each color and size combination is a separate variation. Each variation can have a different quantity and price.<br /></p>
	ModelSku                  string          `json:"model_sku"`                    // [Required] <p>A model SKU (stock keeping unit) is an identifier defined by a seller. It is only intended for the seller's use. Many sellers assign a SKU to an item of a specific type, size, and color, which are variations of one item in Shopee Listings.<br /></p>
	OriginalPrice             float64         `json:"original_price"`               // [Required] <p>The original price of the item before ANY promotion/discount in the listing currency. It returns the subtotal of that specific item if quantity exceeds 1.<br /></p>
	OriginalPricePri          float64         `json:"original_price_pri"`           // [Required] <p>The agreed settlement price of items used as settlement basis, amount is in the primary currency. (Only display for CB SIP affiliate shop.)<br /></p>
	SellingPrice              float64         `json:"selling_price"`                // [Required] <p>For non-bundle deal case, this value will be same like item original_price; For bundle deal case, this value will be price of sum of item price before bundle deal promo but after item promo. It returns the subtotal of that specific item if quantity exceeds 1 (Only display for non cb sip affiliate shop.)<br /></p>
	DiscountedPrice           float64         `json:"discounted_price"`             // [Required] <p>The after-discount price of the item in the listing currency. It returns the subtotal of that specific item if quantity exceeds 1. If there is no discount, this value will be the same as that of original_price.<br /></p>
	SellerDiscount            float64         `json:"seller_discount"`              // [Required] <p>The discount provided by seller for this item<br /></p>
	ShopeeDiscount            float64         `json:"shopee_discount"`              // [Required] <p>The discount provided by Shopee for this item<br /></p>
	DiscountFromCoin          float64         `json:"discount_from_coin"`           // [Required] <p>The offset of this item when the buyer consumed Shopee Coins upon checkout. In case of bundle deal item, this value will return 0. Due to technical restriction, this field will return incorrect value under bundle deal case if we don’t configure it to 0.<br /></p>
	DiscountFromVoucherShopee float64         `json:"discount_from_voucher_shopee"` // [Required] <p>The offset of this item when the buyer use Shopee voucher.&nbsp;</p>
	DiscountFromVoucherSeller float64         `json:"discount_from_voucher_seller"` // [Required] <p>The offset of this item when the buyer use seller-specific voucher.</p>
	ActivityType              string          `json:"activity_type"`                // [Required] <p>The type of the item, default is "". If the item is a bundle item the value is "bundle_deal", and if a add on deal item, the value is "add_on_deal"<br /></p>
	ActivityId                int64           `json:"activity_id"`                  // [Required] <p>If bundle_deal the is id of bundle deal, if add_on_deal this is id of add on deal.<br /></p>
	IsMainItem                bool            `json:"is_main_item"`                 // [Required] <p>Meaning a main or sub item for add_on_deal.<br /></p>
	QuantityPurchased         int64           `json:"quantity_purchased"`           // [Required] <p>This value indicates the number of identical items purchased at the same time by the same buyer from one listing/item.<br /></p>
	IsB2cShopItem             bool            `json:"is_b2c_shop_item"`             // [Required] <p>Indicates whether this is a B2C owned item.<br /></p>
	AmsCommissionFee          float64         `json:"ams_commission_fee"`           // [Required] <p>The amount of affiliate commission fee. Applicable for items sold via the Affiliate Program.<br /></p>
	IsKit                     bool            `json:"is_kit"`                       // [Required] <p>Indicates if this item is a kit. (Only for BR Local Sellers)</p>
	KitItems                  *KitItems       `json:"kit_items"`                    // [Required] <p>Only applicable to BR local sellers</p>
	PromotionList             []ItemPromotion `json:"promotion_list"`               // [Required]
}

type OrderItem struct {
	ItemId                 int64           `json:"item_id"`                  // [Required] Shopee's unique identifier for an item.
	ItemName               string          `json:"item_name"`                // [Required] The name of the item.
	ItemSku                string          `json:"item_sku"`                 // [Required]  A item SKU (stock keeping unit) is an identifier defined by a seller, sometimes called parent SKU. Item SKU can be assigned to an item in Shopee Listings.
	ModelId                int64           `json:"model_id"`                 // [Required] ID of the model that belongs to the same item.
	ModelName              string          `json:"model_name"`               // [Required] Name of the model that belongs to the same item. A seller can offer models of the same item. For example, the seller could create a fixed-priced listing for a t-shirt design and offer the shirt in different colors and sizes. In this case, each color and size combination is a separate model. Each model can have a different quantity and price.
	ModelSku               string          `json:"model_sku"`                // [Required] A model SKU (stock keeping unit) is an identifier defined by a seller. It is only intended for the seller's use. Many sellers assign a SKU to an item of a specific type, size, and color, which are models of one item in Shopee Listings.
	ModelQuantityPurchased int64           `json:"model_quantity_purchased"` // [Required] The number of identical items purchased at the same time by the same buyer from one listing/item.
	ModelOriginalPrice     float64         `json:"model_original_price"`     // [Required] The original price of the item in the listing currency.
	ModelDiscountedPrice   float64         `json:"model_discounted_price"`   // [Required] The after-discount price of the item in the listing currency. If there is no discount, this value will be same as that of model_original_price. In case of bundle deal item, this value will return 0 as by design bundle deal discount will not be breakdown to item/model level. Due to technical restriction, the value will return the price before bundle deal if we don't configure it to 0. Please call GetEscrowDetails if you want to calculate item-level discounted price for bundle deal item.
	Wholesale              bool            `json:"wholesale"`                // [Required] This value indicates whether buyer buy the order item in wholesale price.
	Weight                 float64         `json:"weight"`                   // [Required] The weight of the item
	AddOnDeal              bool            `json:"add_on_deal"`              // [Required] To indicate if this item belongs to an addon deal.
	MainItem               bool            `json:"main_item"`                // [Required] To indicate if this item is main item or sub item. True means main item, false means sub item.
	AddOnDealId            int64           `json:"add_on_deal_id"`           // [Required] A unique ID to distinguish groups of items in Cart, and Order. (e.g. AddOnDeal)
	PromotionType          string          `json:"promotion_type"`           // [Required] <p>Available type：product_promotion, flash_sale, bundle_deal, add_on_deal_main, add_on_deal_sub.</p><p><br /></p><p>For items which attend multiple promotions will only show one promotion, the order of priority is:&nbsp;</p><p>bundle_deal &gt; add_on_deal_main &gt; add_on_deal_sub &gt; product_promotion &gt;flash_sale</p>
	PromotionId            int64           `json:"promotion_id"`             // [Required] The ID of the promotion.
	OrderItemId            int64           `json:"order_item_id"`            // [Required] The identify of order item.
	PromotionGroupId       int64           `json:"promotion_group_id"`       // [Required] The identify of product promotion.
	ImageInfo              *OptionImage    `json:"image_info"`               // [Required] Image info of the product.
	ProductLocationId      string          `json:"product_location_id"`      // [Required] The fulfilment warehouse ID(s) of the items in the order. (Multi-Warehouse sellers only)
	IsPrescriptionItem     bool            `json:"is_prescription_item"`     // [Required] <p>To indicate if this item is prescription item<br /></p>
	IsB2cOwnedItem         bool            `json:"is_b2c_owned_item"`        // [Required] <p>determine if item is B2C_shop_item</p><p>It should be `<b>is_b2c_shop_item</b>` but it was a bug from dev. Then now it's <b>is_b2c_owned_item</b></p>
	ConsultationId         string          `json:"consultation_id"`          // [Required] <p>An identifier of teleconsultation session which buyer did to order this item. Empty if item is not ordered through teleconsultation session</p>
	PromotionList          []ItemPromotion `json:"promotion_list"`           // [Required]
	HotListingItem         bool            `json:"hot_listing_item"`         // [Required] <p>[Only for PH,TH,VN,MY,BR,TW] True if the item is hot listing.</p>
}

type OrderPackage struct {
	PackageNumber          string             `json:"package_number"`           // [Required] Shopee's unique identifier for the package under an order.
	LogisticsStatus        LogisticsStatus    `json:"logistics_status"`         // [Required] The Shopee logistics status for the order. Applicable values: See Data Definition-LogisticsStatus.
	LogisticsChannelId     int64              `json:"logistics_channel_id"`     // [Required] <p>The identity of logistic channel.</p>
	ShippingCarrier        string             `json:"shipping_carrier"`         // [Required] <p>The logistics service provider that the buyer selected for the order to deliver items.&nbsp;</p><p><br /></p><p>Note: If logistics_channel_id is 90021, 90025 or 90026, service_code will be appended,&nbsp;e.g., Entrega Turbo - M1020.</p>
	AllowSelfDesignAwb     bool               `json:"allow_self_design_awb"`    // [Required] <p>To indicate whether the package allows for self-designed AWB, if allow_self_design_awb returns false, it means that the package does not allow for self-designed AWB and only the system-AWB can be used.</p>
	ItemList               []OrderPackageItem `json:"item_list"`                // [Required] The lis of items.
	ParcelChargeableWeight int64              `json:"parcel_chargeable_weight"` // [Required] display weight used to calculate ASF for this parcel
	GroupShipmentId        int64              `json:"group_shipment_id"`        // [Required] <p>The common identifier for multiple orders combined in the same parcel.<br /></p>
	VirtualContactNumber   string             `json:"virtual_contact_number"`   // [Required] <p>[Only for TW non-integrated channel] The virtual phone number to contact the recipient.<br /></p>
	PackageQueryNumber     string             `json:"package_query_number"`     // [Required] <p>[Only for TW non-integrated channel] The query number used in virtual phone number calls to contact the recipient of this package.<br /></p>
	SortingGroup           string             `json:"sorting_group"`            // [Required] <p>[Only for TW 30029 channel] This field indicate the sorting group value of the package. This field is only available for logistics_channel_id = 30029 and after the package has been arranged for shipment.</p>
}

type OrderPackageItem struct {
	ItemId            int64  `json:"item_id"`             // [Required] Shopee's unique identifier for an item.
	ModelId           int64  `json:"model_id"`            // [Required] Shopee's unique identifier for a model.
	ModelQuantity     int64  `json:"model_quantity"`      // [Required] <p>The number of identical items/variations purchased at the same time by the same buyer from one listing/item.</p>
	OrderItemId       int64  `json:"order_item_id"`       // [Required] <p>The identify of order item. For items in one same bundle deal promotion, the order_item_id should share the same id, such as 1,2. For items not in bundle deal promotion, the order_item_id should be the same as item_id.<br /></p>
	PromotionGroupId  int64  `json:"promotion_group_id"`  // [Required] <p>The identify of product promotion.<br /></p>
	ProductLocationId string `json:"product_location_id"` // [Required] <p>The warehouse ID of the item.<br /></p>
}

type OutboundQty struct {
	OutboundTotal    int64 `json:"outbound_total"`    // [Required] <p>Total outbound quantity during the selected time period.</p>
	OutboundSold     int64 `json:"outbound_sold"`     // [Required] <p>"Total sold quantity during the selected time period."</p>
	OutboundReturned int64 `json:"outbound_returned"` // [Required] <p>Total merchant return quantity during the selected time period.</p>
	OutboundDisposed int64 `json:"outbound_disposed"` // [Required] <p>Total disposal quantity during the selected time period.</p>
}

type OutletShopInfo struct {
	OutletShopId int64 `json:"outlet_shop_id"` // [Required] <p>Shop ID of the Outlet Shop.</p>
}

type OverallPerformance struct {
	Rating              int64 `json:"rating"`                // [Required] <p>Overall Performance:&nbsp;</p><p>Poor = 1<br /></p><p>ImprovementNeeded = 2<br /></p><p>Good = 3<br /></p><p>Excellent = 4</p>
	FulfillmentFailed   int64 `json:"fulfillment_failed"`    // [Required] <p>The number of metrics that did not meet target under Fulfillment Performance type.</p>
	ListingFailed       int64 `json:"listing_failed"`        // [Required] <p>The number of metrics that did not meet target under Listing Performance type.<br /></p>
	CustomServiceFailed int64 `json:"custom_service_failed"` // [Required] <p>The number of metrics that did not meet target under Customer Service Performance type.</p>
}

type Package struct {
	ItemList []PackageItem `json:"item_list"` // [Required] <p>The list of items under the same package.</p>
}

type PackageItem struct {
	ItemId           int64 `json:"item_id"`            // [Required] Shopee's unique identifier for an item.
	ModelId          int64 `json:"model_id"`           // [Required] Shopee's unique identifier for a model.
	OrderItemId      int64 `json:"order_item_id"`      // [Required] <p>The identify of order item. For items in one same bundle deal promotion, the order_item_id should share the same id, such as 1,2. For items not in bundle deal promotion, the order_item_id should be the same as item_id.</p>
	PromotionGroupId int64 `json:"promotion_group_id"` // [Required] <p>The identify of product promotion. For items in one same add on deal promotion, the promotion_group_id should share the same id. For items not in add on deal promotion, the promotion_group_id should be 0. And the data is from group_id of shopee.orders.GetOrderDetails.</p>
	ModelQuantity    int64 `json:"model_quantity"`     // [Required] <p>The number of identical items put in the package.<br /></p>
}

type Packages struct {
	OrderSn            string `json:"order_sn"`             // [Required] <p>Shopee's unique identifier for an order.</p>
	PackageNumber      string `json:"package_number"`       // [Required] <p>Shopee's unique identifier for the package under an order</p>
	LogisticsChannelId int64  `json:"logistics_channel_id"` // [Required] <p>The identity of logistic channel. </p><p></p>
	ProductLocationId  string `json:"product_location_id"`  // [Required] Just use this field to pass the next step of Mass ArrangeShipment<br />
	SortingGroup       string `json:"sorting_group"`        // [Required] <p>[Only for TW 30029 channel] This field indicate the sorting group value of the package. This field is only available for logistics_channel_id = 30029 and after the package has been arranged for shipment.</p>
	IsShipmentArranged bool   `json:"is_shipment_arranged"` // [Required] <p>Only effective when the package's logistics_status/fulfillment_status is LOGISTICS_READY.&nbsp;</p><p><br /></p><p>This parameter further distinguishes between two scenarios:</p><p>- true: Package shipment has been arranged (Seller has processed shipment, system is generating tracking number, not yet updated to LOGISTICS_REQUEST_CREATED, no duplicate action needed)</p><p>- false: Package awaiting shipment arrangement (Seller hasn't processed shipment yet, shipping arrangement required)</p>
}

type PackagingFee struct {
	Value float64 `json:"value"` // [Required] <p>The packaging fee price in the seller's local currency.</p>
}

type PageInfo struct {
	HasNextPage bool   `json:"has_next_page"` // [Required]
	NextOffset  string `json:"next_offset"`   // [Required]
}

type Pagination struct {
	Cursor   string `json:"cursor"`    // [Required] <p>Token to retrieve the next page of results. Returns empty if there is no more data.</p>
	PageSize int64  `json:"page_size"` // [Required] <p>Number of records returned per page.</p>
}

type PairIds struct {
	CriteriaId   int64             `json:"criteria_id"`   // [Required]
	CategoryList []PairIdsCategory `json:"category_list"` // [Required] <p>these are the categories that the shop has items, and the criteria will apply to these categories</p>
}

type PairIdsCategory struct {
	CategoryId int64  `json:"category_id"` // [Required] <p>o means this is All category</p>
	Name       string `json:"name"`        // [Required] <p>category name</p>
	ParentId   int64  `json:"parent_id"`   // [Required] <p>the parent category id, 0 means this category is L1 category</p>
}

type PayOrder struct {
	OrderSn  string `json:"order_sn"`  // [Required] Shopee's unique identifier for an order.
	ShopName string `json:"shop_name"` // [Required] Name of the shop.
}

type PaymentInfo struct {
	PaymentMethod            string  `json:"payment_method"`             // [Required] <p>[Only for BR] Payment method used in the order, such as Credit Card, Debit Card, Pix, etc.</p>
	PaymentProcessorRegister string  `json:"payment_processor_register"` // [Required] <p>[Only for BR] CNPJ of the payment processor handling the transaction.</p>
	CardBrand                string  `json:"card_brand"`                 // [Required] <p>[Only for BR] Card brand for credit or debit transactions, such as VISA, MASTER, etc. Empty string for Pix payments.</p>
	TransactionId            string  `json:"transaction_id"`             // [Required] <p>[Only for BR] Payment authorization code generated by the bank or payment processor to validate the transaction.</p>
	PaymentAmount            float64 `json:"payment_amount"`             // [Required] <p>[Only for BR] Amount paid by the corresponding payment method.<br /></p>
}

type Payout struct {
	PayoutInfo            *PayoutInfo         `json:"payout_info"`             // [Required] The information of payout.
	EscrowList            []Escrow            `json:"escrow_list"`             // [Required]
	OfflineAdjustmentList []OfflineAdjustment `json:"offline_adjustment_list"` // [Required] The list of offline adjustments.
}

type PayoutInfo struct {
	FromCurrency   string  `json:"from_currency"`   // [Required] The settlement currency of orders.
	PayoutCurrency string  `json:"payout_currency"` // [Required] The actual currency of payout.
	FromAmount     float64 `json:"from_amount"`     // [Required] The settlement amount.
	PayoutAmount   float64 `json:"payout_amount"`   // [Required] The actual amount of payout.
	ExchangeRate   string  `json:"exchange_rate"`   // [Required] The exchange rate.
	PayoutTime     int64   `json:"payout_time"`     // [Required] The time of payout.
	PayService     string  `json:"pay_service"`     // [Required] The service provider of seller. Available value: payoneer, pingpong, lianlian.
	PayeeId        string  `json:"payee_id"`        // [Required] Seller's account to receive the payout.
}

type PayoutList struct {
	FromCurrency      string  `json:"from_currency"`       // [Required] <p>The settlement currency of orders.<br /></p>
	PayoutCurrency    string  `json:"payout_currency"`     // [Required] <p>The actual currency of payout.<br /></p>
	FromAmount        float64 `json:"from_amount"`         // [Required] <p>The settlement amount.<br /></p>
	PayoutAmount      float64 `json:"payout_amount"`       // [Required] <p>The actual amount of payout.<br /></p>
	ExchangeRate      string  `json:"exchange_rate"`       // [Required] <p>The exchange rate.<br /></p>
	PayoutTime        int64   `json:"payout_time"`         // [Required] <p>The time of payout.<br /></p>
	PayService        string  `json:"pay_service"`         // [Required] <p>The service provider of seller. Available value: payoneer, pingpong, lianlian.<br /></p>
	PayeeId           string  `json:"payee_id"`            // [Required] <p>Seller's account to receive the payout.<br /></p>
	EncryptedPayoutId string  `json:"encrypted_payout_id"` // [Required] <p>payout id used to query API "v2.get_billing_item_info" as request parameters. User can get detailed billing items under current payout<br /></p>
}

type PenaltyPoint struct {
	IssueTime        int64 `json:"issue_time"`         // [Required] <p>The time when penalty points are issued.</p>
	LatestPointNum   int64 `json:"latest_point_num"`   // [Required] <p>The latest penalty points issued under current penalty point record.&nbsp;</p><p><br /></p><p>If seller raised appeal for this penalty point record and the appeal has been approved and Shopee adjusted the penalty point, then the original_point_num returns the penalty point before the adjustment, and latest_point_num returns the penalty point after the adjustment.<br /></p>
	OriginalPointNum int64 `json:"original_point_num"` // [Required] <p>The original penalty points&nbsp;issued under current penalty point record.</p><p><br /></p><p>If seller raised appeal for this penalty point record and the appeal has been approved and Shopee adjusted the penalty point, then the original_point_num returns the penalty point before the adjustment, and latest_point_num returns the penalty point after the adjustment.</p>
	ReferenceId      int64 `json:"reference_id"`       // [Required] <p>Reference ID for this penalty point record.</p>
	ViolationType    int64 `json:"violation_type"`     // [Required] <p>Applicable values:&nbsp;</p><p>5: High Late Shipment Rate</p><p>6: High Non-fulfilment Rate</p><p>7: High number of non-fulfilled orders</p><p>8: High number of late shipped orders</p><p>9: Prohibited Listings</p><p>10: Counterfeit / IP infringement</p><p>11: Spam</p><p>12: Copy/Steal images</p><p>13: Re-uploading deleted listings with no change</p><p>14: Bought counterfeit from mall</p><p>15: Counterfeit caught by Shopee</p><p>16: High percentage of pre-order listings</p><p>17: Confirmed Fraud attempts (total)</p><p>18: Confirmed Fraud attempts per week (All with vouchers only)</p><p>19: Fake return address</p><p>20: Shipping fraud/abuse</p><p>21: High No. of Non-responded Chat</p><p>22: Rude chat replies</p><p>23: Request buyer to cancel order</p><p>24: Rude reply to buyer's review</p><p>25: Violate Return/Refund policy</p><p>101: Tier Reason</p><p>3026: Misuse of Shopee’s IP</p><p>3028: Violate Shop Name Regulations</p><p>3030: Direct transactions outside of the Shopee platform</p><p>3032: Shipping empty / incomplete parcels</p><p>3034: Severe Violations on Shopee Feed</p><p>3036: Severe Violations on Shopee LIVE</p><p>3038: Misuse of Local Vendor Tag</p><p>3040: Use of misleading shop tag in listing image</p><p>3042: Counterfeit / IP Infringement test</p><p>3044: Repeat Offender - IP infringement and Counterfeit listings</p><p>3046: Violation of Live Animals Selling Policy</p><p>3048: Chat Spam</p><p>3050: High Overseas Return Refunds Rate</p><p>3052: Privacy breach in buyer's review reply</p><p>3054: Order Brushing</p><p>3056: porn image</p><p>3058: Incorrect Product Categories</p><p>3060: Extremely High Non-Fulfilment Rate</p><p>3062: Penalty of Affiliate Marketing Solution (AMS) Overdue Invoice Payment</p><p>3064: Government-related listing</p><p>3066: Listing invalid gifted items</p><p>3068: High non-fulfilment rate (Next Day Delivery Orders)</p><p>3070: High Late Shipment Rate (Next Day Delivery Orders)</p><p>3072: OPFR Violation Value</p><p>3074: Direct transactions outside Shopee platform via chat</p><p>3090: Prohibited Listings-Extreme Violations</p><p>3091: Prohibited Listings-High Violations</p><p>3092: Prohibited Listings-Mid Violations</p><p>3093: Prohibited Listings-Low Violations</p><p>3094: Counterfeit Listings-Extreme Violations</p><p>3095: Counterfeit Listings-High Violations</p><p>3096: Counterfeit Listings-Mid Violations</p><p>3097: Counterfeit Listings-Low Violations</p><p>3098: Spam Listings-Extreme Violations</p><p>3099: Spam Listings-High Violations</p><p>3100: Spam Listings-Mid Violations</p><p>3101: Spam Listings-Low Violations</p><p>3145: Return/Refund Rate (Non-integrated Channel)</p><p>4130: Poor Product Quality</p>
}

type Photo struct {
	Url       string `json:"url"`       // [Required] <p>The video url in dispute proof.&nbsp;</p>
	Thumbnail string `json:"thumbnail"` // [Required] <p>The thumbnail of video</p>
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

type PostCommentRequest struct {
	SessionId int64  `json:"session_id"` // [Required] <p>The identifier of livestream session.</p>
	Content   string `json:"content"`    // [Required] <p>The comment content, cannot exceed 150 characters.</p>
}

type PostCommentResponse struct {
	BaseResponse `json:",inline"`        // Common response fields
	Response     PostCommentResponseData `json:"response"` // Actual response data
}

type PostCommentResponseData struct {
	CommentId int64 `json:"comment_id"` // [Required] <p>The identifier of the comment.</p>
}

type PreOrder struct {
	IsPreOrder bool  `json:"is_pre_order"` // [Required]
	DaysToShip int64 `json:"days_to_ship"` // [Required]
}

type PreOrderListing struct {
	ItemId                int64 `json:"item_id"`                  // [Required] <p>Item ID.</p>
	CurrentPreOrderStatus int64 `json:"current_pre_order_status"` // [Required] <p>Current Pre-order Status. Applicable values:&nbsp;</p><p>1: Yes</p><p>2: No<br /></p>
}

type PreOrderListingViolationData struct {
	Date                 string `json:"date"`                    // [Required] <p>Date.</p>
	LiveListingCount     int64  `json:"live_listing_count"`      // [Required] <p>Number of Live Listings.<br /></p>
	PreOrderListingCount int64  `json:"pre_order_listing_count"` // [Required] <p>Number of pre-order Listings.<br /></p>
	PreOrderListingRate  int64  `json:"pre_order_listing_rate"`  // [Required] <p>Pre-order Listing %.<br /></p>
	Target               string `json:"target"`                  // [Required] <p>Target.<br /></p>
}

type Price struct {
	ModelId       int64   `json:"model_id"`       // [Required] ID of model.
	OriginalPrice float64 `json:"original_price"` // [Required] Original price for model.
}

type PriceInfo struct {
	Currency                     string  `json:"currency"`                         // [Required] The three-digit code representing the currency unit used for the item in Shopee Listings.
	OriginalPrice                float64 `json:"original_price"`                   // [Required] The original price of the item in the listing currency.
	CurrentPrice                 float64 `json:"current_price"`                    // [Required] The current price of the item in the listing currency. If product under a onging promotion, current_price will be the promotion price
	InflatedPriceOfOriginalPrice float64 `json:"inflated_price_of_original_price"` // [Required] The After-tax original price of the item in the listing currency.
	InflatedPriceOfCurrentPrice  float64 `json:"inflated_price_of_current_price"`  // [Required] The After-tax current price of the item in the listing currency.
	SipItemPrice                 float64 `json:"sip_item_price"`                   // [Required] The price of the item in sip.If item is for CNSC primary shop, this field will not be returned
	SipItemPriceSource           string  `json:"sip_item_price_source"`            // [Required]  source of sip' price. ( auto or manual).If item is for CNSC SIP primary shop, this field will not be returned
	LocalPrice                   float64 `json:"local_price"`                      // [Required] <p>The original price multiplied by the local adjustment rate equals the local price. The local price is denominated in the local currency and is rounded to two decimal places.</p>
	LocalPromotionPrice          float64 `json:"local_promotion_price"`            // [Required] <p>During the promotion period, the CB price is multiplied by the local adjustment rate. Once the promotion starts, the price remains unchanged. During the promotion, the local_promotion_price= current_price, which is denominated in the local currency and retained to two decimal places.<br /></p>
}

type PriceLimit struct {
	MinLimit float64 `json:"min_limit"` // [Required] Global item price min limit.
	MaxLimit float64 `json:"max_limit"` // [Required] Global item price max limit.
}

type ProductInfo struct {
	ItemId            int64                        `json:"item_id"`             // [Required] <p>ID of this kit item.<br /></p>
	ItemName          string                       `json:"item_name"`           // [Required] <p>The name of this kit item.<br /></p>
	CategoryId        []int64                      `json:"category_id"`         // [Required] <p>The category of this kit item, sync from the category of the main component of this kit item.</p>
	ItemStatus        ItemStatus                   `json:"item_status"`         // [Required] <p>Enumerated type that defines the current status of the item. Applicable values: NORMAL, BANNED, UNLIST,&nbsp;SELLER_DELETE, SHOPEE_DELETE, REVIEWING.<br /></p>
	ItemSku           string                       `json:"item_sku"`            // [Required] <p>An item SKU (stock keeping unit) is an identifier defined by a seller, sometimes called parent SKU. Item SKU can be assigned to an item in Shopee Listings.<br /></p>
	Images            *ItemImage                   `json:"images"`              // [Required] <p>Item images with 1:1 ratio.<br /></p>
	LongImages        *ItemImage                   `json:"long_images"`         // [Required] <p>Item images with 3:4 ratio.<br /></p>
	DescriptionInfo   *ResponseDataDescriptionInfo `json:"description_info"`    // [Required] <p>Rich text description field. Only whitelist sellers can use it.<br /></p>
	Description       string                       `json:"description"`         // [Required] <p>If description_type is normal, description information will be returned through this field, else description will be empty.<br /></p>
	DescriptionType   DescriptionType              `json:"description_type"`    // [Required] <p>Type of description : values: See Data Definition- description_type (normal , extended).<br /></p>
	VideoList         *VideoInfo                   `json:"video_list"`          // [Required] <p>Info of video list.<br /></p>
	Attributes        []Attributes                 `json:"attributes"`          // [Required] <p>The attributes of this kit item, sync from the attributes of the main component of this kit item.<br /></p>
	Weight            string                       `json:"weight"`              // [Required] <p>The weight of this kit item, the unit is KG.</p>
	Dimension         *Dimension                   `json:"dimension"`           // [Required] <p>The dimension of this kit item.</p>
	BrandInfo         *Brand                       `json:"brand_info"`          // [Required] <p>The brand of this kit item, sync from the brand of the main component of this kit item.<br /></p>
	ModelList         []ProductInfoModel           `json:"model_list"`          // [Required] <p>Model info list, model number at most 9.<br /></p>
	PreOrderInfo      *PreOrder                    `json:"pre_order_info"`      // [Required]
	TierVariationList []ProductInfoTierVariation   `json:"tier_variation_list"` // [Required] <p>Variation config of item.<br /></p>
}

type ProductInfoModel struct {
	ModelId       int64            `json:"model_id"`       // [Required] <p>ID of this kit model.<br /></p>
	ModelSku      int64            `json:"model_sku"`      // [Required] <p>Seller SKU of this kit model.<br /></p>
	OriginalPrice float64          `json:"original_price"` // [Required] <p>Original price of this kit model.<br /></p>
	TierIndex     interface{}      `json:"tier_index"`     // [Required] <p>Tier index of this kit model.<br /></p>
	ComponentList []ModelComponent `json:"component_list"` // [Required]
}

type ProductInfoTierVariation struct {
	Name       string                           `json:"name"`        // [Required] <p>Variation name.<br /></p>
	OptionList []ProductInfoTierVariationOption `json:"option_list"` // [Required] <p>Option list.<br /></p>
}

type ProductInfoTierVariationOption struct {
	Option string           `json:"option"` // [Required] <p>Option name.<br /></p>
	Image  []FieldImageInfo `json:"image"`  // [Required]
}

type Promotion struct {
	PromotionType        string                `json:"promotion_type"`          // [Required] Promotion type, Applicable values: See Data Definition- PromotionType.
	PromotionId          int64                 `json:"promotion_id"`            // [Required] The identity of item promotion.
	ModelId              int64                 `json:"model_id"`                // [Required] The identity of product model.
	StartTime            int64                 `json:"start_time"`              // [Required] Promotion start tiem.
	EndTime              int64                 `json:"end_time"`                // [Required] Promotion end item.
	PromotionPriceInfo   []PromotionPriceInfo  `json:"promotion_price_info"`    // [Required] Promotion price info.
	PromotionStaging     string                `json:"promotion_staging"`       // [Required] Could be ongoing/upcoming
	PromotionStockInfoV2 *PromotionStockInfoV2 `json:"promotion_stock_info_v2"` // [Required] <p>new promotion stock<br /></p>
}

type PromotionImages struct {
	ImageIdList []string `json:"image_id_list"` // [Required] Image id list of item.
}

type PromotionPriceInfo struct {
	PromotionPrice float64 `json:"promotion_price"` // [Required] Promotion price.
}

type PromotionStockInfoV2 struct {
	SummaryInfo        interface{} `json:"summary_info"`         // [Required] <p>stock summary info<br /></p>
	TotalReservedStock int64       `json:"total_reserved_stock"` // [Required] <p>Total Stock reserved for promotion<br /></p>
}

type PublishItemToOutletShopResponse struct {
	BaseResponse `json:",inline"` // Common response fields
}

type PublishableShop struct {
	ShopRegion string `json:"shop_region"` // [Required] Region of published shop.
}

type PublishedItem struct {
	ShopRegion string     `json:"shop_region"` // [Required] Region of shop.
	ItemId     int64      `json:"item_id"`     // [Required] Id of published item.
	ItemStatus ItemStatus `json:"item_status"` // [Required] <p>Status of published item.Applicable values: 0.DELETED(Item is deleted by seller himself),1.NORMAL, 2.BANNED,3.REVIEWING,4.INVALID(Shopee Admin deleted),5.INVALID_HIDE(Shopee Admin delete confirmed),6.BLACKLISTED(Offensive_hide),8.NORMAL_UNLIST</p>
}

type Punishment struct {
	IssueTime      int64  `json:"issue_time"`      // [Required] <p>The time when punishment are issued.<br /></p>
	StartTime      int64  `json:"start_time"`      // [Required] <p>Start time in the duration of this punishment record.</p>
	EndTime        int64  `json:"end_time"`        // [Required] <p>End time in the duration of this punishment record.<br /></p>
	PunishmentType int64  `json:"punishment_type"` // [Required] <p>Punishment Type of this punishment record. Applicable values:&nbsp;</p><p>103: Listings not displayed in category browsing<br />104: Listings not displayed in search<br />105: Unable to create new listings<br />106: Unable to edit listings<br />107: Unable to join marketing campaigns<br />108: No shipping subsidies<br />109: Account is suspended<br />600: Listings not displayed in search<br />601: Shop listings hide from recommendation<br />602: Listings not displayed in category browsing<br />1109: Listing Limit is reduced<br />1110: Listing Limit is reduced<br />1111: Listing Limit is reduced<br />1112: Listing Limit is reduced<br /></p><p>2008: Order Limit</p>
	Reason         int64  `json:"reason"`          // [Required] <p>Reason of this punishment record. Applicable values:&nbsp;</p><p>1: Tier 1</p><p>2: Tier 2</p><p>3: Tier 3</p><p>4: Tier 4</p><p>5: Tier 5</p><p>1109: Listing Limit Tier 1</p><p>1110: Listing Limit Tier 2</p><p>1111: Listing Limit POL<br /></p>
	ReferenceId    int64  `json:"reference_id"`    // [Required] <p>Reference ID for this punishment record.<br /></p>
	ListingLimit   int64  `json:"listing_limit"`   // [Required] <p>Return the specific value of listing limit when punishment_type is:&nbsp;</p><p>1109: Listing Limit is reduced<br />1110: Listing Limit is reduced<br />1111: Listing Limit is reduced<br />1112: Listing Limit is reduced</p>
	OrderLimit     string `json:"order_limit"`     // [Required] <p>Return the specific percentage of order limit when punishment_type is:&nbsp;</p><p>2008: Order Limit</p><p><br /></p><p>Daily Order Limit = X % * L28D ADO (Average Daily Order of this Shop in Past 28 Days)</p>
}

type PurchaseLimitInfo struct {
	MinPurchaseLimit int64             `json:"min_purchase_limit"` // [Required] <p>minimum purchase count for each order</p>
	MaxPurchaseLimit *MaxPurchaseLimit `json:"max_purchase_limit"` // [Required]
}

type PushMessage struct {
	Code int64  `json:"code"` // [Required] <p>Shopee's unique identifier for a push notification.<br /></p>
	Data string `json:"data"` // [Required] <p>Main Push message data.<br /></p>
}

type Queries struct {
	OrderSn string `json:"order_sn"` // [Required] Shopee's unique identifier for an order.
}

type QueryBrShopBlockStatusResponse struct {
	BaseResponse `json:",inline"`                   // Common response fields
	Response     QueryBrShopBlockStatusResponseData `json:"response"` // Actual response data
}

type QueryBrShopBlockStatusResponseData struct {
	IsBlock bool `json:"is_block"` // [Required] <p>shop blocked status</p>
}

type QueryBrShopEnrollmentStatusResponse struct {
	BaseResponse `json:",inline"`                        // Common response fields
	Response     QueryBrShopEnrollmentStatusResponseData `json:"response"` // Actual response data
}

type QueryBrShopEnrollmentStatusResponseData struct {
	EnrollmentStatus     int64 `json:"enrollment_status"`      // [Required] <p>1: enable enrollment</p><p>2: disable enrollment</p><p>3: already enrollment</p>
	EnableEnrollmentTime int64 `json:"enable_enrollment_time"` // [Required] <p>The time of this shop able to enroll FBS.</p>
}

type QueryBrShopInvoiceErrorRequest struct {
	PageNo   *int64 `json:"page_no,omitempty"`   // [Optional]
	PageSize *int64 `json:"page_size,omitempty"` // [Optional] <p>max: 100</p>
}

type QueryBrShopInvoiceErrorResponse struct {
	BaseResponse `json:",inline"`                    // Common response fields
	Response     QueryBrShopInvoiceErrorResponseData `json:"response"` // Actual response data
}

type QueryBrShopInvoiceErrorResponseData struct {
	Total int64                                     `json:"total"` // [Required]
	List  []QueryBrShopInvoiceErrorResponseDataList `json:"list"`  // [Required]
}

type QueryBrShopInvoiceErrorResponseDataList struct {
	BizRequestType      int64         `json:"biz_request_type"`      // [Required] <p>1: Inbound</p><p>2: Return From Warehouse</p><p>3:&nbsp;Sales order invoice</p><p>4: Move Transfer</p><p>5：IA</p>
	BizRequestId        string        `json:"biz_request_id"`        // [Required] <p>Return by default. The business FBS request order ID.</p>
	FailReason          string        `json:"fail_reason"`           // [Required] <p>Invoice issuance failed reason.</p>
	FailType            int64         `json:"fail_type"`             // [Required] <p>1: sku tax info error</p><p>2: seller tax info error</p>
	InvoiceDeadlineTime int64         `json:"invoice_deadline_time"` // [Required] <p>The expired time of this failed invoice. If expired, then this request order would be cancelled.</p>
	ShopSkuList         []ListShopSku `json:"shop_sku_list"`         // [Required]
	InvoiceId           string        `json:"invoice_id"`            // [Required] <p>Invoice ID</p>
	ReminderDesc        string        `json:"reminder_desc"`         // [Required] <p>remind seller if this block issue is not solved , it will block the shop or item</p>
}

type QueryBrSkuBlockStatusRequest struct {
	ShopSkuId string `json:"shop_sku_id"` // [Required]
}

type QueryBrSkuBlockStatusResponse struct {
	BaseResponse `json:",inline"`                  // Common response fields
	Response     QueryBrSkuBlockStatusResponseData `json:"response"` // Actual response data
}

type QueryBrSkuBlockStatusResponseData struct {
	ShopSkuId     string `json:"shop_sku_id"`     // [Required] <p>itemID_modelID</p>
	IsBlock       bool   `json:"is_block"`        // [Required] <p>product is blocked and warehouse stock cannot be sold</p>
	ShopItemId    int64  `json:"shop_item_id"`    // [Required] <p>ID of item</p>
	ShopModelId   int64  `json:"shop_model_id"`   // [Required] <p>ID of model</p>
	ShopItemName  string `json:"shop_item_name"`  // [Required] <p>Name of Item</p>
	ShopModelName string `json:"shop_model_name"` // [Required] <p>Name of model</p>
}

type QueryProofRequest struct {
	ReturnSn string `json:"return_sn" url:"return_sn"` // [Required] <p>The serial number of return.</p>
}

type QueryProofResponse struct {
	BaseResponse `json:",inline"`       // Common response fields
	Response     QueryProofResponseData `json:"response"` // Actual response data
}

type QueryProofResponseData struct {
	Image       []Photo `json:"image"`       // [Required]
	Video       []Photo `json:"video"`       // [Required]
	Description string  `json:"description"` // [Required] <p>The text description in the dispute proof.<br /></p>
}

type RecipientAddress struct {
	Name        string       `json:"name"`         // [Required] <p>Recipient's name for the address.<br /></p>
	Phone       string       `json:"phone"`        // [Required] <p>Recipient's phone number input when order was placed.</p><p>[Only for TW non-integrated channel] Will return "****" when the "virtual_contact_number" is available</p>
	Town        string       `json:"town"`         // [Required] <p>The town of the recipient's address. Whether there is a town will depend on the region and/or country.<br /></p>
	District    string       `json:"district"`     // [Required] <p>The district of the recipient's address. Whether there is a district will depend on the region and/or country.<br /></p>
	City        string       `json:"city"`         // [Required] <p>The city of the recipient's address. Whether there is a city will depend on the region and/or country.<br /></p>
	State       string       `json:"state"`        // [Required] <p>The state/province of the recipient's address. Whether there is a state/province will depend on the region and/or country.<br /></p>
	Region      string       `json:"region"`       // [Required] <p>The two-digit code representing the region of the Recipient.<br /></p>
	Zipcode     string       `json:"zipcode"`      // [Required] <p>Recipient's postal code.<br /></p>
	FullAddress string       `json:"full_address"` // [Required] <p>The full address of the recipient, including country, state, even street, and etc.<br /></p>
	Geolocation *Geolocation `json:"geolocation"`  // [Required] <p>Geolocation info. Only available for logistics_channel_id 90026.</p>
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

type RegisterBrandRequest struct {
	OriginalBrandName        string           `json:"original_brand_name"`                  // [Required] Brand name, length<=254.
	CategoryList             []int64          `json:"category_list"`                        // [Required] Category_id list for this brand, please input category in L1 or L2. Max input num of category_id is 50.
	ProductImage             *PromotionImages `json:"product_image"`                        // [Required]
	AppLogoImageId           *string          `json:"app_logo_image_id,omitempty"`          // [Optional] Image_id  of logo for  app client,please input hashcode of this picture.
	BrandWebsite             *string          `json:"brand_website,omitempty"`              // [Optional] Official website of brand, length<=254.
	BrandDescription         *string          `json:"brand_description,omitempty"`          // [Optional] The description of this brand, can input the information, length<=254.
	AdditionalInformation    *string          `json:"additional_information,omitempty"`     // [Optional] Additional notes or comment can seller can add, length<=254.
	PcLogoImageId            *string          `json:"pc_logo_image_id,omitempty"`           // [Optional] Image_id  of logo for  pc client,please input hashcode of this picture.
	BrandRegion              string           `json:"brand_region"`                         // [Required] <p>origin region of brand.</p>
	Licenses                 []Licenses       `json:"licenses,omitempty"`                   // [Optional] <p>For appeal blacklisted brand data</p>
	BrandRegistrationWebsite *string          `json:"brand_registration_website,omitempty"` // [Optional] <p>The link to brand registration website, It is mandatory when brand name hit blacklist.len&lt;254<br /></p>
}

type RegisterBrandResponse struct {
	BaseResponse `json:",inline"`          // Common response fields
	Response     RegisterBrandResponseData `json:"response"` // Actual response data
}

type RegisterBrandResponseData struct {
	BrandId           int64  `json:"brand_id"`            // [Required] The identity of brand.
	OriginalBrandName string `json:"original_brand_name"` // [Required] Brand name
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

type ReplyCommentRequest struct {
	CommentList []Comment `json:"comment_list"` // [Required] The list of comment. The limit is between 1 and 100.
}

type ReplyCommentResponse struct {
	BaseResponse `json:",inline"`         // Common response fields
	Response     ReplyCommentResponseData `json:"response"` // Actual response data
}

type ReplyCommentResponseData struct {
	ResultList []Result `json:"result_list"` // [Required] The result list of the request comment list.
}

type Report struct {
	BroadCir          float64 `json:"broad_cir"`           // [Required] <p>The direct advertising cost of sales, or direct ACOS, measures how much your ad costs relative to the revenue generated from sales of the advertised product. It is the amount spent on the ad divided by the amount of sales revenue for the advertised product that is attributed to the ad. Direct ACOS = expense ÷ direct GMV × 100%.<br /></p>
	BroadGmv          float64 `json:"broad_gmv"`           // [Required] <p>The amount of sales revenue generated from shoppers purchasing products within 7 days of them clicking on your ad.<br /></p>
	BroadOrder        int64   `json:"broad_order"`         // [Required] <p>The number of times shoppers purchased any product from your shop within 7 days of them clicking on your ad.</p>
	BroadOrderAmount  int64   `json:"broad_order_amount"`  // [Required] <p>The total quantity of products purchased by shoppers within 7 days of them clicking on your ad.</p>
	BroadRoi          float64 `json:"broad_roi"`           // [Required] <p>Return on ad spend (ROAS) measures how much revenue is generated by your ad relative to the cost of the ad. It is the amount of sales revenue attributed to your ad divided by the amount spent on the ad. ROAS = GMV ÷ expense. (Note: We recommend monitoring ROAS trends on a weekly basis.)<br /></p>
	Clicks            int64   `json:"clicks"`              // [Required] <p>The number of times shoppers click on your ad. (Note: Shopee filters out repeated clicks from the same shopper that occur within a short time frame.)<br /></p>
	Expense           float64 `json:"expense"`             // [Required] <p>The amount spent on your ad.</p>
	Cpc               float64 `json:"cpc"`                 // [Required] <p>The cost per conversion is how much each conversion costs, on average. It is the amount spent on your ad divided by the number of conversions attributed to the ad. Cost per conversion = expense ÷ conversions.<br /></p>
	Cpdc              float64 `json:"cpdc"`                // [Required] <p>The cost per direct conversion is how much each direct conversion costs, on average. It is the amount spent on your ad divided by the number of direct conversions attributed to the ad. Cost per direct conversion = expense ÷ direct conversions.<br /></p>
	Cr                float64 `json:"cr"`                  // [Required] <p>The conversion rate measures how often shoppers end up purchasing something from your shop after clicking on your ad. It is the number of conversions attributed to your ad divided by the number of clicks on the ad. Conversion rate = conversions ÷ clicks × 100%.<br /></p>
	DirectCr          float64 `json:"direct_cr"`           // [Required] <p>The direct conversion rate measures how often shoppers end up purchasing the advertised product after clicking on the ad. Direct conversion rate is the number of direct conversions divided by the number of clicks. Direct conversion rate = direct conversions ÷ clicks × 100%.<br /></p>
	DirectCir         float64 `json:"direct_cir"`          // [Required] <p>The direct advertising cost of sales, or direct ACOS, measures how much your ad costs relative to the revenue generated from sales of the advertised product. It is the amount spent on the ad divided by the amount of sales revenue for the advertised product that is attributed to the ad. Direct ACOS = expense ÷ direct GMV × 100%.<br /></p>
	DirectOrder       int64   `json:"direct_order"`        // [Required] <p>The number of times shoppers purchased the advertised product within 7 days of them clicking on the ad.<br /></p>
	DirectOrderAmount int64   `json:"direct_order_amount"` // [Required] <p>The total quantity of the advertised product purchased by shoppers within 7 days of them clicking on the ad.<br /></p>
	DirectRoi         float64 `json:"direct_roi"`          // [Required] <p>The direct return on ad spend, or direct ROAS, measures how much revenue is generated from sales of the advertised product, relative to the cost of the ad. It is the amount of sales revenue for the advertised product attributed to the ad, divided by the amount spent on the ad. Direct ROAS = direct GMV ÷ expense.<br /></p>
	Impression        float64 `json:"impression"`          // [Required] <p>The number of times shoppers see your ad.<br /></p>
}

type ReportData struct {
	UploadCost int64 `json:"upload_cost"` // [Required] Time used for uploading the video file via upload_video_part api, in milliseconds. For video upload performance tracking purpose.
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

type RequestBrand struct {
	BrandId *int64 `json:"brand_id,omitempty"` // [Optional] Id of brand.
}

type RequestCertificationInfo struct {
	CertificationList []CertificationInfoCertification `json:"certification_list,omitempty"` // [Optional] <p>Array of certification records for the product, each containing type, certificate number, permit ID, and proof documents.</p><p>  </p><p><br /></p>
}

type RequestComponent struct {
	ComponentItemId  int64  `json:"component_item_id"`            // [Required] <p>ID of the item that composes this kit model.</p>
	ComponentModelId *int64 `json:"component_model_id,omitempty"` // [Optional] <p>ID of the model that composes this kit model.</p>
}

type RequestDimension struct {
	Length int64 `json:"length"` // [Required] <p>The length of the packaging in centimetres (cm).</p>
	Width  int64 `json:"width"`  // [Required] <p>The width of the packaging in centimetres (cm).</p>
	Height int64 `json:"height"` // [Required] <p>The height of the packaging in centimetres (cm).</p>
}

type RequestDiscoveryAdsLocations struct {
	Location string  `json:"location"`  // [Required] <p>daily_discover, you_may_also_like<br /></p>
	Status   string  `json:"status"`    // [Required] <p>toggle on or toggle off</p>
	BidPrice float64 `json:"bid_price"` // [Required] <p>bid price&nbsp;</p>
}

type RequestDropoff struct {
	BranchId       *int64  `json:"branch_id,omitempty"`        // [Optional] The identity of branch. Retrieved from shopee.logistics.GetBranch branch.
	SenderRealName *string `json:"sender_real_name,omitempty"` // [Optional] The real name of sender.
	TrackingNumber *string `json:"tracking_number,omitempty"`  // [Optional] Need input this field when "tracking_number" is returned from "info_need". Please note that this tracking number is assigned by third-party shipping carrier for item shipment.
}

type RequestGlobalModel struct {
	GlobalModelSku string           `json:"global_model_sku"`    // [Required] Sku of global model.
	GlobalModelId  int64            `json:"global_model_id"`     // [Required] ID of global model.
	Weight         *float64         `json:"weight,omitempty"`    // [Optional] <p>The weight of this global model, the unit is KG.</p><p>If don't set the weight of this global model, will use the weight of global item by default.</p><p>If set the dimension of this global model, them must set the weight of this global model.</p>
	Dimension      *Dimension       `json:"dimension,omitempty"` // [Optional] <p>The dimension of this global model.</p><p>If don't set the dimension of this global model, will use the dimension of global item by default.</p>
	PreOrder       *RequestPreOrder `json:"pre_order,omitempty"` // [Optional] <p>Pre-order information of this global model.</p><p><br /></p><p>Notes:&nbsp;</p><p>If don't set the DTS of this global model, will use the DTS of the global item by default.</p>
}

type RequestIdList struct {
	RequestId []int64 `json:"request_id"` // [Required]
}

type RequestImage struct {
	ImageIdList []interface{} `json:"image_id_list"`         // [Required] Image ID.
	ImageRatio  *string       `json:"image_ratio,omitempty"` // [Optional] <p>Ratio of image, <br />OptionalAllowed ratios :<br />"1:1" (default)&nbsp;<br />"3:4"<br /></p>
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

type RequestItemModel struct {
	ModelId             int64   `json:"model_id"`                        // [Required] Shopee's unique identifier for a variation of an item. If there is no variation of this item, you don't need to input this param. Dafault is 0.
	ModelPromotionPrice float64 `json:"model_promotion_price"`           // [Required] The discount price of the item.
	ModelPromotionStock *int64  `json:"model_promotion_stock,omitempty"` // [Optional] <p>The reserved stock of the model, default is no limit, and can not update. To edit the promotion stock, you need to delete the exist discount and re-add again.</p>
}

type RequestItemSetting struct {
	ItemName          *string                   `json:"item_name,omitempty"`           // [Optional] <p>The name of this kit item.<br /></p>
	Images            *PromotionImages          `json:"images,omitempty"`              // [Optional] <p>Item images with 1:1 ratio.<br /></p>
	LongImages        *PromotionImages          `json:"long_images,omitempty"`         // [Optional] <p>Item images with 3:4 ratio.<br /></p>
	VideoUploadId     []string                  `json:"video_upload_id,omitempty"`     // [Optional] <p>Video upload ID returned from video uploading API. Only accept one video_upload_id.<br /></p>
	Description       *string                   `json:"description,omitempty"`         // [Optional] <p>If description_type is normal, description information should be set by this field.<br /></p>
	DescriptionInfo   *DescriptionInfo          `json:"description_info,omitempty"`    // [Optional] <p>Rich text description field. Only whitelist sellers can use it. If you use the field, please upload the description_type=extended otherwise api will return error. If you don't use this field, you don't need to upload the description_type or upload description_type=normal<br /></p>
	DescriptionType   *DescriptionType          `json:"description_type,omitempty"`    // [Optional] <p>See Data Definition- description_type (normal , extended). If you want to use extended_description, this field must be inputed.<br /></p>
	LogisticInfo      []LogisticInfo            `json:"logistic_info,omitempty"`       // [Optional] <p>Logistic channel setting.<br /></p>
	Unlisted          *bool                     `json:"unlisted,omitempty"`            // [Optional] <p>Unlist or not.<br /></p>
	ItemSku           *string                   `json:"item_sku,omitempty"`            // [Optional] <p>SKU tag of item<br /></p>
	Weight            *float64                  `json:"weight,omitempty"`              // [Optional] <p>The weight of this kit item, the unit is KG.<br /></p>
	Dimension         *Dimension                `json:"dimension,omitempty"`           // [Optional] <p>The dimension of this kit item.<br /></p>
	PreOrder          *PreOrder                 `json:"pre_order,omitempty"`           // [Optional] <p>Pre order setting.<br /></p>
	ModelList         []RequestItemSettingModel `json:"model_list,omitempty"`          // [Optional] <p>Model info list, model number at most 9.<br /></p>
	TierVariationList []TierVariation           `json:"tier_variation_list,omitempty"` // [Optional] <p>Tier variation info list.&nbsp;</p><p>Only support one tier variation, and each kit item can have from 1 to 9 kit variations.</p>
}

type RequestItemSettingModel struct {
	ModelId       *int64      `json:"model_id,omitempty"`       // [Optional] <p>ID of this kit model.<br /></p>
	TierIndex     []int64     `json:"tier_index"`               // [Required] <p>Tier index of this kit model.<br /></p>
	ModelSku      *string     `json:"model_sku,omitempty"`      // [Optional] <p>Seller SKU of this kit model, model_sku length information needs to be no more than 100 characters.<br /></p>
	OriginalPrice *float64    `json:"original_price,omitempty"` // [Optional] <p>Original price of this kit model.<br /></p>
	ComponentList []Component `json:"component_list,omitempty"` // [Optional] <p>Please get the amount of item/model that one kit model support from get_kit_item_limit.<br /></p>
}

type RequestItems struct {
	ItemId              int64    `json:"item_id"`                          // [Required]
	PurchaseLimit       int64    `json:"purchase_limit"`                   // [Required] <p>min=0, 0 means no limit</p>
	Models              []Models `json:"models,omitempty"`                 // [Optional] <p>If the item has variation, this param is necessary.</p>
	ItemInputPromoPrice *float64 `json:"item_input_promo_price,omitempty"` // [Optional] <p>promotion price without tax of the item. If the item has no variation, this param is necessary, otherwise don't use this field<br /></p>
	ItemStock           *int64   `json:"item_stock,omitempty"`             // [Optional] <p>min=1, The campaign stock of the item. If the item has no variation, this param is necessary, otherwise don't use this field</p>
}

type RequestModel struct {
	ModelId   int64   `json:"model_id"`   // [Required] <p>ID of model</p>
	TierIndex []int64 `json:"tier_index"` // [Required] <p>Model's tier_variation</p>
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

type RequestPreOrder struct {
	DaysToShip int64 `json:"days_to_ship"` // [Required] <p>Days to ship.</p>
}

type RequestPrice struct {
	GlobalModelId *int64  `json:"global_model_id,omitempty"` // [Optional] ID of global model.
	OriginalPrice float64 `json:"original_price"`            // [Required] Original price of global item.
}

type RequestSelectedKeywords struct {
	EditAction       string   `json:"edit_action"`                   // [Required] <p>The update behaviours such as "add", "delete", "restore", "change_bid_price", "change_match_type"<br /></p>
	Keyword          string   `json:"keyword"`                       // [Required] <p>bid keyword for each campaign</p>
	MatchType        *string  `json:"match_type,omitempty"`          // [Optional] <p>exact, broad; required if changing match type<br /></p>
	BidPricePerClick *float64 `json:"bid_price_per_click,omitempty"` // [Optional] <p>the bid price of keyword; required if changing bid price</p>
}

type RequestStock struct {
	GlobalModelId *int64        `json:"global_model_id,omitempty"` // [Optional] ID of global model.
	SellerStock   []SellerStock `json:"seller_stock,omitempty"`    // [Optional]
}

type RequestTaxInfo struct {
	Ncm               *string        `json:"ncm,omitempty"`                 // [Optional] <p>Mercosur Common Nomenclature, it is a convention between Mercosur member countries to easily recognize goods, services and productive factors negotiated among themselves.&nbsp;(BR region)<br /></p><p><br />NCM must have 8 digits, OR, if your item doesn't have a NCM enter the value "00"</p>
	SameStateCfop     *string        `json:"same_state_cfop,omitempty"`     // [Optional] Tax Code of Operations and Installments for orders that seller and buyer are in the same state. It identifies a specific operation by category at the time of issuing the invoice.
	DiffStateCfop     *string        `json:"diff_state_cfop,omitempty"`     // [Optional] Tax Code of Operations and Installments for orders that seller and buyer are in different states. It identifies a specific operation by category at the time of issuing the invoice.
	Csosn             *string        `json:"csosn,omitempty"`               // [Optional] Code of Operation Status – Simples Nacional, code for company operations to identify the origin of the goods and the taxation regime of the operations.
	Origin            *string        `json:"origin,omitempty"`              // [Optional] Product source, domestic or foreig
	Cest              *string        `json:"cest,omitempty"`                // [Optional] <p>Tax Replacement Specifying Code (CEST), to separate within the same NCM products that do or do not have ICMS tax substitution. (BR region)<br /><br />CEST must have 7 digits, OR, if your item doesn't have a CEST enter the value "00".<br /></p>
	MeasureUnit       *string        `json:"measure_unit,omitempty"`        // [Optional] (BR region)
	InvoiceOption     *InvoiceOption `json:"invoice_option,omitempty"`      // [Optional] Value shuold be one of NO_INVOICES VAT_MARGIN_SCHEME_INVOICES VAT_INVOICES NON_VAT_INVOICES and if value is NON_VAT_INVOICE vat_rate should be null (PL region)
	VatRate           *string        `json:"vat_rate,omitempty"`            // [Optional] Value should be one of 0% 5% 8% 23% NO_VAT_RATE (PL region)
	HsCode            *string        `json:"hs_code,omitempty"`             // [Optional] HS Code. (Only for IN region)
	TaxCode           *string        `json:"tax_code,omitempty"`            // [Optional] Tax Code. (Only for IN region)
	TaxType           *TaxType       `json:"tax_type,omitempty"`            // [Optional] <p>tax_type only for TW whitelist shop. Shopee will referred Tax type when substitute sellers for issuing e-receipts to buyers. All variations share the same tax type. The meaning of value:&nbsp;</p><p>0: no tax type</p><p>1: tax-able</p><p>2: tax-free</p>
	Pis               *string        `json:"pis,omitempty"`                 // [Optional] <p>Only for BR shop.</p><p>PIS - Programa de Integração Social (Social Integration Program). It is a government tax to collect resources for the payment of unemployment insurance and other employee related rights.</p><p>PIS % - the tax applied to this product</p>
	Cofins            *string        `json:"cofins,omitempty"`              // [Optional] <p>Only for BR shop.<br /></p><p>COFINS – Contribuição para Financiamento da Seguridade Social (Contribution for Social Security Funding). It&nbsp;is a government tax to collect resources for public health system and social security.</p><p>COFINS&nbsp;% - the tax applied to this product</p>
	IcmsCst           *string        `json:"icms_cst,omitempty"`            // [Optional] <p>Only for BR shop.<br /></p><p>ICMS - Imposto sobre Circulação de Mercadorias e Serviços (Circulation of Goods and Services Tax).&nbsp;</p><p>CST - Código da Situação Tributária (Tax Situation Code) is represented by a combination of 3 numbers with the purpose of demonstrating the origin of a product and determining the form of taxation that will apply to it. Therefore, each digit in the CST Table has a specific meaning: the first digit indicates the origin of the operation, the second digit represents the ICMS taxation on the operation and the third digit provides additional information about the form of taxation.</p>
	PisCofinsCst      *string        `json:"pis_cofins_cst,omitempty"`      // [Optional] <p>Only for BR shop.</p><p>The CST PIS/Cofins is a code on the Electronic Invoice (NF-e) that identifies the tax situation of PIS (Programa de Integração Social) and Cofins (Contribuição para o Financiamento da Seguridade Social) in sales of goods.</p>
	FederalStateTaxes *string        `json:"federal_state_taxes,omitempty"` // [Optional] <p>Only for BR shop.</p><p>Enter the total percentage of the combination of federal, state, and municipal taxes, using up to two decimals.</p>
	OperationType     *OperationType `json:"operation_type,omitempty"`      // [Optional] <p>Only for BR shop.</p><p>1: Retailer</p><p>2: Manufacturer</p>
	ExTipi            *string        `json:"ex_tipi,omitempty"`             // [Optional] <p>Only for BR shop.<br /></p><p>The EXTIPI field in the NF-e (Nota Fiscal Eletrônica) is used to indicate if there's an exception to the IPI (Imposto sobre Produtos Industrializados) tax rate for a specific product.</p>
	FciNum            *string        `json:"fci_num,omitempty"`             // [Optional] <p>Only for BR shop.<br /></p><p>The FCI Control Number is a unique identifier assigned to each import FCI (Import Content Form). It's mandatory on the corresponding NF-e (electronic invoice) to ensure compliance with Brazilian import tax regulations.</p>
	RecopiNum         *string        `json:"recopi_num,omitempty"`          // [Optional] <p>Only for BR shop.</p><p>RECOPI NACIONAL is a Brazilian government system that facilitates the registration and management of tax-exempt operations involving paper destined for printing books, newspapers, and periodicals (known as "papel imune" in Portuguese).</p>
	AdditionalInfo    *string        `json:"additional_info,omitempty"`     // [Optional] <p>Only for BR shop.</p><p>Include relevant information to display on Invoice.</p>
	GroupItemInfo     *GroupItemInfo `json:"group_item_info,omitempty"`     // [Optional] <p>Only for BR shop.</p><p>Required if the item is a group item.</p>
	ExportCfop        *string        `json:"export_cfop,omitempty"`         // [Optional] <p><br />7101 - for sales of self-produced goods</p><p>7102 - resale of third-party goods</p>
}

type ResponseCertificationInfo struct {
	CertificationList []ResponseCertificationInfoCertification `json:"certification_list"` // [Required] <p>Array of certification records for the product, each containing type, certificate number, permit ID, and proof documents.</p><p><br /></p><p><br /></p>
}

type ResponseCertificationInfoCertification struct {
	PermitId            int64                                               `json:"permit_id"`            // [Required] <p>Permit ID, get from&nbsp;v2.product.get_product_certification_rule<br /></p>
	CertificationNo     string                                              `json:"certification_no"`     // [Required] <p>Certification No.</p>
	ExpiryDate          int64                                               `json:"expiry_date"`          // [Required] <p>expiry timestamp</p>
	CertificationProofs []CertificationInfoCertificationCertificationProofs `json:"certification_proofs"` // [Required] <p>An array of proof documents for the certification; each element represents one proof file.</p>
}

type ResponseData struct {
	CreateTime int64  `json:"create_time"` // [Required] <p>the notification create time</p>
	Content    string `json:"content"`     // [Required] <p>The content of the notification</p>
	Title      string `json:"title"`       // [Required] <p>The content of the notification<br /></p>
	Url        string `json:"url"`         // [Required] <p>Some notification may be attached with URL, it will redirect to seller center</p>
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

type ResponseDataAttribute struct {
	AttributeId        int64                     `json:"attribute_id"`         // [Required] ID of attribute.
	AttributeValueList []AttributeAttributeValue `json:"attribute_value_list"` // [Required] Value list of this attribute.
}

type ResponseDataBooking struct {
	BookingSn     string        `json:"booking_sn"`     // [Required] <p>Shopee's unique identifier for a booking.<br /></p>
	OrderSn       string        `json:"order_sn"`       // [Required] <p>Shopee's unique identifier for an order. Only return if booking_status is MATCHED.<br /></p>
	BookingStatus BookingStatus `json:"booking_status"` // [Required] <p>The booking_status filter for retrieving booking and each one only every request. Available value: READY_TO_SHIP/PROCESSED/SHIPPED/CANCELLED/MATCHED<br /></p>
	NextCursor    string        `json:"next_cursor"`    // [Required] <p>If more is true, you should pass the next_cursor in the next request as cursor. The value of next_cursor will be empty string when more is false.<br /></p>
}

type ResponseDataBrand struct {
	BrandId           int64  `json:"brand_id"`            // [Required]  Id of brand.
	OriginalBrandName string `json:"original_brand_name"` // [Required] Original name of brand
	DisplayBrandName  string `json:"display_brand_name"`  // [Required]  Display name of brand
}

type ResponseDataBundleDealRule struct {
	RuleType           int64            `json:"rule_type"`           // [Required] The bundle deal rule type：FIX_PRICE = 1 ；DISCOUNT_PERCENTAGE = 2； DISCOUNT_VALUE = 3
	DiscountValue      float64          `json:"discount_value"`      // [Required] The deducted price when when buying a bundle deal.Need to input it when the bundle deal rule type is 3
	FixPrice           float64          `json:"fix_price"`           // [Required] The amount of the buyer needs to spend to purchase a bundle deal. Need to input it when the bundle deal rule type is 1
	DiscountPercentage int64            `json:"discount_percentage"` // [Required] The discount that the buyer can get when buying a bundle deal. Need to input it when the bundle deal rule type is 2
	MinAmount          int64            `json:"min_amount"`          // [Required] The quantity of items that need buyer to combine purchased
	AdditionalTiers    *AdditionalTiers `json:"additional_tiers"`    // [Required] <p>Use to create tiered discount for bundle deals, a max of 2 additional tiers are allowed to create bundle deals like buy 2 get 10% off, buy 3 for 15% off, buy 4 for 20% off;&nbsp;For each tier, we will need to set the following 4 values based on bundle deal type&nbsp;+&nbsp; &nbsp; min_amount = IntAttribute()&nbsp;+&nbsp;&nbsp;&nbsp; fix_price = IntAttribute()&nbsp;+&nbsp;&nbsp;&nbsp; discount_percentage = IntAttribute()&nbsp;+&nbsp;&nbsp;&nbsp; discount_value = IntAttribute()&nbsp;&nbsp;Note: for additional tiers, the fix price, discount_percentage, discount_value should be consistent with tier 1<br /></p>
}

type ResponseDataCampaign struct {
	CampaignId        int64             `json:"campaign_id"`        // [Required] <p>The unique identifier for a campaign<br /></p>
	AdType            string            `json:"ad_type"`            // [Required] <p>auto, manual<br /></p>
	CampaignPlacement string            `json:"campaign_placement"` // [Required] <p>search, discovery, all<br /></p>
	AdName            string            `json:"ad_name"`            // [Required] <p>the name of each campaign</p>
	MetricsList       []CampaignMetrics `json:"metrics_list"`       // [Required] <p>performance metric list</p>
}

type ResponseDataComplaintPolicy struct {
	WarrantyTime                WarrantyTime `json:"warranty_time"`                 // [Required] Value should be in one of ONE_YEAR TWO_YEARS OVER_TWO_YEARS.
	ExcludeEntrepreneurWarranty bool         `json:"exclude_entrepreneur_warranty"` // [Required] If True means "I exclude warranty complaints for entrepreneur"
	AdditionalInformation       string       `json:"additional_information"`        // [Required] Additional information for complaint policy
}

type ResponseDataDescriptionInfo struct {
	ExtendedDescription *DescriptionInfoExtendedDescription `json:"extended_description"` // [Required] If description_type is extended , Description information will be returned through this field.
}

type ResponseDataDiscount struct {
	Region          string `json:"region"`            // [Required] <p>The region of SIP affiliate shop.</p>
	Status          string `json:"status"`            // [Required] <p>The status of discount for SIP affiliate shop in current region, can be upcoming/ongoing, excluding expired discounts.</p>
	SipDiscountRate int64  `json:"sip_discount_rate"` // [Required] <p>The discount rate set for SIP affiliate shop in current region.</p>
	StartTime       int64  `json:"start_time"`        // [Required] <p>The start time of discount for SIP affiliate shop in current region,&nbsp;in UNIX seconds.</p>
	EndTime         int64  `json:"end_time"`          // [Required] <p>The end time of discount for SIP affiliate shop in current region,&nbsp;in UNIX seconds.</p>
	CreateTime      int64  `json:"create_time"`       // [Required] <p>The create time of discount for SIP affiliate shop in current region,&nbsp;in UNIX seconds.</p>
	UpdateTime      int64  `json:"update_time"`       // [Required] <p>The latest update time of discount for SIP affiliate shop in current region,&nbsp;in UNIX seconds.</p>
}

type ResponseDataDropoff struct {
	BranchList []Branch `json:"branch_list"` // [Required] List of available dropoff branches info.
	SlugList   []Slug   `json:"slug_list"`   // [Required]  List of available TW 3PL drop-off partners.
}

type ResponseDataDtsLimit struct {
	NonPreOrderDaysToShip int64                              `json:"non_pre_order_days_to_ship"` // [Required] <p>Days to ship for non pre-order products.<br /></p>
	SupportPreOrder       bool                               `json:"support_pre_order"`          // [Required] <p>Whether support pre_order for the category.<br /></p>
	DaysToShipLimit       *WholesalePriceThresholdPercentage `json:"days_to_ship_limit"`         // [Required] <p>Days to ship for pre-order products.<br /></p>
}

type ResponseDataEscrow struct {
	OrderSn           string  `json:"order_sn"`            // [Required] <p>Shopee's unique identifier for an order.</p>
	PayoutAmount      float64 `json:"payout_amount"`       // [Required] The settlement amount
	EscrowReleaseTime int64   `json:"escrow_release_time"` // [Required] The release time
}

type ResponseDataFail struct {
	Id          string `json:"id"`           // [Required] <p>Package Number or Unpackaged SKU ID that failed in generating Shipping Document</p>
	FailError   string `json:"fail_error"`   // [Required] <p>Indicate error type if one element hit error.</p>
	FailMessage string `json:"fail_message"` // [Required] <p>Indicate error details if one element hit error.</p>
}

type ResponseDataFailed struct {
	ItemId      int64  `json:"item_id"`      // [Required] The invalid item id.
	FailError   string `json:"fail_error"`   // [Required] The reason of the fail.
	FailMessage string `json:"fail_message"` // [Required] The detailed reason of the failure and the hints of error fixing
}

type ResponseDataFailure struct {
	ModelId      int64  `json:"model_id"`      // [Required] ID of model.
	FailedReason string `json:"failed_reason"` // [Required] Reason for failure.
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

type ResponseDataImageInfo struct {
	ImageId      string         `json:"image_id"`       // [Required] <p>Id of image</p>
	ImageUrlList []ThumbnailUrl `json:"image_url_list"` // [Required] <p>Image URL of each region<br /></p>
}

type ResponseDataInfoNeeded struct {
	Dropoff []string `json:"dropoff"` // [Required] <p>Could contain 'branch_id', 'sender_real_name' or 'tracking_no'. If it contains 'branch_id', choose one to Init. If it contains 'sender_real_name' or 'tracking_no', should manually input these values in Init API. If it has empty value, developer should still include "dropoff" field in Init API.<br /></p>
	Pickup  []string `json:"pickup"`  // [Required] <p>Could contain 'address_id' and 'pickup_time_id'. Choose one address_id and its corresponding pickup_time_id to Init. If it has empty value, developer should still include "pickup" field in Init API.It could contains "tracking_number" returned from "info_need"for some channels, please also add it when init.<br /></p>
}

type ResponseDataItem struct {
	ItemId     int64      `json:"item_id"`     // [Required] Shopee's unique identifier for an item.
	ItemStatus ItemStatus `json:"item_status"` // [Required] <p>Enumerated type that defines the current status of the item. Applicable values: NORMAL,&nbsp;BANNED, UNLIST, <b><font color="#c24f4a">REVIEWING, SELLER_DELETE, SHOPEE_DELETE</font></b>.<br /></p>
	UpdateTime int64      `json:"update_time"` // [Required] The update time of item.
	Tag        *Tag       `json:"tag"`         // [Required]
}

type ResponseDataItemModel struct {
	ModelId                            int64   `json:"model_id"`                                // [Required] Shopee's unique identifier for a variation of an item.
	ModelName                          string  `json:"model_name"`                              // [Required] Name of the variation that belongs to the same item.
	ModelNormalStock                   int64   `json:"model_normal_stock"`                      // [Required] The current stock quantity of the variation.
	ModelPromotionStock                int64   `json:"model_promotion_stock"`                   // [Required] The reserved stock of the model.
	ModelOriginalPrice                 float64 `json:"model_original_price"`                    // [Required] The original price before discount of the variation.
	ModelPromotionPrice                float64 `json:"model_promotion_price"`                   // [Required] The discount price of the variation.
	ModelInflatedPriceOfOriginalPrice  float64 `json:"model_inflated_price_of_original_price"`  // [Required] <p>The original price after tax of model (Only for taxable Shop).</p>
	ModelInflatedPriceOfPromotionPrice float64 `json:"model_inflated_price_of_promotion_price"` // [Required] <p>The discount price after tax of model (Only for taxable Shop).</p>
	ModelLocalPrice                    float64 `json:"model_local_price"`                       // [Required] <p>The local price of model calculated as:&nbsp;Local Price = CB Original Price × Local Adjustment Rate.<br />Reflects the final local price derived from shop-level adjustment rules&nbsp;and is denominated in local currency.</p>
	ModelLocalPromotionPrice           float64 `json:"model_local_promotion_price"`             // [Required] <p>The local discount price of model calculated as: Local Discount Price = Local Price × Discount Rate.<br />Reflects the final local seller discount price derived from setting a seller discount&nbsp;and is denominated in local currency.</p>
	ModelLocalPriceInflated            float64 `json:"model_local_price_inflated"`              // [Required] <p>The local price after tax of model (Only for taxable Shop).</p>
	ModelLocalPromotionPriceInflated   float64 `json:"model_local_promotion_price_inflated"`    // [Required] <p>The local discount price after tax of model (Only for taxable Shop).</p>
}

type ResponseDataItemPlanAhora struct {
	ItemId               int64 `json:"item_id"`                // [Required] Only applicable for local AR sellers.
	ParticipatePlanAhora bool  `json:"participate_plan_ahora"` // [Required] Only applicable for local AR sellers.
}

type ResponseDataItemSku struct {
	MtskuId            string       `json:"mtsku_id"`             // [Required] <p>mtsku id<br /></p>
	ModelId            string       `json:"model_id"`             // [Required] <p>Warehouse model SKU ID</p><p>For CB global items, this is equal to the global model_id.</p><p>For local items, it differs from model_id; use shop_model_id to match the model_id</p>
	FulfillMappingMode int64        `json:"fulfill_mapping_mode"` // [Required] <p>0-Null；1-Bundle SKU；2-Parent SKU<br /></p>
	ModelName          string       `json:"model_name"`           // [Required] <p>model name<br /></p>
	Barcode            string       `json:"barcode"`              // [Required]
	WhsList            []ItemSkuWhs `json:"whs_list"`             // [Required] <p>Info of whs<br /></p>
	ShopSkuList        []ShopSku    `json:"shop_sku_list"`        // [Required]
}

type ResponseDataItemSkuWhs struct {
	WhsId            string `json:"whs_id"`              // [Required] <p>Whs id<br /></p>
	StartOnHandTotal int64  `json:"start_on_hand_total"` // [Required] <p>Total warehouse inventory at the start time<br /></p>
	InboundTotal     int64  `json:"inbound_total"`       // [Required] <p>Inbound quantity to the warehouse during the selected time period.<br /></p>
	OutboundTotal    int64  `json:"outbound_total"`      // [Required] <p>Outbound quantity from the warehouse during the selected time period<br /></p>
	AdjustTotal      int64  `json:"adjust_total"`        // [Required] <p>Inventory adjustment quantity in the warehouse during the selected time period<br /></p>
	EndOnHandTotal   int64  `json:"end_on_hand_total"`   // [Required] <p>Total warehouse inventory at the end time.<br /></p>
}

type ResponseDataList struct {
	DirectItemId int64 `json:"direct_item_id"` // [Required] <p>Item id of direct shop.</p>
	MainShopId   int64 `json:"main_shop_id"`   // [Required] <p>Id of main shop.</p>
	MainItemId   int64 `json:"main_item_id"`   // [Required] <p>Item id of main shop.</p>
}

type ResponseDataLogisticInfo struct {
	EstimatedShippingFee float64 `json:"estimated_shipping_fee"` // [Required] Estimated shipping fee.
	LogisticName         string  `json:"logistic_name"`          // [Required] Name of logistics channel.
	Enabled              bool    `json:"enabled"`                // [Required] Whether this channel is enabled.
	LogisticId           int64   `json:"logistic_id"`            // [Required] ID of this channel.
	IsFree               bool    `json:"is_free"`                // [Required] Whether cover shipping fee for buyer.
}

type ResponseDataLogisticsChannel struct {
	LogisticsChannelId   int64  `json:"logistics_channel_id"`   // [Required] The identity of logistic channel.
	LogisticsChannelName string `json:"logistics_channel_name"` // [Required] The name of logistic channel.
	ShipmentMethod       string `json:"shipment_method"`        // [Required] <p>The shipment method for bound orders.Available values: pickup, dropoff, self_deliver.</p>
}

type ResponseDataModel struct {
	TierIndex   []interface{}    `json:"tier_index"`   // [Required] Tier index of model. Index starts from 0.
	ModelId     int64            `json:"model_id"`     // [Required] ID of model
	ModelSku    string           `json:"model_sku"`    // [Required] Seller SKU of this model
	PriceInfo   []ModelPriceInfo `json:"price_info"`   // [Required]
	SellerStock []SellerStock    `json:"seller_stock"` // [Required] <p>new stock info<br /></p>
	Weight      float64          `json:"weight"`       // [Required] <p>The weight of this model, the unit is KG.</p><p>If don't set the weight of this model, will use the weight of item by default.</p><p>If set the dimension of this model, them must set the weight of this model.</p>
	Dimension   *Dimension       `json:"dimension"`    // [Required] <p>The dimension of this model.</p><p>If don't set the dimension of this model, will use the dimension of item by default.</p>
}

type ResponseDataModelPriceInfo struct {
	Currency                     string  `json:"currency"`                         // [Required] <p>Currency for the item price.</p>
	CurrentPrice                 float64 `json:"current_price"`                    // [Required] Current price of item.
	OriginalPrice                float64 `json:"original_price"`                   // [Required] Original price of item.
	InflatedPriceOfOriginalPrice float64 `json:"inflated_price_of_original_price"` // [Required] Original price of item after tax.
	InflatedPriceOfCurrentPrice  float64 `json:"inflated_price_of_current_price"`  // [Required] Current price of item after tax.
	SipItemPrice                 float64 `json:"sip_item_price"`                   // [Required] <p>SIP item price. If item is from SIP primary shop, this field will be returned.</p>
	SipItemPriceSource           string  `json:"sip_item_price_source"`            // [Required] <p>SIP item price source, could be manual or auto.If item is from SIP primary shop, this field will be returned.</p>
	SipItemPriceCurrency         string  `json:"sip_item_price_currency"`          // [Required] <p>The currency of sip_item_price.If item is from SIP primary shop, this field will be returned.<br /></p>
	LocalPrice                   float64 `json:"local_price"`                      // [Required] <p>The original price multiplied by the local adjustment rate equals the local price. The local price is denominated in the local currency and is rounded to two decimal places.</p><p>&lt;path&gt;&lt;/path&gt;<br /></p>
	LocalPromotionPrice          float64 `json:"local_promotion_price"`            // [Required] <p>During the promotion period, the CB price is multiplied by the local adjustment rate. Once the promotion starts, the price remains unchanged. During the promotion, the local_promotion_price= current_price, which is denominated in the local currency and retained to two decimal places.<br /></p><p>&lt;path&gt;&lt;/path&gt;<br /></p>
}

type ResponseDataModels struct {
	ItemId                int64                  `json:"item_id"`                  // [Required]
	ModelId               int64                  `json:"model_id"`                 // [Required]
	ModelName             string                 `json:"model_name"`               // [Required]
	Status                int64                  `json:"status"`                   // [Required] <p>the status of model in shop flash sale</p><p>0: disable<br /></p><p>1: enable</p><p>2: delete</p><p>4: system_rejected,&nbsp;the model is rejected by system</p><p>5: manual_rejected, the model is rejected manually</p>
	OriginalPrice         float64                `json:"original_price"`           // [Required]
	InputPromotionPrice   float64                `json:"input_promotion_price"`    // [Required] <p>promotion price without tax</p>
	PromotionPriceWithTax float64                `json:"promotion_price_with_tax"` // [Required] <p>promotion price with tax<br /></p>
	PurchaseLimit         int64                  `json:"purchase_limit"`           // [Required] <p>0 means NO LIMIT</p>
	CampaignStock         int64                  `json:"campaign_stock"`           // [Required]
	Stock                 int64                  `json:"stock"`                    // [Required] <p>Active inventory<br /></p>
	RejectReason          string                 `json:"reject_reason"`            // [Required] <p>if the status is 4 or 5, this field will show the reason why this model was rejected</p>
	UnqualifiedConditions *UnqualifiedConditions `json:"unqualified_conditions"`   // [Required] <p>if the model doesn't meet a criteria, will show the detail in this field<br /></p>
}

type ResponseDataOrder struct {
	OrderSn     string      `json:"order_sn"`     // [Required]  Shopee's unique identifier for an order.
	OrderStatus OrderStatus `json:"order_status"` // [Required] The order_status filter for retriveing orders and each one only every request. Available value: UNPAID/READY_TO_SHIP/PROCESSED/SHIPPED/COMPLETED/IN_CANCEL/CANCELLED
	BookingSn   string      `json:"booking_sn"`   // [Required] <p>Return by default. Shopee's unique identifier for a booking.</p><p>Only returned for advance fulfilment matched order only.<br /></p>
}

type ResponseDataPackage struct {
	OrderSn                      string                    `json:"order_sn"`                        // [Required] <p>Shopee's unique identifier for an order.</p>
	PackageNumber                string                    `json:"package_number"`                  // [Required] <p>Shopee's unique identifier for the package under an order.</p>
	FulfillmentStatus            string                    `json:"fulfillment_status"`              // [Required] <p>The Shopee fulfillment status for the package. Applicable values: See V2.0 Data Definition - PackageFulfillmentStatus.</p>
	UpdateTime                   int64                     `json:"update_time"`                     // [Required] <p>Timestamp that indicates the last time that there was a change in value of package.</p>
	LogisticsChannelId           int64                     `json:"logistics_channel_id"`            // [Required] <p>The identity of logistic channel.</p>
	ShippingCarrier              string                    `json:"shipping_carrier"`                // [Required] <p>The logistics service provider that the buyer selected for the package to deliver items.&nbsp;</p><p><br /></p><p>Note: If logistics_channel_id is 90021, 90025 or 90026, service_code will be appended, e.g., Entrega Turbo - M1020.</p>
	AllowSelfDesignAwb           bool                      `json:"allow_self_design_awb"`           // [Required] <p>To indicate whether the package allows for self-designed AWB, if allow_self_design_awb returns false, it means that the package does not allow for self-designed AWB and only the system-AWB can be used.</p>
	DaysToShip                   int64                     `json:"days_to_ship"`                    // [Required] <p>Shipping preparation time set by the seller when listing item on Shopee.<br /></p>
	ShipByDate                   int64                     `json:"ship_by_date"`                    // [Required] <p>The deadline to ship out the package.</p>
	PendingTerms                 []string                  `json:"pending_terms"`                   // [Required] <p>The list of pending terms. Applicable values:</p><p>- SYSTEM_PENDING: Under Shopee internal processing.</p><p>- KYC_PENDING: Under KYC checking (TW CB order only).</p><p>- ARRANGE_SHIPMENT_PENDING: Temporarily held due to 3PL capacity constraints.</p>
	PendingDescription           []string                  `json:"pending_description"`             // [Required] <p>The value of this field is the description of pending reason corresponding with pending terms. Applicable values:&nbsp;</p><p>- For SYSTEM_PENDING: Order is being processed by Shopee.</p><p>- For KYC_PENDING: Order is pending buyer TW KYC pre-authorization.</p><p>- For ARRANGE_SHIPMENT_PENDING: Allocating delivery resources due to high order volume. Label print will be available within 4 days after buyer paid.</p>
	TrackingNumber               string                    `json:"tracking_number"`                 // [Required] <p>The tracking number of this package.<br /></p>
	TrackingNumberExpirationDate int64                     `json:"tracking_number_expiration_date"` // [Required] <p>[TW only] Tracking number expiration date</p>
	PickupDoneTime               int64                     `json:"pickup_done_time"`                // [Required] <p>The timestamp when pickup is done.<br /></p>
	IsSplitUp                    bool                      `json:"is_split_up"`                     // [Required] <p>To indicate whether this parcel is split.<br /></p>
	ItemList                     []ResponseDataPackageItem `json:"item_list"`                       // [Required] <p>The lis of items in the package.</p>
	RecipientAddress             *RecipientAddress         `json:"recipient_address"`               // [Required] <p>This object contains detailed breakdown for the recipient address.<br />Different parameters might be masked according to each market and kind of seller.<br /><br />For TW region integrated channel orders will be all masked as "****". More details may refer the announcement.</p>
	ParcelChargeableWeightGram   int64                     `json:"parcel_chargeable_weight_gram"`   // [Required] <p>display weight used to calculate ASF for this parcel<br /></p>
	GroupShipmentId              int64                     `json:"group_shipment_id"`               // [Required] <p>The common identifier for multiple orders combined in the same parcel.<br /></p>
	VirtualContactNumber         string                    `json:"virtual_contact_number"`          // [Required] <p>[Only for TW non-integrated channel] The virtual phone number to contact the recipient.<br /></p>
	PackageQueryNumber           string                    `json:"package_query_number"`            // [Required] <p>[Only for TW non-integrated channel] The query number used in virtual phone number calls to contact the recipient of this package.<br /></p>
	SortingGroup                 string                    `json:"sorting_group"`                   // [Required] <p>[Only for TW 30029 channel] This field indicate the sorting group value of the package. This field is only available for logistics_channel_id = 30029 and after the package has been arranged for shipment.</p>
	IsShipmentArranged           bool                      `json:"is_shipment_arranged"`            // [Required] <p>Only effective when the package's logistics_status/fulfillment_status is LOGISTICS_READY.&nbsp;</p><p><br /></p><p>This parameter further distinguishes between two scenarios:</p><p>- true: Package shipment has been arranged (Seller has processed shipment, system is generating tracking number, not yet updated to LOGISTICS_REQUEST_CREATED, no duplicate action needed)</p><p>- false: Package awaiting shipment arrangement (Seller hasn't processed shipment yet, shipping arrangement required)</p>
	StatusInfoTag                *StatusInfoTag            `json:"status_info_tag"`                 // [Required] <p>Package shipping urgency tag information.</p>
	CanSplitOrder                bool                      `json:"can_split_order"`                 // [Required] <p>This field indicates whether this order can be split into multiple packages for separate shipment.</p><p>- true: Support splitting, can call v2.order.split_order to execute</p><p>- false: Does not support splitting</p>
	CanUnsplitOrder              bool                      `json:"can_unsplit_order"`               // [Required] <p>This field indicates whether this order can be unsplit.</p><p>- true: Support unsplitting, can call v2.order.unsplit_order to execute</p><p>- false: Does not support unsplitting</p>
	IsPreOrder                   bool                      `json:"is_pre_order"`                    // [Required] <p>This field indicates whether this order is a pre-order.</p><p>- true: Pre-order</p><p>- false: Non pre-order</p>
	PrescriptionImages           []string                  `json:"prescription_images"`             // [Required] <p>Return prescription images of this order, only for ID and PH whitelist sellers.</p><p><br /></p><p>Please add the prefix to review:</p><p>for ID:&nbsp;https://cf.shopee.co.id/file/+prescription_image</p><p>for PH: https://cf.shopee.ph/file/+prescription_image</p>
	PharmacistName               string                    `json:"pharmacist_name"`                 // [Required] <p>Name of the Pharmacist for Prescription Order.</p>
	PrescriptionApprovalTime     int64                     `json:"prescription_approval_time"`      // [Required] <p>Time of when the prescription is approved.</p>
	PrescriptionRejectionTime    int64                     `json:"prescription_rejection_time"`     // [Required] <p>Time of when the prescription is rejected.</p>
	IsBuyerShopCollection        bool                      `json:"is_buyer_shop_collection"`        // [Required] <p>To indicate if this order is buyer self collection at store order.</p>
	BuyerProofOfCollection       []string                  `json:"buyer_proof_of_collection"`       // [Required] <p>The image url of the proof for buyer self collection at the store.</p>
	PreparationEndTime           int64                     `json:"preparation_end_time"`            // [Required] <p>The system-calculated deadline for package preparation. When the package fulfillment_status/logistics_status changes to "LOGISTICS_READY", the system calculates this time based on the "Preparation Time" configured for the logistics channel of this package.&nbsp;</p><p><br /></p><p>Notes:&nbsp;</p><p>1) Only effective for logistics channels that have Auto Call Driver enabled and Preparation Time configured.</p><p>2) Seller needs to complete packing and waybill printing before this time to ensure the package is ready when the driver arrives.</p><p>3) When this time is reached, the system will automatically arrange shipment and trigger driver dispatch:</p><p>- If driver calling is successful, the package fulfillment_status/logistics_status will change from “LOGISTICS_READY” to “LOGISTICS_REQUEST_CREATED”.</p><p>- If driver calling fails, the package fulfillment_status/logistics_status will remain unchanged, and the seller must arrange shipment manually.</p>
	DriverInfo                   *DriverInfo               `json:"driver_info"`                     // [Required] <p>After the driver is successfully called, the driver's information will be returned.</p><p><br /></p><p>Note: Data availability depends on the specific 3PL provider; certain fields may be omitted due to provider policies, PII restrictions, or data unavailability.</p>
}

type ResponseDataPackageItem struct {
	ItemId            int64  `json:"item_id"`             // [Required] <p>Shopee's unique identifier for an item.</p>
	ModelId           int64  `json:"model_id"`            // [Required] <p>Shopee's unique identifier for a model.</p>
	ItemSku           string `json:"item_sku"`            // [Required] <p>A item SKU (stock keeping unit) is an identifier defined by a seller, sometimes called parent SKU. Item SKU can be assigned to an item in Shopee Listings.<br /></p>
	ModelSku          string `json:"model_sku"`           // [Required] <p>ID of the model that belongs to the same item.<br /></p>
	ModelQuantity     int64  `json:"model_quantity"`      // [Required] <p>The number of identical items/variations purchased at the same time by the same buyer from one listing/item.</p>
	OrderItemId       int64  `json:"order_item_id"`       // [Required] <p>The identify of order item. For items in one same bundle deal promotion, the order_item_id should share the same id, such as 1,2. For items not in bundle deal promotion, the order_item_id should be the same as item_id.</p>
	PromotionGroupId  int64  `json:"promotion_group_id"`  // [Required] <p>The identify of product promotion.</p>
	ProductLocationId string `json:"product_location_id"` // [Required] <p>The warehouse ID of the item.</p>
	ConsultationId    string `json:"consultation_id"`     // [Required] <p>An identifier of teleconsultation session which buyer did to order this item. Empty if item is not ordered through teleconsultation session</p>
}

type ResponseDataPageInfo struct {
	Cursor  int64 `json:"cursor"`   // [Required]
	HasNext bool  `json:"has_next"` // [Required]
}

type ResponseDataPagination struct {
	TotalCount int64  `json:"total_count"` // [Required] <p>Total orders can be returned with your query<br /></p>
	NextCursor string `json:"next_cursor"` // [Required] <p>if packages is not empty or length of packages &lt;= page_size. You should pass the next_cursor in the next request as page_sentinel. <br /></p>
	More       bool   `json:"more"`        // [Required] <p>To indicate, it's a the last page or not<br /></p>
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

type ResponseDataPriceInfo struct {
	CurrentPrice  float64 `json:"current_price"`  // [Required] Current price of item
	OriginalPrice float64 `json:"original_price"` // [Required] Original price of item
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

type ResponseDataShop struct {
	ShopName string `json:"shop_name"` // [Required] <p>Name of the shop.<br /></p>
}

type ResponseDataSort struct {
	SortType int64 `json:"sort_type"` // [Required] <p>As same as request param</p><p></p>
	IsAsc    bool  `json:"is_asc"`    // [Required] <p>As same as request param</p><p></p>
}

type ResponseDataSsp struct {
	SspId             int64              `json:"ssp_id"`             // [Required] <p>Shopee's unique identifier for Shopee Standard Product.</p>
	ProductName       string             `json:"product_name"`       // [Required] <p>Name of Shopee&nbsp;Standard Product.</p>
	CategoryId        int64              `json:"category_id"`        // [Required] <p>Shopee's unique identifier for a category of Shopee&nbsp;Standard Product.</p>
	Description       *Description       `json:"description"`        // [Required] <p>Description of Shopee&nbsp;Standard Product.</p>
	Weight            string             `json:"weight"`             // [Required] <p>The weight of Shopee Standard Product, the unit is KG.</p>
	BrandInfo         *Brand             `json:"brand_info"`         // [Required] <p>Brand of Shopee Standard Product.</p>
	AttributeList     []SspAttribute     `json:"attribute_list"`     // [Required]
	Media             *SspMedia          `json:"media"`              // [Required]
	CompatibilityInfo *CompatibilityInfo `json:"compatibility_info"` // [Required]
	Dimension         *Dimension         `json:"dimension"`          // [Required] <p>The dimension of this Shopee Standard Product.</p>
	GlobalCode        *GlobalCode        `json:"global_code"`        // [Required]
	SalesInfo         *SalesInfo         `json:"sales_info"`         // [Required]
}

type ResponseDataStandardiseTierVariation struct {
	VariationId         int64                                     `json:"variation_id"`          // [Required] <p>Standardise Tier variation ID.<br /></p>
	VariationName       string                                    `json:"variation_name"`        // [Required] <p>Standardise Tier variation Name.<br /></p>
	VariationGroupId    int64                                     `json:"variation_group_id"`    // [Required] <p>Standardise Tier variation Group ID.<br /></p>
	VariationOptionList []StandardiseTierVariationVariationOption `json:"variation_option_list"` // [Required] <p>Standardise Tier variation Options List.<br /></p>
}

type ResponseDataSubItem struct {
	ItemId      int64  `json:"item_id"`      // [Required] Shopee's unique identifier for an item.
	Status      int64  `json:"status"`       // [Required] The status of add on deal item：enable = 1；disable =2
	ModelId     int64  `json:"model_id"`     // [Required] Shopee's unique identifier for a model.
	FailError   string `json:"fail_error"`   // [Required]
	FailMessage string `json:"fail_message"` // [Required]
}

type ResponseDataSuccess struct {
	ItemId    int64       `json:"item_id"`   // [Required] The identity of product item.
	Promotion []Promotion `json:"promotion"` // [Required] Item promotion info list
}

type ResponseDataTaxInfo struct {
	Ncm               string         `json:"ncm"`                 // [Required] <p>Mercosur Common Nomenclature, it is a convention between Mercosur member countries to easily recognize goods, services and productive factors negotiated among themselves.(BR region)</p><p><br /></p><p>Note: ncm = "00" means that this item doesn't have a NCM.</p>
	DiffStateCfop     string         `json:"diff_state_cfop"`     // [Required] Tax Code of Operations and Installments for orders that seller and buyer are in different states. It identifies a specific operation by category at the time of issuing the invoice. (BR region)
	Csosn             string         `json:"csosn"`               // [Required] Code of Operation Status – Simples Nacional, code for company operations to identify the origin of the goods and the taxation regime of the operations. (BR region)
	Origin            string         `json:"origin"`              // [Required] Product source, domestic or foreig (BR region)
	Cest              string         `json:"cest"`                // [Required] <p>Tax Replacement Specifying Code (CEST), to separate within the same NCM products that do or do not have ICMS tax substitution. (BR region)<br /></p><p><br /></p><p>Note: cest = "00" means that this item doesn't have a CEST.<br /></p>
	MeasureUnit       string         `json:"measure_unit"`        // [Required] (BR region)
	InvoiceOption     InvoiceOption  `json:"invoice_option"`      // [Required] Value shuold be one of NO_INVOICES VAT_MARGIN_SCHEME_INVOICES VAT_INVOICES NON_VAT_INVOICES and if value is NON_VAT_INVOICE vat_rate should be null (PL region)
	VatRate           string         `json:"vat_rate"`            // [Required] Value should be one of 0% 5% 8% 23% NO_VAT_RATE (PL region)
	HsCode            string         `json:"hs_code"`             // [Required] HS Code (Only for IN region)
	TaxCode           string         `json:"tax_code"`            // [Required] Tax Code (Only for IN region)
	TaxType           TaxType        `json:"tax_type"`            // [Required] <p>tax_type only for TW whitelist shop. Shopee will referred Tax type when substitute sellers for issuing e-receipts to buyers. All variations share the same tax type. The meaning of value:&nbsp;</p><p>0: no tax type</p><p>1: tax-able</p><p>2: tax-free<br /></p>
	Pis               string         `json:"pis"`                 // [Required] <p>Only for BR shop.</p><p>PIS - Programa de Integração Social (Social Integration Program). It is a government tax to collect resources for the payment of unemployment insurance and other employee related rights.</p><p>PIS % - the tax applied to this product</p>
	Cofins            string         `json:"cofins"`              // [Required] <p>Only for BR shop.<br /></p><p>COFINS – Contribuição para Financiamento da Seguridade Social (Contribution for Social Security Funding). It&nbsp;is a government tax to collect resources for public health system and social security.</p><p>COFINS&nbsp;% - the tax applied to this product</p>
	IcmsCst           string         `json:"icms_cst"`            // [Required] <p>Only for BR shop.<br /></p><p>ICMS - Imposto sobre Circulação de Mercadorias e Serviços (Circulation of Goods and Services Tax).&nbsp;</p><p>CST - Código da Situação Tributária (Tax Situation Code) is represented by a combination of 3 numbers with the purpose of demonstrating the origin of a product and determining the form of taxation that will apply to it. Therefore, each digit in the CST Table has a specific meaning: the first digit indicates the origin of the operation, the second digit represents the ICMS taxation on the operation and the third digit provides additional information about the form of taxation.</p>
	PisCofinsCst      string         `json:"pis_cofins_cst"`      // [Required] <p>Only for BR shop.</p><p>The CST PIS/Cofins is a code on the Electronic Invoice (NF-e) that identifies the tax situation of PIS (Programa de Integração Social) and Cofins (Contribuição para o Financiamento da Seguridade Social) in sales of goods.</p>
	FederalStateTaxes string         `json:"federal_state_taxes"` // [Required] <p>Only for BR shop.</p><p>Enter the total percentage of the combination of federal, state, and municipal taxes, using up to two decimals.</p>
	OperationType     OperationType  `json:"operation_type"`      // [Required] <p>Only for BR shop.</p><p>1: Retailer</p><p>2: Manufacturer</p>
	ExTipi            string         `json:"ex_tipi"`             // [Required] <p>Only for BR shop.<br /></p><p>The EXTIPI field in the NF-e (Nota Fiscal Eletrônica) is used to indicate if there's an exception to the IPI (Imposto sobre Produtos Industrializados) tax rate for a specific product.</p>
	FciNum            string         `json:"fci_num"`             // [Required] <p>Only for BR shop.<br /></p><p>The FCI Control Number is a unique identifier assigned to each import FCI (Import Content Form). It's mandatory on the corresponding NF-e (electronic invoice) to ensure compliance with Brazilian import tax regulations.</p>
	RecopiNum         string         `json:"recopi_num"`          // [Required] <p>Only for BR shop.</p><p>RECOPI NACIONAL is a Brazilian government system that facilitates the registration and management of tax-exempt operations involving paper destined for printing books, newspapers, and periodicals (known as "papel imune" in Portuguese).</p>
	AdditionalInfo    string         `json:"additional_info"`     // [Required] <p>Only for BR shop.</p><p>Include relevant information to display on Invoice.</p>
	GroupItemInfo     *GroupItemInfo `json:"group_item_info"`     // [Required] <p>Only for BR shop.</p><p>Required if the item is a group item.</p>
	ExportCfop        string         `json:"export_cfop"`         // [Required] <p>7101 - for sales of self-produced goods</p><p>7102 - resale of third-party goods</p><p><br /></p><p>a tax code used in Brazil to classify and identify the nature of goods or services transactions for tax purposes. This is used for goods export to other counties</p>
}

type ResponseDataTierVariation struct {
	Name       string                `json:"name"`        // [Required] Variation name
	OptionList []TierVariationOption `json:"option_list"` // [Required] Options of this variation
}

type ResponseDataTierVariationOption struct {
	Option string          `json:"option"` // [Required] Tier option.
	Image  *FieldImageInfo `json:"image"`  // [Required] Image information of tier.
}

type ResponseDataTrackingInfo struct {
	UpdateTime          int64    `json:"update_time"`          // [Required] <p>The timestamps when reverse logistics info has been updated for Normal RR from warehouse to seller, pushed from third party logistics provider to Shopee.</p>
	TrackingDescription string   `json:"tracking_description"` // [Required] <p>The description of reverse logistics tracking info&nbsp;for Normal RR from warehouse to seller, pushed by third party logistics provider to Shopee.</p><p><br /></p><p>These would match the tracking description displayed to sellers on Seller Center return/refund detail page.</p>
	EpopImageList       []string `json:"epop_image_list"`      // [Required] <p>Image URLs of electronic proof of pickup (ePOP) after return parcel has been picked up from the warehouse for Normal RR with warehouse validation.</p>
	EpodImageList       []string `json:"epod_image_list"`      // [Required] <p>Image URLs of electronic proof of delivery (ePOD) after return parcel has been delivered to the seller for Normal RR with warehouse validation.</p>
}

type ResponseDataTrackingNumber struct {
	BindingId               string `json:"binding_id"`                 // [Required] <p>The generated binding ID.</p>
	FirstMileTrackingNumber string `json:"first_mile_tracking_number"` // [Required] <p>The generate first-mile tracking number, value might be blank.</p>
	Status                  string `json:"status"`                     // [Required] <p>The logistics status for first-mile tracking number. Status could be:</p><p>CANCELED<br />CANCELING<br />DELIVERED<br />NOT_AVAILABLE<br />ORDER_CREATED<br />ORDER_RECEIVED<br />PICKED_UP</p><p><br /></p><p>Note:&nbsp;NOT_AVAILABLE status means that Binding ID / First Mile Tracking Number is not yet bound with any order.</p>
	Reason                  string `json:"reason"`                     // [Required] <p>Indicate the reason when Shopee failed to place courier order to 3PL (Kuaidi 100 supporting) or courier company cancelled the order.</p><p><br /></p><p>Note: Will be empty if status is not CANCELED.</p>
	DeclareDate             string `json:"declare_date"`               // [Required] <p>The declare date of binding ID/first-mile tracking number.</p>
}

type ResponseDataVideoInfo struct {
	VideoUrlList     []VideoUrl     `json:"video_url_list"`     // [Required] Video playback URL list.
	ThumbnailUrlList []ThumbnailUrl `json:"thumbnail_url_list"` // [Required] Video thumbnail image URL list.
	Duration         int64          `json:"duration"`           // [Required] Duration of this video, in seconds.
}

type ResponseResult struct {
	RequestId   int64  `json:"request_id"`   // [Required] <p>Unique task identifier that includes one or more tax documents to be downloaded according to the filters sent in the request.<br /></p>
	FailError   string `json:"fail_error"`   // [Required] <p>Indicate error type if one element hit error. Empty if no error happened.</p>
	FailMessage string `json:"fail_message"` // [Required] <p>Indicate error details if one element hit error. Empty if no error happened.<br /></p>
}

type Result struct {
	CommentId   int64  `json:"comment_id"`   // [Required] The identity of comment.
	FailError   string `json:"fail_error"`   // [Required] Indicate error details if one element hit error.
	FailMessage string `json:"fail_message"` // [Required] Indicate error type if one element hit error.
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

type ResultReport struct {
	BroadCir          float64 `json:"broad_cir"`           // [Required] <p>The direct advertising cost of sales, or direct ACOS, measures how much your ad costs relative to the revenue generated from sales of the advertised product. It is the amount spent on the ad divided by the amount of sales revenue for the advertised product that is attributed to the ad. Direct ACOS = expense ÷ direct GMV × 100%.<br /></p>
	BroadGmv          float64 `json:"broad_gmv"`           // [Required] <p>The amount of sales revenue generated from shoppers purchasing products within 7 days of them clicking on your ad.<br /></p>
	BroadOrder        int64   `json:"broad_order"`         // [Required] <p>The number of times shoppers purchased any product from your shop within 7 days of them clicking on your ad.</p>
	BroadOrderAmount  int64   `json:"broad_order_amount"`  // [Required] <p>The total quantity of products purchased by shoppers within 7 days of them clicking on your ad.</p>
	BroadRoi          float64 `json:"broad_roi"`           // [Required] <p>Return on ad spend (ROAS) measures how much revenue is generated by your ad relative to the cost of the ad. It is the amount of sales revenue attributed to your ad divided by the amount spent on the ad. ROAS = GMV ÷ expense. (Note: We recommend monitoring ROAS trends on a weekly basis.)<br /></p>
	Clicks            int64   `json:"clicks"`              // [Required] <p>The number of times shoppers click on your ad. (Note: Shopee filters out repeated clicks from the same shopper that occur within a short time frame.)<br /></p>
	Expense           float64 `json:"expense"`             // [Required] <p>The amount spent on your ad.</p>
	Cpc               float64 `json:"cpc"`                 // [Required] <p>The cost per conversion is how much each conversion costs, on average. It is the amount spent on your ad divided by the number of conversions attributed to the ad. Cost per conversion = expense ÷ conversions.<br /></p>
	Cpdc              float64 `json:"cpdc"`                // [Required] <p>The cost per direct conversion is how much each direct conversion costs, on average. It is the amount spent on your ad divided by the number of direct conversions attributed to the ad. Cost per direct conversion = expense ÷ direct conversions.<br /></p>
	Cr                float64 `json:"cr"`                  // [Required] <p>The conversion rate measures how often shoppers end up purchasing something from your shop after clicking on your ad. It is the number of conversions attributed to your ad divided by the number of clicks on the ad. Conversion rate = conversions ÷ clicks × 100%.<br /></p>
	DirectCr          float64 `json:"direct_cr"`           // [Required] <p>The direct conversion rate measures how often shoppers end up purchasing the advertised product after clicking on the ad. Direct conversion rate is the number of direct conversions divided by the number of clicks. Direct conversion rate = direct conversions ÷ clicks × 100%.<br /></p>
	DirectCir         float64 `json:"direct_cir"`          // [Required] <p>The direct advertising cost of sales, or direct ACOS, measures how much your ad costs relative to the revenue generated from sales of the advertised product. It is the amount spent on the ad divided by the amount of sales revenue for the advertised product that is attributed to the ad. Direct ACOS = expense ÷ direct GMV × 100%.<br /></p>
	DirectOrder       int64   `json:"direct_order"`        // [Required] <p>The number of times shoppers purchased the advertised product within 7 days of them clicking on the ad.<br /></p>
	DirectOrderAmount int64   `json:"direct_order_amount"` // [Required] <p>The total quantity of the advertised product purchased by shoppers within 7 days of them clicking on the ad.<br /></p>
	DirectRoi         float64 `json:"direct_roi"`          // [Required] <p>The direct return on ad spend, or direct ROAS, measures how much revenue is generated from sales of the advertised product, relative to the cost of the ad. It is the amount of sales revenue for the advertised product attributed to the ad, divided by the amount spent on the ad. Direct ROAS = direct GMV ÷ expense.<br /></p>
	Impression        int64   `json:"impression"`          // [Required] <p>The number of times shoppers see your ad.<br /></p>
}

type Return struct {
	Image                    []string         `json:"image"`                       // [Required] Image URLs of return.
	Reason                   string           `json:"reason"`                      // [Required] <p>Indicates the original return reason submitted by the buyer when initiating the return request.</p><p><br /></p><p>Applicable values: See Data Definition- ReturnReason and Reassessed Request Reason.</p><p><br /></p><p>Note: There may be cases where Shopee Agent updates the return request with a "Reassessed Return Reason" after reviewing more details about the buyer's return request and potentially after requesting evidence from the seller.&nbsp;If the platform updates the return reason during this process, the reassessed outcome will be provided separately in the&nbsp;<b>reassessed_request_reason</b>&nbsp;field.</p>
	TextReason               string           `json:"text_reason"`                 // [Required] Reason that buyer provide.
	ReassessedRequestReason  string           `json:"reassessed_request_reason"`   // [Required] <p>Indicates the return reason reassessed by the platform as more suitable.</p><p><br /></p><p>There may be cases where Shopee Agent updates the return request with a "Reassessed Return Reason" after reviewing more details about the buyer's return request and potentially after requesting evidence from the seller.&nbsp;</p><p><br /></p><p>Applicable values: See Data Definition- ReturnReason and Reassessed Request Reason. If no reassessment has been made, the value will be NONE.</p>
	ReturnSn                 string           `json:"return_sn"`                   // [Required] The serial number of return.
	RefundAmount             float64          `json:"refund_amount"`               // [Required] Amount of the refund.
	Currency                 string           `json:"currency"`                    // [Required] Currency of the return.
	CreateTime               int64            `json:"create_time"`                 // [Required] The time of return create.
	UpdateTime               int64            `json:"update_time"`                 // [Required] The time of modify return.
	Status                   string           `json:"status"`                      // [Required] Enumerated type that defines the current status of the return. Applicable values: See Data Definition- ReturnStatus.
	DueDate                  int64            `json:"due_date"`                    // [Required] The last time seller deal with this return.
	TrackingNumber           string           `json:"tracking_number"`             // [Required] The tracking number assigned by the shipping carrier for item shipment.
	DisputeReason            []string         `json:"dispute_reason"`              // [Required] The reason of seller dispute return. While the return has been disputed, this field is useful. Applicable values: See Data Definition- ReturnDisputeReason.
	DisputeTextReason        []string         `json:"dispute_text_reason"`         // [Required] The reason that seller provide. While the return has been disputed, this field is useful.
	NeedsLogistics           bool             `json:"needs_logistics"`             // [Required] Items to be sent back to seller. Can be either integrated/non-integrated.
	AmountBeforeDiscount     float64          `json:"amount_before_discount"`      // [Required] Order price before discount.
	User                     *User            `json:"user"`                        // [Required]
	Item                     []ReturnItem     `json:"item"`                        // [Required]
	OrderSn                  string           `json:"order_sn"`                    // [Required] Shopee's unique identifier for an order.
	ReturnShipDueDate        int64            `json:"return_ship_due_date"`        // [Required] The due date for buyer to ship order.
	ReturnSellerDueDate      int64            `json:"return_seller_due_date"`      // [Required] The due date for seller to deal with this return when buyer have shipped order.
	NegotiationStatus        string           `json:"negotiation_status"`          // [Required] Counter status. See "Data Definition - NegotiationStatus"
	SellerProofStatus        string           `json:"seller_proof_status"`         // [Required] <p>Proof status. See "Data Definition - SellerProofStatus"</p>
	SellerCompensationStatus string           `json:"seller_compensation_status"`  // [Required] Compensation status. See "Data Definition - SellerCompensationStatus"
	ReturnRefundType         string           `json:"return_refund_type"`          // [Required] <p>To indicate whether the return is RRBOC (Return/Refund request raised before Order Complete) or RRAOC&nbsp;(Return/Refund request raised after Order Complete).<br /></p>
	ReturnSolution           int64            `json:"return_solution"`             // [Required] <p>To indicate the most updated solution of the Return/Refund request (NOTE: this is not the solution during negotiation). Applicable value:&nbsp;</p><p>- 0: Return and Refund</p><p>- 1: Refund Only</p>
	IsSellerArrange          bool             `json:"is_seller_arrange"`           // [Required] <p>To indicate whether the return_sn is using the “Seller Arrange” return method. This would only be True for TW and BR.</p>
	IsShippingProofMandatory bool             `json:"is_shipping_proof_mandatory"` // [Required] <p>To indicate whether uploading shipping proof is mandatory for seller to confirm "Arrange Pickup" when is_seller_arrange = true.</p>
	ReturnRefundRequestType  int64            `json:"return_refund_request_type"`  // [Required] <p>To indicate the type of return refund request, whether it is a Normal RR request, an In-transit RR request, and a Return on the Spot:&nbsp;</p><p style>0:&nbsp;Normal RR (RR is raised by the buyer after delivery done / estimated delivery date)</p><p style>1: In-transit RR (RR is raised by the buyer while item is still in-transit to buyer)</p><p style>2: Return-on-the-Spot (RR is raised by the driver after buyer rejected parcel at delivery)</p><p><br /></p><p>For more details, see Data Definition- Return Refund Request Type.</p>
	ValidationType           string           `json:"validation_type"`             // [Required] <p>To indicate whether seller or warehouse will expect to receive the return parcel from buyer and validate the condition of the parcel:&nbsp;</p><p>- seller_validation&nbsp;</p><p>- warehouse_validation</p><p><br /></p><p>For more details, see Data Definition- ValidationType.</p>
	IsArrivedAtWarehouse     int64            `json:"is_arrived_at_warehouse"`     // [Required] <p><b>[Only for validation_type =&nbsp;warehouse_validation]</b> Indicates the parcel’s check-in status at the warehouse. This field helps sellers quickly determine whether the parcel has arrived at the warehouse or has been rejected.&nbsp;</p><p><br /></p><p>Applicable values:</p><p>1: Pending Inbound</p><p>2: Rejected</p><p>3: Inbound</p><p>4: Cancelled</p>
	FollowUpActionList       []FollowUpAction `json:"follow_up_action_list"`       // [Required] <p><b>[Only for validation_type =&nbsp;warehouse_validation]</b>&nbsp;Warehouse handling actions for each item in the parcel.</p>
}

type ReturnAddress struct {
	WhsId string `json:"whs_id"` // [Required] <p>To indicate the warehouse id where item will be returned to.&nbsp;<br /></p><p>Please call&nbsp;v2.shop.get_warehouse_detail to check the detailed warehouse information the item returned to with the field "location_id" of the&nbsp;v2.shop.get_warehouse_detail match to the field"whs_id"of the v2.return.get_return_detail.<br /></p><p><br /></p><p>For fulfillment by Shopee (FBS) &amp; multi warehouse sellers, R/R orders will be returned back to the nearest warehouse of buyer address instead of going back to only 1 default return address like a normal seller.If it's a normal seller, then the field will be response empty.<br /></p>
}

type ReturnItem struct {
	ModelId      int64    `json:"model_id"`       // [Required] Shopee's unique identifier for a variation of an item.
	Name         string   `json:"name"`           // [Required] Name of item in local language.
	Images       []string `json:"images"`         // [Required] Image URLs of item.
	Amount       int64    `json:"amount"`         // [Required] Amount of this item.
	ItemPrice    float64  `json:"item_price"`     // [Required] The price of item.
	IsAddOnDeal  bool     `json:"is_add_on_deal"` // [Required] To indicate if this item belongs to an addon deal.
	IsMainItem   bool     `json:"is_main_item"`   // [Required] To indicate if this item is main item or sub item. True means main item, false means sub item.
	AddOnDealId  int64    `json:"add_on_deal_id"` // [Required] The unique identity of an addon deal.
	ItemId       int64    `json:"item_id"`        // [Required] The id of item.
	ItemSku      string   `json:"item_sku"`       // [Required] The sku of item.
	VariationSku string   `json:"variation_sku"`  // [Required] The variation sku of item
}

type ReturnPickupAddress struct {
	Address  string `json:"address"`  // [Required] <p>To indicate receiver's address</p>
	Name     string `json:"name"`     // [Required] <p>To indicate receiver's name<br /></p>
	Phone    string `json:"phone"`    // [Required] <p>To indicate receiver's phone<br />[Only for TW non-integrated channel] Will return "****" when the "virtual_contact_number" is available<br /></p>
	Town     string `json:"town"`     // [Required] <p>To indicate receiver's town</p>
	District string `json:"district"` // [Required] <p>To indicate receiver's district<br /></p>
	City     string `json:"city"`     // [Required] <p>To indicate receiver's city</p>
	State    string `json:"state"`    // [Required] <p>To indicate receiver's state</p>
	Region   string `json:"region"`   // [Required] <p>To indicate receiver's region</p>
	Zipcode  string `json:"zipcode"`  // [Required] <p>To indicate receiver's zip code<br /></p>
}

type ReturnRefundOrder struct {
	OrderSn        string `json:"order_sn"`        // [Required] <p>Order SN.</p>
	DetailedReason int64  `json:"detailed_reason"` // [Required] <p>Reason. Applicable values:&nbsp;</p><p>1001: Return Refund</p><p>1002: Parcel Split Cancellation</p><p>1003: First Mile Pick up fail</p><p>1004: Order inclusion<br /></p><p>10005: Out of Stock<br />10006: Undeliverable area<br />10007: Cannot support COD<br />10008: Logistics request cancelled<br />10009: Logistics pickup failed<br />10010: Logistics not ready<br />10011: Inactive seller<br />10012: Seller did not ship order<br />10013: Order did not reach warehouse<br /></p><p>10014: Seller asked to cancel<br /></p><p>10015: Non-receipt<br />10016: Wrong item<br />10017: Damaged item<br />10018: Incomplete product<br />10019: Fake item<br />10020: Functional Damage<br />10021: Return Refund</p>
}

type ReturnSortCode struct {
	ReturnFirstSortCode string `json:"return_first_sort_code"` // [Required] <p>The first-level sort code for 3PL doing RTS.<br /></p>
}

type ReverseLogisticsCarrier struct {
	ReverseLogisticsCarrierId   int64  `json:"reverse_logistics_carrier_id"`   // [Required] <p>To indicate the id of the non-integrated reverse logistics channel used by seller.</p>
	ReverseLogisticsCarrierName string `json:"reverse_logistics_carrier_name"` // [Required] <p>To indicate the selected carrier name from the list of carrier names provided.</p>
}

type ReversedTrackingInfo struct {
	UpdateTime  int64  `json:"update_time"` // [Required] <p>The time when the reversed logistics&nbsp;tracking info is updated.</p>
	Description string `json:"description"` // [Required] <p>The description of the reversed logistics tracking info.</p>
}

type SalesInfo struct {
	ModelSettingList     []ModelSetting                         `json:"model_setting_list"`      // [Required]
	StdTierVariationList []ResponseDataStandardiseTierVariation `json:"std_tier_variation_list"` // [Required] <p>Standardise Variation config of Shopee Standard Product.<br /></p>
}

type SampleEvidence struct {
	Type      int64  `json:"type"`      // [Required] <p>The type of sample evidence. Applicable value:&nbsp;</p><p>- 1: Image</p>
	Url       string `json:"url"`       // [Required] <p>The link of sample evidence.</p>
	Thumbnail string `json:"thumbnail"` // [Required] <p>The link of the thumbnail of sample evidence.</p>
}

type SddListing struct {
	ItemId           int64 `json:"item_id"`            // [Required] <p>Item ID.</p>
	CurrentSddStatus int64 `json:"current_sdd_status"` // [Required] <p>Current SDD Status. Applicable values:&nbsp;</p><p>1: Yes</p><p>0: No</p>
}

type SearchAttributeValueListRequest struct {
	AttributeId int64   `json:"attribute_id"`         // [Required]
	ValueName   *string `json:"value_name,omitempty"` // [Optional] <p>search the keywords of the attributes value</p>
	Cursor      int64   `json:"cursor"`               // [Required]
	Limit       int64   `json:"limit"`                // [Required] <p>The range is 1 to 100</p>
}

type SearchAttributeValueListResponse struct {
	BaseResponse `json:",inline"`                     // Common response fields
	Msg          string                               `json:"msg,omitempty"`           //
	DebugMessage string                               `json:"debug_message,omitempty"` //
	Response     SearchAttributeValueListResponseData `json:"response"`                // Actual response data
}

type SearchAttributeValueListResponseData struct {
	ValueList []Value               `json:"value_list"` // [Required]
	PageInfo  *ResponseDataPageInfo `json:"page_info"`  // [Required]
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

type SearchItemRequest struct {
	Offset          *string      `json:"offset,omitempty" url:"offset,omitempty"`                     // [Optional] Specifies the starting entry of data to return in the current call. Default is empty. if data is more than one page, the offset can be some entry to start next call.
	PageSize        int64        `json:"page_size" url:"page_size"`                                   // [Required] the size of one page.
	ItemName        *string      `json:"item_name,omitempty" url:"item_name,omitempty"`               // [Optional] name of item.
	AttributeStatus *int64       `json:"attribute_status,omitempty" url:"attribute_status,omitempty"` // [Optional] 1:get item lack of requires attribute.   2:get item lack of optional attribute.
	ItemSku         *string      `json:"item_sku,omitempty" url:"item_sku,omitempty"`                 // [Optional] <p>sku. If you search for item_sku and item_name at the same time, only the results that match item_sku will be returned. If you search for item_sku and attribute_status at the same time, the results that match both item_sku and attribute_status will be returned.<br /></p>
	ItemStatus      []ItemStatus `json:"item_status,omitempty" url:"item_status,omitempty"`           // [Optional] <p>NORMAL/BANNED/UNLIST/<b><font color="#c24f4a">REVIEWING/SELLER_DELETE/SHOPEE_DELETE</font></b></p><p>If you want to search multiple status, please upload the url like this: item_status=NORMAL&amp;item_status=BANNED<br /></p>
	DeboostOnly     *bool        `json:"deboost_only,omitempty" url:"deboost_only,omitempty"`         // [Optional] <p>If deboost_only is true, then API will return items whose deboost is true, if deboost_only is empty or false, then API will return items whose deboost is true and false simultaneously<br /></p>
}

type SearchItemResponse struct {
	BaseResponse `json:",inline"`       // Common response fields
	Response     SearchItemResponseData `json:"response"` // Actual response data
}

type SearchItemResponseData struct {
	ItemIdList []int64 `json:"item_id_list"` // [Required] List of  item ID.
	TotalCount int64   `json:"total_count"`  // [Required] Total num of items match condation.
	NextOffset string  `json:"next_offset"`  // [Required] If has_next_page is true, this value need set to next request.offset
}

type SearchPackageListRequest struct {
	Filter     *Filter     `json:"filter,omitempty"` // [Optional]
	Pagination *Pagination `json:"pagination"`       // [Required]
	Sort       *Sort       `json:"sort,omitempty"`   // [Optional]
}

type SearchPackageListResponse struct {
	BaseResponse `json:",inline"`              // Common response fields
	Mesage       string                        `json:"mesage,omitempty"` // <p>Indicate error details if hit error. Empty if no error happened.</p>
	Response     SearchPackageListResponseData `json:"response"`         // Actual response data
}

type SearchPackageListResponseData struct {
	PackagesList []Packages              `json:"packages_list"` // [Required]
	Pagination   *ResponseDataPagination `json:"pagination"`    // [Required]
	Sort         *ResponseDataSort       `json:"sort"`          // [Required] <p>As same as request param</p>
}

type SearchUnpackagedModelListRequest struct {
	PageSize        int64   `json:"page_size"`                   // [Required] <p>Each result set is returned as a page of entries. Use the "page_size" filters to control the maximum number of entries to retrieve per page (i.e., per call). This integer value is used to specify the maximum number of entries to return in a single "page" of data. The limit of page_size if between 1 and 48.</p>
	Cursor          *string `json:"cursor,omitempty"`            // [Optional] <p>Specifies the starting entry of data to return in the current call. Default is "". If data is more than one page, the cursor can be some entry to start next call.</p>
	ItemId          *int64  `json:"item_id,omitempty"`           // [Optional] <p>Shopee's unique identifier for an item.</p>
	ItemName        *string `json:"item_name,omitempty"`         // [Optional] <p>Name of the item.</p>
	ModelId         *int64  `json:"model_id,omitempty"`          // [Optional] <p>Shopee's unique identifier for a model under item.</p>
	UnpackagedSkuId *string `json:"unpackaged_sku_id,omitempty"` // [Optional] <p>Unpackaged SKU ID of the model.</p>
}

type SearchUnpackagedModelListResponse struct {
	BaseResponse `json:",inline"`                      // Common response fields
	Response     SearchUnpackagedModelListResponseData `json:"response"` // Actual response data
}

type SearchUnpackagedModelListResponseData struct {
	TotalCount int64                                        `json:"total_count"` // [Required] <p>Total number of models that match the condition.</p>
	NextCursor string                                       `json:"next_cursor"` // [Required] <p>Pass the next_cursor in the next request as cursor to get the next page data.</p>
	ModelList  []SearchUnpackagedModelListResponseDataModel `json:"model_list"`  // [Required] <p>List of models that match the condition.</p>
}

type SearchUnpackagedModelListResponseDataModel struct {
	ItemId          int64  `json:"item_id"`           // [Required] <p>Shopee's unique identifier for an item.</p>
	ItemName        string `json:"item_name"`         // [Required] <p>Name of the item.</p>
	ModelId         int64  `json:"model_id"`          // [Required] <p>Shopee's unique identifier for a model under item.&nbsp;0 for no model item.</p>
	UnpackagedSkuId string `json:"unpackaged_sku_id"` // [Required] <p>Unpackaged SKU ID of the model.</p>
}

type SelectedKeywords struct {
	Keyword          string  `json:"keyword"`             // [Required] <p>bid keyword for each campaign</p>
	MatchType        string  `json:"match_type"`          // [Required] <p>exact, broad<br /></p>
	BidPricePerClick float64 `json:"bid_price_per_click"` // [Required] <p>the bid price of keyword</p>
}

type SellerCompensation struct {
	SellerCompensationStatus  string  `json:"seller_compensation_status"`   // [Required] To indicate whether the seller is eligible for raising a compensation request. See "Data Definition - SellerCompensationStatus"
	SellerCompensationDueDate int64   `json:"seller_compensation_due_date"` // [Required] To indicate the deadline for requesting the compensation
	CompensationAmount        float64 `json:"compensation_amount"`          // [Required] To indicate the compensation amount that the agent decided
}

type SellerProductRebate struct {
	Amount              float64 `json:"amount"`                // [Required] <p>This is the portion of Shopee rebate borne by seller.</p>
	CommissionFeeOffset float64 `json:"commission_fee_offset"` // [Required] <p>This is the offset to gross commission fee, reducing it to the net value.</p>
	ServiceFeeOffset    float64 `json:"service_fee_offset"`    // [Required] <p>This is the offset to gross service fee, reducing it to the net value.</p>
}

type SellerProof struct {
	SellerProofStatus      string `json:"seller_proof_status"`      // [Required] <p>To indicate whether the seller needs to provide evidence when the return status is RETURN_JUDING, RETURN_SELLER_DISPUTE and RETURN_ACCEPTED. Applicable values: See Data Definition- SellerProofStatus.</p>
	SellerEvidenceDeadline int64  `json:"seller_evidence_deadline"` // [Required] <p>To indicate the deadline for submitting the evidence.</p>
}

type SellerStock struct {
	LocationId string `json:"location_id"` // [Required] <p>location id<br /></p>
	Stock      int64  `json:"stock"`       // [Required] <p>stock<br /></p>
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

type SetAppPushConfigRequest struct {
	CallbackUrl       *string `json:"callback_url,omitempty"`         // [Optional] <p>The callback url of push mechanism. It is the address where the Shopee will send the push message to. If you don't set any callback_url before, this parameters is required.<br /></p>
	SetPushConfigOn   []int64 `json:"set_push_config_on,omitempty"`   // [Optional] <p>Turn on push config, Shopee will send the push message into the callback url.</p><p>1=Shop authorization for partners</p><p>2=Shop deauthorization for partners</p><p>3=Order status update push</p><p>4=TrackingNo push</p><p>5=Shopee Updates</p><p>6=Banned item push</p><p>7=item promotion push</p><p>8=reserved stock change push</p><p>9=promotionn update push</p><p>10=webchat push</p><p>11=video upload push</p><p>12=openapi authorization expiry push</p><p>13=brand register result<br /></p>
	SetPushConfigOff  []int64 `json:"set_push_config_off,omitempty"`  // [Optional] <p>Turn off Push config, Shopee won't send the push message into the callback url.</p><p>1=Shop authorization for partners</p><p>2=Shop deauthorization for partners</p><p>3=Order status update push</p><p>4=TrackingNo push</p><p>5=Shopee Updates</p><p>6=Banned item push</p><p>7=item promotion push</p><p>8=reserved stock change push</p><p>9=promotionn update push</p><p>10=webchat push</p><p>11=video upload push</p><p>12=openapi authorization expiry push</p><p>13=brand register result<br /></p>
	BlockedShopIdList []int64 `json:"blocked_shop_id_list,omitempty"` // [Optional] <p>Use this filed to set shops that need to be blocked.Please input no more than 500 shop id.<br /></p>
}

type SetAppPushConfigResponse struct {
	BaseResponse `json:",inline"`             // Common response fields
	Response     SetAppPushConfigResponseData `json:"response"` // Actual response data
}

type SetAppPushConfigResponseData struct {
	Result string `json:"result"` // [Required] <p>Use this field to indicate whether the configuration is set successfully.<br /></p>
}

type SetItemInstallmentStatusRequest struct {
	ItemIdList           []int64 `json:"item_id_list"`                     // [Required] The id array of the item, Max :100
	TenureList           []int64 `json:"tenure_list"`                      // [Required] <p>Staged array, TH must be [3,6,10], TW region tenures must be in [3,6,12,24], [] means closed</p>
	ParticipatePlanAhora *bool   `json:"participate_plan_ahora,omitempty"` // [Optional] Only applicable and required for local AR sellers.
}

type SetItemInstallmentStatusResponse struct {
	BaseResponse `json:",inline"`                     // Common response fields
	Response     SetItemInstallmentStatusResponseData `json:"response"` // Actual response data
}

type SetItemInstallmentStatusResponseData struct {
	ItemInstallmentList []ItemInstallment `json:"item_installment_list"` // [Required]
	ItemPlanAhoraList   []ItemPlanAhora   `json:"item_plan_ahora_list"`  // [Required] Only applicable for local AR sellers.
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

type SetNoteRequest struct {
	OrderSn string `json:"order_sn"` // [Required] Shopee's unique identifier for an order.
	Note    string `json:"note"`     // [Required] The note seller add for reference.
}

type SetNoteResponse struct {
	BaseResponse `json:",inline"` // Common response fields
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

type SetShopInstallmentStatusRequest struct {
	InstallmentStatus int64 `json:"installment_status"` // [Required] <p>The status of installment contains 1 and 0.</p>
}

type SetShopInstallmentStatusResponse struct {
	BaseResponse `json:",inline"`                     // Common response fields
	Response     SetShopInstallmentStatusResponseData `json:"response"` // Actual response data
}

type SetShopInstallmentStatusResponseData struct {
	InstallmentStatus int64 `json:"installment_status"` // [Required]
}

type SetSipDiscountRequest struct {
	Region          string `json:"region"`            // [Required] <p>The region of SIP affiliate shop that needs to set discount.</p>
	SipDiscountRate int64  `json:"sip_discount_rate"` // [Required] <p>The overall market discount rate that will apply to all items for SIP affiliate shop in current region.</p>
}

type SetSipDiscountResponse struct {
	BaseResponse `json:",inline"`           // Common response fields
	Response     SetSipDiscountResponseData `json:"response"` // Actual response data
}

type SetSipDiscountResponseData struct {
	Region          string `json:"region"`            // [Required] <p>The region of SIP affiliate shop.<br /></p>
	Status          string `json:"status"`            // [Required] <p>The status of discount for SIP affiliate shop in current region, can be upcoming/ongoing, excluding expired discounts.</p>
	SipDiscountRate int64  `json:"sip_discount_rate"` // [Required] <p>The discount rate set for SIP affiliate shop in current region.<br /></p>
	StartTime       int64  `json:"start_time"`        // [Required] <p>The start time of discount for SIP affiliate shop in current region,&nbsp;in UNIX seconds.</p><p><br /></p><p>Note: The start time is 30 minutes after the sellers set up the sip_discount_rate.</p>
	EndTime         int64  `json:"end_time"`          // [Required] <p>The end time of discount for SIP affiliate shop in current region,&nbsp;in UNIX seconds.<br /></p><p><br /></p><p>Note:&nbsp;The end time is 180 days after the start time.</p>
	CreateTime      int64  `json:"create_time"`       // [Required] <p>The create time of discount for SIP affiliate shop in current region,&nbsp;in UNIX seconds.<br /></p>
	UpdateTime      int64  `json:"update_time"`       // [Required] <p>The latest update time of discount for SIP affiliate shop in current region,&nbsp;in UNIX seconds.<br /></p>
}

type SetSyncFieldRequest struct {
	ShopSyncList []ShopSync `json:"shop_sync_list"` // [Required] Length limit is [1,50].
}

type SetSyncFieldResponse struct {
	BaseResponse `json:",inline"` // Common response fields
}

type SharedImageInfo struct {
	ImageId string `json:"image_id"` // [Required] Image id.
}

type SharedItem struct {
	ItemId int64 `json:"item_id"` // [Required] Success item id
	Unlist bool  `json:"unlist"`  // [Required] Whether the item is unlisted
}

type SharedOption struct {
	Option string           `json:"option"`          // [Required] <p>Option name.<br /></p>
	Image  *SharedImageInfo `json:"image,omitempty"` // [Optional] <p>Option image.<br /></p>
}

type SharedOrder struct {
	OrderSn       string `json:"order_sn"`       // [Required] <p>Shopee's unique identifier for an order.<br /></p>
	PackageNumber string `json:"package_number"` // [Required] <p>Shopee's unique identifier for the package under an order.<br /></p>
}

type SharedUploadImageResponse struct {
	BaseResponse `json:",inline"`              // Common response fields
	Response     SharedUploadImageResponseData `json:"response"` // Actual response data
}

type SharedUploadImageResponseData struct {
	ImageInfo     *ResponseDataImageInfo             `json:"image_info"`      // [Required]
	ImageInfoList []UploadImageResponseDataImageInfo `json:"image_info_list"` // [Required]
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

type ShippingProofTemplate struct {
	IsTrackingNumberRequired     bool `json:"is_tracking_number_required"`      // [Required] <p>To indicate whether it is mandatory to provide tracking number when uploading shipping proof.</p>
	IsShippingImageFileMandatory bool `json:"is_shipping_image_file_mandatory"` // [Required] <p>To indicate whether it is mandatory to provide shipping image file(s) when uploading shipping proof.</p>
}

type Shop struct {
	SipAffiShops []ShopSipAffiShops `json:"sip_affi_shops"` // [Required] List of SIP affiliate shops.Only primary shop will return this parameter
}

type ShopCategoryGetItemListRequest struct {
	ShopCategoryId int64  `json:"shop_category_id"`    // [Required] ShopCategory's unique identifier.
	PageSize       *int64 `json:"page_size,omitempty"` // [Optional] Specifies the starting entry of data to return in the current call. Default is 1000. The input range of page_size is [0, 1000]
	PageNo         *int64 `json:"page_no,omitempty"`   // [Optional] If many items are available to retrieve, you may need to call this api multiple times to retrieve all the data. And the default will be 0. page_size*page_no should be [0, 2147483446].
}

type ShopCategoryGetItemListResponse struct {
	BaseResponse `json:",inline"`                    // Common response fields
	Response     ShopCategoryGetItemListResponseData `json:"response"` // Actual response data
}

type ShopCategoryGetItemListResponseData struct {
	ItemList   []int64 `json:"item_list"`   // [Required] A list of Shopee's unique identifiers for items.
	TotalCount int64   `json:"total_count"` // [Required] This is to indicate the whole number of items under the shop category.
	More       bool    `json:"more"`        // [Required] This is to indicate whether the item list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of items.
}

type ShopCategorys struct {
	ShopCategoryId int64  `json:"shop_category_id"` // [Required] ShopCategory's unique identifier.
	Status         int64  `json:"status"`           // [Required] ShopCategory's status. Applicable values--1: 'NORMAL', 2: 'INACTIVE', 0: 'DELETED'
	Name           string `json:"name"`             // [Required] ShopCategory's name.
	SortWeight     int64  `json:"sort_weight"`      // [Required] ShopCategory's sort weight.
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

type ShopPublishableStatus struct {
	Region                string `json:"region"`                  // [Required] <p>Region of published shop.<br /></p>
	ShopPublishableStatus bool   `json:"shop_publishable_status"` // [Required] <p>If the shop is publishable, ture means shop is publishable, fals means shop is unpublishable<br /></p>
	UnpublishableReason   string `json:"unpublishable_reason"`    // [Required] <p>Return the unpublishable reason. If the shop is publishable, will return empty for this field.<br /></p>
}

type ShopSipAffiShops struct {
	AffiShopId int64 `json:"affi_shop_id"` // [Required]  Affiliate shop's id.
}

type ShopSku struct {
	ShopSkuId   string `json:"shop_sku_id"`   // [Required] <p>shop level sku_id  shop_sku_id="item_id" _ "model_id"</p>
	ShopItemId  string `json:"shop_item_id"`  // [Required] <p>shop_item_id="item_id" in Product Module</p>
	ShopModelId string `json:"shop_model_id"` // [Required] <p>shop_model_id= item level model_id</p>
}

type ShopSync struct {
	ShopRegion                 string `json:"shop_region"`                    // [Required] TW TH MY BR IN SG VN
	NameAndDescription         bool   `json:"name_and_description"`           // [Required] sync name and description
	MediaInformation           bool   `json:"media_information"`              // [Required] sync media information
	TierVariationNameAndOption bool   `json:"tier_variation_name_and_option"` // [Required] sync tier variation
	Price                      bool   `json:"price"`                          // [Required] sync price
	DaysToShip                 bool   `json:"days_to_ship"`                   // [Required] sync days to ship info
}

type ShopeeStock struct {
	LocationId string `json:"location_id"` // [Required] <p>location id<br /></p>
	Stock      string `json:"stock"`       // [Required] <p>stock<br /></p>
}

type SipAffiShops struct {
	Region     string `json:"region"`       // [Required] Affiliate Shop's area
	AffiShopId int64  `json:"affi_shop_id"` // [Required] Affiliate shop's id
}

type SipItemPrice struct {
	ModelId      *int64  `json:"model_id,omitempty"` // [Optional] 0 for no model item.
	SipItemPrice float64 `json:"sip_item_price"`     // [Required] SIP item price.
}

type Size struct {
	SizeId       string  `json:"size_id"`       // [Required] The identity of size.
	Name         string  `json:"name"`          // [Required] The name of size.
	DefaultPrice float64 `json:"default_price"` // [Required] The pre-defined shipping fee for the specific size.
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

type Sku struct {
	MtskuId            string    `json:"mtsku_id"`             // [Required] <p>mtsku id</p>
	ModelId            string    `json:"model_id"`             // [Required] <p>Warehouse model SKU ID</p><p>For CB global items, this is equal to the global model_id.</p><p>      </p><p>For local items, it differs from model_id; use shop_model_id to match the model_id</p>
	FulfillMappingMode int64     `json:"fulfill_mapping_mode"` // [Required] <p>0-Null；1-Bundle SKU；2-Parent SKU</p>
	ModelName          string    `json:"model_name"`           // [Required] <p>model name</p>
	NotMovingTag       int64     `json:"not_moving_tag"`       // [Required]
	WhsList            []Whs     `json:"whs_list"`             // [Required] <p>Info of whs</p>
	ShopSkuList        []ShopSku `json:"shop_sku_list"`        // [Required]
}

type SkuWhs struct {
	WhsId            string `json:"whs_id"`             // [Required] <p>warehouse ID</p>
	ExpiringQty      int64  `json:"expiring_qty"`       // [Required] <p>Stocks that are expiring soon and should be sold as soon as possible.<br /></p>
	ExpiredQty       int64  `json:"expired_qty"`        // [Required] <p>Stock past expiry date.<br /></p>
	ExpiryBlockedQty int64  `json:"expiry_blocked_qty"` // [Required] <p>Stocks that are too near to expiry and cannot be sold any more.<br /></p>
	DamagedQty       int64  `json:"damaged_qty"`        // [Required] <p>Stock in damaged condition and cannot be sold.<br /></p>
	NormalQty        int64  `json:"normal_qty"`         // [Required] <p>Stocks that are normal.<br /></p>
	TotalQty         int64  `json:"total_qty"`          // [Required] <p>Total stocks on hand.<br /></p>
}

type Slug struct {
	Slug     string `json:"slug"`      // [Required]  The identity of slug.
	SlugName string `json:"slug_name"` // [Required]  The name of slug.
}

type Sort struct {
	SortType  *int64 `json:"sort_type,omitempty"` // [Optional] <p>Use this field to specify which field to use to sort the returned package list. Available values:&nbsp;</p><p>1: ShipByDate&nbsp;&nbsp;<br /></p><p>2: CreateDate<br /></p><p>3: ConfirmedDate<br /></p><p><br /></p><p><font color="#c24f4a">Default value = 1 (ShipByDate)</font><br /></p>
	Ascending *bool  `json:"ascending,omitempty"` // [Optional] <p>Use this field to specify whether the returned package list is sorted in ascending or descending sort_type.</p><p><font color="#c24f4a"><br /></font></p><p><font color="#c24f4a">Default value = true</font></p>
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

type SplitOrderRequest struct {
	OrderSn     string    `json:"order_sn"`     // [Required] Shopee's unique identifier for an order.
	PackageList []Package `json:"package_list"` // [Required] <p>The list of packages that you want to split.&nbsp;</p><p><br /></p><p>Note:&nbsp;</p><p>- Orders that include installation services cannot be split by quantity.</p><p>- When splitting the order, must contain all items in the order in one request.</p><p>- You can split the order into 30 parcels at most in TW and 5 parcels at most in other regions.</p>
}

type SplitOrderResponse struct {
	BaseResponse `json:",inline"`       // Common response fields
	Response     SplitOrderResponseData `json:"response"` // Actual response data
}

type SplitOrderResponseData struct {
	OrderSn     string                          `json:"order_sn"`     // [Required] Shopee's unique identifier for an order.
	PackageList []SplitOrderResponseDataPackage `json:"package_list"` // [Required] The list of package under this order you have split.
}

type SplitOrderResponseDataPackage struct {
	PackageNumber string        `json:"package_number"` // [Required] Shopee's unique identifier for the package under an order.
	ItemList      []PackageItem `json:"item_list"`      // [Required] The list of items under this package.
}

type SpxReceiveStation struct {
	SpxFirstReceiveStation string `json:"spx_first_receive_station"` // [Required] <p>The first pickup station.<br /></p>
}

type Ssp struct {
	SspInfo   *SspInfo `json:"ssp_info"`  // [Required]
	Available bool     `json:"available"` // [Required] <p>Indicate whether the Shopee Standard Product is available to create product (One&nbsp;SSP can only be used to create one product under the shop, so if the SSP has already been used, then available will be false).</p>
}

type SspAttribute struct {
	AttributeId           int64            `json:"attribute_id"`            // [Required] <p>The Identify of each attribute.</p>
	OriginalAttributeName string           `json:"original_attribute_name"` // [Required] <p>The name of each attribute.</p>
	AttributeValues       []AttributeValue `json:"attribute_values"`        // [Required]
}

type SspInfo struct {
	SspId       int64  `json:"ssp_id"`       // [Required] <p>Shopee's unique identifier for Shopee&nbsp;Standard Product.<br /></p>
	ProductName string `json:"product_name"` // [Required] <p>Name of Shopee&nbsp;Standard Product.<br /></p>
	Gtin        string `json:"gtin"`         // [Required] <p>GTIN of Shopee&nbsp;Standard Product.</p>
	Oem         string `json:"oem"`          // [Required] <p>OEM of Shopee&nbsp;Standard Product.</p>
}

type SspMedia struct {
	Image *ItemImage `json:"image"` // [Required]
	Video []Video    `json:"video"` // [Required]
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

type StartQty struct {
	StartOnHandTotal int64 `json:"start_on_hand_total"` // [Required] <p>sku number at the start time<br /></p>
	StartSellable    int64 `json:"start_sellable"`      // [Required] <p>Number of sellable SKUs at the start time<br /></p>
	StartReserved    int64 `json:"start_reserved"`      // [Required] <p>Number of reserved SKUs at the start time.<br /></p>
	StartUnsellable  int64 `json:"start_unsellable"`    // [Required]
}

type StartSessionRequest struct {
	SessionId int64 `json:"session_id"` // [Required] <p>The identifier of livestream session.</p>
	DomainId  int64 `json:"domain_id"`  // [Required] <p>The identifier of the stream domain.</p>
}

type StartSessionResponse struct {
	BaseResponse `json:",inline"` // Common response fields
	Response     interface{}      `json:"response"` // Actual response data
}

type StatusInfoTag struct {
	TagId int64 `json:"tag_id"` // [Required] <p>Shipping urgency tag type, applicable values below:</p><p>0: No tag</p><p>1: Will be cancelled within 1 day</p><p>2: Must ship before the specified timestamp</p><p>3: Shipment delayed</p><p>4: Must ship within the current hour</p><p>5: Will be cancelled at the specified timestamp</p>
}

type Stock struct {
	ModelId     *int64        `json:"model_id,omitempty"` // [Optional] 0 for no model item.
	SellerStock []SellerStock `json:"seller_stock"`       // [Required] <p>new stock info（Please notice that stock(including Seller Stock and Shopee Stock) should be larger than or equal to real-time reserved stock）<br /></p>
}

type StockInfo struct {
	StockType       int64  `json:"stock_type"`        // [Required] The stock type.
	StockLocationId string `json:"stock_location_id"` // [Required] location_id of the stock.
	NormalStock     int64  `json:"normal_stock"`      // [Required] The normal stock quantity of the variation in the listing currency.
	ReservedStock   int64  `json:"reserved_stock"`    // [Required] The reserved stock quantity of the variation in the listing currency.
}

type StockInfoV2 struct {
	SummaryInfo  *SummaryInfo             `json:"summary_info"`  // [Required] <p>stock summary info<br /></p>
	SellerStock  []StockInfoV2SellerStock `json:"seller_stock"`  // [Required] <p>seller stock<br /></p>
	ShopeeStock  []SellerStock            `json:"shopee_stock"`  // [Required] <p>shopee stock<br /></p>
	AdvanceStock *AdvanceStock            `json:"advance_stock"` // [Required] <p>Only for PH/VN/ID/MY local selected shops.</p>
}

type StockInfoV2SellerStock struct {
	LocationId string `json:"location_id"` // [Required] <p>location id<br /></p>
	Stock      int64  `json:"stock"`       // [Required] <p>stock in the current warehouse<br /></p>
	IfSaleable bool   `json:"if_saleable"` // [Required] <p>To return if the stock of the location id is saleable<br /></p>
}

type StreamUrlList struct {
	PushUrl  string `json:"push_url"`  // [Required] <p>The push stream url for the livestream session.</p>
	PushKey  string `json:"push_key"`  // [Required] <p>The push stream key for the livestream session.</p>
	PlayUrl  string `json:"play_url"`  // [Required] <p>The pull stream url of the livestream session.</p>
	DomainId int64  `json:"domain_id"` // [Required] <p>The identifier of the stream domain, need to be passed in request for v2.livestream.start_session.</p>
}

type Style struct {
	TextStyle  []string `json:"text_style,omitempty"`  // [Optional] <p>supports bold and italic<br /></p>
	FontSize   *int64   `json:"font_size,omitempty"`   // [Optional] <p>supports 1 - 108<br /></p>
	TextColor  *string  `json:"text_color,omitempty"`  // [Optional] <p>color string like "#AbCd12"<br /></p>
	ImageWidth *float64 `json:"image_width,omitempty"` // [Optional] <p>supports 0.1-30, in centimeters<br /></p>
	HAlign     *string  `json:"h_align,omitempty"`     // [Optional] <p>text horizontal align, supports left, center and right.<br /></p>
}

type SubItem struct {
	ItemId            *int64   `json:"item_id,omitempty"`              // [Optional] Shopee's unique identifier for an item.
	ModelId           *int64   `json:"model_id,omitempty"`             // [Optional] Shopee's unique identifier for a model.
	Status            *int64   `json:"status,omitempty"`               // [Optional] The status of add on deal item：enable = 1；disable =2
	SubItemInputPrice *float64 `json:"sub_item_input_price,omitempty"` // [Optional] Add-on discount price before tax
	SubItemLimit      *int64   `json:"sub_item_limit,omitempty"`       // [Optional] The purchase limit of sub item.The purchase limit of each sub item. Only the add on discount can be set and the default limit of gift with mini.spend is 1
}

type SubItemPrice struct {
	PromoInputPrice float64 `json:"promo_input_price"` // [Required] Add-on discount price before tax
	PromoPrice      float64 `json:"promo_price"`       // [Required] <p>Add-on discount price after tax<br /></p>
}

type Success struct {
	ModelId    int64  `json:"model_id"`    // [Required] ID of model.
	LocationId string `json:"location_id"` // [Required] <p>location id; This field and the stock field are returned in pairs<br /></p>
	Stock      int64  `json:"stock"`       // [Required] <p>stock;This field is returned if seller stock is used in the request, and normal stock fields are not returned.<br /></p>
}

type SuccessItem struct {
	ItemId         int64            `json:"item_id"`         // [Required] <p>Shopee's unique identifier for an item.<br /></p>
	QualityLevel   int64            `json:"quality_level"`   // [Required] <p>Item's latest content quality level. Applicable values:<br />0: NONE (No quality level for item in SELLER_DELETE / SHOPEE_DELETE / BANNED status)<br />1: TO_BE_IMPROVED<br />2: QUALIFIED<br />3: EXCELLENT<br /></p>
	UnfinishedTask []UnfinishedTask `json:"unfinished_task"` // [Required]
}

type SuccessList struct {
	ItemIdList []int64 `json:"item_id_list"` // [Required] Success item ID.
}

type SuggestedCategory struct {
	CategoryId   int64  `json:"category_id"`   // [Required] <p>ID for Shopee suggested category.</p>
	CategoryName string `json:"category_name"` // [Required] <p>Default name for Shopee suggested category.<br /></p>
}

type SuggestedKeywords struct {
	Keyword      string  `json:"keyword"`       // [Required] <p>Keyword value</p><p>(Only return the highly recommended keywords, will be sightly different from Seller Center)</p>
	QualityScore int64   `json:"quality_score"` // [Required] <p>This is a measure of how attractive your ad is and its relevance to the keyword. The higher the quality score, the higher your ad rank. Ad rank is based on this score and your bid price.<br /></p>
	SearchVolume int64   `json:"search_volume"` // [Required] <p>The number of times the keyword has been searched on Shopee in the last 30 days. The larger the search volume, the more impressions your ad will receive.<br /></p>
	SuggestedBid float64 `json:"suggested_bid"` // [Required] <p>This is bid price suggested by Shopee algorithm for the keyword in local currency.<br /></p>
}

type SummaryInfo struct {
	TotalReservedStock  int64 `json:"total_reserved_stock"`  // [Required] <p>Stock reserved for promotion.<br /></p><p><br /></p><p>Note: For SIP P Item, will return the total reserved stock for P Item and all A Items under the P Item.</p>
	TotalAvailableStock int64 `json:"total_available_stock"` // [Required] <p>Stock can be sold currently<br /></p>
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

type SyncSetting struct {
	AutoSyncDts bool `json:"auto_sync_dts"` // [Required] <p>Auto sync the pre_order setting from main component or not.<br /></p>
}

type Tag struct {
	Kit bool `json:"kit"` // [Required] <p>Indicate if the item is kit item.</p>
}

type Target struct {
	Value      float64 `json:"value"`      // [Required] <p>Value of target.</p>
	Comparator string  `json:"comparator"` // [Required] <p>Comparator of target: &lt;, &lt;=, &gt;, &gt;=, =</p>
}

type TaxInfo struct {
	Ncm               *string        `json:"ncm,omitempty"`                 // [Optional] <p>Mercosur Common Nomenclature, it is a convention between Mercosur member countries to easily recognize goods, services and productive factors negotiated among themselves.&nbsp;(BR region)<br /></p><p><br />NCM must have 8 digits, OR, if your item doesn't have a NCM enter the value "00"<br /></p>
	SameStateCfop     *string        `json:"same_state_cfop,omitempty"`     // [Optional] Tax Code of Operations and Installments for orders that seller and buyer are in the same state. It identifies a specific operation by category at the time of issuing the invoice.(BR region)
	DiffStateCfop     *string        `json:"diff_state_cfop,omitempty"`     // [Optional] Tax Code of Operations and Installments for orders that seller and buyer are in different states. It identifies a specific operation by category at the time of issuing the invoice.(BR region)
	Csosn             *string        `json:"csosn,omitempty"`               // [Optional] Code of Operation Status – Simples Nacional, code for company operations to identify the origin of the goods and the taxation regime of the operations.(BR region)
	Origin            *string        `json:"origin,omitempty"`              // [Optional] <p>Product source, domestic or foreig (BR region).</p><p>|0 - National, except for those indicated in codes 3, 4, 5, and 8|<br /> |1 - Foreign: Direct import, except for that indicated in code 6|<br /> |2 - Foreign: Acquired in the domestic market, except for that indicated in code 7|<br /> |3 - National: Goods or products with Import Content greater than 40% and less than or equal to 70%|<br /> |4 - National: Produced in compliance with the basic production processes outlined in the legislations cited in the Agreements|<br /> |5 - National: Goods or products with Import Content less than or equal to 40%|<br /> |6 - Foreign: Direct import, without a national equivalent, listed by CAMEX and natural gas|<br /> |7 - Foreign: Acquired in the domestic market, without a national equivalent, listed by CAMEX and natural gas|<br /> |8 - National: Goods or products with Import Content greater than 70%|</p>
	Cest              *string        `json:"cest,omitempty"`                // [Optional] <p>Tax Replacement Specifying Code (CEST), to separate within the same NCM products that do or do not have ICMS tax substitution. (BR region)<br /><br />CEST must have 7 digits, OR, if your item doesn't have a CEST enter the value "00".<br /></p>
	MeasureUnit       *string        `json:"measure_unit,omitempty"`        // [Optional] (BR region)
	TaxType           *TaxType       `json:"tax_type,omitempty"`            // [Optional] <p>tax_type only for TW whitelist shop. Shopee will referred Tax type when substitute sellers for issuing e-receipts to buyers. All variations share the same tax type. The meaning of value:&nbsp;</p><p>0: no tax type</p><p>1: tax-able</p><p>2: tax-free<br /></p>
	Pis               *string        `json:"pis,omitempty"`                 // [Optional] <p>Only for BR shop.</p><p>PIS - Programa de Integração Social (Social Integration Program). It is a government tax to collect resources for the payment of unemployment insurance and other employee related rights.</p><p>PIS % - the tax applied to this product</p>
	Cofins            *string        `json:"cofins,omitempty"`              // [Optional] <p>Only for BR shop.<br /></p><p>COFINS – Contribuição para Financiamento da Seguridade Social (Contribution for Social Security Funding). It&nbsp;is a government tax to collect resources for public health system and social security.</p><p>COFINS&nbsp;% - the tax applied to this product</p>
	IcmsCst           *string        `json:"icms_cst,omitempty"`            // [Optional] <p>Only for BR shop.<br /></p><p>ICMS - Imposto sobre Circulação de Mercadorias e Serviços (Circulation of Goods and Services Tax).&nbsp;</p><p>CST - Código da Situação Tributária (Tax Situation Code) is represented by a combination of 3 numbers with the purpose of demonstrating the origin of a product and determining the form of taxation that will apply to it. Therefore, each digit in the CST Table has a specific meaning: the first digit indicates the origin of the operation, the second digit represents the ICMS taxation on the operation and the third digit provides additional information about the form of taxation.</p>
	PisCofinsCst      *string        `json:"pis_cofins_cst,omitempty"`      // [Optional] <p>Only for BR shop.</p><p>The CST PIS/Cofins is a code on the Electronic Invoice (NF-e) that identifies the tax situation of PIS (Programa de Integração Social) and Cofins (Contribuição para o Financiamento da Seguridade Social) in sales of goods.</p>
	FederalStateTaxes *string        `json:"federal_state_taxes,omitempty"` // [Optional] <p>Only for BR shop.</p><p>Enter the total percentage of the combination of federal, state, and municipal taxes, using up to two decimals.<br /></p>
	OperationType     *OperationType `json:"operation_type,omitempty"`      // [Optional] <p>Only for BR shop.</p><p>1: Retailer</p><p>2: Manufacturer</p>
	ExTipi            *string        `json:"ex_tipi,omitempty"`             // [Optional] <p>Only for BR shop.<br /></p><p>The EXTIPI field in the NF-e (Nota Fiscal Eletrônica) is used to indicate if there's an exception to the IPI (Imposto sobre Produtos Industrializados) tax rate for a specific product.<br /></p>
	FciNum            *string        `json:"fci_num,omitempty"`             // [Optional] <p>Only for BR shop.<br /></p><p>The FCI Control Number is a unique identifier assigned to each import FCI (Import Content Form). It's mandatory on the corresponding NF-e (electronic invoice) to ensure compliance with Brazilian import tax regulations.<br /></p>
	RecopiNum         *string        `json:"recopi_num,omitempty"`          // [Optional] <p>Only for BR shop.</p><p>RECOPI NACIONAL is a Brazilian government system that facilitates the registration and management of tax-exempt operations involving paper destined for printing books, newspapers, and periodicals (known as "papel imune" in Portuguese).</p>
	AdditionalInfo    *string        `json:"additional_info,omitempty"`     // [Optional] <p>Only for BR shop.</p><p>Include relevant information to display on Invoice.</p>
	GroupItemInfo     *GroupItemInfo `json:"group_item_info,omitempty"`     // [Optional] <p>Only for BR shop.</p><p>Required if the item is a group item.</p>
	ExportCfop        *string        `json:"export_cfop,omitempty"`         // [Optional] <p>[BR region]</p><p>7101 - for sales of self-produced goods</p><p>7102 - resale of third-party goods</p>
}

type TenureInfoList struct {
	PaymentChannelName string `json:"payment_channel_name"` // [Required] <p>Name of the payment channel that buyer used in checkout.</p>
	InstalmentPlan     string `json:"instalment_plan"`      // [Required] <p>Tenure information. This will have value if payment channel used has tenure information, such as credit card.</p>
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

type ThumbnailUrl struct {
	ImageUrlRegion string `json:"image_url_region"` // [Required] <p>Region of image url<br /></p>
	ImageUrl       string `json:"image_url"`        // [Required] <p>image url<br /></p>
}

type TierVariation struct {
	Name       *string        `json:"name,omitempty"` // [Optional] <p>Tier variation name.<br /></p>
	OptionList []SharedOption `json:"option_list"`    // [Required] <p>Tier variation option info list.<br /></p>
}

type TierVariationOption struct {
	Image  *OptionImage `json:"image"`  // [Required] Image of this option
	Option string       `json:"option"` // [Required] Option name
}

type TimeSlot struct {
	Date         int64    `json:"date"`           // [Required] <p>The date of pickup time. In timestamp.<br /></p>
	TimeText     string   `json:"time_text"`      // [Required] <p>The text description of pickup time. Only applicable for certain channels.<br /></p>
	PickupTimeId string   `json:"pickup_time_id"` // [Required] <p>The identity of pickuptime.<br /></p>
	Flags        []string `json:"flags"`          // [Required] <p>This field will have the value <b><font color="#c24f4a">“recommended”</font></b> for the time slot that Shopee suggests sellers choose.&nbsp;</p><p><br /></p><p>While it is advisable for sellers to choose the recommended time slot, they can also choose other time slots that do not have the recommended flag.</p>
}

type TotalIncome struct {
	PendingAmount   float64 `json:"pending_amount"`    // [Required] <p>Total amount pending release (Local: orders before ESCROW_PAID; CB: orders before ESCROW_PAYOUT).</p><p>&lt;path&gt;&lt;/path&gt;<br /></p>
	ToReleaseAmount float64 `json:"to_release_amount"` // [Required] <p>Amount queued for release in the next payout cycle (CB only). Not applicable for Local shops.</p><p>&lt;path&gt;&lt;/path&gt;<br /></p>
	ReleasedAmount  float64 `json:"released_amount"`   // [Required] <p>Total amount successfully released to the seller.</p><p>&lt;path&gt;&lt;/path&gt;<br /></p>
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

type Transaction struct {
	Status             string     `json:"status"`               // [Required] The status of the transaction，available values: FAILED,COMPLETED,PENDING,INITIAL.
	TransactionType    string     `json:"transaction_type"`     // [Required] The type of transaction.
	Amount             float64    `json:"amount"`               // [Required] The amount of transaction.
	CurrentBalance     float64    `json:"current_balance"`      // [Required] The current balance of this account.
	CreateTime         int64      `json:"create_time"`          // [Required] The create time of the transaction.
	OrderSn            string     `json:"order_sn"`             // [Required] Shopee's unique identifier for an order.
	RefundSn           string     `json:"refund_sn"`            // [Required] The serial number of return.
	WithdrawalType     string     `json:"withdrawal_type"`      // [Required] The type of withdrawal.
	TransactionFee     float64    `json:"transaction_fee"`      // [Required] This field indicates the transaction fee.
	Description        string     `json:"description"`          // [Required] The detailed description of TOPUP SUCCESS and TOPUP FAILED.
	BuyerName          string     `json:"buyer_name"`           // [Required] The name of buyer.
	PayOrderList       []PayOrder `json:"pay_order_list"`       // [Required]
	ShopName           string     `json:"shop_name"`            // [Required] Name of the shop.
	WithdrawalId       int64      `json:"withdrawal_id"`        // [Required] <p>Withdrawal ID when transaction type is withdraw_created, withdrawal_completed, withdrawal_cancelled.</p>
	Reason             string     `json:"reason"`               // [Required] The reason for ADJUSTMENT_ADD and ADJUSTMENT_MINUS.
	RootWithdrawalId   int64      `json:"root_withdrawal_id"`   // [Required] Use this field to indicate the event where a withdrawal is split into several withdrawals due to the withdrawal limit.
	TransactionTabType string     `json:"transaction_tab_type"` // [Required] <p><b>Description:</b></p><p>A new response parameter added after:&nbsp;https://confluence.shopee.io/display/SPCT/%5BPRD%5D+%5BOpen+API%5D+Update+on+New+Open+API+to+fetch+Seller+wallet+Transaction&nbsp;<br /><br />This returns the updated transaction tab types that client can use to specify which transaction types they want to return. It will have the following tab types<br /><br />Default<br />wallet_order_income<br />wallet_adjustment_filter</p><p>wallet_wallet_payment</p><p>wallet_refund_from_order</p><p>wallet_withdrawals</p><p>fast_escrow_repayment</p><p>fast_pay</p><p>seller_loan<br />corporate_loan<br />pix_transactions_filter</p><p>open_finance_transactions_filter&nbsp;<br /><br />Note for BR, currently in SOP live configuration, wallet txn type that linked to&nbsp;pix_transactions_filter&nbsp; and&nbsp;open_finance_transactions_filter&nbsp;&nbsp;are classified as&nbsp;default&nbsp;&nbsp;type tab instead. therefore for Open API client who wants to query these 2 txn can put default as the filter in this type</p>
	MoneyFlow          string     `json:"money_flow"`           // [Required] <p>New response parameter provided after:&nbsp;https://confluence.shopee.io/display/SPCT/%5BPRD%5D+%5BOpen+API%5D+Update+on+New+Open+API+to+fetch+Seller+wallet+Transaction&nbsp;<br /><br />It's to indicate the money flow</p><p>MONEY_IN = addition&nbsp;<br />MONEY_OUT = deduction</p><p>if not specified in request, will return both</p><p>Note special case for TW JKO Pay, we will ignore Money_flow&nbsp;</p>
	OutletShopName     string     `json:"outlet_shop_name"`     // [Required] <p>The outlet shop name where this outlet transaction came from.&nbsp;(In the Original Instant Mart concept, outlet transactions are redirected to Mart.)</p>
}

type Transactions struct {
	Amount                   float64 `json:"amount"`                     // [Required] <p>each transaction's amount<br /></p>
	Currency                 string  `json:"currency"`                   // [Required] <p>transaction currency<br /></p>
	OrderSn                  string  `json:"order_sn"`                   // [Required] <p>transaction currency<br /></p>
	CostHeader               string  `json:"cost_header"`                // [Required] <p>transaction belongs to which type<br /></p>
	Scenario                 string  `json:"scenario"`                   // [Required] <p>transaction detailed scenarios<br /></p>
	Remark                   string  `json:"remark"`                     // [Required] <p>detailed description for this transactions<br /></p>
	Level                    string  `json:"level"`                      // [Required] <p>To define this transaction happen at order level or shop level. e.g. shop level adjustment<br /></p>
	BillingTransactionType   string  `json:"billing_transaction_type"`   // [Required] <p>could be Escrow (Order Income) or Adjustment (for this order)<br /></p>
	BillingTransactionStatus string  `json:"billing_transaction_status"` // [Required] <p>Will be either "To Release" or "Released".<br /></p>
}

type TransitWarehouse struct {
	WarehouseId     string `json:"warehouse_id"`      // [Required] <p>The identity of transit warehouse.<br /></p>
	WarehouseNameEn string `json:"warehouse_name_en"` // [Required] <p>The name of transit warehouse in English.<br /></p>
	WarehouseNameCn string `json:"warehouse_name_cn"` // [Required] <p>The name of transit warehouse in Chinese.<br /></p>
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

type UnbanUserCommentRequest struct {
	SessionId   int64 `json:"session_id"`    // [Required] <p>The identifier of livestream session.</p>
	UnbanUserId int64 `json:"unban_user_id"` // [Required] <p>The user ID that will be unbanned from posting comments.</p>
}

type UnbanUserCommentResponse struct {
	BaseResponse `json:",inline"` // Common response fields
	Response     interface{}      `json:"response"` // Actual response data
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

type UnfinishedTask struct {
	IssueType  int64  `json:"issue_type"` // [Required] <p>Item's content issue. Applicable values:<br /></p><p>1: TOO_FEW_IMAGES<br />2: WRONG_CATEGORY<br />3: TOO_FEW_ATTRIBUTES_FOR_QUALIFIED<br />4: LACK_OF_SIZE_CHART<br />5: LACK_OF_STANDARD_VARIATION<br />6: LACK_BRAND<br />7: TOO_SHORT_DESCRIPTION<br />8: TOO_SHORT_OR_TOO_LONG_NAME<br />9: WRONG_WEIGHT<br />10: LACK_OF_VIDEO<br />11: TOO_FEW_ATTRIBUTES_FOR_EXCELLENT</p>
	Suggestion string `json:"suggestion"` // [Required] <p>System&nbsp;suggestion for item's content issue. Applicable values:</p><p>Add at least 3 images</p><p>Adopt suggested category<br />Add at least 1 attributes<br />Add size chart<br />Adopt the color or size variation<br />Add brand info<br />Add at least 100 characters or 1 image for desc<br />Add characters for name to 25~100<br />Adopt suggested weight<br />Add video<br />Add at least 3 attributes</p>
}

type UnlinkSspRequest struct {
	ItemId int64 `json:"item_id"` // [Required] <p>ID of this item.<br /></p>
}

type UnlinkSspResponse struct {
	BaseResponse `json:",inline"`      // Common response fields
	Response     UnlinkSspResponseData `json:"response"` // Actual response data
}

type UnlinkSspResponseData struct {
	ItemId int64 `json:"item_id"` // [Required] <p>Shopee's unique identifier for an item.<br /></p>
}

type UnlistItemRequest struct {
	ItemList []SharedItem `json:"item_list"` // [Required] Length should be between 1 to 50.
}

type UnlistItemResponse struct {
	BaseResponse `json:",inline"`       // Common response fields
	Response     UnlistItemResponseData `json:"response"` // Actual response data
}

type UnlistItemResponseData struct {
	FailureList []Failure    `json:"failure_list"` // [Required]
	SuccessList []SharedItem `json:"success_list"` // [Required]
}

type UnpackagedSkuRequests struct {
	UnpackagedSkuId *string `json:"unpackaged_sku_id,omitempty"` // [Optional] <p>Unpackaged SKU ID of the model.</p>
	Quantity        *int64  `json:"quantity,omitempty"`          // [Optional] <p>Number of copies for the generated labels (maximum 600 total across all requested SKUs).</p>
}

type UnqualifiedConditions struct {
	UnqualifiedCode int64  `json:"unqualified_code"` // [Required]
	UnqualifiedMsg  string `json:"unqualified_msg"`  // [Required]
}

type UnsplitOrderRequest struct {
	OrderSn string `json:"order_sn"` // [Required] Shopee's unique identifier for an order.
}

type UnsplitOrderResponse struct {
	BaseResponse `json:",inline"` // Common response fields
}

type UnsupportWarehouse struct {
	WarehouseId   int64  `json:"warehouse_id"`   // [Required] <p>Unsupported warehouse ID<br /></p>
	WarehouseName string `json:"warehouse_name"` // [Required] <p>Unsupported warehouse name<br /></p>
}

type UpdateAddOnDealMainItemRequest struct {
	AddOnDealId  int64                          `json:"add_on_deal_id"` // [Required] Shopee's unique identifier for add on deal activity.
	MainItemList []AddBundleDealItemRequestItem `json:"main_item_list"` // [Required] The main items added in this add on deal promotion.
}

type UpdateAddOnDealMainItemResponse struct {
	BaseResponse `json:",inline"`                    // Common response fields
	Response     UpdateAddOnDealMainItemResponseData `json:"response"`                 // Actual response data
	AddOnDealId  int64                               `json:"add_on_deal_id,omitempty"` // Shopee's unique identifier for add on deal activity.
}

type UpdateAddOnDealMainItemResponseData struct {
	MainItemList []AddBundleDealItemRequestItem `json:"main_item_list"` // [Required] The main items added in this add on deal promotion.
}

type UpdateAddOnDealRequest struct {
	AddOnDealId            int64    `json:"add_on_deal_id"`                     // [Required] Shopee's unique identifier for an add on deal activity.
	StartTime              *int64   `json:"start_time,omitempty"`               // [Optional] The time when bundle deal activity start.The start time must be 1 hour than current time.
	EndTime                *int64   `json:"end_time,omitempty"`                 // [Optional] The time when bundle deal activity end. The end time must be later than start time.
	PurchaseMinSpend       *float64 `json:"purchase_min_spend,omitempty"`       // [Optional] The minimum purchase amount that needs to be met to buy the gift with min.Spend
	PerGiftNum             *int64   `json:"per_gift_num,omitempty"`             // [Optional] Number of gifts that buyers can get
	PromotionPurchaseLimit *int64   `json:"promotion_purchase_limit,omitempty"` // [Optional] Max. number of add-on products that a customer can purchase per order.
	SubItemPriority        []int64  `json:"sub_item_priority,omitempty"`        // [Optional] The order of sub item
	AddOnDealName          *string  `json:"add_on_deal_name,omitempty"`         // [Optional] Title of the add on deal
}

type UpdateAddOnDealResponse struct {
	BaseResponse `json:",inline"`            // Common response fields
	Response     UpdateAddOnDealResponseData `json:"response"` // Actual response data
}

type UpdateAddOnDealResponseData struct {
	StartTime              int64   `json:"start_time"`               // [Required] The time when add on deal activity start.
	EndTime                int64   `json:"end_time"`                 // [Required] The time when add on deal activity end
	PromotionType          int64   `json:"promotion_type"`           // [Required] The type of add on deal：add on discount =0；gift with mini spend=1
	PurchaseMinSpend       float64 `json:"purchase_min_spend"`       // [Required] The minimum purchase amount that needs to be met to buy the gift with min.Spend
	AddOnDealId            int64   `json:"add_on_deal_id"`           // [Required] Shopee's unique identifier for an add on deal activity.
	PerGiftNum             int64   `json:"per_gift_num"`             // [Required] Number of gifts that buyers can get
	PromotionPurchaseLimit int64   `json:"promotion_purchase_limit"` // [Required] Max. number of add-on products that a customer can purchase per order.
	AddOnDealName          string  `json:"add_on_deal_name"`         // [Required] Title of the add on deal
}

type UpdateAddOnDealSubItemRequest struct {
	AddOnDealId int64     `json:"add_on_deal_id"` // [Required] Shopee's unique identifier for add on deal activity.
	SubItemList []SubItem `json:"sub_item_list"`  // [Required] The sub items added in this add on deal promotion.
}

type UpdateAddOnDealSubItemResponse struct {
	BaseResponse `json:",inline"`                   // Common response fields
	Response     UpdateAddOnDealSubItemResponseData `json:"response"`                 // Actual response data
	AddOnDealId  int64                              `json:"add_on_deal_id,omitempty"` // Shopee's unique identifier for add on deal activity.
}

type UpdateAddOnDealSubItemResponseData struct {
	SubItemList []UpdateAddOnDealSubItemResponseDataSubItem `json:"sub_item_list"` // [Required] The sub items added in this add on deal promotion.
}

type UpdateAddOnDealSubItemResponseDataSubItem struct {
	ItemId            int64   `json:"item_id"`              // [Required] Shopee's unique identifier for an item.
	Status            int64   `json:"status"`               // [Required] The status of add on deal item：enable = 1；disable =2
	ModelId           int64   `json:"model_id"`             // [Required] Shopee's unique identifier for a model.
	FailError         string  `json:"fail_error"`           // [Required]
	FailMessage       string  `json:"fail_message"`         // [Required]
	SubItemInputPrice float64 `json:"sub_item_input_price"` // [Required] The discounted price of sub item
	SubItemLimit      int64   `json:"sub_item_limit"`       // [Required] The purchase limit of sub item.The purchase limit of each sub item. Only the add on discount can be set and the default limit of gift with mini.spend is 1
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

type UpdateBundleDealItemRequest struct {
	BundleDealId int64                          `json:"bundle_deal_id"` // [Required] Shopee's unique identifier for a bundle deal activity.
	ItemList     []AddBundleDealItemRequestItem `json:"item_list"`      // [Required] The items added in this bundle deal promotion.
}

type UpdateBundleDealItemResponse struct {
	BaseResponse `json:",inline"`                 // Common response fields
	Response     UpdateBundleDealItemResponseData `json:"response"` // Actual response data
}

type UpdateBundleDealItemResponseData struct {
	FailedList  []ResponseDataFailed `json:"failed_list"`  // [Required] Indicate error details.
	SuccessList []interface{}        `json:"success_list"` // [Required] The list of succeed added items
}

type UpdateBundleDealRequest struct {
	RuleType           *int64           `json:"rule_type,omitempty"`           // [Optional] The bundle deal rule type：FIX_PRICE = 1 ；DISCOUNT_PERCENTAGE = 2； DISCOUNT_VALUE = 3
	DiscountValue      *float64         `json:"discount_value,omitempty"`      // [Optional]  The deducted price when when buying a bundle deal. Need to input it when the bundle deal rule type is 3
	FixPrice           *float64         `json:"fix_price,omitempty"`           // [Optional] The amount of the buyer needs to spend to purchase a bundle deal.Need to input it when the bundle deal rule type is 1
	DiscountPercentage *int64           `json:"discount_percentage,omitempty"` // [Optional] The discount that the buyer can get when buying a bundle deal. Need to input it when the bundle deal rule type is 2
	MinAmount          *int64           `json:"min_amount,omitempty"`          // [Optional] The quantity of items that need buyer to combine purchased
	StartTime          *int64           `json:"start_time,omitempty"`          // [Optional] The time when bundle deal activity start.The start time must be later than current time.
	EndTime            *int64           `json:"end_time,omitempty"`            // [Optional] The time when bundle deal activity end. The end time must be later than current time.
	Name               *string          `json:"name,omitempty"`                // [Optional] Title of the bundle deal
	PurchaseLimit      *int64           `json:"purchase_limit,omitempty"`      // [Optional] Maximum number of bundle deals that can be bought by a buyer.
	BundleDealId       int64            `json:"bundle_deal_id"`                // [Required] Shopee's unique identifier for a bundle deal activity.
	AdditionalTiers    *AdditionalTiers `json:"additional_tiers,omitempty"`    // [Optional] <p>Use to create tiered discount for bundle deals, a max of 2 additional tiers are allowed to create bundle deals like buy 2 get 10% off, buy 3 for 15% off, buy 4 for 20% off;&nbsp;For each tier, we will need to set the following 4 values based on bundle deal type&nbsp;+&nbsp; &nbsp; min_amount = IntAttribute()&nbsp;+&nbsp;&nbsp;&nbsp; fix_price = IntAttribute()&nbsp;+&nbsp;&nbsp;&nbsp; discount_percentage = IntAttribute()&nbsp;+&nbsp;&nbsp;&nbsp; discount_value = IntAttribute()Note: for additional tiers, the fix price, discount_percentage, discount_value should be consistent with tier 1<br /></p>
}

type UpdateBundleDealResponse struct {
	BaseResponse `json:",inline"`             // Common response fields
	Response     UpdateBundleDealResponseData `json:"response"` // Actual response data
}

type UpdateBundleDealResponseData struct {
	BundleDealId   int64           `json:"bundle_deal_id"`   // [Required] Shopee's unique identifier for a bundle deal activity.
	Name           string          `json:"name"`             // [Required] Title of the bundle deal
	StartTime      int64           `json:"start_time"`       // [Required] The time when bundle deal activity start.
	EndTime        int64           `json:"end_time"`         // [Required] The time when bundle deal activity end.
	BundleDealRule *BundleDealRule `json:"bundle_deal_rule"` // [Required]
	PurchaseLimit  int64           `json:"purchase_limit"`   // [Required] Maximum number of bundle deals that can be bought by a buyer.
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

type UpdateDiscountItemRequest struct {
	DiscountId int64                           `json:"discount_id"` // [Required] Shopee's unique identifier for a discount activity.
	ItemList   []UpdateDiscountItemRequestItem `json:"item_list"`   // [Required] The items selected to this discount. You can update at most 50 items per call.
}

type UpdateDiscountItemRequestItem struct {
	ItemId             int64                                `json:"item_id"`                        // [Required] Shopee's unique identifier for an item.
	ItemPromotionPrice *float64                             `json:"item_promotion_price,omitempty"` // [Optional] The discount price of the item.
	ModelList          []UpdateDiscountItemRequestItemModel `json:"model_list,omitempty"`           // [Optional] The models selected to this discount.
	PurchaseLimit      *int64                               `json:"purchase_limit,omitempty"`       // [Optional] The max number of this product in the promotion price.
}

type UpdateDiscountItemRequestItemModel struct {
	ModelId             int64   `json:"model_id"`              // [Required] Shopee's unique identifier for a variation of an item. If there is no variation of this item, you don't need to input this param. Dafault is 0.
	ModelPromotionPrice float64 `json:"model_promotion_price"` // [Required] The discount price of the item.
}

type UpdateDiscountItemResponse struct {
	BaseResponse `json:",inline"`               // Common response fields
	Response     UpdateDiscountItemResponseData `json:"response"` // Actual response data
}

type UpdateDiscountItemResponseData struct {
	DiscountId int64   `json:"discount_id"` // [Required] Shopee's unique identifier for a discount activity.
	Count      int64   `json:"count"`       // [Required] The number of items that modify successfully.
	ErrorList  []Error `json:"error_list"`  // [Required] Error list of this discount.
}

type UpdateDiscountRequest struct {
	DiscountId   int64   `json:"discount_id"`             // [Required] Shopee's unique identifier for a discount activity.
	DiscountName *string `json:"discount_name,omitempty"` // [Optional] Title of the discount.
	EndTime      *int64  `json:"end_time,omitempty"`      // [Optional] The time when discount activity end. The end time must be 1 hour later than start time.
	StartTime    *int64  `json:"start_time,omitempty"`    // [Optional] The time when discount activity start. The new start time must later than original start time.
}

type UpdateDiscountResponse struct {
	BaseResponse `json:",inline"`           // Common response fields
	Response     UpdateDiscountResponseData `json:"response"` // Actual response data
}

type UpdateDiscountResponseData struct {
	DiscountId int64 `json:"discount_id"` // [Required] Shopee's unique identifier for a discount activity.
	ModifyTime int64 `json:"modify_time"` // [Required] The time when discount is updated.
}

type UpdateFollowPrizeRequest struct {
	FollowPrizeName *string  `json:"follow_prize_name,omitempty"` // [Optional] <p>The name of the follow prize,The follow prize name length max limit is 20.<br /></p>
	CampaignId      int64    `json:"campaign_id"`                 // [Required] <p>The unique identifier for the created follow prize.<br /></p>
	StartTime       *int64   `json:"start_time,omitempty"`        // [Optional] <p>The timing from when the follow prize is valid,the start time later than the current time.If the start time and end time passed in by the seller overlap with other upcoming/ongoing activities, it will prompt "Another Follow Prize voucher already exists during this time period, please set another period."<br /></p>
	EndTime         *int64   `json:"end_time,omitempty"`          // [Optional] <p>The timing until when the follow prize is still valid,the end time must be greater than the start time by at least 1 day and end time cannot exceed 3 months after start time.If the start time and end time passed in by the seller overlap with other upcoming/ongoing activities, it will prompt "Another Follow Prize voucher already exists during this time period, please set another period."<br /></p>
	UsageQuantity   *int64   `json:"usage_quantity,omitempty"`    // [Optional] <p>Please enter a value between 1 and 200000.</p>
	MinSpend        *float64 `json:"min_spend,omitempty"`         // [Optional] <p>The minimum spend required for using this follow prize.<br /></p>
}

type UpdateFollowPrizeResponse struct {
	BaseResponse `json:",inline"`              // Common response fields
	Response     UpdateFollowPrizeResponseData `json:"response"` // Actual response data
}

type UpdateFollowPrizeResponseData struct {
	CampaginId int64 `json:"campagin_id"` // [Required] <p>The unique identifier for the created follow prize.<br /></p>
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

type UpdateItemListRequest struct {
	SessionId int64                             `json:"session_id"` // [Required] <p>The identifier of livestream session.</p>
	ItemList  []DeleteBundleDealItemRequestItem `json:"item_list"`  // [Required] <p>The list of item with updated order.</p>
}

type UpdateItemListResponse struct {
	BaseResponse `json:",inline"` // Common response fields
	Response     interface{}      `json:"response"` // Actual response data
}

type UpdateItemRequest struct {
	Description          *string                   `json:"description,omitempty"`            // [Optional] Description of item.
	Weight               *float64                  `json:"weight,omitempty"`                 // [Optional] <p>The weight of this item, the unit is KG.</p><p>Updating the weight of this item will overwrite the weight of all models under this item.</p>
	PreOrder             *PreOrder                 `json:"pre_order,omitempty"`              // [Optional] Pre Order setting.
	ItemName             *string                   `json:"item_name,omitempty"`              // [Optional] Item name.
	AttributeList        []Attribute               `json:"attribute_list,omitempty"`         // [Optional] Item attributes.
	Image                *RequestImage             `json:"image,omitempty"`                  // [Optional] Images of item.
	ItemSku              *string                   `json:"item_sku,omitempty"`               // [Optional] SKU tag for item.
	ItemStatus           *ItemStatus               `json:"item_status,omitempty"`            // [Optional] Item status, could be UNLIST or NORMAL.
	LogisticInfo         []LogisticInfo            `json:"logistic_info,omitempty"`          // [Optional] Logistic channel setting.
	Wholesale            []Wholesale               `json:"wholesale,omitempty"`              // [Optional] <p>Wholesale setting.</p><p>If you want to delete it, please pass it with blank.<br /></p>
	ItemId               int64                     `json:"item_id"`                          // [Required] ID of item.
	CategoryId           *int64                    `json:"category_id,omitempty"`            // [Optional] ID of category.
	Dimension            *Dimension                `json:"dimension,omitempty"`              // [Optional] <p>The dimension of this item.</p><p>Updating the dimension of this item will overwrite the dimension of all models under this item.<br /></p>
	Condition            *string                   `json:"condition,omitempty"`              // [Optional] Condition of item, could be NEW or USED.
	VideoUploadId        []string                  `json:"video_upload_id,omitempty"`        // [Optional] <p>Video upload ID returned from video uploading API.</p><p>If you want to delete it, please pass it with blank.</p>
	Brand                *Brand                    `json:"brand,omitempty"`                  // [Optional]
	ItemDangerous        *int64                    `json:"item_dangerous,omitempty"`         // [Optional] This field is only applicable for local sellers in Indonesia and Malaysia. Use this field to identify whether a product is a dangerous product. 0 for non-dangerous product and 1 for dangerous product. For more information, please visit the market's respective Seller Education Hub.
	TaxInfo              *RequestTaxInfo           `json:"tax_info,omitempty"`               // [Optional] Tax information
	ComplaintPolicy      *ComplaintPolicy          `json:"complaint_policy,omitempty"`       // [Optional] Complaint Policy for item. Only required for local PL sellers, ignored otherwise.
	DescriptionInfo      *DescriptionInfo          `json:"description_info,omitempty"`       // [Optional] New description field. Only whitelist sellers can use it. If you use the field, please upload the description_type=extended otherwise api will return error. If you don't use this field, you don't need to upload the description_type or upload description_type=normal
	DescriptionType      *DescriptionType          `json:"description_type,omitempty"`       // [Optional] Values: See Data Definition- description_type (normal , extended). If you want to use extended_description or change description type ,this field must be inputed
	GtinCode             *string                   `json:"gtin_code,omitempty"`              // [Optional] <p>- GTIN is an identifier for trade items, developed by the international organization GS1.<br />- They have 8 to 14 digits. The most common are UPC, EAN, JAN and ISBN.<br />- GTIN will help boost positioning in online marketing channels like Google and Facebook.<br />- That incorporation with GTIN will also aid in Search and Recommendation in Shopee itself allowing buyers to have higher likelihood of finding one's listing.<br /></p><p><br /></p><p>Note: If you want to set “Item without GTIN”, please pass the gtin_code as "00".<br /><br />The validation rule is based on the value return in gtin_validation_rule" field in v2.product.get_item_limit API</p><p>- Mandatory:&nbsp;This field is required and must contain a correctly formatted GTiN number.</p><p>- Flexible: This field is required and must contain either a correctly formatted GTlN number or "00" to declare that the item/model has no valid GTlN.<br />- Optional: This field is optional and can contain a correctly formatted GTiN number, "00" or be omitted entirely.</p>
	DsCatRcmdId          *string                   `json:"ds_cat_rcmd_id,omitempty"`         // [Optional] <p>category recommendation service id<br /></p>
	PromotionImages      *PromotionImages          `json:"promotion_images,omitempty"`       // [Optional] <p>Promotion Image<br />Currently only allow one promoton image<br />You could set promotion image only if the product images' ratio is 3:4<br /></p>
	CompatibilityInfo    *CompatibilityInfo        `json:"compatibility_info,omitempty"`     // [Optional]
	ScheduledPublishTime *int64                    `json:"scheduled_publish_time,omitempty"` // [Optional] <p>Scheduled publish time of this item:&nbsp;</p><p>1) Can only set scheduled_publish_time for item with UNLIST status</p><p>2) Can only set the time from current time +1hour to current time +90days, and the time is only allowed to be accurate to the minute</p>
	AuthorisedBrandId    *int64                    `json:"authorised_brand_id,omitempty"`    // [Optional] <p>ID of authorised reseller brand.<br /></p>
	SizeChartInfo        *SizeChartInfo            `json:"size_chart_info,omitempty"`        // [Optional] <p><br /></p>
	CertificationInfo    *RequestCertificationInfo `json:"certification_info,omitempty"`     // [Optional] <p>For PH product certification input<br />Required for some category and attribute option</p>
	PurchaseLimitInfo    *PurchaseLimitInfo        `json:"purchase_limit_info,omitempty"`    // [Optional] <p>purchase limit info</p>
}

type UpdateItemResponse struct {
	BaseResponse `json:",inline"`       // Common response fields
	Response     UpdateItemResponseData `json:"response"` // Actual response data
}

type UpdateItemResponseData struct {
	Description     string                       `json:"description"`      // [Required] Item description.
	Weight          float64                      `json:"weight"`           // [Required] <p>The weight of this item, the unit is KG.</p>
	PreOrder        *PreOrder                    `json:"pre_order"`        // [Required]
	ItemName        string                       `json:"item_name"`        // [Required] Item name.
	ItemStatus      ItemStatus                   `json:"item_status"`      // [Required] Item status.
	Images          *Images                      `json:"images"`           // [Required] Item images.
	LogisticInfo    []ResponseDataLogisticInfo   `json:"logistic_info"`    // [Required]
	ItemId          int64                        `json:"item_id"`          // [Required] ID of item.
	CategoryId      int64                        `json:"category_id"`      // [Required] ID of item category.
	Dimension       *Dimension                   `json:"dimension"`        // [Required] <p>The dimension of this item.</p>
	Condition       string                       `json:"condition"`        // [Required] Item condition, could be USED or NEW.
	Brand           *Brand                       `json:"brand"`            // [Required]
	ItemDangerous   int64                        `json:"item_dangerous"`   // [Required] This field is only applicable for local sellers in Indonesia and Malaysia. Use this field to identify whether a product is a dangerous product. 0 for non-dangerous product and 1 for dangerous product. For more information, please visit the market's respective Seller Education Hub.
	ComplaintPolicy *ResponseDataComplaintPolicy `json:"complaint_policy"` // [Required] Complaint policy
	DescriptionInfo *DescriptionInfo             `json:"description_info"` // [Required] New description field. Only whitelist sellers can use it. If you use the field, please upload the description_type=extended otherwise api will return error. If you don't use this field, you don't need to upload the description_type or upload description_type=normal
	DescriptionType DescriptionType              `json:"description_type"` // [Required] Values: See Data Definition- description_type (normal , extended).
}

type UpdateKitItemRequest struct {
	ItemId      int64               `json:"item_id"`                // [Required] <p>ID of kit item.<br /></p>
	ItemSetting *RequestItemSetting `json:"item_setting,omitempty"` // [Optional]
	SyncSetting *SyncSetting        `json:"sync_setting,omitempty"` // [Optional]
}

type UpdateKitItemResponse struct {
	BaseResponse `json:",inline"` // Common response fields
}

type UpdateLocalAdjustmentRateRequest struct {
	AdjustmentRate float64 `json:"adjustment_rate"` // [Required] <p>The multiplier used to adjust the cross-border original price to local price</p>
}

type UpdateLocalAdjustmentRateResponse struct {
	BaseResponse `json:",inline"` // Common response fields
}

type UpdateModelRequest struct {
	ItemId int64                     `json:"item_id"` // [Required] ID of item
	Model  []UpdateModelRequestModel `json:"model"`   // [Required] Length should be between 1 to 50
}

type UpdateModelRequestModel struct {
	ModelId     int64      `json:"model_id"`               // [Required] ID of model
	ModelSku    string     `json:"model_sku"`              // [Required] <p>Seller SKU for model, model_sku length information needs to be no more than 100 characters. <b><font color="#c24f4a">CNSC and KRSC sellers are not allowed to update the MPSKU model sku.</font></b></p>
	PreOrder    *PreOrder  `json:"pre_order,omitempty"`    // [Optional]
	GtinCode    *string    `json:"gtin_code,omitempty"`    // [Optional] <p>- GTIN is an identifier for trade items, developed by the international organization GS1.<br />- They have 8 to 14 digits. The most common are UPC, EAN, JAN and ISBN.<br />- GTIN will help boost positioning in online marketing channels like Google and Facebook.<br />- That incorporation with GTIN will also aid in Search and Recommendation in Shopee itself allowing buyers to have higher likelihood of finding one's listing.<br /></p><p><br /></p><p>Note: If you want to set “Item without GTIN”, please pass the gtin_code as "00".<br /><br />The validation rule is based on the value return in gtin_validation_rule" field in v2.product.get_item_limit API</p><p>- <b>Mandatory</b>:&nbsp;This field is required and must contain a correctly formatted GTiN number.</p><p>- <b>Flexible</b>: This field is required and must contain either a correctly formatted GTlN number or "00" to declare that the item/model has no valid GTlN.<br />- <b>Optional</b>: This field is optional and can contain a correctly formatted GTiN number, "00" or be omitted entirely.</p>
	ModelStatus *string    `json:"model_status,omitempty"` // [Optional] <p>can be&nbsp;"NORMAL" or "UNAVAILABLE". <b><font color="#c24f4a">Only CNSC and KRSC sellers can set the model_status.</font></b>&nbsp;Normal models can be sold on the buyer's side, and UNAVAILABLE models cannot be sold on the buyer's side.</p>
	Weight      *float64   `json:"weight,omitempty"`       // [Optional] <p>The weight of this model, the unit is KG.</p><p>If don't set the weight of this model, will use the weight of item by default.</p><p>If set the dimension of this model, them must set the weight of this model.</p>
	Dimension   *Dimension `json:"dimension,omitempty"`    // [Optional] <p>The dimension of this model.</p><p>If don't set the dimension of this model, will use the dimension of item by default.</p>
}

type UpdateModelResponse struct {
	BaseResponse `json:",inline"` // Common response fields
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

type UpdatePriceRequest struct {
	ItemId    int64   `json:"item_id"`    // [Required] ID of item.
	PriceList []Price `json:"price_list"` // [Required] Length should be between 1 to 50.
}

type UpdatePriceResponse struct {
	BaseResponse `json:",inline"`        // Common response fields
	Response     UpdatePriceResponseData `json:"response"` // Actual response data
}

type UpdatePriceResponseData struct {
	FailureList []ResponseDataFailure `json:"failure_list"` // [Required] Fail model list.
	SuccessList []Price               `json:"success_list"` // [Required] Success model list.
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

type UpdateSelfCollectionOrderLogisticsRequest struct {
	PackageNumber                 string   `json:"package_number"`                   // [Required] <p>Shopee's unique identifier for the package under an order.</p>
	SelfCollectionLogisticsAction string   `json:"self_collection_logistics_action"` // [Required] <p>Order logistics action. available values:</p><p>- ready_for_collection</p><p>- order_collected</p>
	EpocImageList                 []string `json:"epoc_image_list,omitempty"`        // [Optional] <p>List of image_id for the proof that buyer already collected the order at the store.&nbsp;</p><p>Required when self_collection_logistics_action is order_collected. Max: 3.</p><p>You can call the v2.media.upload_image to upload image and get the image_id, for this scenario, please pass the business = 1 and scene = 1.</p>
}

type UpdateSelfCollectionOrderLogisticsResponse struct {
	BaseResponse `json:",inline"` // Common response fields
}

type UpdateSessionRequest struct {
	SessionId     int64   `json:"session_id"`            // [Required] <p>The identifier of livestream session.</p>
	Title         string  `json:"title"`                 // [Required] <p>The title of the livestream session, cannot exceed 200 characters.</p>
	Description   *string `json:"description,omitempty"` // [Optional] <p>The description of the livestream session, cannot exceed 200 characters.</p>
	CoverImageUrl string  `json:"cover_image_url"`       // [Required] <p>The cover image url of the livestream session.</p><p>Please call the v2.livestream.upload_image to upload the cover image file and get the cover_image_url.</p>
	IsTest        bool    `json:"is_test"`               // [Required] <p>Indicate whether this livestream session if for testing purpose only.</p>
}

type UpdateSessionResponse struct {
	BaseResponse `json:",inline"` // Common response fields
	Response     interface{}      `json:"response"` // Actual response data
}

type UpdateShippingOrderRequest struct {
	OrderSn       string         `json:"order_sn"`                 // [Required] Shopee's unique identifier for an order.
	PackageNumber *string        `json:"package_number,omitempty"` // [Optional] Shopee's unique identifier for the package under an order. You should't fill the field with empty string when there is't a package number.
	Pickup        *RequestPickup `json:"pickup"`                   // [Required] Required parameter ONLY if GetParameterForInit returns "pickup" or if GetLogisticsInfo returns "pickup" under "info_needed" for the same order. Developer should still include "pickup" field in the call even if "pickup" has empty value.
}

type UpdateShippingOrderResponse struct {
	BaseResponse `json:",inline"` // Common response fields
}

type UpdateShopCategoryRequest struct {
	ShopCategoryId int64   `json:"shop_category_id"`      // [Required] ShopCategory's unique identifier.
	Name           *string `json:"name,omitempty"`        // [Optional] ShopCategory's name.
	SortWeight     *int64  `json:"sort_weight,omitempty"` // [Optional] ShopCategory's sort weight.
	Status         *string `json:"status,omitempty"`      // [Optional] ShopCategory's status. Applicable values: NORMAL, INACTIVE, DELETED.
}

type UpdateShopCategoryResponse struct {
	BaseResponse `json:",inline"`               // Common response fields
	Response     UpdateShopCategoryResponseData `json:"response"` // Actual response data
}

type UpdateShopCategoryResponseData struct {
	ShopCategoryId int64  `json:"shop_category_id"` // [Required] This is to indicate whether the shop categories list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of shop categories
	Name           string `json:"name"`             // [Required] ShopCategory's name.
	SortWeight     int64  `json:"sort_weight"`      // [Required] ShopCategory's sort weight.
	Status         string `json:"status"`           // [Required] ShopCategory's status. Applicable values: NORMAL, INACTIVE, DELETED.
}

type UpdateShopFlashSaleItemsRequest struct {
	FlashSaleId int64                                  `json:"flash_sale_id"` // [Required]
	Items       []UpdateShopFlashSaleItemsRequestItems `json:"items"`         // [Required]
}

type UpdateShopFlashSaleItemsRequestItems struct {
	ItemId              int64         `json:"item_id"`                          // [Required]
	PurchaseLimit       *int64        `json:"purchase_limit,omitempty"`         // [Optional] <p>min=0, 0 means no limit</p><p><br /></p><p>if the item is in enabled status or the item has models in enabled status, you can't set this field</p>
	Models              []ItemsModels `json:"models,omitempty"`                 // [Optional] <p>If the item has variation, this param is necessary, otherwise please don't use this field</p>
	ItemStatus          *ItemStatus   `json:"item_status,omitempty"`            // [Optional] <p>The status of the item. If the item has no variation, this param is necessary, otherwise don't use this field</p><p><br /></p><p>you can use this field to set the status of item</p><p><br /></p><p>0: disable<br /></p><p>1: enable</p>
	ItemInputPromoPrice *float64      `json:"item_input_promo_price,omitempty"` // [Optional] <p>The promotion price of the item. If the item has no variation, you can use this field to update the promotion price of the item, otherwise don't use this field</p><p><br /></p><p>if the item is enabled(item_status&nbsp; = 1) now, you can't set this field, you can only disable the item</p><p>if the item is disabled(item_status&nbsp; = 0) now and you want to set this field, you should also set item_status to 1</p>
	ItemStock           *int64        `json:"item_stock,omitempty"`             // [Optional] <p>min=1, The campaign stock of the item. If the item has no variation, you can use this field to update the campaign stock of the item, otherwise don't use this field</p><p><br /></p><p>if the item is enabled(item_status&nbsp; = 1) now, you can't set this field,&nbsp;you can only disable the item</p><p>if the item is disabled(item_status&nbsp; = 0) now and you want to set this field, you should also set item_status to 1</p>
}

type UpdateShopFlashSaleItemsResponse struct {
	BaseResponse `json:",inline"`                     // Common response fields
	Response     UpdateShopFlashSaleItemsResponseData `json:"response"` // Actual response data
}

type UpdateShopFlashSaleItemsResponseData struct {
	FailedItems []FailedItems `json:"failed_items"` // [Required]
}

type UpdateShopFlashSaleRequest struct {
	FlashSaleId int64 `json:"flash_sale_id"` // [Required]
	Status      int64 `json:"status"`        // [Required] <p>the status of shop flash sale you want to set, you cannot edit the shop flash sale in 'system_rejected' status</p><p><br /></p><p>Disabling this Flash Sale will disable all items in this session</p><p><br /></p><p>1: enable<br /></p><p>2: disbaled</p>
}

type UpdateShopFlashSaleResponse struct {
	BaseResponse `json:",inline"`                // Common response fields
	Response     UpdateShopFlashSaleResponseData `json:"response"` // Actual response data
}

type UpdateShopFlashSaleResponseData struct {
	TimeslotId  int64 `json:"timeslot_id"`   // [Required]
	FlashSaleId int64 `json:"flash_sale_id"` // [Required]
	Status      int64 `json:"status"`        // [Required] <p>the status of shop flash sale</p><p>0: deleted</p><p>1: enabled</p><p>2: disabled</p><p>3: system_rejected, you cannot edit the shop flash sale in 'system_rejected' status</p>
}

type UpdateShowItemRequest struct {
	SessionId int64 `json:"session_id"` // [Required] <p>The identifier of livestream session.</p>
	ItemId    int64 `json:"item_id"`    // [Required] <p>Shopee's unique identifier for an item.</p>
}

type UpdateShowItemResponse struct {
	BaseResponse `json:",inline"` // Common response fields
	Response     interface{}      `json:"response"` // Actual response data
}

type UpdateSipItemPriceRequest struct {
	ItemId       int64          `json:"item_id"`                  // [Required] ID of item.
	SipItemPrice []SipItemPrice `json:"sip_item_price,omitempty"` // [Optional]
}

type UpdateSipItemPriceResponse struct {
	BaseResponse `json:",inline"` // Common response fields
}

type UpdateSizeChartRequest struct {
	GlobalItemId int64  `json:"global_item_id"` // [Required] Id of global item.
	SizeChart    string `json:"size_chart"`     // [Required] Image id of size chart.
}

type UpdateSizeChartResponse struct {
	BaseResponse `json:",inline"` // Common response fields
}

type UpdateStockRequest struct {
	ItemId    int64   `json:"item_id"`    // [Required] ID of item.
	StockList []Stock `json:"stock_list"` // [Required] Length should be between 1 to 50.
}

type UpdateStockResponse struct {
	BaseResponse `json:",inline"`        // Common response fields
	Response     UpdateStockResponseData `json:"response"` // Actual response data
}

type UpdateStockResponseData struct {
	FailureList []ResponseDataFailure `json:"failure_list"` // [Required] Fail model list.
	SuccessList []Success             `json:"success_list"` // [Required] Success model list.
}

type UpdateTierVariationRequest struct {
	ItemId                   int64                      `json:"item_id"`                              // [Required] ID of item.
	ModelList                []RequestModel             `json:"model_list,omitempty"`                 // [Optional] <p>Item's model list</p>
	StandardiseTierVariation []StandardiseTierVariation `json:"standardise_tier_variation,omitempty"` // [Optional] <p>&nbsp;item standardise tier variation<br />&nbsp;There is at least one standardise_tier_variation and&nbsp;tier_variation&nbsp;<br /></p>
}

type UpdateTierVariationRequestModel struct {
	ModelId   int64 `json:"model_id"`   // [Required] <p>ID of model<br /></p>
	TierIndex int64 `json:"tier_index"` // [Required] <p>Model's tier_variation<br /></p>
}

type UpdateTierVariationResponse struct {
	BaseResponse `json:",inline"` // Common response fields
}

type UpdateTopPicksRequest struct {
	TopPicksId  int64        `json:"top_picks_id"`           // [Required] collection id
	Name        *string      `json:"name,omitempty"`         // [Optional] collection name
	ItemIdList  *interface{} `json:"item_id_list,omitempty"` // [Optional] a list of item id, and we will cover old item_ids by new_item_ids
	IsActivated *bool        `json:"is_activated,omitempty"` // [Optional] if true, we will close other collection and open this collection
}

type UpdateTopPicksResponse struct {
	BaseResponse `json:",inline"`           // Common response fields
	Response     UpdateTopPicksResponseData `json:"response"` // Actual response data
}

type UpdateTopPicksResponseData struct {
	CollectionList []Collection `json:"collection_list"` // [Required] The top picks list in this shop.
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

type UpdateVoucherRequest struct {
	VoucherId          int64        `json:"voucher_id"`                     // [Required] The unique identifier of the voucher which is going to be updated.
	VoucherName        *string      `json:"voucher_name,omitempty"`         // [Optional] The name of the voucher
	StartTime          *int64       `json:"start_time,omitempty"`           // [Optional] <p>The timing from when the voucher is valid; so buyer is allowed to claim and to use. Field can only be updated if voucher has not started.</p>
	EndTime            *int64       `json:"end_time,omitempty"`             // [Optional] The timing until when the voucher is still valid. Any time after this end_time, buyer is not allowed to claim or to use.
	UsageQuantity      *int64       `json:"usage_quantity,omitempty"`       // [Optional] The number of times for this particular voucher could be used.
	MinBasketPrice     *float64     `json:"min_basket_price,omitempty"`     // [Optional] The minimum spend required for using this voucher.
	DiscountAmount     *float64     `json:"discount_amount,omitempty"`      // [Optional] The discount amount set for this particular voucher. Only fill in when you are updating a fix amount voucher.
	Percentage         *int64       `json:"percentage,omitempty"`           // [Optional] The discount percentage set for this particular voucher. Only fill in when you are updating a discount percentage voucher or coins cashback voucher.
	MaxPrice           *float64     `json:"max_price,omitempty"`            // [Optional] The max amount of discount/value a user can enjoy by using this particular voucher. Only fill in when you are updating a discount percentage voucher or coins cashback voucher.
	DisplayChannelList *interface{} `json:"display_channel_list,omitempty"` // [Optional] The FE channel where the voucher will be displayed. The available values are: 1: display_all, 2: order page, 3: feed, 4: live streaming,   [] (empty - which is hidden).
	ItemIdList         []int64      `json:"item_id_list,omitempty"`         // [Optional] The list of items which is applicable for the voucher. Only fill in when you are updating a product type of voucher.
	DisplayStartTime   *int64       `json:"display_start_time,omitempty"`   // [Optional] The timing of when voucher is displayed on shop pages; so buyer is allowed to claim.
}

type UpdateVoucherResponse struct {
	BaseResponse `json:",inline"`          // Common response fields
	Response     UpdateVoucherResponseData `json:"response"` // Actual response data
}

type UpdateVoucherResponseData struct {
	VoucherId int64 `json:"voucher_id"` // [Required] The unique identifier of the voucher which is being updated.
}

type UpdatedChannels struct {
	ChannelId          int64                `json:"channel_id"`           // [Required] <p>Logistics channel ID<br /></p>
	ChannelDisplayName string               `json:"channel_display_name"` // [Required] <p>Logistics channel name<br /></p>
	UnsupportWarehouse []UnsupportWarehouse `json:"unsupport_warehouse"`  // [Required] <p>List details of unsupported warehouses<br /></p>
}

type UploadImageRequest struct {
	Image interface{} `json:"image"`           // [Required] <p>image files. Max 10.0 MB each. Image format accepted: JPG, JPEG, PNG. image number should be less than 9<br /></p>
	Scene *string     `json:"scene,omitempty"` // [Optional] The scene where the picture is used, The value range is normal or desc; normal: we will process the image as a square image, it is recommended to use when uploading item image; desc: we will not process the image, it is recommended to use when uploading the image of extend_description, if you do not upload this field, it will be normal.
	Ratio *string     `json:"ratio,omitempty"` // [Optional] <p>only applicable to whitelisted sellers.<br />only support 1:1 and 3:4;&nbsp;</p><p>1:1 by default.<br /></p>
}

type UploadImageResponseDataImageInfo struct {
	Id        int64                  `json:"id"`         // [Required] <p>the index of images</p>
	Error     string                 `json:"error"`      // [Required] <p>Indicate error type if this index's image upload processing hit error. Empty if no error happened for this index's image .<br /></p>
	Message   string                 `json:"message"`    // [Required] <p>Indicate error detail if this index's image upload processing hit error. Empty if no error happened for this index's image .<br /></p>
	ImageInfo *ResponseDataImageInfo `json:"image_info"` // [Required]
}

type UploadInvoiceDocRequest struct {
	OrderSn  string      `json:"order_sn"`  // [Required] Shopee's unique identifier for an order.
	FileType int64       `json:"file_type"` // [Required] <p>the type of invoice file. 1:pdf 2.jpeg 3.png. 4.xml</p>
	File     interface{} `json:"file"`      // [Required] <p>invoice file.&nbsp;File size limit to 1MB.</p>
}

type UploadInvoiceDocResponse struct {
	BaseResponse `json:",inline"` // Common response fields
}

type UploadProofRequest struct {
	ReturnSn    string  `json:"return_sn"`             // [Required] <p>The serial number of return.</p>
	Photo       []Photo `json:"photo,omitempty"`       // [Optional]
	Description *string `json:"description,omitempty"` // [Optional] <p>text description in the dispute proof<br /></p>
}

type UploadProofResponse struct {
	BaseResponse `json:",inline"` // Common response fields
	Response     interface{}      `json:"response"` // Actual response data
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

type UploadShippingProofRequest struct {
	ReturnSn                    string            `json:"return_sn"`                                // [Required] <p>The serial number of return.</p>
	ReverseLogisticsCarrierId   int64             `json:"reverse_logistics_carrier_id"`             // [Required] <p>Unique ID of non-integrated reverse logistics channel used by seller.</p>
	ReverseLogisticsCarrierName *string           `json:"reverse_logistics_carrier_name,omitempty"` // [Optional] <p>Non-integrated reverse logistics channel name used by seller.</p>
	TrackingNumber              *string           `json:"tracking_number,omitempty"`                // [Optional] <p>Tracking number used in seller arrange. Required when&nbsp;is_tracking_number_required = true in v2.returns.get_shipping_carrier.</p>
	ImageIdList                 []SharedImageInfo `json:"image_id_list,omitempty"`                  // [Optional] <p>List of image_id of shipping proof image. Required when&nbsp;is_shipping_image_file_mandatory = true in v2.returns.get_shipping_carrier. Max: 3.</p><p>You can call the v2.media.upload_image to upload image and get the image_id, for this scenario, please pass the business = 2 and scene = 1.</p>
	Remarks                     *string           `json:"remarks,omitempty"`                        // [Optional] <p>Optional remarks</p>
}

type UploadShippingProofResponse struct {
	BaseResponse `json:",inline"` // Common response fields
	Response     interface{}      `json:"response"` // Actual response data
}

type UploadVideoPartRequest struct {
	VideoUploadId string      `json:"video_upload_id"` // [Required] The video_upload_id in the response of initiate_video_upload
	PartSeq       int64       `json:"part_seq"`        // [Required] Sequence of the current part, starts from 0
	ContentMd5    string      `json:"content_md5"`     // [Required] md5 of this part
	PartContent   interface{} `json:"part_content"`    // [Required] The content of this part of file.  Part size should be exactly 4MB, except last part of file.
}

type UploadVideoPartResponse struct {
	BaseResponse `json:",inline"` // Common response fields
}

type User struct {
	Username string `json:"username"` // [Required] <p>Buyer's nickname, will be masked as "****" if it is a non-integrated return in TW region.</p>
	Email    string `json:"email"`    // [Required] <p>Buyer's email, will be empty if it is a non-integrated return in TW region.</p>
	Portrait string `json:"portrait"` // [Required] <p>Buyer's portrait, will be empty if it is a non-integrated return in TW region.</p>
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

type Vehicle struct {
	BrandId     int64  `json:"brand_id"`     // [Required] <p>ID of the brand.<br /></p>
	BrandName   string `json:"brand_name"`   // [Required] <p>Name of the brand.<br /></p>
	ModelId     int64  `json:"model_id"`     // [Required] <p>ID of the model.<br /></p>
	ModelName   string `json:"model_name"`   // [Required] <p>Name of the model.<br /></p>
	YearId      int64  `json:"year_id"`      // [Required] <p>ID of the year.<br /></p>
	YearName    string `json:"year_name"`    // [Required] <p>Name of the year.<br /></p>
	VersionId   int64  `json:"version_id"`   // [Required] <p>ID of the version.<br /></p>
	VersionName string `json:"version_name"` // [Required] <p>Name of the version.<br /></p>
}

type VehicleInfo struct {
	BrandId   int64 `json:"brand_id"`   // [Required] <p>ID of the brand.</p>
	ModelId   int64 `json:"model_id"`   // [Required] <p>ID of the model.<br /></p>
	YearId    int64 `json:"year_id"`    // [Required] <p>ID of the year.<br /></p>
	VersionId int64 `json:"version_id"` // [Required] <p>ID of the version.<br /></p>
}

type Video struct {
	VideoUrl     string `json:"video_url"`     // [Required] <p>Url of video.<br /></p>
	ThumbnallUrl string `json:"thumbnall_url"` // [Required] <p>Thumbnail of video.<br /></p>
	Duration     int64  `json:"duration"`      // [Required] <p>Duration of video.<br /></p>
}

type VideoInfo struct {
	VideoUrl     string `json:"video_url"`     // [Required] Url of video.
	ThumbnailUrl string `json:"thumbnail_url"` // [Required] Thumbnail of video.
	Duration     int64  `json:"duration"`      // [Required] Duration of video.
}

type VideoUrl struct {
	VideoUrlRegion string `json:"video_url_region"` // [Required] The region of this video URL.
	VideoUrl       string `json:"video_url"`        // [Required] Video playback URL.
}

type ViolationListing struct {
	ItemId         int64 `json:"item_id"`         // [Required] <p>Item ID.</p>
	DetailedReason int64 `json:"detailed_reason"` // [Required] <p>Reason. Applicable values:&nbsp;</p><p>1: Prohibited</p><p>2: Counterfeit</p><p>3: Spam</p><p>4: Inappropriate Image</p><p>5: Insufficient Info</p><p>6: Mall Listing Improvement</p><p>7: Other Listing Improvement<br /></p><p>8: PQR Products</p>
	UpdateTime     int64 `json:"update_time"`     // [Required] <p>Updated on.</p>
}

type VolumeLimit struct {
	ItemMaxVolume float64 `json:"item_max_volume"` // [Required] The max volume for an item on this logistic channel.If the value is 0 or null, that means there is no limit for the item weight.
	ItemMinVolume float64 `json:"item_min_volume"` // [Required] The min volume for an item on this logistic channel. If the value is 0 or null, that means there is no limit for the item weight.
}

type Voucher struct {
	VoucherId        int64   `json:"voucher_id"`         // [Required] The unique identifier for a voucher.
	VoucherCode      string  `json:"voucher_code"`       // [Required] Voucher Code
	VoucherName      string  `json:"voucher_name"`       // [Required] Voucher Name
	VoucherType      int64   `json:"voucher_type"`       // [Required] The type of the voucher. The available values are: 1: shop voucher, 2: product voucher.
	RewardType       int64   `json:"reward_type"`        // [Required] The reward type of the voucher. The available values are: 1: fix_amount voucher, 2: discount_percentage voucher, 3: coin_cashback voucher.
	UsageQuantity    int64   `json:"usage_quantity"`     // [Required] The number of times for this particular voucher could be used.
	CurrentUsage     int64   `json:"current_usage"`      // [Required] Up till now, how many times has this particular voucher already been used.
	StartTime        int64   `json:"start_time"`         // [Required] The timing from when the voucher is valid; so buyer is allowed to claim and to use.
	EndTime          int64   `json:"end_time"`           // [Required] The timing until when the voucher is still valid. Any time after this end_time, buyer is not allowed to claim or to use.
	IsAdmin          bool    `json:"is_admin"`           // [Required] If the voucher is created by Shopee or not.
	VoucherPurpose   int64   `json:"voucher_purpose"`    // [Required] The use case for the voucher. The available values are: 0: normal; 1: welcome, 2: referral; 3: shop_follow; 4:shop_game, 5: free_gift, 6: membership
	DiscountAmount   float64 `json:"discount_amount"`    // [Required] The discount amount set for this particular voucher. Only when it is a fix amount voucher, api will return a value.
	Percentage       int64   `json:"percentage"`         // [Required] The discount percentage set for this particular voucher. Only when it is a discount percentage voucher or coins cashback voucher, api will return a value.
	CmtVoucherStatus int64   `json:"cmt_voucher_status"` // [Required] The voucher status in CMT. The available values are: 1:review, 2: approved, 3:reject. Only when this voucher is attending CMT campaign and not being rejected, api will return a value.
	DisplayStartTime int64   `json:"display_start_time"` // [Required] The timing of when voucher is displayed on shop pages; so buyer is allowed to claim.
}

type Warehouse struct {
	WarehouseId     int64           `json:"warehouse_id"`     // [Required] <p>Warehouse address identifier.<br /></p>
	WarehouseName   string          `json:"warehouse_name"`   // [Required] <p>The warehouse name filled in when creating the warehouse.<br /></p>
	WarehouseType   int64           `json:"warehouse_type"`   // [Required] <p>1 means pickup warehouse</p><p>2 means return warehouse<br /></p>
	WarehouseRegion string          `json:"warehouse_region"` // [Required] <p>Region of your warehouse.<br /></p>
	LocationId      string          `json:"location_id"`      // [Required] <p>Location identifier for stocks. Different location_ids represent that your addresses are in different item stocks.<br /></p>
	Address         *Address        `json:"address"`          // [Required]
	EnterpriseInfo  *EnterpriseInfo `json:"enterprise_info"`  // [Required]
}

type WarehouseFilters struct {
	WarehouseName     string `json:"warehouse_name"`      // [Required] <p>The warehouse name filled in when creating the warehouse address.</p>
	WarehouseType     int64  `json:"warehouse_type"`      // [Required] <p>Type of warehouse. Applicable values:</p><p>- 1: Local Warehouse</p><p>- 2: CB Warehouse</p>
	ProductLocationId string `json:"product_location_id"` // [Required] <p>Location identifier for stocks. Different location_ids represent that your addresses are in different item stocks.</p>
	AddressId         int64  `json:"address_id"`          // [Required] <p>Identity of address.</p>
	Address           string `json:"address"`             // [Required] <p>Detail address of your warehouse.</p>
}

type Waybill struct {
	BindingId        string `json:"binding_id"`         // [Required] <p>Binding ID</p>
	ShippingLabelUrl string `json:"shipping_label_url"` // [Required] <p>URL for downloading waybill.</p>
}

type WeightLimit struct {
	WeightMandatory bool `json:"weight_mandatory"` // [Required] <p>weight is mandatory or not<br /></p>
}

type Wholesale struct {
	MinCount  int64   `json:"min_count"`  // [Required] Minimum count of this tier
	MaxCount  int64   `json:"max_count"`  // [Required] Maximum count of this tier
	UnitPrice float64 `json:"unit_price"` // [Required] Unit price of this tier
}

type WholesalePriceThresholdPercentage struct {
	MinLimit int64 `json:"min_limit"` // [Required]
	MaxLimit int64 `json:"max_limit"` // [Required]
}

type Wholesales struct {
	MinCount                 int64   `json:"min_count"`                    // [Required] The min count of this tier wholesale.
	MaxCount                 int64   `json:"max_count"`                    // [Required] The max count of this tier wholesale.
	UnitPrice                float64 `json:"unit_price"`                   // [Required] The current price of the wholesale in the listing currency.If item is in promotion, this price is useless.
	InflatedPriceOfUnitPrice float64 `json:"inflated_price_of_unit_price"` // [Required] The After-tax Price of the wholesale show to buyer.
}

type Whs struct {
	WhsId                      string  `json:"whs_id"`                         // [Required] <p>&nbsp;Warehouse ID</p>
	StockLevel                 int64   `json:"stock_level"`                    // [Required] <p>-1-No need to calculate stock level；0-None；1-Low Stock &amp; No Sellable stock; 2-Low Stock &amp; To replenish; 3-Low Stock &amp; Replenished; 4-Excess<br /></p>
	IrApprovalQty              int64   `json:"ir_approval_qty"`                // [Required] <p>IR approval but no ASN generated will be included</p>
	InTransitPendingPutawayQty int64   `json:"in_transit_pending_putaway_qty"` // [Required] <p>ASN in-transit，ASN pending putaway, Move transfer in transit and Move transfer pending putaway will be included</p>
	SellableQty                int64   `json:"sellable_qty"`                   // [Required] <p>Stocks that are available for sale. This includes warehouse sellable stock, move transfer &amp; rack transfer reserved stock that available for sale, pre-order stock.This quantity may be greater than qty displayed to buyers as it includes stock reserved for upcoming promotions.<br /></p>
	ReservedQty                int64   `json:"reserved_qty"`                   // [Required] <p>Stocks reserved by buyer order, RTS. And also includes rack transfer reserved stock that are not available for sale<br /></p>
	UnsellableQty              int64   `json:"unsellable_qty"`                 // [Required] <p>Stocks in the warehouse that are not available for sale. This includes damaged stocks, isolated stock due to expired/near expiring, in warehouse staging location, missing, etc.<br /></p>
	ExcessStock                int64   `json:"excess_stock"`                   // [Required] <p>Number of units that are above 6 days of sales coverage Days.</p>
	CoverageDays               float64 `json:"coverage_days"`                  // [Required] <p>Days that the current sellable and pending inbound inventory are expected to last based on current selling speed.<br /></p>
	InWhsCoverageDays          float64 `json:"in_whs_coverage_days"`           // [Required] <p>Days that the current sellable inventory are expected to last based on current selling speed.<br /></p>
	SellingSpeed               float64 `json:"selling_speed"`                  // [Required] <p>Average daily sold quantity</p>
	Last_7Sold                 int64   `json:"last_7_sold"`                    // [Required] <p>Sales qty last 7 days</p>
	Last_15Sold                int64   `json:"last_15_sold"`                   // [Required] <p>Sales qty last 15 days<br /></p>
	Last_30Sold                int64   `json:"last_30_sold"`                   // [Required] <p>Sales qty last 30 days<br /></p>
	Last_60Sold                int64   `json:"last_60_sold"`                   // [Required] <p>Sales qty last 60 days<br /></p>
	Last_90Sold                int64   `json:"last_90_sold"`                   // [Required] <p>Sales qty last 90 days<br /></p>
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

type InvoiceOption string

const (
	InvoiceOptionNonVat          InvoiceOption = "NON_VAT_INVOICES"
	InvoiceOptionNoInvoice       InvoiceOption = "NO_INVOICE"
	InvoiceOptionVat             InvoiceOption = "VAT_INVOICES"
	InvoiceOptionVatMarginScheme InvoiceOption = "VAT_MARGIN_SCHEME_INVOICES"
)

type ItemStatus string

const (
	ItemStatusBanned ItemStatus = "BANNED"
	ItemStatusNormal ItemStatus = "NORMAL"
	ItemStatusUnlist ItemStatus = "UNLIST"
)

type CampaignStatus string

const (
	CampaignStatusClosed    CampaignStatus = "closed"
	CampaignStatusDeleted   CampaignStatus = "deleted"
	CampaignStatusEnded     CampaignStatus = "ended"
	CampaignStatusExpired   CampaignStatus = "expired"
	CampaignStatusOngoing   CampaignStatus = "ongoing"
	CampaignStatusPaused    CampaignStatus = "paused"
	CampaignStatusScheduled CampaignStatus = "scheduled"
	CampaignStatusUpcoming  CampaignStatus = "upcoming"
)

type DescriptionElementFieldType string

const (
	DescriptionElementFieldTypeImage DescriptionElementFieldType = "image"
	DescriptionElementFieldTypeText  DescriptionElementFieldType = "text"
)

type TaxType int

const (
	TaxTypeNoTax   TaxType = 0
	TaxTypeTaxable TaxType = 1
	TaxTypeTaxFree TaxType = 2
)

type OperationType string

const (
	OperationTypeRetailer     OperationType = "1"
	OperationTypeManufactorer OperationType = "2"
)

type PromotionStatus string

const (
	PromotionStatusAll      PromotionStatus = "all"
	PromotionStatusExpired  PromotionStatus = "expired"
	PromotionStatusOngoing  PromotionStatus = "ongoing"
	PromotionStatusUpcoming PromotionStatus = "upcoming"
)

type OrderStatus string

const (
	OrderStatusCancelled      OrderStatus = "CANCELLED"
	OrderStatusCompleted      OrderStatus = "COMPLETED"
	OrderStatusInvoicePending OrderStatus = "INVOICE_PENDING"
	OrderStatusInCancel       OrderStatus = "IN_CANCEL"
	OrderStatusProcessed      OrderStatus = "PROCESSED"
	OrderStatusReadyToShip    OrderStatus = "READY_TO_SHIP"
	OrderStatusShipped        OrderStatus = "SHIPPED"
	OrderStatusUnpaid         OrderStatus = "UNPAID"
)

type LogisticsStatus string

const (
	LogisticsStatusCodRejected     LogisticsStatus = "LOGISTICS_COD_REJECTED"
	LogisticsStatusDeliveryDone    LogisticsStatus = "LOGISTICS_DELIVERY_DONE"
	LogisticsStatusDeliveryFailed  LogisticsStatus = "LOGISTICS_DELIVERY_FAILED"
	LogisticsStatusInvalid         LogisticsStatus = "LOGISTICS_INVALID"
	LogisticsStatusLost            LogisticsStatus = "LOGISTICS_LOST"
	LogisticsStatusNotStart        LogisticsStatus = "LOGISTICS_NOT_START"
	LogisticsStatusPickupDone      LogisticsStatus = "LOGISTICS_PICKUP_DONE"
	LogisticsStatusPickupFailed    LogisticsStatus = "LOGISTICS_PICKUP_FAILED"
	LogisticsStatusReady           LogisticsStatus = "LOGISTICS_READY"
	LogisticsStatusRequestCanceled LogisticsStatus = "LOGISTICS_REQUEST_CANCELED"
	LogisticsStatusRequestCreated  LogisticsStatus = "LOGISTICS_REQUEST_CREATED"
)

type ReturnStatus string

const (
	ReturnStatusAccepted      ReturnStatus = "ACCEPTED"
	ReturnStatusCancelled     ReturnStatus = "CANCELLED"
	ReturnStatusClosed        ReturnStatus = "CLOSED"
	ReturnStatusJudging       ReturnStatus = "JUDGING"
	ReturnStatusProcessing    ReturnStatus = "PROCESSING"
	ReturnStatusRefunding     ReturnStatus = "REFUNDING"
	ReturnStatusRequested     ReturnStatus = "REQUESTED"
	ReturnStatusSellerDispute ReturnStatus = "SELLER_DISPUTE"
)

type BookingStatus string

const (
	BookingStatusCancelled   BookingStatus = "CANCELLED"
	BookingStatusMatched     BookingStatus = "MATCHED"
	BookingStatusProcessed   BookingStatus = "PROCESSED"
	BookingStatusReadyToShip BookingStatus = "READY_TO_SHIP"
	BookingStatusShipped     BookingStatus = "SHIPPED"
)

type DescriptionType string

const (
	DescriptionTypeExtended DescriptionType = "extended"
	DescriptionTypeNormal   DescriptionType = "normal"
)

type WarrantyTime string

const (
	WarrantyTimeOneYear      WarrantyTime = "ONE_YEAR"
	WarrantyTimeOverTwoYears WarrantyTime = "OVER_TWO_YEARS"
	WarrantyTimeTwoYears     WarrantyTime = "TWO_YEARS"
)
