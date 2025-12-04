import './App.css'

import Header from './assets/components/Header'
import Carousel from './assets/components/Carousel'

function App() {

  return (
    <body className="bg-custom-background">

      <div className='items-center space-y-2'>
        <Header/>

        <Carousel/>

        <div>
          <p>Ganhe um desconto usando o cupom "SMART"</p>
        </div>
        
      </div>

    </body>
  )
}

export default App
