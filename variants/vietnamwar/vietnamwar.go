package vietnamwar

import (
	"github.com/zond/godip"
	"github.com/zond/godip/graph"
	"github.com/zond/godip/state"
	"github.com/zond/godip/variants/classical"
	"github.com/zond/godip/variants/common"
)

const (
	NorthVietnam godip.Nation = "North Vietnam"
	Thailand     godip.Nation = "Thailand"
	SouthVietnam godip.Nation = "South Vietnam"
	Cambodia     godip.Nation = "Cambodia"
	Laos         godip.Nation = "Laos"
)

var Nations = []godip.Nation{NorthVietnam, Thailand, SouthVietnam, Cambodia, Laos}

var SVGFlags = map[godip.Nation]func() ([]byte, error){
	NorthVietnam: func() ([]byte, error) {
		return Asset("svg/northvietnam.svg")
	},
	Thailand: func() ([]byte, error) {
		return Asset("svg/thailand.svg")
	},
	SouthVietnam: func() ([]byte, error) {
		return Asset("svg/southvietnam.svg")
	},
	Cambodia: func() ([]byte, error) {
		return Asset("svg/cambodia.svg")
	},
	Laos: func() ([]byte, error) {
		return Asset("svg/laos.svg")
	},
}

var VietnamWarVariant = common.Variant{
	Name: "Vietnam War",
	NationColors: map[godip.Nation]string{
		SouthVietnam: "#F0D211",
		NorthVietnam: "#FD0000",
		Laos:         "#03853A",
		Thailand:     "#006FFF",
		Cambodia:     "#708D9E",
	},
	Graph:             func() godip.Graph { return VietnamWarGraph() },
	Start:             VietnamWarStart,
	Blank:             VietnamWarBlank,
	Phase:             classical.NewPhase,
	Parser:            classical.Parser,
	Nations:           Nations,
	PhaseTypes:        classical.PhaseTypes,
	Seasons:           classical.Seasons,
	UnitTypes:         classical.UnitTypes,
	SoloWinner:        common.SCCountWinner(15),
	ProvinceLongNames: provinceLongNames,
	SVGMap: func() ([]byte, error) {
		return Asset("svg/vietnamwarmap.svg")
	},
	SVGVersion:  "6",
	SVGUnits:    classical.SVGUnits,
	SVGFlags:    SVGFlags,
	CreatedBy:   "ThePolice",
	Version:     "1.12",
	Description: "The Indochina Peninsula in 1955: the beginning of Vietnam War.",
	SoloSCCount: func(*state.State) int { return 15 },
	Rules: `First to 15 Supply Centers (SC) wins.
All provinces connected to the Mekong river are coastal: Xuyen, Mekong, Pakxe and Ubon (e.g. Laos can build fleets in Pakxe). 
Two provinces have dual coasts: Xuyen and Mekong (South coast and River).`,
}

func VietnamWarBlank(phase godip.Phase) *state.State {
	return state.New(VietnamWarGraph(), phase, classical.BackupRule, nil)
}

func VietnamWarStart() (result *state.State, err error) {
	startPhase := classical.NewPhase(1955, godip.Spring, godip.Movement)
	result = VietnamWarBlank(startPhase)
	if err = result.SetUnits(map[godip.Province]godip.Unit{
		"han": godip.Unit{godip.Fleet, NorthVietnam},
		"thn": godip.Unit{godip.Army, NorthVietnam},
		"thh": godip.Unit{godip.Army, NorthVietnam},
		"pat": godip.Unit{godip.Fleet, Thailand},
		"ban": godip.Unit{godip.Army, Thailand},
		"loe": godip.Unit{godip.Army, Thailand},
		"cam": godip.Unit{godip.Fleet, SouthVietnam},
		"eas": godip.Unit{godip.Fleet, SouthVietnam},
		"sag": godip.Unit{godip.Army, SouthVietnam},
		"pre": godip.Unit{godip.Fleet, Cambodia},
		"ang": godip.Unit{godip.Army, Cambodia},
		"meo": godip.Unit{godip.Army, Cambodia},
		"nah": godip.Unit{godip.Army, Laos},
		"pak": godip.Unit{godip.Army, Laos},
		"vie": godip.Unit{godip.Army, Laos},
	}); err != nil {
		return
	}
	result.SetSupplyCenters(map[godip.Province]godip.Nation{
		"han": NorthVietnam,
		"thn": NorthVietnam,
		"thh": NorthVietnam,
		"pat": Thailand,
		"ban": Thailand,
		"loe": Thailand,
		"cam": SouthVietnam,
		"eas": SouthVietnam,
		"sag": SouthVietnam,
		"pre": Cambodia,
		"ang": Cambodia,
		"meo": Cambodia,
		"nah": Laos,
		"vie": Laos,
		"pak": Laos,
	})
	return
}

func VietnamWarGraph() *graph.Graph {
	return graph.New().
		// Khao Luang
		Prov("kha").Conn("soa", godip.Sea).Conn("cgo", godip.Sea).Conn("noa", godip.Sea).Conn("syk", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Tonle Sap
		Prov("ton").Conn("pre", godip.Coast...).Conn("meo", godip.Land).Conn("kul", godip.Land).Conn("ang", godip.Land).Conn("pat", godip.Coast...).Conn("stc", godip.Sea).Conn("cgo", godip.Sea).Flag(godip.Coast...).
		// North Vietnam
		Prov("nov").Conn("nol", godip.Land).Conn("thn", godip.Land).Flag(godip.Land).
		// Pleiku
		Prov("ple").Conn("cev", godip.Land).Conn("att", godip.Land).Conn("vir", godip.Land).Conn("phn", godip.Land).Conn("tay", godip.Land).Conn("pha", godip.Land).Flag(godip.Land).
		// Mekong Delta
		Prov("med").Conn("coa", godip.Sea).Conn("pha", godip.Sea).Conn("sag", godip.Sea).Conn("xuy", godip.Sea).Conn("xuy/ec", godip.Sea).Conn("cam", godip.Sea).Conn("soa", godip.Sea).Conn("pac", godip.Sea).Flag(godip.Sea).
		// Phnum
		Prov("phn").Conn("vir", godip.Land).Conn("meo", godip.Land).Conn("tay", godip.Land).Conn("ple", godip.Land).Flag(godip.Land).
		// Nan
		Prov("nan").Conn("chi", godip.Land).Conn("phr", godip.Land).Conn("nah", godip.Land).Flag(godip.Land).
		// Nak
		Prov("nak").Conn("phh", godip.Land).Conn("udo", godip.Land).Conn("ubo", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Phonsavan
		Prov("phs").Conn("lua", godip.Land).Conn("vie", godip.Land).Conn("thh", godip.Land).Conn("nol", godip.Land).Flag(godip.Land).
		// Pakxe
		Prov("pak").Conn("kul", godip.Land).Conn("meo", godip.Land).Conn("meo/river", godip.Sea).Conn("vir", godip.Land).Conn("att", godip.Land).Conn("phh", godip.Land).Conn("ubo", godip.Coast...).Flag(godip.Coast...).SC(Laos).
		// North Gulf of Thailand
		Prov("noa").Conn("kha", godip.Sea).Conn("cgo", godip.Sea).Conn("stc", godip.Sea).Conn("pat", godip.Sea).Conn("ban", godip.Sea).Conn("syk", godip.Sea).Flag(godip.Sea).
		// Buri Ram
		Prov("bur").Conn("kul", godip.Land).Conn("ubo", godip.Land).Conn("loe", godip.Land).Conn("ban", godip.Land).Conn("pat", godip.Land).Conn("ang", godip.Land).Flag(godip.Land).
		// Chiang Mai
		Prov("chi").Conn("phr", godip.Land).Conn("nan", godip.Land).Conn("syk", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Kulen Prum
		Prov("kul").Conn("meo", godip.Land).Conn("pak", godip.Land).Conn("ubo", godip.Land).Conn("bur", godip.Land).Conn("ang", godip.Land).Conn("ton", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Angkor Wat
		Prov("ang").Conn("bur", godip.Land).Conn("pat", godip.Land).Conn("ton", godip.Land).Conn("kul", godip.Land).Flag(godip.Land).SC(Cambodia).
		// Hajnan
		Prov("hai").Conn("scs", godip.Sea).Conn("non", godip.Sea).Conn("son", godip.Sea).Conn("coa", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// South  China Sea
		Prov("scs").Conn("non", godip.Sea).Conn("hai", godip.Sea).Conn("coa", godip.Sea).Conn("pac", godip.Sea).Flag(godip.Sea).
		// Vinh
		Prov("vin").Conn("att", godip.Land).Conn("hue", godip.Coast...).Conn("son", godip.Sea).Conn("thh", godip.Coast...).Conn("nkd", godip.Land).Conn("phh", godip.Land).Flag(godip.Coast...).
		// Saigon
		Prov("sag").Conn("med", godip.Sea).Conn("pha", godip.Coast...).Conn("tay", godip.Land).Conn("xuy", godip.Land).Conn("xuy/ec", godip.Sea).Flag(godip.Coast...).SC(SouthVietnam).
		// Ubon
		Prov("ubo").Conn("bur", godip.Land).Conn("kul", godip.Land).Conn("pak", godip.Coast...).Conn("phh", godip.Land).Conn("nak", godip.Land).Conn("udo", godip.Land).Conn("loe", godip.Land).Flag(godip.Coast...).
		// Luang
		Prov("lua").Conn("phs", godip.Land).Conn("nol", godip.Land).Conn("nah", godip.Land).Conn("vie", godip.Land).Flag(godip.Land).
		// Udon
		Prov("udo").Conn("loe", godip.Land).Conn("ubo", godip.Land).Conn("nak", godip.Land).Conn("phh", godip.Land).Conn("nkd", godip.Land).Conn("vie", godip.Land).Flag(godip.Land).
		// Phan
		Prov("pha").Conn("cev", godip.Land).Conn("ple", godip.Land).Conn("tay", godip.Land).Conn("sag", godip.Coast...).Conn("med", godip.Sea).Conn("coa", godip.Sea).Conn("eas", godip.Coast...).Flag(godip.Coast...).
		// Pacific Ocean
		Prov("pac").Conn("scs", godip.Sea).Conn("coa", godip.Sea).Conn("med", godip.Sea).Conn("soa", godip.Sea).Flag(godip.Sea).
		// Thanh
		Prov("thh").Conn("han", godip.Coast...).Conn("nol", godip.Land).Conn("phs", godip.Land).Conn("vin", godip.Coast...).Conn("son", godip.Sea).Conn("non", godip.Sea).Flag(godip.Coast...).SC(NorthVietnam).
		// North Gulf of Tonkin
		Prov("non").Conn("thh", godip.Sea).Conn("son", godip.Sea).Conn("hai", godip.Sea).Conn("scs", godip.Sea).Conn("thn", godip.Sea).Conn("han", godip.Sea).Flag(godip.Sea).
		// Sa Mau
		Prov("cam").Conn("xuy", godip.Land).Conn("xuy/wc", godip.Sea).Conn("xuy/ec", godip.Sea).Conn("soa", godip.Sea).Conn("med", godip.Sea).Flag(godip.Coast...).SC(SouthVietnam).
		// Sai Yok
		Prov("syk").Conn("phr", godip.Land).Conn("chi", godip.Land).Conn("kha", godip.Coast...).Conn("noa", godip.Sea).Conn("ban", godip.Coast...).Flag(godip.Coast...).
		// Wientian
		Prov("vie").Conn("lua", godip.Land).Conn("udo", godip.Land).Conn("nkd", godip.Land).Conn("phs", godip.Land).Flag(godip.Land).SC(Laos).
		// South Gulf of  Thailand
		Prov("soa").Conn("pac", godip.Sea).Conn("med", godip.Sea).Conn("cam", godip.Sea).Conn("xuy", godip.Sea).Conn("xuy/wc", godip.Sea).Conn("cgo", godip.Sea).Conn("kha", godip.Sea).Flag(godip.Sea).
		// Hue
		Prov("hue").Conn("att", godip.Land).Conn("eas", godip.Coast...).Conn("coa", godip.Sea).Conn("son", godip.Sea).Conn("vin", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Tay Ninh
		Prov("tay").Conn("pha", godip.Land).Conn("ple", godip.Land).Conn("phn", godip.Land).Conn("meo", godip.Land).Conn("xuy", godip.Land).Conn("sag", godip.Land).Flag(godip.Land).
		// Coast of Vietnam
		Prov("coa").Conn("med", godip.Sea).Conn("pac", godip.Sea).Conn("scs", godip.Sea).Conn("hai", godip.Sea).Conn("son", godip.Sea).Conn("hue", godip.Sea).Conn("eas", godip.Sea).Conn("pha", godip.Sea).Flag(godip.Sea).
		// Mekong
		Prov("meo").Conn("kul", godip.Land).Conn("ton", godip.Land).Conn("pre", godip.Land).Conn("xuy", godip.Land).Conn("tay", godip.Land).Conn("phn", godip.Land).Conn("vir", godip.Land).Conn("pak", godip.Land).Flag(godip.Land).SC(Cambodia).
		// Mekong (West Coast)
		Prov("meo/wc").Conn("pre", godip.Sea).Conn("cgo", godip.Sea).Conn("xuy/wc", godip.Sea).Flag(godip.Sea).
		// Mekong (River)
		Prov("meo/river").Conn("xuy/ec", godip.Sea).Conn("pak", godip.Sea).Flag(godip.Sea).
		// Centr Viet
		Prov("cev").Conn("pha", godip.Land).Conn("eas", godip.Land).Conn("att", godip.Land).Conn("ple", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// North Laos
		Prov("nol").Conn("nov", godip.Land).Conn("lua", godip.Land).Conn("phs", godip.Land).Conn("thh", godip.Land).Conn("han", godip.Land).Conn("thn", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Bangkok
		Prov("ban").Conn("noa", godip.Sea).Conn("pat", godip.Coast...).Conn("bur", godip.Land).Conn("loe", godip.Land).Conn("phr", godip.Land).Conn("syk", godip.Coast...).Flag(godip.Coast...).SC(Thailand).
		// Virachey
		Prov("vir").Conn("pak", godip.Land).Conn("meo", godip.Land).Conn("phn", godip.Land).Conn("ple", godip.Land).Conn("att", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Hanoi
		Prov("han").Conn("thh", godip.Coast...).Conn("non", godip.Sea).Conn("thn", godip.Coast...).Conn("nol", godip.Land).Flag(godip.Coast...).SC(NorthVietnam).
		// Thai Nguyen
		Prov("thn").Conn("nov", godip.Land).Conn("nol", godip.Land).Conn("han", godip.Coast...).Conn("non", godip.Sea).Flag(godip.Coast...).SC(NorthVietnam).
		// Phrae
		Prov("phr").Conn("syk", godip.Land).Conn("ban", godip.Land).Conn("loe", godip.Land).Conn("nan", godip.Land).Conn("chi", godip.Land).Flag(godip.Land).
		// Phou Hin
		Prov("phh").Conn("vin", godip.Land).Conn("nkd", godip.Land).Conn("udo", godip.Land).Conn("nak", godip.Land).Conn("ubo", godip.Land).Conn("pak", godip.Land).Conn("att", godip.Land).Flag(godip.Land).
		// Nam Ha
		Prov("nah").Conn("lua", godip.Land).Conn("nan", godip.Land).Flag(godip.Land).SC(Laos).
		// Preah
		Prov("pre").Conn("meo", godip.Land).Conn("meo/wc", godip.Sea).Conn("ton", godip.Coast...).Conn("cgo", godip.Sea).Flag(godip.Coast...).SC(Cambodia).
		// Pattaya
		Prov("pat").Conn("ton", godip.Coast...).Conn("ang", godip.Land).Conn("bur", godip.Land).Conn("ban", godip.Coast...).Conn("noa", godip.Sea).Conn("stc", godip.Sea).Flag(godip.Coast...).SC(Thailand).
		// South Thailand Coast
		Prov("stc").Conn("noa", godip.Sea).Conn("cgo", godip.Sea).Conn("ton", godip.Sea).Conn("pat", godip.Sea).Flag(godip.Sea).
		// Attapu
		Prov("att").Conn("hue", godip.Land).Conn("vin", godip.Land).Conn("phh", godip.Land).Conn("pak", godip.Land).Conn("vir", godip.Land).Conn("ple", godip.Land).Conn("cev", godip.Land).Conn("eas", godip.Land).Flag(godip.Land).
		// East Coast
		Prov("eas").Conn("att", godip.Land).Conn("cev", godip.Land).Conn("pha", godip.Coast...).Conn("coa", godip.Sea).Conn("hue", godip.Coast...).Flag(godip.Coast...).SC(SouthVietnam).
		// Nam Kading
		Prov("nkd").Conn("vie", godip.Land).Conn("udo", godip.Land).Conn("phh", godip.Land).Conn("vin", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// South Gulf of Tonkin
		Prov("son").Conn("vin", godip.Sea).Conn("hue", godip.Sea).Conn("coa", godip.Sea).Conn("hai", godip.Sea).Conn("non", godip.Sea).Conn("thh", godip.Sea).Flag(godip.Sea).
		// Xuyen
		Prov("xuy").Conn("cam", godip.Land).Conn("sag", godip.Land).Conn("tay", godip.Land).Conn("meo", godip.Land).Flag(godip.Land).
		// Xuyen (East Coast)
		Prov("xuy/ec").Conn("cam", godip.Sea).Conn("med", godip.Sea).Conn("sag", godip.Sea).Conn("meo/river", godip.Sea).Flag(godip.Sea).
		// Xuyen (West Coast)
		Prov("xuy/wc").Conn("cam", godip.Sea).Conn("meo/wc", godip.Sea).Conn("cgo", godip.Sea).Conn("soa", godip.Sea).Flag(godip.Sea).
		// Central Gulf of Thailand
		Prov("cgo").Conn("meo", godip.Sea).Conn("meo/wc", godip.Sea).Conn("pre", godip.Sea).Conn("ton", godip.Sea).Conn("stc", godip.Sea).Conn("noa", godip.Sea).Conn("kha", godip.Sea).Conn("soa", godip.Sea).Conn("xuy", godip.Sea).Conn("xuy/wc", godip.Sea).Flag(godip.Sea).
		// Loei
		Prov("loe").Conn("udo", godip.Land).Conn("phr", godip.Land).Conn("ban", godip.Land).Conn("bur", godip.Land).Conn("ubo", godip.Land).Flag(godip.Land).SC(Thailand).
		Done()
}

var provinceLongNames = map[godip.Province]string{
	"soa":       "South Gulf of Thailand",
	"pre":       "Preah",
	"vie":       "Vientiane",
	"ang":       "Angkor Wat",
	"cam":       "Ca Mau",
	"hue":       "Hue",
	"sag":       "Saigon",
	"tay":       "Tay Ninh",
	"son":       "South Gulf of Tonkin",
	"vir":       "Virachey",
	"kha":       "Khao Luang",
	"eas":       "East Coast",
	"hai":       "Hainan",
	"nkd":       "Nam Kading",
	"pat":       "Pattaya",
	"han":       "Hanoi",
	"loe":       "Loei",
	"nak":       "Nak",
	"nah":       "Nam Ha",
	"nan":       "Nan",
	"lua":       "Luang",
	"kul":       "Kulen Prum",
	"ton":       "Tonle Sap",
	"ban":       "Bangkok",
	"phh":       "Phou Hin",
	"pak":       "Pakxe",
	"meo":       "Mekong",
	"nol":       "Dien Bien",
	"med":       "Mekong Delta",
	"non":       "North Gulf of Tonkin",
	"phn":       "Phnum",
	"cgo":       "Central Gulf of Thailand",
	"bur":       "Buri Ram",
	"vin":       "Vinh",
	"pha":       "Phan",
	"thh":       "Thanh",
	"noa":       "North Gulf of Thailand",
	"cev":       "Central Vietnam",
	"ple":       "Pleiku",
	"phs":       "Xiangkhouang",
	"phr":       "Phrae",
	"nov":       "North Vietnam",
	"ubo":       "Ubon",
	"udo":       "Udon",
	"syk":       "Sai Yok",
	"xuy":       "Xuyen",
	"stc":       "South Thailand Coast",
	"pac":       "Pacific Ocean",
	"coa":       "Coast of Vietnam",
	"thn":       "Thai Nguyen",
	"att":       "Attapeu",
	"chi":       "Chiang Mai",
	"scs":       "South China Sea",
	"xuy/ec":    "Xuyen (EC)",
	"xuy/wc":    "Xuyen (WC)",
	"meo/wc":    "Mekong (WC)",
	"meo/river": "Mekong (River)",
}
