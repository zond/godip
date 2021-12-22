package sengoku

import (
	"github.com/zond/godip"
	"github.com/zond/godip/graph"
	"github.com/zond/godip/state"
	"github.com/zond/godip/variants/classical"
	"github.com/zond/godip/variants/common"
)

const (
	Takeda    godip.Nation = "Takeda"
	Mori      godip.Nation = "Mori"
	Chosokabe godip.Nation = "Chosokabe"
	Hojo      godip.Nation = "Hojo"
	Oda       godip.Nation = "Oda"
	Shimazu   godip.Nation = "Shimazu"
	Uesugi    godip.Nation = "Uesugi"
)

var Nations = []godip.Nation{Takeda, Mori, Chosokabe, Hojo, Oda, Shimazu, Uesugi}

var SengokuVariant = common.Variant{
	Name:              "Sengoku",
	Graph:             func() godip.Graph { return SengokuGraph() },
	Start:             SengokuStart,
	Blank:             SengokuBlank,
	Phase:             classical.NewPhase,
	Parser:            classical.Parser,
	Nations:           Nations,
	PhaseTypes:        classical.PhaseTypes,
	Seasons:           classical.Seasons,
	UnitTypes:         classical.UnitTypes,
	SoloWinner:        common.SCCountWinner(19),
	SoloSCCount:       func(*state.State) int { return 19 },
	ProvinceLongNames: provinceLongNames,
	SVGMap: func() ([]byte, error) {
		return Asset("svg/sengokumap.svg")
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
	CreatedBy:   "Benjamin Hester",
	Version:     "1.0",
	Description: "Battle it out during the Sengoku (warring states) period of 16th Century Japan which collapsed the feudal system under the Ashikaga Shogunate. Select one of seven clans to become the new Shogun.",
	Rules:       "The first to 25 Supply Centers (SC) is the winner. Units may be built at any owned SC. There are 6 bridges connecting provinces (dashed lines) across the water.",
}

func SengokuBlank(phase godip.Phase) *state.State {
	return state.New(SengokuGraph(), phase, classical.BackupRule, nil, nil)
}

func SengokuStart() (result *state.State, err error) {
	startPhase := classical.NewPhase(1570, godip.Spring, godip.Movement)
	result = SengokuBlank(startPhase)
	if err = result.SetUnits(map[godip.Province]godip.Unit{
		"sos": godip.Unit{godip.Army, Takeda},
		"kai": godip.Unit{godip.Army, Takeda},
		"nag": godip.Unit{godip.Fleet, Mori},
		"iwm": godip.Unit{godip.Army, Mori},
		"iyo": godip.Unit{godip.Fleet, Chosokabe},
		"toa": godip.Unit{godip.Fleet, Chosokabe},
		"izu": godip.Unit{godip.Fleet, Hojo},
		"saa": godip.Unit{godip.Army, Hojo},
		"mik": godip.Unit{godip.Fleet, Oda},
		"owa": godip.Unit{godip.Army, Oda},
		"osu": godip.Unit{godip.Fleet, Shimazu},
		"sat": godip.Unit{godip.Army, Shimazu},
		"ecg": godip.Unit{godip.Fleet, Uesugi},
		"koz": godip.Unit{godip.Army, Uesugi},
	}); err != nil {
		return
	}
	result.SetSupplyCenters(map[godip.Province]godip.Nation{
		"sos": Takeda,
		"kai": Takeda,
		"nag": Mori,
		"iwm": Mori,
		"iyo": Chosokabe,
		"toa": Chosokabe,
		"saa": Hojo,
		"izu": Hojo,
		"mik": Oda,
		"owa": Oda,
		"osu": Shimazu,
		"sat": Shimazu,
		"ecg": Uesugi,
		"koz": Uesugi,
	})
	return
}

func SengokuGraph() *graph.Graph {
	return graph.New().
		// Mutsu
		Prov("mut").Conn("dew", godip.Coast...).Conn("sht", godip.Land).Conn("hit", godip.Coast...).Conn("npo", godip.Sea).Conn("tsu", godip.Sea).Conn("nso", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Kaga
		Prov("kag").Conn("etc", godip.Land).Conn("not", godip.Coast...).Conn("wak", godip.Sea).Conn("ecz", godip.Coast...).Conn("hid", godip.Land).Flag(godip.Coast...).
		// Sanuki
		Prov("sau").Conn("inl", godip.Sea).Conn("iyo", godip.Coast...).Conn("bit", godip.Coast...).Conn("awa", godip.Coast...).Conn("osa", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Hizen
		Prov("hiz").Conn("tus", godip.Sea).Conn("eas", godip.Sea).Conn("hig", godip.Coast...).Conn("bug", godip.Land).Conn("chi", godip.Coast...).Conn("kan", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Suruga
		Prov("sur").Conn("sos", godip.Land).Conn("too", godip.Coast...).Conn("sab", godip.Sea).Conn("izu", godip.Coast...).Conn("kai", godip.Land).Flag(godip.Coast...).
		// Iwami
		Prov("iwm").Conn("iwc", godip.Sea).Conn("nag", godip.Coast...).Conn("suo", godip.Land).Conn("aki", godip.Land).Conn("izm", godip.Coast...).Flag(godip.Coast...).SC(Mori).
		// Kii
		Prov("kii").Conn("kum", godip.Sea).Conn("tse", godip.Sea).Conn("shi", godip.Coast...).Conn("awa", godip.Coast...).Conn("yat", godip.Land).Conn("kaw", godip.Coast...).Conn("osa", godip.Sea).Conn("kis", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Sagami Bay
		Prov("sab").Conn("pac", godip.Sea).Conn("kum", godip.Sea).Conn("sur", godip.Sea).Conn("saa", godip.Coast...).Conn("izu", godip.Sea).Conn("tse", godip.Sea).Conn("too", godip.Sea).Conn("saa", godip.Sea).Conn("kaz", godip.Sea).Flag(godip.Coast...).
		// Set
		Prov("set").Conn("kaw", godip.Coast...).Conn("yas", godip.Land).Conn("har", godip.Coast...).Conn("osa", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Izu
		Prov("izu").Conn("sur", godip.Coast...).Conn("saa", godip.Coast...).Conn("sab", godip.Sea).Flag(godip.Sea).SC(Hojo).
		// Totomi Sea
		Prov("tse").Conn("sab", godip.Sea).Conn("too", godip.Sea).Conn("mik", godip.Sea).Conn("owa", godip.Sea).Conn("ise", godip.Sea).Conn("shi", godip.Sea).Conn("kii", godip.Sea).Conn("kum", godip.Sea).Flag(godip.Sea).
		// Tosa
		Prov("toa").Conn("awa", godip.Coast...).Conn("iyo", godip.Coast...).Conn("bus", godip.Sea).Conn("hyc", godip.Sea).Conn("tob", godip.Sea).Conn("kis", godip.Sea).Flag(godip.Coast...).SC(Chosokabe).
		// Mino
		Prov("min").Conn("too", godip.Land).Conn("sos", godip.Land).Conn("hid", godip.Land).Conn("ecz", godip.Land).Conn("omi", godip.Land).Conn("ise", godip.Land).Conn("owa", godip.Land).Conn("mik", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Totomi
		Prov("too").Conn("min", godip.Land).Conn("mik", godip.Coast...).Conn("tse", godip.Sea).Conn("sab", godip.Sea).Conn("sur", godip.Coast...).Conn("sos", godip.Land).Flag(godip.Coast...).SC(godip.Neutral).
		// Hyuga
		Prov("hyg").Conn("hig", godip.Land).Conn("sat", godip.Land).Conn("osu", godip.Coast...).Conn("ari", godip.Sea).Conn("hyc", godip.Sea).Conn("bus", godip.Sea).Conn("bug", godip.Coast...).Flag(godip.Coast...).
		// Osumi
		Prov("osu").Conn("hyg", godip.Coast...).Conn("sat", godip.Coast...).Conn("eas", godip.Sea).Conn("ari", godip.Sea).Flag(godip.Coast...).SC(Shimazu).
		// Musashi
		Prov("mus").Conn("koz", godip.Land).Conn("ecg", godip.Land).Conn("nos", godip.Land).Conn("kai", godip.Land).Conn("saa", godip.Land).Conn("shs", godip.Land).Conn("hit", godip.Land).Conn("sht", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Hyuga Coast
		Prov("hyc").Conn("ari", godip.Sea).Conn("spo", godip.Sea).Conn("tob", godip.Sea).Conn("toa", godip.Sea).Conn("bus", godip.Sea).Conn("hyg", godip.Sea).Flag(godip.Sea).
		// Shima
		Prov("shi").Conn("kii", godip.Coast...).Conn("mik", godip.Coast...).Conn("tse", godip.Sea).Conn("ise", godip.Coast...).Conn("yat", godip.Land).Flag(godip.Coast...).
		// Noto
		Prov("not").Conn("kag", godip.Coast...).Conn("etc", godip.Coast...).Conn("toj", godip.Sea).Conn("nso", godip.Sea).Conn("sea", godip.Sea).Conn("wak", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Kumano Sea
		Prov("kum").Conn("sab", godip.Sea).Conn("tse", godip.Sea).Conn("kii", godip.Sea).Conn("kis", godip.Sea).Conn("spo", godip.Sea).Conn("pac", godip.Sea).Flag(godip.Sea).
		// North Sea of Japan
		Prov("nso").Conn("sea", godip.Sea).Conn("not", godip.Sea).Conn("toj", godip.Sea).Conn("ecg", godip.Sea).Conn("dew", godip.Sea).Conn("mut", godip.Sea).Conn("tsu", godip.Sea).Flag(godip.Sea).
		// Kawachi
		Prov("kaw").Conn("osa", godip.Sea).Conn("kii", godip.Coast...).Conn("yat", godip.Land).Conn("yas", godip.Land).Conn("set", godip.Coast...).Flag(godip.Coast...).
		// Kai
		Prov("kai").Conn("saa", godip.Land).Conn("mus", godip.Land).Conn("nos", godip.Land).Conn("sos", godip.Land).Conn("sur", godip.Land).Flag(godip.Land).SC(Takeda).
		// Tosa Bay
		Prov("tob").Conn("kis", godip.Sea).Conn("toa", godip.Sea).Conn("hyc", godip.Sea).Conn("spo", godip.Sea).Flag(godip.Sea).
		// Yamato
		Prov("yat").Conn("kaw", godip.Land).Conn("kii", godip.Land).Conn("shi", godip.Land).Conn("ise", godip.Land).Conn("yas", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Kii Straits
		Prov("kis").Conn("tob", godip.Sea).Conn("spo", godip.Sea).Conn("kum", godip.Sea).Conn("kii", godip.Sea).Conn("osa", godip.Sea).Conn("awa", godip.Sea).Conn("toa", godip.Sea).Flag(godip.Sea).
		// Wakasa Bay
		Prov("wak").Conn("sac", godip.Sea).Conn("tan", godip.Sea).Conn("ecz", godip.Sea).Conn("kag", godip.Sea).Conn("not", godip.Sea).Conn("sea", godip.Sea).Flag(godip.Sea).
		// East China Sea
		Prov("eas").Conn("shu", godip.Sea).Conn("ari", godip.Sea).Conn("osu", godip.Sea).Conn("sat", godip.Sea).Conn("hig", godip.Sea).Conn("hiz", godip.Sea).Conn("tus", godip.Sea).Flag(godip.Sea).
		// Shimosa
		Prov("shs").Conn("kaz", godip.Coast...).Conn("npo", godip.Sea).Conn("hit", godip.Coast...).Conn("mus", godip.Land).Conn("saa", godip.Land).Flag(godip.Coast...).
		// Inland Sea
		Prov("inl").Conn("sau", godip.Sea).Conn("osa", godip.Sea).Conn("bit", godip.Sea).Conn("bin", godip.Sea).Conn("aki", godip.Sea).Conn("suo", godip.Sea).Conn("sus", godip.Sea).Conn("iyo", godip.Sea).Flag(godip.Sea).
		// Harima
		Prov("har").Conn("tan", godip.Land).Conn("biz", godip.Coast...).Conn("osa", godip.Sea).Conn("set", godip.Coast...).Conn("yas", godip.Land).Conn("omi", godip.Land).Flag(godip.Coast...).
		// Owari
		Prov("owa").Conn("tse", godip.Sea).Conn("mik", godip.Coast...).Conn("min", godip.Land).Conn("ise", godip.Coast...).Flag(godip.Coast...).SC(Oda).
		// Bungo
		Prov("bug").Conn("hig", godip.Land).Conn("hyg", godip.Coast...).Conn("iyo", godip.Coast...).Conn("bus", godip.Sea).Conn("sus", godip.Sea).Conn("chi", godip.Coast...).Conn("hiz", godip.Land).Flag(godip.Coast...).SC(godip.Neutral).
		// Pacific Ocean
		Prov("pac").Conn("npo", godip.Sea).Conn("kaz", godip.Sea).Conn("sab", godip.Sea).Conn("kum", godip.Sea).Conn("spo", godip.Sea).Flag(godip.Sea).
		// South Sea of Japan
		Prov("ssj").Conn("tus", godip.Sea).Conn("kan", godip.Sea).Conn("iwc", godip.Sea).Conn("sea", godip.Sea).Flag(godip.Sea).
		// North Pacific Ocean
		Prov("npo").Conn("tsu", godip.Sea).Conn("mut", godip.Sea).Conn("hit", godip.Sea).Conn("shs", godip.Sea).Conn("kaz", godip.Sea).Conn("pac", godip.Sea).Flag(godip.Sea).
		// Nagato
		Prov("nag").Conn("iwc", godip.Sea).Conn("chi", godip.Coast...).Conn("kan", godip.Sea).Conn("sus", godip.Sea).Conn("suo", godip.Coast...).Conn("iwm", godip.Coast...).Flag(godip.Coast...).SC(Mori).
		// Tojama Bay
		Prov("toj").Conn("etc", godip.Sea).Conn("ecg", godip.Sea).Conn("nso", godip.Sea).Conn("not", godip.Sea).Flag(godip.Sea).
		// Mikawa
		Prov("mik").Conn("too", godip.Coast...).Conn("shi", godip.Coast...).Conn("min", godip.Land).Conn("owa", godip.Coast...).Conn("tse", godip.Sea).Flag(godip.Coast...).SC(Oda).
		// Kanmon Straits
		Prov("kan").Conn("nag", godip.Sea).Conn("iwc", godip.Sea).Conn("ssj", godip.Sea).Conn("tus", godip.Sea).Conn("hiz", godip.Sea).Conn("chi", godip.Sea).Conn("sus", godip.Sea).Flag(godip.Sea).
		// South Shinano
		Prov("sos").Conn("nos", godip.Land).Conn("hid", godip.Land).Conn("min", godip.Land).Conn("too", godip.Land).Conn("sur", godip.Land).Conn("kai", godip.Land).Flag(godip.Land).SC(Takeda).
		// Etchu
		Prov("etc").Conn("toj", godip.Sea).Conn("not", godip.Coast...).Conn("kag", godip.Land).Conn("hid", godip.Land).Conn("nos", godip.Land).Conn("ecg", godip.Coast...).Flag(godip.Coast...).
		// Suo Sea
		Prov("sus").Conn("bus", godip.Sea).Conn("iyo", godip.Sea).Conn("inl", godip.Sea).Conn("suo", godip.Sea).Conn("nag", godip.Sea).Conn("kan", godip.Sea).Conn("chi", godip.Sea).Conn("bug", godip.Sea).Flag(godip.Sea).
		// Suo
		Prov("suo").Conn("inl", godip.Sea).Conn("aki", godip.Coast...).Conn("iyo", godip.Coast...).Conn("iwm", godip.Land).Conn("nag", godip.Coast...).Conn("sus", godip.Sea).Flag(godip.Coast...).
		// Bitchu
		Prov("bit").Conn("bin", godip.Coast...).Conn("sau", godip.Coast...).Conn("inl", godip.Sea).Conn("osa", godip.Sea).Conn("biz", godip.Coast...).Conn("hok", godip.Land).Flag(godip.Coast...).
		// Awa
		Prov("awa").Conn("toa", godip.Coast...).Conn("kii", godip.Coast...).Conn("kis", godip.Sea).Conn("osa", godip.Sea).Conn("sau", godip.Coast...).Conn("iyo", godip.Land).Flag(godip.Coast...).
		// Tushima Strait
		Prov("tus").Conn("eas", godip.Sea).Conn("hiz", godip.Sea).Conn("kan", godip.Sea).Conn("ssj", godip.Sea).Flag(godip.Sea).
		// Hida
		Prov("hid").Conn("kag", godip.Land).Conn("ecz", godip.Land).Conn("min", godip.Land).Conn("sos", godip.Land).Conn("nos", godip.Land).Conn("etc", godip.Land).Flag(godip.Land).
		// Echigo
		Prov("ecg").Conn("nos", godip.Land).Conn("mus", godip.Land).Conn("koz", godip.Land).Conn("dew", godip.Coast...).Conn("nso", godip.Sea).Conn("toj", godip.Sea).Conn("etc", godip.Coast...).Flag(godip.Coast...).SC(Uesugi).
		// Hitachi
		Prov("hit").Conn("mut", godip.Coast...).Conn("sht", godip.Land).Conn("mus", godip.Land).Conn("shs", godip.Coast...).Conn("npo", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Izumo
		Prov("izm").Conn("aki", godip.Land).Conn("bin", godip.Land).Conn("hok", godip.Coast...).Conn("sac", godip.Sea).Conn("iwc", godip.Sea).Conn("iwm", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Yamash
		Prov("yas").Conn("ise", godip.Land).Conn("omi", godip.Land).Conn("har", godip.Land).Conn("set", godip.Land).Conn("kaw", godip.Land).Conn("yat", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Dewa
		Prov("dew").Conn("mut", godip.Coast...).Conn("nso", godip.Sea).Conn("ecg", godip.Coast...).Conn("koz", godip.Land).Conn("sht", godip.Land).Flag(godip.Coast...).
		// Satsuma
		Prov("sat").Conn("hyg", godip.Land).Conn("hig", godip.Coast...).Conn("eas", godip.Sea).Conn("osu", godip.Coast...).Flag(godip.Coast...).SC(Shimazu).
		// Sea of Japan
		Prov("sea").Conn("not", godip.Sea).Conn("nso", godip.Sea).Conn("ssj", godip.Sea).Conn("iwc", godip.Sea).Conn("sac", godip.Sea).Conn("wak", godip.Sea).Flag(godip.Sea).
		// Tango
		Prov("tan").Conn("har", godip.Land).Conn("omi", godip.Land).Conn("ecz", godip.Coast...).Conn("wak", godip.Sea).Conn("sac", godip.Sea).Conn("ina", godip.Coast...).Conn("biz", godip.Land).Flag(godip.Coast...).
		// Shimotsuke
		Prov("sht").Conn("mus", godip.Land).Conn("hit", godip.Land).Conn("mut", godip.Land).Conn("dew", godip.Land).Conn("koz", godip.Land).Flag(godip.Land).
		// Sanin Coast
		Prov("sac").Conn("hok", godip.Sea).Conn("ina", godip.Sea).Conn("tan", godip.Sea).Conn("wak", godip.Sea).Conn("sea", godip.Sea).Conn("iwc", godip.Sea).Conn("izm", godip.Sea).Flag(godip.Sea).
		// Iwami Coast
		Prov("iwc").Conn("iwm", godip.Sea).Conn("izm", godip.Sea).Conn("sac", godip.Sea).Conn("sea", godip.Sea).Conn("ssj", godip.Sea).Conn("kan", godip.Sea).Conn("nag", godip.Sea).Flag(godip.Sea).
		// Bingo
		Prov("bin").Conn("bit", godip.Coast...).Conn("hok", godip.Land).Conn("izm", godip.Land).Conn("aki", godip.Coast...).Conn("inl", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Kazusa
		Prov("kaz").Conn("shs", godip.Coast...).Conn("saa", godip.Coast...).Conn("sab", godip.Sea).Conn("pac", godip.Sea).Conn("npo", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Kozuke
		Prov("koz").Conn("mus", godip.Land).Conn("sht", godip.Land).Conn("dew", godip.Land).Conn("ecg", godip.Land).Flag(godip.Land).SC(Uesugi).
		// Iyo
		Prov("iyo").Conn("bus", godip.Sea).Conn("toa", godip.Coast...).Conn("suo", godip.Coast...).Conn("bug", godip.Coast...).Conn("awa", godip.Land).Conn("sau", godip.Coast...).Conn("inl", godip.Sea).Conn("sus", godip.Sea).Flag(godip.Coast...).SC(Chosokabe).
		// Higo
		Prov("hig").Conn("hyg", godip.Land).Conn("bug", godip.Land).Conn("hiz", godip.Coast...).Conn("eas", godip.Sea).Conn("sat", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// North Shinano
		Prov("nos").Conn("sos", godip.Land).Conn("kai", godip.Land).Conn("mus", godip.Land).Conn("ecg", godip.Land).Conn("etc", godip.Land).Conn("hid", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Ise
		Prov("ise").Conn("yas", godip.Land).Conn("yat", godip.Land).Conn("shi", godip.Coast...).Conn("tse", godip.Sea).Conn("owa", godip.Coast...).Conn("min", godip.Land).Conn("omi", godip.Land).Flag(godip.Coast...).SC(godip.Neutral).
		// Bungo Straits
		Prov("bus").Conn("iyo", godip.Sea).Conn("sus", godip.Sea).Conn("bug", godip.Sea).Conn("hyg", godip.Sea).Conn("hyc", godip.Sea).Conn("toa", godip.Sea).Flag(godip.Sea).
		// Echizen
		Prov("ecz").Conn("omi", godip.Land).Conn("min", godip.Land).Conn("hid", godip.Land).Conn("kag", godip.Coast...).Conn("wak", godip.Sea).Conn("tan", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Omi
		Prov("omi").Conn("tan", godip.Land).Conn("har", godip.Land).Conn("yas", godip.Land).Conn("ise", godip.Land).Conn("min", godip.Land).Conn("ecz", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// South Pacific Ocean
		Prov("spo").Conn("pac", godip.Sea).Conn("kum", godip.Sea).Conn("kis", godip.Sea).Conn("tob", godip.Sea).Conn("hyc", godip.Sea).Conn("ari", godip.Sea).Conn("shu", godip.Sea).Flag(godip.Sea).
		// Osaka Bay
		Prov("osa").Conn("kaw", godip.Sea).Conn("set", godip.Sea).Conn("har", godip.Sea).Conn("biz", godip.Sea).Conn("bit", godip.Sea).Conn("inl", godip.Sea).Conn("sau", godip.Sea).Conn("awa", godip.Sea).Conn("kis", godip.Sea).Conn("kii", godip.Sea).Flag(godip.Sea).
		// Bizen
		Prov("biz").Conn("hok", godip.Land).Conn("bit", godip.Coast...).Conn("osa", godip.Sea).Conn("har", godip.Coast...).Conn("tan", godip.Land).Conn("ina", godip.Land).Flag(godip.Coast...).SC(godip.Neutral).
		// Inaba
		Prov("ina").Conn("sac", godip.Sea).Conn("hok", godip.Coast...).Conn("biz", godip.Land).Conn("tan", godip.Coast...).Flag(godip.Coast...).SC(godip.Neutral).
		// Hoki
		Prov("hok").Conn("sac", godip.Sea).Conn("izm", godip.Coast...).Conn("bin", godip.Land).Conn("bit", godip.Land).Conn("biz", godip.Land).Conn("ina", godip.Coast...).Flag(godip.Coast...).
		// Aki
		Prov("aki").Conn("izm", godip.Land).Conn("iwm", godip.Land).Conn("suo", godip.Coast...).Conn("inl", godip.Sea).Conn("bin", godip.Coast...).Flag(godip.Coast...).
		// Tsugaru Straits
		Prov("tsu").Conn("nso", godip.Sea).Conn("mut", godip.Sea).Conn("npo", godip.Sea).Flag(godip.Sea).
		// Ariake Bay
		Prov("ari").Conn("shu", godip.Sea).Conn("spo", godip.Sea).Conn("hyc", godip.Sea).Conn("hyg", godip.Sea).Conn("osu", godip.Sea).Conn("eas", godip.Sea).Flag(godip.Sea).
		// Shuri Straits
		Prov("shu").Conn("spo", godip.Sea).Conn("ari", godip.Sea).Conn("eas", godip.Sea).Flag(godip.Sea).
		// Chikuzen
		Prov("chi").Conn("bug", godip.Coast...).Conn("nag", godip.Coast...).Conn("sus", godip.Sea).Conn("kan", godip.Sea).Conn("hiz", godip.Coast...).Flag(godip.Coast...).
		// Sagami
		Prov("saa").Conn("kai", godip.Land).Conn("izu", godip.Coast...).Conn("sab", godip.Sea).Conn("kaz", godip.Coast...).Conn("shs", godip.Land).Conn("mus", godip.Land).Flag(godip.Coast...).SC(Hojo).
		Done()
}

var provinceLongNames = map[godip.Province]string{
	"mut": "Mutsu",
	"kag": "Kaga",
	"sau": "Sanuki",
	"hiz": "Hizen",
	"sur": "Suruga",
	"iwm": "Iwami",
	"kii": "Kii",
	"sab": "Sagami Bay",
	"set": "Set",
	"izu": "Izu",
	"tse": "Totomi Sea",
	"toa": "Tosa",
	"min": "Mino",
	"too": "Totomi",
	"hyg": "Hyuga",
	"osu": "Osumi",
	"mus": "Musashi",
	"hyc": "Hyuga Coast",
	"shi": "Shima",
	"not": "Noto",
	"kum": "Kumano Sea",
	"nso": "North Sea of Japan",
	"kaw": "Kawachi",
	"kai": "Kai",
	"tob": "Tosa Bay",
	"yat": "Yamato",
	"kis": "Kii Straits",
	"wak": "Wakasa Bay",
	"eas": "East China Sea",
	"shs": "Shimosa",
	"inl": "Inland Sea",
	"har": "Harima",
	"owa": "Owari",
	"bug": "Bungo",
	"pac": "Pacific Ocean",
	"ssj": "South Sea of Japan",
	"npo": "North Pacific Ocean",
	"nag": "Nagato",
	"toj": "Tojama Bay",
	"mik": "Mikawa",
	"kan": "Kanmon Straits",
	"sos": "South Shinano",
	"etc": "Etchu",
	"sus": "Suo Sea",
	"suo": "Suo",
	"bit": "Bitchu",
	"awa": "Awa",
	"tus": "Tushima Strait",
	"hid": "Hida",
	"ecg": "Echigo",
	"hit": "Hitachi",
	"izm": "Izumo",
	"yas": "Yamash",
	"dew": "Dewa",
	"sat": "Satsuma",
	"sea": "Sea of Japan",
	"tan": "Tango",
	"sht": "Shimotsuke",
	"sac": "Sanin Coast",
	"iwc": "Iwami Coast",
	"bin": "Bingo",
	"kaz": "Kazusa",
	"koz": "Kozuke",
	"iyo": "Iyo",
	"hig": "Higo",
	"nos": "North Shinano",
	"ise": "Ise",
	"bus": "Bungo Straits",
	"ecz": "Echizen",
	"omi": "Omi",
	"spo": "South Pacific Ocean",
	"osa": "Osaka Bay",
	"biz": "Bizen",
	"ina": "Inaba",
	"hok": "Hoki",
	"aki": "Aki",
	"tsu": "Tsugaru Straits",
	"ari": "Ariake Bay",
	"shu": "Shuri Straits",
	"chi": "Chikuzen",
	"saa": "Sagami",
}
