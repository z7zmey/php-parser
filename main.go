package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sync"

	"github.com/pkg/profile"
	"github.com/yookoala/realpath"
	"github.com/z7zmey/php-parser/ast"
	"github.com/z7zmey/php-parser/parser"
	"github.com/z7zmey/php-parser/php7"
	"github.com/z7zmey/php-parser/printer"
	"github.com/z7zmey/php-parser/visitor"
)

var wg sync.WaitGroup
var usePhp5 *bool
var dumpType string
var dumpPath *bool
var profiler string
var withFreeFloating *bool
var showResolvedNs *bool
var printBack *bool

type file struct {
	path    string
	content []byte
}

type result struct {
	path   string
	parser parser.Parser
}

func main() {
	usePhp5 = flag.Bool("php5", false, "parse as PHP5")
	dumpPath = flag.Bool("path", false, "print filepath")
	withFreeFloating = flag.Bool("ff", false, "parse and show free floating strings")
	showResolvedNs = flag.Bool("r", false, "resolve names")
	printBack = flag.Bool("pb", false, "print AST back into the parsed file")
	flag.StringVar(&dumpType, "d", "", "dump format: [custom, go, json, pretty_json]")
	flag.StringVar(&profiler, "prof", "", "start profiler: [cpu, mem, trace]")

	flag.Parse()

	if len(flag.Args()) == 0 {
		flag.Usage()
		return
	}

	switch profiler {
	case "cpu":
		defer profile.Start(profile.ProfilePath("."), profile.NoShutdownHook).Stop()
	case "mem":
		defer profile.Start(profile.MemProfile, profile.ProfilePath("."), profile.NoShutdownHook).Stop()
	case "trace":
		defer profile.Start(profile.TraceProfile, profile.ProfilePath("."), profile.NoShutdownHook).Stop()
	}

	numCpu := runtime.GOMAXPROCS(0)

	fileCh := make(chan *file, numCpu)
	resultCh := make(chan result, numCpu)

	// run 4 concurrent parserWorkers
	for i := 0; i < numCpu; i++ {
		go parserWorker(fileCh, resultCh)
	}

	// run printer goroutine
	go printerWorker(resultCh)

	// process files
	processPath(flag.Args(), fileCh)

	// wait the all files done
	wg.Wait()
	close(fileCh)
	close(resultCh)
}

func processPath(pathList []string, fileCh chan<- *file) {
	for _, path := range pathList {
		real, err := realpath.Realpath(path)
		checkErr(err)

		err = filepath.Walk(real, func(path string, f os.FileInfo, err error) error {
			if !f.IsDir() && filepath.Ext(path) == ".php" {
				wg.Add(1)
				content, err := ioutil.ReadFile(path)
				checkErr(err)
				fileCh <- &file{path, content}
			}
			return nil
		})
		checkErr(err)
	}
}

func parserWorker(fileCh <-chan *file, r chan<- result) {
	var parserWorker parser.Parser

	parserWorker = php7.NewParser(nil)

	for {
		f, ok := <-fileCh
		if !ok {
			return
		}

		// if *usePhp5 {
		// 	parserWorker = php5.NewParser(f.content)
		// } else {
		// 	parserWorker = php7.NewParser(f.content)
		// }

		if *withFreeFloating {
			parserWorker.WithFreeFloating()
		}

		abstractSyntaxTree := &ast.AST{
			Positions: ast.NewPositionStorage(make([]ast.Position, 0, 1024)),
			Nodes:     ast.NewNodeStorage(make([]ast.Node, 0, 1024)),
		}

		parserWorker.Parse(f.content, abstractSyntaxTree)

		r <- result{path: f.path, parser: parserWorker}
	}
}

func printerWorker(r <-chan result) {
	var counter int

	for {
		res, ok := <-r
		if !ok {
			return
		}

		counter++

		if *dumpPath {
			fmt.Printf("==> [%d] %s\n", counter, res.path)
		}

		for _, e := range res.parser.GetErrors() {
			fmt.Println(e)
		}

		if *printBack {
			o := bytes.NewBuffer([]byte{})
			p := printer.NewPrinter(o)
			p.Print(res.parser.GetRootNode())

			err := ioutil.WriteFile(res.path, o.Bytes(), 0644)
			checkErr(err)
		}

		var nsResolver *visitor.NamespaceResolver
		if *showResolvedNs {
			nsResolver = visitor.NewNamespaceResolver()
			res.parser.GetRootNode().Walk(nsResolver)
		}

		switch dumpType {
		case "custom":
			dumper := &visitor.Dumper{
				Writer:     os.Stdout,
				Indent:     "| ",
				NsResolver: nsResolver,
			}
			res.parser.GetRootNode().Walk(dumper)
		case "json":
			dumper := &visitor.JsonDumper{
				Writer:     os.Stdout,
				NsResolver: nsResolver,
			}
			res.parser.GetRootNode().Walk(dumper)
		case "pretty_json":
			dumper := &visitor.PrettyJsonDumper{
				Writer:     os.Stdout,
				NsResolver: nsResolver,
			}
			res.parser.GetRootNode().Walk(dumper)
		case "go":
			dumper := &visitor.GoDumper{Writer: os.Stdout}
			res.parser.GetRootNode().Walk(dumper)
		}

		wg.Done()
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
