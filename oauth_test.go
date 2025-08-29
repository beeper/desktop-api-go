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

func TestOAuthAuthorizeWithOptionalParams(t *testing.T) {
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
	_, err := client.OAuth.Authorize(context.TODO(), githubcombeeperdesktopapigo.OAuthAuthorizeParams{
		ClientID:            "client_id",
		CodeChallenge:       "code_challenge",
		RedirectUri:         "https://example.com",
		ResponseType:        githubcombeeperdesktopapigo.OAuthAuthorizeParamsResponseTypeCode,
		CodeChallengeMethod: githubcombeeperdesktopapigo.OAuthAuthorizeParamsCodeChallengeMethodS256,
		Resource:            githubcombeeperdesktopapigo.String("resource"),
		Scope:               githubcombeeperdesktopapigo.String("scope"),
		State:               githubcombeeperdesktopapigo.String("state"),
	})
	if err != nil {
		var apierr *githubcombeeperdesktopapigo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOAuthAuthorizeCallbackWithOptionalParams(t *testing.T) {
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
	_, err := client.OAuth.AuthorizeCallback(context.TODO(), githubcombeeperdesktopapigo.OAuthAuthorizeCallbackParams{
		ClientInfo: githubcombeeperdesktopapigo.OAuthAuthorizeCallbackParamsClientInfo{
			ClientID:      githubcombeeperdesktopapigo.String("clientID"),
			ClientUri:     githubcombeeperdesktopapigo.String("clientURI"),
			Name:          githubcombeeperdesktopapigo.String("name"),
			RemoteAddress: githubcombeeperdesktopapigo.String("remoteAddress"),
			UserAgent:     githubcombeeperdesktopapigo.String("userAgent"),
			Version:       githubcombeeperdesktopapigo.String("version"),
		},
		CodeChallenge:       githubcombeeperdesktopapigo.String("codeChallenge"),
		CodeChallengeMethod: githubcombeeperdesktopapigo.OAuthAuthorizeCallbackParamsCodeChallengeMethodS256,
		Resource:            githubcombeeperdesktopapigo.String("resource"),
		Scopes:              []string{"read"},
		State:               githubcombeeperdesktopapigo.String("state"),
	})
	if err != nil {
		var apierr *githubcombeeperdesktopapigo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOAuthGetUserInfo(t *testing.T) {
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
	_, err := client.OAuth.GetUserInfo(context.TODO())
	if err != nil {
		var apierr *githubcombeeperdesktopapigo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOAuthRegisterClientWithOptionalParams(t *testing.T) {
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
	_, err := client.OAuth.RegisterClient(context.TODO(), githubcombeeperdesktopapigo.OAuthRegisterClientParams{
		ClientName:              "client_name",
		RedirectUris:            []string{"https://example.com"},
		ClientUri:               githubcombeeperdesktopapigo.String("https://example.com"),
		GrantTypes:              []string{"authorization_code"},
		ResponseTypes:           []string{"code"},
		Scope:                   githubcombeeperdesktopapigo.String("scope"),
		TokenEndpointAuthMethod: githubcombeeperdesktopapigo.OAuthRegisterClientParamsTokenEndpointAuthMethodNone,
	})
	if err != nil {
		var apierr *githubcombeeperdesktopapigo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOAuthRevokeTokenWithOptionalParams(t *testing.T) {
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
	err := client.OAuth.RevokeToken(context.TODO(), githubcombeeperdesktopapigo.OAuthRevokeTokenParams{
		Token:         "token",
		TokenTypeHint: githubcombeeperdesktopapigo.OAuthRevokeTokenParamsTokenTypeHintAccessToken,
	})
	if err != nil {
		var apierr *githubcombeeperdesktopapigo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOAuthTokenWithOptionalParams(t *testing.T) {
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
	_, err := client.OAuth.Token(context.TODO(), githubcombeeperdesktopapigo.OAuthTokenParams{
		Code:         "code",
		CodeVerifier: "code_verifier",
		GrantType:    githubcombeeperdesktopapigo.OAuthTokenParamsGrantTypeAuthorizationCode,
		ClientID:     githubcombeeperdesktopapigo.String("client_id"),
		Resource:     githubcombeeperdesktopapigo.String("resource"),
	})
	if err != nil {
		var apierr *githubcombeeperdesktopapigo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOAuthWellKnownAuthorizationServer(t *testing.T) {
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
	_, err := client.OAuth.WellKnownAuthorizationServer(context.TODO())
	if err != nil {
		var apierr *githubcombeeperdesktopapigo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOAuthWellKnownProtectedResource(t *testing.T) {
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
	_, err := client.OAuth.WellKnownProtectedResource(context.TODO())
	if err != nil {
		var apierr *githubcombeeperdesktopapigo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
