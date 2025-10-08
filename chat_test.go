// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package beeperdesktopapi_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/beeper/desktop-api-go"
	"github.com/beeper/desktop-api-go/internal/testutil"
	"github.com/beeper/desktop-api-go/option"
)

func TestChatNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Chats.New(context.TODO(), beeperdesktopapi.ChatNewParams{
		AccountID:      "local-whatsapp_ba_EvYDBBsZbRQAy3UOSWqG0LuTVkc",
		ParticipantIDs: []string{"string"},
		Type:           beeperdesktopapi.ChatNewParamsTypeSingle,
		MessageText:    beeperdesktopapi.String("messageText"),
		Title:          beeperdesktopapi.String("title"),
	})
	if err != nil {
		var apierr *beeperdesktopapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestChatGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Chats.Get(
		context.TODO(),
		"!NCdzlIaMjZUmvmvyHU:beeper.com",
		beeperdesktopapi.ChatGetParams{
			MaxParticipantCount: beeperdesktopapi.Int(50),
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

func TestChatListWithOptionalParams(t *testing.T) {
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
	_, err := client.Chats.List(context.TODO(), beeperdesktopapi.ChatListParams{
		AccountIDs: []string{"whatsapp", "local-whatsapp_ba_EvYDBBsZbRQAy3UOSWqG0LuTVkc", "local-instagram_ba_eRfQMmnSNy_p7Ih7HL7RduRpKFU"},
		Cursor:     beeperdesktopapi.String("1725489123456|c29tZUltc2dQYWdl"),
		Direction:  beeperdesktopapi.ChatListParamsDirectionBefore,
	})
	if err != nil {
		var apierr *beeperdesktopapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestChatArchiveWithOptionalParams(t *testing.T) {
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
	_, err := client.Chats.Archive(
		context.TODO(),
		"!NCdzlIaMjZUmvmvyHU:beeper.com",
		beeperdesktopapi.ChatArchiveParams{
			Archived: beeperdesktopapi.Bool(true),
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
