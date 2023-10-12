import { Fragment, useEffect } from "react"

import { InitialInputPortPage } from "../../wailsjs/go/main/App"

import styles from "../styles/pages/input_port_page_style.module.css" 

const InputPortPage = () => {
  
  useEffect(() => {
    InitialInputPortPage()
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