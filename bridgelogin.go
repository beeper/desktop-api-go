// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package beeperdesktopapi

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/beeper/desktop-api-go/v6/internal/apijson"
	"github.com/beeper/desktop-api-go/v6/internal/requestconfig"
	"github.com/beeper/desktop-api-go/v6/option"
	"github.com/beeper/desktop-api-go/v6/packages/param"
	"github.com/beeper/desktop-api-go/v6/packages/respjson"
	"github.com/beeper/desktop-api-go/v6/shared/constant"
)

// Available bridges, bridge logins, login sessions for connect and reconnect
// flows, and advanced network capabilities.
//
// BridgeLoginService contains methods and other services that help with
// interacting with the beeperdesktop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBridgeLoginService] method instead.
type BridgeLoginService struct {
	Options []option.RequestOption
}

// NewBridgeLoginService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBridgeLoginService(opts ...option.RequestOption) (r BridgeLoginService) {
	r = BridgeLoginService{}
	r.Options = opts
	return
}

// Get one bridge login.
func (r *BridgeLoginService) Get(ctx context.Context, loginID string, query BridgeLoginGetParams, opts ...option.RequestOption) (res *BridgeLogin, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.BridgeID == "" {
		err = errors.New("missing required bridgeID parameter")
		return nil, err
	}
	if loginID == "" {
		err = errors.New("missing required loginID parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/bridges/%s/logins/%s", query.BridgeID, loginID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List bridge logins. A bridge login is a signed-in identity for a bridge and can
// contain one or more chat accounts.
func (r *BridgeLoginService) List(ctx context.Context, bridgeID string, opts ...option.RequestOption) (res *BridgeLoginListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if bridgeID == "" {
		err = errors.New("missing required bridgeID parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/bridges/%s/logins", bridgeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Remove a bridge login from this device or, when supported by the bridge, from
// all devices.
func (r *BridgeLoginService) Remove(ctx context.Context, loginID string, params BridgeLoginRemoveParams, opts ...option.RequestOption) (res *BridgeLoginRemoveResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.BridgeID == "" {
		err = errors.New("missing required bridgeID parameter")
		return nil, err
	}
	if loginID == "" {
		err = errors.New("missing required loginID parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/bridges/%s/logins/%s/remove", params.BridgeID, loginID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

type BridgeLoginListResponse struct {
	Items []BridgeLogin `json:"items" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BridgeLoginListResponse) RawJSON() string { return r.JSON.raw }
func (r *BridgeLoginListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BridgeLoginRemoveResponse struct {
	BridgeID string `json:"bridgeID" api:"required"`
	LoginID  string `json:"loginID" api:"required"`
	// Where this bridge login should be removed.
	//
	// Any of "current-device", "all-devices".
	Scope              BridgeLoginRemoveResponseScope `json:"scope" api:"required"`
	Status             constant.Removed               `json:"status" default:"removed"`
	AffectedAccountIDs []string                       `json:"affectedAccountIDs"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BridgeID           respjson.Field
		LoginID            respjson.Field
		Scope              respjson.Field
		Status             respjson.Field
		AffectedAccountIDs respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BridgeLoginRemoveResponse) RawJSON() string { return r.JSON.raw }
func (r *BridgeLoginRemoveResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Where this bridge login should be removed.
type BridgeLoginRemoveResponseScope string

const (
	BridgeLoginRemoveResponseScopeCurrentDevice BridgeLoginRemoveResponseScope = "current-device"
	BridgeLoginRemoveResponseScopeAllDevices    BridgeLoginRemoveResponseScope = "all-devices"
)

type BridgeLoginGetParams struct {
	// Bridge ID.
	BridgeID string `path:"bridgeID" api:"required" json:"-"`
	paramObj
}

type BridgeLoginRemoveParams struct {
	// Bridge ID.
	BridgeID string `path:"bridgeID" api:"required" json:"-"`
	// Where this bridge login should be removed.
	//
	// Any of "current-device", "all-devices".
	Scope BridgeLoginRemoveParamsScope `json:"scope,omitzero" api:"required"`
	paramObj
}

func (r BridgeLoginRemoveParams) MarshalJSON() (data []byte, err error) {
	type shadow BridgeLoginRemoveParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BridgeLoginRemoveParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Where this bridge login should be removed.
type BridgeLoginRemoveParamsScope string

const (
	BridgeLoginRemoveParamsScopeCurrentDevice BridgeLoginRemoveParamsScope = "current-device"
	BridgeLoginRemoveParamsScopeAllDevices    BridgeLoginRemoveParamsScope = "all-devices"
)
