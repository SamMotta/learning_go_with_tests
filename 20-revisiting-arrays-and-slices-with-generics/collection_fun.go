package main

func Reduce[A, B any](collection []A, f func(B, A) B, initialValue B) (result B) {
	result = initialValue
	for _, x := range collection {
		result = f(result, x)
	}

	return
}

func Find[T any](items []T, predicate func(x T) bool) (value T, found bool) {
	for _, v := range items {
		if predicate(v) {
			return v, true
		}
	}
	return
}
