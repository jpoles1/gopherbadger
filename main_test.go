package main

import (
	"testing"
)

func TestGetCommandOutput(t *testing.T) {
	getCommandOutput("echo test")
	getCommandOutput("echo 'testline\ncoverage: 35%'")
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
	_ = drawBadge(22.7, "test_badge.png")
	_ = drawBadge(88, "test_badge.png")
	_ = drawBadge(66, "test_badge.png")
	if drawBadge(66, "bad_folder/test_badge.png") == nil {
		t.Error("Should respond with error when saving to invalid folder")
	}
	if drawBadge(-34, "test_badge.png") == nil {
		t.Error("Should respond with error when coverage is less than 0")
	}
}
