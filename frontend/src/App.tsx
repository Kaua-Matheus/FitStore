import './App.css'

import { ToastProvider } from './assets/context/useToast.tsx'
import ToastContainer from './assets/components/Toast.tsx'
import { AuthProvider } from './assets/context/auth.tsx'

// Importação Dom
import { BrowserRouter } from "react-router"
import AppRoutes from './routes/appRoutes'

function App() {

  return (
    <BrowserRouter>

      <AuthProvider>
        <ToastProvider>
          <AppRoutes />
          <ToastContainer />
        </ToastProvider>
      </AuthProvider>
      
    </BrowserRouter>
  )
}

export default App
