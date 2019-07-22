package lovense

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type Remote struct {
}

func NewRemote() *Remote {
	return &Remote{}
}

type domainInfo struct {
	DeviceId   string             `json:"deviceId"`
	Domain     string             `json:"domain"`
	HTTPPort   int                `json:"httpPort"`
	HTTPSPort  int                `json:"httpsPort"`
	WSPort     int                `json:"wsPort"`
	WSSPort    int                `json:"wssPort"`
	Platform   string             `json:"platform"`
	AppVersion string             `json:"appVersion"`
	Toys       map[string]toyInfo `json:"toys"`
}

type toyInfo struct {
	Status   string `json:"status"`
	Name     string `json:"name"`
	Battery  int    `json:"battery"`
	ID       string `json:"id"`
	Nickname string `json:"nickname"`
}

// Discover search all available toys on LAN using Lovense Connect app
func (r *Remote) Discover() ([]*Toy, error) {
	resp, err := http.Get("https://api.lovense.com/api/lan/getToys")

	if err != nil {
		return nil, fmt.Errorf("failed to request discover api: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("discover api return bad http code, status code error: %d %s", resp.StatusCode, resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, fmt.Errorf("failed to read discover api body: %v", err)
	}

	domains := make(map[string]domainInfo)

	err = json.Unmarshal(body, &domains)

	if err != nil {
		return nil, fmt.Errorf("failed to parse discover api json: %v", err)
	}

	var toys []*Toy

	for _, domain := range domains {
		host := resolve(domain.Domain)

		for _, toy := range domain.Toys {
			toyStatus, _ := strconv.Atoi(toy.Status)
			toys = append(toys, &Toy{
				ID:     toy.ID,
				Host:   fmt.Sprintf("https://%s:%d", host, domain.HTTPSPort),
				Name:   toy.Name,
				Status: Status(toyStatus),
			})
		}
	}

	return toys, nil
}

// Not all DNS seem to resolve local domain to IP, so we do it manually here :)
func resolve(host string) string {
	host = strings.Replace(host, ".lovense.club", "", -1)
	host = strings.Replace(host, "-", ".", -1)
	return host
}
