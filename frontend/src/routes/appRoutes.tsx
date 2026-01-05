import { Routes, Route } from "react-router"

// Páginas
import Main from "../assets/pages/main"

export default function AppRoutes() {
    return (
        <Routes>
            <Route path="/" element={<Main />}></Route>
        </Routes>
    )
}