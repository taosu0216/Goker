// main包提供了一个简单的程序，该程序创建一个新的shell进程
// 并将新进程的标准输入、输出和错误流连接到当前进程。
package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

// main是程序的入口点。
func main() {
	// 使用"sh"作为参数，通过exec.Command函数创建一个新的shell进程。
	cmd := exec.Command("sh")

	// 将*Cmd实例的SysProcAttr字段设置为一个新的syscall.SysProcAttr实例，
	// 其中Cloneflags字段设置为syscall.CLONE_NEWUTS。这意味着新进程将拥有自己的UTS命名空间，
	// 这是Linux内核的一个特性，可以在不同的进程之间隔离某些系统标识符（如主机名）。
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS |
			syscall.CLONE_NEWIPC |
			syscall.CLONE_NEWPID |
			syscall.CLONE_NEWNS |
			syscall.CLONE_NEWNET |
			syscall.CLONE_NEWUSER,
		UidMappings: []syscall.SysProcIDMap{
			{
				ContainerID: 1,
				HostID:      0,
				Size:        1,
			},
		},
		GidMappings: []syscall.SysProcIDMap{
			{
				ContainerID: 1,
				HostID:      0,
				Size:        1,
			},
		},
	}

	// 将新进程的标准输入、输出和错误流连接到当前进程。
	// 这意味着用户可以像与原始进程交互一样与新的shell进程交互。
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// 通过调用*Cmd实例的Run方法来启动新进程。
	// 如果启动进程时出现错误，它会被记录下来，程序将以非零状态码退出。
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
