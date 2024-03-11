package auth

type AuthController[C any] interface {
	DoLogin(context *C)
}
