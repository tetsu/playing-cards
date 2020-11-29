package main

func main() {
	d := newDeckFromFile("my_deck")

	d.shuffle()

	d.print()
}
