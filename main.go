package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/yookoala/realpath"
	php7parser "github.com/z7zmey/php-parser/php7parser"
)

func main() {
	flag.Parse()

	for _, path := range flag.Args() {
		real, err := realpath.Realpath(path)
		checkErr(err)

		err = filepath.Walk(real, func(path string, f os.FileInfo, err error) error {
			if !f.IsDir() && filepath.Ext(path) == ".php" {
				fmt.Printf("==> %s\n", path)

				src, _ := os.Open(string(path))
				nodes, comments, positions := php7parser.Parse(src, path)

				visitor := Dumper{
					Indent:    "  | ",
					Comments:  comments,
					Positions: positions,
				}
				nodes.Walk(visitor)
			}
			return nil
		})
		checkErr(err)
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
