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

type ImageData = {
  content_type: string
  file_data: string
}

function App() {

  const [products, setProducts] = useState<ProductData[]>([])
  const [images, setImages] = useState<ImageData[]>([])
  // var [data, setData] = useState({})

  useEffect(() => {

    // Teste de API
    // fetch("http://localhost:8080/")
    //   .then(res => res.json())
    //   .then(data => setData(data))
    //   .catch(err => console.error(err));

    // console.log(`Esses são os dados: ${data}`);
    const fetchAll = async () => {

      try {
        // Response Product
        const response_product = await fetch("http://localhost:8080/products");
        const data_product = await response_product.json();

        setProducts(data_product.data);

        const response_image = await fetch("http://localhost:8080/images");
        const data_image = await response_image.json();

        setImages(data_image.data);
        
      } catch(err) {
        console.log(`Erro: ${err}`);
      }

    };

    fetchAll();

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

        <div>
          <h1>Imagens</h1>

          <div className='flex space-x-2'>
            {/* Passamos o Product dentro de um () pois componentes react devem ser introduzidos assim */}

            {
              images.map((img) => (
                <>
                  <p>imagem: {img.file_data}</p> <br />
                  <img src={`${img.file_data}`} alt={`${img.content_type}`} />
                </>
              ))
            }
          </div>

        </div>

        <br />
        <div>
          
          <img src={`http://localhost:8080/files/perfil.png`} alt="Imagem Perfil" />
        </div>

      </div>

    </body>
  )
}

export default App
