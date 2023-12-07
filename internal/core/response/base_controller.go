package response

import (
	"net/http"

	"khrix/egommerce/internal/models"

	"github.com/gin-gonic/gin"
)

func ControllerInputMethod[Input interface{}, Output interface{}](
	context *gin.Context,
	input *Input,
	binder func(obj any) error,
	goFunciton func(channel chan models.Resolve[Output]),
) {
	if err := binder(input); err != nil {
		context.JSON(http.StatusBadRequest, &ResponseResult[*Output]{Result: nil, ErrorMessage: err.Error()})
		return
	}

	ControllerBaseMethod(context, goFunciton)
}

func ControllerBaseMethod[Output interface{}](context *gin.Context, goFunciton func(channel chan models.Resolve[Output])) {
	channel := make(chan models.Resolve[Output])
	defer close(channel)

	go goFunciton(channel)

	resolve := <-channel

	if resolve.Err != nil {
		context.JSON(http.StatusBadRequest, &ResponseResult[*Output]{Result: nil, ErrorMessage: resolve.Err.Error()})
		return
	}

	context.JSON(http.StatusOK, &ResponseResult[*Output]{
		Result: &resolve.Result,
	})
}
