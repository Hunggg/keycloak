package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config represents the structure of your JSON configuration file
type Config struct {
	Owner OwnerConfig
	Database DatabaseConfig
	ECS EcsConfig
}

type OwnerConfig struct {
	AccessKey string `mapstructure:"access_key"`
	SecretKey string `mapstructure:"secret_key"`
	Region string `mapstructure:"region"`
}

type DatabaseConfig struct {
	Postgres PostgresConfig
	DynamoDB DynamoDBConfig
}

type PostgresConfig struct {
	Engine string `mapstructure:"engine"`
	InstanceClass string `mapstructure:"instanceClass"`
	AllocatedStorage int64 `mapstructure:"allocatedStorage"`
	Name string `mapstructure:"name"`
	Password string `mapstructure:"password"`
	UserName string `mapstructure:"username"`
	SkipFinalSnapshot string `mapstructure:"skipFinalSnapshot"`
}

type DynamoDBConfig struct {
}

type EcsConfig struct {
	Cluster ClusterConfig
	KeyCloak KeyCloakContainerConfig `mapstructure:"keycloak_container"`
}

type KeyCloakContainerConfig struct {
	Name string `mapstructure:"name"`
	Family string `mapstructure:"family"`
	Cpu int64 `mapstructure:"cpu"`
	Memory int64 `mapstructure:"memory"`
	Image string `mapstructure:"image"`
	ContainerPort string `mapstructure:"container_port"`
}

type ClusterConfig struct {
	Name string `mapstructure:"name"`
}

func Get_Config() (*Config, error) {
	// Set the path to your JSON configuration file
	configFilePath := "./config.json"

	// Initialize Viper
	viper.SetConfigFile(configFilePath)

	// Read the configuration file
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error reading config file:", err)
		return nil, err
	}

	// Create an instance of the Config struct to store the configuration
	var config Config

	// Unmarshal the configuration into the Config struct
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Println("Error unmarshalling config:", err)
		return nil, err
	}

	return  &config, nil
}
