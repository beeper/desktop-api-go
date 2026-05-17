// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package beeperdesktopapi

import (
	"github.com/beeper/desktop-api-go/v5/option"
)

// Complete first-party Beeper app login
//
// AppLoginService contains methods and other services that help with interacting
// with the beeperdesktop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAppLoginService] method instead.
type AppLoginService struct {
	Options      []option.RequestOption
	Verification AppLoginVerificationService
}

// NewAppLoginService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAppLoginService(opts ...option.RequestOption) (r AppLoginService) {
	r = AppLoginService{}
	r.Options = opts
	r.Verification = NewAppLoginVerificationService(opts...)
	return
}
