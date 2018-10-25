package main

import (
	"github.com/segmentio/aws-okta/cmd"
)

// These are set via linker flags
var (
	Version           = "forked"
	AnalyticsWriteKey = ""
)

func main() {
	// vars set by linker flags must be strings...
	cmd.Execute(Version, AnalyticsWriteKey)
}
