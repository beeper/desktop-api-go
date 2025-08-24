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

func TestMessageDraftMessageWithOptionalParams(t *testing.T) {
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
	_, err := client.Messages.DraftMessage(context.TODO(), githubcombeeperdesktopapigo.MessageDraftMessageParams{
		ChatID:   "!-5hI_iHR5vSDCtI8PzSDQT0H_3I:ba_EvYDBBsZbRQAy3UOSWqG0LuTVkc.local-whatsapp.localhost",
		FocusApp: githubcombeeperdesktopapigo.Bool(true),
		Text:     githubcombeeperdesktopapigo.String("text"),
	})
	if err != nil {
		var apierr *githubcombeeperdesktopapigo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMessageSearchMessagesWithOptionalParams(t *testing.T) {
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
	_, err := client.Messages.SearchMessages(context.TODO(), githubcombeeperdesktopapigo.MessageSearchMessagesParams{
		AccountIDs:         []string{"local-whatsapp_ba_EvYDBBsZbRQAy3UOSWqG0LuTVkc", "local-instagram_ba_eRfQMmnSNy_p7Ih7HL7RduRpKFU"},
		ChatIDs:            []string{"!-5hI_iHR5vSDCtI8PzSDQT0H_3I:ba_EvYDBBsZbRQAy3UOSWqG0LuTVkc.local-whatsapp.localhost"},
		ChatType:           githubcombeeperdesktopapigo.MessageSearchMessagesParamsChatTypeGroup,
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
		Sender:             githubcombeeperdesktopapigo.MessageSearchMessagesParamsSenderMe,
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

func TestMessageSendMessageWithOptionalParams(t *testing.T) {
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
	_, err := client.Messages.SendMessage(context.TODO(), githubcombeeperdesktopapigo.MessageSendMessageParams{
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
