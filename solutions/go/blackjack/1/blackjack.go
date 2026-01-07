package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
    var cardValue int
    
	switch card {
        case "ace":
        	cardValue = 11
        case "two":
        	cardValue =  2
        case "three":
        	cardValue =  3
        case "four":
        	cardValue =  4
        case "five":
        	cardValue =  5
        case "six":
        	cardValue =  6
        case "seven":
        	cardValue =  7
        case "eight":
        	cardValue =  8
        case "nine":
        	cardValue =  9
        case "ten", "jack", "king", "queen":
        	cardValue = 10
        default:
        	cardValue =  0
    }

    return cardValue
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
    var descision string;
    myCardSum := ParseCard(card1) + ParseCard(card2)
    dealerCardValue := ParseCard(dealerCard)
    
	switch {
        case card1 == "ace" && card2 == "ace":
        	descision = "P"
        case myCardSum == 21 && dealerCardValue < 10:
        	descision = "W"
        case myCardSum == 21 && dealerCardValue >= 10:
        	descision = "S"
        case myCardSum >= 17 && myCardSum <= 20:
         	descision = "S"
        case myCardSum >= 12 && myCardSum <= 16 && dealerCardValue < 7:
        	descision = "S"
        case myCardSum >= 12 && myCardSum <= 16 && dealerCardValue >= 7:
        	descision = "H"
        case myCardSum <= 11:
        	descision = "H"    
    }

    return descision
}
