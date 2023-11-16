// // latest 속성 추가 함수
// function markLatestFiles(firstDir: any, secondDir: any) {
//   for (const key in firstDir.files) {
//     if (secondDir.files[key]) {
//       const firstFiles = firstDir.files[key];
//       const secondFiles = secondDir.files[key];

//       for (const firstFile of firstFiles) {
//         const duplicateFile = secondFiles.find(
//           (secondFile) => secondFile.filename === firstFile.filename
//         );

//         if (duplicateFile) {
//           // Marking duplicates by adding the 'duplication' property
//           firstFile.duplication = true;
//           duplicateFile.duplication = true;

//           // Determining the latest file based on filemodtime
//           const firstFileTime = new Date(firstFile.filemodtime).getTime();
//           const duplicateFileTime = new Date(duplicateFile.filemodtime).getTime();

//           if (firstFileTime > duplicateFileTime) {
//             firstFile.latest = 1;
//             duplicateFile.latest = 2;
//           } else if (firstFileTime < duplicateFileTime) {
//             firstFile.latest = 2;
//             duplicateFile.latest = 1;
//           } else {
//             firstFile.latest = 0;
//             duplicateFile.latest = 0;
//           }
//         }
//       }
//     }
//   }
// }

// // latest 속성 추가 함수 호출
// markLatestFiles(null, null);

// console.log("가공 후");
// console.log(Object.keys(FirstDirectory.files).length);
// console.log(Object.keys(SecondDirectory.files).length);

export let a= 2