package ancientmediterranean

import (
	"github.com/zond/godip/graph"
	"github.com/zond/godip/state"
	"github.com/zond/godip/variants/classical"
	"github.com/zond/godip/variants/classical/orders"
	"github.com/zond/godip/variants/common"

	cla "github.com/zond/godip/variants/classical/common"
	dip "github.com/zond/godip/common"
)

const (
	Rome     dip.Nation = "Rome"
	Carthage dip.Nation = "Carthage"
	Greece   dip.Nation = "Greece"
	Persia   dip.Nation = "Persia"
	Egypt    dip.Nation = "Egypt"
)

var AncientMediterraneanVariant = common.Variant{
	Name: "Ancient Mediterranean",
	Graph: func() dip.Graph { return AncientMediterraneanGraph() },
	Start: AncientMediterraneanStart,
	Blank: AncientMediterraneanBlank,
	Phase:             classical.Phase,
	ParseOrders:       orders.ParseAll,
	ParseOrder:        orders.Parse,
	OrderTypes:        orders.OrderTypes(),
	Nations:           []dip.Nation{Rome, Greece, Egypt, Persia, Carthage},
	PhaseTypes:        cla.PhaseTypes,
	Seasons:           cla.Seasons,
	UnitTypes:         cla.UnitTypes,
	SoloSupplyCenters: 18,
	SVGMap: func() ([]byte, error) {
		return Asset("svg/ancientmediterraneanmap.svg")
	},
	SVGVersion: "1",
	SVGUnits: map[dip.UnitType]func() ([]byte, error){
		cla.Army: func() ([]byte, error) {
			return classical.Asset("svg/army.svg")
		},
		cla.Fleet: func() ([]byte, error) {
			return classical.Asset("svg/fleet.svg")
		},
	},
	CreatedBy: "Don Hessong",
	Version: "",
	Description: "Five historical nations battle ",
	Rules: "",
}

func AncientMediterraneanBlank(phase dip.Phase) *state.State {
	return state.New(AncientMediterraneanGraph(), phase, classical.BackupRule)
}

func AncientMediterraneanStart() (result *state.State, err error) {
	if result, err = classical.Start(); err != nil {
		return
	}
	if err = result.SetUnits(map[dip.Province]dip.Unit{
		"nea": dip.Unit{cla.Fleet, Rome},
		"rom": dip.Unit{cla.Army, Rome},
		"rav": dip.Unit{cla.Army, Rome},
		"tha": dip.Unit{cla.Fleet, Carthage},
		"cir": dip.Unit{cla.Army, Carthage},
		"car": dip.Unit{cla.Army, Carthage},
		"spa": dip.Unit{cla.Fleet, Greece},
		"ath": dip.Unit{cla.Army, Greece},
		"mac": dip.Unit{cla.Army, Greece},
		"sid": dip.Unit{cla.Fleet, Persia},
		"ant": dip.Unit{cla.Army, Persia},
		"dam": dip.Unit{cla.Army, Persia},
		"ale": dip.Unit{cla.Fleet, Egypt},
		"mem": dip.Unit{cla.Army, Egypt},
		"the": dip.Unit{cla.Army, Egypt},
	}); err != nil {
		return
	}
	result.SetSupplyCenters(map[dip.Province]dip.Nation{
		"nea": Rome,
		"rom": Rome,
		"rav": Rome,
		"tha": Carthage,
		"cir": Carthage,
		"car": Carthage,
		"spa": Greece,
		"ath": Greece,
		"mac": Greece,
		"sid": Persia,
		"ant": Persia,
		"dam": Persia,
		"ale": Egypt,
		"mem": Egypt,
		"the": Egypt,
	})
	return
}

func AncientMediterraneanGraph() *graph.Graph {
	return graph.New().
		// gau
		Prov("gau").Conn("rha", cla.Land).Conn("mas", cla.Land).Conn("tar", cla.Land).Conn("lus", cla.Land).Flag(cla.Land).
		// rha
		Prov("rha").Conn("sam", cla.Land).Conn("vin", cla.Land).Conn("ven", cla.Land).Conn("etr", cla.Land).Conn("mas", cla.Land).Conn("gau", cla.Land).Flag(cla.Land).
		// sam
		Prov("sam").Conn("che", cla.Land).Conn("dac", cla.Land).Conn("ill", cla.Land).Conn("vin", cla.Land).Conn("rha", cla.Land).Flag(cla.Land).
		// che
		Prov("che").Conn("arm", cla.Land).Conn("sip", cla.Coast...).Conn("bla", cla.Sea).Conn("dac", cla.Coast...).Conn("sam", cla.Land).Flag(cla.Coast...).SC(cla.Neutral).
		// vin
		Prov("vin").Conn("sam", cla.Land).Conn("ill", cla.Land).Conn("dal", cla.Land).Conn("ven", cla.Land).Conn("rha", cla.Land).Flag(cla.Land).SC(cla.Neutral).
		// dac
		Prov("dac").Conn("che", cla.Coast...).Conn("bla",cla.Sea).Conn("byz", cla.Coast...).Conn("mac", cla.Land).Conn("ill", cla.Land).Conn("sam", cla.Land).Flag(cla.Coast...).
		// bla
		Prov("bla").Conn("che", cla.Sea).Conn("sip", cla.Sea).Conn("bit", cla.Sea).Conn("byz", cla.Sea).Conn("dac", cla.Sea).Flag(cla.Sea).
		// sip
		Prov("sip").Conn("che", cla.Coast...).Conn("arm", cla.Land).Conn("cap", cla.Land).Conn("gal", cla.Land).Conn("bit", cla.Coast...).Conn("bla", cla.Sea).Flag(cla.Coast...).SC(cla.Neutral).
		// arm
		Prov("arm").Conn("che", cla.Land).Conn("dam", cla.Land).Conn("cap", cla.Land).Conn("sip", cla.Land).Flag(cla.Land).
		// mas
		Prov("mas").Conn("gau", cla.Land).Conn("rha", cla.Land).Conn("etr", cla.Coast...).Conn("lig", cla.Sea).Conn("tar", cla.Coast...).Flag(cla.Coast...).SC(cla.Neutral).
		// etr
		Prov("etr").Conn("rha", cla.Land).Conn("ven", cla.Land).Conn("rav", cla.Land).Conn("rom", cla.Coast...).Conn("lig", cla.Sea).Conn("mas", cla.Coast...).Flag(cla.Coast...).
		// ven
		Prov("ven").Conn("rha", cla.Land).Conn("vin", cla.Land).Conn("dal", cla.Coast...).Conn("adr", cla.Sea).Conn("rav", cla.Coast...).Conn("etr", cla.Land).Flag(cla.Coast...).
		// dal
		Prov("dal").Conn("ill", cla.Land).Conn("epi", cla.Coast...).Conn("adr", cla.Sea).Conn("ven", cla.Coast...).Conn("vin", cla.Land).Flag(cla.Coast...).SC(cla.Neutral).
		// ill
		Prov("ill").Conn("sam", cla.Land).Conn("dac", cla.Land).Conn("mac", cla.Land).Conn("epi", cla.Land).Conn("dal", cla.Land).Conn("vin", cla.Land).Flag(cla.Land).
		// lus
		Prov("lus").Conn("gau", cla.Land).Conn("tar", cla.Land).Conn("sag", cla.Land).Flag(cla.Land).
		// tar
		Prov("tar").Conn("gau", cla.Land).Conn("mas", cla.Coast...).Conn("lig", cla.Sea).Conn("bal", cla.Sea).Conn("sag", cla.Coast...).Conn("lus", cla.Land).Flag(cla.Coast...).
		// sag
		Prov("sag").Conn("tar", cla.Coast...).Conn("bal", cla.Sea).Conn("ber", cla.Sea).Conn("ibe", cla.Sea).Conn("mau", cla.Coast...).Conn("lus", cla.Land).Flag(cla.Coast...).SC(cla.Neutral).
		// bal
		Prov("bal").Conn("lig", cla.Sea).Conn("ber", cla.Sea).Conn("sag", cla.Sea).Conn("tar", cla.Sea).Flag(cla.Coast...).SC(cla.Neutral).
		// rom
		Prov("rom").Conn("rav", cla.Land).Conn("apu", cla.Land).Conn("nea", cla.Coast...).Conn("tys", cla.Sea).Conn("lig", cla.Sea).Conn("etr", cla.Coast...).Flag(cla.Coast...).SC(Rome).
		// rav
		Prov("rav").Conn("ven", cla.Coast...).Conn("adr", cla.Sea).Conn("apu", cla.Coast...).Conn("rom", cla.Land).Conn("etr", cla.Land).Flag(cla.Coast...).SC(Rome).
		// apu
		Prov("apu").Conn("adr", cla.Sea).Conn("ion", cla.Sea).Conn("nea", cla.Coast...).Conn("rom", cla.Land).Conn("rav", cla.Coast...).Flag(cla.Coast...).
		// nea
		Prov("nea").Conn("apu", cla.Coast...).Conn("ion", cla.Sea).Conn("aus", cla.Sea).Conn("sic", cla.Coast...).Conn("tys", cla.Sea).Conn("rom", cla.Coast...).Flag(cla.Coast...).SC(Rome).
		// sic
		// epi
		// mac
		// ath
		// byz
		// bit
		// mil
		// gal
		// isa
		// cap
		// ant
		// dam
		// sid
		// tye
		// ara
		// jer
		// nab
		// pet
		// sii
		// the
		// ale
		// mem
		// bay
		// cyr
		// mar
		// lep
		// pha
		// sah
		// num
		// cir
		// mau
		// car
		// tha
		// cor
		// sad
		// cre
		// cyp
		// ibe
		// ber
		// lig
		// tys
		// pun
		// adr
		// ion
		// aus
		// got
		// gos
		// lib
		// mes
		// aeg
		// min
		// egy
		// cil
		// syr
		// gop
		// ree
		Done()
}
