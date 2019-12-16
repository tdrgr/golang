package procedural

// CreditCard blah
type CreditCard struct {
	OwnerName       string
	CardNumber      string
	ExpirationMonth int
	ExpirationYear  int
	SecurityCode    int
	AvailableCredit float32
}

// PayWithCredit blah
func PayWithCredit(card *CreditCard, amount float32) bool {
	if card.AvailableCredit > amount {
		card.AvailableCredit -= amount
		return true
	}
	return false
}