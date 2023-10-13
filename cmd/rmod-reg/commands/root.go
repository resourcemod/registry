// Package commands is an entry point for every single cobra command available
package commands

import (
	"github.com/resourcemod/registry/cmd/rmod-reg/commands/serve"
	"github.com/spf13/cobra"
)

var (
	root *cobra.Command
)

// NewRootCommand is root command entry point
func NewRootCommand() *cobra.Command {
	root = &cobra.Command{
		Use:   "rmod-rmod-reg [command] [flags]",
		Short: "rmod-rmod-reg - An opensource plugins rmod-reg package.",
		Long:  "rmod-rmod-reg - Share your plugins using the Resource Mod plugin manager. Store your code safely in your repository using rmod rmod-reg.",
	}

	root.AddCommand(serve.NewServeCommand())

	return root
}
