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
// AppE2eeVerificationQrService contains methods and other services that help with
// interacting with the beeperdesktop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAppE2eeVerificationQrService] method instead.
type AppE2eeVerificationQrService struct {
	Options []option.RequestOption
}

// NewAppE2eeVerificationQrService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAppE2eeVerificationQrService(opts ...option.RequestOption) (r AppE2eeVerificationQrService) {
	r = AppE2eeVerificationQrService{}
	r.Options = opts
	return
}

// Confirm that another device scanned this device QR code.
func (r *AppE2eeVerificationQrService) ConfirmScanned(ctx context.Context, verificationID string, opts ...option.RequestOption) (res *AppE2eeVerificationQrConfirmScannedResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if verificationID == "" {
		err = errors.New("missing required verificationID parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/app/e2ee/verification/%s/qr/confirm-scanned", verificationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Submit the QR code scanned from another signed-in device.
func (r *AppE2eeVerificationQrService) Scan(ctx context.Context, body AppE2eeVerificationQrScanParams, opts ...option.RequestOption) (res *AppE2eeVerificationQrScanResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/app/e2ee/verification/qr/scan"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type AppE2eeVerificationQrConfirmScannedResponse struct {
	// Current onboarding state after the requested step.
	AppState AppE2eeVerificationQrConfirmScannedResponseAppState `json:"appState" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AppState    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppE2eeVerificationQrConfirmScannedResponse) RawJSON() string { return r.JSON.raw }
func (r *AppE2eeVerificationQrConfirmScannedResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current onboarding state after the requested step.
type AppE2eeVerificationQrConfirmScannedResponseAppState struct {
	// Encrypted messaging setup status.
	E2ee AppE2eeVerificationQrConfirmScannedResponseAppStateE2ee `json:"e2ee" api:"required"`
	// Current onboarding state for Beeper Desktop.
	//
	// Any of "needs-login", "initializing", "needs-cross-signing-setup",
	// "needs-verification", "needs-secrets", "needs-first-sync", "ready".
	State string `json:"state" api:"required"`
	// Signed-in account details. Omitted until sign-in is complete.
	Matrix AppE2eeVerificationQrConfirmScannedResponseAppStateMatrix `json:"matrix"`
	// Trusted-device verification progress.
	Verification AppE2eeVerificationQrConfirmScannedResponseAppStateVerification `json:"verification"`
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
func (r AppE2eeVerificationQrConfirmScannedResponseAppState) RawJSON() string { return r.JSON.raw }
func (r *AppE2eeVerificationQrConfirmScannedResponseAppState) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Encrypted messaging setup status.
type AppE2eeVerificationQrConfirmScannedResponseAppStateE2ee struct {
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
	Secrets AppE2eeVerificationQrConfirmScannedResponseAppStateE2eeSecrets `json:"secrets" api:"required"`
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
func (r AppE2eeVerificationQrConfirmScannedResponseAppStateE2ee) RawJSON() string { return r.JSON.raw }
func (r *AppE2eeVerificationQrConfirmScannedResponseAppStateE2ee) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Encrypted messaging keys available on this device.
type AppE2eeVerificationQrConfirmScannedResponseAppStateE2eeSecrets struct {
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
func (r AppE2eeVerificationQrConfirmScannedResponseAppStateE2eeSecrets) RawJSON() string {
	return r.JSON.raw
}
func (r *AppE2eeVerificationQrConfirmScannedResponseAppStateE2eeSecrets) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Signed-in account details. Omitted until sign-in is complete.
type AppE2eeVerificationQrConfirmScannedResponseAppStateMatrix struct {
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
func (r AppE2eeVerificationQrConfirmScannedResponseAppStateMatrix) RawJSON() string {
	return r.JSON.raw
}
func (r *AppE2eeVerificationQrConfirmScannedResponseAppStateMatrix) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Trusted-device verification progress.
type AppE2eeVerificationQrConfirmScannedResponseAppStateVerification struct {
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
	Error AppE2eeVerificationQrConfirmScannedResponseAppStateVerificationError `json:"error"`
	// User ID that started verification.
	From string `json:"from"`
	// Device that started verification.
	FromDevice string `json:"fromDevice"`
	// Other device participating in verification.
	OtherDevice string `json:"otherDevice"`
	// QR code payload to display for verification.
	QrData string `json:"qrData"`
	// Emoji or number comparison data for verification.
	Sas AppE2eeVerificationQrConfirmScannedResponseAppStateVerificationSas `json:"sas"`
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
func (r AppE2eeVerificationQrConfirmScannedResponseAppStateVerification) RawJSON() string {
	return r.JSON.raw
}
func (r *AppE2eeVerificationQrConfirmScannedResponseAppStateVerification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Verification error details, if verification stopped.
type AppE2eeVerificationQrConfirmScannedResponseAppStateVerificationError struct {
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
func (r AppE2eeVerificationQrConfirmScannedResponseAppStateVerificationError) RawJSON() string {
	return r.JSON.raw
}
func (r *AppE2eeVerificationQrConfirmScannedResponseAppStateVerificationError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Emoji or number comparison data for verification.
type AppE2eeVerificationQrConfirmScannedResponseAppStateVerificationSas struct {
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
func (r AppE2eeVerificationQrConfirmScannedResponseAppStateVerificationSas) RawJSON() string {
	return r.JSON.raw
}
func (r *AppE2eeVerificationQrConfirmScannedResponseAppStateVerificationSas) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppE2eeVerificationQrScanResponse struct {
	// Current onboarding state after the requested step.
	AppState AppE2eeVerificationQrScanResponseAppState `json:"appState" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AppState    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppE2eeVerificationQrScanResponse) RawJSON() string { return r.JSON.raw }
func (r *AppE2eeVerificationQrScanResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current onboarding state after the requested step.
type AppE2eeVerificationQrScanResponseAppState struct {
	// Encrypted messaging setup status.
	E2ee AppE2eeVerificationQrScanResponseAppStateE2ee `json:"e2ee" api:"required"`
	// Current onboarding state for Beeper Desktop.
	//
	// Any of "needs-login", "initializing", "needs-cross-signing-setup",
	// "needs-verification", "needs-secrets", "needs-first-sync", "ready".
	State string `json:"state" api:"required"`
	// Signed-in account details. Omitted until sign-in is complete.
	Matrix AppE2eeVerificationQrScanResponseAppStateMatrix `json:"matrix"`
	// Trusted-device verification progress.
	Verification AppE2eeVerificationQrScanResponseAppStateVerification `json:"verification"`
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
func (r AppE2eeVerificationQrScanResponseAppState) RawJSON() string { return r.JSON.raw }
func (r *AppE2eeVerificationQrScanResponseAppState) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Encrypted messaging setup status.
type AppE2eeVerificationQrScanResponseAppStateE2ee struct {
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
	Secrets AppE2eeVerificationQrScanResponseAppStateE2eeSecrets `json:"secrets" api:"required"`
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
func (r AppE2eeVerificationQrScanResponseAppStateE2ee) RawJSON() string { return r.JSON.raw }
func (r *AppE2eeVerificationQrScanResponseAppStateE2ee) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Encrypted messaging keys available on this device.
type AppE2eeVerificationQrScanResponseAppStateE2eeSecrets struct {
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
func (r AppE2eeVerificationQrScanResponseAppStateE2eeSecrets) RawJSON() string { return r.JSON.raw }
func (r *AppE2eeVerificationQrScanResponseAppStateE2eeSecrets) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Signed-in account details. Omitted until sign-in is complete.
type AppE2eeVerificationQrScanResponseAppStateMatrix struct {
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
func (r AppE2eeVerificationQrScanResponseAppStateMatrix) RawJSON() string { return r.JSON.raw }
func (r *AppE2eeVerificationQrScanResponseAppStateMatrix) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Trusted-device verification progress.
type AppE2eeVerificationQrScanResponseAppStateVerification struct {
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
	Error AppE2eeVerificationQrScanResponseAppStateVerificationError `json:"error"`
	// User ID that started verification.
	From string `json:"from"`
	// Device that started verification.
	FromDevice string `json:"fromDevice"`
	// Other device participating in verification.
	OtherDevice string `json:"otherDevice"`
	// QR code payload to display for verification.
	QrData string `json:"qrData"`
	// Emoji or number comparison data for verification.
	Sas AppE2eeVerificationQrScanResponseAppStateVerificationSas `json:"sas"`
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
func (r AppE2eeVerificationQrScanResponseAppStateVerification) RawJSON() string { return r.JSON.raw }
func (r *AppE2eeVerificationQrScanResponseAppStateVerification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Verification error details, if verification stopped.
type AppE2eeVerificationQrScanResponseAppStateVerificationError struct {
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
func (r AppE2eeVerificationQrScanResponseAppStateVerificationError) RawJSON() string {
	return r.JSON.raw
}
func (r *AppE2eeVerificationQrScanResponseAppStateVerificationError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Emoji or number comparison data for verification.
type AppE2eeVerificationQrScanResponseAppStateVerificationSas struct {
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
func (r AppE2eeVerificationQrScanResponseAppStateVerificationSas) RawJSON() string { return r.JSON.raw }
func (r *AppE2eeVerificationQrScanResponseAppStateVerificationSas) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppE2eeVerificationQrScanParams struct {
	// QR code payload scanned from the other device.
	Data string `json:"data" api:"required"`
	paramObj
}

func (r AppE2eeVerificationQrScanParams) MarshalJSON() (data []byte, err error) {
	type shadow AppE2eeVerificationQrScanParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppE2eeVerificationQrScanParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
