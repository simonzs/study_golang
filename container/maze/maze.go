package main

import (
	"fmt"
	"io"
	"os"
)

func readMaze(reader io.Reader) [][]int {
	var r, c int
	fmt.Fscanf(reader, "%d %d", &r, &c)
	maze := make([][]int, r)
	for i := range maze {
		maze[i] = make([]int, c)
		for j := range maze[i] {
			fmt.Fscanf(reader, "%d", &maze[i][j])
		}
	}
	return maze
}

type point struct{ i, j int }

var dir []point = []point{
	{-1, 0}, {0, -1}, {1, 0}, {0, 1},
}

func main() {
	file, err := os.Open("container/maze/maze.in")
	if err != nil {
		panic(err)
	}

	maze := readMaze(file)
	r, c := len(maze), len(maze[0])
	steps := make([][]int, r)
	for i := range steps {
		steps[i] = make([]int, c)
	}

	startI, startJ := 0, 0
	endI, endJ := r-1, c-1

	queue := []point{
		{startI, startJ},
	}

	for len(queue) > 0 {
		// pop from queue
		current := queue[0]
		queue = queue[1:]

		if current.i == endI &&
			current.j == endJ {
			break
		}

		// explore 4 directions
		for _, d := range dir {
			nextI, nextJ :=
				current.i+d.i, current.j+d.j

			if nextI < 0 || nextI >= r ||
				nextJ < 0 || nextJ >= c ||
				maze[nextI][nextJ] == 1 ||
				steps[nextI][nextJ] != 0 ||
				(nextI == startI &&
					nextJ == startJ) {
				continue
			}

			queue = append(queue,
				point{nextI, nextJ})
			steps[nextI][nextJ] =
				steps[current.i][current.j] + 1
		}
	}

	for _, sr := range steps {
		for _, sc := range sr {
			fmt.Printf("%2d ", sc)
		}
		fmt.Println()
	}
}
