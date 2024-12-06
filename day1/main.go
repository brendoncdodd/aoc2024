package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var (
		total  int
		lefts  []int
		rights []int
	)

	file_name := flag.String("i", "", "Filename for input")
	flag.Parse()

	if *file_name == "" {
		fmt.Println("No input filename defined.")
		return
	}

	in_file, err := os.Open(*file_name)
	if err != nil {
		fmt.Printf("%v", fmt.Errorf("Failed to open file.\n\t%w", err))
	}

	input, err := io.ReadAll(in_file)
	if err != nil {
		fmt.Printf("%v", fmt.Errorf("Failed to read file contents.\n\t%w", err))
	}

outer:
	for idx, line := range strings.Split(string(input), "\n") {
		split := strings.Split(line, " ")
		left := 0
		right := 0

		for _, item := range split {
			if item == "" {
				continue
			}

			if left == 0 {
				left, err = strconv.Atoi(item)
				if err != nil {
					fmt.Printf("%v", fmt.Errorf(
						"Failed to get left value at line %d\n\t%w",
						idx,
						err,
					))
					continue outer
				}
				lefts = append(lefts, left)
				fmt.Printf("%d\t", left)
			} else {
				right, err = strconv.Atoi(item)
				if err != nil {
					fmt.Printf("%v", fmt.Errorf(
						"Failed to get right value at line %d\n\t%w",
						idx,
						err,
					))
					continue outer
				}
				rights = append(rights, right)
				fmt.Printf("%d\n", right)
			}

		}
	}

	sort.Ints(lefts)
	sort.Ints(rights)

	fmt.Println("Sorted:")

	for idx, left := range lefts {
		diff := rights[idx] - left
		fmt.Printf("%d\t%d\t%d\n", left, rights[idx], diff)
		if diff < 0 {
			diff = diff * -1
		}
		total += diff
	}

	fmt.Println("Total: ", total)
}
