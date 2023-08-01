package must

func Do(err error) {
	if err != nil {
		panic(err)
	}
}

func Do1[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}
	return t
}

func Do2[T, U any](t T, u U, err error) (T, U) {
	if err != nil {
		panic(err)
	}
	return t, u
}

func Do3[T, U, V any](t T, u U, v V, err error) (T, U, V) {
	if err != nil {
		panic(err)
	}
	return t, u, v
}

func Do4[T, U, V, W any](t T, u U, v V, w W, err error) (T, U, V, W) {
	if err != nil {
		panic(err)
	}
	return t, u, v, w
}

func Do5[T, U, V, W, X any](t T, u U, v V, w W, x X, err error) (T, U, V, W, X) {
	if err != nil {
		panic(err)
	}
	return t, u, v, w, x
}
