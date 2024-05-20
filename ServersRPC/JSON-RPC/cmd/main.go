package main

import (
	"GeoServiseAppDate/JSON-RPC/internal/service"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	geoService := new(service.Geocoder)
	if err := rpc.Register(geoService); err != nil {
		log.Fatal(err)
	}

	l, err := net.Listen("tcp", ":4321")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Сервер запущен, ожидание вызовов... На порту: 4321.")

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go jsonrpc.ServeConn(conn)
	}
}
