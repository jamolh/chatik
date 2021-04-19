package main

import (
	"log"
	"net"
	"os"

	"github.com/jamolh/chatik/db"
	"github.com/jamolh/chatik/grpcmethods"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/grpc"
	"gopkg.in/natefinch/lumberjack.v2"

	_ "github.com/joho/godotenv/autoload" // to load .env
)

var (
	router = httprouter.New()
)

func init() {

	// app logging to file
	log.SetOutput(&lumberjack.Logger{
		Filename:   "logs/app.log",
		MaxSize:    2, // megabytes
		MaxBackups: 30,
		MaxAge:     40,   //days
		Compress:   true, // disabled by default
	})

	log.Println("-------- * ------- Logging -------- * -------")
}

func main() {
	port, found := os.LookupEnv("SERVER_ADDR")
	if !found {
		log.Fatal("chatik run failed:SERVER_ADDR not set")
	}
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("Failed to listen:", err)
	}
	//time.Sleep(time.Minute) // wait for psql docker container will be started
	db.Connect()
	defer db.Close()
	// TODO: do it better!
	// ctx, cancel := context.WithCancel(context.Background())
	// go func() {
	// 	sigint := make(chan os.Signal)
	// 	signal.Notify(sigint, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGINT)
	// 	s := <-sigint
	// 	log.Println("server received signal", s)
	// 	defer cancel()
	// 	if err := srv.Shutdown(ctx); err != nil {
	// 		log.Println("server: couldn't shutdown because of " + err.Error())
	// 	}
	// }()
	s := grpc.NewServer()
	grpcmethods.RegisterService(s)
	log.Println("http server is running on port:", port)
	if err = s.Serve(lis); err != nil {
		log.Fatal("Failed to server", err)
	}
}
