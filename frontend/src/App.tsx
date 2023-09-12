import { Fragment, useEffect, useState } from "react"

import Loading from "./components/common/Loading"
import Layout from "./components/layouts/Layout"

import { OpenDirectory } from "../wailsjs/go/main/App"
import { main } from "../wailsjs/go/models"

function App() {
	const [isLoading ,setIsLoading] = useState<boolean>(true)
  const [resFileData, setResFileData] = useState<main.ResponseFileStruct>()

  useEffect(() => {
		FetchFileData()
  }, []);

	const FetchFileData = async (): Promise<void> => {
		try {
			const res = await OpenDirectory();

			if (res.root_path.length == 0 || !res.root_path) {
				FetchFileData()
				return
			}
			setIsLoading(false);
      setResFileData(res)
      console.log(res.files);
      
		} catch (error) {
			console.error("Error fetching data:", error);
			FetchFileData()
		}
	}

	return (
		<Layout>
			{ isLoading ? 
			<Loading />
			: 
			<Fragment>
				<div id="folderStructure"></div>
			</Fragment>}
		</Layout>
	)
}

export default App
