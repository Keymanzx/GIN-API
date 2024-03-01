package config

import (
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var config *viper.Viper

func Init(env string) error {
	config = viper.New()
	config.SetConfigType("yaml")
	// config.SetConfigName("default")
	// config.AddConfigPath("config/")

	if err := config.ReadInConfig(); err != nil {
		return err
	}

	envConfig := viper.New()
	envConfig.SetConfigType("yaml")
	// envConfig.AddConfigPath("config/")
	// envConfig.SetConfigName(env)

	if err := envConfig.ReadInConfig(); err != nil {
		return err
	}

	config.MergeConfigMap(envConfig.AllSettings())

	return nil
}

func GetMongoDBClient() (*mongo.Client, error) {
	// Connect to MongoDB using configuration settings
	uri := config.GetString("mongodb://localhost:27017")
	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(nil, clientOptions)
	if err != nil {
		return nil, err
	}

	return client, nil
}


func GetConfig() *viper.Viper {
	return config
}

