package osx

import (
	"os"
	"os/exec"
)

func RunCmd(dir string, cmdSlice ...string) error {
	for _, v := range cmdSlice {
		cmd := exec.Command("cmd", "/C", v)
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

func RunBuild(dir string) error {
	return RunCmd(dir, "go mod tidy", "go build .")
}
