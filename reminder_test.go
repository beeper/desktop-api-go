// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcombeeperbeeperdesktopapigo_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/beeper/beeper-desktop-api-go"
	"github.com/beeper/beeper-desktop-api-go/internal/testutil"
	"github.com/beeper/beeper-desktop-api-go/option"
)

func TestReminderClear(t *testing.T) {
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
	_, err := client.Reminders.Clear(context.TODO(), githubcombeeperbeeperdesktopapigo.ReminderClearParams{
		ChatID: "!-5hI_iHR5vSDCtI8PzSDQT0H_3I:ba_EvYDBBsZbRQAy3UOSWqG0LuTVkc.local-whatsapp.localhost",
	})
	if err != nil {
		var apierr *githubcombeeperbeeperdesktopapigo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestReminderSetWithOptionalParams(t *testing.T) {
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
	_, err := client.Reminders.Set(context.TODO(), githubcombeeperbeeperdesktopapigo.ReminderSetParams{
		ChatID: "!-5hI_iHR5vSDCtI8PzSDQT0H_3I:ba_EvYDBBsZbRQAy3UOSWqG0LuTVkc.local-whatsapp.localhost",
		Reminder: githubcombeeperbeeperdesktopapigo.ReminderSetParamsReminder{
			RemindAtMs:               0,
			DismissOnIncomingMessage: githubcombeeperbeeperdesktopapigo.Bool(true),
		},
	})
	if err != nil {
		var apierr *githubcombeeperbeeperdesktopapigo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
