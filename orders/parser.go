package orders

import (
	"fmt"

	dip "github.com/zond/godip"
)

type Parser struct {
	prototypes []dip.Order
}

func NewParser(prototypes []dip.Order) Parser {
	return Parser{prototypes}
}

func (self Parser) OrderTypes() (result []dip.OrderType) {
	result = make([]dip.OrderType, len(self.prototypes))
	for index, prototype := range self.prototypes {
		result[index] = prototype.DisplayType()
	}
	return
}

func (self Parser) Orders() (result []dip.Order) {
	result = make([]dip.Order, len(self.prototypes))
	copy(result, self.prototypes)
	return
}

func (self Parser) ParseAll(orders map[dip.Nation]map[dip.Province][]string) (result map[dip.Province]dip.Adjudicator, err error) {
	merr := MultiError{}
	result = map[dip.Province]dip.Adjudicator{}
	for _, nationOrders := range orders {
		for prov, bits := range nationOrders {
			if parsed, e := self.Parse(append([]string{string(prov)}, bits...)); e == nil {
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

func (self Parser) Parse(bits []string) (result dip.Adjudicator, err error) {
	for _, prototype := range self.Orders() {
		result, err := prototype.Parse(bits)
		if result != nil || err != nil {
			return result, err
		}
	}
	return nil, fmt.Errorf("Invalid order %+v", bits)
}

type MultiError []error

func (self MultiError) Error() string {
	return fmt.Sprint([]error(self))
}
