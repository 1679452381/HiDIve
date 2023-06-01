// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"

	"hiDive-server/app/user/cmd/rpc/internal/logic"
	"hiDive-server/app/user/cmd/rpc/internal/svc"
	"hiDive-server/app/user/cmd/rpc/pb/user"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServer) Ping(ctx context.Context, in *user.Request) (*user.Response, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}

func (s *UserServer) Login(ctx context.Context, in *user.LoginReq) (*user.LoginResp, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}

// rpc register(LoginReq) returns(LoginResp);
func (s *UserServer) GenerateToken(ctx context.Context, in *user.GenerateTokenReq) (*user.GenerateTokenResp, error) {
	l := logic.NewGenerateTokenLogic(ctx, s.svcCtx)
	return l.GenerateToken(in)
}
