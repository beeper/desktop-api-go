// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcombeeperdesktopapigo_test

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

func TestMessageGetAttachment(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := githubcombeeperdesktopapigo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAccessToken("My Access Token"),
	)
	_, err := client.Messages.GetAttachment(context.TODO(), githubcombeeperdesktopapigo.MessageGetAttachmentParams{
		ChatID:    "!-5hI_iHR5vSDCtI8PzSDQT0H_3I:ba_EvYDBBsZbRQAy3UOSWqG0LuTVkc.local-whatsapp.localhost",
		MessageID: "messageID",
	})
	if err != nil {
		var apierr *githubcombeeperdesktopapigo.Error
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
	client := githubcombeeperdesktopapigo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAccessToken("My Access Token"),
	)
	_, err := client.Messages.Search(context.TODO(), githubcombeeperdesktopapigo.MessageSearchParams{
		AccountIDs:         []string{"local-whatsapp_ba_EvYDBBsZbRQAy3UOSWqG0LuTVkc", "local-instagram_ba_eRfQMmnSNy_p7Ih7HL7RduRpKFU"},
		ChatIDs:            []string{"!-5hI_iHR5vSDCtI8PzSDQT0H_3I:ba_EvYDBBsZbRQAy3UOSWqG0LuTVkc.local-whatsapp.localhost"},
		ChatType:           githubcombeeperdesktopapigo.MessageSearchParamsChatTypeGroup,
		DateAfter:          githubcombeeperdesktopapigo.Time(time.Now()),
		DateBefore:         githubcombeeperdesktopapigo.Time(time.Now()),
		EndingBefore:       githubcombeeperdesktopapigo.String("872739"),
		ExcludeLowPriority: githubcombeeperdesktopapigo.Bool(true),
		IncludeMuted:       githubcombeeperdesktopapigo.Bool(true),
		Limit:              githubcombeeperdesktopapigo.Int(50),
		OnlyWithFile:       githubcombeeperdesktopapigo.Bool(true),
		OnlyWithImage:      githubcombeeperdesktopapigo.Bool(true),
		OnlyWithLink:       githubcombeeperdesktopapigo.Bool(true),
		OnlyWithMedia:      githubcombeeperdesktopapigo.Bool(true),
		OnlyWithVideo:      githubcombeeperdesktopapigo.Bool(true),
		Query:              githubcombeeperdesktopapigo.String("dinner"),
		Sender:             githubcombeeperdesktopapigo.MessageSearchParamsSenderMe,
		StartingAfter:      githubcombeeperdesktopapigo.String("196640"),
	})
	if err != nil {
		var apierr *githubcombeeperdesktopapigo.Error
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
	client := githubcombeeperdesktopapigo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAccessToken("My Access Token"),
	)
	_, err := client.Messages.Send(context.TODO(), githubcombeeperdesktopapigo.MessageSendParams{
		ChatID:           "!-5hI_iHR5vSDCtI8PzSDQT0H_3I:ba_EvYDBBsZbRQAy3UOSWqG0LuTVkc.local-whatsapp.localhost",
		ReplyToMessageID: githubcombeeperdesktopapigo.String("replyToMessageID"),
		Text:             githubcombeeperdesktopapigo.String("text"),
	})
	if err != nil {
		var apierr *githubcombeeperdesktopapigo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
