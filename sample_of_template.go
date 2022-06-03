package main

import "strings"

func main() {
	// 3つのPrintable5Timesインターフェースを持つオブジェクトを生成し、それぞれをPrintメソッドに渡す
	printables := []Printable5Times{
		CharDisplay{
			char: "A",
		},
		StringDisplay{
			str: "Hello, world.",
		},
		StringWithPaddingDisplay{
			str: "Goodbye, world.",
		},
	}

	for _, p := range printables {
		Print(p)

		// 改行
		println()
	}
}

// Printable5Times 5回繰り返すインターフェース
type Printable5Times interface {
	PrintBegin()
	PrintContent()
	PrintEnd()
}

// Print 5回繰り返す
// ***テンプレートメソッドです。***
func Print(d Printable5Times) {
	d.PrintBegin()
	for i := 0; i < 5; i++ {
		d.PrintContent()
	}
	d.PrintEnd()
}

// CharDisplay implements Printable5Times
type CharDisplay struct {
	char string
}

var _ Printable5Times = CharDisplay{}

// PrintBegin コンテンツの先頭に出力する。
func (c CharDisplay) PrintBegin() {
	print("<<")
}

// PrintContent コンテンツを出力する。
func (c CharDisplay) PrintContent() {
	print(c.char)
}

// PrintEnd コンテンツの最後に出力する
func (c CharDisplay) PrintEnd() {
	print(">>")
	println()
}

// StringDisplay implements Printable5Times
type StringDisplay struct {
	str string
}

var _ Printable5Times = StringDisplay{}

// PrintBegin コンテンツの先頭に出力する。
func (s StringDisplay) PrintBegin() {
	print(strings.Repeat("-", len(s.str)+2), "\n")
}

// PrintContent コンテンツを出力する。
func (s StringDisplay) PrintContent() {
	print("|"+s.str+"|", "\n")
}

// PrintEnd コンテンツの最後に出力する
func (s StringDisplay) PrintEnd() {
	print(strings.Repeat("-", len(s.str)+2), "\n")
}

// StringWithPaddingDisplay implements Printable5Times
type StringWithPaddingDisplay struct {
	str string
}

var _ Printable5Times = StringWithPaddingDisplay{}

// PrintBegin コンテンツの先頭に出力する。
func (m StringWithPaddingDisplay) PrintBegin() {
	print(strings.Repeat("*", len(m.str)+4), "\n")
	print("*", strings.Repeat(" ", len(m.str)+2), "*", "\n")
}

// PrintContent コンテンツを出力する。
func (m StringWithPaddingDisplay) PrintContent() {
	print("* ", m.str, " *", "\n")
}

// PrintEnd コンテンツの最後に出力する
func (m StringWithPaddingDisplay) PrintEnd() {
	print("*", strings.Repeat(" ", len(m.str)+2), "*", "\n")
	print(strings.Repeat("*", len(m.str)+4), "\n")
}
