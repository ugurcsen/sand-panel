package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/ugurcsen/sand-panel/internal/api/rest"
	"github.com/ugurcsen/sand-panel/internal/core/services/user_service"
	"github.com/ugurcsen/sand-panel/internal/repositories/postgresql"
	"net"
)

var (
	port int
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
		rps := postgresql.NewPostgresqlRepository()
		srv := user_service.NewUserService(rps)
		_ = srv
		tcp, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
		if err != nil {
			cmd.PrintErr(err)
			return
		}
		restApi, err := rest.NewRest(rest.Options{
			Listener: tcp,
		})
		if err != nil {
			panic(err)
		}
		restApi.Listen()
	},
}

func init() {
	serveCmd.PersistentFlags().IntVarP(&port, "port", "p", 8080, "Port number")
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
