package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jamolh/chatik/config"
	"github.com/jamolh/chatik/db"
	"github.com/jamolh/chatik/handlers"
	"github.com/julienschmidt/httprouter"

	_ "github.com/joho/godotenv/autoload" // to load .env
)

var (
	router = httprouter.New()
)

func main() {
	var (
		configPath = flag.String("config", "./config.json", "path of the config file")
		conf       = config.FromFile(*configPath)
	)
	fmt.Println("here we come", conf.Server)

	initRoutes()
	srv := &http.Server{
		Addr:         conf.Server.Addr, //port,
		ReadTimeout:  40 * time.Second,
		WriteTimeout: 60 * time.Second,
		Handler:      router,
	}

	//time.Sleep(time.Minute) // wait for psql docker container will be started
	db.Connect()
	defer db.Close()
	// TODO: do it better!
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		sigint := make(chan os.Signal)
		signal.Notify(sigint, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGINT)
		s := <-sigint
		log.Println("server received signal", s)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			log.Println("server: couldn't shutdown because of " + err.Error())
		}
	}()

	log.Println("http server is running on port", conf.Server.Addr)
	log.Fatal("Closing Server error:", srv.ListenAndServe())

}

// declaring our routes
func initRoutes() {
	path := config.Peek().Server.Path // os.Getenv("PATH")
	fmt.Println("Path:", path)
	router.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Access-Control-Request-Method") != "" {
			// Set CORS headers
			w.Header().Set("Accept", "*/*")
			w.Header().Set("Access-Control-Allow-Methods", "*")
			w.Header().Set("Access-Control-Allow-Headers", "*") // w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, SourceAuthorization, RefreshToken")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
			w.Header().Set("Content-Type", "application/json, text/html")

			return
		}
	})

	router.GET("/", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		fmt.Println("got empty request")
		w.Write([]byte("i'm ok"))
	})
	router.POST("/v1/user/token", handlers.Register)
}
