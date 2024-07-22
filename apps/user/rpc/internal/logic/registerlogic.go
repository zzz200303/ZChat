package logic

import (
	"ZChat/apps/user/model"
	"ZChat/pkg/ctxdata"
	"ZChat/pkg/encrypt"
	"context"
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-queue/kq"
	"log"
	"strconv"
	"time"

	"ZChat/apps/user/rpc/internal/svc"
	"ZChat/apps/user/rpc/pb/user"

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

var (
	ErrPhoneIsRegister = errors.New("手机号已经注册过")
)

func (l *RegisterLogic) Register(in *user.RegisterReq) (*user.RegisterResp, error) {
	// todo: add your logic here and delete this line

	// 1. 验证用户是否注册，根据手机号码验证
	userEntity, err := l.svcCtx.UsersModel.FindOneByName(l.ctx, in.Name)
	if err != nil && err != model.ErrNotFound {
		return nil, err
	}

	if userEntity != nil {
		return nil, ErrPhoneIsRegister
	}

	// 定义用户数据
	userEntity = &model.Users{
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

	fmt.Println(userEntity)

	sqlres, err := l.svcCtx.UsersModel.Insert(l.ctx, userEntity)
	if err != nil {
		return nil, err
	}

	// 生成token
	now := time.Now().Unix()
	token, err := ctxdata.GetJwtToken(l.svcCtx.Config.Jwt.AccessSecret, now, l.svcCtx.Config.Jwt.AccessExpire, userEntity.Id, userEntity.Name)
	if err != nil {
		return nil, err
	}

	//写入kafka消息，通知有新用户
	pusher := kq.NewPusher(l.svcCtx.Config.KqPusherConf.Brokers, l.svcCtx.Config.KqPusherConf.Topic)
	id, err := sqlres.LastInsertId()
	if err != nil {
		return nil, err
	}
	s := strconv.Itoa(int(id))
	if err := pusher.Push(s); err != nil {
		log.Fatal(err)
	}

	return &user.RegisterResp{
		Token:  token,
		Expire: now + l.svcCtx.Config.Jwt.AccessExpire,
	}, nil
}
