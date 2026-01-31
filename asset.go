// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package beeperdesktopapi

import (
	"bytes"
	"context"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"slices"

	"github.com/beeper/desktop-api-go/internal/apiform"
	"github.com/beeper/desktop-api-go/internal/apijson"
	"github.com/beeper/desktop-api-go/internal/apiquery"
	"github.com/beeper/desktop-api-go/internal/requestconfig"
	"github.com/beeper/desktop-api-go/option"
	"github.com/beeper/desktop-api-go/packages/param"
	"github.com/beeper/desktop-api-go/packages/respjson"
)

// Manage assets in Beeper Desktop, like message attachments
//
// AssetService contains methods and other services that help with interacting with
// the beeperdesktop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAssetService] method instead.
type AssetService struct {
	Options []option.RequestOption
}

// NewAssetService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAssetService(opts ...option.RequestOption) (r AssetService) {
	r = AssetService{}
	r.Options = opts
	return
}

// Download a Matrix asset using its mxc:// or localmxc:// URL to the device
// running Beeper Desktop and return the local file URL.
func (r *AssetService) Download(ctx context.Context, body AssetDownloadParams, opts ...option.RequestOption) (res *AssetDownloadResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/assets/download"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Stream a file given an mxc://, localmxc://, or file:// URL. Downloads first if
// not cached. Supports Range requests for seeking in large files.
func (r *AssetService) Serve(ctx context.Context, query AssetServeParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "v1/assets/serve"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return
}

// Upload a file to a temporary location using multipart/form-data. Returns an
// uploadID that can be referenced when sending messages with attachments.
func (r *AssetService) Upload(ctx context.Context, body AssetUploadParams, opts ...option.RequestOption) (res *AssetUploadResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/assets/upload"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Upload a file using a JSON body with base64-encoded content. Returns an uploadID
// that can be referenced when sending messages with attachments. Alternative to
// the multipart upload endpoint.
func (r *AssetService) UploadBase64(ctx context.Context, body AssetUploadBase64Params, opts ...option.RequestOption) (res *AssetUploadBase64Response, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/assets/upload/base64"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AssetDownloadResponse struct {
	// Error message if the download failed.
	Error string `json:"error"`
	// Local file URL to the downloaded asset.
	SrcURL string `json:"srcURL"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		SrcURL      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssetDownloadResponse) RawJSON() string { return r.JSON.raw }
func (r *AssetDownloadResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssetUploadResponse struct {
	// Duration in seconds (audio/videos)
	Duration float64 `json:"duration"`
	// Error message if upload failed
	Error string `json:"error"`
	// Resolved filename
	FileName string `json:"fileName"`
	// File size in bytes
	FileSize float64 `json:"fileSize"`
	// Height in pixels (images/videos)
	Height float64 `json:"height"`
	// Detected or provided MIME type
	MimeType string `json:"mimeType"`
	// Local file URL (file://) for the uploaded asset
	SrcURL string `json:"srcURL"`
	// Unique upload ID for this asset
	UploadID string `json:"uploadID"`
	// Width in pixels (images/videos)
	Width float64 `json:"width"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Duration    respjson.Field
		Error       respjson.Field
		FileName    respjson.Field
		FileSize    respjson.Field
		Height      respjson.Field
		MimeType    respjson.Field
		SrcURL      respjson.Field
		UploadID    respjson.Field
		Width       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssetUploadResponse) RawJSON() string { return r.JSON.raw }
func (r *AssetUploadResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssetUploadBase64Response struct {
	// Duration in seconds (audio/videos)
	Duration float64 `json:"duration"`
	// Error message if upload failed
	Error string `json:"error"`
	// Resolved filename
	FileName string `json:"fileName"`
	// File size in bytes
	FileSize float64 `json:"fileSize"`
	// Height in pixels (images/videos)
	Height float64 `json:"height"`
	// Detected or provided MIME type
	MimeType string `json:"mimeType"`
	// Local file URL (file://) for the uploaded asset
	SrcURL string `json:"srcURL"`
	// Unique upload ID for this asset
	UploadID string `json:"uploadID"`
	// Width in pixels (images/videos)
	Width float64 `json:"width"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Duration    respjson.Field
		Error       respjson.Field
		FileName    respjson.Field
		FileSize    respjson.Field
		Height      respjson.Field
		MimeType    respjson.Field
		SrcURL      respjson.Field
		UploadID    respjson.Field
		Width       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssetUploadBase64Response) RawJSON() string { return r.JSON.raw }
func (r *AssetUploadBase64Response) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssetDownloadParams struct {
	// Matrix content URL (mxc:// or localmxc://) for the asset to download.
	URL string `json:"url,required"`
	paramObj
}

func (r AssetDownloadParams) MarshalJSON() (data []byte, err error) {
	type shadow AssetDownloadParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AssetDownloadParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssetServeParams struct {
	// Asset URL to serve. Accepts mxc://, localmxc://, or file:// URLs.
	URL string `query:"url,required" json:"-"`
	paramObj
}

// URLQuery serializes [AssetServeParams]'s query parameters as `url.Values`.
func (r AssetServeParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AssetUploadParams struct {
	// The file to upload (max 500 MB).
	File io.Reader `json:"file,omitzero,required" format:"binary"`
	// Original filename. Defaults to the uploaded file name if omitted
	FileName param.Opt[string] `json:"fileName,omitzero"`
	// MIME type. Auto-detected from magic bytes if omitted
	MimeType param.Opt[string] `json:"mimeType,omitzero"`
	paramObj
}

func (r AssetUploadParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err == nil {
		err = apiform.WriteExtras(writer, r.ExtraFields())
	}
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

type AssetUploadBase64Params struct {
	// Base64-encoded file content (max ~500MB decoded)
	Content string `json:"content,required"`
	// Original filename. Generated if omitted
	FileName param.Opt[string] `json:"fileName,omitzero"`
	// MIME type. Auto-detected from magic bytes if omitted
	MimeType param.Opt[string] `json:"mimeType,omitzero"`
	paramObj
}

func (r AssetUploadBase64Params) MarshalJSON() (data []byte, err error) {
	type shadow AssetUploadBase64Params
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AssetUploadBase64Params) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
