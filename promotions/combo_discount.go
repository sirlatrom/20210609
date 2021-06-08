package promotions

import "github.com/sirlatrom/20210609/model"

// ComboDiscount finds all matching pairs of given SKU IDs and gives a promotion price for each of them, while returning the remainder unchanged.
var ComboDiscount = func(skuID1, skuID2 string, promoPrice float64) Promotion {
	return func(cart model.Cart) ([]model.OrderLine, model.Cart) {
		orderLines := []model.OrderLine{}
		sku1s := []model.SKU{}
		sku2s := []model.SKU{}
		remainder := model.Cart{}

		for _, sku := range cart {
			switch sku.ID {
			case skuID1:
				sku1s = append(sku1s, sku)
			case skuID2:
				sku2s = append(sku2s, sku)
			default:
				remainder = append(remainder, sku)
			}
		}

		for len(sku1s) > 0 && len(sku2s) > 0 {
			orderLines = append(orderLines, model.OrderLine{
				SKUs:  []model.SKU{sku1s[0], sku2s[0]},
				Price: promoPrice,
			})
			sku1s = sku1s[1:]
			sku2s = sku2s[1:]
		}

		remainder = append(remainder, sku1s...)
		remainder = append(remainder, sku2s...)
		return orderLines, remainder
	}
}
