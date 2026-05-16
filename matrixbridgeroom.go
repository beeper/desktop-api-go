// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package beeperdesktopapi

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/beeper/desktop-api-go/v5/internal/apijson"
	"github.com/beeper/desktop-api-go/v5/internal/apiquery"
	"github.com/beeper/desktop-api-go/v5/internal/requestconfig"
	"github.com/beeper/desktop-api-go/v5/option"
	"github.com/beeper/desktop-api-go/v5/packages/param"
	"github.com/beeper/desktop-api-go/v5/packages/respjson"
)

// Matrix-compatible APIs for accounts and connected network bridges.
//
// MatrixBridgeRoomService contains methods and other services that help with
// interacting with the beeperdesktop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMatrixBridgeRoomService] method instead.
type MatrixBridgeRoomService struct {
	Options []option.RequestOption
}

// NewMatrixBridgeRoomService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMatrixBridgeRoomService(opts ...option.RequestOption) (r MatrixBridgeRoomService) {
	r = MatrixBridgeRoomService{}
	r.Options = opts
	return
}

// Create a direct chat with a user on the remote network.
func (r *MatrixBridgeRoomService) NewDm(ctx context.Context, identifier string, params MatrixBridgeRoomNewDmParams, opts ...option.RequestOption) (res *MatrixBridgeRoomNewDmResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.BridgeID == "" {
		err = errors.New("missing required bridgeID parameter")
		return nil, err
	}
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return nil, err
	}
	path := fmt.Sprintf("_matrix/client/unstable/com.beeper.bridge/%s/_matrix/provision/v3/create_dm/%s", params.BridgeID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Create a group chat on the remote network.
func (r *MatrixBridgeRoomService) NewGroup(ctx context.Context, groupType string, params MatrixBridgeRoomNewGroupParams, opts ...option.RequestOption) (res *MatrixBridgeRoomNewGroupResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.BridgeID == "" {
		err = errors.New("missing required bridgeID parameter")
		return nil, err
	}
	if groupType == "" {
		err = errors.New("missing required groupType parameter")
		return nil, err
	}
	path := fmt.Sprintf("_matrix/client/unstable/com.beeper.bridge/%s/_matrix/provision/v3/create_group/%s", params.BridgeID, groupType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// A successfully resolved identifier.
type MatrixBridgeRoomNewDmResponse struct {
	// The internal user ID of the resolved user.
	ID string `json:"id" api:"required"`
	// The avatar of the user on the remote network.
	AvatarURL string `json:"avatar_url" format:"mxc"`
	// The Matrix room ID of the direct chat with the user.
	DmRoomMxid string `json:"dm_room_mxid" format:"matrix_room_id"`
	// A list of identifiers for the user on the remote network.
	Identifiers []string `json:"identifiers" format:"uri"`
	// The Matrix user ID of the ghost representing the user.
	Mxid string `json:"mxid" format:"matrix_user_id"`
	// The name of the user on the remote network.
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		AvatarURL   respjson.Field
		DmRoomMxid  respjson.Field
		Identifiers respjson.Field
		Mxid        respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MatrixBridgeRoomNewDmResponse) RawJSON() string { return r.JSON.raw }
func (r *MatrixBridgeRoomNewDmResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A successfully created group chat.
type MatrixBridgeRoomNewGroupResponse struct {
	// The internal chat ID of the created group.
	ID string `json:"id" api:"required"`
	// The Matrix room ID of the portal.
	Mxid string `json:"mxid" api:"required" format:"matrix_room_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Mxid        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MatrixBridgeRoomNewGroupResponse) RawJSON() string { return r.JSON.raw }
func (r *MatrixBridgeRoomNewGroupResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MatrixBridgeRoomNewDmParams struct {
	BridgeID string `path:"bridgeID" api:"required" json:"-"`
	// An optional explicit login ID to do the action through.
	LoginID param.Opt[string] `query:"login_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MatrixBridgeRoomNewDmParams]'s query parameters as
// `url.Values`.
func (r MatrixBridgeRoomNewDmParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MatrixBridgeRoomNewGroupParams struct {
	BridgeID string `path:"bridgeID" api:"required" json:"-"`
	// An optional explicit login ID to do the action through.
	LoginID param.Opt[string] `query:"login_id,omitzero" json:"-"`
	// An existing Matrix room ID to bridge to. The other parameters must be already in
	// sync with the room state when using this parameter.
	RoomID param.Opt[string] `json:"room_id,omitzero" format:"matrix_room_id"`
	// The type of group to create.
	Type param.Opt[string] `json:"type,omitzero"`
	// The public username for the created group.
	Username param.Opt[string] `json:"username,omitzero"`
	// The `m.room.avatar` event content for the room.
	Avatar MatrixBridgeRoomNewGroupParamsAvatar `json:"avatar,omitzero"`
	// The `com.beeper.disappearing_timer` event content for the room.
	Disappear MatrixBridgeRoomNewGroupParamsDisappear `json:"disappear,omitzero"`
	// The `m.room.name` event content for the room.
	Name   MatrixBridgeRoomNewGroupParamsName `json:"name,omitzero"`
	Parent any                                `json:"parent,omitzero"`
	// The users to add to the group initially.
	Participants []string `json:"participants,omitzero"`
	// The `m.room.topic` event content for the room.
	Topic MatrixBridgeRoomNewGroupParamsTopic `json:"topic,omitzero"`
	paramObj
}

func (r MatrixBridgeRoomNewGroupParams) MarshalJSON() (data []byte, err error) {
	type shadow MatrixBridgeRoomNewGroupParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MatrixBridgeRoomNewGroupParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [MatrixBridgeRoomNewGroupParams]'s query parameters as
// `url.Values`.
func (r MatrixBridgeRoomNewGroupParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The `m.room.avatar` event content for the room.
type MatrixBridgeRoomNewGroupParamsAvatar struct {
	URL param.Opt[string] `json:"url,omitzero" format:"mxc"`
	paramObj
}

func (r MatrixBridgeRoomNewGroupParamsAvatar) MarshalJSON() (data []byte, err error) {
	type shadow MatrixBridgeRoomNewGroupParamsAvatar
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MatrixBridgeRoomNewGroupParamsAvatar) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The `com.beeper.disappearing_timer` event content for the room.
type MatrixBridgeRoomNewGroupParamsDisappear struct {
	Timer param.Opt[float64] `json:"timer,omitzero"`
	Type  param.Opt[string]  `json:"type,omitzero"`
	paramObj
}

func (r MatrixBridgeRoomNewGroupParamsDisappear) MarshalJSON() (data []byte, err error) {
	type shadow MatrixBridgeRoomNewGroupParamsDisappear
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MatrixBridgeRoomNewGroupParamsDisappear) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The `m.room.name` event content for the room.
type MatrixBridgeRoomNewGroupParamsName struct {
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r MatrixBridgeRoomNewGroupParamsName) MarshalJSON() (data []byte, err error) {
	type shadow MatrixBridgeRoomNewGroupParamsName
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MatrixBridgeRoomNewGroupParamsName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The `m.room.topic` event content for the room.
type MatrixBridgeRoomNewGroupParamsTopic struct {
	Topic param.Opt[string] `json:"topic,omitzero"`
	paramObj
}

func (r MatrixBridgeRoomNewGroupParamsTopic) MarshalJSON() (data []byte, err error) {
	type shadow MatrixBridgeRoomNewGroupParamsTopic
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MatrixBridgeRoomNewGroupParamsTopic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
