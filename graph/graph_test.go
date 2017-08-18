package graph

import (
	"github.com/zond/godip/common"
	"reflect"
	"testing"
)

func assertPath(t *testing.T, g *Graph, src, dst common.Province, found []common.Province) {
	if f := g.Path(src, dst, nil, false); !reflect.DeepEqual(f, found) {
		t.Errorf("%v should have a path between %v and %v like %v but found %v", g, src, dst, found, f)
	}
}

func TestPath(t *testing.T) {
	g := New().
		Prov("a").Conn("f").Conn("h").
		Prov("b").Conn("g").Conn("c").
		Prov("c").Conn("b").Conn("h").Conn("d").Conn("i").
		Prov("d").Conn("c").Conn("h").Conn("e").
		Prov("e").Conn("d").Conn("g").Conn("f").
		Prov("f").Conn("a").Conn("e").
		Prov("g").Conn("b").Conn("h").Conn("e").
		Prov("h").Conn("a").Conn("c").Conn("d").Conn("g").
		Prov("i").Conn("c").
		Done()
	assertPath(t, g, "a", "e", []common.Province{"f", "e"})
	assertPath(t, g, "a", "d", []common.Province{"h", "d"})
	assertPath(t, g, "a", "i", []common.Province{"h", "c", "i"})
}
