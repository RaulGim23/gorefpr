package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	//"os"
	//"os/signal"
	//"syscall"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/oklog/run"
	"github.com/sirupsen/logrus"

	"files/repository/mysqlconnect"
	"files/service/file"
	"files/service/router"
	"files/transport/handler"

	_ "github.com/go-sql-driver/mysql"
)

const (
	timeout = 20 * time.Second
)

func main() {
	var httpServerAddress string

	flag.StringVar(&httpServerAddress, "addr", ":3380", "The http listen address")
	flag.Parse()

	db, err := mysqlconnect.Database("root:@tcp(127.0.0.1:3306)/asd")
	if err != nil {
		log.Printf("database %s", err)

		return
	}
	defer db.Close()

	fileSvc := file.New(mysqlconnect.File(db))

	r := mux.NewRouter()
	logging := logrus.New()
	logging.SetLevel(logrus.DebugLevel)

	// Internal API
	api := r.PathPrefix("/v1").Subrouter()

	muxRouter := router.NewMuxRouter(api, logging)
	handler.NewFile(muxRouter, fileSvc, logging)

	httpd := &http.Server{
		Addr:           httpServerAddress,
		ReadTimeout:    timeout,
		WriteTimeout:   timeout,
		MaxHeaderBytes: 1 << 20,
		Handler:        r,
	}

	var g run.Group

	//ctx, cancel := context.WithCancel(context.Background())
	g.Add(
		func() error {
			return httpd.ListenAndServe()
		},
		func(error) {
			_ = httpd.Shutdown(context.Background())
		},
	)
	/*
		g.Add(
			func() error {
				c := make(chan os.Signal, 1)
				signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
				 <-c
				 cancel()
				 return nil

			},
			func(error) {},
		)
	*/
	_ = g.Run()
}
