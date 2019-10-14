package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
)

type planet struct {
	name     string
	pType    string
	distance string
	mass     string
}

func main() {
	planets := make(map[string]*planet)
	jovians := []*planet{}

	parseCSV("dataset2.csv", func(rec []string) {
		// exclude all but jovian
		if rec[2] != "Jovian" {
			return
		}
		if _, ok := planets[rec[0]]; !ok {
			p := &planet{
				name:     rec[0],
				distance: rec[1],
				pType:    rec[2],
			}
			// add same planet to map and slice
			planets[rec[0]] = p
			jovians = append(jovians, p)
		}
	})

	parseCSV("dataset1.csv", func(rec []string) {
		// assume all we care about will exist
		if p, ok := planets[rec[0]]; ok {
			p.mass = rec[1]
		}
	})

	// we sort the slice in order of largest mass
	sort.Slice(jovians, func(i, j int) bool {
		f1, _ := strconv.ParseFloat(jovians[i].mass, 64)
		f2, _ := strconv.ParseFloat(jovians[j].mass, 64)
		return f1 > f2
	})

	// print out slice
	for _, p := range jovians {
		fmt.Printf("%s: (%s x 10^24 kg)\n", p.name, p.mass)
	}
}

func parseCSV(path string, handler func(rec []string)) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("could not open file %s: %s", path, err)
	}
	r := csv.NewReader(f)
	for {
		record, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break // exit condition
			}
			log.Fatalf("error reading from CSV file: %s", err)
		}
		handler(record)
	}
}
