import { Fragment, useEffect, useState } from "react"

import Loading from "./components/common/Loading"
import Layout from "./components/layouts/Layout"

import "./styles/app_style.css"

import { OpenDirectory } from "../wailsjs/go/main/App"
import { main } from "../wailsjs/go/models"
import { file } from "../wailsjs/go/models"

interface RenameFileData {
  key: string;
  depth: number;
  files: file.File[];
}

function App() {
	const [isLoading ,setIsLoading] = useState<boolean>(true)
  const [projectData, setProjectData] = useState<main.ResponseFileStruct>()
  const [resFileData, setResFileData] = useState<RenameFileData[]>()

  useEffect(() => {
		FetchFileData()
  }, []);

  // 서버로부터 파일 정보와, 선택 된 폴더 경로를 받아옴
	const FetchFileData = async (): Promise<void> => {
		try {
			const res = await OpenDirectory();

			if (res.root_path.length == 0 || !res.root_path) {
				FetchFileData()
				return
			}
			setIsLoading(() => {
        return false
      });
      setProjectData(() => {
        return res
      })
      setResFileData(() => {
        return renameFile(res)
      })
      
		} catch (error) {
			console.error("Error fetching data:", error);
			FetchFileData()
		}
	}

  const renameFile = (fileData: main.ResponseFileStruct): RenameFileData[] => {
    let renameFileData: RenameFileData[] = []
    for (const path_key in fileData.files) {
      let renameKey = path_key.replace(fileData.root_path, "")

      renameKey = renameKey.replace("\\", "/")

      if (renameKey.startsWith("/")) {
        renameKey = renameKey.slice(1)
      }

      renameFileData.push({
        key: renameKey,
        depth: renameKey == "" ? 0: renameKey.split("/").length,
        // @ts-ignore
        files: fileData.files[path_key],
      })
    }
    
    return renameFileData
  }

	return (
		<Layout>
			{ isLoading ? 
			<Loading />
			: 
			<Fragment>
        <div className="main">
          <div className="folderStructure">
            {resFileData && resFileData.length && resFileData.map((DirData) => (
              <div className={`folder ${DirData.key}`} key={DirData.key}>
                <i className="folder_icon"></i> ({DirData.depth}) {DirData.key}: {DirData.files.length}
              </div>
            ))}
          </div>
          <div className="connect_folderStructure">
            연결 된 상대 PC 파일 정보들...
          </div>
        </div>
			</Fragment>}
		</Layout>
	)
}

export default App
