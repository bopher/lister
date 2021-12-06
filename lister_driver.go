package lister

import (
	"fmt"
	"math"
	"strings"

	"github.com/bopher/caster"
)

type lDriver struct {
	page    uint
	limit   uint
	sort    string
	order   string
	search  string
	filters map[string]interface{}

	limits []uint
	sorts  []string
	meta   map[string]interface{}

	total      uint64
	from       uint64
	to         uint64
	pagesCount uint
}

func (this *lDriver) init() {
	this.page = 1
	this.limit = 25
	this.sort = "_id"
	this.order = "asc"
	this.filters = make(map[string]interface{})
	this.limits = []uint{10, 25, 50, 100, 250}
	this.sorts = []string{"_id"}
	this.meta = make(map[string]interface{})
}

func (this *lDriver) SetPage(page uint) {
	if page > 0 {
		if this.pagesCount > 0 && page > this.pagesCount {
			this.page = this.pagesCount
			return
		}
		this.page = page
	}
}

func (this lDriver) Page() uint {
	return this.page
}

func (this *lDriver) SetLimits(limits ...uint) {
	if len(limits) > 0 {
		this.limits = limits
	}
}

func (this lDriver) Limits() []uint {
	return this.limits
}

func (this *lDriver) SetLimit(limit uint) {
	for _, l := range this.limits {
		if l == limit {
			this.limit = limit
		}
	}
}

func (this lDriver) Limit() uint {
	return this.limit
}

func (this *lDriver) SetSorts(sorts ...string) {
	if len(sorts) > 0 {
		this.sorts = sorts
	}
}

func (this lDriver) Sorts() []string {
	return this.sorts
}

func (this *lDriver) SetSort(sort string) {
	for _, s := range this.sorts {
		if s == sort {
			this.sort = sort
		}
	}
}

func (this lDriver) Sort() string {
	return this.sort
}

func (this *lDriver) SetOrder(order interface{}) {
	o := strings.ToLower(fmt.Sprint(order))
	if o == "-1" {
		o = "desc"
	}
	if o == "1" {
		o = "asc"
	}
	if o == "asc" || o == "desc" {
		this.order = o
	}
}

func (this lDriver) Order() string {
	return this.order
}

func (this lDriver) OrderNumeric() int8 {
	if this.order == "desc" {
		return -1
	}
	return 1
}

func (this *lDriver) SetSearch(search string) {
	this.search = search
}

func (this lDriver) Search() string {
	return this.search
}

func (this *lDriver) SetFilters(filters map[string]interface{}) {
	if filters != nil {
		this.filters = filters
	} else {
		this.filters = make(map[string]interface{})
	}
}

func (this lDriver) Filters() map[string]interface{} {
	return this.filters
}

func (this *lDriver) SetFilter(key string, value interface{}) {
	this.filters[key] = value
}

func (this lDriver) Filter(key string) interface{} {
	return this.filters[key]
}

func (this lDriver) HasFilter(key string) bool {
	_, exists := this.filters[key]
	return exists
}

func (this lDriver) CastFilter(key string) caster.Caster {
	return caster.NewCaster(this.filters[key])
}

func (this *lDriver) SetMeta(key string, value interface{}) {
	this.meta[key] = value
}

func (this lDriver) Meta(key string) interface{} {
	return this.meta[key]
}

func (this lDriver) HasMeta(key string) bool {
	_, exists := this.meta[key]
	return exists
}

func (this lDriver) MetaData() map[string]interface{} {
	return this.meta
}

func (this lDriver) CastMeta(key string) caster.Caster {
	return caster.NewCaster(this.meta[key])
}

func (this *lDriver) SetTotal(total uint64) {
	this.total = total
	this.pagesCount = uint(math.Ceil(float64(this.total) / float64(this.limit)))
	if this.page > this.pagesCount {
		this.page = this.pagesCount
	}
	if this.page < 1 {
		this.page = 1
	}

	this.from = (uint64(this.page-1) * uint64(this.limit))

	this.to = this.from + uint64(this.limit)
	if this.to > total {
		this.to = total
	}
}

func (this lDriver) Total() uint64 {
	return this.total
}

func (this lDriver) From() uint64 {
	return this.from
}

func (this lDriver) To() uint64 {
	return this.to
}

func (this lDriver) Pages() uint {
	return this.pagesCount
}

func (this lDriver) Response() map[string]interface{} {
	res := make(map[string]interface{})
	for k, v := range this.meta {
		res[k] = v
	}
	res["page"] = this.page
	res["limit"] = this.limit
	res["sort"] = this.sort
	res["order"] = this.order
	res["search"] = this.search
	res["total"] = this.total
	res["from"] = this.from + 1
	res["to"] = this.to
	res["pages"] = this.pagesCount
	return res
}

func (this lDriver) ResponseWithData(data interface{}) map[string]interface{} {
	res := make(map[string]interface{})
	for k, v := range this.meta {
		res[k] = v
	}
	res["page"] = this.page
	res["limit"] = this.limit
	res["sort"] = this.sort
	res["order"] = this.order
	res["search"] = this.search
	res["total"] = this.total
	res["from"] = this.from + 1
	res["to"] = this.to
	res["pages"] = this.pagesCount
	res["data"] = data
	return res
}
