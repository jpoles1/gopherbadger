package main

import (
	"bufio"
	"fmt"
	"io"
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
		re := regexp.MustCompile("cover(?:ed|(?:age))*:? *(\\d+\\.?\\d*) *%")
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
func main() {
	var coverageFloat float64
	coverageFloat = <-getCommandOutput("make cover")
	drawBadge(coverageFloat, "badge.png")
}
