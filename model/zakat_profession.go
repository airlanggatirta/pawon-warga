package model

type ZakatProfession struct {
	// gorm.Model
	MonthlyIncome int64
	OtherIncome   int64
	Debt          int64
}
