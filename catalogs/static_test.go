package catalogs_test

import (
	"testing"

	"github.com/sirlatrom/20210609/catalogs"
)

func TestStaticNaiveCatalog(t *testing.T) {
	for _, skuID := range []string{"A", "B", "C", "D"} {
		if catalogs.StaticNaiveCatalog.SKUByID(skuID) == nil {
			t.Fatalf("Expected to find SKU by ID %q; got nil", skuID)
		}
	}

	if sku := catalogs.StaticNaiveCatalog.SKUByID("bogus"); sku != nil {
		t.Fatalf("Expected not to find SKU by ID %q; got %#v", "bogus", sku)
	}
}
