package router

import (
	"github.com/VernHe/zinx/ziface"
	"github.com/VernHe/zinx/zlog"
	"github.com/VernHe/zinx/znet"
)

type HelloRouter struct {
	znet.BaseRouter
}

func (hr *HelloRouter) Handle(request ziface.IRequest) {
	zlog.Ins().InfoF(string(request.GetData()))
}
