package main

import (
	"context"
	"fmt"
	"strings"
)

type ctxKey string

const (
	favoriteColorKey ctxKey = "favorite-color"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, favoriteColorKey, "blue")
	anyValue := ctx.Value(favoriteColorKey)

	// This .(string) format attempts to assert that anyValue has a type of string
	// If it succeeds, ok will be true. Otherwise ok will be false.
	strValue, ok := anyValue.(string)
	if !ok {
		// anyValue is not a string!
		fmt.Println(anyValue, "is not a string")
		return
	} else {
		fmt.Println(strValue)
		fmt.Println(strings.HasPrefix(strValue, "b"))
	}
}
