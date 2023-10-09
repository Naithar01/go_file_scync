import { file } from "../../../wailsjs/go/models"

type Props = {
  marginLeft: number 
  paddingLeft: number
  verticalLineHeight: number
  files: file.File[]
}

const FileList = ({marginLeft, paddingLeft, verticalLineHeight, files}: Props) => {
  return (
    <div className="folder_files_wrap" style={{ marginLeft, paddingLeft}}>
      { files && files.length > 0 && files.map((FileData) => {
        return (
          FileData.filename &&
          <div className="folder_files" key={FileData.filename}>
            {/* <div className="verticalLine" style={{ height: verticalLineHeight }}></div> */}
            <div className="file_header">
                <i className="file_header_icon"></i>{FileData.filename}
            </div>
          </div>
        )
      }) }
    </div>
  )
}

export default FileList