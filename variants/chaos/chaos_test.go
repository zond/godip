package chaos

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/zond/godip"
	"github.com/zond/godip/orders"
	"github.com/zond/godip/state"
	"github.com/zond/godip/variants/classical"
	tst "github.com/zond/godip/variants/testing"
)

func TestDefaultBuild(t *testing.T) {
	j, err := Start()
	if err != nil {
		t.Fatal(err)
	}
	j.SetOrders(map[godip.Province]godip.Adjudicator{
		"tri": orders.BuildAnywhere("tri", godip.Fleet, time.Now()),
	})
	j.Next()
	tst.AssertUnit(t, j, "tri", godip.Unit{godip.Fleet, "Trieste"})
	tst.AssertUnit(t, j, "vie", godip.Unit{godip.Army, "Vienna"})
}

// Test to verify that por doesn't get an option to move to mar in
// https://diplicity-engine.appspot.com/Game/ahJzfmRpcGxpY2l0eS1lbmdpbmVyEQsSBEdhbWUYgICAkLyDlQoM/Phase/17/Map
func TestPORConvoyOpts(t *testing.T) {
	g := Graph()
	judge := state.New(g, Phase(1903, godip.Fall, godip.Movement), classical.BackupRule)
	scMap := map[godip.Province]godip.Nation{}
	for _, prov := range g.Provinces() {
		if nat := g.SC(prov); nat != nil {
			scMap[prov.Super()] = *nat
		}
	}
	unitJSON := `[
      {
        "Province": "bud",
        "Unit": {
          "Type": "Army",
          "Nation": "Budapest"
        }
      },
      {
        "Province": "tri",
        "Unit": {
          "Type": "Army",
          "Nation": "Budapest"
        }
      },
      {
        "Province": "mid",
        "Unit": {
          "Type": "Fleet",
          "Nation": "London"
        }
      },
      {
        "Province": "por",
        "Unit": {
          "Type": "Army",
          "Nation": "Spain"
        }
      },
      {
        "Province": "nap",
        "Unit": {
          "Type": "Army",
          "Nation": "Naples"
        }
      },
      {
        "Province": "mar",
        "Unit": {
          "Type": "Army",
          "Nation": "Marseilles"
        }
      },
      {
        "Province": "bur",
        "Unit": {
          "Type": "Army",
          "Nation": "Spain"
        }
      },
      {
        "Province": "lvp",
        "Unit": {
          "Type": "Army",
          "Nation": "Liverpool"
        }
      },
      {
        "Province": "tun",
        "Unit": {
          "Type": "Army",
          "Nation": "Tunis"
        }
      },
      {
        "Province": "gre",
        "Unit": {
          "Type": "Fleet",
          "Nation": "Greece"
        }
      },
      {
        "Province": "lon",
        "Unit": {
          "Type": "Army",
          "Nation": "London"
        }
      },
      {
        "Province": "rum",
        "Unit": {
          "Type": "Army",
          "Nation": "Sevastopol"
        }
      },
      {
        "Province": "nth",
        "Unit": {
          "Type": "Fleet",
          "Nation": "Norway"
        }
      },
      {
        "Province": "kie",
        "Unit": {
          "Type": "Army",
          "Nation": "Sweden"
        }
      },
      {
        "Province": "bel",
        "Unit": {
          "Type": "Fleet",
          "Nation": "London"
        }
      },
      {
        "Province": "con",
        "Unit": {
          "Type": "Army",
          "Nation": "Smyrna"
        }
      },
      {
        "Province": "war",
        "Unit": {
          "Type": "Army",
          "Nation": "Warsaw"
        }
      },
      {
        "Province": "ven",
        "Unit": {
          "Type": "Army",
          "Nation": "Rome"
        }
      },
      {
        "Province": "boh",
        "Unit": {
          "Type": "Army",
          "Nation": "Munich"
        }
      },
      {
        "Province": "sev",
        "Unit": {
          "Type": "Army",
          "Nation": "Sevastopol"
        }
      },
      {
        "Province": "smy",
        "Unit": {
          "Type": "Fleet",
          "Nation": "Smyrna"
        }
      },
      {
        "Province": "spa/sc",
        "Unit": {
          "Type": "Fleet",
          "Nation": "Spain"
        }
      },
      {
        "Province": "bul",
        "Unit": {
          "Type": "Army",
          "Nation": "Greece"
        }
      },
      {
        "Province": "ser",
        "Unit": {
          "Type": "Army",
          "Nation": "Greece"
        }
      },
      {
        "Province": "vie",
        "Unit": {
          "Type": "Army",
          "Nation": "Budapest"
        }
      },
      {
        "Province": "ber",
        "Unit": {
          "Type": "Army",
          "Nation": "Berlin"
        }
      },
      {
        "Province": "mos",
        "Unit": {
          "Type": "Army",
          "Nation": "Moscow"
        }
      },
      {
        "Province": "edi",
        "Unit": {
          "Type": "Army",
          "Nation": "Edinburgh"
        }
      },
      {
        "Province": "hol",
        "Unit": {
          "Type": "Army",
          "Nation": "Holland"
        }
      },
      {
        "Province": "bal",
        "Unit": {
          "Type": "Fleet",
          "Nation": "Norway"
        }
      },
      {
        "Province": "stp",
        "Unit": {
          "Type": "Army",
          "Nation": "Sweden"
        }
      },
      {
        "Province": "ank",
        "Unit": {
          "Type": "Army",
          "Nation": "Smyrna"
        }
      },
      {
        "Province": "pru",
        "Unit": {
          "Type": "Fleet",
          "Nation": "Sweden"
        }
      }
    ]`
	unitSlice := []interface{}{}
	if err := json.Unmarshal([]byte(unitJSON), &unitSlice); err != nil {
		t.Fatal(err)
	}
	for _, unit := range unitSlice {
		unitMap := unit.(map[string]interface{})
		dataMap := unitMap["Unit"].(map[string]interface{})
		judge.SetUnit(godip.Province(unitMap["Province"].(string)), godip.Unit{godip.UnitType(dataMap["Type"].(string)), godip.Nation(dataMap["Nation"].(string))})
	}
	opts := judge.Phase().Options(judge, Spain)
	tst.AssertOpt(t, opts, []string{"por", "Move", "por", "bre"})
	tst.AssertNoOpt(t, opts, []string{"por", "Move", "por", "mar"})
	opts = judge.Phase().Options(judge, London)
	tst.AssertOpt(t, opts, []string{"mid", "Convoy", "mid", "por", "bre"})
}
