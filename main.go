package main

import (
	"forex/go-soft4fx/internal/path"
	"forex/go-soft4fx/internal/simulator/analyze"
	"forex/go-soft4fx/internal/simulator/analyze/export"
	"forex/go-soft4fx/internal/simulator/parser"
	"log"
)

func main() {
	htmlFiles := path.ResultHtmlFiles()
	for _, htmlFile := range htmlFiles {
		sim, err := parser.Parse(htmlFile)
		if err != nil {
			log.Fatal(err)
		}
		r, err := analyze.Analyze(sim)
		if err != nil {
			log.Fatal(err)
		}
		if err = export.Export(r); err != nil {
			log.Fatal(err)
		}
	}
}
