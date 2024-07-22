package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"

	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UsersModel = (*customUsersModel)(nil)

type (
	// UsersModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUsersModel.
	UsersModel interface {
		usersModel
		AllUser(ctx context.Context) ([]*Users, error)
		FindOneByName(ctx context.Context, name string) (*Users, error)
	}

	customUsersModel struct {
		*defaultUsersModel
	}
)

func (c customUsersModel) FindOneByName(ctx context.Context, name string) (*Users, error) {
    usersIdKey := fmt.Sprintf("%s%v", cacheUsersIdPrefix, name)
	var resp Users
	err := c.QueryRowCtx(ctx, &resp, usersIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `name` = ? limit 1", usersRows, c.table)
		return conn.QueryRowCtx(ctx, v, query, name)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}


func (c customUsersModel) AllUser(ctx context.Context) ([]*Users, error) {
	query := fmt.Sprintf("select %s from %s", usersRows, c.table)

	var resp []*Users
	err := c.QueryRowsNoCacheCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

// NewUsersModel returns a model for the database table.
func NewUsersModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UsersModel {
	return &customUsersModel{
		defaultUsersModel: newUsersModel(conn, c, opts...),
	}
}
