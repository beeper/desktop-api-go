// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pagination

import (
	"net/http"
	"reflect"

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
	OldestCursor string `json:"oldestCursor" api:"nullable"`
	NewestCursor string `json:"newestCursor" api:"nullable"`
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

type CursorNoLimit[T any] struct {
	Items        []T    `json:"items"`
	HasMore      bool   `json:"hasMore"`
	OldestCursor string `json:"oldestCursor" api:"nullable"`
	NewestCursor string `json:"newestCursor" api:"nullable"`
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
func (r CursorNoLimit[T]) RawJSON() string { return r.JSON.raw }
func (r *CursorNoLimit[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *CursorNoLimit[T]) GetNextPage() (res *CursorNoLimit[T], err error) {
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

func (r *CursorNoLimit[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &CursorNoLimit[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type CursorNoLimitAutoPager[T any] struct {
	page *CursorNoLimit[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewCursorNoLimitAutoPager[T any](page *CursorNoLimit[T], err error) *CursorNoLimitAutoPager[T] {
	return &CursorNoLimitAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *CursorNoLimitAutoPager[T]) Next() bool {
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

func (r *CursorNoLimitAutoPager[T]) Current() T {
	return r.cur
}

func (r *CursorNoLimitAutoPager[T]) Err() error {
	return r.err
}

func (r *CursorNoLimitAutoPager[T]) Index() int {
	return r.run
}

type CursorSortKey[T any] struct {
	Items   []T  `json:"items"`
	HasMore bool `json:"hasMore"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		HasMore     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r CursorSortKey[T]) RawJSON() string { return r.JSON.raw }
func (r *CursorSortKey[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *CursorSortKey[T]) GetNextPage() (res *CursorSortKey[T], err error) {
	if len(r.Items) == 0 {
		return nil, nil
	}

	if r.JSON.HasMore.Valid() && r.HasMore == false {
		return nil, nil
	}
	items := r.Items
	if items == nil || len(items) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	value := reflect.ValueOf(items[len(items)-1])
	field := value.FieldByName("SortKey")
	err = cfg.Apply(option.WithQuery("cursor", field.Interface().(string)))
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

func (r *CursorSortKey[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &CursorSortKey[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type CursorSortKeyAutoPager[T any] struct {
	page *CursorSortKey[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewCursorSortKeyAutoPager[T any](page *CursorSortKey[T], err error) *CursorSortKeyAutoPager[T] {
	return &CursorSortKeyAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *CursorSortKeyAutoPager[T]) Next() bool {
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

func (r *CursorSortKeyAutoPager[T]) Current() T {
	return r.cur
}

func (r *CursorSortKeyAutoPager[T]) Err() error {
	return r.err
}

func (r *CursorSortKeyAutoPager[T]) Index() int {
	return r.run
}
