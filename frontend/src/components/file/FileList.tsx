import { Fragment } from "react";
import { file } from "../../../wailsjs/go/models";

import FileTree from "./FileTree";

type Props = {
  depth: number
  files: file.File[]
}

const FileList = ({files, depth}: Props) => {
  return (
    <Fragment>
      {files.map((fileItem, index) => (
        fileItem.filename && <FileTree
          key={index}
          name={fileItem.filename}
          depth={depth + 1}
          isLastFile={index === files.length - 1}
        />
      ))}
    </Fragment>
  )
}

export default FileList