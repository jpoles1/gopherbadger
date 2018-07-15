package coverbadge

import (
	"fmt"
	"gopherbadger/logging"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)

type Badge struct {
	CoveragePrefix string
	Style          string
	ImageExtension string
}

func (badge Badge) generateBadgeBadgeURL(coverageFloat float64) string {
	if badge.CoveragePrefix != "" {
		badge.CoveragePrefix += "%20"
	}
	urlTemplate := "https://img.shields.io/badge/%sCoverage-%2.f%%25-brightgreen%s?longCache=true&style=%s"
	return fmt.Sprintf(urlTemplate, badge.CoveragePrefix, coverageFloat, badge.ImageExtension, badge.Style)
}

func (badge Badge) DownloadBadge(filepath string, coverageFloat float64) {
	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		logging.Fatal("Creating file", err)
		return
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(badge.generateBadgeBadgeURL(coverageFloat))
	if err != nil {
		logging.Fatal("Fetching badge image", err)
		return
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		logging.Fatal("Writing file to disk", err)
		return
	}

	return
}

func (badge Badge) WriteBadgeToMd(filepath string, coverageFloat float64) {
	badge.ImageExtension = ".svg"
	badgeURL := badge.generateBadgeBadgeURL(coverageFloat)
	newImageTag := fmt.Sprintf("![gopherbadger-tag-do-not-edit](%s)", badgeURL)
	imageTagRegex := `\!?\[gopherbadger-tag-do-not-edit\](.*)`
	r, err := regexp.Compile(imageTagRegex)
	if err != nil {
		logging.Fatal("Compiling regex: ", err)
		return
	}
	filedata, err := ioutil.ReadFile(filepath)
	logging.Error("Could not open markdown file: ", err)

	var markdownData string
	if string(filedata) == "" {
		markdownData = newImageTag
	} else {
		markdownData = r.ReplaceAllString(string(filedata), newImageTag)
	}
	err = ioutil.WriteFile(filepath, []byte(markdownData), 0644)
	if err != nil {
		logging.Fatal("Error: could not write badge url to markdown file: ", err)
		return
	}
	logging.Success("Wrote badge image to markdown file: " + filepath)
}
