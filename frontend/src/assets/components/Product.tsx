
// Temos que inserir as props para realmente poder passar parâmetros
interface ProductProps {
    Name: string
    Price?: number
    Description?: string
    Image?: string
}

export default function Product({Name, Price, Description, Image}: ProductProps) {

    return (

        // Adicionar responsividade
        // Adicionada alteração de cores para debug
        <div className="
            h-[20vh] w-[22vw] text-sm
            sm:h-[32vh] sm:w-[22vw]
            md:h-[34vh] md:w-[23vw] md:bg-custom-secondary xl:text-md
            lg:h-[36vh] lg:w-[24vw] lg:bg-custom-secondary xl:text-lg
            xl:h-[38vh] xl:w-82 xl:bg-custom-tertiary xl:text-xl
            bg-custom-primary rounded-md border-md p-4
            flex flex-col">

            <div className="text-center mb-2">
                {`${Name}`}
            </div>

            <div className="flex flex-1 items-center justify-center mb-2">
                
                <img className="
                    h-[20vh] w-[20vh]
                    xl:h-[26vh] xl:w-[26vh]
                    rounded-md" 
                src={Image} alt="Imagem produto" />
            </div> 

            <div className="mb-2">
                R$ {`${Price}`}
            </div>

            {/* Descrição:  {`${Description}`} */}

        </div>
    )
}