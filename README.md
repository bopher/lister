# Lister

Lister helps parsing list request (page, limit, sort, order, filters).

## Requirements

### RequestResolver

Request resolver is a function that parse lister fields from request (string, form, etc.). lister contains following resolver by default:

**Note:** You can write your own resolver by implementing `func(lister Lister, data interface{}) error` signature.

**RecordResolver:** this resolver take ListRecord `struct` as input and parse to lister.

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
    "gender": "female",
    "permissions": ["acc", "report"]
  }
}
```

## Create Lister

```go
import "github.com/bopher/lister"
import "fmt"
lst := lister.New()
lst.SetLimits(10, 25, 50, 100)
lst.SetSorts("_id", "name", "last_activity")
lister.JsonStringResolver(lst,`{"page": 2, "limit": 10}`)
lst.SetTotal(/* Get Total Record Count From Somewhere */)
// Do other operations, paginate and fetch record
fmt.Println(lst.ResponseWithData(myData))
```

## Usage

Lister interface contains following methods:

### SetPage

Set current page.

```go
SetPage(page uint)
```

### Page

Get current page.

```go
Page() uint
```

### SetLimits

Set valid limits list.

```go
SetLimits(limits ...uint)
```

### Limits

Get valid limits.

```go
Limits() []uint
```

### SetLimit

Set limit.

```go
SetLimit(limit uint)
```

### Limit

Get limit.

```go
Limit() uint
```

### SetSorts

Set valid sorts list.

```go
SetSorts(sorts ...string)
```

### Sorts

Get valid sorts.

```go
Sorts() []string
```

### SetSort

Set sort.

```go
SetSort(sort string)
```

### Sort

Get sort.

```go
Sort() string
```

### SetOrder

Set order (valid values are `"asc"`, `"desc"`, `"1"`, `"-1"`, `1` and `-1`).

```go
SetOrder(order interface{})
```

### Order

Get order.

```go
Order() string
```

### OrderNumeric

Return order in 1 and -1.

```go
OrderNumeric() int8
```

### SetSearch

Set search phrase.

```go
SetSearch(search string)
```

### Search

Get search phrase.

```go
Search() string
```

### SetFilters

Set filters list.

```go
SetFilters(filters map[string]interface{})
```

### Filters

Get filters list.

```go
Filters() map[string]interface{}
```

### SetFilter

Set filter.

```go
SetFilter(key string, value interface{})
```

### Filter

Get filter.

```go
Filter(key string) interface{}
```

### HasFilter

Check if filter exists.

```go
HasFilter(key string) bool
```

### CastFilter

Parse filter as caster.

```go
CastFilter(key string) caster.Caster
```

### SetMeta

Set meta data.

```go
SetMeta(key string, value interface{})
```

### Meta

Get meta.

```go
Meta(key string) interface{}
```

### HasMeta

Check if meta exists.

```go
HasMeta(key string) bool
```

### CastMeta

Parse meta as caster.

```go
CastMeta(key string) caster.Caster
```

### MetaData

Get meta data list.

```go
MetaData() map[string]interface{}
```

### SetTotal

Set total records count. You must pass total records count to this method for getting paginator information.

**Caution:** Call this method after setting all lister fields(page, limits, etc).

```go
SetTotal(total uint64)
```

### Total

Get total records count.

```go
Total() uint64
```

### From

Get from record position.

```go
From() uint64
```

### To

Get to record position.

```go
To() uint64
```

### Pages

Get total pages count.

```go
Pages() uint
```

### Response

Get response for json, contains pagination information and meta data.

```go
Response() map[string]interface{}
```

### ResponseWithData

Return response with data.

```go
ResponseWithData(data interface{}) map[string]interface{}
```
