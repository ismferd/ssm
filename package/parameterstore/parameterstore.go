// Package parameterstore provide methods to get data from
// AWS Parameterstore
package parameterstore

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/aws/aws-sdk-go/service/ssm/ssmiface"
	"fmt"
)

// New creates an AWS Session Manager Client
func New(config *AWSConfig) *Client {
	c := &Client{
		config: config,
	}

	s := c.newSession(config)
	c.api = ssm.New(s)
	return c
}

// ParameterStore returns a representation of the Secrets Manager API
func (c *Client) ParameterStore() ssmiface.SSMAPI {
	return c.api
}


func (c *Client) newSession(config *AWSConfig) *session.Session {
	// Initialize config with error verbosity
	sess := aws.NewConfig().WithCredentialsChainVerboseErrors(true)

	if config.Region != "" {
		sess = sess.WithRegion(config.Region)
	}

	opts := session.Options{
		Config: *sess,
	}

	return session.Must(session.NewSessionWithOptions(opts))
}

// GetParam lalala
func (c *Client) GetParam(spec *ParemeterString) (string, error){
	//keyname := "/first/try"
	withDecryption := false
	param, err := c.api.GetParameter(&ssm.GetParameterInput{
		Name:           &spec.Name,
		WithDecryption: &withDecryption,
	})
	if err != nil {
		fmt.Println(err)
	}
	value := *param.Parameter.Value
	fmt.Println(value)
}
