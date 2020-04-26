package files

import (
	"strings"
	"testing"
)

func Test_list_register(t *testing.T) {
	l := newList()
	recs := [...]record{
		record{path: "a", size: 10},
		record{path: "b", size: 10},
		record{path: "c", size: 20},
	}
	for _, rec := range recs {
		l.register(rec)
	}
	result10 := strings.Join(l.bySize[10], "")
	expected10 := "ab"
	if result10 != expected10 {
		t.Errorf("expected %s, got %s", result10, expected10)
	}
}
