package main

import (
	"github.com/juanfer2/whorship_band/cmd"
	"github.com/juanfer2/whorship_band/src/shared/infrastructure/config"
)

func main() {
	config.LoadEnv()
	cmd.Execute()
}
