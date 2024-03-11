package libs

type BaseMapper[InputDto any, OutputDto any, R any] interface {
	ToDto(item R) *OutputDto
	ToEntity(item InputDto) *R
}
