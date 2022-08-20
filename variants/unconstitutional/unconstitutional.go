package unconstitutional

import (
	"time"
	"github.com/zond/godip"
	"github.com/zond/godip/graph"
	"github.com/zond/godip/state"
	"github.com/zond/godip/orders"
	"github.com/zond/godip/variants/classical"
	"github.com/zond/godip/variants/common"
	"github.com/zond/godip/variants/hundred"

)

const (
	SouthCarolina        godip.Nation = "South Carolina"
	NewYork              godip.Nation = "New York"
	WesternConfederacy   godip.Nation = "Western Confederacy"
	Pennsylvania         godip.Nation = "Pennsylvania"
	MuskogeeConfederacy  godip.Nation = "Muskogee Confederacy"
	Virginia             godip.Nation = "Virginia"
)

var Nations = []godip.Nation{SouthCarolina, NewYork, WesternConfederacy, Pennsylvania, MuskogeeConfederacy, Virginia}

var SVGFlags = map[godip.Nation]func() ([]byte, error){
	SouthCarolina: func() ([]byte, error) {
		return Asset("svg/southcarolina.svg")
	},
	NewYork: func() ([]byte, error) {
		return Asset("svg/newyork.svg")
	},
	WesternConfederacy: func() ([]byte, error) {
		return Asset("svg/westernconfederacy.svg")
	},
	Pennsylvania: func() ([]byte, error) {
		return Asset("svg/pennsylvania.svg")
	},
	MuskogeeConfederacy: func() ([]byte, error) {
		return Asset("svg/muskogee.svg")
	},
	Virginia: func() ([]byte, error) {
		return Asset("svg/virginia.svg")
	},
}

var UnconstitutionalVariant = common.Variant{
	Name:              "Unconstitutional",
	Graph:             func() godip.Graph { return UnconstitutionalGraph() },
	Start:             UnconstitutionalStart,
	Blank:             UnconstitutionalBlank,
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
		return Asset("svg/unconstitutionalmap.svg")
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
	CreatedBy:   "David E. Cohen",
	Version:     "1.0",
	Description: "Alternative history variant where the US constitution was not ratified (which nearly happened). Operating under the weak Articles of Confederation, States keep their conflicting land claims and border disputes turn into armed conflict. Former slaves control Haiti, and inhabitants of New Orleans, Saint Louis and the Turks and Cacois oppose annexation by the US. Federal government ceases to function, many States have seceded and two groups of Native American tribes, the Western and Muskogee Confederacy, are warning the Americans.",
	Rules:       "First to 18 Supply Centers (SC) wins. Units may be built in any owned SC. Neutral SCs get an army which always holds and disbands when dislodged. This will be rebuilt if the SC is unowned during adjustment. There are four rivers where fleets can navigate to any adjacent province",
}


func NeutralOrders(state state.State) (ret map[godip.Province]godip.Adjudicator) {
	ret = map[godip.Province]godip.Adjudicator{}
	switch state.Phase().Type() {
	case godip.Movement:
		// Strictly this is unnecessary - because hold is the default order.
		for prov, unit := range state.Units() {
			if unit.Nation == godip.Neutral {
				ret[prov] = orders.Hold(prov)
			}
		}
	case godip.Adjustment:
		// Rebuild any missing units.
		for _, prov := range state.Graph().AllSCs() {
			if n, _, ok := state.SupplyCenter(prov); ok && n == godip.Neutral {
				if _, _, ok := state.Unit(prov); !ok {
					ret[prov] = orders.BuildAnywhere(prov, godip.Army, time.Now())
				}
			}
		}
	}
	return
}

func UnconstitutionalBlank(phase godip.Phase) *state.State {
	return state.New(UnconstitutionalGraph(), phase, classical.BackupRule, nil, nil)
}

func UnconstitutionalStart() (result *state.State, err error) {
	startPhase := classical.NewPhase(1805, godip.Spring, godip.Movement)
	result = UnconstitutionalBlank(startPhase)
	if err = result.SetUnits(map[godip.Province]godip.Unit{
		"chn": godip.Unit{godip.Fleet, SouthCarolina},
		"col": godip.Unit{godip.Army, SouthCarolina},
		"bea": godip.Unit{godip.Army, SouthCarolina},
		"nyc": godip.Unit{godip.Fleet, NewYork},
		"loi": godip.Unit{godip.Fleet, NewYork},
		"alb": godip.Unit{godip.Army, NewYork},
		"pro": godip.Unit{godip.Army, WesternConfederacy},
		"kek": godip.Unit{godip.Army, WesternConfederacy},
		"wap": godip.Unit{godip.Army, WesternConfederacy},
		"phi": godip.Unit{godip.Fleet, Pennsylvania},
		"pit": godip.Unit{godip.Army, Pennsylvania},
		"har": godip.Unit{godip.Army, Pennsylvania},
		"yor": godip.Unit{godip.Army, Pennsylvania},
		"mic": godip.Unit{godip.Fleet, MuskogeeConfederacy},
		"tuk": godip.Unit{godip.Army, MuskogeeConfederacy},
		"cus": godip.Unit{godip.Army, MuskogeeConfederacy},
		"wil": godip.Unit{godip.Fleet, Virginia},
		"ale": godip.Unit{godip.Army, Virginia},
		"chv": godip.Unit{godip.Army, Virginia},
		"ric": godip.Unit{godip.Army, Virginia},
	}); err != nil {
		return
	}
	result.SetSupplyCenters(map[godip.Province]godip.Nation{
		"chn": SouthCarolina,
		"col": SouthCarolina,
		"bea": SouthCarolina,
		"nyc": NewYork,
		"loi": NewYork,
		"alb": NewYork,
		"pro": WesternConfederacy,
		"kek": WesternConfederacy,
		"wap": WesternConfederacy,
		"phi": Pennsylvania,
		"pit": Pennsylvania,
		"har": Pennsylvania,
		"yor": Pennsylvania,
		"mic": MuskogeeConfederacy,
		"tuk": MuskogeeConfederacy,
		"cus": MuskogeeConfederacy,
		"wil": Virginia,
		"ale": Virginia,
		"chv": Virginia,
		"ric": Virginia,
		"cho": godip.Neutral,
		"neo": godip.Neutral,
		"sal": godip.Neutral,
		"eas": godip.Neutral,
		"tur": godip.Neutral,
		"por": godip.Neutral,
		"sad": godip.Neutral,
		"chk": godip.Neutral,
		"noc": godip.Neutral,
		"ken": godip.Neutral,
		"wer": godip.Neutral,
		"mar": godip.Neutral,
		"nej": godip.Neutral,
		"mas": godip.Neutral,
		"neh": godip.Neutral,
	})
	return
}

func UnconstitutionalGraph() *graph.Graph {
	return graph.New().
		// Alexandria
		Prov("ale").Conn("she", godip.Land).Conn("chv", godip.Land).Conn("rap", godip.Land).Conn("mar", godip.Land).Conn("upp", godip.Land).Flag(godip.Land).SC(Virginia).
		// New York City
		Prov("nyc").Conn("loi", godip.Coast...).Conn("lis", godip.Sea).Conn("con", godip.Coast...).Conn("alb", godip.Coast...).Conn("cat", godip.Coast...).Conn("nej", godip.Coast...).Conn("nyb", godip.Sea).Flag(godip.Coast...).SC(NewYork).
		// Western Reserve
		Prov("wer").Conn("lyc", godip.Land).Conn("cat", godip.Land).Conn("iro", godip.Land).Conn("det", godip.Land).Conn("ohi", godip.Land).Conn("all", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Delaware
		Prov("der").Conn("nej", godip.Coast...).Conn("phi", godip.Coast...).Conn("yor", godip.Land).Conn("mar", godip.Coast...).Conn("deb", godip.Sea).Flag(godip.Coast...).
		// Philadelphia
		Prov("phi").Conn("der", godip.Coast...).Conn("nej", godip.Coast...).Conn("lyc", godip.Land).Conn("har", godip.Land).Conn("yor", godip.Land).Flag(godip.Coast...).SC(Pennsylvania).
		// Detroit
		Prov("det").Conn("wer", godip.Land).Conn("pot", godip.Land).Conn("kek", godip.Land).Conn("wap", godip.Land).Conn("ohi", godip.Land).Flag(godip.Land).
		// Maryland
		Prov("mar").Conn("upp", godip.Land).Conn("ale", godip.Land).Conn("rap", godip.Coast...).Conn("chb", godip.Sea).Conn("deb", godip.Sea).Conn("der", godip.Coast...).Conn("yor", godip.Land).Flag(godip.Coast...).SC(godip.Neutral).
		// Rappahannock
		Prov("rap").Conn("mar", godip.Coast...).Conn("ale", godip.Land).Conn("chv", godip.Land).Conn("ric", godip.Land).Conn("wil", godip.Coast...).Conn("chb", godip.Sea).Flag(godip.Coast...).
		// Cibao
		Prov("cib").Conn("azu", godip.Land).Conn("sad", godip.Coast...).Conn("sar", godip.Sea).Conn("old", godip.Sea).Conn("art", godip.Coast...).Flag(godip.Coast...).
		// Port au Prince
		Prov("por").Conn("azu", godip.Coast...).Conn("art", godip.Coast...).Conn("win", godip.Sea).Conn("gul", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Turks and Caicos
		Prov("tur").Conn("old", godip.Sea).Conn("sar", godip.Sea).Conn("bah", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Prophetstown
		Prov("pro").Conn("pen", godip.Coast...).Conn("wap", godip.Land).Conn("kek", godip.Land).Conn("pot", godip.Land).Conn("hoc", godip.Land).Conn("ill", godip.Coast...).Conn("sal", godip.Coast...).Conn("chi", godip.Coast...).Flag(godip.Coast...).SC(WesternConfederacy).
		// Azua
		Prov("azu").Conn("cib", godip.Land).Conn("art", godip.Land).Conn("por", godip.Coast...).Conn("win", godip.Sea).Conn("sad", godip.Coast...).Flag(godip.Coast...).
		// Tidewater
		Prov("tid").Conn("ric", godip.Coast...).Conn("noc", godip.Coast...).Conn("out", godip.Sea).Conn("chb", godip.Sea).Conn("wil", godip.Coast...).Flag(godip.Coast...).
		// Gulf of Mexico
		Prov("guc").Conn("win", godip.Sea).Conn("old", godip.Sea).Conn("flo", godip.Sea).Conn("cad", godip.Sea).Conn("neo", godip.Sea).Flag(godip.Sea).
		// Connecticut
		Prov("con").Conn("alb", godip.Land).Conn("nyc", godip.Coast...).Conn("lis", godip.Sea).Conn("mas", godip.Coast...).Flag(godip.Coast...).
		// Tennessee
		Prov("ten").Conn("chk", godip.Land).Conn("fra", godip.Land).Conn("pen", godip.Land).Conn("chi", godip.Land).Conn("nic", godip.Land).Flag(godip.Land).
		// Alabama
		Prov("ala").Conn("tuk", godip.Land).Conn("nic", godip.Land).Conn("chi", godip.Land).Conn("cho", godip.Land).Conn("wef", godip.Land).Flag(godip.Land).
		// Georgia Bight
		Prov("geb").Conn("eas", godip.Sea).Conn("bah", godip.Sea).Conn("sar", godip.Sea).Conn("out", godip.Sea).Conn("noc", godip.Sea).Conn("chn", godip.Sea).Conn("bea", godip.Sea).Conn("ger", godip.Sea).Flag(godip.Sea).
		// Outer Banks
		Prov("out").Conn("noc", godip.Sea).Conn("geb", godip.Sea).Conn("sar", godip.Sea).Conn("chb", godip.Sea).Conn("tid", godip.Sea).Flag(godip.Sea).
		// Lycoming
		Prov("lyc").Conn("all", godip.Land).Conn("har", godip.Land).Conn("phi", godip.Land).Conn("nej", godip.Land).Conn("cat", godip.Land).Conn("wer", godip.Land).Flag(godip.Land).
		// Sargasso Sea
		Prov("sar").Conn("cib", godip.Sea).Conn("sad", godip.Sea).Conn("atl", godip.Sea).Conn("gua", godip.Sea).Conn("nyb", godip.Sea).Conn("chb", godip.Sea).Conn("out", godip.Sea).Conn("geb", godip.Sea).Conn("bah", godip.Sea).Conn("tur", godip.Sea).Conn("old", godip.Sea).Flag(godip.Sea).
		// Caddo
		Prov("cad").Conn("guc", godip.Sea).Conn("neo", godip.Coast...).Conn("qua", godip.Land).Flag(godip.Coast...).
		// Harrisburg
		Prov("har").Conn("yor", godip.Land).Conn("phi", godip.Land).Conn("lyc", godip.Land).Conn("all", godip.Land).Flag(godip.Land).SC(Pennsylvania).
		// Quapow
		Prov("qua").Conn("neo", godip.Coast...).Conn("cad", godip.Land).Conn("cho", godip.Coast...).Conn("sal", godip.Coast...).Conn("osa", godip.Land).Flag(godip.Coast...).
		// Charleston
		Prov("chn").Conn("mid", godip.Land).Conn("col", godip.Land).Conn("bea", godip.Coast...).Conn("geb", godip.Sea).Conn("noc", godip.Coast...).Flag(godip.Coast...).SC(SouthCarolina).
		// Albany
		Prov("alb").Conn("iro", godip.Land).Conn("cat", godip.Coast...).Conn("nyc", godip.Coast...).Conn("con", godip.Land).Conn("mas", godip.Land).Conn("ver", godip.Land).Flag(godip.Coast...).SC(NewYork).
		// Vermont
		Prov("ver").Conn("iro", godip.Land).Conn("alb", godip.Land).Conn("mas", godip.Land).Conn("neh", godip.Land).Flag(godip.Land).
		// Gulf of Gonave
		Prov("gul").Conn("art", godip.Sea).Conn("old", godip.Sea).Conn("win", godip.Sea).Conn("por", godip.Sea).Flag(godip.Sea).
		// Delaware Bay
		Prov("deb").Conn("chb", godip.Sea).Conn("nyb", godip.Sea).Conn("nej", godip.Sea).Conn("der", godip.Sea).Conn("mar", godip.Sea).Flag(godip.Sea).
		// Windward Passage
		Prov("win").Conn("atl", godip.Sea).Conn("sad", godip.Sea).Conn("azu", godip.Sea).Conn("por", godip.Sea).Conn("gul", godip.Sea).Conn("old", godip.Sea).Conn("guc", godip.Sea).Flag(godip.Sea).
		// Midlands
		Prov("mid").Conn("chn", godip.Land).Conn("noc", godip.Land).Conn("nop", godip.Land).Conn("fra", godip.Land).Conn("col", godip.Land).Flag(godip.Land).
		// New Orleans
		Prov("neo").Conn("guc", godip.Sea).Conn("flo", godip.Sea).Conn("cad", godip.Coast...).Conn("cho", godip.Land).Conn("qua", godip.Coast...).Conn("wef", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Shenandoah
		Prov("she").Conn("wet", godip.Land).Conn("ken", godip.Land).Conn("ric", godip.Land).Conn("chv", godip.Land).Conn("ale", godip.Land).Conn("upp", godip.Land).Flag(godip.Land).
		// Osage
		Prov("osa").Conn("qua", godip.Land).Conn("sal", godip.Coast...).Conn("mis", godip.Coast...).Flag(godip.Coast...).
		// Beaufort
		Prov("bea").Conn("geb", godip.Sea).Conn("chn", godip.Coast...).Conn("col", godip.Land).Conn("ger", godip.Coast...).Flag(godip.Coast...).SC(SouthCarolina).
		// New Hampshire
		Prov("neh").Conn("mai", godip.Coast...).Conn("ver", godip.Land).Conn("mas", godip.Coast...).Conn("gua", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Atlantic Ocean
		Prov("atl").Conn("gua", godip.Sea).Conn("sar", godip.Sea).Conn("sad", godip.Sea).Conn("win", godip.Sea).Flag(godip.Sea).
		// New Jersey
		Prov("nej").Conn("nyb", godip.Sea).Conn("nyc", godip.Coast...).Conn("cat", godip.Coast...).Conn("lyc", godip.Land).Conn("phi", godip.Coast...).Conn("der", godip.Coast...).Conn("deb", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Wapakoneta
		Prov("wap").Conn("kek", godip.Land).Conn("pro", godip.Land).Conn("pen", godip.Land).Conn("ken", godip.Land).Conn("ohi", godip.Land).Conn("det", godip.Land).Flag(godip.Land).SC(WesternConfederacy).
		// South Piedmont
		Prov("sou").Conn("ger", godip.Land).Conn("col", godip.Land).Conn("chk", godip.Land).Conn("cus", godip.Land).Conn("sem", godip.Land).Flag(godip.Land).
		// Catskill
		Prov("cat").Conn("nej", godip.Coast...).Conn("nyc", godip.Coast...).Conn("alb", godip.Coast...).Conn("iro", godip.Land).Conn("wer", godip.Land).Conn("lyc", godip.Land).Flag(godip.Coast...).
		// New York Bight
		Prov("nyb").Conn("nej", godip.Sea).Conn("deb", godip.Sea).Conn("chb", godip.Sea).Conn("sar", godip.Sea).Conn("gua", godip.Sea).Conn("lis", godip.Sea).Conn("loi", godip.Sea).Conn("nyc", godip.Sea).Flag(godip.Sea).
		// Artibonite
		Prov("art").Conn("azu", godip.Land).Conn("cib", godip.Coast...).Conn("old", godip.Sea).Conn("gul", godip.Sea).Conn("por", godip.Coast...).Flag(godip.Coast...).
		// Columbia
		Prov("col").Conn("chn", godip.Land).Conn("mid", godip.Land).Conn("fra", godip.Land).Conn("chk", godip.Land).Conn("sou", godip.Land).Conn("ger", godip.Land).Conn("bea", godip.Land).Flag(godip.Land).SC(SouthCarolina).
		// Chesapeake Bay
		Prov("chb").Conn("deb", godip.Sea).Conn("mar", godip.Sea).Conn("rap", godip.Sea).Conn("wil", godip.Sea).Conn("tid", godip.Sea).Conn("out", godip.Sea).Conn("sar", godip.Sea).Conn("nyb", godip.Sea).Flag(godip.Sea).
		// Ohio
		Prov("ohi").Conn("all", godip.Land).Conn("wer", godip.Land).Conn("det", godip.Land).Conn("wap", godip.Land).Conn("ken", godip.Land).Conn("wet", godip.Land).Flag(godip.Land).
		// York
		Prov("yor").Conn("upp", godip.Land).Conn("mar", godip.Land).Conn("der", godip.Land).Conn("phi", godip.Land).Conn("har", godip.Land).Conn("all", godip.Land).Flag(godip.Land).SC(Pennsylvania).
		// Iroquois
		Prov("iro").Conn("alb", godip.Land).Conn("ver", godip.Land).Conn("wer", godip.Land).Conn("cat", godip.Land).Flag(godip.Land).
		// East Florida
		Prov("eas").Conn("geb", godip.Sea).Conn("ger", godip.Coast...).Conn("sem", godip.Coast...).Conn("flo", godip.Sea).Conn("bah", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Tukabatchee
		Prov("tuk").Conn("ala", godip.Land).Conn("wef", godip.Coast...).Conn("mic", godip.Coast...).Conn("cus", godip.Coast...).Conn("nic", godip.Land).Flag(godip.Coast...).SC(MuskogeeConfederacy).
		// Richmond
		Prov("ric").Conn("tid", godip.Coast...).Conn("wil", godip.Coast...).Conn("rap", godip.Land).Conn("chv", godip.Land).Conn("she", godip.Land).Conn("ken", godip.Land).Conn("fra", godip.Land).Conn("nop", godip.Land).Conn("noc", godip.Land).Flag(godip.Coast...).SC(Virginia).
		// Saint Domingue
		Prov("sad").Conn("azu", godip.Coast...).Conn("win", godip.Sea).Conn("atl", godip.Sea).Conn("sar", godip.Sea).Conn("cib", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// North Carolina
		Prov("noc").Conn("out", godip.Sea).Conn("tid", godip.Coast...).Conn("ric", godip.Land).Conn("nop", godip.Land).Conn("mid", godip.Land).Conn("chn", godip.Coast...).Conn("geb", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Potawatomi
		Prov("pot").Conn("kek", godip.Land).Conn("det", godip.Land).Conn("hoc", godip.Land).Conn("pro", godip.Land).Flag(godip.Land).
		// West Florida
		Prov("wef").Conn("cho", godip.Coast...).Conn("neo", godip.Coast...).Conn("flo", godip.Sea).Conn("mic", godip.Coast...).Conn("tuk", godip.Coast...).Conn("ala", godip.Land).Flag(godip.Coast...).
		// Gulf of Maine
		Prov("gua").Conn("mai", godip.Sea).Conn("neh", godip.Sea).Conn("mas", godip.Sea).Conn("lis", godip.Sea).Conn("nyb", godip.Sea).Conn("sar", godip.Sea).Conn("atl", godip.Sea).Flag(godip.Sea).
		// Seminole
		Prov("sem").Conn("cus", godip.Land).Conn("mic", godip.Coast...).Conn("flo", godip.Sea).Conn("eas", godip.Coast...).Conn("ger", godip.Land).Conn("sou", godip.Land).Flag(godip.Coast...).
		// Missouri
		Prov("mis").Conn("osa", godip.Coast...).Conn("sal", godip.Coast...).Conn("ill", godip.Coast...).Conn("hoc", godip.Coast...).Flag(godip.Coast...).
		// Cherokee
		Prov("chk").Conn("ten", godip.Land).Conn("nic", godip.Land).Conn("cus", godip.Land).Conn("sou", godip.Land).Conn("col", godip.Land).Conn("fra", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Saint Louis
		Prov("sal").Conn("osa", godip.Coast...).Conn("qua", godip.Coast...).Conn("cho", godip.Coast...).Conn("chi", godip.Coast...).Conn("pro", godip.Coast...).Conn("ill", godip.Coast...).Conn("mis", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Williamsburg
		Prov("wil").Conn("ric", godip.Coast...).Conn("tid", godip.Coast...).Conn("chb", godip.Sea).Conn("rap", godip.Coast...).Flag(godip.Coast...).SC(Virginia).
		// Long Island
		Prov("loi").Conn("lis", godip.Sea).Conn("nyc", godip.Coast...).Conn("nyb", godip.Sea).Flag(godip.Coast...).SC(NewYork).
		// North Piedmont
		Prov("nop").Conn("mid", godip.Land).Conn("noc", godip.Land).Conn("ric", godip.Land).Conn("fra", godip.Land).Flag(godip.Land).
		// Pittsburgh
		Prov("pit").Conn("all", godip.Land).Conn("all", godip.Land).Conn("wet", godip.Land).Flag(godip.Land).SC(Pennsylvania).
		// Georgia
		Prov("ger").Conn("sou", godip.Land).Conn("sem", godip.Land).Conn("eas", godip.Coast...).Conn("geb", godip.Sea).Conn("bea", godip.Coast...).Conn("col", godip.Land).Flag(godip.Coast...).
		// Kekionga
		Prov("kek").Conn("pot", godip.Land).Conn("pro", godip.Land).Conn("wap", godip.Land).Conn("det", godip.Land).Flag(godip.Land).SC(WesternConfederacy).
		// Allegheny
		Prov("all").Conn("lyc", godip.Land).Conn("wer", godip.Land).Conn("ohi", godip.Land).Conn("wet", godip.Land).Conn("pit", godip.Land).Conn("pit", godip.Land).Conn("wet", godip.Land).Conn("upp", godip.Land).Conn("yor", godip.Land).Conn("har", godip.Land).Flag(godip.Land).
		// Pennyrile
		Prov("pen").Conn("fra", godip.Land).Conn("ken", godip.Land).Conn("wap", godip.Land).Conn("pro", godip.Coast...).Conn("chi", godip.Coast...).Conn("ten", godip.Land).Flag(godip.Coast...).
		// Illiniwek
		Prov("ill").Conn("pro", godip.Coast...).Conn("hoc", godip.Coast...).Conn("mis", godip.Coast...).Conn("sal", godip.Coast...).Flag(godip.Coast...).
		// Miccosukee
		Prov("mic").Conn("tuk", godip.Coast...).Conn("wef", godip.Coast...).Conn("flo", godip.Sea).Conn("sem", godip.Coast...).Conn("cus", godip.Coast...).Flag(godip.Coast...).SC(MuskogeeConfederacy).
		// Kentucky
		Prov("ken").Conn("she", godip.Land).Conn("wet", godip.Land).Conn("ohi", godip.Land).Conn("wap", godip.Land).Conn("pen", godip.Land).Conn("fra", godip.Land).Conn("ric", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Choctaw
		Prov("cho").Conn("wef", godip.Coast...).Conn("ala", godip.Land).Conn("chi", godip.Coast...).Conn("sal", godip.Coast...).Conn("qua", godip.Coast...).Conn("neo", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Long Island Sound
		Prov("lis").Conn("loi", godip.Sea).Conn("nyb", godip.Sea).Conn("gua", godip.Sea).Conn("mas", godip.Sea).Conn("con", godip.Sea).Conn("nyc", godip.Sea).Flag(godip.Sea).
		// Ho Chunk
		Prov("hoc").Conn("mis", godip.Coast...).Conn("ill", godip.Coast...).Conn("pro", godip.Land).Conn("pot", godip.Land).Flag(godip.Coast...).
		// Chickasaw
		Prov("chi").Conn("ten", godip.Land).Conn("pen", godip.Coast...).Conn("pro", godip.Coast...).Conn("sal", godip.Coast...).Conn("cho", godip.Coast...).Conn("ala", godip.Land).Conn("nic", godip.Land).Flag(godip.Coast...).
		// Franklin
		Prov("fra").Conn("pen", godip.Land).Conn("ten", godip.Land).Conn("chk", godip.Land).Conn("col", godip.Land).Conn("mid", godip.Land).Conn("nop", godip.Land).Conn("ric", godip.Land).Conn("ken", godip.Land).Flag(godip.Land).
		// Charlottesville
		Prov("chv").Conn("she", godip.Land).Conn("ric", godip.Land).Conn("rap", godip.Land).Conn("ale", godip.Land).Flag(godip.Land).SC(Virginia).
		// Maine
		Prov("mai").Conn("neh", godip.Coast...).Conn("gua", godip.Sea).Flag(godip.Coast...).
		// Upper Pontomac
		Prov("upp").Conn("yor", godip.Land).Conn("all", godip.Land).Conn("wet", godip.Land).Conn("she", godip.Land).Conn("ale", godip.Land).Conn("mar", godip.Land).Flag(godip.Land).
		// Florida Bight
		Prov("flo").Conn("bah", godip.Sea).Conn("eas", godip.Sea).Conn("sem", godip.Sea).Conn("mic", godip.Sea).Conn("neo", godip.Sea).Conn("wef", godip.Sea).Conn("guc", godip.Sea).Conn("old", godip.Sea).Flag(godip.Sea).
		// Bahama Banks
		Prov("bah").Conn("flo", godip.Sea).Conn("old", godip.Sea).Conn("tur", godip.Sea).Conn("sar", godip.Sea).Conn("geb", godip.Sea).Conn("eas", godip.Sea).Flag(godip.Sea).
		// Old Bahama Channel
		Prov("old").Conn("tur", godip.Sea).Conn("bah", godip.Sea).Conn("flo", godip.Sea).Conn("guc", godip.Sea).Conn("gul", godip.Sea).Conn("win", godip.Sea).Conn("art", godip.Sea).Conn("cib", godip.Sea).Conn("sar", godip.Sea).Flag(godip.Sea).
		// Nickajack
		Prov("nic").Conn("cus", godip.Land).Conn("chk", godip.Land).Conn("ten", godip.Land).Conn("chi", godip.Land).Conn("ala", godip.Land).Conn("tuk", godip.Land).Flag(godip.Land).
		// Westsylvania
		Prov("wet").Conn("all", godip.Land).Conn("ohi", godip.Land).Conn("ken", godip.Land).Conn("she", godip.Land).Conn("upp", godip.Land).Conn("all", godip.Land).Conn("pit", godip.Land).Flag(godip.Land).
		// Cusseta
		Prov("cus").Conn("nic", godip.Land).Conn("tuk", godip.Land).Conn("mic", godip.Land).Conn("sem", godip.Land).Conn("sou", godip.Land).Conn("chk", godip.Land).Flag(godip.Land).SC(MuskogeeConfederacy).
		// Massachuesetts
		Prov("mas").Conn("lis", godip.Sea).Conn("gua", godip.Sea).Conn("neh", godip.Coast...).Conn("ver", godip.Land).Conn("alb", godip.Land).Conn("con", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		Done()
}

var provinceLongNames = map[godip.Province]string{
	"ale": "Alexandria",
	"nyc": "New York City",
	"wer": "Western Reserve",
	"der": "Delaware",
	"phi": "Philadelphia",
	"det": "Detroit",
	"mar": "Maryland",
	"rap": "Rappahannock",
	"cib": "Cibao",
	"por": "Port au Prince",
	"tur": "Turks and Caicos",
	"pro": "Prophetstown",
	"azu": "Azua",
	"tid": "Tidewater",
	"guc": "Gulf of Mexico",
	"con": "Connecticut",
	"ten": "Tennessee",
	"ala": "Alabama",
	"geb": "Georgia Bight",
	"out": "Outer Banks",
	"lyc": "Lycoming",
	"sar": "Sargasso Sea",
	"neo": "New Orleans",
	"har": "Harrisburg",
	"qua": "Quapow",
	"chn": "Charleston",
	"alb": "Albany",
	"ver": "Vermont",
	"win": "Windward Passage",
	"deb": "Delaware Bay",
	"mid": "Midlands",
	"wef": "West Florida",
	"she": "Shenandoah",
	"osa": "Osage",
	"bea": "Beaufort",
	"neh": "New Hampshire",
	"atl": "Atlantic Ocean",
	"nej": "New Jersey",
	"wap": "Wapakoneta",
	"sou": "South Piedmont",
	"cat": "Catskill",
	"nyb": "New York Bight",
	"art": "Artibonite",
	"col": "Columbia",
	"chb": "Chesapeake Bay",
	"ohi": "Ohio",
	"yor": "York",
	"iro": "Iroquois",
	"gua": "Gulf of Maine",
	"eas": "East Florida",
	"tuk": "Tukabatchee",
	"ric": "Richmond",
	"sad": "Saint Domingue",
	"noc": "North Carolina",
	"pot": "Potawatomi",
	"cad": "Caddo",
	"gul": "Gulf of Gonave",
	"sem": "Seminole",
	"mis": "Missouri",
	"chk": "Cherokee",
	"sal": "Saint Louis",
	"wil": "Williamsburg",
	"loi": "Long Island",
	"nop": "North Piedmont",
	"pit": "Pittsburgh",
	"ger": "Georgia",
	"kek": "Kekionga",
	"all": "Allegheny",
	"pen": "Pennyrile",
	"ill": "Illiniwek",
	"mic": "Miccosukee",
	"ken": "Kentucky",
	"cho": "Choctaw",
	"lis": "Long Island Sound",
	"hoc": "Ho Chunk",
	"chi": "Chickasaw",
	"fra": "Franklin",
	"chv": "Charlottesville",
	"mai": "Maine",
	"upp": "Upper Pontomac",
	"flo": "Florida Bight",
	"bah": "Bahama Banks",
	"old": "Old Bahama Channel",
	"nic": "Nickajack",
	"wet": "Westsylvania",
	"cus": "Cusseta",
	"mas": "Massachuesetts",
}
