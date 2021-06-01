package main

import (
	"fmt"
	"reflect"
)

func reflectTest01(b interface{}) {
	// 通过反射获取到传入变量的 type(类型), kind(类别), value(值)
	rTyp := reflect.TypeOf(b)
	fmt.Println("rTyp=", rTyp)

	rVal := reflect.ValueOf(b)
	fmt.Println("rVal=", rVal)
	fmt.Printf("rVal type is %T\n", rVal)

	// 将reflect.Value转换为int
	iV := rVal.Interface()
	num2 := iV.(int)
	fmt.Println("num2=", num2)
	fmt.Printf("num2 type is %T\n", num2)
}

func reflectTest02(b interface{}) {
	rTyp := reflect.TypeOf(b)
	fmt.Println("rTyp=", rTyp)
	rVal := reflect.ValueOf(b)

	iV := rVal.Interface().(Student)
	fmt.Printf("iV=%v, iV type is %T\n", iV, iV)
}

type Student struct {
	Name string
	Age  int
}

func main() {
	// var num int = 100
	// reflectTest01(num)

	var stu = Student{
		Name: "tom",
		Age:  12,
	}
	reflectTest02(stu)
}
