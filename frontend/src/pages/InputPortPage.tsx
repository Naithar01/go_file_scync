import { Fragment, useEffect } from "react"
import { Link } from "react-router-dom"

import { InitialInputPortPage } from "../../wailsjs/go/initial/Initial"
import { SetServerPort } from "../../wailsjs/go/tcpserver/TCPServer"

import styles from "../styles/pages/input_port_page_style.module.css" 

const InputPortPage = () => {
  useEffect(() => {
    InitialInputPortPage()
    // SetServerPort(1)
 
  }, [])
 
  const ConnectServer = (): void => {

  }  

  return (
    <Fragment>
      <div className={styles.input_port_page}>
        Input Port Page

        <Link to="/dir">Hel</Link>
      </div>
    </Fragment>
  )
}

export default InputPortPage