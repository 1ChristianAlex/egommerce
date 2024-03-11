package auth

import (
	"khrix/egommerce/internal/application/user/dto"

	"github.com/gin-gonic/gin"
)

type AuthHelper interface {
	ExtractClaimsFromContext(context *gin.Context) (*dto.UserOutputTokenDto, error)
	JwtMiddleware(context *gin.Context)
}
