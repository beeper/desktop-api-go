// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package beeperdesktopapi

import (
	"github.com/beeper/desktop-api-go/v5/option"
)

// AppLoginVerificationRecoveryKeyResetService contains methods and other services
// that help with interacting with the beeperdesktop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAppLoginVerificationRecoveryKeyResetService] method instead.
type AppLoginVerificationRecoveryKeyResetService struct {
	Options []option.RequestOption
}

// NewAppLoginVerificationRecoveryKeyResetService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewAppLoginVerificationRecoveryKeyResetService(opts ...option.RequestOption) (r AppLoginVerificationRecoveryKeyResetService) {
	r = AppLoginVerificationRecoveryKeyResetService{}
	r.Options = opts
	return
}
