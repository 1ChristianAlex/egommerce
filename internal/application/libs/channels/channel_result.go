package channels

type Resolve[T interface{}] struct {
	Result T
	Err    error
}
