package coverbadge

import (
	"testing"
)
var coverageBadge = Badge{
	CoveragePrefix: "Go",
	Style:          "flat",
	ImageExtension: ".png",
}

func TestDownloadBadge(t *testing.T) {
	coverageBadge.DownloadBadge("test_badge.png", 66)
}
func TestWriteBadgeToMd(t *testing.T) {
	coverageBadge.WriteBadgeToMd("coverage_test.md", 22)
}
func TestGenerateBadgeBadgeURL_1(t *testing.T) {
	testUrl(t, 1.1, "https://img.shields.io/badge/Go%20Coverage-1%25-brightgreen.png?longCache=true&style=flat")
}
func TestGenerateBadgeBadgeURL_10(t *testing.T) {
	testUrl(t, 10.1, "https://img.shields.io/badge/Go%20Coverage-10%25-brightgreen.png?longCache=true&style=flat")
}
func TestGenerateBadgeBadgeURL_100(t *testing.T) {
	testUrl(t, 100.1, "https://img.shields.io/badge/Go%20Coverage-100%25-brightgreen.png?longCache=true&style=flat")
}
func testUrl(t *testing.T, coverageFloat float64, expected string) {
	url := coverageBadge.generateBadgeBadgeURL(coverageFloat)
	if url != expected {
		t.Fatal("url should be", expected, "but is", url)
	}
}
