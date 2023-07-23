package main

import (
	"log"
	"os"
	"path/filepath"

	"kit/logger"
	"kit/tools/osx"
	"kit/tools/timex"
)

func init() {
	logger.InitZap()
}

func main() {
	defer timex.Cost()()
	root := osx.PwdParent()
	log.Printf("【编译二进制文件】上级目录: %s \n", root)
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() || path == root || // 【编译二进制文件】不是目录，或者是根目录，跳过
			!osx.IsFileExists(filepath.Join(path, "go.mod")) { // 【编译二进制文件】go.mod 文件不存在，跳过
			return nil
		}

		err = osx.RunBuild(path)
		if err != nil {
			log.Fatalf("【编译二进制文件】失败：%v", err)
			return err
		}
		log.Printf("【编译二进制文件】编译成功：%s \n", filepath.Base(path))

		return nil
	})
	if err != nil {
		log.Fatalf("【编译二进制文件】失败：%v", err)
		return
	}
}
