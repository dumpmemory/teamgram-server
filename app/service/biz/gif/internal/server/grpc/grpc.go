/*
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package grpc

import (
	"github.com/teamgram/teamgram-server/app/service/biz/gif/gif"
	"github.com/teamgram/teamgram-server/app/service/biz/gif/internal/server/grpc/service"
	"github.com/teamgram/teamgram-server/app/service/biz/gif/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

// New new a grpc server.
func New(ctx *svc.ServiceContext, c zrpc.RpcServerConf) *zrpc.RpcServer {
	s, err := zrpc.NewServer(c, func(grpcServer *grpc.Server) {
		gif.RegisterRPCGifServer(grpcServer, service.New(ctx))
	})
	logx.Must(err)
	return s
}
