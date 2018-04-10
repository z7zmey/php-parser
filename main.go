package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/yookoala/realpath"
	"github.com/z7zmey/php-parser/php5"
	"github.com/z7zmey/php-parser/php7"
	"github.com/z7zmey/php-parser/visitor"
)

var wg sync.WaitGroup
var usePhp5 *bool

func main() {
	usePhp5 = flag.Bool("php5", false, "use PHP5 parser")
	flag.Parse()

	pathCh := make(chan string)
	resultCh := make(chan Parser)

	// run 4 concurrent parsers
	for i := 0; i < 4; i++ {
		go parser(pathCh, resultCh)
	}

	// run printer goroutine
	go printer(resultCh)

	// process files
	processPath(flag.Args(), pathCh)

	// wait the all files done
	wg.Wait()
	close(pathCh)
	close(resultCh)
}

func processPath(pathList []string, pathCh chan<- string) {
	for _, path := range pathList {
		real, err := realpath.Realpath(path)
		checkErr(err)

		err = filepath.Walk(real, func(path string, f os.FileInfo, err error) error {
			if !f.IsDir() && filepath.Ext(path) == ".php" {
				wg.Add(1)
				pathCh <- path
			}
			return nil
		})
		checkErr(err)
	}
}

func parser(pathCh <-chan string, result chan<- Parser) {
	var parser Parser

	for {
		path := <-pathCh
		src, _ := os.Open(path)

		if *usePhp5 {
			parser = php5.NewParser(src, path)
		} else {
			parser = php7.NewParser(src, path)
		}

		parser.Parse()
		result <- parser
	}
}

func printer(result <-chan Parser) {
	for {
		parser := <-result
		fmt.Printf("==> %s\n", parser.GetPath())

		for _, e := range parser.GetErrors() {
			fmt.Println(e)
		}

		nsResolver := visitor.NewNamespaceResolver()
		parser.GetRootNode().Walk(nsResolver)

		dumper := visitor.Dumper{
			Writer:     os.Stdout,
			Indent:     "  | ",
			Comments:   parser.GetComments(),
			Positions:  parser.GetPositions(),
			NsResolver: nsResolver,
		}
		parser.GetRootNode().Walk(dumper)
		wg.Done()
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
