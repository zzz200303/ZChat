package models

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
)

var _ GroupsModel = (*customGroupsModel)(nil)

type (
	// GroupsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGroupsModel.
	GroupsModel interface {
		groupsModel
		Trans(ctx context.Context, fn func(context.Context, sqlx.Session) error) error
		ListByGroupIds(ctx context.Context, ids []string) ([]*Groups, error)
		InsertWithSession(ctx context.Context, session sqlx.Session, groups *Groups) (sql.Result, error)
	}

	customGroupsModel struct {
		*defaultGroupsModel
	}
)

func (c customGroupsModel) InsertWithSession(ctx context.Context, session sqlx.Session, groups *Groups) (sql.Result, error) {
	groupsIdKey := fmt.Sprintf("%s%v", cacheGroupsIdPrefix, groups.Id)
	ret, err := c.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", c.table, groupsRowsExpectAutoSet)
		return session.ExecCtx(ctx, query, groups.Id, groups.Name, groups.Icon, groups.Status, groups.CreatorUid, groups.GroupType, groups.IsVerify, groups.Notification, groups.NotificationUid, groups.CreatedAt, groups.UpdatedAt)
	}, groupsIdKey)
	return ret, err
}

func (c customGroupsModel) Trans(ctx context.Context, fn func(context.Context, sqlx.Session) error) error {
	return c.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})
}

func (c customGroupsModel) ListByGroupIds(ctx context.Context, ids []string) ([]*Groups, error) {
	query := fmt.Sprintf("select %s from %s where `id` in (?)", groupsRows, c.table)

	var resp []*Groups

	idStr := strings.Join(ids, "','")
	err := c.QueryRowsNoCacheCtx(ctx, &resp, query, idStr)

	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

// NewGroupsModel returns a model for the groupsbase table.
func NewGroupsModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) GroupsModel {
	return &customGroupsModel{
		defaultGroupsModel: newGroupsModel(conn, c, opts...),
	}
}
