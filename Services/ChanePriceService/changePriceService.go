package ChanePriceService

import (
	"FarshidTemp/Initializer/database"
	"FarshidTemp/ViewModels"
)

func FindBestPrice(changePrice *ViewModels.ChangePriceViewModel) *ViewModels.ChangePriceViewModel {

	var bestRow ViewModels.ChangePriceQueryResult

	sql := "select  r2.id, r2.amount_type,r2.amount_value,\n        case when r2.amount_type='FIXED' then r2.amount_value + ?\n            else r2.amount_value*?\n                end\n            as Payableprice\n        , ((case when r2.amount_type='FIXED' then r2.amount_value + ? else r2.amount_value * ? end) - r2.amount_value) as markup\nfrom routes r\nright outer join rule_routes rr on r.id = rr.route_id\njoin rule_agencies rag on rr.rule_id = rag.rule_id\njoin rule_airlines ra on rr.rule_id = ra.rule_id\njoin rule_suppliers rs on rr.rule_id = rs.rule_id\njoin rules r2 on ra.rule_id = r2.id\nwhere r.from_id=? and to_id=?\n  and ra.aireline_code=? and rag.agency_id=? and rs.supplier_id=?\norder by 4 desc\n"

	database.DBCon.Raw(sql, changePrice.BasePrice, changePrice.BasePrice, changePrice.BasePrice, changePrice.BasePrice, changePrice.Origin, changePrice.Destination, changePrice.Airline, changePrice.Agency, changePrice.Supplier).Scan(&bestRow)

	if bestRow.Id == "" {
		changePrice.RuleId = ""
		changePrice.PayablePrice = changePrice.BasePrice
		changePrice.Markup = 0
		return changePrice
	} else {
		changePrice.RuleId = bestRow.Id
		changePrice.PayablePrice = bestRow.Payableprice
		changePrice.Markup = bestRow.Payableprice - changePrice.BasePrice
		return changePrice
	}

}
