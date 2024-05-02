package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	name = flag.String("name", "", "Name to find")
	last = flag.Int64("last", 0, "Limit to last X years")
	csvF = flag.Bool("csv", false, "Use CSV output")
	filF = flag.Bool("f", false, "Filter to results with sex 'F'")
	filM = flag.Bool("m", false, "Filter to results with sex 'M'")
)

func main() {
	flag.Parse()

	female, err := os.Open("data/ontario_top_baby_names_female_1917-2022_en_fr.csv")
	if err != nil {
		panic(err)
	}
	male, err := os.Open("data/ontario_top_baby_names_male_1917-2022_en_fr.csv")
	if err != nil {
		panic(err)
	}

	if len(*name) == 0 {
		flag.Usage()
		return
	}

	upperName := strings.ToUpper(*name)
	if *csvF != true {
		fmt.Printf("Parsing name: %s\n\n", upperName)
	}

	if *filF == true || (*filF == false && *filM == false) {
		read(female, upperName, "F")
	}

	if *filM == true || (*filF == false && *filM == false) {
		read(male, upperName, "M")
	}
}

func read(file *os.File, name string, sex string) {
	r := csv.NewReader(file)
	currentYear := time.Now().Year()

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		year, err := strconv.Atoi(record[0])
		if err != nil {
			continue
		}
		if *last > 0 && int64(currentYear-year) > *last {
			continue
		}
		if record[1] == name {
			if *csvF != true {
				fmt.Printf("[%s] Year %s (%s)\n", sex, record[0], record[2])
			} else {
				fmt.Printf("%s,%s,%s\n", sex, record[0], record[2])
			}
		}
	}
}
