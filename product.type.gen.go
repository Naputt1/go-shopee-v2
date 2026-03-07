package goshopee

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

type AdvanceStock struct {
	SellableAdvanceStock  int64 `json:"sellable_advance_stock"`   // [Required] <p>Refers to Advance Fulfillment&nbsp;stock that Seller has shipped out and is available&nbsp;to be used to fulfill an order.<br /></p>
	InTransitAdvanceStock int64 `json:"in_transit_advance_stock"` // [Required] <p>Refers to Advance Fulfillment stock that seller has shipped out and is still in transit and unavailable to be used to fulfill an order.<br /></p>
}

type Aitem struct {
	AshopId          int64          `json:"ashop_id"`           // [Required] <p>ID of SIP Affiliate Shop.</p>
	AshopRegion      string         `json:"ashop_region"`       // [Required] <p>Region of SIP Affiliate Shop.</p>
	AitemId          int64          `json:"aitem_id"`           // [Required] <p>ID of item under SIP Affiliate Shop corresponding to the P Item.</p>
	ModelMappingList []ModelMapping `json:"model_mapping_list"` // [Required] <p>If the P Item does not have model, then the model_mapping_list will not be returned.</p>
}

type Attribute struct {
	AttributeId        int64            `json:"attribute_id"`         // [Required] Attribute ID
	AttributeValueList []AttributeValue `json:"attribute_value_list"` // [Required]
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

type AttributeTreeAttributeValue struct {
	ValueId            int64         `json:"value_id"`             // [Required] <p>Value ID</p>
	Name               string        `json:"name"`                 // [Required] <p>Value name</p>
	ValueUnit          string        `json:"value_unit"`           // [Required] <p>Value unit</p>
	ChildAttributeList []interface{} `json:"child_attribute_list"` // [Required] <p>Child attributes for the value of parent attribute<br />The structure content is the same as attribute_tree<br /></p>
	MultiLang          []MultiLang   `json:"multi_lang"`           // [Required] <p>Translate results for display</p>
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

type Comment struct {
	CommentId int64  `json:"comment_id"` // [Required] The identity of comment.
	Comment   string `json:"comment"`    // [Required] The content of the comment.
}

type CommentReply struct {
	Reply      string `json:"reply"`       // [Required] The content of reply.
	Hidden     bool   `json:"hidden"`      // [Required] The comment reply is hidden or not.
	CreateTime int64  `json:"create_time"` // [Required] <p>The time the seller replied to the comment.</p>
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

type Component struct {
	ComponentItemId  int64  `json:"component_item_id"`            // [Required] <p>ID of the item that composes this kit model.<br /></p>
	ComponentModelId *int64 `json:"component_model_id,omitempty"` // [Optional] <p>ID of the model that composes this kit model.<br /></p>
	Quantity         int64  `json:"quantity"`                     // [Required] <p>The amount of the item/model that composes this kit model.<br /></p>
	MainComponent    *bool  `json:"main_component,omitempty"`     // [Optional] <p>Whether this item/model is the main component for this kit.</p><p>One kit item can only have one item/model as main component.</p>
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

type DirectItem struct {
	DirectShopId int64 `json:"direct_shop_id"` // [Required] <p>Id of direct shop.</p>
	DirectItemId int64 `json:"direct_item_id"` // [Required] <p>Item id of direct shop.</p>
}

type DirectItemPrice struct {
	Region             string           `json:"region"`                // [Required] <p>Region of direct shop.</p>
	HiddenPrice        float64          `json:"hidden_price"`          // [Required]
	ItemModelPriceList []ItemModelPrice `json:"item_model_price_list"` // [Required]
}

type DtsLimit struct {
	DaysToShipLimit       *WholesalePriceThresholdPercentage `json:"days_to_ship_limit"`         // [Required] <p>Pre order limits for the category</p>
	NonPreOrderDaysToShip int64                              `json:"non_pre_order_days_to_ship"` // [Required]
}

type ExtendedDescription struct {
	FieldList []Field `json:"field_list"` // [Required]  Field of extended description.
}

type Failure struct {
	ItemId       int64  `json:"item_id"`       // [Required] <p>Shopee's unique identifier for an item.<br /></p>
	FailedReason string `json:"failed_reason"` // [Required] <p>Item's failure reason.</p>
}

type Field struct {
	FieldType DescriptionElementFieldType `json:"field_type"` // [Required] Type of extended description field ：values: See Data Definition- description_field_type (text , image).
	Text      string                      `json:"text"`       // [Required] If field_type is text， text information will be set by this field.
	ImageInfo *SharedImageInfo            `json:"image_info"` // [Required] If field_type is image，image url will be set by this field.
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
	StockInfoV2           *StockInfoV2                 `json:"stock_info_v2"`            // [Required] <p>new stock info.<br /></p><p>Please check this FAQ for more detail:&nbsp;<a href="https://open.shopee.com/faq?top=162&amp;sub=166&amp;page=1&amp;faq=230" target="_blank" style="font-size:14px;">https://open.shopee.com/faq?top=162&amp;sub=166&amp;page=1&amp;faq=230</a></p>
	GtinCode              string                       `json:"gtin_code"`                // [Required] <p><b><font color="#c24f4a">(Only TW seller and BR local seller available)</font></b> gtin code.</p>
	Weight                string                       `json:"weight"`                   // [Required] <p>The weight of this model, the unit is KG.</p><p>If don't set the weight of this model, will use the weight of item by default.</p>
	Dimension             *Dimension                   `json:"dimension"`                // [Required] <p>The dimension of this model.</p><p>If don't set the dimension of this model, will use the dimension of item by default.</p>
	IsFulfillmentByShopee bool                         `json:"is_fulfillment_by_shopee"` // [Required] <p>whether model is fulfillment by shopee</p>
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

type Image struct {
	ImageIdList []string `json:"image_id_list"`         // [Required] ID of image
	ImageRatio  *string  `json:"image_ratio,omitempty"` // [Optional] <p>Ratio of image, <br />OptionalAllowed ratios :<br />"1:1" (default)&nbsp;<br />"3:4"<br /></p><p>only applicable to whitelisted seller.<br /></p>
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

type ItemImage struct {
	ImageIdList  []string `json:"image_id_list"`  // [Required] <p>List of image id.<br /></p>
	ImageUrlList []string `json:"image_url_list"` // [Required] <p>List of image url.<br /></p>
	ImageRatio   string   `json:"image_ratio"`    // [Required] <p>Image ratio.<br /></p>
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

type ItemModelPrice struct {
	ModelId   int64   `json:"model_id"`   // [Required] <p>Id of main model.</p>
	TierIndex []int64 `json:"tier_index"` // [Required] <p>Tier index of main model. Index starts from 0.</p>
	Price     float64 `json:"price"`      // [Required]
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

type ItemStatusDetails struct {
	ViolationType   string `json:"violation_type"`    // [Required] <p>Violation types defined by Shopee.&nbsp;Applicable values:&nbsp;</p><p>Prohibited Listing</p><p>Counterfeit and IP Infringement</p><p>Spam</p><p>Inappropriate Image</p><p>Insufficient Information</p><p>Mall Listing Improvement</p><p>Other Listing Improvement<br /></p>
	ViolationReason string `json:"violation_reason"`  // [Required] <p>The reason for violation.<br /></p>
	Suggestion      string `json:"suggestion"`        // [Required] <p>Shopee provides you with suggestions for modifying items.<br /></p>
	FixDeadlineTime int64  `json:"fix_deadline_time"` // [Required] <p>Action required deadline. Empty if no deadline.<br /></p>
	UpdateTime      int64  `json:"update_time"`       // [Required] <p>Latest update time.<br /></p>
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

type List struct {
	AttributeTree []AttributeTree `json:"attribute_tree"` // [Required] <p>One category's attribute trees</p>
	CategoryId    int64           `json:"category_id"`    // [Required] <p>Category ID</p>
	Warning       string          `json:"warning"`        // [Required] <p>Warning msg</p>
}

type LogisticInfo struct {
	SizeId      int64   `json:"size_id"`      // [Required] Size ID
	ShippingFee float64 `json:"shipping_fee"` // [Required] Shipping fee
	Enabled     bool    `json:"enabled"`      // [Required] Whether this channel is enabled for this item
	LogisticId  int64   `json:"logistic_id"`  // [Required] Logistic channel ID
	IsFree      bool    `json:"is_free"`      // [Required] Whether cover shipping fee for buyer
}

type MaxPurchaseLimit struct {
	PurchaseLimit int64 `json:"purchase_limit"` // [Required] <p>maximum purchase limit for each order</p>
}

type Media struct {
	ImageUrlList []string `json:"image_url_list"` // [Required] <p>List of image url uploaded by the buyer in the comment.</p>
	VideoUrlList []string `json:"video_url_list"` // [Required] <p>List of video url uploaded by the buyer in the comment.</p>
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

type ModelSetting struct {
	TierIndex interface{} `json:"tier_index"` // [Required] <p>Tier index of this model.<br /></p>
	Gtin      string      `json:"gtin"`       // [Required] <p>GTIN of this model.</p>
}

type PageInfo struct {
	HasNextPage bool   `json:"has_next_page"` // [Required]
	NextOffset  string `json:"next_offset"`   // [Required]
}

type PreOrder struct {
	IsPreOrder bool  `json:"is_pre_order"` // [Required]
	DaysToShip int64 `json:"days_to_ship"` // [Required]
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

type PurchaseLimitInfo struct {
	MinPurchaseLimit int64             `json:"min_purchase_limit"` // [Required] <p>minimum purchase count for each order</p>
	MaxPurchaseLimit *MaxPurchaseLimit `json:"max_purchase_limit"` // [Required]
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

type RequestCertificationInfo struct {
	CertificationList []CertificationInfoCertification `json:"certification_list,omitempty"` // [Optional] <p>Array of certification records for the product, each containing type, certificate number, permit ID, and proof documents.</p><p>  </p><p><br /></p>
}

type RequestComponent struct {
	ComponentItemId  int64  `json:"component_item_id"`            // [Required] <p>ID of the item that composes this kit model.</p>
	ComponentModelId *int64 `json:"component_model_id,omitempty"` // [Optional] <p>ID of the model that composes this kit model.</p>
}

type RequestImage struct {
	ImageIdList []interface{} `json:"image_id_list"`         // [Required] Image ID.
	ImageRatio  *string       `json:"image_ratio,omitempty"` // [Optional] <p>Ratio of image, <br />OptionalAllowed ratios :<br />"1:1" (default)&nbsp;<br />"3:4"<br /></p>
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

type RequestModel struct {
	ModelId   int64   `json:"model_id"`   // [Required] <p>ID of model</p>
	TierIndex []int64 `json:"tier_index"` // [Required] <p>Model's tier_variation</p>
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

type ResponseDataComplaintPolicy struct {
	WarrantyTime                WarrantyTime `json:"warranty_time"`                 // [Required] Value should be in one of ONE_YEAR TWO_YEARS OVER_TWO_YEARS.
	ExcludeEntrepreneurWarranty bool         `json:"exclude_entrepreneur_warranty"` // [Required] If True means "I exclude warranty complaints for entrepreneur"
	AdditionalInformation       string       `json:"additional_information"`        // [Required] Additional information for complaint policy
}

type ResponseDataDtsLimit struct {
	NonPreOrderDaysToShip int64                              `json:"non_pre_order_days_to_ship"` // [Required] <p>Days to ship for non pre-order products.<br /></p>
	SupportPreOrder       bool                               `json:"support_pre_order"`          // [Required] <p>Whether support pre_order for the category.<br /></p>
	DaysToShipLimit       *WholesalePriceThresholdPercentage `json:"days_to_ship_limit"`         // [Required] <p>Days to ship for pre-order products.<br /></p>
}

type ResponseDataFailure struct {
	ModelId      int64  `json:"model_id"`      // [Required] ID of model.
	FailedReason string `json:"failed_reason"` // [Required] Reason for failure.
}

type ResponseDataItem struct {
	ItemId     int64      `json:"item_id"`     // [Required] Shopee's unique identifier for an item.
	ItemStatus ItemStatus `json:"item_status"` // [Required] <p>Enumerated type that defines the current status of the item. Applicable values: NORMAL,&nbsp;BANNED, UNLIST, <b><font color="#c24f4a">REVIEWING, SELLER_DELETE, SHOPEE_DELETE</font></b>.<br /></p>
	UpdateTime int64      `json:"update_time"` // [Required] The update time of item.
	Tag        *Tag       `json:"tag"`         // [Required]
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

type ResponseDataPriceInfo struct {
	CurrentPrice  float64 `json:"current_price"`  // [Required] Current price of item
	OriginalPrice float64 `json:"original_price"` // [Required] Original price of item
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

type Result struct {
	CommentId   int64  `json:"comment_id"`   // [Required] The identity of comment.
	FailError   string `json:"fail_error"`   // [Required] Indicate error details if one element hit error.
	FailMessage string `json:"fail_message"` // [Required] Indicate error type if one element hit error.
}

type SalesInfo struct {
	ModelSettingList     []ModelSetting                         `json:"model_setting_list"`      // [Required]
	StdTierVariationList []ResponseDataStandardiseTierVariation `json:"std_tier_variation_list"` // [Required] <p>Standardise Variation config of Shopee Standard Product.<br /></p>
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

type SellerStock struct {
	LocationId string `json:"location_id"` // [Required] <p>location id<br /></p>
	Stock      int64  `json:"stock"`       // [Required] <p>stock<br /></p>
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

type SipItemPrice struct {
	ModelId      *int64  `json:"model_id,omitempty"` // [Optional] 0 for no model item.
	SipItemPrice float64 `json:"sip_item_price"`     // [Required] SIP item price.
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

type Stock struct {
	ModelId     *int64        `json:"model_id,omitempty"` // [Optional] 0 for no model item.
	SellerStock []SellerStock `json:"seller_stock"`       // [Required] <p>new stock info（Please notice that stock(including Seller Stock and Shopee Stock) should be larger than or equal to real-time reserved stock）<br /></p>
}

type StockInfoV2 struct {
	SummaryInfo  *SummaryInfo             `json:"summary_info"`  // [Required] <p>stock summary Info<br /></p>
	SellerStock  []StockInfoV2SellerStock `json:"seller_stock"`  // [Required] <p>Seller-managed stock<br /></p>
	ShopeeStock  []SellerStock            `json:"shopee_stock"`  // [Required] <p>Shopee warehouse stock<br /></p>
	AdvanceStock *AdvanceStock            `json:"advance_stock"` // [Required] <p>Only for PH/VN/ID/MY local selected shops.</p>
}

type StockInfoV2SellerStock struct {
	LocationId string `json:"location_id"` // [Required] <p>location id<br /></p>
	Stock      int64  `json:"stock"`       // [Required] <p>stock in the current warehouse<br /></p>
	IfSaleable bool   `json:"if_saleable"` // [Required] <p>To return if the stock of the location id is saleable<br /></p>
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

type SummaryInfo struct {
	TotalReservedStock  int64 `json:"total_reserved_stock"`  // [Required] <p>Stock reserved for promotion.<br /></p><p><br /></p><p>Note: For SIP P Item, will return the total reserved stock for P Item and all A Items under the P Item.</p>
	TotalAvailableStock int64 `json:"total_available_stock"` // [Required] <p>Stock can be sold currently<br /></p>
}

type SyncSetting struct {
	AutoSyncDts bool `json:"auto_sync_dts"` // [Required] <p>Auto sync the pre_order setting from main component or not.<br /></p>
}

type Tag struct {
	Kit bool `json:"kit"` // [Required] <p>Indicate if the item is kit item.</p>
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

type TierVariation struct {
	Name       *string        `json:"name,omitempty"` // [Optional] <p>Tier variation name.<br /></p>
	OptionList []SharedOption `json:"option_list"`    // [Required] <p>Tier variation option info list.<br /></p>
}

type TierVariationOption struct {
	Image  *OptionImage `json:"image"`  // [Required] Image of this option
	Option string       `json:"option"` // [Required] Option name
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

type UpdateSipItemPriceRequest struct {
	ItemId       int64          `json:"item_id"`                  // [Required] ID of item.
	SipItemPrice []SipItemPrice `json:"sip_item_price,omitempty"` // [Optional]
}

type UpdateSipItemPriceResponse struct {
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

type UpdateTierVariationResponse struct {
	BaseResponse `json:",inline"` // Common response fields
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

type Wholesale struct {
	MinCount  int64   `json:"min_count"`  // [Required] Minimum count of this tier
	MaxCount  int64   `json:"max_count"`  // [Required] Maximum count of this tier
	UnitPrice float64 `json:"unit_price"` // [Required] Unit price of this tier
}

type Wholesales struct {
	MinCount                 int64   `json:"min_count"`                    // [Required] The min count of this tier wholesale.
	MaxCount                 int64   `json:"max_count"`                    // [Required] The max count of this tier wholesale.
	UnitPrice                float64 `json:"unit_price"`                   // [Required] The current price of the wholesale in the listing currency.If item is in promotion, this price is useless.
	InflatedPriceOfUnitPrice float64 `json:"inflated_price_of_unit_price"` // [Required] The After-tax Price of the wholesale show to buyer.
}
