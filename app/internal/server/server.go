package server

import (
	"github.com/gin-gonic/gin"
	"github.com/maikwork/restPayments/internal/handler"
	"github.com/maikwork/restPayments/internal/model"
	"github.com/maikwork/restPayments/internal/repository"
	"github.com/sirupsen/logrus"
)

type Server struct {
	Config model.Config
	PG     repository.Paymenter
}

func (s *Server) Run() {
	r := gin.Default()

	handler.Handle(r, s.PG)

	if err := r.Run(); err != nil {
		logrus.WithError(err).Fatal(err)
	}
}
