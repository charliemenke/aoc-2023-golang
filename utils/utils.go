package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func ReadInput(filePath string) []string {
    file, err := os.Open(filePath)
    if err != nil {
        log.Fatalln(err)
    }
    defer file.Close()

    inputLines := []string{}
    scan := bufio.NewScanner(file)
    for scan.Scan() {
        inputLines = append(inputLines, scan.Text())
    }
    return inputLines
}

func TimeFunc(funcName string) func() {
    start := time.Now()
    return func() {
        log.Printf("%s took %v", funcName, time.Since(start))
    }
}

func BuidGrid(input []string) [][]string {
    result := make([][]string, len(input))
    for i, row := range input {
        rowVal := []string{}
        for _, c := range row {
            rowVal = append(rowVal, string(c))
        }
        result[i] = rowVal
    }
    return result
}

func PrintGrid(grid [][]string) {
    for row := 0; row < len(grid); row++ {
        for col := 0; col < len(grid[row]); col++ {
            fmt.Print(grid[row][col], " ")
        }
        fmt.Print("\n")
    }
}
