package main

import (
	"fmt"
	"reflect"
)

func reflectType(v interface{}) {
	t := reflect.TypeOf(v)
	value := reflect.ValueOf(v)

	//t.String()
	fmt.Println(t.Name())       //Cat
	fmt.Println(t.Kind())       //struct
	fmt.Printf("type=%T \n", t) //*reflect.rtype
	fmt.Printf("value=%v", value)

}

type Cat struct {
	name string
}

func main() {
	//aa:=123
	//reflectType(aa)
	c1 := Cat{"小白"}

	reflectType(c1)
}
