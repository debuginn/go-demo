package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	ipVerifyList := "192.168.1.0-192.172.3.255"
	ip := "192.170.223.1"
	ipSlice := strings.Split(ipVerifyList, `-`)
	if len(ipSlice) < 0 {
		return
	}
	if ip2Int(ip) >= ip2Int(ipSlice[0]) && ip2Int(ip) <= ip2Int(ipSlice[1]) {
		fmt.Println("ip in iplist")
		return
	}
	fmt.Println("ip not in iplist")
}

func ip2Int(ip string) int64 {
	if len(ip) == 0 {
		return 0
	}
	bits := strings.Split(ip, ".")
	if len(bits) < 4 {
		return 0
	}
	b0 := string2Int(bits[0])
	b1 := string2Int(bits[1])
	b2 := string2Int(bits[2])
	b3 := string2Int(bits[3])

	var sum int64
	sum += int64(b0) << 24
	sum += int64(b1) << 16
	sum += int64(b2) << 8
	sum += int64(b3)

	return sum
}

func string2Int(in string) (out int) {
	out, _ = strconv.Atoi(in)
	return
}

func isBelong(ip, cidr string) bool {
	ipAddr := strings.Split(ip, `.`)
	if len(ipAddr) < 4 {
		return false
	}
	cidrArr := strings.Split(cidr, `/`)
	if len(cidrArr) < 2 {
		return false
	}

	var ipMask string
	cidrArr2 := strings.Split(cidr, `.`)
	if len(cidrArr2) < 4 {
		return false
	}
	if cidrArr2[1] == "0" {
		ipMask = "255.0.0.0"
	}
	if cidrArr2[2] == "0" {
		ipMask = "255.255.0.0"
	}
	if cidrArr2[3] == "0" {
		ipMask = "255.255.255.0"
	}

	var tmp = make([]string, 0)
	for key, value := range strings.Split(ipMask, `.`) {
		iint, _ := strconv.Atoi(value)

		iint2, _ := strconv.Atoi(ipAddr[key])

		tmp = append(tmp, strconv.Itoa(iint&iint2))
	}
	return strings.Join(tmp, `.`) == cidrArr[0]
}
