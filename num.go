package iutils

import (
	"strconv"
	"strings"
	"math"
	"fmt"
)

var tenToAny = map[int]string{
	0:  "0",
	1:  "1",
	2:  "2",
	3:  "3",
	4:  "4",
	5:  "5",
	6:  "6",
	7:  "7",
	8:  "8",
	9:  "9",
	10: "a",
	11: "b",
	12: "c",
	13: "d",
	14: "e",
	15: "f",
	16: "g",
	17: "h",
	18: "i",
	19: "j",
	20: "k",
	21: "l",
	22: "m",
	23: "n",
	24: "o",
	25: "p",
	26: "q",
	27: "r",
	28: "s",
	29: "t",
	30: "u",
	31: "v",
	32: "w",
	33: "x",
	34: "y",
	35: "z",
	36: "A",
	37: "B",
	38: "C",
	39: "D",
	40: "E",
	41: "F",
	42: "G",
	43: "H",
	44: "I",
	45: "J",
	46: "K",
	47: "L",
	48: "M",
	49: "N",
	50: "O",
	51: "P",
	52: "Q",
	53: "R",
	54: "S",
	55: "T",
	56: "U",
	57: "V",
	58: "W",
	59: "X",
	60: "Y",
	61: "Z"}

// 10进制转任意进制
func DecimalToAny(num, n int) string {
	new_num_str := ""
	var remainder int
	var remainder_string string
	for num != 0 {
		remainder = num % n
		if 76 > remainder && remainder > 9 {
			remainder_string = tenToAny[remainder]
		} else {
			remainder_string = strconv.Itoa(remainder)
		}
		new_num_str = remainder_string + new_num_str
		num = num / n
	}
	return new_num_str
}

// map根据value找key
func findKey(in string) int {
	result := -1
	for k, v := range tenToAny {
		if in == v {
			result = k
		}
	}
	return result
}

// 任意进制转10进制
func AnyToDecimal(num string, n int) int64 {
	var new_num float64
	new_num = 0.0
	nNum := len(strings.Split(num, "")) - 1
	for _, value := range strings.Split(num, "") {
		tmp := float64(findKey(value))
		if tmp != -1 {
			new_num = new_num + tmp*math.Pow(float64(n), float64(nNum))
			nNum = nNum - 1
		} else {
			break
		}
	}
	return int64(new_num)
}

// 数字转化为两个字母+
func Num2Letter(i int64) string {
	i += 26
	var j = i / 26
	var k = i % 26
	var s string
	if j > 26 {
		s = Num2Letter(j)
	} else if j == 0 {
		fmt.Println("k=", k)
		return string(rune(k+97))
	} else {
		fmt.Println("j=",j, "k=", k)
		return string(rune(j+96)) + string(rune(k+97))
	}
	fmt.Println("zzk=", k)
	return s + string(rune(k+97))
}
