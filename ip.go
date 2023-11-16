package iutils

import (
	"fmt"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"
	"net"
	"strings"
	"time"
)

// LocalIP 获取机器的IP
func GetLocalIP() string {
	info, _ := net.InterfaceAddrs()
	for _, addr := range info {
		ipNet, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}
		if !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			return ipNet.IP.String()
		}
	}
	return ""
}

/*
验证ip区域是否禁用
forbiddenArea用英文逗号隔开，如：北京,上海
 */
func CheckIpAllowed(forbiddenArea string, ip string) bool {
	if forbiddenArea == "" {
		return true
	}
	areas := strings.Split(forbiddenArea, ",")
	for _, v := range areas {
		if v == "" {
			continue
		}
		if ip == v {
			return false
		}
	}

	resStr1, err1 := httplib.Get("http://whois.pconline.com.cn/ip.jsp?ip=" + ip).SetTimeout(time.Second * 2, time.Second * 2).String()
	resStr2, err2 := httplib.Get(fmt.Sprintf("http://opendata.baidu.com/api.php?query=%s&co=&resource_id=6006&oe=utf8", ip)).SetTimeout(time.Second * 2, time.Second * 2).String()
	if err1 != nil && err2 != nil {
		logs.Error("QueryApiIpErr: err1=", err1, "err2=", err2)
		return false
	}
	for _, v := range areas {
		if v == "" {
			continue
		}
		if strings.Contains(resStr1, v) || strings.Contains(resStr2, v) {
			return false
		}
	}
	return true
}
