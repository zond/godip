package godip

import (
	"sort"
	"testing"
)

func TestNationSorting(t *testing.T) {
	nations := Nations{
		Austria,
		Russia,
		England,
	}
	sort.Sort(nations)
	if nations[0] != Austria {
		t.Errorf("Wanted Austria, got %v", nations[0])
	}
	if nations[1] != England {
		t.Errorf("Wanted England, got %v", nations[1])
	}
	if nations[2] != Russia {
		t.Errorf("Wanted Russia, got %v", nations[2])
	}
}

func TestNationEqual(t *testing.T) {
	n := Nations{
		Austria,
		Russia,
		England,
	}
	o := Nations{
		Austria,
		Russia,
		England,
	}
	if !n.Equal(o) {
		t.Errorf("Wanted equal, wasn't: %+v, %+v", n, o)
	}
	o = Nations{
		Russia,
		Austria,
		England,
	}
	if !n.Equal(o) {
		t.Errorf("Wanted equal, wasn't: %+v, %+v", n, o)
	}
	if n[0] != Austria || n[1] != Russia || n[2] != England || o[0] != Russia || o[1] != Austria || o[2] != England {
		t.Errorf("Changed order during equal, not OK")
	}
	o = Nations{
		Austria,
		Turkey,
		England,
	}
	if n.Equal(o) {
		t.Errorf("Wanted inequal, was: %+v, %+v", n, o)
	}
	o = Nations{
		Austria,
		Russia,
	}
	if n.Equal(o) {
		t.Errorf("Wanted inequal, was: %+v, %+v", n, o)
	}
}
