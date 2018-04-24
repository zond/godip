package ancientmediterranean

import (
	"github.com/zond/godip"
	"github.com/zond/godip/graph"
	"github.com/zond/godip/state"
	"github.com/zond/godip/variants/classical"
	"github.com/zond/godip/variants/classical/orders"
	"github.com/zond/godip/variants/common"
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
	PhaseTypes: classical.PhaseTypes,
	Seasons:    classical.Seasons,
	UnitTypes:  classical.UnitTypes,
	SoloWinner: common.SCCountWinner(18),
	SVGMap: func() ([]byte, error) {
		return Asset("svg/ancientmediterraneanmap.svg")
	},
	SVGVersion: "2",
	SVGUnits: map[godip.UnitType]func() ([]byte, error){
		godip.Army: func() ([]byte, error) {
			return classical.Asset("svg/army.svg")
		},
		godip.Fleet: func() ([]byte, error) {
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
	startPhase := classical.Phase(1, godip.Spring, godip.Movement)
	result = state.New(AncientMediterraneanGraph(), startPhase, classical.BackupRule)
	if err = result.SetUnits(map[godip.Province]godip.Unit{
		"nea": godip.Unit{godip.Fleet, Rome},
		"rom": godip.Unit{godip.Army, Rome},
		"rav": godip.Unit{godip.Army, Rome},
		"tha": godip.Unit{godip.Fleet, Carthage},
		"cir": godip.Unit{godip.Army, Carthage},
		"car": godip.Unit{godip.Army, Carthage},
		"spa": godip.Unit{godip.Fleet, Greece},
		"ath": godip.Unit{godip.Army, Greece},
		"mac": godip.Unit{godip.Army, Greece},
		"sid": godip.Unit{godip.Fleet, Persia},
		"ant": godip.Unit{godip.Army, Persia},
		"dam": godip.Unit{godip.Army, Persia},
		"ale": godip.Unit{godip.Fleet, Egypt},
		"mem": godip.Unit{godip.Army, Egypt},
		"the": godip.Unit{godip.Army, Egypt},
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
		Prov("gau").Conn("rha", godip.Land).Conn("mas", godip.Land).Conn("tar", godip.Land).Conn("lus", godip.Land).Flag(godip.Land).
		// rha
		Prov("rha").Conn("sam", godip.Land).Conn("vin", godip.Land).Conn("ven", godip.Land).Conn("etr", godip.Land).Conn("mas", godip.Land).Conn("gau", godip.Land).Flag(godip.Land).
		// sam
		Prov("sam").Conn("che", godip.Land).Conn("dac", godip.Land).Conn("ill", godip.Land).Conn("vin", godip.Land).Conn("rha", godip.Land).Flag(godip.Land).
		// che
		Prov("che").Conn("arm", godip.Land).Conn("sip", godip.Coast...).Conn("bla", godip.Sea).Conn("dac", godip.Coast...).Conn("sam", godip.Land).Flag(godip.Coast...).SC(godip.Neutral).
		// vin
		Prov("vin").Conn("sam", godip.Land).Conn("ill", godip.Land).Conn("dal", godip.Land).Conn("ven", godip.Land).Conn("rha", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// dac
		Prov("dac").Conn("che", godip.Coast...).Conn("bla", godip.Sea).Conn("byz", godip.Coast...).Conn("mac", godip.Land).Conn("ill", godip.Land).Conn("sam", godip.Land).Flag(godip.Coast...).
		// bla
		Prov("bla").Conn("che", godip.Sea).Conn("sip", godip.Sea).Conn("bit", godip.Sea).Conn("byz", godip.Sea).Conn("dac", godip.Sea).Flag(godip.Sea).
		// sip
		Prov("sip").Conn("che", godip.Coast...).Conn("arm", godip.Land).Conn("cap", godip.Land).Conn("gal", godip.Land).Conn("bit", godip.Coast...).Conn("bla", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// arm
		Prov("arm").Conn("che", godip.Land).Conn("dam", godip.Land).Conn("cap", godip.Land).Conn("sip", godip.Land).Flag(godip.Land).
		// mas
		Prov("mas").Conn("gau", godip.Land).Conn("rha", godip.Land).Conn("etr", godip.Coast...).Conn("lig", godip.Sea).Conn("tar", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// etr
		Prov("etr").Conn("rha", godip.Land).Conn("ven", godip.Land).Conn("rav", godip.Land).Conn("rom", godip.Coast...).Conn("lig", godip.Sea).Conn("mas", godip.Coast...).Flag(godip.Coast...).
		// ven
		Prov("ven").Conn("rha", godip.Land).Conn("vin", godip.Land).Conn("dal", godip.Coast...).Conn("adr", godip.Sea).Conn("rav", godip.Coast...).Conn("etr", godip.Land).Flag(godip.Coast...).
		// dal
		Prov("dal").Conn("ill", godip.Land).Conn("epi", godip.Coast...).Conn("adr", godip.Sea).Conn("ven", godip.Coast...).Conn("vin", godip.Land).Flag(godip.Coast...).SC(godip.Neutral).
		// ill
		Prov("ill").Conn("sam", godip.Land).Conn("dac", godip.Land).Conn("mac", godip.Land).Conn("epi", godip.Land).Conn("dal", godip.Land).Conn("vin", godip.Land).Flag(godip.Land).
		// lus
		Prov("lus").Conn("gau", godip.Land).Conn("tar", godip.Land).Conn("sag", godip.Land).Flag(godip.Land).
		// tar
		Prov("tar").Conn("gau", godip.Land).Conn("mas", godip.Coast...).Conn("lig", godip.Sea).Conn("bal", godip.Sea).Conn("sag", godip.Coast...).Conn("lus", godip.Land).Flag(godip.Coast...).
		// sag
		Prov("sag").Conn("tar", godip.Coast...).Conn("bal", godip.Sea).Conn("ber", godip.Sea).Conn("ibe", godip.Sea).Conn("mau", godip.Coast...).Conn("lus", godip.Land).Flag(godip.Coast...).SC(godip.Neutral).
		// bal
		Prov("bal").Conn("lig", godip.Sea).Conn("ber", godip.Sea).Conn("sag", godip.Sea).Conn("tar", godip.Sea).Flag(godip.Archipelago...).SC(godip.Neutral).
		// rom
		Prov("rom").Conn("rav", godip.Land).Conn("apu", godip.Land).Conn("nea", godip.Coast...).Conn("tys", godip.Sea).Conn("lig", godip.Sea).Conn("etr", godip.Coast...).Flag(godip.Coast...).SC(Rome).
		// rav
		Prov("rav").Conn("ven", godip.Coast...).Conn("adr", godip.Sea).Conn("apu", godip.Coast...).Conn("rom", godip.Land).Conn("etr", godip.Land).Flag(godip.Coast...).SC(Rome).
		// apu
		Prov("apu").Conn("adr", godip.Sea).Conn("ion", godip.Sea).Conn("nea", godip.Coast...).Conn("rom", godip.Land).Conn("rav", godip.Coast...).Flag(godip.Coast...).
		// nea
		Prov("nea").Conn("apu", godip.Coast...).Conn("ion", godip.Sea).Conn("aus", godip.Sea).Conn("sic", godip.Coast...).Conn("tys", godip.Sea).Conn("rom", godip.Coast...).Flag(godip.Coast...).SC(Rome).
		// sic
		Prov("sic").Conn("tys", godip.Sea).Conn("nea", godip.Coast...).Conn("aus", godip.Sea).Conn("pun", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// epi
		Prov("epi").Conn("ill", godip.Land).Conn("mac", godip.Land).Conn("ath", godip.Coast...).Conn("ion", godip.Sea).Conn("adr", godip.Sea).Conn("dal", godip.Coast...).Flag(godip.Coast...).
		// mac
		Prov("mac").Conn("dac", godip.Land).Conn("byz", godip.Coast...).Conn("aeg", godip.Sea).Conn("ath", godip.Coast...).Conn("epi", godip.Land).Conn("ill", godip.Land).Flag(godip.Coast...).SC(Greece).
		// ath
		Prov("ath").Conn("mac", godip.Coast...).Conn("aeg", godip.Sea).Conn("spa", godip.Coast...).Conn("ion", godip.Sea).Conn("spa", godip.Coast...).Conn("epi", godip.Coast...).Flag(godip.Coast...).SC(Greece).
		// spa
		Prov("spa").Conn("ath", godip.Coast...).Conn("aeg", godip.Sea).Conn("mes", godip.Sea).Conn("ion", godip.Sea).Flag(godip.Coast...).SC(Greece).
		// byz
		Prov("byz").Conn("bla", godip.Sea).Conn("bit", godip.Coast...).Conn("gal", godip.Land).Conn("mil", godip.Coast...).Conn("aeg", godip.Sea).Conn("mac", godip.Coast...).Conn("dac", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// bit
		Prov("bit").Conn("bla", godip.Sea).Conn("sip", godip.Coast...).Conn("gal", godip.Land).Conn("byz", godip.Coast...).Flag(godip.Coast...).
		// mil
		Prov("mil").Conn("byz", godip.Coast...).Conn("gal", godip.Land).Conn("isa", godip.Coast...).Conn("cil", godip.Sea).Conn("min", godip.Sea).Conn("aeg", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// gal
		Prov("gal").Conn("sip", godip.Land).Conn("cap", godip.Land).Conn("isa", godip.Land).Conn("mil", godip.Land).Conn("byz", godip.Land).Conn("bit", godip.Land).Flag(godip.Land).
		// isa
		Prov("isa").Conn("gal", godip.Land).Conn("cap", godip.Coast...).Conn("cil", godip.Sea).Conn("min", godip.Sea).Conn("mil", godip.Coast...).Flag(godip.Coast...).
		// cap
		Prov("cap").Conn("sip", godip.Land).Conn("arm", godip.Land).Conn("dam", godip.Land).Conn("ant", godip.Coast...).Conn("cil", godip.Sea).Conn("isa", godip.Coast...).Conn("gal", godip.Land).Flag(godip.Coast...).
		// ant
		Prov("ant").Conn("cap", godip.Coast...).Conn("dam", godip.Land).Conn("sid", godip.Coast...).Conn("cil", godip.Sea).Flag(godip.Coast...).SC(Persia).
		// dam
		Prov("dam").Conn("arm", godip.Land).Conn("ara", godip.Land).Conn("sid", godip.Land).Conn("ant", godip.Land).Conn("cap", godip.Land).Flag(godip.Land).SC(Persia).
		// sid
		Prov("sid").Conn("ant", godip.Coast...).Conn("dam", godip.Land).Conn("ara", godip.Land).Conn("tye", godip.Coast...).Conn("syr", godip.Sea).Conn("cil", godip.Sea).Flag(godip.Coast...).SC(Persia).
		// tye
		Prov("tye").Conn("sid", godip.Coast...).Conn("ara", godip.Land).Conn("jer", godip.Coast...).Conn("syr", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// ara
		Prov("ara").Conn("dam", godip.Land).Conn("nab", godip.Land).Conn("jer", godip.Land).Conn("tye", godip.Land).Conn("sid", godip.Land).Flag(godip.Land).
		// jer
		Prov("jer").Conn("tye", godip.Coast...).Conn("ara", godip.Land).Conn("nab", godip.Land).Conn("pet", godip.Land).Conn("sii", godip.Coast...).Conn("gop", godip.Sea).Conn("syr", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// nab
		Prov("nab").Conn("ara", godip.Land).Conn("ree", godip.Sea).Conn("pet", godip.Coast...).Conn("jer", godip.Land).Flag(godip.Coast...).
		// pet
		Prov("pet").Conn("nab", godip.Coast...).Conn("ree", godip.Sea).Conn("sii", godip.Coast...).Conn("jer", godip.Land).Flag(godip.Coast...).SC(godip.Neutral).
		// sii
		Prov("sii").Conn("jer", godip.Coast...).Conn("pet", godip.Coast...).Conn("ree", godip.Sea).Conn("the", godip.Coast...).Conn("ale", godip.Coast...).Conn("gop", godip.Sea).Flag(godip.Coast...).
		// the
		Prov("the").Conn("sii", godip.Coast...).Conn("ree", godip.Sea).Conn("bay", godip.Coast...).Conn("mem", godip.Coast...).Conn("ale", godip.Coast...).Conn("gop", godip.Sea).Flag(godip.Coast...).SC(Egypt).
		// ale
		Prov("ale").Conn("egy", godip.Sea).Conn("gop", godip.Sea).Conn("sii", godip.Coast...).Conn("the", godip.Coast...).Conn("mem", godip.Coast...).Conn("cyr", godip.Coast...).Conn("lib", godip.Sea).Flag(godip.Coast...).SC(Egypt).
		// mem
		Prov("mem").Conn("ale", godip.Coast...).Conn("the", godip.Coast...).Conn("bay", godip.Coast...).Conn("mar", godip.Land).Conn("cyr", godip.Land).Flag(godip.Coast...).SC(Egypt).
		// bay
		Prov("bay").Conn("mem", godip.Coast...).Conn("the", godip.Coast...).Conn("sah", godip.Land).Conn("pha", godip.Land).Conn("mar", godip.Land).Flag(godip.Coast...).
		// cyr
		Prov("cyr").Conn("lib", godip.Sea).Conn("ale", godip.Coast...).Conn("mem", godip.Land).Conn("mar", godip.Land).Conn("lep", godip.Coast...).Conn("gos", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// mar
		Prov("mar").Conn("cyr", godip.Land).Conn("mem", godip.Land).Conn("bay", godip.Land).Conn("pha", godip.Land).Conn("lep", godip.Land).Flag(godip.Land).
		// lep
		Prov("lep").Conn("gos", godip.Sea).Conn("cyr", godip.Coast...).Conn("mar", godip.Land).Conn("pha", godip.Land).Conn("num", godip.Coast...).Conn("got", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// pha
		Prov("pha").Conn("num", godip.Land).Conn("lep", godip.Land).Conn("mar", godip.Land).Conn("bay", godip.Land).Conn("sah", godip.Land).Conn("cir", godip.Land).Flag(godip.Land).
		// sah
		Prov("sah").Conn("cir", godip.Land).Conn("pha", godip.Land).Conn("bay", godip.Land).Conn("mau", godip.Land).Flag(godip.Land).
		// num
		Prov("num").Conn("got", godip.Sea).Conn("lep", godip.Coast...).Conn("pha", godip.Land).Conn("cir", godip.Land).Conn("tha", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// cir
		Prov("cir").Conn("car", godip.Land).Conn("tha", godip.Land).Conn("num", godip.Land).Conn("pha", godip.Land).Conn("sah", godip.Land).Conn("mau", godip.Land).Flag(godip.Land).SC(Carthage).
		// mau
		Prov("mau").Conn("ibe", godip.Sea).Conn("ber", godip.Sea).Conn("car", godip.Coast...).Conn("cir", godip.Land).Conn("sah", godip.Land).Conn("sag", godip.Coast...).Flag(godip.Coast...).
		// car
		Prov("car").Conn("pun", godip.Sea).Conn("tha", godip.Coast...).Conn("cir", godip.Land).Conn("mau", godip.Coast...).Conn("ber", godip.Sea).Flag(godip.Coast...).SC(Carthage).
		// tha
		Prov("tha").Conn("car", godip.Coast...).Conn("pun", godip.Sea).Conn("got", godip.Sea).Conn("num", godip.Coast...).Conn("cir", godip.Land).Flag(godip.Coast...).SC(Carthage).
		// cor
		Prov("cor").Conn("lig", godip.Sea).Conn("tys", godip.Sea).Conn("sad", godip.Coast...).Flag(godip.Coast...).
		// sad
		Prov("sad").Conn("cor", godip.Coast...).Conn("tys", godip.Sea).Conn("pun", godip.Sea).Conn("ber", godip.Sea).Conn("lig", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// cre
		Prov("cre").Conn("aeg", godip.Sea).Conn("min", godip.Sea).Conn("egy", godip.Sea).Conn("lib", godip.Sea).Conn("mes", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// cyp
		Prov("cyp").Conn("cil", godip.Sea).Conn("syr", godip.Sea).Conn("egy", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// ibe
		Prov("ibe").Conn("sag", godip.Sea).Conn("ber", godip.Sea).Conn("mau", godip.Sea).Flag(godip.Sea).
		// ber
		Prov("ber").Conn("lig", godip.Sea).Conn("sad", godip.Sea).Conn("pun", godip.Sea).Conn("car", godip.Sea).Conn("mau", godip.Sea).Conn("ibe", godip.Sea).Conn("sag", godip.Sea).Conn("bal", godip.Sea).Flag(godip.Sea).
		// lig
		Prov("lig").Conn("etr", godip.Sea).Conn("rom", godip.Sea).Conn("tys", godip.Sea).Conn("cor", godip.Sea).Conn("sad", godip.Sea).Conn("ber", godip.Sea).Conn("bal", godip.Sea).Conn("tar", godip.Sea).Conn("mas", godip.Sea).Flag(godip.Sea).
		// tys
		Prov("tys").Conn("rom", godip.Sea).Conn("nea", godip.Sea).Conn("aus", godip.Sea).Conn("sic", godip.Sea).Conn("pun", godip.Sea).Conn("sad", godip.Sea).Conn("cor", godip.Sea).Conn("lig", godip.Sea).Flag(godip.Sea).
		// pun
		Prov("pun").Conn("tys", godip.Sea).Conn("sic", godip.Sea).Conn("aus", godip.Sea).Conn("got", godip.Sea).Conn("tha", godip.Sea).Conn("car", godip.Sea).Conn("ber", godip.Sea).Conn("sad", godip.Sea).Flag(godip.Sea).
		// adr
		Prov("adr").Conn("dal", godip.Sea).Conn("epi", godip.Sea).Conn("ion", godip.Sea).Conn("apu", godip.Sea).Conn("rav", godip.Sea).Conn("ven", godip.Sea).Flag(godip.Sea).
		// ion
		Prov("ion").Conn("adr", godip.Sea).Conn("epi", godip.Sea).Conn("ath", godip.Sea).Conn("spa", godip.Sea).Conn("mes", godip.Sea).Conn("aus", godip.Sea).Conn("nea", godip.Sea).Conn("apu", godip.Sea).Flag(godip.Sea).
		// aus
		Prov("aus").Conn("ion", godip.Sea).Conn("mes", godip.Sea).Conn("lib", godip.Sea).Conn("got", godip.Sea).Conn("pun", godip.Sea).Conn("sic", godip.Sea).Conn("tys", godip.Sea).Conn("nea", godip.Sea).Flag(godip.Sea).
		// got
		Prov("got").Conn("aus", godip.Sea).Conn("mes", godip.Sea).Conn("lib", godip.Sea).Conn("gos", godip.Sea).Conn("lep", godip.Sea).Conn("num", godip.Sea).Conn("tha", godip.Sea).Conn("pun", godip.Sea).Flag(godip.Sea).
		// gos
		Prov("gos").Conn("lib", godip.Sea).Conn("cyr", godip.Sea).Conn("lep", godip.Sea).Conn("got", godip.Sea).Flag(godip.Sea).
		// lib
		Prov("lib").Conn("mes", godip.Sea).Conn("cre", godip.Sea).Conn("egy", godip.Sea).Conn("ale", godip.Sea).Conn("cyr", godip.Sea).Conn("gos", godip.Sea).Conn("got", godip.Sea).Conn("aus", godip.Sea).Flag(godip.Sea).
		// mes
		Prov("mes").Conn("spa", godip.Sea).Conn("aeg", godip.Sea).Conn("cre", godip.Sea).Conn("lib", godip.Sea).Conn("got", godip.Sea).Conn("aus", godip.Sea).Conn("ion", godip.Sea).Flag(godip.Sea).
		// aeg
		Prov("aeg").Conn("byz", godip.Sea).Conn("mil", godip.Sea).Conn("min", godip.Sea).Conn("cre", godip.Sea).Conn("mes", godip.Sea).Conn("spa", godip.Sea).Conn("ath", godip.Sea).Conn("mac", godip.Sea).Flag(godip.Sea).
		// min
		Prov("min").Conn("mil", godip.Sea).Conn("cil", godip.Sea).Conn("egy", godip.Sea).Conn("cre", godip.Sea).Conn("aeg", godip.Sea).Flag(godip.Sea).
		// egy
		Prov("egy").Conn("cil", godip.Sea).Conn("cyp", godip.Sea).Conn("syr", godip.Sea).Conn("gop", godip.Sea).Conn("ale", godip.Sea).Conn("lib", godip.Sea).Conn("cre", godip.Sea).Conn("min", godip.Sea).Flag(godip.Sea).
		// cil
		Prov("cil").Conn("cap", godip.Sea).Conn("ant", godip.Sea).Conn("sid", godip.Sea).Conn("syr", godip.Sea).Conn("cyp", godip.Sea).Conn("egy", godip.Sea).Conn("min", godip.Sea).Conn("mil", godip.Sea).Conn("isa", godip.Sea).Flag(godip.Sea).
		// syr
		Prov("syr").Conn("cil", godip.Sea).Conn("sid", godip.Sea).Conn("tye", godip.Sea).Conn("jer", godip.Sea).Conn("gop", godip.Sea).Conn("egy", godip.Sea).Conn("cyp", godip.Sea).Flag(godip.Sea).
		// gop
		Prov("gop").Conn("syr", godip.Sea).Conn("jer", godip.Sea).Conn("sii", godip.Sea).Conn("the", godip.Sea).Conn("ale", godip.Sea).Conn("egy", godip.Sea).Flag(godip.Sea).
		// ree
		Prov("ree").Conn("pet", godip.Sea).Conn("nab", godip.Sea).Conn("the", godip.Sea).Conn("sii", godip.Sea).Flag(godip.Sea).
		Done()
}
