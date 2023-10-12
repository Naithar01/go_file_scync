import { Route, Routes } from "react-router-dom"

import MainPage from "./pages/MainPage"

import Layout from "./components/layouts/Layout"
import InputPortPage from "./pages/InputPortPage"

import "./styles/app_style.css"

function App() {
	return (
  <Routes>
    <Route path="/" element={<Layout />}>
      <Route index element={<InputPortPage />} />
      <Route path="dir" element={<MainPage />} />
    </Route>
  </Routes>
	)
}

export default App
