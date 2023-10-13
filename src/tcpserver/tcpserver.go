package tcpserver

import "context"

type TCPServer struct {
	ctx *context.Context
}

func NewTCPServer(ctx *context.Context) *TCPServer {
	return &TCPServer{
		ctx: ctx,
	}
}
