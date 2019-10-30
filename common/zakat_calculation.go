package common

func RoundZakatAmount(zakatAmount int64) int64 {
	var finalAmount int64

	lastThreeDigits := zakatAmount % 1000

	if lastThreeDigits != 0 {
		finalAmount = zakatAmount + 1000 - lastThreeDigits
	} else {
		finalAmount = zakatAmount
	}

	return finalAmount
}
