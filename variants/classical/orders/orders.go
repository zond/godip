package orders

import (
	"github.com/zond/godip/orders"

	dip "github.com/zond/godip/common"
)

var Parser = orders.NewParser([]dip.Order{
	BuildOrder,
	ConvoyOrder,
	DisbandOrder,
	HoldOrder,
	MoveOrder,
	MoveViaConvoyOrder,
	SupportOrder,
})
