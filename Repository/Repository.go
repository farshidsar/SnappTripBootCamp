package Repository

import (
	"FarshidTemp/Initializer/database"
	"FarshidTemp/model"
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v9"
	"strings"
)

var ctx = context.Background()

func AddKeyRedis(key string, value string) bool {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := rdb.Set(ctx, key, value, 0).Err()
	if err != nil {
		//panic(err)
		return false
	}
	return true
}

func GetKeyRedis(key string) string {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	val, err := rdb.Get(ctx, key).Result()
	if err != nil && err == redis.Nil {
		return ""
	} else if err != nil {
		panic(err)
	}
	return val
}

func GetCity(cityCode string) *model.City {

	var city model.City
	item := GetKeyRedis("city_" + cityCode)
	if item != "" {

		json.Unmarshal([]byte(item), &city)
		return &city
	} else {
		database.DBCon.First(&city, "name=?", cityCode)
		b, _ := json.Marshal(city)
		AddKeyRedis("city_"+cityCode, string(b))
		return &city
	}
}

func GetAirplane(airplaneCode string) *model.Airline {

	var airline model.Airline
	item := GetKeyRedis("air_" + airplaneCode)
	if item != "" {

		json.Unmarshal([]byte(item), &airline)
		return &airline
	} else {
		database.DBCon.First(&airline, "code=?", strings.ToUpper(item))
		b, _ := json.Marshal(airline)
		AddKeyRedis("air_"+airplaneCode, string(b))
		return &airline
	}
}

func GetAgency(agencyName string) *model.Agency {

	var agency model.Agency
	item := GetKeyRedis("ag_" + agencyName)
	if item != "" {

		json.Unmarshal([]byte(item), &agency)
		return &agency
	} else {
		database.DBCon.First(&agency, "name=?", item)
		b, _ := json.Marshal(agency)
		AddKeyRedis("ag_"+agencyName, string(b))
		return &agency
	}
}

func GetSupplier(supplierName string) *model.Supplier {

	var supplier model.Supplier
	item := GetKeyRedis("sp_" + supplierName)
	if item != "" {

		json.Unmarshal([]byte(item), &supplier)
		return &supplier
	} else {
		database.DBCon.First(&supplier, "name=?", item)
		b, _ := json.Marshal(supplier)
		AddKeyRedis("sp_"+supplierName, string(b))
		return &supplier
	}
}
