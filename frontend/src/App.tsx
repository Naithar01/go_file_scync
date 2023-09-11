import { Fragment, useEffect, useState } from "react"

import Loading from "./components/common/Loading"
import Layout from "./components/layouts/Layout"

import { ResponseFileData } from "../wailsjs/go/main/App"
import { file } from "../wailsjs/go/models"

function App() {
	const [isLoading ,setIsLoading] = useState<boolean>(true)

  const [data, setData] = useState<file.File[]>([]);
  const [intervalId, setIntervalId] = useState<number>(0);

  useEffect(() => {
		FetchFileData()
  }, []);

	const FetchFileData = async (): Promise<void> => {
		try {
			const res = await ResponseFileData();

			if (!res) {
				FetchFileData()
				return
			}
			setData(res);
			setIsLoading(false);
			console.log(res);
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
