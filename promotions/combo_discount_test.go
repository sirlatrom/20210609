package promotions_test

import (
	"testing"

	"github.com/sirlatrom/20210609/catalogs"
	"github.com/sirlatrom/20210609/model"
	"github.com/sirlatrom/20210609/promotions"
)

func TestComboDiscount(t *testing.T) {
	acComboPromo := promotions.ComboDiscount("A", "C", 60.0)
	cart := model.Cart{ // A+B+C+D = (A+C = 60) + 30 + 15 = 105
		*catalogs.StaticNaiveCatalog.SKUByID("A"),
		*catalogs.StaticNaiveCatalog.SKUByID("B"),
		*catalogs.StaticNaiveCatalog.SKUByID("C"),
		*catalogs.StaticNaiveCatalog.SKUByID("D"),
	}
	totalPrice := 0.0
	result, remainder := acComboPromo(cart)
	for _, orderLine := range result {
		totalPrice += orderLine.Price
	}
	for _, sku := range remainder {
		totalPrice += sku.UnitPrice
	}
	if totalPrice != 105 {
		t.Fatalf("Expected total price 105 != %v", totalPrice)
	}
}
