package promotions

// VolumeDiscount finds all whole N-multiples of the given SKU ID and gives a promotion price for each of them, while returning the remainder unchanged.
var VolumeDiscount = func(skuID string, n int, promoPrice float64) Promotion {
	return func(cart Cart) ([]OrderLine, Cart) {
		orderLines := []OrderLine{}
		volumeCollection := []SKU{}
		remainder := Cart{}
		for _, sku := range cart {
			switch sku.ID {
			case skuID:
				volumeCollection = append(volumeCollection, sku)
				if len(volumeCollection) == n {
					orderLines = append(orderLines, OrderLine{
						SKUs:  volumeCollection,
						Price: promoPrice,
					})
					volumeCollection = []SKU{}
				}
			default:
				remainder = append(remainder, sku)
			}
		}
		remainder = append(remainder, volumeCollection...)
		return orderLines, remainder
	}
}
