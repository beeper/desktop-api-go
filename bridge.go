// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package beeperdesktopapi

import (
	"context"
	"net/http"
	"slices"

	"github.com/beeper/desktop-api-go/v5/internal/apijson"
	"github.com/beeper/desktop-api-go/v5/internal/requestconfig"
	"github.com/beeper/desktop-api-go/v5/option"
	"github.com/beeper/desktop-api-go/v5/packages/respjson"
)

// Manage bridge-backed account types and account availability
//
// BridgeService contains methods and other services that help with interacting
// with the beeperdesktop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBridgeService] method instead.
type BridgeService struct {
	Options []option.RequestOption
}

// NewBridgeService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBridgeService(opts ...option.RequestOption) (r BridgeService) {
	r = BridgeService{}
	r.Options = opts
	return
}

// List bridge-backed account types that can be shown in add-account flows, grouped
// with connected accounts that use the same Account schema as GET /v1/accounts.
func (r *BridgeService) List(ctx context.Context, opts ...option.RequestOption) (res *BridgeListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/bridges"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Bridge-backed account type that can be shown in add-account flows.
type BridgeAvailability struct {
	// Connected accounts for this bridge. Uses the same Account schema as GET
	// /v1/accounts.
	Accounts []Account `json:"accounts" api:"required"`
	// Number of active accounts for this network on this device.
	ActiveAccountCount int64 `json:"activeAccountCount" api:"required"`
	// Bridge metadata for the account. Available in Beeper Desktop v4.2.785+.
	Bridge BridgeAvailabilityBridge `json:"bridge" api:"required"`
	// Human-friendly account type name shown in Beeper Desktop.
	DisplayName string `json:"displayName" api:"required"`
	// Login mode used by Beeper Desktop for this bridge.
	LoginMode string `json:"loginMode" api:"required"`
	// Whether this bridge can currently be used to add an account.
	//
	// Any of "available", "connected", "limit_reached", "temporarily_unavailable".
	Status BridgeAvailabilityStatus `json:"status" api:"required"`
	// Network grouping used for account counts and limits.
	Network string `json:"network"`
	// Human-friendly status text matching Beeper Desktop account management language.
	StatusText string `json:"statusText"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Accounts           respjson.Field
		ActiveAccountCount respjson.Field
		Bridge             respjson.Field
		DisplayName        respjson.Field
		LoginMode          respjson.Field
		Status             respjson.Field
		Network            respjson.Field
		StatusText         respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BridgeAvailability) RawJSON() string { return r.JSON.raw }
func (r *BridgeAvailability) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Bridge metadata for the account. Available in Beeper Desktop v4.2.785+.
type BridgeAvailabilityBridge struct {
	// Bridge instance identifier. Matrix and cloud bridges often use the bridge type
	// (for example matrix or discordgo); local bridges use a local bridge ID (for
	// example local-whatsapp). Available in Beeper Desktop v4.2.785+.
	ID string `json:"id" api:"required"`
	// Bridge provider for the account. Available in Beeper Desktop v4.2.785+.
	//
	// Any of "cloud", "self-hosted", "local", "platform-sdk".
	Provider string `json:"provider" api:"required"`
	// Bridge type, such as matrix, discordgo, slackgo, whatsapp, telegram, or twitter.
	// Available in Beeper Desktop v4.2.785+.
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Provider    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BridgeAvailabilityBridge) RawJSON() string { return r.JSON.raw }
func (r *BridgeAvailabilityBridge) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether this bridge can currently be used to add an account.
type BridgeAvailabilityStatus string

const (
	BridgeAvailabilityStatusAvailable              BridgeAvailabilityStatus = "available"
	BridgeAvailabilityStatusConnected              BridgeAvailabilityStatus = "connected"
	BridgeAvailabilityStatusLimitReached           BridgeAvailabilityStatus = "limit_reached"
	BridgeAvailabilityStatusTemporarilyUnavailable BridgeAvailabilityStatus = "temporarily_unavailable"
)

// Bridge-backed account types and their connected accounts.
type BridgeListResponse struct {
	Items []BridgeAvailability `json:"items" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BridgeListResponse) RawJSON() string { return r.JSON.raw }
func (r *BridgeListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
