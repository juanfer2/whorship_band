package cmd

import (
	"log"
	"os/exec"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "build project",
	Long:  `build project`,
	Run:   buildProject,
}

func buildProject(cmd *cobra.Command, args []string) {
	color.Blue(">>>Building Project start 🚀 <<<")

	_, err := exec.Command("go", "build", "-o", "go_whorship_band/wompi").Output()
	if err != nil {
		color.Red(err.Error())
		log.Fatal(err)
	}
	color.Green("  go_whorship_band/* was created successfully")
	color.Blue(">>>Building Project start 🚀 <<<")
}
