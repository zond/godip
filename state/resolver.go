package state

import (
	"fmt"

	"github.com/zond/godip/common"
)

type resolver struct {
	*State
	deps      []common.Province
	guesses   map[common.Province]error
	resolving map[common.Province]bool
}

func (self *resolver) adjudicate(prov common.Province) (err error) {
	common.Logf("Adj(%v)", prov)
	common.Indent("  ")
	err = self.State.orders[prov].Adjudicate(self)
	common.DeIndent()
	if err == nil {
		common.Logf("%v: Success", prov)
	} else {
		common.Logf("%v: Failure: %v", prov, err)
	}
	return
}

func (self *resolver) Resolve(prov common.Province) (err error) {
	common.Logf("Res(%v) (deps %v)", prov, self.deps)
	common.Indent("  ")
	var ok bool
	if err, ok = self.State.resolutions[prov]; !ok {
		if err, ok = self.guesses[prov]; !ok {
			if self.resolving[prov] {
				common.Logf("Already resolving %v, making negative guess", prov)
				err = fmt.Errorf("Negative guess")
				self.guesses[prov] = err
				self.deps = append(self.deps, prov)
			} else {
				self.resolving[prov] = true
				n_guesses := len(self.guesses)
				err = self.adjudicate(prov)
				delete(self.resolving, prov)
				if _, ok = self.guesses[prov]; ok {
					common.Logf("Guess made for %v, changing guess to positive", prov)
					self.guesses[prov] = nil
					secondErr := self.adjudicate(prov)
					delete(self.guesses, prov)
					if (err == nil) != (secondErr == nil) {
						common.Logf("Calling backup rule with %v", self.deps)
						if err = self.State.backupRule(self, self.deps); err != nil {
							return
						}
						self.deps = nil
						err = self.Resolve(prov)
					} else {
						common.Logf("Only one consistent result, returning %+v", err)
					}
				} else if len(self.guesses) != n_guesses {
					common.Logf("Made new guess, adding %v to deps", prov)
					self.deps = append(self.deps, prov)
				}
			}
		} else {
			common.Logf("Guessed")
		}
		if len(self.guesses) == 0 {
			common.Logf("No guessing, resolving %v", prov)
			self.State.resolutions[prov] = err
		}
	} else {
		common.Logf("Resolved")
	}
	common.DeIndent()
	if err == nil {
		common.Logf("%v: Success (deps %v)", prov, self.deps)
	} else {
		common.Logf("%v: Failure: %v (deps %v)", prov, err, self.deps)
	}
	return
}
