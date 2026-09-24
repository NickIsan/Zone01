package piscine

type food struct {
	preptime int
}

func FoodDeliveryTime(order string) int {
	burger := food{preptime: 15}
	chips := food{preptime: 10}
	nuggets := food{preptime: 12}

	switch order {
	case "burger":
		return burger.preptime
	case "chips":
		return chips.preptime
	case "nuggets":
		return nuggets.preptime
	default:
		return 404
	}
}
