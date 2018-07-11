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
var withMeta *bool
var showResolvedNs *bool

type file struct {
	path    string
	content []byte
}

func main() {
	usePhp5 = flag.Bool("php5", false, "parse as PHP5")
	withMeta = flag.Bool("meta", false, "show meta")
	showResolvedNs = flag.Bool("r", false, "resolve names")
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
	go printer(resultCh)

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

		s, err := os.Stat(real)
		checkErr(err)

		if !s.IsDir() {
			wg.Add(1)
			content, err := ioutil.ReadFile(real)
			checkErr(err)
			fileCh <- &file{real, content}
		} else {
			godirwalk.Walk(real, &godirwalk.Options{
				Unsorted: true,
				Callback: func(osPathname string, de *godirwalk.Dirent) error {
					if !de.IsDir() && filepath.Ext(osPathname) == ".php" {
						wg.Add(1)
						content, err := ioutil.ReadFile(osPathname)
						checkErr(err)
						fileCh <- &file{osPathname, content}
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

		if *withMeta {
			parserWorker.WithMeta()
		}

		parserWorker.Parse()

		result <- parserWorker
	}
}

func printer(result <-chan parser.Parser) {
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
