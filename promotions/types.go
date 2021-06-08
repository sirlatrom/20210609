package promotions

// SKU represents a stock-keeping unit by an ID and a list price.
type SKU struct {
	ID        string
	ListPrice float64
}

// OrderLine represents price for a collection of SKUs, which may differ from the sum of their list prices.
type OrderLine struct {
	SKUs  []SKU
	Price float64
}

// Catalog provides a function to look up an SKU by its ID.
type Catalog interface {
	SKUByID(id string) SKU
}

// Cart is a collection of SKUs in a user's shopping cart.
type Cart []SKU

// Promotion processes the given cart for matching promos, returning an order line for each match and the remaining cart SKUs.
type Promotion func(cart Cart) ([]OrderLine, Cart)
