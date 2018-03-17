package main

import "testing"

func TestGetCommandOutput(t *testing.T) {
	getCommandOutput("echo test")
	//getCommandOutput("echo 'testline\ncoverage: 35%'")
}

func TestDrawBadge(t *testing.T) {
	drawBadge(22.7, "test_badge.png")
	drawBadge(66, "test_badge.png")
	drawBadge(88, "test_badge.png")
	if drawBadge(88, "bad_folder/test_badge.png") == nil {
		t.Error("Should respond with error when saving to invalid folder")
	}
	if drawBadge(-34, "test_badge.png") == nil {
		t.Error("Should respond with error when coverage is less than 0")
	}
}
