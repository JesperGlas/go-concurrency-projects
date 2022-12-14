package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
)

func main() {
	diffMap := qwrtyDist()
	a := "dinamck"
	b := "dynamic"
	result := difference(a, b, diffMap)
	fmt.Printf("Difference between %s and %s: %d\n", a, b, result)
}

// Variables for operation codes
var match int = 0
var skip_a int = 1
var skip_b int = 2

func subsolutionMatrix(a string, b string, R map[string]map[string]int) [][]int {
	/**
		* Function creates a subsolution matrix for the minimum difference between
		*	string a and string b.
		*
		* Args:
		*	a (string) - One of the strings to be compared
		*	b (string) - The other string to be compared
	    *	R (map[string]map[string]int) - Map containing difference cost between chars
		*		Ex. R["-"]["r"] = 2
	*/

	// Get length of strin a (n) and b (m)
	n := len([]rune(a))
	m := len([]rune(b))

	// Initialize subsolution matrix with dimensions n x m
	s := make([][]int, n+1)
	for i := range s {
		s[i] = make([]int, m+1)
	}

	// Fill subsolution matrix
	for i := 0; i < n+1; i++ {
		for j := 0; j < m+1; j++ {
			// Solves subproblem a[:i] - b[:j]

			if i == 0 && j == 0 {
				// First element is considered a match
				s[i][j] = match
			} else if i == 0 {
				// If a is an empty string we can only skip it
				s[0][j] = ((cost(s[0][j-1]) + R["-"][string(b[j-1])]) << 2) + skip_a
			} else if j == 0 {
				// If b is an empty string we can only skip it
				s[i][0] = ((cost(s[i-1][0]) + R[string(a[i-1])]["-"]) << 2) + skip_b
			} else {
				// Complex case, chose minimum value

				// Set subsolution to max int to ensure it's discarded in min check
				s[i][j] = math.MaxInt

				// Possible operations and there cost
				candidates := [3]int{
					((cost(s[i-1][j-1]) + R[string(a[i-1])][string(b[j-1])]) << 2) + match,
					((cost(s[i][j-1]) + R["-"][string(b[j-1])]) << 2) + skip_a,
					((cost(s[i-1][j]) + R[string(a[i-1])]["-"]) << 2) + skip_b,
				}

				// Set subsolution to candidate with min cost
				for _, v := range candidates {
					if cost(v) < cost(s[i][j]) {
						s[i][j] = v
					}
				}
			}
		}
	}
	return s
}

func difference(a string, b string, R map[string]map[string]int) int {
	subsolutions := subsolutionMatrix(a, b, R)
	return cost(subsolutions[len([]rune(a))][len([]rune(b))])
}

func cost(data int) int {
	return data >> 2
}

func op(data int) int {
	return data % 4
}

func printSubsolutionsMatrix(matrix map[string]map[string]int) {
	fmt.Println("Subsolution Matrix:")
	for _, row := range matrix {
		fmt.Printf("[ ")
		for _, val := range row {
			fmt.Printf("%d ", cost(val))
		}
		fmt.Printf("]\n")
	}
}

func qwrtyDist() map[string]map[string]int {
	// Load json file
	filePath := "qwerty.json"
	byteJSON, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal("[FATAL ERROR]\nCould not load json file (Run gen_qwerty.py in strmtch folder to generate one!)\n", err)
	}

	// Unpack json file
	var charMap map[string]map[string]int
	err = json.Unmarshal(byteJSON, &charMap)
	if err != nil {
		fmt.Println("Could not unpack json data..")
		panic(err)
	}
	return charMap
}
