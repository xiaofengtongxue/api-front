package proxy

import (
	"time"
)

type Host struct {
	Name      string `json:"-"`
	Url       string `json:"url"`
	Enable    bool   `json:"enable"`
	Note      string `json:"note"`
	SortIndex int    `json:"sort"`
	Checked   bool   `json:"-"`
}

type Hosts map[string]*Host

func newHosts() Hosts {
	return make(Hosts)
}

func (h *Host) copy() *Host {
	return &Host{
		Name:      h.Name,
		Url:       h.Url,
		Enable:    h.Enable,
		Note:      h.Note,
		SortIndex: h.SortIndex,
	}
}

func (hs Hosts) addNewHost(host *Host) {
	hs[host.Name] = host
}

func (hs Hosts) init() {
	for name, host := range hs {
		host.Name = name
	}
}

func newHost(name string, url string, enable bool) *Host {
	return &Host{
		Name:   name,
		Url:    url,
		Enable: enable,
	}
}

func (hs Hosts) getDefaultHostName() string {
	n := time.Now().UnixNano() % int64(len(hs))
	for name := range hs {
		if n == 0 {
			return name
		}
		n = n - 1
	}
	return ""
}

func (hs Hosts) activeHostsNum() int {
	num := 0
	for _, host := range hs {
		if host.Enable {
			num = num + 1
		}
	}
	return num
}

func (hs Hosts) getHostsWithPref(pref []string) []*Host {

	enableNames := []string{}
	for name, host := range hs {
		if host.Enable {
			enableNames = append(enableNames, name)
		}
	}
	var arr []*Host
	for _, name := range pref {
		if InStringSlice(name, enableNames) {
			h := hs[name].copy()
			h.Checked = true
			arr = append(arr, h)
		}
	}
	for name, host := range hs {
		if host.Enable && !InStringSlice(name, pref) {
			arr = append(arr, host)
		}
	}
	return arr
}
