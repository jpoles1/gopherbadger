package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"github.com/jpoles1/gopherbadger/coverbadge"
	"github.com/jpoles1/gopherbadger/logging"

	"github.com/fatih/color"
)

const toolCoverCommand = "go tool cover -func=coverage.out"

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
		re := regexp.MustCompile(`total:\s*\(statements\)?\s*(\d+\.?\d*)\s*\%`)
		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {
			lineText := scanner.Text()
			match := re.FindStringSubmatch(lineText)
			if len(match) == 2 {
				color.Green(lineText)
				// fmt.Printf("Found coverage = %s\n", match[1])
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
		if err := cmd.Wait(); err != nil {
			log.Fatal(err)
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}(reader)

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	return coverageFloatChannel
}

func containsString(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}

type gopherBadgerConfig struct {
	badgeOutputFlag     bool
	badgeStyleFlag      string
	updateMdFilesFlag   string
	coveragePrefixFlag  string
	coverageCommandFlag string
	manualCoverageFlag  float64
	rootFolderFlag      string
	tagsFlag            string
	shortFlag           bool
}

var badgeStyles = []string{"plastic", "flat", "flat-square", "for-the-badge", "social"}

func main() {
	badgeOutputFlag := flag.Bool("png", true, "Boolean to decide if .png will be generated by the software")
	badgeStyleFlag := flag.String("style", "flat", "Badge style from list: ["+strings.Join(badgeStyles, ",")+"]")
	updateMdFilesFlag := flag.String("md", "", "A list of markdown filepaths for badge updates.")
	coveragePrefixFlag := flag.String("prefix", "Go", "A prefix to specify the coverage in your badge.")
	coverageCommandFlag := flag.String("covercmd", "", "gocover command to run; must print coverage report to stdout")
	manualCoverageFlag := flag.Float64("manualcov", -1.0, "A manually inputted coverage float.")
	rootFolderFlag := flag.String("root", ".", "A folder within your project from which to start recursively scanning and testing.")
	tagsFlag := flag.String("tags", "", "The build tests you'd like to include in your coverage")
	shortFlag := flag.Bool("short", false, "It will skip tests marked as testing.Short()")
	flag.Parse()
	config := gopherBadgerConfig{
		badgeOutputFlag:     *badgeOutputFlag,
		badgeStyleFlag:      *badgeStyleFlag,
		updateMdFilesFlag:   *updateMdFilesFlag,
		coveragePrefixFlag:  *coveragePrefixFlag,
		coverageCommandFlag: *coverageCommandFlag,
		manualCoverageFlag:  *manualCoverageFlag,
		rootFolderFlag:      *rootFolderFlag,
		tagsFlag:            *tagsFlag,
		shortFlag:           *shortFlag,
	}
	badger(config)
}

func badger(config gopherBadgerConfig) {
	if !containsString(badgeStyles, config.badgeStyleFlag) {
		logging.Fatal("Invalid style flag! Must be a member of list: ["+strings.Join(badgeStyles, ", ")+"]", errors.New("Invalid style flag"))
	}

	coverageBadge := coverbadge.Badge{
		CoveragePrefix: config.coveragePrefixFlag,
		Style:          config.badgeStyleFlag,
		ImageExtension: ".png",
	}

	var coverageFloat float64
	coverageCommand := ""

	if config.coverageCommandFlag != "" {
		coverageCommand = config.coverageCommandFlag
		if config.tagsFlag != "" || config.shortFlag {
			log.Println("Warning: When the covercmd flag is used the -tags and -short flags will be ignored.")
		}
	} else {
		flagsCommands := ""
		if config.tagsFlag != "" {
			flagsCommands = flagsCommands + " -tags \"" + config.tagsFlag + "\""
		}
		if config.shortFlag {
			flagsCommands = flagsCommands + " -short"
		}
		coverageCommand = fmt.Sprintf("go test %s/... -coverprofile=coverage.out %s && %s", config.rootFolderFlag, flagsCommands, toolCoverCommand)
	}

	if config.manualCoverageFlag == -1 {
		coverageFloat = <-getCommandOutput(coverageCommand)
	} else {
		coverageFloat = config.manualCoverageFlag
	}
	if config.badgeOutputFlag {
		coverageBadge.DownloadBadge("coverage_badge.png", coverageFloat)
	}
	if config.updateMdFilesFlag != "" {
		for _, filepath := range strings.Split(config.updateMdFilesFlag, ",") {
			coverageBadge.WriteBadgeToMd(filepath, coverageFloat)
		}
	}
}
