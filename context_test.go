package main

import (
	"context"
	"fmt"
	"testing"
)

func TestContext(t *testing.T) {
	ctx := context.Background()
	fmt.Println(ctx)

	todo := context.TODO()
	fmt.Println(todo)
}
