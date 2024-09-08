package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/vaidik-bajpai/green-screen/pkg/discovery"
	"github.com/vaidik-bajpai/green-screen/pkg/discovery/consul"
	"github.com/vaidik-bajpai/green-screen/rating/internal/controller/rating"
	httphandler "github.com/vaidik-bajpai/green-screen/rating/internal/handler/http"
	"github.com/vaidik-bajpai/green-screen/rating/internal/repository/memory"
)

const serviceName = "rating"

func main() {
	var port int
	flag.IntVar(&port, "port", 8082, "port address for the rating service")
	flag.Parse()

	log.Println("Starting the rating service")
	registry, err := consul.NewRegistry("localhost:8500")
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	instanceID := discovery.GenerateInstanceID(serviceName)

	err = registry.Register(ctx, instanceID, serviceName, fmt.Sprintf("localhost:%d", port))
	if err != nil {
		panic(err)
	}

	go func() {
		for {
			err := registry.ReportHealthyState(instanceID, serviceName)
			if err != nil {
				log.Println("Failed to report healthy state: " + err.Error())
			}

			time.Sleep(1 * time.Second)
		}
	}()
	defer registry.Deregister(ctx, instanceID, serviceName)

	repo := memory.New()
	ctrl := rating.New(repo)
	hdl := httphandler.New(ctrl)
	http.Handle("/rating", http.HandlerFunc(hdl.Handle))
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		panic(err)
	}
}
