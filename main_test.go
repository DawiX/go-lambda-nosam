package main

import (
	"context"
	"fmt"
	"testing"

	"dawix/lambda/structs"
)

var evnt = &structs.Request{
	Name: "Test",
}

func TestLambda(t *testing.T) {
	ctx := context.Background()

	result, err := HandleRequest(ctx, *evnt)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
