package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/yookoala/realpath"
	"github.com/z7zmey/php-parser/parser"
	"github.com/z7zmey/php-parser/php5"
	"github.com/z7zmey/php-parser/php7"
	"github.com/z7zmey/php-parser/visitor"
)

var wg sync.WaitGroup
var usePhp5 *bool

func main() {
	usePhp5 = flag.Bool("php5", false, "use PHP5 parserWorker")
	flag.Parse()

	pathCh := make(chan string)
	resultCh := make(chan parser.Parser)

	// run 4 concurrent parserWorkers
	for i := 0; i < 4; i++ {
		go parserWorker(pathCh, resultCh)
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

func parserWorker(pathCh <-chan string, result chan<- parser.Parser) {
	var parserWorker parser.Parser

	for {
		path := <-pathCh
		src, _ := os.Open(path)

		if *usePhp5 {
			parserWorker = php5.NewParser(src, path)
		} else {
			parserWorker = php7.NewParser(src, path)
		}

		parserWorker.Parse()
		result <- parserWorker
	}
}

func printer(result <-chan parser.Parser) {
	for {
		parserWorker := <-result
		fmt.Printf("==> %s\n", parserWorker.GetPath())

		for _, e := range parserWorker.GetErrors() {
			fmt.Println(e)
		}

		nsResolver := visitor.NewNamespaceResolver()
		parserWorker.GetRootNode().Walk(nsResolver)

		dumper := visitor.Dumper{
			Writer:     os.Stdout,
			Indent:     "  | ",
			Comments:   parserWorker.GetComments(),
			Positions:  parserWorker.GetPositions(),
			NsResolver: nsResolver,
		}
		parserWorker.GetRootNode().Walk(dumper)
		wg.Done()
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
