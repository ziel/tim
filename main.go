package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ziel/tim/control"
)

func handleArgs() []string {
	flag.Usage = func() {
		fmt.Printf("Usage: %s [options] file1 file2 [file3]\n", os.Args[0])
		flag.PrintDefaults()
	}

	if len(os.Args) == 2 && os.Args[1] == "help" {
		flag.Usage()
		os.Exit(0)
	}

	flag.Parse()
	return flag.Args()
}

func main() {
	errlog := log.New(os.Stderr, "Erp! ", 0)
	paths := handleArgs()

	if err := control.Init(paths); err != nil {
		errlog.Fatal(err)
	}
}
