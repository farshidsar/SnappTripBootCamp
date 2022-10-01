package migrations

import (
	"FarshidTemp/Initializer/database"
	"FarshidTemp/model"
)

func Migrate() {

	database.DBCon.AutoMigrate(model.Rule{}, model.Agency{}, model.Airline{}, model.Supplier{}, model.City{}, model.Route{})
	//database.DBCon.AutoMigrate(model.Route{}, model.City{})
	//database.DBCon.AutoMigrate(model.Rule{}, model.City{}, model.Route{})

	//Routes Fk
	//database.DBCon.Model(&model.Route{}).AddForeignKey("from_id", "cities(id)", "CASCADE", "RESTRICT")
	//database.DBCon.Model(&model.Route{}).AddForeignKey("to_id", "cities(id)", "CASCADE", "RESTRICT")

	//Routes Fk
	//database.DBCon.Model(&model.Rule{}).AddForeignKey("fk_route_id", "routes(id)", "CASCADE", "RESTRICT")
	//database.DBCon.Model(&model.Rule{}).AddForeignKey("fk_air_line_id", "airlines(code)", "CASCADE", "RESTRICT")
	//database.DBCon.Model(&model.Rule{}).AddForeignKey("fk_agency_id", "agencies(id)", "CASCADE", "RESTRICT")
	//database.DBCon.Model(&model.Rule{}).AddForeignKey("fk_supplier_id", "suppliers(id)", "CASCADE", "RESTRICT")

}
