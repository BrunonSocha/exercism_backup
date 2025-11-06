package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int {"quarter_of_a_dozen": 3, "half_of_a_dozen": 6, "dozen": 12, "small_gross": 120, "gross": 144, "great_gross": 1728}
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	unitV, e := units[unit]
    itemV, ex := bill[item]
    if !e {
        return false
    } else if !ex {
        bill[item] = unitV
    } else {
        bill[item] = itemV+unitV
    }
    return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	unitV, e := units[unit]
    itemV, ex := bill[item]
    if !e || !ex {
        return false
    }
    if itemV - unitV < 0 {
        return false
    }
    if itemV == unitV {
        delete(bill, item)
    } else {
        bill[item] -= unitV
    }
    return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	itemV, ex := bill[item]
    return itemV, ex
}
