package osx

import (
	"log"
	"os"
)

// IsFileExists 判断文件是否存在
func IsFileExists(file string) bool {
	_, err := os.Stat(file)
	if err == nil { // if not error, 文件存在
		return true
	}
	if os.IsNotExist(err) { // 文件不存在
		return false
	}
	log.Fatalf("判断文件是否存在失败: %s", file)
	return false
}
