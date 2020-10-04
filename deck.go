package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Creamos un nuevo tipo de mazo que es un slice de strings
type deck []string

//Método para crear un nuevo mazo
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Picas", "Diamantes", "Corazones", "Tréboles"}
	cardValues := []string{"As", "Dos", "Tres", "Cuatro", "Cinco", "Seis", "Siete", "Ocho", "Nueve", "Diez", "Jack", "Reina", "Rey"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" de "+suit)
		}
	}
	return cards
}

//Imprime el mazo
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

//Crea dos nuevos mazos, el primero desde el inicio de las cartas hasta el número que le pasemos
//y el segundo con el resto de las cartas
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

//Convierte el tipo deck al tipo string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

//Guarda el mazo en un archivo
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

//Genera un mazo con cartas random
func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
