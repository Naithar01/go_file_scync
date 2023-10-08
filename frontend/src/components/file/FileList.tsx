import { file } from "../../../wailsjs/go/models"

type Props = {
  marginLeft: number 
  paddingLeft: number
  fileDepth: number
  files: file.File[]
}

const FileList = ({marginLeft, paddingLeft, fileDepth, files}: Props) => {
  return (
    <div className="folder_files_wrap" style={{ marginLeft, paddingLeft}}>
      { files && files.length > 0 && files.map((FileData) => {
        return (
          FileData.filename &&
          <div className="folder_files" key={FileData.filename}>
            {Array.from({ length: fileDepth + 1 }, (_, index) => (
              <div className="verticalLine" key={index} style={{ height: 14, left: `${(index + 1) * 17}px` }}></div>
            ))}
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