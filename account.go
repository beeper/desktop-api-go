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
	"github.com/beeper/desktop-api-go/v5/shared"
)

// Manage connected chat accounts
//
// AccountService contains methods and other services that help with interacting
// with the beeperdesktop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountService] method instead.
type AccountService struct {
	Options []option.RequestOption
	// Manage contacts on a specific account
	Contacts AccountContactService
}

// NewAccountService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAccountService(opts ...option.RequestOption) (r AccountService) {
	r = AccountService{}
	r.Options = opts
	r.Contacts = NewAccountContactService(opts...)
	return
}

// Get one chat account connected to this Beeper Client API server.
func (r *AccountService) Get(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountID == "" {
		err = errors.New("missing required accountID parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/accounts/%s", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List chat accounts connected to this Beeper Client API server, including bridge,
// network, user identity, and connection status.
func (r *AccountService) List(ctx context.Context, opts ...option.RequestOption) (res *[]Account, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/accounts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// A chat account added to Beeper.
type Account struct {
	// Chat account added to Beeper. Use this to route account-scoped actions. Examples
	// include matrix for Beeper/Matrix, discordgo for a cloud bridge,
	// slackgo.TEAM-USER for workspace-scoped cloud bridges, and local-whatsapp*ba*...
	// for local bridges.
	AccountID string `json:"accountID" api:"required"`
	// Bridge metadata for the account. Available in Beeper Desktop v4.2.785+.
	Bridge AccountBridge `json:"bridge" api:"required"`
	// Current connection status for this account.
	//
	// Any of "connected", "connecting", "backfilling", "connection_required",
	// "reconnect_required", "attention_required", "disconnected", "disabled".
	Status AccountStatus `json:"status" api:"required"`
	// User the account belongs to.
	User shared.User `json:"user" api:"required"`
	// Runtime chat/message capabilities for this connected account, when available.
	Capabilities map[string]any `json:"capabilities"`
	// Bridge login ID for this account, when known. One bridge login can contain
	// multiple chat accounts.
	LoginID string `json:"loginID"`
	// Human-friendly network name for the account. Omitted when the network is
	// unknown.
	Network string `json:"network"`
	// Human-friendly account status text.
	StatusText string `json:"statusText"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountID    respjson.Field
		Bridge       respjson.Field
		Status       respjson.Field
		User         respjson.Field
		Capabilities respjson.Field
		LoginID      respjson.Field
		Network      respjson.Field
		StatusText   respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Account) RawJSON() string { return r.JSON.raw }
func (r *Account) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current connection status for this account.
type AccountStatus string

const (
	AccountStatusConnected          AccountStatus = "connected"
	AccountStatusConnecting         AccountStatus = "connecting"
	AccountStatusBackfilling        AccountStatus = "backfilling"
	AccountStatusConnectionRequired AccountStatus = "connection_required"
	AccountStatusReconnectRequired  AccountStatus = "reconnect_required"
	AccountStatusAttentionRequired  AccountStatus = "attention_required"
	AccountStatusDisconnected       AccountStatus = "disconnected"
	AccountStatusDisabled           AccountStatus = "disabled"
)

// Bridge metadata for the account. Available in Beeper Desktop v4.2.785+.
type AccountBridge struct {
	// Bridge identifier. Beeper Cloud accounts often use the network type (for example
	// matrix or discordgo); on-device accounts use a local bridge ID (for example
	// local-whatsapp). Available in Beeper Desktop v4.2.785+.
	ID string `json:"id" api:"required"`
	// Where this account runs: on this device or in Beeper Cloud. Available in Beeper
	// Desktop v4.2.785+.
	//
	// Any of "cloud", "self-hosted", "local", "platform-sdk".
	Provider AccountBridgeProvider `json:"provider" api:"required"`
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
func (r AccountBridge) RawJSON() string { return r.JSON.raw }
func (r *AccountBridge) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Where this account runs: on this device or in Beeper Cloud. Available in Beeper
// Desktop v4.2.785+.
type AccountBridgeProvider string

const (
	AccountBridgeProviderCloud       AccountBridgeProvider = "cloud"
	AccountBridgeProviderSelfHosted  AccountBridgeProvider = "self-hosted"
	AccountBridgeProviderLocal       AccountBridgeProvider = "local"
	AccountBridgeProviderPlatformSDK AccountBridgeProvider = "platform-sdk"
)

// A chat account added to Beeper.
type AccountGetResponse struct {
	// Chat account added to Beeper. Use this to route account-scoped actions. Examples
	// include matrix for Beeper/Matrix, discordgo for a cloud bridge,
	// slackgo.TEAM-USER for workspace-scoped cloud bridges, and local-whatsapp*ba*...
	// for local bridges.
	AccountID string `json:"accountID" api:"required"`
	// Bridge metadata for the account. Available in Beeper Desktop v4.2.785+.
	Bridge AccountBridge `json:"bridge" api:"required"`
	// Current connection status for this account.
	//
	// Any of "connected", "connecting", "backfilling", "connection_required",
	// "reconnect_required", "attention_required", "disconnected", "disabled".
	Status AccountGetResponseStatus `json:"status" api:"required"`
	// User the account belongs to.
	User shared.User `json:"user" api:"required"`
	// Runtime chat/message capabilities for this connected account, when available.
	Capabilities map[string]any `json:"capabilities"`
	// Bridge login ID for this account, when known. One bridge login can contain
	// multiple chat accounts.
	LoginID string `json:"loginID"`
	// Human-friendly network name for the account. Omitted when the network is
	// unknown.
	Network string `json:"network"`
	// Human-friendly account status text.
	StatusText string `json:"statusText"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountID    respjson.Field
		Bridge       respjson.Field
		Status       respjson.Field
		User         respjson.Field
		Capabilities respjson.Field
		LoginID      respjson.Field
		Network      respjson.Field
		StatusText   respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountGetResponse) RawJSON() string { return r.JSON.raw }
func (r *AccountGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current connection status for this account.
type AccountGetResponseStatus string

const (
	AccountGetResponseStatusConnected          AccountGetResponseStatus = "connected"
	AccountGetResponseStatusConnecting         AccountGetResponseStatus = "connecting"
	AccountGetResponseStatusBackfilling        AccountGetResponseStatus = "backfilling"
	AccountGetResponseStatusConnectionRequired AccountGetResponseStatus = "connection_required"
	AccountGetResponseStatusReconnectRequired  AccountGetResponseStatus = "reconnect_required"
	AccountGetResponseStatusAttentionRequired  AccountGetResponseStatus = "attention_required"
	AccountGetResponseStatusDisconnected       AccountGetResponseStatus = "disconnected"
	AccountGetResponseStatusDisabled           AccountGetResponseStatus = "disabled"
)
