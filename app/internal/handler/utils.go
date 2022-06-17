package handler

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/maikwork/restPayments/internal/model"
	"github.com/sirupsen/logrus"
)

func getQuery(ctx *gin.Context, w string) (interface{}, error) {
	t, b := ctx.GetQuery(w)
	if !b {
		err := fmt.Errorf("Error: the number of parameters")
		logrus.WithError(err).Info("Where is id client")
		return nil, err
	}

	switch w {
	case "email":
		return t, nil
	case "status":
		tmp, err := strconv.Atoi(t)
		if err != nil {
			logrus.WithError(err).Info("Don't currect query parametr")
			return nil, err
		}

		status := model.Status(tmp)
		return status, nil
	default:
		res, err := strconv.Atoi(t)
		if err != nil {
			logrus.WithError(err).Info("Don't currect query parametr")
			return nil, err
		}
		return res, nil
	}
}
