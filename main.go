package main

import (
	"fmt"

	"github.com/ziel/tim/model"
	"github.com/ziel/tim/view"
)

func main() {
	fmt.Println("NARGS:", flag.NArgs())
	m := model.Fake2FileModel()
	view.Display(m)
}
