package ancientmediterranean

import (
	"github.com/zond/godip"
	"github.com/zond/godip/graph"
	"github.com/zond/godip/state"
	"github.com/zond/godip/variants/classical"
	"github.com/zond/godip/variants/classical/orders"
	"github.com/zond/godip/variants/common"

	cla "github.com/zond/godip/variants/classical/common"
)

const (
	Rome     godip.Nation = "Rome"
	Carthage godip.Nation = "Carthage"
	Greece   godip.Nation = "Greece"
	Persia   godip.Nation = "Persia"
	Egypt    godip.Nation = "Egypt"
)

var Nations = []godip.Nation{Rome, Greece, Egypt, Persia, Carthage}

var AncientMediterraneanVariant = common.Variant{
	Name:       "Ancient Mediterranean",
	Graph:      func() godip.Graph { return AncientMediterraneanGraph() },
	Start:      AncientMediterraneanStart,
	Blank:      AncientMediterraneanBlank,
	Phase:      classical.Phase,
	Parser:     orders.ClassicalParser,
	Nations:    Nations,
	PhaseTypes: cla.PhaseTypes,
	Seasons:    cla.Seasons,
	UnitTypes:  cla.UnitTypes,
	SoloWinner: common.SCCountWinner(18),
	SVGMap: func() ([]byte, error) {
		return Asset("svg/ancientmediterraneanmap.svg")
	},
	SVGVersion: "2",
	SVGUnits: map[godip.UnitType]func() ([]byte, error){
		cla.Army: func() ([]byte, error) {
			return classical.Asset("svg/army.svg")
		},
		cla.Fleet: func() ([]byte, error) {
			return classical.Asset("svg/fleet.svg")
		},
	},
	CreatedBy:   "Don Hessong",
	Version:     "",
	Description: "Five historical nations battle for dominance of the Mediterranean.",
	Rules: "Rules are as per classical Diplomacy, with a few parts of the map that have noteworthy connectivity. " +
		"Baleares is an archipelago that can be occupied by armies or fleets. Armies may not move directly from " +
		"the mainland to Baleares, and a fleet in Baleares is able to form part of a convoy chain. " +
		"The canal between Athens and Sparta is passable for armies, and means that Athens only has a single " +
		"coast. Similarly the canal in Byzantium, the Sicilian Straits and the River Nile. There is a four way " +
		"connection between the Ausonian Sea, Messenian Sea, Gulf of Tacape and Libyan Sea. There is another " +
		"four-way connection between Alexandria, Sinai, Thebes and the Gulf of Pelusium. The first to 18 supply " +
		"centers is the winner.",
}

func AncientMediterraneanBlank(phase godip.Phase) *state.State {
	return state.New(AncientMediterraneanGraph(), phase, classical.BackupRule)
}

func AncientMediterraneanStart() (result *state.State, err error) {
	startPhase := classical.Phase(1, cla.Spring, cla.Movement)
	result = state.New(AncientMediterraneanGraph(), startPhase, classical.BackupRule)
	if err = result.SetUnits(map[godip.Province]godip.Unit{
		"nea": godip.Unit{cla.Fleet, Rome},
		"rom": godip.Unit{cla.Army, Rome},
		"rav": godip.Unit{cla.Army, Rome},
		"tha": godip.Unit{cla.Fleet, Carthage},
		"cir": godip.Unit{cla.Army, Carthage},
		"car": godip.Unit{cla.Army, Carthage},
		"spa": godip.Unit{cla.Fleet, Greece},
		"ath": godip.Unit{cla.Army, Greece},
		"mac": godip.Unit{cla.Army, Greece},
		"sid": godip.Unit{cla.Fleet, Persia},
		"ant": godip.Unit{cla.Army, Persia},
		"dam": godip.Unit{cla.Army, Persia},
		"ale": godip.Unit{cla.Fleet, Egypt},
		"mem": godip.Unit{cla.Army, Egypt},
		"the": godip.Unit{cla.Army, Egypt},
	}); err != nil {
		return
	}
	result.SetSupplyCenters(map[godip.Province]godip.Nation{
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
		Prov("dac").Conn("che", cla.Coast...).Conn("bla", cla.Sea).Conn("byz", cla.Coast...).Conn("mac", cla.Land).Conn("ill", cla.Land).Conn("sam", cla.Land).Flag(cla.Coast...).
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
		Prov("bal").Conn("lig", cla.Sea).Conn("ber", cla.Sea).Conn("sag", cla.Sea).Conn("tar", cla.Sea).Flag(cla.Archipelago...).SC(cla.Neutral).
		// rom
		Prov("rom").Conn("rav", cla.Land).Conn("apu", cla.Land).Conn("nea", cla.Coast...).Conn("tys", cla.Sea).Conn("lig", cla.Sea).Conn("etr", cla.Coast...).Flag(cla.Coast...).SC(Rome).
		// rav
		Prov("rav").Conn("ven", cla.Coast...).Conn("adr", cla.Sea).Conn("apu", cla.Coast...).Conn("rom", cla.Land).Conn("etr", cla.Land).Flag(cla.Coast...).SC(Rome).
		// apu
		Prov("apu").Conn("adr", cla.Sea).Conn("ion", cla.Sea).Conn("nea", cla.Coast...).Conn("rom", cla.Land).Conn("rav", cla.Coast...).Flag(cla.Coast...).
		// nea
		Prov("nea").Conn("apu", cla.Coast...).Conn("ion", cla.Sea).Conn("aus", cla.Sea).Conn("sic", cla.Coast...).Conn("tys", cla.Sea).Conn("rom", cla.Coast...).Flag(cla.Coast...).SC(Rome).
		// sic
		Prov("sic").Conn("tys", cla.Sea).Conn("nea", cla.Coast...).Conn("aus", cla.Sea).Conn("pun", cla.Sea).Flag(cla.Coast...).SC(cla.Neutral).
		// epi
		Prov("epi").Conn("ill", cla.Land).Conn("mac", cla.Land).Conn("ath", cla.Coast...).Conn("ion", cla.Sea).Conn("adr", cla.Sea).Conn("dal", cla.Coast...).Flag(cla.Coast...).
		// mac
		Prov("mac").Conn("dac", cla.Land).Conn("byz", cla.Coast...).Conn("aeg", cla.Sea).Conn("ath", cla.Coast...).Conn("epi", cla.Land).Conn("ill", cla.Land).Flag(cla.Coast...).SC(Greece).
		// ath
		Prov("ath").Conn("mac", cla.Coast...).Conn("aeg", cla.Sea).Conn("spa", cla.Coast...).Conn("ion", cla.Sea).Conn("spa", cla.Coast...).Conn("epi", cla.Coast...).Flag(cla.Coast...).SC(Greece).
		// spa
		Prov("spa").Conn("ath", cla.Coast...).Conn("aeg", cla.Sea).Conn("mes", cla.Sea).Conn("ion", cla.Sea).Flag(cla.Coast...).SC(Greece).
		// byz
		Prov("byz").Conn("bla", cla.Sea).Conn("bit", cla.Coast...).Conn("gal", cla.Land).Conn("mil", cla.Coast...).Conn("aeg", cla.Sea).Conn("mac", cla.Coast...).Conn("dac", cla.Coast...).Flag(cla.Coast...).SC(cla.Neutral).
		// bit
		Prov("bit").Conn("bla", cla.Sea).Conn("sip", cla.Coast...).Conn("gal", cla.Land).Conn("byz", cla.Coast...).Flag(cla.Coast...).
		// mil
		Prov("mil").Conn("byz", cla.Coast...).Conn("gal", cla.Land).Conn("isa", cla.Coast...).Conn("cil", cla.Sea).Conn("min", cla.Sea).Conn("aeg", cla.Sea).Flag(cla.Coast...).SC(cla.Neutral).
		// gal
		Prov("gal").Conn("sip", cla.Land).Conn("cap", cla.Land).Conn("isa", cla.Land).Conn("mil", cla.Land).Conn("byz", cla.Land).Conn("bit", cla.Land).Flag(cla.Land).
		// isa
		Prov("isa").Conn("gal", cla.Land).Conn("cap", cla.Coast...).Conn("cil", cla.Sea).Conn("min", cla.Sea).Conn("aeg", cla.Sea).Conn("mil", cla.Coast...).Flag(cla.Coast...).
		// cap
		Prov("cap").Conn("sip", cla.Land).Conn("arm", cla.Land).Conn("dam", cla.Land).Conn("ant", cla.Coast...).Conn("cil", cla.Sea).Conn("isa", cla.Coast...).Conn("gal", cla.Land).Flag(cla.Coast...).
		// ant
		Prov("ant").Conn("cap", cla.Coast...).Conn("dam", cla.Land).Conn("sid", cla.Coast...).Conn("cil", cla.Sea).Flag(cla.Coast...).SC(Persia).
		// dam
		Prov("dam").Conn("arm", cla.Land).Conn("ara", cla.Land).Conn("sid", cla.Land).Conn("ant", cla.Land).Conn("cap", cla.Land).Flag(cla.Land).SC(Persia).
		// sid
		Prov("sid").Conn("ant", cla.Coast...).Conn("dam", cla.Land).Conn("ara", cla.Land).Conn("tye", cla.Coast...).Conn("syr", cla.Sea).Conn("cil", cla.Sea).Flag(cla.Coast...).SC(Persia).
		// tye
		Prov("tye").Conn("sid", cla.Coast...).Conn("ara", cla.Land).Conn("jer", cla.Coast...).Conn("syr", cla.Sea).Flag(cla.Coast...).SC(cla.Neutral).
		// ara
		Prov("ara").Conn("dam", cla.Land).Conn("nab", cla.Land).Conn("jer", cla.Land).Conn("tye", cla.Land).Conn("sid", cla.Land).Flag(cla.Land).
		// jer
		Prov("jer").Conn("tye", cla.Coast...).Conn("ara", cla.Land).Conn("nab", cla.Land).Conn("pet", cla.Land).Conn("sii", cla.Coast...).Conn("gop", cla.Sea).Conn("syr", cla.Sea).Flag(cla.Coast...).SC(cla.Neutral).
		// nab
		Prov("nab").Conn("ara", cla.Land).Conn("ree", cla.Sea).Conn("pet", cla.Coast...).Conn("jer", cla.Land).Flag(cla.Coast...).
		// pet
		Prov("pet").Conn("nab", cla.Coast...).Conn("ree", cla.Sea).Conn("sii", cla.Coast...).Conn("jer", cla.Land).Flag(cla.Coast...).SC(cla.Neutral).
		// sii
		Prov("sii").Conn("jer", cla.Coast...).Conn("pet", cla.Coast...).Conn("ree", cla.Sea).Conn("the", cla.Coast...).Conn("ale", cla.Coast...).Conn("gop", cla.Sea).Flag(cla.Coast...).
		// the
		Prov("the").Conn("sii", cla.Coast...).Conn("ree", cla.Sea).Conn("bay", cla.Coast...).Conn("mem", cla.Coast...).Conn("ale", cla.Coast...).Conn("gop", cla.Sea).Flag(cla.Coast...).SC(Egypt).
		// ale
		Prov("ale").Conn("egy", cla.Sea).Conn("gop", cla.Sea).Conn("sii", cla.Coast...).Conn("the", cla.Coast...).Conn("mem", cla.Coast...).Conn("cyr", cla.Coast...).Conn("lib", cla.Sea).Flag(cla.Coast...).SC(Egypt).
		// mem
		Prov("mem").Conn("ale", cla.Coast...).Conn("the", cla.Coast...).Conn("bay", cla.Coast...).Conn("mar", cla.Land).Conn("cyr", cla.Land).Flag(cla.Coast...).SC(Egypt).
		// bay
		Prov("bay").Conn("mem", cla.Coast...).Conn("the", cla.Coast...).Conn("sah", cla.Land).Conn("pha", cla.Land).Conn("mar", cla.Land).Flag(cla.Coast...).
		// cyr
		Prov("cyr").Conn("lib", cla.Sea).Conn("ale", cla.Coast...).Conn("mem", cla.Land).Conn("mar", cla.Land).Conn("lep", cla.Coast...).Conn("gos", cla.Sea).Flag(cla.Coast...).SC(cla.Neutral).
		// mar
		Prov("mar").Conn("cyr", cla.Land).Conn("mem", cla.Land).Conn("bay", cla.Land).Conn("pha", cla.Land).Conn("lep", cla.Land).Flag(cla.Land).
		// lep
		Prov("lep").Conn("gos", cla.Sea).Conn("cyr", cla.Coast...).Conn("mar", cla.Land).Conn("pha", cla.Land).Conn("num", cla.Coast...).Conn("got", cla.Sea).Flag(cla.Coast...).SC(cla.Neutral).
		// pha
		Prov("pha").Conn("num", cla.Land).Conn("lep", cla.Land).Conn("mar", cla.Land).Conn("bay", cla.Land).Conn("sah", cla.Land).Conn("cir", cla.Land).Flag(cla.Land).
		// sah
		Prov("sah").Conn("cir", cla.Land).Conn("pha", cla.Land).Conn("bay", cla.Land).Conn("mau", cla.Land).Flag(cla.Land).
		// num
		Prov("num").Conn("got", cla.Sea).Conn("lep", cla.Coast...).Conn("pha", cla.Land).Conn("cir", cla.Land).Conn("tha", cla.Coast...).Flag(cla.Coast...).SC(cla.Neutral).
		// cir
		Prov("cir").Conn("car", cla.Land).Conn("tha", cla.Land).Conn("num", cla.Land).Conn("pha", cla.Land).Conn("sah", cla.Land).Conn("mau", cla.Land).Flag(cla.Land).SC(Carthage).
		// mau
		Prov("mau").Conn("ibe", cla.Sea).Conn("ber", cla.Sea).Conn("car", cla.Coast...).Conn("cir", cla.Land).Conn("sah", cla.Land).Conn("sag", cla.Coast...).Flag(cla.Coast...).
		// car
		Prov("car").Conn("pun", cla.Sea).Conn("tha", cla.Coast...).Conn("cir", cla.Land).Conn("mau", cla.Coast...).Conn("ber", cla.Sea).Flag(cla.Coast...).SC(Carthage).
		// tha
		Prov("tha").Conn("car", cla.Coast...).Conn("pun", cla.Sea).Conn("got", cla.Sea).Conn("num", cla.Coast...).Conn("cir", cla.Land).Flag(cla.Coast...).SC(Carthage).
		// cor
		Prov("cor").Conn("lig", cla.Sea).Conn("tys", cla.Sea).Conn("sad", cla.Coast...).Flag(cla.Coast...).
		// sad
		Prov("sad").Conn("cor", cla.Coast...).Conn("tys", cla.Sea).Conn("pun", cla.Sea).Conn("ber", cla.Sea).Conn("lig", cla.Sea).Flag(cla.Coast...).SC(cla.Neutral).
		// cre
		Prov("cre").Conn("aeg", cla.Sea).Conn("min", cla.Sea).Conn("egy", cla.Sea).Conn("lib", cla.Sea).Conn("mes", cla.Sea).Flag(cla.Coast...).SC(cla.Neutral).
		// cyp
		Prov("cyp").Conn("cil", cla.Sea).Conn("syr", cla.Sea).Conn("egy", cla.Sea).Flag(cla.Coast...).SC(cla.Neutral).
		// ibe
		Prov("ibe").Conn("sag", cla.Sea).Conn("ber", cla.Sea).Conn("mau", cla.Sea).Flag(cla.Sea).
		// ber
		Prov("ber").Conn("lig", cla.Sea).Conn("sad", cla.Sea).Conn("pun", cla.Sea).Conn("car", cla.Sea).Conn("mau", cla.Sea).Conn("ibe", cla.Sea).Conn("sag", cla.Sea).Conn("bal", cla.Sea).Flag(cla.Sea).
		// lig
		Prov("lig").Conn("etr", cla.Sea).Conn("rom", cla.Sea).Conn("tys", cla.Sea).Conn("cor", cla.Sea).Conn("sad", cla.Sea).Conn("ber", cla.Sea).Conn("bal", cla.Sea).Conn("tar", cla.Sea).Conn("mas", cla.Sea).Flag(cla.Sea).
		// tys
		Prov("tys").Conn("rom", cla.Sea).Conn("nea", cla.Sea).Conn("aus", cla.Sea).Conn("sic", cla.Sea).Conn("pun", cla.Sea).Conn("sad", cla.Sea).Conn("cor", cla.Sea).Conn("lig", cla.Sea).Flag(cla.Sea).
		// pun
		Prov("pun").Conn("tys", cla.Sea).Conn("sic", cla.Sea).Conn("aus", cla.Sea).Conn("got", cla.Sea).Conn("tha", cla.Sea).Conn("car", cla.Sea).Conn("ber", cla.Sea).Conn("sad", cla.Sea).Flag(cla.Sea).
		// adr
		Prov("adr").Conn("dal", cla.Sea).Conn("epi", cla.Sea).Conn("ion", cla.Sea).Conn("apu", cla.Sea).Conn("rav", cla.Sea).Conn("ven", cla.Sea).Flag(cla.Sea).
		// ion
		Prov("ion").Conn("adr", cla.Sea).Conn("epi", cla.Sea).Conn("ath", cla.Sea).Conn("spa", cla.Sea).Conn("mes", cla.Sea).Conn("aus", cla.Sea).Conn("nea", cla.Sea).Conn("apu", cla.Sea).Flag(cla.Sea).
		// aus
		Prov("aus").Conn("ion", cla.Sea).Conn("mes", cla.Sea).Conn("lib", cla.Sea).Conn("got", cla.Sea).Conn("pun", cla.Sea).Conn("sic", cla.Sea).Conn("tys", cla.Sea).Conn("nea", cla.Sea).Flag(cla.Sea).
		// got
		Prov("got").Conn("aus", cla.Sea).Conn("mes", cla.Sea).Conn("lib", cla.Sea).Conn("gos", cla.Sea).Conn("lep", cla.Sea).Conn("num", cla.Sea).Conn("tha", cla.Sea).Conn("pun", cla.Sea).Flag(cla.Sea).
		// gos
		Prov("gos").Conn("lib", cla.Sea).Conn("cyr", cla.Sea).Conn("lep", cla.Sea).Conn("got", cla.Sea).Flag(cla.Sea).
		// lib
		Prov("lib").Conn("mes", cla.Sea).Conn("cre", cla.Sea).Conn("egy", cla.Sea).Conn("ale", cla.Sea).Conn("cyr", cla.Sea).Conn("gos", cla.Sea).Conn("got", cla.Sea).Conn("aus", cla.Sea).Flag(cla.Sea).
		// mes
		Prov("mes").Conn("spa", cla.Sea).Conn("aeg", cla.Sea).Conn("cre", cla.Sea).Conn("lib", cla.Sea).Conn("got", cla.Sea).Conn("aus", cla.Sea).Conn("ion", cla.Sea).Flag(cla.Sea).
		// aeg
		Prov("aeg").Conn("byz", cla.Sea).Conn("mil", cla.Sea).Conn("min", cla.Sea).Conn("cre", cla.Sea).Conn("mes", cla.Sea).Conn("spa", cla.Sea).Conn("ath", cla.Sea).Conn("mac", cla.Sea).Flag(cla.Sea).
		// min
		Prov("min").Conn("mil", cla.Sea).Conn("cil", cla.Sea).Conn("egy", cla.Sea).Conn("cre", cla.Sea).Conn("aeg", cla.Sea).Flag(cla.Sea).
		// egy
		Prov("egy").Conn("cil", cla.Sea).Conn("cyp", cla.Sea).Conn("syr", cla.Sea).Conn("gop", cla.Sea).Conn("ale", cla.Sea).Conn("lib", cla.Sea).Conn("cre", cla.Sea).Conn("min", cla.Sea).Flag(cla.Sea).
		// cil
		Prov("cil").Conn("cap", cla.Sea).Conn("ant", cla.Sea).Conn("sid", cla.Sea).Conn("syr", cla.Sea).Conn("cyp", cla.Sea).Conn("egy", cla.Sea).Conn("min", cla.Sea).Conn("mil", cla.Sea).Conn("isa", cla.Sea).Flag(cla.Sea).
		// syr
		Prov("syr").Conn("cil", cla.Sea).Conn("sid", cla.Sea).Conn("tye", cla.Sea).Conn("jer", cla.Sea).Conn("gop", cla.Sea).Conn("egy", cla.Sea).Conn("cyp", cla.Sea).Flag(cla.Sea).
		// gop
		Prov("gop").Conn("syr", cla.Sea).Conn("jer", cla.Sea).Conn("sii", cla.Sea).Conn("the", cla.Sea).Conn("ale", cla.Sea).Conn("egy", cla.Sea).Flag(cla.Sea).
		// ree
		Prov("ree").Conn("pet", cla.Sea).Conn("nab", cla.Sea).Conn("the", cla.Sea).Conn("sii", cla.Sea).Flag(cla.Sea).
		Done()
}
