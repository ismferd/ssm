package main

import "github.com/ismferd/ssm/package/parameterstore"

func main() {
	config := &parameterstore.AWSConfig{Region: "us-east-1"}
	p := parameterstore.New(config)
	param := &parameterstore.ParemeterString{Name: "/first/try"}
	p.GetParam(param)
}
