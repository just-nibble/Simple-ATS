package main

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(2)
	var wg sync.WaitGroup
	wg.Add(1)

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
			var file_name string = "bin/files/" + f.Name() + ".txt"

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

			if points <= 2 {
				go SendRejectionEmail(FindEmail(text))
				wg.Wait()
			}
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
	var outp string = "bin/files/" + name + ".txt"
	var input string = "bin/files/" + name
	_, err := exec.Command("pdftotext", input, outp).Output()
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

func SendRejectionEmail(toEmailAddress string) error {
	var from string = os.Getenv("SMTP_FROM")
	var password string = os.Getenv("SMTP_PASSWORD")

	to := []string{toEmailAddress}

	var host string = os.Getenv("SMTP_HOST")
	var port string = os.Getenv("SMTP_PORT")
	var address string = host + ":" + port

	var subject string = "APPLICATION RESPONSE"

	var body string = `
		Hello applicant, I regret to inform you that
		you did not meet the required requirements for this role.

		Do not be discouraged, you are welcome to apply at a later
		time.

		Have a nice day.
	`
	message := []byte(subject + body)

	auth := smtp.PlainAuth("", from, password, host)

	err := smtp.SendMail(address, auth, from, to, message)
	if err != nil {
		return err
	}
	return nil
}

func FindEmail(str string) string {
	var re = regexp.MustCompile(`(?m)mail=[A-Za-z.@0-9]+\,`)
	var output string

	for i, match := range re.FindAllString(str, -1) {
		fmt.Println(match, "found at index", i)
		email := strings.Split(match, "=")[1]

		email = strings.ReplaceAll(email, ",", "")
		output = email
	}
	return output
}
