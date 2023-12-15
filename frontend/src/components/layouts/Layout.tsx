import { Outlet } from "react-router-dom"

import Sidebar from "./Sidebar"

import "../../styles/common/layout_style.css"

const Layout = () => {
  return (
    <div className="layout">
      <main>
        <Outlet />
      </main>
    </div>
  )
}

export default Layout