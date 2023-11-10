package auth

import (
	"errors"
	"net/http"
	"strings"

	"khrix/egommerce/internal/core/response"
	"khrix/egommerce/internal/modules/user/di"
	"khrix/egommerce/internal/modules/user/dto"

	"github.com/gin-gonic/gin"
)

type AuthHelper struct {
	jwtService di.JwtService
}

func NewAuthHelper(jwtService di.JwtService) *AuthHelper {
	return &AuthHelper{
		jwtService: jwtService,
	}
}

func (auth AuthHelper) abortRequest(context *gin.Context, err error) error {
	context.AbortWithStatusJSON(http.StatusUnauthorized, response.ResponseResult[interface{}]{
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
