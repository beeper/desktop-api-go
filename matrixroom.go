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

// MatrixRoomService contains methods and other services that help with interacting
// with the beeperdesktop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMatrixRoomService] method instead.
type MatrixRoomService struct {
	Options     []option.RequestOption
	AccountData MatrixRoomAccountDataService
	State       MatrixRoomStateService
	Events      MatrixRoomEventService
}

// NewMatrixRoomService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMatrixRoomService(opts ...option.RequestOption) (r MatrixRoomService) {
	r = MatrixRoomService{}
	r.Options = opts
	r.AccountData = NewMatrixRoomAccountDataService(opts...)
	r.State = NewMatrixRoomStateService(opts...)
	r.Events = NewMatrixRoomEventService(opts...)
	return
}

// Create a new room with various configuration options.
//
// The server MUST apply the normal state resolution rules when creating the new
// room, including checking power levels for each event. It MUST apply the events
// implied by the request in the following order:
//
// 1. The `m.room.create` event itself. Must be the first event in the room.
//
//  2. An `m.room.member` event for the creator to join the room. This is needed so
//     the remaining events can be sent.
//
//  3. A default `m.room.power_levels` event. Overridden by the
//     `power_level_content_override` parameter.
//
//     In [room versions](https://spec.matrix.org/v1.18/rooms) 1 through 11, the
//     room creator (and not other members) will be given permission to send state
//     events.
//
//     In room versions 12 and later, the room creator is given infinite power level
//     and cannot be specified in the `users` field of `m.room.power_levels`, so is
//     not listed explicitly.
//
//     **Note**: For `trusted_private_chat`, the users specified in the `invite`
//     parameter SHOULD also be appended to `additional_creators` by the server, per
//     the `creation_content` parameter.
//
//     If the room's version is 12 or higher, the power level for sending
//     `m.room.tombstone` events MUST explicitly be higher than `state_default`. For
//     example, set to 150 instead of 100.
//
// 4. An `m.room.canonical_alias` event if `room_alias_name` is given.
//
//  5. Events set by the `preset`. Currently these are the `m.room.join_rules`,
//     `m.room.history_visibility`, and `m.room.guest_access` state events.
//
// 6. Events listed in `initial_state`, in the order that they are listed.
//
//  7. Events implied by `name` and `topic` (`m.room.name` and `m.room.topic` state
//     events).
//
//  8. Invite events implied by `invite` and `invite_3pid` (`m.room.member` with
//     `membership: invite` and `m.room.third_party_invite`).
//
// The available presets do the following with respect to room state:
//
// | Preset                 | `join_rules` | `history_visibility` | `guest_access` | Other                                                            |
// | ---------------------- | ------------ | -------------------- | -------------- | ---------------------------------------------------------------- |
// | `private_chat`         | `invite`     | `shared`             | `can_join`     |                                                                  |
// | `trusted_private_chat` | `invite`     | `shared`             | `can_join`     | All invitees are given the same power level as the room creator. |
// | `public_chat`          | `public`     | `shared`             | `forbidden`    |                                                                  |
//
// The server will create a `m.room.create` event in the room with the requesting
// user as the creator, alongside other keys provided in the `creation_content` or
// implied by behaviour of `creation_content`.
func (r *MatrixRoomService) New(ctx context.Context, body MatrixRoomNewParams, opts ...option.RequestOption) (res *MatrixRoomNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "_matrix/client/v3/createRoom"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// _Note that this API takes either a room ID or alias, unlike_
// `/rooms/{roomId}/join`.
//
// This API starts a user's participation in a particular room, if that user is
// allowed to participate in that room. After this call, the client is allowed to
// see all current state events in the room, and all subsequent events associated
// with the room until the user leaves the room.
//
// After a user has joined a room, the room will appear as an entry in the response
// of the
// [`/initialSync`](https://spec.matrix.org/v1.18/client-server-api/#get_matrixclientv3initialsync)
// and
// [`/sync`](https://spec.matrix.org/v1.18/client-server-api/#get_matrixclientv3sync)
// APIs.
func (r *MatrixRoomService) Join(ctx context.Context, roomIDOrAlias string, params MatrixRoomJoinParams, opts ...option.RequestOption) (res *MatrixRoomJoinResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if roomIDOrAlias == "" {
		err = errors.New("missing required roomIdOrAlias parameter")
		return nil, err
	}
	path := fmt.Sprintf("_matrix/client/v3/join/%s", roomIDOrAlias)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// This API stops a user participating in a particular room.
//
// If the user was already in the room, they will no longer be able to see new
// events in the room. If the room requires an invite to join, they will need to be
// re-invited before they can re-join.
//
// If the user was invited to the room, but had not joined, this call serves to
// reject the invite.
//
// Servers MAY additionally forget the room when this endpoint is called – just as
// if the user had also invoked
// [`/forget`](https://spec.matrix.org/v1.18/client-server-api/#post_matrixclientv3roomsroomidforget).
// Servers that do this, MUST inform clients about this behavior using the
// [`m.forget_forced_upon_leave`](https://spec.matrix.org/v1.18/client-server-api/#mforget_forced_upon_leave-capability)
// capability.
//
// If the server doesn't automatically forget the room, the user will still be
// allowed to retrieve history from the room which they were previously allowed to
// see.
func (r *MatrixRoomService) Leave(ctx context.Context, roomID string, body MatrixRoomLeaveParams, opts ...option.RequestOption) (res *MatrixRoomLeaveResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if roomID == "" {
		err = errors.New("missing required roomId parameter")
		return nil, err
	}
	path := fmt.Sprintf("_matrix/client/v3/rooms/%s/leave", roomID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Information about the newly created room.
type MatrixRoomNewResponse struct {
	// The created room's ID.
	RoomID string `json:"room_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RoomID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MatrixRoomNewResponse) RawJSON() string { return r.JSON.raw }
func (r *MatrixRoomNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MatrixRoomJoinResponse struct {
	// The joined room ID.
	RoomID string `json:"room_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RoomID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MatrixRoomJoinResponse) RawJSON() string { return r.JSON.raw }
func (r *MatrixRoomJoinResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MatrixRoomLeaveResponse = any

type MatrixRoomNewParams struct {
	// This flag makes the server set the `is_direct` flag on the `m.room.member`
	// events sent to the users in `invite` and `invite_3pid`. See
	// [Direct Messaging](https://spec.matrix.org/v1.18/client-server-api/#direct-messaging)
	// for more information.
	IsDirect param.Opt[bool] `json:"is_direct,omitzero"`
	// If this is included, an
	// [`m.room.name`](https://spec.matrix.org/v1.18/client-server-api/#mroomname)
	// event will be sent into the room to indicate the name for the room. This
	// overwrites any
	// [`m.room.name`](https://spec.matrix.org/v1.18/client-server-api/#mroomname)
	// event in `initial_state`.
	Name param.Opt[string] `json:"name,omitzero"`
	// The desired room alias **local part**. If this is included, a room alias will be
	// created and mapped to the newly created room. The alias will belong on the
	// _same_ homeserver which created the room. For example, if this was set to "foo"
	// and sent to the homeserver "example.com" the complete room alias would be
	// `#foo:example.com`.
	//
	// The complete room alias will become the canonical alias for the room and an
	// `m.room.canonical_alias` event will be sent into the room.
	RoomAliasName param.Opt[string] `json:"room_alias_name,omitzero"`
	// The room version to set for the room. If not provided, the homeserver is to use
	// its configured default. If provided, the homeserver will return a 400 error with
	// the errcode `M_UNSUPPORTED_ROOM_VERSION` if it does not support the room
	// version.
	RoomVersion param.Opt[string] `json:"room_version,omitzero"`
	// If this is included, an
	// [`m.room.topic`](https://spec.matrix.org/v1.18/client-server-api/#mroomtopic)
	// event with a `text/plain` mimetype will be sent into the room to indicate the
	// topic for the room. This overwrites any
	// [`m.room.topic`](https://spec.matrix.org/v1.18/client-server-api/#mroomtopic)
	// event in `initial_state`.
	Topic param.Opt[string] `json:"topic,omitzero"`
	// Extra keys, such as `m.federate`, to be added to the content of the
	// [`m.room.create`](https://spec.matrix.org/v1.18/client-server-api/#mroomcreate)
	// event.
	//
	// The server will overwrite the following keys: `creator`, `room_version`. Future
	// versions of the specification may allow the server to overwrite other keys.
	//
	// When using the `trusted_private_chat` preset, the server SHOULD combine
	// `additional_creators` specified here and the `invite` array into the eventual
	// `m.room.create` event's `additional_creators`, deduplicating between the two
	// parameters.
	CreationContent any `json:"creation_content,omitzero"`
	// A list of state events to set in the new room. This allows the user to override
	// the default state events set in the new room. The expected format of the state
	// events are an object with type, state_key and content keys set.
	//
	// Takes precedence over events set by `preset`, but gets overridden by `name` and
	// `topic` keys.
	InitialState []MatrixRoomNewParamsInitialState `json:"initial_state,omitzero"`
	// A list of user IDs to invite to the room. This will tell the server to invite
	// everyone in the list to the newly created room.
	Invite []string `json:"invite,omitzero"`
	// A list of objects representing third-party IDs to invite into the room.
	Invite3pid []MatrixRoomNewParamsInvite3pid `json:"invite_3pid,omitzero"`
	// The power level content to override in the default power level event. This
	// object is applied on top of the generated
	// [`m.room.power_levels`](https://spec.matrix.org/v1.18/client-server-api/#mroompower_levels)
	// event content prior to it being sent to the room. Defaults to overriding
	// nothing.
	PowerLevelContentOverride any `json:"power_level_content_override,omitzero"`
	// Convenience parameter for setting various default state events based on a
	// preset.
	//
	// If unspecified, the server should use the `visibility` to determine which preset
	// to use. A visibility of `public` equates to a preset of `public_chat` and
	// `private` visibility equates to a preset of `private_chat`.
	//
	// Any of "private_chat", "public_chat", "trusted_private_chat".
	Preset MatrixRoomNewParamsPreset `json:"preset,omitzero"`
	// The room's visibility in the server's
	// [published room directory](https://spec.matrix.org/v1.18/client-server-api#published-room-directory).
	// Defaults to `private`.
	//
	// Any of "public", "private".
	Visibility MatrixRoomNewParamsVisibility `json:"visibility,omitzero"`
	paramObj
}

func (r MatrixRoomNewParams) MarshalJSON() (data []byte, err error) {
	type shadow MatrixRoomNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MatrixRoomNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Content, Type are required.
type MatrixRoomNewParamsInitialState struct {
	// The content of the event.
	Content any `json:"content,omitzero" api:"required"`
	// The type of event to send.
	Type string `json:"type" api:"required"`
	// The state_key of the state event. Defaults to an empty string.
	StateKey param.Opt[string] `json:"state_key,omitzero"`
	paramObj
}

func (r MatrixRoomNewParamsInitialState) MarshalJSON() (data []byte, err error) {
	type shadow MatrixRoomNewParamsInitialState
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MatrixRoomNewParamsInitialState) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Address, IDAccessToken, IDServer, Medium are required.
type MatrixRoomNewParamsInvite3pid struct {
	// The invitee's third-party identifier.
	Address string `json:"address" api:"required"`
	// An access token previously registered with the identity server. Servers can
	// treat this as optional to distinguish between r0.5-compatible clients and this
	// specification version.
	IDAccessToken string `json:"id_access_token" api:"required"`
	// The hostname+port of the identity server which should be used for third-party
	// identifier lookups.
	IDServer string `json:"id_server" api:"required"`
	// The kind of address being passed in the address field, for example `email` (see
	// [the list of recognised values](https://spec.matrix.org/v1.18/appendices/#3pid-types)).
	Medium string `json:"medium" api:"required"`
	paramObj
}

func (r MatrixRoomNewParamsInvite3pid) MarshalJSON() (data []byte, err error) {
	type shadow MatrixRoomNewParamsInvite3pid
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MatrixRoomNewParamsInvite3pid) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Convenience parameter for setting various default state events based on a
// preset.
//
// If unspecified, the server should use the `visibility` to determine which preset
// to use. A visibility of `public` equates to a preset of `public_chat` and
// `private` visibility equates to a preset of `private_chat`.
type MatrixRoomNewParamsPreset string

const (
	MatrixRoomNewParamsPresetPrivateChat        MatrixRoomNewParamsPreset = "private_chat"
	MatrixRoomNewParamsPresetPublicChat         MatrixRoomNewParamsPreset = "public_chat"
	MatrixRoomNewParamsPresetTrustedPrivateChat MatrixRoomNewParamsPreset = "trusted_private_chat"
)

// The room's visibility in the server's
// [published room directory](https://spec.matrix.org/v1.18/client-server-api#published-room-directory).
// Defaults to `private`.
type MatrixRoomNewParamsVisibility string

const (
	MatrixRoomNewParamsVisibilityPublic  MatrixRoomNewParamsVisibility = "public"
	MatrixRoomNewParamsVisibilityPrivate MatrixRoomNewParamsVisibility = "private"
)

type MatrixRoomJoinParams struct {
	// Optional reason to be included as the `reason` on the subsequent membership
	// event.
	Reason param.Opt[string] `json:"reason,omitzero"`
	// The servers to attempt to join the room through. One of the servers must be
	// participating in the room.
	Via []string `query:"via,omitzero" json:"-"`
	// A signature of an `m.third_party_invite` token to prove that this user owns a
	// third-party identity which has been invited to the room.
	ThirdPartySigned MatrixRoomJoinParamsThirdPartySigned `json:"third_party_signed,omitzero"`
	paramObj
}

func (r MatrixRoomJoinParams) MarshalJSON() (data []byte, err error) {
	type shadow MatrixRoomJoinParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MatrixRoomJoinParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [MatrixRoomJoinParams]'s query parameters as `url.Values`.
func (r MatrixRoomJoinParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// A signature of an `m.third_party_invite` token to prove that this user owns a
// third-party identity which has been invited to the room.
//
// The properties Token, Mxid, Sender, Signatures are required.
type MatrixRoomJoinParamsThirdPartySigned struct {
	// The state key of the m.third_party_invite event.
	Token string `json:"token" api:"required"`
	// The Matrix ID of the invitee.
	Mxid string `json:"mxid" api:"required"`
	// The Matrix ID of the user who issued the invite.
	Sender string `json:"sender" api:"required"`
	// A signatures object containing a signature of the entire signed object.
	Signatures map[string]map[string]string `json:"signatures,omitzero" api:"required"`
	paramObj
}

func (r MatrixRoomJoinParamsThirdPartySigned) MarshalJSON() (data []byte, err error) {
	type shadow MatrixRoomJoinParamsThirdPartySigned
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MatrixRoomJoinParamsThirdPartySigned) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MatrixRoomLeaveParams struct {
	// Optional reason to be included as the `reason` on the subsequent membership
	// event.
	Reason param.Opt[string] `json:"reason,omitzero"`
	paramObj
}

func (r MatrixRoomLeaveParams) MarshalJSON() (data []byte, err error) {
	type shadow MatrixRoomLeaveParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MatrixRoomLeaveParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
