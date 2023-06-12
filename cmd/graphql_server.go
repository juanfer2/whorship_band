package cmd

import (
	"github.com/spf13/cobra"

	"github.com/juanfer2/whorship_band/src/shared/infrastructure/servers"
)

var serverGraphqlCmd = &cobra.Command{
	Use:   "server_graphql",
	Short: "server about this service",
	Long:  `server about this service`,
	Run:   runServerGraphql,
}

func runServerGraphql(cmd *cobra.Command, args []string) {
	servers.GraphqlServer()
}
