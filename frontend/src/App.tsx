import { Route, Routes } from "react-router-dom"

import MainPage from "./pages/MainPage"

import Layout from "./components/layouts/Layout"
import InputPortPage from "./pages/InputPortPage"
import ConnectServerPage from "./pages/ConnectServerPage"

import "./styles/app_style.css"

function App() {
	return (
  <Routes>
    <Route path="/" element={<InputPortPage />}>
    </Route>
    <Route path="/" element={<Layout />}>
      <Route path="/connect"  element={<ConnectServerPage />} />
      <Route path="/dir"  element={<MainPage />} />
    </Route>
  </Routes>
	)
}

export default App
