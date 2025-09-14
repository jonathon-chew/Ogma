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

	cmd "github.com/jonathon-chew/Lines_Of_Code/cmd"

	Aphrodite "github.com/jonathon-chew/Aphrodite"
)

type LanguageStats struct {
	Name  string
	Files int
	Lines int
}

type Results struct {
	LanguageStats []LanguageStats
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
		cmd.Cmd(os.Args)
	}

	root := "./"
	var LangStats Results
	filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return err // stop on error
		}

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
			for i := range LangStats.LanguageStats {
				if LangStats.LanguageStats[i].Name == extToLang[fileExtension] {
					LangStats.LanguageStats[i].Files += 1
					LangStats.LanguageStats[i].Lines += lines
					found = true
					break
				}
			}

			if !found {
				LangStats.LanguageStats = append(LangStats.LanguageStats, LanguageStats{
					Name:  extToLang[fileExtension],
					Files: 1,
					Lines: lines,
				})
			}

		}

		return nil
	})

	for _, printresult := range LangStats.LanguageStats {
		Aphrodite.PrintColour("Green", fmt.Sprintf("\nName: %s,\nNumber of files: %s,\nNumber of Lines: %s\n", printresult.Name, HumanReadableInt(printresult.Files), HumanReadableInt(printresult.Lines)))
	}
}
