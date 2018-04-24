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
