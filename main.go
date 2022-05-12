package main

func main() {
	cards := newDeckFromFile("my_cardsa")
	cards.print()
	//cards.saveToFile("my_cards")
}
