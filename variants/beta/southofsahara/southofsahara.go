package southofsahara

import (
	"github.com/zond/godip"
	"github.com/zond/godip/graph"
	"github.com/zond/godip/phase"
	"github.com/zond/godip/state"
	"github.com/zond/godip/variants/classical"
	"github.com/zond/godip/variants/common"
	"github.com/zond/godip/variants/hundred"
)

const (
	Bonoman godip.Nation = "Bonoman"
	Benin   godip.Nation = "Benin"
	Jolof   godip.Nation = "Jolof"
	Mali    godip.Nation = "Mali"
	Bornu   godip.Nation = "Bornu"
)

var Nations = []godip.Nation{Bonoman, Benin, Jolof, Mali, Bornu}

var newPhase = phase.Generator(hundred.BuildAnywhereParser, classical.AdjustSCs)

func Phase(year int, season godip.Season, typ godip.PhaseType) godip.Phase {
	return newPhase(year, season, typ)
}

var SouthofSaharaVariant = common.Variant{
	Name:              "SouthofSahara",
	Graph:             func() godip.Graph { return SouthofSaharaGraph() },
	Start:             SouthofSaharaStart,
	Blank:             SouthofSaharaBlank,
	Phase:             Phase,
	Parser:            hundred.BuildAnywhereParser,
	Nations:           Nations,
	PhaseTypes:        classical.PhaseTypes,
	Seasons:           classical.Seasons,
	UnitTypes:         classical.UnitTypes,
	SoloWinner:        common.SCCountWinner(13),
	SoloSCCount:       func(*state.State) int { return 13 },
	ProvinceLongNames: provinceLongNames,
	SVGMap: func() ([]byte, error) {
		return Asset("svg/southofsaharamap.svg")
	},
	SVGVersion: "1",
	SVGUnits: map[godip.UnitType]func() ([]byte, error){
		godip.Army: func() ([]byte, error) {
			return Asset("svg/army.svg")
		},
		godip.Fleet: func() ([]byte, error) {
			return Asset("svg/fleet.svg")
		},
	},
	CreatedBy:   "",
	Version:     "",
	Description: "",
	Rules:       "",
}

func SouthofSaharaBlank(phase godip.Phase) *state.State {
	return state.New(SouthofSaharaGraph(), phase, classical.BackupRule, nil, nil)
}

func SouthofSaharaStart() (result *state.State, err error) {
	startPhase := Phase(1401, godip.Spring, godip.Movement)
	result = SouthofSaharaBlank(startPhase)
	if err = result.SetUnits(map[godip.Province]godip.Unit{
		"bon": godip.Unit{godip.Army, Bonoman},
		"beg": godip.Unit{godip.Army, Bonoman},
		"saa": godip.Unit{godip.Army, Bonoman},
		"edo": godip.Unit{godip.Army, Benin},
		"ife": godip.Unit{godip.Army, Benin},
		"owo": godip.Unit{godip.Army, Benin},
		"kay": godip.Unit{godip.Army, Jolof},
		"bao": godip.Unit{godip.Army, Jolof},
		"sao": godip.Unit{godip.Army, Jolof},
		"kum": godip.Unit{godip.Army, Mali},
		"wal": godip.Unit{godip.Army, Mali},
		"tim": godip.Unit{godip.Army, Mali},
		"jen": godip.Unit{godip.Army, Mali},
		"nji": godip.Unit{godip.Army, Bornu},
		"mas": godip.Unit{godip.Army, Bornu},
		"bil": godip.Unit{godip.Army, Bornu},
	}); err != nil {
		return
	}
	result.SetSupplyCenters(map[godip.Province]godip.Nation{
		"bon": Bonoman,
		"beg": Bonoman,
		"saa": Bonoman,
		"edo": Benin,
		"ife": Benin,
		"owo": Benin,
		"kay": Jolof,
		"bao": Jolof,
		"sao": Jolof,
		"kum": Mali,
		"wal": Mali,
		"tim": Mali,
		"jen": Mali,
		"nji": Bornu,
		"mas": Bornu,
		"bil": Bornu,
	})
	return
}

func SouthofSaharaGraph() *graph.Graph {
	return graph.New().
		// Mossi
		Prov("mos").Conn("gwa", godip.Land).Conn("gwa", godip.Land).Conn("aog", godip.Land).Conn("aog", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Wouri
		Prov("wou").Conn("lun", godip.Coast...).Conn("kon", godip.Land).Conn("bou", godip.Land).Conn("zaz", godip.Land).Conn("owo", godip.Coast...).Conn("big", godip.Sea).Flag(godip.Coast...).
		// Bure
		Prov("bur").Conn("beg", godip.Coast...).Conn("aog", godip.Land).Conn("jen", godip.Land).Conn("nia", godip.Land).Conn("bam", godip.Land).Conn("wan", godip.Coast...).Conn("sob", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Sine
		Prov("sin").Conn("wan", godip.Coast...).Conn("sao", godip.Land).Conn("bao", godip.Land).Conn("kay", godip.Coast...).Conn("sow", godip.Sea).Flag(godip.Coast...).
		// Bight of Biafra
		Prov("big").Conn("lun", godip.Sea).Conn("wou", godip.Sea).Conn("owo", godip.Sea).Conn("edo", godip.Sea).Conn("ije", godip.Sea).Conn("aka", godip.Sea).Conn("etu", godip.Sea).Conn("etu", godip.Sea).Conn("hs1", godip.Sea).Conn("hs2", godip.Sea).Conn("hs3", godip.Sea).Flag(godip.Sea).
		// Jolof Sea
		Prov("jol").Conn("kay", godip.Sea).Conn("waa", godip.Sea).Conn("sow", godip.Sea).Conn("so1", godip.Sea).Conn("so2", godip.Sea).Conn("so3", godip.Sea).Conn("so4", godip.Sea).Conn("so5", godip.Sea).Conn("hs1", godip.Sea).Conn("hs2", godip.Sea).Conn("hs3", godip.Sea).Flag(godip.Sea).
		// Bambuk
		Prov("bam").Conn("bur", godip.Land).Conn("nia", godip.Land).Conn("kum", godip.Land).Conn("awd", godip.Land).Conn("sao", godip.Land).Conn("wan", godip.Land).Flag(godip.Land).
		// Tibesti
		Prov("tib").Conn("mur", godip.Land).Conn("bil", godip.Land).Conn("abe", godip.Land).Conn("so1", godip.Land).Conn("so2", godip.Land).Conn("so3", godip.Land).Conn("so4", godip.Land).Conn("so5", godip.Land).Flag(godip.Land).
		// Ijebu
		Prov("ije").Conn("edo", godip.Coast...).Conn("ife", godip.Land).Conn("oyo", godip.Land).Conn("saa", godip.Land).Conn("bon", godip.Coast...).Conn("aka", godip.Sea).Conn("big", godip.Sea).Flag(godip.Coast...).
		// Waalo
		Prov("waa").Conn("jol", godip.Sea).Conn("kay", godip.Coast...).Conn("tic", godip.Land).Conn("so1", godip.Land).Conn("so2", godip.Land).Conn("so3", godip.Land).Conn("so4", godip.Land).Conn("so5", godip.Land).Flag(godip.Coast...).
		// Ubangi
		Prov("uba").Conn("abe", godip.Land).Conn("bou", godip.Land).Conn("kon", godip.Land).Conn("lun", godip.Land).Flag(godip.Land).
		// Gwandu
		Prov("gwa").Conn("aog", godip.Land).Conn("mos", godip.Land).Conn("mos", godip.Land).Conn("aog", godip.Land).Conn("saa", godip.Land).Conn("oyo", godip.Land).Conn("tak", godip.Land).Conn("gao", godip.Land).Conn("jen", godip.Land).Flag(godip.Land).
		// Ghat
		Prov("gha").Conn("mur", godip.Land).Conn("tav", godip.Land).Conn("gao", godip.Land).Conn("tad", godip.Land).Conn("tak", godip.Land).Conn("aga", godip.Land).Conn("bil", godip.Land).Conn("so1", godip.Land).Conn("so2", godip.Land).Conn("so3", godip.Land).Conn("so4", godip.Land).Conn("so5", godip.Land).Flag(godip.Land).
		// Masseniya
		Prov("mas").Conn("zaz", godip.Land).Conn("bou", godip.Land).Conn("abe", godip.Land).Conn("bil", godip.Land).Conn("nji", godip.Land).Flag(godip.Land).SC(Bornu).
		// Sahara Oasis 1
		Prov("so1").Conn("so2", godip.Land).Conn("so3", godip.Land).Conn("so4", godip.Land).Conn("so5", godip.Land).Conn("tib", godip.Land).Conn("mur", godip.Land).Conn("gha", godip.Land).Conn("tav", godip.Land).Conn("tic", godip.Land).Conn("waa", godip.Land).Conn("jol", godip.Sea).Conn("hs1", godip.Sea).Conn("hs2", godip.Sea).Conn("hs3", godip.Sea).Flag(godip.Land).
		// Etula Eri
		Prov("etu").Conn("big", godip.Sea).Conn("big", godip.Sea).Conn("hs1", godip.Sea).Conn("hs2", godip.Sea).Conn("hs3", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Sahara Oasis 4
		Prov("so4").Conn("so2", godip.Land).Conn("so3", godip.Land).Conn("so1", godip.Land).Conn("so5", godip.Land).Conn("tib", godip.Land).Conn("mur", godip.Land).Conn("gha", godip.Land).Conn("tav", godip.Land).Conn("tic", godip.Land).Conn("waa", godip.Land).Conn("jol", godip.Sea).Conn("hs1", godip.Sea).Conn("hs2", godip.Sea).Conn("hs3", godip.Sea).Flag(godip.Land).
		// Sahara Oasis 5
		Prov("so5").Conn("so2", godip.Land).Conn("so3", godip.Land).Conn("so4", godip.Land).Conn("so1", godip.Land).Conn("tib", godip.Land).Conn("mur", godip.Land).Conn("gha", godip.Land).Conn("tav", godip.Land).Conn("tic", godip.Land).Conn("waa", godip.Land).Conn("jol", godip.Sea).Conn("hs1", godip.Sea).Conn("hs2", godip.Sea).Conn("hs3", godip.Sea).Flag(godip.Land).
		// Murzuk
		Prov("mur").Conn("gha", godip.Land).Conn("bil", godip.Land).Conn("tib", godip.Land).Conn("so1", godip.Land).Conn("so2", godip.Land).Conn("so3", godip.Land).Conn("so4", godip.Land).Conn("so5", godip.Land).Flag(godip.Land).
		// Takedda
		Prov("tak").Conn("gao", godip.Land).Conn("gwa", godip.Land).Conn("oyo", godip.Land).Conn("aga", godip.Land).Conn("gha", godip.Land).Conn("tad", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Gao
		Prov("gao").Conn("tad", godip.Land).Conn("gha", godip.Land).Conn("tav", godip.Land).Conn("tim", godip.Land).Conn("jen", godip.Land).Conn("gwa", godip.Land).Conn("tak", godip.Land).Conn("tad", godip.Land).Flag(godip.Land).
		// Oyo
		Prov("oyo").Conn("ife", godip.Land).Conn("nuf", godip.Land).Conn("aga", godip.Land).Conn("tak", godip.Land).Conn("gwa", godip.Land).Conn("saa", godip.Land).Conn("ije", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Tadmekka
		Prov("tad").Conn("gao", godip.Land).Conn("gao", godip.Land).Conn("tak", godip.Land).Conn("gha", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Ife
		Prov("ife").Conn("oyo", godip.Land).Conn("ije", godip.Land).Conn("edo", godip.Land).Conn("owo", godip.Land).Conn("zaz", godip.Land).Conn("nuf", godip.Land).Flag(godip.Land).SC(Benin).
		// Owo
		Prov("owo").Conn("big", godip.Sea).Conn("wou", godip.Coast...).Conn("zaz", godip.Land).Conn("ife", godip.Land).Conn("edo", godip.Coast...).Flag(godip.Coast...).SC(Benin).
		// Agadez
		Prov("aga").Conn("nuf", godip.Land).Conn("zaz", godip.Land).Conn("nji", godip.Land).Conn("bil", godip.Land).Conn("gha", godip.Land).Conn("tak", godip.Land).Conn("oyo", godip.Land).Flag(godip.Land).
		// Sea of Wangara
		Prov("sow").Conn("jol", godip.Sea).Conn("sob", godip.Sea).Conn("wan", godip.Sea).Conn("sin", godip.Sea).Conn("kay", godip.Sea).Conn("hs1", godip.Sea).Conn("hs2", godip.Sea).Conn("hs3", godip.Sea).Flag(godip.Sea).
		// Jenne
		Prov("jen").Conn("aog", godip.Land).Conn("gwa", godip.Land).Conn("gao", godip.Land).Conn("tim", godip.Land).Conn("wal", godip.Land).Conn("nia", godip.Land).Conn("bur", godip.Land).Flag(godip.Land).SC(Mali).
		// Tavdeni
		Prov("tav").Conn("wal", godip.Land).Conn("tim", godip.Land).Conn("gao", godip.Land).Conn("gha", godip.Land).Conn("tic", godip.Land).Conn("kum", godip.Land).Conn("so1", godip.Land).Conn("so2", godip.Land).Conn("so3", godip.Land).Conn("so4", godip.Land).Conn("so5", godip.Land).Flag(godip.Land).
		// Wangara
		Prov("wan").Conn("sao", godip.Land).Conn("sin", godip.Coast...).Conn("sow", godip.Sea).Conn("sob", godip.Sea).Conn("bur", godip.Coast...).Conn("bam", godip.Land).Flag(godip.Coast...).SC(godip.Neutral).
		// Aogadougou
		Prov("aog").Conn("jen", godip.Land).Conn("bur", godip.Land).Conn("beg", godip.Land).Conn("saa", godip.Land).Conn("gwa", godip.Land).Conn("mos", godip.Land).Conn("mos", godip.Land).Conn("gwa", godip.Land).Flag(godip.Land).
		// Kayor
		Prov("kay").Conn("jol", godip.Sea).Conn("sow", godip.Sea).Conn("sin", godip.Coast...).Conn("bao", godip.Land).Conn("sao", godip.Land).Conn("awd", godip.Land).Conn("tic", godip.Land).Conn("waa", godip.Coast...).Flag(godip.Coast...).SC(Jolof).
		// Bouar
		Prov("bou").Conn("kon", godip.Land).Conn("uba", godip.Land).Conn("abe", godip.Land).Conn("mas", godip.Land).Conn("zaz", godip.Land).Conn("wou", godip.Land).Flag(godip.Land).
		// Awdagost
		Prov("awd").Conn("tic", godip.Land).Conn("kay", godip.Land).Conn("sao", godip.Land).Conn("bam", godip.Land).Conn("kum", godip.Land).Flag(godip.Land).
		// Walata
		Prov("wal").Conn("tav", godip.Land).Conn("kum", godip.Land).Conn("nia", godip.Land).Conn("jen", godip.Land).Conn("tim", godip.Land).Flag(godip.Land).SC(Mali).
		// Lunda
		Prov("lun").Conn("uba", godip.Land).Conn("kon", godip.Land).Conn("wou", godip.Coast...).Conn("big", godip.Sea).Flag(godip.Coast...).
		// Begho
		Prov("beg").Conn("bur", godip.Coast...).Conn("sob", godip.Sea).Conn("aka", godip.Sea).Conn("bon", godip.Coast...).Conn("saa", godip.Land).Conn("aog", godip.Land).Flag(godip.Coast...).SC(Bonoman).
		// Sahara Oasis 2
		Prov("so2").Conn("so5", godip.Land).Conn("so3", godip.Land).Conn("so4", godip.Land).Conn("so1", godip.Land).Conn("tib", godip.Land).Conn("mur", godip.Land).Conn("gha", godip.Land).Conn("tav", godip.Land).Conn("tic", godip.Land).Conn("waa", godip.Land).Conn("jol", godip.Sea).Conn("hs1", godip.Sea).Conn("hs2", godip.Sea).Conn("hs3", godip.Sea).Flag(godip.Land).
		// High Sea 1
		Prov("hs1").Conn("so1", godip.Sea).Conn("so2", godip.Sea).Conn("so3", godip.Sea).Conn("so4", godip.Sea).Conn("so5", godip.Sea).Conn("jol", godip.Sea).Conn("sow", godip.Sea).Conn("sob", godip.Sea).Conn("aka", godip.Sea).Conn("big", godip.Sea).Conn("etu", godip.Sea).Flag(godip.Sea).
		// Timbuktu
		Prov("tim").Conn("wal", godip.Land).Conn("jen", godip.Land).Conn("gao", godip.Land).Conn("tav", godip.Land).Flag(godip.Land).SC(Mali).
		// Sea of Bure
		Prov("sob").Conn("bur", godip.Sea).Conn("wan", godip.Sea).Conn("sow", godip.Sea).Conn("aka", godip.Sea).Conn("beg", godip.Sea).Conn("hs1", godip.Sea).Conn("hs2", godip.Sea).Conn("hs3", godip.Sea).Flag(godip.Sea).
		// Kumbi Saleh
		Prov("kum").Conn("tav", godip.Land).Conn("tic", godip.Land).Conn("awd", godip.Land).Conn("bam", godip.Land).Conn("nia", godip.Land).Conn("wal", godip.Land).Flag(godip.Land).SC(Mali).
		// Sahara Oasis 3
		Prov("so3").Conn("so2", godip.Land).Conn("so5", godip.Land).Conn("so4", godip.Land).Conn("so1", godip.Land).Conn("tib", godip.Land).Conn("mur", godip.Land).Conn("gha", godip.Land).Conn("tav", godip.Land).Conn("tic", godip.Land).Conn("waa", godip.Land).Conn("jol", godip.Sea).Conn("hs1", godip.Sea).Conn("hs2", godip.Sea).Conn("hs3", godip.Sea).Flag(godip.Land).
		// High Sea 2
		Prov("hs2").Conn("so1", godip.Sea).Conn("so2", godip.Sea).Conn("so3", godip.Sea).Conn("so4", godip.Sea).Conn("so5", godip.Sea).Conn("jol", godip.Sea).Conn("sow", godip.Sea).Conn("sob", godip.Sea).Conn("aka", godip.Sea).Conn("big", godip.Sea).Conn("etu", godip.Sea).Flag(godip.Sea).
		// Baol
		Prov("bao").Conn("kay", godip.Land).Conn("sin", godip.Land).Conn("sao", godip.Land).Flag(godip.Land).SC(Jolof).
		// Salaga
		Prov("saa").Conn("aog", godip.Land).Conn("beg", godip.Land).Conn("bon", godip.Land).Conn("ije", godip.Land).Conn("oyo", godip.Land).Conn("gwa", godip.Land).Flag(godip.Land).SC(Bonoman).
		// Akan Sea
		Prov("aka").Conn("big", godip.Sea).Conn("ije", godip.Sea).Conn("bon", godip.Sea).Conn("beg", godip.Sea).Conn("sob", godip.Sea).Conn("hs1", godip.Sea).Conn("hs2", godip.Sea).Conn("hs3", godip.Sea).Flag(godip.Sea).
		// High Sea 3
		Prov("hs3").Conn("so1", godip.Sea).Conn("so2", godip.Sea).Conn("so3", godip.Sea).Conn("so4", godip.Sea).Conn("so5", godip.Sea).Conn("jol", godip.Sea).Conn("sow", godip.Sea).Conn("sob", godip.Sea).Conn("aka", godip.Sea).Conn("big", godip.Sea).Conn("etu", godip.Sea).Flag(godip.Sea).
		// Bilma
		Prov("bil").Conn("aga", godip.Land).Conn("nji", godip.Land).Conn("mas", godip.Land).Conn("abe", godip.Land).Conn("tib", godip.Land).Conn("mur", godip.Land).Conn("gha", godip.Land).Flag(godip.Land).SC(Bornu).
		// Tichitt
		Prov("tic").Conn("awd", godip.Land).Conn("kum", godip.Land).Conn("tav", godip.Land).Conn("waa", godip.Land).Conn("kay", godip.Land).Conn("so1", godip.Land).Conn("so2", godip.Land).Conn("so3", godip.Land).Conn("so4", godip.Land).Conn("so5", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Abesehr
		Prov("abe").Conn("tib", godip.Land).Conn("bil", godip.Land).Conn("mas", godip.Land).Conn("bou", godip.Land).Conn("uba", godip.Land).Flag(godip.Land).
		// Bono Manso
		Prov("bon").Conn("saa", godip.Land).Conn("beg", godip.Coast...).Conn("aka", godip.Sea).Conn("ije", godip.Coast...).Flag(godip.Coast...).SC(Bonoman).
		// Njimi
		Prov("nji").Conn("mas", godip.Land).Conn("bil", godip.Land).Conn("aga", godip.Land).Conn("zaz", godip.Land).Flag(godip.Land).SC(Bornu).
		// Saloum
		Prov("sao").Conn("wan", godip.Land).Conn("bam", godip.Land).Conn("awd", godip.Land).Conn("kay", godip.Land).Conn("bao", godip.Land).Conn("sin", godip.Land).Flag(godip.Land).SC(Jolof).
		// Nufe
		Prov("nuf").Conn("aga", godip.Land).Conn("oyo", godip.Land).Conn("ife", godip.Land).Conn("zaz", godip.Land).Flag(godip.Land).
		// Kongo
		Prov("kon").Conn("bou", godip.Land).Conn("wou", godip.Land).Conn("lun", godip.Land).Conn("uba", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Zazzau
		Prov("zaz").Conn("mas", godip.Land).Conn("nji", godip.Land).Conn("aga", godip.Land).Conn("nuf", godip.Land).Conn("ife", godip.Land).Conn("owo", godip.Land).Conn("wou", godip.Land).Conn("bou", godip.Land).Flag(godip.Land).
		// Edo
		Prov("edo").Conn("ije", godip.Coast...).Conn("big", godip.Sea).Conn("owo", godip.Coast...).Conn("ife", godip.Land).Flag(godip.Coast...).SC(Benin).
		// Niani
		Prov("nia").Conn("bur", godip.Land).Conn("jen", godip.Land).Conn("wal", godip.Land).Conn("kum", godip.Land).Conn("bam", godip.Land).Flag(godip.Land).
		Done()
}

var provinceLongNames = map[godip.Province]string{
	"mos": "Mossi",
	"wou": "Wouri",
	"bur": "Bure",
	"sin": "Sine",
	"big": "Bight of Biafra",
	"jol": "Jolof Sea",
	"bam": "Bambuk",
	"tib": "Tibesti",
	"ije": "Ijebu",
	"waa": "Waalo",
	"uba": "Ubangi",
	"gwa": "Gwandu",
	"gha": "Ghat",
	"mas": "Masseniya",
	"so1": "Sahara Oasis 1",
	"etu": "Etula Eri",
	"so4": "Sahara Oasis 4",
	"so5": "Sahara Oasis 5",
	"mur": "Murzuk",
	"tak": "Takedda",
	"gao": "Gao",
	"oyo": "Oyo",
	"tad": "Tadmekka",
	"ife": "Ife",
	"owo": "Owo",
	"aga": "Agadez",
	"sow": "Sea of Wangara",
	"jen": "Jenne",
	"tav": "Tavdeni",
	"wan": "Wangara",
	"aog": "Aogadougou",
	"kay": "Kayor",
	"bou": "Bouar",
	"awd": "Awdagost",
	"wal": "Walata",
	"lun": "Lunda",
	"beg": "Begho",
	"so2": "Sahara Oasis 2",
	"hs1": "High Sea 1",
	"tim": "Timbuktu",
	"sob": "Sea of Bure",
	"kum": "Kumbi Saleh",
	"so3": "Sahara Oasis 3",
	"hs2": "High Sea 2",
	"bao": "Baol",
	"saa": "Salaga",
	"aka": "Akan Sea",
	"hs3": "High Sea 3",
	"bil": "Bilma",
	"tic": "Tichitt",
	"abe": "Abesehr",
	"bon": "Bono Manso",
	"nji": "Njimi",
	"sao": "Saloum",
	"nuf": "Nufe",
	"kon": "Kongo",
	"zaz": "Zazzau",
	"edo": "Edo",
	"nia": "Niani",
}
