package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/maikwork/restPayments/internal/model"
)

const (
	host   = "postgresdb"
	port   = "5432"
	name   = "test"
	pass   = "test"
	dbname = "payment"
)

func getDSN() string {
	return fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", name, pass, host, port, dbname)
}

func Connection(dbCnf model.DataBase) (*sql.DB, error) {
	fmt.Println(dbCnf.DSN)
	db, err := sql.Open("postgres", getDSN())
	if err != nil {
		log.Print("Can't open the connection")
		return nil, err
	}

	return db, nil
}
