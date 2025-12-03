// React Icons
import { CiSearch } from "react-icons/ci"
import { CiMenuBurger } from "react-icons/ci";

import Button from "./Button";

export default function Header() {
    return (
        <header className='flex items-center justify-between w-full top-0 bg-custom-extra-light py-3 px-10'>

            {/* Adicionar logo */}
            {/* Para colocar imagens é necessário incluir o caminho todo desde o root */}
            <div className="flex items-center space-x-2">
                <img className="h-24 rounded" src="src/assets/images/logos/Dark.png" alt="" />
                <p className="font-bold text-2xl text-custom-extra-dark">FitStore</p>
            </div>

            <div className="flex items-center space-x-12">
                <div className="space-x-2">
                    <Button>Contato</Button>
                    <Button>Sobre Nós</Button>
                    <Button>Promoções</Button>
                </div>

                {/* Para alterar o tamanho de um icone usamos o size */}
                {/* Devemos checar se o tamanho é esse mesmo do icone e o resto*/}
                <div className="flex items-center space-x-2">
                    <CiSearch size={40} className="
                        text-custom-extra-dark hover:text-custom-extra-dark
                        bg-custom-extra-light hover:bg-custom-primary
                        active:bg-custom-secondary transition-colors duration-400
                        text-center p-2 rounded"/>
                    <CiMenuBurger size={40} className="
                        text-custom-extra-dark hover:text-custom-extra-dark
                        bg-custom-extra-light hover:bg-custom-primary
                        active:bg-custom-secondary transition-colors duration-400
                        text-center p-2 rounded"/>
                </div>
            </div>

        </header>
    )
}