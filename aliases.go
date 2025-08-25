// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcombeeperdesktopapigo

import (
	"github.com/beeper/desktop-api-go/internal/apierror"
	"github.com/beeper/desktop-api-go/packages/param"
	"github.com/beeper/desktop-api-go/shared"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type Error = apierror.Error

// A chat account added to Beeper
//
// This is an alias to an internal type.
type Account = shared.Account

// This is an alias to an internal type.
type Attachment = shared.Attachment

// Attachment type.
//
// This is an alias to an internal type.
type AttachmentType = shared.AttachmentType

// Equals "unknown"
const AttachmentTypeUnknown = shared.AttachmentTypeUnknown

// Equals "img"
const AttachmentTypeImg = shared.AttachmentTypeImg

// Equals "video"
const AttachmentTypeVideo = shared.AttachmentTypeVideo

// Equals "audio"
const AttachmentTypeAudio = shared.AttachmentTypeAudio

// Pixel dimensions of the attachment: width/height in px.
//
// This is an alias to an internal type.
type AttachmentSize = shared.AttachmentSize

// This is an alias to an internal type.
type BaseResponse = shared.BaseResponse

// This is an alias to an internal type.
type Chat = shared.Chat

// Chat participants information.
//
// This is an alias to an internal type.
type ChatParticipants = shared.ChatParticipants

// Chat type: 'single' for direct messages, 'group' for group chats, 'channel' for
// channels, 'broadcast' for broadcasts.
//
// This is an alias to an internal type.
type ChatType = shared.ChatType

// Equals "single"
const ChatTypeSingle = shared.ChatTypeSingle

// Equals "group"
const ChatTypeGroup = shared.ChatTypeGroup

// Equals "channel"
const ChatTypeChannel = shared.ChatTypeChannel

// Equals "broadcast"
const ChatTypeBroadcast = shared.ChatTypeBroadcast

// Last read message sortKey (hsOrder). Used to compute 'isUnread'.
//
// This is an alias to an internal type.
type ChatLastReadMessageSortKeyUnion = shared.ChatLastReadMessageSortKeyUnion

// This is an alias to an internal type.
type Message = shared.Message

// A unique key used to sort messages
//
// This is an alias to an internal type.
type MessageSortKeyUnion = shared.MessageSortKeyUnion

// This is an alias to an internal type.
type Reaction = shared.Reaction

// A person on or reachable through Beeper. Values are best-effort and can vary by
// network.
//
// This is an alias to an internal type.
type User = shared.User
