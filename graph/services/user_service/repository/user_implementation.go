package repository

import (
	"context"
	"example/graph/model"
)

type User interface {
	FetchAll(ctx context.Context, search string) ([]model.User, error)
}
