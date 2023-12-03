package main

import (
	// "fmt"
	"log"
	"regexp"
	"strconv"

	// "strconv"
	// "strings"

	"github.com/charliemenke/aoc-2023-golang/utils"
)

type Position struct {
    x int
    y int
}

type PartialPartNum struct {
    val string
    position Position
}

type PartNum struct {
    val int
    start Position
    end Position
}

func (p *PartNum) Equal(e PartNum) bool {
    return p.val == e.val && p.start.x == e.start.x && p.start.y == e.start.y && p.end.x == e.end.x && p.end.y == e.end.y
}

func main() {
    timer := utils.TimeFunc("p2")
    defer timer()
    input := utils.ReadInput("./input.txt")


    a := p2(input)
    log.Println(a)
}

func p1(input []string) int {

    // parse input into a 2d array, replaces anything != [0-9.] as a * (did this incase symbols got more complex in p2
    parsedInput := []string{}
    for _, line := range input {
        reg := regexp.MustCompile(`[^0-9.]+`)
        parsedLine := reg.ReplaceAllString(line, "*")
        parsedInput = append(parsedInput, parsedLine)
    }
    originalGrid := utils.BuidGrid(parsedInput)
    // pad grid on all sides by one more "." so we dont have to worry about array bounderies
    grid := padGrid(originalGrid) 
    utils.PrintGrid(grid)

    result := 0

    // for each 'row', find symbol. Then check all adjacent postions
    partNumbers := []PartNum{}
    for x := 1; x < len(grid) - 1; x++ {
        for y := 1; y < len(grid[x]) -1; y++ {
            // if char is symbol, check all adjacent positions
            if grid[x][y] == "*" {
                // check all adjecent postions
                adjacent := adjacentCords(Position{x: x, y: y}, )
                partialPartNums := []PartialPartNum{}
                // for all adjacent positions, check if any are digits
                for _, p := range adjacent {
                    // if postion is num, record it and its postion
                    if _, err := strconv.Atoi(grid[p.x][p.y]); err == nil {
                        partial := PartialPartNum{
                            val: grid[p.x][p.y],
                            position: p,
                        }
                        partialPartNums = append(partialPartNums, partial)
                    }
                }

                // for all found partial part nums, find begining and end offset for full number
                for _, partialPart := range partialPartNums {
                    part := PartNum{}
                    // find start of part num value (start) - either until "." or index y = 1
                    if partialPart.position.y == 1 {
                        part.start = Position{x: partialPart.position.x, y: 1}
                    } else {
                        part.start = Position{x: partialPart.position.x, y: partialPart.position.y}
                        for startY := partialPart.position.y-1; startY >= 1; startY-- {
                            if _, err := strconv.Atoi(grid[partialPart.position.x][startY]); err != nil {
                                break
                            }
                            part.start = Position{x: partialPart.position.x, y: startY}
                        }
                    }
                    // find end of part num value (end) - either "." or index y = len(grid[x]) - 2
                    if partialPart.position.y == len(grid[x]) - 2 {
                        part.end = Position{x: partialPart.position.x, y: len(grid[partialPart.position.x]) - 2 }
                    } else {
                        part.end = Position{x: partialPart.position.x, y: partialPart.position.y}
                        for endY := partialPart.position.y+1; endY < len(grid[partialPart.position.x]); endY++ {
                            if _, err := strconv.Atoi(grid[partialPart.position.x][endY]); err != nil {
                                break
                            }
                            part.end = Position{x: partialPart.position.x, y: endY}
                        }
                    }
                    // concat nums and conver to int
                    val := ""
                    for valI := part.start.y; valI <= part.end.y; valI++ {
                        val = val + grid[part.start.x][valI]
                    }
                    valParsed, err := strconv.Atoi(val)
                    if err != nil {
                        panic(1)
                    }
                    part.val = valParsed
                    partNumbers = append(partNumbers, part)
                }
            }
        }
    }

    // deduplicate all found part numbers
    // . . . . .
    // . * . . .
    // 4 5 . . .
    // . . . . .
    //
    // above grid would produce 2 Part Numbers as row 2 col 2 has
    // two "partial part nums" of '4' and '5', which then result into two
    // valid full "Part Nums" that really both point to '45' 
    deduped := deDupParts(partNumbers)
    
    // add up sum of part nums
    for _, p := range deduped {
        result = result + p.val
    }
    return result
}

func p2(input []string) int {

    // build grid based on input, this time do not replace all symbols as '*'
    originalGrid := utils.BuidGrid(input)
    // pad grid on all sides with "." so we dont have to worry about array bounderies
    grid := padGrid(originalGrid) 
    utils.PrintGrid(grid)

    result := 0

    // for each 'row' find symbol '*' and check all adjecent poistions
    for x := 1; x < len(grid) - 1; x++ {
        for y := 1; y < len(grid[x]) -1; y++ {
            // if char is symbol, check all adjecent positions
            if grid[x][y] == "*" {
                // check all adjecent postions
                adjacent := adjacentCords(Position{x: x, y: y}, )
                partialPartNums := []PartialPartNum{}
                for _, p := range adjacent {
                    // if postion is num, record it and its postion
                    if _, err := strconv.Atoi(grid[p.x][p.y]); err == nil {
                        partial := PartialPartNum{
                            val: grid[p.x][p.y],
                            position: p,
                        }
                        partialPartNums = append(partialPartNums, partial)
                    }
                }

                // for all found partial part nums, find start and end offsets of full part number
                partNumbers := []PartNum{}
                for _, partialPart := range partialPartNums {
                    part := PartNum{}
                    // find start of part num value (start) - either until "." or index y = 1
                    if partialPart.position.y == 1 {
                        part.start = Position{x: partialPart.position.x, y: 1}
                    } else {
                        part.start = Position{x: partialPart.position.x, y: partialPart.position.y}
                        for startY := partialPart.position.y-1; startY >= 1; startY-- {
                            if _, err := strconv.Atoi(grid[partialPart.position.x][startY]); err != nil {
                                break
                            }
                            part.start = Position{x: partialPart.position.x, y: startY}
                        }
                    }
                    // find end of part num value (end) - either "." or index y = len(grid[x]) - 2
                    if partialPart.position.y == len(grid[x]) - 2 {
                        part.end = Position{x: partialPart.position.x, y: len(grid[partialPart.position.x]) - 2 }
                    } else {
                        part.end = Position{x: partialPart.position.x, y: partialPart.position.y}
                        for endY := partialPart.position.y+1; endY < len(grid[partialPart.position.x]); endY++ {
                            if _, err := strconv.Atoi(grid[partialPart.position.x][endY]); err != nil {
                                break
                            }
                            part.end = Position{x: partialPart.position.x, y: endY}
                        }
                    }
                    // concat nums and conver to int
                    // log.Printf("%+v", part)
                    val := ""
                    for valI := part.start.y; valI <= part.end.y; valI++ {
                        val = val + grid[part.start.x][valI]
                    }
                    // log.Println(val)
                    valParsed, err := strconv.Atoi(val)
                    if err != nil {
                        panic(1)
                    }
                    part.val = valParsed
                    // log.Printf("%+v", part)
                    partNumbers = append(partNumbers, part)
                }
                // dedupe here so we are left with final adjacent part numbers of the specific '*' symbol
                dedupped := deDupParts(partNumbers)
                // only if there are two adjacent part numers, multiply their vals and add to result
                if len(dedupped) == 2 {
                    result = result + (dedupped[0].val * dedupped[1].val)
                }
            }
        }
    }

    return result
}

// pads grid in each direction so we can avoid boundery erros
func padGrid (grid [][]string) [][]string {
    result := make([][]string, len(grid) + 2)
    // add first row
    result[0] = []string{".",".",".",".",".",".",".",".",".",".", ".","."}

    for x := 0; x < len(grid); x++ {
        row := []string{}
        for y := 0; y < len(grid[x]); y++ {
            //log.Printf("x: %d - y: %d",x,y)
            if y == 0 {
                row = append(row, ".")
                row = append(row, grid[x][y])
            } else if y == len(grid[x]) - 1 {
                row = append(row, grid[x][y])
                row = append(row, ".")
            } else {
                row = append(row, grid[x][y])
            }
            //log.Println(row)
        }
        result[x+1] = row
    }
    // add last row
    result[len(result) -1] = []string{".",".",".",".",".",".",".",".",".",".", ".","."}
    return result
}

func adjacentCords(position Position) []Position {
    // adjacent cords are:

    validPositions := []Position{
    // 1. (position.x + 1, position.y) adjacent right
        {
            x: position.x + 1,
            y: position.y,
        },
    // 2. (position.x - 1, position.y) adjacent left
        {
            x: position.x - 1,
            y: position.y,
        },
    // 3. (position.x, position.y + 1) adjacane top
        {
            x: position.x,
            y: position.y + 1,
        },
    // 4. (position.x, position.y - 1) adjacent bottom
        {
            x: position.x,
            y: position.y - 1,
        },
    // 5. (position.x + 1, position.y + 1) adjacent top right
        {
            x: position.x + 1,
            y: position.y + 1,
        },
    // 6. (position.x - 1, position.y + 1) adjacent top left
        {
            x: position.x - 1,
            y: position.y + 1,
        },
    // 7. (position.x + 1) position.y - 1) adjacent bottom right
        {
            x: position.x + 1,
            y: position.y - 1,
        },
    // 8. (position.x - 1) position.y - 1) adjacent bottom left
        {
            x: position.x - 1,
            y: position.y - 1,
        },
    }


    return validPositions
}

func deDupParts(parts []PartNum) []PartNum {
    unique := []PartNum{}
partsloop:
    for _, p := range parts {
        for i, u := range unique {
            if u.Equal(p) {
                unique[i] = p
                continue partsloop
            }
        }
        unique = append(unique, p)
    }
    return unique
}



