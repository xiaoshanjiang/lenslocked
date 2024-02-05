package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "favorite-color", "blue")
	value := ctx.Value("favorite-color")
	fmt.Println(value)
}
