package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	var cardValue int

	switch card{
	case "ace": cardValue = 11
	case "two": cardValue = 2
	case "three": cardValue = 3
	case "four": cardValue = 4
	case "five": cardValue = 5
	case "six": cardValue = 6
	case "seven": cardValue = 7
	case "eight": cardValue = 8
	case "nine": cardValue = 9
	case "ten": cardValue = 10
	case "jack": cardValue = 10
	case "king": cardValue = 10
	case "queen": cardValue = 10
	default: return 0
	}

	return cardValue
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {

	card1AsInt, card2AsInt, dealerCardAsInt := ParseCard(card1), ParseCard(card2), ParseCard(dealerCard)
	var status string
	var checkValueInRange17 bool
	var checkValueInRange12 string

	for _, value := range []int{17, 18, 19, 20}{
		if card1AsInt + card2AsInt == value{
			checkValueInRange17 = true
		}
	} 


	for _, value := range []int{12, 13, 14, 15, 16}{
		if card1AsInt + card2AsInt == value && dealerCardAsInt < 7 {
		 	checkValueInRange12 = "S"
		}
		if card1AsInt + card2AsInt == value && dealerCardAsInt >= 7{
			checkValueInRange12 = "H"
		}
	}

	switch{
	case card1AsInt + card2AsInt <= 11:
		status = "H"
	case card1 == "ace" && card2 == "ace":
		status = "P"
	case checkValueInRange12 == "H":
		status = "H"
	case checkValueInRange12 == "S":
		status = "S"
	case checkValueInRange17:
		status = "S"
	case card1AsInt + card2AsInt <= 11:
		status = "H"
	case card1AsInt + card2AsInt == 21 && dealerCard != "ace" && dealerCard != "jack" && dealerCard != "king" && dealerCard != "queen" && dealerCardAsInt != 10:
		status = "W"
	case card1AsInt + card2AsInt == 21 && dealerCard == "ace" || dealerCard == "queen" || dealerCard == "king" || dealerCard == "jack" || dealerCardAsInt == 10:
		status = "S"
	}
	return status
}