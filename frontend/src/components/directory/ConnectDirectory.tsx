import { RenameFileData } from "../../App"

type Props = {
  resFileData: RenameFileData[]
}

const ConnectDirectory = ({resFileData}: Props) => {
  return (
    <div className="connect_folderStructure">
      연결 된 PC의 정보 * 임시
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
              <i className="folder_icon"></i>{DirData.key}
            </div>
          </div>
        );
      })}
    </div>
  ) 
}

export default ConnectDirectory