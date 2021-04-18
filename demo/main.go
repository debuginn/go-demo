package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func main() {
	for {
		doRequest()
		time.Sleep(time.Second / 10)
	}
}

func doRequest() {
	url := "http://localhost:14082/v1/home/float_info"
	method := "POST"

	payload := strings.NewReader(`{}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("USERID", "3150076059")
	req.Header.Add("LAT", "40.0479625108507")
	req.Header.Add("LNG", "116.3173621961806")
	req.Header.Add("CLIENTIP", "39.107.70.169")
	req.Header.Add("CLIENTID", "180100031052")
	req.Header.Add("LogID", "1")
	req.Header.Add("Content-Type", "text/plain")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
