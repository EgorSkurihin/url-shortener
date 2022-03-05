package main

import (
	"flag"
	"log"

	"github.com/EgorSkurihin/url-shortener/server"
)

func init() {
	flag.Parse()

}

func main() {
	addr := flag.String("Addr", ":8181", "Address of HTTP Server")
	dsn := flag.String("DSN", "root:232323@tcp(localhost:3306)/url_shortener", "Name of data sourse MySQL")
	srv := server.New(*dsn)
	log.Printf("Server started at %s", *addr)
	log.Fatal(srv.Run(*addr))
}
