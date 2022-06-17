package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/maikwork/restPayments/internal/model"
	"github.com/sirupsen/logrus"
)

type Paymenter interface {
	GetAllPay() ([]model.Payment, error)
	GetStatusByID(id int) (model.Status, error)
	GetPaysByIDClient(id int) ([]model.Payment, error)
	GetPaysByEmail(email string) ([]model.Payment, error)
	CreatePayment(model.Payment) error
	CancelPayByID(id int) error
	ChangeStatusByID(id int, status model.Status) error
}

type Repository struct {
	DB *sql.DB
}

func New(d *sql.DB) *Repository {
	return &Repository{DB: d}
}

func (r *Repository) CreatePayment(p model.Payment) error {
	db := r.DB

	t1 := p.Created.Format(time.RFC3339)
	t2 := p.Changed.Format(time.RFC3339)
	db.Exec("insert into payment(id_client, amount, currency, time_created, time_changed, status_pay) "+
		"values($1, $2, $3, $4, $5, $6)", p.Client.ID, p.Amount, p.Currency, t1, t2, p.Status)

	return nil
}

func (r *Repository) GetStatusByID(id int) (model.Status, error) {
	db := r.DB

	row := db.QueryRow("select status_pay from payment where id=$1", id)

	status := 0
	row.Scan(&status)
	return model.Status(status), nil
}

func (r *Repository) GetPaysByIDClient(id int) ([]model.Payment, error) {
	db := r.DB
	result := make([]model.Payment, 0)

	pays, err := db.Query("select * from payment where id_client=$1", id)
	if err != nil {
		logrus.WithError(err).Info(err)
	}

	for pays.Next() {
		result = append(result, getFromScan(pays))
	}
	return result, nil
}

func (r *Repository) GetPaysByEmail(email string) ([]model.Payment, error) {
	db := r.DB
	result := make([]model.Payment, 0)

	pays, err := db.Query("select * from payment where email=$1", email)
	if err != nil {
		logrus.WithError(err).Info(err)
	}

	for pays.Next() {
		result = append(result, getFromScan(pays))
	}
	return result, nil
}

func (r *Repository) CancelPayByID(id int) error {
	var err error

	db := r.DB

	if isChangedStatus(db, id) {
		db.Exec("update payment set status_pay=$1 where id=$2", int(model.CANCEL), id)
		return nil
	}

	err = fmt.Errorf("Сannot be cancelled because it is in the \"unchangeable\" status")

	return err
}

func (r *Repository) GetAllPay() ([]model.Payment, error) {
	db := r.DB
	result := make([]model.Payment, 0)

	pays, err := db.Query("select * from payment join client on payment.id_client=client.id")
	if err != nil {
		logrus.WithError(err).Info(err)
	}

	for pays.Next() {
		result = append(result, getFromScan(pays))
	}
	return result, nil
}

func (r *Repository) ChangeStatusByID(id int, status model.Status) error {
	var err error
	var update string

	db := r.DB

	if isChangedStatus(db, id) {
		update = "update payment set status_pay=$1 where id=$2"
		db.Exec(update, int(status), id)
		return nil
	}

	err = fmt.Errorf("Сannot be cancelled because it is in the \"unchangeable\" status")

	return err
}
