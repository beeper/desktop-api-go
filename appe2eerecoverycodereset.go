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

// First-party sign-in and encrypted messaging setup for Beeper Desktop.
//
// AppE2eeRecoveryCodeResetService contains methods and other services that help
// with interacting with the beeperdesktop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAppE2eeRecoveryCodeResetService] method instead.
type AppE2eeRecoveryCodeResetService struct {
	Options []option.RequestOption
}

// NewAppE2eeRecoveryCodeResetService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAppE2eeRecoveryCodeResetService(opts ...option.RequestOption) (r AppE2eeRecoveryCodeResetService) {
	r = AppE2eeRecoveryCodeResetService{}
	r.Options = opts
	return
}

// Create a new recovery key when the user cannot use the existing one.
func (r *AppE2eeRecoveryCodeResetService) New(ctx context.Context, body AppE2eeRecoveryCodeResetNewParams, opts ...option.RequestOption) (res *AppE2eeRecoveryCodeResetNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/app/e2ee/recovery-code/reset"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Confirm that the new recovery key should be used for this account.
func (r *AppE2eeRecoveryCodeResetService) Confirm(ctx context.Context, body AppE2eeRecoveryCodeResetConfirmParams, opts ...option.RequestOption) (res *AppE2eeRecoveryCodeResetConfirmResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/app/e2ee/recovery-code/reset/confirm"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type AppE2eeRecoveryCodeResetNewResponse struct {
	// Current onboarding state after creating the new recovery key.
	AppState AppE2eeRecoveryCodeResetNewResponseAppState `json:"appState" api:"required"`
	// New recovery key. Show it once and ask the user to save it.
	RecoveryCode string `json:"recoveryCode" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AppState     respjson.Field
		RecoveryCode respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppE2eeRecoveryCodeResetNewResponse) RawJSON() string { return r.JSON.raw }
func (r *AppE2eeRecoveryCodeResetNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current onboarding state after creating the new recovery key.
type AppE2eeRecoveryCodeResetNewResponseAppState struct {
	// Encrypted messaging setup status.
	E2ee AppE2eeRecoveryCodeResetNewResponseAppStateE2ee `json:"e2ee" api:"required"`
	// Current onboarding state for Beeper Desktop.
	//
	// Any of "needs-login", "initializing", "needs-cross-signing-setup",
	// "needs-verification", "needs-secrets", "needs-first-sync", "ready".
	State string `json:"state" api:"required"`
	// Signed-in account details. Omitted until sign-in is complete.
	Matrix AppE2eeRecoveryCodeResetNewResponseAppStateMatrix `json:"matrix"`
	// Trusted-device verification progress.
	Verification AppE2eeRecoveryCodeResetNewResponseAppStateVerification `json:"verification"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		E2ee         respjson.Field
		State        respjson.Field
		Matrix       respjson.Field
		Verification respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppE2eeRecoveryCodeResetNewResponseAppState) RawJSON() string { return r.JSON.raw }
func (r *AppE2eeRecoveryCodeResetNewResponseAppState) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Encrypted messaging setup status.
type AppE2eeRecoveryCodeResetNewResponseAppStateE2ee struct {
	// Whether this account can verify trusted devices.
	CrossSigning bool `json:"crossSigning" api:"required"`
	// Whether the first encrypted message sync is complete.
	FirstSyncDone bool `json:"firstSyncDone" api:"required"`
	// Whether the user confirmed that they saved their recovery key.
	HasBackedUpCode bool `json:"hasBackedUpCode" api:"required"`
	// Whether encrypted messaging setup has started.
	Initialized bool `json:"initialized" api:"required"`
	// Whether encrypted message backup is available.
	KeyBackup bool `json:"keyBackup" api:"required"`
	// Encrypted messaging keys available on this device.
	Secrets AppE2eeRecoveryCodeResetNewResponseAppStateE2eeSecrets `json:"secrets" api:"required"`
	// Whether secure key storage is available.
	SecretStorage bool `json:"secretStorage" api:"required"`
	// Whether this device is trusted for encrypted messages.
	Verified bool `json:"verified" api:"required"`
	// Unix timestamp for when the recovery key was created.
	RecoveryCodeGeneratedAt float64 `json:"recoveryCodeGeneratedAt"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CrossSigning            respjson.Field
		FirstSyncDone           respjson.Field
		HasBackedUpCode         respjson.Field
		Initialized             respjson.Field
		KeyBackup               respjson.Field
		Secrets                 respjson.Field
		SecretStorage           respjson.Field
		Verified                respjson.Field
		RecoveryCodeGeneratedAt respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppE2eeRecoveryCodeResetNewResponseAppStateE2ee) RawJSON() string { return r.JSON.raw }
func (r *AppE2eeRecoveryCodeResetNewResponseAppStateE2ee) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Encrypted messaging keys available on this device.
type AppE2eeRecoveryCodeResetNewResponseAppStateE2eeSecrets struct {
	// Whether the account identity key is available.
	MasterKey bool `json:"masterKey" api:"required"`
	// Whether the encrypted message backup key is available.
	MegolmBackupKey bool `json:"megolmBackupKey" api:"required"`
	// Whether a recovery key is available.
	RecoveryCode bool `json:"recoveryCode" api:"required"`
	// Whether the device trust key is available.
	SelfSigningKey bool `json:"selfSigningKey" api:"required"`
	// Whether the user trust key is available.
	UserSigningKey bool `json:"userSigningKey" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MasterKey       respjson.Field
		MegolmBackupKey respjson.Field
		RecoveryCode    respjson.Field
		SelfSigningKey  respjson.Field
		UserSigningKey  respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppE2eeRecoveryCodeResetNewResponseAppStateE2eeSecrets) RawJSON() string { return r.JSON.raw }
func (r *AppE2eeRecoveryCodeResetNewResponseAppStateE2eeSecrets) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Signed-in account details. Omitted until sign-in is complete.
type AppE2eeRecoveryCodeResetNewResponseAppStateMatrix struct {
	// Current device ID.
	DeviceID string `json:"deviceID" api:"required"`
	// Beeper server URL for this account.
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
func (r AppE2eeRecoveryCodeResetNewResponseAppStateMatrix) RawJSON() string { return r.JSON.raw }
func (r *AppE2eeRecoveryCodeResetNewResponseAppStateMatrix) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Trusted-device verification progress.
type AppE2eeRecoveryCodeResetNewResponseAppStateVerification struct {
	// Verification actions that are valid for the current state.
	//
	// Any of "create", "qr.scan", "accept", "cancel", "qr.confirmScanned",
	// "sas.start", "sas.confirm".
	AvailableActions []string `json:"availableActions" api:"required"`
	// Current trusted-device verification state.
	//
	// Any of "idle", "requested", "ready", "sas_ready", "qr_scanned", "done",
	// "cancelled", "error".
	State string `json:"state" api:"required"`
	// Verification error details, if verification stopped.
	Error AppE2eeRecoveryCodeResetNewResponseAppStateVerificationError `json:"error"`
	// User ID that started verification.
	From string `json:"from"`
	// Device that started verification.
	FromDevice string `json:"fromDevice"`
	// Other device participating in verification.
	OtherDevice string `json:"otherDevice"`
	// QR code payload to display for verification.
	QrData string `json:"qrData"`
	// Emoji or number comparison data for verification.
	Sas AppE2eeRecoveryCodeResetNewResponseAppStateVerificationSas `json:"sas"`
	// Whether emoji comparison is available.
	SupportsSas bool `json:"supportsSAS"`
	// Whether QR code verification is available.
	SupportsScanQrCode bool `json:"supportsScanQRCode"`
	// Verification ID to pass in verification action paths.
	VerificationID string `json:"verificationID"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AvailableActions   respjson.Field
		State              respjson.Field
		Error              respjson.Field
		From               respjson.Field
		FromDevice         respjson.Field
		OtherDevice        respjson.Field
		QrData             respjson.Field
		Sas                respjson.Field
		SupportsSas        respjson.Field
		SupportsScanQrCode respjson.Field
		VerificationID     respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppE2eeRecoveryCodeResetNewResponseAppStateVerification) RawJSON() string { return r.JSON.raw }
func (r *AppE2eeRecoveryCodeResetNewResponseAppStateVerification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Verification error details, if verification stopped.
type AppE2eeRecoveryCodeResetNewResponseAppStateVerificationError struct {
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
func (r AppE2eeRecoveryCodeResetNewResponseAppStateVerificationError) RawJSON() string {
	return r.JSON.raw
}
func (r *AppE2eeRecoveryCodeResetNewResponseAppStateVerificationError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Emoji or number comparison data for verification.
type AppE2eeRecoveryCodeResetNewResponseAppStateVerificationSas struct {
	// Number sequence to compare on both devices.
	Decimals string `json:"decimals" api:"required"`
	// Emoji sequence to compare on both devices.
	Emojis string `json:"emojis" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Decimals    respjson.Field
		Emojis      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppE2eeRecoveryCodeResetNewResponseAppStateVerificationSas) RawJSON() string {
	return r.JSON.raw
}
func (r *AppE2eeRecoveryCodeResetNewResponseAppStateVerificationSas) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppE2eeRecoveryCodeResetConfirmResponse struct {
	// Current onboarding state after the requested step.
	AppState AppE2eeRecoveryCodeResetConfirmResponseAppState `json:"appState" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AppState    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppE2eeRecoveryCodeResetConfirmResponse) RawJSON() string { return r.JSON.raw }
func (r *AppE2eeRecoveryCodeResetConfirmResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current onboarding state after the requested step.
type AppE2eeRecoveryCodeResetConfirmResponseAppState struct {
	// Encrypted messaging setup status.
	E2ee AppE2eeRecoveryCodeResetConfirmResponseAppStateE2ee `json:"e2ee" api:"required"`
	// Current onboarding state for Beeper Desktop.
	//
	// Any of "needs-login", "initializing", "needs-cross-signing-setup",
	// "needs-verification", "needs-secrets", "needs-first-sync", "ready".
	State string `json:"state" api:"required"`
	// Signed-in account details. Omitted until sign-in is complete.
	Matrix AppE2eeRecoveryCodeResetConfirmResponseAppStateMatrix `json:"matrix"`
	// Trusted-device verification progress.
	Verification AppE2eeRecoveryCodeResetConfirmResponseAppStateVerification `json:"verification"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		E2ee         respjson.Field
		State        respjson.Field
		Matrix       respjson.Field
		Verification respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppE2eeRecoveryCodeResetConfirmResponseAppState) RawJSON() string { return r.JSON.raw }
func (r *AppE2eeRecoveryCodeResetConfirmResponseAppState) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Encrypted messaging setup status.
type AppE2eeRecoveryCodeResetConfirmResponseAppStateE2ee struct {
	// Whether this account can verify trusted devices.
	CrossSigning bool `json:"crossSigning" api:"required"`
	// Whether the first encrypted message sync is complete.
	FirstSyncDone bool `json:"firstSyncDone" api:"required"`
	// Whether the user confirmed that they saved their recovery key.
	HasBackedUpCode bool `json:"hasBackedUpCode" api:"required"`
	// Whether encrypted messaging setup has started.
	Initialized bool `json:"initialized" api:"required"`
	// Whether encrypted message backup is available.
	KeyBackup bool `json:"keyBackup" api:"required"`
	// Encrypted messaging keys available on this device.
	Secrets AppE2eeRecoveryCodeResetConfirmResponseAppStateE2eeSecrets `json:"secrets" api:"required"`
	// Whether secure key storage is available.
	SecretStorage bool `json:"secretStorage" api:"required"`
	// Whether this device is trusted for encrypted messages.
	Verified bool `json:"verified" api:"required"`
	// Unix timestamp for when the recovery key was created.
	RecoveryCodeGeneratedAt float64 `json:"recoveryCodeGeneratedAt"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CrossSigning            respjson.Field
		FirstSyncDone           respjson.Field
		HasBackedUpCode         respjson.Field
		Initialized             respjson.Field
		KeyBackup               respjson.Field
		Secrets                 respjson.Field
		SecretStorage           respjson.Field
		Verified                respjson.Field
		RecoveryCodeGeneratedAt respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppE2eeRecoveryCodeResetConfirmResponseAppStateE2ee) RawJSON() string { return r.JSON.raw }
func (r *AppE2eeRecoveryCodeResetConfirmResponseAppStateE2ee) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Encrypted messaging keys available on this device.
type AppE2eeRecoveryCodeResetConfirmResponseAppStateE2eeSecrets struct {
	// Whether the account identity key is available.
	MasterKey bool `json:"masterKey" api:"required"`
	// Whether the encrypted message backup key is available.
	MegolmBackupKey bool `json:"megolmBackupKey" api:"required"`
	// Whether a recovery key is available.
	RecoveryCode bool `json:"recoveryCode" api:"required"`
	// Whether the device trust key is available.
	SelfSigningKey bool `json:"selfSigningKey" api:"required"`
	// Whether the user trust key is available.
	UserSigningKey bool `json:"userSigningKey" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MasterKey       respjson.Field
		MegolmBackupKey respjson.Field
		RecoveryCode    respjson.Field
		SelfSigningKey  respjson.Field
		UserSigningKey  respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppE2eeRecoveryCodeResetConfirmResponseAppStateE2eeSecrets) RawJSON() string {
	return r.JSON.raw
}
func (r *AppE2eeRecoveryCodeResetConfirmResponseAppStateE2eeSecrets) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Signed-in account details. Omitted until sign-in is complete.
type AppE2eeRecoveryCodeResetConfirmResponseAppStateMatrix struct {
	// Current device ID.
	DeviceID string `json:"deviceID" api:"required"`
	// Beeper server URL for this account.
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
func (r AppE2eeRecoveryCodeResetConfirmResponseAppStateMatrix) RawJSON() string { return r.JSON.raw }
func (r *AppE2eeRecoveryCodeResetConfirmResponseAppStateMatrix) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Trusted-device verification progress.
type AppE2eeRecoveryCodeResetConfirmResponseAppStateVerification struct {
	// Verification actions that are valid for the current state.
	//
	// Any of "create", "qr.scan", "accept", "cancel", "qr.confirmScanned",
	// "sas.start", "sas.confirm".
	AvailableActions []string `json:"availableActions" api:"required"`
	// Current trusted-device verification state.
	//
	// Any of "idle", "requested", "ready", "sas_ready", "qr_scanned", "done",
	// "cancelled", "error".
	State string `json:"state" api:"required"`
	// Verification error details, if verification stopped.
	Error AppE2eeRecoveryCodeResetConfirmResponseAppStateVerificationError `json:"error"`
	// User ID that started verification.
	From string `json:"from"`
	// Device that started verification.
	FromDevice string `json:"fromDevice"`
	// Other device participating in verification.
	OtherDevice string `json:"otherDevice"`
	// QR code payload to display for verification.
	QrData string `json:"qrData"`
	// Emoji or number comparison data for verification.
	Sas AppE2eeRecoveryCodeResetConfirmResponseAppStateVerificationSas `json:"sas"`
	// Whether emoji comparison is available.
	SupportsSas bool `json:"supportsSAS"`
	// Whether QR code verification is available.
	SupportsScanQrCode bool `json:"supportsScanQRCode"`
	// Verification ID to pass in verification action paths.
	VerificationID string `json:"verificationID"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AvailableActions   respjson.Field
		State              respjson.Field
		Error              respjson.Field
		From               respjson.Field
		FromDevice         respjson.Field
		OtherDevice        respjson.Field
		QrData             respjson.Field
		Sas                respjson.Field
		SupportsSas        respjson.Field
		SupportsScanQrCode respjson.Field
		VerificationID     respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppE2eeRecoveryCodeResetConfirmResponseAppStateVerification) RawJSON() string {
	return r.JSON.raw
}
func (r *AppE2eeRecoveryCodeResetConfirmResponseAppStateVerification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Verification error details, if verification stopped.
type AppE2eeRecoveryCodeResetConfirmResponseAppStateVerificationError struct {
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
func (r AppE2eeRecoveryCodeResetConfirmResponseAppStateVerificationError) RawJSON() string {
	return r.JSON.raw
}
func (r *AppE2eeRecoveryCodeResetConfirmResponseAppStateVerificationError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Emoji or number comparison data for verification.
type AppE2eeRecoveryCodeResetConfirmResponseAppStateVerificationSas struct {
	// Number sequence to compare on both devices.
	Decimals string `json:"decimals" api:"required"`
	// Emoji sequence to compare on both devices.
	Emojis string `json:"emojis" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Decimals    respjson.Field
		Emojis      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppE2eeRecoveryCodeResetConfirmResponseAppStateVerificationSas) RawJSON() string {
	return r.JSON.raw
}
func (r *AppE2eeRecoveryCodeResetConfirmResponseAppStateVerificationSas) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppE2eeRecoveryCodeResetNewParams struct {
	// Existing recovery key, if the user has it.
	RecoveryCode param.Opt[string] `json:"recoveryCode,omitzero"`
	paramObj
}

func (r AppE2eeRecoveryCodeResetNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AppE2eeRecoveryCodeResetNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppE2eeRecoveryCodeResetNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppE2eeRecoveryCodeResetConfirmParams struct {
	// New recovery key returned by the reset step.
	RecoveryCode string `json:"recoveryCode" api:"required"`
	paramObj
}

func (r AppE2eeRecoveryCodeResetConfirmParams) MarshalJSON() (data []byte, err error) {
	type shadow AppE2eeRecoveryCodeResetConfirmParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppE2eeRecoveryCodeResetConfirmParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
