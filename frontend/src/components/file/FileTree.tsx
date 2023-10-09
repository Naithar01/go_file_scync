type Props = {
  name: string; 
  depth: number; 
  isLastFile: boolean;
}

const FileTree = ({ name, depth, isLastFile }: Props) => {
  return (
    <div className="file" >
      {depth > 0 && (
        <span className="vertical-line">{isLastFile ? "└" : "├"}</span>
      )}
      <span className="vertical-line">{"─".repeat(depth * 1)}</span>
      📄 {name}
    </div>
  )
} 

export default FileTree