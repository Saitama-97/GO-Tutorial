package main

import "fmt"

type Point struct {
	x int
	y int
}

type Player struct {
	Name        string
	HealthPoint int
	MagicPoint  int
}

type Command struct {
	Name    string
	Var     *int
	Comment string
}

type People struct {
	name  string
	child *People
}

func struc1() {
	p1 := Point{
		x: 3,
		y: 4,
	}

	var player1 = Player{
		Name:        "Alan",
		HealthPoint: 10,
		MagicPoint:  10,
	}

	version := 1

	cmd := Command{}
	cmd.Name = "version"
	cmd.Var = &version
	cmd.Comment = "show version"

	relation := &People{
		name: "Grandpa",
		child: &People{
			name: "Father",
			child: &People{
				name: "Me",
			},
		},
	}

	fmt.Println(p1, player1, cmd)
	fmt.Println(relation.child.child.name)
}
