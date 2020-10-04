package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.saveToFile("Mis_cartas")
	cards = newDeckFromFile("Mis_cartas")
	cards.print()
}
