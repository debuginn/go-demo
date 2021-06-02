package topic05

import (
	"fmt"
)

type Params struct {
	A int64
	B int64
}

func Sum(req *Params) {
	req.A = 22
	fmt.Printf("param mem addr:%p, pointer addr:%p, value:%v\n", req, &req, *req)
}

func TPointer() {

	param := &Params{1, 2}

	fmt.Printf("param mem addr:%p, pointer addr:%p, value:%v\n", param, &param, *param)
	Sum(param)
	fmt.Printf("param mem addr:%p, pointer addr:%p, value:%v\n", param, &param, *param)
}
