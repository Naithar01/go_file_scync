type Props = {
  name: string; 
  depth: number; 
  isLastFile: boolean;
  dir_path: string;
  root_path: string;
}

const FileTree = ({ name, depth, isLastFile, dir_path, root_path }: Props) => {
  return (
    // 첫 번째: "File" | 두 번째: Root 경로를 제외한 경로 | 세 번째: 파일이 존재하는 경로 
    <div className={`file ${(depth == 1 ? name : dir_path + "/" + name)} ${(depth == 1 ? root_path + "/" + name : root_path + "/" + dir_path + "/" + name)} `} >
      {depth > 0 && (
        <span className="vertical-line">{isLastFile ? "└" : "├"}</span>
      )}
      <span className="vertical-line">{"─".repeat(depth * 1)}</span>
      📄 {name}
    </div>
  )
} 

export default FileTree 