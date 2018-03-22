package entity

type Payment struct {
	Beneficiary Beneficiary `json:"beneficiary_party"`
	Charge      Charge      `json:"charges_information"`
	Debtor      Debtor      `json:"debtor_party"`
	Fx          Fx          `json:"fx"`
	Sponsor     Sponsor     `json:"sponsor_party"`
	// Other information
}

type Beneficiary struct {
	// Beneficiary information
}

func Store(payment Payment) error {
	// Store on different databases
	// Depend if some information is already stored (like beneficiary), an upsert/fetch may be done.

	return
}

