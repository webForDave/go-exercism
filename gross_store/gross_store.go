package grossstore


func Units() map[string]int {
	measurements := make(map[string]int)
	measurements["quarter_of_a_dozen"] = 3
	measurements["half_of_a_dozen"] = 6
	measurements["dozen"] = 12
	measurements["small_gross"] = 120
	measurements["gross"] = 144
	measurements["great_gross"] = 1728

	return measurements
}

func NewBill() map[string]int {
	return make(map[string]int)
}

func AddItem(bill, units map[string]int, item, unit string) bool {

	// isUnitInUnits stores whether the unit argument exists in the units argument
	_, isUnitInUnits := units[unit]

	// this checks if the unit is not present in units
	if !isUnitInUnits{
		return false
	}

	// isItemInBill stores whether the item argument exists in the bill 
	_, isItemInBill := bill[item]

	if isItemInBill{
		bill[item] = bill[item] + units[unit]
	}else {
		bill[item] = units[unit]
	}

	return true

}

func RemoveItem(bill, units map[string]int, item, unit string) bool {
	isItemInBill := false
	_, isUnitInUnits := units[unit]

	if !isUnitInUnits{
		return false
	}

	for key := range bill{
		if key == item{
			isItemInBill = true
		}
	}

	if !isItemInBill{
		return false 
	}

	originalBillItem := bill[item]
	bill[item] = bill[item] - units[unit]

	if bill[item] < 0{
		bill[item] = originalBillItem
		return false
	}

	if bill[item] == 0{
		delete(bill, item)
		return true 
	}


	return true
}

func GetItem(bill map[string]int, item string) (int, bool) {

	for key, value := range bill{
		if item == key{
			return value, true
		}
	}
	return 0, false
}