package model

import "time"

type Status = uint8

const (
	NEW Status = iota
	SUCCESS
	FAIL
	CANCEL
)

type Client struct {
	ID    int
	Email string
}

type Payment struct {
	ID       int
	Client   Client
	Amount   int
	Currency rune
	Created  time.Time
	Changed  time.Time
	Status   Status
}

func NewPayment(id int, email string, amount int, currency rune) Payment {
	t := time.Now()
	return Payment{
		Client: Client{
			ID:    id,
			Email: email,
		},
		Amount:   amount,
		Currency: currency,
		Created:  t,
		Changed:  t,
		Status:   NEW,
	}
}
