// Package cmd implements the Cobra commands for the charm CLI.
package cmd

import "github.com/charmbracelet/charm/ui/common"

var (
	// Version is the version of the charm CLI.
	Version = ""
	// CommitSHA is the commit SHA of the charm CLI.
	CommitSHA = ""
	// Styles is the default styles for the charm CLI.
	Styles = common.DefaultStyles()
)

func Main() {
	RootCmdExecute()
}
