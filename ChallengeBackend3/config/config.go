package config

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

var path = "config.yml"

type config struct {
	Email map[string]string `json:"Email" yaml:"Email"`
}

var Cfg *config = &config{}

func init() {
	fmt.Printf("Verificando Arquivo de configuração: ")

	_, err := os.Stat(path)

	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("não existe")
		err = createConfigFile()
		loadConfigFile()

		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println("OK")
		loadConfigFile()
	}
}

func createConfigFile() error {
	Cfg.Email = map[string]string{
		"SMTP Server": "smtp.mailtrap.io",
		"Usuario":     "nomedeusuario",
		"Porta":       "2525",
		"Senha":       "senhadousuario",
	}

	bs, _ := yaml.Marshal(Cfg)
	return os.WriteFile(path, bs, 0600)
}

func loadConfigFile() {

	bs, err := os.ReadFile(path)

	if err != nil {
		fmt.Println(err)
	}

	err = yaml.Unmarshal(bs, Cfg)

	for key, value := range Cfg.Email {
		os.Setenv(key, value)
	}

	if err != nil {
		fmt.Println(err)
	}
}
