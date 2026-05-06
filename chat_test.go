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
		AccountID:      "accountID",
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
			MaxParticipantCount: beeperdesktopapi.Int(100),
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

func TestChatUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Chats.Update(
		context.TODO(),
		"!NCdzlIaMjZUmvmvyHU:beeper.com",
		beeperdesktopapi.ChatUpdateParams{
			Description: beeperdesktopapi.String("description"),
			Draft: beeperdesktopapi.ChatUpdateParamsDraft{
				Text: "text",
				Attachments: map[string]beeperdesktopapi.ChatUpdateParamsDraftAttachment{
					"foo": {
						UploadID: "uploadID",
						ID:       beeperdesktopapi.String("id"),
						Duration: beeperdesktopapi.Float(0),
						FileName: beeperdesktopapi.String("fileName"),
						MimeType: beeperdesktopapi.String("mimeType"),
						Size: beeperdesktopapi.ChatUpdateParamsDraftAttachmentSize{
							Height: 0,
							Width:  0,
						},
						Type: "image",
					},
				},
			},
			ImgURL:               beeperdesktopapi.String("imgURL"),
			IsArchived:           beeperdesktopapi.Bool(true),
			IsLowPriority:        beeperdesktopapi.Bool(true),
			IsMuted:              beeperdesktopapi.Bool(true),
			IsPinned:             beeperdesktopapi.Bool(true),
			MessageExpirySeconds: beeperdesktopapi.Int(0),
			Title:                beeperdesktopapi.String("title"),
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
		AccountIDs: []string{"matrix", "discordgo", "local-whatsapp_ba_EvYDBBsZbRQAy3UOSWqG0LuTVkc"},
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
	err := client.Chats.Archive(
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

func TestChatMarkReadWithOptionalParams(t *testing.T) {
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
	_, err := client.Chats.MarkRead(
		context.TODO(),
		"!NCdzlIaMjZUmvmvyHU:beeper.com",
		beeperdesktopapi.ChatMarkReadParams{
			MessageID: beeperdesktopapi.String("1343993"),
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

func TestChatMarkUnreadWithOptionalParams(t *testing.T) {
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
	_, err := client.Chats.MarkUnread(
		context.TODO(),
		"!NCdzlIaMjZUmvmvyHU:beeper.com",
		beeperdesktopapi.ChatMarkUnreadParams{
			MessageID: beeperdesktopapi.String("1343993"),
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

func TestChatNotifyAnyway(t *testing.T) {
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
	_, err := client.Chats.NotifyAnyway(
		context.TODO(),
		"!NCdzlIaMjZUmvmvyHU:beeper.com",
		beeperdesktopapi.ChatNotifyAnywayParams{},
	)
	if err != nil {
		var apierr *beeperdesktopapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestChatSearchWithOptionalParams(t *testing.T) {
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
	_, err := client.Chats.Search(context.TODO(), beeperdesktopapi.ChatSearchParams{
		AccountIDs:         []string{"matrix", "discordgo", "local-whatsapp_ba_EvYDBBsZbRQAy3UOSWqG0LuTVkc"},
		Cursor:             beeperdesktopapi.String("1725489123456|c29tZUltc2dQYWdl"),
		Direction:          beeperdesktopapi.ChatSearchParamsDirectionBefore,
		Inbox:              beeperdesktopapi.ChatSearchParamsInboxPrimary,
		IncludeMuted:       beeperdesktopapi.Bool(true),
		LastActivityAfter:  beeperdesktopapi.Time(time.Now()),
		LastActivityBefore: beeperdesktopapi.Time(time.Now()),
		Limit:              beeperdesktopapi.Int(1),
		Query:              beeperdesktopapi.String("x"),
		Scope:              beeperdesktopapi.ChatSearchParamsScopeTitles,
		Type:               beeperdesktopapi.ChatSearchParamsTypeSingle,
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

func TestChatStartWithOptionalParams(t *testing.T) {
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
	_, err := client.Chats.Start(context.TODO(), beeperdesktopapi.ChatStartParams{
		AccountID: "accountID",
		User: beeperdesktopapi.ChatStartParamsUser{
			ID:          beeperdesktopapi.String("id"),
			Email:       beeperdesktopapi.String("email"),
			FullName:    beeperdesktopapi.String("fullName"),
			PhoneNumber: beeperdesktopapi.String("phoneNumber"),
			Username:    beeperdesktopapi.String("username"),
		},
		AllowInvite: beeperdesktopapi.Bool(true),
		MessageText: beeperdesktopapi.String("messageText"),
	})
	if err != nil {
		var apierr *beeperdesktopapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
