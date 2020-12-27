package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	req, err := http.NewRequest(http.MethodGet, "http://www.debuginn.cn", nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	req.Header.Set("Cookie", "name=foo")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := http.Client{}
	res, err := client.Do(req)

	//defer res.Body.Close()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(body))
}
