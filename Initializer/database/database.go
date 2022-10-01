package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var (
	DBCon *gorm.DB
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "sa123"
	dbname   = "SnappTripBootCamp"
)

func InitDB() {
	var err error
	//psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	//	"password=%s dbname=%s sslmode=disable",
	//	host, port, user, password, dbname)

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	//connect to postgres database
	DBCon, err = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})

	//where myhost is port is the port postgres is running on
	//user is your postgres use name
	//password is your postgres password
	if err != nil {
		log.Fatal("Faild to connect DB")
		panic(err)

		panic("failed to connect database")
	}

}
