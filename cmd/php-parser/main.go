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

	"github.com/z7zmey/php-parser/pkg/errors"
	"github.com/z7zmey/php-parser/pkg/traverser"
	"github.com/z7zmey/php-parser/pkg/visitor"

	"github.com/pkg/profile"
	"github.com/yookoala/realpath"
	"github.com/z7zmey/php-parser/pkg/parser"
)

var wg sync.WaitGroup
var usePhp5 *bool
var dump *bool
var dumpPath *bool
var profiler string
var showResolvedNs *bool
var printBack *bool
var withTokens *bool

type file struct {
	path    string
	content []byte
}

type result struct {
	path      string
	traverser traverser.Traverser
	errors    []*errors.Error
}

func main() {
	ballast := make([]byte, 100<<20)

	usePhp5 = flag.Bool("php5", false, "parse as PHP5")
	dump = flag.Bool("d", false, "dump ast")
	dumpPath = flag.Bool("p", false, "print filepath")
	showResolvedNs = flag.Bool("r", false, "resolve names")
	withTokens = flag.Bool("t", false, "parse free-floating tokens")
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

	runtime.KeepAlive(ballast)
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

	if *usePhp5 {
		parserWorker = parser.NewPHP7Parser()
	} else {
		parserWorker = parser.NewPHP7Parser()
	}

	if *withTokens {
		parserWorker = parserWorker.WithTokens()
	}

	for {
		f, ok := <-fileCh
		if !ok {
			return
		}

		t, e := parserWorker.Parse(f.content)

		r <- result{path: f.path, traverser: t, errors: e}
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

		for _, e := range res.errors {
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

		linker := &visitor.Linker{}

		if *dump {
			res.traverser.Traverse(linker)

			buf, err := json.MarshalIndent(linker.GetRoot(), "", "  ")
			checkErr(err)
			_, err = os.Stdout.Write(buf)
			checkErr(err)
		}

		parser.Reuse(res.traverser)

		wg.Done()
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
