package cmd

import (
	"github.com/charmbracelet/charm/server"
	"github.com/charmbracelet/charm/server/db/sqlite"
	"github.com/spf13/cobra"

	_ "modernc.org/sqlite" // sqlite driver
)

func init() {
	ServeMigrationCmd.AddCommand(RetryMigrationCmd)
}

var RetryMigrationCmd = &cobra.Command{
	Use:     "retry",
	Aliases: []string{"r"},
	Hidden:  false,
	Short:   "Retry a failed migration.",
	Long:    paragraph("Retry a failed migration."),
	RunE:    RetryMigrationCmdRunE,
}

// ServeMigrationCmd migrate server db.
var ServeMigrationCmd = &cobra.Command{
	Use:     "migrate",
	Aliases: []string{"migration"},
	Hidden:  true,
	Short:   "Run the server migration tool.",
	Long:    paragraph("Run the server migration tool to migrate the database."),
	Args:    cobra.NoArgs,
	RunE:    ServeMigrationCmdRunE,
}

func RetryMigrationCmdRunE(cmd *cobra.Command, args []string) error {
	cfg := server.DefaultConfig()
	if len(args) == 0 {
		args = []string{"allow-last-failed"}
	}
	return sqlite.NewDB(cfg.DbDriver, server.GetDBDataSource(cfg), args).Migrate()
}

func ServeMigrationCmdRunE(cmd *cobra.Command, args []string) error {
	cfg := server.DefaultConfig()
	return sqlite.NewDB(cfg.DbDriver, server.GetDBDataSource(cfg), nil).Migrate()
}
