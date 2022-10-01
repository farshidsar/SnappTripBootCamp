package model

import "time"

type Airline struct {
	Code     string `json:"code" gorm:"primaryKey"`
	Name     string
	FullName string
}

type City struct {
	ID       int `json:"id"`
	IsActive bool
	Name     string `gorm:"primaryKey"`
	FullName string
}

type Supplier struct {
	ID       int `json:"id"`
	IsActive bool
	Name     string `gorm:"primaryKey"`
	FullName string
}

type Route struct {
	ID       int    `json:"id" gorm:"primaryKey;autoIncrement"`
	FromId   string `json:"from_id"`
	ToId     string `json:"to_id"`
	FromCity City   `gorm:"foreignKey:from_id;References:Name"`
	ToCity   City   `gorm:"foreignKey:to_id;References:Name"`
}

type Agency struct {
	ID        int `json:"id"`
	IsActive  bool
	Name      string `gorm:"primaryKey"`
	CreatedOn time.Time
}

type ChangePrice struct {
	RuleId       int    `json:"ruleId"`
	Origin       string `json:"origin"`
	Destination  string `json:"destination"`
	Airline      string `json:"airline"`
	Agency       string `json:"agency"`
	Supplier     string `json:"supplier"`
	BasePrice    int    `json:"basePrice"`
	Markup       int    `json:"markup"`
	PayablePrice int    `json:"payablePrice"`
}

type Rule struct {
	ID int `json:"id" gorm:"primaryKey;autoIncrement"`
	//Fk_routeId    int        `json:"fk_routeId"`
	//Fk_airLineId  string     `json:"fk_air_line_id"`
	//Fk_agencyId   int        `json:"fk_agency_id"`
	//Fk_supplierId string     `json:"fk_supplier_id"`
	Routes      []Route    `json:"Routes" gorm:"many2many:rule_routes;foreignKey:ID;joinForeignKey:RuleID;References:ID;joinReferences:RouteID"`
	Airlines    []Airline  `json:"Airlines" gorm:"many2many:rule_airlines;foreignKey:ID;joinForeignKey:RuleID;References:Code;joinReferences:AirelineCode"`
	Agencies    []Agency   `json:"Agencies" gorm:"many2many:rule_agencies;foreignKey:ID;joinForeignKey:RuleID;References:Name;joinReferences:AgencyID"`
	Suppliers   []Supplier `json:"Suppliers" gorm:"many2many:rule_suppliers;foreignKey:ID;joinForeignKey:RuleID;References:Name;joinReferences:SupplierID"`
	AmountType  string
	AmountValue float64
}
