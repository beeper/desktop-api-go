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
	"github.com/beeper/desktop-api-go/v5/packages/param"
	"github.com/beeper/desktop-api-go/v5/packages/respjson"
)

// First-party sign-in and encrypted messaging setup for Beeper Desktop.
//
// AppE2eeVerificationService contains methods and other services that help with
// interacting with the beeperdesktop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAppE2eeVerificationService] method instead.
type AppE2eeVerificationService struct {
	Options []option.RequestOption
	// First-party sign-in and encrypted messaging setup for Beeper Desktop.
	Qr AppE2eeVerificationQrService
	// First-party sign-in and encrypted messaging setup for Beeper Desktop.
	Sas AppE2eeVerificationSaService
}

// NewAppE2eeVerificationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAppE2eeVerificationService(opts ...option.RequestOption) (r AppE2eeVerificationService) {
	r = AppE2eeVerificationService{}
	r.Options = opts
	r.Qr = NewAppE2eeVerificationQrService(opts...)
	r.Sas = NewAppE2eeVerificationSaService(opts...)
	return
}

// Start verifying this device from another signed-in device.
func (r *AppE2eeVerificationService) New(ctx context.Context, body AppE2eeVerificationNewParams, opts ...option.RequestOption) (res *AppE2eeVerificationNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/app/e2ee/verification"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Accept an incoming device verification request.
func (r *AppE2eeVerificationService) Accept(ctx context.Context, verificationID string, opts ...option.RequestOption) (res *AppE2eeVerificationAcceptResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if verificationID == "" {
		err = errors.New("missing required verificationID parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/app/e2ee/verification/%s/accept", verificationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Cancel an active device verification request.
func (r *AppE2eeVerificationService) Cancel(ctx context.Context, verificationID string, body AppE2eeVerificationCancelParams, opts ...option.RequestOption) (res *AppE2eeVerificationCancelResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if verificationID == "" {
		err = errors.New("missing required verificationID parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/app/e2ee/verification/%s/cancel", verificationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type AppE2eeVerificationNewResponse struct {
	// Current onboarding state after starting verification.
	AppState AppE2eeVerificationNewResponseAppState `json:"appState" api:"required"`
	// Verification ID to pass in verification action paths.
	VerificationID string `json:"verificationID" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AppState       respjson.Field
		VerificationID respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppE2eeVerificationNewResponse) RawJSON() string { return r.JSON.raw }
func (r *AppE2eeVerificationNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current onboarding state after starting verification.
type AppE2eeVerificationNewResponseAppState struct {
	// Encrypted messaging setup status.
	E2ee AppE2eeVerificationNewResponseAppStateE2ee `json:"e2ee" api:"required"`
	// Current onboarding state for Beeper Desktop.
	//
	// Any of "needs-login", "initializing", "needs-cross-signing-setup",
	// "needs-verification", "needs-secrets", "needs-first-sync", "ready".
	State string `json:"state" api:"required"`
	// Signed-in account details. Omitted until sign-in is complete.
	Matrix AppE2eeVerificationNewResponseAppStateMatrix `json:"matrix"`
	// Trusted-device verification progress.
	Verification AppE2eeVerificationNewResponseAppStateVerification `json:"verification"`
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
func (r AppE2eeVerificationNewResponseAppState) RawJSON() string { return r.JSON.raw }
func (r *AppE2eeVerificationNewResponseAppState) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Encrypted messaging setup status.
type AppE2eeVerificationNewResponseAppStateE2ee struct {
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
	Secrets AppE2eeVerificationNewResponseAppStateE2eeSecrets `json:"secrets" api:"required"`
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
func (r AppE2eeVerificationNewResponseAppStateE2ee) RawJSON() string { return r.JSON.raw }
func (r *AppE2eeVerificationNewResponseAppStateE2ee) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Encrypted messaging keys available on this device.
type AppE2eeVerificationNewResponseAppStateE2eeSecrets struct {
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
func (r AppE2eeVerificationNewResponseAppStateE2eeSecrets) RawJSON() string { return r.JSON.raw }
func (r *AppE2eeVerificationNewResponseAppStateE2eeSecrets) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Signed-in account details. Omitted until sign-in is complete.
type AppE2eeVerificationNewResponseAppStateMatrix struct {
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
func (r AppE2eeVerificationNewResponseAppStateMatrix) RawJSON() string { return r.JSON.raw }
func (r *AppE2eeVerificationNewResponseAppStateMatrix) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Trusted-device verification progress.
type AppE2eeVerificationNewResponseAppStateVerification struct {
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
	Error AppE2eeVerificationNewResponseAppStateVerificationError `json:"error"`
	// User ID that started verification.
	From string `json:"from"`
	// Device that started verification.
	FromDevice string `json:"fromDevice"`
	// Other device participating in verification.
	OtherDevice string `json:"otherDevice"`
	// QR code payload to display for verification.
	QrData string `json:"qrData"`
	// Emoji or number comparison data for verification.
	Sas AppE2eeVerificationNewResponseAppStateVerificationSas `json:"sas"`
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
func (r AppE2eeVerificationNewResponseAppStateVerification) RawJSON() string { return r.JSON.raw }
func (r *AppE2eeVerificationNewResponseAppStateVerification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Verification error details, if verification stopped.
type AppE2eeVerificationNewResponseAppStateVerificationError struct {
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
func (r AppE2eeVerificationNewResponseAppStateVerificationError) RawJSON() string { return r.JSON.raw }
func (r *AppE2eeVerificationNewResponseAppStateVerificationError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Emoji or number comparison data for verification.
type AppE2eeVerificationNewResponseAppStateVerificationSas struct {
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
func (r AppE2eeVerificationNewResponseAppStateVerificationSas) RawJSON() string { return r.JSON.raw }
func (r *AppE2eeVerificationNewResponseAppStateVerificationSas) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppE2eeVerificationAcceptResponse struct {
	// Current onboarding state after the requested step.
	AppState AppE2eeVerificationAcceptResponseAppState `json:"appState" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AppState    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppE2eeVerificationAcceptResponse) RawJSON() string { return r.JSON.raw }
func (r *AppE2eeVerificationAcceptResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current onboarding state after the requested step.
type AppE2eeVerificationAcceptResponseAppState struct {
	// Encrypted messaging setup status.
	E2ee AppE2eeVerificationAcceptResponseAppStateE2ee `json:"e2ee" api:"required"`
	// Current onboarding state for Beeper Desktop.
	//
	// Any of "needs-login", "initializing", "needs-cross-signing-setup",
	// "needs-verification", "needs-secrets", "needs-first-sync", "ready".
	State string `json:"state" api:"required"`
	// Signed-in account details. Omitted until sign-in is complete.
	Matrix AppE2eeVerificationAcceptResponseAppStateMatrix `json:"matrix"`
	// Trusted-device verification progress.
	Verification AppE2eeVerificationAcceptResponseAppStateVerification `json:"verification"`
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
func (r AppE2eeVerificationAcceptResponseAppState) RawJSON() string { return r.JSON.raw }
func (r *AppE2eeVerificationAcceptResponseAppState) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Encrypted messaging setup status.
type AppE2eeVerificationAcceptResponseAppStateE2ee struct {
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
	Secrets AppE2eeVerificationAcceptResponseAppStateE2eeSecrets `json:"secrets" api:"required"`
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
func (r AppE2eeVerificationAcceptResponseAppStateE2ee) RawJSON() string { return r.JSON.raw }
func (r *AppE2eeVerificationAcceptResponseAppStateE2ee) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Encrypted messaging keys available on this device.
type AppE2eeVerificationAcceptResponseAppStateE2eeSecrets struct {
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
func (r AppE2eeVerificationAcceptResponseAppStateE2eeSecrets) RawJSON() string { return r.JSON.raw }
func (r *AppE2eeVerificationAcceptResponseAppStateE2eeSecrets) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Signed-in account details. Omitted until sign-in is complete.
type AppE2eeVerificationAcceptResponseAppStateMatrix struct {
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
func (r AppE2eeVerificationAcceptResponseAppStateMatrix) RawJSON() string { return r.JSON.raw }
func (r *AppE2eeVerificationAcceptResponseAppStateMatrix) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Trusted-device verification progress.
type AppE2eeVerificationAcceptResponseAppStateVerification struct {
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
	Error AppE2eeVerificationAcceptResponseAppStateVerificationError `json:"error"`
	// User ID that started verification.
	From string `json:"from"`
	// Device that started verification.
	FromDevice string `json:"fromDevice"`
	// Other device participating in verification.
	OtherDevice string `json:"otherDevice"`
	// QR code payload to display for verification.
	QrData string `json:"qrData"`
	// Emoji or number comparison data for verification.
	Sas AppE2eeVerificationAcceptResponseAppStateVerificationSas `json:"sas"`
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
func (r AppE2eeVerificationAcceptResponseAppStateVerification) RawJSON() string { return r.JSON.raw }
func (r *AppE2eeVerificationAcceptResponseAppStateVerification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Verification error details, if verification stopped.
type AppE2eeVerificationAcceptResponseAppStateVerificationError struct {
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
func (r AppE2eeVerificationAcceptResponseAppStateVerificationError) RawJSON() string {
	return r.JSON.raw
}
func (r *AppE2eeVerificationAcceptResponseAppStateVerificationError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Emoji or number comparison data for verification.
type AppE2eeVerificationAcceptResponseAppStateVerificationSas struct {
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
func (r AppE2eeVerificationAcceptResponseAppStateVerificationSas) RawJSON() string { return r.JSON.raw }
func (r *AppE2eeVerificationAcceptResponseAppStateVerificationSas) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppE2eeVerificationCancelResponse struct {
	// Current onboarding state after the requested step.
	AppState AppE2eeVerificationCancelResponseAppState `json:"appState" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AppState    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppE2eeVerificationCancelResponse) RawJSON() string { return r.JSON.raw }
func (r *AppE2eeVerificationCancelResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current onboarding state after the requested step.
type AppE2eeVerificationCancelResponseAppState struct {
	// Encrypted messaging setup status.
	E2ee AppE2eeVerificationCancelResponseAppStateE2ee `json:"e2ee" api:"required"`
	// Current onboarding state for Beeper Desktop.
	//
	// Any of "needs-login", "initializing", "needs-cross-signing-setup",
	// "needs-verification", "needs-secrets", "needs-first-sync", "ready".
	State string `json:"state" api:"required"`
	// Signed-in account details. Omitted until sign-in is complete.
	Matrix AppE2eeVerificationCancelResponseAppStateMatrix `json:"matrix"`
	// Trusted-device verification progress.
	Verification AppE2eeVerificationCancelResponseAppStateVerification `json:"verification"`
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
func (r AppE2eeVerificationCancelResponseAppState) RawJSON() string { return r.JSON.raw }
func (r *AppE2eeVerificationCancelResponseAppState) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Encrypted messaging setup status.
type AppE2eeVerificationCancelResponseAppStateE2ee struct {
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
	Secrets AppE2eeVerificationCancelResponseAppStateE2eeSecrets `json:"secrets" api:"required"`
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
func (r AppE2eeVerificationCancelResponseAppStateE2ee) RawJSON() string { return r.JSON.raw }
func (r *AppE2eeVerificationCancelResponseAppStateE2ee) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Encrypted messaging keys available on this device.
type AppE2eeVerificationCancelResponseAppStateE2eeSecrets struct {
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
func (r AppE2eeVerificationCancelResponseAppStateE2eeSecrets) RawJSON() string { return r.JSON.raw }
func (r *AppE2eeVerificationCancelResponseAppStateE2eeSecrets) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Signed-in account details. Omitted until sign-in is complete.
type AppE2eeVerificationCancelResponseAppStateMatrix struct {
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
func (r AppE2eeVerificationCancelResponseAppStateMatrix) RawJSON() string { return r.JSON.raw }
func (r *AppE2eeVerificationCancelResponseAppStateMatrix) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Trusted-device verification progress.
type AppE2eeVerificationCancelResponseAppStateVerification struct {
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
	Error AppE2eeVerificationCancelResponseAppStateVerificationError `json:"error"`
	// User ID that started verification.
	From string `json:"from"`
	// Device that started verification.
	FromDevice string `json:"fromDevice"`
	// Other device participating in verification.
	OtherDevice string `json:"otherDevice"`
	// QR code payload to display for verification.
	QrData string `json:"qrData"`
	// Emoji or number comparison data for verification.
	Sas AppE2eeVerificationCancelResponseAppStateVerificationSas `json:"sas"`
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
func (r AppE2eeVerificationCancelResponseAppStateVerification) RawJSON() string { return r.JSON.raw }
func (r *AppE2eeVerificationCancelResponseAppStateVerification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Verification error details, if verification stopped.
type AppE2eeVerificationCancelResponseAppStateVerificationError struct {
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
func (r AppE2eeVerificationCancelResponseAppStateVerificationError) RawJSON() string {
	return r.JSON.raw
}
func (r *AppE2eeVerificationCancelResponseAppStateVerificationError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Emoji or number comparison data for verification.
type AppE2eeVerificationCancelResponseAppStateVerificationSas struct {
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
func (r AppE2eeVerificationCancelResponseAppStateVerificationSas) RawJSON() string { return r.JSON.raw }
func (r *AppE2eeVerificationCancelResponseAppStateVerificationSas) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppE2eeVerificationNewParams struct {
	// User ID to verify. Defaults to the signed-in user.
	UserID param.Opt[string] `json:"userID,omitzero"`
	paramObj
}

func (r AppE2eeVerificationNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AppE2eeVerificationNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppE2eeVerificationNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppE2eeVerificationCancelParams struct {
	// Optional cancellation code.
	Code param.Opt[string] `json:"code,omitzero"`
	// Optional user-facing cancellation reason.
	Reason param.Opt[string] `json:"reason,omitzero"`
	paramObj
}

func (r AppE2eeVerificationCancelParams) MarshalJSON() (data []byte, err error) {
	type shadow AppE2eeVerificationCancelParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppE2eeVerificationCancelParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
