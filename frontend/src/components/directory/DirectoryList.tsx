import { RenameFileData } from "../../App"

import FileList from "../file/FileList"; 


type Props = {
  resFileData: RenameFileData[]
}
const DirectoryList = ({resFileData}: Props) => {
  return (
    <div className="folderStructure">
      {resFileData && resFileData.length > 0 && resFileData.map((DirData, index_st) => {
        const marginLeft =  DirData.depth * 6;
        const paddingLeft = DirData.depth * 6;
        const fileMarginLeft =  index_st == 0 ? 6 : DirData.depth * 6;
        const filePaddingLeft = index_st == 0 ? 6 : DirData.depth * 6;

        const matchDirs = resFileData.filter((fileStr) => fileStr.key != DirData.key && fileStr.key.includes(DirData.key) && fileStr.key.startsWith(DirData.key) && fileStr.depth != DirData.depth)

        const verticalLineHeight = (
          (matchDirs).length
        + matchDirs.concat(DirData).flatMap((fileStr) => fileStr.files.length).reduce((acc, currentValue) => acc + currentValue, 0)
        - matchDirs.concat(DirData).length) * 20 + 10;
        
        
        const fileVerticalLineHeight = (matchDirs.length) * 11 + 15;

        return (
          <div className="folder_wrap" key={DirData.key} style={{ marginLeft, paddingLeft }}>
            {DirData.depth > 0 && DirData.files && (
              <div className="verticalLine" style={{ height: verticalLineHeight }}></div>
            )}
            <div className={`folder ${DirData.key}`}>
              <div className="folder_header">
                <i className="folder_header_icon"></i>{DirData.key}
              </div>
              <FileList marginLeft={fileMarginLeft} paddingLeft={filePaddingLeft} verticalLineHeight={fileVerticalLineHeight} files={DirData.files} />
            </div>
          </div>
        );
      })}
    </div>
  ) 
}

export default DirectoryList