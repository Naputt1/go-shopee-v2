package goshopee

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"testing"

	"github.com/caarlos0/env/v11"
	"github.com/jarcoal/httpmock"
	"github.com/joho/godotenv"
)

const (
	maxRetries  = 3
	shopID      = 1234567
	merchantID  = 0
	accessToken = "accesstoken"
)

var (
	client *Client[any]
	app    App
)

func setup() {
	err := godotenv.Load()
	if err != nil {
		slog.Warn("Error loading .env file")
		app = App{
			PartnerID:   12345678,
			PartnerKey:  "hush",
			RedirectURL: "https://example.com/callback",
			APIURL:      "https://partner.test-stable.shopeemobile.com",
		}
	} else {
		if err := env.Parse(&app); err != nil {
			slog.Error("Error parsing env", "error", err)
			os.Exit(1)
		}
	}
	client = NewClient[any](app,
		WithRetry[any](maxRetries))
	httpmock.ActivateNonDefault(client.Client)
}

func teardown() {
	httpmock.DeactivateAndReset()
}

func loadFixture(filename string) []byte {
	f, err := os.ReadFile("fixtures/" + filename)
	if err != nil {
		panic(fmt.Sprintf("Cannot load fixture %v", filename))
	}
	return f
}

func loadMockData(filename string, out interface{}) {
	f, err := os.ReadFile("fixtures/" + filename)
	if err != nil {
		panic(fmt.Sprintf("Cannot load fixture %v", filename))
	}
	if err := json.Unmarshal(f, &out); err != nil {
		panic(fmt.Sprintf("decode mock data error: %s", err))
	}
}

func loadFixtureInterface(filename string) interface{} {
	var out interface{}
	f, err := os.ReadFile("fixtures/" + filename)
	if err != nil {
		return nil
	}
	json.Unmarshal(f, &out)
	return out
}

func TestCheckResponseError(t *testing.T) {
	setup()
	defer teardown()

	rawErr := `{"error":"error_auth","message":"Invalid access token","request_id":"req123"}`
	httpmock.RegisterResponder("GET", "https://partner.test-stable.shopeemobile.com/api/v2/test",
		httpmock.NewStringResponder(401, rawErr))

	resp, _ := client.Client.Get("https://partner.test-stable.shopeemobile.com/api/v2/test")
	defer resp.Body.Close()

	apiErr := CheckResponseError(resp)
	if apiErr == nil {
		t.Fatal("Expected error, got nil")
	}

	if !IsShopeeError(apiErr, ErrErrorAuth) {
		t.Errorf("Expected IsShopeeError to be true for ErrErrorAuth, got %v", apiErr)
	}

	re, ok := apiErr.(ResponseError)
	if !ok {
		t.Fatalf("Expected ResponseError, got %T", apiErr)
	}

	if re.ShopeeError != ErrErrorAuth {
		t.Errorf("Expected ShopeeError %s, got %s", ErrErrorAuth, re.ShopeeError)
	}

	if re.RequestID != "req123" {
		t.Errorf("Expected RequestID req123, got %s", re.RequestID)
	}
}

func TestGenericMeta(t *testing.T) {
	type MyMeta struct {
		ID   int
		Name string
	}

	meta := MyMeta{ID: 123, Name: "Test"}
	app := App{
		APIURL: "https://example.com",
	}

	var capturedMeta MyMeta
	c := NewClient(app,
		WithMeta(meta),
		WithOnTokenRefresh(func(res *RefreshAccessTokenResponse, m MyMeta) {
			capturedMeta = m
		}),
	)

	if c.Meta.ID != 123 {
		t.Errorf("Expected meta ID 123, got %d", c.Meta.ID)
	}

	// Trigger token refresh callback manually if possible, or just check the type safety
	c.OnTokenRefresh(nil, c.Meta)
	if capturedMeta.ID != 123 {
		t.Errorf("Expected captured meta ID 123, got %d", capturedMeta.ID)
	}
}

func TestNewDefaultClient(t *testing.T) {
	app := App{
		APIURL: "https://example.com",
	}

	c := NewDefaultClient(app,
		WithRetryDefault(3),
		WithMetaDefault("some meta"),
	)

	if c.retries != 3 {
		t.Errorf("Expected retries 3, got %d", c.retries)
	}

	if c.Meta != "some meta" {
		t.Errorf("Expected meta 'some meta', got %v", c.Meta)
	}
}
