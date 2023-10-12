import { RenameFileData } from "../../pages/MainPage";

import DirectoryTree from "./DirectoryTree";

type Props = {
  resFileData: RenameFileData[];
};

const DirectoryList = ({ resFileData }: Props) => {
  return (
    <div className="folderStructure">
      {resFileData.map((dirItem, index) => (
        <DirectoryTree key={index} name={dirItem.key} depth={dirItem.depth} files={dirItem.files} isLastDir={index === resFileData.length - 1} />
      ))}
    </div>
  );
};

export default DirectoryList;
