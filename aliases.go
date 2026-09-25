// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package beeperdesktopapi

import (
	"github.com/beeper/desktop-api-go/v6/internal/apierror"
	"github.com/beeper/desktop-api-go/v6/packages/param"
	"github.com/beeper/desktop-api-go/v6/shared"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type Error = apierror.Error

// This is an alias to an internal type.
type APIError = shared.APIError

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

// Attachment transcription if available.
//
// This is an alias to an internal type.
type AttachmentTranscription = shared.AttachmentTranscription

// Capabilities for one attachment message type.
//
// This is an alias to an internal type.
type AttachmentCapabilities = shared.AttachmentCapabilities

// Chat capabilities reported by the platform.
//
// This is an alias to an internal type.
type ChatCapabilities = shared.ChatCapabilities

// Disappearing-message timer capabilities.
//
// This is an alias to an internal type.
type ChatCapabilitiesDisappearingTimer = shared.ChatCapabilitiesDisappearingTimer

// Message request capabilities.
//
// This is an alias to an internal type.
type ChatCapabilitiesMessageRequest = shared.ChatCapabilitiesMessageRequest

// Participant management capabilities.
//
// This is an alias to an internal type.
type ChatCapabilitiesParticipantActions = shared.ChatCapabilitiesParticipantActions

// Current draft object for this chat, or null when no draft is set.
//
// This is an alias to an internal type.
type ChatDraft = shared.ChatDraft

// Chat state update capabilities.
//
// This is an alias to an internal type.
type ChatStateCapabilities = shared.ChatStateCapabilities

// Chat avatar state capability.
//
// This is an alias to an internal type.
type ChatStateCapabilitiesAvatar = shared.ChatStateCapabilitiesAvatar

// Chat description/topic state capability.
//
// This is an alias to an internal type.
type ChatStateCapabilitiesDescription = shared.ChatStateCapabilitiesDescription

// Disappearing-message timer state capability.
//
// This is an alias to an internal type.
type ChatStateCapabilitiesDisappearingTimer = shared.ChatStateCapabilitiesDisappearingTimer

// Chat title state capability.
//
// This is an alias to an internal type.
type ChatStateCapabilitiesTitle = shared.ChatStateCapabilitiesTitle

// This is an alias to an internal type.
type DraftAttachment = shared.DraftAttachment

// Draft attachment type. GIF and recorded audio are mutually exclusive types.
//
// This is an alias to an internal type.
type DraftAttachmentType = shared.DraftAttachmentType

// Equals "file"
const DraftAttachmentTypeFile = shared.DraftAttachmentTypeFile

// Equals "gif"
const DraftAttachmentTypeGif = shared.DraftAttachmentTypeGif

// Equals "recorded_audio"
const DraftAttachmentTypeRecordedAudio = shared.DraftAttachmentTypeRecordedAudio

// Pixel dimensions of the attachment.
//
// This is an alias to an internal type.
type DraftAttachmentSize = shared.DraftAttachmentSize

// Link preview included with a message.
//
// This is an alias to an internal type.
type LinkPreview = shared.LinkPreview

// Preview image dimensions.
//
// This is an alias to an internal type.
type LinkPreviewImgSize = shared.LinkPreviewImgSize

// This is an alias to an internal type.
type Message = shared.Message

// Read receipt state for this message, when available.
//
// This is an alias to an internal type.
type MessageSeenUnion = shared.MessageSeenUnion

// ISO 8601 timestamp.
//
// This is an alias to an internal type.
type MessageSeenByParticipantItemUnion = shared.MessageSeenByParticipantItemUnion

// Message content type. Useful for distinguishing reactions, media messages, and
// state events from regular text messages.
//
// This is an alias to an internal type.
type MessageType = shared.MessageType

// Equals "TEXT"
const MessageTypeText = shared.MessageTypeText

// Equals "NOTICE"
const MessageTypeNotice = shared.MessageTypeNotice

// Equals "IMAGE"
const MessageTypeImage = shared.MessageTypeImage

// Equals "VIDEO"
const MessageTypeVideo = shared.MessageTypeVideo

// Equals "VOICE"
const MessageTypeVoice = shared.MessageTypeVoice

// Equals "AUDIO"
const MessageTypeAudio = shared.MessageTypeAudio

// Equals "FILE"
const MessageTypeFile = shared.MessageTypeFile

// Equals "STICKER"
const MessageTypeSticker = shared.MessageTypeSticker

// Equals "LOCATION"
const MessageTypeLocation = shared.MessageTypeLocation

// Equals "REACTION"
const MessageTypeReaction = shared.MessageTypeReaction

// This is an alias to an internal type.
type Reaction = shared.Reaction

// Message send status for this message, when reported by the bridge.
//
// This is an alias to an internal type.
type SendStatus = shared.SendStatus

// Current status of the message send attempt.
//
// This is an alias to an internal type.
type SendStatusStatus = shared.SendStatusStatus

// Equals "SUCCESS"
const SendStatusStatusSuccess = shared.SendStatusStatusSuccess

// Equals "PENDING"
const SendStatusStatusPending = shared.SendStatusStatusPending

// Equals "FAIL_RETRIABLE"
const SendStatusStatusFailRetriable = shared.SendStatusStatusFailRetriable

// Equals "FAIL_PERMANENT"
const SendStatusStatusFailPermanent = shared.SendStatusStatusFailPermanent

// User the account belongs to.
//
// This is an alias to an internal type.
type User = shared.User
