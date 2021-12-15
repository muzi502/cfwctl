package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

const ipv4_regex = `(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}`

var urlList = []string{
	"http://ip.sb",
	"http://ip.cip.cc",
	"http://myip.ipip.net",
}

func GetPublicIP() string {
	for _, url := range urlList {
		client := &http.Client{}
		request, err := http.NewRequest("GET", url, nil)
		if err != nil {
			continue
		}
		request.Header.Set("User-Agent", "curl/7.54.0")
		resp, err := client.Do(request)
		if resp.StatusCode != 200 && err != nil {
			continue
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		reg := regexp.MustCompile(ipv4_regex)
		ipList := reg.FindAllString(string(body), -1)
		if len(ipList) > 0 {
			fmt.Printf("my public ip is %s\n", ipList[0])
			return ipList[0]
		}
	}
	return ""
}
