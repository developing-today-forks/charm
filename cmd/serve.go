package cmd

import (
	"context"
	"path/filepath"
	"time"

	"github.com/charmbracelet/log"

	"github.com/charmbracelet/charm/server"
	"github.com/charmbracelet/keygen"
	"github.com/spf13/cobra"
)

var (
	serverHTTPPort   int
	serverSSHPort    int
	serverStatsPort  int
	serverHealthPort int
	serverDataDir    string

	// ServeCmd is the cobra.Command to self-host the Charm Cloud.
	ServeCmd = &cobra.Command{
		Use:     "serve",
		Aliases: []string{"server"},
		Hidden:  false,
		Short:   "Start a self-hosted Charm Cloud server.",
		Long:    paragraph("Start the SSH and HTTP servers needed to power a SQLite-backed Charm Cloud."),
		Args:    cobra.NoArgs,
		RunE:    ServeCmdRunE,
	}
)

func init() {
	ServeCmd.AddCommand(
		ServeMigrationCmd,
	)
	ServeCmd.Flags().IntVar(&serverHTTPPort, "http-port", 0, "HTTP port to listen on")
	ServeCmd.Flags().IntVar(&serverSSHPort, "ssh-port", 0, "SSH port to listen on")
	ServeCmd.Flags().IntVar(&serverStatsPort, "stats-port", 0, "Stats port to listen on")
	ServeCmd.Flags().IntVar(&serverHealthPort, "health-port", 0, "Health port to listen on")
	ServeCmd.Flags().StringVar(&serverDataDir, "data-dir", "", "Directory to store SQLite db, SSH keys and file data")
}

func ServeCmdRunE(cmd *cobra.Command, args []string) error {
	return ServeCmdRunEWithContext(context.Background(), cmd, args)
}

func ServeCmdRunEWithContext(ctx context.Context, cmd *cobra.Command, args []string) error {
	cfg := server.DefaultConfig()
	if serverHTTPPort != 0 {
		cfg.HTTPPort = serverHTTPPort
	}
	if serverSSHPort != 0 {
		cfg.SSHPort = serverSSHPort
	}
	if serverStatsPort != 0 {
		cfg.StatsPort = serverStatsPort
	}
	if serverHealthPort != 0 {
		cfg.HealthPort = serverHealthPort
	}
	if serverDataDir != "" {
		cfg.DataDir = serverDataDir
	}
	sp := filepath.Join(cfg.DataDir, ".ssh")
	kp, err := keygen.New(filepath.Join(sp, "charm_server_ed25519"), keygen.WithKeyType(keygen.Ed25519), keygen.WithWrite())
	if err != nil {
		return err
	}
	cfg = cfg.WithKeys(kp.RawAuthorizedKey(), kp.RawPrivateKey())
	s, err := server.NewServer(cfg)
	if err != nil {
		return err
	}

	done := make(chan error, 1)
	go func() {
		done <- s.Start()
	}()

	select {
	case err := <-done:
		if err != nil {
			log.Fatal("error starting server", "err", err)
		}
	case <-ctx.Done():
		log.Info("Context cancelled, shutting down server")
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	return s.Shutdown(shutdownCtx)
}
