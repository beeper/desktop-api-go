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

func TestMessageUpdate(t *testing.T) {
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
	_, err := client.Messages.Update(
		context.TODO(),
		"messageID",
		beeperdesktopapi.MessageUpdateParams{
			ChatID: "!NCdzlIaMjZUmvmvyHU:beeper.com",
			Text:   "x",
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
	_, err := client.Messages.List(
		context.TODO(),
		"!NCdzlIaMjZUmvmvyHU:beeper.com",
		beeperdesktopapi.MessageListParams{
			Cursor:    beeperdesktopapi.String("1725489123456|c29tZUltc2dQYWdl"),
			Direction: beeperdesktopapi.MessageListParamsDirectionBefore,
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
		AccountIDs:         []string{"local-whatsapp_ba_EvYDBBsZbRQAy3UOSWqG0LuTVkc", "local-instagram_ba_eRfQMmnSNy_p7Ih7HL7RduRpKFU"},
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
		Sender:             beeperdesktopapi.String("sender"),
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
	_, err := client.Messages.Send(
		context.TODO(),
		"!NCdzlIaMjZUmvmvyHU:beeper.com",
		beeperdesktopapi.MessageSendParams{
			Attachment: beeperdesktopapi.MessageSendParamsAttachment{
				UploadID: "uploadID",
				Duration: beeperdesktopapi.Float(0),
				FileName: beeperdesktopapi.String("fileName"),
				MimeType: beeperdesktopapi.String("mimeType"),
				Size: beeperdesktopapi.MessageSendParamsAttachmentSize{
					Height: 0,
					Width:  0,
				},
				Type: "gif",
			},
			ReplyToMessageID: beeperdesktopapi.String("replyToMessageID"),
			Text:             beeperdesktopapi.String("text"),
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
