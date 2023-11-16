import { models } from "../../wailsjs/go/models";
import { RenameFileData } from "../pages/MainPage";

export const markDuplicates = (firstDir: RenameFileData, secondDir: RenameFileData) => {
  for (const key in firstDir.files) {
    if (secondDir.files[key]) {
      const firstFiles = firstDir.files[key];
      const secondFiles = secondDir.files[key];

        // @ts-ignore
      for (const firstFile of firstFiles) {
        // @ts-ignore
        const duplicateFile = secondFiles.find(
          (secondFile:models.File) => secondFile.filename === firstFile?.filename
        );

        if (duplicateFile) {
          // Marking duplicates by adding the 'duplication' property
          firstFile.duplication = true;
          duplicateFile.duplication = true;
        }
      }
    }
  }
}