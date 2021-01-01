package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"
)

func mustInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
		return 0
	}
	return i
}

func main() {
	t := time.Now()
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	part1(input)
	part2(input)

	log.Println("success in", time.Since(t))
}

func part1(input []byte) {
	fmt.Println("answer part 1:", calcPart1(input))
}

func part2(input []byte) {
	fmt.Println("answer part 2:", calcPart2(input))
}

func calcPart1(input []byte) int {
	p1, p2 := setupPlayers(input)
	runMatch(p1, p2)
	sum := calculateResult(p1, p2)
	return sum
}

func calculateResult(p1 *player, p2 *player) int {
	var winner *player
	if p1.count > p2.count {
		winner = p1
	} else {
		winner = p2
	}

	var sum int
	for card := range winner.ch {
		sum += card * winner.count
		winner.count--
	}
	return sum
}

type player struct {
	ch    chan int
	count int
	sum   int
}

func (p *player) addCard(card int) {
	p.count++
	p.sum += card
	p.ch <- card
}

func (p *player) drawCard() (int, bool) {
	select {
	case card := <-p.ch:
		p.count--
		p.sum -= card
		return card, true
	default:
		return -1, false
	}
}

func (p *player) cardList() []int {
	var cards []int
	for i := 0; i < p.count; i++ {
		card, _ := p.drawCard()
		cards = append(cards, card)
		p.addCard(card)
	}
	return cards
}

func addCardToTopAgain(p1 *player, card1 int) {
	p1.addCard(card1)
	nCard := p1.count - 1
	for i := 0; i < nCard; i++ {
		p1.ch <- <-p1.ch
	}
}

func setupPlayers(input []byte) (*player, *player) {
	players := strings.Split(string(input), "\n\n")

	deck1 := strings.Split(players[0], "\n")[1:]
	deck2 := strings.Split(players[1], "\n")[1:]

	p1 := &player{
		ch:    make(chan int, len(deck1)+len(deck2)),
		count: 0,
	}
	p2 := &player{
		ch:    make(chan int, len(deck1)+len(deck2)),
		count: 0,
	}

	for _, card := range deck1 {
		p1.addCard(mustInt(card))
	}
	for _, card := range deck2 {
		p2.addCard(mustInt(card))
	}
	return p1, p2
}

func runMatch(p1 *player, p2 *player) {
	defer func() {
		close(p1.ch)
		close(p2.ch)
	}()

	for {
		// Player 1
		card1, ok := p1.drawCard()
		if !ok {
			fmt.Println("done 1")
			return
		}

		// Player 2
		card2, ok := p2.drawCard()
		if !ok {
			// Workaround: since we have drawn a card from the front of the deck, we want to put it back on top
			addCardToTopAgain(p1, card1)
			return
		}

		// Match!
		if card1 > card2 {
			p1.addCard(card1)
			p1.addCard(card2)
		} else {
			p2.addCard(card2)
			p2.addCard(card1)
		}
	}
}

type previousDecks struct {
	list [][]int
}

func (d *previousDecks) addDecks(p1 *player, p2 *player) {
	cards := append(p1.cardList(), -1)
	cards = append(cards, p2.cardList()...)
	d.list = append(d.list, cards)
}

func (d *previousDecks) addAndCheckForLoop(p1 *player, p2 *player) bool {
	cards := append(p1.cardList(), -1)
	cards = append(cards, p2.cardList()...)

	//fmt.Println("checking!")
	//fmt.Println(cards)

equals:
	for _, ints := range d.list {
		if len(ints) != len(cards) {
			// Not identical
			continue
		}

		for i, v := range cards {
			if v != ints[i] {
				// Not identical
				continue equals
			}
		}
		return false
	}

	d.list = append(d.list, cards)
	return true
}

func runMatchRecursive2(p1 *player, p2 *player) bool {
	defer func() {
		close(p1.ch)
		close(p2.ch)
	}()

	//var rounds int
	loopTracker := map[[2]int]*previousDecks{}

	for {
		//fmt.Printf("p1: %d, p2: %d\n", p1.sum, p2.sum)
		sums := [2]int{p1.sum, p2.sum}
		w, exists := loopTracker[sums]
		if !exists {
			cc := &previousDecks{}
			cc.addDecks(p1, p2)
			loopTracker[sums] = cc
		} else {
			ok := w.addAndCheckForLoop(p1, p2)
			if !ok {
				return true
			}
		}

		//// TODO guard for recursion
		//rounds++
		//if rounds%100 == 0 {
		//	fmt.Println(rounds)
		//}

		//if rounds >= 500 {
		//	fmt.Println("reached, player 1 probably won", rounds)
		//	return true
		//}

		// Player 1
		card1, ok := p1.drawCard()
		if !ok {
			return false
		}

		// Player 2
		card2, ok := p2.drawCard()
		if !ok {
			// Workaround: since we have drawn a card from the front of the deck, we want to put it back on top
			addCardToTopAgain(p1, card1)
			return true
		}

		var p1Won bool
		playNewGame := p1.count >= card1 && p2.count >= card2
		if playNewGame {
			// Clarify variables
			numCard1, numCard2 := card1, card2
			deckSize := numCard1 + numCard2

			p1c := &player{
				ch:    make(chan int, deckSize),
				count: 0,
			}
			p2c := &player{
				ch:    make(chan int, deckSize),
				count: 0,
			}
			copyCards(p1, p1c, numCard1)
			copyCards(p2, p2c, numCard2)

			p1Won = runMatchRecursive2(p1c, p2c)
		} else {
			p1Won = card1 > card2
		}

		if p1Won {
			p1.addCard(card1)
			p1.addCard(card2)
		} else {
			p2.addCard(card2)
			p2.addCard(card1)
		}
	}
}
func copyCards(p1 *player, p1c *player, card1 int) {
	for i := 0; i < p1.count; i++ {
		c, ok := p1.drawCard()
		if !ok {
			log.Fatal("failed to draw c in unexpected place")
		}

		// Add the c back to the bottom of the deck
		p1.addCard(c)

		// Also add to player copy
		if i < card1 {
			p1c.addCard(c)
		}
	}
}

func calcPart2(input []byte) int {
	p1, p2 := setupPlayers(input)
	runMatchRecursive2(p1, p2)
	sum := calculateResult(p1, p2)
	return sum
}

// Part 1: 30882 too high
// Part 2: 4.357837586s
