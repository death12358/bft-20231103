package prob_test

import (
	"testing"

	"github.com/adimax2953/bftrtpmodel/prob"

	"github.com/adimax2953/go-tool/randtool"
	"github.com/seehuhn/mt19937"
)

func TestUtil(t *testing.T) {
	// assert := assert.New(t)
	// assert.Equal(int64(0x7fffffffffffffff), prob.Uint32n(9))
	for i := 0; i < 1000; i++ {
		gen := prob.Uint32n(10)
		if gen > 10 {
			println("prob.Uint32n = %v\n", gen)
		}
	}

	var mt19937rand = randtool.New(mt19937.New())
	for i := 0; i < 1000; i++ {
		gen := mt19937rand.Uint32r(0, 10)
		if gen > 10 {
			println("mt19937rand.Uint32r = %v\n", gen)
		}
	}
}
