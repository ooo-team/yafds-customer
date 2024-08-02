package customer

import (
	"context"
	"testing"

	model "github.com/ooo-team/yafds/internal/model/customer"
	"github.com/stretchr/testify/assert"
)

func TestGetDB(t *testing.T) {
	assert := assert.New(t)
	var r repository

	assert.NotNil(r.GetDB(), "Failed to connect to DB")

}

func TestCreate(t *testing.T) {
	ctx := context.Background()

	var r repository
	info := model.CustomerInfo{Phone: "+79999999999", Email: "", Address: ""}

	err := r.Create(ctx, 100, &info)
	if err != nil {
		panic(err.Error())
	}

}
