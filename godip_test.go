package godip

import (
	"encoding/json"
	"reflect"
	"sort"
	"testing"
)

func TestOptionsMarshal(t *testing.T) {
	opts := Options{
		FilteredOptionValue{
			Filter: "MAX:Build:1",
			Value:  Province("spa"),
		}: Options{
			OrderType("Build"): Options{
				UnitType("Army"): Options{
					SrcProvince("spa"): Options{},
				},
			},
		},
	}
	b, err := json.MarshalIndent(opts, "  ", "  ")
	if err != nil {
		t.Fatal(err)
	}
	wanted := `{
    "spa": {
      "Filter": "MAX:Build:1",
      "Next": {
        "Build": {
          "Next": {
            "Army": {
              "Next": {
                "spa": {
                  "Next": {},
                  "Type": "SrcProvince"
                }
              },
              "Type": "UnitType"
            }
          },
          "Type": "OrderType"
        }
      },
      "Type": "Province"
    }
  }`
	if string(b) != wanted {
		t.Errorf("Bad json marshal, got %v, wanted %v", string(b), wanted)
	}
}

func TestOptionsBubbleFilters(t *testing.T) {
	for _, tc := range []struct {
		opts    Options
		bubbled Options
	}{
		{
			opts: Options{
				Province("spa"): Options{
					OrderType("Build"): Options{
						FilteredOptionValue{
							Filter: "MAX:Build:1",
							Value:  UnitType("Army"),
						}: Options{
							SrcProvince("spa"): Options{},
						},
					},
				},
			},
			bubbled: Options{
				FilteredOptionValue{
					Filter: "MAX:Build:1",
					Value:  Province("spa"),
				}: Options{
					OrderType("Build"): Options{
						UnitType("Army"): Options{
							SrcProvince("spa"): Options{},
						},
					},
				},
			},
		},
		{
			opts: Options{
				Province("lon"): Options{
					OrderType("Build"): Options{
						FilteredOptionValue{
							Filter: "MAX:Build:1",
							Value:  UnitType("Army"),
						}: Options{
							SrcProvince("lon"): Options{},
						},
					},
				},
				Province("spa"): Options{
					OrderType("Build"): Options{
						FilteredOptionValue{
							Filter: "MAX:Build:1",
							Value:  UnitType("Army"),
						}: Options{
							SrcProvince("spa"): Options{},
						},
					},
				},
			},
			bubbled: Options{
				FilteredOptionValue{
					Filter: "MAX:Build:1",
					Value:  Province("spa"),
				}: Options{
					OrderType("Build"): Options{
						UnitType("Army"): Options{
							SrcProvince("spa"): Options{},
						},
					},
				},
				FilteredOptionValue{
					Filter: "MAX:Build:1",
					Value:  Province("lon"),
				}: Options{
					OrderType("Build"): Options{
						UnitType("Army"): Options{
							SrcProvince("lon"): Options{},
						},
					},
				},
			},
		},
		{
			opts: Options{
				Province("spa"): Options{
					OrderType("Build"): Options{
						FilteredOptionValue{
							Filter: "MAX:Build:1",
							Value:  UnitType("Army"),
						}: Options{
							SrcProvince("spa"): Options{},
						},
					},
					OrderType("Disband"): Options{
						FilteredOptionValue{
							Filter: "MAX:Disband:1",
							Value:  UnitType("Army"),
						}: Options{
							SrcProvince("spa"): Options{},
						},
					},
				},
			},
			bubbled: Options{
				Province("spa"): Options{
					OrderType("Build"): Options{
						FilteredOptionValue{
							Filter: "MAX:Build:1",
							Value:  UnitType("Army"),
						}: Options{
							SrcProvince("spa"): Options{},
						},
					},
					OrderType("Disband"): Options{
						FilteredOptionValue{
							Filter: "MAX:Disband:1",
							Value:  UnitType("Army"),
						}: Options{
							SrcProvince("spa"): Options{},
						},
					},
				},
			},
		},
		{
			opts: Options{
				Province("spa"): Options{
					OrderType("Build"): Options{
						FilteredOptionValue{
							Filter: "MAX:Build:1",
							Value:  UnitType("Army"),
						}: Options{
							SrcProvince("spa"): Options{},
						},
						FilteredOptionValue{
							Filter: "MAX:Build:1",
							Value:  UnitType("Fleet"),
						}: Options{
							SrcProvince("spa"): Options{},
						},
					},
				},
			},
			bubbled: Options{
				FilteredOptionValue{
					Filter: "MAX:Build:1",
					Value:  Province("spa"),
				}: Options{
					OrderType("Build"): Options{
						UnitType("Army"): Options{
							SrcProvince("spa"): Options{},
						},
						UnitType("Fleet"): Options{
							SrcProvince("spa"): Options{},
						},
					},
				},
			},
		},
		{
			opts: Options{
				Province("spa"): Options{
					OrderType("Build"): Options{
						FilteredOptionValue{
							Filter: "MAX:Build:1",
							Value:  UnitType("Army"),
						}: Options{
							SrcProvince("spa"): Options{},
						},
						UnitType("Fleet"): Options{
							SrcProvince("spa"): Options{},
						},
					},
				},
			},
			bubbled: Options{
				Province("spa"): Options{
					OrderType("Build"): Options{
						FilteredOptionValue{
							Filter: "MAX:Build:1",
							Value:  UnitType("Army"),
						}: Options{
							SrcProvince("spa"): Options{},
						},
						UnitType("Fleet"): Options{
							SrcProvince("spa"): Options{},
						},
					},
				},
			},
		},
	} {
		bubbled := tc.opts.BubbleFilters()
		if !reflect.DeepEqual(bubbled, tc.bubbled) {
			t.Errorf("Got %+v, wanted %+v", bubbled, tc.bubbled)
		}
	}
}

func TestNationSorting(t *testing.T) {
	nations := Nations{
		Austria,
		Russia,
		England,
	}
	sort.Sort(nations)
	if nations[0] != Austria {
		t.Errorf("Wanted Austria, got %v", nations[0])
	}
	if nations[1] != England {
		t.Errorf("Wanted England, got %v", nations[1])
	}
	if nations[2] != Russia {
		t.Errorf("Wanted Russia, got %v", nations[2])
	}
}

func TestNationEqual(t *testing.T) {
	n := Nations{
		Austria,
		Russia,
		England,
	}
	o := Nations{
		Austria,
		Russia,
		England,
	}
	if !n.Equal(o) {
		t.Errorf("Wanted equal, wasn't: %+v, %+v", n, o)
	}
	o = Nations{
		Russia,
		Austria,
		England,
	}
	if !n.Equal(o) {
		t.Errorf("Wanted equal, wasn't: %+v, %+v", n, o)
	}
	if n[0] != Austria || n[1] != Russia || n[2] != England || o[0] != Russia || o[1] != Austria || o[2] != England {
		t.Errorf("Changed order during equal, not OK")
	}
	o = Nations{
		Austria,
		Turkey,
		England,
	}
	if n.Equal(o) {
		t.Errorf("Wanted inequal, was: %+v, %+v", n, o)
	}
	o = Nations{
		Austria,
		Russia,
	}
	if n.Equal(o) {
		t.Errorf("Wanted inequal, was: %+v, %+v", n, o)
	}
}
