import { Fragment, useEffect, useState } from "react"
import { useNavigate } from "react-router-dom"

import { InitialInputPortPage } from "../../wailsjs/go/initial/Initial"
import { SetServerPort } from "../../wailsjs/go/tcpserver/TCPServer"

import Alert from "../components/common/Alert"

import "../styles/pages/input_port_page_style.css" 

const InputPortPage = () => {
  const navigate = useNavigate()
  const [portState, setPortState] = useState<number>()

  useEffect(() => {
    InitialInputPortPage() 
  }, [])

  const ChangePortStateHandler = (e: React.ChangeEvent<HTMLInputElement>): void => {
    const enteredPort: number = +e.target.value;
  
    if (!Number.isNaN(enteredPort)) setPortState(() => enteredPort);
  }

  const StartServerHandler = async (): Promise<void> => {
    if (!portState) {
      return
    }

    const serverListeningState: boolean = await SetServerPort(portState)

    // 서버 실행 실패 시에...
    if (!serverListeningState) {
      return 
    }

    navigate("/connect")
    return
  }  

  return (
    <Fragment>
      <div className="input_port_page">
        <Alert text="서버를 실행하기 위한 PORT를 입력하세요." />
        <div className="input_port_page_port_inp_areas">
          <div className="input_port_page_port_inp_area">
            <input type="text" inputMode="numeric" value={portState} placeholder="포트를 입력하세요." onChange={ChangePortStateHandler}/>
          </div>
          <div className="input_port_page_port_inp_area">
            <button type="button" onClick={StartServerHandler}>서버 시작</button>
          </div>
        </div>
      </div>
    </Fragment>
  )
}

export default InputPortPage