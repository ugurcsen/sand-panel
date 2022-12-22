package cmd

import (
	"github.com/spf13/cobra"
	"github.com/ugurcsen/sand-panel/internal/api"
)

var (
	restPort uint16
	grpcPort uint16
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve [OPTIONS] [COMMANDS]",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("serve called")

		// Rest API
		restChan := make(chan error)
		go func() {
			err := api.RunRest(restPort)
			restChan <- err
		}()
		// GRPC API
		grpcChan := make(chan error)
		go func() {
			err := api.RunGrpc(grpcPort)
			grpcChan <- err
		}()

		// Wait for any error
		err := make(chan error)
		go func() {
			err <- <-restChan
		}()
		go func() {
			err <- <-restChan
		}()
		cmd.PrintErr(<-err)
	},
}

func init() {
	serveCmd.PersistentFlags().Uint16VarP(&restPort, "rest_port", "r", 8000, "Rest Port number")
	serveCmd.PersistentFlags().Uint16VarP(&grpcPort, "grpc_port", "g", 9000, "gRPC Port number")
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
