package customer

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"

	_ "github.com/lib/pq"
	model "github.com/ooo-team/yafds/internal/model/customer"
	def "github.com/ooo-team/yafds/internal/repository"
	"github.com/ooo-team/yafds/internal/repository/customer/converter"
	repoModel "github.com/ooo-team/yafds/internal/repository/customer/model"
	common "github.com/ooo-team/yafds/pkg"
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
func (r *repository) Get(ctx context.Context, customerID uint32) (*model.Customer, error) {

	rows, err := r.GetDB().QueryContext(ctx,
		`select c.phone,
				c.email,
				c.address,
				hc1.created_at,
				hc.modified_at
		   from customers c 
		   join h_customers hc1 
			 on hc1.customer_id = c.id
			and hc1.created_at is not null
		   join h_customers hc 
			 on hc.customer_id = c.id
			and coalesce(hc.modified_at, 
			             TIMESTAMP '1900-01-01 00:00:00') = (select max(coalesce(hc2.modified_at, TIMESTAMP '1900-01-01 00:00:00'))
															   from h_customers hc2 
															  where hc2.customer_id = hc.customer_id)

		  where c.id = $1`, customerID)

	if err != nil {
		log.Println(err.Error())
	}
	defer rows.Close()

	var phone string
	var email string
	var address string
	var created_at time.Time
	var modified_at sql.NullTime

	if !rows.Next() {
		err = &common.NotFoundError{Message: "Could not find customer"}
		log.Println(err.Error())
		return nil, err
	}

	if err := rows.Scan(&phone, &email, &address, &created_at, &modified_at); err != nil {
		log.Println(err.Error())
	}

	info := repoModel.CustomerInfo{Phone: phone, Email: email, Address: address}

	customer := converter.ToCustomerFromRepo(&repoModel.Customer{
		ID:        customerID,
		Info:      info,
		CreatedAt: created_at,
		UpdatedAt: modified_at,
	})

	return customer, nil
}

var _ def.CustomerRepository = (*repository)(nil)
