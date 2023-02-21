package spiceislands

import (
	"github.com/zond/godip"
	"github.com/zond/godip/graph"
	"github.com/zond/godip/state"
	"github.com/zond/godip/phase"
	"github.com/zond/godip/variants/classical"
	"github.com/zond/godip/variants/common"
	"github.com/zond/godip/variants/hundred"
)

const (
	Brunei    godip.Nation = "Sultanate of Brunei"
	DaiViet   godip.Nation = "Kingdom of Dai Viet"
	Malacca   godip.Nation = "Sultanate of Malacca"
	Ternate   godip.Nation = "Sultanate of Ternate"
	Tondo     godip.Nation = "Kingdom of Tondo"
	Ayutthaya godip.Nation = "Kingdom of Ayutthaya"
	Majapahit godip.Nation = "Kingdom of Majapahit"
)

var Nations = []godip.Nation{Brunei, DaiViet, Malacca, Ternate, Tondo, Ayutthaya, Majapahit}

var newPhase = phase.Generator(hundred.BuildAnywhereParser, classical.AdjustSCs)

func Phase(year int, season godip.Season, typ godip.PhaseType) godip.Phase {
	return newPhase(year, season, typ)
}

var SpiceIslandsVariant = common.Variant{
	Name:              "Spice Islands",
	Graph:             func() godip.Graph { return SpiceIslandsGraph() },
	Start:             SpiceIslandsStart,
	Blank:             SpiceIslandsBlank,
	Phase:             classical.NewPhase,
	Parser:            hundred.BuildAnywhereParser,
	Nations:           Nations,
	PhaseTypes:        classical.PhaseTypes,
	Seasons:           classical.Seasons,
	UnitTypes:         classical.UnitTypes,
	SoloWinner:        common.SCCountWinner(18),
	SoloSCCount:       func(*state.State) int { return 18 },
	ProvinceLongNames: provinceLongNames,
	SVGMap: func() ([]byte, error) {
		return Asset("svg/spiceislandsmap.svg")
	},
	SVGVersion: "1",
	SVGUnits:    hundred.SVGUnits,
	CreatedBy:   "David E. Cohen",
	Version:     "2.0",
	Description: "THIS IS A BETA MAP. IT MIGHT BE UPDATED AND CHANGED DURING YOUR GAME, WITHOUT WARNING. IT IS ONLY ACCESSIBLE OR VISIBLE FROM THE DIPLICITY BETA VERSION. This map is a standalone version of a part of the larger East Indies Variant, but playable on a stand alone basis.",
	Rules:       "The first to 18 Supply Centers (SC) is the winner. Units may be built at any owned SC. ",
}

func SpiceIslandsBlank(phase godip.Phase) *state.State {
	return state.New(SpiceIslandsGraph(), phase, classical.BackupRule, nil, nil)
}

func SpiceIslandsStart() (result *state.State, err error) {
	startPhase := classical.NewPhase(1491, godip.Spring, godip.Movement)
	result = SpiceIslandsBlank(startPhase)
	if err = result.SetUnits(map[godip.Province]godip.Unit{
		"tun": godip.Unit{godip.Fleet, Brunei},
		"bru": godip.Unit{godip.Fleet, Brunei},
		"pla": godip.Unit{godip.Fleet, Brunei},
		"fai": godip.Unit{godip.Fleet, DaiViet},
		"han": godip.Unit{godip.Army, DaiViet},
		"hai": godip.Unit{godip.Army, DaiViet},
		"ria": godip.Unit{godip.Fleet, Malacca},
		"mal": godip.Unit{godip.Fleet, Malacca},
		"pah": godip.Unit{godip.Fleet, Malacca},
		"hal": godip.Unit{godip.Fleet, Ternate},
		"bur": godip.Unit{godip.Fleet, Ternate},
		"ser": godip.Unit{godip.Fleet, Ternate},
		"ton": godip.Unit{godip.Fleet, Tondo},
		"nam": godip.Unit{godip.Fleet, Tondo},
		"kas": godip.Unit{godip.Fleet, Tondo},
		"daw": godip.Unit{godip.Fleet, Ayutthaya},
		"ayu": godip.Unit{godip.Army, Ayutthaya},
		"roi": godip.Unit{godip.Army, Ayutthaya},
		"paj": godip.Unit{godip.Fleet, Majapahit},
		"jva": godip.Unit{godip.Fleet, Majapahit},
		"tro": godip.Unit{godip.Fleet, Majapahit},
	}); err != nil {
		return
	}
	result.SetSupplyCenters(map[godip.Province]godip.Nation{
		"tun": Brunei,
		"bru": Brunei,
		"pla": Brunei,
		"fai": DaiViet,
		"han": DaiViet,
		"hai": DaiViet,
		"ria": Malacca,
		"mal": Malacca,
		"pah": Malacca,
		"hal": Ternate,
		"bur": Ternate,
		"ser": Ternate,
		"ton": Tondo,
		"nam": Tondo,
		"kas": Tondo,
		"daw": Ayutthaya,
		"ayu": Ayutthaya,
		"roi": Ayutthaya,
		"paj": Majapahit,
		"jva": Majapahit,
		"tro": Majapahit,
	})
	return
}

func SpiceIslandsGraph() *graph.Graph {
	return graph.New().
		// Taiwan
		Prov("tai").Conn("ecs", godip.Sea).Conn("sot", godip.Sea).Conn("scs", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Namayan
		Prov("nam").Conn("bik", godip.Coast...).Conn("lse", godip.Sea).Conn("kas", godip.Coast...).Conn("ton", godip.Coast...).Conn("sib", godip.Sea).Conn("vis", godip.Sea).Flag(godip.Coast...).SC(Tondo).
		// Chaiya
		Prov("chy").Conn("gos", godip.Sea).Conn("daw", godip.Coast...).Conn("sea", godip.Sea).Conn("som", godip.Sea).Conn("mal", godip.Coast...).Conn("kel", godip.Coast...).Flag(godip.Coast...).
		// Lan Xang
		Prov("lan").Conn("han", godip.Land).Conn("sha", godip.Land).Conn("chn", godip.Land).Conn("roi", godip.Land).Conn("wia", godip.Land).Conn("fai", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Wehali
		Prov("weh").Conn("ban", godip.Sea).Conn("mas", godip.Sea).Conn("soo", godip.Sea).Conn("tim", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// East China Sea
		Prov("ecs").Conn("sot", godip.Sea).Conn("tai", godip.Sea).Conn("scs", godip.Sea).Conn("luo", godip.Sea).Conn("kas", godip.Sea).Conn("lse", godip.Sea).Conn("eao", godip.Sea).Flag(godip.Sea).
		// Khmer
		Prov("khm").Conn("wia", godip.Land).Conn("oce", godip.Coast...).Conn("gos", godip.Sea).Conn("kar", godip.Sea).Conn("chs", godip.Sea).Conn("cmp", godip.Coast...).Conn("chk", godip.Land).Conn("fai", godip.Land).Flag(godip.Coast...).SC(godip.Neutral).
		// Hanoi
		Prov("han").Conn("fai", godip.Land).Conn("chk", godip.Land).Conn("hai", godip.Land).Conn("sha", godip.Land).Conn("lan", godip.Land).Flag(godip.Land).SC(DaiViet).
		// Sulu Sea
		Prov("suu").Conn("pla", godip.Sea).Conn("chs", godip.Sea).Conn("bru", godip.Sea).Conn("tun", godip.Sea).Conn("sui", godip.Sea).Conn("zam", godip.Sea).Conn("vis", godip.Sea).Conn("sib", godip.Sea).Conn("mai", godip.Sea).Flag(godip.Sea).
		// Bikol
		Prov("bik").Conn("nam", godip.Coast...).Conn("vis", godip.Sea).Conn("lse", godip.Sea).Flag(godip.Coast...).
		// Tondo
		Prov("ton").Conn("mai", godip.Sea).Conn("sib", godip.Sea).Conn("nam", godip.Coast...).Conn("kas", godip.Land).Conn("luo", godip.Coast...).Flag(godip.Coast...).SC(Tondo).
		// Timor Sea
		Prov("tim").Conn("eao", godip.Sea).Conn("ser", godip.Sea).Conn("mol", godip.Sea).Conn("bur", godip.Sea).Conn("ban", godip.Sea).Conn("weh", godip.Sea).Conn("soo", godip.Sea).Flag(godip.Sea).
		// Negara Daha
		Prov("neg").Conn("kut", godip.Land).Conn("bru", godip.Land).Conn("suk", godip.Land).Conn("sap", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Sibuyan Sea
		Prov("sib").Conn("vis", godip.Sea).Conn("nam", godip.Sea).Conn("ton", godip.Sea).Conn("mai", godip.Sea).Conn("suu", godip.Sea).Flag(godip.Sea).
		// Sambas
		Prov("sab").Conn("kar", godip.Sea).Conn("jas", godip.Sea).Conn("suk", godip.Coast...).Conn("chs", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Mait Sea
		Prov("mai").Conn("luo", godip.Sea).Conn("scs", godip.Sea).Conn("chs", godip.Sea).Conn("pla", godip.Sea).Conn("suu", godip.Sea).Conn("sib", godip.Sea).Conn("ton", godip.Sea).Flag(godip.Sea).
		// Mindanao
		Prov("mid").Conn("eao", godip.Sea).Conn("buu", godip.Coast...).Conn("zam", godip.Coast...).Conn("sui", godip.Sea).Flag(godip.Coast...).
		// Sukadana
		Prov("suk").Conn("bru", godip.Coast...).Conn("chs", godip.Sea).Conn("sab", godip.Coast...).Conn("jas", godip.Sea).Conn("sap", godip.Coast...).Conn("neg", godip.Land).Flag(godip.Coast...).
		// Trowulan
		Prov("tro").Conn("jva", godip.Coast...).Conn("lum", godip.Coast...).Conn("mas", godip.Sea).Conn("jas", godip.Sea).Flag(godip.Coast...).SC(Majapahit).
		// Jambi
		Prov("jam").Conn("ria", godip.Coast...).Conn("wes", godip.Land).Conn("pae", godip.Coast...).Conn("kar", godip.Sea).Flag(godip.Coast...).
		// Oc Eo
		Prov("oce").Conn("khm", godip.Coast...).Conn("wia", godip.Land).Conn("roi", godip.Land).Conn("ayu", godip.Coast...).Conn("gos", godip.Sea).Flag(godip.Coast...).
		// Pahang
		Prov("pah").Conn("kel", godip.Coast...).Conn("mal", godip.Coast...).Conn("som", godip.Sea).Conn("kar", godip.Sea).Flag(godip.Coast...).SC(Malacca).
		// Shan
		Prov("sha").Conn("ava", godip.Land).Conn("peg", godip.Land).Conn("chn", godip.Land).Conn("lan", godip.Land).Conn("han", godip.Land).Flag(godip.Land).
		// Ayutthaya
		Prov("ayu").Conn("peg", godip.Land).Conn("daw", godip.Coast...).Conn("gos", godip.Sea).Conn("oce", godip.Coast...).Conn("roi", godip.Land).Flag(godip.Coast...).SC(Ayutthaya).
		// Gulf of Tomini
		Prov("got").Conn("luw", godip.Sea).Conn("ban", godip.Sea).Conn("mol", godip.Sea).Conn("hal", godip.Sea).Conn("eao", godip.Sea).Conn("sui", godip.Sea).Conn("mih", godip.Sea).Flag(godip.Sea).
		// Kelantan
		Prov("kel").Conn("kar", godip.Sea).Conn("gos", godip.Sea).Conn("chy", godip.Coast...).Conn("mal", godip.Land).Conn("pah", godip.Coast...).Flag(godip.Coast...).
		// Gulf of Dai Viet
		Prov("god").Conn("hai", godip.Sea).Conn("chk", godip.Sea).Conn("scs", godip.Sea).Flag(godip.Sea).
		// Riau
		Prov("ria").Conn("jam", godip.Coast...).Conn("kar", godip.Sea).Conn("som", godip.Sea).Conn("ped", godip.Coast...).Conn("wes", godip.Land).Flag(godip.Coast...).SC(Malacca).
		// Strait of Taiwan
		Prov("sot").Conn("scs", godip.Sea).Conn("tai", godip.Sea).Conn("ecs", godip.Sea).Flag(godip.Sea).
		// Seram
		Prov("ser").Conn("eao", godip.Sea).Conn("mol", godip.Sea).Conn("tim", godip.Sea).Flag(godip.Coast...).SC(Ternate).
		// Tunku
		Prov("tun").Conn("sui", godip.Sea).Conn("suu", godip.Sea).Conn("bru", godip.Coast...).Conn("kut", godip.Coast...).Flag(godip.Coast...).SC(Brunei).
		// Kasiguran
		Prov("kas").Conn("ecs", godip.Sea).Conn("luo", godip.Coast...).Conn("ton", godip.Land).Conn("nam", godip.Coast...).Conn("lse", godip.Sea).Flag(godip.Coast...).SC(Tondo).
		// Makassar
		Prov("mar").Conn("ban", godip.Sea).Conn("buo", godip.Coast...).Conn("luw", godip.Coast...).Conn("mih", godip.Coast...).Conn("mas", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Champa
		Prov("cmp").Conn("chs", godip.Sea).Conn("scs", godip.Sea).Conn("chk", godip.Coast...).Conn("khm", godip.Coast...).Flag(godip.Coast...).
		// Ava
		Prov("ava").Conn("sea", godip.Sea).Conn("peg", godip.Coast...).Conn("sha", godip.Land).Flag(godip.Coast...).SC(godip.Neutral).
		// Malacca
		Prov("mal").Conn("som", godip.Sea).Conn("pah", godip.Coast...).Conn("kel", godip.Land).Conn("chy", godip.Coast...).Flag(godip.Coast...).SC(Malacca).
		// Java Sea
		Prov("jas").Conn("pae", godip.Sea).Conn("mig", godip.Sea).Conn("paj", godip.Sea).Conn("jva", godip.Sea).Conn("tro", godip.Sea).Conn("mas", godip.Sea).Conn("sap", godip.Sea).Conn("suk", godip.Sea).Conn("sab", godip.Sea).Conn("kar", godip.Sea).Flag(godip.Sea).
		// Palembang
		Prov("pae").Conn("jas", godip.Sea).Conn("kar", godip.Sea).Conn("jam", godip.Coast...).Conn("wes", godip.Coast...).Conn("mig", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Lumajang
		Prov("lum").Conn("jva", godip.Coast...).Conn("soo", godip.Sea).Conn("mas", godip.Sea).Conn("tro", godip.Coast...).Flag(godip.Coast...).
		// Faifo
		Prov("fai").Conn("han", godip.Land).Conn("lan", godip.Land).Conn("wia", godip.Land).Conn("khm", godip.Land).Conn("chk", godip.Land).Flag(godip.Land).SC(DaiViet).
		// Sulawesi Sea
		Prov("sui").Conn("tun", godip.Sea).Conn("kut", godip.Sea).Conn("mas", godip.Sea).Conn("mih", godip.Sea).Conn("got", godip.Sea).Conn("eao", godip.Sea).Conn("mid", godip.Sea).Conn("zam", godip.Sea).Conn("suu", godip.Sea).Flag(godip.Sea).
		// Champa Sea
		Prov("chs").Conn("cmp", godip.Sea).Conn("khm", godip.Sea).Conn("kar", godip.Sea).Conn("sab", godip.Sea).Conn("suk", godip.Sea).Conn("bru", godip.Sea).Conn("suu", godip.Sea).Conn("pla", godip.Sea).Conn("mai", godip.Sea).Conn("scs", godip.Sea).Flag(godip.Sea).
		// Brunei
		Prov("bru").Conn("kut", godip.Land).Conn("tun", godip.Coast...).Conn("suu", godip.Sea).Conn("chs", godip.Sea).Conn("suk", godip.Coast...).Conn("neg", godip.Land).Flag(godip.Coast...).SC(Brunei).
		// Aceh
		Prov("ace").Conn("ped", godip.Coast...).Conn("som", godip.Sea).Conn("sea", godip.Sea).Conn("mig", godip.Sea).Conn("wes", godip.Coast...).Flag(godip.Coast...).
		// Zamboanga
		Prov("zam").Conn("suu", godip.Sea).Conn("sui", godip.Sea).Conn("mid", godip.Coast...).Conn("buu", godip.Coast...).Conn("vis", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Luson Sea
		Prov("lse").Conn("buu", godip.Sea).Conn("eao", godip.Sea).Conn("ecs", godip.Sea).Conn("kas", godip.Sea).Conn("nam", godip.Sea).Conn("bik", godip.Sea).Conn("vis", godip.Sea).Flag(godip.Sea).
		// Western Ocean
		Prov("wes").Conn("ace", godip.Coast...).Conn("mig", godip.Sea).Conn("pae", godip.Coast...).Conn("jam", godip.Land).Conn("ria", godip.Land).Conn("ped", godip.Land).Flag(godip.Coast...).
		// Champassak
		Prov("chk").Conn("han", godip.Land).Conn("fai", godip.Land).Conn("khm", godip.Land).Conn("cmp", godip.Coast...).Conn("scs", godip.Sea).Conn("god", godip.Sea).Conn("hai", godip.Coast...).Flag(godip.Coast...).
		// Sunda
		Prov("sun").Conn("jva", godip.Coast...).Conn("paj", godip.Coast...).Conn("mig", godip.Sea).Flag(godip.Coast...).
		// Buru
		Prov("bur").Conn("tim", godip.Sea).Conn("mol", godip.Sea).Conn("ban", godip.Sea).Flag(godip.Coast...).SC(Ternate).
		// Butuan
		Prov("buu").Conn("lse", godip.Sea).Conn("vis", godip.Sea).Conn("zam", godip.Coast...).Conn("mid", godip.Coast...).Conn("eao", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Luson
		Prov("luo").Conn("scs", godip.Sea).Conn("mai", godip.Sea).Conn("ton", godip.Coast...).Conn("kas", godip.Coast...).Conn("ecs", godip.Sea).Flag(godip.Coast...).
		// Moluccan Sea
		Prov("mol").Conn("hal", godip.Sea).Conn("got", godip.Sea).Conn("ban", godip.Sea).Conn("bur", godip.Sea).Conn("tim", godip.Sea).Conn("ser", godip.Sea).Conn("eao", godip.Sea).Flag(godip.Sea).
		// Javadvipa
		Prov("jva").Conn("lum", godip.Coast...).Conn("tro", godip.Coast...).Conn("jas", godip.Sea).Conn("paj", godip.Coast...).Conn("sun", godip.Coast...).Conn("mig", godip.Sea).Conn("soo", godip.Sea).Flag(godip.Coast...).SC(Majapahit).
		// Sea of India
		Prov("sea").Conn("mig", godip.Sea).Conn("ace", godip.Sea).Conn("som", godip.Sea).Conn("chy", godip.Sea).Conn("daw", godip.Sea).Conn("peg", godip.Sea).Conn("ava", godip.Sea).Flag(godip.Sea).
		// Pajajaran
		Prov("paj").Conn("mig", godip.Sea).Conn("sun", godip.Coast...).Conn("jva", godip.Coast...).Conn("jas", godip.Sea).Flag(godip.Coast...).SC(Majapahit).
		// Haiphong
		Prov("hai").Conn("chk", godip.Coast...).Conn("god", godip.Sea).Conn("han", godip.Land).Flag(godip.Coast...).SC(DaiViet).
		// Dawei
		Prov("daw").Conn("peg", godip.Coast...).Conn("sea", godip.Sea).Conn("chy", godip.Coast...).Conn("gos", godip.Sea).Conn("ayu", godip.Coast...).Flag(godip.Coast...).SC(Ayutthaya).
		// Banda Sea
		Prov("ban").Conn("luw", godip.Sea).Conn("buo", godip.Sea).Conn("mar", godip.Sea).Conn("mas", godip.Sea).Conn("weh", godip.Sea).Conn("tim", godip.Sea).Conn("bur", godip.Sea).Conn("mol", godip.Sea).Conn("got", godip.Sea).Flag(godip.Sea).
		// Karimata Sea
		Prov("kar").Conn("pae", godip.Sea).Conn("jas", godip.Sea).Conn("sab", godip.Sea).Conn("chs", godip.Sea).Conn("khm", godip.Sea).Conn("gos", godip.Sea).Conn("kel", godip.Sea).Conn("pah", godip.Sea).Conn("som", godip.Sea).Conn("ria", godip.Sea).Conn("jam", godip.Sea).Flag(godip.Sea).
		// Halmahera
		Prov("hal").Conn("eao", godip.Sea).Conn("got", godip.Sea).Conn("mol", godip.Sea).Flag(godip.Coast...).SC(Ternate).
		// Roi Et
		Prov("roi").Conn("peg", godip.Land).Conn("ayu", godip.Land).Conn("oce", godip.Land).Conn("wia", godip.Land).Conn("lan", godip.Land).Conn("chn", godip.Land).Flag(godip.Land).SC(Ayutthaya).
		// Minangkabau
		Prov("mig").Conn("soo", godip.Sea).Conn("jva", godip.Sea).Conn("sun", godip.Sea).Conn("paj", godip.Sea).Conn("jas", godip.Sea).Conn("pae", godip.Sea).Conn("wes", godip.Sea).Conn("ace", godip.Sea).Conn("sea", godip.Sea).Flag(godip.Sea).
		// Buton
		Prov("buo").Conn("luw", godip.Coast...).Conn("mar", godip.Coast...).Conn("ban", godip.Sea).Flag(godip.Coast...).
		// Gulf of Siam
		Prov("gos").Conn("ayu", godip.Sea).Conn("daw", godip.Sea).Conn("chy", godip.Sea).Conn("kel", godip.Sea).Conn("kar", godip.Sea).Conn("khm", godip.Sea).Conn("oce", godip.Sea).Flag(godip.Sea).
		// South China Sea
		Prov("scs").Conn("luo", godip.Sea).Conn("ecs", godip.Sea).Conn("tai", godip.Sea).Conn("sot", godip.Sea).Conn("god", godip.Sea).Conn("chk", godip.Sea).Conn("cmp", godip.Sea).Conn("chs", godip.Sea).Conn("mai", godip.Sea).Flag(godip.Sea).
		// Eastern Ocean
		Prov("eao").Conn("ecs", godip.Sea).Conn("lse", godip.Sea).Conn("buu", godip.Sea).Conn("mid", godip.Sea).Conn("sui", godip.Sea).Conn("got", godip.Sea).Conn("hal", godip.Sea).Conn("mol", godip.Sea).Conn("ser", godip.Sea).Conn("tim", godip.Sea).Flag(godip.Sea).
		// Visayas Sea
		Prov("vis").Conn("zam", godip.Sea).Conn("buu", godip.Sea).Conn("lse", godip.Sea).Conn("bik", godip.Sea).Conn("nam", godip.Sea).Conn("sib", godip.Sea).Conn("suu", godip.Sea).Flag(godip.Sea).
		// Pedir
		Prov("ped").Conn("som", godip.Sea).Conn("ace", godip.Coast...).Conn("wes", godip.Land).Conn("ria", godip.Coast...).Flag(godip.Coast...).
		// Minahassa
		Prov("mih").Conn("luw", godip.Coast...).Conn("got", godip.Sea).Conn("sui", godip.Sea).Conn("mas", godip.Sea).Conn("mar", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Luwuk
		Prov("luw").Conn("ban", godip.Sea).Conn("got", godip.Sea).Conn("mih", godip.Coast...).Conn("mar", godip.Coast...).Conn("buo", godip.Coast...).Flag(godip.Coast...).
		// Makassar Strait
		Prov("mas").Conn("mih", godip.Sea).Conn("sui", godip.Sea).Conn("kut", godip.Sea).Conn("sap", godip.Sea).Conn("jas", godip.Sea).Conn("tro", godip.Sea).Conn("lum", godip.Sea).Conn("soo", godip.Sea).Conn("weh", godip.Sea).Conn("ban", godip.Sea).Conn("mar", godip.Sea).Flag(godip.Sea).
		// Straits of Malacca
		Prov("som").Conn("sea", godip.Sea).Conn("ace", godip.Sea).Conn("ped", godip.Sea).Conn("ria", godip.Sea).Conn("kar", godip.Sea).Conn("pah", godip.Sea).Conn("mal", godip.Sea).Conn("chy", godip.Sea).Flag(godip.Sea).
		// Southern Ocean
		Prov("soo").Conn("tim", godip.Sea).Conn("weh", godip.Sea).Conn("mas", godip.Sea).Conn("lum", godip.Sea).Conn("jva", godip.Sea).Conn("mig", godip.Sea).Flag(godip.Sea).
		// Palawan
		Prov("pla").Conn("suu", godip.Sea).Conn("mai", godip.Sea).Conn("chs", godip.Sea).Flag(godip.Coast...).SC(Brunei).
		// Chiangmai
		Prov("chn").Conn("sha", godip.Land).Conn("peg", godip.Land).Conn("roi", godip.Land).Conn("lan", godip.Land).Flag(godip.Land).
		// Kutai
		Prov("kut").Conn("bru", godip.Land).Conn("neg", godip.Land).Conn("sap", godip.Coast...).Conn("mas", godip.Sea).Conn("sui", godip.Sea).Conn("tun", godip.Coast...).Flag(godip.Coast...).
		// Wiangjun
		Prov("wia").Conn("khm", godip.Land).Conn("fai", godip.Land).Conn("lan", godip.Land).Conn("roi", godip.Land).Conn("oce", godip.Land).Flag(godip.Land).
		// Sampit
		Prov("sap").Conn("jas", godip.Sea).Conn("mas", godip.Sea).Conn("kut", godip.Coast...).Conn("neg", godip.Land).Conn("suk", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Pegu
		Prov("peg").Conn("ayu", godip.Land).Conn("roi", godip.Land).Conn("chn", godip.Land).Conn("sha", godip.Land).Conn("ava", godip.Coast...).Conn("sea", godip.Sea).Conn("daw", godip.Coast...).Flag(godip.Coast...).
		Done()
}

var provinceLongNames = map[godip.Province]string{
	"tai": "Taiwan",
	"nam": "Namayan",
	"chy": "Chaiya",
	"lan": "Lan Xang",
	"weh": "Wehali",
	"ecs": "East China Sea",
	"khm": "Khmer",
	"han": "Hanoi",
	"suu": "Sulu Sea",
	"bik": "Bikol",
	"ton": "Tondo",
	"tim": "Timor Sea",
	"neg": "Negara Daha",
	"sib": "Sibuyan Sea",
	"sab": "Sambas",
	"mai": "Mait Sea",
	"mid": "Mindanao",
	"suk": "Sukadana",
	"tro": "Trowulan",
	"jam": "Jambi",
	"oce": "Oc Eo",
	"pah": "Pahang",
	"sha": "Shan",
	"ayu": "Ayutthaya",
	"got": "Gulf of Tomini",
	"kel": "Kelantan",
	"god": "Gulf of Dai Viet",
	"ria": "Riau",
	"sot": "Strait of Taiwan",
	"ser": "Seram",
	"tun": "Tunku",
	"kas": "Kasiguran",
	"mar": "Makassar",
	"cmp": "Champa",
	"ava": "Ava",
	"mal": "Malacca",
	"jas": "Java Sea",
	"pae": "Palembang",
	"lum": "Lumajang",
	"fai": "Faifo",
	"sui": "Sulawesi Sea",
	"chs": "Champa Sea",
	"bru": "Brunei",
	"ace": "Aceh",
	"zam": "Zamboanga",
	"lse": "Luson Sea",
	"wes": "Western Ocean",
	"chk": "Champassak",
	"sun": "Sunda",
	"bur": "Buru",
	"buu": "Butuan",
	"luo": "Luson",
	"mol": "Moluccan Sea",
	"jva": "Javadvipa",
	"sea": "Sea of India",
	"paj": "Pajajaran",
	"hai": "Haiphong",
	"daw": "Dawei",
	"ban": "Banda Sea",
	"kar": "Karimata Sea",
	"hal": "Halmahera",
	"roi": "Roi Et",
	"mig": "Minangkabau",
	"buo": "Buton",
	"gos": "Gulf of Siam",
	"scs": "South China Sea",
	"eao": "Eastern Ocean",
	"vis": "Visayas Sea",
	"ped": "Pedir",
	"mih": "Minahassa",
	"luw": "Luwuk",
	"mas": "Makassar Strait",
	"som": "Straits of Malacca",
	"soo": "Southern Ocean",
	"pla": "Palawan",
	"chn": "Chiangmai",
	"kut": "Kutai",
	"wia": "Wiangjun",
	"sap": "Sampit",
	"peg": "Pegu",
}
