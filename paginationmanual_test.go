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

func TestManualPagination(t *testing.T) {
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
	page, err := client.Messages.Search(context.TODO(), beeperdesktopapi.MessageSearchParams{
		AccountIDs: []string{"local-telegram_ba_QFrb5lrLPhO3OT5MFBeTWv0x4BI"},
		Limit:      beeperdesktopapi.Int(10),
		Query:      beeperdesktopapi.String("deployment"),
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	for _, message := range page.Items {
		t.Logf("%+v\n", message.ID)
	}
	// Prism mock isn't going to give us real pagination
	page, err = page.GetNextPage()
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if page != nil {
		for _, message := range page.Items {
			t.Logf("%+v\n", message.ID)
		}
	}
}
