package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
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
	filR = flag.Bool("r", false, "Filter with regex matching")
)

func main() {
	flag.Parse()

	//
	// Validate data and args
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
	//

	//
	// Set vars
	showF := *filF == true || (*filF == false && *filM == false)
	showM := *filM == true || (*filF == false && *filM == false)

	var filter *regexp.Regexp = nil
	var validName = *name
	//

	// Check for valid regex
	if f, err := regexp.Compile(fmt.Sprintf("(?i)%s", validName)); err != nil || *filF != true {
		// don't use regex if it fails to compile
		validName = strings.ToUpper(validName)

		if err != nil && *csvF != true {
			log.Printf("Failed to compile regex, reverting to string compare mode: %v\n", err)
		}
	} else {
		filter = f
		validName = filter.String()
	}

	if *csvF != true {
		log.Printf(
			"\nMATCHING NAME:\t%s\nUSE REGEX:\t%v\nSHOW MALE:\t%v\nSHOW FEMALE:\t%v\n\n",
			validName, filter != nil, showM, showF,
		)
	}

	if showF {
		read(female, "F", validName, filter)
	}

	if showM {
		read(male, "M", validName, filter)
	}
}

func read(file *os.File, sex string, name string, filter *regexp.Regexp) {
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

		if matchName(record[1], name, filter) {
			if *csvF != true {
				fmt.Printf("[%s] %s (%s)\n", sex, record[0], record[2])
			} else {
				fmt.Printf("%s,%s,%s,%s\n", sex, record[0], record[1], record[2])
			}
		}
	}
}

func matchName(record string, name string, filter *regexp.Regexp) bool {
	if filter == nil {
		return record == name
	} else {
		return filter.MatchString(record)
	}
}
