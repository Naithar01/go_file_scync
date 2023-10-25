type Props = {
  name: string; 
  depth: number; 
  isLastFile: boolean;
  dir_path: string;
}

const FileTree = ({ name, depth, isLastFile, dir_path }: Props) => {
  return (
    <div className={`file ${(depth == 1 ? name : dir_path + "/" + name)}`} >
      {depth > 0 && (
        <span className="vertical-line">{isLastFile ? "└" : "├"}</span>
      )}
      <span className="vertical-line">{"─".repeat(depth * 1)}</span>
      📄 {name}
    </div>
  )
} 

export default FileTree