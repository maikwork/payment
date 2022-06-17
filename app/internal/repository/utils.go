package repository

import (
	"database/sql"
	"time"

	"github.com/maikwork/restPayments/internal/model"
)

func getFromScan(rows *sql.Rows) model.Payment {
	tmp := model.Payment{}

	timeCreat := ""
	timeChang := ""

	rows.Scan(tmp.ID, tmp.Client.ID, tmp.Client.Email, tmp.Amount, tmp.Currency, &timeCreat, &timeChang, tmp.Status)

	tmp.Created, _ = time.Parse(time.RFC3339, timeCreat)
	tmp.Changed, _ = time.Parse(time.RFC3339, timeChang)

	return tmp
}

func isChangedStatus(db *sql.DB, id int) bool {
	var status model.Status
	var tmp int

	row := db.QueryRow("select status_pay from payment where id=$1", id)

	row.Scan(&tmp)

	status = model.Status(tmp)

	s := status == model.SUCCESS
	c := status == model.CANCEL
	if !c || !s {
		return true
	}
	return false
}
