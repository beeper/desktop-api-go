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
	"github.com/beeper/desktop-api-go/shared"
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

// List Chat Accounts connected to this Beeper Desktop instance, including bridge
// metadata and network identity.
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
	// User the account belongs to.
	User shared.User `json:"user" api:"required"`
	// Human-friendly network name for the account. Omitted when the network is
	// unknown.
	Network string `json:"network"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountID   respjson.Field
		Bridge      respjson.Field
		User        respjson.Field
		Network     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Account) RawJSON() string { return r.JSON.raw }
func (r *Account) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Bridge metadata for the account. Available in Beeper Desktop v4.2.785+.
type AccountBridge struct {
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
func (r AccountBridge) RawJSON() string { return r.JSON.raw }
func (r *AccountBridge) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
