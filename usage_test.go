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
		AccountIDs:   []string{"matrix", "discordgo", "local-whatsapp_ba_EvYDBBsZbRQAy3UOSWqG0LuTVkc"},
		IncludeMuted: beeperdesktopapi.Bool(true),
		Limit:        beeperdesktopapi.Int(3),
		Type:         beeperdesktopapi.ChatSearchParamsTypeSingle,
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", page)
}
