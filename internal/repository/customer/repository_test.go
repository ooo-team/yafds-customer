package customer

import (
	"log"

	"context"
	"testing"

	"github.com/joho/godotenv"

	"github.com/stretchr/testify/assert"

	model "github.com/ooo-team/yafds/internal/model/customer"
	// common "github.com/ooo-team/yafds/internal/model/common"
)

func TestGetDB(t *testing.T) {
	assert := assert.New(t)
	var r repository

	assert.NotNil(r.GetDB(), "zalupa")

}

func TestCreate(t *testing.T) {

	if err := godotenv.Load("/home/dimoha_zadira/yafds/.env"); err != nil {
		log.Print("No .env file found")
	}
	var r repository
	var ctx context.Context

	r.Create(ctx, 100, &model.CustomerInfo{
		Phone:   "+79999999999",
		Email:   "DimohaZadira@gmail.com",
		Address: "zalupkina, 24",
	})

}
