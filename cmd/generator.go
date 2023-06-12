package cmd

import (
	"github.com/juanfer2/whorship_band/cmd/migrations"
	"github.com/spf13/cobra"
	// "github.com/juanfer2/ayenda_service/src/config"
	// "github.com/juanfer2/ayenda_service/src/models"
)

var generatorCmd = &cobra.Command{
	Use:     "generate",
	Aliases: []string{"g"},
	Short:   "Generate files",
}

func init() {
	migrateCmd.AddCommand(migrations.CreateMigrationCmd, migrations.RunMigrationCmd,
		migrations.DownMigrationCmd)
}
