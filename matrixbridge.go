// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package beeperdesktopapi

import (
	"github.com/beeper/desktop-api-go/v5/option"
)

// Matrix-compatible APIs for connected network bridges.
//
// MatrixBridgeService contains methods and other services that help with
// interacting with the beeperdesktop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMatrixBridgeService] method instead.
type MatrixBridgeService struct {
	Options []option.RequestOption
	// Matrix-compatible APIs for accounts and connected network bridges.
	Auth MatrixBridgeAuthService
	// Matrix-compatible APIs for accounts and connected network bridges.
	Contacts MatrixBridgeContactService
	// Matrix-compatible APIs for accounts and connected network bridges.
	Users MatrixBridgeUserService
	// Matrix-compatible APIs for accounts and connected network bridges.
	Rooms MatrixBridgeRoomService
	// Matrix-compatible APIs for accounts and connected network bridges.
	Capabilities MatrixBridgeCapabilityService
}

// NewMatrixBridgeService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMatrixBridgeService(opts ...option.RequestOption) (r MatrixBridgeService) {
	r = MatrixBridgeService{}
	r.Options = opts
	r.Auth = NewMatrixBridgeAuthService(opts...)
	r.Contacts = NewMatrixBridgeContactService(opts...)
	r.Users = NewMatrixBridgeUserService(opts...)
	r.Rooms = NewMatrixBridgeRoomService(opts...)
	r.Capabilities = NewMatrixBridgeCapabilityService(opts...)
	return
}
