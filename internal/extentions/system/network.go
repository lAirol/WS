package system

import (
	"WS/internal/config/constants"
	"WS/internal/modules/response"
	"encoding/json"
	"fmt"
	"github.com/shirou/gopsutil/net"
	"net/http"
	"time"
)

type NetWorks struct {
	Target int                 `json:"target"`
	Type   string              `json:"type"`
	Info   map[int]NetWorkInfo `json:"info"`
}
type NetWorkInfo struct {
	Sent float64 `json:"sent"`
	Recv float64 `json:"recv"`
}

func (si *SysInfo) StartNetwork() {
	nwi := NetWorks{Info: make(map[int]NetWorkInfo), Target: constants.Constants.WsConst.SysInfo, Type: "INTERNET"}
	for {
		// Получаем начальные значения
		json := nwi.GetNetworkInfo()
		si.SendMessageAll(json)
	}
}

func (nw *NetWorks) GetNetworkInfo() []byte {
	initial, err := net.IOCounters(true)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	time.Sleep(SystemInfo.UpdateTime)

	final, err := net.IOCounters(true)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	for i := range initial {
		bytesSentPerSec := float64(final[i].BytesSent - initial[i].BytesSent)
		bytesRecvPerSec := float64(final[i].BytesRecv - initial[i].BytesRecv)
		info := NetWorkInfo{
			Sent: bytesSentPerSec,
			Recv: bytesRecvPerSec,
		}
		nw.Info[i] = info
	}

	data, err := json.Marshal(nw)
	return data
}

func GetNetInterfacesParams(w http.ResponseWriter, r *http.Request) {
	initial, _ := net.IOCounters(true) // TODO добавить имя
	type params struct {
		Count int            `json:"count"`
		Name  map[int]string `json:"name"`
	}
	paramsInfo := params{Count: len(initial), Name: make(map[int]string)}
	for i := range initial {
		paramsInfo.Name[i] = initial[i].Name
	}
	response.JSONResponse(w, paramsInfo)
}
