// Code generated. DO NOT EDIT.
package goshopee

type LivestreamService interface {
	// UploadImage Upload an image as the live stream cover.(For TW, ID, TH, PH, MY, SG, VN)
	// https://open.shopee.com/documents/v2/v2.livestream.upload_image?module=125&type=1
	UploadImage(sid uint64, req UploadImageRequest, tok string) (*SharedUploadImageResponse, error)
	// CreateSession Create a new live stream, include basic information, like cover, title, description, type (test live or normal live). (For TW, ID, TH, PH, MY, SG, VN)
	// https://open.shopee.com/documents/v2/v2.livestream.create_session?module=125&type=1
	CreateSession(sid uint64, req CreateSessionRequest, tok string) (*CreateSessionResponse, error)
	// UpdateSession Update live stream information, including cover, title, description, and type (test live or normal live). (For TW, ID, TH, PH, MY, SG, VN)
	// https://open.shopee.com/documents/v2/v2.livestream.update_session?module=125&type=1
	UpdateSession(sid uint64, req UpdateSessionRequest, tok string) (*UpdateSessionResponse, error)
	// StartSession Start Live. (For TW, ID, TH, PH, MY, SG, VN)
	// https://open.shopee.com/documents/v2/v2.livestream.start_session?module=125&type=1
	StartSession(sid uint64, req StartSessionRequest, tok string) (*StartSessionResponse, error)
	// EndSession End Live. (For TW, ID, TH, PH, MY, SG, VN)
	// https://open.shopee.com/documents/v2/v2.livestream.end_session?module=125&type=1
	EndSession(sid uint64, req EndSessionRequest, tok string) (*EndSessionResponse, error)
	// GetSessionDetail Get basic information about the live streaming room, including cover, title, description, type (test live or normal live), create time, update time, stream url, etc. (For TW, ID, TH, PH, MY, SG, VN)
	// https://open.shopee.com/documents/v2/v2.livestream.get_session_detail?module=125&type=1
	GetSessionDetail(sid uint64, opt GetSessionDetailRequest, tok string) (*GetSessionDetailResponse, error)
	// AddItemList Add items to item bag. (For TW, ID, TH, PH, MY, SG, VN)
	// https://open.shopee.com/documents/v2/v2.livestream.add_item_list?module=125&type=1
	AddItemList(sid uint64, req LivestreamAddItemListRequest, tok string) (*LivestreamAddItemListResponse, error)
	// DeleteItemList Delete items from item bag. (For TW, ID, TH, PH, MY, SG, VN)
	// https://open.shopee.com/documents/v2/v2.livestream.delete_item_list?module=125&type=1
	DeleteItemList(sid uint64, req LivestreamDeleteItemListRequest, tok string) (*LivestreamDeleteItemListResponse, error)
	// UpdateItemList Update the order of items in item bag. (For TW, ID, TH, PH, MY, SG, VN)
	// https://open.shopee.com/documents/v2/v2.livestream.update_item_list?module=125&type=1
	UpdateItemList(sid uint64, req UpdateItemListRequest, tok string) (*UpdateItemListResponse, error)
	// GetItemCount Get the number of items in the shopping bag, including the current number of items in the shopping bag, the upper limit of the number, etc. (For TW, ID, TH, PH, MY, SG, VN)
	// https://open.shopee.com/documents/v2/v2.livestream.get_item_count?module=125&type=1
	GetItemCount(sid uint64, opt GetItemCountRequest, tok string) (*GetItemCountResponse, error)
	// GetItemList Get the detail information of item in item bag, including item id, item serial number, etc.(For TW, ID, TH, PH, MY, SG, VN)
	// https://open.shopee.com/documents/v2/v2.livestream.get_item_list?module=125&type=1
	GetItemList(sid uint64, opt LivestreamGetItemListRequest, tok string) (*LivestreamGetItemListResponse, error)
	// UpdateShowItem Set the showing item. (For TW, ID, TH, PH, MY, SG, VN)
	// https://open.shopee.com/documents/v2/v2.livestream.update_show_item?module=125&type=1
	UpdateShowItem(sid uint64, req UpdateShowItemRequest, tok string) (*UpdateShowItemResponse, error)
	// DeleteShowItem Unshow showing item. (For TW, ID, TH, PH, MY, SG, VN)
	// https://open.shopee.com/documents/v2/v2.livestream.delete_show_item?module=125&type=1
	DeleteShowItem(sid uint64, req DeleteShowItemRequest, tok string) (*DeleteShowItemResponse, error)
	// GetShowItem Get the showing item. (For TW, ID, TH, PH, MY, SG, VN)
	// https://open.shopee.com/documents/v2/v2.livestream.get_show_item?module=125&type=1
	GetShowItem(sid uint64, opt GetShowItemRequest, tok string) (*GetShowItemResponse, error)
	// GetLikeItemList Get the item list of My Likes tab.(For TW, ID, TH, PH, MY, SG, VN)
	// https://open.shopee.com/documents/v2/v2.livestream.get_like_item_list?module=125&type=1
	GetLikeItemList(sid uint64, opt GetLikeItemListRequest, tok string) (*GetLikeItemListResponse, error)
	// GetRecentItemList Get the item list of the Recently tab. (For TW, ID, TH, PH, MY, SG, VN)
	// https://open.shopee.com/documents/v2/v2.livestream.get_recent_item_list?module=125&type=1
	GetRecentItemList(sid uint64, opt GetRecentItemListRequest, tok string) (*GetRecentItemListResponse, error)
	// GetItemSetList Get the product set of the live stream, including the product set name, id, and item number. (For TW, ID, TH, PH, MY, SG, VN)
	// https://open.shopee.com/documents/v2/v2.livestream.get_item_set_list?module=125&type=1
	GetItemSetList(sid uint64, opt GetItemSetListRequest, tok string) (*GetItemSetListResponse, error)
	// GetItemSetItemList Get the item list of the product set, including item name, id, etc. (For TW, ID, TH, PH, MY, SG, VN)
	// https://open.shopee.com/documents/v2/v2.livestream.get_item_set_item_list?module=125&type=1
	GetItemSetItemList(sid uint64, opt GetItemSetItemListRequest, tok string) (*GetItemSetItemListResponse, error)
	// ApplyItemSet Add product set to item bag. (For TW, ID, TH, PH, MY, SG, VN)
	// https://open.shopee.com/documents/v2/v2.livestream.apply_item_set?module=125&type=1
	ApplyItemSet(sid uint64, req ApplyItemSetRequest, tok string) (*ApplyItemSetResponse, error)
	// GetSessionMetric Get real-time indicator data of the live stream room, including the number of likes, comments, shares, views, etc.(For TW, ID, TH, PH, MY, SG, VN)
	// https://open.shopee.com/documents/v2/v2.livestream.get_session_metric?module=125&type=1
	GetSessionMetric(sid uint64, opt GetSessionMetricRequest, tok string) (*GetSessionMetricResponse, error)
	// GetSessionItemMetric Get real-time indicator data of live stream products, including product clicks, add-to-cart, etc. (For TW, ID, TH, PH, MY, SG, VN)
	// https://open.shopee.com/documents/v2/v2.livestream.get_session_item_metric?module=125&type=1
	GetSessionItemMetric(sid uint64, opt GetSessionItemMetricRequest, tok string) (*GetSessionItemMetricResponse, error)
	// GetLatestCommentList Get live stream room comments in the last 10 seconds, including user id, user name, comment id, comment content, and comment time. (For TW, ID, TH, PH, MY, SG, VN)
	// https://open.shopee.com/documents/v2/v2.livestream.get_latest_comment_list?module=125&type=1
	GetLatestCommentList(sid uint64, opt GetLatestCommentListRequest, tok string) (*GetLatestCommentListResponse, error)
	// PostComment Post comment in the live stream as streamer. (For TW, ID, TH, PH, MY, SG, VN)
	// https://open.shopee.com/documents/v2/v2.livestream.post_comment?module=125&type=1
	PostComment(sid uint64, req PostCommentRequest, tok string) (*PostCommentResponse, error)
	// BanUserComment Ban the user from posting comments. (For TW, ID, TH, PH, MY, SG, VN)
	// https://open.shopee.com/documents/v2/v2.livestream.ban_user_comment?module=125&type=1
	BanUserComment(sid uint64, req BanUserCommentRequest, tok string) (*BanUserCommentResponse, error)
	// UnbanUserComment Unban a user from posting comments. (For TW, ID, TH, PH, MY, SG, VN)
	// https://open.shopee.com/documents/v2/v2.livestream.unban_user_comment?module=125&type=1
	UnbanUserComment(sid uint64, req UnbanUserCommentRequest, tok string) (*UnbanUserCommentResponse, error)
}

type LivestreamServiceOp struct {
	client *Client
}

func (s *LivestreamServiceOp) UploadImage(sid uint64, req UploadImageRequest, tok string) (*SharedUploadImageResponse, error) {
	path := "/livestream/upload_image"
	resp := new(SharedUploadImageResponse)
	err := s.client.WithMerchant(sid, tok).Post(path, req, resp)
	return resp, err
}

func (s *LivestreamServiceOp) CreateSession(sid uint64, req CreateSessionRequest, tok string) (*CreateSessionResponse, error) {
	path := "/livestream/create_session"
	resp := new(CreateSessionResponse)
	err := s.client.WithMerchant(sid, tok).Post(path, req, resp)
	return resp, err
}

func (s *LivestreamServiceOp) UpdateSession(sid uint64, req UpdateSessionRequest, tok string) (*UpdateSessionResponse, error) {
	path := "/livestream/update_session"
	resp := new(UpdateSessionResponse)
	err := s.client.WithMerchant(sid, tok).Post(path, req, resp)
	return resp, err
}

func (s *LivestreamServiceOp) StartSession(sid uint64, req StartSessionRequest, tok string) (*StartSessionResponse, error) {
	path := "/livestream/start_session"
	resp := new(StartSessionResponse)
	err := s.client.WithMerchant(sid, tok).Post(path, req, resp)
	return resp, err
}

func (s *LivestreamServiceOp) EndSession(sid uint64, req EndSessionRequest, tok string) (*EndSessionResponse, error) {
	path := "/livestream/end_session"
	resp := new(EndSessionResponse)
	err := s.client.WithMerchant(sid, tok).Post(path, req, resp)
	return resp, err
}

func (s *LivestreamServiceOp) GetSessionDetail(sid uint64, opt GetSessionDetailRequest, tok string) (*GetSessionDetailResponse, error) {
	path := "/livestream/get_session_detail"
	resp := new(GetSessionDetailResponse)
	err := s.client.WithMerchant(sid, tok).Get(path, resp, opt)
	return resp, err
}

func (s *LivestreamServiceOp) AddItemList(sid uint64, req LivestreamAddItemListRequest, tok string) (*LivestreamAddItemListResponse, error) {
	path := "/livestream/add_item_list"
	resp := new(LivestreamAddItemListResponse)
	err := s.client.WithMerchant(sid, tok).Post(path, req, resp)
	return resp, err
}

func (s *LivestreamServiceOp) DeleteItemList(sid uint64, req LivestreamDeleteItemListRequest, tok string) (*LivestreamDeleteItemListResponse, error) {
	path := "/livestream/delete_item_list"
	resp := new(LivestreamDeleteItemListResponse)
	err := s.client.WithMerchant(sid, tok).Post(path, req, resp)
	return resp, err
}

func (s *LivestreamServiceOp) UpdateItemList(sid uint64, req UpdateItemListRequest, tok string) (*UpdateItemListResponse, error) {
	path := "/livestream/update_item_list"
	resp := new(UpdateItemListResponse)
	err := s.client.WithMerchant(sid, tok).Post(path, req, resp)
	return resp, err
}

func (s *LivestreamServiceOp) GetItemCount(sid uint64, opt GetItemCountRequest, tok string) (*GetItemCountResponse, error) {
	path := "/livestream/get_item_count"
	resp := new(GetItemCountResponse)
	err := s.client.WithMerchant(sid, tok).Get(path, resp, opt)
	return resp, err
}

func (s *LivestreamServiceOp) GetItemList(sid uint64, opt LivestreamGetItemListRequest, tok string) (*LivestreamGetItemListResponse, error) {
	path := "/livestream/get_item_list"
	resp := new(LivestreamGetItemListResponse)
	err := s.client.WithMerchant(sid, tok).Get(path, resp, opt)
	return resp, err
}

func (s *LivestreamServiceOp) UpdateShowItem(sid uint64, req UpdateShowItemRequest, tok string) (*UpdateShowItemResponse, error) {
	path := "/livestream/update_show_item"
	resp := new(UpdateShowItemResponse)
	err := s.client.WithMerchant(sid, tok).Post(path, req, resp)
	return resp, err
}

func (s *LivestreamServiceOp) DeleteShowItem(sid uint64, req DeleteShowItemRequest, tok string) (*DeleteShowItemResponse, error) {
	path := "/livestream/delete_show_item"
	resp := new(DeleteShowItemResponse)
	err := s.client.WithMerchant(sid, tok).Post(path, req, resp)
	return resp, err
}

func (s *LivestreamServiceOp) GetShowItem(sid uint64, opt GetShowItemRequest, tok string) (*GetShowItemResponse, error) {
	path := "/livestream/get_show_item"
	resp := new(GetShowItemResponse)
	err := s.client.WithMerchant(sid, tok).Get(path, resp, opt)
	return resp, err
}

func (s *LivestreamServiceOp) GetLikeItemList(sid uint64, opt GetLikeItemListRequest, tok string) (*GetLikeItemListResponse, error) {
	path := "/livestream/get_like_item_list"
	resp := new(GetLikeItemListResponse)
	err := s.client.WithMerchant(sid, tok).Get(path, resp, opt)
	return resp, err
}

func (s *LivestreamServiceOp) GetRecentItemList(sid uint64, opt GetRecentItemListRequest, tok string) (*GetRecentItemListResponse, error) {
	path := "/livestream/get_recent_item_list"
	resp := new(GetRecentItemListResponse)
	err := s.client.WithMerchant(sid, tok).Get(path, resp, opt)
	return resp, err
}

func (s *LivestreamServiceOp) GetItemSetList(sid uint64, opt GetItemSetListRequest, tok string) (*GetItemSetListResponse, error) {
	path := "/livestream/get_item_set_list"
	resp := new(GetItemSetListResponse)
	err := s.client.WithMerchant(sid, tok).Get(path, resp, opt)
	return resp, err
}

func (s *LivestreamServiceOp) GetItemSetItemList(sid uint64, opt GetItemSetItemListRequest, tok string) (*GetItemSetItemListResponse, error) {
	path := "/livestream/get_item_set_item_list"
	resp := new(GetItemSetItemListResponse)
	err := s.client.WithMerchant(sid, tok).Get(path, resp, opt)
	return resp, err
}

func (s *LivestreamServiceOp) ApplyItemSet(sid uint64, req ApplyItemSetRequest, tok string) (*ApplyItemSetResponse, error) {
	path := "/livestream/apply_item_set"
	resp := new(ApplyItemSetResponse)
	err := s.client.WithMerchant(sid, tok).Post(path, req, resp)
	return resp, err
}

func (s *LivestreamServiceOp) GetSessionMetric(sid uint64, opt GetSessionMetricRequest, tok string) (*GetSessionMetricResponse, error) {
	path := "/livestream/get_session_metric"
	resp := new(GetSessionMetricResponse)
	err := s.client.WithMerchant(sid, tok).Get(path, resp, opt)
	return resp, err
}

func (s *LivestreamServiceOp) GetSessionItemMetric(sid uint64, opt GetSessionItemMetricRequest, tok string) (*GetSessionItemMetricResponse, error) {
	path := "/livestream/get_session_item_metric"
	resp := new(GetSessionItemMetricResponse)
	err := s.client.WithMerchant(sid, tok).Get(path, resp, opt)
	return resp, err
}

func (s *LivestreamServiceOp) GetLatestCommentList(sid uint64, opt GetLatestCommentListRequest, tok string) (*GetLatestCommentListResponse, error) {
	path := "/livestream/get_latest_comment_list"
	resp := new(GetLatestCommentListResponse)
	err := s.client.WithMerchant(sid, tok).Get(path, resp, opt)
	return resp, err
}

func (s *LivestreamServiceOp) PostComment(sid uint64, req PostCommentRequest, tok string) (*PostCommentResponse, error) {
	path := "/livestream/post_comment"
	resp := new(PostCommentResponse)
	err := s.client.WithMerchant(sid, tok).Post(path, req, resp)
	return resp, err
}

func (s *LivestreamServiceOp) BanUserComment(sid uint64, req BanUserCommentRequest, tok string) (*BanUserCommentResponse, error) {
	path := "/livestream/ban_user_comment"
	resp := new(BanUserCommentResponse)
	err := s.client.WithMerchant(sid, tok).Post(path, req, resp)
	return resp, err
}

func (s *LivestreamServiceOp) UnbanUserComment(sid uint64, req UnbanUserCommentRequest, tok string) (*UnbanUserCommentResponse, error) {
	path := "/livestream/unban_user_comment"
	resp := new(UnbanUserCommentResponse)
	err := s.client.WithMerchant(sid, tok).Post(path, req, resp)
	return resp, err
}
