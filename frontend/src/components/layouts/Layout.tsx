import Sidebar from "./Sidebar"

import "../../styles/common/layout_style.css"

type Props = {
  children: React.ReactNode
}

const Layout = ({children}: Props) => {
  return (
    <div className="layout">
      <Sidebar />
      {children}
    </div>
  )
}

export default Layout