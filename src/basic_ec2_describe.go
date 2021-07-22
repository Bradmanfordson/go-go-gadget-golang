package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

var role_arn string = "fill in with your role arn"

func main() {
	sess := session.Must(session.NewSession())

	creds := stscreds.NewCredentials(
		sess,
		role_arn,
  )

	ec2_service := ec2.New(
		sess, 
    &aws.Config{
			Credentials: creds,
			Region:      aws.String("us-east-1"),
		},
	)

	result, err := ec2_service.DescribeInstances(nil)

	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println("Success", result)
	}
}
