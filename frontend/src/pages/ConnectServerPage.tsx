import { Fragment, useEffect, useState } from "react"

import { CustomErrorDialog } from "../../wailsjs/go/main/App"
import { InitialConnectServerPage } from "../../wailsjs/go/initial/Initial"
import { StartClient } from "../../wailsjs/go/tcpclient/TCPClient"
import { GetPort } from "../../wailsjs/go/tcpserver/TCPServer"

import Alert from "../components/common/Alert"

import "../styles/pages/connect_server_page_style.css"

const ConnectServerPage = () => {
  const [portState, setPortState] = useState<number>()

  useEffect(() => {
    InitialConnectServerPage()
  }, [])

  const ChangePortStateHandler = (e: React.ChangeEvent<HTMLInputElement>): void => {
    const enteredPort: number = +e.target.value;
  
    if (!Number.isNaN(enteredPort)) setPortState(() => enteredPort);
  }

  // 현재 실행 중인 서버의 Port는 접속 불가능 하게
  // 연결 성공 시에 폴더 구조 페이지로 이동
  const StartClientHandler = async () => {
    if (await GetPort() == portState) {
      CustomErrorDialog("현재 PC에서 실행 중인 서버에 접속할 수 없습니다.")
    }

    
  }

  return (
    <Fragment>
      <div className="connect_server_page">
        <Alert text="연결할 PC의 포트를 입력하세요." />
      </div>
      <div className="connect_server_page_port_inp_areas">
        <div className="connect_server_page_port_inp_area">
          <input type="text" inputMode="numeric" value={portState} placeholder="포트를 입력하세요." onChange={ChangePortStateHandler}/>
        </div>
        <div className="connect_server_page_port_inp_area">
          <button type="button" onClick={StartClientHandler}>서버 연결</button>
        </div>
      </div>
    </Fragment>
  )
}

export default ConnectServerPage