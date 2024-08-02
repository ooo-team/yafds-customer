package customer

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"

	_ "github.com/lib/pq"
	common "github.com/ooo-team/yafds/internal/app/common"
	model "github.com/ooo-team/yafds/internal/model/customer"
	def "github.com/ooo-team/yafds/internal/repository"
	"github.com/ooo-team/yafds/internal/repository/customer/converter"
	repoModel "github.com/ooo-team/yafds/internal/repository/customer/model"
)

type repository struct {
	db *sql.DB
}

func NewRepository() *repository {
	return &repository{}
}

func (r *repository) GetDB() *sql.DB {

	if r.db != nil {
		return r.db
	}
	var err error

	host := common.LoadEnvVar("dbHost")
	port, err := strconv.Atoi(common.LoadEnvVar("dbPort"))
	if err != nil {
		panic("cannot convert string dbPort to int")
	}
	user := common.LoadEnvVar("dbUser")
	password := common.LoadEnvVar("dbPassword")
	dbname := common.LoadEnvVar("dbName")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	r.db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	return r.db
}

func (r *repository) Create(ctx context.Context, customerID uint32, info *model.CustomerInfo) error {
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
	tx, err := r.GetDB().BeginTx(ctx, nil)

	if err != nil {
		return err
	}
	defer tx.Rollback()

	tx.ExecContext(ctx,
		`insert into customers
		(
		id,
		phone,
		email,
		address
		)
		values($1, $2, $3, $4)`,
		repo_entity.ID,
		repo_entity.Info.Phone,
		repo_entity.Info.Email,
		repo_entity.Info.Address,
	)

	tx.ExecContext(ctx,
		`insert into h_customers 
		(
		customer_id, 
		created_at, 
		modified_at
		) 
		values ($1, $2, $3)`,
		repo_entity.ID,
		repo_entity.CreatedAt,
		repo_entity.UpdatedAt)

	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
}

// Get implements repository.CustomerRepository.
func (r *repository) Get(ctx context.Context, customerID uint32, need_metainfo bool) (*model.CustomerInfo, error) {

	rows, err := r.GetDB().QueryContext(ctx,
		`select c.phone,
				c.email,
				c.address
			from customers c where c.id = $1`, customerID)

	if err != nil {
		log.Println(err.Error())
	}
	defer rows.Close()

	var phone string
	var email string
	var address string

	if !rows.Next() {
		log.Println("Could not find customer")
		return nil, &common.NotFoundError{Message: "Could not find customer"}
	}

	if err := rows.Scan(&phone, &email, &address); err != nil {
		log.Println(err.Error())
	}

	info := repoModel.CustomerInfo{Phone: phone, Email: email, Address: address}

	customer_info := converter.ToCustomerInfoFromRepo(info)
	return &customer_info, nil
}

var _ def.CustomerRepository = (*repository)(nil)
