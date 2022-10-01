package ViewModels

type ChangePriceQueryResult struct {
	Markup       int    `json:"markup" db:"markup" column:"markup" gorm:"column:markup"`
	Payableprice int    `json:"payableprice" db:"payableprice" column:"payableprice" gorm:"column:payableprice"`
	Amount_value int    `json:"amount_value" db:"amount_value" column:"amount_value" gorm:"column:amount_value"`
	Amount_type  string `json:"amount_type" db:"amount_type" column:"amount_type" gorm:"column:amount_type"`
	Id           string `json:"id" db:"id" column:"id" gorm:"column:id"`
}
