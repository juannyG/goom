package product

import (
	"testing"
)

func TestProductSerialization(t *testing.T) {
	dw := DBWrapper{}
	db := dw.InitDB()
	defer db.Close()

	expected_product := Product{"TestP", true, true}
	got_product, _ := GetProduct(dw.DB, "TestM", "TestP")
	if expected_product != got_product {
		t.Errorf("Wanted: %+v\nGot: %+v\n", expected_product, got_product)
	}

	dw.tearDown()
}

func TestProductNotFound(t *testing.T) {
	dw := DBWrapper{}
	db := dw.InitDB()
	defer db.Close()

	_, err := GetProduct(dw.DB, "OtherM", "NotP")
	if err == nil {
		t.Errorf("Expected error but got Product")
	}

	dw.tearDown()
}

func TestProductNotFoundIfMerchantInactive(t *testing.T) {
	dw := DBWrapper{}
	db := dw.InitDB()
	defer db.Close()

	id := dw.insertMerchant("NewM", false)
	dw.insertProduct(id, "OtherP", true, true)
	_, err := GetProduct(dw.DB, "NewM", "OtherP")
	if err == nil {
		t.Errorf("Expected error but got Product")
	}

	dw.tearDown()
}
