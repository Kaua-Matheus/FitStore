import './App.css'

import { useEffect, useState } from "react"

// Importações de componentes
import Header from './assets/components/Header'
import Carousel from './assets/components/Carousel'
import Product from './assets/components/Product'


// Devemos colocar as definições de interface em outro arquivo separado
type ProductData = {
  product: {
    "id": string,
    "product_name": string,
    "product_price": number,
    "product_description": string,
  }
  url_image: string
}

function App() {

  const [products, setProducts] = useState<ProductData[]>([])

  useEffect(() => {

    const fetchAll = async () => {

      try {
        // Response Product
        const response_product = await fetch("http://localhost:8080/product");
        const data_product = await response_product.json();

        setProducts(data_product.data);
        
      } catch(err) {
        console.log(`Erro: ${err}`);
      }

    };

    fetchAll();

  }, [])

  return (
    <main className="bg-custom-background space-y-2"> 

      <Header/>

      <div className='flex flex-col items-center space-y-6 mx-4'>

        <Carousel/>

        <div>
          <p>Ganhe um desconto usando o cupom "SMART"</p>
        </div>

        <div>
          <h1>Produtos</h1>

          <div className='flex space-x-2'>
            {/* Passamos o Product dentro de um () pois componentes react devem ser introduzidos assim */}

            {
              products.map((prod, index) => (
                <Product key={index} Name={prod.product.product_name} Price={prod.product.product_price} Image={prod.url_image}></Product>
              ))
            }
          </div>

        </div>

        <br />
        <div>
          Teste de Imagem
          <img src={`http://localhost:8080/files/Perfil/perfil.png`} alt="Imagem Perfil" />
        </div>

      </div>

    </main>
  )
}

export default App
