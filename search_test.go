// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package beeperdesktopapi_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/beeper/desktop-api-go"
	"github.com/beeper/desktop-api-go/internal/testutil"
	"github.com/beeper/desktop-api-go/option"
)

func TestSearchChatsWithOptionalParams(t *testing.T) {
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
	_, err := client.Search.Chats(context.TODO(), beeperdesktopapi.SearchChatsParams{
		AccountIDs:         []string{"local-whatsapp_ba_EvYDBBsZbRQAy3UOSWqG0LuTVkc", "local-telegram_ba_QFrb5lrLPhO3OT5MFBeTWv0x4BI"},
		Cursor:             beeperdesktopapi.String("1725489123456|c29tZUltc2dQYWdl"),
		Direction:          beeperdesktopapi.SearchChatsParamsDirectionBefore,
		Inbox:              beeperdesktopapi.SearchChatsParamsInboxPrimary,
		IncludeMuted:       beeperdesktopapi.Bool(true),
		LastActivityAfter:  beeperdesktopapi.Time(time.Now()),
		LastActivityBefore: beeperdesktopapi.Time(time.Now()),
		Limit:              beeperdesktopapi.Int(1),
		Query:              beeperdesktopapi.String("x"),
		Scope:              beeperdesktopapi.SearchChatsParamsScopeTitles,
		Type:               beeperdesktopapi.SearchChatsParamsTypeSingle,
		UnreadOnly:         beeperdesktopapi.Bool(true),
	})
	if err != nil {
		var apierr *beeperdesktopapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSearchContacts(t *testing.T) {
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
	_, err := client.Search.Contacts(
		context.TODO(),
		"local-whatsapp_ba_EvYDBBsZbRQAy3UOSWqG0LuTVkc",
		beeperdesktopapi.SearchContactsParams{
			Query: "x",
		},
	)
	if err != nil {
		var apierr *beeperdesktopapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
