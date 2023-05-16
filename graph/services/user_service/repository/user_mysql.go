package repository

import (
	"context"
	"database/sql"
	"example/config"
	"example/graph/model"
)

type user struct {
	DB *sql.DB
}

func NewPostsRepository(conf config.Config) *user {
	return &user{conf.DB}
}
func (p *user) FetchAll(ctx context.Context, search string) ([]model.UserInfo, error) {
	var result []model.UserInfo
	rows, err := p.DB.Query("SELECT id,name FROM user where role like '%" + search + "%'")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var userInfo model.UserInfo
		err = rows.Scan(&userInfo.ID, &userInfo.Name)
		if err != nil {
			return result, err
		}

		result = append(result, userInfo)
	}

	return result, nil
}
