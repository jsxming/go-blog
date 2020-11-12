package main

import (
	"blog/pkg/util"
	"encoding/json"
	"fmt"
	"time"
)

type AA struct {
	Test  string `json:"test"`
	Test2 string `json:"Test2"`
	Empty string `json:"empty,omitempty"` //omitempty 表示忽略空字段
}

type MyTime time.Time

func (t *MyTime) UnmarshalJSON(data []byte) error {
	str := string(data)
	if str == "null" || str == "" {
		return nil
	}
	tempTime, err := time.Parse(util.TimeLayout, str)
	*t = MyTime(tempTime)
	return err
}

//func (t *MyTime) MarshalJSON()([]byte,error)  {
//
//}

func main() {
	a := AA{
		Test:  "sdf",
		Test2: "bbb",
	}
	ad, _ := json.Marshal(a)

	var bb AA

	json.Unmarshal(ad, &bb)

	fmt.Println(string(ad))
	fmt.Println(bb)
}
