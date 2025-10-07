// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package beeperdesktopapi

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/beeper-desktop-api-go/internal/apijson"
	"github.com/stainless-sdks/beeper-desktop-api-go/internal/requestconfig"
	"github.com/stainless-sdks/beeper-desktop-api-go/option"
	"github.com/stainless-sdks/beeper-desktop-api-go/packages/respjson"
	"github.com/stainless-sdks/beeper-desktop-api-go/shared"
)

// Accounts operations
//
// AccountService contains methods and other services that help with interacting
// with the beeperdesktop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountService] method instead.
type AccountService struct {
	Options []option.RequestOption
}

// NewAccountService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAccountService(opts ...option.RequestOption) (r AccountService) {
	r = AccountService{}
	r.Options = opts
	return
}

// List connected Beeper accounts available on this device
func (r *AccountService) List(ctx context.Context, opts ...option.RequestOption) (res *[]Account, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/get-accounts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// A chat account added to Beeper
type Account struct {
	// Chat account added to Beeper. Use this to route account-scoped actions.
	AccountID string `json:"accountID,required"`
	// Display-only human-readable network name (e.g., 'WhatsApp', 'Messenger'). You
	// MUST use 'accountID' to perform actions.
	Network string `json:"network,required"`
	// A person on or reachable through Beeper. Values are best-effort and can vary by
	// network.
	User shared.User `json:"user,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountID   respjson.Field
		Network     respjson.Field
		User        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Account) RawJSON() string { return r.JSON.raw }
func (r *Account) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
