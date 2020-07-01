package classical

import (
	"testing"

	"github.com/zond/godip"
	"github.com/zond/godip/orders"
)

func TestPhaseMessage(t *testing.T) {
	s, err := ClassicalVariant.Start()
	if err != nil {
		t.Fatal(err)
	}
	checkPhase := func(typ godip.PhaseType) {
		if s.Phase().Type() != typ {
			t.Fatalf("Wanted an %v phase, got %v", typ, s.Phase().Type())
		}
	}
	checkPhase(godip.Movement)
	s.SetOrders(map[godip.Province]godip.Adjudicator{
		"vie": orders.Move("vie", "tyr"),
	})
	if err := s.Next(); err != nil {
		t.Fatal(err)
	}
	checkPhase(godip.Retreat)
	if unit, _, found := s.Unit("tyr"); !found || unit.Nation != godip.Austria {
		t.Fatalf("Wanted Austrian unit in tyr, got %+v", unit)
	}
	if err := s.Next(); err != nil {
		t.Fatal(err)
	}
	checkPhase(godip.Movement)
	s.SetOrders(map[godip.Province]godip.Adjudicator{
		"tyr": orders.Move("tyr", "ven"),
		"tri": orders.SupportMove("tri", "tyr", "ven"),
	})
	if err := s.Next(); err != nil {
		t.Fatal(err)
	}
	checkPhase(godip.Retreat)
	if unit, _, found := s.Unit("ven"); !found || unit.Nation != godip.Austria {
		t.Fatalf("Wanted Austrian unit in ven, got %+v", unit)
	}
	s.SetOrders(map[godip.Province]godip.Adjudicator{
		"ven": orders.Move("ven", "pie"),
	})
	if err := s.Next(); err != nil {
		t.Fatal(err)
	}
	if unit, _, found := s.Unit("pie"); !found || unit.Nation != godip.Italy {
		t.Fatalf("Wanted Italian unit in pie, got %+v", unit)
	}
	checkPhase(godip.Adjustment)
	p := s.Phase()
	if p.Type() != godip.Adjustment {
		t.Fatalf("Wanted an adjustment phase, got %v", p.Type())
	}
	ver := func(nat godip.Nation, want []string) {
		wantMap := map[string]bool{}
		for _, s := range want {
			wantMap[s] = true
		}
		msgs := p.Messages(s, nat)
		if len(msgs) != len(want) {
			t.Errorf("Wanted %v messages for %v, got %v", len(want), nat, len(msgs))
			return
		}
		for _, msg := range msgs {
			if !wantMap[msg] {
				t.Errorf("Found message %v for %v, didn't find it among %+v", msg, nat, want)
			}
		}
	}
	ver(godip.Italy, []string{"MustDisband:1", "OtherMayBuild:Austria:1", "OtherMayBuild:Turkey:0", "OtherMayBuild:England:0", "OtherMayBuild:Russia:0", "OtherMayBuild:Germany:0", "OtherMayBuild:France:0"})
	ver(godip.Austria, []string{"MayBuild:1", "OtherMustDisband:Italy:1", "OtherMayBuild:Turkey:0", "OtherMayBuild:England:0", "OtherMayBuild:Russia:0", "OtherMayBuild:Germany:0", "OtherMayBuild:France:0"})
	ver(godip.Turkey, []string{"MayBuild:0", "OtherMustDisband:Italy:1", "OtherMayBuild:Austria:1", "OtherMayBuild:England:0", "OtherMayBuild:Russia:0", "OtherMayBuild:Germany:0", "OtherMayBuild:France:0"})

	s = Blank(NewPhase(1903, godip.Fall, godip.Adjustment))
	s.SetSC("lon", godip.France)
	s.SetSC("ber", godip.France)
	s.SetSC("mos", godip.France)
	s.SetSC("con", godip.France)
	ver(godip.France, []string{"MayBuild:3", "OtherMayBuild:Austria:0", "OtherMayBuild:England:0", "OtherMayBuild:Germany:0", "OtherMayBuild:Italy:0", "OtherMayBuild:Turkey:0", "OtherMayBuild:Russia:0"})
	ver(godip.Italy, []string{"MayBuild:0", "OtherMayBuild:Austria:0", "OtherMayBuild:England:0", "OtherMayBuild:Germany:0", "OtherMayBuild:France:3", "OtherMayBuild:Turkey:0", "OtherMayBuild:Russia:0"})

	s.SetUnit("par", godip.Unit{godip.Army, godip.France})
	ver(godip.France, []string{"MayBuild:2", "OtherMayBuild:Austria:0", "OtherMayBuild:England:0", "OtherMayBuild:Germany:0", "OtherMayBuild:Italy:0", "OtherMayBuild:Turkey:0", "OtherMayBuild:Russia:0"})
	ver(godip.Italy, []string{"MayBuild:0", "OtherMayBuild:Austria:0", "OtherMayBuild:England:0", "OtherMayBuild:Germany:0", "OtherMayBuild:France:2", "OtherMayBuild:Turkey:0", "OtherMayBuild:Russia:0"})
}
