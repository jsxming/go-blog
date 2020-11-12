package main

import (
	"blog/pkg/util"
	"fmt"
	"time"
)

func main() {
	t := time.Unix(1598677860, 0)
	s := util.ToTimeString(t.Unix())

	fmt.Println(s)
}
