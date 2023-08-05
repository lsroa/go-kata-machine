package mazesolver_test

import (
	ms "lsroa/go-machine/pkg/mazesolver"
	"testing"
)

func TestMazesolver(t *testing.T) {

	maze := []string{
		"xxxxxxxxxx x",
		"x        x x",
		"x        x x",
		"x xxxxxxxx x",
		"x          x",
		"x xxxxxxxxxx",
	}

	want := []ms.Point{
		{10, 0},
		{10, 1},
		{10, 2},
		{10, 3},
		{10, 4},
		{9, 4},
		{8, 4},
		{7, 4},
		{6, 4},
		{5, 4},
		{4, 4},
		{3, 4},
		{2, 4},
		{1, 4},
		{1, 5},
	}

	solver := ms.New(maze, "x")
	solver.Solve(maze, ms.Point{10, 0}, ms.Point{1, 5})

	got := solver.Path

	for i, point := range want {
		if len(*got) != len(want) {
			t.Errorf("got: %v, want: %v", got, want)
			break
		}

		if (*got)[i] != point {
			t.Errorf("got: %v, want: %v", (*got)[i], point)
		}
	}
}
