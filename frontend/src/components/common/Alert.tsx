import "../../styles/common/alert_style.css"

type Props = {
  text: string
}

const Alert = ({text}: Props) => {
  return (
    <div className="page_alert">
      {text}
    </div>
  )
}

export default Alert