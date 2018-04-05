package state

import (
	"testing"
	"time"

	"github.com/zond/godip"
	"github.com/zond/godip/graph"
)

type testOrder int

func (self testOrder) Options(v godip.Validator, nation godip.Nation, src godip.Province) (result godip.Options) {
	return nil
}
func (self testOrder) DisplayType() godip.OrderType {
	return ""
}
func (self testOrder) Type() godip.OrderType {
	return ""
}
func (self testOrder) Flags() map[godip.Flag]bool {
	return nil
}
func (self testOrder) Parse(parts []string) (godip.Adjudicator, error) {
	return nil, nil
}
func (self testOrder) At() time.Time {
	return time.Now()
}
func (self testOrder) Targets() []godip.Province {
	return nil
}
func (self testOrder) Adjudicate(godip.Resolver) error {
	return nil
}
func (self testOrder) Validate(godip.Validator) (godip.Nation, error) {
	return "", nil
}
func (self testOrder) Execute(godip.State) {
}

/*
     C
 A B
     D
*/
func testGraph() godip.Graph {
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

func assertOrderLocation(t *testing.T, j *State, prov godip.Province, order godip.Order, ok bool) {
	if o, _, k := j.Order(prov); o != order || k != ok {
		t.Errorf("Wrong order, wanted %v, %v at %v but got %v, %v", order, ok, prov, o, k)
	}
}

func TestStateLocations(t *testing.T) {
	j := New(testGraph(), nil, nil)
	j.SetOrders(map[godip.Province]godip.Adjudicator{
		"a":    testOrder(1),
		"b/ec": testOrder(2),
	})
	j.SetOrders(map[godip.Province]godip.Adjudicator{
		"b": testOrder(2),
	})
	assertOrderLocation(t, j, "a", nil, false)
	assertOrderLocation(t, j, "b", testOrder(2), true)
	assertOrderLocation(t, j, "b/sc", testOrder(2), true)
	assertOrderLocation(t, j, "b/ec", testOrder(2), true)
	assertOrderLocation(t, j, "b/nc", testOrder(2), true)
}
