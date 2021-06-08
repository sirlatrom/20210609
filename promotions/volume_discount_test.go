package promotions_test

import (
	"testing"

	"github.com/sirlatrom/20210609/catalogs"
	"github.com/sirlatrom/20210609/model"
	"github.com/sirlatrom/20210609/promotions"
)

func TestVolumeDiscount(t *testing.T) {
	acComboPromo := promotions.VolumeDiscount("D", 3, 40.0)
	cart := model.Cart{ // 4*A + 4*B + 4*C + 4*D = 4*50 + 4*30 + 4*20 + (3*D=40) + 15 =
		*catalogs.StaticNaiveCatalog.SKUByID("A"),
		*catalogs.StaticNaiveCatalog.SKUByID("B"),
		*catalogs.StaticNaiveCatalog.SKUByID("C"),
		*catalogs.StaticNaiveCatalog.SKUByID("D"),
		*catalogs.StaticNaiveCatalog.SKUByID("A"),
		*catalogs.StaticNaiveCatalog.SKUByID("B"),
		*catalogs.StaticNaiveCatalog.SKUByID("C"),
		*catalogs.StaticNaiveCatalog.SKUByID("D"),
		*catalogs.StaticNaiveCatalog.SKUByID("A"),
		*catalogs.StaticNaiveCatalog.SKUByID("B"),
		*catalogs.StaticNaiveCatalog.SKUByID("C"),
		*catalogs.StaticNaiveCatalog.SKUByID("D"),
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
	if totalPrice != 455 {
		t.Fatalf("Expected total price 455 != %v", totalPrice)
	}
}
