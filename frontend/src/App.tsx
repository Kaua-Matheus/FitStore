import './App.css'

import Header from './assets/components/Header'

function App() {

  return (
    <>
      {/* Utilização do header pré definido */}
      <Header/>
    
      <div className="bg-custom-background">
        <main>
          Olá
        </main>
      </div>
    </>
  )
}

export default App
