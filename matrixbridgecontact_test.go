// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package beeperdesktopapi_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/beeper/desktop-api-go/v5"
	"github.com/beeper/desktop-api-go/v5/internal/testutil"
	"github.com/beeper/desktop-api-go/v5/option"
)

func TestMatrixBridgeContactListWithOptionalParams(t *testing.T) {
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
	_, err := client.Matrix.Bridges.Contacts.List(
		context.TODO(),
		"bridgeID",
		beeperdesktopapi.MatrixBridgeContactListParams{
			LoginID: beeperdesktopapi.String("bcc68892-b180-414f-9516-b4aadf7d0496"),
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
