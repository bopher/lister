package lister

import (
	"encoding/base64"
	"encoding/json"

	"github.com/bopher/utils"
	"github.com/gofiber/fiber/v2"
)

type ListerRequest struct {
	Page    uint           `json:"page" form:"page" xml:"page"`
	Limit   uint           `json:"limit" form:"limit" xml:"limit"`
	Sort    string         `json:"sort" form:"sort" xml:"sort"`
	Order   string         `json:"order" form:"order" xml:"order"`
	Search  string         `json:"search" form:"search" xml:"search"`
	Filters map[string]any `json:"filters" form:"filters" xml:"filters"`
}

// RequestResolver
type RequestResolver func(lister Lister, data any) error

// RecordResolver fill lister from ListerRecord
func RecordResolver(lister Lister, data any) error {
	if rec, ok := data.(ListerRequest); ok {
		lister.SetPage(rec.Page)
		lister.SetLimit(rec.Limit)
		lister.SetSort(rec.Sort)
		lister.SetOrder(rec.Order)
		lister.SetSearch(rec.Search)
		lister.SetFilters(rec.Filters)
		return nil
	}
	return utils.TaggedError([]string{"RecordResolver"}, "data is not valid!")
}

// Base64Resolver parse data from base64 encoded json string
func Base64Resolver(lister Lister, data any) error {
	if qs, ok := data.(string); ok {
		base64decoded := make([]byte, base64.StdEncoding.EncodedLen(len(qs)))
		if _, err := base64.StdEncoding.Decode(base64decoded, []byte(qs)); err == nil {
			return JsonStringResolver(lister, string(base64decoded))
		} else {
			return utils.TaggedError([]string{"Base64Resolver"}, err.Error())
		}
	}
	return utils.TaggedError([]string{"Base64Resolver"}, "data is not valid!")
}

// JsonStringResolver parse parameters from json string
func JsonStringResolver(lister Lister, data any) error {
	if qs, ok := data.(string); ok {
		record := ListerRequest{}
		if err := json.Unmarshal([]byte(qs), &record); err == nil {
			return RecordResolver(lister, record)
		} else {
			return utils.TaggedError([]string{"JsonStringResolver"}, err.Error())
		}
	}
	return utils.TaggedError([]string{"JsonStringResolver"}, "data is not valid!")
}

// FiberFormResolver parse parameters from fiber context
func FiberFormResolver(lister Lister, data any) error {
	if ctx, ok := data.(*fiber.Ctx); ok {
		record := ListerRequest{}
		if err := ctx.BodyParser(&record); err == nil {
			return RecordResolver(lister, record)
		} else {
			return utils.TaggedError([]string{"FiberFormResolver"}, err.Error())
		}
	}
	return utils.TaggedError([]string{"FiberFormResolver"}, "data is not valid!")
}
