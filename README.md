# Goker
## Cgroup
```shell
mkdir cgroup-demo
mount -t cgroup -o none,name=cgroup-demo cgroup-demo ./cgroup-demo
```
-t cgroup: 指定文件系统类型为cgroup
除了cgroup还有
- ext4：这是 Linux 中最常用的文件系统类型。
- ntfs：这是 Windows 系统中的文件系统类型。
- vfat：这是老式的 FAT32 文件系统类型。
- iso9660：这是 CD-ROM 或 DVD-ROM 的文件系统类型。
- nfs：这是网络文件系统 (Network File System) 的类型。
- tmpfs：这是一个临时文件系统，通常用于存储 /tmp 目录的内容。

```text
root@taosu:~# ls cgroup-demo/
cgroup.clone_children  cgroup.procs  cgroup.sane_behavior  notify_on_release  release_agent  tasks
```
- cgroup.clone_children：这个文件的值为 1 时，表示子 cgroup 会继承父 cgroup 的资源限制。如果值为 0，则表示子 cgroup 不会继承父 cgroup 的资源限制。
- cgroup.procs：这个文件中存储了当前 cgroup 中的所有进程 ID。
- cgroup.sane_behavior：这个文件的值为 1 时，表示 cgroup 的行为是合理的。如果值为 0，则表示 cgroup 的行为是不合理的。
- notify_on_release：这个文件的值为 1 时，表示当 cgroup 被释放时，会向 release_agent 指定的程序发送一个通知。如果值为 0，则表示不发送通知。
- release_agent：这个文件中存储了一个可执行程序的路径。当 cgroup 被释放时，会执行这个可执行程序。
- tasks：这个文件中存储了当前 cgroup 中的所有进程 ID。

## 进程资源
```shell
# /sys/fs/cgroup目录下的memory文件
ls /sys/fs/cgroup
cd memory
# 在当前目录创建目录,会在目录中自动生成相应的文件
mkdir cgroup-demo-memory
# 
#root@taosu:/sys/fs/cgroup/memory# ls cgroup-demo-memory
#cgroup.clone_children  memory.kmem.failcnt             memory.kmem.tcp.max_usage_in_bytes  memory.memsw.failcnt             memory.oom_control          memory.usage_in_bytes
#cgroup.event_control   memory.kmem.limit_in_bytes      memory.kmem.tcp.usage_in_bytes      memory.memsw.limit_in_bytes      memory.pressure_level       memory.use_hierarchy
#cgroup.procs           memory.kmem.max_usage_in_bytes  memory.kmem.usage_in_bytes          memory.memsw.max_usage_in_bytes  memory.soft_limit_in_bytes  notify_on_release
#memory.failcnt         memory.kmem.tcp.failcnt         memory.limit_in_bytes               memory.memsw.usage_in_bytes      memory.stat                 tasks
#memory.force_empty     memory.kmem.tcp.limit_in_bytes  memory.max_usage_in_bytes           memory.move_charge_at_immigrate  memory.swappiness


# 只需要将进程id写入tasks文件中即可,然后再在memory.limit_in_bytes中设置限制
```