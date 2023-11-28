import { Fragment, useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { useSyncFileDataContext } from "../contexts/SyncFileDataContext";

import { GetRootPath, OpenDirectory, ParseRootPath } from "../../wailsjs/go/main/App";
import { InitialSnycDirectoryListPage } from "../../wailsjs/go/initial/Initial";
import { models } from "../../wailsjs/go/models";
import { EventsOn } from "../../wailsjs/runtime/runtime"
import { SendDirectoryContent } from "../../wailsjs/go/tcpserver/TCPServer";
import { markDuplicates } from "../utils/duplication"
import { markLatests } from "../utils/latest";

import DirectoryList from "../components/directory/DirectoryList";
import Loading from "../components/common/Loading";
import Modal from "../components/common/Modal";

export interface RenameFileData {
  key: string;
  depth: number;
  files: models.File[];
}

const MainPage = () => {
	const navigate = useNavigate()
  const { updateSynchronizedFiles } = useSyncFileDataContext()
  
  const [isLoading ,setIsLoading] = useState<boolean>(true)
  // Sync Files Modal - Files Info
  const [isOpen, setIsOpen] = useState<boolean>(false)
  const ToggleSyncModal = (): void => {
    setIsOpen((prev) => {
      return !prev
    })
    return 
  }

  // Root PC
  const [rootPath, setRootPath] = useState<string>("")
  const [resFile, setResFile] = useState<models.ResponseFileStruct>()
  const [resFileData, setResFileData] = useState<RenameFileData[]>()

  // Connected PC 
  const [connectedClientFile, setConnectedClientFile] = useState<models.ResponseFileStruct>()
  const [connectedClientFileData, setConnectedClientFileData] = useState<RenameFileData[]>()

  // 각 PC의 폴더 동기화 시작을 RUNTIME으로 받으면 
  EventsOn("start_sync_files", async () => {
    setIsLoading(() => true)
    setIsOpen(() => true)
  })

  // 상대 PC 서버 종료 시에 페이지 이동 
  EventsOn("server_shutdown", function() {
    navigate("/connect")
  })

  // 상대 PC 서버로부터 단일 파일 정보를 받아오면
  //  RootPath로 지정 된 폴더에 파일들을 다시 조회 
  EventsOn("receive_file", async () => {
    setIsLoading(() => true)

    const res = await ParseRootPath(await GetRootPath())

    setResFile(() => res)
    await SendDirectoryContent(res)
  })

  // 상대 PC 서버로부터 폴더 정보를 받아오면 
  //  정보를 변수로 저장
  EventsOn("connectedDirectoryData", async (data: models.ResponseFileStruct) => {
    setConnectedClientFile(() => data)
  })

  // 페이지 입장 시에 페이지 셋팅
  useEffect(() => {
    InitialSnycDirectoryListPage()
    FetchFileData()
  }, []);

  // 파일 송/수신 작업이 이루어졌을 때 
  // 페이지 입장 시에 폴더 정보를 가져 왔을 때 
  useEffect(() => {
    if (resFile && connectedClientFile) {
      let res = resFile
      let conRes = connectedClientFile

      markDuplicates(res, conRes)
      markLatests(res, conRes)
      // Check File Name
      // DUPLICATION이 TRUE이면서 LATEST가 1이면 
      // DUPLICATION이 FALSE이면서 LATEST가 0이면
      let synchronized_file_name_list: {filename: string, filepath: string}[] = []
      for (const directory_path in res.files) {
        const directory_path_file = res.files[directory_path]
      
        directory_path_file.forEach(file => {
          if (file.filename && file.duplication && file.latest == 1) {
            synchronized_file_name_list.push({filename: file.filename, filepath: file.directorypath})
          } else if (file.filename && !file.duplication && file.latest == 0) {
            synchronized_file_name_list.push({filename: file.filename, filepath: file.directorypath})
          }
        })
      }
      updateSynchronizedFiles(synchronized_file_name_list)

      setResFileData(() => renameFile(res))
      setConnectedClientFileData(() => renameFile(conRes))

      setIsLoading(() => false)
    }
  }, [resFile, connectedClientFile])

  // 서버로부터 파일 정보와, 선택 된 폴더 경로를 받아옴
  // 마지막에 가공 전 데이터를 보냄 
	const FetchFileData = async (): Promise<void> => {
		try {
			const res = await OpenDirectory();

			if (res.root_path.length == 0 || !res.root_path) {
				FetchFileData()
				return
			}
      setRootPath(() => res.root_path)
      setResFile(() => res)
      await SendDirectoryContent(res)
		} catch (error) {
			console.error("Error fetching data:", error);
			FetchFileData()
		}
	}

  // 이 함수는 주어진 파일 데이터를 가공하여 경로를 수정하고, 각 파일의 깊이를 계산한 후 이 정보를 배열로 정리하여 반환합니다.
  // 반환된 데이터는 RenameFileData 형식의 배열입니다.
  const renameFile = (fileData: models.ResponseFileStruct): RenameFileData[] => {
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
        files: fileData.files[path_key],
      })
    }
    return renameFileData
  }

  return (
    <Fragment>
      { isLoading ? 
      <Loading />
      :
      <div className="main">
        {resFileData && rootPath && <DirectoryList resFileData={resFileData} type='server' root_path={rootPath} />}
        {connectedClientFileData && <DirectoryList resFileData={connectedClientFileData} type='conn' root_path=""/>}
      </div>
      }

      {/* Sync Files Modal */}
      { isOpen && 
        <Modal isOpen={isOpen} onClose={ToggleSyncModal} header_title="파일 동기화 목록">
          
        </Modal> 
      }
    </Fragment>
  )
}

export default MainPage