// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package beeperdesktopapi

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/beeper/desktop-api-go/v5/internal/requestconfig"
	"github.com/beeper/desktop-api-go/v5/option"
)

// Matrix-compatible APIs for accounts and connected network bridges.
//
// MatrixBridgeCapabilityService contains methods and other services that help with
// interacting with the beeperdesktop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMatrixBridgeCapabilityService] method instead.
type MatrixBridgeCapabilityService struct {
	Options []option.RequestOption
}

// NewMatrixBridgeCapabilityService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMatrixBridgeCapabilityService(opts ...option.RequestOption) (r MatrixBridgeCapabilityService) {
	r = MatrixBridgeCapabilityService{}
	r.Options = opts
	return
}

// Get bridge capabilities
func (r *MatrixBridgeCapabilityService) Get(ctx context.Context, bridgeID string, opts ...option.RequestOption) (res *MatrixBridgeCapabilityGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if bridgeID == "" {
		err = errors.New("missing required bridgeID parameter")
		return nil, err
	}
	path := fmt.Sprintf("_matrix/client/unstable/com.beeper.bridge/%s/_matrix/provision/v3/capabilities", bridgeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type MatrixBridgeCapabilityGetResponse map[string]any
