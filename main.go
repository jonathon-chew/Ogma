package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"log"
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
}

/* Convert a int into a string, but make it human readbale by working backwards and applying commas in the right place to split up the number
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

	if len(os.Args[1:]) >= 1 {
		cmd.Cmd(os.Args[1:])
	}

	root := "./"
	var LangStats []LanguageStats
	filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return err // stop on error
		}

		// Pass back a pointer to a file and an error if it fails
		PointerToFile, OpenFileError := os.Open(path)
		if OpenFileError != nil {
			log.Printf("error opening the file %s", path)
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

		fileExtension := strings.ReplaceAll(filepath.Ext(d.Name()), ".", "")

		_, isMapContainsKey := extToLang[fileExtension]

		if strings.Contains(filepath.Ext(d.Name()), ".") && isMapContainsKey {

			var found bool = false
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
		}
		return nil
	})

	var biggestLangLength int
	for _, longestLang := range LangStats {
		if len(longestLang.Name) > biggestLangLength {
			biggestLangLength = len(longestLang.Name)
		}
	}

	var biggestNumberOfFilesLength int
	for _, longestLang := range LangStats {
		if len(longestLang.Name) > biggestNumberOfFilesLength {
			biggestNumberOfFilesLength = len(longestLang.Name)
		}
	}

	// cdbiggestLangLength = biggestLangLength + 4
	biggestNumberOfFilesLength = len(HumanReadableInt(biggestNumberOfFilesLength))

	for _, printresult := range LangStats {
		sentence := fmt.Sprintf("\nName: %%-%ds    No. files: %%-%ds   No. Lines: %%s\n", biggestLangLength, biggestNumberOfFilesLength)
		Aphrodite.PrintColour("Green", fmt.Sprintf(sentence, printresult.Name, HumanReadableInt(printresult.Files), HumanReadableInt(printresult.Lines)))
	}
}
