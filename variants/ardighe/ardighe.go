package ardighe

import (
	"github.com/zond/godip"
	"github.com/zond/godip/graph"
	"github.com/zond/godip/state"
	"github.com/zond/godip/variants/classical"
	"github.com/zond/godip/variants/common"
)

const (
	Connacht godip.Nation = "Connacht"
	Ulaidh   godip.Nation = "Ulaidh"
	Midhe    godip.Nation = "Midhe"
	Laighin  godip.Nation = "Laighin"
	Mumhan   godip.Nation = "Mumhan"
)

var Nations = []godip.Nation{Connacht, Ulaidh, Midhe, Laighin, Mumhan}

var ardigheVariant = common.Variant{
	Name:       "Ard Rí",
	Graph:      func() godip.Graph { return ardigheGraph() },
	Start:      ardigheStart,
	Blank:      ardigheBlank,
	Phase:      classical.NewPhase,
	Parser:     classical.Parser,
	Nations:    Nations,
	PhaseTypes: classical.PhaseTypes,
	Seasons:    classical.Seasons,
	UnitTypes:  classical.UnitTypes,
	SoloWinner: common.SCCountWinner(15),
	SVGMap: func() ([]byte, error) {
		return Asset("svg/ardighemap.svg")
	},
	SVGVersion: "1",
	SVGUnits: map[godip.UnitType]func() ([]byte, error){
		godip.Army: func() ([]byte, error) {
			return classical.Asset("svg/army.svg")
		},
		godip.Fleet: func() ([]byte, error) {
			return classical.Asset("svg/fleet.svg")
		},
	},
	CreatedBy:   "Stuart John Bernard",
	Version:     "",
	Description: "",
	Rules: "",
}

func ardigheBlank(phase godip.Phase) *state.State {
	return state.New(ardigheGraph(), phase, classical.BackupRule, nil)
}

func ardigheStart() (result *state.State, err error) {
	startPhase := classical.NewPhase(379, godip.Fall, godip.Adjustment)
	result = ardigheBlank(startPhase)
	}); err != nil {
		return
	}
	result.SetSupplyCenters(map[godip.Province]godip.Nation{
		"cru": Connacht,
		"tua": Connacht,
		"mon": Connacht,
		"aic": Ulaidh,
		"mag": Ulaidh,
		"ema": Ulaidh,
		"tem": Midhe,
		"uis": Midhe,
		"gua": Midhe,
		"naa": Laighin,
		"alm": Laighin,
		"aie": Laighin,
		"anu": Mumhan,
		"eog": Mumhan,
		"edg": Mumhan,
	})
	return
}

func ardigheGraph() *graph.Graph {
	return graph.New().
		// Taiden
		Prov("tai").Conn("cru", godip.Coast...).Conn("gam", godip.Coast...).Conn("nat", godip.Sea).Flag(godip.Coast...).
		// Ailech
		Prov("aic").Conn("nia", godip.Land).Conn("ema", godip.Land).Conn("air", godip.Land).Conn("fri", godip.Coast...).Conn("nat", godip.Sea).Conn("cnl", godip.Coast...).Flag(godip.Coast...).SC(Ulaidh).
		// Cruachu
		Prov("cru").Conn("gam", godip.Coast...).Conn("tai", godip.Coast...).Conn("nat", godip.Sea).Conn("cor", godip.Coast...).Conn("eir", godip.Land).Conn("tua", godip.Land).Conn("con", godip.Land).Flag(godip.Coast...).SC(Connacht).
		// Gamain Raige
		Prov("gam").Conn("cru", godip.Coast...).Conn("con", godip.Coast...).Conn("mnb", godip.Sea).Conn("nat", godip.Sea).Conn("tai", godip.Coast...).Flag(godip.Coast...).
		// Erainn
		Prov("era").Conn("alm", godip.Coast...).Conn("edg", godip.Land).Conn("cia", godip.Land).Conn("anu", godip.Coast...).Conn("sou", godip.Sea).Conn("eri", godip.Sea).Flag(godip.Coast...).
		// Mona Gha
		Prov("mon").Conn("nia", godip.Coast...).Conn("mnb", godip.Sea).Conn("con", godip.Coast...).Conn("ere", godip.Land).Conn("ema", godip.Land).Flag(godip.Coast...).SC(Connacht).
		// Midhe
		Prov("mid").Conn("esc", godip.Land).Conn("lai", godip.Land).Conn("tem", godip.Land).Conn("gua", godip.Land).Conn("asa", godip.Land).Conn("uis", godip.Land).Flag(godip.Land).
		// North Atlantic
		Prov("nat").Conn("sou", godip.Sea).Conn("mum", godip.Sea).Conn("sha", godip.Sea).Conn("cor", godip.Sea).Conn("cru", godip.Sea).Conn("tai", godip.Sea).Conn("gam", godip.Sea).Conn("mnb", godip.Sea).Conn("cnl", godip.Sea).Conn("aic", godip.Sea).Conn("fri", godip.Sea).Conn("nch", godip.Sea).Flag(godip.Sea).
		// North Channel
		Prov("nch").Conn("nat", godip.Sea).Conn("fri", godip.Sea).Conn("eri", godip.Sea).Flag(godip.Sea).
		// Muma
		Prov("mum").Conn("anu", godip.Coast...).Conn("cia", godip.Land).Conn("eog", godip.Coast...).Conn("sha", godip.Sea).Conn("nat", godip.Sea).Conn("sou", godip.Sea).Flag(godip.Coast...).
		// Uis- neach
		Prov("uis").Conn("asa", godip.Land).Conn("ere", godip.Land).Conn("tua", godip.Land).Conn("esc", godip.Land).Conn("mid", godip.Land).Flag(godip.Land).SC(Midhe).
		// Mona Bay
		Prov("mnb").Conn("nia", godip.Sea).Conn("cnl", godip.Sea).Conn("nat", godip.Sea).Conn("gam", godip.Sea).Conn("con", godip.Sea).Conn("mon", godip.Sea).Flag(godip.Sea).
		// Erin Sea
		Prov("eri").Conn("nch", godip.Sea).Conn("fri", godip.Sea).Conn("mag", godip.Sea).Conn("tmb", godip.Sea).Conn("nsb", godip.Sea).Conn("rat", godip.Sea).Conn("alm", godip.Sea).Conn("era", godip.Sea).Conn("sou", godip.Sea).Flag(godip.Sea).
		// Shannon
		Prov("sha").Conn("cor", godip.Sea).Conn("nat", godip.Sea).Conn("mum", godip.Sea).Conn("eog", godip.Sea).Flag(godip.Sea).
		// Frida
		Prov("fri").Conn("nat", godip.Sea).Conn("aic", godip.Coast...).Conn("mag", godip.Coast...).Conn("eri", godip.Sea).Conn("nch", godip.Sea).Flag(godip.Coast...).
		// Almu
		Prov("alm").Conn("era", godip.Coast...).Conn("eri", godip.Sea).Conn("rat", godip.Coast...).Conn("lai", godip.Land).Conn("aie", godip.Land).Conn("edg", godip.Land).Flag(godip.Coast...).SC(Laighin).
		// Ciann
		Prov("cia").Conn("mog", godip.Land).Conn("eir", godip.Land).Conn("eog", godip.Land).Conn("mum", godip.Land).Conn("anu", godip.Land).Conn("era", godip.Land).Conn("edg", godip.Land).Flag(godip.Land).
		// Eremon
		Prov("ere").Conn("con", godip.Land).Conn("tua", godip.Land).Conn("uis", godip.Land).Conn("asa", godip.Land).Conn("ema", godip.Land).Conn("mon", godip.Land).Flag(godip.Land).
		// Guara
		Prov("gua").Conn("ema", godip.Land).Conn("asa", godip.Land).Conn("mid", godip.Land).Conn("tem", godip.Coast...).Conn("tmb", godip.Sea).Conn("air", godip.Coast...).Flag(godip.Coast...).SC(Midhe).
		// Escir Riada
		Prov("esc").Conn("mog", godip.Land).Conn("lai", godip.Land).Conn("mid", godip.Land).Conn("uis", godip.Land).Conn("tua", godip.Land).Conn("eir", godip.Land).Flag(godip.Land).
		// Niall
		Prov("nia").Conn("cnl", godip.Coast...).Conn("mnb", godip.Sea).Conn("mon", godip.Coast...).Conn("ema", godip.Land).Conn("aic", godip.Land).Flag(godip.Coast...).
		// Rath- drum
		Prov("rat").Conn("alm", godip.Coast...).Conn("eri", godip.Sea).Conn("nsb", godip.Sea).Conn("naa", godip.Coast...).Conn("lai", godip.Land).Flag(godip.Coast...).
		// Mag Ruth
		Prov("mag").Conn("eri", godip.Sea).Conn("fri", godip.Coast...).Conn("air", godip.Coast...).Conn("tmb", godip.Sea).Flag(godip.Coast...).SC(Ulaidh).
		// Naas
		Prov("naa").Conn("tem", godip.Coast...).Conn("lai", godip.Land).Conn("rat", godip.Coast...).Conn("nsb", godip.Sea).Flag(godip.Coast...).SC(Laighin).
		// Conn
		Prov("con").Conn("ere", godip.Land).Conn("mon", godip.Coast...).Conn("mnb", godip.Sea).Conn("gam", godip.Coast...).Conn("cru", godip.Land).Conn("tua", godip.Land).Flag(godip.Coast...).
		// Naas Bay
		Prov("nsb").Conn("tem", godip.Sea).Conn("naa", godip.Sea).Conn("rat", godip.Sea).Conn("eri", godip.Sea).Conn("tmb", godip.Sea).Flag(godip.Sea).
		// Anu
		Prov("anu").Conn("mum", godip.Coast...).Conn("sou", godip.Sea).Conn("era", godip.Coast...).Conn("cia", godip.Land).Flag(godip.Coast...).SC(Mumhan).
		// Emain Macha
		Prov("ema").Conn("gua", godip.Land).Conn("air", godip.Land).Conn("aic", godip.Land).Conn("nia", godip.Land).Conn("mon", godip.Land).Conn("ere", godip.Land).Conn("asa", godip.Land).Flag(godip.Land).SC(Ulaidh).
		// Conall
		Prov("cnl").Conn("nia", godip.Coast...).Conn("aic", godip.Coast...).Conn("nat", godip.Sea).Conn("mnb", godip.Sea).Flag(godip.Coast...).
		// Cora Baiscin
		Prov("cor").Conn("eog", godip.Coast...).Conn("eir", godip.Land).Conn("cru", godip.Coast...).Conn("nat", godip.Sea).Conn("sha", godip.Sea).Flag(godip.Coast...).
		// Temuir
		Prov("tem").Conn("nsb", godip.Sea).Conn("tmb", godip.Sea).Conn("gua", godip.Coast...).Conn("mid", godip.Land).Conn("lai", godip.Land).Conn("naa", godip.Coast...).Flag(godip.Coast...).SC(Midhe).
		// Temuir Bay
		Prov("tmb").Conn("eri", godip.Sea).Conn("mag", godip.Sea).Conn("air", godip.Sea).Conn("gua", godip.Sea).Conn("tem", godip.Sea).Conn("nsb", godip.Sea).Flag(godip.Sea).
		// Ail- end
		Prov("aie").Conn("edg", godip.Land).Conn("alm", godip.Land).Conn("lai", godip.Land).Conn("mog", godip.Land).Flag(godip.Land).SC(Laighin).
		// Airgyallia
		Prov("air").Conn("aic", godip.Land).Conn("ema", godip.Land).Conn("gua", godip.Coast...).Conn("tmb", godip.Sea).Conn("mag", godip.Coast...).Flag(godip.Coast...).
		// Laigin
		Prov("lai").Conn("alm", godip.Land).Conn("rat", godip.Land).Conn("naa", godip.Land).Conn("tem", godip.Land).Conn("mid", godip.Land).Conn("esc", godip.Land).Conn("mog", godip.Land).Conn("aie", godip.Land).Flag(godip.Land).
		// Eir Craide
		Prov("eir").Conn("tua", godip.Land).Conn("cru", godip.Land).Conn("cor", godip.Land).Conn("eog", godip.Land).Conn("cia", godip.Land).Conn("mog", godip.Land).Conn("esc", godip.Land).Flag(godip.Land).
		// Mogh
		Prov("mog").Conn("esc", godip.Land).Conn("eir", godip.Land).Conn("cia", godip.Land).Conn("edg", godip.Land).Conn("aie", godip.Land).Conn("lai", godip.Land).Flag(godip.Land).
		// South Atlantic
		Prov("sou").Conn("eri", godip.Sea).Conn("era", godip.Sea).Conn("anu", godip.Sea).Conn("mum", godip.Sea).Conn("nat", godip.Sea).Flag(godip.Sea).
		// Asail
		Prov("asa").Conn("gua", godip.Land).Conn("ema", godip.Land).Conn("ere", godip.Land).Conn("uis", godip.Land).Conn("mid", godip.Land).Flag(godip.Land).
		// Tuathal
		Prov("tua").Conn("eir", godip.Land).Conn("esc", godip.Land).Conn("uis", godip.Land).Conn("ere", godip.Land).Conn("con", godip.Land).Conn("cru", godip.Land).Flag(godip.Land).SC(Connacht).
		// Eoghan Mor
		Prov("eog").Conn("cor", godip.Coast...).Conn("sha", godip.Sea).Conn("mum", godip.Coast...).Conn("cia", godip.Land).Conn("eir", godip.Land).Flag(godip.Coast...).SC(Mumhan).
		// Edghanacata
		Prov("edg").Conn("aie", godip.Land).Conn("mog", godip.Land).Conn("cia", godip.Land).Conn("era", godip.Land).Conn("alm", godip.Land).Flag(godip.Land).SC(Mumhan).
		Done()
}
