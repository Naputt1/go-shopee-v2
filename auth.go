package goshopee

import "fmt"

// https://open.shopee.com/documents?module=87&type=2&id=58&version=2
type AuthService interface {
	// GetAuthURL returns the URL to authorize the app.
	// Path: /api/v2/shop/auth_partner
	GetAuthURL() (string, error)

	// GetCancelAuthURL returns the URL to cancel the authorization.
	// Path: /api/v2/shop/cancel_auth_partner
	GetCancelAuthURL() (string, error)

	// GetAccessToken gets the access token.
	// Path: /api/v2/auth/token/get
	GetAccessToken(sid uint64, aid uint64, code string) (*AccessTokenResponse, error)

	// RefreshAccessToken refreshes the access token.
	// Path: /api/v2/auth/access_token/get
	RefreshAccessToken(sid uint64, aid uint64, refresh string) (*RefreshAccessTokenResponse, error)
}

type AccessTokenResponse struct {
	BaseResponse

	AccessToken string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpireIn int `json:"expire_in"`
	MerchantIDList []uint64 `json:"merchant_id_list,omitempty"`
	ShopIDList []uint64 `json:"shop_id_list,omitempty"`
}

type RefreshAccessTokenResponse struct {
	BaseResponse

	AccessToken string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpireIn int `json:"expire_in"`
	PartnerID uint64 `json:"partner_id"`
	MerchantID uint64 `json:"merchant_id"`
	ShopID uint64 `json:"shop_id"`
}

type AuthServiceOp[T any] struct {
	client *Client[T]
}

// GetAuthURL returns the URL to authorize the app.
// Path: /api/v2/shop/auth_partner
func (s *AuthServiceOp[T]) GetAuthURL() (string, error) {
	rurl := s.client.app.RedirectURL
	path := "/api/v2/shop/auth_partner"
	sign, ts, _ := s.client.Util.Sign(path)
	aurl := fmt.Sprintf("%s%s?partner_id=%d&timestamp=%d&sign=%s&redirect=%s", s.client.app.APIURL, path, s.client.app.PartnerID, ts, sign, rurl)
	return aurl, nil
}

// GetCancelAuthURL returns the URL to cancel the authorization.
// Path: /api/v2/shop/cancel_auth_partner
func (s *AuthServiceOp[T]) GetCancelAuthURL() (string, error) {
	rurl := s.client.app.RedirectURL
	path := "/api/v2/shop/cancel_auth_partner"
	sign, ts, _ := s.client.Util.Sign(path)
	aurl := fmt.Sprintf("%s%s?partner_id=%d&timestamp=%d&sign=%s&redirect=%s", s.client.app.APIURL, path, s.client.app.PartnerID, ts, sign, rurl)
	return aurl, nil
}

// GetAccessToken gets the access token.
// Path: /api/v2/auth/token/get
func (s *AuthServiceOp[T]) GetAccessToken(sid uint64, aid uint64, code string) (*AccessTokenResponse, error) {
	path := "/auth/token/get"
	params := map[string]interface{}{
		"code":       code,
		"partner_id": s.client.app.PartnerID,
	}
	if sid!=0{
		params["shop_id"]=sid
	}else if aid!=0{
		params["main_account_id"]=aid
	}

	resp := new(AccessTokenResponse)
	err := s.client.Post(path, params, resp)
	return resp, err
}

// RefreshAccessToken refreshes the access token.
// Path: /api/v2/auth/access_token/get
func (s *AuthServiceOp[T]) RefreshAccessToken(sid uint64, aid uint64, refresh string) (*RefreshAccessTokenResponse, error) {
	path := "/auth/access_token/get"
	params := map[string]interface{}{
		"refresh_token": refresh,
		"partner_id":    s.client.app.PartnerID,
	}
	if sid!=0{
		params["shop_id"]=sid
	}else if aid!=0{
		params["main_account_id"]=aid
	}

	resp := new(RefreshAccessTokenResponse)
	err := s.client.Post(path, params, resp)
	return resp, err
}
