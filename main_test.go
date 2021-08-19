package main

import (
	"testing"
)

func TestGetCommandOutput(t *testing.T) {
	getCommandOutput("echo test", false)
	getCommandOutput("echo 'testline\ncoverage: 35%'", false)
}

func TestBadger(t *testing.T) {
	defaultConfig := gopherBadgerConfig{
		badgeOutputFlag:     true,
		badgeStyleFlag:      "flat",
		updateMdFilesFlag:   "",
		coveragePrefixFlag:  "Go",
		coverageCommandFlag: "echo 'testline\ncoverage: 35%'",
		manualCoverageFlag:  90,
		rootFolderFlag:      ".",
		tagsFlag:            "",
		shortFlag:           false,
	}
	badger(defaultConfig)
}

func TestDrawBadge(t *testing.T) {
	if err := drawBadge(22.7, "test_badge.png"); err != nil {
		t.Errorf("error drawing 22.7 coverage: %s", err.Error())
	}
	if err := drawBadge(88, "test_badge.png"); err != nil {
		t.Errorf("error drawing 88% coverage: %s", err.Error())
	}
	if err := drawBadge(66, "test_badge.png"); err != nil {
		t.Errorf("error drawing 66% coverage: %s", err.Error())
	}
	if drawBadge(66, "bad_folder/test_badge.png") == nil {
		t.Error("Should respond with error when saving to invalid folder")
	}
	if drawBadge(-34, "test_badge.png") == nil {
		t.Error("Should respond with error when coverage is less than 0")
	}
}
