package orders

import (
	"fmt"

	dip "github.com/zond/godip/common"
)

var orderTypes []dip.OrderType

var generators []func() dip.Order

func OrderTypes() (result []dip.OrderType) {
	result = make([]dip.OrderType, len(generators))
	for index, gen := range generators {
		result[index] = gen().DisplayType()
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
	return fmt.Sprint([]error(self))
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
	for _, generator := range generators {
		result, err := generator().Parse(bits)
		if result != nil || err != nil {
			return result, err
		}
	}
	return nil, fmt.Errorf("Invalid order %+v", bits)
}
