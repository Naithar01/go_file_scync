import { RenameFileData } from "../../App"

import FileList from "../file/FileList"; 


type Props = {
  resFileData: RenameFileData[]
}
const DirectoryList = ({resFileData}: Props) => {
  return (
    <div className="folderStructure">
      {resFileData && resFileData.length > 0 && resFileData.map((DirData, index_st) => {
        const marginLeft =  DirData.depth > 1 ? DirData.depth * 9 : DirData.depth * 6;
        const paddingLeft = DirData.depth > 1 ? DirData.depth * 9 : DirData.depth * 6;
        const fileMarginLeft =  index_st == 0 ? 6 : DirData.depth * 6;
        const filePaddingLeft = index_st == 0 ? 6 : DirData.depth * 6;

        return (
          <div className={`folder_wrap ${DirData.depth}`} key={DirData.key} style={{ marginLeft, paddingLeft }}>
            {Array.from({ length: DirData.depth }, (_, index) => (
              <div className="verticalLine" key={index} style={{ height: 14, left: `${(index + 1) * 17}px` }}></div>
            ))}
            <div className={`folder ${DirData.key}`}>
              <div className="folder_header">
                <i className="folder_header_icon"></i>{DirData.key}
              </div>
              <FileList marginLeft={fileMarginLeft} paddingLeft={filePaddingLeft} fileDepth={DirData.depth} files={DirData.files} />
            </div>
          </div>
        );
      })}
    </div>
  ) 
}

export default DirectoryList