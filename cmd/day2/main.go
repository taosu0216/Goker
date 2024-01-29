package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"strconv"
	"syscall"
)

const (
	cgroupMemoryHierarchyMount = "/sys/fs/cgroup/memory"
)

func main() {
	if os.Args[0] == "/proc/self/exe" {
		// 容器进程
		fmt.Printf("current pid %d", syscall.Getpid())
		fmt.Println()
		cmd := exec.Command("sh", "-c", `stress --vm-bytes 200m --vm-keep -m 1`)
		cmd.SysProcAttr = &syscall.SysProcAttr{}
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		// 启动容器进程
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}
	cmd := exec.Command("/proc/self/exe")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("current pid %d", syscall.Getpid())

	// 在系统默认创建挂载了memory subsystem的Hierarchy上创建cgroup
	newCgroup := path.Join(cgroupMemoryHierarchyMount, "cgroup-demo-memory")
	if err = os.Mkdir(newCgroup, 0755); err != nil {
		log.Fatal("mkdir", err)
	}

	// 将容器进程加入到这个cgroup中
	// 将容器进程的PID写入到cgroup的PID列表中
	if err = ioutil.WriteFile(path.Join(newCgroup, "tasks"), []byte(strconv.Itoa(cmd.Process.Pid)), 0644); err != nil {
		log.Fatal("write tasks", err)
	}
	if err = ioutil.WriteFile(path.Join(newCgroup, "memory.limit_in_bytes"), []byte("100m"), 0644); err != nil {
		log.Fatal("write memory.limit_in_bytes", err)
	}
	_, _ = cmd.Process.Wait()
}
