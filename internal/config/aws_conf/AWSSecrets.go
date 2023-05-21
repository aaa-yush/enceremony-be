package aws_conf

import (
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"os"
)

type ConfigurationService interface {
	GetConfigString(envKey, secretName string) (string, error)
}

type AWSConfigurationService struct {
}

func NewConfigurationService() ConfigurationService {
	return &AWSConfigurationService{}
}

// GetConfigString ...
// In stage/dev export the config json as a environment variable
// (in your shell/docker env ) $ export SUMMARY_ENVS={"app": "my-app", "mysql-pass" : "123"}
// GetConfigString("SUMMARY_ENVS", "")
// In prod when function won't find the envKey and it calls the SecretManager service to get the
// config json
// GetConfigString("", "summary")
func (a *AWSConfigurationService) GetConfigString(envKey, secretName string) (string, error) {
	// Try reading the environment variables, if it's empty we assumes secrets manager to have it
	jsonFromEnv := os.Getenv(envKey)
	if jsonFromEnv == "" {
		var err error
		jsonFromEnv, err = GetAWSSecret(secretName)
		if err != nil {
			return "", err
		}
	}
	return jsonFromEnv, nil
}

// GetAWSSecret Use this code snippet in your app.
// If you need more information about configurations or implementing the sample code, visit the AWS docs:
// https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/setting-up.html
func GetAWSSecret(secretName string) (string, error) {
	region := "ap-south-1"

	//Create a Secrets Manager client
	svc := secretsmanager.New(session.New(),
		aws.NewConfig().WithRegion(region))
	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secretName),
		VersionStage: aws.String("AWSCURRENT"), // VersionStage defaults to AWSCURRENT if unspecified
	}

	// In this sample we only handle the specific exceptions for the 'GetSecretValue' API.
	// See https://docs.aws.amazon.com/secretsmanager/latest/apireference/API_GetSecretValue.html

	result, err := svc.GetSecretValue(input)
	if err != nil {
		return "", err
	}

	// Decrypts secret using the associated KMS CMK.
	// Depending on whether the secret is a string or binary, one of these fields will be populated.
	var secretString string
	if result.SecretString != nil {
		secretString = *result.SecretString
	} else {
		return "", errors.New("empty string from aws config api")
	}
	return secretString, nil
	// Your code goes here.
}
