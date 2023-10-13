// Package serve is an entry point and help tools for serve commands
package serve

import (
	"github.com/resourcemod/registry/internal/handlers/serve"
	"github.com/spf13/cobra"
)

// NewServeCommand is serve command entry point
func NewServeCommand() *cobra.Command {
	var (
		command *cobra.Command
		host    string
		port    string
		ui      bool
		static  string
	)
	command = &cobra.Command{
		Use:     "serve",
		Short:   "Serve ResourceMod rmod-reg api and UI app",
		Long:    "Launches the ResourceMod rmod-reg in server mode, as well as running the UI interface for convenient use.",
		Example: getExample(),
		RunE: func(cmd *cobra.Command, args []string) error {
			err := serve.HandleServeCommand(command.Context(), host, port, static, ui)
			return err
		},
	}
	command.Flags().StringVarP(&port, "port", "p", "8888", "Server port.")
	command.Flags().StringVarP(&host, "addr", "a", "0.0.0.0", "Server host.")
	command.Flags().BoolVarP(&ui, "ui", "u", true, "Run UI app? Default=true.")
	command.Flags().StringVarP(&static, "static", "s", "./web/app/dist", "UI app static files folder.")

	return command
}
