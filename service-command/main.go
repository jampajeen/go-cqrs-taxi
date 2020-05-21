package main

import (
	"fmt"
	"os"
	"time"

	"github.com/jampajeen/go-cqrs-taxi/db"
	"github.com/jampajeen/go-cqrs-taxi/event"
	Log "github.com/jampajeen/go-cqrs-taxi/logger"
	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gopkg.in/go-playground/validator.v9"
)

type Config struct {
	MysqlHost     string `envconfig:"MYSQL_HOST"`
	MysqlPort     int    `envconfig:"MYSQL_PORT"`
	MysqlUser     string `envconfig:"MYSQL_USER"`
	MysqlPassword string `envconfig:"MYSQL_PASSWORD"`
	MysqlDB       string `envconfig:"MYSQL_DATABASE"`
	NatsAddress   string `envconfig:"NATS_ADDRESS"`
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
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
		// log.Println(err)
		Log.Error(err)
		os.Exit(1)
	}
	db.SetRepository(repo)

	// Connect to Nats
	es, err := event.NewNats(fmt.Sprintf("nats://%s", cfg.NatsAddress))
	if err != nil {
		Log.Fatal(err)
	}
	event.SetEventStore(es)
	defer event.Close()

	e := echo.New()
	e.Use(middleware.JWT([]byte("secret")))
	e.Validator = &CustomValidator{validator: validator.New()}

	// API For Internal system
	e.POST("/command/Broadcast", broadcastHandler)
	e.POST("/command/InsertTaxi", insertTaxiHandler)
	e.PUT("/command/UpdateTaxi", updateTaxiHandler)
	e.PUT("/command/UpdateTaxiLocation", updateTaxiLocationHandler)

	// API For App
	e.POST("/command/TaxiCancelProposal", taxiCancelProposalHandler)
	e.POST("/command/TaxiDropOff", taxiDropOffHandler)
	e.POST("/command/TaxiPickUp", taxiPickUpHandler)
	e.POST("/command/TaxiPropose", taxiProposeHandler)
	e.POST("/command/UserAcceptProposal", userAcceptProposalHandler)
	e.POST("/command/UserCancelBooking", userCancelBookingHandler)
	e.POST("/command/UserCancelRequestBooking", userCancelRequestBookingHandler)
	e.POST("/command/UserRequestBooking", userRequestBookingHandler)

	port := fmt.Sprintf(":%d", 8081)
	e.Logger.Fatal(e.Start(port))
}
