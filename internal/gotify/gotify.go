package gotify

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/eterline/echidna/internal/settings"
)

type IPInfo struct {
	Query       string  `json:"query"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	Zip         string  `json:"zip"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	TZ          string  `json:"timezone"`
	ISP         string  `json:"isp"`
	Org         string  `json:"org"`
	As          string  `json:"as"`
}

const msgText = `
%s have visited you.
Host: %s
------------------------
UserAgnet: %s
%s: %s
%s: %s | %s
ZIP: %s
Location: %v | %v
ISP: %s %s
`

func SendMessage(r *http.Request, c settings.Config) {
	ipData := checkIp(r.RemoteAddr)
	msg := fmt.Sprintf(msgText,
		ipData.Query, c.Host, r.UserAgent(),
		ipData.CountryCode, ipData.Country,
		ipData.Region, ipData.RegionName, ipData.City,
		ipData.Zip,
		ipData.Lat, ipData.Lon,
		ipData.ISP, ipData.Org,
	)
	remote := fmt.Sprintf("%smessage?token=%s", c.Gotify.URL, c.Gotify.ApiKey)

	http.PostForm(remote,
		url.Values{"message": {msg}, "title": {"Echidna catcher"}})
}

func StartMessage(c settings.Config) {
	if c.StartMsg {
		remote := fmt.Sprintf("%smessage?token=%s", c.Gotify.URL, c.Gotify.ApiKey)
		http.PostForm(remote,
			url.Values{"message": {"App has been started."}, "title": {"Echidna catcher"}})
	}
}

func checkIp(addr string) IPInfo {
	var data IPInfo
	ip := strings.Split(addr, `:`)[0]

	response, err := http.Get(fmt.Sprintf("http://ip-api.com/json/%s", ip))
	if err != nil {
		return data
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return data
	}
	err = json.Unmarshal(responseData, &data)
	if err != nil {
		return data
	}
	fmt.Println(data)
	return data
}
