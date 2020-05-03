package main

// Useage:
// referenceFinder search C:\\Users\\Me\Desktop\\FolderToExplore searchText --search-sub -e ".exe, .bat specificFileName"
// referenceFinder search "C:\Users\Me\Documents\Coding_Projects\Golang\project\resources\app" --search-sub --excludeDirs "main"
// referenceFinder search "C:\Users\ME\Documents\Coding_Projects\Golang\" --find "func main()" --search-sub -f ".exe" -d "git"

import (
	"os"
	"path/filepath"
	"strings"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	// App
	app = kingpin.New("referenceFinder", "Search across multiple files")

	// Global flag
	searchSubdirectories = app.Flag("search-sub", "Search subdirectories aswell").Bool()
	excludedFileFlag     = app.Flag("excludeFiles", "Exclude files e.g -e \".exe, .bat, fileName1.txt\"").Short('f').String()
	excludedFolderFlag   = app.Flag("excludeDirs", "Exclude specific subdirs e.g -e \"folderName1, folderName2\"").Short('d').String()

	// Search flags
	searchFlag  = app.Command("search", "Search within files")
	searchPaths = searchFlag.Arg("filePath", "Path to the directories you want to search through").Strings()
	searchText  = searchFlag.Flag("find", "Text you want to find").Short('f').Short('s').Required().String()

	// Replace flags TODO
	replaceFlag = app.Command("replace", "Replace within files")

	// Search variables
	excludedFiles []string
	excludedDirs  []string
)

func main() {
	app.HelpFlag.Short('h')

	// Prase cli flags
	parsed := kingpin.MustParse(app.Parse(os.Args[1:]))

	// Parse excluded lists
	excludedFiles = strings.Split(strings.ReplaceAll(*excludedFileFlag, ", ", ","), ",")
	if excludedFiles[0] == "" {
		excludedFiles = nil
	}
	excludedDirs = strings.Split(strings.ReplaceAll(*excludedFolderFlag, ", ", ","), ",")
	if excludedDirs[0] == "" {
		excludedDirs = nil
	}

	// Execute selected command
	switch parsed {
	case "search":
		Search()
	case "replace":
		Replace()
	}

}

// Replace replaces file contents within given paths for the given values
func Replace() {

}

// IsIgnored returns true if the file or folder should be ignored
func IsIgnored(basePath string, subPath string, info os.FileInfo) bool {

	// Subdirectories
	p, f := filepath.Split(subPath)
	if !*searchSubdirectories {
		// long term for checking whether the path is still within the selected dir
		if basePath[len(basePath)-1:len(basePath)] == string(filepath.Separator) && p != basePath || basePath[len(basePath)-1:len(basePath)] != string(filepath.Separator) && p[:len(p)-1] != basePath {
			return true
		}
	}

	// File searching
	for i := 0; i < len(excludedFiles); i++ {
		if strings.Contains(f, excludedFiles[i]) {
			return true
		}
	}
	// Ignored directory?
	for i := 0; i < len(excludedDirs); i++ {
		if strings.Contains(p, excludedDirs[i]) {
			return true
		}
	}

	return false
}
