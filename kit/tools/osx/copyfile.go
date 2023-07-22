package osx

import (
	"fmt"
	"io"
	"log"
	"os"
)

var (
	ErrStatModeIsNotRegular = fmt.Errorf("文件不是普通文件")
	ErrAlreadyExists        = fmt.Errorf("文件已存在")
)

// CopyFileBuf 复制文件 src -> dst bufSize 为缓冲区大小
func CopyFileBuf(src, dst string, bufSize int64) error {
	srcStat, err := os.Stat(src)
	if err != nil {
		return err
	}

	if !srcStat.Mode().IsRegular() {
		log.Printf("文件不是普通文件: %s", src)
		return ErrStatModeIsNotRegular
	}
	source, err := os.Open(src)
	if err != nil {
		log.Printf("打开文件失败: %s", src)
		return err
	}

	defer func(source *os.File) { _ = source.Close() }(source)

	_, err = os.Stat(dst)
	if err == nil {
		log.Printf("文件已存在: %s", dst)
		return ErrAlreadyExists
	}

	destination, err := os.Create(dst)
	if err != nil {
		log.Printf("创建文件失败: %s", dst)
		return err
	}
	defer func() { _ = destination.Close() }()

	buf := make([]byte, bufSize)
	for {
		n, err := source.Read(buf)
		if err != nil && err != io.EOF {
			log.Printf("复制文件失败: %s", dst)
			return err
		}
		if n == 0 {
			break
		}

		if _, err := destination.Write(buf[:n]); err != nil {
			log.Printf("复制文件失败: %s", dst)
			return err
		}
	}
	return err
}

// CopyFile 复制文件 src -> dst
func CopyFile(src, dst string) {
	existFlag := IsFileExists(dst)
	if existFlag {
		return
	}
	_ = CopyFileBuf(src, dst, 204800)
}
