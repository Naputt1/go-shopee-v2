package goshopee

import (
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
)

func Test_GetCategory(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/product/get_category", app.APIURL),
		httpmock.NewBytesResponder(200, loadFixture("category_list.json")))

	res, err := client.Product.GetCategory(shopID, "zh-hant", accessToken)
	if err != nil {
		t.Errorf("Product.GetCategory error: %s", err)
	}

	t.Logf("Product.GetCategory: %#v", res)

	var expectedID uint64 = 123
	if res.Response.CategoryList[0].CategoryID != expectedID {
		t.Errorf("CategoryID returned %+v, expected %+v", res.Response.CategoryList[0].CategoryID, expectedID)
	}
}

func Test_GetCategoryTree(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/product/get_category_tree", app.APIURL),
		httpmock.NewBytesResponder(200, loadFixture("category_list.json")))

	res, err := client.Product.GetCategoryTree(shopID, "zh-hant", accessToken)
	if err != nil {
		t.Errorf("Product.GetCategoryTree error: %s", err)
	}

	t.Logf("Product.GetCategoryTree: %#v", res)

	var expectedID uint64 = 123
	if res.Response.CategoryList[0].CategoryID != expectedID {
		t.Errorf("CategoryID returned %+v, expected %+v", res.Response.CategoryList[0].CategoryID, expectedID)
	}
}

func Test_GetBrandList(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/product/get_brand_list", app.APIURL),
		httpmock.NewBytesResponder(200, loadFixture("brand_list.json")))

	res, err := client.Product.GetBrandList(shopID, 123, 1, 0, 10, accessToken)
	if err != nil {
		t.Errorf("Product.GetBrandList error: %s", err)
	}

	t.Logf("Product.GetBrandList: %#v", res)

	var expectedID uint64 = 2500139861
	if res.Response.BrandList[0].BrandID != expectedID {
		t.Errorf("BrandID returned %+v, expected %+v", res.Response.BrandList[0].BrandID, expectedID)
	}
}

func Test_GetDTSLimit(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/product/get_dts_limit", app.APIURL),
		httpmock.NewBytesResponder(200, loadFixture("dts_limit.json")))

	res, err := client.Product.GetDTSLimit(shopID, 123, accessToken)
	if err != nil {
		t.Errorf("Product.GetDTSLimit error: %s", err)
	}

	t.Logf("Product.GetDTSLimit: %#v", res)

	var expected int = 7
	if res.Response.DaysToShipLimit.MaxLimit != expected {
		t.Errorf("DaysToShipLimit.MaxLimit returned %+v, expected %+v", res.Response.DaysToShipLimit.MaxLimit, expected)
	}
}

func Test_GetAttributes(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/product/get_attributes", app.APIURL),
		httpmock.NewBytesResponder(200, loadFixture("attributes.json")))

	res, err := client.Product.GetAttributes(shopID, 123, "en", accessToken)
	if err != nil {
		t.Errorf("Product.GetAttributes error: %s", err)
	}

	t.Logf("Product.GetAttributes: %#v", res)

	var expectedID uint64 = 123
	if res.Response.AttributeList[0].AttributeID != expectedID {
		t.Errorf("AttributeList[0].AttributeID returned %+v, expected %+v", res.Response.AttributeList[0].AttributeID, expectedID)
	}
}

func Test_GetAttributeTree(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/product/get_attribute_tree", app.APIURL),
		httpmock.NewBytesResponder(200, loadFixture("attributeTree.json")))

	res, err := client.Product.GetAttributeTree(shopID, []uint64{123}, "en", accessToken)
	if err != nil {
		t.Errorf("Product.GetAttributeTree error: %s", err)
	}

	t.Logf("Product.GetAttributeTree: %#v", res)

	var expectedID uint64 = 123
	if res.Response.List[0].CategoryID != expectedID {
		t.Errorf("List[0].CategoryID returned %+v, expected %+v", res.Response.List[0].CategoryID, expectedID)
	}
}

func Test_GetSizeChartList(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/product/get_size_chart_list", app.APIURL),
		httpmock.NewBytesResponder(200, loadFixture("support_size_chart.json"))) // reusing fixture for simplicity if structure matches

	res, err := client.Product.GetSizeChartList(shopID, 123, accessToken)
	if err != nil {
		t.Errorf("Product.GetSizeChartList error: %s", err)
	}

	t.Logf("Product.GetSizeChartList: %#v", res)
}

func Test_SupportSizeChart_Deprecated(t *testing.T) {
	setup()
	defer teardown()

	_, err := client.Product.SupportSizeChart(shopID, 123, accessToken)
	if err == nil {
		t.Errorf("Product.SupportSizeChart should return error as it is deprecated")
	}
}

func Test_UpdateSizeChart_Deprecated(t *testing.T) {
	setup()
	defer teardown()

	_, err := client.Product.UpdateSizeChart(shopID, 123, "test", accessToken)
	if err == nil {
		t.Errorf("Product.UpdateSizeChart should return error as it is deprecated")
	}
}

func Test_GetItemList(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/product/get_item_list", app.APIURL),
		httpmock.NewBytesResponder(200, loadFixture("get_item_list_resp.json")))

	res, err := client.Product.GetItemList(shopID, GetItemListRequest{
		Offset:     0,
		PageSize:   10,
		ItemStatus: []ItemStatus{ItemStatusNormal},
	}, accessToken)
	if err != nil {
		t.Errorf("Product.GetItemList error: %s", err)
	}

	t.Logf("Product.GetItemList: %#v", res)

	var expectedID uint64 = 12345
	if res.Response.Item[0].ItemID != expectedID {
		t.Errorf("ItemID returned %+v, expected %+v", res.Response.Item[0].ItemID, expectedID)
	}
}

func Test_AddItem(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/product/add_item", app.APIURL),
		httpmock.NewBytesResponder(200, loadFixture("add_item_resp.json")))

	var req AddItemRequest
	loadMockData("add_item_req.json", &req)

	res, err := client.Product.AddItem(shopID, req, accessToken)
	if err != nil {
		t.Errorf("Product.AddItem error: %s", err)
	}

	t.Logf("Product.AddItem: %#v", res)

	var expectedID uint64 = 3000142341
	if res.Response.ItemID != expectedID {
		t.Errorf("ItemID returned %+v, expected %+v", res.Response.ItemID, expectedID)
	}
}

func Test_InitTierVariation(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/product/init_tier_variation", app.APIURL),
		httpmock.NewBytesResponder(200, loadFixture("init_tier_variation_resp.json")))

	var req InitTierVariationRequest
	// We might need to adjust mock data loading due to struct changes
	loadMockData("init_tier_variation_req.json", &req)

	res, err := client.Product.InitTierVariation(shopID, req, accessToken)
	if err != nil {
		t.Errorf("Product.InitTierVariation error: %s", err)
	}

	t.Logf("Product.InitTierVariation: %#v", res)

	var expectedID uint64 = 12345
	if res.Response.Model[0].ModelID != expectedID {
		t.Errorf("ModelID returned %+v, expected %+v", res.Response.Model[0].ModelID, expectedID)
	}
}

func Test_AddModel(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/product/add_model", app.APIURL),
		httpmock.NewBytesResponder(200, loadFixture("add_model_resp.json")))

	var req AddModelRequest
	loadMockData("add_model_req.json", &req)

	res, err := client.Product.AddModel(shopID, req, accessToken)
	if err != nil {
		t.Errorf("Product.AddModel error: %s", err)
	}

	t.Logf("Product.AddModel: %#v", res)
}

func Test_GetModelList(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/product/get_model_list", app.APIURL),
		httpmock.NewBytesResponder(200, loadFixture("get_model_list_resp.json")))

	res, err := client.Product.GetModelList(shopID, 123, accessToken)
	if err != nil {
		t.Errorf("Product.GetModelList error: %s", err)
	}

	t.Logf("Product.GetModelList: %#v", res)

	var expected uint64 = 2000458802
	if res.Response.Model[0].ModelID != expected {
		t.Errorf("ModelID returned %+v, expected %+v", res.Response.Model[0].ModelID, expected)
	}
}

func Test_GetItemBaseInfo(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/product/get_item_base_info", app.APIURL),
		httpmock.NewBytesResponder(200, loadFixture("get_item_base_info_resp.json")))

	res, err := client.Product.GetItemBaseInfo(shopID, []uint64{123, 356}, accessToken)
	if err != nil {
		t.Errorf("Product.GetItemBaseInfo error: %s", err)
	}

	t.Logf("Product.GetItemBaseInfo: %#v", res)

	var expected uint64 = 10311771375
	if res.Response.ItemList[0].ItemID != expected {
		t.Errorf("ItemID returned %+v, expected %+v", res.Response.ItemList[0].ItemID, expected)
	}
}

func Test_DeleteItem(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/product/delete_item", app.APIURL),
		httpmock.NewBytesResponder(200, loadFixture("response.json")))

	res, err := client.Product.DeleteItem(shopID, 34001, accessToken)
	if err != nil {
		t.Errorf("Product.DeleteItem error: %s", err)
	}

	t.Logf("Product.DeleteItem: %#v", res)
}

func Test_UpdateItem(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/product/update_item", app.APIURL),
		httpmock.NewBytesResponder(200, loadFixture("update_item_resp.json")))

	var req UpdateItemRequest
	loadMockData("update_item_req.json", &req)

	res, err := client.Product.UpdateItem(shopID, req, accessToken)
	if err != nil {
		t.Errorf("Product.UpdateItem error: %s", err)
	}

	t.Logf("Product.UpdateItem: %#v", res)
}

func Test_UnlistItem(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/product/unlist_item", app.APIURL),
		httpmock.NewBytesResponder(200, loadFixture("unlist_item_resp.json")))

	var req UnlistItemRequest
	loadMockData("unlist_item_req.json", &req)

	res, err := client.Product.UnlistItem(shopID, req, accessToken)
	if err != nil {
		t.Errorf("Product.UnlistItem error: %s", err)
	}

	t.Logf("Product.UnlistItem: %#v", res)
}

func Test_DeleteModel(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/product/delete_model", app.APIURL),
		httpmock.NewBytesResponder(200, loadFixture("delete_model_resp.json")))

	res, err := client.Product.DeleteModel(shopID, 34001, 123, accessToken)
	if err != nil {
		t.Errorf("Product.DeleteModel error: %s", err)
	}

	t.Logf("Product.DeleteModel: %#v", res)
}

func Test_UpdateModel(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/product/update_model", app.APIURL),
		httpmock.NewBytesResponder(200, loadFixture("response.json")))

	var req UpdateModelRequest
	loadMockData("update_model_req.json", &req)

	res, err := client.Product.UpdateModel(shopID, req, accessToken)
	if err != nil {
		t.Errorf("Product.UpdateModel error: %s", err)
	}

	t.Logf("Product.UpdateModel: %#v", res)
}

func Test_UpdatePrice(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/product/update_price", app.APIURL),
		httpmock.NewBytesResponder(200, loadFixture("update_price_resp.json")))

	var req UpdatePriceRequest
	loadMockData("update_price_req.json", &req)

	res, err := client.Product.UpdatePrice(shopID, req, accessToken)
	if err != nil {
		t.Errorf("Product.UpdatePrice error: %s", err)
	}

	t.Logf("Product.UpdatePrice: %#v", res)
}

func Test_UpdateStock(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/product/update_stock", app.APIURL),
		httpmock.NewBytesResponder(200, loadFixture("update_stock_resp.json")))

	var req UpdateStockRequest
	loadMockData("update_stock_req.json", &req)

	res, err := client.Product.UpdateStock(shopID, req, accessToken)
	if err != nil {
		t.Errorf("Product.UpdateStock error: %s", err)
	}

	t.Logf("Product.UpdateStock: %#v", res)
}

func Test_CategoryRecommend(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/product/category_recommend", app.APIURL),
		httpmock.NewBytesResponder(200, loadFixture("category_recommend_resp.json")))

	res, err := client.Product.CategoryRecommend(shopID, "test", accessToken)
	if err != nil {
		t.Errorf("Product.CategoryRecommend error: %s", err)
	}

	t.Logf("Product.CategoryRecommend: %#v", res)
}

func Test_GetItemPromotion(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/product/get_item_promotion", app.APIURL),
		httpmock.NewBytesResponder(200, loadFixture("get_item_promotion_resp.json")))

	res, err := client.Product.GetItemPromotion(shopID, []uint64{123}, accessToken)
	if err != nil {
		t.Errorf("Product.GetItemPromotion error: %s", err)
	}

	t.Logf("Product.GetItemPromotion: %#v", res)
}

func Test_UpdateTierVariation(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/product/update_tier_variation", app.APIURL),
		httpmock.NewBytesResponder(200, loadFixture("response.json")))

	var req UpdateTierVariationRequest
	loadMockData("update_tier_variation_req.json", &req)

	res, err := client.Product.UpdateTierVariation(shopID, req, accessToken)
	if err != nil {
		t.Errorf("Product.UpdateTierVariation error: %s", err)
	}

	t.Logf("Product.UpdateTierVariation: %#v", res)
}

func Test_SearchItem(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/product/search_item", app.APIURL),
		httpmock.NewBytesResponder(200, loadFixture("get_item_list_resp.json"))) // reuse similar fixture

	res, err := client.Product.SearchItem(shopID, SearchItemRequest{PageSize: 10}, accessToken)
	if err != nil {
		t.Errorf("Product.SearchItem error: %s", err)
	}

	t.Logf("Product.SearchItem: %#v", res)
}

func Test_BoostItem(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/product/boost_item", app.APIURL),
		httpmock.NewBytesResponder(200, loadFixture("response.json")))

	res, err := client.Product.BoostItem(shopID, []uint64{123}, accessToken)
	if err != nil {
		t.Errorf("Product.BoostItem error: %s", err)
	}

	t.Logf("Product.BoostItem: %#v", res)
}

func Test_GetBoostedList(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/product/get_boosted_list", app.APIURL),
		httpmock.NewBytesResponder(200, loadFixture("response.json")))

	res, err := client.Product.GetBoostedList(shopID, accessToken)
	if err != nil {
		t.Errorf("Product.GetBoostedList error: %s", err)
	}

	t.Logf("Product.GetBoostedList: %#v", res)
}

func Test_GetComment(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/product/get_comment", app.APIURL),
		httpmock.NewBytesResponder(200, loadFixture("response.json")))

	res, err := client.Product.GetComment(shopID, GetCommentRequest{ItemID: 123}, accessToken)
	if err != nil {
		t.Errorf("Product.GetComment error: %s", err)
	}

	t.Logf("Product.GetComment: %#v", res)
}

func Test_ReplyComment(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/product/reply_comment", app.APIURL),
		httpmock.NewBytesResponder(200, loadFixture("response.json")))

	res, err := client.Product.ReplyComment(shopID, ReplyCommentRequest{CommentID: 123, ReplyText: "test"}, accessToken)
	if err != nil {
		t.Errorf("Product.ReplyComment error: %s", err)
	}

	t.Logf("Product.ReplyComment: %#v", res)
}
