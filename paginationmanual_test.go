// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package beeperdesktopapi_test

import (
	"context"
	"os"
	"testing"

	"github.com/beeper/desktop-api-go/v5"
	"github.com/beeper/desktop-api-go/v5/internal/testutil"
	"github.com/beeper/desktop-api-go/v5/option"
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
		AccountIDs: []string{"discordgo", "local-whatsapp_ba_EvYDBBsZbRQAy3UOSWqG0LuTVkc"},
		Limit:      beeperdesktopapi.Int(10),
		Query:      beeperdesktopapi.String("oauth"),
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	for _, message := range page.Items {
		t.Logf("%+v\n", message.ID)
	}
	// The mock server isn't going to give us real pagination
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
