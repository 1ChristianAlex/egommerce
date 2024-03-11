package base_controller

type ResponseResult[T interface{}] struct {
	Result       T      `json:"result" `
	ErrorMessage string `json:"errorMessage" `
}
