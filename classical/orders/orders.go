package orders

import (
	"fmt"
	"time"

	cla "github.com/zond/godip/classical/common"
	dip "github.com/zond/godip/common"
)

var orderTypes []dip.OrderType

var generators []func() dip.Order

func OrderTypes() (result []dip.OrderType) {
	result = make([]dip.OrderType, len(generators))
	for index, gen := range generators {
		result[index] = gen().Type()
	}
	return
}

func Orders() (result []dip.Order) {
	result = make([]dip.Order, len(generators))
	for index, gen := range generators {
		result[index] = gen()
	}
	return
}

type MultiError []error

func (self MultiError) Error() string {
	return fmt.Sprint(self)
}

func ParseAll(orders map[dip.Nation]map[dip.Province][]string) (result map[dip.Province]dip.Adjudicator, err error) {
	merr := MultiError{}
	result = map[dip.Province]dip.Adjudicator{}
	for _, nationOrders := range orders {
		for prov, bits := range nationOrders {
			if parsed, e := Parse(append([]string{string(prov)}, bits...)); e == nil {
				result[prov] = parsed
			} else {
				merr = append(merr, e)
			}
		}
	}
	if len(merr) > 0 {
		err = merr
	}
	return
}

func Parse(bits []string) (result dip.Adjudicator, err error) {
	if len(bits) > 1 {
		switch dip.OrderType(bits[1]) {
		case (&build{}).DisplayType():
			if len(bits) == 3 {
				result = Build(dip.Province(bits[0]), dip.UnitType(bits[2]), time.Now())
			}
		case (&convoy{}).DisplayType():
			if len(bits) == 4 {
				result = Convoy(dip.Province(bits[0]), dip.Province(bits[2]), dip.Province(bits[3]))
			}
		case (&disband{}).DisplayType():
			if len(bits) == 2 {
				result = Disband(dip.Province(bits[0]), time.Now())
			}
		case (&hold{}).DisplayType():
			if len(bits) == 2 {
				result = Hold(dip.Province(bits[0]))
			}
		case (&move{}).DisplayType():
			if len(bits) == 3 {
				result = Move(dip.Province(bits[0]), dip.Province(bits[2]))
			}
		case (&move{flags: map[dip.Flag]bool{cla.ViaConvoy: true}}).DisplayType():
			if len(bits) == 3 {
				result = Move(dip.Province(bits[0]), dip.Province(bits[2])).ViaConvoy()
			}
		case (&support{}).DisplayType():
			if len(bits) == 4 {
				if bits[2] == bits[3] {
					result = SupportHold(dip.Province(bits[0]), dip.Province(bits[2]))
				} else {
					result = SupportMove(dip.Province(bits[0]), dip.Province(bits[2]), dip.Province(bits[3]))
				}
			}
		}
	}
	if result == nil {
		err = fmt.Errorf("Invalid order %+v", bits)
	}
	return
}
