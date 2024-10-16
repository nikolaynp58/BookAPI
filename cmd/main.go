package main

import (
	db "bookAPI/db"
	gen "bookAPI/gen/book"
	"bookAPI/gen/http/book/server"
	service "bookAPI/service"
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	goahttp "goa.design/goa/v3/http"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
		panic(err)
	}

	storage, err := db.NewStorage()
	if err != nil {
		panic(err)
	}

	service := service.NewBookService(storage)
	endpoints := gen.NewEndpoints(service)

	mux := goahttp.NewMuxer()
	dec := goahttp.RequestDecoder
	enc := goahttp.ResponseEncoder
	svr := server.New(endpoints, mux, dec, enc, nil, nil)

	server.Mount(mux, svr)
	httpsvr := &http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	if err := httpsvr.ListenAndServe(); err != nil {
		fmt.Println("Failed to start server", err)
		panic(err)
	}
}
