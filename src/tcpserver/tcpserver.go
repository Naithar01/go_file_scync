package tcpserver

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go_file_sync/src/file"
	"go_file_sync/src/logs"
	"go_file_sync/src/models"
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
		logs.CustomErrorDialog(*t.ctx, "Port is already in use")
		return false
	}
	// logs.CustomInfoDialog(*t.ctx, "TCP server is listening")
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

// 클라이언트로부터 받은 메시지를 Type에 맞게 동작함
func (t *TCPServer) handleMessage(buffer []byte, n int) {
	var message models.Message

	err := json.Unmarshal(buffer[:n], &message)
	if err != nil {
		logs.PrintMsgLog(fmt.Sprintf("데이터 수신에 실패하였습니다.: %s\n", err.Error()))
		return
	}

	logs.PrintMsgLog(fmt.Sprintf("서버로부터 받은 헤더: %s\n", message.Type))
	switch message.Type {
	case "auto connect":
		var AutoConnect models.AutoConnect
		json.Unmarshal(buffer[:n], &AutoConnect)

		logs.PrintMsgLog("상대 PC 자동 연결")
		runtime.EventsEmit(*t.ctx, "server_auto_connect", AutoConnect.Content)
	}
}

// 클라이언트로부터 메시지를 받음
func (t *TCPServer) ReceiveMessages() {
	for t.client != nil {
		var buffer []byte
		tempBuffer := make([]byte, 1024) // Temporary buffer

		for t.client != nil {
			n, err := t.client.Read(tempBuffer)
			if err != nil {
				if err != io.EOF {
					logs.PrintMsgLog(fmt.Sprintf("메시지 받기 실패 에러: %s\n", err.Error()))
				}
				return
			}

			buffer = append(buffer, tempBuffer[:n]...)

			// Attempt to decode the received data as JSON
			var message models.Message
			decoder := json.NewDecoder(bytes.NewReader(buffer))
			if err := decoder.Decode(&message); err == nil {
				// Successfully decoded a JSON object
				t.handleMessage(buffer[:decoder.InputOffset()], len(buffer[:decoder.InputOffset()]))
				// Trim processed data from the buffer
				buffer = buffer[decoder.InputOffset():]
			}
		}
	}
}

// 선택 된 폴더의 내용을 클라이언트한테 보냄
func (t *TCPServer) SendDirectoryContent(files models.ResponseFileStruct) {
	if t.client == nil {
		return
	}

	t.m.Lock()
	defer t.m.Unlock()

	message := models.Message{
		Type:    "directory",
		Content: files,
	}

	// JSON 직렬화
	writeData, err := json.Marshal(message)
	if err != nil {
		logs.CustomErrorDialog(*t.ctx, "데이터 전송에 실패하였습니다.")
		logs.PrintMsgLog(fmt.Sprintf("데이터 전송에 실패하였습니다.: %s\n", err.Error()))
	}

	_, err = t.client.Write(writeData)
	if err != nil {
		logs.PrintMsgLog(fmt.Sprintf("Error sending close signal: %s\n", err.Error()))
	}
}

// 프로그램 종료 시에 종료 문구를 보냄
func (t *TCPServer) Shutdown(ctx context.Context) {
	if t.client == nil {
		return
	}

	t.m.Lock()
	defer t.m.Unlock()

	message := models.Message{
		Type:    "close server",
		Content: nil,
	}

	// JSON 직렬화
	writeData, err := json.Marshal(message)
	if err != nil {
		logs.CustomErrorDialog(*t.ctx, "데이터 전송에 실패하였습니다.")
		logs.PrintMsgLog(fmt.Sprintf("데이터 전송에 실패하였습니다.: %s\n", err.Error()))
	}

	_, err = t.client.Write(writeData)
	if err != nil {
		logs.PrintMsgLog(fmt.Sprintf("Error sending close signal: %s\n", err.Error()))
	}
	t.client.Close()
	t.listener.Close()
	t.client = nil
	t.listener = nil
}

// 선택한 파일을 클라이언트한테 보냄
func (t *TCPServer) SendFile(file_path string, file_name string) error {
	if t.client == nil {
		return nil
	}

	t.m.Lock()
	defer t.m.Unlock()

	dialog, err := runtime.MessageDialog(*t.ctx, runtime.MessageDialogOptions{
		Type:    runtime.QuestionDialog,
		Title:   "파일 전송",
		Message: "선택한 파일을 전송하시겠습니까?",
	})
	if err != nil {
		return errors.New(err.Error())
	}

	// "No" 클릭 시에 팝업 종료
	if dialog != "Yes" {
		return nil
	}

	file_content, err := file.ReadFile(file_path)
	if err != nil {
		return err
	}

	message := models.FileData{
		Type: "file",
		Content: models.ReadFile{
			FileName: file_name,
			FileData: file_content,
		},
	}

	// JSON 직렬화
	writeData, err := json.Marshal(message)
	if err != nil {
		logs.CustomErrorDialog(*t.ctx, "데이터 전송에 실패하였습니다.")
		logs.PrintMsgLog(fmt.Sprintf("데이터 전송에 실패하였습니다.: %s\n", err.Error()))
	}

	_, err = t.client.Write(writeData)
	if err != nil {
		logs.PrintMsgLog(fmt.Sprintf("Error sending close signal: %s\n", err.Error()))
	}

	return nil
}
