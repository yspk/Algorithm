package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	"math/big"
	"github.com/ethereum/go-ethereum/common/math"
	"strconv"
)

var (
	expDiffPeriod = big.NewInt(100000)
	big1          = big.NewInt(1)
	big2          = big.NewInt(2)
	big9          = big.NewInt(9)
	big10         = big.NewInt(10)
	big12         = big.NewInt(12)
	bigMinus99    = big.NewInt(-99)
	big2999999    = big.NewInt(2999999)
)

// calcDifficultyByzantium is the difficulty adjustment algorithm. It returns
// the difficulty that a new block should have when created at time given the
// parent block's time and difficulty. The calculation uses the Byzantium rules.
//func calcDifficultyByzantium2(time uint64, parent *types.Header) *big.Int {
//	// https://github.com/ethereum/EIPs/issues/100.
//	// algorithm:
//	// diff = (parent_diff +
//	//         (parent_diff / 2048 * max((2 if len(parent.uncles) else 1) - ((timestamp - parent.timestamp) // 9), -99))
//	//        ) + 2^(periodCount - 2)
//
//	bigTime := new(big.Int).SetUint64(time)
//	bigParentTime := new(big.Int).Set(parent.Time)
//
//	// holds intermediate values to make the algo easier to read & audit
//	x := new(big.Int)
//	y := new(big.Int)
//
//	// (2 if len(parent_uncles) else 1) - (block_timestamp - parent_timestamp) // 9
//	x.Sub(bigTime, bigParentTime)
//	x.Div(x, big9)
//	if parent.UncleHash == types.EmptyUncleHash {
//		x.Sub(big1, x)
//	} else {
//		x.Sub(big2, x)
//	}
//	// max((2 if len(parent_uncles) else 1) - (block_timestamp - parent_timestamp) // 9, -99)
//	if x.Cmp(bigMinus99) < 0 {
//		x.Set(bigMinus99)
//	}
//	// parent_diff + (parent_diff / 2048 * max((2 if len(parent.uncles) else 1) - ((timestamp - parent.timestamp) // 9), -99))
//	y.Div(parent.Difficulty, params.DifficultyBoundDivisor)
//	x.Mul(y, x)
//	x.Add(parent.Difficulty, x)
//
//	// minimum difficulty can ever be (before exponential factor)
//	if x.Cmp(params.MinimumDifficulty) < 0 {
//		x.Set(params.MinimumDifficulty)
//	}
//	// calculate a fake block number for the ice-age delay:
//	// https://github.com/ethereum/EIPs/pull/669
//	// fake_block_number = max(0, block.number - 3_000_000)
//	fakeBlockNumber := new(big.Int)
//	if parent.Number.Cmp(big2999999) >= 0 {
//		fakeBlockNumber = fakeBlockNumber.Sub(parent.Number, big2999999) // Note, parent is 1 less than the actual block number
//	}
//	// for the exponential factor
//	periodCount := fakeBlockNumber
//	periodCount.Div(periodCount, expDiffPeriod)
//
//	// the exponential factor, commonly referred to as "the bomb"
//	// diff = diff + 2^(periodCount - 2)
//	if periodCount.Cmp(big1) > 0 {
//		y.Sub(periodCount, big2)
//		y.Exp(big2, y, nil)
//		x.Add(x, y)
//	}
//	return x
//}

func calcDifficultySimpleChain(time uint64, parent *types.Header) *big.Int {
	// diff = parent_diff
	//            + parent_diff * 5 / 100000000 * ( MIN ( timestamp - parent.timestamp , 900 ) ) ^ 2
	//            - parent_diff * 1 / 10000 * ( MIN ( timestamp - parent.timestamp , 900 ) )
	//            + parent_diff * 1/1000
	x := big.NewInt(0)

	big5 := big.NewInt(5)
	big900 := big.NewInt(900)
	big1000 := big.NewInt(1000)
	big10000 := big.NewInt(10000)
	big100000000 := big.NewInt(100000000)
	bigTime := new(big.Int).SetUint64(time)

	yn := big.NewInt(0)
	y1 := big.NewInt(0)
	y2 := big.NewInt(0)
	y3 := big.NewInt(0)

	x.Sub(bigTime, parent.Time)

	timeDiff := x

	if timeDiff.Cmp(big900) > 0 {
		timeDiff.Set(big900)
	}

	y1.Mul(timeDiff, timeDiff)
	y1.Mul(y1, parent.Difficulty)
	y1.Mul(y1, big5)
	y1.Div(y1, big100000000)

	//y1.Mul(y1,timeDiff)
	//y1.Div(y1.Mul( parent.Difficulty , big5 ),big100000000)
	//y1.Mul(y1,y1)

	y2.Mul(parent.Difficulty, timeDiff)
	y2.Div(y2, big10000)

	y3.Div(parent.Difficulty, big1000)

	yn.Add(yn, parent.Difficulty)
	yn.Add(yn, y1)
	yn.Sub(yn, y2)
	yn.Add(yn, y3)

	return yn
}

func main() {
	h := &types.Header{}
	//h.Time = big.NewInt(1508131362)
	h.Time,_ =  hexutil.DecodeBig("0x59e44222")
	//h.UncleHash
	//h.Number = big.NewInt(4370001)
	h.Number,_ = hexutil.DecodeBig("0x42ae51")
	//h.Difficulty = big.NewInt(2991422903560207)
	h.Difficulty,_ = hexutil.DecodeBig("0xaa0aeeb8a280f")
	h.UncleHash = common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347")

	time,_ := strconv.ParseUint("59e44227",16,64)
	x := calcDifficultyFrontier(time, h)
	fmt.Println("calcDifficultyFrontier", hexutil.EncodeBig(x))

	x = calcDifficultyByzantium(time, h)
	fmt.Println("calcDifficultyByzantium", hexutil.EncodeBig(x))

	x = calcDifficultyHomestead(time, h)
	fmt.Println("calcDifficultyHomestead", hexutil.EncodeBig(x))

	//x = calcDifficultyByzantium2(1508131367, h)
	//fmt.Println("calcDifficultyByzantium2", x)

	//x = calcDifficultySimpleChain(1542866839, h)
	//fmt.Println("calcDifficultySimpleChain", hexutil.EncodeBig(x))
}

func calcDifficultyByzantium(time uint64, parent *types.Header) *big.Int {
	// https://github.com/ethereum/EIPs/issues/100.
	// algorithm:
	// diff = (parent_diff +
	//         (parent_diff / 2048 * max((2 if len(parent.uncles) else 1) - ((timestamp - parent.timestamp) // 9), -99))
	//        ) + 2^(periodCount - 2)

	bigTime := new(big.Int).SetUint64(time)
	bigParentTime := new(big.Int).Set(parent.Time)

	// holds intermediate values to make the algo easier to read & audit
	x := new(big.Int)
	y := new(big.Int)

	// (2 if len(parent_uncles) else 1) - (block_timestamp - parent_timestamp) // 9
	x.Sub(bigTime, bigParentTime)
	x.Div(x, big9)
	if parent.UncleHash == types.EmptyUncleHash {
		x.Sub(big1, x)
	} else {
		x.Sub(big2, x)
	}
	// max((2 if len(parent_uncles) else 1) - (block_timestamp - parent_timestamp) // 9, -99)
	if x.Cmp(bigMinus99) < 0 {
		x.Set(bigMinus99)
	}
	// parent_diff + (parent_diff / 2048 * max((2 if len(parent.uncles) else 1) - ((timestamp - parent.timestamp) // 9), -99))
	y.Div(parent.Difficulty, params.DifficultyBoundDivisor)
	x.Mul(y, x)
	x.Add(parent.Difficulty, x)

	// minimum difficulty can ever be (before exponential factor)
	if x.Cmp(params.MinimumDifficulty) < 0 {
		x.Set(params.MinimumDifficulty)
	}
	// calculate a fake block number for the ice-age delay:
	// https://github.com/ethereum/EIPs/pull/669
	// fake_block_number = max(0, block.number - 3_000_000)
	fakeBlockNumber := new(big.Int)
	if parent.Number.Cmp(big2999999) >= 0 {
		fakeBlockNumber = fakeBlockNumber.Sub(parent.Number, big2999999) // Note, parent is 1 less than the actual block number
	}
	// for the exponential factor
	periodCount := fakeBlockNumber
	periodCount.Div(periodCount, expDiffPeriod)

	// the exponential factor, commonly referred to as "the bomb"
	// diff = diff + 2^(periodCount - 2)
	if periodCount.Cmp(big1) > 0 {
		y.Sub(periodCount, big2)
		y.Exp(big2, y, nil)
		x.Add(x, y)
	}
	return x
}

// calcDifficultyHomestead is the difficulty adjustment algorithm. It returns
// the difficulty that a new block should have when created at time given the
// parent block's time and difficulty. The calculation uses the Homestead rules.
func calcDifficultyHomestead(time uint64, parent *types.Header) *big.Int {
	// https://github.com/ethereum/EIPs/blob/master/EIPS/eip-2.md
	// algorithm:
	// diff = (parent_diff +
	//         (parent_diff / 2048 * max(1 - (block_timestamp - parent_timestamp) // 10, -99))
	//        ) + 2^(periodCount - 2)

	bigTime := new(big.Int).SetUint64(time)
	bigParentTime := new(big.Int).Set(parent.Time)

	// holds intermediate values to make the algo easier to read & audit
	x := new(big.Int)
	y := new(big.Int)

	// 1 - (block_timestamp - parent_timestamp) // 10
	x.Sub(bigTime, bigParentTime)
	x.Div(x, big10)
	x.Sub(big1, x)

	// max(1 - (block_timestamp - parent_timestamp) // 10, -99)
	if x.Cmp(bigMinus99) < 0 {
		x.Set(bigMinus99)
	}
	// (parent_diff + parent_diff // 2048 * max(1 - (block_timestamp - parent_timestamp) // 10, -99))
	y.Div(parent.Difficulty, params.DifficultyBoundDivisor)
	x.Mul(y, x)
	x.Add(parent.Difficulty, x)

	// minimum difficulty can ever be (before exponential factor)
	if x.Cmp(params.MinimumDifficulty) < 0 {
		x.Set(params.MinimumDifficulty)
	}
	// for the exponential factor
	periodCount := new(big.Int).Add(parent.Number, big1)
	periodCount.Div(periodCount, expDiffPeriod)

	// the exponential factor, commonly referred to as "the bomb"
	// diff = diff + 2^(periodCount - 2)
	if periodCount.Cmp(big1) > 0 {
		y.Sub(periodCount, big2)
		y.Exp(big2, y, nil)
		x.Add(x, y)
	}
	return x
}

// calcDifficultyFrontier is the difficulty adjustment algorithm. It returns the
// difficulty that a new block should have when created at time given the parent
// block's time and difficulty. The calculation uses the Frontier rules.
func calcDifficultyFrontier(time uint64, parent *types.Header) *big.Int {
	diff := new(big.Int)
	adjust := new(big.Int).Div(parent.Difficulty, params.DifficultyBoundDivisor)
	bigTime := new(big.Int)
	bigParentTime := new(big.Int)

	bigTime.SetUint64(time)
	bigParentTime.Set(parent.Time)

	if bigTime.Sub(bigTime, bigParentTime).Cmp(params.DurationLimit) < 0 {
		diff.Add(parent.Difficulty, adjust)
	} else {
		diff.Sub(parent.Difficulty, adjust)
	}
	if diff.Cmp(params.MinimumDifficulty) < 0 {
		diff.Set(params.MinimumDifficulty)
	}

	periodCount := new(big.Int).Add(parent.Number, big1)
	periodCount.Div(periodCount, expDiffPeriod)
	if periodCount.Cmp(big1) > 0 {
		// diff = diff + 2^(periodCount - 2)
		expDiff := periodCount.Sub(periodCount, big2)
		expDiff.Exp(big2, expDiff, nil)
		diff.Add(diff, expDiff)
		diff = math.BigMax(diff, params.MinimumDifficulty)
	}
	return diff
}