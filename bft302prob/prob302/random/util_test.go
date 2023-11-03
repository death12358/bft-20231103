package random_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/adimax2953/go-tool/randtool"
	"github.com/seehuhn/mt19937"
)

func TestUtil(t *testing.T) {
	// assert := assert.New(t)
	// assert.Equal(int64(0x7fffffffffffffff), prob.Uintn(9))
	for i := 0; i < 1000; i++ {
		rand.Seed(time.Now().UnixNano())

		gen := rand.Intn(10)
		if gen > 10 {
			println("prob.Uintn = %v\n", gen)
		}
	}

	var mt19937rand = randtool.New(mt19937.New())
	for i := 0; i < 1000; i++ {
		gen := mt19937rand.Uint32r(0, 10)
		if gen > 10 {
			println("mt19937rand.Uintr = %v\n", gen)
		}
	}
}
