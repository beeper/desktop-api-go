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

func TestChatArchiveChatWithOptionalParams(t *testing.T) {
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
	_, err := client.Chats.ArchiveChat(context.TODO(), githubcombeeperdesktopapigo.ChatArchiveChatParams{
		ChatID:   "!-5hI_iHR5vSDCtI8PzSDQT0H_3I:ba_EvYDBBsZbRQAy3UOSWqG0LuTVkc.local-whatsapp.localhost",
		Archived: githubcombeeperdesktopapigo.Bool(true),
	})
	if err != nil {
		var apierr *githubcombeeperdesktopapigo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestChatFindChatsWithOptionalParams(t *testing.T) {
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
	_, err := client.Chats.FindChats(context.TODO(), githubcombeeperdesktopapigo.ChatFindChatsParams{
		AccountIDs:         []string{"local-whatsapp_ba_EvYDBBsZbRQAy3UOSWqG0LuTVkc", "slackgo.T031TC83W"},
		EndingBefore:       githubcombeeperdesktopapigo.String("872739"),
		Inbox:              githubcombeeperdesktopapigo.ChatFindChatsParamsInboxPrimary,
		IncludeMuted:       githubcombeeperdesktopapigo.Bool(true),
		LastActivityAfter:  githubcombeeperdesktopapigo.Time(time.Now()),
		LastActivityBefore: githubcombeeperdesktopapigo.Time(time.Now()),
		Limit:              githubcombeeperdesktopapigo.Int(1),
		ParticipantQuery:   githubcombeeperdesktopapigo.String("participantQuery"),
		Query:              githubcombeeperdesktopapigo.String("query"),
		StartingAfter:      githubcombeeperdesktopapigo.String("196640"),
		Type:               githubcombeeperdesktopapigo.ChatFindChatsParamsTypeSingle,
		UnreadOnly:         githubcombeeperdesktopapigo.Bool(true),
	})
	if err != nil {
		var apierr *githubcombeeperdesktopapigo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestChatGetChatWithOptionalParams(t *testing.T) {
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
	_, err := client.Chats.GetChat(context.TODO(), githubcombeeperdesktopapigo.ChatGetChatParams{
		ChatID:              "!-5hI_iHR5vSDCtI8PzSDQT0H_3I:ba_EvYDBBsZbRQAy3UOSWqG0LuTVkc.local-whatsapp.localhost",
		MaxParticipantCount: githubcombeeperdesktopapigo.Int(50),
	})
	if err != nil {
		var apierr *githubcombeeperdesktopapigo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestChatGetLinkToChatWithOptionalParams(t *testing.T) {
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
	_, err := client.Chats.GetLinkToChat(context.TODO(), githubcombeeperdesktopapigo.ChatGetLinkToChatParams{
		ChatID:         "!-5hI_iHR5vSDCtI8PzSDQT0H_3I:ba_EvYDBBsZbRQAy3UOSWqG0LuTVkc.local-whatsapp.localhost",
		MessageSortKey: githubcombeeperdesktopapigo.String("messageSortKey"),
	})
	if err != nil {
		var apierr *githubcombeeperdesktopapigo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
