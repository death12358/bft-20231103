package random_test

import (
	"testing"

	"github.com/adimax2953/bftrtpmodel/bft302prob/prob302/random"
)

// func TestGenerate_random_numbers(t *testing.T) {
// 	prob302.Test_generate_random_numbers()
// }

func TestGetRandom(t *testing.T) {
	random.Test_GetRandom(10)
}

func TestRandomFloat64(t *testing.T) {
	random.Test_RandomFloat64()
}

// func TestGenRandArray(t *testing.T) {
// 	random.Test_GenRandArray([]int32{1, 2, 3, 4, 5, 6}, 3)
// }

func TestShuffle(t *testing.T) {
	random.Test_Shuffle([]int{1, 2, 3, 4})
}
func TestMT19937(t *testing.T) {
	random.Test_mt19937()
}
