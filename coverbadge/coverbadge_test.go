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
