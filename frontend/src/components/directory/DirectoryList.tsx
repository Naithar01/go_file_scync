import { RenameFileData } from "../../App"

import FileList from "../file/FileList"; 


type Props = {
  resFileData: RenameFileData[]
}
const DirectoryList = ({resFileData}: Props) => {
  return (
    <div className="folderStructure">
      {resFileData && resFileData.length > 0 && resFileData.map((DirData, index_st) => {
        const marginLeft =  `${DirData.depth * 6}px`;
        const paddingLeft = `${DirData.depth * 6}px`;
        const fileMarginLeft =  index_st == 0 ? `6px` : `${DirData.depth * 6}px`;
        const filePaddingLeft = index_st == 0 ? `6px` : `${DirData.depth * 6}px`;
        const verticalLineHeight = `${((resFileData.filter((fileStr) => fileStr.key.includes(DirData.key) && fileStr.key.startsWith(DirData.key))).length - 1) * 20 + 10}px`;
        const fileVerticalLineHeight = `${((resFileData.filter((fileStr) => fileStr.key.includes(DirData.key) && fileStr.key.startsWith(DirData.key))).length - 1) * 11 + 10}px`;

        return (
          <div className="folder_wrap" key={DirData.key} style={{ marginLeft, paddingLeft }}>
            {/* {DirData.depth > 0 && (
              <div className="verticalLine" style={{ height: verticalLineHeight }}></div>
            )} */}
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