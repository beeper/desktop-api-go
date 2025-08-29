// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcombeeperdesktopapigo

import (
	"context"
	"net/http"
	"net/url"

	"github.com/beeper/desktop-api-go/internal/apijson"
	"github.com/beeper/desktop-api-go/internal/apiquery"
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

// Starts the OAuth2 Authorization Code flow with PKCE. Renders an HTML consent
// page where the user can approve the requested scopes.
func (r *OAuthService) Authorize(ctx context.Context, query OAuthAuthorizeParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/html")}, opts...)
	path := "oauth/authorize"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Completes the OAuth2 approval initiated by the consent screen and returns an
// authorization code if approved.
func (r *OAuthService) AuthorizeCallback(ctx context.Context, body OAuthAuthorizeCallbackParams, opts ...option.RequestOption) (res *OAuthAuthorizeCallbackResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "oauth/authorize/callback"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns information about the authenticated user/token
func (r *OAuthService) GetUserInfo(ctx context.Context, opts ...option.RequestOption) (res *UserInfo, err error) {
	opts = append(r.Options[:], opts...)
	path := "oauth/userinfo"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Registers a new OAuth2 public client using Dynamic Client Registration (RFC
// 7591). Returns client metadata.
func (r *OAuthService) RegisterClient(ctx context.Context, body OAuthRegisterClientParams, opts ...option.RequestOption) (res *OAuthRegisterClientResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "oauth/register"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
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

// Exchanges an authorization code (PKCE) for an access token. Supports the
// Authorization Code grant with PKCE.
func (r *OAuthService) Token(ctx context.Context, body OAuthTokenParams, opts ...option.RequestOption) (res *OAuthTokenResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "oauth/token"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// RFC 8414 authorization server metadata document.
func (r *OAuthService) WellKnownAuthorizationServer(ctx context.Context, opts ...option.RequestOption) (res *OAuthWellKnownAuthorizationServerResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := ".well-known/oauth-authorization-server"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// RFC 9728 protected resource metadata document.
func (r *OAuthService) WellKnownProtectedResource(ctx context.Context, opts ...option.RequestOption) (res *OAuthWellKnownProtectedResourceResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := ".well-known/oauth-protected-resource"
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

type OAuthAuthorizeCallbackResponse struct {
	// Authorization code to exchange for a token
	Code string `json:"code"`
	// Error code if authorization was denied or failed
	Error string `json:"error"`
	State string `json:"state"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Error       respjson.Field
		State       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OAuthAuthorizeCallbackResponse) RawJSON() string { return r.JSON.raw }
func (r *OAuthAuthorizeCallbackResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OAuthRegisterClientResponse struct {
	AuthorizationEndpoint string `json:"authorization_endpoint,required" format:"uri"`
	ClientID              string `json:"client_id,required"`
	ClientIDIssuedAt      int64  `json:"client_id_issued_at,required"`
	ClientName            string `json:"client_name,required"`
	// Any of "authorization_code".
	GrantTypes   []string `json:"grant_types,required"`
	RedirectUris []string `json:"redirect_uris,required" format:"uri"`
	// Any of "code".
	ResponseTypes []string `json:"response_types,required"`
	Scope         string   `json:"scope,required"`
	TokenEndpoint string   `json:"token_endpoint,required" format:"uri"`
	// Any of "none".
	TokenEndpointAuthMethod OAuthRegisterClientResponseTokenEndpointAuthMethod `json:"token_endpoint_auth_method,required"`
	ClientUri               string                                             `json:"client_uri" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AuthorizationEndpoint   respjson.Field
		ClientID                respjson.Field
		ClientIDIssuedAt        respjson.Field
		ClientName              respjson.Field
		GrantTypes              respjson.Field
		RedirectUris            respjson.Field
		ResponseTypes           respjson.Field
		Scope                   respjson.Field
		TokenEndpoint           respjson.Field
		TokenEndpointAuthMethod respjson.Field
		ClientUri               respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OAuthRegisterClientResponse) RawJSON() string { return r.JSON.raw }
func (r *OAuthRegisterClientResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OAuthRegisterClientResponseTokenEndpointAuthMethod string

const (
	OAuthRegisterClientResponseTokenEndpointAuthMethodNone OAuthRegisterClientResponseTokenEndpointAuthMethod = "none"
)

type OAuthTokenResponse struct {
	AccessToken string `json:"access_token,required"`
	ExpiresIn   int64  `json:"expires_in,required"`
	Scope       string `json:"scope,required"`
	// Any of "Bearer".
	TokenType OAuthTokenResponseTokenType `json:"token_type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccessToken respjson.Field
		ExpiresIn   respjson.Field
		Scope       respjson.Field
		TokenType   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OAuthTokenResponse) RawJSON() string { return r.JSON.raw }
func (r *OAuthTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OAuthTokenResponseTokenType string

const (
	OAuthTokenResponseTokenTypeBearer OAuthTokenResponseTokenType = "Bearer"
)

type OAuthWellKnownAuthorizationServerResponse struct {
	AuthorizationEndpoint string `json:"authorization_endpoint,required" format:"uri"`
	// Any of "S256".
	CodeChallengeMethodsSupported []string `json:"code_challenge_methods_supported,required"`
	// Any of "authorization_code".
	GrantTypesSupported  []string `json:"grant_types_supported,required"`
	Issuer               string   `json:"issuer,required" format:"uri"`
	RegistrationEndpoint string   `json:"registration_endpoint,required" format:"uri"`
	// Any of "code".
	ResponseTypesSupported []string `json:"response_types_supported,required"`
	RevocationEndpoint     string   `json:"revocation_endpoint,required" format:"uri"`
	// Any of "read", "write".
	ScopesSupported      []string `json:"scopes_supported,required"`
	ServiceDocumentation string   `json:"service_documentation,required" format:"uri"`
	TokenEndpoint        string   `json:"token_endpoint,required" format:"uri"`
	// Any of "none".
	TokenEndpointAuthMethodsSupported []string `json:"token_endpoint_auth_methods_supported,required"`
	UserinfoEndpoint                  string   `json:"userinfo_endpoint,required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AuthorizationEndpoint             respjson.Field
		CodeChallengeMethodsSupported     respjson.Field
		GrantTypesSupported               respjson.Field
		Issuer                            respjson.Field
		RegistrationEndpoint              respjson.Field
		ResponseTypesSupported            respjson.Field
		RevocationEndpoint                respjson.Field
		ScopesSupported                   respjson.Field
		ServiceDocumentation              respjson.Field
		TokenEndpoint                     respjson.Field
		TokenEndpointAuthMethodsSupported respjson.Field
		UserinfoEndpoint                  respjson.Field
		ExtraFields                       map[string]respjson.Field
		raw                               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OAuthWellKnownAuthorizationServerResponse) RawJSON() string { return r.JSON.raw }
func (r *OAuthWellKnownAuthorizationServerResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OAuthWellKnownProtectedResourceResponse struct {
	AuthorizationServers []string `json:"authorization_servers,required" format:"uri"`
	Resource             string   `json:"resource,required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AuthorizationServers respjson.Field
		Resource             respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OAuthWellKnownProtectedResourceResponse) RawJSON() string { return r.JSON.raw }
func (r *OAuthWellKnownProtectedResourceResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OAuthAuthorizeParams struct {
	// Client identifier
	ClientID string `query:"client_id,required" json:"-"`
	// PKCE code challenge (S256)
	CodeChallenge string `query:"code_challenge,required" json:"-"`
	// Redirect URI to receive the authorization code
	RedirectUri string `query:"redirect_uri,required" format:"uri" json:"-"`
	// Must be "code"
	//
	// Any of "code".
	ResponseType OAuthAuthorizeParamsResponseType `query:"response_type,omitzero,required" json:"-"`
	// Requested resource (RFC 8707)
	Resource param.Opt[string] `query:"resource,omitzero" json:"-"`
	// Space-delimited scopes to request
	Scope param.Opt[string] `query:"scope,omitzero" json:"-"`
	// Opaque value to maintain state between the request and callback
	State param.Opt[string] `query:"state,omitzero" json:"-"`
	// PKCE method; only S256 is supported
	//
	// Any of "S256".
	CodeChallengeMethod OAuthAuthorizeParamsCodeChallengeMethod `query:"code_challenge_method,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OAuthAuthorizeParams]'s query parameters as `url.Values`.
func (r OAuthAuthorizeParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Must be "code"
type OAuthAuthorizeParamsResponseType string

const (
	OAuthAuthorizeParamsResponseTypeCode OAuthAuthorizeParamsResponseType = "code"
)

// PKCE method; only S256 is supported
type OAuthAuthorizeParamsCodeChallengeMethod string

const (
	OAuthAuthorizeParamsCodeChallengeMethodS256 OAuthAuthorizeParamsCodeChallengeMethod = "S256"
)

type OAuthAuthorizeCallbackParams struct {
	// PKCE code challenge from the authorize request
	CodeChallenge param.Opt[string] `json:"codeChallenge,omitzero"`
	// Requested resource (RFC 8707)
	Resource   param.Opt[string]                      `json:"resource,omitzero"`
	State      param.Opt[string]                      `json:"state,omitzero"`
	ClientInfo OAuthAuthorizeCallbackParamsClientInfo `json:"clientInfo,omitzero"`
	// PKCE method; only S256 is supported
	//
	// Any of "S256".
	CodeChallengeMethod OAuthAuthorizeCallbackParamsCodeChallengeMethod `json:"codeChallengeMethod,omitzero"`
	// Requested scopes
	//
	// Any of "read", "write".
	Scopes []string `json:"scopes,omitzero"`
	paramObj
}

func (r OAuthAuthorizeCallbackParams) MarshalJSON() (data []byte, err error) {
	type shadow OAuthAuthorizeCallbackParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OAuthAuthorizeCallbackParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OAuthAuthorizeCallbackParamsClientInfo struct {
	ClientID      param.Opt[string] `json:"clientID,omitzero"`
	ClientUri     param.Opt[string] `json:"clientURI,omitzero"`
	Name          param.Opt[string] `json:"name,omitzero"`
	RemoteAddress param.Opt[string] `json:"remoteAddress,omitzero"`
	UserAgent     param.Opt[string] `json:"userAgent,omitzero"`
	Version       param.Opt[string] `json:"version,omitzero"`
	paramObj
}

func (r OAuthAuthorizeCallbackParamsClientInfo) MarshalJSON() (data []byte, err error) {
	type shadow OAuthAuthorizeCallbackParamsClientInfo
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OAuthAuthorizeCallbackParamsClientInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PKCE method; only S256 is supported
type OAuthAuthorizeCallbackParamsCodeChallengeMethod string

const (
	OAuthAuthorizeCallbackParamsCodeChallengeMethodS256 OAuthAuthorizeCallbackParamsCodeChallengeMethod = "S256"
)

type OAuthRegisterClientParams struct {
	// Human-readable client name
	ClientName string `json:"client_name,required"`
	// Allowed redirect URIs
	RedirectUris []string `json:"redirect_uris,omitzero,required" format:"uri"`
	// Client homepage
	ClientUri param.Opt[string] `json:"client_uri,omitzero" format:"uri"`
	// Default scopes
	Scope param.Opt[string] `json:"scope,omitzero"`
	// Supported grant types
	//
	// Any of "authorization_code".
	GrantTypes []string `json:"grant_types,omitzero"`
	// Supported response types
	//
	// Any of "code".
	ResponseTypes []string `json:"response_types,omitzero"`
	// Auth method for token endpoint; public clients use none
	//
	// Any of "none".
	TokenEndpointAuthMethod OAuthRegisterClientParamsTokenEndpointAuthMethod `json:"token_endpoint_auth_method,omitzero"`
	paramObj
}

func (r OAuthRegisterClientParams) MarshalJSON() (data []byte, err error) {
	type shadow OAuthRegisterClientParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OAuthRegisterClientParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Auth method for token endpoint; public clients use none
type OAuthRegisterClientParamsTokenEndpointAuthMethod string

const (
	OAuthRegisterClientParamsTokenEndpointAuthMethodNone OAuthRegisterClientParamsTokenEndpointAuthMethod = "none"
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

type OAuthTokenParams struct {
	// Authorization code returned by the authorize step
	Code string `json:"code,required"`
	// PKCE code verifier
	CodeVerifier string `json:"code_verifier,required"`
	// OAuth2 grant type; only authorization_code is supported
	//
	// Any of "authorization_code".
	GrantType OAuthTokenParamsGrantType `json:"grant_type,omitzero,required"`
	// Client ID (optional for public PKCE clients)
	ClientID param.Opt[string] `json:"client_id,omitzero"`
	// Resource parameter (RFC 8707)
	Resource param.Opt[string] `json:"resource,omitzero"`
	paramObj
}

func (r OAuthTokenParams) MarshalJSON() (data []byte, err error) {
	type shadow OAuthTokenParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OAuthTokenParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OAuth2 grant type; only authorization_code is supported
type OAuthTokenParamsGrantType string

const (
	OAuthTokenParamsGrantTypeAuthorizationCode OAuthTokenParamsGrantType = "authorization_code"
)
