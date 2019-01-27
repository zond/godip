package twentytwenty

import (
	"github.com/zond/godip"
	"github.com/zond/godip/graph"
	"github.com/zond/godip/orders"
	"github.com/zond/godip/state"
	"github.com/zond/godip/variants/classical"
	"github.com/zond/godip/variants/common"
)

const (
	Brazil      godip.Nation = "Brazil"
	Canada      godip.Nation = "Canada"
	Australia   godip.Nation = "Australia"
	Italy       godip.Nation = "Italy"
	USA         godip.Nation = "USA"
	Kenya       godip.Nation = "Kenya"
	Egypt       godip.Nation = "Egypt"
	Thailand    godip.Nation = "Thailand"
	Turkey      godip.Nation = "Turkey"
	SouthAfrica godip.Nation = "South Africa"
	India       godip.Nation = "India"
	Russia      godip.Nation = "Russia"
	Pakistan    godip.Nation = "Pakistan"
	China       godip.Nation = "China"
	UK          godip.Nation = "UK"
	Japan       godip.Nation = "Japan"
	Germany     godip.Nation = "Germany"
	Argentina   godip.Nation = "Argentina"
	Nigeria     godip.Nation = "Nigeria"
	Spain       godip.Nation = "Spain"
)

var Nations = []godip.Nation{Brazil, Canada, Australia, Italy, USA, Kenya, Egypt, Thailand, Turkey, SouthAfrica, India, Russia, Pakistan, China, UK, Japan, Germany, Argentina, Nigeria, Spain}

// A function that declares a solo winner if a nation has over 49 SCs, or they lead by at least max(2020-year, 1) SCs.
func TwentyTwentyWinner(s *state.State) godip.Nation {
	// Create a map from nation to count of owned SCs.
	scCount := map[godip.Nation]int{}
	for _, nat := range s.SupplyCenters() {
		if nat != "" {
			scCount[nat] = scCount[nat] + 1
		}
	}
	// Find the leading two players.
	highestSCCount := 0
	secondHighestSCCount := 0
	var leader godip.Nation
	for nat, count := range scCount {
		if count > highestSCCount {
			leader = nat
			secondHighestSCCount = highestSCCount
			highestSCCount = count
		} else if count == highestSCCount {
			leader = ""
			secondHighestSCCount = highestSCCount
		} else if count > secondHighestSCCount {
			secondHighestSCCount = count
		}
	}
	// If there's a leader then check if we have a winner.
	if leader != "" {
		// Return the nation if they have more than half the map.
		if highestSCCount >= 49 {
			return leader
		}
		// Return the nation if they have a big enough lead.
		if highestSCCount-secondHighestSCCount > 2020-s.Phase().Year() {
			return leader
		}
	}
	return ""
}

var BuildAnyHomeCenterParser = orders.NewParser([]godip.Order{
	orders.MoveOrder,
	orders.MoveViaConvoyOrder,
	orders.HoldOrder,
	orders.SupportOrder,
	orders.BuildAnyHomeCenterOrder,
	orders.DisbandOrder,
	orders.ConvoyOrder,
})

var TwentyTwentyVariant = common.Variant{
	Name:       "Twenty Twenty",
	Graph:      func() godip.Graph { return TwentyTwentyGraph() },
	Start:      TwentyTwentyStart,
	Blank:      TwentyTwentyBlank,
	Phase:      classical.NewPhase,
	Parser:     BuildAnyHomeCenterParser,
	Nations:    Nations,
	PhaseTypes: classical.PhaseTypes,
	Seasons:    classical.Seasons,
	UnitTypes:  classical.UnitTypes,
	SoloWinner: TwentyTwentyWinner,
	SVGMap: func() ([]byte, error) {
		return Asset("svg/twentytwentymap.svg")
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
	CreatedBy:   "TTTPPP",
	Version:     "1",
	Description: "Twenty nations complete to conquer the world by the year 2020.",
	Rules: "The rules are mostly standard. Nations may build in any captured " +
		"home centers (note - they may not build in captured neutral supply " +
		"centers). To win a nation needs to own more supply centers than any " +
		"opponent. In the first year they need 20 more supply centers than any " +
		"other player, but this target is reduced by 1 each year. So to win in " +
		"the year 2015 a player needs at least 6 more supply centers than any " +
		"other player, and in 2020 and beyond they need a lead of a single " +
		"supply center. Alternatively, if a player manages to get to 49 centers " +
		"(i.e. they own over half the map) then they automatically win. There " +
		"are six bridges connecting regions for armies (and fleets). These are " +
		"Anchorage-Vladivostok, New Orleans-Cuba, Cuba-Dominican Republic, " +
		"Ethiopia-Yemen, Korea-Nagisaki and Indonesia-Darwin. Thirteen regions " +
		"have multiple coasts. These are Whitehorse, Los Angeles, Mexico, " +
		"Colombia, Bordeaux, Milan, Rome, Finland, Bulgaria, Ankara, Iraq, " +
		"Mecca and Shanyang.",
}

func TwentyTwentyBlank(phase godip.Phase) *state.State {
	return state.New(TwentyTwentyGraph(), phase, classical.BackupRule, nil)
}

func TwentyTwentyStart() (result *state.State, err error) {
	startPhase := classical.NewPhase(2001, godip.Spring, godip.Movement)
	result = TwentyTwentyBlank(startPhase)
	if err = result.SetUnits(map[godip.Province]godip.Unit{
		"rec": godip.Unit{godip.Fleet, Brazil},
		"rio": godip.Unit{godip.Fleet, Brazil},
		"mac": godip.Unit{godip.Army, Brazil},
		"van": godip.Unit{godip.Fleet, Canada},
		"iqa": godip.Unit{godip.Fleet, Canada},
		"mot": godip.Unit{godip.Fleet, Canada},
		"pet": godip.Unit{godip.Fleet, Australia},
		"bri": godip.Unit{godip.Fleet, Australia},
		"syd": godip.Unit{godip.Fleet, Australia},
		"roe": godip.Unit{godip.Fleet, Italy},
		"mil": godip.Unit{godip.Army, Italy},
		"nap": godip.Unit{godip.Army, Italy},
		"was": godip.Unit{godip.Fleet, USA},
		"atl": godip.Unit{godip.Army, USA},
		"okl": godip.Unit{godip.Army, USA},
		"anc": godip.Unit{godip.Army, USA},
		"mom": godip.Unit{godip.Fleet, Kenya},
		"nai": godip.Unit{godip.Army, Kenya},
		"mar": godip.Unit{godip.Army, Kenya},
		"ale": godip.Unit{godip.Fleet, Egypt},
		"cai": godip.Unit{godip.Army, Egypt},
		"asw": godip.Unit{godip.Army, Egypt},
		"chm": godip.Unit{godip.Army, Thailand},
		"bnk": godip.Unit{godip.Fleet, Thailand},
		"hat": godip.Unit{godip.Fleet, Thailand},
		"ist": godip.Unit{godip.Fleet, Turkey},
		"ank": godip.Unit{godip.Army, Turkey},
		"diy": godip.Unit{godip.Army, Turkey},
		"cap": godip.Unit{godip.Fleet, SouthAfrica},
		"dur": godip.Unit{godip.Fleet, SouthAfrica},
		"pre": godip.Unit{godip.Army, SouthAfrica},
		"mum": godip.Unit{godip.Fleet, India},
		"bna": godip.Unit{godip.Fleet, India},
		"ned": godip.Unit{godip.Army, India},
		"mos": godip.Unit{godip.Army, Russia},
		"oms": godip.Unit{godip.Army, Russia},
		"irk": godip.Unit{godip.Army, Russia},
		"vla": godip.Unit{godip.Army, Russia},
		"kar": godip.Unit{godip.Fleet, Pakistan},
		"isl": godip.Unit{godip.Army, Pakistan},
		"lah": godip.Unit{godip.Army, Pakistan},
		"yum": godip.Unit{godip.Army, China},
		"bao": godip.Unit{godip.Army, China},
		"cho": godip.Unit{godip.Army, China},
		"bei": godip.Unit{godip.Army, China},
		"lon": godip.Unit{godip.Fleet, UK},
		"dub": godip.Unit{godip.Fleet, UK},
		"edi": godip.Unit{godip.Fleet, UK},
		"sap": godip.Unit{godip.Fleet, Japan},
		"tok": godip.Unit{godip.Fleet, Japan},
		"nag": godip.Unit{godip.Fleet, Japan},
		"ham": godip.Unit{godip.Fleet, Germany},
		"mun": godip.Unit{godip.Army, Germany},
		"ben": godip.Unit{godip.Army, Germany},
		"men": godip.Unit{godip.Fleet, Argentina},
		"bue": godip.Unit{godip.Fleet, Argentina},
		"com": godip.Unit{godip.Fleet, Argentina},
		"lag": godip.Unit{godip.Fleet, Nigeria},
		"abu": godip.Unit{godip.Army, Nigeria},
		"kan": godip.Unit{godip.Army, Nigeria},
		"mad": godip.Unit{godip.Fleet, Spain},
		"cad": godip.Unit{godip.Fleet, Spain},
		"bar": godip.Unit{godip.Army, Spain},
	}); err != nil {
		return
	}
	result.SetSupplyCenters(map[godip.Province]godip.Nation{
		"rec": Brazil,
		"rio": Brazil,
		"mac": Brazil,
		"van": Canada,
		"iqa": Canada,
		"mot": Canada,
		"pet": Australia,
		"bri": Australia,
		"syd": Australia,
		"roe": Italy,
		"mil": Italy,
		"nap": Italy,
		"was": USA,
		"atl": USA,
		"okl": USA,
		"anc": USA,
		"mom": Kenya,
		"nai": Kenya,
		"mar": Kenya,
		"ale": Egypt,
		"cai": Egypt,
		"asw": Egypt,
		"chm": Thailand,
		"bnk": Thailand,
		"hat": Thailand,
		"ist": Turkey,
		"ank": Turkey,
		"diy": Turkey,
		"cap": SouthAfrica,
		"dur": SouthAfrica,
		"pre": SouthAfrica,
		"mum": India,
		"bna": India,
		"ned": India,
		"mos": Russia,
		"oms": Russia,
		"irk": Russia,
		"vla": Russia,
		"kar": Pakistan,
		"isl": Pakistan,
		"lah": Pakistan,
		"yum": China,
		"bao": China,
		"cho": China,
		"bei": China,
		"lon": UK,
		"dub": UK,
		"edi": UK,
		"sap": Japan,
		"tok": Japan,
		"nag": Japan,
		"ham": Germany,
		"mun": Germany,
		"ben": Germany,
		"men": Argentina,
		"bue": Argentina,
		"com": Argentina,
		"lag": Nigeria,
		"abu": Nigeria,
		"kan": Nigeria,
		"mad": Spain,
		"cad": Spain,
		"bar": Spain,
	})
	return
}

func TwentyTwentyGraph() *graph.Graph {
	return graph.New().
		// East Mediterranean Sea
		Prov("ems").Conn("bnh", godip.Sea).Conn("ale", godip.Sea).Conn("cai", godip.Sea).Conn("mec", godip.Sea).Conn("mec/nc", godip.Sea).Conn("ira", godip.Sea).Conn("ira/wc", godip.Sea).Conn("aeg", godip.Sea).Conn("grc", godip.Sea).Conn("ion", godip.Sea).Flag(godip.Sea).
		// Scotia Sea
		Prov("sco").Conn("spo", godip.Sea).Conn("soo", godip.Sea).Conn("cap", godip.Sea).Conn("esa", godip.Sea).Conn("wsa", godip.Sea).Flag(godip.Sea).
		// Ghana
		Prov("gha").Conn("abu", godip.Land).Conn("nig", godip.Land).Conn("cot", godip.Coast...).Conn("gog", godip.Sea).Conn("lag", godip.Coast...).Flag(godip.Coast...).
		// Rostov-on-Don
		Prov("ros").Conn("mos", godip.Land).Conn("ukr", godip.Coast...).Conn("bla", godip.Sea).Conn("aze", godip.Coast...).Conn("ast", godip.Land).Conn("oms", godip.Land).Flag(godip.Coast...).
		// Albania
		Prov("alb").Conn("ser", godip.Land).Conn("cro", godip.Coast...).Conn("ion", godip.Sea).Conn("grc", godip.Coast...).Flag(godip.Coast...).
		// Bordeau
		Prov("bod").Conn("bar", godip.Land).Conn("lyo", godip.Land).Conn("pai", godip.Land).Conn("mad", godip.Land).Flag(godip.Land).
		// Bordeau (West Coast)
		Prov("bod/wc").Conn("pai", godip.Sea).Conn("ena", godip.Sea).Conn("mad", godip.Sea).Flag(godip.Sea).
		// Bordeau (East Coast)
		Prov("bod/ec").Conn("bar", godip.Sea).Conn("ble", godip.Sea).Conn("lyo", godip.Sea).Flag(godip.Sea).
		// Gulf of Mannar
		Prov("gur").Conn("mio", godip.Sea).Conn("eio", godip.Sea).Conn("bob", godip.Sea).Conn("kol", godip.Sea).Conn("hyd", godip.Sea).Conn("bna", godip.Sea).Flag(godip.Sea).
		// Chongqing
		Prov("cho").Conn("lan", godip.Land).Conn("lij", godip.Land).Conn("kum", godip.Land).Conn("hon", godip.Land).Conn("bei", godip.Land).Conn("bao", godip.Land).Flag(godip.Land).SC(China).
		// Cuba
		Prov("cub").Conn("bet", godip.Sea).Conn("gux", godip.Sea).Conn("cas", godip.Sea).Conn("dom", godip.Coast...).Conn("neo", godip.Coast...).Flag(godip.Coast...).
		// Cape Town
		Prov("cap").Conn("esa", godip.Sea).Conn("sco", godip.Sea).Conn("soo", godip.Sea).Conn("por", godip.Coast...).Conn("zim", godip.Coast...).Flag(godip.Coast...).SC(SouthAfrica).
		// Uzbekistan
		Prov("uzb").Conn("tur", godip.Land).Conn("afg", godip.Land).Conn("alm", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Mongolia
		Prov("mog").Conn("yak", godip.Land).Conn("irk", godip.Land).Conn("ast", godip.Land).Conn("alm", godip.Land).Conn("uru", godip.Land).Conn("yum", godip.Land).Conn("bao", godip.Land).Conn("bei", godip.Land).Conn("she", godip.Land).Flag(godip.Land).
		// Kashmir
		Prov("kam").Conn("mum", godip.Coast...).Conn("ned", godip.Land).Conn("nep", godip.Land).Conn("tib", godip.Land).Conn("isl", godip.Land).Conn("lah", godip.Land).Conn("kar", godip.Coast...).Conn("ara", godip.Sea).Flag(godip.Coast...).
		// Hong Kong
		Prov("hon").Conn("kum", godip.Coast...).Conn("yes", godip.Sea).Conn("bei", godip.Coast...).Conn("cho", godip.Land).Flag(godip.Coast...).
		// Antarctica
		Prov("ant").Conn("spo", godip.Sea).Conn("tas", godip.Sea).Conn("mio", godip.Sea).Conn("moc", godip.Sea).Conn("soo", godip.Sea).Conn("spo", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Red Sea
		Prov("red").Conn("yem", godip.Sea).Conn("mec", godip.Sea).Conn("mec/sc", godip.Sea).Conn("cai", godip.Sea).Conn("asw", godip.Sea).Conn("sud", godip.Sea).Conn("eth", godip.Sea).Conn("wio", godip.Sea).Conn("lac", godip.Sea).Conn("ara", godip.Sea).Flag(godip.Sea).
		// Philippines
		Prov("phi").Conn("ecs", godip.Sea).Conn("scs", godip.Sea).Conn("npo", godip.Sea).Conn("arc", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Ottawa
		Prov("ott").Conn("lab", godip.Sea).Conn("iqa", godip.Coast...).Conn("yeo", godip.Land).Conn("chg", godip.Land).Conn("ney", godip.Land).Conn("mot", godip.Coast...).Flag(godip.Coast...).
		// London
		Prov("lon").Conn("ena", godip.Sea).Conn("pai", godip.Coast...).Conn("nos", godip.Sea).Conn("edi", godip.Coast...).Conn("dub", godip.Coast...).Flag(godip.Coast...).SC(UK).
		// Mombasa
		Prov("mom").Conn("eth", godip.Coast...).Conn("mar", godip.Land).Conn("nai", godip.Land).Conn("tan", godip.Coast...).Conn("moc", godip.Sea).Conn("wio", godip.Sea).Flag(godip.Coast...).SC(Kenya).
		// Yemen
		Prov("yem").Conn("red", godip.Sea).Conn("ara", godip.Sea).Conn("oma", godip.Coast...).Conn("riy", godip.Land).Conn("mec", godip.Land).Conn("mec/sc", godip.Sea).Conn("eth", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Afghanistan
		Prov("afg").Conn("uru", godip.Land).Conn("alm", godip.Land).Conn("uzb", godip.Land).Conn("tur", godip.Land).Conn("mas", godip.Land).Conn("kar", godip.Land).Conn("lah", godip.Land).Conn("isl", godip.Land).Conn("ksi", godip.Land).Flag(godip.Land).
		// South China Sea
		Prov("scs").Conn("npo", godip.Sea).Conn("phi", godip.Sea).Conn("ecs", godip.Sea).Conn("tai", godip.Sea).Conn("yes", godip.Sea).Conn("vie", godip.Sea).Conn("bnk", godip.Sea).Conn("chm", godip.Sea).Conn("npo", godip.Sea).Flag(godip.Sea).
		// South Mid Atlantic
		Prov("sma").Conn("wsa", godip.Sea).Conn("gin", godip.Sea).Conn("sen", godip.Sea).Conn("nma", godip.Sea).Conn("pag", godip.Sea).Conn("bue", godip.Sea).Flag(godip.Sea).
		// Hydrabad
		Prov("hyd").Conn("kol", godip.Coast...).Conn("ned", godip.Land).Conn("mum", godip.Land).Conn("bna", godip.Coast...).Conn("gur", godip.Sea).Flag(godip.Coast...).
		// Sao Paulo
		Prov("sao").Conn("rio", godip.Coast...).Conn("pag", godip.Coast...).Conn("nma", godip.Sea).Flag(godip.Coast...).
		// Lahore
		Prov("lah").Conn("isl", godip.Land).Conn("afg", godip.Land).Conn("kar", godip.Land).Conn("kam", godip.Land).Flag(godip.Land).SC(Pakistan).
		// Mandalay
		Prov("man").Conn("bad", godip.Coast...).Conn("bob", godip.Sea).Conn("and", godip.Sea).Conn("ran", godip.Coast...).Conn("kum", godip.Land).Conn("lij", godip.Land).Conn("bhu", godip.Land).Flag(godip.Coast...).SC(godip.Neutral).
		// New York
		Prov("ney").Conn("mot", godip.Coast...).Conn("ott", godip.Land).Conn("chg", godip.Land).Conn("was", godip.Coast...).Conn("che", godip.Sea).Conn("gos", godip.Sea).Flag(godip.Coast...).
		// Rome
		Prov("roe").Conn("nap", godip.Land).Conn("mil", godip.Land).Flag(godip.Land).SC(Italy).
		// Rome (West Coast)
		Prov("roe/wc").Conn("nap", godip.Sea).Conn("mil/wc", godip.Sea).Conn("ble", godip.Sea).Flag(godip.Sea).
		// Rome (East Coast)
		Prov("roe/ec").Conn("nap", godip.Sea).Conn("ion", godip.Sea).Conn("mil/ec", godip.Sea).Flag(godip.Sea).
		// Western North Atlantic
		Prov("wna").Conn("des", godip.Sea).Conn("lab", godip.Sea).Conn("mot", godip.Sea).Conn("gos", godip.Sea).Conn("che", godip.Sea).Conn("bet", godip.Sea).Conn("dom", godip.Sea).Conn("ena", godip.Sea).Flag(godip.Sea).
		// Atlanta
		Prov("atl").Conn("neo", godip.Land).Conn("was", godip.Land).Conn("chg", godip.Land).Conn("okl", godip.Land).Conn("neo", godip.Land).Flag(godip.Land).SC(USA).
		// Macapa
		Prov("mac").Conn("ven", godip.Coast...).Conn("bra", godip.Land).Conn("rec", godip.Coast...).Conn("gub", godip.Sea).Flag(godip.Coast...).SC(Brazil).
		// Almaty
		Prov("alm").Conn("uru", godip.Land).Conn("mog", godip.Land).Conn("ast", godip.Land).Conn("aze", godip.Land).Conn("tur", godip.Land).Conn("uzb", godip.Land).Conn("afg", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Alexandria
		Prov("ale").Conn("sud", godip.Land).Conn("asw", godip.Land).Conn("cai", godip.Coast...).Conn("ems", godip.Sea).Conn("bnh", godip.Coast...).Flag(godip.Coast...).SC(Egypt).
		// Paris
		Prov("pai").Conn("nos", godip.Sea).Conn("lon", godip.Coast...).Conn("ena", godip.Sea).Conn("bod", godip.Land).Conn("bod/wc", godip.Sea).Conn("lyo", godip.Land).Conn("beg", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Korea
		Prov("kor").Conn("ecs", godip.Sea).Conn("soj", godip.Sea).Conn("she", godip.Land).Conn("she/sc", godip.Sea).Conn("she/ec", godip.Sea).Conn("yes", godip.Sea).Conn("nag", godip.Coast...).Flag(godip.Coast...).
		// Ionian Sea
		Prov("ion").Conn("nap", godip.Sea).Conn("tri", godip.Sea).Conn("bnh", godip.Sea).Conn("ems", godip.Sea).Conn("grc", godip.Sea).Conn("alb", godip.Sea).Conn("cro", godip.Sea).Conn("mil", godip.Sea).Conn("mil/ec", godip.Sea).Conn("roe", godip.Sea).Conn("roe/ec", godip.Sea).Flag(godip.Sea).
		// Anchorage
		Prov("anc").Conn("arc", godip.Sea).Conn("arc", godip.Sea).Conn("whi", godip.Land).Conn("whi/nc", godip.Sea).Conn("whi/wc", godip.Sea).Conn("vla", godip.Coast...).Flag(godip.Coast...).SC(USA).
		// Adelaide
		Prov("ade").Conn("syd", godip.Coast...).Conn("bri", godip.Land).Conn("bri", godip.Land).Conn("dar", godip.Land).Conn("pet", godip.Coast...).Conn("tas", godip.Sea).Flag(godip.Coast...).
		// Chicago
		Prov("chg").Conn("atl", godip.Land).Conn("was", godip.Land).Conn("ney", godip.Land).Conn("ott", godip.Land).Conn("min", godip.Land).Conn("los", godip.Land).Conn("okl", godip.Land).Flag(godip.Land).
		// Ankara
		Prov("ank").Conn("ira", godip.Land).Conn("diy", godip.Land).Conn("ist", godip.Land).Flag(godip.Land).SC(Turkey).
		// Ankara (North Coast)
		Prov("ank/nc").Conn("diy", godip.Sea).Conn("bla", godip.Sea).Conn("ist", godip.Sea).Flag(godip.Sea).
		// Ankara (South Coast)
		Prov("ank/sc").Conn("aeg", godip.Sea).Conn("ira/wc", godip.Sea).Conn("ist", godip.Sea).Flag(godip.Sea).
		// Chad
		Prov("cha").Conn("bnh", godip.Land).Conn("tri", godip.Land).Conn("nig", godip.Land).Conn("gab", godip.Land).Conn("car", godip.Land).Conn("sud", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Tibet
		Prov("tib").Conn("nep", godip.Land).Conn("bhu", godip.Land).Conn("lij", godip.Land).Conn("lan", godip.Land).Conn("ksi", godip.Land).Conn("isl", godip.Land).Conn("kam", godip.Land).Flag(godip.Land).
		// Lagos
		Prov("lag").Conn("kan", godip.Land).Conn("abu", godip.Land).Conn("gha", godip.Coast...).Conn("gog", godip.Sea).Conn("gab", godip.Coast...).Flag(godip.Coast...).SC(Nigeria).
		// Recife
		Prov("rec").Conn("nma", godip.Sea).Conn("cab", godip.Sea).Conn("gub", godip.Sea).Conn("mac", godip.Coast...).Conn("bra", godip.Coast...).Flag(godip.Coast...).SC(Brazil).
		// Mauritania
		Prov("mau").Conn("sen", godip.Coast...).Conn("alg", godip.Land).Conn("mor", godip.Coast...).Conn("can", godip.Sea).Conn("cab", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Tanzania
		Prov("tan").Conn("moq", godip.Coast...).Conn("moc", godip.Sea).Conn("mom", godip.Coast...).Conn("nai", godip.Land).Conn("mar", godip.Land).Conn("eth", godip.Land).Conn("drc", godip.Land).Conn("zam", godip.Land).Flag(godip.Coast...).
		// North Pacific Ocean
		Prov("npo").Conn("arc", godip.Sea).Conn("mpo", godip.Sea).Conn("mex", godip.Sea).Conn("mex/wc", godip.Sea).Conn("los", godip.Sea).Conn("los/wc", godip.Sea).Conn("hat", godip.Sea).Conn("ind", godip.Sea).Conn("bob", godip.Sea).Conn("mpo", godip.Sea).Conn("arc", godip.Sea).Conn("phi", godip.Sea).Conn("scs", godip.Sea).Conn("scs", godip.Sea).Conn("chm", godip.Sea).Flag(godip.Sea).
		// St. Petersburg
		Prov("stp").Conn("lap", godip.Sea).Conn("fin", godip.Land).Conn("fin/nc", godip.Sea).Conn("bea", godip.Land).Conn("mos", godip.Land).Conn("oms", godip.Coast...).Flag(godip.Coast...).
		// Minneapolis
		Prov("min").Conn("van", godip.Coast...).Conn("arc", godip.Sea).Conn("los", godip.Land).Conn("los/wc", godip.Sea).Conn("chg", godip.Land).Flag(godip.Coast...).
		// Labrador Sea
		Prov("lab").Conn("des", godip.Sea).Conn("grd", godip.Sea).Conn("iqa", godip.Sea).Conn("ott", godip.Sea).Conn("mot", godip.Sea).Conn("wna", godip.Sea).Flag(godip.Sea).
		// Bhutan
		Prov("bhu").Conn("nep", godip.Land).Conn("bad", godip.Land).Conn("man", godip.Land).Conn("lij", godip.Land).Conn("tib", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Laccadive Sea
		Prov("lac").Conn("ara", godip.Sea).Conn("red", godip.Sea).Conn("wio", godip.Sea).Conn("mio", godip.Sea).Conn("bna", godip.Sea).Conn("mum", godip.Sea).Flag(godip.Sea).
		// Tokyo
		Prov("tok").Conn("nag", godip.Coast...).Conn("ecs", godip.Sea).Conn("bes", godip.Sea).Conn("sap", godip.Coast...).Conn("soj", godip.Sea).Flag(godip.Coast...).SC(Japan).
		// Gulf of St. Lawrence
		Prov("gos").Conn("ney", godip.Sea).Conn("che", godip.Sea).Conn("wna", godip.Sea).Conn("mot", godip.Sea).Flag(godip.Sea).
		// Azerbaijan
		Prov("aze").Conn("teh", godip.Land).Conn("tur", godip.Land).Conn("alm", godip.Land).Conn("ast", godip.Land).Conn("ros", godip.Coast...).Conn("bla", godip.Sea).Conn("diy", godip.Coast...).Flag(godip.Coast...).
		// Gabon
		Prov("gab").Conn("kan", godip.Land).Conn("lag", godip.Coast...).Conn("gog", godip.Sea).Conn("ang", godip.Coast...).Conn("drc", godip.Land).Conn("car", godip.Land).Conn("cha", godip.Land).Conn("nig", godip.Land).Flag(godip.Coast...).
		// Algeria
		Prov("alg").Conn("nig", godip.Land).Conn("tri", godip.Coast...).Conn("ble", godip.Sea).Conn("mor", godip.Coast...).Conn("mau", godip.Land).Flag(godip.Coast...).SC(godip.Neutral).
		// Taiwan
		Prov("tai").Conn("scs", godip.Sea).Conn("ecs", godip.Sea).Conn("yes", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Beijing
		Prov("bei").Conn("bao", godip.Land).Conn("cho", godip.Land).Conn("hon", godip.Coast...).Conn("yes", godip.Sea).Conn("she", godip.Land).Conn("she/sc", godip.Sea).Conn("mog", godip.Land).Flag(godip.Coast...).SC(China).
		// Hat Yai
		Prov("hat").Conn("npo", godip.Sea).Conn("chm", godip.Coast...).Conn("ran", godip.Coast...).Conn("and", godip.Sea).Conn("ind", godip.Coast...).Flag(godip.Coast...).SC(Thailand).
		// Mashad
		Prov("mas").Conn("kar", godip.Coast...).Conn("afg", godip.Land).Conn("tur", godip.Land).Conn("teh", godip.Coast...).Conn("peg", godip.Sea).Flag(godip.Coast...).
		// Mozambique Channel
		Prov("moc").Conn("tan", godip.Sea).Conn("moq", godip.Sea).Conn("pre", godip.Sea).Conn("dur", godip.Sea).Conn("soo", godip.Sea).Conn("ant", godip.Sea).Conn("mio", godip.Sea).Conn("wio", godip.Sea).Conn("mom", godip.Sea).Flag(godip.Sea).
		// Kashi
		Prov("ksi").Conn("lan", godip.Land).Conn("uru", godip.Land).Conn("afg", godip.Land).Conn("isl", godip.Land).Conn("tib", godip.Land).Flag(godip.Land).
		// Greenland
		Prov("grd").Conn("arc", godip.Sea).Conn("iqa", godip.Coast...).Conn("lab", godip.Sea).Conn("des", godip.Sea).Conn("nos", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// New Orleans
		Prov("neo").Conn("atl", godip.Land).Conn("atl", godip.Land).Conn("okl", godip.Land).Conn("los", godip.Land).Conn("los/ec", godip.Sea).Conn("gux", godip.Sea).Conn("bet", godip.Sea).Conn("was", godip.Coast...).Conn("cub", godip.Coast...).Flag(godip.Coast...).
		// Dublin
		Prov("dub").Conn("ena", godip.Sea).Conn("lon", godip.Coast...).Conn("edi", godip.Coast...).Conn("des", godip.Sea).Flag(godip.Coast...).SC(UK).
		// Baltic Sea
		Prov("bat").Conn("ben", godip.Sea).Conn("pol", godip.Sea).Conn("bea", godip.Sea).Conn("fin", godip.Sea).Conn("fin/sc", godip.Sea).Conn("swe", godip.Sea).Conn("dem", godip.Sea).Flag(godip.Sea).
		// Guiana Basin
		Prov("gub").Conn("ven", godip.Sea).Conn("mac", godip.Sea).Conn("rec", godip.Sea).Conn("cab", godip.Sea).Conn("ena", godip.Sea).Conn("dom", godip.Sea).Conn("cas", godip.Sea).Flag(godip.Sea).
		// Norway
		Prov("now").Conn("swe", godip.Coast...).Conn("fin", godip.Land).Conn("fin/nc", godip.Sea).Conn("bes", godip.Sea).Conn("arc", godip.Sea).Conn("nos", godip.Sea).Conn("ska", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Eastern South Atlantic
		Prov("esa").Conn("cot", godip.Sea).Conn("gin", godip.Sea).Conn("wsa", godip.Sea).Conn("sco", godip.Sea).Conn("cap", godip.Sea).Conn("zim", godip.Sea).Conn("ang", godip.Sea).Conn("gog", godip.Sea).Flag(godip.Sea).
		// Vladivostok
		Prov("vla").Conn("lap", godip.Sea).Conn("yak", godip.Coast...).Conn("she", godip.Land).Conn("she/ec", godip.Sea).Conn("soj", godip.Sea).Conn("anc", godip.Coast...).Flag(godip.Coast...).SC(Russia).
		// Bulgaria
		Prov("bul").Conn("ist", godip.Land).Conn("rma", godip.Land).Conn("ser", godip.Land).Conn("grc", godip.Land).Flag(godip.Land).
		// Bulgaria (South Coast)
		Prov("bul/sc").Conn("ist", godip.Sea).Conn("grc", godip.Sea).Conn("aeg", godip.Sea).Flag(godip.Sea).
		// Bulgaria (East Coast)
		Prov("bul/ec").Conn("ist", godip.Sea).Conn("bla", godip.Sea).Conn("rma", godip.Sea).Flag(godip.Sea).
		// Kumming
		Prov("kum").Conn("hon", godip.Coast...).Conn("cho", godip.Land).Conn("lij", godip.Land).Conn("man", godip.Land).Conn("ran", godip.Land).Conn("vie", godip.Coast...).Conn("yes", godip.Sea).Flag(godip.Coast...).
		// Caribbean Sea
		Prov("cas").Conn("gux", godip.Sea).Conn("pam", godip.Sea).Conn("col", godip.Sea).Conn("col/nc", godip.Sea).Conn("ven", godip.Sea).Conn("gub", godip.Sea).Conn("dom", godip.Sea).Conn("cub", godip.Sea).Flag(godip.Sea).
		// Morocco
		Prov("mor").Conn("ble", godip.Sea).Conn("bar", godip.Coast...).Conn("cad", godip.Coast...).Conn("can", godip.Sea).Conn("mau", godip.Coast...).Conn("alg", godip.Coast...).Flag(godip.Coast...).
		// Serbia
		Prov("ser").Conn("alb", godip.Land).Conn("grc", godip.Land).Conn("bul", godip.Land).Conn("rma", godip.Land).Conn("hun", godip.Land).Conn("cro", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Mumbai
		Prov("mum").Conn("kam", godip.Coast...).Conn("ara", godip.Sea).Conn("lac", godip.Sea).Conn("bna", godip.Coast...).Conn("hyd", godip.Land).Conn("ned", godip.Land).Flag(godip.Coast...).SC(India).
		// Belgium
		Prov("beg").Conn("nos", godip.Sea).Conn("pai", godip.Coast...).Conn("lyo", godip.Land).Conn("ham", godip.Coast...).Conn("ska", godip.Sea).Flag(godip.Coast...).
		// Abula
		Prov("abu").Conn("gha", godip.Land).Conn("lag", godip.Land).Conn("kan", godip.Land).Conn("nig", godip.Land).Flag(godip.Land).SC(Nigeria).
		// Hungary
		Prov("hun").Conn("ukr", godip.Land).Conn("pol", godip.Land).Conn("cze", godip.Land).Conn("cro", godip.Land).Conn("ser", godip.Land).Conn("rma", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Ethiopia
		Prov("eth").Conn("mom", godip.Coast...).Conn("wio", godip.Sea).Conn("red", godip.Sea).Conn("sud", godip.Coast...).Conn("car", godip.Land).Conn("drc", godip.Land).Conn("tan", godip.Land).Conn("mar", godip.Land).Conn("yem", godip.Coast...).Flag(godip.Coast...).
		// Sapporo
		Prov("sap").Conn("soj", godip.Sea).Conn("tok", godip.Coast...).Conn("bes", godip.Sea).Flag(godip.Coast...).SC(Japan).
		// Sudan
		Prov("sud").Conn("ale", godip.Land).Conn("bnh", godip.Land).Conn("cha", godip.Land).Conn("car", godip.Land).Conn("eth", godip.Coast...).Conn("red", godip.Sea).Conn("asw", godip.Coast...).Flag(godip.Coast...).
		// Islamabad
		Prov("isl").Conn("kam", godip.Land).Conn("tib", godip.Land).Conn("ksi", godip.Land).Conn("afg", godip.Land).Conn("lah", godip.Land).Flag(godip.Land).SC(Pakistan).
		// Aegean Sea
		Prov("aeg").Conn("ank", godip.Sea).Conn("ank/sc", godip.Sea).Conn("ist", godip.Sea).Conn("bul", godip.Sea).Conn("bul/sc", godip.Sea).Conn("grc", godip.Sea).Conn("ems", godip.Sea).Conn("ira", godip.Sea).Conn("ira/wc", godip.Sea).Flag(godip.Sea).
		// Ukraine
		Prov("ukr").Conn("ros", godip.Coast...).Conn("mos", godip.Land).Conn("bea", godip.Land).Conn("pol", godip.Land).Conn("hun", godip.Land).Conn("rma", godip.Coast...).Conn("bla", godip.Sea).Flag(godip.Coast...).
		// Denmark
		Prov("dem").Conn("bat", godip.Sea).Conn("swe", godip.Coast...).Conn("ska", godip.Sea).Conn("ham", godip.Coast...).Conn("ben", godip.Coast...).Flag(godip.Coast...).
		// Bering Sea
		Prov("bes").Conn("arc", godip.Sea).Conn("now", godip.Sea).Conn("fin", godip.Sea).Conn("fin/nc", godip.Sea).Conn("lap", godip.Sea).Conn("soj", godip.Sea).Conn("sap", godip.Sea).Conn("tok", godip.Sea).Conn("ecs", godip.Sea).Conn("arc", godip.Sea).Flag(godip.Sea).
		// Los Angeles
		Prov("los").Conn("mex", godip.Land).Conn("neo", godip.Land).Conn("okl", godip.Land).Conn("chg", godip.Land).Conn("min", godip.Land).Flag(godip.Land).
		// Los Angeles (West Coast)
		Prov("los/wc").Conn("npo", godip.Sea).Conn("mex/wc", godip.Sea).Conn("min", godip.Sea).Conn("arc", godip.Sea).Flag(godip.Sea).
		// Los Angeles (East Coast)
		Prov("los/ec").Conn("mex/ec", godip.Sea).Conn("gux", godip.Sea).Conn("neo", godip.Sea).Flag(godip.Sea).
		// Buenos Aires
		Prov("bue").Conn("peu", godip.Land).Conn("men", godip.Land).Conn("com", godip.Coast...).Conn("wsa", godip.Sea).Conn("sma", godip.Sea).Conn("pag", godip.Coast...).Flag(godip.Coast...).SC(Argentina).
		// Banghazi
		Prov("bnh").Conn("cha", godip.Land).Conn("sud", godip.Land).Conn("ale", godip.Coast...).Conn("ems", godip.Sea).Conn("ion", godip.Sea).Conn("tri", godip.Coast...).Flag(godip.Coast...).
		// Paraguay
		Prov("pag").Conn("bue", godip.Coast...).Conn("sma", godip.Sea).Conn("nma", godip.Sea).Conn("sao", godip.Coast...).Conn("peu", godip.Land).Flag(godip.Coast...).
		// Omsk
		Prov("oms").Conn("stp", godip.Coast...).Conn("mos", godip.Land).Conn("ros", godip.Land).Conn("ast", godip.Land).Conn("irk", godip.Coast...).Conn("lap", godip.Sea).Flag(godip.Coast...).SC(Russia).
		// Rangoon
		Prov("ran").Conn("man", godip.Coast...).Conn("and", godip.Sea).Conn("hat", godip.Coast...).Conn("chm", godip.Land).Conn("vie", godip.Land).Conn("kum", godip.Land).Flag(godip.Coast...).
		// Panama
		Prov("pam").Conn("cas", godip.Sea).Conn("gux", godip.Sea).Conn("mex", godip.Land).Conn("mex/wc", godip.Sea).Conn("mex/ec", godip.Sea).Conn("mpo", godip.Sea).Conn("pab", godip.Sea).Conn("col", godip.Land).Conn("col/nc", godip.Sea).Conn("col/wc", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Madrid
		Prov("mad").Conn("bar", godip.Land).Conn("bod", godip.Land).Conn("bod/wc", godip.Sea).Conn("ena", godip.Sea).Conn("cad", godip.Coast...).Flag(godip.Coast...).SC(Spain).
		// Zambia
		Prov("zam").Conn("moq", godip.Land).Conn("tan", godip.Land).Conn("drc", godip.Land).Conn("ang", godip.Land).Conn("zim", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Istanbul
		Prov("ist").Conn("bul", godip.Land).Conn("bul/sc", godip.Sea).Conn("bul/ec", godip.Sea).Conn("aeg", godip.Sea).Conn("ank", godip.Land).Conn("ank/nc", godip.Sea).Conn("ank/sc", godip.Sea).Conn("bla", godip.Sea).Flag(godip.Coast...).SC(Turkey).
		// Arabian Sea
		Prov("ara").Conn("peg", godip.Sea).Conn("oma", godip.Sea).Conn("yem", godip.Sea).Conn("red", godip.Sea).Conn("lac", godip.Sea).Conn("mum", godip.Sea).Conn("kam", godip.Sea).Conn("kar", godip.Sea).Flag(godip.Sea).
		// Indonesia
		Prov("ind").Conn("and", godip.Sea).Conn("bob", godip.Sea).Conn("npo", godip.Sea).Conn("hat", godip.Coast...).Conn("dar", godip.Coast...).Flag(godip.Coast...).
		// Croatia
		Prov("cro").Conn("cze", godip.Land).Conn("mil", godip.Land).Conn("mil/ec", godip.Sea).Conn("ion", godip.Sea).Conn("alb", godip.Coast...).Conn("ser", godip.Land).Conn("hun", godip.Land).Flag(godip.Coast...).
		// Port Elizabeth
		Prov("por").Conn("dur", godip.Coast...).Conn("pre", godip.Land).Conn("zim", godip.Land).Conn("cap", godip.Coast...).Conn("soo", godip.Sea).Flag(godip.Coast...).
		// Tasman Sea
		Prov("tas").Conn("spo", godip.Sea).Conn("mpo", godip.Sea).Conn("syd", godip.Sea).Conn("ade", godip.Sea).Conn("pet", godip.Sea).Conn("eio", godip.Sea).Conn("mio", godip.Sea).Conn("ant", godip.Sea).Flag(godip.Sea).
		// Niger
		Prov("nig").Conn("gab", godip.Land).Conn("cha", godip.Land).Conn("tri", godip.Land).Conn("alg", godip.Land).Conn("gha", godip.Land).Conn("abu", godip.Land).Conn("kan", godip.Land).Flag(godip.Land).
		// Lanzhou
		Prov("lan").Conn("cho", godip.Land).Conn("bao", godip.Land).Conn("yum", godip.Land).Conn("uru", godip.Land).Conn("ksi", godip.Land).Conn("tib", godip.Land).Conn("lij", godip.Land).Flag(godip.Land).
		// Kano
		Prov("kan").Conn("gab", godip.Land).Conn("nig", godip.Land).Conn("abu", godip.Land).Conn("lag", godip.Land).Flag(godip.Land).SC(Nigeria).
		// Aswan
		Prov("asw").Conn("sud", godip.Coast...).Conn("red", godip.Sea).Conn("cai", godip.Coast...).Conn("ale", godip.Land).Flag(godip.Coast...).SC(Egypt).
		// Cadiz
		Prov("cad").Conn("mad", godip.Coast...).Conn("ena", godip.Sea).Conn("can", godip.Sea).Conn("mor", godip.Coast...).Conn("bar", godip.Coast...).Flag(godip.Coast...).SC(Spain).
		// Denmark Strait
		Prov("des").Conn("dub", godip.Sea).Conn("edi", godip.Sea).Conn("nos", godip.Sea).Conn("grd", godip.Sea).Conn("lab", godip.Sea).Conn("wna", godip.Sea).Conn("ena", godip.Sea).Flag(godip.Sea).
		// Chesapeake Bay
		Prov("che").Conn("was", godip.Sea).Conn("bet", godip.Sea).Conn("wna", godip.Sea).Conn("gos", godip.Sea).Conn("ney", godip.Sea).Flag(godip.Sea).
		// Poland
		Prov("pol").Conn("hun", godip.Land).Conn("ukr", godip.Land).Conn("bea", godip.Coast...).Conn("bat", godip.Sea).Conn("ben", godip.Coast...).Conn("mun", godip.Land).Conn("cze", godip.Land).Flag(godip.Coast...).
		// Yumen
		Prov("yum").Conn("uru", godip.Land).Conn("lan", godip.Land).Conn("bao", godip.Land).Conn("mog", godip.Land).Flag(godip.Land).SC(China).
		// North Mid Atlantic
		Prov("nma").Conn("cab", godip.Sea).Conn("rec", godip.Sea).Conn("bra", godip.Sea).Conn("rio", godip.Sea).Conn("sao", godip.Sea).Conn("pag", godip.Sea).Conn("sma", godip.Sea).Conn("sen", godip.Sea).Flag(godip.Sea).
		// Panama Basin
		Prov("pab").Conn("men", godip.Sea).Conn("peu", godip.Sea).Conn("col", godip.Sea).Conn("col/wc", godip.Sea).Conn("pam", godip.Sea).Conn("mpo", godip.Sea).Conn("chc", godip.Sea).Flag(godip.Sea).
		// Senegal
		Prov("sen").Conn("gin", godip.Coast...).Conn("mau", godip.Coast...).Conn("cab", godip.Sea).Conn("nma", godip.Sea).Conn("sma", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Rio de Janeiro
		Prov("rio").Conn("nma", godip.Sea).Conn("bra", godip.Coast...).Conn("sao", godip.Coast...).Flag(godip.Coast...).SC(Brazil).
		// Zimbabwe
		Prov("zim").Conn("ang", godip.Coast...).Conn("esa", godip.Sea).Conn("cap", godip.Coast...).Conn("por", godip.Land).Conn("pre", godip.Land).Conn("moq", godip.Land).Conn("zam", godip.Land).Flag(godip.Coast...).
		// Bay of Bengal
		Prov("bob").Conn("bad", godip.Sea).Conn("kol", godip.Sea).Conn("gur", godip.Sea).Conn("eio", godip.Sea).Conn("pet", godip.Sea).Conn("dar", godip.Sea).Conn("bri", godip.Sea).Conn("mpo", godip.Sea).Conn("npo", godip.Sea).Conn("ind", godip.Sea).Conn("and", godip.Sea).Conn("man", godip.Sea).Flag(godip.Sea).
		// Gulf of Guinea
		Prov("gog").Conn("ang", godip.Sea).Conn("gab", godip.Sea).Conn("lag", godip.Sea).Conn("gha", godip.Sea).Conn("cot", godip.Sea).Conn("esa", godip.Sea).Flag(godip.Sea).
		// Karachi
		Prov("kar").Conn("kam", godip.Coast...).Conn("lah", godip.Land).Conn("afg", godip.Land).Conn("mas", godip.Coast...).Conn("peg", godip.Sea).Conn("ara", godip.Sea).Flag(godip.Coast...).SC(Pakistan).
		// Persian Gulf
		Prov("peg").Conn("teh", godip.Sea).Conn("ira", godip.Sea).Conn("ira/ec", godip.Sea).Conn("riy", godip.Sea).Conn("oma", godip.Sea).Conn("ara", godip.Sea).Conn("kar", godip.Sea).Conn("mas", godip.Sea).Flag(godip.Sea).
		// Canaries
		Prov("can").Conn("cab", godip.Sea).Conn("mau", godip.Sea).Conn("mor", godip.Sea).Conn("cad", godip.Sea).Conn("ena", godip.Sea).Flag(godip.Sea).
		// Brisbane
		Prov("bri").Conn("bob", godip.Sea).Conn("dar", godip.Coast...).Conn("ade", godip.Land).Conn("ade", godip.Land).Conn("syd", godip.Coast...).Conn("mpo", godip.Sea).Flag(godip.Coast...).SC(Australia).
		// Kolkata
		Prov("kol").Conn("ned", godip.Land).Conn("hyd", godip.Coast...).Conn("gur", godip.Sea).Conn("bob", godip.Sea).Conn("bad", godip.Coast...).Conn("nep", godip.Land).Flag(godip.Coast...).
		// Dominican Republic
		Prov("dom").Conn("wna", godip.Sea).Conn("bet", godip.Sea).Conn("cub", godip.Coast...).Conn("cas", godip.Sea).Conn("gub", godip.Sea).Conn("ena", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// East China Sea
		Prov("ecs").Conn("phi", godip.Sea).Conn("arc", godip.Sea).Conn("bes", godip.Sea).Conn("tok", godip.Sea).Conn("nag", godip.Sea).Conn("soj", godip.Sea).Conn("kor", godip.Sea).Conn("yes", godip.Sea).Conn("tai", godip.Sea).Conn("scs", godip.Sea).Flag(godip.Sea).
		// Bangladesh
		Prov("bad").Conn("bob", godip.Sea).Conn("man", godip.Coast...).Conn("bhu", godip.Land).Conn("nep", godip.Land).Conn("kol", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Whitehorse
		Prov("whi").Conn("yeo", godip.Land).Conn("anc", godip.Land).Conn("van", godip.Land).Conn("yeo", godip.Land).Flag(godip.Land).
		// Whitehorse (North Coast)
		Prov("whi/nc").Conn("yeo", godip.Sea).Conn("arc", godip.Sea).Conn("anc", godip.Sea).Flag(godip.Sea).
		// Whitehorse (West Coast)
		Prov("whi/wc").Conn("van", godip.Sea).Conn("arc", godip.Sea).Conn("anc", godip.Sea).Flag(godip.Sea).
		// Mid Indian Ocean
		Prov("mio").Conn("gur", godip.Sea).Conn("bna", godip.Sea).Conn("lac", godip.Sea).Conn("wio", godip.Sea).Conn("moc", godip.Sea).Conn("ant", godip.Sea).Conn("tas", godip.Sea).Conn("eio", godip.Sea).Flag(godip.Sea).
		// Mendoza
		Prov("men").Conn("bue", godip.Land).Conn("peu", godip.Coast...).Conn("pab", godip.Sea).Conn("chc", godip.Sea).Conn("com", godip.Coast...).Flag(godip.Coast...).SC(Argentina).
		// Tehran
		Prov("teh").Conn("aze", godip.Land).Conn("diy", godip.Land).Conn("ira", godip.Land).Conn("ira/ec", godip.Sea).Conn("peg", godip.Sea).Conn("mas", godip.Coast...).Conn("tur", godip.Land).Flag(godip.Coast...).
		// Democratic Republic of the Congo
		Prov("drc").Conn("gab", godip.Land).Conn("ang", godip.Land).Conn("zam", godip.Land).Conn("tan", godip.Land).Conn("eth", godip.Land).Conn("car", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Iqaluit
		Prov("iqa").Conn("arc", godip.Sea).Conn("yeo", godip.Coast...).Conn("ott", godip.Coast...).Conn("lab", godip.Sea).Conn("grd", godip.Coast...).Flag(godip.Coast...).SC(Canada).
		// Perth
		Prov("pet").Conn("bob", godip.Sea).Conn("eio", godip.Sea).Conn("tas", godip.Sea).Conn("ade", godip.Coast...).Conn("dar", godip.Coast...).Flag(godip.Coast...).SC(Australia).
		// Southern Ocean
		Prov("soo").Conn("ant", godip.Sea).Conn("moc", godip.Sea).Conn("dur", godip.Sea).Conn("por", godip.Sea).Conn("cap", godip.Sea).Conn("sco", godip.Sea).Conn("spo", godip.Sea).Flag(godip.Sea).
		// Sydney
		Prov("syd").Conn("mpo", godip.Sea).Conn("bri", godip.Coast...).Conn("ade", godip.Coast...).Conn("tas", godip.Sea).Flag(godip.Coast...).SC(Australia).
		// Romania
		Prov("rma").Conn("bla", godip.Sea).Conn("ukr", godip.Coast...).Conn("hun", godip.Land).Conn("ser", godip.Land).Conn("bul", godip.Land).Conn("bul/ec", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Brasilia
		Prov("bra").Conn("rio", godip.Coast...).Conn("nma", godip.Sea).Conn("rec", godip.Coast...).Conn("mac", godip.Land).Flag(godip.Coast...).
		// West Indian Ocean
		Prov("wio").Conn("mio", godip.Sea).Conn("lac", godip.Sea).Conn("red", godip.Sea).Conn("eth", godip.Sea).Conn("mom", godip.Sea).Conn("moc", godip.Sea).Flag(godip.Sea).
		// Norwegian Sea
		Prov("nos").Conn("pai", godip.Sea).Conn("beg", godip.Sea).Conn("ska", godip.Sea).Conn("now", godip.Sea).Conn("arc", godip.Sea).Conn("grd", godip.Sea).Conn("des", godip.Sea).Conn("edi", godip.Sea).Conn("lon", godip.Sea).Flag(godip.Sea).
		// Czech Rep.
		Prov("cze").Conn("cro", godip.Land).Conn("hun", godip.Land).Conn("pol", godip.Land).Conn("mun", godip.Land).Conn("mil", godip.Land).Flag(godip.Land).
		// Laptev Sea
		Prov("lap").Conn("stp", godip.Sea).Conn("oms", godip.Sea).Conn("irk", godip.Sea).Conn("yak", godip.Sea).Conn("vla", godip.Sea).Conn("soj", godip.Sea).Conn("bes", godip.Sea).Conn("fin", godip.Sea).Conn("fin/nc", godip.Sea).Flag(godip.Sea).
		// Tripoli
		Prov("tri").Conn("bnh", godip.Coast...).Conn("ion", godip.Sea).Conn("nap", godip.Coast...).Conn("ble", godip.Sea).Conn("alg", godip.Coast...).Conn("nig", godip.Land).Conn("cha", godip.Land).Flag(godip.Coast...).
		// Balearic Sea
		Prov("ble").Conn("nap", godip.Sea).Conn("roe", godip.Sea).Conn("roe/wc", godip.Sea).Conn("mil", godip.Sea).Conn("mil/wc", godip.Sea).Conn("lyo", godip.Sea).Conn("bod", godip.Sea).Conn("bod/ec", godip.Sea).Conn("bar", godip.Sea).Conn("mor", godip.Sea).Conn("alg", godip.Sea).Conn("tri", godip.Sea).Flag(godip.Sea).
		// Yellow Sea
		Prov("yes").Conn("bei", godip.Sea).Conn("hon", godip.Sea).Conn("kum", godip.Sea).Conn("vie", godip.Sea).Conn("scs", godip.Sea).Conn("tai", godip.Sea).Conn("ecs", godip.Sea).Conn("kor", godip.Sea).Conn("she", godip.Sea).Conn("she/sc", godip.Sea).Flag(godip.Sea).
		// Darwin
		Prov("dar").Conn("bri", godip.Coast...).Conn("bob", godip.Sea).Conn("pet", godip.Coast...).Conn("ade", godip.Land).Conn("ind", godip.Coast...).Flag(godip.Coast...).
		// Skaggerak
		Prov("ska").Conn("now", godip.Sea).Conn("nos", godip.Sea).Conn("beg", godip.Sea).Conn("ham", godip.Sea).Conn("dem", godip.Sea).Conn("swe", godip.Sea).Flag(godip.Sea).
		// Greece
		Prov("grc").Conn("ser", godip.Land).Conn("alb", godip.Coast...).Conn("ion", godip.Sea).Conn("ems", godip.Sea).Conn("aeg", godip.Sea).Conn("bul", godip.Land).Conn("bul/sc", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Nagasaki
		Prov("nag").Conn("tok", godip.Coast...).Conn("soj", godip.Sea).Conn("ecs", godip.Sea).Conn("kor", godip.Coast...).Flag(godip.Coast...).SC(Japan).
		// Cairo
		Prov("cai").Conn("mec", godip.Land).Conn("mec/nc", godip.Sea).Conn("mec/sc", godip.Sea).Conn("ems", godip.Sea).Conn("ale", godip.Coast...).Conn("asw", godip.Coast...).Conn("red", godip.Sea).Flag(godip.Coast...).SC(Egypt).
		// Arctic Ocean
		Prov("arc").Conn("npo", godip.Sea).Conn("los", godip.Sea).Conn("los/wc", godip.Sea).Conn("min", godip.Sea).Conn("van", godip.Sea).Conn("whi", godip.Sea).Conn("whi/nc", godip.Sea).Conn("whi/wc", godip.Sea).Conn("anc", godip.Sea).Conn("anc", godip.Sea).Conn("yeo", godip.Sea).Conn("iqa", godip.Sea).Conn("grd", godip.Sea).Conn("nos", godip.Sea).Conn("now", godip.Sea).Conn("bes", godip.Sea).Conn("ecs", godip.Sea).Conn("phi", godip.Sea).Conn("npo", godip.Sea).Flag(godip.Sea).
		// Peru
		Prov("peu").Conn("bue", godip.Land).Conn("pag", godip.Land).Conn("col", godip.Land).Conn("col/wc", godip.Sea).Conn("pab", godip.Sea).Conn("men", godip.Coast...).Flag(godip.Coast...).
		// Chiang Mai
		Prov("chm").Conn("scs", godip.Sea).Conn("bnk", godip.Coast...).Conn("vie", godip.Land).Conn("ran", godip.Land).Conn("hat", godip.Coast...).Conn("npo", godip.Sea).Flag(godip.Coast...).SC(Thailand).
		// Andaman Sea
		Prov("and").Conn("bob", godip.Sea).Conn("ind", godip.Sea).Conn("hat", godip.Sea).Conn("ran", godip.Sea).Conn("man", godip.Sea).Flag(godip.Sea).
		// Sweden
		Prov("swe").Conn("now", godip.Coast...).Conn("ska", godip.Sea).Conn("dem", godip.Coast...).Conn("bat", godip.Sea).Conn("fin", godip.Land).Conn("fin/sc", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Sea of Japan
		Prov("soj").Conn("bes", godip.Sea).Conn("lap", godip.Sea).Conn("vla", godip.Sea).Conn("she", godip.Sea).Conn("she/ec", godip.Sea).Conn("kor", godip.Sea).Conn("ecs", godip.Sea).Conn("nag", godip.Sea).Conn("tok", godip.Sea).Conn("sap", godip.Sea).Flag(godip.Sea).
		// Durban
		Prov("dur").Conn("pre", godip.Coast...).Conn("por", godip.Coast...).Conn("soo", godip.Sea).Conn("moc", godip.Sea).Flag(godip.Coast...).SC(SouthAfrica).
		// Iraq
		Prov("ira").Conn("ank", godip.Land).Conn("mec", godip.Land).Conn("riy", godip.Land).Conn("teh", godip.Land).Conn("diy", godip.Land).Flag(godip.Land).
		// Iraq (West Coast)
		Prov("ira/wc").Conn("ank/sc", godip.Sea).Conn("aeg", godip.Sea).Conn("ems", godip.Sea).Conn("mec/nc", godip.Sea).Flag(godip.Sea).
		// Iraq (East Coast)
		Prov("ira/ec").Conn("riy", godip.Sea).Conn("peg", godip.Sea).Conn("teh", godip.Sea).Flag(godip.Sea).
		// Cabo Verde
		Prov("cab").Conn("nma", godip.Sea).Conn("sen", godip.Sea).Conn("mau", godip.Sea).Conn("can", godip.Sea).Conn("ena", godip.Sea).Conn("gub", godip.Sea).Conn("rec", godip.Sea).Flag(godip.Sea).
		// Washington DC
		Prov("was").Conn("che", godip.Sea).Conn("ney", godip.Coast...).Conn("chg", godip.Land).Conn("atl", godip.Land).Conn("neo", godip.Coast...).Conn("bet", godip.Sea).Flag(godip.Coast...).SC(USA).
		// Yellowknife
		Prov("yeo").Conn("ott", godip.Land).Conn("iqa", godip.Coast...).Conn("arc", godip.Sea).Conn("whi", godip.Land).Conn("whi/nc", godip.Sea).Conn("van", godip.Land).Flag(godip.Coast...).
		// Guinea
		Prov("gin").Conn("sen", godip.Coast...).Conn("sma", godip.Sea).Conn("wsa", godip.Sea).Conn("esa", godip.Sea).Conn("cot", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Mexico
		Prov("mex").Conn("los", godip.Land).Conn("pam", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Mexico (West Coast)
		Prov("mex/wc").Conn("los/wc", godip.Sea).Conn("npo", godip.Sea).Conn("mpo", godip.Sea).Conn("pam", godip.Sea).Flag(godip.Sea).
		// Mexico (East Coast)
		Prov("mex/ec").Conn("gux", godip.Sea).Conn("los/ec", godip.Sea).Conn("pam", godip.Sea).Flag(godip.Sea).
		// Nepal
		Prov("nep").Conn("tib", godip.Land).Conn("kam", godip.Land).Conn("ned", godip.Land).Conn("kol", godip.Land).Conn("bad", godip.Land).Conn("bhu", godip.Land).Flag(godip.Land).
		// Riyadh
		Prov("riy").Conn("mec", godip.Land).Conn("yem", godip.Land).Conn("oma", godip.Coast...).Conn("peg", godip.Sea).Conn("ira", godip.Land).Conn("ira/ec", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Milan
		Prov("mil").Conn("lyo", godip.Land).Conn("roe", godip.Land).Conn("cro", godip.Land).Conn("cze", godip.Land).Flag(godip.Land).SC(Italy).
		// Milan (West Coast)
		Prov("mil/wc").Conn("lyo", godip.Sea).Conn("ble", godip.Sea).Conn("roe/wc", godip.Sea).Flag(godip.Sea).
		// Milan (East Coast)
		Prov("mil/ec").Conn("roe/ec", godip.Sea).Conn("ion", godip.Sea).Conn("cro", godip.Sea).Flag(godip.Sea).
		// Diyarbakir
		Prov("diy").Conn("bla", godip.Sea).Conn("ank", godip.Land).Conn("ank/nc", godip.Sea).Conn("ira", godip.Land).Conn("teh", godip.Land).Conn("aze", godip.Coast...).Flag(godip.Coast...).SC(Turkey).
		// Barcelona
		Prov("bar").Conn("bod", godip.Land).Conn("bod/ec", godip.Sea).Conn("mad", godip.Land).Conn("cad", godip.Coast...).Conn("mor", godip.Coast...).Conn("ble", godip.Sea).Flag(godip.Coast...).SC(Spain).
		// Mozambique
		Prov("moq").Conn("tan", godip.Coast...).Conn("zam", godip.Land).Conn("zim", godip.Land).Conn("pre", godip.Coast...).Conn("moc", godip.Sea).Flag(godip.Coast...).
		// New Delhi
		Prov("ned").Conn("kol", godip.Land).Conn("nep", godip.Land).Conn("kam", godip.Land).Conn("mum", godip.Land).Conn("hyd", godip.Land).Flag(godip.Land).SC(India).
		// Black Sea
		Prov("bla").Conn("diy", godip.Sea).Conn("aze", godip.Sea).Conn("ros", godip.Sea).Conn("ukr", godip.Sea).Conn("rma", godip.Sea).Conn("bul", godip.Sea).Conn("bul/ec", godip.Sea).Conn("ist", godip.Sea).Conn("ank", godip.Sea).Conn("ank/nc", godip.Sea).Flag(godip.Sea).
		// Moscow
		Prov("mos").Conn("ros", godip.Land).Conn("oms", godip.Land).Conn("stp", godip.Land).Conn("bea", godip.Land).Conn("ukr", godip.Land).Flag(godip.Land).SC(Russia).
		// Hamburg
		Prov("ham").Conn("ben", godip.Land).Conn("dem", godip.Coast...).Conn("ska", godip.Sea).Conn("beg", godip.Coast...).Conn("lyo", godip.Land).Conn("mun", godip.Land).Flag(godip.Coast...).SC(Germany).
		// Central African Republic
		Prov("car").Conn("cha", godip.Land).Conn("gab", godip.Land).Conn("drc", godip.Land).Conn("eth", godip.Land).Conn("sud", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Nairobi
		Prov("nai").Conn("mom", godip.Land).Conn("mar", godip.Land).Conn("tan", godip.Land).Flag(godip.Land).SC(Kenya).
		// Astana
		Prov("ast").Conn("ros", godip.Land).Conn("aze", godip.Land).Conn("alm", godip.Land).Conn("mog", godip.Land).Conn("irk", godip.Land).Conn("oms", godip.Land).Flag(godip.Land).
		// Bermuda Triangle
		Prov("bet").Conn("neo", godip.Sea).Conn("gux", godip.Sea).Conn("cub", godip.Sea).Conn("dom", godip.Sea).Conn("wna", godip.Sea).Conn("che", godip.Sea).Conn("was", godip.Sea).Flag(godip.Sea).
		// Chilean Coast
		Prov("chc").Conn("pab", godip.Sea).Conn("mpo", godip.Sea).Conn("spo", godip.Sea).Conn("com", godip.Sea).Conn("men", godip.Sea).Flag(godip.Sea).
		// Turkmenistan
		Prov("tur").Conn("afg", godip.Land).Conn("uzb", godip.Land).Conn("alm", godip.Land).Conn("aze", godip.Land).Conn("teh", godip.Land).Conn("mas", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Cote D'Ivoire
		Prov("cot").Conn("esa", godip.Sea).Conn("gog", godip.Sea).Conn("gha", godip.Coast...).Conn("gin", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Vietnam
		Prov("vie").Conn("kum", godip.Coast...).Conn("ran", godip.Land).Conn("chm", godip.Land).Conn("bnk", godip.Coast...).Conn("scs", godip.Sea).Conn("yes", godip.Sea).Flag(godip.Coast...).
		// Marsabit
		Prov("mar").Conn("nai", godip.Land).Conn("mom", godip.Land).Conn("eth", godip.Land).Conn("tan", godip.Land).Flag(godip.Land).SC(Kenya).
		// Vancouver
		Prov("van").Conn("whi", godip.Land).Conn("whi/wc", godip.Sea).Conn("arc", godip.Sea).Conn("min", godip.Coast...).Conn("yeo", godip.Land).Flag(godip.Coast...).SC(Canada).
		// Bangalore
		Prov("bna").Conn("gur", godip.Sea).Conn("hyd", godip.Coast...).Conn("mum", godip.Coast...).Conn("lac", godip.Sea).Conn("mio", godip.Sea).Flag(godip.Coast...).SC(India).
		// Edinburgh
		Prov("edi").Conn("des", godip.Sea).Conn("dub", godip.Coast...).Conn("lon", godip.Coast...).Conn("nos", godip.Sea).Flag(godip.Coast...).SC(UK).
		// Belarus
		Prov("bea").Conn("bat", godip.Sea).Conn("pol", godip.Coast...).Conn("ukr", godip.Land).Conn("mos", godip.Land).Conn("stp", godip.Land).Conn("fin", godip.Land).Conn("fin/sc", godip.Sea).Flag(godip.Coast...).
		// Comodoro Rivadavia
		Prov("com").Conn("bue", godip.Coast...).Conn("men", godip.Coast...).Conn("chc", godip.Sea).Conn("spo", godip.Sea).Conn("wsa", godip.Sea).Flag(godip.Coast...).SC(Argentina).
		// Oman
		Prov("oma").Conn("yem", godip.Coast...).Conn("ara", godip.Sea).Conn("peg", godip.Sea).Conn("riy", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Angola
		Prov("ang").Conn("gog", godip.Sea).Conn("esa", godip.Sea).Conn("zim", godip.Coast...).Conn("zam", godip.Land).Conn("drc", godip.Land).Conn("gab", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Oklahoma
		Prov("okl").Conn("los", godip.Land).Conn("neo", godip.Land).Conn("atl", godip.Land).Conn("chg", godip.Land).Flag(godip.Land).SC(USA).
		// Baotou
		Prov("bao").Conn("bei", godip.Land).Conn("mog", godip.Land).Conn("yum", godip.Land).Conn("lan", godip.Land).Conn("cho", godip.Land).Flag(godip.Land).SC(China).
		// Finland
		Prov("fin").Conn("now", godip.Land).Conn("swe", godip.Land).Conn("bea", godip.Land).Conn("stp", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Finland (North Coast)
		Prov("fin/nc").Conn("lap", godip.Sea).Conn("bes", godip.Sea).Conn("now", godip.Sea).Conn("stp", godip.Sea).Flag(godip.Sea).
		// Finland (South Coast)
		Prov("fin/sc").Conn("swe", godip.Sea).Conn("bat", godip.Sea).Conn("bea", godip.Sea).Flag(godip.Sea).
		// South Pacific Ocean
		Prov("spo").Conn("mpo", godip.Sea).Conn("ant", godip.Sea).Conn("soo", godip.Sea).Conn("sco", godip.Sea).Conn("wsa", godip.Sea).Conn("com", godip.Sea).Conn("chc", godip.Sea).Conn("mpo", godip.Sea).Conn("tas", godip.Sea).Conn("ant", godip.Sea).Flag(godip.Sea).
		// Berlin
		Prov("ben").Conn("ham", godip.Land).Conn("mun", godip.Land).Conn("pol", godip.Coast...).Conn("bat", godip.Sea).Conn("dem", godip.Coast...).Flag(godip.Coast...).SC(Germany).
		// Lijiang
		Prov("lij").Conn("tib", godip.Land).Conn("bhu", godip.Land).Conn("man", godip.Land).Conn("kum", godip.Land).Conn("cho", godip.Land).Conn("lan", godip.Land).Flag(godip.Land).
		// Mid Pacific Ocean
		Prov("mpo").Conn("spo", godip.Sea).Conn("chc", godip.Sea).Conn("pab", godip.Sea).Conn("pam", godip.Sea).Conn("mex", godip.Sea).Conn("mex/wc", godip.Sea).Conn("npo", godip.Sea).Conn("spo", godip.Sea).Conn("npo", godip.Sea).Conn("bob", godip.Sea).Conn("bri", godip.Sea).Conn("syd", godip.Sea).Conn("tas", godip.Sea).Flag(godip.Sea).
		// Eastern North Atlantic
		Prov("ena").Conn("lon", godip.Sea).Conn("dub", godip.Sea).Conn("des", godip.Sea).Conn("wna", godip.Sea).Conn("dom", godip.Sea).Conn("gub", godip.Sea).Conn("cab", godip.Sea).Conn("can", godip.Sea).Conn("cad", godip.Sea).Conn("mad", godip.Sea).Conn("bod", godip.Sea).Conn("bod/wc", godip.Sea).Conn("pai", godip.Sea).Flag(godip.Sea).
		// Colombia
		Prov("col").Conn("ven", godip.Land).Conn("pam", godip.Land).Conn("peu", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Colombia (North Coast)
		Prov("col/nc").Conn("ven", godip.Sea).Conn("cas", godip.Sea).Conn("pam", godip.Sea).Flag(godip.Sea).
		// Colombia (West Coast)
		Prov("col/wc").Conn("pam", godip.Sea).Conn("pab", godip.Sea).Conn("peu", godip.Sea).Flag(godip.Sea).
		// Urumqi
		Prov("uru").Conn("afg", godip.Land).Conn("ksi", godip.Land).Conn("lan", godip.Land).Conn("yum", godip.Land).Conn("mog", godip.Land).Conn("alm", godip.Land).Flag(godip.Land).
		// Western South Atlantic
		Prov("wsa").Conn("sma", godip.Sea).Conn("bue", godip.Sea).Conn("com", godip.Sea).Conn("spo", godip.Sea).Conn("sco", godip.Sea).Conn("esa", godip.Sea).Conn("gin", godip.Sea).Flag(godip.Sea).
		// Mecca
		Prov("mec").Conn("yem", godip.Land).Conn("riy", godip.Land).Conn("ira", godip.Land).Conn("cai", godip.Land).Flag(godip.Land).
		// Mecca (North Coast)
		Prov("mec/nc").Conn("ira/wc", godip.Sea).Conn("ems", godip.Sea).Conn("cai", godip.Sea).Flag(godip.Sea).
		// Mecca (South Coast)
		Prov("mec/sc").Conn("yem", godip.Sea).Conn("cai", godip.Sea).Conn("red", godip.Sea).Flag(godip.Sea).
		// Pretoria
		Prov("pre").Conn("dur", godip.Coast...).Conn("moc", godip.Sea).Conn("moq", godip.Coast...).Conn("zim", godip.Land).Conn("por", godip.Land).Flag(godip.Coast...).SC(SouthAfrica).
		// Lyon
		Prov("lyo").Conn("pai", godip.Land).Conn("bod", godip.Land).Conn("bod/ec", godip.Sea).Conn("ble", godip.Sea).Conn("mil", godip.Land).Conn("mil/wc", godip.Sea).Conn("mun", godip.Land).Conn("ham", godip.Land).Conn("beg", godip.Land).Flag(godip.Coast...).
		// Naples
		Prov("nap").Conn("roe", godip.Land).Conn("roe/wc", godip.Sea).Conn("roe/ec", godip.Sea).Conn("ble", godip.Sea).Conn("tri", godip.Coast...).Conn("ion", godip.Sea).Flag(godip.Coast...).SC(Italy).
		// Venezuela
		Prov("ven").Conn("mac", godip.Coast...).Conn("gub", godip.Sea).Conn("cas", godip.Sea).Conn("col", godip.Land).Conn("col/nc", godip.Sea).Flag(godip.Coast...).
		// Montreal
		Prov("mot").Conn("ney", godip.Coast...).Conn("gos", godip.Sea).Conn("wna", godip.Sea).Conn("lab", godip.Sea).Conn("ott", godip.Coast...).Flag(godip.Coast...).SC(Canada).
		// Irkutsk
		Prov("irk").Conn("oms", godip.Coast...).Conn("ast", godip.Land).Conn("mog", godip.Land).Conn("yak", godip.Coast...).Conn("lap", godip.Sea).Flag(godip.Coast...).SC(Russia).
		// Bangkok
		Prov("bnk").Conn("scs", godip.Sea).Conn("vie", godip.Coast...).Conn("chm", godip.Coast...).Flag(godip.Coast...).SC(Thailand).
		// Shenyang
		Prov("she").Conn("bei", godip.Land).Conn("kor", godip.Land).Conn("vla", godip.Land).Conn("yak", godip.Land).Conn("mog", godip.Land).Flag(godip.Land).
		// Shenyang (South Coast)
		Prov("she/sc").Conn("bei", godip.Sea).Conn("yes", godip.Sea).Conn("kor", godip.Sea).Flag(godip.Sea).
		// Shenyang (East Coast)
		Prov("she/ec").Conn("kor", godip.Sea).Conn("soj", godip.Sea).Conn("vla", godip.Sea).Flag(godip.Sea).
		// Yakutsk
		Prov("yak").Conn("mog", godip.Land).Conn("she", godip.Land).Conn("vla", godip.Coast...).Conn("lap", godip.Sea).Conn("irk", godip.Coast...).Flag(godip.Coast...).
		// East Indian Ocean
		Prov("eio").Conn("pet", godip.Sea).Conn("bob", godip.Sea).Conn("gur", godip.Sea).Conn("mio", godip.Sea).Conn("tas", godip.Sea).Flag(godip.Sea).
		// Munich
		Prov("mun").Conn("lyo", godip.Land).Conn("cze", godip.Land).Conn("pol", godip.Land).Conn("ben", godip.Land).Conn("ham", godip.Land).Flag(godip.Land).SC(Germany).
		// Gulf of Mexico
		Prov("gux").Conn("cas", godip.Sea).Conn("cub", godip.Sea).Conn("bet", godip.Sea).Conn("neo", godip.Sea).Conn("los", godip.Sea).Conn("los/ec", godip.Sea).Conn("mex", godip.Sea).Conn("mex/ec", godip.Sea).Conn("pam", godip.Sea).Flag(godip.Sea).
		Done()
}
