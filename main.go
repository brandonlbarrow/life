package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	live int = 1
	dead int = 0
)

type block [][]int

func (b block) Print() {
	for i := range b {
		for ii := range b[i] {
			fmt.Print(b[i][ii], "\t")
		}
		fmt.Println()
	}
}

func main() {
	b := makeBlock(20)
	fmt.Println("start")
	for {
		b = generation(b)
		b.Print()
		time.Sleep(1 * time.Second)
	}
}

func generation(b block) block {
	for i := range b[0:len(b):len(b)] {
		for j := range b[i:len(b[i]):len(b[i])] {
			n := neighbors(i, j, b)
			newCell := evaluate(b[i][j], n)
			b[i][j] = newCell
		}
		fmt.Println()
	}
	return b
}

// top left
// b[i-1][j-1]
// top
// b[i-1][j]
// top right
// b[i-1][j+1]
// left
// b[i][j-1]
// right
// b[i][j+1]
// bottom left
// b[i+1][j+1]
// bottom
// b[i+1][j]
// bottom right
// b[i+1][j+1]
func neighbors(i, j int, b block) int {
	if i > 0 && i < len(b) && j > 0 && j < len(b[i]) {
		return b[i-1][j-1] + b[i-1][j] + b[i-1][j+1] + b[i][j-1] + b[i][j+1] + b[i+1][j-1] + b[i+1][j] + b[i+1][j+1]
	} else {
		return 0
	}
}

// Each cell with one or no neighbors dies, as if by solitude.
// Each cell with two or three neighbors survives.
// Each cell with four or more neighbors dies, as if by overpopulation.
// Each dead cell with three neighbors becomes alive.
func evaluate(cell, liveNeighbors int) int {
	switch cell {
	case live:
		switch {
		case liveNeighbors < 2 || liveNeighbors > 3:
			return dead
		}
	case dead:
		switch {
		case liveNeighbors == 3:
			return live
		}
	}
	return cell
}

func makeBlock(size int) block {
	block := make([][]int, 0)
	for i := 0; i < size; i++ {
		out := make([]int, 0)
		for j := 0; j < size; j++ {
			out = append(out, generate0or1())
		}
		block = append(block, out)
	}
	return block
}

func generate0or1() int {
	return rand.Intn(2)
}
