package nonempty

type NonEmpty[A any] struct {
	head A
	tail []A
}

func NewNonEmpty[A any](a A, slice []A) NonEmpty[A] {
	return NonEmpty[A]{
		head: a,
		tail: slice,
	}
}

func (ne NonEmpty[_]) Length() int {
	return 1 + len(ne.tail)
}
