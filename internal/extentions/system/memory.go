package system

import (
	"fmt"
	"golang.org/x/sys/unix"
)

type MemoryInfo struct {
	FreeRAM  uint64    `json:"free_ram"`
	Uptime   int64     `json:"uptime"`
	Procs    uint16    `json:"procs"`
	Loads    [3]uint64 `json:"loads"`
	Freeswap uint64    `json:"freeswap"`
	Freehigh uint64    `json:"freehigh"`
}

func (si *SysInfo) GetMemInfo() {
	memInfo := &MemoryInfo{}
	var info unix.Sysinfo_t
	err := unix.Sysinfo(&info)
	if err != nil {
		fmt.Printf("Ошибка получения информации: %v\n", err)
		return
	}

	memInfo.FreeRAM = info.Freeram * uint64(info.Unit)
	memInfo.Uptime = info.Uptime
	memInfo.Procs = info.Procs
	memInfo.Loads = info.Loads
	memInfo.Freeswap = info.Freeswap * uint64(info.Unit)
	memInfo.Freehigh = info.Freehigh * uint64(info.Unit)

	fmt.Println(memInfo)
}
