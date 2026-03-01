// Code generated. DO NOT EDIT.
package goshopee

type PushService interface {
	// SetAppPushConfig you can turn on or turn off your app push config setting through this open api
	// https://open.shopee.com/documents/v2/v2.push.set_app_push_config?module=105&type=1
	SetAppPushConfig(sid uint64, req SetAppPushConfigRequest, tok string) (*SetAppPushConfigResponse, error)
	// GetAppPushConfig you can get your app current push config setting through this api
	// https://open.shopee.com/documents/v2/v2.push.get_app_push_config?module=105&type=1
	GetAppPushConfig(sid uint64, tok string) (*GetAppPushConfigResponse, error)
	// GetLostPushMessage Get the lost push messages that were lost within 3 days of the current time and not confirmed to have been consumed
	// https://open.shopee.com/documents/v2/v2.push.get_lost_push_message?module=105&type=1
	GetLostPushMessage(sid uint64, tok string) (*GetLostPushMessageResponse, error)
	// ConfirmConsumedLostPushMessage Confirm consumed lost push message
	// https://open.shopee.com/documents/v2/v2.push.confirm_consumed_lost_push_message?module=105&type=1
	ConfirmConsumedLostPushMessage(sid uint64, req ConfirmConsumedLostPushMessageRequest, tok string) (*ConfirmConsumedLostPushMessageResponse, error)
}

type PushServiceOp struct {
	client *Client
}

func (s *PushServiceOp) SetAppPushConfig(sid uint64, req SetAppPushConfigRequest, tok string) (*SetAppPushConfigResponse, error) {
	path := "/push/set_app_push_config"
	resp := new(SetAppPushConfigResponse)
	err := s.client.WithMerchant(sid, tok).Post(path, req, resp)
	return resp, err
}

func (s *PushServiceOp) GetAppPushConfig(sid uint64, tok string) (*GetAppPushConfigResponse, error) {
	path := "/push/get_app_push_config"
	resp := new(GetAppPushConfigResponse)
	err := s.client.WithMerchant(sid, tok).Get(path, resp, nil)
	return resp, err
}

func (s *PushServiceOp) GetLostPushMessage(sid uint64, tok string) (*GetLostPushMessageResponse, error) {
	path := "/push/get_lost_push_message"
	resp := new(GetLostPushMessageResponse)
	err := s.client.WithMerchant(sid, tok).Get(path, resp, nil)
	return resp, err
}

func (s *PushServiceOp) ConfirmConsumedLostPushMessage(sid uint64, req ConfirmConsumedLostPushMessageRequest, tok string) (*ConfirmConsumedLostPushMessageResponse, error) {
	path := "/push/confirm_consumed_lost_push_message"
	resp := new(ConfirmConsumedLostPushMessageResponse)
	err := s.client.WithMerchant(sid, tok).Post(path, req, resp)
	return resp, err
}
