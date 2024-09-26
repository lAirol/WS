package constants

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"time"
)

type SiteConstants struct {
	SystemWatchInterval time.Duration `json:"system_watch_interval"`
	AdminTimeoutSeconds time.Duration `json:"admin_timeout_seconds"`
	ItemsPerPage        int           `json:"items_per_page"`
	PrimaryColor        string        `json:"primary_color"`
	SecondaryColor      string        `json:"secondary_color"`
	WsConst             struct {
		Points  int `json:"points"`
		Chat    int `json:"chat"`
		SysInfo int `json:"sys_info"`
	} `json:"ws_const"`
}

var Constants *SiteConstants

func LoadSiteConstants(filename string) (*SiteConstants, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var constants SiteConstants
	if err := json.Unmarshal(bytes, &constants); err != nil {
		return nil, err
	}

	return &constants, nil
}

func SaveSiteConstants(filename string, constants *SiteConstants) error {
	bytes, err := json.MarshalIndent(constants, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, bytes, 0644)
}
