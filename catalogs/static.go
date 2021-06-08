package catalogs

import (
	"github.com/sirlatrom/20210609/model"
)

type staticNaiveCatalog struct {
	SKUs map[string]model.SKU
}

func (c staticNaiveCatalog) SKUByID(id string) *model.SKU {
	sku, exists := c.SKUs[id]
	if !exists {
		return nil
	}
	return &sku
}

// StaticNaiveCatalog has a constant set of SKUs with fixed unit prices.
var StaticNaiveCatalog = staticNaiveCatalog{
	SKUs: map[string]model.SKU{
		"A": {
			ID:        "A",
			UnitPrice: 50.0,
		},
		"B": {
			ID:        "B",
			UnitPrice: 30.0,
		},
		"C": {
			ID:        "C",
			UnitPrice: 20.0,
		},
		"D": {
			ID:        "D",
			UnitPrice: 15.0,
		},
	},
}
