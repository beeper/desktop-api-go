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
	"github.com/beeper/desktop-api-go/v5/packages/param"
	"github.com/beeper/desktop-api-go/v5/packages/respjson"
	"github.com/beeper/desktop-api-go/v5/shared/constant"
)

// Available bridges, bridge logins, login sessions for connect and reconnect
// flows, and advanced network capabilities.
//
// BridgeLoginSessionService contains methods and other services that help with
// interacting with the beeperdesktop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBridgeLoginSessionService] method instead.
type BridgeLoginSessionService struct {
	Options []option.RequestOption
	// Available bridges, bridge logins, login sessions for connect and reconnect
	// flows, and advanced network capabilities.
	Steps BridgeLoginSessionStepService
}

// NewBridgeLoginSessionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewBridgeLoginSessionService(opts ...option.RequestOption) (r BridgeLoginSessionService) {
	r = BridgeLoginSessionService{}
	r.Options = opts
	r.Steps = NewBridgeLoginSessionStepService(opts...)
	return
}

// Start a temporary bridge login session to connect a new chat account or
// reconnect an existing bridge login. Omit loginID and accountID to connect a new
// account.
func (r *BridgeLoginSessionService) New(ctx context.Context, bridgeID string, body BridgeLoginSessionNewParams, opts ...option.RequestOption) (res *LoginSession, err error) {
	opts = slices.Concat(r.Options, opts)
	if bridgeID == "" {
		err = errors.New("missing required bridgeID parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/bridges/%s/login-sessions", bridgeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get the current state of a temporary bridge login session.
func (r *BridgeLoginSessionService) Get(ctx context.Context, loginSessionID string, query BridgeLoginSessionGetParams, opts ...option.RequestOption) (res *LoginSession, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.BridgeID == "" {
		err = errors.New("missing required bridgeID parameter")
		return nil, err
	}
	if loginSessionID == "" {
		err = errors.New("missing required loginSessionID parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/bridges/%s/login-sessions/%s", query.BridgeID, loginSessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Cancel a temporary bridge login session.
func (r *BridgeLoginSessionService) Cancel(ctx context.Context, loginSessionID string, body BridgeLoginSessionCancelParams, opts ...option.RequestOption) (res *BridgeLoginSessionCancelResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.BridgeID == "" {
		err = errors.New("missing required bridgeID parameter")
		return nil, err
	}
	if loginSessionID == "" {
		err = errors.New("missing required loginSessionID parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/bridges/%s/login-sessions/%s", body.BridgeID, loginSessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type BridgeLoginSessionCancelResponse struct {
	BridgeID       string             `json:"bridgeID" api:"required"`
	LoginSessionID string             `json:"loginSessionID" api:"required"`
	Status         constant.Cancelled `json:"status" default:"cancelled"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BridgeID       respjson.Field
		LoginSessionID respjson.Field
		Status         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BridgeLoginSessionCancelResponse) RawJSON() string { return r.JSON.raw }
func (r *BridgeLoginSessionCancelResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BridgeLoginSessionNewParams struct {
	// Existing chat account ID to reconnect. Omit to connect a new account.
	AccountID param.Opt[string] `json:"accountID,omitzero"`
	// Optional flow ID returned by the list login flows endpoint. If omitted, Beeper
	// chooses the default flow.
	FlowID param.Opt[string] `json:"flowID,omitzero"`
	// Existing bridge login ID to reconnect. Omit to connect a new account.
	LoginID param.Opt[string] `json:"loginID,omitzero"`
	paramObj
}

func (r BridgeLoginSessionNewParams) MarshalJSON() (data []byte, err error) {
	type shadow BridgeLoginSessionNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BridgeLoginSessionNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BridgeLoginSessionGetParams struct {
	// Bridge ID.
	BridgeID string `path:"bridgeID" api:"required" json:"-"`
	paramObj
}

type BridgeLoginSessionCancelParams struct {
	// Bridge ID.
	BridgeID string `path:"bridgeID" api:"required" json:"-"`
	paramObj
}
