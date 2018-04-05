package orders

import (
	"fmt"

	"github.com/zond/godip"
)

type Parser struct {
	prototypes []godip.Order
}

func NewParser(prototypes []godip.Order) Parser {
	return Parser{prototypes}
}

func (self Parser) OrderTypes() (result []godip.OrderType) {
	result = make([]godip.OrderType, len(self.prototypes))
	for index, prototype := range self.prototypes {
		result[index] = prototype.DisplayType()
	}
	return
}

func (self Parser) Orders() (result []godip.Order) {
	result = make([]godip.Order, len(self.prototypes))
	copy(result, self.prototypes)
	return
}

func (self Parser) ParseAll(orders map[godip.Nation]map[godip.Province][]string) (result map[godip.Province]godip.Adjudicator, err error) {
	merr := MultiError{}
	result = map[godip.Province]godip.Adjudicator{}
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

func (self Parser) Parse(bits []string) (result godip.Adjudicator, err error) {
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
