package graph

import (
	"reflect"
	"testing"

	"github.com/zond/godip"
)

func assertPath(t *testing.T, g *Graph, src, dst godip.Province, found []godip.Province) {
	if f := g.Path(src, dst, nil); !reflect.DeepEqual(f, found) {
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
	assertPath(t, g, "a", "e", []godip.Province{"f", "e"})
	assertPath(t, g, "a", "d", []godip.Province{"h", "d"})
	assertPath(t, g, "a", "i", []godip.Province{"h", "c", "i"})
}
