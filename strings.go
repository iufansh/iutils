package iutils

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
	"time"
	"unicode"
)

//RandomStr 获取一个随机字符串
func RandomStr() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

//字串截取,截取前面
func SubString(s string, pos, length int) string {
	if pos < 0 {
		pos = 0
	}
	runes := []rune(s)
	var l int
	if length == 0 {
		l = len(s)
	} else if length > 0 && pos+length <= len(s) {
		l = pos + length
	} else if length > 0 && pos+length > len(s) {
		l = len(s)
	} else if length < 0 && pos+length >= 0 {
		l = pos
		pos = pos + length
	} else if length < 0 && pos+length < 0 {
		l = pos
		pos = 0
	} else {
		l = pos + length
	}
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

// 格式化字符串，lim位后添加换行
func WrapString(s string, lim uint) string {
	// Initialize a buffer with a slightly larger size to account for breaks
	init := make([]byte, 0, len(s))
	buf := bytes.NewBuffer(init)

	var current uint
	var wordBuf, spaceBuf bytes.Buffer

	for _, char := range s {
		if char == '\n' {
			if wordBuf.Len() == 0 {
				if current+uint(spaceBuf.Len()) > lim {
					current = 0
				} else {
					current += uint(spaceBuf.Len())
					spaceBuf.WriteTo(buf)
				}
				spaceBuf.Reset()
			} else {
				current += uint(spaceBuf.Len() + wordBuf.Len())
				spaceBuf.WriteTo(buf)
				spaceBuf.Reset()
				wordBuf.WriteTo(buf)
				wordBuf.Reset()
			}
			buf.WriteRune(char)
			current = 0
		} else if unicode.IsSpace(char) {
			if spaceBuf.Len() == 0 || wordBuf.Len() > 0 {
				current += uint(spaceBuf.Len() + wordBuf.Len())
				spaceBuf.WriteTo(buf)
				spaceBuf.Reset()
				wordBuf.WriteTo(buf)
				wordBuf.Reset()
			}
			spaceBuf.WriteRune(char)
		} else {

			wordBuf.WriteRune(char)

			if current+uint(spaceBuf.Len()+wordBuf.Len()) > lim && uint(wordBuf.Len()) < lim {
				buf.WriteRune('\n')
				current = 0
				spaceBuf.Reset()
			}
		}
	}

	if wordBuf.Len() == 0 {
		if current+uint(spaceBuf.Len()) <= lim {
			spaceBuf.WriteTo(buf)
		}
	} else {
		spaceBuf.WriteTo(buf)
		wordBuf.WriteTo(buf)
	}

	return buf.String()
}

/**
 * 判断是否包含中文
 */
func HasChinese(str string) bool {
	var count int
	for _, v := range str {
		if unicode.Is(unicode.Han, v) {
			count++
			break
		}
	}
	return count > 0
}

/**
 * 判断是否英文单词
 */
func IsEnglishWord(word string) bool {
	for _, v := range word {
		if !unicode.IsLetter(v) {
			return false
		}
	}
	return true
}

// 删除字符串中的多余空格，有多个空格时，仅保留一个空格
func DeleteExtraSpace(s string) string {
	//删除字符串中的多余空格，有多个空格时，仅保留一个空格
	s1 := strings.Replace(s, "  ", " ", -1)      //替换tab为空格
	regstr := "\\s{2,}"                          //两个及两个以上空格的正则表达式
	reg, _ := regexp.Compile(regstr)             //编译正则表达式
	s2 := make([]byte, len(s1))                  //定义字符数组切片
	copy(s2, s1)                                 //将字符串复制到切片
	spc_index := reg.FindStringIndex(string(s2)) //在字符串中搜索
	for len(spc_index) > 0 {                     //找到适配项
		s2 = append(s2[:spc_index[0]+1], s2[spc_index[1]:]...) //删除多余空格
		spc_index = reg.FindStringIndex(string(s2))            //继续在字符串中搜索
	}
	return string(s2)
}

//利用正则表达式压缩字符串，去除空格或制表符
func DeleteAllSpace(str string) string {
	if str == "" {
		return ""
	}
	//匹配一个或多个空白符的正则表达式
	reg := regexp.MustCompile("\\s+")
	return reg.ReplaceAllString(str, "")
}

// 首字母转大写
func UpperFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

// 首字母转小写
func LowerFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}