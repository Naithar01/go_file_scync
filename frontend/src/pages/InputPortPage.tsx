import { Fragment, useEffect, useState } from "react"

import { InitialInputPortPage } from "../../wailsjs/go/initial/Initial"
import { SetServerPort } from "../../wailsjs/go/tcpserver/TCPServer"

import styles from "../styles/pages/input_port_page_style.module.css" 

const InputPortPage = () => {
  const [portState, setPortState] = useState<number>()

  useEffect(() => {
    InitialInputPortPage()
    // SetServerPort(1)
 
  }, [])

  const ChangePortStateHandler = (e: React.ChangeEvent<HTMLInputElement>): void => {
    const enteredPort: number = +e.target.value;
  
    setPortState(() => enteredPort);
  }

  const ConnectServerHandler = (): void => {
    alert(portState)
  }  

  return (
    <Fragment>
      <div className={styles.input_port_page}>
        <div className={styles.input_port_page_port_inp_areas}>
          <div className={styles.input_port_page_port_inp_area}>
            <input type="text" inputMode="numeric" placeholder="Enter Server Port" onChange={ChangePortStateHandler}/>
          </div>
          <div className={styles.input_port_page_port_inp_area}>
            <button type="button" onClick={ConnectServerHandler}>Start Server</button>
          </div>
        </div>
      </div>
    </Fragment>
  )
}

export default InputPortPage