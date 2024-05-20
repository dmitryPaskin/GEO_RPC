package main

import (
	"GeoServiseAppDate/RPC/internal/service"
	"log"
	"net"
	"net/rpc"
)

func main() {
	geoService := new(service.Geocoder)
	if err := rpc.Register(geoService); err != nil {
		log.Fatal(err)
	}

	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Сервер запущен, ожидание вызовов... На порту: 1234.")

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go rpc.ServeConn(conn)
	}

}
