import './App.css'

import { useEffect, useState } from "react"

// Importações de componentes
import Header from './assets/components/Header'
import Carousel from './assets/components/Carousel'
import Product from './assets/components/Product'


// Devemos colocar as definições de interface em outro arquivo separado
type ProductData = {
  product_name: string
  product_price: number
}

function App() {

  const [products, setProducts] = useState<ProductData[]>([])
  // var [data, setData] = useState({})

  useEffect(() => {

    // Teste de API
    // fetch("http://localhost:8080/")
    //   .then(res => res.json())
    //   .then(data => setData(data))
    //   .catch(err => console.error(err));

    // console.log(`Esses são os dados: ${data}`);
    const fetchProducts = async () => {

      try {
        const response = await fetch("http://localhost:8080/products");

        // Inserimos await pois a operação de conversão de json não é instantânea
        const data = await response.json();

        setProducts(data.data);
        console.log(data);
        
      } catch(err) {
        console.log(`Erro: ${err}`);
      }

    };

    fetchProducts();

  }, [])

  return (
    <body className="bg-custom-background">

      <div className='items-center space-y-2'>
        <Header/>

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
                <Product key={index} Name={prod.product_name} Price={prod.product_price}></Product>
              ))
            }
          </div>

        </div>

      </div>

    </body>
  )
}

export default App
