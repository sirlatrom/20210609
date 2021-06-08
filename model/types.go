package model

import (
	"fmt"
	"strings"
)

// SKU represents a stock-keeping unit by an ID and a unit price.
type SKU struct {
	ID        string
	UnitPrice float64
}

func (s SKU) String() string {
	return fmt.Sprintf("%s = %v", s.ID, s.UnitPrice)
}

// OrderLine represents price for a collection of SKUs, which may differ from the sum of their unit prices.
type OrderLine struct {
	SKUs  []SKU
	Price float64
}

func (o OrderLine) String() string {
	skuIDs := []string{}
	for _, sku := range o.SKUs {
		skuIDs = append(skuIDs, sku.ID)
	}
	return fmt.Sprintf("%s = %v", strings.Join(skuIDs, " + "), o.Price)
}

// Catalog provides a function to look up an SKU by its ID.
type Catalog interface {
	SKUByID(id string) *SKU
}

// Cart is a collection of SKUs in a user's shopping cart.
type Cart []SKU
