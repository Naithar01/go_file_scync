import { Fragment, useEffect, useState } from "react"

import Loading from "./components/common/Loading"
import Layout from "./components/layouts/Layout"

import { ResponseFileData } from "../wailsjs/go/main/App"
import { main } from "../wailsjs/go/models"

function App() {
	const [isLoading ,setIsLoading] = useState<boolean>(true)
  const [resFileData, setResFileData] = useState<main.ResponseFileStruct>()

  useEffect(() => {
		FetchFileData()
  }, []);

	const FetchFileData = async (): Promise<void> => {
		try {
			const res = await ResponseFileData();

			if (res.root_path.length == 0 || !res.root_path) {
				FetchFileData()
				return
			}
			setIsLoading(false);
      setResFileData(res)
      alert(res.root_path)
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
				Loadding Success
			</Fragment>}
		</Layout>
	)
}

export default App
