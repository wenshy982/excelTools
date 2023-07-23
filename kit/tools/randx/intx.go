package randx

import (
	cRand "crypto/rand"
	"math/big"
)

var IntMap = map[int][]int64{
	0: {0, 0},
	1: {8, 1},
	2: {89, 10},
	3: {899, 100},
	4: {8999, 1000},
	5: {89999, 10000},
	6: {899999, 100000},
	7: {8999999, 1000000},
	8: {89999999, 10000000},
}

// Int 生成x位随机数
func Int(x int) string {
	if x <= 0 {
		x = 0
	}
	if x > 8 {
		x = 8
	}
	r := IntMap[x]
	n, _ := cRand.Int(cRand.Reader, big.NewInt(r[0])) // 生成x位随机数 0 - r[0]
	n.Add(n, big.NewInt(r[1]))                        // 加 r[1] 保证随机数为x位
	return n.String()
}
