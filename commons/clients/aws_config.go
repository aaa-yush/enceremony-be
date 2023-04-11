package clients

import (
	"context"
	"enceremony-be/commons/clients/models"

	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
)

const awsRegion = "ap-south-1"

var awsConfigOnce sync.Once
var awsCfg aws.Config
var errFetchConfig error

func getAWSConfig(config *models.AwsConf) (aws.Config, error) {
	awsConfigOnce.Do(func() {
		customResolver := aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
			if endpoint := config.Endpoint; endpoint != "" {
				return aws.Endpoint{
					PartitionID:   "aws",
					URL:           endpoint,
					SigningRegion: awsRegion,
				}, nil
			}

			// returning EndpointNotFoundError will allow the service to fallback to it's default resolution
			return aws.Endpoint{}, &aws.EndpointNotFoundError{}
		})

		awsCfg, errFetchConfig = awsConfig.LoadDefaultConfig(context.TODO(),
			awsConfig.WithRegion(awsRegion),
			awsConfig.WithEndpointResolver(customResolver),
		)

	})

	return awsCfg, errFetchConfig
}
