package logic

import (
	"ZeZeIM/apps/user/models"
	"ZeZeIM/pkg/ctxdata"
	"ZeZeIM/pkg/encrypt"
	"ZeZeIM/pkg/wuid"
	"context"
	"database/sql"
	"github.com/pkg/errors"
	"time"

	"ZeZeIM/apps/user/rpc/internal/svc"
	"ZeZeIM/apps/user/rpc/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterReq) (*user.RegisterResp, error) {
	// todo: add your logic here and delete this line

	// 1. 验证用户是否注册，根据手机号码验证
	userEntity, err := l.svcCtx.UsersModel.FindByName(l.ctx, in.Name)
	if err != nil && !errors.Is(err, models.ErrNotFound) {
		return nil, err
	}

	if userEntity != nil {
		return nil, err
	}

	// 定义用户数据
	userEntity = &models.Users{
		Id:   wuid.GenUid(l.svcCtx.Config.Mysql.DataSource),
		Name: in.Name,
	}

	if len(in.Password) > 0 {
		genPassword, err := encrypt.GenPasswordHash([]byte(in.Password))
		if err != nil {
			return nil, err
		}
		userEntity.Password = sql.NullString{
			String: string(genPassword),
			Valid:  true,
		}
	}

	_, err = l.svcCtx.UsersModel.Insert(l.ctx, userEntity)
	if err != nil {
		return nil, err
	}

	// 生成token
	now := time.Now().Unix()
	token, err := ctxdata.GetJwtToken(l.svcCtx.Config.Jwt.AccessSecret, now, l.svcCtx.Config.Jwt.AccessExpire,
		userEntity.Id)
	if err != nil {
		return nil, err
	}

	return &user.RegisterResp{
		Token:  token,
		Expire: now + l.svcCtx.Config.Jwt.AccessExpire,
	}, nil
}
