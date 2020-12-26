package main

import (
	"encoding/json"
	"fmt"
	"time"
)

var str = `[
	{
		"code":1,
		"p_code":0,
		"name":"c1"
	},
{
"code":4,
"p_code":1,
"name":"c4"
},
{
"code":2,
"p_code":1,
"name":"c2"
},
{
"code":5,
"p_code":3,
"name":"c5"
},
{
"code":3,
"p_code":1,
"name":"c3"
},
{
"code":6,
"p_code":3,
"name":"c6"
},
{
"code":7,
"p_code":4,
"name":"c7"
},
{
"code":8,
"p_code":2,
"name":"c8"
},
{
"code":9,
"p_code":2,
"name":"c9"
},
{
"code":10,
"p_code":9,
"name":"c10"
},
{
"code":11,
"p_code":9,
"name":"c11"
},
{
"code":12,
"p_code":5,
"name":"c12"
}
]`

/*数据结构*/
type CodeList struct {
	Code   int
	P_Code int
	Name   string
}
type Ret struct {
	Code   int
	P_Code int
	Name   string
	Child  []CodeList
}

/*const CHOOSENODE = 4*/
var ChooseNode = 4

func main() {
	cond := make([]CodeList, 0)
	err := json.Unmarshal([]byte(str), &cond)
	exitOnError(err)
	ret := make([]CodeList, 0)
	/*选择需要开始遍历的节点*/
	//tmp := CodeList{}
	//tmp.Code = ChooseNode
	//tmp.Code = ""
	//find(v,cond,&ret)
	for _, v := range cond {
		/*选择需要开始遍历的节点*/
		if v.Code == ChooseNode {
			find(v, cond, &ret)
			break
		}
	}
}

/*
查询指定节点的子节点
*/
func findChile(currentCode CodeList, cond []CodeList) (child []CodeList) {
	for _, v := range cond {
		if currentCode.Code == v.P_Code {
			child = append(child, v)
		}
	}
	return
}

/*递归函数*/
func find(currentCode CodeList, cond []CodeList, Ret *[]CodeList) {
	//
	tmp := findChile(currentCode, cond)
	fmt.Println("child", tmp)
	time.Sleep(1 * time.Second)
	if len(tmp) == 0 {
		return
	}
	for _, v := range tmp {
		for _, vv := range cond {
			if v.P_Code == vv.Code {
				*Ret = append(*Ret, v)
			}
		}
		find(v, cond, Ret)
	}
}

/*错误处理函数*/
func exitOnError(err error) {
	if err != nil {
		panic(err)
	}
}
