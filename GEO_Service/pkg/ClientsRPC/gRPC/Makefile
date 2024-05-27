PHONY: generate
generate:
		mkdir -p pkg/GEOgRPC
		protoc --go_out=pkg/GEOgRPC --go_opt=paths=source_relative \
				--go-grpc_out=pkg/GEOgRPC --go-grpc_opt=paths=source_relative \
				proto/geo.proto