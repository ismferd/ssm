
package parameterstore

import (
	 "github.com/aws/aws-sdk-go/service/ssm/ssmiface"
)

// AWSConfig store configuration used to initialize
// secrets manager client.
type AWSConfig struct {
	Region string
}

// Client represents an AWS Secrets Manager client
//
// maps to ProviderServices
type Client struct {
	config *AWSConfig
	api    ssmiface.SSMAPI
}

// SecretString is a concret representation
// of an AWS Secrets Manager Secret String
type ParemeterString struct {
	Name string
}
