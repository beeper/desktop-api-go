// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package beeperdesktopapi_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/beeper-desktop-api-go"
	"github.com/stainless-sdks/beeper-desktop-api-go/internal/testutil"
	"github.com/stainless-sdks/beeper-desktop-api-go/option"
)

func TestChatReminderNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Chats.Reminders.New(context.TODO(), beeperdesktopapi.ChatReminderNewParams{
		ChatID: "!NCdzlIaMjZUmvmvyHU:beeper.com",
		Reminder: beeperdesktopapi.ChatReminderNewParamsReminder{
			RemindAtMs:               0,
			DismissOnIncomingMessage: beeperdesktopapi.Bool(true),
		},
	})
	if err != nil {
		var apierr *beeperdesktopapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestChatReminderDelete(t *testing.T) {
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
	_, err := client.Chats.Reminders.Delete(context.TODO(), beeperdesktopapi.ChatReminderDeleteParams{
		ChatID: "!NCdzlIaMjZUmvmvyHU:beeper.com",
	})
	if err != nil {
		var apierr *beeperdesktopapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
