package osx

import (
	"os"
	"path/filepath"
	"time"

	"go.uber.org/zap"
)

// Pwd 获取当前目录
func Pwd() string {
	dir, err := os.Getwd()
	if err != nil {
		zap.L().Fatal("【获取当前目录】失败", zap.Error(err))
	}
	zap.L().Info("【获取当前目录】成功", zap.String("dir", dir))
	return dir
}

// DirParent 获取上级目录
func DirParent(dir string) string {
	return filepath.Dir(dir)
}

// PwdParent 获取当前目录的上级目录
func PwdParent() string {
	return DirParent(Pwd())
}

// MkdirAll 创建目录
func MkdirAll(dir string) error {
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		zap.L().Error("【创建目录】创建失败", zap.Error(err))
		return err
	}
	return nil
}

// RemoveAll 删除所有文件或目录
func RemoveAll(path string) error {
	err := os.RemoveAll(path)
	if err != nil {
		zap.L().Error("【删除文件或目录】删除失败", zap.Error(err))
		return err
	}
	return nil
}

// RemoveTimeoutDir 删除超时目录
func RemoveTimeoutDir(dir string, timeout time.Duration) error {
	fi, err := os.Stat(dir)
	if err != nil || !fi.IsDir() {
		zap.L().Error("【删除目录】失败", zap.Error(err))
	}
	isTimeout := fi.ModTime().Add(timeout).Before(time.Now())
	if !isTimeout {
		return nil
	}
	err = RemoveAll(dir)
	if err != nil {
		return err
	}
	return nil
}

// IsPathExist 判断目录或者文件是否存在
func IsPathExist(dir string) bool {
	_, err := os.Stat(dir)
	if err != nil {
		return false
	}
	return true
}

// InitDir 初始化目录
func InitDir(dir string) {
	// 判断目录是否存在，则清空
	if IsPathExist(dir) {
		_ = RemoveAll(dir)
	}
	_ = MkdirAll(dir)
}
