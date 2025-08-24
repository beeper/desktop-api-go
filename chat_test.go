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

func TestChatGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Chats.Get(context.TODO(), githubcombeeperbeeperdesktopapigo.ChatGetParams{
		ChatID:              "!-5hI_iHR5vSDCtI8PzSDQT0H_3I:ba_EvYDBBsZbRQAy3UOSWqG0LuTVkc.local-whatsapp.localhost",
		MaxParticipantCount: githubcombeeperbeeperdesktopapigo.Int(50),
	})
	if err != nil {
		var apierr *githubcombeeperbeeperdesktopapigo.Error
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
	client := githubcombeeperbeeperdesktopapigo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAccessToken("My Access Token"),
	)
	_, err := client.Chats.Archive(context.TODO(), githubcombeeperbeeperdesktopapigo.ChatArchiveParams{
		ChatID:   "!-5hI_iHR5vSDCtI8PzSDQT0H_3I:ba_EvYDBBsZbRQAy3UOSWqG0LuTVkc.local-whatsapp.localhost",
		Archived: githubcombeeperbeeperdesktopapigo.Bool(true),
	})
	if err != nil {
		var apierr *githubcombeeperbeeperdesktopapigo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestChatFindWithOptionalParams(t *testing.T) {
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
	_, err := client.Chats.Find(context.TODO(), githubcombeeperbeeperdesktopapigo.ChatFindParams{
		AccountIDs:         []string{"local-whatsapp_ba_EvYDBBsZbRQAy3UOSWqG0LuTVkc", "slackgo.T031TC83W"},
		EndingBefore:       githubcombeeperbeeperdesktopapigo.String("872739"),
		Inbox:              githubcombeeperbeeperdesktopapigo.ChatFindParamsInboxPrimary,
		IncludeMuted:       githubcombeeperbeeperdesktopapigo.Bool(true),
		LastActivityAfter:  githubcombeeperbeeperdesktopapigo.Time(time.Now()),
		LastActivityBefore: githubcombeeperbeeperdesktopapigo.Time(time.Now()),
		Limit:              githubcombeeperbeeperdesktopapigo.Int(1),
		ParticipantQuery:   githubcombeeperbeeperdesktopapigo.String("participantQuery"),
		Query:              githubcombeeperbeeperdesktopapigo.String("query"),
		StartingAfter:      githubcombeeperbeeperdesktopapigo.String("196640"),
		Type:               githubcombeeperbeeperdesktopapigo.ChatFindParamsTypeSingle,
		UnreadOnly:         githubcombeeperbeeperdesktopapigo.Bool(true),
	})
	if err != nil {
		var apierr *githubcombeeperbeeperdesktopapigo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestChatGetLinkWithOptionalParams(t *testing.T) {
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
	_, err := client.Chats.GetLink(context.TODO(), githubcombeeperbeeperdesktopapigo.ChatGetLinkParams{
		ChatID:         "!-5hI_iHR5vSDCtI8PzSDQT0H_3I:ba_EvYDBBsZbRQAy3UOSWqG0LuTVkc.local-whatsapp.localhost",
		MessageSortKey: githubcombeeperbeeperdesktopapigo.String("messageSortKey"),
	})
	if err != nil {
		var apierr *githubcombeeperbeeperdesktopapigo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
