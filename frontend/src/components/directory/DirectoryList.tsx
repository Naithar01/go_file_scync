import { RenameFileData } from "../../pages/MainPage";

import DirectoryTree from "./DirectoryTree";

type Props = {
  resFileData: RenameFileData[];
  type: string;
};

const DirectoryList = ({ resFileData, type }: Props) => {
  return (
    <div className={type == 'server' ? 'folderStructure' : 'connect_folderStructure'}>
      {resFileData.map((dirItem, index) => (
        <DirectoryTree key={index} name={dirItem.key} depth={dirItem.depth} files={dirItem.files} isLastDir={index === resFileData.length - 1} />
      ))}
    </div>
  );
};

export default DirectoryList;
