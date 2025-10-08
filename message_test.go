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

func TestMessageListWithOptionalParams(t *testing.T) {
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
	_, err := client.Messages.List(context.TODO(), beeperdesktopapi.MessageListParams{
		ChatID:    "!NCdzlIaMjZUmvmvyHU:beeper.com",
		Cursor:    beeperdesktopapi.String("821744079"),
		Direction: beeperdesktopapi.MessageListParamsDirectionBefore,
	})
	if err != nil {
		var apierr *beeperdesktopapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMessageSearchWithOptionalParams(t *testing.T) {
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
	_, err := client.Messages.Search(context.TODO(), beeperdesktopapi.MessageSearchParams{
		AccountIDs:         []string{"whatsapp", "local-whatsapp_ba_EvYDBBsZbRQAy3UOSWqG0LuTVkc", "local-instagram_ba_eRfQMmnSNy_p7Ih7HL7RduRpKFU"},
		ChatIDs:            []string{"!NCdzlIaMjZUmvmvyHU:beeper.com", "1231073"},
		ChatType:           beeperdesktopapi.MessageSearchParamsChatTypeGroup,
		Cursor:             beeperdesktopapi.String("1725489123456|c29tZUltc2dQYWdl"),
		DateAfter:          beeperdesktopapi.Time(time.Now()),
		DateBefore:         beeperdesktopapi.Time(time.Now()),
		Direction:          beeperdesktopapi.MessageSearchParamsDirectionBefore,
		ExcludeLowPriority: beeperdesktopapi.Bool(true),
		IncludeMuted:       beeperdesktopapi.Bool(true),
		Limit:              beeperdesktopapi.Int(20),
		MediaTypes:         []string{"any"},
		Query:              beeperdesktopapi.String("dinner"),
		Sender:             beeperdesktopapi.MessageSearchParamsSenderMe,
	})
	if err != nil {
		var apierr *beeperdesktopapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMessageSendWithOptionalParams(t *testing.T) {
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
	_, err := client.Messages.Send(context.TODO(), beeperdesktopapi.MessageSendParams{
		ChatID:           beeperdesktopapi.String("!NCdzlIaMjZUmvmvyHU:beeper.com"),
		ReplyToMessageID: beeperdesktopapi.String("replyToMessageID"),
		Text:             beeperdesktopapi.String("text"),
	})
	if err != nil {
		var apierr *beeperdesktopapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
