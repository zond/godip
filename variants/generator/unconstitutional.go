package unconstitutional

import (
	"github.com/zond/godip"
	"github.com/zond/godip/graph"
	"github.com/zond/godip/state"
	"github.com/zond/godip/variants/classical"
	"github.com/zond/godip/variants/common"
)

const (
	Pennsylvania  godip.Nation = "Pennsylvania"
	SouthCarolina godip.Nation = "SouthCarolina"
)

var Nations = []godip.Nation{Pennsylvania, SouthCarolina}

var UnconstitutionalVariant = common.Variant{
	Name:              "Unconstitutional",
	Graph:             func() godip.Graph { return UnconstitutionalGraph() },
	Start:             UnconstitutionalStart,
	Blank:             UnconstitutionalBlank,
	Phase:             classical.NewPhase,
	Parser:            classical.Parser,
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
			return classical.Asset("svg/army.svg")
		},
		godip.Fleet: func() ([]byte, error) {
			return classical.Asset("svg/fleet.svg")
		},
	},
	CreatedBy:   "",
	Version:     "",
	Description: "",
	Rules:       "",
}

func UnconstitutionalBlank(phase godip.Phase) *state.State {
	return state.New(UnconstitutionalGraph(), phase, classical.BackupRule, nil, nil)
}

func UnconstitutionalStart() (result *state.State, err error) {
	startPhase := classical.NewPhase(1805, godip.Spring, godip.Movement)
	result = UnconstitutionalBlank(startPhase)
	if err = result.SetUnits(map[godip.Province]godip.Unit{
		"cho": godip.Unit{godip.Army, Pennsylvania},
		"sal": godip.Unit{godip.Army, SouthCarolina},
	}); err != nil {
		return
	}
	result.SetSupplyCenters(map[godip.Province]godip.Nation{
		"cho": Pennsylvania,
		"sal": SouthCarolina,
	})
	return
}

func UnconstitutionalGraph() *graph.Graph {
	return graph.New().
		// Alexandria
		Prov("ale").Conn("she", godip.Land).Conn("chv", godip.Land).Conn("rap", godip.Land).Conn("mar", godip.Land).Conn("upp", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// New York City
		Prov("nyc").Conn("loi", godip.Coast...).Conn("lis", godip.Sea).Conn("con", godip.Coast...).Conn("alb", godip.Land).Conn("cat", godip.Land).Conn("nej", godip.Coast...).Conn("nyb", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Western Reserve
		Prov("wer").Conn("lyc", godip.Land).Conn("cat", godip.Land).Conn("iro", godip.Land).Conn("det", godip.Land).Conn("ohi", godip.Land).Conn("all", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Delaware
		Prov("der").Conn("nej", godip.Coast...).Conn("phi", godip.Land).Conn("yor", godip.Land).Conn("mar", godip.Coast...).Conn("deb", godip.Sea).Flag(godip.Coast...).
		// Philadelphia
		Prov("phi").Conn("der", godip.Land).Conn("nej", godip.Land).Conn("lyc", godip.Land).Conn("har", godip.Land).Conn("yor", godip.Land).Flag(godip.Land).SC(godip.Neutral).
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
		Prov("pro").Conn("pen", godip.Land).Conn("wap", godip.Land).Conn("kek", godip.Land).Conn("pot", godip.Land).Conn("hoc", godip.Land).Conn("ill", godip.Land).Conn("sal", godip.Land).Conn("chi", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Azua
		Prov("azu").Conn("cib", godip.Land).Conn("art", godip.Land).Conn("por", godip.Coast...).Conn("gul", godip.Sea).Conn("sad", godip.Coast...).Flag(godip.Coast...).
		// Tidewater
		Prov("tid").Conn("ric", godip.Land).Conn("noc", godip.Coast...).Conn("out", godip.Sea).Conn("chb", godip.Sea).Conn("wil", godip.Coast...).Flag(godip.Coast...).
		// Gulf of Mexico
		Prov("guc").Conn("gul", godip.Sea).Conn("old", godip.Sea).Conn("flo", godip.Sea).Conn("wef", godip.Sea).Conn("neo", godip.Sea).Flag(godip.Sea).
		// Connecticut
		Prov("con").Conn("alb", godip.Land).Conn("nyc", godip.Coast...).Conn("lis", godip.Sea).Conn("mas", godip.Coast...).Flag(godip.Coast...).
		// Tennessee
		Prov("ten").Conn("chk", godip.Land).Conn("fra", godip.Land).Conn("pen", godip.Land).Conn("chi", godip.Land).Conn("nic", godip.Land).Flag(godip.Land).
		// Alabama
		Prov("ala").Conn("tuk", godip.Land).Conn("nic", godip.Land).Conn("chi", godip.Land).Conn("cho", godip.Land).Conn("cad", godip.Land).Flag(godip.Land).
		// Georgia Bight
		Prov("geb").Conn("eas", godip.Sea).Conn("bah", godip.Sea).Conn("sar", godip.Sea).Conn("out", godip.Sea).Conn("noc", godip.Sea).Conn("chn", godip.Sea).Conn("bea", godip.Sea).Conn("ger", godip.Sea).Flag(godip.Sea).
		// Outer Banks
		Prov("out").Conn("noc", godip.Sea).Conn("geb", godip.Sea).Conn("sar", godip.Sea).Conn("chb", godip.Sea).Conn("tid", godip.Sea).Flag(godip.Sea).
		// Lycoming
		Prov("lyc").Conn("all", godip.Land).Conn("har", godip.Land).Conn("phi", godip.Land).Conn("nej", godip.Land).Conn("cat", godip.Land).Conn("wer", godip.Land).Flag(godip.Land).
		// Sargasso Sea
		Prov("sar").Conn("cib", godip.Sea).Conn("sad", godip.Sea).Conn("atl", godip.Sea).Conn("mai", godip.Sea).Conn("nyb", godip.Sea).Conn("chb", godip.Sea).Conn("out", godip.Sea).Conn("geb", godip.Sea).Conn("bah", godip.Sea).Conn("tur", godip.Sea).Conn("old", godip.Sea).Flag(godip.Sea).
		// New Orleans
		Prov("neo").Conn("guc", godip.Sea).Conn("wef", godip.Coast...).Conn("qua", godip.Land).Flag(godip.Coast...).
		// Harrisburg
		Prov("har").Conn("yor", godip.Land).Conn("phi", godip.Land).Conn("lyc", godip.Land).Conn("all", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Quapow
		Prov("qua").Conn("neo", godip.Land).Conn("wef", godip.Land).Conn("cho", godip.Land).Conn("sal", godip.Land).Conn("osa", godip.Land).Flag(godip.Land).
		// Charleston
		Prov("chn").Conn("mid", godip.Land).Conn("col", godip.Land).Conn("bea", godip.Coast...).Conn("geb", godip.Sea).Conn("noc", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Albany
		Prov("alb").Conn("iro", godip.Land).Conn("cat", godip.Land).Conn("nyc", godip.Land).Conn("con", godip.Land).Conn("mas", godip.Land).Conn("ver", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Vermont
		Prov("ver").Conn("iro", godip.Land).Conn("alb", godip.Land).Conn("mas", godip.Land).Conn("neh", godip.Land).Flag(godip.Land).
		// Windward Passage
		Prov("win").Conn("art", godip.Sea).Conn("old", godip.Sea).Conn("gul", godip.Sea).Conn("por", godip.Sea).Flag(godip.Sea).
		// Delaware Bay
		Prov("deb").Conn("chb", godip.Sea).Conn("nyb", godip.Sea).Conn("nej", godip.Sea).Conn("der", godip.Sea).Conn("mar", godip.Sea).Flag(godip.Sea).
		// Midlands
		Prov("mid").Conn("chn", godip.Land).Conn("noc", godip.Land).Conn("nop", godip.Land).Conn("fra", godip.Land).Conn("col", godip.Land).Flag(godip.Land).
		// West Florida
		Prov("wef").Conn("guc", godip.Sea).Conn("flo", godip.Sea).Conn("cad", godip.Coast...).Conn("cho", godip.Land).Conn("qua", godip.Land).Conn("neo", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Shenandoah
		Prov("she").Conn("wet", godip.Land).Conn("ken", godip.Land).Conn("ric", godip.Land).Conn("chv", godip.Land).Conn("ale", godip.Land).Conn("upp", godip.Land).Flag(godip.Land).
		// Osage
		Prov("osa").Conn("qua", godip.Land).Conn("sal", godip.Land).Conn("mis", godip.Land).Flag(godip.Land).
		// Beaufort
		Prov("bea").Conn("geb", godip.Sea).Conn("chn", godip.Coast...).Conn("col", godip.Land).Conn("ger", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// New Hampshire
		Prov("neh").Conn("gua", godip.Coast...).Conn("ver", godip.Land).Conn("mas", godip.Coast...).Conn("mai", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Atlantic Ocean
		Prov("atl").Conn("mai", godip.Sea).Conn("sar", godip.Sea).Conn("sad", godip.Sea).Conn("gul", godip.Sea).Flag(godip.Sea).
		// New Jersey
		Prov("nej").Conn("nyb", godip.Sea).Conn("nyc", godip.Coast...).Conn("cat", godip.Land).Conn("lyc", godip.Land).Conn("phi", godip.Land).Conn("der", godip.Coast...).Conn("deb", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Wapakoneta
		Prov("wap").Conn("kek", godip.Land).Conn("pro", godip.Land).Conn("pen", godip.Land).Conn("ken", godip.Land).Conn("ohi", godip.Land).Conn("det", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// South Piedmont
		Prov("sou").Conn("ger", godip.Land).Conn("col", godip.Land).Conn("chk", godip.Land).Conn("cus", godip.Land).Conn("sem", godip.Land).Flag(godip.Land).
		// Catskill
		Prov("cat").Conn("nej", godip.Land).Conn("nyc", godip.Land).Conn("alb", godip.Land).Conn("iro", godip.Land).Conn("wer", godip.Land).Conn("lyc", godip.Land).Flag(godip.Land).
		// New York Bight
		Prov("nyb").Conn("nej", godip.Sea).Conn("deb", godip.Sea).Conn("chb", godip.Sea).Conn("sar", godip.Sea).Conn("mai", godip.Sea).Conn("lis", godip.Sea).Conn("loi", godip.Sea).Conn("nyc", godip.Sea).Flag(godip.Sea).
		// Artibonite
		Prov("art").Conn("azu", godip.Land).Conn("cib", godip.Coast...).Conn("old", godip.Sea).Conn("win", godip.Sea).Conn("por", godip.Coast...).Flag(godip.Coast...).
		// Columbia
		Prov("col").Conn("chn", godip.Land).Conn("mid", godip.Land).Conn("fra", godip.Land).Conn("chk", godip.Land).Conn("sou", godip.Land).Conn("ger", godip.Land).Conn("bea", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Chesapeake Bay
		Prov("chb").Conn("deb", godip.Sea).Conn("mar", godip.Sea).Conn("rap", godip.Sea).Conn("wil", godip.Sea).Conn("tid", godip.Sea).Conn("out", godip.Sea).Conn("sar", godip.Sea).Conn("nyb", godip.Sea).Flag(godip.Sea).
		// Ohio
		Prov("ohi").Conn("all", godip.Land).Conn("wer", godip.Land).Conn("det", godip.Land).Conn("wap", godip.Land).Conn("ken", godip.Land).Conn("wet", godip.Land).Flag(godip.Land).
		// York
		Prov("yor").Conn("upp", godip.Land).Conn("mar", godip.Land).Conn("der", godip.Land).Conn("phi", godip.Land).Conn("har", godip.Land).Conn("all", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Iroquois
		Prov("iro").Conn("alb", godip.Land).Conn("ver", godip.Land).Conn("wer", godip.Land).Conn("cat", godip.Land).Flag(godip.Land).
		// Gulf of Maine
		Prov("gua").Conn("neh", godip.Coast...).Conn("mai", godip.Sea).Flag(godip.Coast...).
		// East Florida
		Prov("eas").Conn("geb", godip.Sea).Conn("ger", godip.Coast...).Conn("sem", godip.Coast...).Conn("flo", godip.Sea).Conn("bah", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Tukabatchee
		Prov("tuk").Conn("ala", godip.Land).Conn("cad", godip.Land).Conn("mic", godip.Land).Conn("cus", godip.Land).Conn("nic", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Richmond
		Prov("ric").Conn("tid", godip.Land).Conn("wil", godip.Land).Conn("rap", godip.Land).Conn("chv", godip.Land).Conn("she", godip.Land).Conn("ken", godip.Land).Conn("fra", godip.Land).Conn("nop", godip.Land).Conn("noc", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Saint Domingue
		Prov("sad").Conn("azu", godip.Coast...).Conn("gul", godip.Sea).Conn("atl", godip.Sea).Conn("sar", godip.Sea).Conn("cib", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// North Carolina
		Prov("noc").Conn("out", godip.Sea).Conn("tid", godip.Coast...).Conn("ric", godip.Land).Conn("nop", godip.Land).Conn("mid", godip.Land).Conn("chn", godip.Coast...).Conn("geb", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Potawatomi
		Prov("pot").Conn("kek", godip.Land).Conn("det", godip.Land).Conn("hoc", godip.Land).Conn("pro", godip.Land).Flag(godip.Land).
		// Caddo
		Prov("cad").Conn("cho", godip.Land).Conn("wef", godip.Coast...).Conn("flo", godip.Sea).Conn("mic", godip.Coast...).Conn("tuk", godip.Land).Conn("ala", godip.Land).Flag(godip.Coast...).
		// Gulf of Gonave
		Prov("gul").Conn("atl", godip.Sea).Conn("sad", godip.Sea).Conn("azu", godip.Sea).Conn("por", godip.Sea).Conn("win", godip.Sea).Conn("old", godip.Sea).Conn("guc", godip.Sea).Flag(godip.Sea).
		// Seminole
		Prov("sem").Conn("cus", godip.Land).Conn("mic", godip.Coast...).Conn("flo", godip.Sea).Conn("eas", godip.Coast...).Conn("ger", godip.Land).Conn("sou", godip.Land).Flag(godip.Coast...).
		// Missouri
		Prov("mis").Conn("osa", godip.Land).Conn("sal", godip.Land).Conn("ill", godip.Land).Conn("hoc", godip.Land).Flag(godip.Land).
		// Cherokee
		Prov("chk").Conn("ten", godip.Land).Conn("nic", godip.Land).Conn("cus", godip.Land).Conn("sou", godip.Land).Conn("col", godip.Land).Conn("fra", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Saint Louis
		Prov("sal").Conn("osa", godip.Land).Conn("qua", godip.Land).Conn("cho", godip.Land).Conn("chi", godip.Land).Conn("pro", godip.Land).Conn("ill", godip.Land).Conn("mis", godip.Land).Flag(godip.Land).SC(SouthCarolina).
		// Williamsburg
		Prov("wil").Conn("ric", godip.Land).Conn("tid", godip.Coast...).Conn("chb", godip.Sea).Conn("rap", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Long Island
		Prov("loi").Conn("lis", godip.Sea).Conn("nyc", godip.Coast...).Conn("nyb", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// North Piedmont
		Prov("nop").Conn("mid", godip.Land).Conn("noc", godip.Land).Conn("ric", godip.Land).Conn("fra", godip.Land).Flag(godip.Land).
		// Pittsburgh
		Prov("pit").Conn("all", godip.Land).Conn("all", godip.Land).Conn("wet", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Georgia
		Prov("ger").Conn("sou", godip.Land).Conn("sem", godip.Land).Conn("eas", godip.Coast...).Conn("geb", godip.Sea).Conn("bea", godip.Coast...).Conn("col", godip.Land).Flag(godip.Coast...).
		// Kekionga
		Prov("kek").Conn("pot", godip.Land).Conn("pro", godip.Land).Conn("wap", godip.Land).Conn("det", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Allegheny
		Prov("all").Conn("lyc", godip.Land).Conn("wer", godip.Land).Conn("ohi", godip.Land).Conn("wet", godip.Land).Conn("pit", godip.Land).Conn("pit", godip.Land).Conn("wet", godip.Land).Conn("upp", godip.Land).Conn("yor", godip.Land).Conn("har", godip.Land).Flag(godip.Land).
		// Pennyrile
		Prov("pen").Conn("fra", godip.Land).Conn("ken", godip.Land).Conn("wap", godip.Land).Conn("pro", godip.Land).Conn("chi", godip.Land).Conn("ten", godip.Land).Flag(godip.Land).
		// Illiniwek
		Prov("ill").Conn("pro", godip.Land).Conn("hoc", godip.Land).Conn("mis", godip.Land).Conn("sal", godip.Land).Flag(godip.Land).
		// Miccosukee
		Prov("mic").Conn("tuk", godip.Land).Conn("cad", godip.Coast...).Conn("flo", godip.Sea).Conn("sem", godip.Coast...).Conn("cus", godip.Land).Flag(godip.Coast...).SC(godip.Neutral).
		// Kentucky
		Prov("ken").Conn("she", godip.Land).Conn("wet", godip.Land).Conn("ohi", godip.Land).Conn("wap", godip.Land).Conn("pen", godip.Land).Conn("fra", godip.Land).Conn("ric", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Choctaw
		Prov("cho").Conn("cad", godip.Land).Conn("ala", godip.Land).Conn("chi", godip.Land).Conn("sal", godip.Land).Conn("qua", godip.Land).Conn("wef", godip.Land).Flag(godip.Land).SC(Pennsylvania).
		// Long Island Sound
		Prov("lis").Conn("loi", godip.Sea).Conn("nyb", godip.Sea).Conn("mai", godip.Sea).Conn("mas", godip.Sea).Conn("con", godip.Sea).Conn("nyc", godip.Sea).Flag(godip.Sea).
		// Ho Chunk
		Prov("hoc").Conn("mis", godip.Land).Conn("ill", godip.Land).Conn("pro", godip.Land).Conn("pot", godip.Land).Flag(godip.Land).
		// Chickasaw
		Prov("chi").Conn("ten", godip.Land).Conn("pen", godip.Land).Conn("pro", godip.Land).Conn("sal", godip.Land).Conn("cho", godip.Land).Conn("ala", godip.Land).Conn("nic", godip.Land).Flag(godip.Land).
		// Franklin
		Prov("fra").Conn("pen", godip.Land).Conn("ten", godip.Land).Conn("chk", godip.Land).Conn("col", godip.Land).Conn("mid", godip.Land).Conn("nop", godip.Land).Conn("ric", godip.Land).Conn("ken", godip.Land).Flag(godip.Land).
		// Charlottesville
		Prov("chv").Conn("she", godip.Land).Conn("ric", godip.Land).Conn("rap", godip.Land).Conn("ale", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Maine
		Prov("mai").Conn("gua", godip.Sea).Conn("neh", godip.Sea).Conn("mas", godip.Sea).Conn("lis", godip.Sea).Conn("nyb", godip.Sea).Conn("sar", godip.Sea).Conn("atl", godip.Sea).Flag(godip.Sea).
		// Upper Pontomac
		Prov("upp").Conn("yor", godip.Land).Conn("all", godip.Land).Conn("wet", godip.Land).Conn("she", godip.Land).Conn("ale", godip.Land).Conn("mar", godip.Land).Flag(godip.Land).
		// Florida Bight
		Prov("flo").Conn("bah", godip.Sea).Conn("eas", godip.Sea).Conn("sem", godip.Sea).Conn("mic", godip.Sea).Conn("cad", godip.Sea).Conn("wef", godip.Sea).Conn("guc", godip.Sea).Conn("old", godip.Sea).Flag(godip.Sea).
		// Bahama Banks
		Prov("bah").Conn("flo", godip.Sea).Conn("old", godip.Sea).Conn("tur", godip.Sea).Conn("sar", godip.Sea).Conn("geb", godip.Sea).Conn("eas", godip.Sea).Flag(godip.Sea).
		// Old Bahama Channel
		Prov("old").Conn("tur", godip.Sea).Conn("bah", godip.Sea).Conn("flo", godip.Sea).Conn("guc", godip.Sea).Conn("gul", godip.Sea).Conn("win", godip.Sea).Conn("art", godip.Sea).Conn("cib", godip.Sea).Conn("sar", godip.Sea).Flag(godip.Sea).
		// Nickajack
		Prov("nic").Conn("cus", godip.Land).Conn("chk", godip.Land).Conn("ten", godip.Land).Conn("chi", godip.Land).Conn("ala", godip.Land).Conn("tuk", godip.Land).Flag(godip.Land).
		// Westsylvania
		Prov("wet").Conn("all", godip.Land).Conn("ohi", godip.Land).Conn("ken", godip.Land).Conn("she", godip.Land).Conn("upp", godip.Land).Conn("all", godip.Land).Conn("pit", godip.Land).Flag(godip.Land).
		// Cusseta
		Prov("cus").Conn("nic", godip.Land).Conn("tuk", godip.Land).Conn("mic", godip.Land).Conn("sem", godip.Land).Conn("sou", godip.Land).Conn("chk", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Massachuesetts
		Prov("mas").Conn("lis", godip.Sea).Conn("mai", godip.Sea).Conn("neh", godip.Coast...).Conn("ver", godip.Land).Conn("alb", godip.Land).Conn("con", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
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
