import { file } from "../../../wailsjs/go/models";

import FileList from "../file/FileList";

type Props = {
  name: string;
  depth: number;
  files: file.File[];
  isLastDir: boolean;
}

const DirectoryTree = ({ name, depth, files, isLastDir }: Props) => {
  return (
  <div className="folder" >
    {depth > 0 && (
      <span className="vertical-line">{isLastDir ? "└" : "├"}</span>
    )}
    <span className="vertical-line">{"─".repeat(depth * 1)}</span>
    📁 {name}
    <FileList files={files} depth={depth} />
  </div>
  )
}

export default DirectoryTree