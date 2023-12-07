package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var err error
	scanner := bufio.NewScanner(os.Stdin)
	hands := []HandCard{}
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Fields(line)
		hand := HandCard{hand: split[0]}
		hand.bid, err = strconv.Atoi(split[1])
		if err != nil {
			panic(err)
		}
		hands = append(hands, hand)
		sort.Slice(hands, func(i, j int) bool {
			return hands[i].isLesser(hands[j])
		})
	}
	total := 0
	for i := range hands {
		total += hands[i].bid * (i + 1)
	}
	fmt.Println(total)
}

type (
	HandCardType int
	HandCard     struct {
		hand string
		bid  int
	}
)

const (
	HighCard  HandCardType = 1
	OnePair   HandCardType = 2
	TwoPair   HandCardType = 3
	ThreeKind HandCardType = 4
	FullHouse HandCardType = 5
	FourKind  HandCardType = 6
	FiveKind  HandCardType = 7
)

var CardStrength = map[byte]int{
	'2': 1,
	'3': 2,
	'4': 3,
	'5': 4,
	'6': 5,
	'7': 6,
	'8': 7,
	'9': 8,
	'T': 9,
	'J': 10,
	'Q': 11,
	'K': 12,
	'A': 13,
}

func (h *HandCard) isLesser(dst HandCard) bool {
	if h.getType() == dst.getType() {
		for i := 0; i < len(h.hand); i++ {
			if CardStrength[h.hand[i]] == CardStrength[dst.hand[i]] {
				continue
			}
			return CardStrength[h.hand[i]] < CardStrength[dst.hand[i]]
		}
		fmt.Println("edge cases")
	}
	return h.getType() < dst.getType()
}

func (h *HandCard) getType() (card_type HandCardType) {
	cards := make(map[rune]int, 5)
	for _, v := range h.hand {
		cards[v]++
	}
	distinct_cards := len(cards)

	if distinct_cards == 1 {
		return FiveKind // AAAAA
	}
	if distinct_cards == 2 {
		if maxValueInMap(cards) == 4 {
			return FourKind // AAAAK
		} else {
			return FullHouse // AAAKK
		}
	}
	if distinct_cards == 3 {
		if maxValueInMap(cards) == 3 {
			return ThreeKind // AAAKQ
		} else {
			return TwoPair // AAKKQ
		}
	}
	if distinct_cards == 4 {
		return OnePair // AAKQJ
	}

	return HighCard
}

func maxValueInMap(m map[rune]int) int {
	max := 0
	for _, v := range m {
		if v > max {
			max = v
		}
	}
	return max
}
