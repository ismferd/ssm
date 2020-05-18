
package parameterstore

import (
	 "github.com/aws/aws-sdk-go/service/ssm/ssmiface"
)

// AWSConfig store configuration used to initialize
// AWS SSM client.
type AWSConfig struct {
	Region string
}

// Client represents an AWS SSM client
// maps to ProviderServices
type Client struct {
	config *AWSConfig
	api    ssmiface.SSMAPI
}

// ParameterString is a concret representation
// of an AWS SSM Parameter Store String in plain text
type ParemeterString struct {
	Name string
}
