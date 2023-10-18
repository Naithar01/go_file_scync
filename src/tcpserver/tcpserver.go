package tcpserver

import (
	"context"
	"fmt"
	"go_file_sync/src/logs"
	"net"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// TCPServer는 서버 및 클라이언트 연결을 관리합니다.
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
		ctx: ctx,
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
	logs.PrintMsgLog(fmt.Sprintf("서버 실행중... 포트: %d\n", t.port))

	go t.acceptConnections()

	return nil
}

// acceptConnections는 클라이언트 연결을 수락하고 처리합니다.
func (t *TCPServer) acceptConnections() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			continue
		}

		// 이미 연결된 클라이언트가 있다면 새로운 클라이언트를 거절
		if t.client != nil {
			logs.PrintMsgLog("클라이언트가 이미 연결 중")
			conn.Close()
			continue
		}

		t.client = conn
		logs.PrintMsgLog("클라이언트 연결 됨")

		// 클라이언트로부터 연결이 성공적으로 수락되면 View로 이벤트를 보냄
		runtime.EventsEmit(*t.ctx, "server_accept_success", true)

		// 클라이언트 연결 끊김 이벤트 핸들러 등록
		runtime.EventsOn(*t.ctx, "client_server_disconnect", func(_ ...interface{}) {
			t.client.Close()
			t.client = nil
			logs.PrintMsgLog("클라이언트 연결 끊김")
		})
	}
}

// CloseServerAndDisconnectClient는 현재 실행 중인 서버 및 클라이언트 연결을 모두 닫습니다.
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
	logs.PrintMsgLog("프로그램 종료")
	if t.client != nil {
		_, err := t.client.Write([]byte("close server"))
		if err != nil {
			logs.PrintMsgLog(fmt.Sprintf("Error sending close signal: %s\n", err.Error()))
		}
		t.client.Close()
		t.client = nil
	}
}
