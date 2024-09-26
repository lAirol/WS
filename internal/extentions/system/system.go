package system

import (
	"WS/internal/config/constants"
	"WS/internal/modules/response"
	"WS/internal/modules/users"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"os/exec"
	"time"
)

var SystemInfo SysInfo

type SysInfo struct {
	Sys
	UpdateTime time.Duration
}

type Sys struct {
	AdditionalInfo
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

type AdditionalInfo struct {
	Target int    `json:"target"`
	Type   string `json:"type"`
}

type SignedConn struct {
	//SignedUsers
}

//func (sm *SysInfo) NewSystemMonitor(SysInfo) *SysInfo {
//	return sm
//}

func InitSystemWatch(interval time.Duration) {
	SystemInfo.UpdateTime = interval
	SystemInfo.start() // тут реализовать отправку пользователям данных
}

func (si *SysInfo) start() {
	for {
		cmd := exec.Command("mpstat", "-P", "ALL", "-o", "JSON", "1", "1")
		output, err := cmd.Output()
		json.Unmarshal(output, &si.Sys)
		go func() {
			if len(users.CurrClients.AdminsClients) > 0 {
				si.Sys.Target = constants.Constants.WsConst.SysInfo
				si.Type = "CPU"
				output, _ = json.Marshal(si.Sys)
				if err != nil {
					log.Println("Ошибка монитринга: ", err)
				}
				si.SendMessageAll(output)
			}
		}()
		//time.Sleep(si.UpdateTime)
	}
}

func (si *SysInfo) SendMessageAll(output []byte) {
	for _, client := range users.CurrClients.AdminsClients {
		si.SendMessage(client, output) // добавить go если надо
	}
}

func (si *SysInfo) SendMessage(client *users.AdminClient, message []byte) {
	if client.Conn != nil {
		client.Mu.Lock()
		defer client.Mu.Unlock()
		err := client.Conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func GetCpuCount(w http.ResponseWriter, r *http.Request) {
	response.JSONResponse(w, SystemInfo.Sysstat.Hosts[0].NumberOfCpus)
}

func GetNodeName(w http.ResponseWriter, r *http.Request) {
	response.JSONResponse(w, SystemInfo.Sysstat.Hosts[0].Nodename)
}

func GetSysName(w http.ResponseWriter, r *http.Request) {
	response.JSONResponse(w, SystemInfo.Sysstat.Hosts[0].Sysname)
}

func GetCpuInfo(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		Number int `json:"number"`
	}
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	// Access the CPU load based on the parsed number
	cpuLoad := SystemInfo.Sysstat.Hosts[0].Statistics[0].CpuLoad[requestData.Number+1]
	response.JSONResponse(w, cpuLoad)
}

func GetSystemInfo(w http.ResponseWriter, r *http.Request) {
	response.JSONResponse(w, SystemInfo)
}
