package quorum

import (
	"reflect"
	"testing"

	"github.com/davecgh/go-spew/spew"

	"github.com/simcord-llc/bitbon-system-blockchain/common"
)

func TestGenerateSequence(t *testing.T) {
	type testData struct {
		hash   common.Hash
		size   int
		wanted []uint64
	}

	var data []testData = []testData{
		{
			hash:   common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
			size:   2,
			wanted: []uint64{0, 1},
		},
		{
			hash:   common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000001"),
			size:   2,
			wanted: []uint64{1, 0},
		},
		{
			hash:   common.HexToHash("0x95f1cc9a7334691121538b1b7ae9db19e542c942b465641918cd100b6a9e47f6"),
			size:   40,
			wanted: []uint64{0, 5, 34, 20, 38, 25, 29, 7, 8, 9, 10, 35, 16, 13, 14, 15, 12, 24, 18, 2, 27, 19, 22, 23, 17, 6, 21, 3, 28, 33, 30, 1, 32, 31, 26, 11, 36, 37, 4, 39},
		},
	}

	for _, d := range data {
		res := generateSequence(d.hash, d.size)
		spew.Dump(res)

		if !reflect.DeepEqual(res, d.wanted) {
			t.Errorf("Got unexpected result. Have %v, wanted %v", res, d.wanted)
		}
	}
}

var result []uint64

func BenchmarkGenerateSequence(b *testing.B) {
	hash := common.HexToHash("0x95f1cc9a7334691121538b1b7ae9db19e542c942b465641918cd100b6a9e47f6")

	for i := 0; i < b.N; i++ {
		result = generateSequence(hash, i)
	}
}

func BenchmarkShuffle2(b *testing.B) {
	hash := common.HexToHash("0x95f1cc9a7334691121538b1b7ae9db19e542c942b465641918cd100b6a9e47f6")

	for i := 0; i < b.N; i++ {
		slice := generateIndexes(i)
		result = shuffle2(hash, slice)
	}
}

func BenchmarkShuffle3(b *testing.B) {
	hash := common.HexToHash("0x95f1cc9a7334691121538b1b7ae9db19e542c942b465641918cd100b6a9e47f6")

	for i := 0; i < b.N; i++ {
		slice := generateIndexes(i)
		result = shuffle3(hash, slice)
	}
}
