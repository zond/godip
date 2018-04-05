package state

import (
	"fmt"

	"github.com/zond/godip"
)

type resolver struct {
	*State
	deps      []godip.Province
	guesses   map[godip.Province]error
	resolving map[godip.Province]bool
}

func (self *resolver) adjudicate(prov godip.Province) (err error) {
	godip.Logf("Adj(%v)", prov)
	godip.Indent("  ")
	err = self.State.orders[prov].Adjudicate(self)
	godip.DeIndent()
	if err == nil {
		godip.Logf("%v: Success", prov)
	} else {
		godip.Logf("%v: Failure: %v", prov, err)
	}
	return
}

func (self *resolver) Resolve(prov godip.Province) (err error) {
	godip.Logf("Res(%v) (deps %v)", prov, self.deps)
	godip.Indent("  ")
	var ok bool
	if err, ok = self.State.resolutions[prov]; !ok {
		if err, ok = self.guesses[prov]; !ok {
			if self.resolving[prov] {
				godip.Logf("Already resolving %v, making negative guess", prov)
				err = fmt.Errorf("Negative guess")
				self.guesses[prov] = err
				self.deps = append(self.deps, prov)
			} else {
				self.resolving[prov] = true
				n_guesses := len(self.guesses)
				err = self.adjudicate(prov)
				delete(self.resolving, prov)
				if _, ok = self.guesses[prov]; ok {
					godip.Logf("Guess made for %v, changing guess to positive", prov)
					self.guesses[prov] = nil
					secondErr := self.adjudicate(prov)
					delete(self.guesses, prov)
					if (err == nil) != (secondErr == nil) {
						godip.Logf("Calling backup rule with %v", self.deps)
						if err = self.State.backupRule(self, self.deps); err != nil {
							return
						}
						self.deps = nil
						err = self.Resolve(prov)
					} else {
						godip.Logf("Only one consistent result, returning %+v", err)
					}
				} else if len(self.guesses) != n_guesses {
					godip.Logf("Made new guess, adding %v to deps", prov)
					self.deps = append(self.deps, prov)
				}
			}
		} else {
			godip.Logf("Guessed")
		}
		if len(self.guesses) == 0 {
			godip.Logf("No guessing, resolving %v", prov)
			self.State.resolutions[prov] = err
		}
	} else {
		godip.Logf("Resolved")
	}
	godip.DeIndent()
	if err == nil {
		godip.Logf("%v: Success (deps %v)", prov, self.deps)
	} else {
		godip.Logf("%v: Failure: %v (deps %v)", prov, err, self.deps)
	}
	return
}
