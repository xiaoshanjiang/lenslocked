package main

import (
	"context"
	"fmt"
)

type ctxKey string

const (
	favoriteColorKey ctxKey = "favorite-color"
)

func main() {
	ctx := context.Background()
	// Our code uses our unexported `ctxKey` type. Even though the value still
	// appears to be a string with the contents "favorite-color", Go and the
	// context package treat this different from a string with the value
	// "favorite-color"
	ctx = context.WithValue(ctx, favoriteColorKey, "blue")

	// This key has a type of string, not ctxKey.
	ctx = context.WithValue(ctx, "favorite-color", 0xFF0000)

	// Each key has a unique type, so the keys won't match and we will get
	// unique values for each key.
	value1 := ctx.Value(favoriteColorKey)
	value2 := ctx.Value("favorite-color")
	fmt.Println(value1)
	fmt.Println(value2)
}
