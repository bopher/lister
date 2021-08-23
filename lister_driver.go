package lister

import (
	"math"
	"strings"
)

// lister driver
type lister struct {
	page    uint
	limit   uint
	sort    string
	order   string
	search  string
	filters map[string]interface{}

	resolver RequestResolver
	limits   []uint
	sorts    []string
	meta     map[string]interface{}

	total      uint64
	from       uint64
	to         uint64
	pagesCount uint
}

// init initialize lister driver
func (l *lister) init() {
	l.page = 1
	l.limit = 25
	l.sort = "id"
	l.order = "asc"
	l.filters = make(map[string]interface{})
	l.resolver = FiberFormResolver
	l.limits = []uint{10, 25, 50, 100, 250}
	l.sorts = []string{"id"}
	l.meta = make(map[string]interface{})
}

// SetResolver set request resolver function
func (l *lister) SetResolver(resolver RequestResolver) {
	if resolver != nil {
		l.resolver = resolver
	}
}

// Parse parse request using registered resolver
func (l *lister) Parse(data interface{}) bool {
	return l.resolver(l, data)
}

// SetPage set current page
func (l *lister) SetPage(page uint) {
	if page > 0 {
		if l.pagesCount > 0 && page > l.pagesCount {
			l.page = l.pagesCount
			return
		}
		l.page = page
	}
}

// GetPage get current page
func (l *lister) GetPage() uint {
	return l.page
}

// SetValidLimits set valid limits list
func (l *lister) SetValidLimits(limits ...uint) {
	if len(limits) > 0 {
		l.limits = limits
	}
}

// GetValidLimits get valid limits
func (l *lister) GetValidLimits() []uint {
	return l.limits
}

// SetLimit set limit
func (l *lister) SetLimit(limit uint) {
	for _, lmt := range l.limits {
		if lmt == limit {
			l.limit = limit
		}
	}
}

// GetLimit get limit
func (l *lister) GetLimit() uint {
	return l.limit
}

// SetValidSorts set valid sorts list
func (l *lister) SetValidSorts(sorts ...string) {
	if len(sorts) > 0 {
		l.sorts = sorts
	}
}

// GetValidSort get valid sorts
func (l *lister) GetValidSort() []string {
	return l.sorts
}

// SetSort set sort
func (l *lister) SetSort(sort string) {
	for _, srt := range l.sorts {
		if srt == sort {
			l.sort = sort
		}
	}
}

// GetSort get sort
func (l *lister) GetSort() string {
	return l.sort
}

// SetOrder set order
func (l *lister) SetOrder(order interface{}) {
	if str, ok := order.(string); ok {
		str = strings.ToLower(str)
		if str == "desc" || str == "-1" {
			l.order = "desc"
		} else if str == "asc" || str == "1" {
			l.order = "asc"
		}
	} else if num, ok := order.(int); ok {
		if num == -1 {
			l.order = "desc"
		} else if num == 1 {
			l.order = "asc"
		}
	}
}

// GetOrder get order
func (l *lister) GetOrder() string {
	return l.order
}

// GetNumericOrder return order in 1 and -1
func (l *lister) GetNumericOrder() int8 {
	if l.order == "desc" {
		return -1
	}
	return 1
}

// SetSearch set search phrase
func (l *lister) SetSearch(search string) {
	l.search = search
}

// GetSearch get search phrase
func (l *lister) GetSearch() string {
	return l.search
}

// SetFilters set filters list
func (l *lister) SetFilters(filters map[string]interface{}) {
	if filters != nil {
		l.filters = filters
	}
}

// GetFilters get filters list
func (l *lister) GetFilters() map[string]interface{} {
	return l.filters
}

// SetFilter set filter
func (l *lister) SetFilter(key string, value interface{}) {
	l.filters[key] = value
}

// GetFilter get filter
func (l *lister) GetFilter(key string) interface{} {
	return l.filters[key]
}

// HasFilter check if filter exists
func (l *lister) HasFilter(key string) bool {
	_, ok := l.filters[key]
	return ok
}

// SliceFilter get slice filter or return fallback if filter not exists
func (l *lister) SliceFilter(key string, fallback []interface{}) ([]interface{}, bool) {
	if val, ok := l.filters[key]; ok {
		if sliceVal, ok := val.([]interface{}); ok {
			return sliceVal, true
		}
	}
	return fallback, false
}

// StringFilter get string filter or return fallback if filter not exists
func (l *lister) StringFilter(key string, fallback string) (string, bool) {
	if val, ok := l.filters[key]; ok {
		if strVal, ok := val.(string); ok {
			return strVal, true
		}
	}
	return fallback, false
}

// StringSliceFilter get string slice filter or return fallback if filter not exists
func (l *lister) StringSliceFilter(key string, fallback []string) ([]string, bool) {
	if val, ok := l.filters[key]; ok {
		if strVal, ok := val.([]string); ok {
			return strVal, true
		}
	}
	return fallback, false
}

// BoolFilter get bool filter or return fallback if filter not exists
func (l *lister) BoolFilter(key string, fallback bool) (bool, bool) {
	if val, ok := l.filters[key]; ok {
		if boolVal, ok := val.(bool); ok {
			return boolVal, true
		}
	}
	return fallback, false
}

// BoolSliceFilter get bool slice filter or return fallback if filter not exists
func (l *lister) BoolSliceFilter(key string, fallback []bool) ([]bool, bool) {
	if val, ok := l.filters[key]; ok {
		if boolVal, ok := val.([]bool); ok {
			return boolVal, true
		}
	}
	return fallback, false
}

// Float64Filter get float64 filter or return fallback if filter not exists
func (l *lister) Float64Filter(key string, fallback float64) (float64, bool) {
	if val, ok := l.filters[key]; ok {
		if floatVal, ok := val.(float64); ok {
			return floatVal, true
		}
	}
	return fallback, false
}

// Float64SliceFilter get float64 slice filter or return fallback if filter not exists
func (l *lister) Float64SliceFilter(key string, fallback []float64) ([]float64, bool) {
	if val, ok := l.filters[key]; ok {
		if floatVal, ok := val.([]float64); ok {
			return floatVal, true
		}
	}
	return fallback, false
}

// Int64Filter get int64 filter or return fallback if filter not exists
func (l *lister) Int64Filter(key string, fallback int64) (int64, bool) {
	if val, ok := l.filters[key]; ok {
		if intVal, ok := val.(int64); ok {
			return intVal, true
		}
	}
	return fallback, false
}

// Int64SliceFilter get int64 slice filter or return fallback if filter not exists
func (l *lister) Int64SliceFilter(key string, fallback []int64) ([]int64, bool) {
	if val, ok := l.filters[key]; ok {
		if intVal, ok := val.([]int64); ok {
			return intVal, true
		}
	}
	return fallback, false
}

// SetMeta set meta data
func (l *lister) SetMeta(key string, value interface{}) {
	l.meta[key] = value
}

// GetMeta get meta
func (l *lister) GetMeta(key string) interface{} {
	return l.meta[key]
}

// HasMeta check if meta exists
func (l *lister) HasMeta(key string) bool {
	_, ok := l.meta[key]
	return ok
}

// SliceMeta get slice meta or return fallback if meta not exists
func (l *lister) SliceMeta(key string, fallback []interface{}) ([]interface{}, bool) {
	if val, ok := l.meta[key]; ok {
		if sliceVal, ok := val.([]interface{}); ok {
			return sliceVal, true
		}
	}
	return fallback, false
}

// StringMeta get string meta or return fallback if meta not exists
func (l *lister) StringMeta(key string, fallback string) (string, bool) {
	if val, ok := l.meta[key]; ok {
		if strVal, ok := val.(string); ok {
			return strVal, true
		}
	}
	return fallback, false
}

// StringSliceMeta get string slice slice meta or return fallback if meta not exists
func (l *lister) StringSliceMeta(key string, fallback []string) ([]string, bool) {
	if val, ok := l.meta[key]; ok {
		if strVal, ok := val.([]string); ok {
			return strVal, true
		}
	}
	return fallback, false
}

// BoolMeta get bool meta or return fallback if meta not exists
func (l *lister) BoolMeta(key string, fallback bool) (bool, bool) {
	if val, ok := l.meta[key]; ok {
		if boolVal, ok := val.(bool); ok {
			return boolVal, true
		}
	}
	return fallback, false
}

// BoolSliceMeta get bool slice meta or return fallback if meta not exists
func (l *lister) BoolSliceMeta(key string, fallback []bool) ([]bool, bool) {
	if val, ok := l.meta[key]; ok {
		if boolVal, ok := val.([]bool); ok {
			return boolVal, true
		}
	}
	return fallback, false
}

// Float64Meta get float64 meta or return fallback if meta not exists
func (l *lister) Float64Meta(key string, fallback float64) (float64, bool) {
	if val, ok := l.meta[key]; ok {
		if floatVal, ok := val.(float64); ok {
			return floatVal, true
		}
	}
	return fallback, false
}

// Float64SliceMeta get float64 slice meta or return fallback if meta not exists
func (l *lister) Float64SliceMeta(key string, fallback []float64) ([]float64, bool) {
	if val, ok := l.meta[key]; ok {
		if floatVal, ok := val.([]float64); ok {
			return floatVal, true
		}
	}
	return fallback, false
}

// Int64Meta get int64 meta or return fallback if meta not exists
func (l *lister) Int64Meta(key string, fallback int64) (int64, bool) {
	if val, ok := l.meta[key]; ok {
		if intVal, ok := val.(int64); ok {
			return intVal, true
		}
	}
	return fallback, false
}

// Int64SliceMeta get int64 slice meta or return fallback if meta not exists
func (l *lister) Int64SliceMeta(key string, fallback []int64) ([]int64, bool) {
	if val, ok := l.meta[key]; ok {
		if intVal, ok := val.([]int64); ok {
			return intVal, true
		}
	}
	return fallback, false
}

// MetaData get meta data list
func (l *lister) MetaData() map[string]interface{} {
	return l.meta
}

// SetTotal Set total records count
func (l *lister) SetTotal(total uint64) {
	l.total = total
	l.pagesCount = uint(math.Ceil(float64(l.total) / float64(l.limit)))
	if l.page > l.pagesCount {
		l.page = l.pagesCount
	}
	if l.page < 1 {
		l.page = 1
	}

	l.from = (uint64(l.page-1) * uint64(l.limit))

	l.to = l.from + uint64(l.limit)
	if l.to > total {
		l.to = total
	}
}

// GetTotal get total records count
func (l *lister) GetTotal() uint64 {
	return l.total
}

// From get from record position
func (l *lister) From() uint64 {
	return l.from
}

// To get to record position
func (l *lister) To() uint64 {
	return l.to
}

// Pages get total pages count
func (l *lister) Pages() uint {
	return l.pagesCount
}

// Response get response for json
// contains pagination information and meta data
func (l *lister) Response() map[string]interface{} {
	res := make(map[string]interface{})
	for k, v := range l.meta {
		res[k] = v
	}
	res["page"] = l.page
	res["limit"] = l.limit
	res["sort"] = l.sort
	res["order"] = l.order
	res["search"] = l.search
	res["total"] = l.total
	res["from"] = l.from + 1
	res["to"] = l.to
	res["pages"] = l.pagesCount
	return res
}

// ResponseWithData return response with data
func (l *lister) ResponseWithData(data interface{}) map[string]interface{} {
	res := make(map[string]interface{})
	for k, v := range l.meta {
		res[k] = v
	}
	res["page"] = l.page
	res["limit"] = l.limit
	res["sort"] = l.sort
	res["order"] = l.order
	res["search"] = l.search
	res["total"] = l.total
	res["from"] = l.from + 1
	res["to"] = l.to
	res["pages"] = l.pagesCount
	res["data"] = data
	return res
}
