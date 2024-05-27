package main

import (
	"GeoServiseAppDate/ServersRPC/gRPC/internal/service"
	proto "GeoServiseAppDate/ServersRPC/gRPC/pkg/GEOgRPC/api/GeoService"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Error while listening on port: %v", err)
	}

	server := grpc.NewServer()
	reflection.Register(server)
	proto.RegisterGeoServiceServer(server, &service.Geocoder{})

	log.Printf("server listening at: %v", ":50051")

	if err = server.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
