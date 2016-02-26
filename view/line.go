package view

import "github.com/nsf/termbox-go"

/*

note:
(rune index)

'┌'

'┐'

'└'

'┘'

'┼'

'├'

'┤'

'┬'

'┴'

'─'

'│'

'┤'
'┈'
'║'
╢
╡
╞
═

╠
*/

const LINE_CHR_VERTICAL = '│'
const LINE_CHR_HORIZONTAL = '─'
const LINE_CHR_DBLHORIZONTAL = '═'

// todo: docs
func lineVertical(x, y0, y1 int) {
	fg, bg := lineColor()

	for i := y0; i < y1; i++ {
		termbox.SetCell(x, i, LINE_CHR_VERTICAL, fg, bg)
	}
}

// todo: docs
func lineHorizontal(y, x0, x1 int) {
	fg, bg := lineColor()

	for i := x0; i <= x1; i++ {
		termbox.SetCell(i, y, LINE_CHR_HORIZONTAL, fg, bg)
	}
}

// todo: docs
func lineDblHorizontal(y, x0, x1 int) {
	fg, bg := lineColor()

	for i := x0; i <= x1; i++ {
		termbox.SetCell(i, y, LINE_CHR_DBLHORIZONTAL, fg, bg)
	}
}
