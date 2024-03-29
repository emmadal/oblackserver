package helpers

import (
	"oblackserver/db"
	"os"
	"github.com/joho/godotenv"
)

// EnvCloudName return the cloud name
func EnvCloudName() string {
	defer db.RecoverEnv()
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}
	cloudName := os.Getenv("CloudName")
	return cloudName
}

// EnvCloudSecretAPI return the cloud secret key
func EnvCloudSecretAPI() string {
	defer db.RecoverEnv()
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}
	cloudSecret := os.Getenv("CloudSecret")
	return cloudSecret
}

// EnvCloudAPIKey return the cloud API key
func EnvCloudAPIKey() string {
	defer db.RecoverEnv()
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}
	cloudAPIKey := os.Getenv("CloudAPIKey")
	return cloudAPIKey
}

// EnvCloudFolder return the cloud folder to store images
func EnvCloudFolder() string {
	defer db.RecoverEnv()
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}
	cloudFolder := os.Getenv("CloudFolder")
	return cloudFolder
}

// EnvSecretKey return the secret key
func EnvSecretKey() string {
	defer db.RecoverEnv()
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}
	secret := os.Getenv("Secret_Key")
	return secret
}

// EnvDomainNameKey return the domain name
func EnvDomainNameKey() string {
	defer db.RecoverEnv()
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}
	secret := os.Getenv("DomainName")
	return secret
}

// EnvMode return the status mode of server (dev/production)
func EnvMode() bool {
	defer db.RecoverEnv()
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}
	mode := os.Getenv("mode")
	if mode == "production" {
		return true
	}
	return false
}
