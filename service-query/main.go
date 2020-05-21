package main

import (
	"fmt"
	"os"
	"time"

	"github.com/jampajeen/go-cqrs-taxi/db"
	"github.com/jampajeen/go-cqrs-taxi/event"
	Log "github.com/jampajeen/go-cqrs-taxi/logger"
	"github.com/jampajeen/go-cqrs-taxi/search"
	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Config struct {
	MysqlHost            string `envconfig:"MYSQL_HOST"`
	MysqlPort            int    `envconfig:"MYSQL_PORT"`
	MysqlUser            string `envconfig:"MYSQL_USER"`
	MysqlPassword        string `envconfig:"MYSQL_PASSWORD"`
	MysqlDB              string `envconfig:"MYSQL_DATABASE"`
	NatsAddress          string `envconfig:"NATS_ADDRESS"`
	ElasticsearchAddress string `envconfig:"ELASTICSEARCH_ADDRESS"`
}

type RetryFunc func(int) error

func ForeverRetry(d time.Duration, f RetryFunc) {
	for i := 0; ; i++ {
		err := f(i)
		if err == nil {
			return
		}
		time.Sleep(d)
	}
}

func main() {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		Log.Fatal(err)
	}

	repo, err := db.NewMysql(cfg.MysqlUser, cfg.MysqlPassword, cfg.MysqlHost, cfg.MysqlPort, cfg.MysqlDB)
	if err != nil {
		Log.Error(err)
		os.Exit(1)
	}
	db.SetRepository(repo)

	// Connect to ElasticSearch
	ForeverRetry(2*time.Second, func(_ int) error {
		es, err := search.NewElastic(fmt.Sprintf("http://%s", cfg.ElasticsearchAddress))
		if err != nil {
			Log.Error(err)
			return err
		}
		search.SetRepository(es)
		return nil
	})
	defer search.Close()

	// Connect to Nats
	es, err := event.NewNats(fmt.Sprintf("nats://%s", cfg.NatsAddress))
	if err != nil {
		Log.Fatal(err)
	}
	handleEvents(es)
	event.SetEventStore(es)
	defer event.Close()

	// Import all taxi location on startup
	importTaxies()

	e := echo.New()
	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte("secret"),
		Skipper: func(c echo.Context) bool {
			if c.Path() == "/query/Test" {
				return true
			}
			return false
		},
	}))

	e.GET("/query/Bookings/:id", bookingsHandler)
	e.GET("/query/Payments/:id", paymentsHandler)
	e.GET("/query/Taxies/:id", taxiesHandler)
	e.GET("/query/Drivers/:id", driversHandler)
	e.GET("/query/Users/:id", usersHandler)

	e.GET("/query/SearchTaxies", searchTaxiesHandler)
	e.POST("/query/Test", testHandler)

	port := fmt.Sprintf(":%d", 8082)
	e.Logger.Fatal(e.Start(port))
}
