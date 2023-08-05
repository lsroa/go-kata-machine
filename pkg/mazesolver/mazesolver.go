package mazesolver

import "fmt"

type Point struct {
	X int
	Y int
}

func (a *Point) add(b Point) Point {
	return Point{
		X: a.X + b.X,
		Y: a.Y + b.Y,
	}
}

var DIR = []Point{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

type MazeSolver struct {
	wall  string
	end   *Point
	start *Point
	Path  *[]Point
	seen  [][]bool
	maze  []string
}

func New(input []string, wall string) *MazeSolver {

	dx := len(input[0])
	dy := len(input)

	var seen = make([][]bool, dy)
	for i := range seen {
		seen[i] = make([]bool, dx)
	}

	return &MazeSolver{
		wall: wall,
		end:  &Point{},
		maze: input,
		Path: &[]Point{},
		seen: seen,
	}
}

func (m *MazeSolver) walk(current Point) bool {
	if current.X < 0 || current.Y < 0 || current.Y >= len(m.maze) || current.X >= len(m.maze[0]) {
		println("out of bounds")
		return false
	}

	if string(m.maze[current.Y][current.X]) == m.wall {
		println("its a wall")
		return false
	}

	if current == *m.end {
		println("solved")
		*m.Path = append(*m.Path, current)
		return true
	}

	if m.seen[current.Y][current.X] {
		fmt.Println("repeated")
		return false
	}

	// precondition
	*m.Path = append(*m.Path, current)
	m.seen[current.Y][current.X] = true

	// Recursion
	for _, dir := range DIR {
		if m.walk(current.add(dir)) {
			return true
		}
	}

	// postcondition
	*m.Path = (*m.Path)[:len(*m.Path)-1]
	return false
}

func (m *MazeSolver) Solve(maze []string, start, end Point) {
	dy := len(maze)
	dx := len(maze[0])
	m.end = &end

	seen := make([][]bool, dy)
	for i := range seen {
		seen[i] = make([]bool, dx)
	}

	m.walk(start)
}
