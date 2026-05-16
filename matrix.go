// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package beeperdesktopapi

import (
	"github.com/beeper/desktop-api-go/v5/option"
)

// Matrix-compatible APIs for accounts, rooms, and connected network bridges.
//
// MatrixService contains methods and other services that help with interacting
// with the beeperdesktop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMatrixService] method instead.
type MatrixService struct {
	Options []option.RequestOption
	Users   MatrixUserService
	Rooms   MatrixRoomService
	// Matrix-compatible APIs for connected network bridges.
	Bridges MatrixBridgeService
}

// NewMatrixService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMatrixService(opts ...option.RequestOption) (r MatrixService) {
	r = MatrixService{}
	r.Options = opts
	r.Users = NewMatrixUserService(opts...)
	r.Rooms = NewMatrixRoomService(opts...)
	r.Bridges = NewMatrixBridgeService(opts...)
	return
}
