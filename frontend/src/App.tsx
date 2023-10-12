import { Route, Routes } from "react-router-dom"

import MainPage from "./pages/MainPage"

import "./styles/app_style.css"
import { Link } from "react-router-dom"

function App() {
	return (
    <Routes>
      <Route path="/">
        <Route index element={<>여기에 Port 입력 페이지 <Link to="/dir">Dir</Link></>} />
        <Route path="dir" element={<MainPage />} />
      </Route>
    </Routes>
	)
}

export default App
