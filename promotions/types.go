package promotions

import "github.com/sirlatrom/20210609/model"

// Promotion processes the given cart for matching promos, returning an order line for each match and the remaining cart SKUs.
type Promotion func(cart model.Cart) ([]model.OrderLine, model.Cart)
