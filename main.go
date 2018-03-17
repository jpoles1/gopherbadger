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

func main() {
	cmd := exec.Command("bash", "-c", "make cover")
	cmd.Stderr = os.Stderr
	stdout, err := cmd.StdoutPipe()
	if nil != err {
		log.Fatalf("Error obtaining stdout: %s", err.Error())
	}
	reader := bufio.NewReader(stdout)
	go func(reader io.Reader) {
		re := regexp.MustCompile("cover(?:ed|(?:age))*:? *(\\d+\\.?\\d*) *%")
		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {
			lineText := scanner.Text()
			match := re.FindStringSubmatch(lineText)
			if len(match) == 2 {
				color.Green(lineText)
				fmt.Printf("Found coverage = %s%\n", match[1])
				coverageFloat, err := strconv.ParseFloat(match[1], 32)
				errCheck("Parsing coverage to float", err)
				drawBadge(coverageFloat)
			} else {
				fmt.Println(lineText)
			}
		}
	}(reader)
	if err := cmd.Start(); nil != err {
		log.Fatalf("Error starting program: %s, %s", cmd.Path, err.Error())
	}
	cmd.Wait()

}
