package cmd

import (
	"github.com/juanfer2/whorship_band/src/shared/infrastructure/persistence/postgres"
	"github.com/spf13/cobra"
	// "github.com/juanfer2/ayenda_service/src/config"
	// "github.com/juanfer2/ayenda_service/src/models"
)

var createDatBaseCmd = &cobra.Command{
	Use:   "db:create",
	Short: "Create database",
	Long:  `Create database`,
	Run:   createDB,
}

func createDB(cmd *cobra.Command, args []string) {
	configClient := postgres.CreateConfigClient()
	postgresClient := postgres.PostgresClient{ConfigClient: configClient}
	postgresClient.CreateDataBase()
}
