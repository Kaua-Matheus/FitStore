import Button from "./Button";

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
            h-[50vh] w-[30vw]
            bg-custom-light-gray rounded-sm p-4 shadow-md
            flex flex-col">

            <div className="flex justify-center mb-2">
                
                <img className="
                    h-[24vh] w-[24vh]
                    rounded-sm" 
                src={Image} alt="Imagem produto" />
            </div> 
        
            <div className="flex flex-1 text-center font-light">
                {`${Name}`}
            </div>

            <div className="flex flex-col text-center justify-center">
                <p className="font-light text-custom-secondary">R$ {`${Price?.toFixed(2)}`}</p>
                <p className="font-bold text-custom-primary">R$ {`${Price?.toFixed(2)}`}</p>
            </div>

            <div className="text-center">
                Stars
            </div>

            <div className="flex justify-center">
                <Button type="primary">Comprar</Button>
            </div>

        </div>
    )
}