package ginx

import (
	"errors"
	"fmt"
	"log"
	"net"

	"github.com/gin-gonic/gin"
	"github.com/oschwald/geoip2-golang"
)

// GetClientIP 实现获取客户端IP的功能
func GetClientIP(c *gin.Context) string {
	ClientIP := c.ClientIP()
	fmt.Println("ClientIP:", ClientIP)
	RemoteIP := c.RemoteIP()
	fmt.Println("RemoteIP:", RemoteIP)
	ip := c.Request.Header.Get("X-Forwarded-For")
	if ip == "127.0.0.1" || ip == "::1" || ip == "" {
		ip = c.Request.Header.Get("X-Real-Ip")
	}
	if ip == "" {
		ip = "127.0.0.1"
	}
	if !isLocalIP(RemoteIP) {
		ip = RemoteIP
	}
	if !isLocalIP(ClientIP) {
		ip = ClientIP
	}
	return parseIP(ip)
}

func GetCountryByIP(ip string) (string, error) {
	if ip == "" {
		return "", errors.New("ip 为空")
	}
	if ip == "127.0.0.1" || ip == "localhost" || ip == "::1" {
		return "内部IP", nil
	}
	db, err := geoip2.Open("kit/ginx/GeoLite2-City.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *geoip2.Reader) { _ = db.Close() }(db)
	record, err := db.City(net.ParseIP(ip))
	if err != nil {
		log.Fatal(err)
	}
	return record.City.Names["zh-CN"], nil
}

// isLocalIP 判断字符串是否是本地IP
func isLocalIP(ip string) bool {
	return ip == "127.0.0.1" || ip == "::1" || ip == "localhost"
}

func parseIP(ipStr string) string {
	ip := net.ParseIP(ipStr)
	if ip == nil {
		fmt.Printf("%q is not a valid IP address\n", ipStr)
	} else {
		fmt.Printf("The IP address is %v\n", ip)
		if ip.IsPrivate() {
			fmt.Println("This is a private network IP address.")
		} else {
			fmt.Println("This is a public network IP address.")
		}
	}
	return string(ip)
}
