package nonempty

import "errors"

type NonEmpty[A any] struct {
	head A
	tail []A
}

func NewNonEmpty[A any](slice []A) (NonEmpty[A], error) {
	if len(slice) == 0 {
		return NonEmpty[A]{}, errors.New("empty slice")
	}

	return NonEmpty[A]{
		head: slice[0],
		tail: slice[1:],
	}, nil
}

func (ne NonEmpty[_]) Length() int {
	return 1 + len(ne.tail)
}
