package boshcli

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
)

// Info describes a target BOSH environment
type Info struct {
	Name               string `json:"name"`
	UUID               string `json:"uuid"`
	Version            string `json:"version"`
	User               string `json:"user"`
	CPI                string `json:"cpi"`
	UserAuthentication struct {
		Type    string `json:"type"`
		Options struct {
			URL  string   `json:"url"`
			Urls []string `json:"urls"`
		} `json:"options"`
	} `json:"user_authentication"`
	Features struct {
		DNS struct {
			Status bool `json:"status"`
			Extras struct {
				DomainName string `json:"domain_name"`
			} `json:"extras"`
		} `json:"dns"`
		CompiledPackageCache struct {
			Status bool `json:"status"`
			Extras struct {
				Provider interface{} `json:"provider"`
			} `json:"extras"`
		} `json:"compiled_package_cache"`
		Snapshots struct {
			Status bool `json:"status"`
		} `json:"snapshots"`
		ConfigServer struct {
			Status bool `json:"status"`
			Extras struct {
				Urls []string `json:"urls"`
			} `json:"extras"`
		} `json:"config_server"`
	} `json:"features"`
}

// GetInfo from target BOSH environment
func GetInfo() (info *Info) {
	info = &Info{}
	cmd := exec.Command("sh", "-c", "bosh curl /info")
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("%s\n", stdoutStderr)
		log.Fatal(err)
	}

	if err := json.Unmarshal([]byte(stdoutStderr), info); err != nil {
		log.Fatal(err)
	}

	return
}
