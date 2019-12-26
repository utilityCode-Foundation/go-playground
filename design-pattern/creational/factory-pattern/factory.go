package factory

import (
	"errors"
)

func GetPaymentMethod(m string) (PaymentMethod, error) {
	switch m {
	case CASH:
		return new(CashPM), nil
	case DEBITCARD:
		return new(DebitCardPM), nil
	default:
		return nil, errors.New("Wrong type:"+m)
	}
}
