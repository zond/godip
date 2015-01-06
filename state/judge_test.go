package state

import (
	"github.com/zond/godip/common"
	"github.com/zond/godip/graph"
	"testing"
	"time"
)

type testOrder int

func (self testOrder) Type() common.OrderType {
	return ""
}
func (self testOrder) Flags() map[common.Flag]bool {
	return nil
}
func (self testOrder) At() time.Time {
	return time.Now()
}
func (self testOrder) Targets() []common.Province {
	return nil
}
func (self testOrder) Adjudicate(common.Resolver) error {
	return nil
}
func (self testOrder) Validate(common.Validator) error {
	return nil
}
func (self testOrder) Execute(common.State) {
}

/*
     C
 A B
     D
*/
func testGraph() common.Graph {
	return graph.New().
		Prov("a").Conn("b").Conn("b/sc").Conn("b/nc").
		Prov("b").Conn("a").Conn("c").Conn("d").
		Prov("b/sc").Conn("a").Conn("d").
		Prov("b/nc").Conn("a").Conn("c").
		Prov("b/ec").Conn("c").Conn("d").
		Prov("c").Conn("b/nc").Conn("b/ec").
		Prov("d").Conn("b/sc").Conn("b/ec").
		Done()
}

func assertOrderLocation(t *testing.T, j *State, prov common.Province, order common.Order, ok bool) {
	if o, _, k := j.Order(prov); o != order || k != ok {
		t.Errorf("Wrong order, wanted %v, %v at %v but got %v, %v", order, ok, prov, o, k)
	}
}

func TestStateLocations(t *testing.T) {
	j := New(testGraph(), nil, nil)
	j.SetOrders(map[common.Province]common.Adjudicator{
		"a":    testOrder(1),
		"b/ec": testOrder(2),
	})
	j.SetOrders(map[common.Province]common.Adjudicator{
		"b": testOrder(2),
	})
	assertOrderLocation(t, j, "a", nil, false)
	assertOrderLocation(t, j, "b", testOrder(2), true)
	assertOrderLocation(t, j, "b/sc", testOrder(2), true)
	assertOrderLocation(t, j, "b/ec", testOrder(2), true)
	assertOrderLocation(t, j, "b/nc", testOrder(2), true)
}
