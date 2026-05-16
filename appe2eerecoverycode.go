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
// AppE2eeRecoveryCodeService contains methods and other services that help with
// interacting with the beeperdesktop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAppE2eeRecoveryCodeService] method instead.
type AppE2eeRecoveryCodeService struct {
	Options []option.RequestOption
	// First-party sign-in and encrypted messaging setup for Beeper Desktop.
	Reset AppE2eeRecoveryCodeResetService
}

// NewAppE2eeRecoveryCodeService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAppE2eeRecoveryCodeService(opts ...option.RequestOption) (r AppE2eeRecoveryCodeService) {
	r = AppE2eeRecoveryCodeService{}
	r.Options = opts
	r.Reset = NewAppE2eeRecoveryCodeResetService(opts...)
	return
}

// Record that the user saved their recovery key.
func (r *AppE2eeRecoveryCodeService) MarkBackedUp(ctx context.Context, opts ...option.RequestOption) (res *AppE2eeRecoveryCodeMarkBackedUpResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/app/e2ee/recovery-code/mark-backed-up"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Unlock encrypted messages with the user recovery key.
func (r *AppE2eeRecoveryCodeService) Verify(ctx context.Context, body AppE2eeRecoveryCodeVerifyParams, opts ...option.RequestOption) (res *AppE2eeRecoveryCodeVerifyResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/app/e2ee/recovery-code/verify"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type AppE2eeRecoveryCodeMarkBackedUpResponse struct {
	// Current onboarding state after the requested step.
	AppState AppE2eeRecoveryCodeMarkBackedUpResponseAppState `json:"appState" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AppState    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppE2eeRecoveryCodeMarkBackedUpResponse) RawJSON() string { return r.JSON.raw }
func (r *AppE2eeRecoveryCodeMarkBackedUpResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current onboarding state after the requested step.
type AppE2eeRecoveryCodeMarkBackedUpResponseAppState struct {
	// Encrypted messaging setup status.
	E2ee AppE2eeRecoveryCodeMarkBackedUpResponseAppStateE2ee `json:"e2ee" api:"required"`
	// Current onboarding state for Beeper Desktop.
	//
	// Any of "needs-login", "initializing", "needs-cross-signing-setup",
	// "needs-verification", "needs-secrets", "needs-first-sync", "ready".
	State string `json:"state" api:"required"`
	// Signed-in account details. Omitted until sign-in is complete.
	Matrix AppE2eeRecoveryCodeMarkBackedUpResponseAppStateMatrix `json:"matrix"`
	// Trusted-device verification progress.
	Verification AppE2eeRecoveryCodeMarkBackedUpResponseAppStateVerification `json:"verification"`
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
func (r AppE2eeRecoveryCodeMarkBackedUpResponseAppState) RawJSON() string { return r.JSON.raw }
func (r *AppE2eeRecoveryCodeMarkBackedUpResponseAppState) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Encrypted messaging setup status.
type AppE2eeRecoveryCodeMarkBackedUpResponseAppStateE2ee struct {
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
	Secrets AppE2eeRecoveryCodeMarkBackedUpResponseAppStateE2eeSecrets `json:"secrets" api:"required"`
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
func (r AppE2eeRecoveryCodeMarkBackedUpResponseAppStateE2ee) RawJSON() string { return r.JSON.raw }
func (r *AppE2eeRecoveryCodeMarkBackedUpResponseAppStateE2ee) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Encrypted messaging keys available on this device.
type AppE2eeRecoveryCodeMarkBackedUpResponseAppStateE2eeSecrets struct {
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
func (r AppE2eeRecoveryCodeMarkBackedUpResponseAppStateE2eeSecrets) RawJSON() string {
	return r.JSON.raw
}
func (r *AppE2eeRecoveryCodeMarkBackedUpResponseAppStateE2eeSecrets) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Signed-in account details. Omitted until sign-in is complete.
type AppE2eeRecoveryCodeMarkBackedUpResponseAppStateMatrix struct {
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
func (r AppE2eeRecoveryCodeMarkBackedUpResponseAppStateMatrix) RawJSON() string { return r.JSON.raw }
func (r *AppE2eeRecoveryCodeMarkBackedUpResponseAppStateMatrix) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Trusted-device verification progress.
type AppE2eeRecoveryCodeMarkBackedUpResponseAppStateVerification struct {
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
	Error AppE2eeRecoveryCodeMarkBackedUpResponseAppStateVerificationError `json:"error"`
	// User ID that started verification.
	From string `json:"from"`
	// Device that started verification.
	FromDevice string `json:"fromDevice"`
	// Other device participating in verification.
	OtherDevice string `json:"otherDevice"`
	// QR code payload to display for verification.
	QrData string `json:"qrData"`
	// Emoji or number comparison data for verification.
	Sas AppE2eeRecoveryCodeMarkBackedUpResponseAppStateVerificationSas `json:"sas"`
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
func (r AppE2eeRecoveryCodeMarkBackedUpResponseAppStateVerification) RawJSON() string {
	return r.JSON.raw
}
func (r *AppE2eeRecoveryCodeMarkBackedUpResponseAppStateVerification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Verification error details, if verification stopped.
type AppE2eeRecoveryCodeMarkBackedUpResponseAppStateVerificationError struct {
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
func (r AppE2eeRecoveryCodeMarkBackedUpResponseAppStateVerificationError) RawJSON() string {
	return r.JSON.raw
}
func (r *AppE2eeRecoveryCodeMarkBackedUpResponseAppStateVerificationError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Emoji or number comparison data for verification.
type AppE2eeRecoveryCodeMarkBackedUpResponseAppStateVerificationSas struct {
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
func (r AppE2eeRecoveryCodeMarkBackedUpResponseAppStateVerificationSas) RawJSON() string {
	return r.JSON.raw
}
func (r *AppE2eeRecoveryCodeMarkBackedUpResponseAppStateVerificationSas) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppE2eeRecoveryCodeVerifyResponse struct {
	// Current onboarding state after the requested step.
	AppState AppE2eeRecoveryCodeVerifyResponseAppState `json:"appState" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AppState    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppE2eeRecoveryCodeVerifyResponse) RawJSON() string { return r.JSON.raw }
func (r *AppE2eeRecoveryCodeVerifyResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current onboarding state after the requested step.
type AppE2eeRecoveryCodeVerifyResponseAppState struct {
	// Encrypted messaging setup status.
	E2ee AppE2eeRecoveryCodeVerifyResponseAppStateE2ee `json:"e2ee" api:"required"`
	// Current onboarding state for Beeper Desktop.
	//
	// Any of "needs-login", "initializing", "needs-cross-signing-setup",
	// "needs-verification", "needs-secrets", "needs-first-sync", "ready".
	State string `json:"state" api:"required"`
	// Signed-in account details. Omitted until sign-in is complete.
	Matrix AppE2eeRecoveryCodeVerifyResponseAppStateMatrix `json:"matrix"`
	// Trusted-device verification progress.
	Verification AppE2eeRecoveryCodeVerifyResponseAppStateVerification `json:"verification"`
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
func (r AppE2eeRecoveryCodeVerifyResponseAppState) RawJSON() string { return r.JSON.raw }
func (r *AppE2eeRecoveryCodeVerifyResponseAppState) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Encrypted messaging setup status.
type AppE2eeRecoveryCodeVerifyResponseAppStateE2ee struct {
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
	Secrets AppE2eeRecoveryCodeVerifyResponseAppStateE2eeSecrets `json:"secrets" api:"required"`
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
func (r AppE2eeRecoveryCodeVerifyResponseAppStateE2ee) RawJSON() string { return r.JSON.raw }
func (r *AppE2eeRecoveryCodeVerifyResponseAppStateE2ee) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Encrypted messaging keys available on this device.
type AppE2eeRecoveryCodeVerifyResponseAppStateE2eeSecrets struct {
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
func (r AppE2eeRecoveryCodeVerifyResponseAppStateE2eeSecrets) RawJSON() string { return r.JSON.raw }
func (r *AppE2eeRecoveryCodeVerifyResponseAppStateE2eeSecrets) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Signed-in account details. Omitted until sign-in is complete.
type AppE2eeRecoveryCodeVerifyResponseAppStateMatrix struct {
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
func (r AppE2eeRecoveryCodeVerifyResponseAppStateMatrix) RawJSON() string { return r.JSON.raw }
func (r *AppE2eeRecoveryCodeVerifyResponseAppStateMatrix) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Trusted-device verification progress.
type AppE2eeRecoveryCodeVerifyResponseAppStateVerification struct {
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
	Error AppE2eeRecoveryCodeVerifyResponseAppStateVerificationError `json:"error"`
	// User ID that started verification.
	From string `json:"from"`
	// Device that started verification.
	FromDevice string `json:"fromDevice"`
	// Other device participating in verification.
	OtherDevice string `json:"otherDevice"`
	// QR code payload to display for verification.
	QrData string `json:"qrData"`
	// Emoji or number comparison data for verification.
	Sas AppE2eeRecoveryCodeVerifyResponseAppStateVerificationSas `json:"sas"`
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
func (r AppE2eeRecoveryCodeVerifyResponseAppStateVerification) RawJSON() string { return r.JSON.raw }
func (r *AppE2eeRecoveryCodeVerifyResponseAppStateVerification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Verification error details, if verification stopped.
type AppE2eeRecoveryCodeVerifyResponseAppStateVerificationError struct {
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
func (r AppE2eeRecoveryCodeVerifyResponseAppStateVerificationError) RawJSON() string {
	return r.JSON.raw
}
func (r *AppE2eeRecoveryCodeVerifyResponseAppStateVerificationError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Emoji or number comparison data for verification.
type AppE2eeRecoveryCodeVerifyResponseAppStateVerificationSas struct {
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
func (r AppE2eeRecoveryCodeVerifyResponseAppStateVerificationSas) RawJSON() string { return r.JSON.raw }
func (r *AppE2eeRecoveryCodeVerifyResponseAppStateVerificationSas) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppE2eeRecoveryCodeVerifyParams struct {
	// Recovery key saved by the user.
	RecoveryCode string `json:"recoveryCode" api:"required"`
	paramObj
}

func (r AppE2eeRecoveryCodeVerifyParams) MarshalJSON() (data []byte, err error) {
	type shadow AppE2eeRecoveryCodeVerifyParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppE2eeRecoveryCodeVerifyParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
