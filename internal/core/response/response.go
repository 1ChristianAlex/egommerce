package response

type ResponseResult[T interface{}] struct {
	Result       T      `json:"result" `
	ErrorMessage string `json:"errorMessage" `
}
