package auth

import (
	"errors"
	"net/http"
	"strings"

	base_controller "khrix/egommerce/internal/adapter/libs/base_controller"
	"khrix/egommerce/internal/application/user/dto"
	"khrix/egommerce/internal/port/auth"

	"github.com/gin-gonic/gin"
)

type AuthHelper struct {
	jwtService auth.JwtService
}

func NewAuthHelper(jwtService auth.JwtService) *AuthHelper {
	return &AuthHelper{
		jwtService: jwtService,
	}
}

func (auth AuthHelper) abortRequest(context *gin.Context, err error) error {
	context.AbortWithStatusJSON(http.StatusUnauthorized, base_controller.ResponseResult[interface{}]{
		ErrorMessage: err.Error(),
	})

	return err
}

func (auth AuthHelper) JwtMiddleware(context *gin.Context) {
	defer context.Next()

	tokenString, stringTokenErr := auth.getTokenHeader(context)

	if stringTokenErr != nil {
		return
	}

	if errValid := auth.jwtService.IsValid(tokenString); errValid != nil {
		auth.abortRequest(context, errValid)
		return
	}
}

func (auth AuthHelper) getTokenHeader(context *gin.Context) (string, error) {
	bearer := context.GetHeader("Authorization")

	if bearer == "" {
		return "", auth.abortRequest(context, errors.New("token not provied"))
	}

	jwtToken := strings.Split(bearer, " ")

	if len(jwtToken) != 2 {
		return "", auth.abortRequest(context, errors.New("incorrectly formatted authorization header"))
	}

	return jwtToken[1], nil
}

func (auth AuthHelper) ExtractClaimsFromContext(context *gin.Context) (*dto.UserOutputTokenDto, error) {
	stringToken, stringTokenErr := auth.getTokenHeader(context)

	if stringTokenErr != nil {
		return nil, stringTokenErr
	}

	parsedToken, parseError := auth.jwtService.FromClains(stringToken)

	if parseError != nil {
		return nil, errors.New("invalid token")
	}

	return parsedToken, nil
}
