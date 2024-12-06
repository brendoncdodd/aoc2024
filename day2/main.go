package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var total int
	var reports []string

	assessments := make(map[int]bool)
	assessment_desc := map[bool]string{true: "SAFE", false: "UNSAFE"}

	file_name := flag.String("i", "", "Filename for input")
	flag.Parse()

	if *file_name == "" {
		fmt.Println("No input filename defined.")
		return
	}

	in_file, err := os.Open(*file_name)
	if err != nil {
		fmt.Printf("%v", fmt.Errorf("Failed to open file.\n\t%w\n", err))
	}

	sc := bufio.NewScanner(in_file)

	for sc.Scan() {
		if err := sc.Err(); err != nil {
			fmt.Printf("%v", fmt.Errorf(
				"Scanning File:\n\t%s\n\t%w\n",
				sc.Text(),
				err,
			))
			continue
		}
		reports = append(reports, sc.Text())
	}

outer:
	for report, line := range reports {
		var levels []int

		for i, level_str := range strings.Split(line, " ") {
			level, err := strconv.Atoi(level_str)
			if err != nil {
				fmt.Println(fmt.Errorf(
					"Bad level line %d level %d.\n\t%w\n",
					report,
					i,
					err,
				))
				err = nil
				continue
			}
			levels = append(levels, level)
		}

		if len(levels) == 0 {
			fmt.Printf("No data at line %d.", report)
			continue
		}

		assessments[report] = true
		if levels[1] == levels[0] {
			assessments[report] = false
			break
		}
		if levels[1] > levels[0] {
			for i := 1; i < len(levels); i++ {
				if levels[i] <= levels[i-1] {
					assessments[report] = false
					continue outer
				}
				if levels[i] > levels[i-1]+3 {
					assessments[report] = false
					continue outer
				}
			}
		}
		if levels[1] < levels[0] {
			for i := 1; i < len(levels); i++ {
				if levels[i] >= levels[i-1] {
					assessments[report] = false
					continue outer
				}
				if levels[i] < levels[i-1]-3 {
					assessments[report] = false
					continue outer
				}
			}
		}
	}

	for i, report := range reports {
		if assessments[i] {
			total++
		}
		fmt.Printf("%d: %s\t%s\n",
			i,
			report,
			assessment_desc[assessments[i]],
		)
	}

	fmt.Printf("Total: %d\n", total)
}
