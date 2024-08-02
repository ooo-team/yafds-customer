package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ooo-team/yafds/internal/app"
	common "github.com/ooo-team/yafds/internal/app/common"
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
	fmt.Println("customer web app")
}
