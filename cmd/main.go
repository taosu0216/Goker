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
		Cloneflags:
		//创建一个新的UTS命名空间，隔离主机名和域名
		syscall.CLONE_NEWUTS |
			// 创建一个新的IPC命名空间，隔离System V IPC和POSIX消息队列
			syscall.CLONE_NEWIPC |
			//创建一个新的PID命名空间，使得每个命名空间都有自己独立的PID空间
			syscall.CLONE_NEWPID |
			//创建一个新的NS命名空间，隔离不同进程的mount点
			//TODO: 这里有小坑,仅仅确立新的NS命名空间还不能完全隔离不同进程的mount点,还得自己手动设置当前紫禁城的mount挂载点,才能与其他进程隔离
			syscall.CLONE_NEWNS |
			//创建一个新的NET命名空间，隔离网络设备、协议栈、端口等网络资源。
			syscall.CLONE_NEWNET |
			//创建一个新的用户命名空间，隔离用户和组ID
			syscall.CLONE_NEWUSER,
		UidMappings: []syscall.SysProcIDMap{
			{
				// 容器的用户ID
				ContainerID: 1,
				// 主机的用户ID
				HostID: 0,
				// 映射的ID数量
				Size: 1,
			},
		},
		GidMappings: []syscall.SysProcIDMap{
			{
				// 容器的组ID
				ContainerID: 1,
				// 主机的组ID
				HostID: 0,
				// 映射的ID数量
				Size: 1,
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
