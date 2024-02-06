// Package cmd implements the Cobra commands for the charm CLI.
package cmd

import (
	"fmt"
	"os"
	"runtime/debug"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"

	"github.com/charmbracelet/charm/client"
	"github.com/charmbracelet/charm/ui"
	"github.com/charmbracelet/charm/ui/common"
)

var (
	RootCmd = &cobra.Command{
		Use:                   "charm",
		Short:                 "Do Charm stuff",
		Long:                  Styles.Paragraph.Render(fmt.Sprintf("Do %s stuff. Run without arguments for a TUI or use the sub-commands like a pro.", Styles.Keyword.Render("Charm"))),
		DisableFlagsInUseLine: true,
		RunE:                  RootCmdRunE,
	}
)

func init() {
	if len(CommitSHA) >= 7 {
		vt := RootCmd.VersionTemplate()
		RootCmd.SetVersionTemplate(vt[:len(vt)-1] + " (" + CommitSHA[0:7] + ")\n")
	}
	if Version == "" {
		if info, ok := debug.ReadBuildInfo(); ok && info.Main.Sum != "" {
			Version = info.Main.Version
		} else {
			Version = "unknown (built from source)"
		}
	}
	RootCmd.Version = Version
	RootCmd.CompletionOptions.HiddenDefaultCmd = true

	RootCmd.AddCommand(
		BioCmd,
		IDCmd,
		JWTCmd,
		KeysCmd,
		LinkCmd("charm"),
		NameCmd,
		BackupKeysCmd,
		ImportKeysCmd,
		KeySyncCmd,
		CompletionCmd,
		ServeCmd,
		PostNewsCmd,
		KVCmd,
		FSCmd,
		CryptCmd,
		MigrateAccountCmd,
		WhereCmd,
		ManCmd,
	)
}

func RootCmdExecute() {
	if err := RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func RootCmdRunE(cmd *cobra.Command, args []string) error {
	if common.IsTTY() {
		cfg, err := client.ConfigFromEnv()
		if err != nil {
			log.Fatal(err)
		}

		// Log to file, if set
		if cfg.Logfile != "" {
			f, err := os.OpenFile(cfg.Logfile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0o644)
			if err != nil {
				return err
			}
			if cfg.Debug {
				log.SetLevel(log.DebugLevel)
			}
			log.SetOutput(f)
			log.SetPrefix("charm")

			defer f.Close() // nolint: errcheck
		}

		p := ui.NewProgram(cfg)
		if _, err := p.Run(); err != nil {
			return err
		}
	}

	return cmd.Help()
}
