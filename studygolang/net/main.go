package main

import "fmt"

func main() {
	//ip, ipNet, err := net.ParseCIDR("192.0.2.1/24")
	//if err != nil {
	//	return
	//}
	//
	//fmt.Printf("ip:%v,  ipNet:%v", ip, ipNet)
	//
	//fmtUrl := 111
	//_ = fmt.Errorf("%v", fmtUrl)

	str := FormatImei("861600323744643")
	fmt.Printf("%s", str)
}

func FormatImei(imei string) (formatImei string) {
	if len(imei) < 15 {
		return imei
	}
	formatImei = imei[:7] + "****" + imei[len(imei)-4:]
	return
}
