package system

import (
	"WS/internal/handlers"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"time"
)

var SystemInfo SysInfo

type SysInfo struct {
	Sysstat struct {
		Hosts []struct {
			Nodename     string `json:"nodename"`
			Sysname      string `json:"sysname"`
			Release      string `json:"release"`
			Machine      string `json:"machine"`
			NumberOfCpus int    `json:"number-of-cpus"`
			Date         string `json:"date"`
			Statistics   []struct {
				Timestamp string `json:"timestamp"`
				CpuLoad   []struct {
					CPUNum string  `json:"cpu"`
					Usr    float64 `json:"usr"`
					Nice   float64 `json:"nice"`
					Sys    float64 `json:"sys"`
					IOWait float64 `json:"iowait"`
					IRQ    float64 `json:"irq"`
					Soft   float64 `json:"soft"`
					Steal  float64 `json:"steal"`
					Guest  float64 `json:"guest"`
					GNice  float64 `json:"gnice"`
					Idle   float64 `json:"idle"`
				} `json:"cpu-load"`
			} `json:"statistics"`
		} `json:"hosts"`
	} `json:"sysstat"`
}

type SignedConn struct {
	//SignedUsers
}

//func (sm *SysInfo) NewSystemMonitor(SysInfo) *SysInfo {
//	return sm
//}

func InitSystemWatch(interval time.Duration) {
	SystemInfo.start(interval) // тут реализовать отправку пользователям данных
}

func (si *SysInfo) start(interval time.Duration) {
	for {
		cmd := exec.Command("mpstat", "-P", "ALL", "-o", "JSON")
		output, err := cmd.Output()
		json.Unmarshal(output, si)
		fmt.Println(si.Sysstat.Hosts[0].Statistics[0])
		if err != nil {
			log.Println("Ошибка монитринга: ", err)
		}
		time.Sleep(interval)
		sender := handlers.Sender{
			Server: true,
			Msg:    output,
		}
		handlers.Broadcast <- sender
	}
}
