package main

import (
	//https://www.wsj.com
	//https://www.leatherheadsports.com
	//https://www.beardbrand.com
	//https://rebel8.com/collections/mens

	//"github.com/gorilla/csrf"
	//import "gopkg.in/hlandau/passlib.v1"
	//go get -u github.com/matcornic/hermes
	//"github.com/jasonlvhit/gocron"

	// "github.com/dimiro1/health"
	//https://medium.com/southbridge/nginx-as-reverse-proxy-a62815edd8c1
	//https://ru.wikipedia.org/wiki/Почтовый_адрес

	//"github.com/gorilla/mux"
	//https://github.com/Shopify/themekit/blob/9b04cd2c921985db4cc34ab33805c9aa21cc0b50/docs/_sass/marketing_assets/modules/_forms.scss

	"log"
	"os"
	"os/signal"
	"context"
	"time"
	"net/http"
	"store"
)

func main() {
	publicUrl := os.Getenv("PUBLIC_URL")


	wait := time.Second * 1

	s, shutdownStore, err := store.NewStore(&store.Config{
		ServerURL: publicUrl,
		HR: []string {
			"denisxy12@gmail.com",
			"kosmo_polit@rambler.ru",
		},
	})

	if err != nil {
		log.Fatal(err)
	}

	srv := &http.Server{
		Addr:         "0.0.0.0:8082",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      s, // Pass our instance of gorilla/mux in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	log.Printf("Launch App server")

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	shutdownStore(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	//if your application should wait for other services
	<-ctx.Done()
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}
