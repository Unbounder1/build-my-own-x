package main

import "fmt"

func main() {
	var i Item
	fmt.Printf("i: %#v\n", i)

	i = Item{10, 20}
	fmt.Printf("i: %#v\n", i)

	// use %#v for debugging/logging

	i = Item{
		X: 11,
		Y: 22,
	}
	fmt.Printf("i: %#v\n", i)

	fmt.Println(NewItem(10, 20))
	fmt.Println(NewItem(10, 200))

	i.Move(4, 4)
	fmt.Printf("i (move): %#v\n", i)

	p1 := Player{
		Name: "parzival",
	}
	fmt.Printf("p1: %+v\n", p1)
	fmt.Println("p1.X:", p1.X) // if two names with the same names, compilationError, be more specific and use ie p1.Item.X

	p1.Move(100, 200)
	fmt.Printf("p1 (move): %+v\n", p1)

	fmt.Println(p1.Found("copper")) // nil
	fmt.Println(p1.Found("copper")) // nil
	fmt.Println(p1.Found("gold"))   // unknown key: "gold"
	fmt.Println("keys:", p1.Keys)   // keys: [copper]

	ms := []Mover{
		&i, &p1,
	}

	moveAll(ms, 50, 70)
	for _, m := range ms {
		fmt.Println(m)
	}

}

/*
	Interfaces

- Set of methods (and types)
- We define interfaces as "what you need", not "what you provide"
  - Interfaces are small (stdlib average 2 methods per interface)
  - Interface with more than 4 methods is probably wrong

- Don't start with interfaces, start with concrete types
*/
type Mover interface {
	Move(int, int)
}

func moveAll(ms []Mover, dx, dy int) {
	for _, m := range ms {
		m.Move(dx, dy)
	}
}

// Add a "Keys" field to Player with is a slice of strings
// Add a found(key string) method to player
// err if key is not one of "jade, "copper", "crystal"
// should add a key onlly one
type Player struct {
	Name string
	Keys []string
	Item // player embeds item
}

func (p *Player) Found(key string) error {
	if key != "copper" && key != "jade" && key != "crystal" {
		return fmt.Errorf("Unknown Key: %s", key)
	}

	// slices.COntains(p.Keys, key)
	for _, k := range p.Keys {
		if k == key {
			return nil
		}
	}
	p.Keys = append(p.Keys, key)

	return nil
}

// Moves i by delta x and delta y
// i is called the reciever, must use pointer semantics
// now i is a pointer receiver
/* Value vs pointer reciever
- in general, use value semantics
- try to keep same semantics on all methods
- When you must use pointer reciever
	- If you have a lock field (copying a lock is not good do not do)
	- If you need to mutate the struct
	- Decoding/unmarshaling
*/
func (i *Item) Move(dx, dy int) { // Cannot use any inheritance, extended vs embedded
	i.X += dx
	i.Y += dy
}

// func NewItem(x, y int) Item
// func NewItem(x, y int) *Item
// func NewItem(x, y int) *Item, error
// func NewItem(x, y int) Item, error

func NewItem(x, y int) (*Item, error) {
	if x < 0 || x > maxX || y < 0 || y > maxY {
		return nil, fmt.Errorf("Out of bounds")
	}
	i := Item{
		X: x,
		Y: y,
	}
	// the go compiler does escape analysis and will allocate i on the heap
	return &i, nil
}

const (
	maxX = 600
	maxY = 400
)

type Item struct {
	X int
	Y int
}
