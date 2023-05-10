package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/service/s3"

	"dawix/lambda/structs"
	"dawix/lambda/utilities"
)

func HandleRequest(ctx context.Context, myevent structs.Request) (structs.Response, error) {
	profile := utilities.GetEnv("PROFILE", "")
	region := utilities.GetEnv("REGION", "")
	sess, err := utilities.InitSession(profile, region)
	if err != nil {
		utilities.ExitErrorf("Unable to create session buckets, %s", err)
	}

	svc := s3.New(sess)

	input := &s3.ListBucketsInput{}
	result, err := svc.ListBuckets(input)
	if err != nil {
		utilities.ExitErrorf("Unable to list buckets, %v", err)
	}

	fmt.Printf("%s", result.Buckets)

	return structs.Response{
		Message: fmt.Sprintf("Hello, %s", myevent.Name),
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
