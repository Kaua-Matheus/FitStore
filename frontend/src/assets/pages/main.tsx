import '../../App.css'

// React
import { useEffect, useState } from "react"

// Componentes
import Header from '../components/Header'
import Carousel from '../components/Carousel'
import Product from '../components/Product'
import Footer from '../components/Footer'

type ProductData = {
  product: {
    "id": string,
    "product_name": string,
    "product_price": number,
    "product_description": string,
  }
  url_image: string
}

export default function Main() {

  const [products, setProducts] = useState<ProductData[]>([])

  useEffect(() => {

    const fetchAll = async () => {

      try {
        // Response Product
        const response_product = await fetch("http://localhost:8080/product");
        const data_product = await response_product.json();

        setProducts(data_product.data);

      } catch (err) {
        console.log(`Erro: ${err}`);
      }

    };

    fetchAll();

  }, [])

  return (
    <main className="bg-custom-light space-y-32">

      <Header />

      <div className='flex flex-col items-center space-y-6 mx-4'>

        <Carousel />

        <div className='flex w-[90vw] h-[5vh] rounded-lg 
            justify-center items-center text-custom-light-gray font-semibold bg-custom-dark'>
          <p>15% de desconto no próximo pedido "CUPOM"</p>
        </div>

        <div>
          <h1 className='font-fugaz'>Produtos</h1>

          <div className='flex space-x-2'>

            {
              products.map((prod, index) => (
                <Product key={index} Name={prod.product.product_name} Price={prod.product.product_price} Image={prod.url_image}></Product>
              ))
            }
          </div>

        </div>

      </div>

      <Footer />

    </main>
  )
}