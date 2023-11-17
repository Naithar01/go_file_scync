import { models } from "../../wailsjs/go/models";
import { RenameFileData } from "../pages/MainPage";

export const markDuplicates = (firstDir: models.ResponseFileStruct, secondDir: models.ResponseFileStruct) => {
  for (const key in firstDir.files) {
    const reNameFirstDirKey = key.replace(firstDir.root_path, "");
    for (const key_nd in secondDir.files) {
      const reNameSecondDirKey = key_nd.replace(secondDir.root_path, "");
      if (reNameFirstDirKey == reNameSecondDirKey) {
        // 해당 폴더 내의 파일 비교
        const filesInFirstDir = firstDir.files[key];
        const filesInSecondDir = secondDir.files[key_nd];

        filesInFirstDir.forEach(file1 => {
          const duplicateFile = filesInSecondDir.find(file2 => file1.filename === file2.filename);

          if (duplicateFile) {
            // 중복된 파일이면 duplication 속성을 true로 설정
            file1.duplication = true;
            duplicateFile.duplication = true;
          }
        });
      }
    }
  }
}