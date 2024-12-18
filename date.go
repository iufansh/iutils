package iutils

import (
	"time"

	"github.com/pkg/errors"
)

// 废弃，直接用本机时间了
func GetNetTime() (time.Time, error) {
	// {"api":"mtop.common.getTimestamp","v":"*","ret":["SUCCESS::接口调用成功"],"data":{"t":"1555314256704"}}
	// resp, err := http.Get("http://api.m.taobao.com/rest/api3.do?api=mtop.common.getTimestamp")
	// if err != nil {
	// 	return time.Now(), errors.New("获取网络时间请求失败")
	// }
	// defer resp.Body.Close()
	// body, _ := ioutil.ReadAll(resp.Body)
	// bodys := string(body)

	// if !strings.Contains(bodys, "SUCCESS::接口调用成功") {
	// 	return time.Now(), errors.New("接口调用失败：" + bodys)
	// }
	// regRule := "^{([\\s\\S]+)\"t\":\"([\\d]{13})\"([\\s\\S]+)}$"
	// reg := regexp.MustCompile(regRule)
	// results := reg.FindStringSubmatch(bodys)
	// if len(results) != 4 {
	// 	return time.Now(), errors.New("返回数据解析错误：" + bodys)
	// }
	// fmt.Println(results[2])
	// t, err := strconv.ParseInt(results[2], 10, 64)
	// return time.Unix(0, t*int64(time.Millisecond)), nil
	return time.Now(), errors.New("获取网络时间请求失败")
}

// 计算两个日期相差几天
func TimeSub(t1, t2 time.Time) int {
	t1 = time.Date(t1.Year(), t1.Month(), t1.Day(), 0, 0, 0, 0, time.Local)
	t2 = time.Date(t2.Year(), t2.Month(), t2.Day(), 0, 0, 0, 0, time.Local)

	return int(t1.Sub(t2).Hours() / 24)
}

// 格式2006-01-02
func ParseDate(d time.Time) time.Time {
	if d.IsZero() {
		return d
	}
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, time.Local)
}

// 格式2006-01-02
func FormatDate(d time.Time) string {
	if d.IsZero() {
		return ""
	}
	return d.Format("2006-01-02")
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

// 判断时间是当年的第几周,周一为第一天
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

// 判断时间是当年的第几周,周日为第一天
func WeekByDate2(t time.Time) int {
	yearDay := t.YearDay()
	yearFirstDay := t.AddDate(0, 0, -yearDay+1)
	firstDayInWeek := int(yearFirstDay.Weekday())

	// 今年第一周有几天
	firstWeekDays := 7 - firstDayInWeek     // 7~1
	week := (yearDay-firstWeekDays+6)/7 + 1 // 填充再整除
	return week
}

/*
*
获取本周周一的日期
*/
func GetFirstDateOfWeek() time.Time {
	now := time.Now()

	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}

	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
}

/**
 * @Author Dong
 * @Description 获得当前月的初始和结束日期
 * @Date 16:29 2020/8/6
 * @Param  * @param null
 * @return
 **/
func GetMonthDay() (string, string) {
	now := time.Now()
	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()

	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)
	f := firstOfMonth.Unix()
	l := lastOfMonth.Unix()
	return time.Unix(f, 0).Format("2006-01-02") + " 00:00:00", time.Unix(l, 0).Format("2006-01-02") + " 23:59:59"
}

/**
 * @Author Dong
 * @Description 获得服务当前时区当前周的初始和结束日期，周一为一周开始
 * @Date 16:32 2020/8/6
 * @Param  * @param null
 * @return
 **/
func GetWeekDay() (string, string) {
	return GetWeekDays(time.Now(), time.Monday)
	//now := time.Now()
	//offset := int(time.Monday - now.Weekday())
	////周日做特殊判断 因为time.Monday = 0
	//if offset > 0 {
	//	offset = -6
	//}
	//
	//lastoffset := int(time.Saturday - now.Weekday())
	////周日做特殊判断 因为time.Monday = 0
	//if lastoffset == 6 {
	//	lastoffset = -1
	//}
	//
	//firstOfWeek := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	//lastOfWeeK := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, lastoffset+1)
	//f := firstOfWeek.Unix()
	//l := lastOfWeeK.Unix()
	//return time.Unix(f, 0).Format("2006-01-02") + " 00:00:00", time.Unix(l, 0).Format("2006-01-02") + " 23:59:59"
}

/**
 * @Author Dong
 * @Description 获得服务当前时区当前周的初始和结束日期，周日为一周开始
 * @Date 16:32 2020/8/6
 * @Param  * @param null
 * @return
 **/
func GetWeekDay2() (string, string) {
	return GetWeekDays(time.Now(), time.Sunday)
	//now := time.Now()
	//offset := int(time.Sunday - now.Weekday())
	//
	//lastoffset := int(time.Saturday - now.Weekday())
	//
	//firstOfWeek := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	//lastOfWeeK := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, lastoffset)
	//f := firstOfWeek.Unix()
	//l := lastOfWeeK.Unix()
	//return time.Unix(f, 0).Format("2006-01-02") + " 00:00:00", time.Unix(l, 0).Format("2006-01-02") + " 23:59:59"
}

func GetWeekDays(t time.Time, weekday time.Weekday) (string, string) {
	now := t
	var offset int
	var lastoffset int
	if weekday == time.Sunday {
		offset = int(time.Sunday - now.Weekday())
		lastoffset = int(time.Saturday - now.Weekday())
	} else if weekday == time.Monday {
		offset := int(time.Monday - now.Weekday())
		//周日做特殊判断 因为time.Monday = 0
		if offset > 0 {
			offset = -6
		}

		lastoffset := int(time.Saturday - now.Weekday())
		//周日做特殊判断 因为time.Monday = 0
		if lastoffset == 6 {
			lastoffset = -1
		}
	} else {
		return "", ""
	}

	firstOfWeek := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	lastOfWeeK := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, lastoffset)
	f := firstOfWeek.Unix()
	l := lastOfWeeK.Unix()
	return time.Unix(f, 0).Format("2006-01-02") + " 00:00:00", time.Unix(l, 0).Format("2006-01-02") + " 23:59:59"
}

/**
 * @Author Dong
 * @Description //获得当前季度的初始和结束日期
 * @Date 16:33 2020/8/6
 * @Param  * @param null
 * @return
 **/
func GetQuarterDay() (string, string) {
	year := time.Now().Format("2006")
	month := int(time.Now().Month())
	var firstOfQuarter string
	var lastOfQuarter string
	if month >= 1 && month <= 3 {
		//1月1号
		firstOfQuarter = year + "-01-01 00:00:00"
		lastOfQuarter = year + "-03-31 23:59:59"
	} else if month >= 4 && month <= 6 {
		firstOfQuarter = year + "-04-01 00:00:00"
		lastOfQuarter = year + "-06-30 23:59:59"
	} else if month >= 7 && month <= 9 {
		firstOfQuarter = year + "-07-01 00:00:00"
		lastOfQuarter = year + "-09-30 23:59:59"
	} else {
		firstOfQuarter = year + "-10-01 00:00:00"
		lastOfQuarter = year + "-12-31 23:59:59"
	}
	return firstOfQuarter, lastOfQuarter
}
