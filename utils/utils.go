package utils

import (
	"bufio"
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
