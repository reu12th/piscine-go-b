package piscine

type food struct {
	preptime int
}

func FoodDeliveryTime(order string) int {
	var min food

	switch order {
	case "burger":
		min.preptime = 15
	case "chips":
		min.preptime = 10
	case "nuggets":
		min.preptime = 12
	default:
		return 404
	}
	return min.preptime
}
