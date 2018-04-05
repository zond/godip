package orders

import (
	"github.com/zond/godip/orders"

	dip "github.com/zond/godip/common"
)

var ClassicalParser = orders.NewParser([]dip.Order{
	orders.BuildOrder,
	orders.ConvoyOrder,
	orders.DisbandOrder,
	orders.HoldOrder,
	orders.MoveOrder,
	orders.MoveViaConvoyOrder,
	orders.SupportOrder,
})
