package orders

import (
	"github.com/zond/godip/orders"

	dip "github.com/zond/godip/common"
)

var generators = []func() dip.Order{
	BuildGenerator,
	ConvoyGenerator,
	DisbandGenerator,
	HoldGenerator,
	MoveGenerator,
	MoveViaConvoyGenerator,
	SupportGenerator,
}

var ClassicalParser = orders.NewParser(generators)
