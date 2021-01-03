package main

import (
	"github.com/ed-fx/go-soft4fx/internal/path"
	"github.com/ed-fx/go-soft4fx/internal/simulator"
	"github.com/ed-fx/go-soft4fx/internal/simulator/analyze"
	"github.com/ed-fx/go-soft4fx/internal/simulator/analyze/export"
	"github.com/ed-fx/go-soft4fx/internal/simulator/parser"
	"log"
)

func main() {
	htmlFiles := path.ResultHtmlFiles()
	if len(htmlFiles) == 0 {
		log.Println("No soft4fx result files found")
		return
	}

	sims := parse(htmlFiles)
	results, err := analyze.Analyze(sims)
	if err != nil {
		log.Fatal(err)
	}

	err = export.AggregateResult(results)
	if err != nil {
		log.Fatal(err)
	}
}

func parse(htmlFiles []string) (sims []*simulator.Simulator) {
	sims = make([]*simulator.Simulator, len(htmlFiles))
	var i = 0
	for _, htmlFile := range htmlFiles {
		if sim, err := parser.Parse(htmlFile); err != nil {
			log.Fatal(err)
			return
		} else {
			sims[i] = sim
			i++
		}
	}
	return
}
