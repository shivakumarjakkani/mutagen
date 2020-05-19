package compose

import (
	"fmt"

	"github.com/spf13/cobra"
)

func startMain(_ *cobra.Command, arguments []string) {
	// Handle top-level help and version flags.
	handleTopLevelHelp()
	handleTopLevelVersion()

	// TODO: Implement.
	fmt.Println("start not yet implemented")
}

var startCommand = &cobra.Command{
	Use:          "start",
	Run:          startMain,
	SilenceUsage: true,
}

var startConfiguration struct {
	// help indicates the presence of the -h/--help flag.
	help bool
}

func init() {
	// Avoid Cobra's built-in help functionality that's triggered when the
	// -h/--help flag is present and instead just redirect control to the
	// nominal entry point. We'll use the -h/--help flag that we create below to
	// determine when help functionality needs to be displayed.
	startCommand.SetHelpFunc(startMain)

	// Grab a handle for the command line flags.
	flags := startCommand.Flags()

	// Wire up start command flags.
	flags.BoolVarP(&startConfiguration.help, "help", "h", false, "")
	// TODO: Wire up remaining flags.
}