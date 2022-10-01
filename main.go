package main

import (
	"FarshidTemp/Controllers"
	"FarshidTemp/Initializer"
	"FarshidTemp/Initializer/database"
	migrations "FarshidTemp/migration"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "sa123"
	dbname   = "SnappTripBootCamp"
)

//	func main() {
//		var a = 2
//		var bb = 3
//		var aaa = fmt.Sprintf("%s", a+bb)
//		bbb := fmt.Sprintf("%d", int(a*bb))
//		fmt.Println(reflect.TypeOf(bbb))
//		fmt.Println("aaa%s", a+bb)
//		fmt.Println(aaa)
//		fmt.Println(bbb)
//	}
func main() {
	//database.InitDB()
	migrations.Migrate()
	ginInitialize()
	//ConnectToDataBaseSample()
}

func init() {
	Initializer.LoadEnvVariables()
	database.InitDB()
}
func ginInitialize() {
	r := gin.Default()
	r.POST("/rule", Controllers.RuleCreate)
	r.POST("/c", Controllers.ChangePrice)
	r.Run() // listen and serve on 0.0.0.0:8080
}

func ConnectToDataBaseSample() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	fmt.Println(psqlInfo)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}
