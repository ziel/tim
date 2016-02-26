package view

import "github.com/nsf/termbox-go"

// todo: docs
// todo: relocate?
func clearRegion(x0, x1, y0, y1 int) {
	clr := termbox.ColorDefault

	for i := x0; i <= x1; i++ {
		for j := y0; i <= y1; i++ {
			termbox.SetCell(i, j, ' ', clr, clr)
		}
	}
}

func clearCell(x, y int) {
	clr := termbox.ColorDefault
	termbox.SetCell(x, y, ' ', clr, clr)
}
