package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/karrick/godirwalk"
	"github.com/pkg/profile"
	"github.com/yookoala/realpath"
	"github.com/z7zmey/php-parser/parser"
	"github.com/z7zmey/php-parser/php5"
	"github.com/z7zmey/php-parser/php7"
	"github.com/z7zmey/php-parser/visitor"
)

var wg sync.WaitGroup
var usePhp5 *bool
var dumpType string
var profiler string
var showComments *bool
var showResolvedNs *bool

func main() {
	usePhp5 = flag.Bool("php5", false, "parse as PHP5")
	showComments = flag.Bool("c", false, "show comments")
	showResolvedNs = flag.Bool("r", false, "resolve names")
	flag.StringVar(&dumpType, "d", "", "dump format: [custom, go, json, pretty_json]")
	flag.StringVar(&profiler, "prof", "", "start profiler: [cpu, mem]")

	flag.Parse()

	switch profiler {
	case "cpu":
		defer profile.Start(profile.ProfilePath("."), profile.NoShutdownHook).Stop()
	case "mem":
		defer profile.Start(profile.MemProfile, profile.ProfilePath("."), profile.NoShutdownHook).Stop()
	}

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

		s, err := os.Stat(real)
		checkErr(err)

		if !s.IsDir() {
			wg.Add(1)
			pathCh <- real
		} else {
			godirwalk.Walk(real, &godirwalk.Options{
				Unsorted: true,
				Callback: func(osPathname string, de *godirwalk.Dirent) error {
					if !de.IsDir() && filepath.Ext(osPathname) == ".php" {
						wg.Add(1)
						pathCh <- osPathname
					}
					return nil
				},
				ErrorCallback: func(osPathname string, err error) godirwalk.ErrorAction {
					return godirwalk.SkipNode
				},
			})
		}
	}
}

func parserWorker(pathCh <-chan string, result chan<- parser.Parser) {
	var parserWorker parser.Parser

	for {
		path, ok := <-pathCh
		if !ok {
			return
		}

		src, err := os.Open(path)
		checkErr(err)

		if *usePhp5 {
			parserWorker = php5.NewParser(src, path)
		} else {
			parserWorker = php7.NewParser(src, path)
		}

		parserWorker.Parse()

		src.Close()

		result <- parserWorker
	}
}

func printer(result <-chan parser.Parser) {
	var counter int

	for {
		parserWorker, ok := <-result
		if !ok {
			return
		}

		counter++

		fmt.Printf("==> [%d] %s\n", counter, parserWorker.GetPath())

		for _, e := range parserWorker.GetErrors() {
			fmt.Println(e)
		}

		var nsResolver *visitor.NamespaceResolver
		if *showResolvedNs {
			nsResolver = visitor.NewNamespaceResolver()
			parserWorker.GetRootNode().Walk(nsResolver)
		}

		var comments parser.Comments
		if *showComments {
			comments = parserWorker.GetComments()
		}

		switch dumpType {
		case "custom":
			dumper := &visitor.Dumper{
				Writer:     os.Stdout,
				Indent:     "| ",
				Comments:   comments,
				NsResolver: nsResolver,
			}
			parserWorker.GetRootNode().Walk(dumper)
		case "json":
			dumper := &visitor.JsonDumper{
				Writer:     os.Stdout,
				Comments:   comments,
				NsResolver: nsResolver,
			}
			parserWorker.GetRootNode().Walk(dumper)
		case "pretty_json":
			dumper := &visitor.PrettyJsonDumper{
				Writer:     os.Stdout,
				Comments:   comments,
				NsResolver: nsResolver,
			}
			parserWorker.GetRootNode().Walk(dumper)
		case "go":
			dumper := &visitor.GoDumper{Writer: os.Stdout}
			parserWorker.GetRootNode().Walk(dumper)
		}

		wg.Done()
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
