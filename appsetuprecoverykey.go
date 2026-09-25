// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package beeperdesktopapi

import (
	"context"
	"net/http"
	"slices"

	"github.com/beeper/desktop-api-go/v6/internal/apijson"
	"github.com/beeper/desktop-api-go/v6/internal/requestconfig"
	"github.com/beeper/desktop-api-go/v6/option"
	"github.com/beeper/desktop-api-go/v6/packages/param"
	"github.com/beeper/desktop-api-go/v6/packages/respjson"
)

// Manage recovery-key setup for encrypted messages
//
// AppSetupRecoveryKeyService contains methods and other services that help with
// interacting with the beeperdesktop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAppSetupRecoveryKeyService] method instead.
type AppSetupRecoveryKeyService struct {
	Options []option.RequestOption
	// First-party sign-in and encrypted messaging setup for Beeper Desktop and Beeper
	// Server.
	Reset AppSetupRecoveryKeyResetService
}

// NewAppSetupRecoveryKeyService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAppSetupRecoveryKeyService(opts ...option.RequestOption) (r AppSetupRecoveryKeyService) {
	r = AppSetupRecoveryKeyService{}
	r.Options = opts
	r.Reset = NewAppSetupRecoveryKeyResetService(opts...)
	return
}

// Unlock encrypted messages with the user recovery key.
func (r *AppSetupRecoveryKeyService) Verify(ctx context.Context, body AppSetupRecoveryKeyVerifyParams, opts ...option.RequestOption) (res *AppSetupRecoveryKeyVerifyResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/app/setup/verification/recovery-key"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type AppSetupRecoveryKeyVerifyResponse struct {
	// Current app sign-in and encrypted messaging setup state.
	Session AppSetupRecoveryKeyVerifyResponseSession `json:"session" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Session     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppSetupRecoveryKeyVerifyResponse) RawJSON() string { return r.JSON.raw }
func (r *AppSetupRecoveryKeyVerifyResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current app sign-in and encrypted messaging setup state.
type AppSetupRecoveryKeyVerifyResponseSession struct {
	// Encrypted messaging setup status.
	E2EE AppSetupRecoveryKeyVerifyResponseSessionE2EE `json:"e2ee" api:"required"`
	// Current sign-in and encrypted messaging setup state for Beeper Desktop or Beeper
	// Server.
	//
	// Any of "needs-login", "initializing", "needs-cross-signing-setup",
	// "needs-verification", "needs-secrets", "needs-first-sync", "ready".
	State string `json:"state" api:"required"`
	// Signed-in account details. Omitted until sign-in is complete.
	Matrix AppSetupRecoveryKeyVerifyResponseSessionMatrix `json:"matrix"`
	// Trusted device verification progress.
	Verification AppSetupRecoveryKeyVerifyResponseSessionVerification `json:"verification"`
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
func (r AppSetupRecoveryKeyVerifyResponseSession) RawJSON() string { return r.JSON.raw }
func (r *AppSetupRecoveryKeyVerifyResponseSession) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Encrypted messaging setup status.
type AppSetupRecoveryKeyVerifyResponseSessionE2EE struct {
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
	Secrets AppSetupRecoveryKeyVerifyResponseSessionE2EESecrets `json:"secrets" api:"required"`
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
func (r AppSetupRecoveryKeyVerifyResponseSessionE2EE) RawJSON() string { return r.JSON.raw }
func (r *AppSetupRecoveryKeyVerifyResponseSessionE2EE) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Encrypted messaging keys available on this device.
type AppSetupRecoveryKeyVerifyResponseSessionE2EESecrets struct {
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
func (r AppSetupRecoveryKeyVerifyResponseSessionE2EESecrets) RawJSON() string { return r.JSON.raw }
func (r *AppSetupRecoveryKeyVerifyResponseSessionE2EESecrets) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Signed-in account details. Omitted until sign-in is complete.
type AppSetupRecoveryKeyVerifyResponseSessionMatrix struct {
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
func (r AppSetupRecoveryKeyVerifyResponseSessionMatrix) RawJSON() string { return r.JSON.raw }
func (r *AppSetupRecoveryKeyVerifyResponseSessionMatrix) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Trusted device verification progress.
type AppSetupRecoveryKeyVerifyResponseSessionVerification struct {
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
	Error AppSetupRecoveryKeyVerifyResponseSessionVerificationError `json:"error"`
	// Other device participating in verification.
	OtherDevice AppSetupRecoveryKeyVerifyResponseSessionVerificationOtherDevice `json:"otherDevice"`
	// Other Beeper user participating in verification.
	OtherUserID string `json:"otherUserID"`
	// QR verification data.
	QR AppSetupRecoveryKeyVerifyResponseSessionVerificationQR `json:"qr"`
	// Emoji or number comparison data for verification.
	SAS AppSetupRecoveryKeyVerifyResponseSessionVerificationSAS `json:"sas"`
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
		QR               respjson.Field
		SAS              respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppSetupRecoveryKeyVerifyResponseSessionVerification) RawJSON() string { return r.JSON.raw }
func (r *AppSetupRecoveryKeyVerifyResponseSessionVerification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Verification error details, if verification stopped.
type AppSetupRecoveryKeyVerifyResponseSessionVerificationError struct {
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
func (r AppSetupRecoveryKeyVerifyResponseSessionVerificationError) RawJSON() string {
	return r.JSON.raw
}
func (r *AppSetupRecoveryKeyVerifyResponseSessionVerificationError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Other device participating in verification.
type AppSetupRecoveryKeyVerifyResponseSessionVerificationOtherDevice struct {
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
func (r AppSetupRecoveryKeyVerifyResponseSessionVerificationOtherDevice) RawJSON() string {
	return r.JSON.raw
}
func (r *AppSetupRecoveryKeyVerifyResponseSessionVerificationOtherDevice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// QR verification data.
type AppSetupRecoveryKeyVerifyResponseSessionVerificationQR struct {
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
func (r AppSetupRecoveryKeyVerifyResponseSessionVerificationQR) RawJSON() string { return r.JSON.raw }
func (r *AppSetupRecoveryKeyVerifyResponseSessionVerificationQR) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Emoji or number comparison data for verification.
type AppSetupRecoveryKeyVerifyResponseSessionVerificationSAS struct {
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
func (r AppSetupRecoveryKeyVerifyResponseSessionVerificationSAS) RawJSON() string { return r.JSON.raw }
func (r *AppSetupRecoveryKeyVerifyResponseSessionVerificationSAS) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppSetupRecoveryKeyVerifyParams struct {
	// Recovery key saved by the user.
	RecoveryKey string `json:"recoveryKey" api:"required"`
	paramObj
}

func (r AppSetupRecoveryKeyVerifyParams) MarshalJSON() (data []byte, err error) {
	type shadow AppSetupRecoveryKeyVerifyParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppSetupRecoveryKeyVerifyParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
