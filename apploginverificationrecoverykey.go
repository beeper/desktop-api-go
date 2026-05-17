// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package beeperdesktopapi

import (
	"context"
	"net/http"
	"slices"

	"github.com/beeper/desktop-api-go/v5/internal/apijson"
	"github.com/beeper/desktop-api-go/v5/internal/requestconfig"
	"github.com/beeper/desktop-api-go/v5/option"
	"github.com/beeper/desktop-api-go/v5/packages/param"
	"github.com/beeper/desktop-api-go/v5/packages/respjson"
)

// First-party sign-in and encrypted messaging setup for Beeper Desktop and Beeper
// Server.
//
// AppLoginVerificationRecoveryKeyService contains methods and other services that
// help with interacting with the beeperdesktop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAppLoginVerificationRecoveryKeyService] method instead.
type AppLoginVerificationRecoveryKeyService struct {
	Options []option.RequestOption
	// First-party sign-in and encrypted messaging setup for Beeper Desktop and Beeper
	// Server.
	Reset AppLoginVerificationRecoveryKeyResetService
}

// NewAppLoginVerificationRecoveryKeyService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAppLoginVerificationRecoveryKeyService(opts ...option.RequestOption) (r AppLoginVerificationRecoveryKeyService) {
	r = AppLoginVerificationRecoveryKeyService{}
	r.Options = opts
	r.Reset = NewAppLoginVerificationRecoveryKeyResetService(opts...)
	return
}

// Unlock encrypted messages with the user recovery key.
func (r *AppLoginVerificationRecoveryKeyService) Verify(ctx context.Context, body AppLoginVerificationRecoveryKeyVerifyParams, opts ...option.RequestOption) (res *AppLoginVerificationRecoveryKeyVerifyResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/app/setup/verification/recovery-key"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type AppLoginVerificationRecoveryKeyVerifyResponse struct {
	// Current app sign-in and encrypted messaging setup state.
	Session AppLoginVerificationRecoveryKeyVerifyResponseSession `json:"session" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Session     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppLoginVerificationRecoveryKeyVerifyResponse) RawJSON() string { return r.JSON.raw }
func (r *AppLoginVerificationRecoveryKeyVerifyResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current app sign-in and encrypted messaging setup state.
type AppLoginVerificationRecoveryKeyVerifyResponseSession struct {
	// Encrypted messaging setup status.
	E2EE AppLoginVerificationRecoveryKeyVerifyResponseSessionE2EE `json:"e2ee" api:"required"`
	// Current sign-in and encrypted messaging setup state for Beeper Desktop or Beeper
	// Server.
	//
	// Any of "needs-login", "initializing", "needs-cross-signing-setup",
	// "needs-verification", "needs-secrets", "needs-first-sync", "ready".
	State string `json:"state" api:"required"`
	// Signed-in account details. Omitted until sign-in is complete.
	Matrix AppLoginVerificationRecoveryKeyVerifyResponseSessionMatrix `json:"matrix"`
	// Trusted device verification progress.
	Verification AppLoginVerificationRecoveryKeyVerifyResponseSessionVerification `json:"verification"`
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
func (r AppLoginVerificationRecoveryKeyVerifyResponseSession) RawJSON() string { return r.JSON.raw }
func (r *AppLoginVerificationRecoveryKeyVerifyResponseSession) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Encrypted messaging setup status.
type AppLoginVerificationRecoveryKeyVerifyResponseSessionE2EE struct {
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
	Secrets AppLoginVerificationRecoveryKeyVerifyResponseSessionE2EESecrets `json:"secrets" api:"required"`
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
func (r AppLoginVerificationRecoveryKeyVerifyResponseSessionE2EE) RawJSON() string { return r.JSON.raw }
func (r *AppLoginVerificationRecoveryKeyVerifyResponseSessionE2EE) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Encrypted messaging keys available on this device.
type AppLoginVerificationRecoveryKeyVerifyResponseSessionE2EESecrets struct {
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
func (r AppLoginVerificationRecoveryKeyVerifyResponseSessionE2EESecrets) RawJSON() string {
	return r.JSON.raw
}
func (r *AppLoginVerificationRecoveryKeyVerifyResponseSessionE2EESecrets) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Signed-in account details. Omitted until sign-in is complete.
type AppLoginVerificationRecoveryKeyVerifyResponseSessionMatrix struct {
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
func (r AppLoginVerificationRecoveryKeyVerifyResponseSessionMatrix) RawJSON() string {
	return r.JSON.raw
}
func (r *AppLoginVerificationRecoveryKeyVerifyResponseSessionMatrix) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Trusted device verification progress.
type AppLoginVerificationRecoveryKeyVerifyResponseSessionVerification struct {
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
	Error AppLoginVerificationRecoveryKeyVerifyResponseSessionVerificationError `json:"error"`
	// Other device participating in verification.
	OtherDevice AppLoginVerificationRecoveryKeyVerifyResponseSessionVerificationOtherDevice `json:"otherDevice"`
	// Other Beeper user participating in verification.
	OtherUserID string `json:"otherUserID"`
	// QR verification data.
	Qr AppLoginVerificationRecoveryKeyVerifyResponseSessionVerificationQr `json:"qr"`
	// Emoji or number comparison data for verification.
	SAS AppLoginVerificationRecoveryKeyVerifyResponseSessionVerificationSAS `json:"sas"`
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
func (r AppLoginVerificationRecoveryKeyVerifyResponseSessionVerification) RawJSON() string {
	return r.JSON.raw
}
func (r *AppLoginVerificationRecoveryKeyVerifyResponseSessionVerification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Verification error details, if verification stopped.
type AppLoginVerificationRecoveryKeyVerifyResponseSessionVerificationError struct {
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
func (r AppLoginVerificationRecoveryKeyVerifyResponseSessionVerificationError) RawJSON() string {
	return r.JSON.raw
}
func (r *AppLoginVerificationRecoveryKeyVerifyResponseSessionVerificationError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Other device participating in verification.
type AppLoginVerificationRecoveryKeyVerifyResponseSessionVerificationOtherDevice struct {
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
func (r AppLoginVerificationRecoveryKeyVerifyResponseSessionVerificationOtherDevice) RawJSON() string {
	return r.JSON.raw
}
func (r *AppLoginVerificationRecoveryKeyVerifyResponseSessionVerificationOtherDevice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// QR verification data.
type AppLoginVerificationRecoveryKeyVerifyResponseSessionVerificationQr struct {
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
func (r AppLoginVerificationRecoveryKeyVerifyResponseSessionVerificationQr) RawJSON() string {
	return r.JSON.raw
}
func (r *AppLoginVerificationRecoveryKeyVerifyResponseSessionVerificationQr) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Emoji or number comparison data for verification.
type AppLoginVerificationRecoveryKeyVerifyResponseSessionVerificationSAS struct {
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
func (r AppLoginVerificationRecoveryKeyVerifyResponseSessionVerificationSAS) RawJSON() string {
	return r.JSON.raw
}
func (r *AppLoginVerificationRecoveryKeyVerifyResponseSessionVerificationSAS) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppLoginVerificationRecoveryKeyVerifyParams struct {
	// Recovery key saved by the user.
	RecoveryKey string `json:"recoveryKey" api:"required"`
	paramObj
}

func (r AppLoginVerificationRecoveryKeyVerifyParams) MarshalJSON() (data []byte, err error) {
	type shadow AppLoginVerificationRecoveryKeyVerifyParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppLoginVerificationRecoveryKeyVerifyParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
