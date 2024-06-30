package must

func Must(err error) {
	if err != nil {
		panic(err)
	}
}

func MustReturn[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}
