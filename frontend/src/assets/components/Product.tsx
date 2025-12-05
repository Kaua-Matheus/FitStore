
// Temos que inserir as props para realmente poder passar parâmetros
interface ProductProps {
    Name: string
    Price?: number
    Description?: string
}

export default function Product({Name, Price, Description}: ProductProps) {

    return (

        // Adicionar responsividade
        // Adicionada alteração de cores para debug
        <div className="
            h-[20vh] w-[25vw]
            sm:h-36 sm:w-44
            md:h-40 md:w-48 md:bg-custom-secondary
            lg:h-44 lg:w-52 lg:bg-custom-secondary
            xl:h-48 xl:w-56 xl:bg-custom-tertiary
            bg-custom-primary rounded-md border-md p-3">
            <img src="" alt="Imagem Produto" />

            Nome: {`${Name}`}
            
            {/* Estamos usando um breakline, mas o correto é de outro jeito */}
            <br/>

            Preço: {`${Price}`}

            <br/>

            Descrição:  {`${Description}`}
        </div>
    )
}