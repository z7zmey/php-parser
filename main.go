package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/yookoala/realpath"
	"github.com/z7zmey/php-parser/parser"
)

func main() {
	flag.Parse()

	for _, path := range flag.Args() {
		real, err := realpath.Realpath(path)
		checkErr(err)
		fmt.Printf("==> %s\n", real)

		src, _ := os.Open(string(real))
		rootnode := parser.Parse(src, real)

		rootnode.Walk(dumper{"  | "})
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
