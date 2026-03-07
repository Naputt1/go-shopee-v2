// Package goshopee realizes Shopee API v2
// https://open.shopee.com/documents?module=87&type=2&id=58&version=2
package goshopee

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/google/go-querystring/query"
)

const (
	UserAgent = "goshopee.v2/1.0.0"

	defaultHttpTimeout = 10

	// BEGIN GENERATED ERRORS
	ErrADDONDEALDTSERROR                                      = "ADD_ON_DEAL_DTS_ERROR"                                            // The Days to ship is longer than that of main item
	ErrAddOnAddOnDealEndTimeError                             = "add_on.add_on_deal_end_time_error"                                // The end time should be 1 hour later than start time
	ErrAddOnAddOnDealErrorUnknown                             = "add_on.add_on_deal_error_unknown"                                 // The add on deal created by admin can't be edited
	ErrAddOnAddOnDealExpired                                  = "add_on.add_on_deal_expired"                                       // The expired add on deal can't be edited
	ErrAddOnAddOnDealItemUnderBlockCategories                 = "add_on.add_on_deal_item_under_block_categories"                   // This product is under a category that is prohibited from promotions due to regulations in your country
	ErrAddOnAddOnDealItemUnderBtoc                            = "add_on.add_on_deal_item_under_btoc"                               // There are B2C items that can't be added into your promotion
	ErrAddOnAddOnDealNameInvalid                              = "add_on.add_on_deal_name_invalid"                                  // The name of add on deal can't exceed 25 characters
	ErrAddOnAddOnDealNotExists                                = "add_on.add_on_deal_not_exists"                                    // The add on deal not existed
	ErrAddOnAddOnDealOverSize                                 = "add_on.add_on_deal_over_size"                                     // A maximum of 1000 add on deals can be created
	ErrAddOnAddOnDealSetError                                 = "add_on.add_on_deal_set_error"                                     // Please input per gift number between 1 to 50 | Please input add on deal purchase limit between 1 to 100
	ErrAddOnAddOnDealStartTimeError                           = "add_on.add_on_deal_start_time_error"                              // The start time should be 1 hour later than current time
	ErrAddOnAddOnDealTimeError                                = "add_on.add_on_deal_time_error"                                    // The promotion period can't exceed 6 months
	ErrAddOnAddOnDealTypeError                                = "add_on.add_on_deal_type_error"                                    // Please input correct promotion type: Add-on discount=0; Gift with min.spend=1
	ErrAddOnAddOnDeleteError                                  = "add_on.add_on_delete_error"                                       // Only upcoming add on deal can be deleted
	ErrAddOnAddOnDuplicateItem                                = "add_on.add_on_duplicate_item"                                     // The item is already existed in main item | The item is already existed in sub item
	ErrAddOnAddOnEditError                                    = "add_on.add_on_edit_error"                                         // Please input correct promotion purchase limit | The purchase min spend and per gift num is not needed in add on discount | The purchase limit is not needed in gift with mini.spend | The start time of upcoming add on deal can't be shorted | Promotion end time can only be changed to an earlier timing. | Please enter a discount price less than the original price | Please input correct priority
	ErrAddOnAddOnEndError                                     = "add_on.add_on_end_error"                                          // Only the ongoing promotion can be ended
	ErrAddOnAddOnItemDeleteError                              = "add_on.add_on_item_delete_error"                                  // At least one sub item should be added in add on deal
	ErrAddOnAddOnItemError                                    = "add_on.add_on_item_error"                                         // The item does not existed
	ErrAddOnAddOnItemExceedDiscountLimitError                 = "add_on.add_on_item_exceed_discount_limit_error"                   // The overall item level promotion limit has been reached.
	ErrAddOnAddOnItemInvalid                                  = "add_on.add_on_item_invalid"                                       // The status of this product is abnormal
	ErrAddOnAddOnItemNoStock                                  = "add_on.add_on_item_no_stock"                                      // The product is no stock
	ErrAddOnAddOnItemOverlap                                  = "add_on.add_on_item_overlap"                                       // This product is already in add on deal main item list | This product is already in add on deal sub item list | This product is already in other add on deal main item list | This product is already in other bundle deal
	ErrAddOnAddOnMainItemOverSize                             = "add_on.add_on_main_item_over_size"                                // A maximum of 1000 main item can be added
	ErrAddOnAddOnMpqError                                     = "add_on.add_on_mpq_error"                                          // Please input the value that greater than than the sum of each product minimum purchase quantity | The Minimum purchase quantity of this product is exceeded the purchase limit of add on deal
	ErrAddOnAddOnNoItem                                       = "add_on.add_on_no_item"                                            // This item doesn't existed
	ErrAddOnAddOnNoLogistic                                   = "add_on.add_on_no_logistic"                                        // This item doesn't has shipping channel, please set shipping channel firstly
	ErrAddOnAddOnPageNoError                                  = "add_on.add_on_page_no_error"                                      // Please input page number between 1-1000
	ErrAddOnAddOnPageSizeError                                = "add_on.add_on_page_size_error"                                    // Please input page type between 1-100
	ErrAddOnAddOnPurchaseLimitError                           = "add_on.add_on_purchase_limit_error"                               // The purchase limit of this product should exceeded the minimum purchase quantity
	ErrAddOnAddOnSetError                                     = "add_on.add_on_set_error"                                          // Please input purchase min spend
	ErrAddOnAddOnShippingChannelError                         = "add_on.add_on_shipping_channel_error"                             // This product does not share a common shipping channel with the first enabled product.
	ErrAddOnAddOnStatusError                                  = "add_on.add_on_status_error"                                       // Please input correct promotion status: all, ongoing, upcoming, expired | Please input correct promotion status:all: 0, ongoing: 1, upcoming: 2, expired: 3
	ErrAddOnAddOnSubItemOverSize                              = "add_on.add_on_sub_item_over_size"                                 // A maximum of 100 sub item can be added
	ErrAddOnAddOnUpdateError                                  = "add_on.add_on_update_error"                                       // The {value} of ongoing add on deal can't be edited
	ErrAdsAccountStatusNotNormal                              = "ads.account.status_not_normal"                                    // Account status is not normal.
	ErrAdsAdsAccountInvalidStatus                             = "ads.ads_account.invalid_status"                                   // Ads account status is Invalid.
	ErrAdsApiHttpMethodNotAllowed                             = "ads.api.http_method_not_allowed"                                  // Only the supported HTTP method for the API is allowed.
	ErrAdsCampaignErrorDailyBudgetRange                       = "ads.campaign.error_daily_budget_range"                            // The budget set is invalid.
	ErrAdsCampaignErrorDateSetting                            = "ads.campaign.error_date_setting"                                  // The start/end date of campaign should be larger or equal to today. End date to be empty for unlimited duration.
	ErrAdsCampaignInvalidStartDate                            = "ads.campaign.invalid_start_date"                                  // The start date has to be today for no end date (unlimited duration).
	ErrAdsDeprecated                                          = "ads_deprecated"                                                   // This feature has been deprecated
	ErrAdsEditInvalidAction                                   = "ads.edit.invalid_action"                                          // Edit Action is invalid as per the current status of the campaign.
	ErrAdsErrorNotWhitelistedForProductGms                    = "ads_error_not_whitelisted_for_product_gms"                        // This shop is not whitelisted for Product GMS.
	ErrAdsErrorProductGmsCampaignNotFound                     = "ads_error_product_gms_campaign_not_found"                         // Product GMS Campaign cannot be found.
	ErrAdsInvalidRoasTarget                                   = "ads_invalid_roas_target"                                          // ROAS Target should be greater than equal to zero.
	ErrAdsItemAdultItem                                       = "ads.item.adult_item"                                              // Item is an adult item cannot have this campaign setting.
	ErrAdsItemBlockedForPromotion                             = "ads.item.blocked_for_promotion"                                   // Item is blocked for promotion.
	ErrAdsItemInvalidStatus                                   = "ads.item.invalid_status"                                          // Item status is invalid.
	ErrAdsItemSoldOut                                         = "ads.item.sold_out"                                                // Item is sold out.
	ErrAdsKeywordAlreadyExists                                = "ads.keyword.already_exists"                                       // Keyword already exists.
	ErrAdsKeywordIsBlacklisted                                = "ads.keyword.is_blacklisted"                                       // Keyword is Blacklisted.
	ErrAdsKeywordIsReserved                                   = "ads.keyword.is_reserved"                                          // Keyword is reserved.
	ErrAdsKeywordItemNotMatch                                 = "ads.keyword.item_not_match"                                       // Item doesn't belong to the shop.
	ErrAdsKeywordNoUpdate                                     = "ads.keyword.no_update"                                            // Keyword has no update.
	ErrAdsKeywordNotFound                                     = "ads.keyword.not_found"                                            // Keyword is not found.
	ErrAdsKeywordPriceTooHigh                                 = "ads.keyword.price_too_high"                                       // Price of Keyword is too high.
	ErrAdsKeywordPriceTooLow                                  = "ads.keyword.price_too_low"                                        // Price of Keyword is too low.
	ErrAdsPerformanceErrorDateInFuture                        = "ads.performance.error_date_in_future"                             // The date must be today or earlier than today up to 6 months ago.
	ErrAdsPerformanceErrorDateRangeTooLong                    = "ads.performance.error_date_range_too_long"                        // Date range can't be longer than 1 month.
	ErrAdsPerformanceErrorDateTooOld                          = "ads.performance.error_date_too_old"                               // The date can't be earlier than 6 months ago.
	ErrAdsPerformanceErrorSameStartEndDates                   = "ads.performance.error_same_start_end_dates"                       // Start date can't be equal to end date. For single date's hourly performance data, please query hourly api.
	ErrAdsPerformanceErrorStartLaterThanEnd                   = "ads.performance.error_start_later_than_end"                       // The start date can't be later than the end date.
	ErrAdsRateLimitCampaignLevel                              = "ads.rate_limit.campaign_level"                                    // Too many requests at the moment, please try again later.
	ErrAdsRateLimitExceedApi                                  = "ads.rate_limit.exceed_api"                                        // Too many requests at the moment, please try again later.
	ErrAdsRateLimitExceedPartnerApi                           = "ads.rate_limit.exceed_partner_api"                                // Too many requests, please reduce the request rate and try again later.
	ErrAdsRateLimitExceedShopApi                              = "ads.rate_limit.exceed_shop_api"                                   // Too many requests for the shop at the moment, please try again later.
	ErrAdsShopCampaignNotMatch                                = "ads.shop.campaign_not_match"                                      // Invalid campaign ID for this shop.
	ErrAdsShopInHolidayMode                                   = "ads.shop.in_holiday_mode"                                         // Shop is in holiday mode.
	ErrAdsShopStatusNotNormal                                 = "ads.shop.status_not_normal"                                       // Shop status is not normal.
	ErrBookingBookingListInvalidTime                          = "booking.booking_list_invalid_time"                                // Start time must be earlier than end time and diff in 15days.
	ErrBundleBundeDealMultiTierInvalidDiscount                = "bundle.bunde_deal_multi_tier_invalid_discount"                    // Higher tier should have a higher or same discount or bundle price
	ErrBundleBundleDealDeleteError                            = "bundle.bundle_deal_delete_error"                                  // The expired item in bundle deal can't be delete
	ErrBundleBundleDealDeleteFailed                           = "bundle.bundle_deal_delete_failed"                                 // Only the upcoming bundle deal can be deleted
	ErrBundleBundleDealDiscountValueInvalidRange              = "bundle.bundle_deal_discount_value_invalid_range"                  // Discount value should be within Price range.
	ErrBundleBundleDealEndTimeError                           = "bundle.bundle_deal_end_time_error"                                // Promotion end time can only be changed to an earlier timing.
	ErrBundleBundleDealErrorTotalCountLimit                   = "bundle.bundle_deal_error_total_count_limit"                       // A maximum of 1000 bundle deals can be created
	ErrBundleBundleDealErrorUnknown                           = "bundle.bundle_deal_error_unknown"                                 // Unknown error occurred in the service
	ErrBundleBundleDealExceedRemainingTimeLimit               = "bundle.bundle_deal_exceed_remaining_time_limit"                   // The bundle deal end time must earlier than the max allowed current time
	ErrBundleBundleDealFixPriceInvalidRange                   = "bundle.bundle_deal_fix_price_invalid_range"                       // Fix price should be within Price range.
	ErrBundleBundleDealInvaildDiscountPercentage              = "bundle.bundle_deal_invaild_discount_percentage"                   // Discount percentage should not less than 1%% and not more than 99%%
	ErrBundleBundleDealItemAlreadyExist                       = "bundle.bundle_deal_item_already_exist"                            // The item is already added
	ErrBundleBundleDealItemError                              = "bundle.bundle_deal_item_error"                                    // Failed to set bundle deal item
	ErrBundleBundleDealItemExceedDiscountLimitError           = "bundle.bundle_deal_item_exceed_discount_limit_error"              // The overall item level promotion limit has been reached.
	ErrBundleBundleDealItemInvalid                            = "bundle.bundle_deal_item_invalid"                                  // The item is not exist or abnormal
	ErrBundleBundleDealItemNoCommonLogistics                  = "bundle.bundle_deal_item_no_common_logistics"                      // bundle deal item has no common logistics
	ErrBundleBundleDealItemNotExist                           = "bundle.bundle_deal_item_not_exist"                                // Please update item of bundel deal
	ErrBundleBundleDealItemTotalError                         = "bundle.bundle_deal_item_total_error"                              // A maximum of 1000 items in bundle deal can be created
	ErrBundleBundleDealItemUnderBlockCategories               = "bundle.bundle_deal_item_under_block_categories"                   // 	This product is under a category that is prohibited from promotions due to regulations in your country
	ErrBundleBundleDealItemUnderBtocTag                       = "bundle.bundle_deal_item_under_btoc_tag"                           // There are B2C items that can't be added into your promotion
	ErrBundleBundleDealMaxSpanBeyond                          = "bundle.bundle_deal_max_span_beyond"                               // The promotion period can't exceed 6 months
	ErrBundleBundleDealMinSpanWithin                          = "bundle.bundle_deal_min_span_within"                               // The promotion period must be greater than 1 hour
	ErrBundleBundleDealMultiTierFeatureToggle                 = "bundle.bundle_deal_multi_tier_feature_toggle"                     // Cannot create multi-tier bundle deals
	ErrBundleBundleDealMultiTierInvalidDiscountValuePerItem   = "bundle.bundle_deal_multi_tier_invalid_discount_value_per_item"    // higher tier should have higher discount per item (discount per item = discount value off / item amount)
	ErrBundleBundleDealMultiTierInvalidDiscountWithEqualError = "bundle.bundle_deal_multi_tier_invalid_discount_with_equal_error"  // Higher tier discount must be equivalent or greater than lower tier
	ErrBundleBundleDealMultiTierInvalidFixPricePerItem        = "bundle.bundle_deal_multi_tier_invalid_fix_price_per_item"         // Higher tier should have lower price per item (price per item = fixed bundle price / item amount)
	ErrBundleBundleDealNameLength                             = "bundle.bundle_deal_name_length"                                   // Name length can't be less than 1 or greater than 25
	ErrBundleBundleDealNoAuthUpdate                           = "bundle.bundle_deal_no_auth_update"                                // The shopee created bundle deal can't be edited
	ErrBundleBundleDealNoCommonShippingChannel                = "bundle.bundle_deal_no_common_shipping_channel"                    // This product does not share a common shipping channel with the first enabled product.
	ErrBundleBundleDealNoShippingChannel                      = "bundle.bundle_deal_no_shipping_channel"                           // This product does not set shipping channel.
	ErrBundleBundleDealNotExists                              = "bundle.bundle_deal_not_exists"                                    // Bundle deal doesn't existed
	ErrBundleBundleDealOngoingPromotionInfoLock               = "bundle.bundle_deal_ongoing_promotion_info_lock"                   // Cannot change rule_type, fix_price, discount_percentage, discount_value or tiers number if Bundle deal is ongoing.
	ErrBundleBundleDealOngoingPromotionInfoLockMinAmount      = "bundle.bundle_deal_ongoing_promotion_info_lock_min_amount"        // Cannot change min_amount if Bundle deal is ongoing and has enabled items
	ErrBundleBundleDealOverMinAmount                          = "bundle.bundle_deal_over_min_amount"                               // The item amount can't exceed 3000
	ErrBundleBundleDealOverTier                               = "bundle.bundle_deal_over_tier"                                     // Max bundle tier allowed is 3
	ErrBundleBundleDealOverlapError                           = "bundle.bundle_deal_overlap_error"                                 // The item is already in add on deal | The item is already in other bundle deal
	ErrBundleBundleDealRuleTypeError                          = "bundle.bundle_deal_rule_type_error"                               // The rule type of ongoing bundle deal can't be edited
	ErrBundleBundleDealRuleTypeMixUse                         = "bundle.bundle_deal_rule_type_mix_use"                             // Please use consistent rule type
	ErrBundleBundleDealRuleTypeValue                          = "bundle.bundle_deal_rule_type_value"                               // Please use correct bundle deal rule type
	ErrBundleBundleDealRuleTypeValueError                     = "bundle.bundle_deal_rule_type_value_error"                         // Wrong rule value
	ErrBundleBundleDealStartTimeError                         = "bundle.bundle_deal_start_time_error"                              // Promotion start time can only be changed to a later timing. | The start time should be later than current time | The start_time of ongoing promotion can't be updated
	ErrBundleBundleDealStatusError                            = "bundle.bundle_deal_status_error"                                  // Item status is abnormal
	ErrBundleBundleDealStockError                             = "bundle.bundle_deal_stock_error"                                   // Insufficient item stock
	ErrBundleBundleDealStopFailed                             = "bundle.bundle_deal_stop_failed"                                   // Only the ongoing bundle deal can be ended
	ErrBundleBundleDealTierInvalid                            = "bundle.bundle_deal_tier_invalid"                                  // The discount off for tier 1 must be < 90%%
	ErrBundleBundleDealUnderMinAmount                         = "bundle.bundle_deal_under_min_amount"                              // The item amount should be greater than 1
	ErrBundleBundleDealUnderMinAmountMultiTier                = "bundle.bundle_deal_under_min_amount_multi_tier"                   // The item amount should be greater than the previous tier
	ErrBundleBundleExpired                                    = "bundle.bundle_expired"                                            // bundle deal expired can not update
	ErrBundleProductNotExist                                  = "bundle.product_not_exist"                                         // product id not exists
	ErrChannelErrorServer                                     = "channel_error_server"                                             // Get installment channel info failed.
	ErrCommonBatchApiAllFailed                                = "common.batch_api_all_failed"                                      // Failed, please check result_list for more details.
	ErrCommonErrorNotFound                                    = "common.error_not_found"                                           // The information you queried is not found. | supplier order item income not found   | Unknown error: ***. Something wrong. Please try later.
	ErrCommonErrorParam                                       = "common.error_param"                                               // Wrong parameters, detail: {msg}. | The Wrong parameters, detail: {msg}.
	ErrCommonErrorPermission                                  = "common.error_permission"                                          // You don't have permission for the operation
	ErrCommonErrorServer                                      = "common.error_server"                                              // Something wrong. Please try later. | Something is wrong. Please try later.
	ErrCommonErrorShop                                        = "common.error_shop"                                                // Shopid is invalid | Shopid is invalid.
	ErrCommonInvalidShop                                      = "common.invalid_shop"                                              // Shop id is invalid. Please check your shop.
	ErrDecodedFailedError                                     = "decoded_failed_error"                                             // The escrow item extinfo decoded failed.
	ErrDiscountDiscountEndTimeSmallerThanNow                  = "discount.discount_end_time_smaller_than_now"                      // Discount end time should later than the current time.
	ErrDiscountDiscountEndTimeSmallerThanStartTime            = "discount.discount_end_time_smaller_than_start_time"               // Discount end time should later than start time.
	ErrDiscountDiscountIsEnd                                  = "discount.discount_is_end"                                         // The end discount cannot be updated or deleted.
	ErrDiscountDiscountIsOngoing                              = "discount.discount_is_ongoing"                                     // Only the ongoing discount can be ended.
	ErrDiscountDiscountItemNotExist                           = "discount.discount_item_not_exist"                                 // The item is not in this discount, please check it again.
	ErrDiscountDiscountNeedPromotionPrice                     = "discount.discount_need_promotion_price"                           // Please input the discount price.
	ErrDiscountDiscountNotOngoing                             = "discount.discount_not_ongoing"                                    // The ongoing discount cannot be deleted.
	ErrDiscountDiscountPeriodTooLong                          = "discount.discount_period_too_long"                                // Promotion duration must be within 180 days.
	ErrDiscountDiscountPeriodTooShort                         = "discount.discount_period_too_short"                               // The end time should be later than 1 hour of start time.
	ErrDiscountDiscountStartTimeSmallerThanNow                = "discount.discount_start_time_smaller_than_now"                    // Discount start time is less than the current time | Discount start time is less than the current time.
	ErrDiscountEndDiscountTimeError                           = "discount.end_discount_time_error"                                 // Only can end the discount after 1 hour of start time.
	ErrDiscountErrorAssignedPromoStock                        = "discount.error_assigned_promo_stock"                              // The reserved stock of added products or disabled products can't be edited
	ErrDiscountErrorB2cItemNotAllowed                         = "discount.error_b2c_item_not_allowed"                              // This product is a B2C product that is prohibited from promotions in FBS shops
	ErrDiscountErrorHolidayMode                               = "discount.error_holiday_mode"                                      // Cannot add discount as holiday mode is turned on
	ErrDiscountErrorItemAbnormal                              = "discount.error_item_abnormal"                                     // The item is abnormal
	ErrDiscountErrorItemNotEnoughStock                        = "discount.error_item_not_enough_stock"                             // Promotion Sock should be less than current stock
	ErrDiscountErrorItemUnderBlockCategories                  = "discount.error_item_under_block_categories"                       // This product is under a category that is prohibited from promotions due to regulations in your country
	ErrDiscountErrorStockLessThanMpq                          = "discount.error_stock_less_than_mpq"                               // Promotion Sock should be more than {number}
	ErrDiscountErrorUpdateAdminDiscount                       = "discount.error_update_admin_discount"                             // The discount created by admin can't be edited
	ErrDiscountErrorUpdateStreamingDiscount                   = "discount.error_update_streaming_discount"                         // The discount created by livestreaming can't be edited
	ErrDiscountExceedDiscountItemBatchSize                    = "discount.exceed_discount_item_batch_size"                         // You can only add up to 50 items at a time.
	ErrDiscountExceedMaxDiscountCount                         = "discount.exceed_max_discount_count"                               // The upcoming and ongoing discount cannot exceed 1000.
	ErrDiscountExtendDiscountEndTimeError                     = "discount.extend_discount_end_time_error"                          // Promotion end time can only be changed to an earlier timing.
	ErrDiscountFailModelPriceCheck                            = "discount.fail_model_price_check"                                  // The difference between the maximum price and the minimum price in one item is too large.
	ErrDiscountGetModelOrginPriceError                        = "discount.get_model_orgin_price_error"                             // There are something wrong when get the original price of model, please try again.
	ErrDiscountItemExceedDiscountLimit                        = "discount.item_exceed_discount_limit"                              // The overall item level promotion limit has been reached.
	ErrDiscountItemIdNotExist                                 = "discount.item_id_not_exist"                                       // The Item id you provided is not found, please check your item id.
	ErrDiscountItemIdRepeated                                 = "discount.item_id_repeated"                                        // The item id can't be upload due to it has duplicated.
	ErrDiscountItemInOtherPromotion                           = "discount.item_in_other_promotion"                                 // Item is already in other overlapped discount.
	ErrDiscountItemInPromotion                                = "discount.item_in_promotion"                                       // Item is already in this discount.
	ErrDiscountItemInPromotionTooMany                         = "discount.item_in_promotion_too_many"                              // A maximum of 1000 items can be uploaded in one promotion
	ErrDiscountItemInPromotionTooManyWhitelist                = "discount.item_in_promotion_too_many_whitelist"                    // A maximum of 15,000 items can be uploaded in one promotion
	ErrDiscountItemNeedModelId                                = "discount.item_need_model_id"                                      // Please input model id for the item with models.
	ErrDiscountItemNotInPromotion                             = "discount.item_not_in_promotion"                                   // Item is not found in this discount.
	ErrDiscountItemPurchaseLimitError                         = "discount.item_purchase_limit_error"                               // Purchase limit of this product must be greater than {number}.
	ErrDiscountItemStatusAbnormal                             = "discount.item_status_abnormal"                                    // The item status is abnormal
	ErrDiscountItemStatusInvalid                              = "discount.item_status_invalid"                                     // item status is invalid.
	ErrDiscountModelIdNotExist                                = "discount.model_id_not_exist"                                      // The model id you provided is not found. Please check your model id.
	ErrDiscountModelIdRepeated                                = "discount.model_id_repeated"                                       // The model id can't be upload due to it has duplicated.
	ErrDiscountModelInPromotion                               = "discount.model_in_promotion"                                      // The model that you provided is already in this discount.
	ErrDiscountModelInPromotionTooMany                        = "discount.model_in_promotion_too_many"                             // A maximum of 50,000 models can be uploaded in one promotion
	ErrDiscountModelNotInPromotion                            = "discount.model_not_in_promotion"                                  // The model that you provided is not in this discount.
	ErrDiscountOngoingDiscountPeriodTooLong                   = "discount.ongoing_discount_period_too_long"                        // Promotion period must be less than 180 days, please end it within 30 days.
	ErrDiscountPromotionPriceForVn                            = "discount.promotion_price_for_vn"                                  // The discount price has exceeded 50% of original price.
	ErrDiscountPromotionPriceHigherInputPrice                 = "discount.promotion_price_higher_input_price"                      // Discount price is greater than original price.
	ErrDiscountPromotionPriceTooLow                           = "discount.promotion_price_too_low"                                 // Discount price is less than the minimum national amount.
	ErrDiscountPromotoinPriceTooHigh                          = "discount.promotoin_price_too_high"                                // Discount price is greater than the maximum national amount.
	ErrDiscountShortDiscountStartTime                         = "discount.short_discount_start_time"                               // Promotion start time can only be changed to a later timing.
	ErrDiscountUpdateOngoingDiscountStartTime                 = "discount.update_ongoing_discount_start_time"                      // The start time can‘t be changed when discount is in ongoing status.
	ErrDisputeCancelIneligible                                = "dispute.cancel.ineligible"                                        // Cannot cancel dispute when compensation status is already in COMPENSATION_APPROVED | Cannot cancel dispute when compensation status is already in COMPENSATION_CANCELLED | Cannot cancel dispute when compensation status is not in COMPENSATION_REQUESTED | Cannot cancel dispute when compensation status is already in COMPENSATION_REJECTED | Cannot cancel dispute when return status is not in ACCEPTED | Cannot cancel dispute because return SN is not eligible for compensation dispute
	ErrDisputeRequestParamError                               = "dispute.request.param.error"                                      // Can't raise dispute because the dispute reason is empty | dispute.request.param.error | Input invalid email=[] | can't raise dispute because the requirement and dispute reason ID does not match | Invalid images, images max size is 3 per module | Unable to raise dispute as mandatory module index is missing | can't raise dispute because mandatory evidence is missing | Can't raise dispute because the return_sn and dispute_reason_id does not match. | Invalid images, images max size is 3. | Key: 'XXX' Error:Field validation for 'XX' failed on the 'required' tag | The return_sn you requested was not queried, please check it | return_sn is a required parameter
	ErrDisputeRrNotAllow                                      = "dispute.rr.not.allow"                                             // dispute.rr.not.allow
	ErrERRORMERCHANTUNSETCURRENCYINFO                         = "ERROR_MERCHANT_UNSET_CURRENCY_INFO"                               // The merchant has not set the currency. Please log in to CBSC to make the settings.
	ErrErrData                                                = "err_data"                                                         // Cannot accept your own offer.
	ErrError                                                  = "error"                                                            // no supported warehouse address
	ErrErrorAccount                                           = "error_account"                                                    // Core server error.
	ErrErrorAlreadyCompleted                                  = "error_already_completed"                                          // Upload already completed.
	ErrErrorArchivedOrderUpdate                               = "error_archived_order_update"                                      // Cannot update old order
	ErrErrorAttribute                                         = "error attribute"                                                  // Your attribute info is invalid. Please fill in the category related attributes with correct attribute value.
	ErrErrorAttributeFdaError                                 = "error_attribute_fda_error"                                        // Attribute FDA value is invalid
	ErrErrorAuth                                              = "error_auth"                                                       // Invalid access_token. | No permission to current api. | Your shop can not use model level dts | The location_id input is not matched the shop's location_id(more/wrong). Please double check. | Lack of location_id, please double check. | Please wait for the holiday mode set then to edit item. Please try later. | Total stock must be more than reserved stock. | You do not have right to use seller location_id, please only fill seller_stock filed. | Stock should be larger than reserved stock. | cnsc shop not upgraded, can not edit item. | The current item belong to the full FBS or B2C shop, so normal stock must be equal to 0 | The registered phone number of your shop is abnormal (Malaysian phone number was detected), please update the phone number immediately. | shop is upgrading , can not operate | not all shop is upgraded and confirmed, can not edit item | partner does not have permission to operate shop | can not publish product to sip affiliated shop | shop is not upgraded or lack service fee rate | tw shop not smid verified | Invalid partner_id or shopid. | You have no valid partnerid. | You are not authorized
	ErrErrorAuthModelIsPff                                    = "error_auth_model_is_pff"                                          // Item/Model Fulfillment by Shopee  can not be update.
	ErrErrorAuthProductIsPff                                  = "error_auth_product_is_pff"                                        // Item/Model Fulfillment by Shopee  can not be update.
	ErrErrorAuthShopNotFound                                  = "error_auth_shop_not_found"                                        // Shop is not found.
	ErrErrorBanned                                            = "error_banned"                                                     // Core server error.
	ErrErrorBoostItemAllFailed                                = "error_boost_item_all_failed"                                      // Boost item all failed.
	ErrErrorBoostItemFailed                                   = "error_boost_item_failed"                                          // Boost item failed. please try later.
	ErrErrorBrand                                             = "error brand"                                                      // Your brand is invalid. Please use the valid brand under your category.
	ErrErrorBrandForbidden                                    = "error_brand_forbidden"                                            // The brand is forbidden.
	ErrErrorBusi                                              = "error_busi"                                                       // logistic must be free | The merchant/shop has multi warehouse, please input location id
	ErrErrorBusiAddGlobalItemFailed                           = "error_busi_add_global_item_failed"                                // Add global item failed, please try later. {{.error_info}}
	ErrErrorBusiAddItemFailed                                 = "error_busi_add_item_failed"                                       // Add item failed. please try later {{.error_info}}.
	ErrErrorBusiAttributeError                                = "error_busi_attribute_error"                                       // Attribute NCC value is invalid | Attribute NCC is mandatory | Attribute BSMI value is invalid | Attribute BSMI is mandatory
	ErrErrorBusiBlacklist                                     = "error_busi_blacklist"                                             // This brand name can't be used in several selected categories. You can contact Shopee CS for brand appeal offline if need.
	ErrErrorBusiBlacklistNeedLicense                          = "error_busi_blacklist_need_license"                                // Risky brand name, please provide related license and brand registration website to register this specific brand.
	ErrErrorBusiCannotDeleteAllModel                          = "error_busi_cannot_delete_all_model"                               // Cannot delete all models by this api.
	ErrErrorBusiCannotDeleteDefaultModel                      = "error_busi_cannot_delete_default_model"                           // Can't delete default model.
	ErrErrorBusiCannotDeleteGlobalItem                        = "error_busi_cannot_delete_global_item"                             // Delete global item failed,{{.error_info}}ｼ継lease try later.
	ErrErrorBusiCannotDelistReviewingOrBannedItem             = "error_busi_cannot_delist_reviewing_or_banned_item"                // Banned and Reviewing Products cannot be delisted
	ErrErrorBusiCannotEditVsku                                = "error_busi_cannot_edit_vsku"                                      // Can not use OpenAPI to edit/create VSKU, please connect with your manager
	ErrErrorBusiCannotGetCategoryPath                         = "error_busi_cannot_get_category_path"                              // Get category path failed. please try later.
	ErrErrorBusiCannotPublish                                 = "error_busi_cannot_publish"                                        // Publish failed,  {{.error_info}}.
	ErrErrorBusiCannotUpdateField                             = "error_busi_cannot_update_field"                                   // Update item failed , {{.error_info}}.
	ErrErrorBusiCannotUpdateModelForNoTierItem                = "error_busi_cannot_update_model_for_no_tier_item"                  // Item without tier_variation. Please use init_tier_variation api to upgrade the structure.
	ErrErrorBusiDeleteGlobalDefaultModelFailed                = "error_busi_delete_global_default_model_failed"                    // Can't delete global default model.
	ErrErrorBusiDeleteGlobalModelFalied                       = "error_busi_delete_global_model_falied"                            // Delete global model failed, please try later, {{.error_info}}.
	ErrErrorBusiDuplicated                                    = "error_busi_duplicated"                                            // Brand already exists. If you want to add available category for this brand, you can change the category.
	ErrErrorBusiErrorSipError                                 = "error_busi_error_sip_error"                                       // Call sip api with error.
	ErrErrorBusiGlobalItemDimensionOverLimit                  = "error_busi_global_item_dimension_over_limit"                      // Global item dimension is over limit.
	ErrErrorBusiGlobalItemModelNotAllowEmpty                  = "error_busi_global_item_model_not_allow_empty"                     // Model can not be empty.
	ErrErrorBusiGlobalItemModelPriceNoMatch                   = "error_busi_global_item_model_price_no_match"                      // Global_model_id and global_item_id not match.
	ErrErrorBusiGlobalItemModelStockNoMatch                   = "error_busi_global_item_model_stock_no_match"                      // Global_model_id and global_item_id not match.
	ErrErrorBusiGlobalItemNotFound                            = "error_busi_global_item_not_found"                                 // Global item not found.
	ErrErrorBusiGlobalItemPriceShouldBiggerThanZero           = "error_busi_global_item_price_should_bigger_than_zero"             // Original_price must bigger than 0.
	ErrErrorBusiGlobalItemSkuOverLimit                        = "error_busi_global_item_sku_over_limit"                            // Global item sku is over limit.
	ErrErrorBusiGlobalItemTierOptionOverLimit                 = "error_busi_global_item_tier_option_over_limit"                    // Option of tier_variation is over limit.
	ErrErrorBusiGlobalItemWeightOverLimit                     = "error_busi_global_item_weight_over_limit"                         // Global item weight is over limit.
	ErrErrorBusiGlobalModelNotFound                           = "error_busi_global_model_not_found"                                // Global model no found.
	ErrErrorBusiGlobalTierOptionImageNotEqualToTierOneOption  = "error_busi_global_tier_option_image_not_equal_to_tier_one_option" // Image num of tier_variation option is not equal to option num of tier_variation.
	ErrErrorBusiGlobalTierVariationOverLimit                  = "error_busi_global_tier_variation_over_limit"                      // Global item tier_variation is over limit.
	ErrErrorBusiInvalidAccountStatus                          = "error_busi_invalid_account_status"                                // Account status is invalid.
	ErrErrorBusiInvalidShopStatus                             = "error_busi_invalid_shop_status"                                   // Shop status invalid.
	ErrErrorBusiItemDaysToShipInvalid                         = "error_busi_item_days_to_ship_invalid"                             // Global item days_to_ship is invalid.
	ErrErrorBusiItemNotFound                                  = "error_busi_item_not_found"                                        // Global_item_id not found.
	ErrErrorBusiItemPriceHigherThanPrice                      = "error_busi_item_price_higher_than_price"                          // Sip_item_price should less than the original_price.
	ErrErrorBusiItemStatusInvalid                             = "error_busi_item_status_invalid"                                   // Invalid item status.
	ErrErrorBusiMpskuCategoryProhibited                       = "error_busi_mpsku_category_prohibited"                             // {{.error_info}}
	ErrErrorBusiNotAllowEditItemPrice                         = "error_busi_not_allow_edit_item_price"                             // Cannot edit sip_item_price.
	ErrErrorBusiPendingQc                                     = "error_busi_pending_qc"                                            // The brand you registered is in the inspection process. Please contact Shopee Customer Service for more information.
	ErrErrorBusiPriceLowerThenWholesalePrice                  = "error_busi_price_lower_then_wholesale_price"                      // Original_price should bigger then wholesale price
	ErrErrorBusiRegionNotSupported                            = "error_busi_region_not_supported"                                  // Unsupport region.
	ErrErrorBusiShopNotUpgradedToCnsc                         = "error_busi_shop_not_upgraded_to_cnsc"                             // SIP ASHOP has no permission to all this api or Shop not upgraded or lack of service fee rate.
	ErrErrorBusiShopRegionNotMatch                            = "error_busi_shop_region_not_match"                                 // Shop and shop region not match.
	ErrErrorBusiUpdateGlobalItemFailed                        = "error_busi_update_global_item_failed"                             // Update global item failed, please try later.
	ErrErrorBusiUpdateItemFailed                              = "error_busi_update_item_failed"                                    // Update item failed, please try later.
	ErrErrorBusiUpdatePriceFailed                             = "error_busi_update_price_failed"                                   // Update original_price failed, please try later.
	ErrErrorBusiUpdateSizeChartFailed                         = "error_busi_update_size_chart_failed"                              // Update size chart failed, please try later.
	ErrErrorBusiUpdateSockFailed                              = "error_busi_update_sock_failed"                                    // Update normal_stock failed, please try later.
	ErrErrorBusiUpdateStockFailed                             = "error_busi_update_stock_failed"                                   // Update stock failed, please check failure_list for detailed reason
	ErrErrorCannotDeleteItem                                  = "error_cannot_delete_item"                                         // Delete item failed. please try later.
	ErrErrorCannotGetPublishResult                            = "error_cannot_get_publish_result"                                  // Get publish result failed, please try later.
	ErrErrorCannotUpdatePriceInPromotion                      = "error_cannot_update_price_in_promotion"                           // Price cannot be changed when model is under promotion.
	ErrErrorCanntBeNoVariationInPromotion                     = "error_cannt_be_no_variation_in_promotion"                         // Item cannot be without model when item is under promotion.
	ErrErrorCanntChangeTierVariationInPromotion               = "error_cannt_change_tier_variation_in_promotion"                   // Structure of tier_variation cannot be edited when item is under promotion.
	ErrErrorCanntDeleteInPromotion                            = "error_cannt_delete_in_promotion"                                  // Item cannot be deleted when item is under promotion.
	ErrErrorCanntDeleteOptionInPromotion                      = "error_cannt_delete_option_in_promotion"                           // Option of tier_variation cannot be deleted when item is under promotion.
	ErrErrorCanntEditDescriptionInPromotion                   = "error_cannt_edit_description_in_promotion"                        // Description of item cannot be edited when item is under promotion..
	ErrErrorCanntEditEstimatedDaysInPromotion                 = "error_cannt_edit_estimated_days_in_promotion"                     // Days_to_ship cannot be edited when item is under promotion.
	ErrErrorCanntEditImageInPromotion                         = "error_cannt_edit_image_in_promotion"                              // Image of item cannot be edited when item is under promotion..
	ErrErrorCanntEditNameInPromotion                          = "error_cannt_edit_name_in_promotion"                               // Item_name cannot be edited when item is under promotion..
	ErrErrorCanntEditPreOrderInPromotion                      = "error_cannt_edit_pre_order_in_promotion"                          // Item cannot be changed to pre-order when item is under promotion..
	ErrErrorCanntEditPriceInPromotion                         = "error_cannt_edit_price_in_promotion"                              // Original_price cannot be edited when item is under promotion.
	ErrErrorCanntEditStockInPromotion                         = "error_cannt_edit_stock_in_promotion"                              // Normal_stock cannot be edited when item is under promotion.
	ErrErrorCanntInitTierInPromotion                          = "error_cannt_init_tier_in_promotion"                               // Tier_variation cannot be changed when item is under promotion.
	ErrErrorCanntUnlistedInPromotion                          = "error_cannt_unlisted_in_promotion"                                // Item cannot be unlisted when item is under promotion.
	ErrErrorCategory                                          = "error category"                                                   // Your category is invalid or prohibited.
	ErrErrorCategoryDts                                       = "error_category_dts"                                               // The current day_to_ship is bigger than category's max days_to_ship.
	ErrErrorCategoryIsBlock                                   = "error_category_is_block"                                          // Category is restricted
	ErrErrorCategoryLevel                                     = "error_category_level"                                             // Interal error, please contact openapi team.
	ErrErrorCategoryNotAllowed                                = "error_category_not_allowed"                                       // {{.error_info}}
	ErrErrorCategoryPathCountLimit                            = "error_category_path_count_limit"                                  // Interal error, please contact openapi team.
	ErrErrorChannelUnsupport                                  = "error_channel_unsupport"                                          // This logistic channel/logistic type doesn’t support retrieving info from Open API. Please use v2.logistics.download_shipping_document API instead.
	ErrErrorCheckLanguageFail                                 = "error_check_language_fail"                                        // check language fail, {{.error_info}}
	ErrErrorCheckLucFail                                      = "error_check_luc_fail"                                             // {{.error_info}}
	ErrErrorCnscShopBlockUpdateTierVariation                  = "error_cnsc_shop_block_update_tier_variation"                      // CNCS shop can not modify tier_variation structure.
	ErrErrorConfig                                            = "error_config"                                                     // Something is wrong, please try again later.
	ErrErrorData                                              = "error_data"                                                       // parse data failed | data not exist | error data | File cannot be opened | lack of param | Something wrong. Please try later. [20] | Inner error, please try later. [5]. | Query offer info failed. Please try later. | Query order info failed. Please try later. | Query proof info failed. Please try later. | Query return info failed. Please try later. | Get user info failed. Please try later. | The return you queried doesn't exist. | The return case is missing initial evidence from the buyer and will be auto-cancelled if the buyer does not provide evidence by the due date. | Invalid return status: {status} | Shopee is reviewing the case and will get back to you. | Shopee is reviewing the case and will get back to you.			 | Type of return does not allow seller to offer refund | Accept offer is not available for this return | The description cannot exceed 200 characters | The last session(session_id:{{sid}}) is still ongoing, please end it first | You do not have permission to livestream | The description contains sensitive content | The title contains sensitive content | The session(session_id:{{sid}}) is not exist | The title cannot exceed 200 characters | The user(user_id:{{uid}}) has been frozen | The session(session_id:{{sid}}) is not belong to you | The session(session_id:{{sid}}) is ended | The session(session_id:{{sid}}) is ongoing | Cannot end session because it's started through Shopee App or Live PC | The session(session_id:{{sid}}) is not ongoing | Item(item_id:{{item_id}}) is duplicated in the request | Item(item_id:{{item_id}}) is not exist | You are not allowed to view my like items | Too many items for current session | Item(item_id:{{item_id}}) is not exist in the shopping bag but exist in this request, please use add API If you want to add item | Item(item_id:{{item_id}}) is exist in the shopping bag but not exist in this request, please use delete API If you want to delete item | The item set {{item_set_id}} you provided is not exist | The content cannot exceed 150 characters | The content contains sensitive content | You can not ban yourself
	ErrErrorDataCheck                                         = "error_data_check"                                                 // Failed to change Shop Name for shop name have been changed in recent 30 days. Please change the shop name at least 30 days later than your last modify. | Failed to change shop name for shop name should be between {min} and {max} characters. Please correct and try it again. | This name is not applicable. Please choose another name. | Failed to change shop logo for the image url is not Shopee image url. Please use Shopee image url. | Failed to change shop logo for something wrong happened with the image url. Please try it again later. | This content is not applicable. Please check it before uploading
	ErrErrorDatabase                                          = "error_database"                                                   // Query data from database failed. | Database connection issue. Please try again later.
	ErrErrorDeleted                                           = "error_deleted"                                                    // Core server error.
	ErrErrorDescHashTagOverLimit                              = "error_desc_hash_tag_over_limit"                                   // Count of hash tags is more than 18
	ErrErrorDescImageNoPass                                   = "error_desc_image_no_pass"                                         // {{.error_info}}
	ErrErrorDescLenNoPass                                     = "error_desc_len_no_pass"                                           // {{.error_info}}
	ErrErrorDescLengthMinLimit                                = "error_desc_length_min_limit"                                      // Description length is less than the min limit.
	ErrErrorDisabled                                          = "error_disabled"                                                   // Core server error.
	ErrErrorDtsRangeNotAllowed                                = "error_dts_range_not_allowed"                                      // {{.error_info}}
	ErrErrorDuplicateModelid                                  = "error_duplicate_modelid"                                          // The model_id is duplicate, model_id:{{.model_id}}
	ErrErrorDuplicatedBrand                                   = "error_duplicated_brand"                                           // Brand already exists
	ErrErrorDuplicatedVariationOption                         = "error_duplicated_variation_option"                                // Tier variation option duplicated.
	ErrErrorEditItemPriceForItemHasModel                      = "error_edit_item_price_for_item_has_model"                         // Can't edit item price directly while item has models.
	ErrErrorEditItemStockForItemHasModel                      = "error_edit_item_stock_for_item_has_model"                         // Can't to edit item stock directly while item has models.
	ErrErrorEstimatedDaysLimit                                = "error_estimated_days_limit"                                       // Days_to_ship limitation exceeded.
	ErrErrorException                                         = "error_exception"                                                  // Core server error.
	ErrErrorExist                                             = "error_exist"                                                      // Core server error.
	ErrErrorExpire                                            = "error_expire"                                                     // Core server error.
	ErrErrorFileSize                                          = "error_file_size"                                                  // File size is too large. Video size should be less than 30M.
	ErrErrorFlashSaleDaysToShipLock                           = "error_flash_sale_days_to_ship_lock"                               // Days_to ship cannot be changed when item is under ongoing/upcoming flash sale.
	ErrErrorForbiddenCategory                                 = "error_forbidden_category"                                         // The category is forbidden.
	ErrErrorFraud                                             = "error_fraud"                                                      // Core server error.
	ErrErrorGetBoostedItemInfoFailed                          = "error_get_boosted_item_info_failed"                               // Get boosted item infomation failed. please try later.
	ErrErrorGetGlobalItemLimitFailed                          = "error_get_global_item_limit_failed"                               // Get global item limit failed. please try later.
	ErrErrorGetGlobalItemListFailed                           = "error_get_global_item_list_failed"                                // Get global item list failed. Please try later.
	ErrErrorGetItemInstallment                                = "error_get_item_installment"                                       // get item installment fail
	ErrErrorGetParnterTokenFailed                             = "error_get_parnter_token_failed"                                   // Cannot get partner token.
	ErrErrorGetShopFail                                       = "error_get_shop_fail"                                              // Get shop failed. please try later.
	ErrErrorGlobalItemModelNotMatch                           = "error_global_item_model_not_match"                                // Global_model_id and global_item_id not match.
	ErrErrorGlobalItemShopNotMatch                            = "error_global_item_shop_not_match"                                 // Global item not belong to shop.
	ErrErrorHolidayModeChangeStock                            = "error_holiday_mode_change_stock"                                  // Cannot change stock in holiday mode.
	ErrErrorHolidayOnAddItem                                  = "error_holiday_on_add_item"                                        // Shop is under vocation mode.
	ErrErrorHolidayOnDelItem                                  = "error_holiday_on_del_item"                                        // Shop is under vocation mode.
	ErrErrorImageHightNoPass                                  = "error_image_hight_no_pass"                                        // {{.error_info}}
	ErrErrorImageNumMin                                       = "error_image_num_min"                                              // {{.error_info}}
	ErrErrorImageQuantityNoPass                               = "error_image_quantity_no_pass"                                     // {{.error_info}}
	ErrErrorImageUnavailable                                  = "error_image_unavailable"                                          // Image is invalid: single image url length is less than 32. | image is unavailable | Image is unavailable
	ErrErrorImageWidthNoPass                                  = "error_image_width_no_pass"                                        // {{.error_info}}
	ErrErrorInItemPromotionDeleteLock                         = "error_in_item_promotion_delete_lock"                              // Can't delete item when item is under promotion.
	ErrErrorInItemPromotionDescriptionLock                    = "error_in_item_promotion_description_lock"                         // Can't update description when item is under promotion.
	ErrErrorInItemPromotionImageItemLock                      = "error_in_item_promotion_image_item_lock"                          // Can't update images when item is under promotion.
	ErrErrorInItemPromotionItemPriceLock                      = "error_in_item_promotion_item_price_lock"                          // Can't update price when item is under promotion.
	ErrErrorInItemPromotionNameItemLock                       = "error_in_item_promotion_name_item_lock"                           // Can't update item_name when item is under promotion.
	ErrErrorInItemPromotionNomodelToModels                    = "error_in_item_promotion_nomodel_to_models"                        // Can't to edit item stock directly while item has models.
	ErrErrorInItemPromotionRemoveModel                        = "error_in_item_promotion_remove_model"                             // Can't change tier_variation when item is under promotion.
	ErrErrorInItemPromotionUnlsitLock                         = "error_in_item_promotion_unlsit_lock"                              // Can't unlist item when item is under promotion.
	ErrErrorInModelPromotionDeleteLock                        = "error_in_model_promotion_delete_lock"                             // Can't delete item when item model is under promotion.
	ErrErrorIncalidBrand                                      = "error_incalid_brand"                                              // Brand ID or brand name required | {{.error_info}}
	ErrErrorIncalidCategory                                   = "error_incalid_category"                                           // Category IDs for L1 and L2 do not match.
	ErrErrorInner                                             = "error_inner"                                                      // Our system is taking some time to respond, please try later. | System error, please try again later or contact the OpenAPI support team. | Update item failed {{.error_info}} | Invalid stock location ID | Inner error, please try later. [0].
	ErrErrorInvalidAttribute                                  = "error_invalid_attribute"                                          // Mandatory attribute information required
	ErrErrorInvalidAttributeValue                             = "error_invalid_attribute_value"                                    // Invalid attribute value.
	ErrErrorInvalidBrand                                      = "error_invalid_brand"                                              // Invalid brand | Brand ID value should be "0". | Brand name required | Brand ID required | Brand information required
	ErrErrorInvalidCategory                                   = "error_invalid_category"                                           // Invalid category ID {{.error_info}} | Invalid category. | Category is blocked for CB seller.
	ErrErrorInvalidCategoryAttribute                          = "error_invalid_category_attribute"                                 // Category and attribute not match.
	ErrErrorInvalidDaysToShip                                 = "error_invalid_days_to_ship"                                       // Invalid days_to_ship.
	ErrErrorInvalidGlobalItemId                               = "error_invalid_global_item_id"                                     // Invalid global_item_id.
	ErrErrorInvalidItemList                                   = "error_invalid_item_list"                                          // Invalid item_id_list, please verify the ID and try again
	ErrErrorInvalidLanguage                                   = "error_invalid_language"                                           // Invalid language.
	ErrErrorInvalidLogisticInfo                               = "error_invalid_logistic_info"                                      // invalid logistic info , {{.error_info}}
	ErrErrorInvalidPartSeq                                    = "error_invalid_part_seq"                                           // Invalid part_seq.
	ErrErrorInvalidPartSize                                   = "error_invalid_part_size"                                          // Invalid part_size.
	ErrErrorInvalidPrice                                      = "error_invalid_price"                                              // Invalid price, please use the correct format
	ErrErrorInvalidPriceForLogistic                           = "error_invalid_price_for_logistic"                                 // Shipping channel cannot be enabled as product price exceeds limit.
	ErrErrorInvalidUploadId                                   = "error_invalid_upload_id"                                          // Invalid upload_id.
	ErrErrorInvitation                                        = "error_invitation"                                                 // Core server error.
	ErrErrorItemInPromotion                                   = "error_item_in_promotion"                                          // item is in promotion can not set category
	ErrErrorItemInstallmentSetting                            = "error_item_installment_setting"                                   // The credit card installment tenures is wrong.
	ErrErrorItemModelNotMatch                                 = "error_item_model_not_match"                                       // Input item_id and model_id are not match.
	ErrErrorItemNameEmpty                                     = "error_item_name_empty"                                            // Item name could not be empty.
	ErrErrorItemNameIsTooShort                                = "error_item_name_is_too_short"                                     // Item_name length is less than min limit.
	ErrErrorItemNotBelongShop                                 = "error_item_not_belong_shop"                                       // Item not belong to shop.
	ErrErrorItemNotFound                                      = "error_item_not_found"                                             // Product not found | Item_id is not found.
	ErrErrorItemOrVariationNotFound                           = "error_item_or_variation_not_found"                                // Item or model doesn't exist.
	ErrErrorItemUneditable                                    = "error_item_uneditable"                                            // Can't edit this item. item status can not support editing.
	ErrErrorLessRequiredAttribute                             = "error_less_required_attribute"                                    // Mandatory attribute information required
	ErrErrorLessRequiredBrand                                 = "error_less_required_brand"                                        // Brand information required
	ErrErrorLimit                                             = "error_limit"                                                      // Core server error.
	ErrErrorLocked                                            = "error_locked"                                                     // Core server error.
	ErrErrorMarshal                                           = "error_marshal"                                                    // Interal error, please contact openapi team
	ErrErrorMerchantNotFound                                  = "error_merchant_not_found"                                         // Merchant_id is not found.
	ErrErrorMessageCensored                                   = "error_message_censored"                                           // Core server error.
	ErrErrorModelCountOverLimit                               = "error_model_count_over_limit"                                     // The count of model is too large. Should be under 20 (50 for TW)
	ErrErrorModelDuplicateName                                = "error_model_duplicate_name"                                       // Item tier_variation have duplicate name.
	ErrErrorModelDuplicateTierVariationIndex                  = "error_model_duplicate_tier_variation_index"                       // Item have duplicate tier_index.
	ErrErrorModelEmptyTierIndexNonempty                       = "error_model_empty_tier_index_nonempty"                            // Interal error, please contact openapi team.
	ErrErrorModelInvalidModelId                               = "error_model_invalid_model_id"                                     // Invalid model_id.
	ErrErrorModelNonemptyItemtierEmptyIndex                   = "error_model_nonempty_itemtier_empty_index"                        // Item tier_index is required.
	ErrErrorModelRemoveModelInPromotion                       = "error_model_remove_model_in_promotion"                            // Model cannot be deleted when item/model is under promotion or groupbuy.
	ErrErrorModelTierIndexBound                               = "error_model_tier_index_bound"                                     // Invalid item iter_index.
	ErrErrorModelTierIndexLevelMismatchVarLevel               = "error_model_tier_index_level_mismatch_var_level"                  // Item tier_index level mismatch tier variation list level.
	ErrErrorModelUpdateNameModelInPromotion                   = "error_model_update_name_model_in_promotion"                       // tier_variation name and oprion cannot be editted when item/model is under promotion.
	ErrErrorModelUpdateStatusModelInPromotion                 = "error_model_update_status_model_in_promotion"                     // Cannot change model's status when item/model is promotion.
	ErrErrorModelUpdateStockModelInPromotion                  = "error_model_update_stock_model_in_promotion"                      // Model stock cannot be editted when item/model is promotion.
	ErrErrorNameLengthLimit                                   = "error_name_length_limit"                                          // Exceeded item_name length limitation.
	ErrErrorNeedOtp                                           = "error_need_otp"                                                   // Core server error.
	ErrErrorNeedVoiceOtp                                      = "error_need_voice_otp"                                             // Core server error.
	ErrErrorNegotiationStatus                                 = "error_negotiation_status"                                         // The negotiation status cannot support this action
	ErrErrorNetwork                                           = "error_network"                                                    // Inner http call failed | This API is only available for TH, VN, PH region | Connection issue. Please try again later. | Internal api called failed. Please try again later. | Core server error. | Inner error, please try later. [4].
	ErrErrorNilItemInReq                                      = "error_nil_item_in_req"                                            // Interal error, please contact openapi team.
	ErrErrorNilNameNewItem                                    = "error_nil_name_new_item"                                          // Item_name cannot be empty.
	ErrErrorNilNameNewItemModel                               = "error_nil_name_new_item_model"                                    // Item tier_variation name cannot be empty.
	ErrErrorNilShopidOrItemid                                 = "error_nil_shopid_or_itemid"                                       // Query information failed.
	ErrErrorNoBuyerOffer                                      = "error_no_buyer_offer"                                             // cannot accept refund offer because there is no counter proposal from buyer | cannot accept refund offer because there is no proposal from buyer
	ErrErrorNoDtsRangeNotAllowed                              = "error_no_dts_range_not_allowed"                                   // {{.error_info}}
	ErrErrorNotExists                                         = "error_not_exists"                                                 // The user does not exist. | Core server error.
	ErrErrorNotFound                                          = "error_not_found"                                                  // Wrong parameters, detail: {msg}. | Fail to get order fulfillment group, please try again later.	 | The information you queried is not found.
	ErrErrorNotLogin                                          = "error_not_login"                                                  // Core server error.
	ErrErrorOngoingDispute                                    = "error_ongoing_dispute"                                            // cannot offer refund to buyer when return has ongoing dispute
	ErrErrorOrder                                             = "error_order"                                                      // Something wrong. Please try again later. | System error, please try again later. | Invoice is not available for this order. Please try again later.
	ErrErrorOther                                             = "error_other"                                                      // Rescheduling dependency error
	ErrErrorOutStockm                                         = "error_out_stockm"                                                 // Core server error.
	ErrErrorParam                                             = "error_param"                                                      // There is no access_token in query. | Invalid partner_id. | There is no partner_id in query. | There is no sign in query. | no timestamp | Invalid timestamp | parameter invalid | The information you queried is not found. | Wrong parameters, detail: {msg}. | request not from gateway | Invalid category. | status is required. | get items offset over limit, please use the next field | {{.error_info}} | Can not update item with stock less than reserved stock | Can not update item with different stock structure. Can only update seller stock with location id when existing seller stock have location id. Can only update seller stock without location id when existing seller stock without location id. | Can not update item with stock less than reserve stock | [Gateway] illegal request. | Invalid logistic info, {{.error_info}} | Invalid attribute. {{.error_info}} | Invalid wholesale setting. | Parameter is not match the constraints, {{.error_info}}. | Invalid Weight. | Image not exist. | dimension is required | invalid additional information | all BR tax field should be empty or be filled at same time | Attribute format is invalid. NCC field only allows Eng Alphanumeric input | NCC filed only allows character length less than 50. | Invalid parameter for product. | invalid wholesale setting | The level of tier-variation not change. | The level of tier-variation over 2. | Model tier_index error. | Count of tier_variation options should be under 50. Count of 2-level variations combinations should be under 50. | Canot change the level of tier-variation. | Item without tier_variation. Please use init_tier_variation api to upgrade the structure. | Model_id and item_id not match. | Repeat model_id. | Wrong model_id. | Repeat item_id. | request must set one of [item_name,attribute_status] | brand image is required | Invalid request item_id | The item_id_list is required | Request items over limit, should be less than 50 | item_status does not match latest violation | update_time_from is required | update_time_to is required | update_time_to is later than now, only display item updated from update_time_from to now | category is prohibited | Update_time_to should be later than update_time_from. | Invalid dimension. | You have 3PF shops, please upload stock with location id | Partner_id is not found. | Invalid parameter. | The shop id you provided is not found. Please check your shop id. | The merchant id you provided is not found. Please check your merchant id. | The size of package list is over limit. | Order SN is a required parameter. Please check your request data. | The Order SN is duplicated. | OFG list type is not TO_PROCESS | Shipping parameters can only be obtained when package is ready to be shipped | Package cannot be rescheduled | Package is not under the specified fulfilment channel | The pickup_time_id received is invalid and not expected. Please check and reselect one from the time_slot_list. | Parameter sender_real_name is invalid.			 | Shop id: {shop_id} didn't set any default 7-11 return store, please go back to Shopee Seller Center to complete the setting.			 | Shop id: {shop_id} didn't set {return_store_name} return store, please go back to Shopee Seller Center to complete the setting. | Shop id: {shop_id} didn't set any default Shopee Express return store, please go back to Shopee Seller Center to complete the setting. | Shop id: {shop_id},return store:{return_store_name} not found. | Invalid package number, please get package number from field `package_number` instead of `forder_id`. | Please do not request with tracking_no for this channel which does not require tracking_no. | Slug not allow empty. | Slug not required. | Not allowed to self design AWB | This order logistics is not allowed to get. | Cannot set inbound pickup address for this address. | Inbound pickup address cannot be the same as Shopee warehouse address. | Inbound pickup address must be in São Paulo state. | Can't delete default address/pickup address/return address/inbound pickup address. | Shipping options of Direct Shop will follow Local Shop. | All modifications to the operating hours were unsuccessful. Please check the result_list and try again. | Can not get airway bill for warehouse booking orders.	 | Can not arrange shipment for warehouse booking orders	 | File size exceeds max allowable size. | File type is not .kml type | Dropoff FM tracking code must be between 6 to 40 characters			 | FM code cannot be empty. | Format wrong. FM tracking number only contains character and number.			 | Pickup FM tracking code must start with "CNF" or "KRF" + 18 numbers			 | Self Deliver FM tracking code must start with \"CNF\" or \"KRF\" + 18 numbers | Invalid promotion ID. The promotion may have been recently added and not mapped to the system yet. Please raise a ticket on Open Platform for further verification. | Missing order_sn. Please include a valid order_sn to get escrow details. | The request body is not a valid json format			 | Invalid or missing page_size parameter , please enter a valid value that is with 1-100, indicating the number of records within a page. | This API is no longer available in this region. If this is unexpected, please reach out to Customer Support for assistance. | This payment API is only applicable for cross boarder shop | Invalid page_size parameter. Please provide a valid value between 1 and 100. If left empty, the default value of 40 will be applied. | The selected payout start time is not supported. Please choose a start time no earlier than May 1, 2022.			 | shop_id is required			 | release_time_from and/or release_time_to parameter is missing or invalid. Please provide both a valid start date, and end date to query. | release_time_from and/or release_time_to parameter is invalid. Please provide both a valid start date, and end date to query where start date cannot be later than the end date. | shop_id parameter is missing or invalid. Please include a valid shop_id in the request to retrieve escrow records. | billing_transaction_info_type is required, format should be correct and input within 1-2 | Invalid or missing cursor parameter. Please enter an empty string i.e. "" for the first request or use the next_cursor value returned from the previous response for subsequent requests. | cursor is required and format should be correct. Please check | page_size is required, format should be correct and input within 1-100 | invalid payout_time format, please check | invalid time range, please ensure payout_time_to is greater than payout_time_from but within 15 days | payout_time_to is required | billing_transaction_info_type provided is invalid, please use one of the support billing item type values: 1 - To Release, 2 - Released. | payout ids{id1,id2,...} not found, please recheck if the id is correct or not | payout ids can not be empty when querying Released info | payout ids length exceed limit | order_sn_list is required, format should be string[] | Invalid Time Range: Only records from the past 5 years are supported for the TW region. | Shop is not a mart shop. Please use Mart Shop ID. | Invalid time range, Weekly reports must span from Day 1 (Monday) to Day 7 (Sunday), and monthly reports must cover from 1st to last day of the month. | Wrong income statement type. Please provide the correct parameter. | For local sellers, please input weekly or monthly.  | Wrong document ID provided. Please provide the correct ID in integer format. | release_time_from and/or release_time_to parameter is missing or invalid. Please provide both a valid start date, and end date to query where start date cannot be later than the end date. | Please enter a valid time range input, where the start date should not be later than the end date. | The income_status provided is invalid or not applicable. Please use one of the supported values: 0 - To Release, 1 - Released, 2 - Pending. Local shops do not have the “To Release” status. | release_time_from and/or release_time_to parameter is missing or invalid. Please provide both a valid start date and end date to query, where the start date cannot be later than the end date, and the date range must not exceed 14 days. | Access blocked for Mart shops, please retrieve income details using Outlet Shop ID instead. | The maximum number of the toggle-on automatic category cannot exceed 200. | ShopCategory name length cannot exceed 40 characters. | Sort_weight should be between [0, 2147483546]. | ShopCategory name is duplicated. | ShopCategory name should not be empty. | Syntax error. Please check the format of your Request Parameter. | This shop category doesn't exist. | Shop is not found. | You've reached the maximum number 1500 of categories, please delete some categories to create new			 | Automatic & shopee category cannot be deleted. | The maximum number of the toggle-on customized category cannot exceed 100. | Shopee 24 category cannot edit the name and sort_weight. | Nothing changes on the shop category update operation. | The maximum number of the toggle-on Shopee category cannot exceed 50. | ShopCategory's status can only be: NORMAL, INACTIVE, DELETED. | The total item number has exceed its limit number : 5000. | Automatic & shopee category cannot add items. | Page number(page_no)  should be  [1, 2147483446/ page_size]. | Automatic & shopee category cannot delete items | At most 100 items can be deleted per operation. | Page size should be [0, 1000]. | Core server error. | The period between create_time_from and created_time_of must not more than 15 days. | The update_time_from must be after create_time_from. | Inner error, please try later. [1]. | Return SN or ID is invalid. | All images download fail | This proposed solution is not one of the available solutions for this return. | The proposed adjusted refund amount can not be lower than min refund amount | Please fill a value for proposed adjusted refund amount | The proposed adjusted refund amount can not exceed max refund amount | The proposed adjusted refund amount must be positive | The proposed adjusted refund amount only accepts {digit} digit(s) after decimal point | The proposed solution cannot have adjusted refund amount | Wrong parameters. | this callback_url is invalid, please check your callback_url. | Your app's callback_url is empty, please input a callback_url before you turn on push config. | The call back url can not be Shopee intranet address, please change another new call back url | Shopee have sent a test push to this call back url, but we didn't get any response in 3 seconds with 2xx code, please check your service.  | Shopee have sent a test push to this call back url, the response code we get from this callback_url is not 2xx code, please check your service.  | The image may contain sensitive content, please use another one | No image found in the request | The image size cannot exceed {{image_size_limit}} bytes | Invalid image format, only support jpg/jpeg/png/gif) | Parse data error: {{err_msg}} | Invalid cover_image_url | Invalid title | Invalid session_id | Invalid domain_id | Invalid item_list, item_id and shop_id must greater than 0 | Invalid page_size, page_size must larger than 0 and cannot exceed {{limit}} | Invalid item_id | Invalid shop_id | Invalid item_set_ids | Cannot send empty comment | Invalid {{uid_param}} | The user(user_id:{{uid}}) is not exist
	ErrErrorParamCategoryNotSupportPreOrder                   = "error_param_category_not_support_pre_order"                       // Category does not support pre-order.
	ErrErrorParamDtsExceedsMaxLimit                           = "error_param_dts_exceeds_max_limit"                                // Days_to_ship exceeds max limit
	ErrErrorParamFormat                                       = "error_param_format"                                               // Core server error.
	ErrErrorParamItemStatus                                   = "error_param_item_status"                                          // Item status error.
	ErrErrorParamNameItem                                     = "error_param_name_item"                                            // There's no 'item_list' parameter in the query, please check.
	ErrErrorParamShopIdNotFound                               = "error_param_shop_id_not_found"                                    // Shop_id is not found.
	ErrErrorParamValidate                                     = "error_param_validate"                                             // This is not a valid GTIN. Please, inform a valid number. | Wholesale cannot be used in this category and attributes. | Your current merchant is a special merchant.Please ensure all your shop are under same region
	ErrErrorPasswordChange                                    = "error_password_change"                                            // Core server error.
	ErrErrorPerm                                              = "error_perm"                                                       // Core server error. | invalid dispute_reason. | there is invalid params. | invalid email : {email}.
	ErrErrorPermNonAdmin                                      = "error_perm_non_admin"                                             // Don't have permission.
	ErrErrorPermission                                        = "error_permission"                                                 // Sorry you don't have the permission, detail: {msg}. | Global product related shop is not upgraded. | Sorry you don't have the permission, detail: This shop isn't eligible to set up Regular Operating Hours. | Sorry you don't have the permission, detail: This shop isn't eligible to set up Instant Operating Hours. | Sorry you don't have the permission, detail: This shop isn't eligible to set up Special Operating Hours. | Special Operating Hours not found for this shop. | You don't have permission for the operation, detail: {msg}. | Please link shop and partner on seller center.
	ErrErrorPriceAgpTooLarge                                  = "error_price_agp_too_large"                                        // the max price is 7 lager the min price
	ErrErrorPriceExceedMaxLimitt                              = "error_price_exceed_max_limitt"                                    // Original_price is bigger than max price limit.
	ErrErrorPriceExceedMinLimitt                              = "error_price_exceed_min_limitt"                                    // Original_price is less than min price limit.
	ErrErrorPriceLimitNoPass                                  = "error_price_limit_no_pass"                                        // {{.error_info}}
	ErrErrorPriceOutOfRange                                   = "error_price_out_of_range"                                         // Price is out of range.
	ErrErrorPriceShouldBeSameForWholesales                    = "error_price_should_be_same_for_wholesales"                        // All model price should be the same when the item has been set wholesales.
	ErrErrorPromotionCantnotUpdateStock                       = "error_promotion_cantnot_update_stock"                             // Cannot change stock when item is under promotion.
	ErrErrorQueryConditionListLimit                           = "error_query_condition_list_limit"                                 // Interal error, please contact openapi team.
	ErrErrorQueryLimitMax                                     = "error_query_limit_max"                                            // Limit is above the maximum allowed for this query API.
	ErrErrorQueryOverItemidSize                               = "error_query_over_itemid_size"                                     // Interal error, please contact openapi team.
	ErrErrorQueryQueryLimitTooLarge                           = "error_query_query_limit_too_large"                                // Interal error, please contact openapi team.
	ErrErrorReachShopItemLimit                                = "error_reach_shop_item_limit"                                      // Item published item count reaches limit.
	ErrErrorRelatedProductInPromotion                         = "error_related_product_in_promotion"                               // Asku has upcoming or ongoing promotion, can't update global product priceｼ pls update price  in shop sku
	ErrErrorRepeatedMtsku                                     = "error_repeated_mtsku"                                             // A similar product has already been uploaded
	ErrErrorReturnRequestType                                 = "error_return_request_type"                                        // Action cannot be performed by shop because the Return Refund Request Type = {return_type}
	ErrErrorReturnStatus                                      = "error_return_status"                                              // Cannot raise dispute because return SN is not in REQUESTED, RETURN_PROCESSING, or RETURN_ACCEPTED | The return status cannot support this action
	ErrErrorSellerUnderPenalty                                = "error_seller_under_penalty"                                       // The shop is under penalty.
	ErrErrorServer                                            = "error_server"                                                     // Something wrong. Please try later.	 | Interal error, please contact openapi team. | The current item belong to the full FBS shop, so normal stock must be equal to 0 | System error. Please try again later. | Unknown error. | Unknown internal error. | Got an error when trying to generate image, please try again | Error Details : {msg} | Internal server error			 | The inner server err when send http request. | error server | The API is not supported for current region | Too many requests, please try again later | Server internal error: {{err_msg}}
	ErrErrorService                                           = "error_service"                                                    // Unknown error, please contact developer. | The shop category id is not found.
	ErrErrorSetNormalUnlistedItem                             = "error_set_normal_unlisted_item"                                   // Cannot change unlisted item status to normal directly, need to publishdelisted item first.
	ErrErrorShop                                              = "error_shop"                                                       // shopid is invalid | System error, please try again later. | Something wrong. Please try again later. | Shop is not eligible to upload file as instant channel is not enabled
	ErrErrorShopNotFound                                      = "error_shop_not_found"                                             // Shop is not found.
	ErrErrorSign                                              = "error_sign"                                                       // Wrong sign.
	ErrErrorSlashPriceItemDeleteLock                          = "error_slash_price_item_delete_lock"                               // In slash sale, item cannot be deleted.
	ErrErrorSlashPriceLoad                                    = "error_slash_price_load"                                           // Failed to load slash price.
	ErrErrorSlashPriceModelsDiff                              = "error_slash_price_models_diff"                                    // In slash sale, the model price should be the same.
	ErrErrorSlashPriceModelsToNomodelLock                     = "error_slash_price_models_to_nomodel_lock"                         // In slash sale, can not add model to a non-model item.
	ErrErrorSlashPriceNomodelToModelsLock                     = "error_slash_price_nomodel_to_models_lock"                         // In slash sale, can't add model.
	ErrErrorSlashPriceNotLowest                               = "error_slash_price_not_lowest"                                     // In slash sale, price should not be lower or same as slash price.
	ErrErrorStatus                                            = "error_status"                                                     // The order has been cancelled. | Buyer hasn’t completed pre-authorization yet. | Status verification failed, shipping documents of this package cannot be printed now.
	ErrErrorStockBiggerThenLimit                              = "error_stock_bigger_then_limit"                                    // Normal_stock/ seller_stock is bigger than max limit.
	ErrErrorStockLessThenMinLimit                             = "error_stock_less_then_min_limit"                                  // Normal_stock/ seller_stock is less than min limit.
	ErrErrorStockLimitNoPass                                  = "error_stock_limit_no_pass"                                        // {{.error_info}}
	ErrErrorSystemBusy                                        = "error_system_busy"                                                // Our system is taking some time to respond, please try later.
	ErrErrorTierImgNotAllower                                 = "error_tier_img_not_allower"                                       // Not allowed to update tier-variation image.
	ErrErrorTierImgOldApp                                     = "error_tier_img_old_app"                                           // Interal error, please contact openapi team.
	ErrErrorTierImgPartial                                    = "error_tier_img_partial"                                           // Interal error, please contact openapi team.
	ErrErrorTierIndex                                         = "error_tier_index"                                                 // The tier_index is not exist
	ErrErrorTierOptTooMany                                    = "error_tier_opt_too_many"                                          // Count of tier_variation option is larger than 20.
	ErrErrorTierOptValTooLong                                 = "error_tier_opt_val_too_long"                                      // {{.error_info}}
	ErrErrorTierVarIs                                         = "error_tier_var_is_"                                               // Item without tier-variation. Please use init_tier_variation api to upgrade the structure.
	ErrErrorTierVarLevelNotSame                               = "error_tier_var_level_not_same"                                    // Item tier-variation level not same
	ErrErrorTierVarNameTooLong                                = "error_tier_var_name_too_long"                                     // {{.error_info}}
	ErrErrorTierVarTooMany                                    = "error_tier_var_too_many"                                          // Count of tier_variation is larger than 50.
	ErrErrorTitleCharacterForbidden                           = "error_title_character_forbidden"                                  // Item_name contains forbidden characters.
	ErrErrorTitleExceedsMaxLength                             = "error_title_exceeds_max_length"                                   // The length of item_name is bigger than max limit.
	ErrErrorTitleLenNoPass                                    = "error_title_len_no_pass"                                          // {{.error_info}}
	ErrErrorUnknown                                           = "error_unknown"                                                    // {{.error_info}} | Core server error.
	ErrErrorUnlistInPromotion                                 = "error_unlist_in_promotion"                                        // Cannot unlist item when item is under promotion.
	ErrErrorUnlistItemAllFailed                               = "error_unlist_item_all_failed"                                     // Unlist item all failed.
	ErrErrorUnlistItemFail                                    = "error_unlist_item_fail"                                           // Please upload your products to UNLIST status. Products will be published automatically by Shopee at the official launch date.
	ErrErrorUnlistItemFailed                                  = "error_unlist_item_failed"                                         // Unlist item failed. please try later. | Unlist item failed. plaese try later.
	ErrErrorUpdatMtskuFail                                    = "error_updat_mtsku_fail"                                           // {{.error_info}}
	ErrErrorUpdateItemPrice                                   = "error_update_item_price"                                          // Update item price fail, operation not allowed | Update item price fail, reach item price edit limit for each day | Update item price fail, item price should less than the price | Update item price fail, mtsku not init
	ErrErrorUpdatePriceFail                                   = "error_update_price_fail"                                          // Update price failed, please try later.
	ErrErrorUpdateTimeRange                                   = "error_update_time_range"                                          // Update_time_to should be later than update_time_from.
	ErrErrorUserRefreshToken                                  = "error_user_refresh_token"                                         // Your refresh token is error, please check refresh token or user_id.
	ErrErrorValidation                                        = "error_validation"                                                 // Action cannot be performed by shop because the validation_type of Return/Refund request = warehouse_validation
	ErrErrorValueIdMustEqualZero                              = "error_value_id_must_equal_zero"                                   // Value_id must equal 0.
	ErrErrorValueNameRequired                                 = "error_value_name_required"                                        // Value_name is required.
	ErrErrorVersion                                           = "error_version"                                                    // Core server error.
	ErrErrorVideoInfoNotFound                                 = "error_video_info_not_found"                                       // Video_info not found.
	ErrErrorWholeSaleMinCountIncorrect                        = "error_whole_sale_min_count_incorrect"                             // Interal error, please contact openapi team.
	ErrErrorWholeSalePriceSettingIncorrect                    = "error_whole_sale_price_setting_incorrect"                         // Wholesale price can't more than original price.
	ErrErrorWholesalePriceLessThanRatioLimit                  = "error_wholesale_price_less_than_ratio_limit"                      // Wholesale price is less than ratio limit.
	ErrErrorWholesaleRangeNotAllowed                          = "error_wholesale_range_not_allowed"                                // {{.error_info}}
	ErrErrorWmsShopBlockUpateStock                            = "error_wms_shop_block_upate_stock"                                 // Warehouse shop can't update stock.
	ErrErrorWrongAttrsnapshot                                 = "error_wrong_attrsnapshot"                                         // Invalid attribute.
	ErrErrorWrongModelid                                      = "error_wrong_modelid"                                              // The model_id is wrong or not related to this item, model_id:{{.model_id}}
	ErrExceedMaxLimit                                         = "exceed_max_limit"                                                 // Your request exceeds the max limit of 50 orders.
	ErrFbsErrParam                                            = "fbs_err_param"                                                    // %v is invalid
	ErrFbsErrRecordNotFound                                   = "fbs_err_record_not_found"                                         // %v is not found
	ErrFbsErrRegionNotBr                                      = "fbs_err_region_not_br"                                            // Region is not BR
	ErrFbsErrShopSkuNotRelated                                = "fbs_err_shop_sku_not_related"                                     // Shop and SKU is not related
	ErrFbsErrSystem                                           = "fbs_err_system"                                                   // System error, please try again later
	ErrFirstmileAreaNotSupport                                = "firstmile.area_not_support"                                       // Area is invalid.Now we only support area in {area}.
	ErrFirstmileAuth                                          = "firstmile.auth"                                                   // The shop has no permission for first mile support.
	ErrFirstmileBatchApiAllFailed                             = "firstmile.batch_api_all_failed"                                   // Failed, please check result_list for more details.
	ErrFirstmileBusinessError                                 = "firstmile.business_error"                                         // There is common business error
	ErrFirstmileDateError                                     = "firstmile.date_error"                                             // Parameter declare_date is wrong, should be later than or equal to today.
	ErrFirstmileDateFormatError                               = "firstmile.date_format_error"                                      // Format for param date is wrong, should be like '2000-01-01'.
	ErrFirstmileDateRangeError                                = "firstmile.date_range_error"                                       // The "from_date" time must be smaller than the "to_date" time.
	ErrFirstmileError                                         = "firstmile.error"                                                  // One FM code cannot be predeclared with different shipping methods | The SLS TN does not belong to FM enabled TWS. | The SLS TN is tied to another FM code | The FM code does not exist | System error (TWS fail), please contact Shopee | The logistics status is not supported for binding | The SLS TN does not exist | The SLS TN is not a CB TN | The FM status is incorrect. | System error, please try again later/ contact Shopee | You used a Delivered FM code, please use a new FM code. | Your predeclared order creation time is later than pick up time, please use a new FM code. | The shipping method and logistics id are not matched | System error, please contact Shopee | FM Code WHSID can only bind SLS TN with the same WHSID
	ErrFirstmileErrorNetwork                                  = "firstmile.error_network"                                          // Get waybill fail due to the network error.
	ErrFirstmileErrorParam                                    = "firstmile.error_param"                                            // The ShopID and Pickup Label are not matched.
	ErrFirstmileInvalidOrderStatus                            = "firstmile.invalid_order_status"                                   // The status of order(order_sn: {order_sn}) is invalid. Please check the order status.
	ErrFirstmileInvalidPackage                                = "firstmile.invalid_package"                                        // Package is not exist.
	ErrFirstmileInvalidResponseField                          = "firstmile.invalid_response_field"                                 // Invalid response field({field}).
	ErrFirstmileMaxNumberLimit                                = "firstmile.max_number_limit"                                       // The first mile code number you apply is more than {num}.
	ErrFirstmileOrderHasBeenSplit                             = "firstmile.order_has_been_split"                                   // The order(order_sn: {order_sn}) has been split
	ErrFirstmileOrderNotExist                                 = "firstmile.order_not_exist"                                        // Order is not exist.
	ErrFirstmilePackageHasBind                                = "firstmile.package_has_bind"                                       // Package (order_sn: {order_sn}, package_number: {package_number}) has been bind to first mile code.
	ErrFirstmilePackageHasNotBind                             = "firstmile.package_has_not_bind"                                   // The SLS TN has not been tied to any first mile code.
	ErrFirstmileParamDuplication                              = "firstmile.param_duplication"                                      // Duplicate order_sn {order_sn} or package_number: {package_number}.
	ErrFirstmileShipmentAuth                                  = "firstmile.shipment_auth"                                          // The shop has no permission for the first mile shipment method: {method}.
	ErrFirstmileShipmentPreDeclarePermission                  = "firstmile.shipment_pre_declare_permission"                        // The shop has no permission for first mile pre declare.
	ErrFirstmileSystemError                                   = "firstmile.system_error"                                           // System error. Please try again later.
	ErrFirstmileWrongShipmentMethod                           = "firstmile.wrong_shipment_method"                                  // Shipment method is invalid. Please check the shipment method.
	ErrFollowPrizeCampaignNone                                = "follow_prize.campaign_none"                                       // The promotion id is not existed
	ErrFollowPrizeCampaignNumMaxLimit                         = "follow_prize.campaign_num_max_limit"                              // Max number of follow Prize reached (1000 ongoing and upcoming follow Prize)
	ErrFollowPrizeCampaignOverlap                             = "follow_prize.campaign_overlap"                                    // Another Follow Prize voucher already exists during this time period, please set another period.
	ErrFollowPrizeDeleteType                                  = "follow_prize.delete_type"                                         // Only the upcoming promotions can be deleted
	ErrFollowPrizeEndTimeMaxLimit                             = "follow_prize.end_time_max_limit"                                  // End time cannot exceed 6 months after start time
	ErrFollowPrizeEndTimeMinLimit                             = "follow_prize.end_time_min_limit"                                  // The end time must be greater than the start time by at least 1 day
	ErrFollowPrizeEndType                                     = "follow_prize.end_type"                                            // Only the ongoing follow prize can be ended
	ErrFollowPrizeMinValueBasketPrice                         = "Follow_prize_min_value_basket_price"                              // Minimum basket price can not be smaller than 49 for coin cashback voucher in TH
	ErrFollowPrizeNameLengthLimit                             = "follow_prize.name_length_limit"                                   // please input up to 20 characters of follow prize name
	ErrFollowPrizePercentageRANGE                             = "follow_prize.percentage_RANGE"                                    // Please enter a value between 1 and 99
	ErrFollowPrizeQuotaOutRange                               = "follow_prize.quota_out_range"                                     // Please enter a value between 1 and 200000
	ErrFollowPrizeRewardTypeCoinMx                            = "follow_prize.reward_type_coin_mx"                                 // Can't create coins cashback in MX
	ErrFollowPrizeStartTimeMinLimit                           = "follow_prize.start_time_min_limit"                                // Please enter a start time later than the current time
	ErrFollowPrizeTimeFutureLimit                             = "follow_prize.time_future_limit"                                   // The start/end time should not exceed 2037-12-31 23:59:59
	ErrFollowPrizeUpdateEndTimeLater                          = "follow_prize.update_end_time_later"                               // Promotion end time can only be changed to an earlier timing.
	ErrFollowPrizeUpdateExpiredCampaign                       = "follow_prize.update_expired_campaign"                             // The expired follow prize can't be updated.
	ErrFollowPrizeUpdateMinSpendOngoing                       = "follow_prize.update_min_spend_ongoing"                            // Minimum Basket Price can't be updated which is ongoing
	ErrFollowPrizeUpdatePrizeNameOngoing                      = "follow_prize.update_prize_name_ongoing"                           // Cannot rename follow prize name which is ongoing
	ErrFollowPrizeUpdateQuantityReduce                        = "follow_prize.update_quantity_reduce"                              // Cannot reduce dispatch quantity.
	ErrFollowPrizeUpdateStartTimeEarlier                      = "follow_prize.update_start_time_earlier"                           // Promotion start time can only be changed to a later timing.
	ErrFollowPrizeUpdateStartTimeOngoing                      = "follow_prize.update_start_time_ongoing"                           // Cannot update start time which is ongoing
	ErrFormatUnsupport                                        = "format.unsupport"                                                 // The Image Format Not Support
	ErrGetItemErrorServer                                     = "get_item_error_server"                                            // Get item failed.
	ErrImageUploadFail                                        = "image.upload.fail"                                                // image upload fail
	ErrIncomeErrorServer                                      = "income_error_server"                                              // Get escrow income failed.
	ErrIncomeNotFound                                         = "income_not_found"                                                 // The escrow income not found.
	ErrInternalError                                          = "internal_error"                                                   // Something wrong. Please try again later.
	ErrInternalErrorServer                                    = "internal_error_server"                                            // The inner server err when warp http req.
	ErrInvalidParam                                           = "invalid_param"                                                    // Parameters in the request are invalid, {msg}
	ErrInvalidRequest                                         = "invalid_request"                                                  // Parse the request failed
	ErrItemPriceTooLow                                        = "item_price_too_low"                                               // The credit card installment item's price should be grater than or equal to 1000.
	ErrLackOfInvoiceData                                      = "lack_of_invoice_data"                                             // Pending invoice data, can not arrange shipment.
	ErrLogisticsADDRESSNOTSUPPORT                             = "logistics.ADDRESS_NOT_SUPPORT"                                    // Address is not supported.
	ErrLogisticsAddressMissLongLat                            = "logistics.address_miss_long_lat"                                  // Address miss long lat.
	ErrLogisticsAddressNotFound                               = "logistics.address_not_found"                                      // Address is not found. | The address is invalid. Please check the address.
	ErrLogisticsAddressNotSupported                           = "logistics.address_not_supported"                                  // Address is not supported current operation.
	ErrLogisticsBlockBuyerName                                = "logistics.block_buyer_name"                                       // Buyer name is prohibited.
	ErrLogisticsBusinessProcessError                          = "logistics.business_process_error"                                 // Business process failed.
	ErrLogisticsBusinessValidationError                       = "logistics.business_validation_error"                              // Business validation failed.
	ErrLogisticsBuyerAddressNotFound                          = "logistics.buyer_address_not_found"                                // Buyer address is not found.
	ErrLogisticsCanNotPrintCombineOrder                       = "logistics.can_not_print_combine_order"                            // This order is part of a combined parcel, please use Seller Center instead.
	ErrLogisticsCanNotPrintJitOrder                           = "logistics.can_not_print_jit_order"                                // This shipping channel only supports document printing in Shopee seller center
	ErrLogisticsCanNotPrintToLabel                            = "logistics.can_not_print_to_label"                                 // You do not have permission to call this API because you are either not enabled for channel 30029 or do not have any ongoing orders in this channel.
	ErrLogisticsCanNotUpdateCombineOrder                      = "logistics.can_not_update_combine_order"                           // Can not update combine order through Open API, please use Shopee Seller Center instead.
	ErrLogisticsCancelForderFailed                            = "logistics.cancel_forder_failed"                                   // Cancel forder failed.
	ErrLogisticsCancelNotAllowed                              = "logistics.cancel_not_allowed"                                     // Cancel is not allowed.
	ErrLogisticsCannotDeleteAddress                           = "logistics.cannot_delete_address"                                  // Sorry you can't delete the default address/pickup address/return address.
	ErrLogisticsCategoryProhibited                            = "logistics.category_prohibited"                                    // There are contraband goods in the order.
	ErrLogisticsDefaultWarehouseNotSet                        = "logistics.default_warehouse_not_set"                              // You have not set warehouse.
	ErrLogisticsDownloadLater                                 = "logistics.download_later"                                         // Processing, please download later. Detail: {msg}
	ErrLogisticsError                                         = "logistics_error"                                                  // Fail to update address because fail to update the impacted channels. | Channel is invalid. Please check the channel id. | Sorry that this masked channel can't be enabled. | Sorry that this channel can’t be disabled since it is force enabled. | Sorry that this channel can't be enabled since it doesn’t pass the validation (whitelist validation / exclusive validation / serviceable validation ). | Sorry that this masked channel can't be preferred. | Sorry that this channel can't be enabled(disabled)/preferred/ COD enabled(COD disabled).
	ErrLogisticsErrorAddress                                  = "logistics.error_address"                                          // Address is invalid. Please check the address.
	ErrLogisticsErrorBookingOrder                             = "logistics.error_booking_order"                                    // This order is matched order with booking, don't need arrange shipment through seller side.
	ErrLogisticsErrorBuyerAddressid                           = "logistics.error_buyer_addressid"                                  // Buyer addressid is invalid. Please check the buyer addressid.
	ErrLogisticsErrorCannotCancelLogisticTxn                  = "logistics.error_cannot_cancel_logistic_txn"                       // Cancel operation is not supported for current logistics txn.
	ErrLogisticsErrorChannelExist                             = "logistics.error_channel_exist"                                    // Fail to find the fulfillment channel of this order.
	ErrLogisticsErrorConnection                               = "logistics.error_connection"                                       // Connection to external system is failed.
	ErrLogisticsErrorConsignment                              = "logistics.error_consignment"                                      // Fail to get consignment number.
	ErrLogisticsErrorConsignmentNoAccepted                    = "logistics.error_consignment_no_accepted"                          // Consignment number is not allow to update.
	ErrLogisticsErrorCoreServer                               = "logistics.error_core_server"                                      // Fail to call core server.
	ErrLogisticsErrorCutoffTime                               = "logistics.error_cutoff_time"                                      // Pickup time is not invalid. Please check the pickup time.
	ErrLogisticsErrorData                                     = "logistics.error_data"                                             // Store status is unavailable,pending buyer update.
	ErrLogisticsErrorDecimal                                  = "logistics.error_decimal"                                          // Receive invalid decimal.
	ErrLogisticsErrorDuplicateConsignment                     = "logistics.error_duplicate_consignment"                            // Consignment number is duplicated.
	ErrLogisticsErrorDuplicatedConsignment                    = "logistics.error_duplicated_consignment"                           // Consignment number is duplicate.
	ErrLogisticsErrorExpired                                  = "logistics.error_expired"                                          // Order is expired.
	ErrLogisticsErrorFailedGetNearestWhsIdn                   = "logistics.error_failed_get_nearest_whs_idn"                       // Faile to get the nearest whsId.
	ErrLogisticsErrorFormat                                   = "logistics.error_format"                                           // Content format from external system is invalid. Please check the format.
	ErrLogisticsErrorInvalidChannelId                         = "logistics.error_invalid_channel_id"                               // Channel id is invalid. Please check the channel.
	ErrLogisticsErrorIp                                       = "logistics.error_ip"                                               // Ip address: {ip} is deny.
	ErrLogisticsErrorItemNotFound                             = "logistics.error_item_not_found"                                   // Can not find the item.
	ErrLogisticsErrorLimit                                    = "logistics.error_limit"                                            // The batch request reach limit 50. | Can not update order logistics in current status. | Parcel count should not exceed limit. | Order status not support update logistics. | You don’t ship order yet.
	ErrLogisticsErrorLoginfo                                  = "logistics.error_loginfo"                                          // Logistic info is invalid. Please check the logistic info.
	ErrLogisticsErrorLogisticTxnNotFound                      = "logistics.error_logistic_txn_not_found"                           // Logistics record is not found.
	ErrLogisticsErrorLogisticsShopRequirePickupAddress        = "logistics.error_logistics_shop_require_pickup_address"            // Pickup address is required.
	ErrLogisticsErrorMessageType                              = "logistics.error_message_type"                                     // Message type is invalid. Please check the message type.
	ErrLogisticsErrorMethod                                   = "logistics.error_method"                                           // Operation is not supported.
	ErrLogisticsErrorMissingCdt                               = "logistics.error_missing_cdt"                                      // Cdt settings is not found.
	ErrLogisticsErrorNetwork                                  = "logistics.error_network"                                          // System error. Please try again later.
	ErrLogisticsErrorNoPickup                                 = "logistics.error_no_pickup"                                        // Pickup is not supported.
	ErrLogisticsErrorNoPickupAddress                          = "logistics.error_no_pickup_address"                                // Pickup address is missing.
	ErrLogisticsErrorNotExist                                 = "logistics.error_not_exist"                                        // Forder is not exist. | Address does not exist.
	ErrLogisticsErrorOperationFailed                          = "logistics.error_operation_failed"                                 // Operation failed.
	ErrLogisticsErrorOrderContext                             = "logistics.error_order_context"                                    // Order context is invalid. Please check the order context.
	ErrLogisticsErrorOrderNotFound                            = "logistics.error_order_not_found"                                  // Order is not found.
	ErrLogisticsErrorOrderState                               = "logistics.error_order_state"                                      // Operation is not allow with current order status.
	ErrLogisticsErrorOther                                    = "logistics.error_other"                                            // System error, please try again later. | Pick up address not support COD. | Can not support deliver address. | Can not support pick up address.
	ErrLogisticsErrorParam                                    = "logistics.error_param"                                            // The order is being allocated, please wait until the allocate is completed. | Duplicate package number | Package is not under the specified warehouse | Package exceeds limit. | The address is invalid. Please check the address. | Please wait for buyer to reselect the delivery convenience store. | Logistic status is not supported. | Deliver address not support COD. | Delivery address do not support COD. | This address do not support delivery. | Delivery address is not exist. | Buyer address is not exist. | Delivery store is not exist. | Delivery address is not completed. | This order is not allowed to cancel. | Error, please check the channel. | Fast Escrow ID was not found ! | COD is currently not supported in the area. Please check Shopee Help Center for COD serviceable area. | System error, please try again later. | The forder_ids should belong to one order/orderId and forderIds cannot be empty at the same time/parse forderId error. | Items in the same bundle should be in the same parcel. | Seller address is invalid. Please check your address. | Order list length exceeds limit. | Order is not exist. | This order can not be splited. | Order is not splittable. | Parcels of this order can not be combined. | Order/Forder can not update. | Order has been shipped. | Order has been splitted. | The order status is not ready to ship. | The buyer address is invalid. Please check the buyer address. | Shop not found. | Only support local warehouse, this shop do not support local warehouse. | This shop do not support split parcels. | Size setting error. | Invalid logistics channel. Please check the logictics channel. | Thaipost do not support this country. | Can not update order logistics. | Warehouse address not supported. | The EMS tracking number must begin with “E” | The EMS tracking number must end with “TH” | The Registered Mail tracking number must begin with “E”, “R” or “O” | The Registered Mail tracking number must end with “TH” | Please fill in a valid tracking number (13 digits) | The tracking number has been used by other orders, please fill in a valid one. | Please fill in the tracking number (13 digits) | Package exceeds weight limit. Please adjust or contact Shopee CS for more info. | Pickup address do not support COD. | This address do not support pickup. | Buyer's delivery address is not serviceable. Please contact customer service. | Seller address is not completed. | Channel does not support dropoff. Please select the pickup option and try again. | Please input a unique tracking number. | Seller Info Error. Please check and input your invoice info. | Tracking Number Error. Please input a valid tracking number. | Tracking Number Format Error. Please input a valid tracking number. | Please input a valid seller's name. | Invalid pickup time. Please select another pickup time. | Pickup address is not serviceable. Please select another address. | Seller Info Error. Please check and input your tax number. | Seller address is not exist. Please try other address. | Sender address is not exist. | Seller Address Error. Please input the geo-location of your address. | Selected pickup timeslot is unavailable. Please reselect. | Channel does not support pickup. Please select the dropoff option and try again. | Pickup address is not serviceable. Please select another address.			 | You can not get order logistics of warehouse order. | Invalid logistics_channel_id. Please check the logistics_channel_id.
	ErrLogisticsErrorPharmacistnameRequired                   = "logistics.error_pharmacistname_required"                          // Pharmacist name is required. Please enter a valid name.
	ErrLogisticsErrorPhone                                    = "logistics.error_phone"                                            // Phone number is invalid. Please check the phone number.
	ErrLogisticsErrorPickupTime                               = "logistics.error_pickup_time"                                      // Pickup time is invalid. Please check the pick up time.
	ErrLogisticsErrorReturnTrackingNumNotFound                = "logistics.error_return_tracking_num_not_found"                    // Return tracking number is not found.
	ErrLogisticsErrorSellerNeverApplyTrackingNoBefore         = "logistics.error_seller_never_apply_tracking_no_before"            // History is not found.
	ErrLogisticsErrorSenderAddress                            = "logistics.error_sender_address"                                   // Sender address is invalid. Please check the sender address. | Pickup address not valid. Please add/tag a new pickup_address in Seller Center to continue shipping
	ErrLogisticsErrorSetAsfByWeight                           = "logistics.error_set_asf_by_weight"                                // Fail to set actual shipping fee.
	ErrLogisticsErrorSetSellerAddress                         = "logistics.error_set_seller_address"                               // Fail to set seller address.
	ErrLogisticsErrorState                                    = "logistics.error_state"                                            // Currently, operation is not supported with current logistics record status.
	ErrLogisticsErrorStatus                                   = "logistics.error_status"                                           // Operation is not supported with current logistics record status.
	ErrLogisticsErrorStatusLimit                              = "logistics.error_status_limit"                                     // The order status is invalid. Please check the order status.
	ErrLogisticsErrorStoreNotFound                            = "logistics.error_store_not_found"                                  // Store is not found.
	ErrLogisticsErrorThirdPartyServer                         = "logistics.error_third_party_server"                               // Failed to call external system.
	ErrLogisticsErrorTimeout                                  = "logistics.error_timeout"                                          // Timeout to call external system.
	ErrLogisticsErrorToken                                    = "logistics.error_token"                                            // Token is invalid. Please check the token.
	ErrLogisticsErrorTooManyInvokeFunction                    = "logistics.error_too_many_invoke_function"                         // Failed to get lock.
	ErrLogisticsErrorTracenoRequired                          = "logistics.error_traceno_required"                                 // Trace number is required.
	ErrLogisticsErrorUnknown                                  = "logistics.error_unknown"                                          // Unknown exception, exception details is: {details}.
	ErrLogisticsErrorUnsupported                              = "logistics.error_unsupported"                                      // Operation is not supported.
	ErrLogisticsErrorUnsupportedAddress                       = "logistics.error_unsupported_address"                              // Address is not supported.
	ErrLogisticsErrorUpdate                                   = "logistics.error_update"                                           // Fail to update.
	ErrLogisticsErrorUpdateLogisticUnsupported                = "logistics.error_update_logistic_unsupported"                      // Currently, logistics is not supported to update.
	ErrLogisticsErrorZipcodeNotFound                          = "logistics.error_zipcode_not_found"                                // Zipcode is not found.
	ErrLogisticsFeatureNotSupported                           = "logistics.feature_not_supported"                                  // Feature is not supported.
	ErrLogisticsHasItemInBundle                               = "logistics.has_item_in_bundle"                                     // channel is not allow to disable due to there is item in the bundle
	ErrLogisticsInnerError                                    = "logistics.inner_error"                                            // System error. Please try again later.
	ErrLogisticsInvalidAddressType                            = "logistics.invalid_address_type"                                   // Address_type ({address_type}) is invalid. Please check the address_type. | CN Address Only Allow Set Seller Return Flag
	ErrLogisticsInvalidAddressVersion                         = "logistics.invalid_address_version"                                // Some addresses are no longer supported per new regulation, please update to avoid wrong fulfilment
	ErrLogisticsInvalidError                                  = "logistics.invalid_error"                                          // The order is not exist.
	ErrLogisticsInvalidSize                                   = "logistics.invalid_size"                                           // Size is invalid. Please check the
	ErrLogisticsItemIsInBundle                                = "logistics.item_is_in_bundle"                                      // Channel is not allow to disable due to there is item in the bundle.
	ErrLogisticsLackOfInvoiceData                             = "logistics.lack_of_invoice_data"                                   // Pending invoice data, can not arrange shipment.
	ErrLogisticsLogisticOrderIsLockedOnCreating               = "logistics.logistic_order_is_locked_on_creating"                   // Fail to get the lock.
	ErrLogisticsMaxDimensionIsLimited                         = "logistics.max dimension is limited"                               // Max dimension or weight is limited.
	ErrLogisticsMinimumActivationRule                         = "logistics.minimum_activation_rule"                                // {parent_channel} cannot be turned off. Please enable at least 1 of the following channel(s): {seller channel name list}
	ErrLogisticsMissingCdt                                    = "logistics.missing_cdt"                                            // Cdt settings is missing.
	ErrLogisticsNoAvailableTimeSlot                           = "logistics.no_available_time_slot"                                 // No available time slot
	ErrLogisticsNoBuyerStore                                  = "logistics.no buyer_store"                                         // Buyer store is missing.
	ErrLogisticsNoProportion                                  = "logistics.no_proportion"                                          // Can Not Found Proportion.
	ErrLogisticsNoSupportedDropoffBranch                      = "logistics.no_supported_dropoff_branch"                            // No supported drop-off branch
	ErrLogisticsNoSupportedPickupAddress                      = "logistics.no_supported_pickup_address"                            // No supported pickup address
	ErrLogisticsNoValidShippingParameters                     = "logistics.no_valid_shipping_parameters"                           // No valid shipping parameters. Please contact support
	ErrLogisticsNotAllowToUpdateChannel                       = "logistics.not_allow_to_update_channel"                            // You are not allowed to enable or disable channel.
	ErrLogisticsNotAvailableForSelfCollection                 = "logistics.not_available_for_self_collection"                      // Buyer self collection unavailable for this order
	ErrLogisticsNotFound                                      = "logistics.not_found"                                              // Not found.
	ErrLogisticsNotReadyForSelfCollection                     = "logistics.not_ready_for_self_collection"                          // Parcel is not ready for collection yet
	ErrLogisticsOrderFinalized                                = "logistics.order_finalized"                                        // Order is already in finalized status.
	ErrLogisticsOrderHasNoSellerAddress                       = "logistics.order_has_no_seller_address"                            // This order has no seller address.
	ErrLogisticsOrderNotExist                                 = "logistics.order_not_exist"                                        // The order_sn {order_sn} is not exist.
	ErrLogisticsOrderNotFound                                 = "logistics.order not found"                                        // Relative order is not found.
	ErrLogisticsOrderNotReadyToUpdate                         = "logistics.order_not_ready_to_update"                              // The order status is not ready to update.
	ErrLogisticsOrderPrescriptionUnchecked                    = "logistics.order_prescription_unchecked"                           // This order hasn't passed prescription check
	ErrLogisticsOrderStatusError                              = "logistics.order_status_error"                                     // Order status does not support awb printing.
	ErrLogisticsPackageAlreadyShipped                         = "logistics.package_already_shipped"                                // This package has already been arranged shipment and you can't call this API again.
	ErrLogisticsPackageCanNotPrint                            = "logistics.package_can_not_print"                                  // The package can not print now.
	ErrLogisticsPackageNotExist                               = "logistics.package_not_exist"                                      // The package is not exist.
	ErrLogisticsPackageNumberNotExist                         = "logistics.package_number_not_exist"                               // Please request with package_number for this split order.
	ErrLogisticsPackageNumberNotFound                         = "logistics.package_number_not_found"                               // The package_number {package_number} is not exist.
	ErrLogisticsPackagePrintFailed                            = "logistics.package_print_failed"                                   // Some package failed to print, please try again later. Detail: {msg}
	ErrLogisticsPackagesCanNotDownloadTogether                = "logistics.packages_can_not_download_together"                     // Packages can not download together. Detail: {msg}
	ErrLogisticsParcelAmountOverLimit                         = "logistics.parcel_amount_over_limit"                               // The order amount is over limit.
	ErrLogisticsPartnerShopFail                               = "logistics.partner_shop_fail"                                      // Get partner shop failed.
	ErrLogisticsPickupAddressUnsupported                      = "logistics.pickup_address_unsupported"                             // Pickup address is not supported.
	ErrLogisticsPickupWhitelistNotSupported                   = "logistics.pickup_whitelist_not_supported"                         // Permission is deny for pickup.
	ErrLogisticsPriceError                                    = "logistics.price_error"                                            // The price is invalid. Please check the price.
	ErrLogisticsQuantityError                                 = "logistics.quantity_error"                                         // The quantity is invalid. Please check the quantity.
	ErrLogisticsRcvStoreIdNotFound                            = "logistics.RcvStoreId not found"                                   // Receiver storeid is not found.
	ErrLogisticsRemoveOfgFromGroupFailed                      = "logistics.remove_ofg_from_group_failed"                           // Remove OFG from GROUP error, please try again.
	ErrLogisticsRequirePickupAddress                          = "logistics.require_pickup_address"                                 // Please go to set pickup address first.
	ErrLogisticsSdFileExpired                                 = "logistics.sd_file_expired"                                        // The shipping document file is expired, please get a new one and retry.
	ErrLogisticsSdJobPrintFailed                              = "logistics.sd_job_print_failed"                                    // Some package failed to print, please call create_shipping_document_job again. Detail: {{msg}}
	ErrLogisticsSelfCollectOutboundFailed                     = "logistics.self_collect_outbound_failed"                           // Prepare order for buyer collection
	ErrLogisticsSelfCollectionPinInvalid                      = "logistics.self_collection_pin_invalid"                            // Invalid collection PIN. Please verify and try again.
	ErrLogisticsSellerAddressNotSupportCod                    = "logistics.seller_address_not_support_cod"                         // Cod pickup is not supported with the seller address.
	ErrLogisticsShipOrderInvalidJobParam                      = "logistics.ship_order_invalid_job_param"                           // Parameter sender_real_name or tracking_number is required.
	ErrLogisticsShipOrderLongSenderRealName                   = "logistics.ship_order_long_sender_real_name"                       // Parameter sender_real_name can not be longer than 64.
	ErrLogisticsShipOrderLongTrackingNumber                   = "logistics.ship_order_long_tracking_number"                        // Parameter tracking_number can not be longer than 64.
	ErrLogisticsShipOrderNeedAddressPickupTime                = "logistics.ship_order_need_address_pickup_time"                    // Parameter address_id and pickup_time_id are required.
	ErrLogisticsShipOrderNeedBranchId                         = "logistics.ship_order_need_branch_id"                              // Parameter branch_id is required.
	ErrLogisticsShipOrderNeedPacakgeNumber                    = "logistics.ship_order_need_pacakge_number"                         // Please request with package_number for this split order.
	ErrLogisticsShipOrderNeedSenderRealName                   = "logistics.ship_order_need_sender_real_name"                       // Parameter sender_real_name can not be null.
	ErrLogisticsShipOrderNeedTrackingNumber                   = "logistics.ship_order_need_tracking_number"                        // Parameter tracking_number can not be null.
	ErrLogisticsShipOrderNotNeedPacakgeNumber                 = "logistics.ship_order_not_need_pacakge_number"                     // Please don’t request with package_number for this unsplit order.
	ErrLogisticsShipOrderNotReadyToShip                       = "logistics.ship_order_not_ready_to_ship"                           // The order is not ready to ship.
	ErrLogisticsShipOrderOnlySupportOneType                   = "logistics.ship_order_only_support_one_type"                       // Please select just one way to ship order: pickup or dropoff or non_integrated.
	ErrLogisticsShipOrderPffInit                              = "logistics.ship_order_pff_init"                                    // You can not ship warehouse order.
	ErrLogisticsShipOrderPickupTimeInvalid                    = "logistics.ship_order_pickup_time_invalid"                         // Invalid pickup time id. Please check the pick up time id.
	ErrLogisticsShipOrderUnsupportDropoff                     = "logistics.ship_order_unsupport_dropoff"                           // This order does not support ship with dropoff.
	ErrLogisticsShipOrderUnsupportNonIntegrated               = "logistics.ship_order_unsupport_non_integrated"                    // This order does not support ship with non_integrated.
	ErrLogisticsShipOrderUnsupportPickup                      = "logistics.ship_order_unsupport_pickup"                            // This order does not support ship with pickup.
	ErrLogisticsShippingDocumentShouldPrintFirst              = "logistics.shipping_document_should_print_first"                   // The package can not print now, please create shipping document first.Detail: {msg}
	ErrLogisticsShippingDocumentTypeInvalid                   = "logistics.shipping_document_type_invalid"                         // The shipping_document_type is not invalid. Please check the shipping_document_type.
	ErrLogisticsShopNotFound                                  = "logistics.shop_not_found"                                         // Shop is not found.
	ErrLogisticsShopNotSupportWms                             = "logistics.shop_not_support_wms"                                   // Wms is not supported for the shop.
	ErrLogisticsSlsCalculateFail                              = "logistics.sls calculate fail"                                     // Fail to calculate estimated shipping fee.
	ErrLogisticsTrackingNumberInvalid                         = "logistics.tracking_number_invalid"                                // The tracking number is invalid. Please check the tracking number.
	ErrLogisticsUnknownError                                  = "logistics.unknown_error"                                          // Unknown error, please contact Shopee to check. Detail: {msg}
	ErrLogisticsUpdateNotAllowed                              = "logistics.update_not_allowed"                                     // Currently, update is not allow.
	ErrLogisticsUpdateShippingOrderNoParamChange              = "logistics.update_shipping_order_no_param_change"                  // You can’t update shipping with all the same parameters.
	ErrLogisticsWarehouseStockInsufficient                    = "logistics.warehouse_stock_insufficient"                           // Failed to assign a warehouse with sufficient stock
	ErrMarketingErrorNoPremission                             = "marketing_error_no_premission"                                    // You don't have permission for the operation
	ErrNOTALLOWTOEDITADDRESSREGION                            = "NOT_ALLOW_TO_EDIT_ADDRESS_REGION"                                 // The param is invalid. not allow to edit address region.
	ErrNoProof                                                = "no.proof"                                                         // no proof
	ErrNumberError                                            = "number.error"                                                     // not support entry count
	ErrOrderCancelOrderInvalidStatus                          = "order.cancel_order_invalid_status"                                // Invalid order_status.The order status should be IN_CANCEL.
	ErrOrderDownloadInvoiceError                              = "order.download_invoice_error"                                     // Download invoice failed.
	ErrOrderErrorLimit                                        = "order.error_limit"                                                // Can not cancel this order.
	ErrOrderErrorParam                                        = "order. error_param"                                               // You can not cancel warehouse order.
	ErrOrderErrorPermission                                   = "order. error_permission"                                          // Please link shop and partner on seller center first.
	ErrOrderErrorServer                                       = "order_error_server"                                               // The escrow order extinfo decoded failed.
	ErrOrderNotFound                                          = "order_not_found"                                                  // Order SN provided is invalid or does not belong to you. Please enter a valid Order SN.
	ErrOrderOrderCannotSplit                                  = "order.order_cannot_split"                                         // You don’t have the permission to split order.
	ErrOrderOrderCannotSplitPff                               = "order.order_cannot_split_pff"                                     // Cannot split this warehouse order.
	ErrOrderOrderCannotUndoSplit                              = "order.order_cannot_undo_split"                                    // Cannot undo split this order.
	ErrOrderOrderHasBeenSplit                                 = "order.order_has_been_split"                                       // This order has already been split.
	ErrOrderOrderListInvalidTime                              = "order.order_list_invalid_time"                                    // Start time must be earlier than end time and diff in 15days.
	ErrOrderOrderSplitDuplicateItem                           = "order.order_split_duplicate_item"                                 // Cannot split order with duplicate items.
	ErrOrderOrderSplitFailed                                  = "order.order_split_failed"                                         // Split order failed, please try again. | Cannot split order paid by Credit Card Installment payment method or SPayLater payment method with instalment plan greater then 1x.
	ErrOrderOrderSplitGetPackageInfoFailed                    = "order.order_split_get_package_info_failed"                        // Split order succeed, but failed to get the new package info.
	ErrOrderOrderSplitItemNotExist                            = "order.order_split_item_not_exist"                                 // Cannot split order with invalid items.
	ErrOrderOrderSplitPackageMissingItem                      = "order.order_split_package_missing_item"                           // Cannot split order with missing items.
	ErrOrderOrderSplitWrongItemPosition                       = "order.order_split_wrong_item_position"                            // Cannot split order with items in one bundle-deal or add-on deal into different packages.
	ErrOrderOrderStatusInvalid                                = "order.order_status_invalid"                                       // Invalid order_status.The order status should be IN_CANCEL.
	ErrOrderOrderUnsplitFailed                                = "order.order_unsplit_failed"                                       // Undo split failed, please try again later.
	ErrOrderPrescriptionOrderStatusInvalid                    = "order.prescription_order_status_invalid"                          // Prescription order is already completed or canceled
	ErrOrderShopidInvalid                                     = "order_shopid_invalid"                                             // The escrow order's shopid is invalid.
	ErrOrderUploadInvoiceError                                = "order.upload_invoice_error"                                       // Upload invoice failed, please try again later. | File error.
	ErrProductCmtNotExist                                     = "product.cmt_not_exist"                                            // The comment is not exist. Please check the comment.
	ErrProductCommentLengthInvalid                            = "product.comment_length_invalid"                                   // Comment length should be between 1 and 500. Please update the comment.
	ErrProductDuplicateCommentId                              = "product.duplicate_comment_id"                                     // Duplicate comment id { cmm_id}.
	ErrProductDuplicateRequest                                = "product.duplicate_request"                                        // You has replied this comment already.
	ErrProductErrorBusi                                       = "product.error_busi"                                               // The GTIN code is mandatory, please check and upload again. | Please input the correct TS Mark (TD Mark) to upload your product, and refer to the SEH article - https://seller.shopee.tw/edu/article/19732 if you have any questions. | Medicine list ID is mandatory for products in Prescription/OTC category. | Please input the correct medicine list ID. | Upload failed, please upload a more standard size chart image.
	ErrProductErrorNotExist                                   = "product.error_not_exist"                                          // The Comment you replied does not exist.
	ErrProductErrorPermission                                 = "product.error_permission"                                         // Reply comment failed. Shop token invalid, please try to re-auth to partner.
	ErrProductGetCmtFailed                                    = "product.get_cmt_failed"                                           // Get comment failed. Please try again later.
	ErrProductGetTokenFailed                                  = "product.get_token_failed"                                         // Get token failed.
	ErrProductItemCmtCombineWrong                             = "product.item_cmt_combine_wrong"                                   // Item id of cmt is not the same as item id get from param.
	ErrProductItemNotExist                                    = "product.item_not_exist"                                           // Item (ID {item_id}) is not exist.
	ErrProofOverdue                                           = "proof.overdue"                                                    // proof is overdue
	ErrRegionNotSupport                                       = "region_not_support"                                               // Region has no not support credit card installment.
	ErrReturnInfoErrorServer                                  = "return_info_error_server"                                         // Get return or refund info failed.
	ErrReturnStatusIllegal                                    = "return.status.illegal"                                            // dms return status error
	ErrReturnsDisputeIsDryRun                                 = "returns.dispute_is_dry_run"                                       // Dispute validation passed, but will not submit as this is a dry run
	ErrReturnsErrorBlankTrackingNo                            = "returns.error_blank_tracking_no"                                  // Tracking Number field cannot be blank
	ErrReturnsErrorFilenumber                                 = "returns.error_filenumber"                                         // Number of image files exceeds three. Please only upload max three image files
	ErrReturnsErrorLogisticsStatusAudit                       = "returns.error_logistics_status_audit"                             // Failed to get return logistics status audit
	ErrReturnsErrorOrderFulfilmentGroup                       = "returns.error_order_fulfilment_group"                             // Failed to get order fulfilment group
	ErrReturnsErrorPostReturnLogisticsData                    = "returns.error_post_return_logistics_data"                         // Failed to get post return logistics data
	ErrReturnsErrorReverseLogistics                           = "returns.error_reverse_logistics"                                  // The Return SN does not have reverse logistics status
	ErrReturnsErrorReverseLogisticsTrackingInfo               = "returns.error_reverse_logistics_tracking_info"                    // Failed to get reverse logistics tracking info
	ErrReturnsErrorTrackingNoAlphanumeric                     = "returns.error_tracking_no_alphanumeric"                           // Tracking number should not contain special characters; for eg. !; @; $; %; &; *
	ErrReturnsErrorTrackingNoLength                           = "returns.error_tracking_no_length"                                 // Tracking number should be 5-30 characters
	ErrReturnsExceedShippingProof                             = "returns.exceed_shipping_proof"                                    // Shipping Proof have already been uploaded
	ErrReturnsInvalidCarrierId                                = "returns.invalid_carrier_id"                                       // Reverse Logistics carrier ID is invalid
	ErrReturnsInvalidImageId                                  = "returns.invalid_image_id"                                         // Provided image ID is invalid
	ErrReturnsInvalidTrackingNo                               = "returns.invalid_tracking_no"                                      // Tracking Number is invalid
	ErrReturnsMissingCarrierId                                = "returns.missing_carrier_id"                                       // Reverse Logistics carrier ID is a required parameter
	ErrReturnsMissingCarrierName                              = "returns.missing_carrier_name"                                     // Reverse Logistics carrier name is a required parameter
	ErrReturnsMissingImage                                    = "returns.missing_image"                                            // Shipping Proof image is a required parameter
	ErrReturnsMissingReturnSn                                 = "returns.missing_return_sn"                                        // Return SN is a required parameter
	ErrReturnsMissingReverseLogisticsCarrierList              = "returns.missing_reverse_logistics_carrier_list"                   // No reverse logistics carrier list available
	ErrReturnsMissingTrackingNo                               = "returns.missing_tracking_no"                                      // Tracking Number is a required parameter
	ErrReturnsNewSellerArrangeFlowDisabled                    = "returns.new_seller_arrange_flow_disabled"                         // New seller arrange flow is not enabled
	ErrReturnsReturnNotFound                                  = "returns.return_not_found"                                         // Return not found
	ErrReturnsnIllegal                                        = "returnsn.illegal"                                                 // Return status does not allow initiating a dispute case（From dispute side） | return sn is illegal
	ErrRraocRefundNotAllowed                                  = "rraoc_refund_not_allowed"                                         // Type of return does not allow seller to offer refund
	ErrSIPDISCOUNTERRORDISCOUNTRATEFORVN                      = "SIP_DISCOUNT_ERROR_DISCOUNT_RATE_FOR_VN"                          // You can not give more than 50% off discount in VN
	ErrSIPDISCOUNTERRORINVALIDREGION                          = "SIP_DISCOUNT_ERROR_INVALID_REGION"                                // Please input a valid region where you have opened a shop
	ErrSIPDISCOUNTERRORUPDATELOCKED                           = "SIP_DISCOUNT_ERROR_UPDATE_LOCKED"                                 // You can not edit the promotion within 15 minutes after the update
	ErrServerInternalError                                    = "server_internal_error"                                            // The server encountered an internal error, please retry or contact the server administrator
	ErrShopFlashSaleAlreadyExist                              = "shop_flash_sale_already_exist"                                    // shop flash sale already exist
	ErrShopFlashSaleExceedMaxItemCriteria                     = "shop_flash_sale_exceed_max_item_criteria"                         // exceed the max number of item limit per promotion
	ErrShopFlashSaleExceedMaxItemLimit                        = "shop_flash_sale_exceed_max_item_limit"                            // cannot enabled more than 50 items
	ErrShopFlashSaleInHolidayMode                             = "shop_flash_sale_in_holiday_mode"                                  // Shop in holiday mode
	ErrShopFlashSaleIsNotEnabledOrUpcoming                    = "shop_flash_sale_is_not_enabled_or_upcoming"                       // this shop flash sale is not enabled or upcoming
	ErrShopFlashSaleNotExist                                  = "shop_flash_sale_not_exist"                                        // shop flash sale not exist
	ErrShopFlashSaleNotMeetShopCriteria                       = "shop_flash_sale.not_meet_shop_criteria"                           // not meet shop criteria
	ErrShopFlashSaleParamError                                = "shop_flash_sale_param_error"                                      // The Wrong parameters, detail: {msg}.
	ErrShopFlashSaleShopInactive                              = "shop_flash_sale_shop_inactive"                                    // shop is inactive
	ErrShopManualOff                                          = "shop_manual_off"                                                  // Shop installment status is off.
	ErrShopNotInWhitelist                                     = "shop_not_in_whitelist"                                            // Shop is not in credit card installment whitelist.
	ErrShopidInvalid                                          = "shopid_invalid"                                                   // Shopid is invalid.
	ErrSmidNotVerified                                        = "smid_not_verified"                                                // Your SMID has not verified.
	ErrTenuresInvalid                                         = "tenures_invalid"                                                  // The credit card installment tenures is wrong.
	ErrTimeInvalid                                            = "time_invalid"                                                     // The param error: create_time_from bigger than create_time_to. | The param error: time is invalid
	ErrTimePeriodTooLarge                                     = "time_period_too_large"                                            // The param error: time period too large.
	ErrTopPickExceedMaxTopPickCount                           = "top_pick.exceed_max_top_pick_count"                               // The created top picks can not more than {num}.
	ErrTopPickTopPickDeleteStatusError                        = "top_pick.top_pick_delete_status_error"                            // The enabled top-picks can not be deleted.
	ErrTopPickTopPickItemIdDuplication                        = "top_pick.top_pick_item_id_duplication"                            // The top-picks has duplicative item.
	ErrTopPickTopPickItemIdInvalid                            = "top_pick.top_pick_item_id_invalid"                                // Item {itemid} is wrong, please check it.
	ErrTopPickTopPickItemIdNotExist                           = "top_pick.top_pick_item_id_not_exist"                              // Item {itemid} is not belong to shop, please check it.
	ErrTopPickTopPickNameDuplication                          = "top_pick.top_pick_name_duplication"                               // The Top-picks name is exist, please change to other name.
	ErrUnsupportRegionForRegisterBrand                        = "unsupport_region_for_register_brand"                              // The market this shop belongs to do not support brand registration.
	ErrUseridInvalid                                          = "userid_invalid"                                                   // Userid is invalid.
	ErrVoucherDeleteVoucherError                              = "voucher.delete_voucher_error"                                     // Can only delete upcoming voucher.
	ErrVoucherEndVoucherStatusError                           = "voucher.end_voucher_status_error"                                 // Can only end the ongoing voucher.
	ErrVoucherFixAmountLowQualtiyErrorr                       = "voucher.fix_amount_low_qualtiy_errorr"                            // Voucher cannot be saved as it is lower than {limit}%. Please adjust the percentage amount to be more than or equal to {limit}%.
	ErrVoucherOperateVoucherPermissionError                   = "voucher.operate_voucher_permission_error"                         // The Voucher created in Shopee admin/Shop Game/Follow Prize/Member ship/Voucher campaign can't be updated
	ErrVoucherOversizeList                                    = "voucher_oversize_list"                                            // You can only add up to 1000 products
	ErrVoucherPercentageLowQualityError                       = "voucher.percentage_low_quality_error"                             // Voucher cannot be saved as it is lower than {limit}.00%. Please adjust the percentage amount to be more than or equal to {limit}.00%.
	ErrVoucherUpdateEndVoucherError                           = "voucher.update_end_voucher_error"                                 // Cannot update status for expired voucher.
	ErrVoucherVoucherBtocItemBlockedError                     = "voucher.voucher_btoc_item_blocked_error"                          // Some items cannot be uploaded as they are b2c products in FBS shop
	ErrVoucherVoucherCodeDuplicateError                       = "voucher.voucher_code_duplicate_error"                             // Duplicated voucher code.
	ErrVoucherVoucherCodeError                                = "voucher.voucher_code_error"                                       // Voucher code doesn't meet the requirement. Only support 0-9 or A-Z or a-z. Up to 5 characters.
	ErrVoucherVoucherCoinError                                = "voucher.voucher_coin_error"                                       // Cannot create coin cashback voucher.
	ErrVoucherVoucherDisplayChannelError                      = "voucher.voucher_display_channel_error"                            // Voucher display channel is set wrongly.
	ErrVoucherVoucherEndTimeError                             = "voucher.voucher_end_time_error"                                   // Voucher valid period can not be longer than 3 months.
	ErrVoucherVoucherEndTimeReduceError                       = "voucher.voucher_end_time_reduce_error"                            // Only allow to shorten the voucher end_time. Cannot extend the end_time.
	ErrVoucherVoucherItemBlockedFromPromotionError            = "voucher.voucher_item_blocked_from_promotion_error"                // Some items cannot be uploaded as they are prohibited from promotions due to regulations
	ErrVoucherVoucherItemCountLimitError                      = "voucher.voucher_item_count_limit_error"                           // The number of items in item_id_list exceeds the limit.
	ErrVoucherVoucherItemDuplicationError                     = "voucher.voucher_item_duplication_error"                           // Duplicated items in item_id_list.
	ErrVoucherVoucherItemNotExist                             = "voucher.voucher_item_not_exist"                                   // Some item in item_id_list doesn't exist
	ErrVoucherVoucherItemStatusError                          = "voucher.voucher_item_status_error"                                // Not all the items in item_id_list are in normal status.
	ErrVoucherVoucherItemStockError                           = "voucher.voucher_item_stock_error"                                 // Some item in item_id_list has 0 stock.
	ErrVoucherVoucherMaxDiscountLowQuality                    = "voucher.voucher_max_discount_low_quality"                         // To attract buyers to use your voucher, please increase maximum discount price to > {percentage_input*min_spend*0.01} or reduce minimum basket price to < {max_discount*100/percentage_input} to ensure buyers can enjoy the percentage discount configured.
	ErrVoucherVoucherMinPriceError                            = "voucher.voucher_min_price_error"                                  // Min_basket_price must be higher than the fix amount if Min_basket_price is not 0.
	ErrVoucherVoucherOngoingParamError                        = "voucher.voucher_ongoing_param_error"                              // Only limited fields can be modified when the voucher is ongoing. Fileds allow modification: voucher_name,usage_quantity,end_time,display_channel_list, item_id_list
	ErrVoucherVoucherOverLimit                                = "voucher.voucher_over_limit"                                       // The total number of ongoing voucher and upcoming voucher cannot exceed 1000.
	ErrVoucherVoucherQuantityReduceError                      = "voucher.voucher_quantity_reduce_error"                            // Only allow to increase the voucher quantity.
	ErrVoucherVoucherStartTimeError                           = "voucher.voucher_start_time_error"                                 // Voucher start time is earlier than current time.
	ErrVoucherVoucherTimePeriodError                          = "voucher.voucher_time_period_error"                                // Voucher end time must be 1 hour later than the start time.
	ErrWarehouseErrorCanNotFindWarehouse                      = "warehouse.error_can_not_find_warehouse"                           // This error will show if there is no legal warehouse address for given shop id
	ErrWarehouseErrorNotInWhitelist                           = "warehouse.error_not_in_whitelist"                                 // This error will show if your shop has no permission to access multi-warehouse
	ErrWarehouseErrorRegionCanNotBlank                        = "warehouse.error_region_can_not_blank"                             // This error will show if region is missing.
	ErrWarehouseErrorRegionNotValid                           = "warehouse.error_region_not_valid"                                 // This error will show if region is not in ID, PH, VN, SG, TW, MX
	ErrWarehouseErrorShopIdCanNotBlank                        = "warehouse.error_shop_id_can_not_blank"                            // This error will show if shop id is missing.
	Err_10002                                                 = "10002"                                                            // error_shop_not_exists
	Err_10011                                                 = "10011"                                                            // err_service_internal
// END GENERATED ERRORS
)

// App represents basic app settings such as Api key, secret, scope, and redirect url.
// See oauth.go for OAuth related helper functions.
type App struct {
	PartnerID   int    `env:"SHOPEE_PARTNER_ID"`
	PartnerKey  string `env:"SHOPEE_PARTNER_KEY"`
	RedirectURL string `env:"SHOPEE_REDIRECT_URL"`
	APIURL      string `env:"SHOPEE_API_URL"`
}

type RateLimitInfo struct {
	RequestCount      int
	BucketSize        int
	RetryAfterSeconds float64
}

// Client manages communication with the Shopify API.
type Client[T any] struct {
	Client *http.Client
	log    LeveledLoggerInterface

	app App

	// Base URL for API requests.
	baseURL *url.URL

	// max number of retries, defaults to 0 for no retries see WithRetry option
	retries  int
	attempts int

	RateLimits RateLimitInfo

	ShopID      uint64
	MerchantID  uint64
	AccessToken string

	RefreshToken   string
	OnTokenRefresh func(res *RefreshAccessTokenResponse, meta T)
	Meta           T

	// Services used for communicating with the API
	// BEGIN GENERATED SERVICES
	Util          UtilService
	Auth          AuthService
	Product       ProductService
	GlobalProduct GlobalProductService
	MediaSpace    MediaSpaceService
	Media         MediaService
	Shop          ShopService
	Merchant      MerchantService
	Order         OrderService
	Logistics     LogisticsService
	FirstMile     FirstMileService
	Payment       PaymentService
	Discount      DiscountService
	BundleDeal    BundleDealService
	AddOnDeal     AddOnDealService
	Voucher       VoucherService
	ShopFlashSale ShopFlashSaleService
	FollowPrize   FollowPrizeService
	TopPicks      TopPicksService
	ShopCategory  ShopCategoryService
	Returns       ReturnsService
	AccountHealth AccountHealthService
	Ads           AdsService
	Public        PublicService
	Push          PushService
	SBS           SBSService
	FBS           FBSService
	Livestream    LivestreamService
	// END GENERATED SERVICES
}

// DefaultClient is a non-generic version of the Client
type DefaultClient = Client[any]

// NewClient returns a new Shopify API client with an already authenticated shopname and
// token. The shopName parameter is the shop's myshopify domain,
// e.g. "theshop.myshopify.com", or simply "theshop"
// a.NewClient(shopName, token, opts) is equivalent to NewClient(a, shopName, token, opts)
func NewClient[T any](app App, opts ...Option[T]) *Client[T] {
	baseURL, err := url.Parse(app.APIURL)
	if err != nil {
		panic(err)
	}

	c := &Client[T]{
		Client:  &http.Client{Timeout: time.Duration(defaultHttpTimeout) * time.Second},
		log:     &LeveledLogger{},
		app:     app,
		baseURL: baseURL,
	}

	// BEGIN GENERATED SERVICES INIT
	c.Util = &UtilServiceOp[T]{client: c}
	c.Auth = &AuthServiceOp[T]{client: c}
	c.Product = &ProductServiceOp[T]{client: c}
	c.GlobalProduct = &GlobalProductServiceOp[T]{client: c}
	c.MediaSpace = &MediaSpaceServiceOp[T]{client: c}
	c.Media = &MediaServiceOp[T]{client: c}
	c.Shop = &ShopServiceOp[T]{client: c}
	c.Merchant = &MerchantServiceOp[T]{client: c}
	c.Order = &OrderServiceOp[T]{client: c}
	c.Logistics = &LogisticsServiceOp[T]{client: c}
	c.FirstMile = &FirstMileServiceOp[T]{client: c}
	c.Payment = &PaymentServiceOp[T]{client: c}
	c.Discount = &DiscountServiceOp[T]{client: c}
	c.BundleDeal = &BundleDealServiceOp[T]{client: c}
	c.AddOnDeal = &AddOnDealServiceOp[T]{client: c}
	c.Voucher = &VoucherServiceOp[T]{client: c}
	c.ShopFlashSale = &ShopFlashSaleServiceOp[T]{client: c}
	c.FollowPrize = &FollowPrizeServiceOp[T]{client: c}
	c.TopPicks = &TopPicksServiceOp[T]{client: c}
	c.ShopCategory = &ShopCategoryServiceOp[T]{client: c}
	c.Returns = &ReturnsServiceOp[T]{client: c}
	c.AccountHealth = &AccountHealthServiceOp[T]{client: c}
	c.Ads = &AdsServiceOp[T]{client: c}
	c.Public = &PublicServiceOp[T]{client: c}
	c.Push = &PushServiceOp[T]{client: c}
	c.SBS = &SBSServiceOp[T]{client: c}
	c.FBS = &FBSServiceOp[T]{client: c}
	c.Livestream = &LivestreamServiceOp[T]{client: c}
	// END GENERATED SERVICES INIT

	// apply any options
	for _, opt := range opts {
		opt(c)
	}

	return c
}

// NewDefaultClient returns a new Shopify API client with any as meta type
func NewDefaultClient(app App, opts ...DefaultOption) *DefaultClient {
	return NewClient(app, opts...)
}

// A general response error that follows a similar layout to Shopify's response
// errors, i.e. either a single message or a list of messages.
type ResponseError struct {
	Status      int
	Message     string
	Errors      []string
	ShopeeError string
	RequestID   string
}

// GetStatus returns http  response status
func (e ResponseError) GetStatus() int {
	return e.Status
}

// GetMessage returns response error message
func (e ResponseError) GetMessage() string {
	return e.Message
}

// GetErrors returns response errors list
func (e ResponseError) GetErrors() []string {
	return e.Errors
}

// GetShopeeError returns the raw error code from Shopee
func (e ResponseError) GetShopeeError() string {
	return e.ShopeeError
}

// GetRequestID returns the request ID from Shopee
func (e ResponseError) GetRequestID() string {
	return e.RequestID
}

func (e ResponseError) Error() string {
	msg := e.Message
	if msg == "" && len(e.Errors) > 0 {
		sort.Strings(e.Errors)
		msg = strings.Join(e.Errors, ", ")
	}

	if msg == "" {
		msg = "Unknown Error"
	}

	if e.ShopeeError != "" {
		msg = fmt.Sprintf("shopee: %s [%s]", e.ShopeeError, msg)
	}

	if e.RequestID != "" {
		msg = fmt.Sprintf("%s (RequestID: %s)", msg, e.RequestID)
	}

	return msg
}

// IsShopeeError checks if the error is a Shopee ResponseError with the given error code.
func IsShopeeError(err error, shopeeErrCode string) bool {
	if re, ok := err.(ResponseError); ok {
		return re.ShopeeError == shopeeErrCode
	}
	if re, ok := err.(*ResponseError); ok {
		return re.ShopeeError == shopeeErrCode
	}
	return false
}

// ResponseDecodingError occurs when the response body from Shopify could
// not be parsed.
type ResponseDecodingError struct {
	Body    []byte
	Message string
	Status  int
}

func (e ResponseDecodingError) Error() string {
	return e.Message
}

// An error specific to a rate-limiting response. Embeds the ResponseError to
// allow consumers to handle it the same was a normal ResponseError.
type RateLimitError struct {
	ResponseError
	RetryAfter int
}

// Creates an API request. A relative URL can be provided in urlStr, which will
// be resolved to the BaseURL of the Client. Relative URLS should always be
// specified without a preceding slash. If specified, the value pointed to by
// body is JSON encoded and included as the request body.
func (c *Client[T]) NewRequest(method, relPath string, body, options, headers interface{}) (*http.Request, error) {
	rel, err := url.Parse(relPath)
	if err != nil {
		return nil, err
	}

	// Make the full url based on the relative path
	u := c.baseURL.ResolveReference(rel)

	// Add custom options
	if options != nil {
		optionsQuery, err := query.Values(options)
		if err != nil {
			return nil, err
		}

		// Shopee V2 GET requests expect lists as JSON array strings
		// e.g. item_id_list=[123,456] instead of item_id_list=123&item_id_list=456
		jsonListParams := []string{
			"category_id_list",
			"main_item_id",
			"direct_item_id",
			"shop_id_list",
			"enabled_channel_id_list",
		}
		for _, param := range jsonListParams {
			if values, ok := optionsQuery[param]; ok && len(values) > 1 {
				optionsQuery.Set(param, "["+strings.Join(values, ",")+"]")
			} else if ok && len(values) == 1 && strings.Contains(values[0], ",") {
				// Already joined or potentially a single value that needs brackets
				// but usually go-querystring with a slice of 1 element still needs brackets for Shopee
				if !strings.HasPrefix(values[0], "[") {
					optionsQuery.Set(param, "["+values[0]+"]")
				}
			} else if ok && len(values) == 1 {
				// Single value in a slice is encoded as "val", but Shopee might want "[val]"
				if !strings.HasPrefix(values[0], "[") {
					optionsQuery.Set(param, "["+values[0]+"]")
				}
			}
		}

		// item_id_list for get_item_base_info seems to want comma-separated WITHOUT brackets
		// because of error: strconv.ParseUint: parsing "[844121713"
		if values, ok := optionsQuery["item_id_list"]; ok && len(values) > 0 {
			optionsQuery.Set("item_id_list", strings.Join(values, ","))
		}

		for k, values := range u.Query() {
			for _, v := range values {
				optionsQuery.Add(k, v)
			}
		}
		u.RawQuery = optionsQuery.Encode()
	}

	// A bit of JSON ceremony
	var js []byte = nil
	if body != nil {
		js, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), bytes.NewBuffer(js))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", UserAgent)

	c.makeSignature(req)

	return req, nil
}

func (c *Client[T]) WithShop(sid uint64, tok string) *Client[T] {
	c.ShopID = sid
	c.AccessToken = tok
	return c
}

func (c *Client[T]) WithMerchant(mid uint64, tok string) *Client[T] {
	c.MerchantID = mid
	c.AccessToken = tok
	return c
}

func (c *Client[T]) WithToken(tok string) *Client[T] {
	c.AccessToken = tok
	return c
}

func (c *Client[T]) WithRefreshToken(tok string) *Client[T] {
	c.RefreshToken = tok
	return c
}

func (c *Client[T]) WithOnTokenRefresh(fn func(res *RefreshAccessTokenResponse, meta T)) *Client[T] {
	c.OnTokenRefresh = fn
	return c
}

func (c *Client[T]) WithMeta(meta T) *Client[T] {
	c.Meta = meta
	return c
}

// https://open.shopee.com/documents?module=87&type=2&id=58&version=2
func (c *Client[T]) makeSignature(req *http.Request) (string, int64) {
	ts := time.Now().Unix()
	path := req.URL.Path

	var baseStr string

	u := req.URL

	query := u.Query()
	query.Set("partner_id", fmt.Sprintf("%v", c.app.PartnerID))

	isPublicApi := false
	if strings.Contains(path, "/auth/token/get") || strings.Contains(path, "/auth/access_token/get") {
		isPublicApi = true
	}

	if c.ShopID != 0 && !isPublicApi {
		// Shop APIs: partner_id, api path, timestamp, access_token, shop_id
		baseStr = fmt.Sprintf("%d%s%d%s%d", c.app.PartnerID, path, ts, c.AccessToken, c.ShopID)
		query.Set("shop_id", fmt.Sprintf("%v", c.ShopID))
		query.Set("access_token", c.AccessToken)
	} else if c.MerchantID != 0 && !isPublicApi {
		// Merchant APIs: partner_id, api path, timestamp, access_token, merchant_id
		baseStr = fmt.Sprintf("%d%s%d%s%d", c.app.PartnerID, path, ts, c.AccessToken, c.MerchantID)
		query.Set("merchant_id", fmt.Sprintf("%v", c.MerchantID))
		query.Set("access_token", c.AccessToken)
	} else {
		// Public APIs: partner_id, api path, timestamp
		baseStr = fmt.Sprintf("%d%s%d", c.app.PartnerID, path, ts)
	}
	h := hmac.New(sha256.New, []byte(c.app.PartnerKey))
	h.Write([]byte(baseStr))
	result := hex.EncodeToString(h.Sum(nil))

	query.Set("timestamp", fmt.Sprintf("%v", ts))
	query.Set("sign", result)

	u.RawQuery = query.Encode()
	req.URL = u

	return result, ts
}

// doGetHeaders executes a request, decoding the response into `v` and also returns any response headers.
func (c *Client[T]) doGetHeaders(req *http.Request, v interface{}, skipBody bool) (http.Header, error) {
	var resp *http.Response
	var err error

	retries := c.retries
	c.attempts = 0
	c.logRequest(req, skipBody)

	for {
		c.attempts++

		fmt.Printf("DEBUG: doGetHeaders calling Do, URL: %s\n", req.URL.String())
		resp, err = c.Client.Do(req)
		c.logResponse(resp)
		if err != nil {
			return nil, err //http client errors, not api responses
		}

		respErr := CheckResponseError(resp)
		if respErr == nil {
			break // no errors, break out of the retry loop
		}

		fmt.Printf("DEBUG: doGetHeaders respErr: %T, value: %v\n", respErr, respErr)

		// handle auto token refresh if refresh token is present and the error is token related
		if c.RefreshToken != "" && !strings.Contains(req.URL.Path, "/auth/access_token/get") {
			var shopeeErr string
			if re, ok := respErr.(ResponseError); ok {
				shopeeErr = re.ShopeeError
			} else if re, ok := respErr.(*ResponseError); ok {
				shopeeErr = re.ShopeeError
			}

			if shopeeErr == "error_invalid_access_token" || shopeeErr == "error_access_token_expired" || shopeeErr == "invalid_access_token" || shopeeErr == "invalid_acceess_token" {
				fmt.Printf("DEBUG: Entering refresh token logic for error: %s\n", shopeeErr)
				refreshRes, err := c.Auth.RefreshAccessToken(c.ShopID, c.MerchantID, c.RefreshToken)
				if err == nil {
					fmt.Printf("DEBUG: Token refresh successful. New AccessToken: %s...\n", refreshRes.AccessToken[:10])
					c.AccessToken = refreshRes.AccessToken
					c.RefreshToken = refreshRes.RefreshToken
					if c.OnTokenRefresh != nil {
						c.OnTokenRefresh(refreshRes, c.Meta)
					}

					// resign the request
					c.makeSignature(req)

					// retry scenario, close resp and any continue will retry
					resp.Body.Close()
					continue
				} else {
					fmt.Printf("DEBUG: Token refresh failed: %v\n", err)
				}
			} else if shopeeErr != "" {
				fmt.Printf("DEBUG: Shopee error %s does not match refreshable errors\n", shopeeErr)
			}
		} else {
			if strings.Contains(req.URL.Path, "/auth/access_token/get") {
				// skip refresh logic for refresh call itself
			} else if c.RefreshToken == "" {
				fmt.Println("DEBUG: Skipping refresh logic because RefreshToken is empty in client")
			}
		}

		// retry scenario, close resp and any continue will retry
		resp.Body.Close()

		if retries <= 1 {
			return nil, respErr
		}

		if rateLimitErr, isRetryErr := respErr.(RateLimitError); isRetryErr {
			// back off and retry

			wait := time.Duration(rateLimitErr.RetryAfter) * time.Second
			c.log.Debugf("rate limited waiting %s", wait.String())
			time.Sleep(wait)
			retries--
			continue
		}

		var doRetry bool
		switch resp.StatusCode {
		case http.StatusServiceUnavailable:
			c.log.Debugf("service unavailable, retrying")
			doRetry = true
			retries--
		}

		if doRetry {
			continue
		}

		// no retry attempts, just return the err
		return nil, respErr
	}

	c.logResponse(resp)
	defer resp.Body.Close()

	if v != nil {
		decoder := json.NewDecoder(resp.Body)
		err := decoder.Decode(v)
		if err != nil {
			return nil, err
		}
	}

	return resp.Header, nil
}

// skipBody: if upload image, skip log its binary
func (c *Client[T]) logRequest(req *http.Request, skipBody bool) {
	if req == nil {
		return
	}
	if req.URL != nil {
		c.log.Debugf("%s: %s", req.Method, req.URL.String())
	}
	if !skipBody {
		c.logBody(&req.Body, "SENT: %s")
	}
}

func (c *Client[T]) logResponse(res *http.Response) {
	if res == nil {
		return
	}
	c.log.Debugf("RECV %d: %s", res.StatusCode, res.Status)
	c.logBody(&res.Body, "RESP: %s")
}

func (c *Client[T]) logBody(body *io.ReadCloser, format string) {
	if body == nil || *body == nil {
		return
	}
	b, _ := io.ReadAll(*body)
	if len(b) > 0 {
		c.log.Debugf(format, string(b))
	}
	*body = io.NopCloser(bytes.NewBuffer(b))
}

func wrapSpecificError(r *http.Response, err ResponseError) error {
	// TODO: check rate-limit error for shopee
	if err.Status == http.StatusTooManyRequests {
		f, _ := strconv.ParseFloat(r.Header.Get("Retry-After"), 64)
		return RateLimitError{
			ResponseError: err,
			RetryAfter:    int(f),
		}
	}

	// if err.Status == http.StatusSeeOther {
	// todo
	// The response to the request can be found under a different URL in the
	// Location header and can be retrieved using a GET method on that resource.
	// }

	if err.Status == http.StatusNotAcceptable {
		err.Message = http.StatusText(err.Status)
	}

	return err
}

// shopee error maybe return status=200
// eg. {"error":"error_incalid_category.","message":"Invalid category ID","request_id":"2069449bd255af166cb52b0e15189d6d"}
// {"error":"error_category_is_block.","message":"Category is restricted","request_id":"97994a47af37a22da79cb910bfd9841a"}
func CheckResponseError(r *http.Response) error {
	shopeeError := struct {
		Error     string `json:"error"`
		Message   string `json:"message"`
		RequestID string `json:"request_id"`
	}{}

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	defer func() {
		// already read out, reload for next process
		r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	}()

	if len(bodyBytes) > 0 {
		err := json.Unmarshal(bodyBytes, &shopeeError)
		if err != nil {
			// If status is 200 OK, and it doesn't look like JSON error, assume success
			if r.StatusCode == http.StatusOK {
				return nil
			}
			return ResponseDecodingError{
				Body:    bodyBytes,
				Message: err.Error(),
				Status:  r.StatusCode,
			}
		}
	}

	if shopeeError.Error == "" && http.StatusOK <= r.StatusCode && r.StatusCode < http.StatusMultipleChoices {
		return nil
	}

	responseError := ResponseError{
		Status:      r.StatusCode,
		Message:     shopeeError.Message,
		ShopeeError: shopeeError.Error,
		RequestID:   shopeeError.RequestID,
	}

	return wrapSpecificError(r, responseError)
}

// CreateAndDo performs a web request to Shopify with the given method (GET,
// POST, PUT, DELETE) and relative path (e.g. "/admin/orders.json").
// The data, options and resource arguments are optional and only relevant in
// certain situations.
// If the data argument is non-nil, it will be used as the body of the request
// for POST and PUT requests.
// The options argument is used for specifying request options such as search
// parameters like created_at_min
// Any data returned from Shopify will be marshalled into resource argument.
func (c *Client[T]) CreateAndDo(method, relPath string, data, options, headers, resource interface{}) error {
	_, err := c.createAndDoGetHeaders(method, relPath, data, options, headers, resource)
	if err != nil {
		return err
	}
	return nil
}

// createAndDoGetHeaders creates an executes a request while returning the response headers.
func (c *Client[T]) createAndDoGetHeaders(method, relPath string, data, options, headers, resource interface{}) (http.Header, error) {
	if strings.HasPrefix(relPath, "/") {
		// make sure it's a relative path
		relPath = strings.TrimLeft(relPath, "/")
	}

	relPath = path.Join("api/v2", relPath)

	req, err := c.NewRequest(method, relPath, data, options, headers)
	if err != nil {
		return nil, err
	}

	return c.doGetHeaders(req, resource, false)
}

// Get performs a GET request for the given path and saves the result in the
// given resource.
func (c *Client[T]) Get(path string, resource, options interface{}) error {
	return c.CreateAndDo("GET", path, nil, options, nil, resource)
}

// Post performs a POST request for the given path and saves the result in the
// given resource.
func (c *Client[T]) Post(path string, data, resource interface{}) error {
	return c.CreateAndDo("POST", path, data, nil, nil, resource)
}

// Put performs a PUT request for the given path and saves the result in the
// given resource.
func (c *Client[T]) Put(path string, data, resource interface{}) error {
	return c.CreateAndDo("PUT", path, data, nil, nil, resource)
}

// Delete performs a DELETE request for the given path
func (c *Client[T]) Delete(path string) error {
	return c.CreateAndDo("DELETE", path, nil, nil, nil, nil)
}

// Upload performs a Upload request for the given path and saves the result in the
// given resource.
func (c *Client[T]) Upload(relPath, fieldname, filename string, resource interface{}) error {
	req, err := c.NewfileUploadRequest(relPath, fieldname, filename)
	if err != nil {
		return err
	}

	if _, err := c.doGetHeaders(req, resource, true); err != nil {
		return err
	}
	return nil
}

// UploadFromReader performs an upload from an io.Reader
func (c *Client[T]) UploadFromReader(relPath, fieldname, filename string, reader io.Reader, resource interface{}) error {
	req, err := c.NewUploadFromReaderRequest(relPath, fieldname, filename, reader)
	if err != nil {
		return err
	}

	if _, err := c.doGetHeaders(req, resource, true); err != nil {
		return err
	}
	return nil
}

// NewfileUploadRequest creates a new file upload http request with optional extra params
func (c *Client[T]) NewfileUploadRequest(relPath, paramName, filename string) (*http.Request, error) {
	if strings.HasPrefix(relPath, "/") {
		// make sure it's a relative path
		relPath = strings.TrimLeft(relPath, "/")
	}

	relPath = path.Join("api/v2", relPath)

	rel, err := url.Parse(relPath)
	if err != nil {
		return nil, err
	}

	// Make the full url based on the relative path
	u := c.baseURL.ResolveReference(rel)
	uri := u.String()

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile(paramName, filepath.Base(filename))
	if err != nil {
		return nil, err
	}
	if _, err = io.Copy(part, file); err != nil {
		return nil, err
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", uri, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", UserAgent)

	c.makeSignature(req)

	return req, nil
}

// NewUploadFromReaderRequest creates a new file upload http request from an io.Reader
func (c *Client[T]) NewUploadFromReaderRequest(relPath, paramName, filename string, reader io.Reader) (*http.Request, error) {
	if strings.HasPrefix(relPath, "/") {
		// make sure it's a relative path
		relPath = strings.TrimLeft(relPath, "/")
	}

	relPath = path.Join("api/v2", relPath)

	rel, err := url.Parse(relPath)
	if err != nil {
		return nil, err
	}

	// Make the full url based on the relative path
	u := c.baseURL.ResolveReference(rel)
	uri := u.String()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile(paramName, filepath.Base(filename))
	if err != nil {
		return nil, err
	}
	if _, err = io.Copy(part, reader); err != nil {
		return nil, err
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", uri, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", UserAgent)

	c.makeSignature(req)

	return req, nil
}

// BoolString is a custom type that handles Shopee's string boolean values (e.g., "TRUE", "FALSE")
type BoolString bool

func (bs *BoolString) UnmarshalJSON(data []byte) error {
	s := strings.Trim(string(data), "\"")
	if strings.ToUpper(s) == "TRUE" {
		*bs = true
		return nil
	}
	if strings.ToUpper(s) == "FALSE" {
		*bs = false
		return nil
	}

	// Fallback to standard bool unmarshal
	var b bool
	if err := json.Unmarshal(data, &b); err != nil {
		return err
	}
	*bs = BoolString(b)
	return nil
}

func (bs BoolString) String() string {
	if bs {
		return "TRUE"
	}
	return "FALSE"
}
