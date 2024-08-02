package main

import (
	"context"
	"log"

	"github.com/ooo-team/yafds/internal/app"
	common "github.com/ooo-team/yafds/pkg"
)

func init() {
	common.InitEnv()
}

func main() {

	ctx := context.Background()
	a, err := app.NewApp(ctx)

	if err != nil {
		log.Panic(err.Error())
	}
	a.Run()
}
