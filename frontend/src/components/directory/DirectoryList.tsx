import { RenameFileData } from "../../App";
import { file } from "../../../wailsjs/go/models";

type Props = {
  resFileData: RenameFileData[];
};

//  "└─" : "├─"}
// 폴더1/
// ├─파일1.txt
// ├─파일2.txt
// └─하위폴더1/
// ├─서브파일1.txt
// └─서브파일2.txt
// ├─하위 폴더2/
// └─파일3.txt
// └──하위 하위 폴더1/

const Folder = ({
  name,
  depth,
  files,
  isLastDir,
}: {
  name: string;
  depth: number;
  files: file.File[];
  isLastDir: boolean;
}) => (
  <div className="folder" >
    {depth > 0 && (
      <span className="vertical-line">{isLastDir ? "└" : "├"}</span>
    )}
    <span className="vertical-line">{"─".repeat(depth * 1)}</span>
    📁 {name}
    {files.map((fileItem, index) => (
      fileItem.filename && <File
        key={index}
        name={fileItem.filename}
        depth={depth + 1}
        isLastFile={index === files.length - 1}
      />
    ))}
  </div>
);

const File = ({ name, depth, isLastFile }: { name: string; depth: number; isLastFile: boolean }) => (
  <div className="file" >
    {depth > 0 && (
      <span className="vertical-line">{isLastFile ? "└" : "├"}</span>
    )}
    <span className="vertical-line">{"─".repeat(depth * 1)}</span>
    📄 {name}
  </div>
);

const DirectoryList = ({ resFileData }: Props) => {
  return (
    <div className="folderStructure">
      {resFileData.map((dirItem, index) => (
        <Folder key={index} name={dirItem.key} depth={dirItem.depth} files={dirItem.files} isLastDir={index === resFileData.length - 1} />
      ))}
    </div>
  );
};

export default DirectoryList;
