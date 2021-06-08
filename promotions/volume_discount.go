package promotions

import "github.com/sirlatrom/20210609/model"

// VolumeDiscount finds all whole N-multiples of the given SKU ID and gives a promotion price for each of them, while returning the remainder unchanged.
var VolumeDiscount = func(skuID string, n int, promoPrice float64) Promotion {
	return func(cart model.Cart) ([]model.OrderLine, model.Cart) {
		orderLines := []model.OrderLine{}
		volumeCollection := []model.SKU{}
		remainder := model.Cart{}
		for _, sku := range cart {
			switch sku.ID {
			case skuID:
				volumeCollection = append(volumeCollection, sku)
				if len(volumeCollection) == n {
					orderLines = append(orderLines, model.OrderLine{
						SKUs:  volumeCollection,
						Price: promoPrice,
					})
					volumeCollection = []model.SKU{}
				}
			default:
				remainder = append(remainder, sku)
			}
		}
		remainder = append(remainder, volumeCollection...)
		return orderLines, remainder
	}
}
