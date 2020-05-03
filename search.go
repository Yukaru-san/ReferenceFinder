package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// Search searches files within given paths for the given values
func Search() {
	for _, path := range *searchPaths {

		filepath.Walk(path, func(subPath string, info os.FileInfo, err error) error {

			// Check if file should be ignored
			if info.IsDir() || IsIgnored(path, subPath, info) {
				return nil
			}

			// Read current file
			fileData, err := ioutil.ReadFile(subPath)
			if err != nil {
				fmt.Println(err.Error())
				return err
			}
			fileText := string(fileData)

			// Split fileText into lines
			lines := strings.Split(fileText, "\n")

			// Iterate through all lines
			result := findReferencesInLines(lines, *searchText)

			// Print results (if any)
			if len(result) > 0 {
				fmt.Println(subPath[len(path):])
				for i := 0; i < len(result); i++ {
					fmt.Printf("   > %d:%d\n", result[i][0], result[i][1])
				}

			}

			return nil
		})
	}
}

// Finds all lines mentioning the requested searchText
func findReferencesInLines(lines []string, searchText string) [][]int {
	var findings [][]int

	// Loop through all lines (slice entries)
	re := regexp.MustCompile(searchText)
	for l := 0; l < len(lines); l++ {

		// Find all references
		result := re.FindAllStringIndex(lines[l], -1)

		// Print results (if any)
		if len(result) > 0 {
			for i := 0; i < len(result); i++ {
				findings = append(findings, []int{l - 1, result[i][0]})
			}
		}
	}

	return findings
}
