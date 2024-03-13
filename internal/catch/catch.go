package catch

import (
	"errors"
	"math/rand"
	"time"
)

type CatchAlgorithm int

const (
	BOOTDEV CatchAlgorithm = 0
	GEN1    CatchAlgorithm = 1
)

type CatchResult struct {
	wobble1 bool
	wobble2 bool
	wobble3 bool
}

func (c *CatchResult) IsFullCatch() bool {
	return c.wobble1 && c.wobble2 && c.wobble3
}

func (c *CatchResult) Wobbles() (bool, bool, bool) {
	return c.wobble1, c.wobble2, c.wobble3
}

type Catcher interface {
	CanCatch(catchRate int) *CatchResult
}

func NewCatcher(alg CatchAlgorithm) (Catcher, error) {
	switch alg {
	case BOOTDEV:
		{
			c := BootDevCatcher{}
			return &c, nil
		}
	}

	return nil, errors.New("not implemented")
}

type BootDevCatcher struct{}

func (b *BootDevCatcher) CanCatch(catchRate int) *CatchResult {
	catchResult := CatchResult{}

	if catchRate < getRandValue255() {
		return &catchResult
	}
	catchResult.wobble1 = true
	if catchRate < getRandValue255() {
		return &catchResult
	}
	catchResult.wobble2 = true

	if catchRate < getRandValue255() {
		return &catchResult
	}
	catchResult.wobble3 = true

	return &catchResult
}

func getRandValue255() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(255)
}
