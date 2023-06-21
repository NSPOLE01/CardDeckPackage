package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// creates a new type "deck"
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Clubs", "Hearts"}
	cardNumbers := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	for _, suit := range cardSuits {
		for _, number := range cardNumbers {
			cards = append(cards, number+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")

}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)

}

func newDeckFromFile(fileName string) deck {
	byteSlice, er := ioutil.ReadFile(fileName)
	if er != nil {
		fmt.Println("Error: ", er)
		os.Exit(1)

	}
	deckString := strings.Split(string(byteSlice), ",")
	return deck(deckString)

}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for index := range d {
		newIndex := r.Intn(len(d) - 1)
		d[index], d[newIndex] = d[newIndex], d[index]

	}

}
