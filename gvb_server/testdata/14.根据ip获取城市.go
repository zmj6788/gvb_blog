package main

import (
	"fmt"
	"net"

	geoip2db "github.com/cc14514/go-geoip2-db"
)

func main() {
	fmt.Print(GetAddr(
		"123.52.105.90"))
}

func GetAddr(ip string) string {
	parseIP := net.ParseIP(ip)
	// 判断是否是内网ip
	if IsIntranetIP(parseIP) {
		return "内网地址"
	}
	db, _ := geoip2db.NewGeoipDbByStatik()
	defer db.Close()
	record, err := db.City(net.ParseIP(ip))
	if err != nil {
		return "错误的地址"
	}
	var province string
	// 省份是否存在
	if len(record.Subdivisions) > 0 {
		province = record.Subdivisions[0].Names["zh-CN"]
	}
	city := record.City.Names["zh-CN"]
	return fmt.Sprintf("%s-%s", province, city)
}

func IsIntranetIP(ip net.IP) bool {
	if ip.IsLoopback() {
		return true
	}

	ip4 := ip.To4()
	if ip4 == nil {
		return true
	}
	// 192.168
	// 172.16 - 172.31
	// 10
	// 169.254
	return (ip4[0] == 192 && ip4[1] == 168) ||
		(ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 32) ||
		(ip4[0] == 10) ||
		(ip4[0] == 169 && ip4[1] == 254)
}
