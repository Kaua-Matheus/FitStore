//import { Default } from './assets/consumers/goApplication';

// Componentes
import Header from './assets/components/Header'
import Footer from './assets/components/Footer'

export default async function Home() {

  // Consulta informação do backend e printa
  //const response = await Default();
  //console.log(response);

  return (
    <body className='mx-32 p-5 space-y-5'>
      <Header/>

      <main>
        <div>
          <div className='flex justify-center'>
            Olá
          </div>
        </div>
      </main>

      <Footer/>
    </body>
  )
}