package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

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
