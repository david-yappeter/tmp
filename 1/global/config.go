package global

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"

	"gopkg.in/yaml.v2"
)

const (
	ApiVersion = "1.0.0"

	EnvironmentDevelopment = "development"
	EnvironmentProduction  = "production"
	EnvironmentTesting     = "testing"

	StaticStorageFs = "public"
)

type PostgresConfig struct {
	Host     string `yaml:"host"`
	Port     uint16 `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

type JwtConfig struct {
	SecretKey string `yaml:"secret_key"`
}

type YamlConfig struct {
	MigrationDir       string         `yaml:"-"`
	AppName            string         `yaml:"app_name"`
	Environment        string         `yaml:"environment"`
	IsDebug            bool           `yaml:"debug"`
	Port               uint           `yaml:"port"`
	CorsAllowedOrigins []string       `yaml:"cors_allowed_origins"`
	Postgres           PostgresConfig `yaml:"postgres"`
	JwtConfig          string         `yaml:"jwt"`
}

var config YamlConfig

func init() {
	if err := LoadConfig(); err != nil {
		panic(err)
	}
}

func LoadConfig() error {
	baseDir, err := os.Getwd()
	if err != nil {
		return err
	}

	baseDir = strings.TrimRight(strings.ReplaceAll(baseDir, "\\", "/"), "/")

	config.MigrationDir = fmt.Sprintf("%s/database/migration", baseDir)

	if _, err := os.Stat(fmt.Sprintf("%s/conf.yml", baseDir)); err != nil {
		_, filename, _, _ := runtime.Caller(0)
		baseDir = path.Join(path.Dir(filename), "../")
	}

	yamlFilePath := fmt.Sprintf("%s/conf.yml", baseDir)
	if _, err := os.Stat(yamlFilePath); err != nil {
		return err
	}

	yamlFile, err := os.ReadFile(yamlFilePath)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return err
	}

	return nil
}

func GetConfig() YamlConfig {
	return config
}

func GetJwtSecretKey() string {
	return config.JwtConfig
}

func GetAppName() string {
	return config.AppName
}

func GetPostgresConfig() PostgresConfig {
	return config.Postgres
}

func GetMigrationDir() string {
	return config.MigrationDir
}

func SetEnvironment(environment string) {
	switch environment {
	case EnvironmentDevelopment, EnvironmentProduction, EnvironmentTesting:
		config.Environment = environment
	}
}

func EnableDebug() {
	config.IsDebug = true
}

func DisableDebug() {
	config.IsDebug = false
}

func IsProduction() bool {
	return config.Environment == EnvironmentProduction
}

func IsDevelopment() bool {
	return config.Environment == EnvironmentDevelopment
}

func IsTesting() bool {
	return config.Environment == EnvironmentTesting
}

func IsDebug() bool {
	return config.IsDebug
}
