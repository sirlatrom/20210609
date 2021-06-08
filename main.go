package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/sirlatrom/20210609/catalogs"
	"github.com/sirlatrom/20210609/model"
	"github.com/sirlatrom/20210609/promotions"
)

var (
	// ThreeAsPromo gives a reduced price for 3*A SKUs
	ThreeAsPromo = promotions.VolumeDiscount("A", 3, 130.0)
	// TwoBsPromo gives a reduced price for 2*B SKUs
	TwoBsPromo = promotions.VolumeDiscount("B", 2, 45.0)
	// CDComboPromo gives a reduced price for a pair of C+D SKUs
	CDComboPromo = promotions.ComboDiscount("C", "D", 30.0)
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanLines)
	for s.Scan() {
		cart := model.Cart{}
		line := s.Text()
		fmt.Printf("Cart as input: %q\n", line)
		for _, r := range line {
			skuID := string(r)
			sku := catalogs.StaticNaiveCatalog.SKUByID(skuID)
			if sku != nil {
				cart = append(cart, *sku)
			}
		}
		allPromos := []promotions.Promotion{ThreeAsPromo, TwoBsPromo, CDComboPromo}
		totalPrice := 0.0
		remainder := cart
		for _, promo := range allPromos {
			result, promoRemainder := promo(remainder)
			for _, match := range result {
				totalPrice += match.Price
				fmt.Printf("  - %v (subtotal %v)\n", match, totalPrice)
			}
			remainder = promoRemainder
		}
		for _, sku := range remainder {
			totalPrice += sku.UnitPrice
			fmt.Printf("  - %v (subtotal %v)\n", sku, totalPrice)
		}
		fmt.Printf("Total price: %v\n", totalPrice)
	}
}
