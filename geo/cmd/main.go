package main

import (
	cnt "geo/internal/controller"
	"geo/internal/grpc/geo"
	"geo/internal/router"
	"geo/internal/service"
	geopr "github.com/fapcon/MSHUGOprotos/protos/geo/gen"
	"github.com/go-chi/chi"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"sync"
)

func main() {
	geoservice := service.NewGeoService()
	geohandle := cnt.NewHandleGeo(geoservice)
	r := router.Route(geohandle)

	w := sync.WaitGroup{}
	w.Add(2)

	go func(r *chi.Mux) {
		defer w.Done()
		http.ListenAndServe(":8081", r)
	}(r)

	go func() {
		defer w.Done()
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
	}()
	w.Wait()
}
