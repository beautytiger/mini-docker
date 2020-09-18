package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	// 测试时，这里面很多trick
	// 重置主机名时，使用hostname -b, 不能使用 hostnamectl
	// 使用root运行运行，最好先build，再运行二进制
	// pstree -pl
	// readlink /proc/{pid}/ns/uts 查看进程两侧的uts namespace
	// 查看当前进程 id： echo $$
	cmd := exec.Command("sh")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
