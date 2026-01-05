import './App.css'

// Importação Dom
import { BrowserRouter } from "react-router"
import AppRoutes from './routes/appRoutes'

function App() {

  return (
    <BrowserRouter>
      <AppRoutes />
    </BrowserRouter>
  )
}

export default App
