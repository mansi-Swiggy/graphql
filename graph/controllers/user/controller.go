package user

import (
	"context"
	"errors"
	"example/graph/model"
	"example/graph/services/user_service/repository"
	"example/model/response"
	"net/http"
	"strings"

	"github.com/vektah/gqlparser/gqlerror"
	"go.uber.org/zap"
)

type UserRunner interface {
	CreateUser(ctx context.Context, input model.UserInput) (*model.User, error)
	GetUser(ctx context.Context, request model.UserInput) (interface{}, error)
}
type userUsecase struct {
	repo repository.User
}

func NewPostUseCase(repo repository.User) *userUsecase {
	return &userUsecase{repo: repo}
}

func (oc *userUsecase) GetUser(ctx context.Context, request model.UserInput) (interface{}, error) {
	if len(strings.TrimSpace(request.UserRole)) == 0 {
		return response.NewHTTPResponse(http.StatusBadRequest, "invalid params", nil), errors.New("invalid params")
	}
	// Call the Get all  User method
	response, err := oc.repo.FetchAll(ctx, request.UserRole)
	zap.S().Infow("Get User Request",
		zap.Any("request", request),
		zap.Any("response", response),
		zap.Error(err),
	)
	if err != nil {
		return nil, gqlerror.Errorf("Error while parsing response ")
	}
	return response, nil
}
