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

			// Split fileText into lines (\r\n = windows, \n = linux/mac)
			lines := strings.Split(strings.Join(strings.Split(fileText, "\r\n"), "\n"), "\n")

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

// Replace replaces file contents within given paths for the given values
func Replace(inputs []string) {

	// Used for output
	replacedSomething := false

	// Always exclude these on replacements (safety first)
	excludedFiles = append(excludedFiles, []string{".exe", ".git", "LICENSE"}...)

	// Initial input check
	if len(inputs) != 2 {
		fmt.Println("Replace usage:\nreplace <oldText> <newText>")
		return
	}

	for _, path := range *replacePaths {

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
			lines := strings.Split(strings.Join(strings.Split(fileText, "\r\n"), "\n"), "\n")

			// Iterate through all lines
			findings := findReferencesInLines(lines, inputs[0])

			// Replace and give feedback
			if len(findings) > 0 {

				// Feedback
				if len(findings) == 1 {
					fmt.Println(subPath[len(path):], "- replaced 1 occurrence")
				} else {
					fmt.Println(subPath[len(path):], "- replaced", len(findings), "occurrences")
				}

				// Replace
				for i := 0; i < len(findings); i++ {
					// Original line excluding the part to remove
					before := lines[findings[i][0]-1][:findings[i][1]]
					after := lines[findings[i][0]-1][findings[i][1]+len(inputs[0]):]

					// Newly built file
					lines[i] = before + inputs[1] + after
					newFile := strings.Join(lines, "\n")
					_ = newFile

					// Save
					ioutil.WriteFile(subPath, []byte(newFile), info.Mode())
				}
				replacedSomething = true
			}

			return nil
		})
	}

	if !replacedSomething {
		fmt.Println("No files matched your search.")
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
				findings = append(findings, []int{l + 1, result[i][0]})
			}
		}
	}

	return findings
}
