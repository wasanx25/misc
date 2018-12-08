package main

import (
	"fmt"
)

type Gopher struct {
	Gopher string `json:"gopher"`
}
/*
aaa

bbb
 */

// TODO これはなんかをやっている
//aa
func main() {
	const gopher = "GOPHER"
	gogopher := GOPHER()
	gogopher.Gopher = gopher
	fmt.Println(gogopher)
}

// FIXME あかんやつ
// fixme why
func GOPHER() (gopher *Gopher) {
	gopher = &Gopher{Gopher: "gopher"}
	return
}
