import { useState } from "react";

import { file } from "../../../wailsjs/go/models";

import FileList from "../file/FileList";

type Props = {
  name: string;
  depth: number;
  files: file.File[];
  isLastDir: boolean;
}

const DirectoryTree = ({ name, depth, files, isLastDir }: Props) => {
  const [isExpanded, setExpanded] = useState(false);

  const toggleExpansion = () => {
    setExpanded(!isExpanded);
  };

  return (
  <div className={`folder ${name}`}>
    {depth > 0 && (
      <span className="vertical-line">
        {isLastDir ? "└" : "├"}
      </span>
    )}
    <span className="vertical-line">{"─".repeat(depth * 1)}</span>
    <span className="folder-icon" onClick={toggleExpansion}>
      {isExpanded ? "📂" : "📁"} {name}
    </span>
    {isExpanded && <FileList files={files} depth={depth} dir_path={name} />}
  </div>
  )
}

export default DirectoryTree