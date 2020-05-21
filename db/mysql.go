package db

import (
	"context"
	"fmt"

	Log "github.com/jampajeen/go-cqrs-taxi/logger"
	"github.com/jampajeen/go-cqrs-taxi/schema"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

type MysqlRepository struct {
	db *gorm.DB
}

func NewMysql(DBUser string, DBPassword string, DBHost string, DBPort int, DBName string) (*MysqlRepository, error) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", DBUser, DBPassword, DBHost, DBPort, DBName)
	Log.Info("connectionString: %s", connectionString)
	db, err = gorm.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate( // TODO: don't use in production
		&schema.Booking{},
		&schema.User{},
		&schema.Payment{},
		&schema.Taxi{},
		&schema.Driver{},
		&schema.TaxiDropOffEvent{},
		&schema.TaxiPickUpEvent{},
		&schema.UserAcceptProposalEvent{},
		&schema.TaxiCancelProposalEvent{},
		&schema.TaxiProposeEvent{},
		&schema.UserCancelBookingEvent{},
		&schema.UserCancelRequestBookingEvent{},
		&schema.UserRequestBookingEvent{},
	)
	return &MysqlRepository{
		db,
	}, nil
}

func (r *MysqlRepository) Close() {

}

func (r *MysqlRepository) Insert(ctx context.Context, item interface{}) error {
	d := r.db.Create(item)
	return d.Error
}
func (r *MysqlRepository) Update(ctx context.Context, item interface{}) error {
	d := r.db.Save(item)
	return d.Error
}
func (r *MysqlRepository) Delete(ctx context.Context, id string, item interface{}) error {
	d := r.db.Where("ID = ?", id).Delete(item)
	return d.Error
}
func (r *MysqlRepository) Find(ctx context.Context, id string, out interface{}) error {
	d := r.db.Where("ID = ?", id).First(out)
	return d.Error
}
func (r *MysqlRepository) FindAll(ctx context.Context, out interface{}) error {
	d := r.db.Find(out)
	return d.Error
}
