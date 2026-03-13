package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
	} `yaml:"database"`
}

func Load() (*Config, error) {
	// Cross-check: Uso de os.Getenv para capturar a variável de ambiente
	env := os.Getenv("env")
	if env == "" {
		env = "local"
	}

	// Cross-check: Caminho dinâmico baseado no ambiente
	fileName := "configs/" + env + ".yml"
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var cfg Config
	// Cross-check: Uso de endereço &cfg para decodificação
	err = yaml.NewDecoder(file).Decode(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
