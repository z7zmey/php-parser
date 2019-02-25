package main

import (
	"bufio"
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
	"github.com/z7zmey/php-parser/parser"
	"github.com/z7zmey/php-parser/php5"
	"github.com/z7zmey/php-parser/php7"
	"github.com/z7zmey/php-parser/printer"
	"github.com/z7zmey/php-parser/visitor"
)

var wg sync.WaitGroup
var usePhp5 *bool
var dumpType string
var profiler string
var withFreeFloating *bool
var showResolvedNs *bool
var printBack *bool

type file struct {
	path    string
	content []byte
}

func main() {
	usePhp5 = flag.Bool("php5", false, "parse as PHP5")
	withFreeFloating = flag.Bool("ff", false, "parse and show free floating strings")
	showResolvedNs = flag.Bool("r", false, "resolve names")
	printBack = flag.Bool("pb", false, "print AST back into the parsed file")
	flag.StringVar(&dumpType, "d", "", "dump format: [custom, go, json, pretty_json]")
	flag.StringVar(&profiler, "prof", "", "start profiler: [cpu, mem, trace]")

	flag.Parse()

	switch profiler {
	case "cpu":
		defer profile.Start(profile.ProfilePath("."), profile.NoShutdownHook).Stop()
	case "mem":
		defer profile.Start(profile.MemProfile, profile.ProfilePath("."), profile.NoShutdownHook).Stop()
	case "trace":
		defer profile.Start(profile.TraceProfile, profile.ProfilePath("."), profile.NoShutdownHook).Stop()
	}

	numCpu := runtime.NumCPU()

	fileCh := make(chan *file, numCpu)
	resultCh := make(chan parser.Parser, numCpu)

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

func parserWorker(fileCh <-chan *file, result chan<- parser.Parser) {
	var parserWorker parser.Parser

	for {
		f, ok := <-fileCh
		if !ok {
			return
		}

		src := bytes.NewReader(f.content)

		if *usePhp5 {
			parserWorker = php5.NewParser(src, f.path)
		} else {
			parserWorker = php7.NewParser(src, f.path)
		}

		if *withFreeFloating {
			parserWorker.WithFreeFloating()
		}

		parserWorker.Parse()

		result <- parserWorker
	}
}

func printerWorker(result <-chan parser.Parser) {
	var counter int

	w := bufio.NewWriter(os.Stdout)

	for {
		parserWorker, ok := <-result
		if !ok {
			w.Flush()
			return
		}

		counter++

		fmt.Fprintf(w, "==> [%d] %s\n", counter, parserWorker.GetPath())

		for _, e := range parserWorker.GetErrors() {
			fmt.Fprintln(w, e)
		}

		if *printBack {
			o := bytes.NewBuffer([]byte{})
			p := printer.NewPrinter(o)
			p.Print(parserWorker.GetRootNode())

			err := ioutil.WriteFile(parserWorker.GetPath(), o.Bytes(), 0644)
			checkErr(err)
		}

		var nsResolver *visitor.NamespaceResolver
		if *showResolvedNs {
			nsResolver = visitor.NewNamespaceResolver()
			parserWorker.GetRootNode().Walk(nsResolver)
		}

		switch dumpType {
		case "custom":
			dumper := &visitor.Dumper{
				Writer:     os.Stdout,
				Indent:     "| ",
				NsResolver: nsResolver,
			}
			parserWorker.GetRootNode().Walk(dumper)
		case "json":
			dumper := &visitor.JsonDumper{
				Writer:     os.Stdout,
				NsResolver: nsResolver,
			}
			parserWorker.GetRootNode().Walk(dumper)
		case "pretty_json":
			dumper := &visitor.PrettyJsonDumper{
				Writer:     os.Stdout,
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
