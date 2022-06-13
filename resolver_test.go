package lister_test

import (
	"testing"

	"github.com/bopher/lister"
)

func TestRecordResolver(t *testing.T) {
	l := lister.New()
	l.SetSorts("_id", "title")
	lister.RecordResolver(l, lister.ListerRequest{
		Page:   10,
		Limit:  100,
		Sort:   "title",
		Order:  "desc",
		Search: "John",
		Filters: map[string]any{
			"username": "JackMa",
		},
	})

	if l.Page() != 10 ||
		l.Limit() != 100 ||
		l.Sort() != "title" ||
		l.Order() != "desc" ||
		l.Search() != "John" ||
		l.CastFilter("username").StringSafe("ss") != "JackMa" {
		t.Fail()
	}
}
