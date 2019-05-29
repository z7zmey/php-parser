package main

import (
	"encoding/json"
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
	"github.com/z7zmey/php-parser/graph"
	"github.com/z7zmey/php-parser/parser"
	"github.com/z7zmey/php-parser/parser/php7"
)

var wg sync.WaitGroup
var usePhp5 *bool
var dump *bool
var dumpPath *bool
var profiler string
var showResolvedNs *bool
var printBack *bool

type file struct {
	path    string
	content []byte
}

type result struct {
	path   string
	parser parser.Parser
	ast    *graph.AST
}

func main() {
	usePhp5 = flag.Bool("php5", false, "parse as PHP5")
	dump = flag.Bool("d", false, "dump ast")
	dumpPath = flag.Bool("p", false, "print filepath")
	showResolvedNs = flag.Bool("r", false, "resolve names")
	printBack = flag.Bool("pb", false, "print AST back into the parsed file")
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

	parserWorker = php7.NewParser()

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

		abstractSyntaxTree := &graph.AST{
			Nodes:     &graph.NodeStorage{},
			Edges:     &graph.EdgeStorage{},
			Positions: &graph.PositionStorage{},
			Tokens:    &graph.TokenStorage{},
		}

		parserWorker.Parse(f.content, abstractSyntaxTree)

		r <- result{path: f.path, parser: parserWorker, ast: abstractSyntaxTree}
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

		// if *printBack {
		// 	o := bytes.NewBuffer([]byte{})
		// 	p := printer.NewPrinter(o)
		// 	p.Print(res.parser.GetRootNode())

		// 	err := ioutil.WriteFile(res.path, o.Bytes(), 0644)
		// 	checkErr(err)
		// }

		// var nsResolver *visitor.NamespaceResolver
		// if *showResolvedNs {
		// 	nsResolver = visitor.NewNamespaceResolver()
		// 	res.parser.GetRootNode().Walk(nsResolver)
		// }

		if *dump {
			buf, err := json.MarshalIndent(res.ast.Nested(), "", "  ")
			checkErr(err)
			_, err = os.Stdout.Write(buf)
			checkErr(err)
		}

		wg.Done()
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
