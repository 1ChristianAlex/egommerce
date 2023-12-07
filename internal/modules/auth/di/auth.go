package di

import (
	"khrix/egommerce/internal/modules/user/dto"

	"github.com/gin-gonic/gin"
)

type AuthHelper interface {
	ExtractClaimsFromContext(context *gin.Context) (*dto.UserOutputTokenDto, error)
	JwtMiddleware(context *gin.Context)
}
