package interest

// InterestRate returns the rate in percentage based on the balance amount
func InterestRate(balance float64) float32 {
	if balance < 0{
		return 3.213
	} else if balance >= 0 && balance < 1000{
		return 0.5
	} else if balance >= 1000 && balance < 5000{
		return 1.621
	}
	return 2.475
}

// Interest calculates the annual simple interest earned on the balance amount
func Interest(balance float64) float64{
	rate := InterestRate(balance)
	numberOfYears := 1

	return balance/100 * float64(rate) * float64(numberOfYears)
}

// AnualBalanceUpdate returns the sum of the balance and interest earned on the balance
func AnnualBalanceUpdate(balance float64) float64{
	interest := Interest(balance)
	return interest + balance
}

// YearsBeforeDesiredBalance returns the number of years it will 
// take to earn a target balance
func YearsBeforeDesiredBalance(balance float64, targetBalance float64) int {
    currentBalance := balance
    years := 0
    
    for currentBalance < targetBalance {
        interest := Interest(currentBalance)
        currentBalance = currentBalance + interest
        years++
    }
    
    return years
}