package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/yookoala/realpath"
)

func main() {
	yyDebug = 0
	yyErrorVerbose = true

	flag.Parse()

	for _, path := range flag.Args() {
		real, err := realpath.Realpath(path)
		checkErr(err)
		fmt.Printf("\n==> %s", real)

		src, _ := os.Open(string(real))
		rootnode := parse(src, real)
		fmt.Println(rootnode)
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
