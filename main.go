package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ziel/tim/control"
	"github.com/ziel/tim/model"
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
	filePaths := handleArgs()

	m, err := model.Factory(filePaths)

	if err != nil {
		errlog.Fatal(err)
	}

	cerr := control.Init(m)

	if cerr != nil {
		errlog.Fatal(cerr)
	}
}
