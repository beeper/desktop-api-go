// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package beeperdesktopapi_test

import (
	"context"
	"os"
	"testing"

	"github.com/stainless-sdks/beeper-desktop-api-go"
	"github.com/stainless-sdks/beeper-desktop-api-go/internal/testutil"
	"github.com/stainless-sdks/beeper-desktop-api-go/option"
)

func TestUsage(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := beeperdesktopapi.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAccessToken("My Access Token"),
	)
	page, err := client.Chats.Search(context.TODO(), beeperdesktopapi.ChatSearchParams{
		IncludeMuted: beeperdesktopapi.Bool(true),
		Limit:        beeperdesktopapi.Int(3),
		Type:         beeperdesktopapi.ChatSearchParamsTypeSingle,
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", page)
}
