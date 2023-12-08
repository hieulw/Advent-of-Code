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
	withJoker := true

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
	}

	if withJoker {
		CardStrength['J'] = 0
	}
	sort.Slice(hands, func(i, j int) bool {
		return hands[i].isLesser(hands[j], withJoker)
	})

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
	HighCard HandCardType = iota
	OnePair
	TwoPair
	ThreeKind
	FullHouse
	FourKind
	FiveKind
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

func (h *HandCard) isLesser(dst HandCard, withJoker bool) bool {
	scoreTypeSrc := h.getType()
	scoreTypeDst := dst.getType()
	if withJoker {
		scoreTypeSrc = h.getJokerType()
		scoreTypeDst = dst.getJokerType()
	}
	if scoreTypeSrc == scoreTypeDst {
		for i := 0; i < len(h.hand); i++ {
			if CardStrength[h.hand[i]] == CardStrength[dst.hand[i]] {
				continue
			}
			return CardStrength[h.hand[i]] < CardStrength[dst.hand[i]]
		}
	}
	return scoreTypeSrc < scoreTypeDst
}

func (h *HandCard) getJokerType() (card_type HandCardType) {
	cards := make(map[rune]int)
	for _, v := range h.hand {
		cards[v]++
	}
	distinct_cards := len(cards)

	if strings.ContainsRune(h.hand, 'J') {
		if distinct_cards == 2 {
			return FiveKind // AAAAJ|AAAJJ -> AAAAA
		}
		if distinct_cards == 3 {
			if cards['J'] == 3 || cards['J'] == 2 {
				return FourKind // JJJKQ|JJKKQ -> KKKKQ
			}
			if cards[mostCard(cards)] == 3 {
				return FourKind // JKKKQ -> KKKKQ
			}
			return FullHouse // JKKQQ -> KKKQQ
		}
		if distinct_cards == 4 {
			return ThreeKind // JQKAA -> AQKAA
		}
		if distinct_cards == 5 {
			return OnePair // TJQKA -> TTQKA
		}
	}
	return h.getType()
}

func (h *HandCard) getType() (card_type HandCardType) {
	cards := make(map[rune]int)
	for _, v := range h.hand {
		cards[v]++
	}
	distinct_cards := len(cards)

	if distinct_cards == 1 {
		return FiveKind // AAAAA
	}
	if distinct_cards == 2 {
		if cards[mostCard(cards)] == 4 {
			return FourKind // AAAAK
		} else {
			return FullHouse // AAAKK
		}
	}
	if distinct_cards == 3 {
		if cards[mostCard(cards)] == 3 {
			return ThreeKind // AAAKQ
		} else {
			return TwoPair // AAKKQ
		}
	}
	if distinct_cards == 4 {
		return OnePair // AAKQJ
	}

	return HighCard // TJQKA
}

func mostCard(m map[rune]int) rune {
	var maxCard rune
	maxValue := 0
	for k := range m {
		if m[k] > maxValue {
			maxValue, maxCard = m[k], k
		}
	}
	return maxCard
}
