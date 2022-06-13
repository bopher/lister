package lister

import (
	"fmt"
	"math"
	"strings"

	"github.com/bopher/caster"
	"github.com/bopher/utils"
)

type lDriver struct {
	page    uint
	limit   uint
	sort    string
	order   string
	search  string
	filters map[string]any

	limits []uint
	sorts  []string
	meta   map[string]any

	total      uint64
	from       uint64
	to         uint64
	pagesCount uint
}

func (ld *lDriver) init() {
	ld.page = 1
	ld.limit = 25
	ld.sort = "_id"
	ld.order = "asc"
	ld.filters = make(map[string]any)
	ld.limits = []uint{10, 25, 50, 100, 250}
	ld.sorts = []string{"_id"}
	ld.meta = make(map[string]any)
}

func (ld *lDriver) SetPage(page uint) {
	if page > 0 {
		if ld.pagesCount > 0 && page > ld.pagesCount {
			ld.page = ld.pagesCount
			return
		}
		ld.page = page
	}
}

func (ld lDriver) Page() uint {
	return ld.page
}

func (ld *lDriver) SetLimits(limits ...uint) {
	if len(limits) > 0 {
		ld.limits = limits
	}
}

func (ld lDriver) Limits() []uint {
	return ld.limits
}

func (ld *lDriver) SetLimit(limit uint) {
	if utils.Contains[uint](ld.limits, limit) {
		ld.limit = limit
	}
}

func (ld lDriver) Limit() uint {
	return ld.limit
}

func (ld *lDriver) SetSorts(sorts ...string) {
	if len(sorts) > 0 {
		ld.sorts = sorts
	}
}

func (ld lDriver) Sorts() []string {
	return ld.sorts
}

func (ld *lDriver) SetSort(sort string) {
	if utils.Contains[string](ld.sorts, sort) {
		ld.sort = sort
	}
}

func (ld lDriver) Sort() string {
	return ld.sort
}

func (ld *lDriver) SetOrder(order any) {
	o := strings.ToLower(fmt.Sprint(order))
	if o == "-1" {
		o = "desc"
	}
	if o == "1" {
		o = "asc"
	}
	if o == "asc" || o == "desc" {
		ld.order = o
	}
}

func (ld lDriver) Order() string {
	return ld.order
}

func (ld lDriver) OrderNumeric() int8 {
	if ld.order == "desc" {
		return -1
	}
	return 1
}

func (ld *lDriver) SetSearch(search string) {
	ld.search = search
}

func (ld lDriver) Search() string {
	return ld.search
}

func (ld *lDriver) SetFilters(filters map[string]any) {
	if filters != nil {
		ld.filters = filters
	} else {
		ld.filters = make(map[string]any)
	}
}

func (ld lDriver) Filters() map[string]any {
	return ld.filters
}

func (ld *lDriver) SetFilter(key string, value any) {
	ld.filters[key] = value
}

func (ld lDriver) Filter(key string) any {
	return ld.filters[key]
}

func (ld lDriver) HasFilter(key string) bool {
	_, exists := ld.filters[key]
	return exists
}

func (ld lDriver) CastFilter(key string) caster.Caster {
	return caster.NewCaster(ld.filters[key])
}

func (ld *lDriver) SetMeta(key string, value any) {
	ld.meta[key] = value
}

func (ld lDriver) Meta(key string) any {
	return ld.meta[key]
}

func (ld lDriver) HasMeta(key string) bool {
	_, exists := ld.meta[key]
	return exists
}

func (ld lDriver) MetaData() map[string]any {
	return ld.meta
}

func (ld lDriver) CastMeta(key string) caster.Caster {
	return caster.NewCaster(ld.meta[key])
}

func (ld *lDriver) SetTotal(total uint64) {
	ld.total = total
	ld.pagesCount = uint(math.Ceil(float64(ld.total) / float64(ld.limit)))
	if ld.page > ld.pagesCount {
		ld.page = ld.pagesCount
	}
	if ld.page < 1 {
		ld.page = 1
	}

	ld.from = (uint64(ld.page-1) * uint64(ld.limit))

	ld.to = ld.from + uint64(ld.limit)
	if ld.to > total {
		ld.to = total
	}
}

func (ld lDriver) Total() uint64 {
	return ld.total
}

func (ld lDriver) From() uint64 {
	return ld.from
}

func (ld lDriver) To() uint64 {
	return ld.to
}

func (ld lDriver) Pages() uint {
	return ld.pagesCount
}

func (ld lDriver) SQLSortOrder() string {
	return fmt.Sprintf(" ORDER BY %s %s LIMIT %d, %d", ld.sort, ld.order, ld.from-1, ld.limit)
}

func (ld lDriver) Response() map[string]any {
	res := make(map[string]any)
	for k, v := range ld.meta {
		res[k] = v
	}
	res["page"] = ld.page
	res["limit"] = ld.limit
	res["sort"] = ld.sort
	res["order"] = ld.order
	res["search"] = ld.search
	res["total"] = ld.total
	res["from"] = ld.from + 1
	res["to"] = ld.to
	res["pages"] = ld.pagesCount
	return res
}

func (ld lDriver) ResponseWithData(data any) map[string]any {
	res := make(map[string]any)
	for k, v := range ld.meta {
		res[k] = v
	}
	res["page"] = ld.page
	res["limit"] = ld.limit
	res["sort"] = ld.sort
	res["order"] = ld.order
	res["search"] = ld.search
	res["total"] = ld.total
	res["from"] = ld.from + 1
	res["to"] = ld.to
	res["pages"] = ld.pagesCount
	res["data"] = data
	return res
}
