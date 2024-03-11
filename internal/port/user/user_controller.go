package user

type UserController[C any] interface {
	CreateNewUser(context *C)
}
