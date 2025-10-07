// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package beeperdesktopapi

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/beeper-desktop-api-go/internal/apijson"
	"github.com/stainless-sdks/beeper-desktop-api-go/internal/requestconfig"
	"github.com/stainless-sdks/beeper-desktop-api-go/option"
	"github.com/stainless-sdks/beeper-desktop-api-go/packages/respjson"
)

// Operations related to the current access token
//
// TokenService contains methods and other services that help with interacting with
// the beeperdesktop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTokenService] method instead.
type TokenService struct {
	Options []option.RequestOption
}

// NewTokenService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTokenService(opts ...option.RequestOption) (r TokenService) {
	r = TokenService{}
	r.Options = opts
	return
}

// Returns information about the authenticated user/token
func (r *TokenService) Info(ctx context.Context, opts ...option.RequestOption) (res *UserInfo, err error) {
	opts = append(r.Options[:], opts...)
	path := "oauth/userinfo"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
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
