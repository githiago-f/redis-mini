package main

import (
	"fmt"
	"net"
	"os"

	"github.com/githiago-f/redis-mini/core"
	"github.com/spf13/cobra"
)

var command = &cobra.Command{
	Use:  "redis-mini",
	Long: fmt.Sprintf("Redis copy from %v v%v", core.AUTHOR, core.VERSION),
	Run:  run,
}

func main() {
	command.Flags().Int16P("port", "p", 6379, "Port to listen (Default: 6379)")
	command.Flags().StringP("address", "a", "localhost", "Host on wich this server will run")

	if err := command.Execute(); err != nil {
		core.Logger.Error(err)
		os.Exit(1)
	}
}

func run(cmd *cobra.Command, args []string) {
	port, _ := cmd.Flags().GetInt16("port")
	address, _ := cmd.Flags().GetString("address")
	if address == "" {
		core.Logger.Printf("Cannot listen in address %v", address)
		os.Exit(1)
	}

	receivedPort := fmt.Sprintf("%d", port)
	address += ":" + receivedPort

	server, err := net.Listen("tcp", address)
	core.Logger.Printf("Running server on %v", address)

	if err != nil {
		core.Logger.Error(err)
		os.Exit(1)
	}

	defer server.Close()

	for {
		con, err := server.Accept()
		if err != nil {
			core.Logger.Error(err)
			os.Exit(1)
		}

		go HandleConnection(con)
	}
}
