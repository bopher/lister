# Lister

Lister helps parsing list request (page, limit, sort, order, filters) and generating paginator information.

## Requirements

### RequestResolver

Request resolver is a function that parse lister fields from request (string, form, etc.). lister contains following resolver by default:

**Note:** You can write your own resolver by implementing `func(lister Lister, data interface{}) bool` signature.

**ListerRecordResolver:** this resolver take ListRecord struct as input and parse to lister.

**Base64Resolver:** this resolver parse lister fields from Base64 encoded json string.

**JsonStringResolver:** this resolver parse lister fields from json string.

**FiberFormResolver:** this resolver parse lister fields from goFiber request context (json, form and xml supported).

#### Request Fields Signature

```json
{
    "page": 1,
    "limit": 10,
    "sort": "name",
    "order": "asc",
    "search": "john",
    "filters": {
        "minAge": 25,
        "gender": "female"
    }
}
```

## Create Lister

```go
import "github.com/bopher/lister"
import "fmt"
lst := lister.New()
lst.SetResolver(lister.JsonStringResolver)
lst.SetValidLimits(10, 25, 50, 100)
lst.SetValidSorts("id", "name", "last_activity")
lst.Parse(`{"page": 2, "limit": 10}`)
lst.SetTotal(/* Get Total Record Count From Somewhere */)
// Do other operations, paginate and fetch record
fmt.Println(lst.ResponseWithData(myData))
```

## Usage

Lister interface contains following methods:

### SetResolver

Set request resolver function.

```go
// Signature:
SetResolver(resolver RequestResolver)
```

### Parse

Parse request using registered resolver.

```go
// Signature:
Parse(data interface{}) bool
```

### SetPage

Set current page.

```go
// Signature:
SetPage(page uint)
```

### GetPage

Get current page.

```go
// Signature:
GetPage() uint
```

### SetValidLimits

Set valid limits list.

```go
// Signature:
SetValidLimits(limits ...uint8)
```

### GetValidLimits

Get valid limits.

```go
// Signature:
GetValidLimits() []uint8
```

### SetLimit

Set limit.

```go
// Signature:
SetLimit(limit uint8)
```

### GetLimit

Get limit.

```go
// Signature:
GetLimit() uint8
```

### SetValidSorts

Set valid sorts list.

```go
// Signature:
SetValidSorts(sorts ...string)
```

### GetValidSort

Get valid sorts list.

```go
// Signature:
GetValidSort() []string
```

### SetSort

Set sort.

```go
// Signature:
SetSort(sort string)
```

### GetSort

Get sort.

```go
// Signature:
GetSort() string
```

### SetOrder

Set order. Valid values are `"asc"`, `"desc"`, `"1"`, `"-1"`, `1` and `-1`.

```go
// Signature:
SetOrder(order interface{})
```

### GetOrder

Get order.

```go
// Signature:
GetOrder() string
```

### GetNumericOrder

Return order in numeric format (1 and -1).

```go
// Signature:
GetNumericOrder() int8
```

### SetSearch

Set search phrase.

```go
// Signature:
SetSearch(search string)
```

### GetSearch

Get search phrase.

```go
// Signature:
GetSearch() string
```

### SetFilters

Set filters list.

```go
// Signature:
SetFilters(filters map[string]interface{})
```

### GetFilters

Get filters list.

```go
// Signature:
GetFilters() map[string]interface{}
```

### SetFilter

Set filter.

```go
SetFilter(key string, value interface{})
```

### GetFilter

Get filter.

```go
// Signature:
GetFilter(key string) interface{}
```

### HasFilter

Check if filter exists.

```go
// Signature:
HasFilter(key string) bool
```

### Get Filter By Type

For getting filters with type you can use helper getter methods. getter methods accept a fallback value and returns fallback if value not exists or not in type. getter methods follow ok pattern. Getter methods list:

```go
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
```

### SetMeta

Set meta data.

```go
// Signature:
SetMeta(key string, value interface{})
```

### GetMeta

Get meta.

```go
// Signature:
GetMeta(key string) interface{}
```

### HasMeta

Check if meta exists.

```go
// Signature:
HasMeta(key string) bool
```

### MetaData

Get meta data list.

```go
// Signature:
MetaData() map[string]interface{}
```

### Get Meta By Type

For getting meta with type you can use helper getter methods. getter methods accept a fallback value and returns fallback if value not exists or not in type. getter methods follow ok pattern. Getter methods list:

```go
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
```

### SetTotal

Set total records count. You must pass total records count to this method for getting paginator information.

**Caution:** Call this method after setting all lister fields(page, limits, etc).

```go
// Signature:
SetTotal(total uint64)
```

### GetTotal

Get total records count.

```go
// Signature:
GetTotal() uint64
```

### From

Get from record position.

```go
// Signature:
From() uint64
```

### To

Get to record position.

```go
// Signature:
To() uint64
```

### Pages

Get total pages count.

```go
// Signature:
Pages() uint
```

### Response

Get response for json, contains pagination information and meta data.

```go
// Signature:
Response() map[string]interface{}
```

### ResponseWithData

return response with data field.

```go
// Signature:
ResponseWithData(data interface{}) map[string]interface{}
```
