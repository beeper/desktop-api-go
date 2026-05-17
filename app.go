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

// Manage Beeper app login and encrypted messaging setup
//
// AppService contains methods and other services that help with interacting with
// the beeperdesktop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAppService] method instead.
type AppService struct {
	Options []option.RequestOption
	// Complete first-party Beeper app login
	Login AppLoginService
	// Manage device verification transactions
	Verifications AppVerificationService
}

// NewAppService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAppService(opts ...option.RequestOption) (r AppService) {
	r = AppService{}
	r.Options = opts
	r.Login = NewAppLoginService(opts...)
	r.Verifications = NewAppVerificationService(opts...)
	return
}

// Return the current Beeper Desktop or Beeper Server sign-in and encrypted
// messaging setup state. This endpoint is public before sign-in so apps can
// discover that sign-in is needed; after sign-in, pass a read token.
func (r *AppService) Session(ctx context.Context, opts ...option.RequestOption) (res *AppSessionResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/app/setup"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type AppSessionResponse struct {
	// Encrypted messaging setup status.
	E2EE AppSessionResponseE2EE `json:"e2ee" api:"required"`
	// Current sign-in and encrypted messaging setup state for Beeper Desktop or Beeper
	// Server.
	//
	// Any of "needs-login", "initializing", "needs-cross-signing-setup",
	// "needs-verification", "needs-secrets", "needs-first-sync", "ready".
	State AppSessionResponseState `json:"state" api:"required"`
	// Signed-in account details. Omitted until sign-in is complete.
	Matrix AppSessionResponseMatrix `json:"matrix"`
	// Trusted device verification progress.
	Verification AppSessionResponseVerification `json:"verification"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		E2EE         respjson.Field
		State        respjson.Field
		Matrix       respjson.Field
		Verification respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppSessionResponse) RawJSON() string { return r.JSON.raw }
func (r *AppSessionResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Encrypted messaging setup status.
type AppSessionResponseE2EE struct {
	// Whether this account can verify trusted devices.
	CrossSigning bool `json:"crossSigning" api:"required"`
	// Whether the first encrypted message sync is complete.
	FirstSyncDone bool `json:"firstSyncDone" api:"required"`
	// Whether the user confirmed that they saved their recovery key.
	HasBackedUpRecoveryKey bool `json:"hasBackedUpRecoveryKey" api:"required"`
	// Whether encrypted messaging setup has started.
	Initialized bool `json:"initialized" api:"required"`
	// Whether encrypted message backup is available.
	KeyBackup bool `json:"keyBackup" api:"required"`
	// Encrypted messaging keys available on this device.
	Secrets AppSessionResponseE2EESecrets `json:"secrets" api:"required"`
	// Whether secure key storage is available.
	SecretStorage bool `json:"secretStorage" api:"required"`
	// Whether this device is trusted for encrypted messages.
	Verified bool `json:"verified" api:"required"`
	// Unix timestamp for when the recovery key was created.
	RecoveryKeyGeneratedAt float64 `json:"recoveryKeyGeneratedAt"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CrossSigning           respjson.Field
		FirstSyncDone          respjson.Field
		HasBackedUpRecoveryKey respjson.Field
		Initialized            respjson.Field
		KeyBackup              respjson.Field
		Secrets                respjson.Field
		SecretStorage          respjson.Field
		Verified               respjson.Field
		RecoveryKeyGeneratedAt respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppSessionResponseE2EE) RawJSON() string { return r.JSON.raw }
func (r *AppSessionResponseE2EE) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Encrypted messaging keys available on this device.
type AppSessionResponseE2EESecrets struct {
	// Whether the account identity key is available.
	MasterKey bool `json:"masterKey" api:"required"`
	// Whether the encrypted message backup key is available.
	MegolmBackupKey bool `json:"megolmBackupKey" api:"required"`
	// Whether a recovery key is available.
	RecoveryKey bool `json:"recoveryKey" api:"required"`
	// Whether the device trust key is available.
	SelfSigningKey bool `json:"selfSigningKey" api:"required"`
	// Whether the user trust key is available.
	UserSigningKey bool `json:"userSigningKey" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MasterKey       respjson.Field
		MegolmBackupKey respjson.Field
		RecoveryKey     respjson.Field
		SelfSigningKey  respjson.Field
		UserSigningKey  respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppSessionResponseE2EESecrets) RawJSON() string { return r.JSON.raw }
func (r *AppSessionResponseE2EESecrets) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current sign-in and encrypted messaging setup state for Beeper Desktop or Beeper
// Server.
type AppSessionResponseState string

const (
	AppSessionResponseStateNeedsLogin             AppSessionResponseState = "needs-login"
	AppSessionResponseStateInitializing           AppSessionResponseState = "initializing"
	AppSessionResponseStateNeedsCrossSigningSetup AppSessionResponseState = "needs-cross-signing-setup"
	AppSessionResponseStateNeedsVerification      AppSessionResponseState = "needs-verification"
	AppSessionResponseStateNeedsSecrets           AppSessionResponseState = "needs-secrets"
	AppSessionResponseStateNeedsFirstSync         AppSessionResponseState = "needs-first-sync"
	AppSessionResponseStateReady                  AppSessionResponseState = "ready"
)

// Signed-in account details. Omitted until sign-in is complete.
type AppSessionResponseMatrix struct {
	// Current device ID.
	DeviceID string `json:"deviceID" api:"required"`
	// Beeper homeserver URL for this account.
	Homeserver string `json:"homeserver" api:"required"`
	// Signed-in Beeper user ID.
	UserID string `json:"userID" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DeviceID    respjson.Field
		Homeserver  respjson.Field
		UserID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppSessionResponseMatrix) RawJSON() string { return r.JSON.raw }
func (r *AppSessionResponseMatrix) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Trusted device verification progress.
type AppSessionResponseVerification struct {
	// Verification ID to pass in verification action paths.
	ID string `json:"id" api:"required"`
	// Verification actions that are valid for the current state.
	//
	// Any of "accept", "cancel", "qr.confirmScanned", "sas.start", "sas.confirm".
	AvailableActions []string `json:"availableActions" api:"required"`
	// Whether this device started or received the verification.
	//
	// Any of "incoming", "outgoing".
	Direction string `json:"direction" api:"required"`
	// Verification methods supported for this transaction.
	//
	// Any of "qr", "sas".
	Methods []string `json:"methods" api:"required"`
	// Why this verification exists.
	//
	// Any of "login", "device".
	Purpose string `json:"purpose" api:"required"`
	// Current trusted-device verification state.
	//
	// Any of "requested", "ready", "sas_ready", "qr_scanned", "done", "cancelled",
	// "error".
	State string `json:"state" api:"required"`
	// Verification error details, if verification stopped.
	Error AppSessionResponseVerificationError `json:"error"`
	// Other device participating in verification.
	OtherDevice AppSessionResponseVerificationOtherDevice `json:"otherDevice"`
	// Other Beeper user participating in verification.
	OtherUserID string `json:"otherUserID"`
	// QR verification data.
	Qr AppSessionResponseVerificationQr `json:"qr"`
	// Emoji or number comparison data for verification.
	SAS AppSessionResponseVerificationSAS `json:"sas"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		AvailableActions respjson.Field
		Direction        respjson.Field
		Methods          respjson.Field
		Purpose          respjson.Field
		State            respjson.Field
		Error            respjson.Field
		OtherDevice      respjson.Field
		OtherUserID      respjson.Field
		Qr               respjson.Field
		SAS              respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppSessionResponseVerification) RawJSON() string { return r.JSON.raw }
func (r *AppSessionResponseVerification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Verification error details, if verification stopped.
type AppSessionResponseVerificationError struct {
	// Verification error code.
	Code string `json:"code" api:"required"`
	// User-facing verification error message.
	Reason string `json:"reason" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Reason      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppSessionResponseVerificationError) RawJSON() string { return r.JSON.raw }
func (r *AppSessionResponseVerificationError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Other device participating in verification.
type AppSessionResponseVerificationOtherDevice struct {
	// Other device ID.
	ID string `json:"id" api:"required"`
	// Other device display name, if known.
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppSessionResponseVerificationOtherDevice) RawJSON() string { return r.JSON.raw }
func (r *AppSessionResponseVerificationOtherDevice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// QR verification data.
type AppSessionResponseVerificationQr struct {
	// QR code payload to display for verification.
	Data string `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppSessionResponseVerificationQr) RawJSON() string { return r.JSON.raw }
func (r *AppSessionResponseVerificationQr) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Emoji or number comparison data for verification.
type AppSessionResponseVerificationSAS struct {
	// Emoji sequence to compare on both devices.
	Emojis string `json:"emojis" api:"required"`
	// Number sequence to compare on both devices.
	Decimals string `json:"decimals"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Emojis      respjson.Field
		Decimals    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppSessionResponseVerificationSAS) RawJSON() string { return r.JSON.raw }
func (r *AppSessionResponseVerificationSAS) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
