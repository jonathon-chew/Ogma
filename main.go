package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	cmd "github.com/jonathon-chew/Ogma/cmd"

	Aphrodite "github.com/jonathon-chew/Aphrodite"
)

type LanguageStats struct {
	Name  string
	Files int
	Lines int
}

var extToLang = map[string]string{
	"py":    "Python",
	"js":    "JavaScript",
	"java":  "Java",
	"go":    "Golang",
	"rs":    "Rust",
	"cpp":   "C++",
	"cc":    "C++",
	"cxx":   "C++",
	"c":     "C",
	"cs":    "C#",
	"php":   "PHP",
	"rb":    "Ruby",
	"ts":    "TypeScript",
	"swift": "Swift",
	"kt":    "Kotlin",
	"scala": "Scala",
	"r":     "R",
	"dart":  "Dart",
	"hs":    "Haskell",
	"m":     "Objective-C",
	"qml":   "QML",
	"jl":    "Julia",
	"sh":    "Shell",
	"pl":    "Perl",
	"lua":   "Lua",
	"sql":   "SQL",
	"mod":   "Golang",
	"sum":   "Golang",
	"html":  "HTML",
	"ccs":   "CCS",
	"ps1":   "Powershell",
	"psm1":  "Powershell",
	"psd1":  "Powershell",
	"md":    "Markdown",
}

/*
Convert a int into a string, but make it human readbale by working backwards and applying commas in the right place to split up the number
*/
func HumanReadableInt(initalInt int) string {
	convertedNumber := strconv.Itoa(initalInt)
	var humanReadbleNumber string
	count := 0

	if len(convertedNumber) <= 3 {
		return convertedNumber
	}

	for i := len(convertedNumber) - 1; i >= 0; i-- {
		humanReadbleNumber = string(convertedNumber[i]) + humanReadbleNumber
		count++
		if count%3 == 0 && i != 0 {
			humanReadbleNumber = "," + humanReadbleNumber
		}
	}

	return humanReadbleNumber
}

func main() {

	var cmdFlags cmd.Flags
	if len(os.Args[1:]) >= 1 {
		cmdFlags = cmd.Cmd(os.Args[1:])
	}

	root := "./"
	var LangStats []LanguageStats

	filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return err // stop on error
		}

		for _, ignoreFolder := range cmdFlags.IgnoreFolders {
			if strings.Contains(path, ignoreFolder+"/") || strings.Contains(path, ignoreFolder+"\\") {
				return nil
			}
		}

		for _, ignoreFile := range cmdFlags.IgnoreFiles {
			if strings.Contains(d.Name(), ignoreFile) {
				return nil
			}
		}

		// Pass back a pointer to a file and an error if it fails
		PointerToFile, OpenFileError := os.Open(path)
		if OpenFileError != nil {
			fmt.Print("error opening the file " + path + "\n")
			return nil
		}

		defer PointerToFile.Close()

		var lines int
		scanner := bufio.NewScanner(PointerToFile)
		for scanner.Scan() {
			if scanner.Text() != "\n" {
				lines++
			}
		}

		// Get the file extension
		fileExtension := strings.ReplaceAll(filepath.Ext(d.Name()), ".", "")

		// Check if the extension is a known one
		_, mapContainsKey := extToLang[fileExtension]

		// If the file has an extension AND is a known one
		if strings.Contains(filepath.Ext(d.Name()), ".") && mapContainsKey {

			var found bool = false
			// Loop through Language stats and if it exists add to it, else add it on
			for i := range LangStats {
				if LangStats[i].Name == extToLang[fileExtension] {
					LangStats[i].Files += 1
					LangStats[i].Lines += lines
					found = true
					break
				}
			}

			if !found {
				LangStats = append(LangStats, LanguageStats{
					Name:  extToLang[fileExtension],
					Files: 1,
					Lines: lines,
				})
			}
		} else { // Add the extension as the language
			if strings.Contains(filepath.Ext(d.Name()), ".") {

				var found bool = false
				// Loop through Language stats and if it exists add to it, else add it on
				for i := range LangStats {
					if LangStats[i].Name == filepath.Ext(d.Name())[1:] {
						LangStats[i].Files += 1
						LangStats[i].Lines += lines
						found = true
						break
					}
				}

				if !found {
					LangStats = append(LangStats, LanguageStats{
						Name:  filepath.Ext(d.Name())[1:],
						Files: 1,
						Lines: lines,
					})
				}
			}
		}
		return nil
	})

	// Logic for parsing out the contents well
	// this maybe extracted later for a table implimentation
	var biggestLangLength int = len("No. files:")
	for _, longestLang := range LangStats {
		if len(longestLang.Name) > biggestLangLength {
			biggestLangLength = len(longestLang.Name)
		}
	}

	var biggestNumberOfFilesLength int = len("No. Lines:")
	for _, longestLang := range LangStats {
		if len(longestLang.Name) > biggestNumberOfFilesLength {
			biggestNumberOfFilesLength = len(longestLang.Name)
		}
	}

	biggestNumberOfFilesLength = len(HumanReadableInt(biggestNumberOfFilesLength))

	var totalLines, totalFiles int

	header := fmt.Sprintf("Name: %%-%ds No. files: %%-%ds No. Lines: %%s\n", biggestLangLength, biggestNumberOfFilesLength)
	Aphrodite.PrintBold("Cyan", fmt.Sprintf(header, " ", " ", " "))

	for _, printresult := range LangStats {
		sentence := fmt.Sprintf("%%-%ds %%-%ds %%s\n", biggestLangLength+len("No. files:"), biggestNumberOfFilesLength+len("No. Lines:"))
		Aphrodite.PrintColour("Green", fmt.Sprintf(sentence, printresult.Name, HumanReadableInt(printresult.Files), HumanReadableInt(printresult.Lines)))
		totalFiles += printresult.Files
		totalLines += printresult.Lines
	}

	Aphrodite.PrintBoldHighIntensity("Yellow", "\n\nTotal Lines: "+HumanReadableInt(totalLines)+" Total Files: "+HumanReadableInt(totalFiles)+"\n")
}
