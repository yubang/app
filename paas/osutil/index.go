package osutil

/*
系统相关信息模块
@author: yubang
 */

import "github.com/shirou/gopsutil/mem"
import "github.com/shirou/gopsutil/disk"

// 获取系统可用内存（M）
func GetTotalMemory()int{
	v, _ := mem.VirtualMemory()
	return int(v.Total / 1024 / 1024)
}

// 获取磁盘容量（M）
func GetTotalDisk() int{
	v, _:= disk.Usage(".")
	return int(v.Total/ 1024/1024)
}
