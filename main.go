package main

import (
	"fmt"

	omikuji "github.com/saki-engineering/omikuji-go"
)

func main() {
	box := omikuji.New()
	defer box.Finish()

	fmt.Println("DrawStatus()")
	if s, err := box.DrawStatus(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(s)
	}

	fmt.Println("DrawPaper()")
	if p, err := box.DrawPaper(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(p)
	}
}
