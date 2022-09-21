package main

import "fmt"

//生物接口
type IBeing interface {
	Live()
	Dead()
}

//动物接口嵌入生物接口
type IAnimal interface {
	IBeing
	Hunting()
}

//植物接口嵌入生物接口
type IPlant interface {
	IBeing
	Growing()
}

type Tiger struct{}

func (t *Tiger) Live() {
	fmt.Println("老虎活着称大王！！！")
}
func (t *Tiger) Dead() {
	fmt.Println("老虎战斗死了！！！")
}
func (t *Tiger) Hunting() {
	fmt.Println("老虎捕猎！！！")
}

type Flower struct{}

func (t *Flower) Live() {
	fmt.Println("花儿享受阳光！！！")
}
func (t *Flower) Dead() {
	fmt.Println("花儿落下死了！！！")
}
func (t *Flower) Growing() {
	fmt.Println("花儿茁壮成长！！！")
}

//声明一个interface{}的切片
var earth []interface{}

func addBeing(b interface{}) {
	earth = append(earth, b)
}
func main() {
	tiger := new(Tiger)
	//flower := new(Flower)

	tiger.Live()
	tiger.Hunting()
	tiger.Dead()
	//addBeing(tiger)
	//addBeing(flower)
	//
	////遍历
	//for _, being := range earth {
	//	if animal, ok := being.(IAnimal); ok {
	//		animal.Live()
	//		//如果是动物则捕猎
	//		animal.Hunting()
	//		animal.Dead()
	//	}
	//	if plant, ok := being.(IPlant); ok {
	//		plant.Live()
	//		//如果是植物则成长
	//		plant.Growing()
	//		plant.Dead()
	//
	//	}
	//
	//}
}
