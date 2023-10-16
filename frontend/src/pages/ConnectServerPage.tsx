import { Fragment, useEffect, useState } from "react"
import { useNavigate } from "react-router-dom"


import { CustomErrorDialog } from "../../wailsjs/go/main/App"
import { InitialConnectServerPage } from "../../wailsjs/go/initial/Initial"
import { StartClient } from "../../wailsjs/go/tcpclient/TCPClient"
import { GetPort } from "../../wailsjs/go/tcpserver/TCPServer"
import { EventsOn } from "../../wailsjs/runtime/runtime"

import Alert from "../components/common/Alert"

import "../styles/pages/connect_server_page_style.css"
import Loading from "../components/common/Loading"

const ConnectServerPage = () => {
  const navigate = useNavigate()
  const [portState, setPortState] = useState<number>()
  // 상대 PC가 실행 중인 PC에 연결을 했는지...
  const [connectListeningIsLoading, setConnectListeningIsLoading] = useState<boolean>(false)
  const [acceptSuccessState, setAcceptSuccessState] = useState<Boolean>(false)

  useEffect(() => {
    InitialConnectServerPage()
  }, [])

  const ChangePortStateHandler = (e: React.ChangeEvent<HTMLInputElement>): void => {
    const enteredPort: number = +e.target.value;
  
    if (!Number.isNaN(enteredPort)) setPortState(() => enteredPort);
  }

  // 현재 실행 중인 서버의 Port는 접속 불가능 하게
  const StartClientHandler = async () => {
    if (await GetPort() == portState) {
      CustomErrorDialog("현재 PC에서 실행 중인 서버에 접속할 수 없습니다.")
      return
    }

    if (!portState) {
      return
    }

    const serverConnectState = await StartClient("127.0.0.1", portState)

    if (!serverConnectState) {
      return
    }

    setConnectListeningIsLoading(true)

    // 상대 PC로부터 연결을 받은 상태이며 서버를 연결 한다면...
    if (acceptSuccessState) {
      navigate("/dir")
    }
  }

  // 상대 PC 연결 이후에 로딩 상태라면...
  // 상대 PC 연결을 안 하고 연결을 받았다면...
  EventsOn("server_accept_success", () => {
    setAcceptSuccessState(true)
    if (connectListeningIsLoading) {
      navigate("/dir")
    } else {
      alert("Accept Success")
    }
  })

  return (
    <Fragment>
      { !connectListeningIsLoading ? 
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
      :
      <Fragment>
        <Loading />
      </Fragment>
      }
    </Fragment>
  )
}

export default ConnectServerPage