package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

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
	currentOS := runtime.GOOS
	fileExt := getFileExtension(currentOS)
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() || path == root {
			// 【编译二进制文件】不是目录，或者是根目录，跳过
			return nil
		}

		goModPath := filepath.Join(path, "go.mod")
		if _, err := os.Stat(goModPath); err == nil {
			moduleName, err := getModuleName(goModPath)
			if err != nil {
				return err
			}

			if err := runCommandsInDir(path); err != nil {
				return err
			}

			binaryPath := filepath.Join(path, moduleName+fileExt)
			if _, err := os.Stat(binaryPath); os.IsNotExist(err) {
				return fmt.Errorf("【编译二进制文件】编译失败：%s", moduleName)
			}

			log.Printf("【编译二进制文件】编译成功：%s \n", moduleName)
		}
		return filepath.SkipDir
	})
	if err != nil {
		log.Fatalf("【编译二进制文件】失败：%v", err)
		return
	}
}

func getFileExtension(os string) string {
	if os == "windows" {
		return ".exe"
	}
	return ""
}

func getModuleName(goModPath string) (string, error) {
	content, err := os.ReadFile(goModPath)
	if err != nil {
		return "", err
	}

	lines := strings.Split(string(content), "\n")
	if len(lines) > 0 {
		moduleLine := lines[0]
		if strings.HasPrefix(moduleLine, "module") {
			moduleName := strings.TrimSpace(strings.TrimPrefix(moduleLine, "module"))
			return moduleName, nil
		}
	}

	return "", fmt.Errorf("module name not found in %s", goModPath)
}

func runCommandsInDir(dir string) error {
	cmdSlice := []string{"go mod tidy", "go build ."}

	for _, cmdStr := range cmdSlice {
		cmd := exec.Command("cmd", "/C", cmdStr)
		cmd.Dir = dir
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			return err
		}
	}

	return nil
}
