package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/charliemenke/aoc-2023-golang/utils"
)

func main() {
    timer := utils.TimeFunc("p1")
    defer timer()
    input := utils.ReadInput("./input.txt")

    a := p2(input)
    log.Println(a)
}

func p1(input []string) int {
    // sum of game ids that meet condition
    result := 0
    
    maxRed := 12
    maxGreen := 13
    maxBlue := 14

    for gameId, gameSets := range input {
        parsedGamesSets := strings.Split(strings.TrimPrefix(gameSets, fmt.Sprintf("Game %d: ", gameId+1)), "; ")
        b, _ := json.Marshal(parsedGamesSets)
        log.Printf("%v", string(b))
        
        isPossible := true
        // for each game set, check dice value. if color exceeds maximum exit 
        // early and move on to next set. If no early exist, add gameId to result sume
gamesSets:
        for _, gameSet := range parsedGamesSets {
            for _, gameSetDice := range strings.Split(gameSet, ", ") {
                digit, err := strconv.ParseInt(strings.Split(gameSetDice, " ")[0], 10, 0)
                if err != nil {
                    log.Fatal(err)
                }
                //log.Printf("checking: %s", gameSetDice)
                //log.Printf("has value of: %d", digit)
                if strings.HasSuffix(gameSetDice, "red") {
                    if maxRed < int(digit) {
                        isPossible = false
                        break gamesSets
                    }
                } else if strings.HasSuffix(gameSetDice, "green") {
                    if maxGreen < int(digit) {
                        isPossible = false
                        break gamesSets
                    }
                } else if strings.HasSuffix(gameSetDice, "blue") {
                    if maxBlue < int(digit) {
                        isPossible = false
                        break gamesSets
                    }
                } else {
                    panic(1)
                }
            }
        }
        if isPossible {
            log.Printf("%s is possible", parsedGamesSets)
            result = result + (gameId + 1)
        }
    }

    return result
} 

func p2(input []string) int {
    result := 0
    for gameId, gameSets := range input {
        parsedGamesSets := strings.Split(strings.TrimPrefix(gameSets, fmt.Sprintf("Game %d: ", gameId+1)), "; ")
        //b, _ := json.Marshal(parsedGamesSets)
        //log.Printf("%v", string(b))
        maxRed := 0
        maxGreen := 0
        maxBlue := 0
        // for each game set, track largest value by color
        for _, gameSet := range parsedGamesSets {
            for _, gameSetDice := range strings.Split(gameSet, ", ") {
                digit, err := strconv.ParseInt(strings.Split(gameSetDice, " ")[0], 10, 0)
                if err != nil {
                    log.Fatal(err)
                }
                //log.Printf("checking: %s", gameSetDice)
                //log.Printf("has value of: %d", digit)
                if strings.HasSuffix(gameSetDice, "red") {
                    if int(digit) > maxRed {
                        maxRed = int(digit)
                    }
                } else if strings.HasSuffix(gameSetDice, "green") {
                    if int(digit) > maxGreen {
                        maxGreen = int(digit)
                    }
                } else if strings.HasSuffix(gameSetDice, "blue") {
                    if int(digit) > maxBlue {
                        maxBlue = int(digit)
                    }
                } else {
                    panic(1)
                }
            }
        }
        minSetPower := maxRed * maxGreen * maxBlue
        result = result + minSetPower
    }

    return result
} 

