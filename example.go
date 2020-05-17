package main

import (
  "github.com/ismferd/ssm/package/parameterstore"
  "fmt"
)

func main() {
	config := &parameterstore.AWSConfig{Region: "us-east-1"}
	p := parameterstore.New(config)
	param := &parameterstore.ParemeterString{Name: "/first/try"}
	value, _ := p.GetParam(param)
	fmt.Println(value)
}
