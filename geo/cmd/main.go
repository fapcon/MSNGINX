package main

import (
	"geo/internal/grpc/geo"
	geopr "github.com/fapcon/MSHUGOprotos/protos/geo/gen"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {

	listen, err := net.Listen("tcp", ":44973")
	if err != nil {
		log.Fatalf("Ошибка при прослушивании порта: %v", err)
	}

	server := grpc.NewServer()
	geopr.RegisterGeoServiceServer(server, &geo.ServerGeo{})

	log.Println("Запуск gRPC сервера geo...")
	if err := server.Serve(listen); err != nil {
		log.Fatalf("Ошибка при запуске сервера: %v", err)
	}
}
