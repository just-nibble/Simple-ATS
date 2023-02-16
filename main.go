package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	var points int = 0
	// Reading all files of directory
	files, err := os.ReadDir("./bin/files")
	if err != nil {
		log.Fatal(err)
	}
	// Iterating through all the files
	for _, f := range files {
		// checking for file extensions
		ext := filepath.Ext(f.Name())
		// if extension relevant
		if ext == ".pdf" {
			//Convert to .txt
			ExtractText(f.Name())
			var file_name string = f.Name() + ".txt"

			var text string = ReadText(file_name)
			// checking what skills the candidate has
			if frontEndSkills(text) {
				points += 3
			}
			if backEndSkills(text) {
				points += 1
			}
			if otherSkills(text) {
				points += 1
			}
			fmt.Println("-------------------------------------")
			fmt.Printf("File: %s \n", f.Name())
			fmt.Printf("Candidate is Suitable, with score: %d/5 \n", points)

		}
	}
}

// func for checking if candidate has front end skills
func frontEndSkills(text string) bool {
	isFrontEnd := false

	// Some front end skills that we are checking
	skills := []string{"JavaScript", "JS", "jQuery", "VUE", "React", "CSS", "SCSS"}

	// Iterate through skills
	for _, skill := range skills {
		if strings.Contains(text, skill) {
			isFrontEnd = true
		}
	}
	return isFrontEnd
}

// func for checking if candidate has back end skills
func backEndSkills(text string) bool {
	isBackEnd := false

	// Some front end skills that we are checking
	skills := []string{"Python", "C++", "C#", "Node", "ASP.NET", "Django", "PHP", "Laravel"}

	// Iterate through skills
	for _, skill := range skills {
		if strings.Contains(text, skill) {
			isBackEnd = true
		}
	}
	return isBackEnd
}

// func for checking if candidate has other end skills
func otherSkills(text string) bool {
	isSkilled := false

	// Some front end skills that we are checking
	skills := []string{"TDD", "Agile", "Git", "OOP", "SQL", "JIRA"}

	// Iterate through skills
	for _, skill := range skills {
		if strings.Contains(text, skill) {
			isSkilled = true
		}
	}
	return isSkilled
}

func ExtractText(name string) error {
	outp := name + ".txt"
	_, err := exec.Command("pdftotext", name, outp).Output()
	if err != nil {
		panic(err)
	}
	return nil
}

func ReadText(name string) string {
	file, err := os.ReadFile(name)
	if err != nil {
		fmt.Println("File reading error", err)
		panic(err)
	}
	var output string = string(file)
	err = os.Remove(name)
	if err != nil {
		panic(err)
	}
	return output
}
