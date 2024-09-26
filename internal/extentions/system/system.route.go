package system

import (
	"net/http"
)

func RegisterRoutes() {
	http.HandleFunc("/system/GetCpuCount", GetCpuCount)
	http.HandleFunc("/system/GetNodeName", GetNodeName)
	http.HandleFunc("/system/GetSysName", GetSysName)
	http.HandleFunc("/system/GetCpuInfo", GetCpuInfo)
	http.HandleFunc("/system/GetSystemInfo", GetSystemInfo)
}
