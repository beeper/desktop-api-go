// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package beeperdesktopapi

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/beeper/desktop-api-go/v5/internal/apijson"
	"github.com/beeper/desktop-api-go/v5/internal/requestconfig"
	"github.com/beeper/desktop-api-go/v5/option"
	"github.com/beeper/desktop-api-go/v5/packages/respjson"
)

// Available bridges, bridge logins, login sessions for connect and reconnect
// flows, and advanced network capabilities.
//
// BridgeLoginFlowService contains methods and other services that help with
// interacting with the beeperdesktop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBridgeLoginFlowService] method instead.
type BridgeLoginFlowService struct {
	Options []option.RequestOption
}

// NewBridgeLoginFlowService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBridgeLoginFlowService(opts ...option.RequestOption) (r BridgeLoginFlowService) {
	r = BridgeLoginFlowService{}
	r.Options = opts
	return
}

// List connect and reconnect flow options for a bridge. Use a flowID when creating
// a bridge login session.
func (r *BridgeLoginFlowService) List(ctx context.Context, bridgeID string, opts ...option.RequestOption) (res *BridgeLoginFlowListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if bridgeID == "" {
		err = errors.New("missing required bridgeID parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/bridges/%s/login-flows", bridgeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type BridgeLoginFlowListResponse struct {
	Items []LoginFlow `json:"items" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BridgeLoginFlowListResponse) RawJSON() string { return r.JSON.raw }
func (r *BridgeLoginFlowListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
