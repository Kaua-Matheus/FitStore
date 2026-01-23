// React
import { Routes, Route } from "react-router"

// Páginas
import Main from "../pages/Main"
import Contato from "../pages/Contato"

export default function AppRoutes() {
    return (
        <Routes>
            <Route path="/" element={<Main />}></Route>
            <Route path="/contato" element={<Contato />}></Route>
        </Routes>
    )
}