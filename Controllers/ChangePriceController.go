package Controllers

import (
	"FarshidTemp/Services/ChanePriceService"
	"FarshidTemp/ViewModels"
	"github.com/gin-gonic/gin"
)

func ChangePrice(c *gin.Context) {
	//Get data of Req Body
	var changePrice []ViewModels.ChangePriceViewModel
	var rr []ViewModels.ChangePriceViewModel

	c.Bind(&changePrice)

	//Create Rule
	for _, item := range changePrice {
		rslt := ChanePriceService.FindBestPrice(&item)
		rr = append(rr, *rslt)
	}

	//Return it
	c.JSON(200, rr)

}
