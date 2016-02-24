package view

// todo: docs
type Rect struct {
	X0, Y0, X1, Y1 int
}

func NewRect(X0, Y0, X1, Y1 int) *Rect {
	return &Rect{X0: X0, Y0: Y0, X1: X1, Y1: Y1}
}

func (r *Rect) Width() int {
	return r.X1 - r.X0
}

func (r *Rect) Height() int {
	return r.Y1 - r.Y0
}
