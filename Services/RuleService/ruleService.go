package RuleService

import (
	"FarshidTemp/Initializer/database"
	"FarshidTemp/Repository"
	"FarshidTemp/ViewModels"
	"FarshidTemp/model"
	"strings"
)

func CreateRule(rule *ViewModels.RuleViewModel) (bool, string) {
	var rule3 model.Rule
	var errMsg string
	for _, item := range rule.Routes {

		if item.Origin == "" {
			item.Origin = " "
		}
		if item.Destination == "" {
			item.Destination = " "
		}

		fromCity := Repository.GetCity(item.Origin)
		toCity := Repository.GetCity(item.Destination)

		if item.Origin != "" && fromCity.Name == "" {
			errMsg = "Cannot find origin city in database."
			return false, errMsg
		}

		if item.Origin != "" && fromCity.Name == "" {
			errMsg = "Cannot find destination city in database."
			return false, errMsg
		}

		route := model.Route{
			//ID:       1,
			FromId:   item.Origin,
			ToId:     item.Destination,
			FromCity: *fromCity,
			ToCity:   *toCity,
		}
		rule3.Routes = append(rule3.Routes, route)

	}
	for _, item := range rule.Airlines {

		airline := Repository.GetAirplane(strings.ToUpper(item))

		if item != "" && airline.Name == "" {
			errMsg = "Cannot find airline in database."
			return false, errMsg
		}

		rule3.Airlines = append(rule3.Airlines, *airline)

	}
	for _, item := range rule.Agencies {

		agency := Repository.GetAgency(item)

		if item != "" && agency.Name == "" {
			errMsg = "Cannot find agency in database."
			return false, errMsg
		}

		rule3.Agencies = append(rule3.Agencies, *agency)

	}
	for _, item := range rule.Suppliers {

		var supplier = Repository.GetSupplier(item)

		if item != "" && supplier.Name == "" {
			errMsg = "Cannot find supplier in database."
			return false, errMsg
		}

		rule3.Suppliers = append(rule3.Suppliers, *supplier)

	}
	rule3.AmountType = rule.AmountType
	rule3.AmountValue = rule.AmountValue

	result := database.DBCon.Create(&rule3)
	if result.Error != nil {
		errMsg = "Somting happend in save DB"
		return false, errMsg
	}
	return true, ""
}
