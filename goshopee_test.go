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
	client *Client
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
	client = NewClient(app,
		WithRetry(maxRetries))
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

