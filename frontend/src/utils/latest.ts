import { models } from "../../wailsjs/go/models";

// latest는 integer 타입
// 0은 겹치는 파일 명이 없음
// 1은 겹치는 파일 중에 최신 파일명
// 2는 겹치는 파일 중에 최신 파일명이 아님.
// 최신 파일의 여부는 filemodtime를 기준으로 함
// 비교를 할 때는 duplicateFile가 True인 파일 끼리 비교
export const markLatests = (firstDir: models.ResponseFileStruct, secondDir: models.ResponseFileStruct) => {
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

          if (duplicateFile && file1.duplication && duplicateFile.duplication) {
            // 중복된 파일이면 filemodtime 비교하여 최신 파일 표시
            if (file1.filemodtime > duplicateFile.filemodtime) {
              file1.latest = 1;
              duplicateFile.latest = 2;
            } else if (file1.filemodtime < duplicateFile.filemodtime) {
              file1.latest = 2;
              duplicateFile.latest = 1;
            } else {
              // 동일한 경우(시간이 같은 경우) latest를 0으로 설정
              file1.latest = 0;
              duplicateFile.latest = 0;
            }
          }
        });
      }
    }
  }
}