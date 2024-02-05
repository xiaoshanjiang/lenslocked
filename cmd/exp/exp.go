package main

import (
	stdctx "context"
	"fmt"

	"github.com/xiaoshanjiang/lenslocked/context"
	"github.com/xiaoshanjiang/lenslocked/models"
)

type ctxKey string

const (
	favoriteColorKey ctxKey = "favorite-color"
)

func main() {
	ctx := stdctx.Background()

	user := models.User{
		Email: "jon@show.io",
	}
	ctx = context.WithUser(ctx, &user)

	retrievedUser := context.User(ctx)
	fmt.Println(retrievedUser.Email)
}
