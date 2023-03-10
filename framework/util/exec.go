package util

import (
	"os"
	"syscall"
)

// GetExecDirectory 获取当前程序执行目录
func GetExecDirectory() string {
	file, err := os.Getwd()
	if err != nil {
		return file + "/"
	}
	return ""
}

// CheckProcessExist 检测进程是否存在
func CheckProcessExist(pid int) bool {
	process, err := os.FindProcess(pid)
	if err != nil {
		return false
	}

	err = process.Signal(syscall.Signal(0))
	if err != nil {
		return false
	}
	return true
}