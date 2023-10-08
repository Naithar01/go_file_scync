import { RenameFileData } from "../../App"

import { file } from "../../../wailsjs/go/models"

type Props = {
  resFileData: RenameFileData[]
}

const Folder = ({ name, depth, files }: { name: string; depth: number, files: file.File[] }) => (
  <div className="folder" style={{ marginLeft: `${depth * 10}px` }}>
    {depth > 0 ? <span className="vertical-line">│</span> : <span style={{paddingLeft: 10}}></span>} 📁 {name}
    { files.map((fileItem, index) => {
      return (
        fileItem.filename && <File key={index} name={fileItem.filename} depth={fileItem.depth} />
      )
    }) }
  </div>
);

const File = ({ name, depth }: { name: string; depth: number }) => (
  <div className="file" style={{ marginLeft: `${depth * 10}px` }}>
    {depth > 0 ? <span className="vertical-line">│</span> : <span style={{paddingLeft: 10}}></span>} 📄 {name}
  </div>
);

const DirectoryList = ({ resFileData }: Props) => {
  return (
    <div>
      {resFileData.map((dirItem, index) => {
          return (
            <Folder key={index} name={dirItem.key} depth={dirItem.depth} files={dirItem.files} />
          )
        })
      };
    </div>
  );
};

export default DirectoryList