package cmd

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/charm/ui/link"
	"github.com/charmbracelet/charm/ui/link/detached"
	"github.com/charmbracelet/charm/ui/linkgen"
	detachedGen "github.com/charmbracelet/charm/ui/linkgen/detached"
	"github.com/muesli/reflow/indent"
	"github.com/spf13/cobra"
)

// LinkCmd is the cobra.Command to manage user account linking. Pass the name
// of the parent command.
func LinkCmd(parentName string) *cobra.Command {
	var detach bool
	var outFilePath string
	var keys string
	cmd := &cobra.Command{
		Use:     "link [code]",
		Short:   "Link multiple machines to your Charm account",
		Long:    paragraph("It’s easy to " + keyword("link") + " multiple machines or keys to your Charm account. Just run " + code(parentName+" link") + " on a machine connected to the account to want to link to start the process."),
		Example: indent.String(fmt.Sprintf("%s link\n%s link XXXXXX", parentName, parentName), 2),
		Args:    cobra.RangeArgs(0, 1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if !detach && outFilePath != "" {
				return fmt.Errorf("cannot use out file with interactive mode")
			}
			// Log to file if specified in the environment
			cfg := getCharmConfig()
			if cfg.Logfile != "" {
				f, err := tea.LogToFile(cfg.Logfile, "charm")
				if err != nil {
					return err
				}
				defer f.Close() //nolint:errcheck
			}
			if detach {
				if outFilePath != "" {
					return detachedGen.LinkGen(cfg, parentName, outFilePath, keys)
				} else if len(args) == 0 {
					return fmt.Errorf("code or out file required for detach mode")
				} else {
					return detached.Link(cfg, parentName, args[0])
				}
			}
			var p *tea.Program
			switch len(args) {
			case 0:
				// Initialize a linking session
				p = linkgen.NewProgram(cfg, parentName)
			default:
				// Join in on a linking session
				p = link.NewProgram(cfg, args[0])
			}
			if _, err := p.Run(); err != nil {
				return err
			}
			return nil
		},
	}
	cmd.Flags().BoolVarP(&detach, "detach", "d", false, "Run in headless mode")
	cmd.Flags().BoolVar(&detach, "auto-approve", false, "Run in headless mode")
	cmd.Flags().BoolVar(&detach, "headless", false, "Run in headless mode")
	cmd.Flags().StringVarP(&outFilePath, "file", "f", "", "Write the generated link to a file")
	cmd.Flags().StringVar(&outFilePath, "out-file", "", "Write the generated link to a file")
	cmd.Flags().StringVarP(&outFilePath, "out", "o", "", "Write the generated link to a file")
	cmd.Flags().StringVarP(&keys, "keys", "k", "", "Link a key to your account")
	return cmd
}
