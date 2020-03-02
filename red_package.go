package iutils

import (
	"time"
	"math/rand"
)

// 随机红包（二倍均值法）
// remainCount: 剩余红包数
// remainMoney: 剩余红包金额（单位：分)
func RandomMoney(remainCount, remainMoney int) int {
	if remainCount == 1 {
		return remainMoney
	}

	rand.Seed(time.Now().UnixNano())

	var min = 1
	max := remainMoney / remainCount * 2
	money := rand.Intn(max) + min
	return money
}
