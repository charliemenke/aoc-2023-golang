package main

import (
	"encoding/json"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/charliemenke/aoc-2023-golang/utils"
)

type Set = map[int]bool

type Card struct {
    Id int
    ToMatchNums Set
    Nums Set
    WinningNums Set
    CardVal int
    CopiesWon []int
    MaxCopiesId int // just so i dont have to handle in in the loop
    Scratched bool
}
// func to create Card struct based on numbers to match, card numbers, and the 
// maximum cards provided in input
func CreateCard(id int, toMatchNums []int, nums []int, maxCopies int) Card {
    // use Set for both toMatchNums and Nums for instant lookup
    toMatchSet := Set{}
    for _, n := range toMatchNums {
        toMatchSet[n] = true
    }
    numsSet := Set{}
    for _, n := range nums {
        numsSet[n] = true
    }
    return Card{
        Id: id,
        ToMatchNums: toMatchSet,
        Nums: numsSet,
        Scratched: false,
        MaxCopiesId: maxCopies,
    }
}
// calculate winning numbers by doing an intersection on WinningNums and Nums
func (c *Card) SetWinningNums() Set {
    // set card state to scratched true
    c.Scratched = true
    // do an intersection between WinningNums & Nums
    intersection := map[int]bool{}
    for toMatch := range c.ToMatchNums {
        // will result in false in c.Nums does not have 'toMatch'
        if c.Nums[toMatch] {
            intersection[toMatch] = true
        }
    }
    c.WinningNums = intersection
    return intersection 
}

// for part 1. card val starts at 1 and for each WinningNums past len(1), double it
func (c *Card) SetCardVal() int {
    if len(c.WinningNums) == 0 {
        if !c.Scratched {
            log.Println("Card is not scratched, did you forget to run SetWinningNums()?")
        }
        c.CardVal = 0
        return 0
    }
    val := 0
    // val is 2^len(WinningNums)-1
    if len(c.WinningNums) > 0 {
        val = int(math.Pow(2, float64(len(c.WinningNums)-1)))
    }
    c.CardVal = val
    return val
}

// for part 2. add each of the won copy card's Ids to its parent card
// MaxCopiesId is used to make sure we dont add copy card id's that do 
// no exist
func (c *Card) SetWonCopiesNum() []int {
    if len(c.WinningNums) == 0 {
        if !c.Scratched {
            log.Println("Card is not scratched, did you forget to run SetWinningNums()?")
        }
        c.CopiesWon = []int{}
        return []int{}
    }
    val := []int{}
    for i := 1; i <= len(c.WinningNums); i++ {
        wonCopyId := i + c.Id
        if wonCopyId > c.MaxCopiesId {
            break
        }
        val = append(val, wonCopyId)
    }
    c.CopiesWon = val
    return val
} 


func main() {
    timer := utils.TimeFunc("p2")
    defer timer()
    input := utils.ReadInput("./input.txt")

    a := p2(input)
    log.Println(a)
}


func p1(input []string) int {
    result := 0
    
    cardsProcessed := 0
    for cardId, cardInput := range input {
        parsedCard := strings.Split(cardInput[strings.Index(cardInput, ":")+1:], " | ")

        reg := regexp.MustCompile(`\d+`)
        toMatchNumsStr := reg.FindAllString(parsedCard[0], -1)
        numsStr := reg.FindAllString(parsedCard[1], -1)

        toMatchNums := []int{}
        for _, n := range toMatchNumsStr {
            if v, err := strconv.Atoi(n); err == nil {
                toMatchNums = append(toMatchNums, v)
            } else {
                log.Fatal(err)
            }
        }
        nums := []int{}
        for _, n := range numsStr {
            if v, err := strconv.Atoi(n); err == nil {
                nums = append(nums, v)
            } else {
                log.Fatal(err)
            }
        }

        card := CreateCard(cardId+1, toMatchNums, nums, len(input))
        _ = card.SetWinningNums()
        _ = card.SetCardVal()

        result += card.CardVal
        cardsProcessed++
    }

    log.Printf("cards processed: %d", cardsProcessed)
    
    return result
}

func p2(input []string) int {
    // result is the num of of cards proccessed
    result := 0

    // map to track previous proccessed cards values
    // key = cardId, value = array of ids of copies won
    // {
    //   1: [2,3,4,5]
    //   2: [3,4]
    //   3: [4, 5]
    //   4: [5]
    //   5: []
    //   6: []
    // }
    cardLookUp := map[int][]int{}

    originalCards := []Card{}
    for cardId, cardInput := range input {
        parsedCard := strings.Split(cardInput[strings.Index(cardInput, ":")+1:], " | ")

        reg := regexp.MustCompile(`\d+`)
        toMatchNumsStr := reg.FindAllString(parsedCard[0], -1)
        numsStr := reg.FindAllString(parsedCard[1], -1)

        // convert all num strings to ints
        toMatchNums := []int{}
        for _, n := range toMatchNumsStr {
            if v, err := strconv.Atoi(n); err == nil {
                toMatchNums = append(toMatchNums, v)
            } else {
                log.Fatal(err)
            }
        }
        nums := []int{}
        for _, n := range numsStr {
            if v, err := strconv.Atoi(n); err == nil {
                nums = append(nums, v)
            } else {
                log.Fatal(err)
            }
        }

        // create card struct from input
        card := CreateCard(cardId+1, toMatchNums, nums, len(input))
        // calculate winning numbers
        _ = card.SetWinningNums()
        // calculate card copy ids won
        _ = card.SetWonCopiesNum()
        // append cards to array to process through copies later
        originalCards = append(originalCards, card)
        // since we are going through every card this first pass, we can completly fill
        // out the lookup table now so we dont have to recompute "card.SetWinningNums/SetWonCopiesNum"
        // later.
        cardLookUp[card.Id] = card.CopiesWon
        // add one to result since we are tracking all cards collected
        result++
    }

    // go through each original card and add up all our won copies
    for _, card := range originalCards {
        // for each original card's won copies, add 1 to result and recursivly sum
        // each copies won copies
        for _, id := range card.CopiesWon {
            result++
            CountCopies(&cardLookUp, id, &result)
        }
    }
    return result
}

func CountCopies(lookup *map[int][]int, cardId int, result *int) {
    // base case, exit when we reached a card with no won copies
    if len((*lookup)[cardId]) == 0 {
        return 
    }
    // else, add length of won copy's copies,
    // and recursivly dig through lookup table
    *result += len((*lookup)[cardId])
    for _, id := range (*lookup)[cardId] {
        CountCopies(lookup, id, result)
    }
}

