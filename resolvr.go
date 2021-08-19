package lister

import (
	"encoding/base64"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

type ListerRequest struct {
	Page    uint                   `json:"page" form:"page" xml:"page"`
	Limit   uint8                  `json:"limit" form:"limit" xml:"limit"`
	Sort    string                 `json:"sort" form:"sort" xml:"sort"`
	Order   string                 `json:"order" form:"order" xml:"order"`
	Search  string                 `json:"search" form:"search" xml:"search"`
	Filters map[string]interface{} `json:"filters" form:"filters" xml:"filters"`
}

// RequestResolver
type RequestResolver func(lister Lister, data interface{}) bool

// ListerRecordResolver fill lister from ListerRecord
func ListerRecordResolver(lister Lister, data interface{}) bool {
	if rec, ok := data.(ListerRequest); ok {
		lister.SetPage(rec.Page)
		lister.SetLimit(rec.Limit)
		lister.SetSort(rec.Sort)
		lister.SetOrder(rec.Order)
		lister.SetSearch(rec.Search)
		lister.SetFilters(rec.Filters)
		return true
	}
	return false
}

// Base64Resolver parse data from base64 encoded json string
func Base64Resolver(lister Lister, data interface{}) bool {
	if qs, ok := data.(string); ok {
		var base64decoded []byte
		if _, err := base64.StdEncoding.Decode(base64decoded, []byte(qs)); err == nil {
			return JsonStringResolver(lister, string(base64decoded))
		}
	}
	return false
}

// JsonStringResolver parse parameters from json string
func JsonStringResolver(lister Lister, data interface{}) bool {
	if qs, ok := data.(string); ok {
		record := ListerRequest{}
		if err := json.Unmarshal([]byte(qs), &record); err == nil {
			return ListerRecordResolver(lister, record)
		}
	}
	return false
}

// FiberFormResolver parse parameters from fiber context
func FiberFormResolver(lister Lister, data interface{}) bool {
	if ctx, ok := data.(*fiber.Ctx); ok {
		record := new(ListerRequest)
		if err := ctx.BodyParser(&record); err == nil {
			return ListerRecordResolver(lister, record)
		}
	}
	return false
}
