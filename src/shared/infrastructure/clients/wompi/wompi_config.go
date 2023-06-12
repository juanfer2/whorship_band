package wompi_client

import (
	"io/ioutil"
	"log"
	"strings"

	"github.com/juanfer2/whorship_band/src/shared/utilities"
	"gopkg.in/yaml.v3"
)

type ConfigClient struct {
	Url        string `yaml:"url"`
	PublicKey  string `yaml:"public_key"`
	PrivateKey string `yaml:"private_key"`
}

func NewConfigClient() *ConfigClient {
	configClient := ConfigClient{}
	root := utilities.GetRootFolder()
	enviroment := strings.ToLower(utilities.GetEnv("MODE"))

	yamlFile, err := ioutil.ReadFile(root +
		"/src/shared/infrastructure/clients/wompi/config/wompi_credentias_" + enviroment + ".yml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, &configClient)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return &configClient
}
