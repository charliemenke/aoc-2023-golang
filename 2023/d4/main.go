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
    CopiesWon int
    Scratched bool
}
func CreateCard(id int, toMatchNums []int, nums []int) Card {
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
    }
}
func (c *Card) PrintToMatchNums() {
    tm := []int{}
    for n := range c.ToMatchNums {
        tm = append(tm, n)
    }
    b, _ := json.Marshal(tm)
    log.Printf("%v", string(b))

}
func (c *Card) PrintNums() {
    nums := []int{}
    for n := range c.Nums {
        nums = append(nums, n)
    }
    b, _ := json.Marshal(nums)
    log.Printf("%v", string(b))

}
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

func (c *Card) SetWonCopiesNum() int {
    if len(c.WinningNums) == 0 {
        if !c.Scratched {
            log.Println("Card is not scratched, did you forget to run SetWinningNums()?")
        }
        c.CopiesWon = 0
        return 0
    }
    val := len(c.WinningNums)
    c.CopiesWon = val
    return val
} 


func main() {
    timer := utils.TimeFunc("p1")
    defer timer()
    input := utils.ReadInput("./input.txt")

    a := p1(input)
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
        //a, _ := json.Marshal(toMatchNumsStr)
        //b, _ := json.Marshal(numsStr)
        //log.Printf("%v | %v", string(a), string(b))

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

        card := CreateCard(cardId+1, toMatchNums, nums)
        _ = card.SetWinningNums()
        _ = card.SetCardVal()

        //log.Printf("card %d", card.Id)
        //card.PrintToMatchNums()
        //card.PrintNums()
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
    // key = cardId, value = winningCopies
    cardLookUp := map[int]int{}

    originalCards := []Card{}
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

        card := CreateCard(cardId+1, toMatchNums, nums)
        _ = card.SetWinningNums()
        _ = card.SetWonCopiesNum()
        originalCards = append(originalCards, card)
        result++
    }

    for _, card := range originalCards {
        // for each original card, solve won copies
        for i := card.Id; i < card.Id + card.CopiesWon; i++ {
            // check if card copy value is in lookup table
            if v, exists := cardLookUp[i]; exists {

            } else {
            // if not, calculate and add to look up table

            }
        }
    }



    return result
}

