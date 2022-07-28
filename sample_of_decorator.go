package main

import "fmt"

func main() {
	fmt.Println("test")
}

// type Dispaly struct {
// }

type Display interface {
	getColumns() int
	getRows() int
	getRowText() string
	show() string
}

type StringDisplay struct {
	Text string
}

func (d *StringDisplay) getColumns() int {
	return len(d.text)
}

func (d *StringDisplay) getRows() int {
	return 1
}

func (d *StringDisplay) getRowText() string {
	return d.text
}

func (d *StringDisplay) show() string {
	return d.text
}

type Border struct {
	StringDisplay
}

func (b *Border) getColumns() int {
	return 1 + len(b.Text) + 1
}

func (b *Border) getRows() int {
	return 1 + b.StringDisplay.getRows() + 1
}

func (b *Border) getRowText() string {
	return "|" + b.Text + "|"
}

type SideBoarder struct {
}

type FullBoarder struct {
}
