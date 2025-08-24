// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcombeeperbeeperdesktopapigo

import (
	"context"
	"net/http"

	"github.com/beeper/desktop-api-go/internal/apijson"
	"github.com/beeper/desktop-api-go/internal/requestconfig"
	"github.com/beeper/desktop-api-go/option"
	"github.com/beeper/desktop-api-go/packages/param"
	"github.com/beeper/desktop-api-go/packages/respjson"
)

// OAuth2 authentication and token management
//
// OAuthService contains methods and other services that help with interacting with
// the beeperdesktop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOAuthService] method instead.
type OAuthService struct {
	Options []option.RequestOption
}

// NewOAuthService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewOAuthService(opts ...option.RequestOption) (r OAuthService) {
	r = OAuthService{}
	r.Options = opts
	return
}

// Returns information about the authenticated user/token
func (r *OAuthService) GetUserInfo(ctx context.Context, opts ...option.RequestOption) (res *UserInfo, err error) {
	opts = append(r.Options[:], opts...)
	path := "oauth/userinfo"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Revoke an access token or refresh token (RFC 7009)
func (r *OAuthService) RevokeToken(ctx context.Context, body OAuthRevokeTokenParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "oauth/revoke"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type UserInfo struct {
	// Issued at timestamp (Unix epoch seconds)
	Iat float64 `json:"iat,required"`
	// Granted scopes
	Scope string `json:"scope,required"`
	// Subject identifier (token ID)
	Sub string `json:"sub,required"`
	// Token type
	//
	// Any of "access".
	TokenUse UserInfoTokenUse `json:"token_use,required"`
	// Audience (client ID)
	Aud string `json:"aud"`
	// Client identifier
	ClientID string `json:"client_id"`
	// Expiration timestamp (Unix epoch seconds)
	Exp float64 `json:"exp"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Iat         respjson.Field
		Scope       respjson.Field
		Sub         respjson.Field
		TokenUse    respjson.Field
		Aud         respjson.Field
		ClientID    respjson.Field
		Exp         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserInfo) RawJSON() string { return r.JSON.raw }
func (r *UserInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Token type
type UserInfoTokenUse string

const (
	UserInfoTokenUseAccess UserInfoTokenUse = "access"
)

type OAuthRevokeTokenParams struct {
	// The token to revoke
	Token string `json:"token,required"`
	// Token type hint (RFC 7009)
	//
	// Any of "access_token".
	TokenTypeHint OAuthRevokeTokenParamsTokenTypeHint `json:"token_type_hint,omitzero"`
	paramObj
}

func (r OAuthRevokeTokenParams) MarshalJSON() (data []byte, err error) {
	type shadow OAuthRevokeTokenParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OAuthRevokeTokenParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Token type hint (RFC 7009)
type OAuthRevokeTokenParamsTokenTypeHint string

const (
	OAuthRevokeTokenParamsTokenTypeHintAccessToken OAuthRevokeTokenParamsTokenTypeHint = "access_token"
)
