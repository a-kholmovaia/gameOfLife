package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Universe struct {
	currentGen [][]string
}

func initUniverse(n int) ([][]string, int) {
	m := map[int]string{0: " ", 1: "O"}
	universe := make([][]string, n)
	alive := 0
	for i := 0; i < n; i++ {
		temp := make([]string, 0)
		for j := 0; j < n; j++ {
			el := m[rand.Intn(2)]
			if el == "O" {
				alive += 1
			}
			temp = append(temp, el)
		}
		universe[i] = append(universe[i], temp...)
	}
	return universe, alive
}

func getNextGen(gen [][]string, n int) ([][]string, int) {
	m := map[int]string{0: " ", 1: "O"}
	alive := 0
	universe := make([][]string, n)
	for i := 0; i < n; i++ {
		temp := make([]string, 0)
		for j := 0; j < n; j++ {
			el := m[isAlive(gen, i, j, n)]
			temp = append(temp, el)
			if el == "O" {
				alive += 1
			}
		}
		universe[i] = append(universe[i], temp...)
	}
	return universe, alive
}

func printUniverse(gen [][]string, n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			el := gen[i][j]
			fmt.Print(el)
		}
		fmt.Println()
	}
}

func isAlive(gen [][]string, row int, col int, n int) int {
	neighbors := 0
	neighbors = countAliveN(gen, row, col, n)
	if gen[row][col] == "O" {
		if 1 < neighbors && 4 > neighbors {
			return 1
		}
	} else {
		if neighbors == 3 {
			return 1
		}
	}
	return 0
}

func countAliveN(gen [][]string, row int, col int, n int) int {
	neighbors := 0
	for i := 0; i < 3; i++ {
		col_check := col - 1
		if gen[mod(row-1, n)][mod(col_check+i, n)] == "O" {
			neighbors += 1
		}
		if gen[mod(row+1, n)][mod(col_check+i, n)] == "O" {
			neighbors += 1
		}
		if i != 1 {
			if gen[row][mod(col_check+i, n)] == "O" {
				neighbors += 1
			}
		}
	}
	return neighbors
}

func mod(a, b int) int {
	return (a%b + b) % b
}

func main() {
	var n, alive int
	fmt.Scan(&n)
	var universe Universe
	universe.currentGen, alive = initUniverse(n)
	fmt.Printf("Generation #%d\nAlive: %d\n", 1, alive)
	printUniverse(universe.currentGen, n)
	i := 0
	for true {
		universe.currentGen, alive = getNextGen(universe.currentGen, n)
		time.Sleep(50 * time.Millisecond)
		fmt.Print("\033[H\033[2J")
		fmt.Printf("Generation #%d\nAlive: %d\n", i, alive)
		printUniverse(universe.currentGen, n)
		i += 1
		if alive == 0 {
			break
		}
	}
}
