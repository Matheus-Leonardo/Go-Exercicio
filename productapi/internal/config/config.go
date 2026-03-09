package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// Config define a estrutura para mapear o arquivo YAML
type Config struct {
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
	} `yaml:"database"`
}

// Load lê o arquivo de configuração baseado na variável de ambiente "env"
// Cross-check: Retorna um ponteiro *Config para evitar cópia de memória
func Load() (*Config, error) {
	env := os.Getenv("env")
	if env == "" {
		env = "local" // Default para desenvolvimento local
	}

	fileName := env + ".yml"
	// O caminho assume que você executará o binário da raiz de /productapi
	path := filepath.Join("configs", fileName)

	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config
	// Cross-check: Uso do endereço & para que o Unmarshal preencha a struct original
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
