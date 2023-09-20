import { Fragment, useEffect, useState } from "react"

import Loading from "./components/common/Loading"
import Layout from "./components/layouts/Layout"

import "./styles/app_style.css"

import { OpenDirectory } from "../wailsjs/go/main/App"
import { main } from "../wailsjs/go/models"
import { file } from "../wailsjs/go/models"

function App() {
	const [isLoading ,setIsLoading] = useState<boolean>(true)
  const [projectData, setProjectData] = useState< main.ResponseFileStruct>()
  const [resFileData, setResFileData] = useState<{[key: string]: file.File[]}>()

  useEffect(() => {
		// FetchFileData()
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
        return renameFileData(res)
      })
      
		} catch (error) {
			console.error("Error fetching data:", error);
			FetchFileData()
		}
	}

  // 폴더 경로를 기준으로, root path는 지워주는 함수
  const renameFileData = (fileData: main.ResponseFileStruct): { [key: string]: file.File[] } => {
    let renameFileData: { [key: string]: file.File[] } = {}
    for (const path_key in fileData.files) {
      let renameKey = path_key.replace(fileData.root_path, "")

      if (renameKey.startsWith("\\")) {
        renameKey = renameKey.slice(1)
      }
      // @ts-ignore
      renameFileData[renameKey] = fileData.files[path_key]
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
          <div id="folderStructure">
            {JSON.stringify(resFileData) }
          </div>
          <div id="connect_folderStructure">
            연결 된 상대 PC 파일 정보들...
          </div>
        </div>
			</Fragment>}
		</Layout>
	)
}

export default App
