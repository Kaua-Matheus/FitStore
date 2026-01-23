// React-Router
import { NavLink } from "react-router"

// React Icons
import { CiSearch, CiMenuBurger, CiShoppingCart } from "react-icons/ci"

// Componentes
import Button from "../elements/Button";
import Logo from "../elements/Logo";
import User from "../elements/User"

export default function Header() {
    return (
        <header className='flex fixed items-center shadow-md justify-between w-full top-0 bg-custom-light-gray py-3 px-10'>

            {/* <div className="flex items-center space-x-2">
                <img className="h-24 rounded" src="src/assets/images/logos/Dark.png" alt="" />
                <p className="font-bold text-2xl text-custom-extra-dark">FitStore</p>
            </div> */}

            <NavLink to="/">
                <Logo />
            </NavLink>

            <div className="flex items-center space-x-12">
                <div className="space-x-2">
                    <NavLink to="/contato">
                        <Button type="secondary">Contato</Button>
                    </NavLink>
                    <Button type="secondary">Sobre Nós</Button>
                    <Button type="secondary">Promoções</Button>
                </div>

                <div className="flex items-center space-x-2">
                    <CiSearch size={40} className="
                        text-custom-dark
                        bg-custom-light-gray
                        active:bg-custom-secondary transition-colors duration-400
                        hover:shadow-md cursor-pointer
                        text-center p-2 rounded"/>
                    <CiShoppingCart size={40} className="
                        text-custom-dark
                        bg-custom-light-gray
                        active:bg-custom-secondary transition-colors duration-400
                        hover:shadow-md cursor-pointer
                        text-center p-2 rounded"/>
                    <User />
                    <CiMenuBurger size={40} className="
                        text-custom-dark
                        bg-custom-light-gray
                        active:bg-custom-secondary transition-colors duration-400
                        hover:shadow-md cursor-pointer
                        text-center p-2 rounded"/>
                </div>
            </div>

        </header>
    )
}