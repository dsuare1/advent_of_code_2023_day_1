package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("calibration_values.txt")
	if err != nil {
		log.Fatalln("Error opening file: ", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0

	for scanner.Scan() {
		re := regexp.MustCompile(`(?i)[a-z]+`)
		digitsOnly := re.ReplaceAll(scanner.Bytes(), []byte(""))
		if err != nil {
			log.Fatalln("Error converting to int: ", err)
		}

		if len(digitsOnly) == 1 {
			asStr := string(digitsOnly)
			combined := fmt.Sprintf("%s%s", asStr, asStr)
			combinedAsInts, err := strconv.Atoi(combined)
			if err != nil {
				log.Fatalln("Error converting combined to int: ", err)
			}
			sum += combinedAsInts
		} else {
			combined := fmt.Sprintf("%s%s", digitsOnly[:1], digitsOnly[len(digitsOnly)-1:])
			combinedAsInts, err := strconv.Atoi(combined)
			if err != nil {
				log.Fatalln("Error converting combined to int: ", err)
			}

			sum += combinedAsInts
		}
	}

	log.Println("Done; sum: ", sum)
}
