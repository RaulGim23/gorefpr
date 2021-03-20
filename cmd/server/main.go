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
)

func main() {
	//var httpServerAddress string
	//
	//flag.StringVar(&httpServerAddress,"addr",":3380","The http listen address")
	//flag.Parse()
	//
	//db, err := sqlite.Database("database.sqlite3")
	//if err != nil {
	//	log.Printf("database %s",err)
	//	return
	//}
	//defer db.Close()
	//
	//usrSvc := user.New(sqlite.User(db), []byte{})
	//
	//r := mux.NewRouter()
	//log := logrus.New()
	//log.SetLevel(logrus.DebugLevel)
	//
	//// Internal API
	//api := r.PathPrefix("/v1").Subrouter()
	//// Middleware goes here
	//
	////
	//muxRouter := router.NewMuxRouter(api, log)
	//handler.NewUser(muxRouter, usrSvc, log)
	//
	//httpd := &http.Server{
	//	Addr: httpServerAddress,
	//	ReadTimeout:    10 * time.Second,
	//	WriteTimeout:   10 * time.Second,
	//	MaxHeaderBytes: 1 << 20,
	//	Handler:        r,
	//}
	//
	//var g run.Group
	//
	//ctx, _ := context.WithCancel(context.Background())
	//g.Add(
	//	func() error {
	//		return httpd.ListenAndServe()
	//	},
	//	func(error) {
	//		httpd.Shutdown(ctx)
	//	},
	//)
	///*
	//	g.Add(
	//		func() error {
	//			c := make(chan os.Signal, 1)
	//			signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	//			 <-c
	//			 cancel()
	//			 return nil
	//
	//		},
	//		func(error) {},
	//	)
	//*/
	//g.Run()
}