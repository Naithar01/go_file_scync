package tcpserver

import (
	"context"
	"encoding/json"
	"fmt"
	"go_file_sync/src/logs"
	"io"
	"net"
	"sync"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// TCPServer는 서버 및 클라이언트 연결을 관리합니다.
type TCPServer struct {
	ctx *context.Context
	m   sync.Mutex

	port     int
	listener net.Listener
	client   net.Conn
}

type Message struct {
	Type    string      `json:"type"`
	Content interface{} `json:"content"`
}

// NewTCPServer는 새 TCPServer 인스턴스를 생성합니다.
func NewTCPServer(ctx *context.Context) *TCPServer {
	return &TCPServer{
		ctx: ctx,
	}
}

// 현재 설정된 포트를 반환합니다.
func (t *TCPServer) GetPort() int {
	return t.port
}

// 실제 서버를 실행하는 부분
func (t *TCPServer) bindServer() error {
	listenAddress := fmt.Sprintf(":%d", t.port)
	listener, err := net.Listen("tcp", listenAddress)
	if err != nil {
		return err
	}
	t.listener = listener
	logs.PrintMsgLog(fmt.Sprintf("서버 실행중... 포트: %d\n", t.port))
	return nil
}

// startServer는 서버를 시작
func (t *TCPServer) startServer() error {
	if err := t.bindServer(); err != nil {
		return err
	}

	go t.acceptConnections()

	return nil
}

// 서버를 지정된 포트에서 실행합니다.
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

// 클라이언트 연결을 수락하고 처리
func (t *TCPServer) acceptConnections() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			continue
		}

		if t.client != nil {
			t.client.Close()
			t.client = conn
		}

		t.m.Lock()
		t.client = conn
		t.m.Unlock()
		logs.PrintMsgLog("상대 PC로부터 연결을 받음")

		// 클라이언트로부터 연결이 성공적으로 수락되면 View로 이벤트를 보냄
		runtime.EventsEmit(*t.ctx, "server_accept_success", true)

		// 클라이언트 연결 끊김 이벤트 핸들러 등록
		runtime.EventsOn(*t.ctx, "server_shutdown", func(_ ...interface{}) {
			if t.client != nil {
				t.client.Close()
			}
		})

		go t.ReceiveMessages()
	}
}

func (t *TCPServer) handleMessage(buffer []byte, n int) {
	var message Message

	err := json.Unmarshal(buffer[:n], &message)
	if err != nil {
		runtime.MessageDialog(*t.ctx, runtime.MessageDialogOptions{
			Type:          runtime.ErrorDialog,
			Title:         "Error",
			Message:       "데이터 수신에 실패하였습니다.",
			Buttons:       nil,
			DefaultButton: "",
			CancelButton:  "",
		})
		logs.PrintMsgLog(fmt.Sprintf("데이터 수신에 실패하였습니다.: %s\n", err.Error()))
		return
	}

	logs.PrintMsgLog(fmt.Sprintf("서버로부터 받은 헤더: %s\n", message.Type))
	switch message.Type {
	case "auto connect":
		logs.PrintMsgLog("상대 PC 자동 연결")
		runtime.EventsEmit(*t.ctx, "server_auto_connect", message.Content)
	}
}

func (t *TCPServer) ReceiveMessages() {
	for t.client != nil {
		buffer := make([]byte, 1024)

		n, err := t.client.Read(buffer)
		if err != nil {
			if err != io.EOF {
				logs.PrintMsgLog(fmt.Sprintf("메시지 받기 실패 에러: %s\n", err.Error()))
			}
			return
		}

		t.handleMessage(buffer, n)
	}
}

// 프로그램 종료 시에 종료 문구를 보냄
func (t *TCPServer) Shutdown(ctx context.Context) {
	t.m.Lock()
	defer t.m.Unlock()
	if t.client != nil {
		message := Message{
			Type:    "close server",
			Content: nil,
		}

		// JSON 직렬화
		writeData, err := json.Marshal(message)
		if err != nil {
			runtime.MessageDialog(*t.ctx, runtime.MessageDialogOptions{
				Type:          runtime.ErrorDialog,
				Title:         "Error",
				Message:       "데이터 전송에 실패하였습니다.",
				Buttons:       nil,
				DefaultButton: "",
				CancelButton:  "",
			})
			logs.PrintMsgLog(fmt.Sprintf("데이터 전송에 실패하였습니다.: %s\n", err.Error()))
		}

		_, err = t.client.Write(writeData)
		if err != nil {
			logs.PrintMsgLog(fmt.Sprintf("Error sending close signal: %s\n", err.Error()))
		}
		t.client.Close()
		t.listener.Close()
	}
}
