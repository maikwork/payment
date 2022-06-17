package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/maikwork/restPayments/internal/database"
	"github.com/maikwork/restPayments/internal/helper"
	"github.com/maikwork/restPayments/internal/repository"
	"github.com/maikwork/restPayments/internal/server"
	"github.com/sirupsen/logrus"
)

func listenSignal() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT)

	select {
	case <-c:
		fmt.Println("Exit to use ctrl + c")
		os.Exit(1)
	}
}

func main() {
	path := flag.String("c", "./cmd/config.yaml", "Path to config")
	flag.Parse()

	cnf, err := helper.ReadConfig(*path)
	if err != nil {
		logrus.WithError(err).Fatal("Don't work the function ReadConfig")
	}

	logrus.Info("config: ", cnf)

	db, err := database.Connection(cnf.DB)
	if err != nil {
		logrus.WithError(err).Fatal("Don't work the function ReadConfig")
	}
	defer db.Close()

	s := server.Server{
		Config: *cnf,
		PG:     repository.New(db),
	}

	// go listenSignal()
	s.Run()
}
