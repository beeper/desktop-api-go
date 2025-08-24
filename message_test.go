// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcombeeperbeeperdesktopapigo_test

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

func TestMessageDraftWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := githubcombeeperbeeperdesktopapigo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAccessToken("My Access Token"),
	)
	_, err := client.Messages.Draft(context.TODO(), githubcombeeperbeeperdesktopapigo.MessageDraftParams{
		ChatID:   "!-5hI_iHR5vSDCtI8PzSDQT0H_3I:ba_EvYDBBsZbRQAy3UOSWqG0LuTVkc.local-whatsapp.localhost",
		FocusApp: githubcombeeperbeeperdesktopapigo.Bool(true),
		Text:     githubcombeeperbeeperdesktopapigo.String("text"),
	})
	if err != nil {
		var apierr *githubcombeeperbeeperdesktopapigo.Error
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
	client := githubcombeeperbeeperdesktopapigo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAccessToken("My Access Token"),
	)
	_, err := client.Messages.Search(context.TODO(), githubcombeeperbeeperdesktopapigo.MessageSearchParams{
		AccountIDs:         []string{"local-whatsapp_ba_EvYDBBsZbRQAy3UOSWqG0LuTVkc", "local-instagram_ba_eRfQMmnSNy_p7Ih7HL7RduRpKFU"},
		ChatIDs:            []string{"!-5hI_iHR5vSDCtI8PzSDQT0H_3I:ba_EvYDBBsZbRQAy3UOSWqG0LuTVkc.local-whatsapp.localhost"},
		ChatType:           githubcombeeperbeeperdesktopapigo.MessageSearchParamsChatTypeGroup,
		DateAfter:          githubcombeeperbeeperdesktopapigo.Time(time.Now()),
		DateBefore:         githubcombeeperbeeperdesktopapigo.Time(time.Now()),
		EndingBefore:       githubcombeeperbeeperdesktopapigo.String("872739"),
		ExcludeLowPriority: githubcombeeperbeeperdesktopapigo.Bool(true),
		IncludeMuted:       githubcombeeperbeeperdesktopapigo.Bool(true),
		Limit:              githubcombeeperbeeperdesktopapigo.Int(50),
		OnlyWithFile:       githubcombeeperbeeperdesktopapigo.Bool(true),
		OnlyWithImage:      githubcombeeperbeeperdesktopapigo.Bool(true),
		OnlyWithLink:       githubcombeeperbeeperdesktopapigo.Bool(true),
		OnlyWithMedia:      githubcombeeperbeeperdesktopapigo.Bool(true),
		OnlyWithVideo:      githubcombeeperbeeperdesktopapigo.Bool(true),
		Query:              githubcombeeperbeeperdesktopapigo.String("dinner"),
		Sender:             githubcombeeperbeeperdesktopapigo.MessageSearchParamsSenderMe,
		StartingAfter:      githubcombeeperbeeperdesktopapigo.String("196640"),
	})
	if err != nil {
		var apierr *githubcombeeperbeeperdesktopapigo.Error
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
	client := githubcombeeperbeeperdesktopapigo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAccessToken("My Access Token"),
	)
	_, err := client.Messages.Send(context.TODO(), githubcombeeperbeeperdesktopapigo.MessageSendParams{
		ChatID:           "!-5hI_iHR5vSDCtI8PzSDQT0H_3I:ba_EvYDBBsZbRQAy3UOSWqG0LuTVkc.local-whatsapp.localhost",
		ReplyToMessageID: githubcombeeperbeeperdesktopapigo.String("replyToMessageID"),
		Text:             githubcombeeperbeeperdesktopapigo.String("text"),
	})
	if err != nil {
		var apierr *githubcombeeperbeeperdesktopapigo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
