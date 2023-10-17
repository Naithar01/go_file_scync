package tcpserver

import (
	"context"
	"fmt"
	"net"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// TCPServer는 서버 및 클라이언트 연결을 관리합니다.
// 클라이언트 접속 유무는 client filed의 nil 확인
type TCPServer struct {
	ctx                  *context.Context
	port                 int
	listener             net.Listener
	serverListeningState bool
	client               net.Conn
}

// NewTCPServer는 새 TCPServer 인스턴스를 생성합니다.
func NewTCPServer(ctx *context.Context) *TCPServer {
	return &TCPServer{
		ctx:                  ctx,
		listener:             nil,
		serverListeningState: false,
		client:               nil,
	}
}

// GetPort는 현재 설정된 포트를 반환합니다.
func (t *TCPServer) GetPort() int {
	return t.port
}

// ReStartServer는 실행 중인 서버를 종료하고 앱을 다시로드합니다.
func (t *TCPServer) ReStartServer() {
	t.CloseServerAndDisconnectClient()
	runtime.WindowReload(*t.ctx)
}

// SetServerPort는 서버를 지정된 포트에서 실행합니다.
func (t *TCPServer) SetServerPort(port int) bool {
	t.port = port

	err := t.startServer()

	if err != nil {
		runtime.MessageDialog(*t.ctx, runtime.MessageDialogOptions{
			Type:          runtime.ErrorDialog,
			Title:         "Error",
			Message:       "Port is already in use",
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

// startServer는 서버를 시작합니다.
func (t *TCPServer) startServer() error {
	listenAddress := fmt.Sprintf(":%d", t.port)
	listener, err := net.Listen("tcp", listenAddress)
	if err != nil {
		return err
	}
	t.listener = listener
	fmt.Printf("TCP server is listening on port %d\n", t.port)

	go t.acceptConnections()

	runtime.EventsOn(*t.ctx, "client_server_disconnect", func(optionalData ...interface{}) {
		t.client.Close()
		t.client = nil
		fmt.Println("Client Disconnect")
	})

	return nil
}

// acceptConnections는 클라이언트 연결을 수락하고 처리합니다.
func (t *TCPServer) acceptConnections() {
	conn, err := t.listener.Accept()
	if err != nil {
		return
	}

	// 이미 연결된 클라이언트가 있다면 새로운 클라이언트를 거절
	if t.client != nil {
		conn.Close()
		return
	}

	t.client = conn
	fmt.Println("Client Connected")

	// 클라이언트로부터 연결이 성공적으로 수락되면 View로 이벤트를 보냄
	runtime.EventsEmit(*t.ctx, "server_accept_success", true)
}

// Close는 현재 실행 중인 서버 및 클라이언트 연결을 모두 닫습니다.
func (t *TCPServer) CloseServerAndDisconnectClient() {
	if t.listener != nil {
		t.listener.Close()
	}

	if t.client != nil {
		t.client.Close()
	}
}

// 프로그램 종료 시에 종료 문구를 보내야지 됨
func (t *TCPServer) Shutdown(ctx context.Context) {
	fmt.Println("ShutDown Program")
	if t.client != nil {
		_, err := t.client.Write([]byte("close server"))
		if err != nil {
			fmt.Println("Error sending close signal:", err)
		}
		t.client.Close()
	}
}
