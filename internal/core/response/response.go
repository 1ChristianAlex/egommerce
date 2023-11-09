package response

type ResponseResult[T any] struct {
	Result       T      `json:"result" `
	ErrorMessage string `json:"errorMessage" `
}
