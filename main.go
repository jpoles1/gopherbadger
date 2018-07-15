package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strconv"

	"github.com/fatih/color"
)

func getCommandOutput(commandString string) chan float64 {
	cmd := exec.Command("bash", "-c", commandString)
	cmd.Stderr = os.Stderr
	stdout, err := cmd.StdoutPipe()
	if nil != err {
		log.Fatalf("Error obtaining stdout: %s", err.Error())
	}
	reader := bufio.NewReader(stdout)
	coverageFloatChannel := make(chan float64)
	go func(reader io.Reader) {
		re := regexp.MustCompile("total:\\s*\\(statements\\)?\\s*(\\d+\\.?\\d*)\\s*\\%")
		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {
			lineText := scanner.Text()
			match := re.FindStringSubmatch(lineText)
			if len(match) == 2 {
				color.Green(lineText)
				//fmt.Printf("Found coverage = %s%\n", match[1])
				coverageValue, err := strconv.ParseFloat(match[1], 32)
				errCheck("Parsing coverage to float", err)
				if err == nil {
					coverageFloatChannel <- coverageValue
				}
				break
			} else {
				fmt.Println(lineText)
			}
		}
	}(reader)
	if err := cmd.Start(); nil != err {
		log.Fatalf("Error starting program: %s, %s", cmd.Path, err.Error())
	}
	cmd.Wait()
	return coverageFloatChannel
}
func writeBadgeToMd(coverageFloat float64, filepath string) {
	badgeURL := "https://img.shields.io/badge/Coverage-%2.f%%-brightgreen.svg?longCache=true&style=flat"
	newImageTag := fmt.Sprintf("![gopherbadger-tag-do-not-edit](%s)", fmt.Sprintf(badgeURL, coverageFloat))
	imageTagRegex := `\!?\[gopherbadger-tag-do-not-edit\](.*)`
	r, err := regexp.Compile(imageTagRegex)
	if err != nil {
		log.Fatal("Compiling regex: ", err)
	}
	filedata, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal("Could not open markdown file: ", err)
	}
	markdownData := r.ReplaceAllString(string(filedata), newImageTag)
	err = ioutil.WriteFile(filepath, []byte(markdownData), 0644)
	if err != nil {
		log.Fatal("Error: could not write shield url to markdown file: ", err)
	}
}
func main() {
	var coverageFloat float64
	coverageFloat = <-getCommandOutput("go test ./... -coverprofile=coverage.out && go tool cover -func=coverage.out")
	drawBadge(coverageFloat, "coverage_badge.png")
	writeBadgeToMd(coverageFloat, "./README.md")
}
