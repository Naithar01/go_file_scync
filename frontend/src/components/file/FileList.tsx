import { Fragment } from "react";
import { models } from "../../../wailsjs/go/models";

import FileTree from "./FileTree";

type Props = {
  depth: number;
  files: models.File[];
  dir_path: string;
  root_path: string;
}

const FileList = ({files, depth, dir_path, root_path}: Props) => {
  return (
    <Fragment>
      {files.map((fileItem, index) => (
        fileItem.filename && <FileTree
          key={index}
          name={fileItem.filename}
          depth={depth + 1}
          isLastFile={index === files.length - 1}
          dir_path={dir_path}
          root_path={root_path}
          duplication={fileItem.duplication}
        />
      ))}
    </Fragment>
  )
}

export default FileList