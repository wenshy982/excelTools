package xls

import (
	"errors"

	"log"
)

var SIPSites = []string{"SG", "MY", "TH", "VN", "PH", "TW", "BR", "MX", "CO", "CL"} // 所有站点

// InSIP 判断是否为SIP站点
func InSIP(site string) bool {
	for _, v := range SIPSites {
		if v == site {
			return true
		}
	}
	return false
}

// IndexOf 获取字符串在切片中的索引
func IndexOf(s []string, v string, skipIndex int) int {
	for i, val := range s {
		if i < skipIndex {
			continue
		}
		if v == val {
			return i
		}
	}
	return -1
}

var ErrNotFound = errors.New("not found")

// GetFinalListRowByTitle 从输入文件读取指定列
func GetFinalListRowByTitle(inputTitle, inputRow, templateRow []string) ([]string, error) {
	for i, v := range templateRow {
		if v == "" {
			continue
		}
		index := IndexOf(inputTitle, v, 0)
		if index == -1 {
			log.Fatalf("【读取输入文件】读取失败，未找到 %s 列 \n", v)
			return nil, ErrNotFound
		}
		// Region列特殊处理 因为有多个Region列
		// 如果是Region列，且不是SIP站点，则跳过
		if v == "Region" {
			if !InSIP(inputRow[index]) {
				index = IndexOf(inputTitle, v, index+1)
				if index == -1 {
					log.Fatalf("【读取输入文件】读取失败，未找到 %s 列 \n", v)
					return nil, ErrNotFound
				}
			}
		}
		templateRow[i] = inputRow[index]
	}

	return templateRow, nil
}
