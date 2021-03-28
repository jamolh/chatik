package config

import (
	"encoding/json"
	"io/ioutil"
	"log"

	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

// cfg is instance of config.
var (
	cfg Config
)


//FromFile parse config from config file
func FromFile(path string) *Config {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("config, Reading config failed:", err)
	}

	if err := json.Unmarshal(b, &cfg); err != nil {
		log.Fatal("config, Parcing config failed:", err)
	}

	// app logging to file
	log.SetOutput(&lumberjack.Logger{
		Filename:   "logs/app.log",
		MaxSize:    2, // megabytes
		MaxBackups: 30,
		MaxAge:     40,   //days
		Compress:   true, // disabled by default
	})

	log.Println("-------- * ------- Logging -------- * -------")
	return &cfg
}

// Peek provides secure access to config options.
func Peek() *Config {
	return &cfg
}

// Config holds all config info.
type Config struct {
	Server   server   `json:"server"`   // Service holds service info
	Database database `json:"database"` // Database contains a dataaccess info
	Services services `json:"services"`
}

type server struct {
	Name   string `json:"name"`
	Addr   string `json:"addr"`
	APIKey string `json:"API-Key"`
	Path   string `json:"path"`
}

type services struct {
	AuthService service `json:"auth_service"`
}

type service struct {
	Addr string `json:"addr"`
	Name string `json:"name"`
	Path string `json:"path"`
}

// Database holds dataaccess info.
type database struct {
	Addr         string `json:"addr"`
	DatabaseName string `json:"dbname"`
	Password     string `json:"pass"`
	User         string `json:"user"`
}
