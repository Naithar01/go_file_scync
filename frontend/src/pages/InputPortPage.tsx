import { Fragment, useEffect } from "react"

import { InitialInputPortPage } from "../../wailsjs/go/initial/Initial"
import { SetServerPort } from "../../wailsjs/go/tcpserver/TCPServer"

import styles from "../styles/pages/input_port_page_style.module.css" 
import { Link } from "react-router-dom"

const InputPortPage = () => {
  
  useEffect(() => {
    InitialInputPortPage()
    SetServerPort(1)
 
  }, [])
 
  return (
    <Fragment>
      <div className={styles.input_port_page}>
        Input Port Page
      </div>
    </Fragment>
  )
}

export default InputPortPage