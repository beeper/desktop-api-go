// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pagination

import (
	"net/http"

	"github.com/beeper/desktop-api-go/internal/apijson"
	"github.com/beeper/desktop-api-go/internal/requestconfig"
	"github.com/beeper/desktop-api-go/option"
	"github.com/beeper/desktop-api-go/packages/param"
	"github.com/beeper/desktop-api-go/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type CursorSearch[T any] struct {
	Items        []T    `json:"items"`
	HasMore      bool   `json:"hasMore"`
	OldestCursor string `json:"oldestCursor,nullable"`
	NewestCursor string `json:"newestCursor,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items        respjson.Field
		HasMore      respjson.Field
		OldestCursor respjson.Field
		NewestCursor respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r CursorSearch[T]) RawJSON() string { return r.JSON.raw }
func (r *CursorSearch[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *CursorSearch[T]) GetNextPage() (res *CursorSearch[T], err error) {
	if len(r.Items) == 0 {
		return nil, nil
	}

	if r.JSON.HasMore.Valid() && r.HasMore == false {
		return nil, nil
	}
	next := r.OldestCursor
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("cursor", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *CursorSearch[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &CursorSearch[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type CursorSearchAutoPager[T any] struct {
	page *CursorSearch[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewCursorSearchAutoPager[T any](page *CursorSearch[T], err error) *CursorSearchAutoPager[T] {
	return &CursorSearchAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *CursorSearchAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Items) == 0 {
		return false
	}
	if r.idx >= len(r.page.Items) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Items) == 0 {
			return false
		}
	}
	r.cur = r.page.Items[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *CursorSearchAutoPager[T]) Current() T {
	return r.cur
}

func (r *CursorSearchAutoPager[T]) Err() error {
	return r.err
}

func (r *CursorSearchAutoPager[T]) Index() int {
	return r.run
}

type CursorList[T any] struct {
	Items        []T    `json:"items"`
	HasMore      bool   `json:"hasMore"`
	OldestCursor string `json:"oldestCursor,nullable"`
	NewestCursor string `json:"newestCursor,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items        respjson.Field
		HasMore      respjson.Field
		OldestCursor respjson.Field
		NewestCursor respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r CursorList[T]) RawJSON() string { return r.JSON.raw }
func (r *CursorList[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *CursorList[T]) GetNextPage() (res *CursorList[T], err error) {
	if len(r.Items) == 0 {
		return nil, nil
	}

	if r.JSON.HasMore.Valid() && r.HasMore == false {
		return nil, nil
	}
	next := r.OldestCursor
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("cursor", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *CursorList[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &CursorList[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type CursorListAutoPager[T any] struct {
	page *CursorList[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewCursorListAutoPager[T any](page *CursorList[T], err error) *CursorListAutoPager[T] {
	return &CursorListAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *CursorListAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Items) == 0 {
		return false
	}
	if r.idx >= len(r.page.Items) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Items) == 0 {
			return false
		}
	}
	r.cur = r.page.Items[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *CursorListAutoPager[T]) Current() T {
	return r.cur
}

func (r *CursorListAutoPager[T]) Err() error {
	return r.err
}

func (r *CursorListAutoPager[T]) Index() int {
	return r.run
}
