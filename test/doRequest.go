package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strings"
	"time"
)

func main() {
	for {
		//go doRequest()
		//go doRequest()
		//go doRequest()
		//go doRequest()
		//go doRequest()
		go doReq()
		time.Sleep(time.Second / 10)
	}
}

func doRequest() {
	url := "http://10.221.162.170:7200/client/api"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("url", "http://api.m.mi.com/v1/priorbuy/register_list")
	_ = writer.WriteField("secret", "8CFiclCFwKACNxNfSVJYTlY7kwGKvqOtNYjIw0we0ua51VMNqawQ2GVSS/1Bwc4vQn93cAIH/u0wrO5HjyAllB2j3/oAi7/eiD18piN0lMpOSfUR1J6+/gwa0OX3yif/UEuRAg7kwKj4SLATX4Y/yZhd4jo8cEBJVP6NU/i8eNg=")
	_ = writer.WriteField("client_id", "180100031052")
	_ = writer.WriteField("api_id", "39122")
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	req.Header.Add("Server", "migoServer")
	req.Header.Add("Cookie", "gosessid=9ad711a3a3c6c531ce68b8e864b0b795643da82e")

	req.Header.Set("Content-Type", writer.FormDataContentType())
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
	//if len(string(body)) <= 132 {
	//	fmt.Printf("rate limit len:%d\n", len(string(body)))
	//} else {
	//	fmt.Printf("result len:%d\n", len(string(body)))
	//}
	fmt.Printf("result:%s \n\n\n", string(body))

}

func doReq() {
	url := "http://localhost:11984/v1/user/shop_user"
	method := "POST"

	payload := strings.NewReader(`{
    
}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("USERID", "2199087")
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
	fmt.Printf("%s\n", string(body))
}
