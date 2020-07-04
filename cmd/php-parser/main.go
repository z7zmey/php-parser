package main

import (
	"bytes"
	"flag"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"sync"
	"time"

	"github.com/pkg/profile"
	"github.com/yookoala/realpath"

	"github.com/z7zmey/php-parser/pkg/ast"
	"github.com/z7zmey/php-parser/pkg/ast/traverser"
	"github.com/z7zmey/php-parser/pkg/ast/visitor"
	"github.com/z7zmey/php-parser/pkg/errors"
	"github.com/z7zmey/php-parser/pkg/parser"
	"github.com/z7zmey/php-parser/pkg/printer"
)

var wg sync.WaitGroup
var phpVersion string
var profiler string
var dump *bool
var withFreeFloating *bool
var showResolvedNs *bool
var printBack *bool
var printPath *bool
var printErrors *bool
var printExecTime *bool

type file struct {
	path    string
	content []byte
}

type result struct {
	path     string
	rootNode ast.Vertex
	errors   []*errors.Error
}

func main() {
	start := time.Now()

	printExecTime = flag.Bool("time", false, "print execution time")
	withFreeFloating = flag.Bool("ff", false, "parse and show free floating strings")
	showResolvedNs = flag.Bool("r", false, "resolve names")
	printBack = flag.Bool("pb", false, "print AST back into the parsed file")
	printPath = flag.Bool("p", false, "print filepath")
	printErrors = flag.Bool("e", false, "print errors")
	dump = flag.Bool("d", false, "dump")
	flag.StringVar(&profiler, "prof", "", "start profiler: [cpu, mem, trace]")
	flag.StringVar(&phpVersion, "phpver", "7.4", "php version")

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

	elapsed := time.Since(start)
	if *printExecTime {
		log.Printf("took: %s", elapsed)
	}
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
	for {
		f, ok := <-fileCh
		if !ok {
			return
		}

		parserErrors := []*errors.Error{}
		cfg := parser.Config{
			ErrorHandlerFunc: func(e *errors.Error) {
				parserErrors = append(parserErrors, e)
			},
		}
		rootNode, err := parser.Parse(f.content, phpVersion, cfg)
		if err != nil {
			panic(err.Error())
		}

		r <- result{path: f.path, rootNode: rootNode, errors: parserErrors}
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

		if *printPath {
			_, _ = io.WriteString(os.Stderr, "==> [" + strconv.Itoa(counter) + "] " + res.path + "\n")
		}

		if *printErrors {
			for _, e := range res.errors {
				_, _ = io.WriteString(os.Stderr, "==> " + e.String() + "\n")
			}
		}

		if *printBack {
			o := bytes.NewBuffer([]byte{})
			p := printer.NewPrinter(o)
			p.Print(res.rootNode)

			err := ioutil.WriteFile(res.path, o.Bytes(), 0644)
			checkErr(err)
		}

		if *showResolvedNs {
			v := visitor.NewNamespaceResolver()
			t := traverser.NewDFS(v)
			t.Traverse(res.rootNode)
			for _, n := range v.ResolvedNames {
				_, _ = io.WriteString(os.Stderr, "===> " + n + "\n")
			}
		}

		if *dump == true {
			v := visitor.NewDump(os.Stdout)
			t := traverser.NewDFS(v)
			t.Traverse(res.rootNode)
		}

		wg.Done()
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
