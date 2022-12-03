package logic

import (
	"context"
	"errors"
	"time"

	"cloud-disk/core/define"
	"cloud-disk/core/helper"
	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"
	"cloud-disk/core/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type MailCodeSendRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMailCodeSendRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MailCodeSendRegisterLogic {
	return &MailCodeSendRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MailCodeSendRegisterLogic) MailCodeSendRegister(req *types.MailCodeSendRequest) (resp *types.MailCodeSendReply, err error) {
	// todo: add your logic here and delete this line
	// email 不存在
	cnt, err := l.svcCtx.Engine.Where("email = ?", req.Email).Count(new(models.UserBasic))
	if err != nil {
		return
	}
	if cnt > 0 {
		err = errors.New("该邮箱已经注册")
		return
	}
	//获取验证码
	code := helper.RandCode()
	//存储验证码
	l.svcCtx.RDB.Set(l.ctx, req.Email, code, time.Second*time.Duration(define.CodeExpire))
	//发送验证码
	err = helper.MailSendCode(req.Email, code)
	if err != nil {
		return nil, err
	}
	return
}
