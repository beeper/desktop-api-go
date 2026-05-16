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

func TestMatrixRoomNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Matrix.Rooms.New(context.TODO(), beeperdesktopapi.MatrixRoomNewParams{
		CreationContent: map[string]any{
			"m.federate": false,
		},
		InitialState: []beeperdesktopapi.MatrixRoomNewParamsInitialState{{
			Content:  map[string]any{},
			Type:     "type",
			StateKey: beeperdesktopapi.String("state_key"),
		}},
		Invite: []string{"string"},
		Invite3pid: []beeperdesktopapi.MatrixRoomNewParamsInvite3pid{{
			Address:       "cheeky@monkey.com",
			IDAccessToken: "abc123_OpaqueString",
			IDServer:      "matrix.org",
			Medium:        "email",
		}},
		IsDirect:                  beeperdesktopapi.Bool(true),
		Name:                      beeperdesktopapi.String("The Grand Duke Pub"),
		PowerLevelContentOverride: map[string]any{},
		Preset:                    beeperdesktopapi.MatrixRoomNewParamsPresetPublicChat,
		RoomAliasName:             beeperdesktopapi.String("thepub"),
		RoomVersion:               beeperdesktopapi.String("1"),
		Topic:                     beeperdesktopapi.String("All about happy hour"),
		Visibility:                beeperdesktopapi.MatrixRoomNewParamsVisibilityPublic,
	})
	if err != nil {
		var apierr *beeperdesktopapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMatrixRoomJoinWithOptionalParams(t *testing.T) {
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
	_, err := client.Matrix.Rooms.Join(
		context.TODO(),
		"!monkeys:matrix.org",
		beeperdesktopapi.MatrixRoomJoinParams{
			Via:    []string{"string"},
			Reason: beeperdesktopapi.String("Looking for support"),
			ThirdPartySigned: beeperdesktopapi.MatrixRoomJoinParamsThirdPartySigned{
				Token:  "random8nonce",
				Mxid:   "bob",
				Sender: "alice",
				Signatures: map[string]map[string]string{
					"example.org": {
						"ed25519:0": "some9signature",
					},
				},
			},
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

func TestMatrixRoomLeaveWithOptionalParams(t *testing.T) {
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
	_, err := client.Matrix.Rooms.Leave(
		context.TODO(),
		"!nkl290a:matrix.org",
		beeperdesktopapi.MatrixRoomLeaveParams{
			Reason: beeperdesktopapi.String("Saying farewell - thanks for the support!"),
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
