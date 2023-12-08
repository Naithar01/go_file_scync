package tcpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"go_file_sync/src/global"
	"go_file_sync/src/logs"
	"go_file_sync/src/models"
	"io"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
	"sync"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// TCPClient는 서버에 대한 클라이언트 연결을 관리합니다.
type TCPClient struct {
	ctx *context.Context
	m   sync.Mutex
	wg  sync.WaitGroup

	conn net.Conn
	ip   string
	port int
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
	if c.conn != nil {
		return false
	}

	c.m.Lock()
	c.ip = ip
	c.port = port

	conn, err := c.connectToServer(ip, port)
	if err != nil {
		logs.CustomErrorDialog(*c.ctx, "Could not connect to the server")
		return false
	}
	c.conn = conn
	c.m.Unlock()

	// 클라이언트와 서버 연결 성공
	logs.PrintMsgLog("상대 PC 서버에 연결 성공")

	go c.ReceiveMessages() // 클라이언트가 메시지를 받을 수 있도록 고루틴 시작
	return true
}

func (c *TCPClient) handleMessage(buffer []byte, n int) {
	var message models.Message

	err := json.Unmarshal(buffer[:n], &message)
	if err != nil {
		logs.PrintMsgLog(fmt.Sprintf("데이터 수신에 실패하였습니다.: %s\n", err.Error()))
		return
	}

	logs.PrintMsgLog(fmt.Sprintf("서버로부터 받은 헤더: %s\n", message.Type))
	switch message.Type {
	case "close server":
		var ShutdownMessage models.ServerShutdownMessage
		json.Unmarshal(buffer[:n], &ShutdownMessage)

		c.conn.Close()
		c.conn = nil
		logs.PrintMsgLog("상대 PC로부터 연결 해제 - 서버 종료")
		runtime.EventsEmit(*c.ctx, "server_shutdown", message.Content)
	case "directory":
		var DirectoryContent models.DirectoryContent
		json.Unmarshal(buffer[:n], &DirectoryContent)

		logs.PrintMsgLog("상대 PC로부터 폴더 정보를 받음")
		runtime.EventsEmit(*c.ctx, "connectedDirectoryData", DirectoryContent.Content)
	case "file":
		var FileData models.FileData
		json.Unmarshal(buffer[:n], &FileData)
		global.GetRootPath()

		root_path := global.GetRootPath()
		if len(root_path) == 0 {
			return
		}

		filePath := filepath.Join(root_path, FileData.Content.FileName)
		err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm)
		if err != nil {
			logs.PrintMsgLog(err.Error())
			return
		}

		err = ioutil.WriteFile(filePath, FileData.Content.FileData, 0644)
		if err != nil {
			logs.CustomErrorDialog(*c.ctx, "파일 수신에 실패하였습니다.")
			return
		}

		logs.PrintMsgLog("상대 PC로부터 파일 데이터를 받음")
		runtime.EventsEmit(*c.ctx, "receive_file", nil)
	case "start_sync_files":
		c.wg.Add(1)
		var FileDataInfo models.StartSyncFilesContent
		json.Unmarshal(buffer[:n], &FileDataInfo)
		// PC UI를 로딩 상태로 업데이트
		// PC 폴더 정보를 전송
		runtime.EventsEmit(*c.ctx, "start_sync_files", true)
		runtime.EventsEmit(*c.ctx, "start_together_sync_files", true)
		c.wg.Done()
		c.wg.Wait()
		c.SendStartFileEvent()
	case "start_together_sync_files":
		c.wg.Add(1)
		var FileDataInfo models.StartSyncFilesContent
		json.Unmarshal(buffer[:n], &FileDataInfo)
		c.wg.Done()
		c.wg.Wait()
		c.SendStartFileEvent()
	case "send_sync_file":
		var FileData models.FileData
		json.Unmarshal(buffer[:n], &FileData)
		fmt.Println(FileData.Content.FileName)
	}
}

func (c *TCPClient) ReceiveMessages() {
	for c.conn != nil {
		var buffer []byte
		tempBuffer := make([]byte, 1024) // Temporary buffer

		for c.conn != nil {
			n, err := c.conn.Read(tempBuffer)
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
				c.handleMessage(buffer[:decoder.InputOffset()], len(buffer[:decoder.InputOffset()]))
				// Trim processed data from the buffer
				buffer = buffer[decoder.InputOffset():]
			}
		}
	}
}

// 클라이언트가 서버에 연결한 이후 서버에 현재 PC에서 실행중인 포트를 보냄
func (c *TCPClient) SendAutoConnectServer(port int) {
	message := models.Message{
		Type: "auto connect",
		Content: models.AutoConnectContent{
			IP:   global.GetOutboundIP().String(),
			PORT: port,
		},
	}

	// JSON 직렬화
	writeData, err := json.Marshal(message)
	if err != nil {
		logs.CustomErrorDialog(*c.ctx, "데이터가 올바르지 않습니다.")
		return
	}

	_, err = c.conn.Write(writeData)
	if err != nil {
		logs.PrintMsgLog(fmt.Sprintf("데이터 전송에 실패하였습니다.: %s\n", err.Error()))
	}
}

func (c *TCPClient) SendStartFileEvent() {
	message := models.Message{
		Type:    "send_sync_files",
		Content: nil,
	}

	// JSON 직렬화
	writeData, err := json.Marshal(message)
	if err != nil {
		logs.CustomErrorDialog(*c.ctx, "데이터가 올바르지 않습니다.")
		return
	}

	_, err = c.conn.Write(writeData)
	if err != nil {
		logs.PrintMsgLog(fmt.Sprintf("데이터 전송에 실패하였습니다.: %s\n", err.Error()))
	}
}
