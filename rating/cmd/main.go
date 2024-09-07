package main

import (
	"log"
	"net/http"

	"github.com/vaidik-bajpai/green-screen/rating/internal/controller/rating"
	httphandler "github.com/vaidik-bajpai/green-screen/rating/internal/handler/http"
	"github.com/vaidik-bajpai/green-screen/rating/internal/repository/memory"
)

func main() {
	log.Println("Starting the rating service")
	repo := memory.New()
	ctrl := rating.New(repo)
	hdl := httphandler.New(ctrl)
	http.Handle("/rating", http.HandlerFunc(hdl.Handle))
	if err := http.ListenAndServe(":8082", nil); err != nil {
		panic(err)
	}
}
