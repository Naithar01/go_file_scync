package tcpserver

import (
	"context"
	"fmt"
	"net"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// netstat -tuln

// serverListeningState: 서버 실행 유무
// clientListeningState: 클라이언트 접속 유무
type TCPServer struct {
	ctx                  *context.Context
	port                 int
	listener             net.Listener
	serverListeningState bool
	clientListeningState bool
}

func NewTCPServer(ctx *context.Context) *TCPServer {
	return &TCPServer{
		ctx:                  ctx,
		listener:             nil,
		serverListeningState: false,
		clientListeningState: false,
	}
}

func (t *TCPServer) GetPort() int {
	return t.port
}

// 실행되고 있는 서버 리스너 닫기, 앱 재실행
func (t *TCPServer) ReStartServer() {
	t.Close()
	runtime.WindowReload(*t.ctx)
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
	t.serverListeningState = true
	runtime.MessageDialog(*t.ctx, runtime.MessageDialogOptions{
		Type:          runtime.InfoDialog,
		Title:         "Server Start Success",
		Message:       "TCP server is listening",
		Buttons:       nil,
		DefaultButton: "",
		CancelButton:  "",
	})
	return true
}

func (t *TCPServer) startServer() error {
	listenAddress := fmt.Sprintf(":%d", t.port)
	listener, err := net.Listen("tcp", listenAddress)
	if err != nil {
		return err
	}
	t.listener = listener

	// fmt.Printf("TCP server is listening on port %d\n", t.port) ...Port로 서버 실행 중

	go t.acceptConnections()

	return nil
}

// 클라이언트 연결 수락
func (t *TCPServer) acceptConnections() {
	for !t.clientListeningState {
		conn, err := t.listener.Accept()
		if err != nil {
			// fmt.Println("Error accepting connection:", err) 클라이언트로부터 연결 받기 실패 오류
			continue
		}
		t.clientListeningState = true
		go t.handleConnection(conn)
	}
}

// 클라이언트 연결 수락 시에 어떤 로직을 할지,
func (t *TCPServer) handleConnection(conn net.Conn) {
	defer conn.Close()
}

func (t *TCPServer) Close() {
	if t.listener != nil {
		t.listener.Close()
	}
}
