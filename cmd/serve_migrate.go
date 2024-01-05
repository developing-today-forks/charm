package cmd

import (
	"github.com/charmbracelet/charm/server"
	"github.com/charmbracelet/charm/server/db/sqlite"
	"github.com/spf13/cobra"

	_ "modernc.org/sqlite" // sqlite driver
)

// ServeMigrationCmd migrate server db.
var ServeMigrationCmd = &cobra.Command{
	Use:     "migrate",
	Aliases: []string{"migration"},
	Hidden:  true,
	Short:   "Run the server migration tool.",
	Long:    paragraph("Run the server migration tool to migrate the database."),
	Args:    cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg := server.DefaultConfig()
		return sqlite.NewDB(cfg.DbDriver, server.GetDBDataSource(cfg)).Migrate()
	},
}
