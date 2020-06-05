package iutils

import (
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func GetNetTime() (time.Time, error) {
	// {"api":"mtop.common.getTimestamp","v":"*","ret":["SUCCESS::接口调用成功"],"data":{"t":"1555314256704"}}
	resp, err := http.Get("http://api.m.taobao.com/rest/api3.do?api=mtop.common.getTimestamp")
	if err != nil {
		return time.Now(), errors.New("获取网络时间请求失败")
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	bodys := string(body)

	if !strings.Contains(bodys, "SUCCESS::接口调用成功") {
		return time.Now(), errors.New("接口调用失败：" + bodys)
	}
	regRule := "^{([\\s\\S]+)\"t\":\"([\\d]{13})\"([\\s\\S]+)}$"
	reg := regexp.MustCompile(regRule)
	results := reg.FindStringSubmatch(bodys)
	if len(results) != 4 {
		return time.Now(), errors.New("返回数据解析错误：" + bodys)
	}
	fmt.Println(results[2])
	t, err := strconv.ParseInt(results[2], 10, 64)
	return time.Unix(0, t*int64(time.Millisecond)), nil
}

// 计算两个日期相差几天
func TimeSub(t1, t2 time.Time) int {
	t1 = time.Date(t1.Year(), t1.Month(), t1.Day(), 0, 0, 0, 0, time.Local)
	t2 = time.Date(t2.Year(), t2.Month(), t2.Day(), 0, 0, 0, 0, time.Local)

	return int(t1.Sub(t2).Hours() / 24)
}

// 格式2006-01-02 15:04:05
func FormatDatetime(d time.Time) string {
	if d.IsZero() {
		return ""
	}
	return d.Format("2006-01-02 15:04:05")
}

// 格式20060102150405
func FormatDatetimeKeepNumber(d time.Time) string {
	if d.IsZero() {
		return ""
	}
	return d.Format("20060102150405")
}

//判断时间是当年的第几周,周一为第一天
func WeekByDate(t time.Time) int {
	yearDay := t.YearDay()
	yearFirstDay := t.AddDate(0, 0, -yearDay+1)
	firstDayInWeek := int(yearFirstDay.Weekday())

	//今年第一周有几天
	firstWeekDays := 1
	if firstDayInWeek != 0 {
		firstWeekDays = 7 - firstDayInWeek + 1
	}
	var week int
	if yearDay <= firstWeekDays {
		week = 1
	} else {
		week = (yearDay-firstWeekDays-1)/7 + 2
	}
	return week
}
