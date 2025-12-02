import './App.css'

import Header from './assets/components/Header'

function App() {

  return (
    <div className="space-y-2">

      {/* Utilização do header pré definido */}
      <Header/>
    
      <div className="bg-custom-background">
        <main>
          Olá
        </main>
      </div>

    </div>
  )
}

export default App
