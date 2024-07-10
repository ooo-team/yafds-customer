package customer

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	_ "github.com/lib/pq"
	model "github.com/ooo-team/yafds/internal/model/customer"
	def "github.com/ooo-team/yafds/internal/repository"
	"github.com/ooo-team/yafds/internal/repository/customer/converter"
	repoModel "github.com/ooo-team/yafds/internal/repository/customer/model"
)

func load_env_variable(var_name string) string {
	var_, exists := os.LookupEnv(var_name)

	err_msg := fmt.Sprintf("Env variable %s is not set", var_name)
	if !exists {
		panic(err_msg)
	}
	return var_
}

func (r *repository) GetDB() *sql.DB {

	if r.db != nil {
		return r.db
	}
	var err error

	host := load_env_variable("dbHost")
	port, err := strconv.Atoi(load_env_variable("dbPort"))
	if err != nil {
		panic("cannot convert string dbPort to int")
	}
	user := load_env_variable("dbUser")
	password := load_env_variable("dbPassword")
	dbname := load_env_variable("dbName")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	r.db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	return r.db
}

type repository struct {
	db *sql.DB
}

// Create implements repository.CustomerRepository.
func (r *repository) Create(ctx context.Context, customerID uint64, info *model.CustomerInfo) error {
	var time_ = time.Now()
	repo_entity := repoModel.Customer{
		ID: customerID,
		Info: repoModel.CustomerInfo{
			Phone:   info.Phone,
			Email:   info.Email,
			Address: info.Address,
		},
		CreatedAt: time_,
		UpdatedAt: sql.NullTime{Time: time_, Valid: false},
	}
	// tx, err := r.GetDB().BeginTx(ctx, nil)

	// if err != nil {
	// 	return err
	// }
	// defer tx.Rollback()

	_, err := r.GetDB().Exec("INSERT INTO Customers"+
		"("+
		"id,"+
		"phone,"+
		" email,"+
		"address"+
		") "+
		"VALUES($1, $2, $3, $4)",
		repo_entity.ID,
		repo_entity.Info.Phone,
		repo_entity.Info.Email,
		repo_entity.Info.Address,
	)
	if err != nil {
		panic(err)
	}

	// tx.ExecContext(ctx, "INSERT INTO HCustomers (customer_id, created_at, modofied_at) VALUES ($1, $2, $3)", repo_entity.ID, repo_entity.CreatedAt, repo_entity.UpdatedAt)

	// Commit the transaction.
	// if err = tx.Commit(); err != nil {
	// 	return err
	// }
	return nil
}

// Get implements repository.CustomerRepository.
func (r *repository) Get(ctx context.Context, customerID uint64) (*model.Customer, error) {
	customer := &repoModel.Customer{
		ID:        228,
		Info:      repoModel.CustomerInfo{Phone: "+79999999999", Email: "email", Address: "address"},
		CreatedAt: time.Now(),
		UpdatedAt: sql.NullTime{Time: time.Now(), Valid: true},
	}

	return converter.ToCustomerFromRepo(customer), nil
}

var _ def.CustomerRepository = (*repository)(nil)
