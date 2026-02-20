// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package beeperdesktopapi

import (
	"context"
	"net/http"
	"slices"

	"github.com/beeper/desktop-api-go/internal/apijson"
	"github.com/beeper/desktop-api-go/internal/requestconfig"
	"github.com/beeper/desktop-api-go/option"
	"github.com/beeper/desktop-api-go/packages/respjson"
)

// InfoService contains methods and other services that help with interacting with
// the beeperdesktop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInfoService] method instead.
type InfoService struct {
	Options []option.RequestOption
}

// NewInfoService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewInfoService(opts ...option.RequestOption) (r InfoService) {
	r = InfoService{}
	r.Options = opts
	return
}

// Returns app, platform, server, and endpoint discovery metadata for this Beeper
// Desktop instance.
func (r *InfoService) Get(ctx context.Context, opts ...option.RequestOption) (res *InfoGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/info"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type InfoGetResponse struct {
	App       InfoGetResponseApp       `json:"app,required"`
	Endpoints InfoGetResponseEndpoints `json:"endpoints,required"`
	Platform  InfoGetResponsePlatform  `json:"platform,required"`
	Server    InfoGetResponseServer    `json:"server,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		App         respjson.Field
		Endpoints   respjson.Field
		Platform    respjson.Field
		Server      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InfoGetResponse) RawJSON() string { return r.JSON.raw }
func (r *InfoGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InfoGetResponseApp struct {
	// App bundle identifier
	BundleID string `json:"bundle_id,required"`
	// App name
	Name string `json:"name,required"`
	// App version
	Version string `json:"version,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BundleID    respjson.Field
		Name        respjson.Field
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InfoGetResponseApp) RawJSON() string { return r.JSON.raw }
func (r *InfoGetResponseApp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InfoGetResponseEndpoints struct {
	// MCP endpoint
	Mcp   string                        `json:"mcp,required"`
	OAuth InfoGetResponseEndpointsOAuth `json:"oauth,required"`
	// OpenAPI spec endpoint
	Spec string `json:"spec,required"`
	// WebSocket events endpoint
	WsEvents string `json:"ws_events,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Mcp         respjson.Field
		OAuth       respjson.Field
		Spec        respjson.Field
		WsEvents    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InfoGetResponseEndpoints) RawJSON() string { return r.JSON.raw }
func (r *InfoGetResponseEndpoints) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InfoGetResponseEndpointsOAuth struct {
	// OAuth authorization endpoint
	AuthorizationEndpoint string `json:"authorization_endpoint,required"`
	// OAuth introspection endpoint
	IntrospectionEndpoint string `json:"introspection_endpoint,required"`
	// OAuth dynamic client registration endpoint
	RegistrationEndpoint string `json:"registration_endpoint,required"`
	// OAuth token revocation endpoint
	RevocationEndpoint string `json:"revocation_endpoint,required"`
	// OAuth token endpoint
	TokenEndpoint string `json:"token_endpoint,required"`
	// OAuth userinfo endpoint
	UserinfoEndpoint string `json:"userinfo_endpoint,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AuthorizationEndpoint respjson.Field
		IntrospectionEndpoint respjson.Field
		RegistrationEndpoint  respjson.Field
		RevocationEndpoint    respjson.Field
		TokenEndpoint         respjson.Field
		UserinfoEndpoint      respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InfoGetResponseEndpointsOAuth) RawJSON() string { return r.JSON.raw }
func (r *InfoGetResponseEndpointsOAuth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InfoGetResponsePlatform struct {
	// CPU architecture
	Arch string `json:"arch,required"`
	// Operating system identifier
	Os string `json:"os,required"`
	// Runtime release version
	Release string `json:"release"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Arch        respjson.Field
		Os          respjson.Field
		Release     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InfoGetResponsePlatform) RawJSON() string { return r.JSON.raw }
func (r *InfoGetResponsePlatform) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InfoGetResponseServer struct {
	// Base URL of the Connect server
	BaseURL string `json:"base_url,required"`
	// Listening host
	Hostname string `json:"hostname,required"`
	// Whether MCP endpoint is enabled
	McpEnabled bool `json:"mcp_enabled,required"`
	// Listening port
	Port int64 `json:"port,required"`
	// Whether remote access is enabled
	RemoteAccess bool `json:"remote_access,required"`
	// Server status
	Status string `json:"status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BaseURL      respjson.Field
		Hostname     respjson.Field
		McpEnabled   respjson.Field
		Port         respjson.Field
		RemoteAccess respjson.Field
		Status       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InfoGetResponseServer) RawJSON() string { return r.JSON.raw }
func (r *InfoGetResponseServer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
