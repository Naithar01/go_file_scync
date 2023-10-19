package tcpclient

import (
	"context"
	"encoding/json"
	"fmt"
	"go_file_sync/src/logs"
	"io"
	"net"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// TCPClient는 서버에 대한 클라이언트 연결을 관리합니다.
type TCPClient struct {
	ctx  *context.Context
	conn net.Conn
	ip   string
	port int
}

type Message struct {
	Type    string      `json:"type"`
	Content interface{} `json:"content"`
}

// NewTCPClient는 새 TCPClient 인스턴스를 생성합니다.
func NewTCPClient(ctx *context.Context) *TCPClient {
	return &TCPClient{
		ctx: ctx,
	}
}

func (c *TCPClient) connectToServer(ip string, port int) (net.Conn, error) {
	serverAddress := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.Dial("tcp", serverAddress)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

// 서버에 연결을 시도하고 클라이언트를 초기화
func (c *TCPClient) StartClient(ip string, port int) bool {
	c.ip = ip
	c.port = port

	conn, err := c.connectToServer(ip, port)
	if err != nil {
		runtime.MessageDialog(*c.ctx, runtime.MessageDialogOptions{
			Type:          runtime.ErrorDialog,
			Title:         "Error",
			Message:       "Could not connect to the server",
			Buttons:       nil,
			DefaultButton: "",
			CancelButton:  "",
		})
		return false
	}
	c.conn = conn

	// 클라이언트와 서버 연결 성공
	logs.PrintMsgLog("상대 PC 서버에 연결 성공")

	go c.ReceiveMessages() // 클라이언트가 메시지를 받을 수 있도록 고루틴 시작
	return true
}

func (c *TCPClient) handleMessage(buffer []byte, n int) {
	var message Message

	err := json.Unmarshal(buffer[:n], &message)
	if err != nil {
		runtime.MessageDialog(*c.ctx, runtime.MessageDialogOptions{
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
	case "close server":
		logs.PrintMsgLog("상대 PC로부터 연결 해제 - 서버 종료")
		runtime.EventsEmit(*c.ctx, "server_shutdown", true)
		c.conn.Close()
	}
}

func (c *TCPClient) ReceiveMessages() {
	for c.conn != nil {
		buffer := make([]byte, 1024)

		n, err := c.conn.Read(buffer)
		if err != nil {
			if err != io.EOF {
				logs.PrintMsgLog(fmt.Sprintf("메시지 받기 실패 에러: %s\n", err.Error()))
			}
			return
		}

		c.handleMessage(buffer, n)
	}
}

// 클라이언트가 서버에 연결한 이후 서버에 현재 PC에서 실행중인 포트를 보냄
func (c *TCPClient) SendAutoConnectServer(port int) {
	message := Message{
		Type:    "auto connect",
		Content: port,
	}

	// JSON 직렬화
	writeData, err := json.Marshal(message)
	if err != nil {
		runtime.MessageDialog(*c.ctx, runtime.MessageDialogOptions{
			Type:          runtime.ErrorDialog,
			Title:         "Error",
			Message:       "데이터가 올바르지 않습니다.",
			Buttons:       nil,
			DefaultButton: "",
			CancelButton:  "",
		})
	}

	_, err = c.conn.Write(writeData)
	if err != nil {
		logs.PrintMsgLog(fmt.Sprintf("데이터 전송에 실패하였습니다.: %s\n", err.Error()))
	}
}
