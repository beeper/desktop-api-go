// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcombeeperdesktopapigo_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/beeper/desktop-api-go"
	"github.com/beeper/desktop-api-go/internal/testutil"
	"github.com/beeper/desktop-api-go/option"
)

func TestReminderClearChatReminder(t *testing.T) {
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
	_, err := client.Reminders.ClearChatReminder(context.TODO(), githubcombeeperdesktopapigo.ReminderClearChatReminderParams{
		ChatID: "!-5hI_iHR5vSDCtI8PzSDQT0H_3I:ba_EvYDBBsZbRQAy3UOSWqG0LuTVkc.local-whatsapp.localhost",
	})
	if err != nil {
		var apierr *githubcombeeperdesktopapigo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestReminderSetChatReminderWithOptionalParams(t *testing.T) {
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
	_, err := client.Reminders.SetChatReminder(context.TODO(), githubcombeeperdesktopapigo.ReminderSetChatReminderParams{
		ChatID: "!-5hI_iHR5vSDCtI8PzSDQT0H_3I:ba_EvYDBBsZbRQAy3UOSWqG0LuTVkc.local-whatsapp.localhost",
		Reminder: githubcombeeperdesktopapigo.ReminderSetChatReminderParamsReminder{
			RemindAtMs:               0,
			DismissOnIncomingMessage: githubcombeeperdesktopapigo.Bool(true),
		},
	})
	if err != nil {
		var apierr *githubcombeeperdesktopapigo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
