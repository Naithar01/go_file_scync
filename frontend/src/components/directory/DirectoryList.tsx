import { RenameFileData } from "../../pages/MainPage";

import DirectoryTree from "./DirectoryTree";

type Props = {
  resFileData: RenameFileData[];
  type: string;
  root_path: string;
};

const DirectoryList = ({ resFileData, type, root_path }: Props) => {
  return (
    <div className={type == 'server' ? 'folderStructure' : 'connect_folderStructure'}>
      {resFileData.map((dirItem, index) => (
        <DirectoryTree key={index} name={dirItem.key} depth={dirItem.depth} files={dirItem.files} isLastDir={index === resFileData.length - 1} root_path={root_path} />
      ))}
    </div>
  );
};

export default DirectoryList;
