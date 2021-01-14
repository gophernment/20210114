package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("./oscar_age_male.csv")
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(f)

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	const nameColum = 3
	nameCount := map[string]int{}

	for _, record := range records {
		nameCount[record[nameColum]]++
	}

	for name, count := range nameCount {
		if count > 1 {
			fmt.Println(name, count)
		}
	}
}
