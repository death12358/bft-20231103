package spin_test

import (
	"fmt"
	"testing"

	"github.com/adimax2953/bftrtpmodel/bft302prob/prob302/random"
	"github.com/adimax2953/bftrtpmodel/bft302prob/prob302/spin"
	"github.com/adimax2953/bftrtpmodel/bft302prob/prob302/tables"
)

func init() {
	tables.TableInit()
}

var TestSpinIn = spin.SpinIn{
	RTP:           97,
	RTPflow:       5,
	TotalBet:      15,
	HitFishList:   [1]int32{13},
	MultipleLimit: 99999999999999999,
	FreeGameTimes: 0,
}

func TestFGSpin(t *testing.T) {
	for i := 0; i < 1; i++ {
		spinOut := TestSpinIn.NGSpinCalc()
		fgTimes := spinOut.FreeGameTimes
		for fgTimes > 0 {
			TestSpinIn.FreeGameTimes = fgTimes
			spinOut = TestSpinIn.FGSpinCalc()
			fgTimes = spinOut.FreeGameTimes
		}
	}
}
func TestNGSpin(t *testing.T) {
	// js, err := json.Marshal(tables.FishDeadProb)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Printf("js_FishDeadProb:%+v\n", string(js))

	// js, err = json.Marshal(tables.FishPayTable)
	// if err != nil {
	// 	fmt.Println(err)

	// }
	// fmt.Printf("FishPayTable:%+v", string(js))
	times := 0
	for i := 0; i < 1000000; i++ {
		spinOut := TestSpinIn.NGSpinCalc()
		if spinOut.Odds[0] != 0 {
			// js, err := json.Marshal(TestSpinIn.NGSpinCalc())
			// if err != nil {
			// 	fmt.Printf("err:%v\n", err)
			// }
			// fmt.Printf("%v\n", string(js))
			times++
		}
	}
	fmt.Println(times)
}

func TestNewSpinInOut(t *testing.T) {
	spin.Test_NewSpinInOut()
}

func TestGenRandArray(t *testing.T) {
	weightArray := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var arraySizze int32 = 10
	fmt.Printf("%v", random.GenRandArray(weightArray, arraySizze))
}

// func TestWriteMathData(t *testing.T) {
// 	prob302.WriteMathData("97")
// }

// func TestPlotBalance(t *testing.T) {
// 	prob302.PlotBalance("97")
// }
