import { Fragment, useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";

import { OpenDirectory } from "../../wailsjs/go/main/App";
import { InitialSnycDirectoryListPage } from "../../wailsjs/go/initial/Initial";
import { file, main } from "../../wailsjs/go/models";
import { EventsOn } from "../../wailsjs/runtime/runtime"

import DirectoryList from "../components/directory/DirectoryList";
import Loading from "../components/common/Loading";
import { SendDirectory } from "../../wailsjs/go/tcpserver/TCPServer";

export interface RenameFileData {
  key: string;
  depth: number;
  files: file.File[];
}

const MainPage = () => {
	const navigate = useNavigate()
  
  const [isLoading ,setIsLoading] = useState<boolean>(true)
  const [resFileData, setResFileData] = useState<RenameFileData[]>()
  const [connectedClientFileData, setConnectedClientFileData] = useState<RenameFileData[]>()

  EventsOn("server_shutdown", function() {
    navigate("/connect")
  })

  // const [isOpen, setIsOpen] = useState<boolean>(true)
  // const ToggleModal = (): void => {
  //   setIsOpen((prev) => {
  //     return !prev
  //   })
  //   return 
  // }

  useEffect(() => {
    InitialSnycDirectoryListPage()
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
      setResFileData(() => {
        return renameFile(res)
      })

      // 선택한 폴더의 내용을 상대 PC에게 보내줌
      await SendDirectory(res)
		} catch (error) {
			console.error("Error fetching data:", error);
			FetchFileData()
		}
	}

  // 이 함수는 주어진 파일 데이터를 가공하여 경로를 수정하고, 각 파일의 깊이를 계산한 후 이 정보를 배열로 정리하여 반환합니다.
  // 반환된 데이터는 RenameFileData 형식의 배열입니다.
  const renameFile = (fileData: main.ResponseFileStruct): RenameFileData[] => {
    let renameFileData: RenameFileData[] = []
    for (const path_key in fileData.files) {
      let renameKey = path_key.replace(fileData.root_path, "")
      
      renameKey = renameKey.replace(/\\/g, '/')

      if (renameKey.startsWith("/")) {
        renameKey = renameKey.slice(1)
      }

      renameFileData.push({
        key: renameKey == "" ? path_key : renameKey,
        depth: renameKey == "" ? 0: renameKey.split("/").length,
        // @ts-ignore
        files: fileData.files[path_key],
      })
    }
    return renameFileData
  }

  // 폴더 정보를 상대 PC로 부터 받아오면 
    //  로딩 종료
    //  정보를 변수로 저장
  EventsOn("connectedDirectoryData", async (data: main.ResponseFileStruct) => {
    setConnectedClientFileData(() => {
      return renameFile(data)
    })
    setIsLoading(() => false)
  })

  return (
    <Fragment>
      { isLoading ? 
      <Loading />
      :
      <div className="main">
        {resFileData && <DirectoryList resFileData={resFileData} type='server' />}
        {connectedClientFileData && <DirectoryList resFileData={connectedClientFileData} type='conn' />}
      </div>
      }
    </Fragment>
  )
}

export default MainPage