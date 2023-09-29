import { RenameFileData } from "../../App"

type Props = {
  resFileData: RenameFileData[]
}

const Directory = ({resFileData}: Props) => {
  console.log(resFileData);
  
  return (
    <div className="folderStructure">
      {resFileData && resFileData.length > 0 && resFileData.map((DirData) => {
        const marginLeft = `${DirData.depth * 6}px`;
        const paddingLeft = `${DirData.depth * 6}px`;
        const verticalLineHeight = `${((resFileData.filter((fileStr) => fileStr.key.includes(DirData.key) && fileStr.key.startsWith(DirData.key))).length - 1) * 20 + 10}px`;

        return (
          <div className="folder_wrap" key={DirData.key} style={{ marginLeft, paddingLeft }}>
            {DirData.depth > 0 && (
              <div className="verticalLine" style={{ height: verticalLineHeight }}></div>
            )}
            <div className={`folder ${DirData.key}`}>
              <div className="folder_header">
                <i className="folder_header_icon"></i>{DirData.key}
              </div>
              { DirData.files && DirData.files.length > 0 && DirData.files.map((FielData) => {
                return (
                  <div className="folder_files_wrap">
                    
                  </div>
                )
              }) }
            </div>
          </div>
        );
      })}
    </div>
  ) 
}

export default Directory