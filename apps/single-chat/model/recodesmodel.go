package model

//
//import (
//	"context"
//	"fmt"
//	"github.com/zeromicro/go-zero/core/stores/sqlc"
//	"github.com/zeromicro/go-zero/core/stores/sqlx"
//)
//
//var _ RecodesModel = (*customRecodesModel)(nil)
//
//type (
//	// RecodesModel is an interface to be customized, add more methods here,
//	// and implement the added methods in customRecodesModel.
//	RecodesModel interface {
//		recodesModel
//		withSession(session sqlx.Session) RecodesModel
//		SelectRecordList(ctx context.Context, uid, friendUid int64) ([]*Recodes, error)
//	}
//
//	customRecodesModel struct {
//		*defaultRecodesModel
//	}
//)
//
//func (m *customRecodesModel) SelectRecordList(ctx context.Context, uid, friendUid int64) ([]*Recodes, error) {
//	query := fmt.Sprintf("select %s from %s where `from` = ? and `to` = ?", recodesRows, m.table)
//	var resp []*Recodes
//	err := m.conn.QueryRowsCtx(ctx, &resp, query, uid, friendUid)
//	switch err {
//	case nil:
//		return resp, nil
//	case sqlc.ErrNotFound:
//		return nil, ErrNotFound
//	default:
//		return nil, err
//	}
//}
//
//// NewRecodesModel returns a model for the database table.
//func NewRecodesModel(conn sqlx.SqlConn) RecodesModel {
//	return &customRecodesModel{
//		defaultRecodesModel: newRecodesModel(conn),
//	}
//}
//
//func (m *customRecodesModel) withSession(session sqlx.Session) RecodesModel {
//	return NewRecodesModel(sqlx.NewSqlConnFromSession(session))
//}
