package graph

import (
	"reflect"
	"testing"

	"github.com/zond/godip"
)

func assertPath(t *testing.T, g *Graph, first, last godip.Province, reverse bool, found []godip.Province) {
	if f := g.Path(first, last, reverse, nil); !reflect.DeepEqual(f, found) {
		t.Errorf("%v should have a path (reverse set to %v) between %v and %v like %v but found %v", g, reverse, first, last, found, f)
	}
}

func TestPath(t *testing.T) {
	g := New().
		Prov("a").Conn("f").Conn("h").
		Prov("b").Conn("g").Conn("c").
		Prov("c").Conn("b").Conn("h").Conn("d").Conn("i").
		Prov("d").Conn("c").Conn("h").Conn("e").
		Prov("e").Conn("d").Conn("g").Conn("f").Conn("j").
		Prov("f").Conn("a").Conn("e").
		Prov("g").Conn("b").Conn("h").Conn("e").
		Prov("h").Conn("a").Conn("c").Conn("d").Conn("g").
		Prov("i").Conn("c").
		Prov("j").Conn("g").
		Done()
	// The shortest path from a to e is via f.
	assertPath(t, g, "a", "e", false, []godip.Province{"f", "e"})
	// The shortest path from a to d is via h.
	assertPath(t, g, "a", "d", false, []godip.Province{"h", "d"})
	// The only edge to i is from c, and the shortest route to c is via h.
	assertPath(t, g, "a", "i", false, []godip.Province{"h", "c", "i"})

	// There is a directed edge from e to j.
	assertPath(t, g, "a", "j", false, []godip.Province{"f", "e", "j"})
	// There is a directed edge from j to g.
	assertPath(t, g, "j", "a", false, []godip.Province{"g", "h", "a"})

	// Check we can navigate both directed paths in reverse too.
	assertPath(t, g, "a", "j", true, []godip.Province{"h", "g", "j"})
	assertPath(t, g, "j", "a", true, []godip.Province{"e", "f", "a"})
}
