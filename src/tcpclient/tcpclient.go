package tcpclient

import (
	"context"
	"fmt"
	"net"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// connectState: 서버에 연결 되었다는 유무
type TCPClient struct {
	ctx          *context.Context
	conn         net.Conn
	ip           string
	port         int
	connectState bool
}

func NewTCPClient(ctx *context.Context) *TCPClient {
	return &TCPClient{
		ctx:          ctx,
		conn:         nil,
		connectState: false,
	}
}

// 클라이언트 연결 시작
func (c *TCPClient) StartClient(ip string, port int) bool {
	c.ip = ip
	c.port = port

	err := c.connectToServer()

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

	fmt.Println("서버에 연결 성공")
	c.connectState = true
	runtime.MessageDialog(*c.ctx, runtime.MessageDialogOptions{
		Type:          runtime.InfoDialog,
		Title:         "Connected",
		Message:       "Connected to the server",
		Buttons:       nil,
		DefaultButton: "",
		CancelButton:  "",
	})

	go c.ReceiveMessages() // 클라이언트가 메시지를 받을 수 있도록 고루틴 시작
	return true
}

func (c *TCPClient) connectToServer() error {
	serverAddress := fmt.Sprintf("%s:%d", c.ip, c.port)
	conn, err := net.Dial("tcp", serverAddress)
	if err != nil {
		return err
	}
	c.conn = conn

	// 클라이언트와 서버 연결 성공
	return nil
}

func (c *TCPClient) ReceiveMessages() {
	for {
		buffer := make([]byte, 1024)
		n, err := c.conn.Read(buffer)
		if err != nil {
			fmt.Println("메시지 받기 실패 에러: ", err)
			c.Close()
			return
		}

		message := string(buffer[:n])
		fmt.Println("받은 문구:", message)

		if message == "close server" {
			fmt.Println("서버 닫힘, 연결 끊기")
			runtime.EventsEmit(*c.ctx, "client_server_disconnect", true)
			c.Close()
		}
	}
}

func (c *TCPClient) Close() {
	if c.conn != nil {
		c.connectState = false
		c.conn.Close()
	}
}
