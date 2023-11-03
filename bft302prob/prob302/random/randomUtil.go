package random

import (
	crand "crypto/rand"
	"fmt"
	"math/big"
	"math/rand"

	"os"
	"sync"
	"time"

	"github.com/adimax2953/go-tool/randtool"
	"github.com/seehuhn/mt19937"
)

var rngPool sync.Pool

var mt19937rand = randtool.New(mt19937.New())

func init() {
	b := new(big.Int).SetUint64(uint64(time.Now().UTC().UnixNano() / int64(os.Getpid())))
	sd, _ := crand.Int(crand.Reader, b)
	x := sd.Uint64() + 0x9E3779B97F4A7C15
	x ^= x >> 30 * 0xBF58476D1CE4E5B9
	x ^= x << 27 * 0x94D049BB133111EB
	x ^= x >> 31
	seed := int64(x)

	mt19937rand.Seed(seed)

}

// Uint32 - returns pseudorandom uint32.
func Uint32() uint32 {
	// v := rngPool.Get()
	// if v == nil {
	// 	v = &RNG{}
	// }
	// r := v.(*RNG)
	// x := r.Uint32()
	// rngPool.Put(r)
	// return x
	return mt19937rand.Uint32()
	// return rand.Uint32()
}

// It is safe calling this function from concurrent goroutines.
func RandomInt() int {
	// v := rngPool.Get()
	// if v == nil {
	// 	v = &RNG{}
	// }
	// r := v.(*RNG)
	// x := r.RandomInt32()
	// rngPool.Put(r)
	// return x

	return mt19937rand.Int()
	// return rand.RandomInt32()
}

func RandomFloat64() float64 {
	// v := rngPool.Get()
	// if v == nil {
	// 	v = &RNG{}
	// }
	// r := v.(*RNG)
	// x := r.RandomInt32()
	// rngPool.Put(r)
	// return x

	return mt19937rand.Float64()
	// return rand.RandomInt32()
}

// 隨機產生[0,N-1]間的一個正整數
func GetRandom(maxN int32) int32 {
	x := Uint32()
	// See http://lemire.me/blog/2016/06/27/a-fast-alternative-to-the-modulo-reduction/
	return int32((uint64(x) * uint64(maxN)) >> 32)
}

// // 隨機產生[0,N-1]間的一個正整數
// func GetRandom(maxN int) int {
// 	x := RandomInt()
// 	// See http://lemire.me/blog/2016/06/27/a-fast-alternative-to-the-modulo-reduction/
// 	return int(int32((uint32(x) * uint32(maxN)) >> 16))
// }

// GenRandArray - safe
func GenRandArray(weightArray []int32, arraySizze int32) int32 {
	var resultNum int32
	var sumWeight int32
	var sumArray []int32
	sumArray = make([]int32, arraySizze)
	var i int32

	for i = 0; i < arraySizze; i++ {
		sumWeight += weightArray[i]
		sumArray[i] = sumWeight
	}

	var randNum int32
	randNum = GetRandom(sumWeight)

	for i = 0; i < arraySizze; i++ {
		if randNum < sumArray[i] {
			resultNum = i
			break
		}
	}

	return resultNum
}

// RNG is a pseudorandom number generator.
//
// It is unsafe to call RNG methods from concurrent goroutines.
// type RNG struct {
// 	state uint32
// }

// // GetRandom -
// func (r *RNG) GetRandom(maxN int32) int32 {
// 	x := r.Uint32()
// 	// See http://lemire.me/blog/2016/06/27/a-fast-alternative-to-the-modulo-reduction/
// 	return int32((uint64(x) * uint64(maxN)) >> 32)
// }

// // GenRandArray -
// func (r *RNG) GenRandArray(weightArray []int32, arraySizze int32) uint32 {
// 	var resultNum uint32
// 	var sumWeight uint32
// 	var sumArray []uint32
// 	sumArray = make([]uint32, arraySizze)
// 	var i int32

// 	for i = 0; i < arraySizze; i++ {
// 		sumWeight += uint32(weightArray[i])
// 		sumArray[i] = sumWeight
// 	}

// 	var randNum uint32
// 	randNum = r.Uint32n(sumWeight)

// 	for i = 0; i < arraySizze; i++ {
// 		if randNum < sumArray[i] {
// 			resultNum = uint32(i)
// 			break
// 		}
// 	}

// 	return resultNum
// }

// // Uint32 returns pseudorandom uint32.
// //
// // It is unsafe to call this method from concurrent goroutines.
// func (r *RNG) Uint32() uint32 {
// 	if r.state == 0 {
// 		r.state = getRandomUint()
// 	}

// 	// See https://en.wikipedia.org/wiki/Xorshift
// 	x := r.state
// 	x ^= x << 13
// 	x ^= x >> 17
// 	x ^= x << 5
// 	r.state = x
// 	return x
// }

// // Uint32n returns pseudorandom uint32 in the range [0..maxN).
// //
// // It is unsafe to call this method from concurrent goroutines.
// func (r *RNG) Uint32n(maxN uint32) uint32 {
// 	x := r.Uint32()
// 	// See http://lemire.me/blog/2016/06/27/a-fast-alternative-to-the-modulo-reduction/
// 	return uint32((uint64(x) * uint64(maxN)) >> 32)
// }

func getRandomUint() int {
	// x := time.Now().UnixNano()
	x := uint64(time.Now().UTC().UnixNano() / int64(os.Getpid()))
	x ^= x << 13
	x ^= x >> 7
	x ^= x << 17
	return int((x >> 32) ^ x)
}

// Shuffle -打亂陣列
func Shuffle(nums []int) []int {
	for i := len(nums); i > 0; i-- {
		last := i - 1
		idx := rand.Intn(i)
		nums[last], nums[idx] = nums[idx], nums[last]
	}
	return nums
}

// BossRange -
var BossRange = make([]int, 100*10000)

// ShufflebossRange -
func ShufflebossRange() {

	for i := 0; i < 100*10000; i++ {
		BossRange[i] = int(i)
	}

	BossRange = Shuffle(BossRange)
}
func Test_generate_random_numbers() {
	fmt.Printf("RandomInt():%v\n", RandomInt())
	fmt.Printf("RandomFloat64() %v\n", RandomFloat64())
	fmt.Printf("GetRandom(10):%v\n", GetRandom(10))
	fmt.Printf("GetRandom(10):%v\n", GetRandom(10))
	fmt.Printf("GenRandArray([]int32{1, 2, 3}, 3):%v\n", GenRandArray([]int32{1, 2, 3}, 3))
	// r := RNG{2}
	// fmt.Printf("r.Uint32():%v\n", r.Uint32())
	// fmt.Printf("r.Uint32n(10):%v\n", r.Uint32n(10))

	fmt.Printf("getRandomUint32():%v\n", getRandomUint())

	fmt.Printf("Shuffle(nums []int{1,2,3,4}:%v\n", Shuffle([]int{1, 2, 3, 4}))
}

func Test_RandomFloat64() {
	distribution := make(map[int]int)
	for i := 0; i < 10000; i++ {
		for a := 1; a <= 10; a++ {
			if RandomFloat64() < 0.1*float64(a) {
				distribution[a]++
			}
		}
	}
	fmt.Println(distribution)
}

func Test_GetRandom(n int32) {
	distribution := make(map[int32]int)
	for i := 0; i < 100000000; i++ {
		distribution[GetRandom(n)]++
	}
	fmt.Println(distribution)
}

func Test_GenRandArray(weightArray []int32, size int32) {
	distribution := make(map[int32]int)
	for i := 0; i < 10000000; i++ {
		distribution[GenRandArray(weightArray, size)]++
	}
	fmt.Println(distribution)
}

//var rng = RNG{state: 2}

func Test_Shuffle(originalSlice []int) {
	distribution := make(map[string]int)
	for i := 0; i < 10000000; i++ {
		testSlice := make([]int, len(originalSlice))
		copy(testSlice, originalSlice)
		Shuffle(testSlice)
		resultStr := fmt.Sprint(testSlice)
		distribution[resultStr]++
	}
	fmt.Println(distribution)
}
func Test_mt19937() {
	fmt.Printf("mt19937rand:%v\n", mt19937rand)

	fmt.Printf("mt19937rand:%v\n", mt19937rand)
	fmt.Printf("mt19937rand:%v\n", mt19937rand)
	fmt.Printf("mt19937rand:%v\n", mt19937rand)
	fmt.Printf("mt19937rand:%v\n", mt19937rand)
}
