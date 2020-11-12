package main

import "fmt"

/**
切片的便利
 */
func rangeSlice(list []int)  {
	for index,value :=range list{
		fmt.Printf("index=%d,value=%d",index,value)
		fmt.Println()
	}
}

func main() {
	b:=make([]int,0,30)
	b = append(b,10,20)
	rangeSlice(b)

}