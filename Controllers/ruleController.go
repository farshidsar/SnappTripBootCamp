package Controllers

import (
	"FarshidTemp/Initializer/database"
	"FarshidTemp/Services/RuleService"

	"FarshidTemp/ViewModels"
	"FarshidTemp/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func RuleCreate(c *gin.Context) {
	//Get data of Req Body
	var rule ViewModels.RuleViewModel

	c.Bind(&rule)

	//Create Rule

	result, errMSG := RuleService.CreateRule(&rule)

	//Return it

	if result {
		c.JSON(200, gin.H{
			"status":  "SUCCESS",
			"message": nil,
		})
	} else {
		c.JSON(500, gin.H{
			"status":  "FAILED",
			"message": errMSG,
		})
	}

	var rr model.Rule
	database.DBCon.Preload(clause.Associations).Preload("Routes." + clause.Associations).First(&rr)

	var route model.Route
	database.DBCon.Preload(clause.Associations).First(&route)

}
