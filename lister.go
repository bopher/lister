package lister

import "github.com/bopher/caster"

type Lister interface {
	// SetPage set current page
	SetPage(page uint)
	// Page get current page
	Page() uint
	// SetLimits set valid limits list
	SetLimits(limits ...uint)
	// Limits get valid limits
	Limits() []uint
	// SetLimit set limit
	SetLimit(limit uint)
	// Limit get limit
	Limit() uint
	// SetSorts set valid sorts list
	SetSorts(sorts ...string)
	// Sorts get valid sorts
	Sorts() []string
	// SetSort set sort
	SetSort(sort string)
	// Sort get sort
	Sort() string
	// SetOrder set order (valid values are "asc", "desc", "1", "-1", 1 and -1)
	SetOrder(order any)
	// Order get order
	Order() string
	// OrderNumeric return order in 1 and -1
	OrderNumeric() int8
	// SetSearch set search phrase
	SetSearch(search string)
	// Search get search phrase
	Search() string
	// SetFilters set filters list
	SetFilters(filters map[string]any)
	// Filters get filters list
	Filters() map[string]any
	// SetFilter set filter
	SetFilter(key string, value any)
	// Filter get filter
	Filter(key string) any
	// HasFilter check if filter exists
	HasFilter(key string) bool
	// CastFilter parse filter as caster
	CastFilter(key string) caster.Caster
	// SetMeta set meta data
	SetMeta(key string, value any)
	// Meta get meta
	Meta(key string) any
	// HasMeta check if meta exists
	HasMeta(key string) bool
	// MetaData get meta data list
	MetaData() map[string]any
	// CastMeta parse meta as caster
	CastMeta(key string) caster.Caster
	// SetTotal Set total records count
	SetTotal(total uint64)
	// Total get total records count
	Total() uint64
	// From get from record position
	From() uint64
	// To get to record position
	To() uint64
	// Pages get total pages count
	Pages() uint
	// SQLSortOrder get sql order and limit command as string
	SQLSortOrder() string
	// Response get response for json, contains pagination information and meta data
	Response() map[string]any
	// ResponseWithData return response with data
	ResponseWithData(data any) map[string]any
}
