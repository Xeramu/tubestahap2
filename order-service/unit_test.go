package main

import "testing"

type MockValidator struct{}

func (m MockValidator) CheckUser(userID int, token string) bool {
	return true
}

func TestCreateOrder(t *testing.T) {
	mock := MockValidator{}

	req := Order{
		UserID:         1,
		NamaBarang:     "Laptop",
		Berat:          2,
		Dimensi:        "10x10",
		Jenis:          "Elektronik",
		AlamatPengirim: "Bandung",
		AlamatPenerima: "Jakarta",
	}

	o, err := CreateOrder(req, "dummy", mock)

	if err != nil {
		t.Fatal(err)
	}

	if o.Status != "created" {
		t.Fail()
	}
}