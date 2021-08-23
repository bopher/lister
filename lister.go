package lister

// Lister interface
// this interface helps parsing list request and generating paginator information
type Lister interface {
	// SetResolver set request resolver function
	SetResolver(resolver RequestResolver)
	// Parse parse request using registered resolver
	Parse(data interface{}) bool
	// SetPage set current page
	SetPage(page uint)
	// GetPage get current page
	GetPage() uint
	// SetValidLimits set valid limits list
	SetValidLimits(limits ...uint)
	// GetValidLimits get valid limits
	GetValidLimits() []uint
	// SetLimit set limit
	SetLimit(limit uint)
	// GetLimit get limit
	GetLimit() uint
	// SetValidSorts set valid sorts list
	SetValidSorts(sorts ...string)
	// GetValidSort get valid sorts
	GetValidSort() []string
	// SetSort set sort
	SetSort(sort string)
	// GetSort get sort
	GetSort() string
	// SetOrder set order
	// Valid values are "asc", "desc", "1", "-1", 1 and -1
	SetOrder(order interface{})
	// GetOrder get order
	GetOrder() string
	// GetNumericOrder return order in 1 and -1
	GetNumericOrder() int8
	// SetSearch set search phrase
	SetSearch(search string)
	// GetSearch get search phrase
	GetSearch() string
	// SetFilters set filters list
	SetFilters(filters map[string]interface{})
	// GetFilters get filters list
	GetFilters() map[string]interface{}
	// SetFilter set filter
	SetFilter(key string, value interface{})
	// GetFilter get filter
	GetFilter(key string) interface{}
	// HasFilter check if filter exists
	HasFilter(key string) bool
	// SliceFilter get slice filter or return fallback if filter not exists
	SliceFilter(key string, fallback []interface{}) ([]interface{}, bool)
	// StringFilter get string filter or return fallback if filter not exists
	StringFilter(key string, fallback string) (string, bool)
	// StringSliceFilter get string slice filter or return fallback if filter not exists
	StringSliceFilter(key string, fallback []string) ([]string, bool)
	// BoolFilter get bool filter or return fallback if filter not exists
	BoolFilter(key string, fallback bool) (bool, bool)
	// BoolSliceFilter get bool slice filter or return fallback if filter not exists
	BoolSliceFilter(key string, fallback []bool) ([]bool, bool)
	// Float64Filter get float64 filter or return fallback if filter not exists
	Float64Filter(key string, fallback float64) (float64, bool)
	// Float64SliceFilter get float64 slice filter or return fallback if filter not exists
	Float64SliceFilter(key string, fallback []float64) ([]float64, bool)
	// Int64Filter get int64 filter or return fallback if filter not exists
	Int64Filter(key string, fallback int64) (int64, bool)
	// Int64SliceFilter get int64 slice filter or return fallback if filter not exists
	Int64SliceFilter(key string, fallback []int64) ([]int64, bool)
	// SetMeta set meta data
	SetMeta(key string, value interface{})
	// GetMeta get meta
	GetMeta(key string) interface{}
	// HasMeta check if meta exists
	HasMeta(key string) bool
	// SliceMeta get slice meta or return fallback if meta not exists
	SliceMeta(key string, fallback []interface{}) ([]interface{}, bool)
	// StringMeta get string meta or return fallback if meta not exists
	StringMeta(key string, fallback string) (string, bool)
	// StringSliceMeta get string slice slice meta or return fallback if meta not exists
	StringSliceMeta(key string, fallback []string) ([]string, bool)
	// BoolMeta get bool meta or return fallback if meta not exists
	BoolMeta(key string, fallback bool) (bool, bool)
	// BoolSliceMeta get bool slice meta or return fallback if meta not exists
	BoolSliceMeta(key string, fallback []bool) ([]bool, bool)
	// Float64Meta get float64 meta or return fallback if meta not exists
	Float64Meta(key string, fallback float64) (float64, bool)
	// Float64SliceMeta get float64 slice meta or return fallback if meta not exists
	Float64SliceMeta(key string, fallback []float64) ([]float64, bool)
	// Int64Meta get int64 meta or return fallback if meta not exists
	Int64Meta(key string, fallback int64) (int64, bool)
	// Int64SliceMeta get int64 slice meta or return fallback if meta not exists
	Int64SliceMeta(key string, fallback []int64) ([]int64, bool)
	// MetaData get meta data list
	MetaData() map[string]interface{}
	// SetTotal Set total records count
	SetTotal(total uint64)
	// GetTotal get total records count
	GetTotal() uint64
	// From get from record position
	From() uint64
	// To get to record position
	To() uint64
	// Pages get total pages count
	Pages() uint
	// Response get response for json
	// contains pagination information and meta data
	Response() map[string]interface{}
	// ResponseWithData return response with data
	ResponseWithData(data interface{}) map[string]interface{}
}
