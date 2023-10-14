import { Fragment, useEffect, useState } from "react"

import { InitialConnectServerPage } from "../../wailsjs/go/initial/Initial"

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
          <button type="button" onClick={() => {}}>서버 연결</button>
        </div>
      </div>
    </Fragment>
  )
}

export default ConnectServerPage