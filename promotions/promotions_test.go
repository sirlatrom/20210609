package promotions_test

import (
	"testing"

	"github.com/sirlatrom/20210609/catalogs"
	"github.com/sirlatrom/20210609/model"
	"github.com/sirlatrom/20210609/promotions"
)

var (
	ThreeAsPromo = promotions.VolumeDiscount("A", 3, 130.0)
	TwoBsPromo   = promotions.VolumeDiscount("B", 2, 45.0)
	CDComboPromo = promotions.ComboDiscount("C", "D", 30.0)
)

func scenarioA(catalog model.Catalog) model.Cart {
	return model.Cart{
		*catalog.SKUByID("A"),
		*catalog.SKUByID("B"),
		*catalog.SKUByID("C"),
	}
}

func scenarioB(catalog model.Catalog) model.Cart {
	return model.Cart{
		*catalog.SKUByID("A"),
		*catalog.SKUByID("A"),
		*catalog.SKUByID("B"),
		*catalog.SKUByID("A"),
		*catalog.SKUByID("B"),
		*catalog.SKUByID("A"),
		*catalog.SKUByID("B"),
		*catalog.SKUByID("A"),
		*catalog.SKUByID("B"),
		*catalog.SKUByID("B"),
		*catalog.SKUByID("C"),
	}
}

func scenarioC(catalog model.Catalog) model.Cart {
	return model.Cart{
		*catalog.SKUByID("A"),
		*catalog.SKUByID("A"),
		*catalog.SKUByID("D"),
		*catalog.SKUByID("A"),
		*catalog.SKUByID("B"),
		*catalog.SKUByID("B"),
		*catalog.SKUByID("C"),
		*catalog.SKUByID("B"),
		*catalog.SKUByID("B"),
		*catalog.SKUByID("B"),
	}
}

func TestPromos(t *testing.T) {
	catalog := catalogs.StaticNaiveCatalog
	allPromos := []promotions.Promotion{ThreeAsPromo, TwoBsPromo, CDComboPromo}
	testcases := []struct {
		description        string
		cart               model.Cart
		promos             []promotions.Promotion
		expectedTotalPrice float64
		expectedMatches    int
	}{
		{
			description:        "Scenario A with all promos",
			cart:               scenarioA(catalog),
			promos:             allPromos,
			expectedTotalPrice: 100,
			expectedMatches:    0,
		},
		{
			description:        "Scenario B with all promos",
			cart:               scenarioB(catalog),
			promos:             allPromos,
			expectedTotalPrice: 370,
			expectedMatches:    3, // once 3*A + twice 2*B
		},
		{
			description:        "Scenario C with all promos",
			cart:               scenarioC(catalog),
			promos:             allPromos,
			expectedTotalPrice: 280,
			expectedMatches:    4, // once 3*A + twice 2*B + once C+D
		},
		{
			description:        "Scenario B with 3*A=130",
			cart:               scenarioB(catalog),
			promos:             []promotions.Promotion{ThreeAsPromo},
			expectedTotalPrice: 400, // 130 + 2*50 + 5*30 + 20 = 400
			expectedMatches:    1,   // once 3*A
		},
		{
			description:        "Scenario B with 2*B=45",
			cart:               scenarioB(catalog),
			promos:             []promotions.Promotion{TwoBsPromo},
			expectedTotalPrice: 390, // 5*50 + 45 + 45 + 30 + 20 = 390
			expectedMatches:    2,   // twice 2*B
		},
		{
			description:        "Scenario B with C+D=30",
			cart:               scenarioB(catalog),
			promos:             []promotions.Promotion{CDComboPromo},
			expectedTotalPrice: 420, // 5*50 + 5*30 + 20 = 420
			expectedMatches:    0,   // no C+D
		},
		{
			description:        "Scenario C with 3*A=130",
			cart:               scenarioC(catalog),
			promos:             []promotions.Promotion{ThreeAsPromo},
			expectedTotalPrice: 315, // 130 + 2*50 + 5*30 + 20 + 15 = 315
			expectedMatches:    1,   // once 3*A
		},
		{
			description:        "Scenario C with 2*B=45",
			cart:               scenarioC(catalog),
			promos:             []promotions.Promotion{TwoBsPromo},
			expectedTotalPrice: 305, // 3*50 + 45 + 45 + 30 + 20 + 15 = 305
			expectedMatches:    2,   // twice 2*B
		},
		{
			description:        "Scenario C with C+D=30",
			cart:               scenarioC(catalog),
			promos:             []promotions.Promotion{CDComboPromo},
			expectedTotalPrice: 330, // 3*50 + 5*30 + 20 + 15 = 330
			expectedMatches:    1,   // once C+D
		},
	}
	for _, testcase := range testcases {
		matches := 0
		totalPrice := 0.0
		remainder := testcase.cart
		for _, promo := range testcase.promos {
			result, promoRemainder := promo(remainder)
			for _, match := range result {
				totalPrice += match.Price
			}
			remainder = promoRemainder
			matches += len(result)
		}
		for _, sku := range remainder {
			totalPrice += sku.UnitPrice
		}
		if totalPrice != testcase.expectedTotalPrice {
			t.Fatalf("Expected total price %v != %v in %q", testcase.expectedTotalPrice, totalPrice, testcase.description)
		}
		if matches != testcase.expectedMatches {
			t.Fatalf("Expected %v matches != %v in %q", testcase.expectedMatches, matches, testcase.description)
		}
	}
}
