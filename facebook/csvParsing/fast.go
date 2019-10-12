package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
)

type dinosaur struct {
	name      string
	legLen    string
	diet      string
	strideLen string
	stance    string
	speed     float64
}

func (d *dinosaur) computeSpeed() {
	slen, _ := strconv.ParseFloat(d.strideLen, 64)
	llen, _ := strconv.ParseFloat(d.legLen, 64)
	d.speed = ((slen / llen) - 1) * math.Sqrt(llen*g)
}

const g = 9.81 // gravitational const

func main() {
	m := make(map[string]*dinosaur)

	bipedals := []*dinosaur{}

	// parse dataset 2
	parseFile("dataset2.csv", func(record []string) {
		if record[2] != "bipedal" {
			return // ignore all but bipedals
		}
		d := &dinosaur{
			name:      record[0],
			strideLen: record[1],
			stance:    record[2],
		}
		m[record[0]] = d
		bipedals = append(bipedals, d)
	})

	// parse dataset 1, computing the speed
	parseFile("dataset1.csv", func(record []string) {
		d, ok := m[record[0]]
		if !ok {
			return // assume we dont care if they werent in dataset2
		}
		d.legLen = record[1]
		d.diet = record[2]
		d.computeSpeed()
	})

	// sort by fastest
	sort.Slice(bipedals, func(i, j int) bool {
		return bipedals[i].speed > bipedals[j].speed
	})

	// print out
	for _, d := range bipedals {
		fmt.Printf("%s\n", d.name)
	}
}

// parse CSV file
func parseFile(path string, handler func([]string)) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("could not open csv file: %s", err)
	}
	r := csv.NewReader(f)
	// chop off csv header row
	if _, err := r.Read(); err != nil {
		log.Fatal(err)
	}
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		handler(record)
	}
}
