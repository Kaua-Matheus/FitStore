import Button from "./Button";

export default function Header() {
    return (
        <header className='flex justify-between w-full top-0 bg-custom-extra-dark py-3 px-10'>

            <p className="text-center">
                FitStore
            </p>

            <div className="space-x-2">
                <Button>Contato</Button>
                <Button>Sobre Nós</Button>
                <Button>Início</Button>
            </div>

        </header>
    )
}