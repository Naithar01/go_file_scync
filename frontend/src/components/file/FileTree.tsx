import { SendFile } from "../../../wailsjs/go/tcpserver/TCPServer";

type Props = {
  name: string; 
  depth: number; 
  isLastFile: boolean;
  dir_path: string;
  root_path: string;
}

const FileTree = ({ name, depth, isLastFile, dir_path, root_path }: Props) => {

  const SendFileHandler = (e: React.MouseEvent<HTMLDivElement>): void => {
    if (!e.currentTarget.className) {
      return 
    }
   
    // "file" 문자열을 제거한 파일 명 
    const file_name: string = e.currentTarget.className.slice(5)
    
    // 파일이 존재하는 경로 ( 최상위 폴더의 경우 Root와 Dir은 같기에 하나만 합침)
    const file_path: string = root_path == dir_path ? root_path + "/" + file_name : root_path + "/" + dir_path  + "/" + file_name

    SendFile(file_path)
  
  }
  
  return (
    // 첫 번째: "File" | 두 번째: Root 경로를 제외한 경로 | 세 번째: 파일이 존재하는 경로 
    <div className={`file ${(depth == 1 ? name : dir_path + "/" + name)}`} onClick={SendFileHandler} >
      {depth > 0 && (
        <span className="vertical-line">{isLastFile ? "└" : "├"}</span>
      )}
      <span className="vertical-line">{"─".repeat(depth * 1)}</span>
      📄 {name}
    </div>
  )
} 

export default FileTree 