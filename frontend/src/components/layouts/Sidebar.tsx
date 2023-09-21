import { Fragment, useState } from "react"

import { FontAwesomeIcon } from "@fortawesome/react-fontawesome"
import { faBars } from "@fortawesome/free-solid-svg-icons"

import "../../styles/common/sidebar_style.css"

const Sidebar = () => {
  return (
    <Fragment>
      <nav className="sidebar_nav">
        <div className="sidebar_nav_toggle_icon">
          <button className="sidebar_nav_toggle_btn" type="button">
            <FontAwesomeIcon icon={faBars} />
          </button>
        </div>
        <p className="sidebar_nav_logo">폴더 동기화</p>
      </nav>
      <div className="sidebar">

      </div>
    </Fragment>
  )
}

export default Sidebar