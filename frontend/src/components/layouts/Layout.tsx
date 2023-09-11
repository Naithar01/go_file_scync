import "../../styles/common/layout_style.css"

type Props = {
  children: React.ReactNode
}

const Layout = ({children}: Props) => {
  return (
    <div className="layout">
      {children}
    </div>
  )
}

export default Layout