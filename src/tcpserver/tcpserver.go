package tcpserver

import (
	"context"
	"fmt"
	"net"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type TCPServer struct {
	ctx      *context.Context
	port     int
	listener net.Listener
}

func NewTCPServer(ctx *context.Context) *TCPServer {
	return &TCPServer{
		ctx: ctx,
	}
}

// TCP 서버 실행 성공 시에 True, 중복되는 PORT 사용 시에는 False, 에러 문구
func (t *TCPServer) SetServerPort(port int) bool {
	t.port = port

	err := t.startServer()

	if err != nil {
		runtime.MessageDialog(*t.ctx, runtime.MessageDialogOptions{
			Type:          runtime.ErrorDialog,
			Title:         "Error",
			Message:       "Port is already use",
			Buttons:       nil,
			DefaultButton: "",
			CancelButton:  "",
		})
		return false
	}

	return true
}

func (t *TCPServer) startServer() error {
	listenAddress := fmt.Sprintf(":%d", t.port)
	listener, err := net.Listen("tcp", listenAddress)
	if err != nil {
		return err
	}
	t.listener = listener

	fmt.Printf("TCP server is listening on port %d\n", t.port)

	go t.acceptConnections()

	return nil
}

func (t *TCPServer) acceptConnections() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go t.handleConnection(conn)
	}
}

func (t *TCPServer) handleConnection(conn net.Conn) {
	defer conn.Close()

	// Handle the connection here, e.g., read and write data.
}

func (t *TCPServer) Close() {
	if t.listener != nil {
		t.listener.Close()
	}
}
