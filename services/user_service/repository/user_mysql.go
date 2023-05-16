package repository

import (
	"context"
	"example/model"
)

type user struct {
}

func NewPostsRepository() *user {
	return &user{}
}
func (p *user) FetchAll(ctx context.Context, search string) ([]model.UserInfo, error) {
	// var result []model.UserInfo
	result := []model.UserInfo{
		{
			ID:   1,
			Name: "Aditya",
		},
	}

	// rows, er
	// if err != nil {
	// 	return nil, err
	// }

	// for rows.Next() {
	// 	var userInfo model.UserInfo
	// 	err = rows.Scan(&userInfo.ID, &userInfo.Name)
	// 	if err != nil {
	// 		return result, err
	// 	}

	// 	result = append(result, userInfo)
	// }

	return result, nil
}
