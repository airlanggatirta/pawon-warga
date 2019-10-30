package model

type ZakatMaal struct {
	// gorm.Model
	Savings      int64
	Investations int64
	FixedAssets  int64
	Debts        int64
	Others       int64
}
