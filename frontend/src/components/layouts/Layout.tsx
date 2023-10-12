import Sidebar from "./Sidebar"

import "../../styles/common/layout_style.css"
import { Outlet } from "react-router-dom"

const Layout = () => {
  return (
    <div className="layout">
      <Sidebar />
      <main>
        <Outlet />
      </main>
    </div>
  )
}

export default Layout