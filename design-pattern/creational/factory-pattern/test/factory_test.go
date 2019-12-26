package test

import (
	factory "go-playground/design-pattern/creational/factory-pattern"
	"strings"
	"testing"
)

func TestCreatePaymentMethodCash(t *testing.T) {
	payment, err := factory.GetPaymentMethod(factory.CASH)
	if err != nil {
		t.Fatal("A payment method of type 'Cash' must exist")
	}
	msg := payment.Pay(10.30)
	if !strings.Contains(msg, "paid using cash") {
		t.Error("The cash payment method message wasn't correct")
	}
	t.Log("LOG:", msg)
}

func TestGetPaymentMethodDebitCard(t *testing.T) {
	payment, err := factory.GetPaymentMethod(factory.DEBITCARD)
	if err != nil {
		t.Error("A payment method of type 'DebitCard' must exist")
	}
	msg := payment.Pay(22.30)
	if !strings.Contains(msg, "paid using debit card") {
		t.Error("The debit card payment method message wasn't correct")
	}
	t.Log("LOG:", msg)
}

func TestGetPaymentMethodNonExistent(t *testing.T) {
	_, err := factory.GetPaymentMethod("abc")
	if err == nil {
		t.Error("A payment method abc must return error")
	}
	t.Log("LOG:", err)
}
