# CardDeckPackage

<p align="center">
    <img width="200" src="[https://www.thoughtco.com/thmb/Qx4zhTIddLwa9jXfelCbnSuhbkM=/1500x0/filters:no_upscale():max_bytes(150000):strip_icc()/close-up-of-cards-on-white-background-767988479-5c4bd7bb4cedfd0001ddb36e.jpg]" alt="Deck of Cards">
</p>

A simple package developed in GoLang with several functions for a deck of cards.

The package contains 6 specific functions:
1. newDeck() -> Generates a new deck of cards containing all 52 standard cards
2. print() -> Prints out all of the cards in the current hand
3. deal() -> Assigns cards at random based on the given hand size
4. saveToFile() -> Creates a text file with a string representation of all the cards in the current hand
5. newDeckFromFile() -> Creates a new deck using the cards in the specified file
6. shuffle() -> Takes the current deck, and shifts cards around at random in order to mix up the order
