package main

import (
	"context"
	"log"

	common "github.com/ooo-team/yafds-common/pkg"
	"github.com/ooo-team/yafds-customer/internal/app"
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
