import './App.css'

import { ToastProvider } from './assets/context/useToast.tsx'
import ToastContainer from './assets/components/utils/Toast.tsx'
import { AuthProvider } from './assets/context/useAuth.tsx'

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
