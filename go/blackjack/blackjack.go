package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	//panic("Please implement the ParseCard function")
	//ace	11	eight	8
	//two	2	nine	9
	//three	3	ten	10
	//four	4	jack	10
	//five	5	queen	10
	//six	6	king	10
	//seven	7	other	0
	var cardValue int
	switch card {
	case "ace":
		cardValue = 11
	case "two":
		cardValue = 2
	case "three":
		cardValue = 3
	case "four":
		cardValue = 4
	case "five":
		cardValue = 5
	case "six":
		cardValue = 6
	case "seven":
		cardValue = 7
	case "eight":
		cardValue = 8
	case "nine":
		cardValue = 9
	case "ten", "jack", "queen", "king":
		cardValue = 10
	default:
		cardValue = 0
	}
	return cardValue
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	//panic("Please implement the FirstTurn function")
	card1Value := ParseCard(card1)
	card2Value := ParseCard(card2)
	dealerCardValue := ParseCard(dealerCard)

	var result string
	switch {
	case card1 == "ace" && card2 == "ace":
		result = "P"
	case card1Value+card2Value <= 11 || (card1Value+card2Value >= 12 && card1Value+card2Value <= 16 && dealerCardValue >= 7):
		result = "H"
	case (card1Value+card2Value >= 17 && card1Value+card2Value <= 20) || (card1Value+card2Value >= 12 && card1Value+card2Value <= 16 && dealerCardValue < 7) || (card1Value+card2Value == 21 && dealerCardValue >= 10):
		result = "S"
	case card1Value+card2Value == 21 && dealerCardValue < 10:
		result = "W"
	}
	return result
}
