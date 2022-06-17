package handler

import (
	"encoding/json"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/maikwork/restPayments/internal/model"
	"github.com/maikwork/restPayments/internal/repository"
	"github.com/sirupsen/logrus"
)

//POST cancel payment by id
func postCancelByID(r *gin.Engine, p repository.Paymenter) {
	r.POST("/cancel/:id", func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			logrus.WithError(err).Info("Don't currect id")
		}

		if err := p.CancelPayByID(id); err != nil {
			data, _ := json.Marshal(err.Error())
			ctx.Writer.Write(data)
		}
	})
}

//GET status by id
func getStatusByID(r *gin.Engine, p repository.Paymenter) {
	r.GET("/status/:id", func(ctx *gin.Context) {
		if ctx.FullPath() == "/payment/:id" {
			id, err := strconv.Atoi(ctx.Param("id"))
			if err != nil {
				logrus.WithError(err).Info("Don't currect id")
				return
			}

			status, err := p.GetStatusByID(id)
			if err != nil {
				logrus.WithError(err).Info(err)
			}
			data, _ := json.Marshal(status)
			ctx.Writer.Write(data)
		}
	})
}

//GET all payments
func getALLPayments(r *gin.Engine, p repository.Paymenter) {
	r.GET("/payement/", func(ctx *gin.Context) {
		if ctx.FullPath() == "/payment/" {
			pays, err := p.GetAllPay()
			if err != nil {
				logrus.WithError(err).Info(err)
			}

			data, _ := json.Marshal(pays)
			ctx.Writer.Write(data)
		}
	})
}

//GET payments by client_id
func getPaymentsByClientID(r *gin.Engine, p repository.Paymenter) {
	r.GET("/payment/client/:id", func(ctx *gin.Context) {
		if ctx.FullPath() == "payment/client/:id" {
			id, err := strconv.Atoi(ctx.Param("id"))
			if err != nil {
				logrus.WithError(err).Info("Don't currect id")
				return
			}

			pays, _ := p.GetPaysByIDClient(id)
			data, _ := json.Marshal(pays)
			ctx.Writer.Write(data)
		}
	})
}

//GET payments by client email
func getPaymentsByClientEmail(r *gin.Engine, p repository.Paymenter) {
	r.GET("/payment/client/:email", func(ctx *gin.Context) {
		if ctx.FullPath() == "/payment/client/:email" {
			email := ctx.Param("email")

			d, err := p.GetPaysByEmail(email)
			if err != nil {
				logrus.WithError(err).Info(err)
			}

			data, _ := json.Marshal(d)
			ctx.Writer.Write(data)
		}
	})
}

//POST create payment
func postCreatePayment(r *gin.Engine, p repository.Paymenter) {
	r.POST("/payment/", func(ctx *gin.Context) {
		tmp, err := getQuery(ctx, "id")
		if err != nil {
		}

		id := tmp.(int)

		tmp, err = getQuery(ctx, "email")
		if err != nil {
		}

		email := tmp.(string)

		tmp, err = getQuery(ctx, "amount")
		if err != nil {
		}

		amount := tmp.(int)

		tmp, err = getQuery(ctx, "currency")
		if err != nil {
		}

		currency := tmp.(rune)

		p.CreatePayment(model.NewPayment(id, email, amount, currency))
	})
}

// GET change status by id
func getChangeStatusByID(r *gin.Engine, p repository.Paymenter) {
	path := "/status/:id"
	r.POST(path, func(ctx *gin.Context) {
		if ctx.FullPath() == path {
			tmp, err := getQuery(ctx, "status")
			if err != nil {
			}

			s := tmp.(model.Status)

			param := ctx.Params

			id, err := strconv.Atoi(param.ByName("id"))
			if err != nil || id < 0 {
				logrus.WithError(err).Info(err)
			}
			p.ChangeStatusByID(id, s)
		}
	})
}

func Handle(r *gin.Engine, p repository.Paymenter) *gin.Engine {
	handls := [...]func(*gin.Engine, repository.Paymenter){
		postCancelByID, postCreatePayment,
		getStatusByID, getPaymentsByClientID, getPaymentsByClientEmail,
		getALLPayments,
	}

	for _, h := range handls {
		h(r, p)
	}
	return r
}
