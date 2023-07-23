package randx

import (
	"strconv"
	"time"
)

const timeStringFormat = "20060102150405" // 时间格式 年月日时分秒 14位

// OrderNo 生成时间格式订单号 年月日时分秒(14位) + 随机数(6位)
func OrderNo() string {
	return time.Now().Format(timeStringFormat) + Int(6) // 加 100000 保证随机数为6位
}

// GenOrderNo 生成时间戳订单号，1970年1月1号到现在的毫秒数
func GenOrderNo() string {
	return strconv.FormatInt(time.Now().UnixMilli(), 10) + Int(3)
}
