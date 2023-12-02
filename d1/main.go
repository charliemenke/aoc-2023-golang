package main

import (
    "log"
    "strconv"

    "github.com/charliemenke/aoc-2023-golang/utils"
)

type WordNum struct {
    len int
    intVal []rune
    val rune
    name string
}



func main() {
    timer := utils.TimeFunc("p2")
    defer timer()
    input := utils.ReadInput("./input.txt")

    a := p2(input)
    log.Println(a)
}

// 1abc2
// pqr3stu8vwx
// a1b2c3d4e5f
// treb7uchet
func p1(input []string) int {
    result := 0
    for _, row := range input {
        rowNum := []rune{}
        for _, c := range row {
            _, err := strconv.ParseInt(string(c),10,0)
            if err == nil {
                rowNum = append(rowNum, c)
            }
        }
        if len(rowNum) >= 2 {
            firstAndLast := []rune{rowNum[0], rowNum[len(rowNum) -1]}
            rowVal, err := strconv.ParseInt(string(firstAndLast), 10, 0)
            if err != nil {
                log.Fatal(err)
            }
            // log.Printf("row val: %d", rowVal)
            result = result + int(rowVal)
        } else if len(rowNum) == 1 {
            rowVal, err := strconv.ParseInt(string(append(rowNum, rowNum...)), 10, 0)
            if err != nil {
                log.Fatal(err)
            }
            // log.Printf("row val: %d", rowVal)
            result = result + int(rowVal)
        }
    }
    return result
}


func p2(input []string) int64 {
    defer utils.TimeFunc("p1")
    word2digit := map[rune][]WordNum{
        'o': {WordNum{len: 3, intVal: []rune{'n', 'e'}, val: '1', name: "one"}},
        't': {WordNum{len: 3, intVal: []rune{'w','o'}, val: '2', name: "two"}, WordNum{len: 5, intVal: []rune{'h','r','e','e'}, val: '3', name: "three"}},
        'f': {WordNum{len: 4, intVal: []rune{'o','u','r'}, val: '4', name: "four"}, WordNum{len: 4, intVal: []rune{'i','v','e'}, val: '5', name: "five"}},
        's': {WordNum{len: 3, intVal: []rune{'i','x'}, val: '6', name: "six"}, WordNum{len: 5, intVal: []rune{'e','v','e','n'}, val: '7', name: "seven"}},
        'e': {WordNum{len: 5, intVal: []rune{'i','g','h','t'}, val: '8', name: "eight"}},
        'n': {WordNum{len: 4, intVal: []rune{'i','n','e'}, val: '9', name: "nine"}},
    }
    result := int64(0)
    // for each line of input
    for _, row := range input {
        rowNum := []rune{}
        // for each char in line
        for i, c := range row {
            _, err := strconv.ParseInt(string(c),10,0)
            if err == nil {
                rowNum = append(rowNum, c)
            } else {
                // if char is apart of 1-9 characters check if word exists
                if numPossiblies, exists := word2digit[c]; exists {
                    for _, numPossibility := range numPossiblies {
                        //log.Printf("checking if %s exists, needs value of %d", numPossibility.name, numPossibility.intVal)
                        isMatch := true
                        if len(row) >= i+numPossibility.len-1 {
                            //0 1 2 3 4 5 5 7 8 9 10 11
                            //e i g h t w o t h r e  e
                            //for _, n := range row[i+1:i+numPossibility.len] {
                            //    log.Printf("%s", string(n))
                            //}
                            var iter string
                            if i+numPossibility.len >= len(row) {
                                iter = row[i+1:]
                            } else {
                                iter = row[i+1:i+numPossibility.len]
                            }
                            for iteri, c2 := range iter {
                                if numPossibility.intVal[iteri] != c2 {
                                   isMatch = false
                                   break
                               }
                            }
                            if isMatch {
                                //log.Printf("i: %d, end: %d", i, numPossibility.len)
                                //log.Printf("found match for %s: %+v", numPossibility.name, row[i:i+numPossibility.len])
                                rowNum = append(rowNum, numPossibility.val)
                            }
                        }
                    } 
                }
            }
        }
        if len(rowNum) >= 2 {
            //for _, n := range rowNum {
            //    log.Printf("%s", string(n))
            //}
            firstAndLast := []rune{rowNum[0], rowNum[len(rowNum)-1]}
            rowVal, err := strconv.ParseInt(string(firstAndLast), 10, 0)
            if err != nil {
                log.Fatal(err)
            }
            //log.Printf("line num: %d - row val: %d", ln, rowVal)
            result = result + rowVal
        } else if len(rowNum) == 1 {
            firstAndLast := []rune{rowNum[0], rowNum[0]}
            rowVal, err := strconv.ParseInt(string(firstAndLast), 10, 0)
            if err != nil {
                log.Fatal(err)
            }
            //log.Printf("line num: %d - row val: %d", ln, rowVal)
            result = result + rowVal
        } else {
            panic(0)
        }
    }
    return result

}
