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
