package config

import (
	"github.com/Yumax-panda/TT-Blog/storage"
	"github.com/spf13/viper"
)

type Config struct {
	DB struct {
		Username string
		Password string
		Hostname string
		Port     int
		Name     string
	}
	S3 struct {
		Bucket         string
		Region         string
		Endpoint       string
		AccessKey      string
		SecretKey      string
		ForcePathStyle bool
	}
	ExternalAuth struct {
		Discord struct {
			ClientID       string
			ClientSecret   string
			AllowedGuildID string
		}
	}
}

func init() {
	// Database
	viper.SetDefault("db.username", "root")
	viper.SetDefault("db.password", "password")
	viper.SetDefault("db.hostname", "mysql")
	viper.SetDefault("db.port", 3306)
	viper.SetDefault("db.name", "ttblog")

	// S3
	viper.SetDefault("s3.bucket", "")
	viper.SetDefault("s3.region", "")
	viper.SetDefault("s3.endpoint", "")
	viper.SetDefault("s3.accessKey", "")
	viper.SetDefault("s3.secretKey", "")
	viper.SetDefault("storage.s3.forcePathStyle", false)

	// External Auth
	viper.SetDefault("externalAuth.discord.clientId", "")
	viper.SetDefault("externalAuth.discord.clientSecret", "")
	viper.SetDefault("externalAuth.discord.allowedGuildId", "")
}

func (c Config) getFileStorage() (storage.FileStorage, error) {
	return storage.NewS3FileStorage(
		c.S3.Bucket,
		c.S3.Region,
		c.S3.Endpoint,
		c.S3.AccessKey,
		c.S3.SecretKey,
		c.S3.ForcePathStyle,
	)
}
