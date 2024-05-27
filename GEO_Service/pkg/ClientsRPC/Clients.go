package ClientsRPC

import (
	"GeoServiseAppDate/internal/service"
	"GeoServiseAppDate/pkg/ClientsRPC/gRPC"
	"fmt"
	"os"
)

type FactoryRPC interface {
	service.Service
}

func GetRPCProtocol() (FactoryRPC, error) {
	protocol := os.Getenv("RPC_PROTOCOL")

	switch protocol {
	case "ServersRPC":
		return &RPC{}, nil
	case "JSON-ServersRPC":
		return &JSONRPC{}, nil
	case "gRPC":
		return &gRPC.GRPC{}, nil
	}
	return nil, fmt.Errorf("%s", "protocol not found")
}
