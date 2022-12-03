package logic

import (
	"context"
	"errors"
	"fmt"

	"cloud-disk/core/helper"
	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"
	"cloud-disk/core/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.LoginRequest) (resp *types.LoginReply, err error) {
	// todo: add your logic here and delete this line
	user := new(models.UserBasic)
	fmt.Println("user:", user)
	//1.数据库里查询当前用户
	has, err := models.Engine.Where("name = ? And password = ?", req.Name, helper.Md5(req.Password)).Get(user)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("用户名或密码错误")
	}
	//2.如果查询到，返回用户信息
	//3.生产token
	token, err := helper.GenerateToken(user.Id, user.Identity, user.Name)
	if err != nil {
		return nil, err

	}
	fmt.Println(token)
	resp = new(types.LoginReply)
	resp.Token = token
	fmt.Println("user:", user)

	return
}
