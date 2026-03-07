package goshopee

type ShopCategoryService interface {
	// AddShopCategory Use this call to add a new shop collecion
	// Path: /api/v2/shop_category/add_shop_category
	// https://open.shopee.com/documents/v2/v2.shop_category.add_shop_category?module=101&type=1
	AddShopCategory(sid uint64, req AddShopCategoryRequest, tok string) (*AddShopCategoryResponse, error)
	// GetShopCategoryList Use this call to get list of shop categories
	// Path: /api/v2/shop_category/get_shop_category_list
	// https://open.shopee.com/documents/v2/v2.shop_category.get_shop_category_list?module=101&type=1
	GetShopCategoryList(sid uint64, req GetShopCategoryListRequest, tok string) (*GetShopCategoryListResponse, error)
	// DeleteShopCategory Use this call to delete a existing shop collecion
	// Path: /api/v2/shop_category/delete_shop_category
	// https://open.shopee.com/documents/v2/v2.shop_category.delete_shop_category?module=101&type=1
	DeleteShopCategory(sid uint64, req DeleteShopCategoryRequest, tok string) (*DeleteShopCategoryResponse, error)
	// UpdateShopCategory Use this call to update a existing collecion
	// Path: /api/v2/shop_category/update_shop_category
	// https://open.shopee.com/documents/v2/v2.shop_category.update_shop_category?module=101&type=1
	UpdateShopCategory(sid uint64, req UpdateShopCategoryRequest, tok string) (*UpdateShopCategoryResponse, error)
	// AddItemList Use this call to add items list to certain shop_category
	// Path: /api/v2/shop_category/add_item_list
	// https://open.shopee.com/documents/v2/v2.shop_category.add_item_list?module=101&type=1
	AddItemList(sid uint64, req AddItemListRequest, tok string) (*AddItemListResponse, error)
	// GetItemList Use this call to get items list of certain shop_category
	// Path: /api/v2/shop_category/get_item_list
	// https://open.shopee.com/documents/v2/v2.shop_category.get_item_list?module=101&type=1
	GetItemList(sid uint64, req ShopCategoryGetItemListRequest, tok string) (*ShopCategoryGetItemListResponse, error)
	// DeleteItemList Use this api to delete items from shop category
	// Path: /api/v2/shop_category/delete_item_list
	// https://open.shopee.com/documents/v2/v2.shop_category.delete_item_list?module=101&type=1
	DeleteItemList(sid uint64, req DeleteItemListRequest, tok string) (*DeleteItemListResponse, error)
}

type ShopCategoryServiceOp[T any] struct {
	client *Client[T]
}

// AddShopCategory Use this call to add a new shop collecion
// Path: /api/v2/shop_category/add_shop_category
// https://open.shopee.com/documents/v2/v2.shop_category.add_shop_category?module=101&type=1
func (s *ShopCategoryServiceOp[T]) AddShopCategory(sid uint64, req AddShopCategoryRequest, tok string) (*AddShopCategoryResponse, error) {
	path := "/shop_category/add_shop_category"
	resp := new(AddShopCategoryResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

// GetShopCategoryList Use this call to get list of shop categories
// Path: /api/v2/shop_category/get_shop_category_list
// https://open.shopee.com/documents/v2/v2.shop_category.get_shop_category_list?module=101&type=1
func (s *ShopCategoryServiceOp[T]) GetShopCategoryList(sid uint64, req GetShopCategoryListRequest, tok string) (*GetShopCategoryListResponse, error) {
	path := "/shop_category/get_shop_category_list"
	resp := new(GetShopCategoryListResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

// DeleteShopCategory Use this call to delete a existing shop collecion
// Path: /api/v2/shop_category/delete_shop_category
// https://open.shopee.com/documents/v2/v2.shop_category.delete_shop_category?module=101&type=1
func (s *ShopCategoryServiceOp[T]) DeleteShopCategory(sid uint64, req DeleteShopCategoryRequest, tok string) (*DeleteShopCategoryResponse, error) {
	path := "/shop_category/delete_shop_category"
	resp := new(DeleteShopCategoryResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

// UpdateShopCategory Use this call to update a existing collecion
// Path: /api/v2/shop_category/update_shop_category
// https://open.shopee.com/documents/v2/v2.shop_category.update_shop_category?module=101&type=1
func (s *ShopCategoryServiceOp[T]) UpdateShopCategory(sid uint64, req UpdateShopCategoryRequest, tok string) (*UpdateShopCategoryResponse, error) {
	path := "/shop_category/update_shop_category"
	resp := new(UpdateShopCategoryResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

// AddItemList Use this call to add items list to certain shop_category
// Path: /api/v2/shop_category/add_item_list
// https://open.shopee.com/documents/v2/v2.shop_category.add_item_list?module=101&type=1
func (s *ShopCategoryServiceOp[T]) AddItemList(sid uint64, req AddItemListRequest, tok string) (*AddItemListResponse, error) {
	path := "/shop_category/add_item_list"
	resp := new(AddItemListResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

// GetItemList Use this call to get items list of certain shop_category
// Path: /api/v2/shop_category/get_item_list
// https://open.shopee.com/documents/v2/v2.shop_category.get_item_list?module=101&type=1
func (s *ShopCategoryServiceOp[T]) GetItemList(sid uint64, req ShopCategoryGetItemListRequest, tok string) (*ShopCategoryGetItemListResponse, error) {
	path := "/shop_category/get_item_list"
	resp := new(ShopCategoryGetItemListResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

// DeleteItemList Use this api to delete items from shop category
// Path: /api/v2/shop_category/delete_item_list
// https://open.shopee.com/documents/v2/v2.shop_category.delete_item_list?module=101&type=1
func (s *ShopCategoryServiceOp[T]) DeleteItemList(sid uint64, req DeleteItemListRequest, tok string) (*DeleteItemListResponse, error) {
	path := "/shop_category/delete_item_list"
	resp := new(DeleteItemListResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}
