package customer

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
	common "github.com/ooo-team/yafds-common/pkg"
	commonRepo "github.com/ooo-team/yafds-common/pkg/repository"
	model "github.com/ooo-team/yafds-customer/internal/model/customer"
	def "github.com/ooo-team/yafds-customer/internal/repository"
	"github.com/ooo-team/yafds-customer/internal/repository/customer/converter"
	repoModel "github.com/ooo-team/yafds-customer/internal/repository/customer/model"
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

	r.db = commonRepo.GetDB()

	return r.db
}

func (r *repository) Create(ctx context.Context, customerID uint32, info *model.CustomerInfo) error {
	var time_ = time.Now()
	repoEntity := repoModel.Customer{
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
	defer func() {
		tx.Rollback()
	}()

	_, err = tx.ExecContext(ctx,
		`insert into customers
		(
		id,
		phone,
		email,
		address
		)
		values($1, $2, $3, $4)`,
		repoEntity.ID,
		repoEntity.Info.Phone,
		repoEntity.Info.Email,
		repoEntity.Info.Address,
	)

	if err != nil {
		log.Println(err.Error())
		return err
	}

	_, err = tx.ExecContext(ctx,
		`insert into h_customers 
		(
		customer_id, 
		createdAt, 
		modifiedAt
		) 
		values ($1, $2, $3)`,
		repoEntity.ID,
		repoEntity.CreatedAt,
		repoEntity.UpdatedAt)

	if err != nil {
		log.Println(err.Error())
		return err
	}

	if err = tx.Commit(); err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

// Get implements repository.CustomerRepository.
func (r *repository) Get(ctx context.Context, customerID uint32) (*model.Customer, error) {

	queryText := `
	select c.phone,
			c.email,
			c.address,
			hc1.createdAt,
			hc.modifiedAt
		from customers c 
		join h_customers hc1 
			on hc1.customer_id = c.id
		and hc1.createdAt is not null
		join h_customers hc 
			on hc.customer_id = c.id
		and coalesce(hc.modifiedAt, 
						TIMESTAMP '0001-01-01 00:00:00') = (select max(coalesce(hc2.modifiedAt, TIMESTAMP '0001-01-01 00:00:00'))
															from h_customers hc2 
															where hc2.customer_id = hc.customer_id)

		  where c.id = $1`
	log.Printf(queryText, customerID)
	rows, err := r.GetDB().QueryContext(ctx, queryText, customerID)

	if err != nil {
		log.Println(err.Error())
	}
	defer rows.Close()

	var phone string
	var email string
	var address string
	var createdAt time.Time
	var modifiedAt sql.NullTime

	if !rows.Next() {
		err = &common.NotFoundError{Message: "Could not find customer"}
		log.Println(err.Error())
		return nil, err
	}

	if err := rows.Scan(&phone, &email, &address, &createdAt, &modifiedAt); err != nil {
		log.Println(err.Error())
	}

	info := repoModel.CustomerInfo{Phone: phone, Email: email, Address: address}

	customer := converter.ToCustomerFromRepo(&repoModel.Customer{
		ID:        customerID,
		Info:      info,
		CreatedAt: createdAt,
		UpdatedAt: modifiedAt,
	})

	return customer, nil
}

var _ def.CustomerRepository = (*repository)(nil)
