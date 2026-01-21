import './App.css'

import { ToastProvider } from './assets/context/useToast.tsx'
import ToastContainer from './assets/components/Toast.tsx'

// Importação Dom
import { BrowserRouter } from "react-router"
import AppRoutes from './routes/appRoutes'

function App() {

  return (
    <BrowserRouter>
      <ToastProvider>
        <AppRoutes />
        <ToastContainer />
      </ToastProvider>
    </BrowserRouter>
  )
}

export default App
